# SPA Rewrite Plan

## Overview
Rewrite the Go/Gotea server-rendered training app as a static SvelteKit SPA. No backend, no build step for content. Deploy on SiteStakk.

## Tech Stack
- SvelteKit 2 + adapter-static (SPA mode)
- Svelte 5 (runes)
- Tailwind CSS v4 + @tailwindcss/typography
- marked (client-side markdown parsing)
- js-yaml (YAML parsing)
- highlight.js (syntax highlighting)
- KaTeX (math rendering)

## Content Strategy
- Build-time manifest script lists all content files → `content-manifest.json`
- SPA fetches manifest → fetches YAML/markdown → parses client-side
- Content directory copies unchanged to `static/content/`
- Image paths work as-is

## Routing
```
/                              → Marketing landing page
/training                      → Redirects to first (or last) module/page
/training/[module]/[page]      → Specific page
```

## State Management
- localStorage-backed Svelte store
- Pages viewed, quiz scores, last position, font size, module completion

## What We Drop
- Go server, WebSocket, Gotea, morphdom
- Authentication, cohorts, sessions
- Presenter mode, live polls, follower mode
- Annotation canvas, SQLite analytics, admin panel

## Implementation Phases

### Phase 1: Skeleton
- SvelteKit project setup with adapter-static
- Tailwind v4 config (port theme)
- Content manifest generation script
- Copy content/ to static/content/
- Basic routing: /, /training/[module]/[page]
- TypeScript types for all content

### Phase 2: Content Loading & Parsing
- loader.ts: fetch course.yaml, module YAML, markdown
- parser.ts: frontmatter extraction, custom block extraction (regex), marked rendering
- Section splitting at H2 boundaries
- Verify all markdown files parse correctly

### Phase 3: Core Training UI
- Sidebar.svelte: module list, page list, navigation
- PageContent.svelte: narrative HTML + block interleaving
- PageNav.svelte: previous/next with slide support
- Callout.svelte: simplest block
- Code syntax highlighting with highlight.js
- KaTeX math rendering

### Phase 4: Quiz & Simple Blocks
- Quiz.svelte: full interaction (select, feedback, retry)
- AnnotatedImage.svelte: hotspot toggle
- Exercise.svelte: display-only
- Progress store (localStorage) with quiz score tracking

### Phase 5: Agent Demo (most complex)
- Port AgentState and playback engine from Go to TypeScript
- AgentDemo.svelte sidebar layout
- Chat messages with animation
- Scratchpad file explorer
- System prompt modal, full context view
- Tool call rendering and scratchpad side effects
- External agent YAML loading

### Phase 6: Progress Tracking & Navigation
- Module completion detection
- Page/slide viewed tracking
- Prerequisite enforcement
- "Continue" button on landing page
- Font size preference persistence

### Phase 7: Marketing Landing Page
- Hero with dynamic Start/Continue button
- Module overview grid
- Responsive design

### Phase 8: Polish & Deploy
- Responsive design, loading states, error handling
- Scroll-to-top on navigation
- Performance testing
- Deploy to SiteStakk
