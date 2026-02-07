# Agent Block — Design Document

## What It Is

An `agent` block embeddable in training markdown that gives learners a real AI chat interface to practice concepts from the lesson. The course creator configures each agent instance declaratively — different system prompts, available tools, guided prompt sequences or free-range mode, and visibility controls that can reveal or hide the "machinery" behind the AI (system prompt, tool calls, full context window) for pedagogical purposes.

The agent renders in a persistent right sidebar, not inline with the narrative content. This keeps the lesson material and the practice environment side-by-side.

---

## Markdown Block Format

### Minimal Example

````markdown
```agent
id: dns-lookup-practice
title: "DNS Resolution Assistant"
system: |
  You are a networking tutor. Help the student understand DNS resolution.
  When they ask you to look up a domain, walk through the resolution process
  step by step before giving the answer.
```
````

### Full-Featured Example

````markdown
```agent
id: prompt-engineering-lab
title: "Prompt Engineering Sandbox"
model: llama-3.3-70b-versatile
max_turns: 20
temperature: 0.7

system: |
  You are a helpful assistant. You have access to a scratchpad where you
  can read and write files. Help the student with whatever they ask.

# Guided mode: learner steps through these prompts in order.
# Each step has a prompt (what gets sent) and a note (shown to the learner
# explaining the pedagogical purpose of this step).
mode: guided
steps:
  - note: "Let's start with a simple, vague prompt and see what happens."
    prompt: "Tell me about dogs."
  - note: "Now let's add specificity. Notice how the output changes."
    prompt: "List 5 health considerations for adopting an adult rescue greyhound."
  - note: "Now try adding a persona and output format constraint."
    prompt: "As a veterinarian, create a table comparing the top 3 joint supplements for greyhounds, with columns for name, active ingredient, and typical dosage."
  - note: "Free turn — write your own prompt applying what you've learned."
    prompt: ""  # Empty = learner writes their own

# Tools available to the agent
tools:
  - scratchpad    # Built-in: read/write files to cohort scratchpad
  - web_search    # Built-in: simulated or real web search

# What the learner can see — the teaching controls
visibility:
  system_prompt: visible       # visible | hidden | toggleable
  tool_calls: toggleable       # visible | hidden | toggleable
  full_context: hidden         # visible | hidden | toggleable
  token_count: visible         # visible | hidden
  temperature: visible         # visible | hidden

# Sidebar appearance
sidebar:
  width: 40%           # default 40%, range 30-60%
  start_open: true     # whether sidebar is open when page loads
```
````

### Free-Range Mode (Default)

````markdown
```agent
id: go-tutor
title: "Go Tutor"
system: |
  You are an expert Go programmer acting as a tutor. The student is
  learning about HTTP handlers. Guide them but don't give away answers
  immediately — use Socratic questioning.
mode: free
tools:
  - scratchpad
visibility:
  system_prompt: toggleable
  tool_calls: visible
```
````

---

## Modes

### `free` (Default)

Standard chat interface. The learner types whatever they want. The conversation continues until they navigate away or hit `max_turns`.

### `guided`

The course creator defines a sequence of steps. Each step has:

- **`note`** — Displayed to the learner above the chat input, explaining the purpose of this step. Think of it as the instructor whispering in their ear.
- **`prompt`** — The text that will be sent to the agent. Three variants:
  - **Pre-filled**: The prompt text is shown in the input field, read-only. The learner clicks "Send" (or a "Next Step" button) to dispatch it. They can see exactly what's being sent.
  - **Editable pre-fill**: `editable: true` on the step. The prompt is pre-filled but the learner can modify it before sending.
  - **Empty**: `prompt: ""` — The learner writes their own. This is a "now you try" step.

The learner progresses through steps in order. They can't skip ahead, but they can go back and review previous exchanges (read-only). After completing all steps, the guided section ends and the agent optionally unlocks into free-range mode (`free_after_guided: true`).

### Step Schema

```yaml
steps:
  - note: "Explanation shown to learner"
    prompt: "Text sent to the agent"
    editable: false           # default false; true lets learner modify
    free_after_guided: true   # only on last step; unlocks free mode after
```

---

## Tools

Tools are declared by name in the agent block. The server provides tool definitions to the Groq API and handles tool call execution server-side.

### Built-in Tools

#### `scratchpad`

Read and write files to a per-cohort shared directory. This is the primary "hands-on" tool — learners and the agent can collaborate on files.

```
Tool: scratchpad_read
Parameters: { "filename": "notes.md" }
Returns: file contents or "File not found"

Tool: scratchpad_write
Parameters: { "filename": "notes.md", "content": "..." }
Returns: "Written N bytes to notes.md"

Tool: scratchpad_list
Parameters: {}
Returns: list of files in the scratchpad
```

Implementation: Files stored at `data/scratchpads/{cohort_id}/`. Simple filesystem operations. Files are capped at a reasonable size (e.g. 50KB) to prevent abuse. The scratchpad directory is displayed as a collapsible panel below the chat in the sidebar, showing current files with ability to view them.

#### `web_search` (Optional, Later Phase)

Could be real (via a search API) or simulated (course creator provides canned results for expected queries). The simulated version is interesting pedagogically — you can control exactly what the agent "finds".

### Custom Tools (Later Phase)

Course creators could define custom tools with fixed responses to demonstrate tool-use concepts:

```yaml
tools:
  - name: get_weather
    description: "Get current weather for a city"
    parameters:
      city: { type: string, description: "City name" }
    mock_response: |
      {"temperature": 22, "conditions": "partly cloudy", "city": "${city}"}
```

This is powerful for teaching about function calling — the learner sees the tool definition, watches the model decide to call it, sees the response injected, and sees how the model incorporates it. All without any real API needed.

---

## Visibility Controls

This is the key pedagogical differentiator. The course creator controls what "backstage" information the learner can see, turning the chat into a teaching tool about how AI systems work, not just a chat window.

### Visibility Levels

- **`visible`** — Always shown. Can't be hidden.
- **`hidden`** — Never shown. Learner doesn't know it exists.
- **`toggleable`** — Hidden by default with a toggle button to reveal. The learner can peek behind the curtain.

### What Can Be Controlled

| Element | Key | What It Shows |
|---|---|---|
| System prompt | `system_prompt` | The full system message, displayed in a collapsible panel at the top of the sidebar |
| Tool calls | `tool_calls` | When the agent calls a tool, show the function name, arguments, and response |
| Full context | `full_context` | The complete messages array being sent to the API on each turn — shows how context accumulates |
| Token count | `token_count` | Running token count for the conversation (input + output) |
| Temperature | `temperature` | The temperature setting, potentially with a slider to adjust it live |
| Model name | `model_name` | Which model is being used |

### Rendering Implications

When `tool_calls: visible`, tool invocations appear as distinct styled blocks in the conversation:

```
🔧 Tool Call: scratchpad_write
   filename: "solution.go"
   content: "package main..."
   
📎 Tool Result: Written 142 bytes to solution.go
```

When `full_context: toggleable` and toggled on, a panel shows the raw JSON messages array, updating after each turn. This is gold for teaching about context windows, token limits, and how conversation history works.

---

## Sidebar Layout

### Layout Architecture

When a page has an agent block, the layout shifts from single-column to split-pane:

```
┌─────────────────────────────────────────────────────────┐
│  Header / Navigation                                     │
├───────────────────────────────┬──────────────────────────┤
│                               │                          │
│   Lesson Content              │   Agent Sidebar          │
│   (narrative + inline blocks) │                          │
│                               │ ┌──────────────────────┐ │
│                               │ │ Title Bar + Controls │ │
│                               │ ├──────────────────────┤ │
│                               │ │ [System Prompt]      │ │
│                               │ │ (if visible)         │ │
│                               │ ├──────────────────────┤ │
│                               │ │ [Guided Step Note]   │ │
│                               │ │ (if guided mode)     │ │
│                               │ ├──────────────────────┤ │
│                               │ │                      │ │
│                               │ │ Chat Messages        │ │
│                               │ │                      │ │
│                               │ │                      │ │
│                               │ ├──────────────────────┤ │
│                               │ │ [Scratchpad Panel]   │ │
│                               │ │ (collapsible)        │ │
│                               │ ├──────────────────────┤ │
│                               │ │ Input + Send         │ │
│                               │ │ [Token Count] [Temp] │ │
│                               │ └──────────────────────┘ │
│                               │                          │
├───────────────────────────────┴──────────────────────────┤
│  Page Navigation (prev/next)                             │
└─────────────────────────────────────────────────────────┘
```

### Sidebar Behaviour

- **Only one agent per page.** If a markdown file has multiple agent blocks, that's a content validation error caught at startup.
- **Sidebar presence changes layout.** The render function checks whether the current page has an agent block. If yes, the content area gets `width: 60%` and the sidebar gets `width: 40%` (configurable).
- **Collapsible.** A toggle button in the header allows minimising the sidebar to a thin strip with just the agent title, freeing up screen space for reading.
- **Scrolls independently.** The lesson content and the chat have independent scroll contexts.
- **Persists only for the current page.** Navigating away destroys the agent state. Navigating back starts fresh. This is intentional — each agent instance is a self-contained exercise.

---

## State Model

### Agent State (Per-Session)

```go
type AgentState struct {
    // Config — copied from the parsed AgentBlock at activation time
    Config    AgentBlock

    // Conversation
    Messages  []AgentMessage        // Display messages (what the learner sees)
    APIMessages []groq.ChatMessage  // Full API message history (may include hidden tool calls)
    
    // Guided mode
    CurrentStep   int     // Index into Config.Steps
    StepCompleted bool    // Whether current step's response has been received
    GuidedDone    bool    // All steps completed
    FreeUnlocked  bool    // Guided finished and free_after_guided is true
    
    // UI state
    SidebarOpen     bool
    ScratchpadOpen  bool
    ShowSystem      bool    // Current toggle state (only relevant if toggleable)
    ShowToolCalls   bool
    ShowFullContext  bool
    
    // Operational
    Loading       bool      // Waiting for Groq API response
    Error         string    // Last error message, if any
    TotalTokens   int       // Running token count
    InputDraft    string    // Current input field content (for editable guided steps)
}

type AgentMessage struct {
    Role      string    // "user", "assistant", "tool_call", "tool_result", "system", "step_note"
    Content   string
    ToolName  string    // For tool_call/tool_result messages
    ToolArgs  string    // JSON string of tool arguments
    Timestamp time.Time
    StepIndex int       // Which guided step this belongs to (-1 for free)
}
```

### Integration with Main Model

```go
type Model struct {
    gt.Router
    // ... existing fields ...
    
    // Agent state — nil if current page has no agent block,
    // or if agent hasn't been activated yet
    Agent *AgentState
}
```

The agent state is initialised when the learner navigates to a page that has an agent block. **It is ephemeral — navigating away from the page destroys the state entirely.** Navigating back to the same page starts a fresh agent session. There is no persistence of agent conversations across page navigations. This is deliberate: each page's agent is a self-contained practice environment tied to that lesson's content. The `NEXT_PAGE` / `PREV_PAGE` / `NAV_PAGE` handlers nil out `model.Agent` as part of navigation.

---

## Message Handlers

```go
// In Update() MessageMap:

// Chat interaction
"AGENT_SEND":           m.handleAgentSend,         // Send a message (free mode or empty guided step)
"AGENT_SEND_STEP":      m.handleAgentSendStep,     // Send the pre-filled guided step prompt
"AGENT_EDIT_DRAFT":     m.handleAgentEditDraft,     // Update input draft text

// Guided mode navigation  
"AGENT_NEXT_STEP":      m.handleAgentNextStep,      // Advance to next guided step

// Visibility toggles
"AGENT_TOGGLE_SYSTEM":  m.handleAgentToggleSystem,
"AGENT_TOGGLE_TOOLS":   m.handleAgentToggleTools,
"AGENT_TOGGLE_CONTEXT": m.handleAgentToggleContext,

// Sidebar controls
"AGENT_TOGGLE_SIDEBAR":    m.handleAgentToggleSidebar,
"AGENT_TOGGLE_SCRATCHPAD": m.handleAgentToggleScratchpad,

// Agent management
"AGENT_RESET":          m.handleAgentReset,         // Clear conversation, restart
"AGENT_SET_TEMP":       m.handleAgentSetTemp,       // Adjust temperature (if visible)
```

### The Send Flow

```go
func (m *Model) handleAgentSend(msg gt.Message, s gt.State) gt.Response {
    model := s.(*Model)
    if model.Agent == nil || model.Agent.Loading {
        return gt.Respond()
    }
    
    userText := msg.ArgsToString()
    if userText == "" {
        return gt.Respond()
    }
    
    // Add user message to display history
    model.Agent.Messages = append(model.Agent.Messages, AgentMessage{
        Role:    "user",
        Content: userText,
        Timestamp: time.Now(),
    })
    
    // Add to API messages
    model.Agent.APIMessages = append(model.Agent.APIMessages, groq.ChatMessage{
        Role:    "user",
        Content: userText,
    })
    
    // Set loading state — the UI will show a spinner
    model.Agent.Loading = true
    model.Agent.InputDraft = ""
    
    // Dispatch async API call via goroutine.
    // When complete, it sends a message back into this session.
    go func() {
        resp, tokens, err := callGroqAPI(model.Agent)
        // Send the result back as a message to this session
        session.Send(gt.NewMessage("AGENT_RESPONSE", AgentResponsePayload{
            Response: resp,
            Tokens:   tokens,
            Error:    err,
        }))
    }()
    
    return gt.Respond()
}

func (m *Model) handleAgentResponse(msg gt.Message, s gt.State) gt.Response {
    model := s.(*Model)
    var payload AgentResponsePayload
    msg.MustDecodeArgs(&payload)
    
    model.Agent.Loading = false
    
    if payload.Error != nil {
        model.Agent.Error = payload.Error.Error()
        return gt.Respond()
    }
    
    model.Agent.TotalTokens += payload.Tokens
    
    // Process response — may include tool calls
    for _, choice := range payload.Response.Choices {
        msg := choice.Message
        
        if msg.ToolCalls != nil {
            // Handle tool calls: execute them, add results to API messages,
            // and make another API call with the results.
            // This is a loop that continues until the model responds with
            // plain text (no more tool calls).
            model.processToolCalls(msg.ToolCalls)
            // Recurse: call API again with tool results
            go func() {
                resp, tokens, err := callGroqAPI(model.Agent)
                session.Send(gt.NewMessage("AGENT_RESPONSE", AgentResponsePayload{
                    Response: resp, Tokens: tokens, Error: err,
                }))
            }()
            return gt.Respond()
        }
        
        // Plain text response
        model.Agent.Messages = append(model.Agent.Messages, AgentMessage{
            Role:      "assistant",
            Content:   msg.Content,
            Timestamp: time.Now(),
        })
        model.Agent.APIMessages = append(model.Agent.APIMessages, groq.ChatMessage{
            Role:    "assistant",
            Content: msg.Content,
        })
    }
    
    // If in guided mode, mark step as completed
    if model.Agent.Config.Mode == "guided" && !model.Agent.GuidedDone {
        model.Agent.StepCompleted = true
    }
    
    return gt.Respond()
}
```

### Tool Call Processing

```go
func (m *Model) processToolCalls(calls []groq.ToolCall) {
    for _, call := range calls {
        // Add tool call to display messages (if visibility allows)
        m.Agent.Messages = append(m.Agent.Messages, AgentMessage{
            Role:     "tool_call",
            ToolName: call.Function.Name,
            ToolArgs: call.Function.Arguments,
            Timestamp: time.Now(),
        })
        
        // Execute the tool
        result := m.executeToolCall(call)
        
        // Add tool result to display messages
        m.Agent.Messages = append(m.Agent.Messages, AgentMessage{
            Role:     "tool_result",
            ToolName: call.Function.Name,
            Content:  result,
            Timestamp: time.Now(),
        })
        
        // Add to API messages (tool call + result)
        m.Agent.APIMessages = append(m.Agent.APIMessages, groq.ChatMessage{
            Role:       "assistant",
            ToolCalls:  []groq.ToolCall{call},
        })
        m.Agent.APIMessages = append(m.Agent.APIMessages, groq.ChatMessage{
            Role:       "tool",
            ToolCallID: call.ID,
            Content:    result,
        })
    }
}

func (m *Model) executeToolCall(call groq.ToolCall) string {
    switch call.Function.Name {
    case "scratchpad_read":
        return m.executeScratchpadRead(call.Function.Arguments)
    case "scratchpad_write":
        return m.executeScratchpadWrite(call.Function.Arguments)
    case "scratchpad_list":
        return m.executeScratchpadList()
    default:
        // Check for custom mock tools defined in the agent block
        if mock, ok := m.Agent.Config.MockTools[call.Function.Name]; ok {
            return mock.Render(call.Function.Arguments)
        }
        return "Unknown tool: " + call.Function.Name
    }
}
```

---

## Groq API Integration

### Client

```go
// groq/client.go

type Client struct {
    APIKey     string
    HTTPClient *http.Client
    BaseURL    string  // https://api.groq.com/openai/v1
}

type ChatRequest struct {
    Model       string        `json:"model"`
    Messages    []ChatMessage `json:"messages"`
    Tools       []Tool        `json:"tools,omitempty"`
    Temperature float64       `json:"temperature,omitempty"`
    MaxTokens   int           `json:"max_tokens,omitempty"`
}

type ChatMessage struct {
    Role       string     `json:"role"`
    Content    string     `json:"content,omitempty"`
    ToolCalls  []ToolCall `json:"tool_calls,omitempty"`
    ToolCallID string     `json:"tool_call_id,omitempty"`
}

type ChatResponse struct {
    Choices []Choice `json:"choices"`
    Usage   Usage    `json:"usage"`
}

type Usage struct {
    PromptTokens     int `json:"prompt_tokens"`
    CompletionTokens int `json:"completion_tokens"`
    TotalTokens      int `json:"total_tokens"`
}
```

### Configuration

API key set via environment variable `GROQ_API_KEY`. Default model: `llama-3.3-70b-versatile`. The model can be overridden per agent block.

### No Streaming — Gotea Handles This Naturally

The agent's API interaction is just state as far as Gotea is concerned. The flow is:

1. User sends message → handler sets `Loading: true` → re-render shows spinner
2. Handler spawns a goroutine that makes the Groq API call
3. Goroutine completes → sends `AGENT_RESPONSE` message back into the session
4. Response handler sets `Loading: false`, appends the message → re-render shows the response

There's no special async plumbing needed. Gotea's message-driven architecture means the API call result is just another message that triggers a state change and re-render. morphdom diffs the DOM and the chat updates.

Groq is fast enough (1-3s typically) that a loading spinner is perfectly fine for a learning context. Streaming would add complexity for marginal UX benefit here.

---

## Scratchpad

### Per-Cohort Shared Storage

```
data/
└── scratchpads/
    └── {cohort_id}/
        └── {agent_id}/
            ├── notes.md
            ├── solution.go
            └── ...
```

The cohort ID comes from the session or a URL parameter. If no cohort system exists yet, it defaults to a single shared scratchpad per agent ID.

### Why Per-Cohort?

In an instructor-led class, all learners in the same cohort share a scratchpad. This enables:

- The instructor can seed the scratchpad with starter files before the lesson.
- Learners can see each other's work (collaborative exercises).
- The agent can read files that the instructor placed there.

For self-paced learning, each learner effectively gets their own "cohort of one."

### Scratchpad UI

A collapsible panel in the sidebar, below the chat:

```
┌─ Scratchpad ──────────── [▼ collapse] ┐
│                                        │
│  📄 notes.md          (2.1 KB)  [view] │
│  📄 solution.go       (0.8 KB)  [view] │
│  📄 config.yaml       (0.3 KB)  [view] │
│                                        │
└────────────────────────────────────────┘
```

Clicking "view" shows file contents in a modal or expands inline. The learner can also write/edit files directly via the scratchpad UI (sends a `SCRATCHPAD_WRITE` message), not just through the agent.

---

## Parsing

### Block Struct

```go
type AgentBlock struct {
    ID    string `yaml:"id"`
    Title string `yaml:"title"`
    
    // Model configuration
    Model       string  `yaml:"model"`       // default: llama-3.3-70b-versatile
    MaxTurns    int     `yaml:"max_turns"`   // default: 50
    Temperature float64 `yaml:"temperature"` // default: 0.7
    MaxTokens   int     `yaml:"max_tokens"`  // max tokens per response, default: 2048
    
    // System prompt
    System string `yaml:"system"`
    
    // Mode
    Mode  string      `yaml:"mode"`  // "free" (default) or "guided"
    Steps []AgentStep `yaml:"steps"` // Only used in guided mode
    
    // Tools
    Tools     []string             `yaml:"tools"`      // Built-in tool names
    MockTools map[string]MockTool  `yaml:"mock_tools"` // Custom mock tools
    
    // Visibility
    Visibility AgentVisibility `yaml:"visibility"`
    
    // Sidebar config
    Sidebar AgentSidebarConfig `yaml:"sidebar"`
    
    // Post-guided behaviour
    FreeAfterGuided bool `yaml:"free_after_guided"`
}

type AgentStep struct {
    Note     string `yaml:"note"`
    Prompt   string `yaml:"prompt"`
    Editable bool   `yaml:"editable"` // default: false
}

type AgentVisibility struct {
    SystemPrompt string `yaml:"system_prompt"` // visible | hidden | toggleable
    ToolCalls    string `yaml:"tool_calls"`
    FullContext  string `yaml:"full_context"`
    TokenCount   string `yaml:"token_count"`
    Temperature  string `yaml:"temperature"`
    ModelName    string `yaml:"model_name"`
}

type AgentSidebarConfig struct {
    Width     string `yaml:"width"`      // CSS width, default "40%"
    StartOpen bool   `yaml:"start_open"` // default: true
}

type MockTool struct {
    Name         string            `yaml:"name"`
    Description  string            `yaml:"description"`
    Parameters   map[string]Param  `yaml:"parameters"`
    MockResponse string            `yaml:"mock_response"` // supports ${param} interpolation
}
```

### Goldmark Extension

The `agent` fenced block parser extracts the YAML into an `AgentBlock` struct, same as other blocks. The placeholder marker is emitted into the HTML, but the **render function treats it differently**: instead of rendering inline, it sets a flag that activates the sidebar layout.

### Validation at Startup

- Agent ID must be unique within a module.
- If `mode: guided`, must have at least one step.
- `visibility` values must be one of `visible`, `hidden`, `toggleable`.
- Referenced built-in tools must be from the known set.
- At most one agent block per page.

---

## Rendering

### Layout Decision

```go
func (m *Model) renderCurrentPage(s gt.State) []byte {
    model := s.(*Model)
    page := model.currentPage()
    
    // Check if this page has an agent block
    agentBlock := findAgentBlock(page.Blocks)
    
    if agentBlock != nil {
        // Two-column layout
        return renderLayout(model,
            h.Div(a.Attrs(a.Class("split-layout")),
                h.Div(a.Attrs(a.Class("content-pane")),
                    renderPageContent(page, model.ActiveQuiz, model.ActiveHotspot),
                    renderPageNav(model),
                ),
                renderAgentSidebar(model.Agent),
            ),
        ).Bytes()
    }
    
    // Standard single-column layout
    return renderLayout(model,
        renderPageContent(page, model.ActiveQuiz, model.ActiveHotspot),
        renderPageNav(model),
    ).Bytes()
}
```

### Agent Sidebar Render

```go
func renderAgentSidebar(state *AgentState) h.Element {
    if state == nil {
        return h.Nothing()
    }
    
    children := []h.Element{
        renderAgentTitleBar(state),
    }
    
    if !state.SidebarOpen {
        // Collapsed — just show title bar with expand button
        return h.Aside(a.Attrs(a.Class("agent-sidebar collapsed")), children...)
    }
    
    // System prompt panel (if visible or toggled on)
    if shouldShow(state.Config.Visibility.SystemPrompt, state.ShowSystem) {
        children = append(children, renderSystemPromptPanel(state))
    }
    
    // Toggle button for system prompt (if toggleable)
    if state.Config.Visibility.SystemPrompt == "toggleable" {
        children = append(children, renderToggleButton("AGENT_TOGGLE_SYSTEM",
            "System Prompt", state.ShowSystem))
    }
    
    // Guided step note
    if state.Config.Mode == "guided" && !state.GuidedDone {
        step := state.Config.Steps[state.CurrentStep]
        children = append(children, renderStepNote(step, state.CurrentStep,
            len(state.Config.Steps)))
    }
    
    // Chat messages
    children = append(children, renderAgentMessages(state))
    
    // Full context panel (if visible or toggled on)
    if shouldShow(state.Config.Visibility.FullContext, state.ShowFullContext) {
        children = append(children, renderFullContextPanel(state))
    }
    
    // Scratchpad panel
    if hasScatchpadTool(state.Config.Tools) {
        children = append(children, renderScratchpadPanel(state))
    }
    
    // Input area
    children = append(children, renderAgentInput(state))
    
    // Status bar (token count, temperature, model)
    children = append(children, renderAgentStatusBar(state))
    
    return h.Aside(a.Attrs(a.Class("agent-sidebar")), children...)
}
```

### Chat Message Rendering

```go
func renderAgentMessages(state *AgentState) h.Element {
    var msgs []h.Element
    
    for _, m := range state.Messages {
        switch m.Role {
        case "user":
            msgs = append(msgs, renderUserMessage(m))
        case "assistant":
            msgs = append(msgs, renderAssistantMessage(m))
        case "tool_call":
            if shouldShow(state.Config.Visibility.ToolCalls, state.ShowToolCalls) {
                msgs = append(msgs, renderToolCallMessage(m))
            }
        case "tool_result":
            if shouldShow(state.Config.Visibility.ToolCalls, state.ShowToolCalls) {
                msgs = append(msgs, renderToolResultMessage(m))
            }
        }
    }
    
    if state.Loading {
        msgs = append(msgs, renderLoadingIndicator())
    }
    
    return h.Div(a.Attrs(a.Class("agent-messages"), a.Id("agent-messages")),
        msgs...,
    )
}
```

### Guided Mode Input

In guided mode, the input area changes depending on the current step:

```go
func renderAgentInput(state *AgentState) h.Element {
    // If loading, disable everything
    if state.Loading {
        return renderDisabledInput("Thinking...")
    }
    
    // Guided mode — waiting for the user to advance
    if state.Config.Mode == "guided" && !state.GuidedDone && !state.FreeUnlocked {
        step := state.Config.Steps[state.CurrentStep]
        
        if state.StepCompleted {
            // Response received, show "Next Step" button
            if state.CurrentStep < len(state.Config.Steps)-1 {
                return h.Div(a.Attrs(a.Class("agent-input guided")),
                    h.Button(a.Attrs(
                        a.Class("btn-next-step"),
                        a.OnClick(gt.SendBasicMessageNoArgs("AGENT_NEXT_STEP")),
                    ), h.Text("Next Step →")),
                )
            }
            // Last step completed
            if state.Config.FreeAfterGuided {
                return h.Div(a.Attrs(a.Class("agent-input guided")),
                    h.Button(a.Attrs(
                        a.Class("btn-next-step"),
                        a.OnClick(gt.SendBasicMessageNoArgs("AGENT_NEXT_STEP")),
                    ), h.Text("Continue to Free Mode →")),
                )
            }
            return h.Div(a.Attrs(a.Class("agent-input guided-complete")),
                h.P(a.Attrs(), h.Text("✓ All steps completed")),
            )
        }
        
        if step.Prompt == "" {
            // Empty prompt — learner writes their own
            return renderFreeInput(state)
        }
        
        if step.Editable {
            // Editable pre-fill
            return renderEditableStepInput(state, step)
        }
        
        // Read-only pre-fill — just a send button
        return h.Div(a.Attrs(a.Class("agent-input guided")),
            h.Div(a.Attrs(a.Class("prefilled-prompt")),
                h.Text(step.Prompt),
            ),
            h.Button(a.Attrs(
                a.Class("btn-send"),
                a.OnClick(gt.SendBasicMessage("AGENT_SEND_STEP",
                    fmt.Sprintf("%d", state.CurrentStep))),
            ), h.Text("Send")),
        )
    }
    
    // Free mode
    return renderFreeInput(state)
}
```

---

## Scroll-to-Bottom After Render

One client-side JS addition needed: after morphdom patches, scroll the agent messages container to the bottom if new messages were added.

```javascript
// In training.js
document.addEventListener('gotea:render', function() {
    const msgContainer = document.getElementById('agent-messages');
    if (msgContainer) {
        msgContainer.scrollTop = msgContainer.scrollHeight;
    }
});
```

---

## Instructor Follow-Along Mode

When learners are in follow-along mode (following the instructor's navigation), the agent sidebar shows a **read-only mirror** of the instructor's agent session. This is the same broadcast pattern used for page navigation, extended to agent state.

### How It Works

The instructor is the only one interacting with the agent. Their session makes the Groq API calls, processes tool responses, and advances through guided steps. After each state change, `app.Broadcast()` fires, and every learner's render cycle picks up the instructor's agent state.

### Shared State

The instructor's `AgentState` lives on a shared struct (same pattern as `sharedPresentation`):

```go
// Package-level shared state for instructor-led sessions
var sharedPresentation struct {
    CurrentModule string
    CurrentPage   int
    Agent         *AgentState  // Instructor's live agent state
}
```

### Render Decision

```go
func (m *Model) getAgentStateForRender() *AgentState {
    if m.FollowInstructor && sharedPresentation.Agent != nil {
        return sharedPresentation.Agent  // Read-only: show instructor's session
    }
    return m.Agent  // Self-paced: show own session
}
```

### Learner UI Differences in Follow Mode

When following the instructor, the learner sees:

- The full chat history as the instructor builds it (messages appear in real-time via broadcast).
- Tool calls, system prompt panels, etc. — all controlled by the instructor's visibility toggles.
- The guided step notes (so learners can read the pedagogical context).
- **No input field.** The chat is read-only. The input area is replaced with a banner: "Your instructor is leading this demo."
- The loading spinner when the instructor's API call is in-flight.

This means the instructor can narrate as they go: "Watch what happens when I send this vague prompt... now compare that to this more specific one..." — and every learner sees the same conversation building in real-time.

### Instructor Agent Messages

The instructor's agent messages go through the normal `AGENT_SEND` / `AGENT_RESPONSE` flow, but the handlers also update `sharedPresentation.Agent` and call `app.Broadcast()`:

```go
func (m *Model) handleAgentResponse(msg gt.Message, s gt.State) gt.Response {
    model := s.(*Model)
    // ... normal response processing ...
    
    // If this is the instructor session, broadcast the updated state
    if model.IsInstructor {
        sharedPresentation.Agent = model.Agent
        app.Broadcast()
    }
    
    return gt.Respond()
}
```

### Transition to Self-Paced

When the instructor releases control (exits follow-along mode), learners can interact with the agent themselves. At that point, each learner gets a **fresh** `AgentState` initialised from the agent block config on their current page — they don't inherit the instructor's conversation history. The instructor's demo is gone; now they practice on their own.

```
training-app/
├── groq/
│   ├── client.go              # Groq API HTTP client
│   ├── types.go               # Request/response types
│   └── tools.go               # Tool definitions for the API
├── agent/
│   ├── state.go               # AgentState struct and initialisation
│   ├── tools.go               # Tool execution (scratchpad, mock tools)
│   └── scratchpad.go          # Scratchpad filesystem operations
├── agent_handlers.go          # Message handlers for agent interactions
├── agent_renderers.go         # Render functions for the sidebar UI
├── static/
│   └── css/
│       └── agent.css          # Sidebar and chat UI styles
```

---

## Example Training Content

Here's what a lesson on prompt engineering might look like, demonstrating the full feature set:

````markdown
---
title: "Prompt Engineering Fundamentals"
duration: 30m
---

# Prompt Engineering

The way you phrase a request to an AI model dramatically affects the quality
of the response. In this lesson, you'll experiment with different prompting
strategies using a live AI agent.

## Why Prompts Matter

A language model doesn't "understand" your intent — it predicts the most
likely continuation given the text you provide. A vague prompt produces
vague output. A specific, well-structured prompt produces focused, useful
output.

## The System Prompt

Look at the sidebar. You'll notice the agent has a **system prompt** that
you can toggle open. This is the instruction set that shapes the agent's
behaviour before your conversation even starts. As you work through the
guided exercises, pay attention to how the system prompt influences
responses.

## Let's Practice

The agent on the right will walk you through a series of prompts. Each
step builds on the last. Pay attention to how small changes in phrasing
produce very different outputs.

```agent
id: prompt-eng-101
title: "Prompt Practice"
model: llama-3.3-70b-versatile
temperature: 0.7

system: |
  You are a helpful assistant with expertise in many topics.
  Always be thorough and specific in your responses.
  When asked about a topic, provide concrete examples.

mode: guided
steps:
  - note: |
      We're starting with a deliberately vague prompt. Notice how the model
      gives a broad, unfocused response. There's no way for it to know what
      you actually need.
    prompt: "Tell me about dogs."

  - note: |
      Now we've added specificity — a particular breed, a particular context
      (adoption), and a concrete ask (health considerations). Compare this
      response to the previous one.
    prompt: "What are the top 5 health considerations when adopting an adult rescue greyhound?"

  - note: |
      Here we're adding a **persona** (veterinarian) and an **output format**
      (table with specific columns). These constraints dramatically improve
      the usefulness of the output.
    prompt: "As a veterinarian, create a table comparing the top 3 joint supplements for greyhounds, with columns for name, active ingredient, and typical dosage."

  - note: |
      Your turn. Using what you've learned about specificity, personas, and
      format constraints, write a prompt about a topic of your choice.
      Try to be as precise as possible.
    prompt: ""

  - note: |
      Now toggle open the system prompt panel and read it. Then write a prompt
      that *conflicts* with the system instructions. What happens? This
      demonstrates the relationship between system prompts and user prompts.
    prompt: ""
    editable: true

free_after_guided: true

visibility:
  system_prompt: toggleable
  tool_calls: hidden
  full_context: toggleable
  token_count: visible
  temperature: visible

sidebar:
  width: 45%
  start_open: true
```

## Key Takeaways

After completing the guided exercises, reflect on these patterns:

1. **Specificity wins.** Vague inputs produce vague outputs.
2. **Personas focus expertise.** Telling the model *who* to be shapes *what* it says.
3. **Format constraints** (tables, lists, JSON) make outputs immediately useful.
4. **System prompts set the baseline** that user prompts build on.
````

---

## Implementation Phases

This feature is large enough to phase internally:

### Phase A — Core Chat

- `AgentBlock` parsing and validation.
- Groq API client (non-streaming).
- Basic `AgentState` with free-mode chat.
- Sidebar layout and message rendering.
- `AGENT_SEND` / `AGENT_RESPONSE` message flow.

### Phase B — Guided Mode

- Step sequencing with notes.
- Pre-filled prompts (read-only and editable).
- Step navigation (next step, completion detection).
- `free_after_guided` unlock.

### Phase C — Visibility Controls

- System prompt panel (visible/hidden/toggleable).
- Tool call display with styled blocks.
- Full context panel showing raw messages array.
- Token count and temperature display.

### Phase D — Tools & Scratchpad

- Scratchpad read/write/list implementation.
- Scratchpad UI panel in sidebar.
- Tool definitions sent to Groq API.
- Tool call execution loop (call → result → re-call).
- Mock tools with template interpolation.

### Phase E — Instructor Follow-Along & Polish

- Instructor agent broadcast (shared state + `Broadcast()`).
- Read-only learner view in follow mode.
- Scroll-to-bottom JS.
- Sidebar collapse/expand.
- Startup content validation for agent blocks.
- CSS styling and theming for chat UI.
