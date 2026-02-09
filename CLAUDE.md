# Claude.md

## Project Overview

This is a server-side rendered training/learning application built in Go using the **Gotea (go-tea)** framework. Training content is authored as markdown files with custom fenced code blocks for interactive elements (quizzes, diagrams, annotated images, etc.). The server parses all content at startup, manages learner state, and renders HTML over WebSocket using the TEA (The Elm Architecture) pattern.

**This is not a traditional web app.** There is no REST API, no client-side framework, no SPA router. Gotea keeps all state on the server. The browser is a "dumb terminal" that sends messages via WebSocket and receives HTML patches applied by morphdom.

## Tech Stack

- **Go** — all application logic
- **Gotea (`github.com/jpincas/go-tea`)** — server-side TEA framework with WebSocket + morphdom
- **Goldmark** — markdown parsing with custom extensions
- **Chroma** — syntax highlighting for code blocks
- **Tailwind CSS v4** — utility-first CSS via `@tailwindcss/cli`, compiled from `css/main.css` source
- **@tailwindcss/typography** — `prose` classes for Goldmark-rendered narrative HTML
- **KaTeX.js / asciinema-player** — client-side rendering for math, terminal replays (the only JS in the app)
- **SQLite (modernc.org/sqlite)** — optional persistence for learner analytics

## Architecture Essentials

### How Gotea Works

Use the `/gotea` skill for detailed API reference and architectural guidance. Key points:

- The `Model` struct implements `gt.State`. It **must** embed `gt.Router`.
- `Init()` creates a fresh model per session and registers routes.
- `Update()` returns a `gt.MessageMap` — a map of message names to handler functions.
- `Render()` produces `[]byte` of HTML. After any message, the full state re-renders and morphdom diffs the DOM.
- Messages are triggered from HTML via `gt.SendBasicMessage()`, `gt.SendBasicMessageNoArgs()`, etc.
- All state mutations happen in message handlers. Never mutate state outside of them.
- JSON numbers arriving in messages are `float64`. Use `msg.ArgsToInt()` for ints.
- Messages use `SCREAMING_SNAKE_CASE` by convention.
- Use pointer receivers on all `State` interface methods.

### Content Pipeline

```
Markdown files (content/)
    ↓ startup parsing
Goldmark + custom extensions
    ↓ extracts
Page{ NarrativeHTML []byte, Blocks []Block }
    ↓ at render time
Gotea h.Element tree (interactive blocks) + h.UnsafeRaw() (narrative HTML)
    ↓ via WebSocket
Browser (morphdom patches DOM)
```

The key architectural decision: **Goldmark does not render interactive blocks to HTML.** It extracts them as Go structs and leaves placeholder markers (`<div data-block-index="N"></div>`) in the HTML output. At render time, the Gotea render functions interleave the narrative HTML segments with live `h.Element` trees that have proper message handlers.

### State Ownership

- **Shared (read-only at runtime):** `CourseGraph`, `Modules` map, parsed `Page` data. Loaded once at startup, referenced by all sessions.
- **Per-session (mutable):** `CurrentModule`, `CurrentPage`, `Progress`, `ActiveQuiz`, `Preferences`. Each WebSocket connection gets its own `Model` instance via `Init()`.

## Key Directories

```
main.go                    — Entry point, content loading, server start
model.go                   — Model struct, Init, Update, Render
handlers.go                — Message handler implementations
render.go                  — Render functions for pages, layout, nav
blocks.go                  — Block interface and concrete types
block_renderers.go         — Gotea render functions for each block type
course.go                  — CourseGraph, prerequisite logic
parsing/                   — Goldmark extensions and content parsing
  parser.go                — Directory walker, orchestrates parsing
  frontmatter.go           — YAML frontmatter extraction
  goldmark_extensions.go   — Custom fenced block extensions
content/                   — Markdown training content (THE source of truth)
css/main.css               — Tailwind v4 source CSS (theme, plugins, minimal overrides) — NOT the output
static/                    — Built output and vendored assets (do NOT edit files here directly)
  css/main.css             — Built Tailwind output
  js/main.js               — Built JS bundle (gotea client + training init)
  vendor/                  — Vendored third-party JS/CSS (KaTeX)
js/                        — JS source files (built by Parcel into static/js/main.js)
  main.js                  — Entry point: imports gotea client and training.js
  training.js              — Client-side init for KaTeX, asciinema after morphdom patches
package.json               — Build config: `npm run build` builds both JS and CSS
```

## Coding Conventions

### Go Style

- Standard `gofmt` formatting.
- Errors are returned, not panicked (except `MustDecodeArgs` in message handlers where failure is a programming error).
- Table-driven tests preferred.
- No globals except `globalCourse` and `globalModules` (the shared read-only content loaded at startup) and the `app` reference needed for broadcasting.

### Gotea Patterns

- **Render functions are pure.** They take state and return `h.Element`. No side effects.
- **Message handlers are the only place state changes.** They receive the message and the state, mutate it, and return a `gt.Response`.
- **Use component namespacing** (`gt.ComponentID`) when the same interactive block type appears multiple times on a page. Each quiz block instance needs unique message names to avoid collisions.
- **`h.UnsafeRaw()`** is used for Goldmark-rendered narrative HTML. This is safe because the content comes from our own markdown files, not user input.

### Styling with Tailwind v4

All styling uses **Tailwind utility classes applied directly in Go render code** via `a.Class("...")`. There is no custom CSS file with semantic classes.

- **Source CSS:** `css/main.css` — Tailwind v4 configuration (imports, theme, plugins, minimal overrides). This is NOT the output file.
- **Built CSS:** `static/css/main.css` — Generated by `npm run build:css`. Do not edit directly.
- **Tailwind scans Go files:** The `@source "../**/*.go"` directive in `css/main.css` tells Tailwind to scan all `.go` files for utility classes.
- **Theme:** Dark theme using zinc scale. Custom accent color `#00d9c0` defined as `--color-accent` in `@theme`. Use `text-accent`, `bg-accent`, `border-accent`, etc.
- **Fonts:** Plus Jakarta Sans (sans) and JetBrains Mono (mono), loaded via Google Fonts `<link>` in the HTML `<head>`.
- **Typography plugin:** Narrative HTML (from Goldmark) is wrapped in `prose prose-invert prose-zinc` classes. Interactive blocks within prose use `not-prose` to opt out.
- **`content-prose` class:** Minimal CSS overrides in `css/main.css` for Goldmark-specific elements (inline code accent color, link colors, pre block borders). This is the only custom CSS.
- **Preserved class names:** Some classes are required by client-side JS and must not be removed:
  - `external` — used by gotea client to identify links that should not be intercepted for client-side routing
  - `katex-block` — used by KaTeX for math rendering
  - `asciinema-player` — used by asciinema-player for terminal replays
- **Custom utilities:** `font-small`, `font-medium`, `font-large` and `animate-pulse-soft` are defined as `@utility` in the CSS source.
- **No semantic CSS classes.** All styling is Tailwind utilities in Go code. If you need a new style, use Tailwind classes. Only add to `css/main.css` if Tailwind genuinely cannot express it.

### Content Authoring

- Custom blocks are YAML inside fenced code blocks (e.g., ` ```quiz `).
- Every quiz block must have a unique `id` field.
- Block IDs should be kebab-case and scoped to their module: `tcp-handshake-q1`.
- Frontmatter is optional on pages, required in `_module.yaml`.

## Common Tasks

### Adding a New Block Type

1. Define the struct in `blocks.go` implementing the `Block` interface.
2. Add a parser case in `parsing/goldmark_extensions.go` that recognises the new fenced language tag and extracts the YAML into the struct.
3. Add a render function in `block_renderers.go` that produces the `h.Element` tree.
4. Add the case to the `renderBlock()` switch in `render.go`.
5. If the block needs client-side JS (like KaTeX), add initialisation in `js/training.js`.
6. Add tests: parsing test (markdown in → struct out) and render test (struct in → HTML contains expected elements).

### Adding a New Message

1. Add the handler function in `handlers.go`.
2. Register it in `Update()` in `model.go`.
3. Trigger it from a render function using `gt.SendBasicMessage("MESSAGE_NAME", args)`.
4. Test with `tester.NewSession` — dispatch the message and assert state changes.

### Modifying the Course Structure

Course structure is driven entirely by the `content/` directory and `_module.yaml` files. To add a module, create a new directory with a `_module.yaml` and markdown files. To change prerequisites, edit the `prerequisites` field in `_module.yaml`. The course graph is rebuilt at startup.

## Development

**After making Go changes, restart the server:**
```bash
./restart.sh
```
This script kills any existing server on port 8080, rebuilds, and starts fresh. Always run this after Go code changes.

**After changing Tailwind classes in Go files or `css/main.css`, rebuild CSS:**
```bash
npm run build:css
```
This compiles `css/main.css` → `static/css/main.css`. You must do this whenever you add new Tailwind utility classes to Go render code, since Tailwind needs to scan the files and generate the corresponding CSS.

**After changing any JS source file in `js/`, rebuild the client JS:**
```bash
npm run build:js
```
Parcel bundles `js/main.js` (the entry point) and all its imports (including `js/training.js`) into a single `static/js/main.js`. **Never put JS files directly into `static/js/`** — all JS source goes in the `js/` directory and gets bundled by the build.

**Build everything (JS + CSS):**
```bash
npm run build
```

## Testing

```bash
# Run all tests
go test ./...

# Run with verbose output
go test -v ./...

# Run a specific package
go test ./parsing/

# Run the app locally (or use ./restart.sh)
go run .
# Visit http://localhost:8080
```

Use `tester.NewSession(t, &Model{})` for integration tests that exercise the full message → state → render cycle. Remember: dispatch arguments must be `float64` for numbers (JSON encoding).

## Things To Watch Out For

- **morphdom and client-side JS:** After morphdom patches the DOM, client-side libraries (KaTeX) need to re-initialise for new elements. The `js/training.js` source file handles this (bundled into `static/js/main.js` by Parcel). If you add a new client-side rendered block, add its init logic there.
- **Block index stability:** The `data-block-index` placeholder system assumes blocks appear in order in the markdown. Don't reorder blocks between parse and render.
- **Shared vs per-session state:** Never modify `globalCourse` or `globalModules` in a message handler. They're shared across all sessions. Only modify fields on the `Model` instance.
- **Quiz answer payloads:** Quiz answers include both `blockIndex` (position on the page) and `answer` (chosen option index). Both are needed because multiple quizzes can appear on one page.
- **Component namespacing:** If the same block type appears twice on a page with interactive elements, their messages will collide unless you use `gt.ComponentID`. Each block instance should get a unique component ID derived from its block ID.

## Subagents for Content Authoring

Two subagent instruction files live in `.claude/agents/`. These are NOT skills — they are launched as independent agents via the Task tool.

### course-director

`.claude/agents/course-director.md` — The course architect. Designs overall course structure, researches AI trends and skills gaps, writes module briefs, and launches module-builder subagents. Use for course-level decisions.

Launch: `Task(subagent_type="general-purpose", prompt="Read .claude/agents/course-director.md and follow its instructions.\n\n[TASK]")`

### module-builder

`.claude/agents/module-builder.md` — Builds and maintains individual modules. Creates directory structure, `_module.yaml`, markdown pages, callouts, diagrams. Uses `/quiz`, `/agent-demo`, `/mermaid`, `/excalidraw`, and `/image-generator` skills for interactive blocks and visuals. Directed by the course-director or the user.

Launch: `Task(subagent_type="general-purpose", prompt="Read .claude/agents/module-builder.md and follow its instructions.\n\n[MODULE BRIEF]")`

### Hierarchy

```
course-director (subagent)
  └── module-builder (subagent, one per module, can run in parallel)
        ├── /quiz (skill, for quiz blocks)
        ├── /agent-demo (skill, for agent demo blocks)
        ├── /mermaid (skill, for structured diagrams — flowcharts, sequences, timelines)
        ├── /excalidraw (skill, for freeform drawings — sketches, concept maps, visual metaphors)
        └── /image-generator (skill, for infographic-style visuals — conceptual graphics, process visuals)
```

## Reference

- **Gotea docs:** Use the `/gotea` skill for the complete API reference, architecture guidance, and code review. It loads the full Gotea documentation automatically.
- **Course materials reference:** I've placed the materials for the course I wish to implement in the folder `ai-course`
- **Plan:** See `plan.md` for the full architecture plan, phasing, and design decisions.
- **Goldmark:** https://github.com/yuin/goldmark
- **KaTeX:** https://katex.org/
