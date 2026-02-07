package main

import (
	"fmt"
	"sync"

	"github.com/google/uuid"
)

// LiveSession represents an active presentation for a cohort.
// The presenter controls the navigation state; all followers in the cohort see what the presenter sees.
type LiveSession struct {
	sync.RWMutex
	CohortID      uuid.UUID
	PresenterSID  uuid.UUID // Session ID of presenter
	CurrentModule string
	CurrentPage   int
	CurrentSlide  int

	// Poll state for this session
	Poll *LivePollState

	// Annotations keyed by "module:page"
	Annotations map[string][]AnnotationStroke

	// Agent state for live presentation of agent blocks
	Agent *AgentState
}

// AnnotationStroke represents a single drawn stroke on the canvas
type AnnotationStroke struct {
	Tool   string            `json:"tool"`   // "pen" or "highlighter"
	Points []AnnotationPoint `json:"points"`
}

// AnnotationPoint is a single coordinate in a stroke (0-1 percentage-based)
type AnnotationPoint struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
}

// LivePollState tracks an active poll within a live session
type LivePollState struct {
	Active      bool
	QuizID      string
	BlockIndex  int
	Responses   map[int]int        // answer index -> count
	Responded   map[uuid.UUID]bool // who has responded
	ShowResults bool
	Closed      bool
}

// Registry of active live sessions, keyed by cohort ID (one per cohort)
var liveSessionRegistry = struct {
	sync.RWMutex
	byCohort map[uuid.UUID]*LiveSession
}{byCohort: make(map[uuid.UUID]*LiveSession)}

// GetLiveSessionForCohort returns the live session for a cohort, or nil if none
func GetLiveSessionForCohort(cohortID uuid.UUID) *LiveSession {
	liveSessionRegistry.RLock()
	defer liveSessionRegistry.RUnlock()
	return liveSessionRegistry.byCohort[cohortID]
}

// CreateLiveSessionForCohort creates a new live session for a cohort.
// Returns an error if the cohort already has an active session.
func CreateLiveSessionForCohort(cohortID uuid.UUID, presenterSID uuid.UUID, module string, page int, slide int) (*LiveSession, error) {
	liveSessionRegistry.Lock()
	defer liveSessionRegistry.Unlock()

	if _, exists := liveSessionRegistry.byCohort[cohortID]; exists {
		return nil, fmt.Errorf("cohort %s already has an active live session", cohortID)
	}

	session := &LiveSession{
		CohortID:      cohortID,
		PresenterSID:  presenterSID,
		CurrentModule: module,
		CurrentPage:   page,
		CurrentSlide:  slide,
		Annotations:   make(map[string][]AnnotationStroke),
	}
	liveSessionRegistry.byCohort[cohortID] = session
	return session, nil
}

// DeleteLiveSessionForCohort removes the live session for a cohort
func DeleteLiveSessionForCohort(cohortID uuid.UUID) {
	liveSessionRegistry.Lock()
	delete(liveSessionRegistry.byCohort, cohortID)
	liveSessionRegistry.Unlock()
}

// UpdateNavigation updates the session's navigation state
func (s *LiveSession) UpdateNavigation(module string, page int, slide int) {
	s.Lock()
	s.CurrentModule = module
	s.CurrentPage = page
	s.CurrentSlide = slide
	s.Unlock()
}

// GetNavigation returns the current module, page, and slide
func (s *LiveSession) GetNavigation() (string, int, int) {
	s.RLock()
	defer s.RUnlock()
	return s.CurrentModule, s.CurrentPage, s.CurrentSlide
}

// StartPoll starts a new poll for a quiz
func (s *LiveSession) StartPoll(quizID string, blockIndex int) {
	s.Lock()
	defer s.Unlock()
	s.Poll = &LivePollState{
		Active:      true,
		QuizID:      quizID,
		BlockIndex:  blockIndex,
		Responses:   make(map[int]int),
		Responded:   make(map[uuid.UUID]bool),
		ShowResults: false,
		Closed:      false,
	}
}

// RecordPollResponse records a response to the active poll
func (s *LiveSession) RecordPollResponse(learnerID uuid.UUID, answerIndex int) bool {
	s.Lock()
	defer s.Unlock()
	if s.Poll == nil || !s.Poll.Active || s.Poll.Closed {
		return false
	}
	if s.Poll.Responded[learnerID] {
		return false // Already responded
	}
	s.Poll.Responses[answerIndex]++
	s.Poll.Responded[learnerID] = true
	return true
}

// ClosePoll closes the poll (no more responses) and shows results
func (s *LiveSession) ClosePoll() {
	s.Lock()
	defer s.Unlock()
	if s.Poll != nil {
		s.Poll.Closed = true
		s.Poll.ShowResults = true
	}
}

// EndPoll completely ends the poll
func (s *LiveSession) EndPoll() {
	s.Lock()
	defer s.Unlock()
	s.Poll = nil
}

// TogglePollResults toggles visibility of poll results
func (s *LiveSession) TogglePollResults() {
	s.Lock()
	defer s.Unlock()
	if s.Poll != nil {
		s.Poll.ShowResults = !s.Poll.ShowResults
	}
}

// annotationKey builds the map key for annotations
func annotationKey(module string, page int) string {
	return fmt.Sprintf("%s:%d", module, page)
}

// AddAnnotation appends a stroke to the annotations for a given page
func (s *LiveSession) AddAnnotation(module string, page int, stroke AnnotationStroke) {
	s.Lock()
	defer s.Unlock()
	key := annotationKey(module, page)
	s.Annotations[key] = append(s.Annotations[key], stroke)
}

// ClearAnnotations removes all annotations for a given page
func (s *LiveSession) ClearAnnotations(module string, page int) {
	s.Lock()
	defer s.Unlock()
	delete(s.Annotations, annotationKey(module, page))
}

// GetAnnotations returns the strokes for a given page
func (s *LiveSession) GetAnnotations(module string, page int) []AnnotationStroke {
	s.RLock()
	defer s.RUnlock()
	return s.Annotations[annotationKey(module, page)]
}

// SetAgent stores the presenter's agent state for followers to mirror
func (s *LiveSession) SetAgent(state *AgentState) {
	s.Lock()
	defer s.Unlock()
	s.Agent = state
}

// GetAgent returns the presenter's agent state, or nil
func (s *LiveSession) GetAgent() *AgentState {
	s.RLock()
	defer s.RUnlock()
	return s.Agent
}

// ClearAgent removes the agent state from the live session
func (s *LiveSession) ClearAgent() {
	s.Lock()
	defer s.Unlock()
	s.Agent = nil
}
