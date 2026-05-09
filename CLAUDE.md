# Claude.md

This project has two distinct concerns:

1. **[SPA Application](#part-1-spa-application)** — The SvelteKit SPA that renders the training site (marketing page, course content, interactive blocks)
2. **[Course Content](#part-2-course-content)** — The markdown-based training material in `content/` (modules, pages, quizzes, diagrams, agent demos)

Each work cycle should go like this:

1. Check that there are no uncommitted changes
2. Complete work
3. Build and preview: `cd spa && npm run build` then `cd build && sitestakk dev --port 3001 --no-open`
4. Go through any iterations with user, rebuilding and restarting the dev server after changes
5. Confirm finalisation of block of work with user
6. Deploy: `cd spa/build && sitestakk deploy`
7. Commit

---

# Part 1: SPA Application

Everything below this line is about the **SvelteKit SPA, its components, stores, content loading, and build pipeline**. All SPA code lives in the `spa/` directory.

## Project Overview

This is a static SPA for "Learn AI with Jon" — a free, self-paced AI training course for professionals. It's built with SvelteKit in SPA mode (no server at runtime) and deployed to SiteStakk as static files.

The SPA fetches course content (markdown + YAML) at runtime, parses it client-side, and renders interactive training pages with quizzes, scripted AI demo walkthroughs, callouts, and more. Progress is tracked in localStorage.

There is also a marketing landing page at `/` and a contact form at `/contact` (plain HTML, processed by SiteStakk).

**Live site:** https://learnai.jonathanpincas.com

## Tech Stack

- **SvelteKit 2** with `@sveltejs/adapter-static` — SPA mode, no SSR
- **Svelte 5** with runes (`$state`, `$derived`, `$effect`)
- **Tailwind CSS v4** + `@tailwindcss/typography`
- **marked** — client-side markdown parsing
- **js-yaml** — YAML parsing for course structure and custom blocks
- **highlight.js** — syntax highlighting for code blocks (github-dark theme)
- **SiteStakk** — static hosting with forms, favicons, sitemaps

## Architecture

### Content Pipeline (Runtime)

```
content-manifest.json (generated at build time)
    ↓ SPA fetches on load
course.yaml + _module.yaml files
    ↓ fetched per module
Markdown pages (.md)
    ↓ fetched and parsed client-side
    ↓ frontmatter extracted (js-yaml)
    ↓ custom fenced blocks extracted (regex → YAML → typed objects)
    ↓ remaining markdown rendered (marked + highlight.js)
ParsedPage { meta, narrativeHTML, blocks[] }
    ↓ Svelte components
PageContent interleaves narrative HTML with block components
```

The same block placeholder pattern as the old Go app: custom fenced blocks are replaced with `<div data-block-index="N"></div>` markers, and `PageContent.svelte` splits the HTML at those markers and interleaves Svelte block components.

### Key Directories

```
spa/
  src/
    routes/
      +layout.svelte              — Root layout (CSS + fonts)
      +page.svelte                — Marketing landing page (/)
      training/
        +layout.svelte            — Training shell (sidebar + main)
        +page.svelte              — Course overview with progress
        [module]/[page]/
          +page.svelte            — Content page renderer
    lib/
      content/
        types.ts                  — TypeScript interfaces for all content
        loader.ts                 — Fetches and caches YAML/markdown files
        parser.ts                 — Markdown parsing + custom block extraction
      stores/
        progress.ts               — localStorage-backed progress store
        course.ts                 — Course data store (manifest, metadata, quiz index)
        agent.ts                  — Agent demo playback state machine
        completion.ts             — Module completion and unlock logic
      components/
        layout/
          Sidebar.svelte          — Dark sidebar with module/page nav
          PageContent.svelte      — Narrative HTML + block interleaving
          PageNav.svelte          — Previous/Next navigation
        blocks/
          Quiz.svelte             — Interactive quiz with feedback
          Callout.svelte          — Info/warning/tip/danger/note boxes
          AgentDemo.svelte        — Agent demo launcher + state management
          AnnotatedImage.svelte   — Image with positioned hotspots
          Exercise.svelte         — Display-only exercise block
          agent/                  — Agent demo sub-components
            AgentSidebar.svelte
            AgentTitleBar.svelte
            AgentChatMessages.svelte
            AgentChatInput.svelte
            AgentFullContextView.svelte
            AgentStatusBar.svelte
            AgentSystemPromptModal.svelte
            AgentWorkspace.svelte
  static/
    content/                      — Copied from content/ at build time
    content-manifest.json         — Generated at build time
    sitestakk.toml                — SiteStakk deployment config
    contact.html                  — Contact form (plain HTML, SiteStakk processes)
    thanks.html                   — Thank you page after form submit
    favicon.png                   — Source for auto-generated favicons
    images/
      jon.png                     — Headshot used as logo/branding
  scripts/
    generate-manifest.js          — Builds content-manifest.json + copies content/
  build/                          — Production build output (deploy from here)
```

### State Management

- **Course data** (`course.ts`): Loaded once on training layout mount. Contains manifest, module metadata, quiz IDs per module.
- **Progress** (`progress.ts`): localStorage-backed. Tracks pages viewed, quiz scores, last position. Persists across sessions.
- **Completion** (`completion.ts`): Pure functions that check module completion against `_module.yaml` criteria (all pages viewed, required quizzes passed). Used for module locking — modules unlock linearly.
- **Agent playback** (`agent.ts`): Port of the Go playback engine. State machine that processes script events (user, assistant, tool_call, tool_result, note, clear, compaction). Manages chat messages, context, scratchpad, animations.

### Custom Fenced Blocks

The parser recognises these block types in markdown (YAML inside triple backticks):

- **`quiz`** — Multiple-choice questions with feedback
- **`callout`** — Info/warning/tip/danger/note boxes
- **`agent`** / **`agent-demo`** — Scripted AI conversation walkthroughs (inline or external YAML via `path:`)
- **`annotated-image`** — Image with positioned hotspot tooltips
- **`exercise`** — Display-only code exercise

### Agent Demo Playback Engine

The most complex component. `agent.ts` contains the full state machine ported from the Go `agent_playback.go`. Key concepts:

- **Event groups**: A user message + subsequent tool_call/tool_result/assistant events form a group. The group animates in with staggered delays.
- **Tool side effects**: `scratchpad_write`, `move_file`, `create_folder`, `delete_file` mutate the virtual scratchpad.
- **Workspace mode**: Agent demos with scratchpad files open in a full-screen workspace (file explorer + editor + chat sidebar).
- **External YAML**: Some agent demos reference external YAML files via `path:` — these are fetched on demand.

## Styling

- **Theme**: Dark marketing site (`#020617` slate-950), white training content area, dark slate-950 sidebar
- **Accent**: Teal `#14b8a6` with hover `#0d9488`
- **Fonts**: Plus Jakarta Sans (body), JetBrains Mono (code)
- **CSS source**: `spa/src/app.css` — Tailwind v4 config with theme, custom utilities, prose overrides
- **Code blocks**: Dark background (`#0f172a`), github-dark highlight.js theme
- **Custom utilities**: `text-gradient`, `hero-grid`, `font-small/medium/large`, `animate-pulse-soft`

## Development

```bash
cd spa

# Development (Vite dev server with HMR)
npm run dev

# Production build
npm run build
# This runs: generate-manifest → vite build → copy index.html to 404.html

# Preview with SiteStakk (from build output)
cd build && sitestakk dev --port 3001 --no-open

# Deploy
cd build && sitestakk deploy

# Type check
npx svelte-check --tsconfig ./tsconfig.json
```

The `404.html` is a copy of `index.html` — SiteStakk serves it for all unmatched routes, enabling SPA client-side routing.

## SiteStakk Configuration

Config lives in `spa/static/sitestakk.toml` (copied to build automatically):

```toml
name = "learn-ai-with-jon"
domain = "learnai.jonathanpincas.com"

[images]
optimize = false

[forms.contact]
```

- **Forms**: The contact form is plain HTML (`static/contact.html`). SiteStakk detects `data-sitestakk="contact"` at deploy time and injects the action URL and honeypot. Submissions are stored and queryable via `sitestakk forms submissions <form-id>`.
- **Favicon**: `static/favicon.png` — SiteStakk auto-generates all favicon variants.
- **SPA fallback**: `404.html` serves the SPA shell for all routes.
- **Deploy from**: Always deploy from `spa/build/`, not `spa/`.

## Things To Watch Out For

- **Content manifest**: Must be regenerated when content files change (`npm run generate-manifest` or `npm run build`).
- **Static HTML pages**: The contact form and thank you page are plain HTML in `static/`, NOT Svelte routes. This is intentional — SiteStakk needs to scan the HTML at deploy time to inject form processing.
- **Agent demo external YAML**: Two demos use `path:` references to external YAML files. The parser fetches these on demand at runtime.
- **Module locking**: Modules unlock linearly. Completion requires all pages viewed + required quizzes passed per `_module.yaml` completion criteria.
- **Quiz IDs**: Extracted at course load time via regex scan of all markdown files (not full parsing). Stored in `course.quizIdsByModule`.

---

# Part 2: Course Content

Everything below this line is about **authoring and editing the training course content** — the markdown files, module structure, quizzes, diagrams, and agent demos that live under `content/`.

## Module Map

The course has 10 modules delivered in order. **When the user refers to a module by number, use this map.** The directory name is the key used in code and file paths.

| # | Directory | Title |
|---|-----------|-------|
| 1 | `module-opening` | Opening: The February 2026 Moment |
| 2 | `module-llms` | How LLMs Actually Work |
| 3 | `module-context` | Context — The Most Important Concept |
| 4 | `module-prompting` | The Art of Prompting |
| 5 | `module-files` | Files — The Unit of Work |
| 6 | `module-writing` | Document Creation and Data Analysis |
| 7 | `module-advanced` | Advanced Patterns |
| 8 | `module-delegation` | Delegation & The AI-First Philosophy |
| 9 | `module-risks` | Risks, Responsibility, and Realistic Expectations |
| 10 | `module-synthesis` | Putting It Together |

The module order is defined in `content/course.yaml`. Prerequisites follow the same linear sequence.

## Content Directory Structure

```
content/
  course.yaml              — Module order and course metadata
  module-opening/          — Module 1
    _module.yaml           — Module metadata (title, prerequisites)
    01-welcome.md          — Pages numbered sequentially
    02-what-youll-learn.md
    03-demo.md
    images/                — Module-specific images and diagrams
  module-llms/             — Module 2
    _module.yaml
    01-prediction-engines.md
    ...
  (same pattern for all modules)
```

Pages within a module are ordered by their numeric filename prefix (01-, 02-, etc.).

## Content Authoring Conventions

- Custom blocks are YAML inside fenced code blocks (e.g., ` ```quiz `).
- Every quiz block must have a unique `id` field.
- Block IDs should be kebab-case and scoped to their module: `tcp-handshake-q1`.
- Frontmatter is optional on pages, required in `_module.yaml`.

## Modifying the Course Structure

Course structure is driven entirely by the `content/` directory and `_module.yaml` files. To add a module, create a new directory with a `_module.yaml` and markdown files. To change prerequisites, edit the `prerequisites` field in `_module.yaml`. After changes, rebuild: `cd spa && npm run build`.

## Subagents for Content Authoring

Two subagent instruction files live in `.claude/agents/`. These are NOT skills — they are launched as independent agents via the Agent tool.

### course-director

`.claude/agents/course-director.md` — The course architect. Designs overall course structure, researches AI trends and skills gaps, writes module briefs, and launches module-builder subagents. Use for course-level decisions.

### module-builder

`.claude/agents/module-builder.md` — Builds and maintains individual modules. Creates directory structure, `_module.yaml`, markdown pages, callouts, diagrams. Uses `/quiz`, `/agent-demo`, `/mermaid`, `/excalidraw`, and `/image-generator` skills for interactive blocks and visuals. Directed by the course-director or the user.

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

## Reference (Course Content)

- **Course materials reference:** I've placed the materials for the course I wish to implement in the folder `ai-course`
- **Plan:** See `plan.md` for the full architecture plan, phasing, and design decisions.
