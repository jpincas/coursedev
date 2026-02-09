---
name: mermaid
description: Create Mermaid diagrams for training course modules. Writes .mmd source files and renders them to SVG using the project's dark theme. Use when a lesson page needs a structured diagram (flowchart, sequence, timeline, etc.).
argument-hint: "[module path and diagram description]"
---

# Mermaid Diagram Author

You create **Mermaid diagrams** — structured visual diagrams rendered as SVG files for training course lesson pages. Mermaid is best for diagrams with defined structure: flowcharts, sequences, timelines, class diagrams, state diagrams, and similar.

## Before Starting

1. **Understand the context** — read the page or module brief to know what concept the diagram should illustrate.
2. **Check existing diagrams** — look in the module's `images/` directory. Avoid duplicating or conflicting with existing visuals.
3. **Choose the right diagram type** — match the concept to the best Mermaid diagram type (see Diagram Types below).

## What You Produce

Two files in the module's `images/` directory, plus a markdown reference:

1. **`{name}.mmd`** — Mermaid source file (the editable source of truth)
2. **`{name}.svg`** — Rendered SVG (what the browser displays)
3. **Markdown reference** — `![Alt text](/content/module-{name}/images/{name}.svg)` for the page

## File Placement

```
content/module-{name}/images/
├── diagram-name.mmd          # Mermaid source
├── diagram-name.svg          # Rendered SVG
└── ...
```

## Rendering Command

The project has `@mermaid-js/mermaid-cli` installed and a theme config at `mermaid-config.json`. Render with:

```bash
npx mmdc -i content/module-{name}/images/{diagram}.mmd -o content/module-{name}/images/{diagram}.svg -c mermaid-config.json -b transparent
```

**Always use these flags:**
- `-c mermaid-config.json` — applies the project's dark theme (zinc scale, accent `#00d9c0`, Plus Jakarta Sans font)
- `-b transparent` — transparent background (the SVG sits on the page's dark background)

## Theme (Automatic)

The `mermaid-config.json` at the project root handles all theming. You do NOT need to add theme directives in `.mmd` files. The config provides:

- **Background:** transparent (page background is zinc-900 `#18181b`)
- **Node fill:** zinc-800 `#27272a` with zinc-700 `#3f3f46` borders
- **Text:** zinc-50 `#fafafa`
- **Edges/lines:** zinc-400 `#a1a1aa`
- **Accent color:** `#00d9c0` (used for actor borders, active tasks, today line)
- **Font:** Plus Jakarta Sans, system-ui, sans-serif
- **Flowchart curves:** basis (smooth)

## Diagram Types and When to Use Them

### Flowchart (`flowchart`)
**Use for:** Processes, decision trees, data flows, hierarchies, relationships.

```mermaid
flowchart TD
    A["Step 1<br/><small>Description</small>"] --> B["Step 2"]
    B --> C{"Decision?"}
    C -->|Yes| D["Path A"]
    C -->|No| E["Path B"]
```

**Direction options:** `TD` (top-down), `LR` (left-right), `TB` (top-bottom), `RL` (right-left)

**Node shapes:**
- `["text"]` — rectangle (default)
- `{"text"}` — diamond (decision)
- `("text")` — rounded rectangle
- `(["text"])` — stadium shape
- `[["text"]]` — subroutine
- `(("text"))` — circle

**Tips:**
- Use `<br/>` for line breaks within nodes
- Use `<small>...</small>` for secondary text (descriptions, notes)
- Keep node text concise — 3-5 words max per line
- Use descriptive edge labels for decisions: `-->|Yes|` not `-->|Y|`

### Sequence Diagram (`sequenceDiagram`)
**Use for:** Interactions between components, API calls, message passing, protocols.

```mermaid
sequenceDiagram
    participant U as User
    participant A as AI Agent
    participant T as Tool

    U->>A: "Analyse this data"
    A->>T: read_file(data.csv)
    T-->>A: file contents
    A->>U: Analysis results
```

**Arrow types:**
- `->>` solid with arrowhead
- `-->>` dotted with arrowhead
- `->>+` activates participant
- `->>-` deactivates participant

### Timeline (`timeline`)
**Use for:** Historical progressions, evolution of concepts, phases.

```mermaid
timeline
    title Evolution of AI Tools
    2022 : ChatGPT launches
         : Text-only chatbots
    2023 : GPT-4 with vision
         : Code generation tools
    2024 : Agentic AI emerges
         : Multi-step tool use
```

### State Diagram (`stateDiagram-v2`)
**Use for:** States and transitions, lifecycle diagrams, mode changes.

```mermaid
stateDiagram-v2
    [*] --> Draft
    Draft --> Review : Submit
    Review --> Published : Approve
    Review --> Draft : Reject
    Published --> [*]
```

### Gantt Chart (`gantt`)
**Use for:** Project timelines, phase breakdowns, parallel workstreams.

```mermaid
gantt
    title Training Programme
    dateFormat YYYY-MM-DD
    section Foundation
        Opening          :a1, 2026-01-01, 1d
        LLM Fundamentals :a2, after a1, 2d
    section Core Skills
        Context          :b1, after a2, 2d
        Prompting        :b2, after b1, 2d
```

### Other Supported Types

- **`classDiagram`** — class relationships, type hierarchies
- **`erDiagram`** — entity relationships, data models
- **`pie`** — proportional breakdowns (use sparingly)
- **`mindmap`** — topic exploration, brainstorming maps
- **`graph`** — generic directed/undirected graphs

## Design Principles

### 1. One Concept Per Diagram
Each diagram should illustrate ONE clear concept. If a diagram tries to show everything, it shows nothing. Split complex systems into multiple focused diagrams.

### 2. Minimal Nodes
3-7 nodes is ideal. More than 10 nodes and the diagram becomes hard to read. If you need more, consider splitting into multiple diagrams or using a different representation.

### 3. Clear Labels
- Node labels: 2-5 words, noun phrases
- Edge labels: 1-3 words, verb phrases or conditions
- Use `<small>` for supplementary detail that aids understanding

### 4. Direction Matters
- **Top-down (`TD`):** hierarchies, priorities, processes with clear sequence
- **Left-right (`LR`):** workflows, data pipelines, temporal progressions
- Choose based on what the reader expects for this type of information

### 5. Consistent Node Shapes
Within a diagram, use shapes consistently:
- Rectangles for steps/processes/entities
- Diamonds for decisions only
- Rounded rectangles for start/end points
- Don't mix shapes arbitrarily

## Naming Convention

```
{concept}-{descriptor}.mmd
```

Examples:
- `context-hierarchy.mmd` — shows the context priority hierarchy
- `file-workflow.mmd` — shows the file-based AI workflow
- `training-process.mmd` — shows how LLM training works
- `mcp-connections.mmd` — shows MCP server connections

Use kebab-case. Keep names descriptive but concise.

## Workflow

When asked to create a Mermaid diagram:

1. **Understand the concept** — what should the reader learn from this diagram?
2. **Choose the diagram type** — flowchart, sequence, timeline, etc.
3. **Draft the `.mmd` source** — write minimal, clear Mermaid syntax.
4. **Write the file** — save to `content/module-{name}/images/{diagram}.mmd`
5. **Render to SVG** — run `npx mmdc -i {input}.mmd -o {output}.svg -c mermaid-config.json -b transparent`
6. **Verify the SVG was created** — check the file exists and has content.
7. **Output the markdown reference** — `![Alt text](/content/module-{name}/images/{diagram}.svg)`

## Reference Examples

### Example 1: Simple Hierarchy (flowchart TD)

Source: `context-hierarchy.mmd`
```
flowchart TB
    A["1. System Prompt<br/><small>Highest priority - sets the rules</small>"] --> B["2. User Instructions<br/><small>Your direct requests and guidance</small>"]
    B --> C["3. Provided Context<br/><small>Files, examples, reference materials</small>"]
    C --> D["4. Conversation History<br/><small>What's been said before</small>"]
```

### Example 2: Process Flow (flowchart LR)

Source: `file-workflow.mmd`
```
flowchart LR
    A[Input Files<br/>Data, Documents] --> B[AI Processing]
    B --> C[Output Files<br/>Deliverables]
```

### Example 3: Hub-and-Spoke (flowchart LR)

Source: `mcp-connections.mmd`
```
flowchart LR
    A["AI"] <--> B["MCP"]
    B <--> C["Files"]
    B <--> D["Databases"]
    B <--> E["Web APIs"]
    B <--> F["Email"]
    B <--> G["Calendar"]
```

## Common Mistakes to Avoid

- **Don't add `%%{init:...}%%` theme directives** — the `mermaid-config.json` handles all theming
- **Don't use very long node text** — keep labels short, use `<small>` for detail
- **Don't forget the render step** — always run `npx mmdc` to generate the SVG
- **Don't use diagram types not supported by mmdc** — stick to the types listed above
- **Don't create diagrams for things better shown as text** — a 2-item list doesn't need a flowchart
