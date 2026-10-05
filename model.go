package main

import (
	"log"
	"time"

	"github.com/google/uuid"
	gt "github.com/jpincas/go-tea"
)

// Model is the application state. It MUST embed gt.Router for routing to work.
type Model struct {
	gt.Router

	// --- Session identity ---
	SessionID uuid.UUID  // Gotea session ID (always set)
	StudentID *uuid.UUID // Set after login (nil for owner or unauthenticated)
	CohortID  *uuid.UUID // Set after login (nil for owner)
	IsOwner   bool       // True if logged in via admin key

	// --- Authentication flow state ---
	AuthStage     AuthStage // Current step in auth flow
	PendingCohort *Cohort   // Validated cohort during auth
	AuthError     string    // Error message to display

	// --- Per-session learner state (loaded after auth) ---
	StudentName   string // Display name
	StudentEmail  string // Email address
	CurrentModule string
	CurrentPage   int
	CurrentSlide  int // Which section within current page (0 = first)
	Progress      map[string]*ModuleProgress
	ActiveQuizzes map[int]*QuizState // blockIndex -> state (multiple quizzes per page)
	ActiveHotspot string             // ID of currently expanded hotspot
	Preferences   LearnerPreferences

	// --- Live session mode ---
	IsPresenting       bool       // This session is driving a live session
	PresentingCohortID *uuid.UUID // Which cohort the admin is presenting to
	FollowingLive      bool       // Position driven by cohort's live session

	// --- Agent block state ---
	Agent *AgentState // nil when current page has no agent block
}

// AuthStage represents the current step in the authentication flow
type AuthStage string

const (
	AuthStageNone       AuthStage = ""            // Check session or show code entry
	AuthStageEnterCode  AuthStage = "enter_code"  // Waiting for cohort code
	AuthStageEnterEmail AuthStage = "enter_email" // Waiting for email/name
	AuthStageError      AuthStage = "error"       // Show error, allow retry
)

// IsAuthenticated returns true if the user is logged in (as student or owner)
func (m *Model) IsAuthenticated() bool {
	return m.StudentID != nil || m.IsOwner
}

// Module represents a training module with metadata and pages
type Module struct {
	ID    string
	Meta  ModuleMeta
	Pages []Page
}

// ModuleMeta contains module-level metadata from _module.yaml
type ModuleMeta struct {
	Title             string             `yaml:"title"`
	Description       string             `yaml:"description"`
	Prerequisites     []string           `yaml:"prerequisites"`
	Difficulty        string             `yaml:"difficulty"`
	Roles             []string           `yaml:"roles"`
	EstimatedDuration string             `yaml:"estimated_duration"`
	Completion        CompletionCriteria `yaml:"completion"`
}

// CompletionCriteria defines when a module is considered complete
type CompletionCriteria struct {
	RequireAllPages bool    `yaml:"require_all_pages"`
	RequireQuizzes  bool    `yaml:"require_quizzes"`
	MinQuizScore    float64 `yaml:"min_quiz_score"`
}

// PageSection represents a chunk of a page split at H2 boundaries
type PageSection struct {
	Title         string  // H2 text (or frontmatter title for first section)
	NarrativeHTML []byte  // This section's HTML with 0-based block markers
	Blocks        []Block // Subset of blocks in this section
}

// Page represents a single page within a module
type Page struct {
	Filename      string
	Meta          PageMeta
	NarrativeHTML []byte       // Goldmark-rendered HTML with block placeholders
	Blocks        []Block      // Extracted interactive blocks
	Sections      []PageSection // nil if page has no H2s (renders as single chunk)
}

// PageMeta contains page-level metadata from frontmatter
type PageMeta struct {
	Title    string   `yaml:"title"`
	Duration string   `yaml:"duration"`
	Tags     []string `yaml:"tags"`
	Notes    string   `yaml:"notes"` // Instructor notes
}

// ModuleProgress tracks a learner's progress through a module
type ModuleProgress struct {
	Started      bool
	PagesViewed  map[int]bool         // Which pages have been viewed
	SlidesViewed map[int]map[int]bool // pageIdx -> slideIdx -> viewed
	QuizScores   map[string]QuizScore // quiz ID -> score
	CompletedAt  *time.Time
}

// QuizScore records a learner's performance on a quiz
type QuizScore struct {
	Correct  bool
	Attempts int
}

// QuizState tracks the current state of quiz interaction
type QuizState struct {
	QuizID      string
	BlockIndex  int
	ChosenIdx   int   // For single answer
	ChosenIdxes []int // For multi-select
	Answered    bool
	Correct     bool
}

// LearnerPreferences stores user preferences
type LearnerPreferences struct {
	FontSize string // "small", "medium", "large"
	Theme    string // "light", "dark"
}

// ============================================================================
// Live Session Helpers
// ============================================================================

// SyncToLiveSession updates the live session navigation if this user is presenting.
// Call this after any navigation change to broadcast to followers.
func (m *Model) SyncToLiveSession() {
	if !m.IsPresenting || m.PresentingCohortID == nil {
		return
	}
	session := GetLiveSessionForCohort(*m.PresentingCohortID)
	if session == nil {
		return
	}
	session.UpdateNavigation(m.CurrentModule, m.CurrentPage, m.CurrentSlide)
	app.Broadcast()
}

// OnDisconnect is called by Gotea when the WebSocket session disconnects.
// It saves admin position and cleans up live sessions if this user was presenting.
func (m *Model) OnDisconnect() {
	// Save admin position so they can reconnect seamlessly
	if m.IsOwner {
		SaveAdminPosition(m.SessionID, m.CurrentModule, m.CurrentPage, m.CurrentSlide)
	}

	// Clean up live session if this user was presenting
	if m.IsPresenting && m.PresentingCohortID != nil {
		// Only delete if the session's presenter is still us (guards against race with takeover)
		session := GetLiveSessionForCohort(*m.PresentingCohortID)
		if session != nil && session.PresenterSID == m.SessionID {
			log.Printf("Presenter disconnected — cleaning up live session for cohort %s", m.PresentingCohortID)
			DeleteLiveSessionForCohort(*m.PresentingCohortID)
			m.IsPresenting = false
			m.PresentingCohortID = nil
			// Broadcast in goroutine so followers learn the session ended
			go app.Broadcast()
		}
	}
}

// syncFromLiveSession syncs this student's position from the live session.
// Called at the top of Render() so followers advance with the instructor.
func (m *Model) syncFromLiveSession() {
	if !m.FollowingLive || m.CohortID == nil {
		return
	}
	session := GetLiveSessionForCohort(*m.CohortID)
	if session == nil {
		// Session ended — stop following
		m.FollowingLive = false
		return
	}
	module, page, slide := session.GetNavigation()
	if module == m.CurrentModule && page == m.CurrentPage && slide == m.CurrentSlide {
		return // No change
	}
	// Advance student's actual position
	m.CurrentModule = module
	m.CurrentPage = page
	m.CurrentSlide = slide
	m.ActiveQuizzes = make(map[int]*QuizState)
	m.ActiveHotspot = ""
	m.Agent = nil
	m.markSlideViewed()
	m.checkAndMarkCompletion()
	// Persist asynchronously (writes are idempotent upserts, safe from goroutine)
	if m.StudentID != nil {
		sid := *m.StudentID
		go func() {
			LogDBError("SaveStudentNavigation", SaveStudentNavigation(sid, module, page))
			LogDBError("SavePageView", SavePageView(sid, module, page))
		}()
	}
}

// ============================================================================
// Navigation Helpers
// ============================================================================

// currentModule returns the current module or nil
func (m *Model) currentModule() *Module {
	if m.CurrentModule == "" {
		return nil
	}
	return globalModules[m.CurrentModule]
}

// currentPage returns the current page or nil
func (m *Model) currentPage() *Page {
	mod := m.currentModule()
	if mod == nil || m.CurrentPage < 0 || m.CurrentPage >= len(mod.Pages) {
		return nil
	}
	return &mod.Pages[m.CurrentPage]
}

// currentPageSectionCount returns the number of sections in the current page, or 0 if none
func (m *Model) currentPageSectionCount() int {
	page := m.currentPage()
	if page == nil || page.Sections == nil {
		return 0
	}
	return len(page.Sections)
}

// ============================================================================
// Progress Helpers
// ============================================================================

// ensureProgress ensures a ModuleProgress exists for the given module
func (m *Model) ensureProgress(moduleID string) *ModuleProgress {
	if m.Progress[moduleID] == nil {
		m.Progress[moduleID] = &ModuleProgress{
			PagesViewed:  make(map[int]bool),
			SlidesViewed: make(map[int]map[int]bool),
			QuizScores:   make(map[string]QuizScore),
		}
	}
	if m.Progress[moduleID].SlidesViewed == nil {
		m.Progress[moduleID].SlidesViewed = make(map[int]map[int]bool)
	}
	return m.Progress[moduleID]
}

// markPageViewed marks the current page as viewed
func (m *Model) markPageViewed() {
	if m.CurrentModule == "" {
		return
	}
	progress := m.ensureProgress(m.CurrentModule)
	progress.Started = true
	progress.PagesViewed[m.CurrentPage] = true
}

// markSlideViewed marks the current slide as viewed.
// For pages without sections, it immediately marks the page as viewed.
// For sectioned pages, it marks the page as viewed only after all slides are seen.
func (m *Model) markSlideViewed() {
	if m.CurrentModule == "" {
		return
	}
	page := m.currentPage()
	if page == nil {
		return
	}
	progress := m.ensureProgress(m.CurrentModule)
	progress.Started = true

	// Pages without sections: behave like markPageViewed
	if page.Sections == nil {
		progress.PagesViewed[m.CurrentPage] = true
		return
	}

	// Track this slide
	if progress.SlidesViewed[m.CurrentPage] == nil {
		progress.SlidesViewed[m.CurrentPage] = make(map[int]bool)
	}
	progress.SlidesViewed[m.CurrentPage][m.CurrentSlide] = true

	// Check if all slides of this page are now viewed
	allViewed := true
	for i := range page.Sections {
		if !progress.SlidesViewed[m.CurrentPage][i] {
			allViewed = false
			break
		}
	}
	if allViewed {
		progress.PagesViewed[m.CurrentPage] = true
	}
}

// isModuleComplete checks if a module meets its completion criteria
func (m *Model) isModuleComplete(moduleID string) bool {
	mod, ok := globalModules[moduleID]
	if !ok {
		return false
	}
	progress := m.Progress[moduleID]
	if progress == nil {
		return false
	}

	// Check all pages viewed if required
	if mod.Meta.Completion.RequireAllPages {
		for i := range mod.Pages {
			if !progress.PagesViewed[i] {
				return false
			}
		}
	}

	// Check quizzes if required
	if mod.Meta.Completion.RequireQuizzes {
		totalQuizzes := 0
		correctQuizzes := 0
		for _, page := range mod.Pages {
			for _, block := range page.Blocks {
				if _, ok := block.(*QuizBlock); ok {
					totalQuizzes++
				}
			}
		}
		for _, score := range progress.QuizScores {
			if score.Correct {
				correctQuizzes++
			}
		}

		if totalQuizzes > 0 {
			scoreRatio := float64(correctQuizzes) / float64(totalQuizzes)
			if scoreRatio < mod.Meta.Completion.MinQuizScore {
				return false
			}
		}
	}

	return true
}

// isPageComplete checks if a specific page is complete (viewed + all quizzes passed)
func (m *Model) isPageComplete(moduleID string, pageIdx int) bool {
	mod := globalModules[moduleID]
	if mod == nil || pageIdx < 0 || pageIdx >= len(mod.Pages) {
		return false
	}
	progress := m.Progress[moduleID]
	if progress == nil || !progress.PagesViewed[pageIdx] {
		return false
	}

	page := mod.Pages[pageIdx]
	for _, block := range page.Blocks {
		if quiz, ok := block.(*QuizBlock); ok {
			score, exists := progress.QuizScores[quiz.ID]
			if !exists || !score.Correct {
				return false
			}
		}
	}
	return true
}

// pageHasIncompleteQuiz checks if a viewed page has quizzes that aren't passed
func (m *Model) pageHasIncompleteQuiz(moduleID string, pageIdx int) bool {
	mod := globalModules[moduleID]
	if mod == nil || pageIdx < 0 || pageIdx >= len(mod.Pages) {
		return false
	}
	progress := m.Progress[moduleID]
	if progress == nil || !progress.PagesViewed[pageIdx] {
		return false
	}

	page := mod.Pages[pageIdx]
	for _, block := range page.Blocks {
		if quiz, ok := block.(*QuizBlock); ok {
			score, exists := progress.QuizScores[quiz.ID]
			if !exists || !score.Correct {
				return true
			}
		}
	}
	return false
}

// moduleNeedsQuizzes checks if all pages are viewed but quizzes are blocking completion
func (m *Model) moduleNeedsQuizzes(moduleID string) bool {
	mod := globalModules[moduleID]
	if mod == nil {
		return false
	}
	progress := m.Progress[moduleID]
	if progress == nil || progress.CompletedAt != nil {
		return false
	}

	for i := range mod.Pages {
		if !progress.PagesViewed[i] {
			return false
		}
	}

	if !mod.Meta.Completion.RequireQuizzes {
		return false
	}

	totalQuizzes := 0
	correctQuizzes := 0
	for _, page := range mod.Pages {
		for _, block := range page.Blocks {
			if _, ok := block.(*QuizBlock); ok {
				totalQuizzes++
			}
		}
	}
	for _, score := range progress.QuizScores {
		if score.Correct {
			correctQuizzes++
		}
	}

	if totalQuizzes == 0 {
		return false
	}
	scoreRatio := float64(correctQuizzes) / float64(totalQuizzes)
	return scoreRatio < mod.Meta.Completion.MinQuizScore
}

// currentPageHasUnansweredRequiredQuizzes checks if the current page has quizzes
// that must be answered correctly before advancing (when module has RequireQuizzes)
func (m *Model) currentPageHasUnansweredRequiredQuizzes() bool {
	mod := m.currentModule()
	if mod == nil || !mod.Meta.Completion.RequireQuizzes {
		return false
	}
	page := m.currentPage()
	if page == nil {
		return false
	}
	progress := m.Progress[m.CurrentModule]
	for _, block := range page.Blocks {
		if quiz, ok := block.(*QuizBlock); ok {
			if progress == nil {
				return true
			}
			score, exists := progress.QuizScores[quiz.ID]
			if !exists || !score.Correct {
				return true
			}
		}
	}
	return false
}

// checkAndMarkCompletion checks if current module is complete and marks it
func (m *Model) checkAndMarkCompletion() {
	if m.CurrentModule == "" {
		return
	}
	progress := m.Progress[m.CurrentModule]
	if progress == nil || progress.CompletedAt != nil {
		return // No progress or already completed
	}
	if m.isModuleComplete(m.CurrentModule) {
		now := time.Now()
		progress.CompletedAt = &now
	}
}

// checkAllModulesCompletion checks all modules for completion (useful after loading state)
func (m *Model) checkAllModulesCompletion() {
	for moduleID, progress := range m.Progress {
		if progress == nil || progress.CompletedAt != nil {
			continue
		}
		if m.isModuleComplete(moduleID) {
			now := time.Now()
			progress.CompletedAt = &now
			if m.StudentID != nil {
				LogDBError("SaveModuleCompletion", SaveModuleCompletion(*m.StudentID, moduleID))
			}
		}
	}
}

// getPresentingSession returns the live session this user is presenting, or nil
func (m *Model) getPresentingSession() *LiveSession {
	if !m.IsPresenting || m.PresentingCohortID == nil {
		return nil
	}
	return GetLiveSessionForCohort(*m.PresentingCohortID)
}

// getFollowingSession returns the live session this user is following, or nil
func (m *Model) getFollowingSession() *LiveSession {
	if !m.FollowingLive || m.CohortID == nil {
		return nil
	}
	return GetLiveSessionForCohort(*m.CohortID)
}
