package main

import (
	"log"
	"os"
	"strings"
	"time"

	"github.com/google/uuid"
	gt "github.com/jpincas/go-tea"
)

// ============================================================================
// Authentication handlers
// ============================================================================

func (m *Model) handleSubmitCohortCode(msg gt.Message, s gt.State) gt.Response {
	mdl := model(s)

	var payload struct {
		Code string `json:"code"`
	}
	msg.MustDecodeArgs(&payload)

	log.Printf("SUBMIT_COHORT_CODE received: payload=%+v", payload)

	code := strings.TrimSpace(strings.ToLower(payload.Code))
	if code == "" {
		mdl.AuthStage = AuthStageError
		mdl.AuthError = "Please enter a cohort code."
		return gt.Respond()
	}

	// Check if this is the admin key
	adminKey := os.Getenv("COURSEDEV_ADMIN_KEY")
	log.Printf("SUBMIT_COHORT_CODE: code='%s' adminKey='%s' match=%v", code, adminKey, code == strings.ToLower(adminKey))
	if adminKey != "" && code == strings.ToLower(adminKey) {
		// Admin login - skip email entry
		log.Printf("SUBMIT_COHORT_CODE: Admin login successful!")
		mdl.IsOwner = true
		mdl.AuthStage = AuthStageNone
		mdl.StudentName = "Admin"

		// Set starting position for owner
		if len(globalCourse.ModuleOrder) > 0 {
			mdl.CurrentModule = globalCourse.ModuleOrder[0]
			mdl.CurrentPage = 0
		}

		return gt.Respond()
	}

	// Look up cohort by code
	cohort, err := GetCohortByCode(code)
	if err != nil {
		LogDBError("GetCohortByCode", err)
		mdl.AuthStage = AuthStageError
		mdl.AuthError = "An error occurred. Please try again."
		return gt.Respond()
	}

	if cohort == nil {
		mdl.AuthStage = AuthStageError
		mdl.AuthError = "That code doesn't match any active cohort."
		return gt.Respond()
	}

	if cohort.IsExpired() {
		mdl.AuthStage = AuthStageError
		mdl.AuthError = "This cohort has expired. Please contact your training manager."
		return gt.Respond()
	}

	// Valid cohort - advance to email entry
	mdl.PendingCohort = cohort
	mdl.AuthStage = AuthStageEnterEmail
	mdl.AuthError = ""

	return gt.Respond()
}

func (m *Model) handleSubmitLogin(msg gt.Message, s gt.State) gt.Response {
	mdl := model(s)

	if mdl.PendingCohort == nil {
		mdl.AuthStage = AuthStageEnterCode
		return gt.Respond()
	}

	var payload struct {
		Email string `json:"email"`
		Name  string `json:"name"`
	}
	msg.MustDecodeArgs(&payload)

	email := strings.TrimSpace(strings.ToLower(payload.Email))
	name := strings.TrimSpace(payload.Name)

	if email == "" {
		mdl.AuthStage = AuthStageError
		mdl.AuthError = "Please enter your email address."
		return gt.Respond()
	}

	if name == "" {
		mdl.AuthStage = AuthStageError
		mdl.AuthError = "Please enter your name."
		return gt.Respond()
	}

	// Check if student already exists in this cohort
	student, err := GetStudentByEmail(mdl.PendingCohort.ID, email)
	if err != nil {
		LogDBError("GetStudentByEmail", err)
		mdl.AuthStage = AuthStageError
		mdl.AuthError = "An error occurred. Please try again."
		return gt.Respond()
	}

	if student == nil {
		// New student - check if cohort has capacity
		count, err := GetCohortStudentCount(mdl.PendingCohort.ID)
		if err != nil {
			LogDBError("GetCohortStudentCount", err)
			mdl.AuthStage = AuthStageError
			mdl.AuthError = "An error occurred. Please try again."
			return gt.Respond()
		}

		if count >= mdl.PendingCohort.MaxStudents {
			mdl.AuthStage = AuthStageError
			mdl.AuthError = "This cohort has reached its maximum capacity. Please contact your training manager."
			return gt.Respond()
		}

		// Create new student
		student, err = CreateStudent(mdl.PendingCohort.ID, email, name)
		if err != nil {
			LogDBError("CreateStudent", err)
			mdl.AuthStage = AuthStageError
			mdl.AuthError = "An error occurred. Please try again."
			return gt.Respond()
		}
	}

	// Create session mapping
	if err := CreateSessionMapping(mdl.SessionID, student.ID); err != nil {
		LogDBError("CreateSessionMapping", err)
		mdl.AuthStage = AuthStageError
		mdl.AuthError = "An error occurred. Please try again."
		return gt.Respond()
	}

	// Load student state
	mdl.loadStudentState(student)
	mdl.PendingCohort = nil

	return gt.Respond()
}

func (m *Model) handleLogout(msg gt.Message, s gt.State) gt.Response {
	mdl := model(s)

	// Clear session mapping
	if mdl.StudentID != nil {
		_ = DeleteSessionMapping(mdl.SessionID)
	}

	// Reset auth state
	mdl.StudentID = nil
	mdl.CohortID = nil
	mdl.IsOwner = false
	mdl.StudentName = ""
	mdl.StudentEmail = ""
	mdl.AuthStage = AuthStageEnterCode
	mdl.AuthError = ""
	mdl.PendingCohort = nil
	mdl.Progress = make(map[string]*ModuleProgress)
	mdl.CurrentModule = ""
	mdl.CurrentPage = 0

	// Stop presenting if they were
	if mdl.IsPresenting && mdl.PresentingCohortID != nil {
		DeleteLiveSessionForCohort(*mdl.PresentingCohortID)
		mdl.IsPresenting = false
		mdl.PresentingCohortID = nil
	}

	// Leave any session they were following
	mdl.FollowingLive = false

	return gt.Respond()
}

// ============================================================================
// Navigation handlers
// ============================================================================

func (m *Model) handleNextPage(msg gt.Message, s gt.State) gt.Response {
	mdl := model(s)

	// Must be authenticated
	if !mdl.IsAuthenticated() {
		return gt.Respond()
	}

	mod := mdl.currentModule()
	if mod == nil {
		return gt.Respond()
	}

	// Block advancement if required quizzes are unanswered (students only)
	if !mdl.IsOwner && mdl.currentPageHasUnansweredRequiredQuizzes() {
		return gt.Respond()
	}

	if mdl.CurrentPage < len(mod.Pages)-1 {
		mdl.CurrentPage++
		mdl.ActiveQuizzes = make(map[int]*QuizState)
		mdl.ActiveHotspot = ""
		mdl.markPageViewed()
		mdl.checkAndMarkCompletion()

		// Persist navigation and page view (only for students)
		if mdl.StudentID != nil {
			LogDBError("SaveStudentNavigation", SaveStudentNavigation(*mdl.StudentID, mdl.CurrentModule, mdl.CurrentPage))
			LogDBError("SavePageView", SavePageView(*mdl.StudentID, mdl.CurrentModule, mdl.CurrentPage))
		}

		// If presenting, sync to live session and broadcast
		mdl.SyncToLiveSession()
	}

	return gt.Respond()
}

func (m *Model) handlePrevPage(msg gt.Message, s gt.State) gt.Response {
	mdl := model(s)

	// Must be authenticated
	if !mdl.IsAuthenticated() {
		return gt.Respond()
	}

	if mdl.CurrentPage > 0 {
		mdl.CurrentPage--
		mdl.ActiveQuizzes = make(map[int]*QuizState)
		mdl.ActiveHotspot = ""

		// Persist navigation (only for students)
		if mdl.StudentID != nil {
			LogDBError("SaveStudentNavigation", SaveStudentNavigation(*mdl.StudentID, mdl.CurrentModule, mdl.CurrentPage))
		}

		// If presenting, sync to live session and broadcast
		mdl.SyncToLiveSession()
	}

	return gt.Respond()
}

func (m *Model) handleNavModule(msg gt.Message, s gt.State) gt.Response {
	mdl := model(s)
	moduleID := msg.ArgsToString()

	// Must be authenticated
	if !mdl.IsAuthenticated() {
		return gt.Respond()
	}

	// Check prerequisites (owner bypasses)
	if !mdl.IsOwner && !globalCourse.PrerequisitesMet(moduleID, mdl.Progress) {
		return gt.Respond()
	}

	// Check and persist completion of current module BEFORE switching
	if mdl.CurrentModule != "" && mdl.StudentID != nil {
		oldProgress := mdl.Progress[mdl.CurrentModule]
		wasComplete := oldProgress != nil && oldProgress.CompletedAt != nil
		mdl.checkAndMarkCompletion()
		if oldProgress != nil && !wasComplete && oldProgress.CompletedAt != nil {
			LogDBError("SaveModuleCompletion", SaveModuleCompletion(*mdl.StudentID, mdl.CurrentModule))
		}
	}

	mdl.CurrentModule = moduleID
	mdl.CurrentPage = 0
	mdl.ActiveQuizzes = make(map[int]*QuizState)
	mdl.ActiveHotspot = ""
	mdl.markPageViewed()
	mdl.checkAndMarkCompletion()

	// Persist navigation and page view (only for students)
	if mdl.StudentID != nil {
		LogDBError("SaveStudentNavigation", SaveStudentNavigation(*mdl.StudentID, mdl.CurrentModule, mdl.CurrentPage))
		LogDBError("SavePageView", SavePageView(*mdl.StudentID, mdl.CurrentModule, mdl.CurrentPage))
	}

	// If presenting, sync to live session and broadcast
	mdl.SyncToLiveSession()

	// Navigate to course view
	mdl.SetNewRoute("/")

	return gt.Respond()
}

func (m *Model) handleNavPage(msg gt.Message, s gt.State) gt.Response {
	mdl := model(s)
	pageIdx := msg.ArgsToInt()

	// Must be authenticated
	if !mdl.IsAuthenticated() {
		return gt.Respond()
	}

	mod := mdl.currentModule()
	if mod == nil {
		return gt.Respond()
	}

	if pageIdx >= 0 && pageIdx < len(mod.Pages) {
		mdl.CurrentPage = pageIdx
		mdl.ActiveQuizzes = make(map[int]*QuizState)
		mdl.ActiveHotspot = ""
		mdl.markPageViewed()
		mdl.checkAndMarkCompletion()

		// Persist navigation and page view (only for students)
		if mdl.StudentID != nil {
			LogDBError("SaveStudentNavigation", SaveStudentNavigation(*mdl.StudentID, mdl.CurrentModule, mdl.CurrentPage))
			LogDBError("SavePageView", SavePageView(*mdl.StudentID, mdl.CurrentModule, mdl.CurrentPage))
		}

		// If presenting, sync to live session and broadcast
		mdl.SyncToLiveSession()

		// Navigate to course view
		mdl.SetNewRoute("/")
	}

	return gt.Respond()
}

// ============================================================================
// Quiz handlers
// ============================================================================

func (m *Model) handleQuizAnswer(msg gt.Message, s gt.State) gt.Response {
	mdl := model(s)

	// Must be authenticated
	if !mdl.IsAuthenticated() {
		return gt.Respond()
	}

	page := mdl.currentPage()
	if page == nil {
		return gt.Respond()
	}

	// Decode answer payload
	var payload struct {
		BlockIndex int `json:"blockIndex"`
		Answer     int `json:"answer"`
	}
	msg.MustDecodeArgs(&payload)

	// Get the quiz block
	if payload.BlockIndex < 0 || payload.BlockIndex >= len(page.Blocks) {
		return gt.Respond()
	}

	quiz, ok := page.Blocks[payload.BlockIndex].(*QuizBlock)
	if !ok {
		return gt.Respond()
	}

	// Check if answer is correct
	correct := payload.Answer == quiz.Answer

	if mdl.ActiveQuizzes == nil {
		mdl.ActiveQuizzes = make(map[int]*QuizState)
	}
	mdl.ActiveQuizzes[payload.BlockIndex] = &QuizState{
		QuizID:     quiz.ID,
		BlockIndex: payload.BlockIndex,
		ChosenIdx:  payload.Answer,
		Answered:   true,
		Correct:    correct,
	}

	// Record score
	progress := mdl.ensureProgress(mdl.CurrentModule)
	existing := progress.QuizScores[quiz.ID]
	newScore := QuizScore{
		Correct:  correct,
		Attempts: existing.Attempts + 1,
	}
	progress.QuizScores[quiz.ID] = newScore

	// Persist quiz score (only for students)
	if mdl.StudentID != nil {
		LogDBError("SaveQuizScore", SaveQuizScore(*mdl.StudentID, quiz.ID, correct, newScore.Attempts))

		// Check if module is now complete
		wasComplete := progress.CompletedAt != nil
		mdl.checkAndMarkCompletion()
		if !wasComplete && progress.CompletedAt != nil {
			LogDBError("SaveModuleCompletion", SaveModuleCompletion(*mdl.StudentID, mdl.CurrentModule))
		}
	}

	// If there's an active poll in the live session for this quiz, record the response
	if session := mdl.getFollowingSession(); session != nil {
		if session.RecordPollResponse(mdl.SessionID, payload.Answer) {
			app.Broadcast() // Update poll results for everyone
		}
	}

	return gt.Respond()
}

func (m *Model) handleQuizRetry(msg gt.Message, s gt.State) gt.Response {
	mdl := model(s)
	blockIndex := msg.ArgsToInt()
	if mdl.ActiveQuizzes != nil {
		delete(mdl.ActiveQuizzes, blockIndex)
	}
	return gt.Respond()
}

// ============================================================================
// Interactive block handlers
// ============================================================================

func (m *Model) handleHotspotToggle(msg gt.Message, s gt.State) gt.Response {
	mdl := model(s)
	hotspotID := msg.ArgsToString()

	if mdl.ActiveHotspot == hotspotID {
		mdl.ActiveHotspot = ""
	} else {
		mdl.ActiveHotspot = hotspotID
	}

	return gt.Respond()
}

// ============================================================================
// Preference handlers
// ============================================================================

func (m *Model) handleSetFontSize(msg gt.Message, s gt.State) gt.Response {
	mdl := model(s)
	size := msg.ArgsToString()

	switch size {
	case "small", "medium", "large":
		mdl.Preferences.FontSize = size
		if mdl.StudentID != nil {
			LogDBError("SaveStudentPreferences", SaveStudentPreferences(*mdl.StudentID, mdl.Preferences.FontSize, mdl.Preferences.Theme))
		}
	}

	return gt.Respond()
}

func (m *Model) handleSetTheme(msg gt.Message, s gt.State) gt.Response {
	mdl := model(s)
	theme := msg.ArgsToString()

	switch theme {
	case "light", "dark":
		mdl.Preferences.Theme = theme
		if mdl.StudentID != nil {
			LogDBError("SaveStudentPreferences", SaveStudentPreferences(*mdl.StudentID, mdl.Preferences.FontSize, mdl.Preferences.Theme))
		}
	}

	return gt.Respond()
}

// ============================================================================
// Live session handlers
// ============================================================================

func (m *Model) handleStartPresenting(msg gt.Message, s gt.State) gt.Response {
	mdl := model(s)

	// Must be owner to present
	if !mdl.IsOwner {
		return gt.Respond()
	}

	// Already presenting?
	if mdl.IsPresenting {
		return gt.Respond()
	}

	// Decode the cohort ID from the message
	cohortIDStr := msg.ArgsToString()
	cohortID, err := uuid.Parse(cohortIDStr)
	if err != nil {
		return gt.Respond()
	}

	// Create a new live session for this cohort
	_, err = CreateLiveSessionForCohort(cohortID, mdl.SessionID, mdl.CurrentModule, mdl.CurrentPage)
	if err != nil {
		return gt.Respond() // Cohort already has a session
	}

	mdl.IsPresenting = true
	mdl.PresentingCohortID = &cohortID

	// Broadcast so everyone knows a new session is available
	app.Broadcast()

	return gt.Respond()
}

func (m *Model) handleStopPresenting(msg gt.Message, s gt.State) gt.Response {
	mdl := model(s)

	if !mdl.IsPresenting || mdl.PresentingCohortID == nil {
		return gt.Respond()
	}

	// Delete the live session
	DeleteLiveSessionForCohort(*mdl.PresentingCohortID)
	mdl.IsPresenting = false
	mdl.PresentingCohortID = nil

	// Broadcast so everyone knows the session ended
	app.Broadcast()

	return gt.Respond()
}

func (m *Model) handleJoinSession(msg gt.Message, s gt.State) gt.Response {
	mdl := model(s)

	// Can't join if already presenting
	if mdl.IsPresenting {
		return gt.Respond()
	}

	// Student must have a cohort
	if mdl.CohortID == nil {
		return gt.Respond()
	}

	// Verify the cohort has an active session
	if GetLiveSessionForCohort(*mdl.CohortID) == nil {
		return gt.Respond()
	}

	mdl.FollowingLive = true
	mdl.ActiveQuizzes = make(map[int]*QuizState)
	mdl.ActiveHotspot = ""

	return gt.Respond()
}

func (m *Model) handleLeaveSession(msg gt.Message, s gt.State) gt.Response {
	mdl := model(s)
	mdl.FollowingLive = false
	return gt.Respond()
}

// ============================================================================
// Live poll handlers
// ============================================================================

func (m *Model) handleStartPoll(msg gt.Message, s gt.State) gt.Response {
	mdl := model(s)

	// Must be presenting to start a poll
	session := mdl.getPresentingSession()
	if session == nil {
		return gt.Respond()
	}

	var payload struct {
		BlockIndex int `json:"blockIndex"`
	}
	msg.MustDecodeArgs(&payload)

	page := mdl.currentPage()
	if page == nil {
		return gt.Respond()
	}

	if payload.BlockIndex < 0 || payload.BlockIndex >= len(page.Blocks) {
		return gt.Respond()
	}

	quiz, ok := page.Blocks[payload.BlockIndex].(*QuizBlock)
	if !ok {
		return gt.Respond()
	}

	session.StartPoll(quiz.ID, payload.BlockIndex)
	app.Broadcast()

	return gt.Respond()
}

func (m *Model) handleEndPoll(msg gt.Message, s gt.State) gt.Response {
	mdl := model(s)

	session := mdl.getPresentingSession()
	if session == nil {
		return gt.Respond()
	}

	session.ClosePoll()
	app.Broadcast()

	return gt.Respond()
}

func (m *Model) handleClosePoll(msg gt.Message, s gt.State) gt.Response {
	mdl := model(s)

	session := mdl.getPresentingSession()
	if session == nil {
		return gt.Respond()
	}

	session.EndPoll()
	app.Broadcast()

	return gt.Respond()
}

func (m *Model) handleTogglePollResults(msg gt.Message, s gt.State) gt.Response {
	mdl := model(s)

	session := mdl.getPresentingSession()
	if session == nil {
		return gt.Respond()
	}

	session.TogglePollResults()
	app.Broadcast()

	return gt.Respond()
}

// ============================================================================
// Annotation handlers
// ============================================================================

func (m *Model) handleAddAnnotation(msg gt.Message, s gt.State) gt.Response {
	mdl := model(s)

	session := mdl.getPresentingSession()
	if session == nil {
		return gt.Respond()
	}

	var stroke AnnotationStroke
	msg.MustDecodeArgs(&stroke)

	if len(stroke.Points) == 0 {
		return gt.Respond()
	}

	session.AddAnnotation(mdl.CurrentModule, mdl.CurrentPage, stroke)
	app.Broadcast()

	return gt.Respond()
}

func (m *Model) handleClearAnnotations(msg gt.Message, s gt.State) gt.Response {
	mdl := model(s)

	session := mdl.getPresentingSession()
	if session == nil {
		return gt.Respond()
	}

	session.ClearAnnotations(mdl.CurrentModule, mdl.CurrentPage)
	app.Broadcast()

	return gt.Respond()
}

// ============================================================================
// Admin handlers - Cohort management
// ============================================================================

func (m *Model) handleCreateCohort(msg gt.Message, s gt.State) gt.Response {
	mdl := model(s)

	// Must be owner to create cohorts
	if !mdl.IsOwner {
		return gt.Respond()
	}

	var payload struct {
		Name        string `json:"name"`
		ExpiresAt   string `json:"expiresAt"` // ISO date string
		MaxStudents int    `json:"maxStudents"`
	}
	msg.MustDecodeArgs(&payload)

	name := strings.TrimSpace(payload.Name)
	if name == "" {
		return gt.Respond()
	}

	expiresAt, err := time.Parse("2006-01-02", payload.ExpiresAt)
	if err != nil {
		// Default to 30 days from now
		expiresAt = time.Now().AddDate(0, 0, 30)
	}
	// Set to end of day
	expiresAt = expiresAt.Add(23*time.Hour + 59*time.Minute + 59*time.Second)

	maxStudents := payload.MaxStudents
	if maxStudents <= 0 {
		maxStudents = 25 // Default
	}

	_, err = CreateCohort(name, expiresAt, maxStudents)
	if err != nil {
		LogDBError("CreateCohort", err)
	}

	return gt.Respond()
}

func (m *Model) handleUpdateCohort(msg gt.Message, s gt.State) gt.Response {
	mdl := model(s)

	// Must be owner to update cohorts
	if !mdl.IsOwner {
		return gt.Respond()
	}

	var payload struct {
		ID          string `json:"id"`
		Name        string `json:"name"`
		ExpiresAt   string `json:"expiresAt"`
		MaxStudents int    `json:"maxStudents"`
	}
	msg.MustDecodeArgs(&payload)

	cohortID, err := uuid.Parse(payload.ID)
	if err != nil {
		return gt.Respond()
	}

	name := strings.TrimSpace(payload.Name)
	if name == "" {
		return gt.Respond()
	}

	expiresAt, err := time.Parse("2006-01-02", payload.ExpiresAt)
	if err != nil {
		return gt.Respond()
	}
	expiresAt = expiresAt.Add(23*time.Hour + 59*time.Minute + 59*time.Second)

	maxStudents := payload.MaxStudents
	if maxStudents <= 0 {
		maxStudents = 25
	}

	err = UpdateCohort(cohortID, name, expiresAt, maxStudents)
	if err != nil {
		LogDBError("UpdateCohort", err)
	}

	return gt.Respond()
}

func (m *Model) handleDeleteCohort(msg gt.Message, s gt.State) gt.Response {
	mdl := model(s)

	// Must be owner to delete cohorts
	if !mdl.IsOwner {
		return gt.Respond()
	}

	cohortIDStr := msg.ArgsToString()
	cohortID, err := uuid.Parse(cohortIDStr)
	if err != nil {
		return gt.Respond()
	}

	err = DeleteCohort(cohortID)
	if err != nil {
		LogDBError("DeleteCohort", err)
	}

	return gt.Respond()
}
