---
title: "Persistent Memory"
duration: "15m"
tags: [memory, preferences, context, CLAUDE.md]
---

# Persistent Memory

Tell AI about yourself once. It remembers forever.

## The Problem

Every new conversation starts fresh. AI doesn't know:
- Who you are
- What you're working on
- How you like things done
- Your role and context

So you explain it. Again. And again.

## The Solution

**Persistent instruction files** — markdown files that load automatically at the start of every session. Different platforms have different names, but the concept is universal.

Write it once. Every session starts with that context already loaded.

## The Landscape

**CLAUDE.md (Claude Code)**
Markdown files automatically loaded at the start of every Claude Code session. They live in your project root and can import subdirectory files for organisation.

**AGENTS.md (Cursor, Zed, others)**
Proposed by OpenAI in August 2025, AGENTS.md has been adopted by **60,000+ open-source projects**. The standard for AI coding tools.

**Claude Projects (non-coding)**
For non-coding work, Claude Projects offer persistent system prompts with attached knowledge bases. Documents, style guides, and templates persist across conversations.

**ChatGPT Custom Instructions**
User-level persistent instructions. What GPT should know about you and how it should respond.

```callout
type: info
title: "The Highest-Leverage Place to Invest"
content: "Persistent instructions override per-conversation prompts. They shape every interaction. This makes them the highest-leverage place to invest in prompt quality."
```

## What Works: CLAUDE.md Best Practices

**Keep under 300 lines**
Research shows frontier models follow roughly **150-200 instructions** before adherence drops uniformly. More instructions means diluted effectiveness.

**Use progressive disclosure**
Import subdirectory files for different contexts. `#import .claude/coding-style.md` keeps the main file focused.

**Include what's universally applicable**
- Build commands (`npm run build:css`, `./restart.sh`)
- Code style essentials (formatting, naming conventions)
- Key architectural patterns (state ownership, render rules)
- Workflow rules (when to run tests, git commit patterns)

**Commit to version control**
The whole team benefits. New developers onboard faster. Everyone gets the same instructions.

## For Developers: The CLAUDE.md Structure

```markdown
# Project Overview
[2-3 sentence description of what this codebase does]

## Tech Stack
- Framework/language
- Key libraries
- Build tools

## Architecture
[Key patterns — state model, component structure, etc.]

## Development
Build: npm run build
Test: go test ./...
Run: ./restart.sh

## Conventions
[Code style, naming, file organisation]
```

That's it. Short, factual, universally relevant.

## For Non-Developers: Claude Projects

Upload documents AI should always reference:
- Style guides
- Brand voice documents
- Template examples
- Process documentation
- Organisation context

**File format matters.**
Markdown, Word, and PDF use similar knowledge base space. HTML uses **twice as much** due to tag interpretation. Prefer clean formats.

**Organise by function.**
Create focused projects by department or use case. Don't dump everything into one massive knowledge base.

## What Goes in Persistent Memory

**About you**
Your role, responsibilities, expertise area.

**Your preferences**
- Formatting preferences (bullet points vs paragraphs)
- Language preferences (UK vs US English)
- Tone preferences (formal vs casual)

**Your context**
Current projects, team information, company background.

**Standing instructions**
Things you always want done a certain way.

## Example: Personal Assistant Context

```
# About Me
Senior Product Manager at TechCorp.
Working on the mobile app team.
Reports to: VP of Product

# My Preferences
- Bullet points over paragraphs
- UK English spelling
- Concise, direct tone
- Include action items at the end of summaries

# Current Focus
Q1 launch of payment features.
Key stakeholders: Engineering, Design, Finance.
Main risk: Third-party payment provider integration.

# Standing Instructions
Always format dates as DD/MM/YYYY.
When drafting emails, keep under 200 words.
```

## Benefits

**No re-explaining**
"I'm a product manager working on..." is already known.

**Consistent output**
Your preferences are applied automatically.

**Accumulated context**
Add information over time. Context grows.

**Project awareness**
AI knows what you're working on without being told each time.

```callout
type: tip
title: "Build Incrementally"
content: "You don't need to write everything at once. Start with basics. Each time you find yourself explaining something repeatedly, add it to your persistent memory."
```

```quiz
id: persistent-memory-benefit
type: multiple-choice
question: "What problem does persistent memory solve?"
options:
  - "AI running out of context window space"
  - "Having to re-explain who you are and your preferences in every conversation"
  - "AI forgetting what it said earlier in the same conversation"
  - "Slow AI response times"
answer: 1
explanation: "Persistent memory solves the problem of starting every conversation fresh. Your context, preferences, and standing instructions load automatically, so you don't re-explain them every time."
```

```quiz
id: persistent-memory-adoption
type: multiple-choice
question: "Why are persistent instructions the highest-leverage place to invest in prompt quality?"
options:
  - "They make AI run faster"
  - "They override per-conversation prompts and shape every interaction"
  - "They require less maintenance than regular prompts"
  - "They work with all AI models"
answer: 1
explanation: "Persistent instructions override per-conversation prompts, meaning they affect every single interaction. This makes them far more valuable than crafting individual prompts — the investment compounds."
```

```agent
id: persistent-memory-demo
title: "The Power of Persistent Memory"
model_label: "Claude"

system: |
  You are a helpful coding assistant. You help developers implement features
  in their projects. Write clean, working code.

scratchpad:
  "CLAUDE.md": |
    # Project Conventions
    - Go with standard gofmt formatting
    - All handlers use pointer receivers
    - British English in all comments and documentation
    - Table-driven tests required for all new functions
    - Use Tailwind utility classes, no custom CSS
    - Error messages should be user-friendly, not technical
  "feature-request.md": |
    Add a /health endpoint that returns the server status including uptime,
    version, and database connection state.

tools:
  - scratchpad_read
  - scratchpad_write

visibility:
  system_prompt: toggleable
  tool_calls: visible
  full_context: hidden

sidebar:
  width: "45%"
  start_open: true

script:
  - type: note
    text: "Watch how a single configuration file transforms every future AI interaction."

  - type: user
    content: "Add a health check endpoint to the server"

  - type: assistant
    content: "I'll create a health check endpoint for you. Here's a simple implementation:"

  - type: tool_call
    tool: scratchpad_write
    args:
      filename: "health.go"
      content: |
        package main

        import (
            "encoding/json"
            "net/http"
            "time"
        )

        var startTime = time.Now()

        // HealthResponse contains server status information
        type HealthResponse struct {
            Status   string `json:"status"`
            Uptime   string `json:"uptime"`
            Version  string `json:"version"`
            Database string `json:"database"`
        }

        func healthHandler(w http.ResponseWriter, r *http.Request) {
            uptime := time.Since(startTime).String()

            // Check if database is connected
            dbStatus := "connected"
            if err := db.Ping(); err != nil {
                dbStatus = "disconnected"
            }

            response := HealthResponse{
                Status:   "ok",
                Uptime:   uptime,
                Version:  "1.0.0",
                Database: dbStatus,
            }

            w.Header().Set("Content-Type", "application/json")
            json.NewEncoder(w).Encode(response)
        }

  - type: tool_result
    tool: scratchpad_write
    content: "Written to health.go (36 lines)"

  - type: assistant
    content: "Done! I've created a /health endpoint that returns server status, uptime, version, and database connection state."

  - type: note
    text: "The output works, but it doesn't match our project conventions. Notice the American English ('connected'), the value receiver pattern, and the absence of tests. Now let's load our persistent instructions..."

  - type: user
    content: "I've set up our project conventions. Please read the CLAUDE.md file and then implement the same health endpoint."

  - type: tool_call
    tool: scratchpad_read
    args:
      filename: "CLAUDE.md"

  - type: tool_result
    tool: scratchpad_read
    content: |
      # Project Conventions
      - Go with standard gofmt formatting
      - All handlers use pointer receivers
      - British English in all comments and documentation
      - Table-driven tests required for all new functions
      - Use Tailwind utility classes, no custom CSS
      - Error messages should be user-friendly, not technical

  - type: tool_call
    tool: scratchpad_read
    args:
      filename: "feature-request.md"

  - type: tool_result
    tool: scratchpad_read
    content: |
      Add a /health endpoint that returns the server status including uptime,
      version, and database connection state.

  - type: assistant
    content: "Perfect, I can see your project conventions. Let me implement the health endpoint following your standards — British English, pointer receivers, and with comprehensive tests."

  - type: tool_call
    tool: scratchpad_write
    args:
      filename: "health.go"
      content: |
        package main

        import (
            "encoding/json"
            "net/http"
            "time"
        )

        var startTime = time.Now()

        // HealthResponse contains server status information
        type HealthResponse struct {
            Status   string `json:"status"`
            Uptime   string `json:"uptime"`
            Version  string `json:"version"`
            Database string `json:"database"`
        }

        // HealthHandler returns the current server health status
        func (m *Model) HealthHandler(w http.ResponseWriter, r *http.Request) {
            uptime := time.Since(startTime).String()

            // Check whether the database is connected
            dbStatus := "connected"
            if err := db.Ping(); err != nil {
                dbStatus = "disconnected"
                // User-friendly error logged, but don't expose technical details
            }

            response := HealthResponse{
                Status:   "ok",
                Uptime:   uptime,
                Version:  "1.0.0",
                Database: dbStatus,
            }

            w.Header().Set("Content-Type", "application/json")
            json.NewEncoder(w).Encode(response)
        }

  - type: tool_result
    tool: scratchpad_write
    content: "Written to health.go (39 lines)"

  - type: tool_call
    tool: scratchpad_write
    args:
      filename: "health_test.go"
      content: |
        package main

        import (
            "net/http"
            "net/http/httptest"
            "testing"
        )

        func TestHealthHandler(t *testing.T) {
            tests := []struct {
                name           string
                dbConnected    bool
                expectedStatus int
                checkDB        bool
            }{
                {
                    name:           "healthy server with database",
                    dbConnected:    true,
                    expectedStatus: http.StatusOK,
                    checkDB:        true,
                },
                {
                    name:           "healthy server without database",
                    dbConnected:    false,
                    expectedStatus: http.StatusOK,
                    checkDB:        true,
                },
            }

            for _, tt := range tests {
                t.Run(tt.name, func(t *testing.T) {
                    req := httptest.NewRequest("GET", "/health", nil)
                    w := httptest.NewRecorder()

                    model := &Model{}
                    model.HealthHandler(w, req)

                    if w.Code != tt.expectedStatus {
                        t.Errorf("expected status %d, got %d", tt.expectedStatus, w.Code)
                    }

                    if contentType := w.Header().Get("Content-Type"); contentType != "application/json" {
                        t.Errorf("expected Content-Type application/json, got %s", contentType)
                    }
                })
            }
        }

  - type: tool_result
    tool: scratchpad_write
    content: "Written to health_test.go (46 lines)"

  - type: assistant
    content: "Implemented! I've created the health endpoint following your conventions: pointer receiver on the handler, British English in comments ('whether' instead of 'if'), user-friendly error handling, and comprehensive table-driven tests covering both database states."

  - type: note
    text: "Same task, dramatically different output. The CLAUDE.md file ensures consistent, convention-following code in every conversation. Write once, benefit forever."
```
