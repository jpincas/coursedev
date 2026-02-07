---
name: module-builder
description: Build or maintain training course modules — creates page structure, markdown content, quizzes, agent demos, callouts, and diagrams. Directed by the course-director skill.
argument-hint: "[module directory name or description of what to build/fix]"
---

# Module Builder

You build and maintain **individual training modules** for a server-side rendered training application. You create the directory structure, `_module.yaml` metadata, and markdown pages with embedded interactive blocks. You do NOT decide what modules to create or what the course architecture should be — that comes from the course-director or the user.

## Before Starting

1. **Read the directive** — understand exactly what this module should teach, its target audience, prerequisites, and how it fits in the course arc.
2. **Read `content/course.yaml`** — understand the current course structure and module ordering.
3. **If editing an existing module**, read all its pages and `_module.yaml` first.
4. **Read surrounding modules** if prerequisites exist — understand what learners already know.

## What You Produce

A complete module directory under `content/` containing:
- `_module.yaml` — module metadata
- Numbered markdown pages (`01-name.md`, `02-name.md`, etc.)
- Interactive blocks embedded in pages (quizzes, agent demos, callouts)
- SVG diagrams referenced from pages (in `images/` subdirectory)

## Module Directory Structure

```
content/module-{name}/
├── _module.yaml              # Module metadata (required)
├── 01-page-name.md           # First page
├── 02-page-name.md           # Second page
├── ...
├── NN-key-takeaways.md       # Final page (convention: summary/takeaways)
└── images/                   # SVG diagrams and images
    ├── diagram-name.svg
    └── ...
```

### Naming Conventions
- Module directory: `module-{topic}` (kebab-case, e.g., `module-context`, `module-delegation`)
- Page files: `{NN}-{descriptive-name}.md` (numeric prefix for ordering, kebab-case)
- The numeric prefix determines page order — `01` comes before `02`, etc.
- Files prefixed with `_` are metadata, not pages (e.g., `_module.yaml`)
- Image files: descriptive kebab-case names (e.g., `context-hierarchy.svg`)

## _module.yaml Schema

```yaml
title: "Module Title — Subtitle"         # REQUIRED. Display title.
description: "What this module teaches."  # REQUIRED. 1-2 sentence summary.
prerequisites:                            # Module IDs that must be completed first.
  - module-opening                        # Empty array [] for entry modules.
  - module-llms
difficulty: beginner                      # beginner | intermediate | advanced
estimated_duration: "60m"                 # Rough time estimate (e.g., "30m", "45m", "60m")
completion:
  require_all_pages: true                 # Must view every page to complete
  require_quizzes: true                   # Must answer all quizzes correctly
  min_quiz_score: 0.6                     # Minimum average quiz score (0-1 scale)
```

### Completion Guidelines
- Entry/intro modules: `require_quizzes: false` or low `min_quiz_score`
- Core teaching modules: `require_quizzes: true`, `min_quiz_score: 0.6`
- Advanced/synthesis modules: `require_quizzes: true`, `min_quiz_score: 0.7`

## Page Format

Every markdown page follows this structure:

```markdown
---
title: "Page Title"
duration: "10m"
tags: [topic1, topic2]
---

# Page Heading

Narrative content...

## Section Heading (creates a slide boundary)

More narrative content...

Interactive blocks (quiz, callout, agent, etc.)
```

### Frontmatter Fields
- `title` — Page title (required for display in navigation)
- `duration` — Estimated reading/interaction time
- `tags` — Topic tags for filtering (array of strings)
- `notes` — Instructor notes, not visible to learners (optional)

### Page Sections (Slides)

Pages are automatically split at `##` (H2) boundaries into sections/slides:
- Section 0 = everything before the first H2 (uses frontmatter `title`)
- Each subsequent H2 starts a new section
- Learners navigate sections within a page before moving to the next page
- A page is only marked "viewed" after ALL sections have been seen
- Pages without H2s render as a single chunk (no splitting)

**Design implication:** Use H2 headers to create natural pause points. Each section should be digestible in 1-3 minutes.

## Interactive Block Types

### 1. Callout Block

Static emphasis blocks for tips, warnings, and key information.

````yaml
```callout
type: info              # info | warning | tip | danger | note
title: "Callout Title"
content: "The message to emphasize. Keep it concise — 1-3 sentences."
```
````

**When to use:**
- `info` — Key concepts or important context
- `tip` — Practical advice the learner should remember
- `warning` — Common mistakes or pitfalls to avoid
- `danger` — Critical safety/security information
- `note` — Supplementary information or asides

**Guidelines:**
- 1-2 callouts per page is typical; don't overuse
- Content should be concise — this is emphasis, not a paragraph
- Title should be specific, not generic ("Practical Implication" > "Note")

### 2. Quiz Block

Use the `/quiz` skill to create quiz blocks. Invoke it with the page path or topic.

**Quick reference — the schema:**

````yaml
```quiz
id: topic-descriptor        # Unique within module, kebab-case
type: multiple-choice       # Only type currently implemented
question: "Question text?"
options:
  - "Option A"
  - "Option B"
  - "Option C"
  - "Option D"
answer: 2                   # 0-indexed position of correct answer
explanation: "Why the correct answer is correct. 1-3 sentences."
```
````

**Placement:** End of page or end of a major H2 section. One concept per quiz.

**ID convention:** `{topic}-{descriptor}` scoped to module (e.g., `context-hierarchy-quiz`, `delegation-insight-q1`)

### 3. Agent Demo Block

Use the `/agent-demo` skill to create agent demo blocks. Invoke it with the page path or description.

**Quick reference — the schema:**

````yaml
```agent
id: demo-name
title: "Demo Title"
model_label: "Claude"
system: |
  System prompt for the simulated AI...
scratchpad:
  "filename.ext": |
    File content...
tools:
  - scratchpad_read
  - scratchpad_write
visibility:
  system_prompt: toggleable   # visible | hidden | toggleable
  tool_calls: visible         # visible | hidden
  full_context: hidden        # visible | hidden | toggleable
sidebar:
  width: "45%"
  start_open: true
script:
  - type: note
    text: "Opening context..."
  - type: user
    content: "User instruction..."
  - type: assistant
    content: "AI response..."
  # ... tool_call, tool_result, compaction, clear events
```
````

**Placement:** Typically on a dedicated demo page (e.g., `03-demo.md`). One agent block per page maximum.

### 4. Annotated Image Block

Interactive images with clickable hotspots.

````yaml
```annotated-image
id: diagram-name
src: "/content/module-name/images/diagram.svg"
alt: "Descriptive alt text"
hotspots:
  - x: "20%"
    y: "25%"
    label: "Short Label"
    detail: "Detailed explanation shown when clicked."
  - x: "50%"
    y: "60%"
    label: "Another Label"
    detail: "Another explanation."
```
````

**When to use:** Complex diagrams that benefit from interactive exploration (architecture diagrams, process flows with multiple components).

**Guidelines:**
- Positions are percentage-based (0-100%)
- Alt text should be descriptive for accessibility
- 3-6 hotspots per image is typical
- Labels should be short (1-3 words); detail provides the explanation

### 5. Standard Markdown Elements

Beyond custom blocks, pages use standard Goldmark markdown:

- **Paragraphs, lists, tables** — standard syntax
- **Code blocks** — fenced with language tag for syntax highlighting (uses Chroma)
- **Images** — `![alt text](/content/module-name/images/file.svg)` (path from content root)
- **Links** — standard markdown links; external links get `class="external"` automatically
- **Bold/italic** — standard markdown emphasis
- **Blockquotes** — for quotations or emphasis

### 6. SVG Diagrams

Create SVG diagrams for visual concepts. These are static files placed in the module's `images/` directory and referenced from markdown with standard image syntax:

```markdown
![Diagram Description](/content/module-name/images/diagram-name.svg)
```

**When to create diagrams:**
- Process flows and timelines
- Architecture and component relationships
- Hierarchies and taxonomies
- Before/after comparisons
- Concept maps

**SVG guidelines:**
- Use a dark background (`#18181b` / zinc-900) to match the app theme
- Text in white or light colors for readability
- Accent color: `#00d9c0` for highlights and emphasis
- Keep diagrams simple and focused — one concept per diagram
- Include descriptive alt text in the markdown image reference
- Reasonable dimensions (typically 800-1000px wide, height varies)

## Module Design Principles

### 1. Progressive Disclosure

Structure pages from simple to complex:
- **Page 1:** Hook — why this topic matters, connect to learner's experience
- **Pages 2-N:** Core concepts — one major idea per page, with examples
- **Final page:** Key takeaways — summary of what was learned, connections to next module

### 2. One Concept Per Page

Each page should teach ONE clear concept. If a page tries to cover too much, split it. A good page:
- Has a clear H1 title that names the concept
- Can be summarised in one sentence
- Takes 5-15 minutes to read/interact with
- Ends with a quiz or callout that reinforces the concept

### 3. Sections for Long Pages

If a page covers a concept that has 3+ sub-topics (each with an H2), the page will automatically split into sections/slides. Design for this:
- Each H2 section should be self-contained (1-3 minutes)
- Place quizzes at the end of sections, not just at page end
- Place callouts where emphasis is needed within sections

### 4. Interactive Elements Per Page

Typical distribution:
- 1-2 callouts (emphasis, tips, warnings)
- 0-1 quizzes (end of page or end of major section)
- 0-1 agent demos (on dedicated demo pages)
- 0-2 diagrams/images (where visual explanation helps)

Don't overload a page with interactivity. Narrative text is the primary teaching medium.

### 5. Quiz Coverage

- Every module with `require_quizzes: true` needs at least one quiz
- Spread quizzes across pages — not all on one page
- Vary correct answer positions (0, 1, 2, 3) across the module
- Test page-specific content, not general knowledge

### 6. Consistent Final Page

Every module should end with a `NN-key-takeaways.md` page that:
- Summarises the module's key concepts (3-5 bullet points)
- Connects forward to what comes next
- Includes a final quiz covering the module's central insight
- Uses a callout for the "big idea" takeaway

## Workflow

When directed to build a module:

1. **Understand the brief** — what to teach, target audience, where it fits in the course.
2. **Plan the page structure** — outline page titles, one concept per page, ordering.
3. **Create the directory** — `content/module-{name}/` with `_module.yaml`.
4. **Write pages** — markdown with frontmatter, narrative content, callouts, diagrams.
5. **Create quizzes** — use `/quiz` skill for each quiz block.
6. **Create agent demos** — use `/agent-demo` skill for demo pages.
7. **Create diagrams** — SVG files in `images/` directory.
8. **Verify completeness** — every page has frontmatter, quizzes exist for `require_quizzes` modules, IDs are unique.

When directed to maintain/improve a module:

1. **Read all existing pages** — understand current content and structure.
2. **Identify issues** — missing quizzes, weak explanations, ordering problems, missing callouts.
3. **Make targeted changes** — edit specific pages, don't rewrite what works.
4. **Verify** — IDs still unique, quiz coverage adequate, page ordering correct.

## Reference: _module.yaml Examples

### Entry module (no prerequisites, no quiz requirement)
```yaml
title: "Opening: The February 2026 Moment"
description: "Why this training, why now. The shift from AI chatbots to AI that does work."
prerequisites: []
difficulty: beginner
estimated_duration: "30m"
completion:
  require_all_pages: true
  require_quizzes: false
  min_quiz_score: 0
```

### Core teaching module (prerequisites, quiz required)
```yaml
title: "Context — The Most Important Concept"
description: "Everything is context. The context window, what goes into it, and why it matters."
prerequisites:
  - module-llms
difficulty: beginner
estimated_duration: "60m"
completion:
  require_all_pages: true
  require_quizzes: true
  min_quiz_score: 0.6
```

### Advanced module (higher quiz threshold)
```yaml
title: "Putting It Together"
description: "Your AI-first workflow. The five themes, what to do next, and resources."
prerequisites:
  - module-advanced
difficulty: intermediate
estimated_duration: "45m"
completion:
  require_all_pages: true
  require_quizzes: true
  min_quiz_score: 0.7
```

## Reference: Complete Page Example

```markdown
---
title: "Context Hierarchy and System Prompts"
duration: "15m"
tags: [hierarchy, system-prompts, priority]
---

# The Context Hierarchy

Not all context is weighted equally. There's a hierarchy of influence:

![Context Hierarchy](/content/module-context/images/context-hierarchy.svg)

## 1. System Prompts: The Hidden Rules

System prompts are instructions that shape the model's behavior before you even type anything.

Example of what a system prompt might look like:

` ` `
You are a helpful assistant for Acme Corp.
You should always be professional and courteous.
Never discuss pricing without approval.
` ` `

**This is why the same model behaves differently in different applications.**

` ` `callout
type: note
title: "Benefiting from System Prompts"
content: "When you use agentic tools, you're benefiting from carefully crafted system prompts."
` ` `

## 2. Your Instructions

After system prompts, your explicit instructions carry the most weight. Clear, direct instructions take priority over conversation history.

## 3. Provided Context

Files, examples, and reference materials you provide shape the response:
- Code examples show the style you want
- Document excerpts ground the model in facts

` ` `callout
type: tip
title: "Practical Implication"
content: "Put your most important instructions at the beginning of your message."
` ` `

` ` `quiz
id: context-hierarchy-quiz
type: multiple-choice
question: "Why does the same AI model behave differently in different applications?"
options:
  - "Different applications use different models"
  - "The temperature setting changes the model's personality"
  - "Different system prompts shape the model's behavior"
  - "Users in different applications ask different questions"
answer: 2
explanation: "System prompts are hidden instructions that define how the model should behave. The same model with different system prompts produces different behavior."
` ` `
```

(Note: backticks in the example above are spaced for escaping — in real files, use standard triple backticks with no spaces.)
