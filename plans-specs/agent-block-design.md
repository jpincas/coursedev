# Agent Block — Design Document

## What It Is

An `agent` block embeddable in training markdown that gives learners a **scripted, interactive walkthrough** of an AI conversation. The course creator choreographs the entire exchange — user messages, assistant responses, tool calls, compactions, resets — and the learner steps through it like a guided demo. The sidebar becomes a conversation player, not a live chat client.

This is a deliberate design choice. Scripted conversations give the course creator full control over what the learner sees, ensuring demos always illustrate the intended concept clearly. No API variance, no hallucinations derailing a lesson, no rate limits during a live class.

The agent renders in a persistent right sidebar alongside the lesson content.

---

## Concepts

### Script Events

A script is an ordered list of events. Each event has a `type` that determines what it does:

| Type | What It Does |
|---|---|
| `note` | Displays a pedagogical note to the learner. The instructor's voice. Doesn't appear in the chat or context — it's outside the simulation. |
| `user` | Adds a user message to the chat and context window. |
| `assistant` | Adds an assistant response to the chat and context window. |
| `tool_call` | Adds a tool invocation to the chat (if visible) and context window. |
| `tool_result` | Adds a tool response to the chat (if visible) and context window. If the tool is `scratchpad_write`, also mutates the scratchpad state. |
| `compaction` | Replaces all context messages before this point with a summary. Demonstrates how real systems manage token limits. |
| `clear` | Resets conversation and optionally scratchpad back to initial state. A fresh start mid-script. |

### Step Grouping

Events fire in logical groups so the learner isn't clicking "Next" for every individual message. The grouping rules:

- A **`note`** is always its own step. The learner reads it, then clicks to advance.
- A **`user`** event starts a group. All subsequent `tool_call`, `tool_result`, and `assistant` events auto-advance (with a brief staggered delay for visual pacing) until hitting the next `note`, `user`, `compaction`, or `clear`.
- A **`compaction`** is its own step.
- A **`clear`** is its own step.

So the learner's experience is: read a note → click Next → watch an exchange play out with natural pacing → read the next note → click Next → watch the next exchange. Like stepping through a debugger, but for conversations.

### Scratchpad

The scratchpad is a simulated filesystem — a `map[string]string` on the agent state. The block can declare initial files, and `scratchpad_write` tool results in the script mutate the map. The scratchpad panel in the sidebar shows the current file state, and learners can click to view any file. It's props on a stage, not a real filesystem.

### Visibility Controls

The course creator controls what "backstage" information the learner sees, turning the sidebar into a teaching tool about how AI systems work:

- **`visible`** — Always shown.
- **`hidden`** — Never shown. Learner doesn't know it exists.
- **`toggleable`** — Hidden by default with a toggle to reveal.

---

## Markdown Block Format

### Minimal Example

````markdown
```agent
id: basic-chat-demo
title: "Simple Conversation"

system: |
  You are a helpful assistant.

script:
  - type: user
    content: "What's the capital of France?"

  - type: assistant
    content: "The capital of France is Paris."
```
````

### Full-Featured Example

````markdown
```agent
id: context-window-demo
title: "How Context Windows Work"

# Model label shown in status bar (purely cosmetic — no real API call)
model_label: llama-3.3-70b

# System prompt — the learner may or may not see this depending on visibility
system: |
  You are a helpful assistant with access to a scratchpad
  for reading and writing files.

# Initial scratchpad state — files that exist before the demo starts
scratchpad:
  "project-brief.md": |
    # Project Atlas
    A distributed caching layer for the analytics pipeline.
    Target latency: <50ms p99. Budget: $200k.
  "meeting-notes.md": |
    # Standup 2025-01-15
    - Cache hit rate dropped to 60% after deploy
    - Sarah investigating memory pressure on node-3

# Tools the agent "has" — for display in system prompt / context panel
tools:
  - scratchpad

# What the learner can see
visibility:
  system_prompt: toggleable
  tool_calls: visible
  full_context: visible
  token_count: visible
  model_name: visible

# Sidebar appearance
sidebar:
  width: 45%
  start_open: true

# The choreographed conversation
script:
  - type: note
    text: |
      This demo shows how an AI agent uses tools. Open the scratchpad
      panel to see the pre-loaded files, then click Next to begin.

  - type: user
    content: "What's the project budget?"

  - type: tool_call
    tool: scratchpad_read
    args:
      filename: "project-brief.md"

  - type: tool_result
    tool: scratchpad_read
    content: |
      # Project Atlas
      A distributed caching layer for the analytics pipeline.
      Target latency: <50ms p99. Budget: $200k.

  - type: assistant
    content: "The budget for Project Atlas is $200k."

  - type: note
    text: |
      Notice how the tool call and result both appear in the context
      panel. The model 'sees' the file contents as part of the
      conversation history. Now watch it synthesise across documents.

  - type: user
    content: "Write a summary combining both documents."

  - type: tool_call
    tool: scratchpad_read
    args:
      filename: "meeting-notes.md"

  - type: tool_result
    tool: scratchpad_read
    content: |
      # Standup 2025-01-15
      - Cache hit rate dropped to 60% after deploy
      - Sarah investigating memory pressure on node-3

  - type: tool_call
    tool: scratchpad_write
    args:
      filename: "summary.md"
      content: |
        # Project Atlas Summary
        ## Overview
        Distributed caching layer, $200k budget, <50ms p99 target.
        ## Current Status
        Cache hit rate at 60% post-deploy. Memory pressure on node-3
        under investigation by Sarah.

  - type: tool_result
    tool: scratchpad_write
    content: "Written 204 bytes to summary.md"

  - type: assistant
    content: "I've created summary.md combining the key points from both documents. Check the scratchpad to see it."

  - type: note
    text: |
      Check the scratchpad — there are now three files. The agent
      synthesised a new document from the existing ones.

      Now watch what happens when we compact the context. The detailed
      message history gets replaced with a summary.

  - type: compaction
    summary: |
      [Earlier: user asked about project budget ($200k). Agent read
      project-brief.md and meeting-notes.md, then created summary.md
      combining key points from both documents.]

  - type: note
    text: |
      Look at the context panel. The detailed back-and-forth is gone,
      replaced by a summary. What happens when we ask about specifics
      that were in the compacted messages?

  - type: user
    content: "What was the exact cache hit rate from the standup notes?"

  - type: assistant
    content: |
      I know I reviewed some standup notes earlier, but I don't have the
      specific numbers in my current context. Let me check the file again.

  - type: tool_call
    tool: scratchpad_read
    args:
      filename: "meeting-notes.md"

  - type: tool_result
    tool: scratchpad_read
    content: |
      # Standup 2025-01-15
      - Cache hit rate dropped to 60% after deploy
      - Sarah investigating memory pressure on node-3

  - type: assistant
    content: "The cache hit rate dropped to 60% after the deploy."

  - type: note
    text: |
      Interesting — the model had to re-read the file because the
      specific number was lost during compaction. The summary only
      said "key points" without exact figures. This is the fundamental
      tradeoff of compaction: you save tokens but lose detail.

  - type: clear
    reset_scratchpad: true

  - type: note
    text: |
      We've reset everything — conversation cleared, scratchpad back
      to its starting state. That concludes the context window demo.
```
````

---

## Script Event Schemas

### `note`

```yaml
- type: note
  text: "Instructional text shown to the learner."
```

Displayed as a styled callout above the chat area (or between chat messages as a divider). Not part of the simulated conversation — it's the course creator speaking directly to the learner.

### `user`

```yaml
- type: user
  content: "The user's message text."
```

Appears as a user chat bubble. Added to the context messages array.

### `assistant`

```yaml
- type: assistant
  content: "The assistant's response text."
  tokens: 45    # Optional: simulated token count for this response
```

Appears as an assistant chat bubble. Added to the context messages array. The optional `tokens` field increments the simulated token counter for realism.

### `tool_call`

```yaml
- type: tool_call
  tool: scratchpad_read
  args:
    filename: "notes.md"
```

Appears in the chat as a styled tool invocation block (if `tool_calls` visibility allows). Added to context messages. The `args` field is a free-form map rendered as formatted key-value pairs.

### `tool_result`

```yaml
- type: tool_result
  tool: scratchpad_read
  content: "File contents here..."
```

Appears in the chat as a styled tool response block (if `tool_calls` visibility allows). Added to context messages.

**Side effects:** If `tool` is `scratchpad_write`, the corresponding `args` from the preceding `tool_call` are used to update the scratchpad map. The content of the `tool_result` is what's shown in the chat ("Written N bytes to file.md"), while the actual file content comes from the `tool_call`'s args. This keeps the scratchpad in sync with what the script describes happening.

### `compaction`

```yaml
- type: compaction
  summary: "Summary text that replaces prior context messages."
```

Replaces all context messages (except the system prompt) with a single summary message. The chat display shows a visual divider: "── Context compacted ──". The full context panel updates to show just the system prompt + summary + any messages after this point.

The simulated token count drops to reflect the shorter context.

### `clear`

```yaml
- type: clear
  reset_scratchpad: false    # default: false
  note: "Starting fresh."    # optional: shown as a divider in chat
```

Resets the conversation: chat messages cleared, context messages reset to just the system prompt. If `reset_scratchpad: true`, the scratchpad map resets to the initial state declared in the block's `scratchpad` field (or empty if none). The optional `note` is shown as a divider in the chat area.

---

## State Model

### Agent State (Per-Session, Ephemeral)

```go
type AgentState struct {
    // Config — copied from the parsed AgentBlock at page load
    Config AgentBlock

    // Script playback position
    ScriptIndex int  // Index of the next event to process

    // Accumulated display state
    ChatMessages    []ChatMessage    // What's shown in the chat panel
    ContextMessages []ContextMessage // What's shown in the full context panel
    CurrentNote     string           // Currently displayed note text (empty if none)

    // Simulated scratchpad — map of filename to content
    Scratchpad        map[string]string // Current state
    InitialScratchpad map[string]string // Starting state (for reset)

    // Simulated metrics
    TokenCount int

    // UI toggles
    SidebarOpen    bool
    ScratchpadOpen bool
    ShowSystem     bool
    ShowToolCalls  bool
    ShowFullContext bool
    ViewingFile    string // Filename currently being viewed (empty if none)
}

type ChatMessage struct {
    Type     string // "user", "assistant", "tool_call", "tool_result",
                    // "compaction_divider", "clear_divider"
    Content  string
    ToolName string            // For tool_call and tool_result
    ToolArgs map[string]string // For tool_call
    GroupIdx int               // Position within its event group (for staggered animation)
}

type ContextMessage struct {
    Role    string // "system", "user", "assistant", "tool_call",
                   // "tool_result", "compaction_summary"
    Content string
}
```

### Integration with Main Model

```go
type Model struct {
    gt.Router
    // ... existing fields ...

    // Agent state — nil when current page has no agent block
    Agent *AgentState
}
```

Agent state is **ephemeral**. It is initialised fresh when navigating to a page with an agent block and nil'd out when navigating away. No persistence across page transitions.

### Initialisation

When the learner navigates to a page with an agent block, the page navigation handler initialises the agent state:

```go
func initAgentState(block AgentBlock) *AgentState {
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
        Scratchpad:        current,
        InitialScratchpad: initial,
        SidebarOpen:       block.Sidebar.StartOpen,
        ShowSystem:        block.Visibility.SystemPrompt == "visible",
        ShowToolCalls:     block.Visibility.ToolCalls == "visible",
        ShowFullContext:    block.Visibility.FullContext == "visible",
    }

    // Seed context with system prompt
    state.ContextMessages = []ContextMessage{
        {Role: "system", Content: block.System},
    }

    // Include initial token estimate for system prompt
    state.TokenCount = estimateTokens(block.System)

    return state
}
```

### Navigation Handlers Clear Agent State

```go
func (m *Model) handleNextPage(msg gt.Message, s gt.State) gt.Response {
    model := s.(*Model)
    mod := model.Modules[model.CurrentModule]

    if model.CurrentPage < len(mod.Pages)-1 {
        model.CurrentPage++
        model.ActiveQuiz = nil
        model.ActiveHotspot = ""
        model.Agent = nil // Clear agent state
    }

    // Check if new page has an agent block — initialise if so
    newPage := mod.Pages[model.CurrentPage]
    if agentBlock := findAgentBlock(newPage.Blocks); agentBlock != nil {
        model.Agent = initAgentState(*agentBlock)
    }

    return gt.Respond()
}
```

---

## Message Handlers

```go
// In Update() MessageMap:

// Script advancement
"AGENT_ADVANCE":           m.handleAgentAdvance,

// Visibility toggles
"AGENT_TOGGLE_SYSTEM":     m.handleAgentToggleSystem,
"AGENT_TOGGLE_TOOLS":      m.handleAgentToggleTools,
"AGENT_TOGGLE_CONTEXT":    m.handleAgentToggleContext,

// Sidebar controls
"AGENT_TOGGLE_SIDEBAR":    m.handleAgentToggleSidebar,
"AGENT_TOGGLE_SCRATCHPAD": m.handleAgentToggleScratchpad,

// Scratchpad navigation
"AGENT_VIEW_FILE":         m.handleAgentViewFile,

// Reset
"AGENT_RESET":             m.handleAgentReset,
```

### The Advance Handler

This is the core of the playback system. When the learner clicks "Next," it processes the next logical group of events:

```go
func (m *Model) handleAgentAdvance(msg gt.Message, s gt.State) gt.Response {
    model := s.(*Model)
    agent := model.Agent
    if agent == nil {
        return gt.Respond()
    }

    script := agent.Config.Script
    if agent.ScriptIndex >= len(script) {
        return gt.Respond() // Script complete
    }

    event := script[agent.ScriptIndex]

    switch event.Type {

    case "note":
        agent.CurrentNote = event.Text
        agent.ScriptIndex++

    case "user":
        agent.CurrentNote = ""
        agent.ScriptIndex = processEventGroup(agent, agent.ScriptIndex)

    case "compaction":
        agent.CurrentNote = ""
        processCompaction(agent, event)
        agent.ScriptIndex++

    case "clear":
        agent.CurrentNote = ""
        processClear(agent, event)
        agent.ScriptIndex++
    }

    // If instructor, broadcast
    if model.IsInstructor {
        sharedPresentation.Agent = model.Agent
        app.Broadcast()
    }

    return gt.Respond()
}
```

### Event Group Processing

```go
// processEventGroup processes a user-initiated group: the user message
// and all subsequent tool_call, tool_result, assistant events until
// hitting a boundary (note, user, compaction, clear, or end of script).
// Returns the new script index.
func processEventGroup(agent *AgentState, startIdx int) int {
    script := agent.Config.Script
    idx := startIdx
    groupPosition := 0

    for idx < len(script) {
        event := script[idx]

        switch event.Type {
        case "user":
            if idx > startIdx {
                return idx // Next user message = new group
            }
            appendUserEvent(agent, event, groupPosition)
            groupPosition++
            idx++

        case "assistant":
            appendAssistantEvent(agent, event, groupPosition)
            groupPosition++
            idx++

        case "tool_call":
            appendToolCallEvent(agent, event, groupPosition)
            groupPosition++
            idx++

        case "tool_result":
            appendToolResultEvent(agent, event, groupPosition)
            applyToolSideEffects(agent, script, idx)
            groupPosition++
            idx++

        default:
            return idx // Boundary: note, compaction, clear
        }
    }

    return idx
}
```

### Event Processors

```go
func appendUserEvent(agent *AgentState, event ScriptEvent, groupIdx int) {
    agent.ChatMessages = append(agent.ChatMessages, ChatMessage{
        Type:     "user",
        Content:  event.Content,
        GroupIdx: groupIdx,
    })
    agent.ContextMessages = append(agent.ContextMessages, ContextMessage{
        Role:    "user",
        Content: event.Content,
    })
    agent.TokenCount += estimateTokens(event.Content)
}

func appendAssistantEvent(agent *AgentState, event ScriptEvent, groupIdx int) {
    agent.ChatMessages = append(agent.ChatMessages, ChatMessage{
        Type:     "assistant",
        Content:  event.Content,
        GroupIdx: groupIdx,
    })
    agent.ContextMessages = append(agent.ContextMessages, ContextMessage{
        Role:    "assistant",
        Content: event.Content,
    })
    tokens := event.Tokens
    if tokens == 0 {
        tokens = estimateTokens(event.Content)
    }
    agent.TokenCount += tokens
}

func appendToolCallEvent(agent *AgentState, event ScriptEvent, groupIdx int) {
    agent.ChatMessages = append(agent.ChatMessages, ChatMessage{
        Type:     "tool_call",
        ToolName: event.Tool,
        ToolArgs: event.Args,
        GroupIdx: groupIdx,
    })
    argsJSON, _ := json.Marshal(event.Args)
    agent.ContextMessages = append(agent.ContextMessages, ContextMessage{
        Role:    "tool_call",
        Content: fmt.Sprintf("%s(%s)", event.Tool, string(argsJSON)),
    })
}

func appendToolResultEvent(agent *AgentState, event ScriptEvent, groupIdx int) {
    agent.ChatMessages = append(agent.ChatMessages, ChatMessage{
        Type:     "tool_result",
        ToolName: event.Tool,
        Content:  event.Content,
        GroupIdx: groupIdx,
    })
    agent.ContextMessages = append(agent.ContextMessages, ContextMessage{
        Role:    "tool_result",
        Content: event.Content,
    })
}

func applyToolSideEffects(agent *AgentState, script []ScriptEvent, idx int) {
    event := script[idx]
    if event.Tool != "scratchpad_write" {
        return
    }
    // Find the preceding tool_call to get the write args
    for i := idx - 1; i >= 0; i-- {
        if script[i].Type == "tool_call" && script[i].Tool == "scratchpad_write" {
            filename := script[i].Args["filename"]
            content := script[i].Args["content"]
            agent.Scratchpad[filename] = content
            return
        }
    }
}

func processCompaction(agent *AgentState, event ScriptEvent) {
    agent.ContextMessages = []ContextMessage{
        agent.ContextMessages[0], // Keep system prompt
        {Role: "compaction_summary", Content: event.Summary},
    }
    agent.ChatMessages = append(agent.ChatMessages, ChatMessage{
        Type:    "compaction_divider",
        Content: event.Summary,
    })
    agent.TokenCount = estimateTokens(agent.Config.System) +
        estimateTokens(event.Summary)
}

func processClear(agent *AgentState, event ScriptEvent) {
    agent.ChatMessages = nil
    agent.ContextMessages = []ContextMessage{
        {Role: "system", Content: agent.Config.System},
    }
    agent.TokenCount = estimateTokens(agent.Config.System)

    if event.Note != "" {
        agent.ChatMessages = append(agent.ChatMessages, ChatMessage{
            Type:    "clear_divider",
            Content: event.Note,
        })
    }

    if event.ResetScratchpad {
        agent.Scratchpad = make(map[string]string)
        for k, v := range agent.InitialScratchpad {
            agent.Scratchpad[k] = v
        }
    }
}

func estimateTokens(text string) int {
    return len(text) / 4 // Rough approximation
}
```

---

## Parsing

### Block Struct

```go
type AgentBlock struct {
    ID    string `yaml:"id"`
    Title string `yaml:"title"`

    // Display label (cosmetic)
    ModelLabel string `yaml:"model_label"`

    // System prompt
    System string `yaml:"system"`

    // Initial scratchpad files
    Scratchpad map[string]string `yaml:"scratchpad"`

    // Tool names — for display purposes
    Tools []string `yaml:"tools"`

    // Visibility controls
    Visibility AgentVisibility `yaml:"visibility"`

    // Sidebar config
    Sidebar AgentSidebarConfig `yaml:"sidebar"`

    // The choreographed conversation
    Script []ScriptEvent `yaml:"script"`
}

type ScriptEvent struct {
    Type string `yaml:"type"` // note, user, assistant, tool_call,
                              // tool_result, compaction, clear

    // For note
    Text string `yaml:"text"`

    // For user, assistant
    Content string `yaml:"content"`
    Tokens  int    `yaml:"tokens"`

    // For tool_call, tool_result
    Tool string            `yaml:"tool"`
    Args map[string]string `yaml:"args"`

    // For compaction
    Summary string `yaml:"summary"`

    // For clear
    ResetScratchpad bool   `yaml:"reset_scratchpad"`
    Note            string `yaml:"note"`
}

type AgentVisibility struct {
    SystemPrompt string `yaml:"system_prompt"` // visible | hidden | toggleable
    ToolCalls    string `yaml:"tool_calls"`    // visible | hidden | toggleable
    FullContext  string `yaml:"full_context"`  // visible | hidden | toggleable
    TokenCount   string `yaml:"token_count"`   // visible | hidden
    ModelName    string `yaml:"model_name"`    // visible | hidden
}

type AgentSidebarConfig struct {
    Width     string `yaml:"width"`      // CSS value, default "40%"
    StartOpen bool   `yaml:"start_open"` // default: true
}
```

### Goldmark Extension

Same pattern as other block types: the `agent` fenced block parser extracts YAML into an `AgentBlock` struct, emits a placeholder marker in the HTML. The render function detects the agent block and activates the sidebar layout.

### Validation at Startup

- Agent ID must be unique within a module.
- At most one agent block per page.
- `visibility` values must be one of `visible`, `hidden`, `toggleable`.
- Script must not be empty.
- Script must start with a `note` or `user` event (not `tool_result` or `assistant`).
- Every `tool_result` must be preceded by a `tool_call` (not necessarily immediately — another `tool_call`/`tool_result` pair can intervene for parallel calls).
- `tool_call` args for `scratchpad_write` must have `filename` and `content` keys.
- `compaction` events must have a non-empty `summary`.
- Referenced tool names in script events must be in the block's `tools` list.

---

## Rendering

### Layout Decision

```go
func (m *Model) renderCurrentPage(s gt.State) []byte {
    model := s.(*Model)
    page := model.currentPage()

    agentBlock := findAgentBlock(page.Blocks)

    if agentBlock != nil {
        agentState := model.getAgentStateForRender()
        return renderLayout(model,
            h.Div(a.Attrs(a.Class("split-layout")),
                h.Div(a.Attrs(a.Class("content-pane")),
                    renderPageContent(page, model.ActiveQuiz, model.ActiveHotspot),
                    renderPageNav(model),
                ),
                renderAgentSidebar(agentState, model.FollowInstructor),
            ),
        ).Bytes()
    }

    return renderLayout(model,
        renderPageContent(page, model.ActiveQuiz, model.ActiveHotspot),
        renderPageNav(model),
    ).Bytes()
}

func (m *Model) getAgentStateForRender() *AgentState {
    if m.FollowInstructor && sharedPresentation.Agent != nil {
        return sharedPresentation.Agent
    }
    return m.Agent
}
```

### Agent Sidebar

```go
func renderAgentSidebar(state *AgentState, followMode bool) h.Element {
    if state == nil {
        return h.Nothing()
    }

    if !state.SidebarOpen {
        return h.Aside(a.Attrs(a.Class("agent-sidebar collapsed")),
            renderAgentTitleBar(state),
        )
    }

    children := []h.Element{
        renderAgentTitleBar(state),
    }

    // System prompt panel
    if shouldShow(state.Config.Visibility.SystemPrompt, state.ShowSystem) {
        children = append(children, renderSystemPromptPanel(state))
    }
    if state.Config.Visibility.SystemPrompt == "toggleable" {
        children = append(children, renderToggleButton(
            "AGENT_TOGGLE_SYSTEM", "System Prompt", state.ShowSystem))
    }

    // Current note
    if state.CurrentNote != "" {
        children = append(children, renderCurrentNote(state.CurrentNote))
    }

    // Chat messages
    children = append(children, renderChatMessages(state))

    // Full context panel
    if shouldShow(state.Config.Visibility.FullContext, state.ShowFullContext) {
        children = append(children, renderFullContextPanel(state))
    }
    if state.Config.Visibility.FullContext == "toggleable" {
        children = append(children, renderToggleButton(
            "AGENT_TOGGLE_CONTEXT", "Full Context", state.ShowFullContext))
    }

    // Scratchpad panel
    if containsTool(state.Config.Tools, "scratchpad") {
        children = append(children, renderScratchpadPanel(state))
    }

    // Advance control (or follow-mode banner)
    if followMode {
        children = append(children, renderFollowModeBanner())
    } else {
        children = append(children, renderAdvanceControl(state))
    }

    // Status bar
    children = append(children, renderStatusBar(state))

    return h.Aside(a.Attrs(a.Class("agent-sidebar")), children...)
}
```

### Chat Message Rendering

```go
func renderChatMessages(state *AgentState) h.Element {
    var msgs []h.Element

    for _, m := range state.ChatMessages {
        switch m.Type {
        case "user":
            msgs = append(msgs, renderUserBubble(m))

        case "assistant":
            msgs = append(msgs, renderAssistantBubble(m))

        case "tool_call":
            if shouldShow(state.Config.Visibility.ToolCalls, state.ShowToolCalls) {
                msgs = append(msgs, renderToolCallBubble(m))
            }

        case "tool_result":
            if shouldShow(state.Config.Visibility.ToolCalls, state.ShowToolCalls) {
                msgs = append(msgs, renderToolResultBubble(m))
            }

        case "compaction_divider":
            msgs = append(msgs, h.Div(a.Attrs(a.Class("chat-divider compaction")),
                h.Span(a.Attrs(), h.Text("── Context compacted ──")),
            ))

        case "clear_divider":
            msgs = append(msgs, h.Div(a.Attrs(a.Class("chat-divider clear")),
                h.Span(a.Attrs(), h.Text(m.Content)),
            ))
        }
    }

    return h.Div(a.Attrs(a.Class("chat-messages"), a.Id("agent-messages")),
        msgs...,
    )
}
```

### Tool Call / Result Bubbles

```go
func renderToolCallBubble(m ChatMessage) h.Element {
    var argLines []h.Element
    for k, v := range m.ToolArgs {
        argLines = append(argLines,
            h.Div(a.Attrs(a.Class("tool-arg")),
                h.Span(a.Attrs(a.Class("tool-arg-key")), h.Text(k+": ")),
                h.Span(a.Attrs(a.Class("tool-arg-val")), h.Text(v)),
            ),
        )
    }

    return h.Div(a.Attrs(a.Class("chat-msg tool-call")),
        h.Div(a.Attrs(a.Class("tool-header")),
            h.Span(a.Attrs(a.Class("tool-icon")), h.Text("🔧")),
            h.Span(a.Attrs(), h.Text("Tool Call: "+m.ToolName)),
        ),
        h.Div(a.Attrs(a.Class("tool-args")), argLines...),
    )
}

func renderToolResultBubble(m ChatMessage) h.Element {
    return h.Div(a.Attrs(a.Class("chat-msg tool-result")),
        h.Div(a.Attrs(a.Class("tool-header")),
            h.Span(a.Attrs(a.Class("tool-icon")), h.Text("📎")),
            h.Span(a.Attrs(), h.Text("Result: "+m.ToolName)),
        ),
        h.Pre(a.Attrs(a.Class("tool-result-content")), h.Text(m.Content)),
    )
}
```

### Full Context Panel

```go
func renderFullContextPanel(state *AgentState) h.Element {
    var entries []h.Element

    for _, cm := range state.ContextMessages {
        roleClass := "context-role-" + cm.Role
        entries = append(entries,
            h.Div(a.Attrs(a.Class("context-entry "+roleClass)),
                h.Div(a.Attrs(a.Class("context-role-label")),
                    h.Text(cm.Role)),
                h.Pre(a.Attrs(a.Class("context-content")),
                    h.Text(cm.Content)),
            ),
        )
    }

    return h.Div(a.Attrs(a.Class("full-context-panel")),
        h.Div(a.Attrs(a.Class("panel-header")), h.Text("Full Context")),
        h.Div(a.Attrs(a.Class("context-entries")), entries...),
    )
}
```

### Advance Control

```go
func renderAdvanceControl(state *AgentState) h.Element {
    atEnd := state.ScriptIndex >= len(state.Config.Script)

    if atEnd {
        return h.Div(a.Attrs(a.Class("agent-advance complete")),
            h.Span(a.Attrs(), h.Text("✓ Demo complete")),
            h.Button(a.Attrs(
                a.Class("btn-reset"),
                a.OnClick(gt.SendBasicMessageNoArgs("AGENT_RESET")),
            ), h.Text("Restart")),
        )
    }

    next := state.Config.Script[state.ScriptIndex]
    buttonText := "Next"
    switch next.Type {
    case "user":
        buttonText = "▶ Run next exchange"
    case "compaction":
        buttonText = "Compact context"
    case "clear":
        buttonText = "Clear & reset"
    }

    return h.Div(a.Attrs(a.Class("agent-advance")),
        h.Button(a.Attrs(
            a.Class("btn-advance"),
            a.OnClick(gt.SendBasicMessageNoArgs("AGENT_ADVANCE")),
        ), h.Text(buttonText)),
        h.Span(a.Attrs(a.Class("step-counter")),
            h.Text(fmt.Sprintf("Step %d / %d",
                state.ScriptIndex, len(state.Config.Script)))),
    )
}

func renderFollowModeBanner() h.Element {
    return h.Div(a.Attrs(a.Class("agent-follow-banner")),
        h.Text("Your instructor is leading this demo"),
    )
}
```

### Scratchpad Panel

```go
func renderScratchpadPanel(state *AgentState) h.Element {
    if !state.ScratchpadOpen {
        fileCount := len(state.Scratchpad)
        return h.Div(a.Attrs(a.Class("scratchpad-panel collapsed")),
            h.Button(a.Attrs(
                a.Class("panel-toggle"),
                a.OnClick(gt.SendBasicMessageNoArgs("AGENT_TOGGLE_SCRATCHPAD")),
            ), h.Text(fmt.Sprintf("Scratchpad (%d files) ▸", fileCount))),
        )
    }

    var files []h.Element
    for filename, content := range state.Scratchpad {
        isViewing := state.ViewingFile == filename
        fileEl := h.Div(a.Attrs(a.Class("scratchpad-file")),
            h.Div(a.Attrs(a.Class("file-header")),
                h.Span(a.Attrs(a.Class("file-icon")), h.Text("📄")),
                h.Span(a.Attrs(a.Class("file-name")), h.Text(filename)),
                h.Button(a.Attrs(
                    a.Class("btn-view-file"),
                    a.OnClick(gt.SendBasicMessage("AGENT_VIEW_FILE", filename)),
                ), h.Text(ternary(isViewing, "hide", "view"))),
            ),
        )
        if isViewing {
            fileEl = h.Div(a.Attrs(a.Class("scratchpad-file viewing")),
                h.Div(a.Attrs(a.Class("file-header")),
                    h.Span(a.Attrs(a.Class("file-icon")), h.Text("📄")),
                    h.Span(a.Attrs(a.Class("file-name")), h.Text(filename)),
                    h.Button(a.Attrs(
                        a.Class("btn-view-file"),
                        a.OnClick(gt.SendBasicMessage("AGENT_VIEW_FILE", "")),
                    ), h.Text("hide")),
                ),
                h.Pre(a.Attrs(a.Class("file-content")), h.Text(content)),
            )
        }
        files = append(files, fileEl)
    }

    return h.Div(a.Attrs(a.Class("scratchpad-panel")),
        h.Div(a.Attrs(a.Class("panel-header")),
            h.Text("Scratchpad"),
            h.Button(a.Attrs(
                a.Class("panel-toggle"),
                a.OnClick(gt.SendBasicMessageNoArgs("AGENT_TOGGLE_SCRATCHPAD")),
            ), h.Text("▾")),
        ),
        h.Div(a.Attrs(a.Class("scratchpad-files")), files...),
    )
}
```

### Status Bar

```go
func renderStatusBar(state *AgentState) h.Element {
    var items []h.Element

    if state.Config.Visibility.TokenCount == "visible" {
        items = append(items,
            h.Span(a.Attrs(a.Class("status-item")),
                h.Text(fmt.Sprintf("~%d tokens", state.TokenCount))))
    }
    if state.Config.Visibility.ModelName == "visible" && state.Config.ModelLabel != "" {
        items = append(items,
            h.Span(a.Attrs(a.Class("status-item model")),
                h.Text(state.Config.ModelLabel)))
    }

    return h.Div(a.Attrs(a.Class("agent-status-bar")), items...)
}
```

---

## Staggered Playback (Visual Pacing)

When a group fires (user → tool_call → tool_result → assistant), displaying everything instantly feels abrupt. CSS animation creates natural pacing without complicating server-side logic.

All events in a group are added to the DOM simultaneously, but each gets a sequential animation delay based on `GroupIdx`:

```css
.chat-msg {
    opacity: 0;
    animation: fadeSlideIn 300ms ease forwards;
}

.chat-msg[data-group-idx="0"] { animation-delay: 0ms; }
.chat-msg[data-group-idx="1"] { animation-delay: 400ms; }
.chat-msg[data-group-idx="2"] { animation-delay: 800ms; }
.chat-msg[data-group-idx="3"] { animation-delay: 1200ms; }
.chat-msg[data-group-idx="4"] { animation-delay: 1600ms; }
.chat-msg[data-group-idx="5"] { animation-delay: 2000ms; }

@keyframes fadeSlideIn {
    from { opacity: 0; transform: translateY(8px); }
    to   { opacity: 1; transform: translateY(0); }
}
```

The render function emits `data-group-idx` from the `GroupIdx` field on `ChatMessage`. Only messages from the most recent advance get animation — older messages have `GroupIdx` reset to -1 (no delay).

### Scroll-to-Bottom

```javascript
document.addEventListener('gotea:render', function() {
    const msgContainer = document.getElementById('agent-messages');
    if (msgContainer) {
        // Delay to account for staggered animations
        setTimeout(function() {
            msgContainer.scrollTop = msgContainer.scrollHeight;
        }, 2500);
    }
});
```

---

## Instructor Follow-Along Mode

When learners are in follow-along mode, the agent sidebar shows a **read-only mirror** of the instructor's agent session. Same broadcast pattern as page navigation.

### Shared State

```go
var sharedPresentation struct {
    CurrentModule string
    CurrentPage   int
    Agent         *AgentState // Instructor's live agent state
}
```

### Behaviour

- The instructor clicks "Next" — their handler updates their `AgentState`, copies it to `sharedPresentation.Agent`, and calls `app.Broadcast()`.
- Every learner's render picks up the instructor's state via `getAgentStateForRender()`.
- Learners see the same chat, same notes, same scratchpad mutations, same context panel — all in real-time.
- The advance button is replaced with "Your instructor is leading this demo."
- The instructor can narrate alongside: "Watch what happens when this vague prompt is sent..."

### On Release

When the instructor releases follow mode, each learner gets a **fresh** `AgentState` initialised from the block config. They don't inherit the instructor's conversation — they start their own walkthrough from step 0.

---

## File Structure Additions

```
training-app/
├── agent.go                   # AgentBlock, AgentState, ScriptEvent structs
│                              #   initAgentState, findAgentBlock
├── agent_handlers.go          # Message handlers (advance, toggles, reset, view file)
├── agent_playback.go          # Event processing (processEventGroup, processCompaction,
│                              #   processClear, append*Event, applyToolSideEffects)
├── agent_renderers.go         # All render functions for the sidebar UI
├── parsing/
│   └── goldmark_extensions.go # Add agent block parser case
├── static/
│   └── css/
│       └── agent.css          # Sidebar layout, chat bubbles, tool blocks,
│                              #   panels, animations, status bar
```

---

## Implementation Phases

### Phase A — Core Playback

- `AgentBlock` and `ScriptEvent` parsing from markdown (Goldmark extension).
- `AgentState` struct and `initAgentState`.
- `AGENT_ADVANCE` handler with event group processing.
- Sidebar layout (split-pane when agent block present).
- Chat message rendering (user + assistant bubbles).
- Note rendering.
- Advance button with step counter.
- Agent state lifecycle (init on page nav, nil on nav away).
- `AGENT_RESET` handler.

### Phase B — Tools & Scratchpad

- Tool call / tool result chat bubbles.
- Scratchpad initial state from block config.
- Scratchpad panel (file list, view/hide file contents).
- `scratchpad_write` side effects.
- Scratchpad reset on `clear` events.

### Phase C — Context & Visibility

- Full context panel rendering.
- Visibility toggle buttons and state.
- System prompt panel.
- Compaction event processing and context panel update.
- Token count display (estimated).
- Model name display in status bar.

### Phase D — Instructor Follow-Along

- Instructor advance broadcasts shared agent state.
- Learner read-only view in follow mode.
- Follow-mode banner replaces advance button.
- Fresh state on follow mode release.

### Phase E — Polish

- CSS staggered animation for event groups.
- Scroll-to-bottom JS with animation delay.
- Sidebar collapse / expand.
- Startup content validation for agent blocks.
- Clear event dividers.
- Contextual advance button text.

---

## Example: Full Lesson on Tool Use

````markdown
---
title: "How AI Agents Use Tools"
duration: 25m
notes: "Ensure scratchpad panel is visible during demo"
---

# How AI Agents Use Tools

A language model on its own can only generate text. But when we give it
**tools** — functions it can call — it becomes an agent that can interact
with the world. In this demo, you'll see exactly how tool calling works
from the inside.

## The Anatomy of a Tool Call

When a model decides to use a tool, it doesn't execute anything itself.
Instead, it outputs a structured request: "I'd like to call function X
with these arguments." The **orchestrator** (the system running the model)
executes the function and feeds the result back into the conversation.

The model then uses that result to formulate its response.

```agent
id: tool-use-demo
title: "Tool Calling Internals"
model_label: llama-3.3-70b

system: |
  You are a helpful assistant. You have a scratchpad tool for
  reading and writing files. Use it when the user asks you to
  save, read, or work with files.

scratchpad:
  "readme.md": |
    # My Project
    A simple web server written in Go.

tools:
  - scratchpad

visibility:
  system_prompt: toggleable
  tool_calls: visible
  full_context: toggleable
  token_count: visible
  model_name: visible

sidebar:
  width: 45%
  start_open: true

script:
  - type: note
    text: |
      Toggle open the system prompt to see the model's instructions.
      Notice it's told about the scratchpad tool. Also check the
      scratchpad panel — there's already a file there.

  - type: user
    content: "What files do I have?"

  - type: tool_call
    tool: scratchpad_list
    args: {}

  - type: tool_result
    tool: scratchpad_list
    content: "readme.md"

  - type: assistant
    content: "You have one file: readme.md. Would you like me to read it?"

  - type: note
    text: |
      The model chose to call `scratchpad_list` to answer your question.
      Watch the tool call and result in the chat — the model received a
      plain text list of filenames, then used that to write its response.

  - type: user
    content: "Yes, read it and then add a section about installation."

  - type: tool_call
    tool: scratchpad_read
    args:
      filename: "readme.md"

  - type: tool_result
    tool: scratchpad_read
    content: |
      # My Project
      A simple web server written in Go.

  - type: tool_call
    tool: scratchpad_write
    args:
      filename: "readme.md"
      content: |
        # My Project
        A simple web server written in Go.

        ## Installation
        ```
        go install github.com/example/myproject@latest
        ```

  - type: tool_result
    tool: scratchpad_write
    content: "Written 112 bytes to readme.md"

  - type: assistant
    content: |
      Done! I've read the file and added an Installation section.
      Check the scratchpad to see the updated readme.md.

  - type: note
    text: |
      Two things to notice:

      1. The model made TWO tool calls in sequence — read then write.
         The orchestrator executed each one and fed the results back.

      2. Check the scratchpad — readme.md has been updated. The model
         composed new content that preserved the original and added to it.

      Now toggle open the full context panel and scroll through. You can
      see the complete message history the model receives on every turn.

  - type: compaction
    summary: "[Earlier: user asked about files (has readme.md). Assistant read it and added an Installation section.]"

  - type: note
    text: |
      We just compacted the context. The detailed messages have been
      replaced with a summary. The model loses the specifics but
      retains the gist. This is the fundamental tradeoff: compaction
      saves tokens but loses detail.

      That concludes the tool calling demo.
```

## Key Takeaways

- Models don't execute tools. They **request** tool calls.
- An orchestrator executes the tools and feeds results back.
- Each tool call and result becomes part of the context window.
- Compaction trades detail for token efficiency.
````
