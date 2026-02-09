---
name: agent-demo
description: Create agent demo blocks — scripted AI conversation walkthroughs embedded in lesson markdown pages. Use when authoring new agent demos or editing existing ones.
argument-hint: "[page path or description of the demo to create]"
---

# Agent Demo Block Author

You create **agent demo blocks** — scripted, interactive AI conversation walkthroughs embedded in training course lesson pages. These are NOT live AI chats. They are carefully choreographed demonstrations where every message, tool call, and response is pre-written to teach a specific concept.

## Before Starting

1. **Read the target page** if a path is given. Understand the page's narrative, what it teaches, and where the agent block fits.
2. **Read surrounding pages** in the same module to understand the pedagogical arc.
3. **Read the module's `_module.yaml`** for context on the module's goals.

## What You Produce

A complete ` ```agent ` fenced code block in valid YAML, ready to paste into a markdown lesson page. The block contains:
- Metadata (id, title, model label, system prompt)
- Initial files (the simulated filesystem)
- Visibility and sidebar configuration
- A script of events (the choreographed conversation)

## Complete YAML Schema

```yaml
id: <string>           # REQUIRED. Unique kebab-case ID scoped to module (e.g., "context-window-demo")
title: <string>        # REQUIRED. Display title in the agent sidebar header
model_label: <string>  # Optional. Cosmetic model name (e.g., "Claude", "GPT-4")
system: |              # Optional. System prompt for the simulated AI
  The AI's instructions...

scratchpad:            # Optional. Initial files in the simulated filesystem (YAML key is "scratchpad" but students see it as a file explorer)
  "filename.txt": |    # Keys are filenames (quoted), values are file content
    File content here...
  "data/report.csv": | # Paths with "/" create folder structure in the explorer
    Header,Value
    Row1,Data

tools:                 # Optional. List of tool names the AI "has access to"
  - scratchpad_read    # Read a file
  - scratchpad_write   # Create or overwrite a file
  - list_files         # List files in a directory
  - move_file          # Move/rename a file
  - create_folder      # Create a new folder
  - delete_file        # Delete a file

visibility:            # Optional. Controls what learners can see
  system_prompt: hidden    # visible | hidden | toggleable (default: hidden)
  tool_calls: hidden       # visible | hidden (default: hidden)
  full_context: hidden     # visible | hidden | toggleable (default: hidden)
  token_count: hidden      # visible | hidden (default: hidden)
  model_name: hidden       # visible | hidden (default: hidden)

sidebar:               # Optional. Layout config
  width: "40%"         # CSS width (default: "40%")
  start_open: true     # Whether sidebar starts expanded (default: true)

script:                # REQUIRED. Non-empty array of events
  - type: <event_type>
    ...
```

## Script Event Types

### `note` — Instructor Commentary

Displayed as an amber callout. Not part of the simulated conversation. The course creator's voice speaking directly to the learner. Used to set context before a conversation starts, draw attention to something that just happened, or provide a teaching moment.

```yaml
- type: note
  text: "Watch how the AI reads every file before synthesising a response."
```

### `user` — User Message

A message from the "user" in the simulated conversation. **Starts a new event group** — all subsequent `tool_call`, `tool_result`, and `assistant` events auto-advance with staggered animation until a boundary (next `note`, `user`, `compaction`, `clear`, or end of script).

```yaml
- type: user
  content: "Analyse this data and write a summary report."
```

### `assistant` — AI Response

The AI's response text. Part of the current event group.

```yaml
- type: assistant
  content: "I'll review the materials and create a structured report."
  tokens: 45    # Optional explicit token count (otherwise estimated at ~4 chars/token)
```

### `tool_call` — Tool Invocation

The AI calling a tool. Part of the current event group.

```yaml
- type: tool_call
  tool: scratchpad_read
  args:
    filename: "data.csv"
```

For `scratchpad_write`, the `args` MUST include `filename` and `content`:
```yaml
- type: tool_call
  tool: scratchpad_write
  args:
    filename: "report.md"
    content: |
      # Report
      Content written by the AI...
```

### `tool_result` — Tool Output

The result of a tool call. Part of the current event group. For `scratchpad_write`, the preceding `tool_call`'s args are used to actually create/update the file in the filesystem.

```yaml
- type: tool_result
  tool: scratchpad_read
  content: |
    The file content that was read...
```

```yaml
- type: tool_result
  tool: scratchpad_write
  content: "Written to report.md (18 lines)"
```

### `compaction` — Context Compaction

Replaces all prior context messages (except system prompt) with a summary. Teaches token management. The advance button text changes to "Compact context".

```yaml
- type: compaction
  summary: |
    [Earlier: user asked about project budget. Agent read 3 files
     and created a preliminary summary.]
```

### `clear` — Reset Conversation

Clears chat messages and context. Optionally resets files to initial state. Button text changes to "Clear & reset".

```yaml
- type: clear
  reset_scratchpad: true   # Optional, default false
  note: "Starting fresh for the next scenario."  # Optional
```

## Validation Rules

These rules are enforced at parse time. Violations cause a build error:

1. `script` must not be empty
2. `script` must start with a `note` or `user` event
3. Every `tool_result` must be preceded by a `tool_call` with the same tool name
4. `scratchpad_write` tool_calls must have `filename` and `content` in args
4b. `move_file` tool_calls must have `source` and `destination` in args
4c. `create_folder` tool_calls must have `path` in args
4d. `delete_file` tool_calls must have `filename` in args
5. `compaction.summary` must not be empty
6. At most one agent block per page
7. Agent ID must be unique within the module

## Event Grouping Model

This is critical to understand for pacing:

- A `user` event starts a **group**. Everything after it auto-advances with staggered animation delays until a boundary.
- Boundaries that stop auto-advance: `note`, `user`, `compaction`, `clear`, end of script.
- Within a group, the learner sees messages appear one after another with a brief delay — it feels like watching a real conversation unfold.
- `note` events are always their own step. The learner clicks "Next" to move past them.
- `compaction` and `clear` are their own steps too.

**Implication for script design:** After a `user` message, all the tool calls, tool results, and the final assistant response should flow as one group. If you want to pause and add commentary mid-conversation, insert a `note` — that breaks the group and forces the learner to click again.

## Design Principles

### 1. Files Should Feel Real

Don't use placeholder content. Create realistic files that look like actual project artifacts — with imperfections, varying formats, realistic data. The filesystem IS the demo's stage set. If the files feel fake, the whole demo feels fake.

Good: A CSV with inconsistent formatting and a few blank cells
Bad: A perfectly formatted CSV with generic "Item 1", "Item 2" data

### 2. Notes Are the Teaching Voice

Notes should be conversational, direct, and purposeful. They:
- Set up what the learner is about to see ("Watch how...")
- Point out what just happened ("Notice that the AI didn't...")
- Connect to the broader lesson ("This is delegation, not prompting")

Don't over-narrate. 2-4 notes per demo is typical. Let the conversation speak for itself.

### 3. One Demo, One Teaching Point

Each agent demo should clearly demonstrate ONE concept. Don't try to show everything at once. Common patterns:
- **Delegation demo:** User gives a high-level task, AI figures out the steps
- **Tool use demo:** Focus on how the AI reads, processes, and creates files
- **Context window demo:** Use `compaction` to show token management
- **Iteration demo:** User gives feedback, AI refines its output
- **Error recovery demo:** AI makes a mistake, user corrects, AI adapts

### 4. User Messages Should Model Good Practice

The user messages ARE the teaching. If the demo is about delegation, the user message should demonstrate good delegation. If it's about specificity, the user message should be specific. The learner will unconsciously adopt the patterns they see.

### 5. AI Responses Should Be Realistic

Don't write an impossibly perfect AI. Write responses that:
- Acknowledge what they're about to do before doing it
- Show a realistic tool-call pattern (read before write, read multiple files)
- Produce output that's good but not superhuman
- Match the claimed model's actual capabilities and tone

### 6. Visibility Settings Match the Lesson

- **Opening demo (delegation focus):** tool_calls visible, system_prompt toggleable
- **How LLMs work lesson:** full_context toggleable, token_count visible
- **Prompting lesson:** system_prompt visible
- **Advanced tooling lesson:** tool_calls visible, full_context visible

### 7. Keep It Tight

Demos should take 2-5 minutes to step through. More than ~20 script events gets tedious. If you need to show a longer interaction, use `compaction` or `clear` to create act breaks.

## Workflow

When the user asks you to create an agent demo:

1. **Understand the context:** Read the page, surrounding pages, and module metadata.
2. **Identify the teaching point:** What ONE concept should this demo illustrate?
3. **Design the files:** What files does the AI need to work with? Make them realistic.
4. **Write the system prompt:** Match the demo's scenario. Keep it focused.
5. **Choreograph the script:**
   - Open with 1-2 notes to set context
   - Write the user message (this models the skill being taught)
   - Write the AI's work sequence (tool calls, responses)
   - Close with 1-2 notes connecting back to the lesson
6. **Set visibility:** Match the lesson's pedagogical goals.
7. **Output the complete block** as a fenced code block ready to paste.

## Complete Reference Example

Below is a full, self-contained agent demo block. Study the patterns: realistic project files in a folder, opening notes that set context, a single user instruction that models delegation, a realistic AI work sequence (list files, read all, then synthesise), and closing notes that connect back to the lesson.

````yaml
```agent
id: delegation-demo
title: "From Chaos to Report"
model_label: "Claude"

system: |
  You are a project management assistant. You help professionals create clear,
  actionable reports from project materials. Be thorough but concise. Write in
  a professional tone suitable for executive audiences.

scratchpad:
  "project/meeting-notes.txt": |
    Project Phoenix - Team Standup Jan 15
    Attendees: Sarah (PM), Dev team, Lisa (Design)
    - API migration 60% complete, targeting Feb 1
    - Design review blocked on brand guidelines
    - Client wants mobile dashboard added to Phase 1
    - Contractor rates up 15% — budget concern
    - Sarah to escalate timeline risk to leadership
    - Next milestone review: Jan 22
  "project/budget-q1.csv": |
    Category,Planned,Actual,Variance
    Development,45000,38000,-7000
    Design,15000,12500,-2500
    Infrastructure,8000,9200,+1200
    Contractors,20000,27500,+7500
    Testing,7000,4200,-2800
    Total,95000,91400,-3600
  "project/client-emails.txt": |
    From: David Chen <d.chen@acmecorp.com>
    Date: Jan 14
    Subject: RE: Project Phoenix Timeline

    Sarah — The board wants mobile dashboard in Phase 1.
    Revised timeline? We need this by March 15 latest.
    ---
    From: Sarah Lopez <s.lopez@team.com>
    Date: Jan 12

    David — Assessing mobile addition now. Adds ~3 weeks.
    Revised plan by end of next week.

tools:
  - list_files
  - scratchpad_read
  - scratchpad_write

visibility:
  system_prompt: toggleable
  tool_calls: visible
  full_context: hidden
  token_count: hidden
  model_name: visible

sidebar:
  width: "45%"

script:
  - type: note
    text: "This is a live walkthrough of AI doing real work. The file explorer shows three project files — meeting notes, budget data, and client emails. Messy, disorganised, typical."

  - type: note
    text: "Now watch. One instruction. No templates, no step-by-step breakdown. Just describe the outcome you need."

  - type: user
    content: "I've got a project folder with meeting notes, budget data, and client emails. I need an executive status report for leadership by end of day. Pull together the key points, flag the risks, and make it professional."

  - type: tool_call
    tool: list_files
    args:
      path: "project"

  - type: tool_result
    tool: list_files
    content: |
      3 files in project/:
      - project/budget-q1.csv
      - project/client-emails.txt
      - project/meeting-notes.txt

  - type: assistant
    content: "I'll review all the project materials and create a professional executive status report. Let me start by reading through each file."

  - type: tool_call
    tool: scratchpad_read
    args:
      filename: "project/meeting-notes.txt"

  - type: tool_result
    tool: scratchpad_read
    content: |
      Project Phoenix - Team Standup Jan 15
      Attendees: Sarah (PM), Dev team, Lisa (Design)
      - API migration 60% complete, targeting Feb 1
      - Design review blocked on brand guidelines
      - Client wants mobile dashboard added to Phase 1
      - Contractor rates up 15% — budget concern
      - Sarah to escalate timeline risk to leadership
      - Next milestone review: Jan 22

  - type: tool_call
    tool: scratchpad_read
    args:
      filename: "project/budget-q1.csv"

  - type: tool_result
    tool: scratchpad_read
    content: |
      Category,Planned,Actual,Variance
      Development,45000,38000,-7000
      Design,15000,12500,-2500
      Infrastructure,8000,9200,+1200
      Contractors,20000,27500,+7500
      Testing,7000,4200,-2800
      Total,95000,91400,-3600

  - type: tool_call
    tool: scratchpad_read
    args:
      filename: "project/client-emails.txt"

  - type: tool_result
    tool: scratchpad_read
    content: |
      From: David Chen <d.chen@acmecorp.com>
      Date: Jan 14
      Subject: RE: Project Phoenix Timeline

      Sarah — The board wants mobile dashboard in Phase 1.
      Revised timeline? We need this by March 15 latest.
      ---
      From: Sarah Lopez <s.lopez@team.com>
      Date: Jan 12

      David — Assessing mobile addition now. Adds ~3 weeks.
      Revised plan by end of next week.

  - type: assistant
    content: "I've reviewed all three documents. Clear themes: a scope change risk from the mobile dashboard request, a design blockage, and contractor cost pressure. Let me synthesise this into a structured executive report."

  - type: tool_call
    tool: scratchpad_write
    args:
      filename: "exec-status-report.md"
      content: |
        # Project Phoenix — Executive Status Report
        Prepared: January 15, 2026

        ## Summary
        Project Phoenix is 60% through API migration but faces timeline risk
        from a client-requested scope change (mobile dashboard). Current spend
        is under budget but contractor overruns signal pressure.

        ## Key Metrics
        - Budget: 91,400 of 95,000 spent (3.8% under)
        - Timeline: Original launch Mar 1 — AT RISK
        - Scope: Mobile dashboard added to Phase 1, not yet planned

        ## Critical Risks
        1. SCOPE CHANGE: Mobile dashboard adds ~3 weeks. Client deadline
           March 15. Recommend presenting revised timeline with phased options.
        2. DESIGN BLOCKAGE: Blocked on brand guidelines. Creates downstream
           delays to frontend and testing. Escalate to client this week.
        3. CONTRACTOR COSTS: 7,500 over plan (37.5% variance). Will accelerate
           with expanded scope. Review rates and consider fixed-price.

        ## Recommended Next Steps
        1. Present revised timeline to client by Jan 22
        2. Escalate brand guidelines blocker
        3. Review contractor cost structure

  - type: tool_result
    tool: scratchpad_write
    content: "Written to exec-status-report.md (26 lines)"

  - type: assistant
    content: "Done. I've created exec-status-report.md — a structured executive status report covering project summary, key metrics, three critical risks with recommendations, and prioritised next steps. Ready for leadership review."

  - type: note
    text: "Three scattered files became one professional report. No copy-paste. No back-and-forth. No templates. The AI read every file, identified the themes, and delivered a finished document."

  - type: note
    text: "Notice what you didn't do: you didn't specify which files to read first, how to structure the report, or what sections to include. You described the outcome. AI figured out the rest. That's delegation, not prompting."
```
````

Key patterns to notice in this example:
- **Files in a `project/` folder** — realistic directory structure, not scattered at root
- **`list_files` before reading** — the AI discovers the directory contents, then reads each file
- **Files have realistic imperfections** — inconsistent formats, real-looking data, email thread with dates
- **Opening notes** set up what the learner should pay attention to (2 notes)
- **Single user message** models the skill being taught (delegation = describe outcome, not steps)
- **AI reads ALL files before writing** — realistic tool-call pattern, not shortcutting
- **tool_result content matches file content** — the read results should echo what's in the files
- **Closing notes** connect back to the lesson's teaching point (2 notes)
- **Visibility** matches a delegation-focused lesson: tool_calls visible (see the work), system_prompt toggleable (can peek), full_context hidden (not the focus)

## File System Tools

Six tools have **real side effects** — they actually mutate the filesystem during playback. Note: tool names use the internal `scratchpad_` prefix in YAML but display as `read_file`/`write_file` to students:

### `scratchpad_read` — Read a file
```yaml
- type: tool_call
  tool: scratchpad_read
  args:
    filename: "data/report.csv"    # Supports folder paths
```
No side effect (read-only). The `tool_result` content should echo the file content.

### `scratchpad_write` — Create or overwrite a file
```yaml
- type: tool_call
  tool: scratchpad_write
  args:
    filename: "reports/summary.md"  # Creates file; parent folders auto-expand in explorer
    content: |
      # Summary
      Content here...
```
**Side effect:** Creates/overwrites the file in the scratchpad. Parent folders auto-expand.

### `list_files` — List directory contents
```yaml
- type: tool_call
  tool: list_files
  args:
    path: "data"                    # Or "." for root
```
No side effect. The `tool_result` content should list the files (you write it manually).

### `move_file` — Move or rename a file
```yaml
- type: tool_call
  tool: move_file
  args:
    source: "old-report.docx"
    destination: "archive/old-report.docx"
```
**Side effect:** Removes the source file, creates the destination. Parent folders auto-expand.

### `create_folder` — Create a folder
```yaml
- type: tool_call
  tool: create_folder
  args:
    path: "reports"
```
**Side effect:** Registers the folder as expanded in the explorer. Parent folders also expand.

### `delete_file` — Delete a file
```yaml
- type: tool_call
  tool: delete_file
  args:
    filename: "old-draft.docx"
```
**Side effect:** Removes the file. If it was being previewed, the preview clears.

## Folder Paths in Scratchpad

File keys can contain `/` to represent folder structure. The workspace explorer renders these as a tree with expandable folders:

```yaml
scratchpad:
  "data/sales.csv": |
    Product,Revenue
    Widget,12000
  "data/support.csv": |
    Ticket,Status
    T-001,Open
  "reports/summary.md": |
    # Q4 Summary
    ...
```

This creates two folders (`data/`, `reports/`) with files inside them. All folders from the initial scratchpad are auto-expanded on load. Folders created during playback (via `create_folder` or `scratchpad_write` to a new path) also auto-expand.

**When to use folders:** Use folder structure when the demo involves organising files, working with project directories, or when realistic file layout adds to the teaching point. Keep files flat when folder structure would just add noise.

## Cosmetic Tool Types

Beyond the six tools above, you can use any tool name. These appear in the UI with generic icons but have no server-side behavior:

- `web_search` — simulated web search (args: `query`)
- `fetch_url` — simulated URL fetch (args: `url`)
- Any domain-specific tool name that makes sense for the scenario
