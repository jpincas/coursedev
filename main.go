package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/google/uuid"
	gt "github.com/jpincas/go-tea"
	"github.com/yagniltd/coursedev/parsing"
)

// Global shared state - loaded once at startup, read-only during runtime
var (
	globalCourse  *CourseGraph
	globalModules map[string]*Module
	app           *gt.Application
)

func main() {
	// Load and parse all content at startup
	var err error
	globalCourse, globalModules, err = LoadContent("content")
	if err != nil {
		log.Fatalf("Failed to load content: %v", err)
	}

	fmt.Printf("Loaded %d modules\n", len(globalModules))
	for id, mod := range globalModules {
		fmt.Printf("  - %s: %s (%d pages)\n", id, mod.Meta.Title, len(mod.Pages))
	}

	// Initialize the database
	if err := InitDB("coursedev.db"); err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}
	defer CloseDB()

	// Serve content directory for static assets (SVG diagrams, images)
	http.Handle("/content/", http.StripPrefix("/content/", http.FileServer(http.Dir("content"))))

	// Start the Gotea app
	app = gt.NewApp(&Model{})
	fmt.Println("Starting training app at http://localhost:8080")
	app.Start(8080, "static")
}

// LoadContent walks the content directory and parses all modules
func LoadContent(contentDir string) (*CourseGraph, map[string]*Module, error) {
	parsedCourse, parsedModules, err := parsing.LoadCourse(contentDir)
	if err != nil {
		return nil, nil, err
	}

	// Convert parsing types to main types
	course := &CourseGraph{
		ModuleOrder:   parsedCourse.ModuleOrder,
		Prerequisites: parsedCourse.Prerequisites,
	}

	modules := make(map[string]*Module)
	for id, pm := range parsedModules {
		mod := &Module{
			ID: pm.ID,
			Meta: ModuleMeta{
				Title:             pm.Meta.Title,
				Description:       pm.Meta.Description,
				Prerequisites:     pm.Meta.Prerequisites,
				Difficulty:        pm.Meta.Difficulty,
				Roles:             pm.Meta.Roles,
				EstimatedDuration: pm.Meta.EstimatedDuration,
				Completion: CompletionCriteria{
					RequireAllPages: pm.Meta.Completion.RequireAllPages,
					RequireQuizzes:  pm.Meta.Completion.RequireQuizzes,
					MinQuizScore:    pm.Meta.Completion.MinQuizScore,
				},
			},
			Pages: make([]Page, len(pm.Pages)),
		}

		for i, pp := range pm.Pages {
			mod.Pages[i] = Page{
				Filename: pp.Filename,
				Meta: PageMeta{
					Title:    pp.Meta.Title,
					Duration: pp.Meta.Duration,
					Tags:     pp.Meta.Tags,
					Notes:    pp.Meta.Notes,
				},
				NarrativeHTML: pp.NarrativeHTML,
				Blocks:        convertBlocks(pp.Blocks),
				Sections:      convertSections(pp.Sections),
			}
		}

		modules[id] = mod
	}

	return course, modules, nil
}

// convertSections converts parsing.PageSection to main.PageSection
func convertSections(parsingSections []parsing.PageSection) []PageSection {
	if parsingSections == nil {
		return nil
	}
	sections := make([]PageSection, len(parsingSections))
	for i, ps := range parsingSections {
		sections[i] = PageSection{
			Title:         ps.Title,
			NarrativeHTML: ps.NarrativeHTML,
			Blocks:        convertBlocks(ps.Blocks),
		}
	}
	return sections
}

// convertBlocks converts parsing.Block to main.Block
func convertBlocks(parsingBlocks []parsing.Block) []Block {
	blocks := make([]Block, len(parsingBlocks))
	for i, pb := range parsingBlocks {
		switch b := pb.(type) {
		case *parsing.QuizBlock:
			blocks[i] = &QuizBlock{
				ID:          b.ID,
				Type:        b.Type,
				Question:    b.Question,
				Options:     b.Options,
				Answer:      b.Answer,
				Answers:     b.Answers,
				Explanation: b.Explanation,
			}
		case *parsing.CalloutBlock:
			blocks[i] = &CalloutBlock{
				Type:    b.Type,
				Title:   b.Title,
				Content: b.Content,
			}
		case *parsing.AnnotatedImageBlock:
			hotspots := make([]Hotspot, len(b.Hotspots))
			for j, h := range b.Hotspots {
				hotspots[j] = Hotspot{
					X:      h.X,
					Y:      h.Y,
					Label:  h.Label,
					Detail: h.Detail,
				}
			}
			blocks[i] = &AnnotatedImageBlock{
				ID:       b.ID,
				Src:      b.Src,
				Alt:      b.Alt,
				Hotspots: hotspots,
			}
		case *parsing.ExerciseBlock:
			blocks[i] = &ExerciseBlock{
				ID:       b.ID,
				Language: b.Language,
				Prompt:   b.Prompt,
				Starter:  b.Starter,
				Validation: ExerciseValidation{
					Type:     b.Validation.Type,
					Expected: b.Validation.Expected,
					Keywords: b.Validation.Keywords,
				},
			}
		case *parsing.AgentBlock:
			blocks[i] = convertAgentBlock(b)
		}
	}
	return blocks
}

// convertAgentBlock converts a parsing.AgentBlock to a main.AgentBlock
func convertAgentBlock(b *parsing.AgentBlock) *AgentBlock {
	script := make([]ScriptEvent, len(b.Script))
	for j, se := range b.Script {
		args := make(map[string]string)
		for k, v := range se.Args {
			args[k] = v
		}
		script[j] = ScriptEvent{
			Type:            se.Type,
			Text:            se.Text,
			Content:         se.Content,
			Tokens:          se.Tokens,
			Tool:            se.Tool,
			Args:            args,
			Summary:         se.Summary,
			ResetScratchpad: se.ResetScratchpad,
			Note:            se.Note,
		}
	}
	scratchpad := make(map[string]string)
	for k, v := range b.Scratchpad {
		scratchpad[k] = v
	}
	return &AgentBlock{
		ID:         b.ID,
		Title:      b.Title,
		ModelLabel: b.ModelLabel,
		System:     b.System,
		Scratchpad: scratchpad,
		Tools:      b.Tools,
		Visibility: AgentVisibility{
			SystemPrompt: b.Visibility.SystemPrompt,
			ToolCalls:    b.Visibility.ToolCalls,
			FullContext:  b.Visibility.FullContext,
			TokenCount:   b.Visibility.TokenCount,
			ModelName:    b.Visibility.ModelName,
		},
		Sidebar: AgentSidebarConfig{
			Width:     b.Sidebar.Width,
			StartOpen: *b.Sidebar.StartOpen,
		},
		Script: script,
	}
}

// model is a helper for type assertion in handlers
func model(s gt.State) *Model {
	return s.(*Model)
}

// Init creates a new session state. Called once per WebSocket connection.
func (m *Model) Init(sid uuid.UUID) gt.State {
	newModel := &Model{
		SessionID:     sid,
		Progress:      make(map[string]*ModuleProgress),
		ActiveQuizzes: make(map[int]*QuizState),
		Preferences: LearnerPreferences{
			FontSize: "medium",
			Theme:    "light",
		},
		AuthStage: AuthStageEnterCode, // Default to showing code entry
	}

	// Check if there's an existing session → student mapping
	student, err := GetStudentBySessionID(sid)
	if err != nil {
		LogDBError("GetStudentBySessionID", err)
	}

	if student != nil {
		// Existing session - verify cohort is still valid
		cohort, err := GetCohortByID(student.CohortID)
		if err != nil {
			LogDBError("GetCohortByID", err)
		}

		if cohort != nil && !cohort.IsExpired() {
			// Valid session - load student state
			log.Printf("Existing session found for student: %s (cohort: %s)", student.Email, cohort.Name)
			newModel.loadStudentState(student)
		} else {
			// Cohort expired or deleted - clear session and show error
			log.Printf("Cohort expired or deleted for student: %s", student.Email)
			_ = DeleteSessionMapping(sid)
			newModel.AuthStage = AuthStageError
			newModel.AuthError = "Your cohort has expired. Please contact your training manager."
		}
	}

	// Register routes
	newModel.Register("/", newModel.renderCurrentPage)
	newModel.Register("/modules", newModel.renderModuleSelector)
	newModel.Register("/admin", newModel.renderAdmin)

	return newModel
}

// loadStudentState loads a student's state into the model
func (m *Model) loadStudentState(student *Student) {
	m.StudentID = &student.ID
	m.CohortID = &student.CohortID
	m.StudentName = student.Name
	m.StudentEmail = student.Email
	m.CurrentModule = student.CurrentModule
	m.CurrentPage = student.CurrentPage
	m.Preferences.FontSize = student.FontSize
	m.Preferences.Theme = student.Theme
	m.AuthStage = AuthStageNone // Authenticated

	// Load progress from database
	state, err := LoadStudentState(student.ID)
	if err != nil {
		LogDBError("LoadStudentState", err)
		return
	}

	if state != nil {
		// Restore module progress
		for moduleID, progDB := range state.Progress {
			prog := &ModuleProgress{
				Started:     progDB.Started,
				PagesViewed: progDB.PagesViewed,
				QuizScores:  make(map[string]QuizScore),
				CompletedAt: progDB.CompletedAt,
			}
			if prog.PagesViewed == nil {
				prog.PagesViewed = make(map[int]bool)
			}
			m.Progress[moduleID] = prog
		}

		// Restore quiz scores - associate with correct module
		for _, qs := range state.QuizScores {
			for moduleID, mod := range globalModules {
				for _, page := range mod.Pages {
					for _, block := range page.Blocks {
						if quiz, ok := block.(*QuizBlock); ok && quiz.ID == qs.QuizID {
							if m.Progress[moduleID] == nil {
								m.Progress[moduleID] = &ModuleProgress{
									PagesViewed: make(map[int]bool),
									QuizScores:  make(map[string]QuizScore),
								}
							}
							m.Progress[moduleID].QuizScores[qs.QuizID] = QuizScore{
								Correct:  qs.Correct,
								Attempts: qs.Attempts,
							}
							break
						}
					}
				}
			}
		}
	}

	// Validate module still exists
	if m.CurrentModule != "" {
		if _, ok := globalModules[m.CurrentModule]; !ok {
			// Module no longer exists, reset to first
			if len(globalCourse.ModuleOrder) > 0 {
				m.CurrentModule = globalCourse.ModuleOrder[0]
				m.CurrentPage = 0
			}
		}
	}

	// Set starting position if none
	if m.CurrentModule == "" && len(globalCourse.ModuleOrder) > 0 {
		m.CurrentModule = globalCourse.ModuleOrder[0]
		m.CurrentPage = 0
	}

	// Check all modules for completion
	m.checkAllModulesCompletion()

	// Init agent state if current page has an agent block
	m.initAgentIfNeeded()
}

// Update returns all message handlers
func (m *Model) Update() gt.MessageMap {
	return gt.MessageMap{
		// Authentication
		"SUBMIT_COHORT_CODE": m.handleSubmitCohortCode,
		"SUBMIT_LOGIN":       m.handleSubmitLogin,
		"LOGOUT":             m.handleLogout,

		// Navigation
		"NEXT_PAGE":  m.handleNextPage,
		"PREV_PAGE":  m.handlePrevPage,
		"NAV_MODULE": m.handleNavModule,
		"NAV_PAGE":   m.handleNavPage,
		"NAV_SLIDE":  m.handleNavSlide,

		// Quiz interaction
		"QUIZ_ANSWER": m.handleQuizAnswer,
		"QUIZ_RETRY":  m.handleQuizRetry,

		// Interactive blocks
		"HOTSPOT_TOGGLE": m.handleHotspotToggle,

		// Preferences
		"SET_FONT_SIZE": m.handleSetFontSize,
		"SET_THEME":     m.handleSetTheme,

		// Live session management
		"START_PRESENTING": m.handleStartPresenting,
		"STOP_PRESENTING":  m.handleStopPresenting,
		"JOIN_SESSION":     m.handleJoinSession,
		"LEAVE_SESSION":    m.handleLeaveSession,

		// Live polling
		"START_POLL":          m.handleStartPoll,
		"END_POLL":            m.handleEndPoll,
		"CLOSE_POLL":          m.handleClosePoll,
		"TOGGLE_POLL_RESULTS": m.handleTogglePollResults,

		// Annotations
		"ADD_ANNOTATION":    m.handleAddAnnotation,
		"CLEAR_ANNOTATIONS": m.handleClearAnnotations,

		// Admin - Cohort management
		"CREATE_COHORT": m.handleCreateCohort,
		"UPDATE_COHORT": m.handleUpdateCohort,
		"DELETE_COHORT": m.handleDeleteCohort,

		// Agent block
		"AGENT_ADVANCE":           m.handleAgentAdvance,
		"AGENT_RESET":             m.handleAgentReset,
		"AGENT_TOGGLE_SYSTEM":     m.handleAgentToggleSystem,
		"AGENT_TOGGLE_TOOLS":      m.handleAgentToggleTools,
		"AGENT_TOGGLE_CONTEXT":    m.handleAgentToggleContext,
		"AGENT_TOGGLE_SIDEBAR":    m.handleAgentToggleSidebar,
		"AGENT_TOGGLE_SCRATCHPAD": m.handleAgentToggleScratchpad,
		"AGENT_VIEW_FILE":         m.handleAgentViewFile,
		"AGENT_OPEN_WORKSPACE":    m.handleAgentOpenWorkspace,
		"AGENT_CLOSE_WORKSPACE":   m.handleAgentCloseWorkspace,
	}
}

// Render returns the full HTML page
func (m *Model) Render() []byte {
	m.syncFromLiveSession()
	html := renderLayout(m).Bytes()

	// Reset animation indices after rendering so subsequent re-renders
	// (from toggles, etc.) don't replay chat message animations.
	if m.Agent != nil {
		for i := range m.Agent.ChatMessages {
			if m.Agent.ChatMessages[i].GroupIdx >= 0 {
				m.Agent.ChatMessages[i].GroupIdx = -1
			}
		}
	}

	return append([]byte("<!DOCTYPE html>"), html...)
}

// RenderError handles errors during rendering
func (m *Model) RenderError(err error) []byte {
	return renderError(err).Bytes()
}

// OnRouteChange is called when the route changes
func (m *Model) OnRouteChange(path string) {
	// Handle join links: /join/xxx-xxxx-xxx
	if strings.HasPrefix(path, "/join/") {
		code := strings.TrimPrefix(path, "/join/")
		code = strings.TrimSpace(strings.ToLower(code))

		if code != "" && !m.IsAuthenticated() {
			// Look up the cohort
			cohort, err := GetCohortByCode(code)
			if err != nil {
				LogDBError("GetCohortByCode (join link)", err)
			}

			if cohort != nil && !cohort.IsExpired() {
				// Valid cohort - advance to email entry
				m.PendingCohort = cohort
				m.AuthStage = AuthStageEnterEmail
				m.AuthError = ""
			} else if cohort != nil && cohort.IsExpired() {
				m.AuthStage = AuthStageError
				m.AuthError = "This cohort has expired. Please contact your training manager."
			} else {
				m.AuthStage = AuthStageError
				m.AuthError = "That code doesn't match any active cohort."
			}
		}

		// Redirect to home (the auth flow will render)
		m.SetNewRoute("/")
		return
	}

	// Check all modules for completion when navigating
	// This ensures completion is marked even when navigating via routes
	m.checkAllModulesCompletion()
}
