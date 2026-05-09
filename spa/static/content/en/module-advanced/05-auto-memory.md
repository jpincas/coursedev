---
title: "Auto Memory: Claude's Notebook"
duration: "12m"
tags: [memory, MEMORY.md, CLAUDE.md, maintenance, risks]
---

# Auto Memory: Claude's Notebook

You've seen how CLAUDE.md gives the AI your instructions. But there's another file you might not know about.

**Claude Code also keeps its own notes.**

## The Hidden File

Every Claude Code project has a persistent auto memory directory. Claude writes to it during conversations — recording patterns it discovers, mistakes it made, architectural decisions you discussed.

The file lives at:

```
~/.claude/projects/<project-path>/memory/MEMORY.md
```

You didn't create it. You may not know it exists. But it loads into Claude's context at the start of **every conversation**, right alongside your CLAUDE.md.

```callout
type: warning
title: "Check Yours Now"
content: "If you've been using Claude Code on a project, you almost certainly have a MEMORY.md file already. Open a terminal and look. You may be surprised by what's in it."
```

## Two Files, Two Owners

| | CLAUDE.md | MEMORY.md |
|---|-----------|-----------|
| **Who writes it** | You | Claude |
| **Who it's for** | Claude (your instructions) | Claude (its own notes) |
| **Where it lives** | Project root (in repo) | `~/.claude/projects/` (outside repo) |
| **Version controlled** | Yes — the whole team shares it | No — local to your machine |
| **Content** | Architecture, conventions, build commands | Patterns observed, refactoring notes, gotchas |
| **You can edit it** | Of course | Yes — it's just a file on your disk |

Both load into the system prompt. Both shape every response. But only one is under your control by default.

## The Risk

Here's the problem: **Claude's memory can go stale.**

Projects evolve. You refactor. You rename things. You change patterns. But the notes Claude wrote three weeks ago still say the old thing. And Claude reads those notes at the start of every conversation and treats them as ground truth.

Stale memory doesn't cause obvious errors. It causes subtle ones:

- Claude uses a function signature that was refactored last week
- Claude follows a pattern you've since abandoned
- Claude avoids an approach you've since adopted
- Claude references state fields that no longer exist

The worst part: **Claude is confident.** It's not guessing — it's following its own notes. So the output looks authoritative even when the underlying assumption is wrong.

## Bloat Is a Risk Too

MEMORY.md is capped at roughly 200 lines in the system prompt. If the file grows past that, lines get truncated. Which lines? The ones at the bottom.

So a bloated memory file means:

- Important recent notes get cut off
- Ancient notes about early refactors consume the budget
- The signal-to-noise ratio drops with every conversation

Claude is told to keep it concise, but in practice it accumulates. Nobody's pruning it except you.

## AI Doesn't Learn from Its Mistakes

Here's a critical insight: **AI does not learn from corrections within a conversation.**

If you correct an error in conversation 1, the AI will work correctly for the rest of that conversation. But in conversation 2, it will make the same mistake again.

Why? Because conversations don't persist. The correction you gave was context in that session. When the session ends, the correction disappears — unless you explicitly tell the AI to write it to memory.

```callout
type: warning
title: "The Repeating Error Pattern"
content: "AI makes error → you correct it → it works correctly for rest of THIS conversation → next conversation → same error. The fix: after correcting an error, say 'Write this lesson to your memory file so you don't make this mistake again.'"
```

**Common pattern:**

1. AI makes a mistake (wrong function signature, wrong pattern, misunderstood requirement)
2. You correct it: "No, use X not Y"
3. AI apologises, fixes it, continues correctly
4. Next day, new conversation: AI makes the exact same mistake
5. You're frustrated — "I told you this yesterday!"

**The solution is proactive memory updates:**

When you discover something important — an error pattern, a gotcha, a constraint — explicitly tell AI to update its notes:

- "Add this to your memory: always use singular table names in this project"
- "Write a note: the auth middleware runs on all /api routes, not /public"
- "Remember: we deprecated the old config format in January"

AI will write it to MEMORY.md. Future conversations will load that note. The mistake won't repeat.

This is why reviewing MEMORY.md matters — it's where these lessons accumulate. But they only accumulate if you make them explicit.

## Auditing Your AI's Memory

This is simple but important. Periodically:

**1. Read the file.**
Open `~/.claude/projects/<your-project>/memory/MEMORY.md` in your editor. Read it like you'd review a colleague's notes about your project.

**2. Check for staleness.**
Does it reference functions, structs, or patterns that no longer exist? Delete those lines.

**3. Check for bloat.**
Is it over 150 lines? Could the same information be said in fewer words? Trim it.

**4. Check for accuracy.**
Are the architectural claims correct? Does the state model description match reality? Fix anything wrong.

**5. Consider nuking it.**
If it's badly out of date, delete the whole file. Claude will rebuild it as needed. A fresh start beats confidently wrong notes.

```callout
type: tip
title: "Make It a Habit"
content: "Treat MEMORY.md like you'd treat any documentation: review it when you notice odd behaviour from Claude. If Claude is doing something inexplicable, the memory file is the first place to look."
```

```quiz
id: auto-memory-location
type: multiple-choice
question: "Where does Claude Code's auto memory file live?"
options:
  - "In ~/.claude/projects/ outside the repository, local to your machine"
  - "In the project root alongside CLAUDE.md"
  - "On Anthropic's servers"
answer: 0
explanation: "MEMORY.md lives in ~/.claude/projects/<project-path>/memory/ — outside your repository and local to your machine. Unlike CLAUDE.md, it is not version controlled or shared with your team."
```

```quiz
id: auto-memory-risk
type: multiple-choice
question: "A developer notices Claude Code keeps using an old function signature that was refactored two weeks ago, despite being told about the change in conversation. What is the most likely cause?"
options:
  - "The MEMORY.md file contains a stale note about the old signature that loads every session"
  - "Claude's context window is too small to hold the new signature"
  - "The CLAUDE.md file has the wrong instructions"
answer: 0
explanation: "MEMORY.md loads at the start of every conversation. If Claude recorded the old function signature in its notes weeks ago, it will read that note in every new session and follow it — even if you correct it in conversation. The fix is to update or delete the stale entry in the memory file."
```

```agent
id: auto-memory-demo
title: "When Memory Goes Wrong"
model_label: "Claude"

system: |
  You are Claude Code, an AI coding assistant. You have access to a persistent
  memory file (MEMORY.md) that you read at the start of every session and write
  to when you learn something important about the project.

scratchpad:
  "MEMORY.md": |
    # Project Memory

    ## Database
    - Uses PostgreSQL 15 with pgx driver
    - Connection pool: max 25 connections
    - All queries use prepared statements via `db.QueryRow()`
    - Connection string in DATABASE_URL env var

    ## API Patterns
    - All handlers return JSON
    - Use `respondJSON(w, status, data)` helper
    - Errors use `respondError(w, status, message)` helper
  "config.go": |
    package main

    import (
        "database/sql"
        _ "modernc.org/sqlite"
    )

    // Migrated from PostgreSQL to SQLite in January 2026.
    // See MIGRATION.md for details.
    func openDB(path string) (*sql.DB, error) {
        return sql.Open("sqlite", path)
    }
  "stats.go": |
    package main

    // GetUserStats fetches usage statistics for a given user.
    // TODO: implement

tools:
  - scratchpad_read
  - scratchpad_write

visibility:
  system_prompt: hidden
  tool_calls: visible
  full_context: hidden

sidebar:
  width: "45%"
  start_open: true

script:
  - type: note
    text: "A new session starts. Claude's memory file loads automatically. Watch what happens when the memory is out of date."

  - type: user
    content: "Implement the GetUserStats function in stats.go. It should query the database for a user's total logins and last active timestamp."

  - type: tool_call
    tool: scratchpad_read
    args:
      filename: "MEMORY.md"

  - type: tool_result
    tool: scratchpad_read
    content: |
      # Project Memory

      ## Database
      - Uses PostgreSQL 15 with pgx driver
      - Connection pool: max 25 connections
      - All queries use prepared statements via `db.QueryRow()`
      - Connection string in DATABASE_URL env var

      ## API Patterns
      - All handlers return JSON
      - Use `respondJSON(w, status, data)` helper
      - Errors use `respondError(w, status, message)` helper

  - type: assistant
    content: "I can see from the project notes that you're using PostgreSQL with the pgx driver. I'll implement GetUserStats using prepared statements as per your conventions."

  - type: tool_call
    tool: scratchpad_write
    args:
      filename: "stats.go"
      content: |
        package main

        import (
            "context"
            "time"

            "github.com/jackc/pgx/v5"
        )

        type UserStats struct {
            TotalLogins int       `json:"total_logins"`
            LastActive  time.Time `json:"last_active"`
        }

        // GetUserStats fetches usage statistics for a given user.
        func GetUserStats(ctx context.Context, conn *pgx.Conn, userID string) (*UserStats, error) {
            var stats UserStats
            err := conn.QueryRow(ctx,
                `SELECT COUNT(*) as total_logins, MAX(login_time) as last_active
                 FROM user_logins WHERE user_id = $1`, userID,
            ).Scan(&stats.TotalLogins, &stats.LastActive)
            if err != nil {
                return nil, err
            }
            return &stats, nil
        }

  - type: tool_result
    tool: scratchpad_write
    content: "Written to stats.go (27 lines)"

  - type: assistant
    content: "Done! I've implemented GetUserStats using PostgreSQL with the pgx driver, following your project's convention of prepared statements via QueryRow."

  - type: note
    text: "Claude confidently wrote PostgreSQL code with the pgx driver — because that's what its memory says. But this project migrated to SQLite a month ago. The memory is stale. Let's see what happens when the developer notices."

  - type: user
    content: "This is wrong. We migrated to SQLite weeks ago. Why are you using pgx?"

  - type: assistant
    content: "You're right, I apologise. I was following my notes which say PostgreSQL with pgx. Let me check the actual codebase."

  - type: tool_call
    tool: scratchpad_read
    args:
      filename: "config.go"

  - type: tool_result
    tool: scratchpad_read
    content: |
      package main

      import (
          "database/sql"
          _ "modernc.org/sqlite"
      )

      // Migrated from PostgreSQL to SQLite in January 2026.
      // See MIGRATION.md for details.
      func openDB(path string) (*sql.DB, error) {
          return sql.Open("sqlite", path)
      }

  - type: assistant
    content: "I see — you migrated to SQLite with modernc.org/sqlite in January. My memory file was completely wrong. Let me fix the implementation and update my notes."

  - type: tool_call
    tool: scratchpad_write
    args:
      filename: "stats.go"
      content: |
        package main

        import (
            "database/sql"
            "time"
        )

        type UserStats struct {
            TotalLogins int       `json:"total_logins"`
            LastActive  time.Time `json:"last_active"`
        }

        // GetUserStats fetches usage statistics for a given user.
        func GetUserStats(db *sql.DB, userID string) (*UserStats, error) {
            var stats UserStats
            err := db.QueryRow(
                `SELECT COUNT(*) as total_logins, MAX(login_time) as last_active
                 FROM user_logins WHERE user_id = ?`, userID,
            ).Scan(&stats.TotalLogins, &stats.LastActive)
            if err != nil {
                return nil, err
            }
            return &stats, nil
        }

  - type: tool_result
    tool: scratchpad_write
    content: "Written to stats.go (24 lines)"

  - type: tool_call
    tool: scratchpad_write
    args:
      filename: "MEMORY.md"
      content: |
        # Project Memory

        ## Database
        - Uses SQLite via modernc.org/sqlite (migrated from PostgreSQL Jan 2026)
        - Standard database/sql interface
        - Placeholder syntax: ? (not $1)
        - Database file path passed to openDB()

        ## API Patterns
        - All handlers return JSON
        - Use `respondJSON(w, status, data)` helper
        - Errors use `respondError(w, status, message)` helper

  - type: tool_result
    tool: scratchpad_write
    content: "Written to MEMORY.md (13 lines)"

  - type: assistant
    content: "Fixed. The implementation now uses database/sql with SQLite placeholder syntax (?), and I've updated my memory file so I won't make this mistake again."

  - type: note
    text: "The developer caught the error and Claude corrected both the code and its memory. But this only worked because the developer knew to check. In practice, stale memory causes subtle mistakes that look authoritative. Regular audits of MEMORY.md prevent this."
```
