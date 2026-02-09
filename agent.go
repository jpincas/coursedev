package main

import "strings"

// AgentState holds per-session ephemeral state for an agent block walkthrough.
// Initialised fresh when navigating to a page with an agent block, nil'd when navigating away.
type AgentState struct {
	// Config — copied from the parsed AgentBlock at page load
	Config *AgentBlock

	// Script playback position — index of the next event to process
	ScriptIndex int

	// Accumulated display state
	ChatMessages    []ChatMessage
	ContextMessages []ContextMessage
	CurrentNote     string

	// Simulated scratchpad — map of filename to content
	Scratchpad        map[string]string
	InitialScratchpad map[string]string

	// Simulated metrics
	TokenCount int

	// UI toggles
	WorkspaceOpen   bool
	SidebarOpen     bool
	ScratchpadOpen  bool
	ShowSystem      bool
	ShowToolCalls   bool
	ShowFullContext  bool
	ViewingFile     string
	ExpandedFolders map[string]bool // Track which folder paths are expanded in explorer
}

// ChatMessage is a single entry in the agent chat panel
type ChatMessage struct {
	Type     string            // "user", "assistant", "tool_call", "tool_result", "compaction_divider", "clear_divider"
	Content  string
	ToolName string
	ToolArgs map[string]string
	GroupIdx int // Position within its event group (for staggered animation); -1 = already displayed
}

// ContextMessage is a single entry in the full context panel
type ContextMessage struct {
	Role    string // "system", "user", "assistant", "tool_call", "tool_result", "compaction_summary"
	Content string
}

// initAgentState creates a fresh AgentState from an AgentBlock config
func initAgentState(block *AgentBlock) *AgentState {
	initial := make(map[string]string)
	for k, v := range block.Scratchpad {
		initial[k] = v
	}
	current := make(map[string]string)
	for k, v := range block.Scratchpad {
		current[k] = v
	}

	state := &AgentState{
		Config:            block,
		ScriptIndex:       0,
		ChatMessages:      nil,
		Scratchpad:        current,
		InitialScratchpad: initial,
		ExpandedFolders:   make(map[string]bool),
		SidebarOpen:       block.Sidebar.StartOpen,
		ShowSystem:        false, // modal starts closed; opened via title bar icon
		ShowToolCalls:     block.Visibility.ToolCalls == "visible",
		ShowFullContext:    false, // toggled via title bar icon; always available
	}

	// Auto-expand all folders that exist in the initial scratchpad
	for path := range current {
		parts := strings.Split(path, "/")
		for i := 1; i < len(parts); i++ {
			state.ExpandedFolders[strings.Join(parts[:i], "/")] = true
		}
	}

	// Seed context with system prompt
	if block.System != "" {
		state.ContextMessages = []ContextMessage{
			{Role: "system", Content: block.System},
		}
		state.TokenCount = estimateTokens(block.System)
	}

	return state
}

// findAgentBlock scans a slice of blocks for the first *AgentBlock
func findAgentBlock(blocks []Block) *AgentBlock {
	for _, b := range blocks {
		if ab, ok := b.(*AgentBlock); ok {
			return ab
		}
	}
	return nil
}

// estimateTokens provides a rough token count estimate (~4 chars per token)
func estimateTokens(text string) int {
	return len(text) / 4
}

// initAgentIfNeeded checks the current page for an agent block and initialises state if found.
// When presenting, also pushes the initial agent state to the live session so followers can see it.
func (m *Model) initAgentIfNeeded() {
	page := m.currentPage()
	if page == nil {
		return
	}
	if ab := findAgentBlock(page.Blocks); ab != nil {
		m.Agent = initAgentState(ab)
		if m.IsPresenting {
			if session := m.getPresentingSession(); session != nil {
				session.SetAgent(m.Agent)
			}
		}
	}
}

// getAgentStateForRender returns the agent state to render.
// When following a live session, returns the instructor's agent state.
func (m *Model) getAgentStateForRender() *AgentState {
	if m.FollowingLive {
		if session := m.getFollowingSession(); session != nil {
			return session.GetAgent()
		}
	}
	return m.Agent
}
