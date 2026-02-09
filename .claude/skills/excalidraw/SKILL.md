---
name: excalidraw
description: Create Excalidraw drawings for training course modules. Produces hand-drawn style diagrams as .excalidraw source files with rendered SVG output. Use for freeform conceptual diagrams, architectural sketches, annotated illustrations, and visual metaphors that don't fit structured Mermaid charts.
argument-hint: "[module path and drawing description]"
---

# Excalidraw Drawing Author

You create **Excalidraw drawings** — freeform, hand-drawn style diagrams for training course lesson pages. Excalidraw is best for conceptual illustrations, architectural sketches, annotated diagrams, visual metaphors, and anything that benefits from a whiteboard-style aesthetic that structured chart tools like Mermaid can't express.

## When to Use Excalidraw vs Mermaid

| Use Excalidraw | Use Mermaid (`/mermaid`) |
|---|---|
| Conceptual/metaphorical diagrams | Flowcharts with defined logic |
| Freeform spatial layouts | Sequence diagrams |
| Annotated illustrations | Timelines with dates |
| Before/after comparisons | Entity relationship diagrams |
| Mind maps with custom layout | State machines |
| Architecture sketches | Gantt charts |
| Visual metaphors | Any diagram with strict structure |

**Rule of thumb:** If the diagram has defined nodes and edges that follow a pattern, use Mermaid. If it's more like something you'd sketch on a whiteboard to explain a concept, use Excalidraw.

## Before Starting

1. **Understand the concept** — what should the learner take away from this visual?
2. **Check existing diagrams** — look in the module's `images/` directory for what already exists.
3. **Plan the layout** — sketch mentally how the elements should be arranged spatially.

## What You Produce

Two files in the module's `images/` directory, plus a markdown reference:

1. **`{name}.excalidraw`** — Excalidraw JSON source (editable in excalidraw.com)
2. **`{name}.svg`** — Hand-crafted SVG matching the app theme (what the browser displays)
3. **Markdown reference** — `![Alt text](/content/module-{name}/images/{name}.svg)` for the page

## File Placement

```
content/module-{name}/images/
├── drawing-name.excalidraw    # Excalidraw source (editable)
├── drawing-name.svg           # Rendered SVG (served)
└── ...
```

## Two-File Approach

You produce BOTH an `.excalidraw` source file and an `.svg` file:

- **`.excalidraw`** — The editable source. Humans can open this at excalidraw.com to make adjustments. This is the canonical source of truth for the drawing's structure.
- **`.svg`** — The rendered output. You create this directly as hand-crafted SVG that matches the app's dark theme. This is what the browser displays.

Both files represent the same diagram. The SVG must visually match the Excalidraw source.

## App Theme

All drawings must match the application's dark theme:

| Element | Color | Use |
|---|---|---|
| Background | `#18181b` (zinc-900) | SVG/drawing background |
| Shape fill | `#27272a` (zinc-800) | Default shape background |
| Shape stroke | `#3f3f46` (zinc-700) | Default borders |
| Text | `#fafafa` (zinc-50) | Primary text, labels |
| Secondary text | `#a1a1aa` (zinc-400) | Descriptions, annotations |
| Accent | `#00d9c0` | Highlights, emphasis, key elements |
| Accent hover | `#0f766e` | Secondary accent |
| Connectors | `#a1a1aa` (zinc-400) | Arrows, lines |
| Font | Plus Jakarta Sans, system-ui, sans-serif | All text |

## Excalidraw JSON Format

The `.excalidraw` file is a JSON document with this structure:

```json
{
  "type": "excalidraw",
  "version": 2,
  "source": "https://excalidraw.com",
  "elements": [ ... ],
  "appState": {
    "viewBackgroundColor": "#18181b",
    "gridSize": null
  }
}
```

### Element Types

Every element shares these base properties:

```json
{
  "id": "unique-id-string",
  "type": "rectangle",
  "x": 100,
  "y": 100,
  "width": 200,
  "height": 80,
  "angle": 0,
  "strokeColor": "#3f3f46",
  "backgroundColor": "#27272a",
  "fillStyle": "solid",
  "strokeWidth": 2,
  "roughness": 1,
  "opacity": 100,
  "groupIds": [],
  "roundness": { "type": 3 },
  "seed": 12345,
  "version": 1,
  "versionNonce": 1,
  "isDeleted": false,
  "boundElements": null,
  "updated": 1700000000000,
  "link": null,
  "locked": false
}
```

**Key properties:**
- `roughness`: `0` = clean lines, `1` = slightly rough (hand-drawn), `2` = very rough. Use `1` for the hand-drawn Excalidraw aesthetic.
- `fillStyle`: `"solid"` (filled), `"hachure"` (hatched lines), `"cross-hatch"` (crossed lines). Use `"solid"` for the dark theme.
- `roundness`: `{ "type": 3 }` for rounded corners, `null` for sharp corners.
- `seed`: Random integer for hand-drawn randomisation. Use different values per element.

### Rectangle

```json
{
  "type": "rectangle",
  "x": 100, "y": 100,
  "width": 200, "height": 80,
  "strokeColor": "#3f3f46",
  "backgroundColor": "#27272a",
  "fillStyle": "solid",
  "strokeWidth": 2,
  "roughness": 1,
  "roundness": { "type": 3 }
}
```

### Ellipse

```json
{
  "type": "ellipse",
  "x": 100, "y": 100,
  "width": 150, "height": 150,
  "strokeColor": "#00d9c0",
  "backgroundColor": "transparent",
  "fillStyle": "solid",
  "strokeWidth": 2,
  "roughness": 1
}
```

### Diamond

```json
{
  "type": "diamond",
  "x": 100, "y": 100,
  "width": 120, "height": 120,
  "strokeColor": "#3f3f46",
  "backgroundColor": "#27272a",
  "fillStyle": "solid",
  "strokeWidth": 2,
  "roughness": 1
}
```

### Text

```json
{
  "type": "text",
  "x": 120, "y": 120,
  "width": 160, "height": 25,
  "text": "Label Text",
  "fontSize": 20,
  "fontFamily": 1,
  "textAlign": "center",
  "verticalAlign": "middle",
  "strokeColor": "#fafafa",
  "backgroundColor": "transparent",
  "containerId": null,
  "originalText": "Label Text",
  "lineHeight": 1.25
}
```

**`fontFamily` values:**
- `1` — Virgil (Excalidraw's hand-drawn font, matches the aesthetic)
- `2` — Helvetica (clean sans-serif)
- `3` — Cascadia (monospace, for code)

Use `1` (Virgil) for labels and `3` (Cascadia) for code/technical text.

**Binding text to shapes:** Set `containerId` to the shape's `id`, and on the shape, set `boundElements: [{ "id": "text-id", "type": "text" }]`. This centres the text within the shape.

### Arrow

```json
{
  "type": "arrow",
  "x": 300, "y": 140,
  "width": 100, "height": 0,
  "points": [[0, 0], [100, 0]],
  "strokeColor": "#a1a1aa",
  "backgroundColor": "transparent",
  "strokeWidth": 2,
  "roughness": 1,
  "startBinding": {
    "elementId": "source-shape-id",
    "focus": 0,
    "gap": 5
  },
  "endBinding": {
    "elementId": "target-shape-id",
    "focus": 0,
    "gap": 5
  },
  "startArrowhead": null,
  "endArrowhead": "arrow"
}
```

**Arrowhead options:** `null` (none), `"arrow"` (standard), `"bar"` (flat end), `"dot"` (circle), `"triangle"` (filled).

**Binding:** Set `startBinding` / `endBinding` to attach arrows to shapes. The `focus` value (-1 to 1) controls which side of the shape the arrow connects to.

### Line

Same as arrow but `"type": "line"` and no arrowhead properties.

### Freedraw

```json
{
  "type": "freedraw",
  "x": 100, "y": 100,
  "points": [[0, 0], [5, 2], [10, 1], ...],
  "strokeColor": "#00d9c0",
  "backgroundColor": "transparent",
  "strokeWidth": 2,
  "roughness": 0
}
```

Use sparingly — for emphasis marks, underlines, or decorative elements.

## SVG Creation Guidelines

When creating the SVG file, follow these guidelines to match the app's visual style:

### SVG Structure

```svg
<svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 {width} {height}"
     style="max-width: {width}px; background-color: transparent;">
  <style>
    text { font-family: 'Plus Jakarta Sans', system-ui, sans-serif; }
    .label { fill: #fafafa; font-size: 16px; }
    .small { fill: #a1a1aa; font-size: 13px; }
    .accent { fill: #00d9c0; }
    .shape { fill: #27272a; stroke: #3f3f46; stroke-width: 2; }
    .shape-accent { fill: #27272a; stroke: #00d9c0; stroke-width: 2; }
    .connector { stroke: #a1a1aa; stroke-width: 2; fill: none; }
    .arrow { fill: #a1a1aa; }
  </style>

  <!-- Shapes, text, connectors -->
</svg>
```

### Hand-Drawn Aesthetic (Optional)

To mimic the Excalidraw hand-drawn look in SVG, add slight imperfections:
- Use `rx="4" ry="4"` on rectangles for rounded corners
- Add very slight path variations on connectors (not perfectly straight)
- Use a hand-drawn font if available, or the standard Plus Jakarta Sans

For this project, **clean SVGs with rounded corners are preferred** over exaggerated hand-drawn effects. The Excalidraw source provides the hand-drawn version; the SVG should be clean and readable.

### Dimensions

- **Width:** 600-1000px typical (the image sits within the content prose column)
- **Height:** varies by content, typically 300-600px
- **Padding:** 20-40px around content
- Use `viewBox` for scalability, `max-width` in style for reasonable sizing

### Arrowheads

Define a reusable arrowhead marker:

```svg
<defs>
  <marker id="arrowhead" markerWidth="10" markerHeight="7"
          refX="10" refY="3.5" orient="auto">
    <polygon points="0 0, 10 3.5, 0 7" class="arrow" />
  </marker>
</defs>

<line x1="100" y1="50" x2="300" y2="50"
      class="connector" marker-end="url(#arrowhead)" />
```

## Naming Convention

```
{concept}-{descriptor}.excalidraw
{concept}-{descriptor}.svg
```

Examples:
- `mindset-shift.excalidraw` / `mindset-shift.svg`
- `ai-workflow-overview.excalidraw` / `ai-workflow-overview.svg`
- `context-mental-model.excalidraw` / `context-mental-model.svg`

Use kebab-case. Keep names descriptive but concise.

## Design Principles

### 1. Visual Hierarchy
Use size, colour, and position to show what's most important:
- **Accent colour (`#00d9c0`)** for the key element or concept
- **Larger shapes/text** for primary concepts
- **Smaller, secondary-coloured text** for supporting details
- **Spatial position** — important things at top or centre

### 2. Whitespace
Don't crowd elements. Leave generous spacing between shapes. The dark background provides natural contrast — use it.

### 3. Alignment
Even in freeform drawings, align elements where it makes sense. Use consistent spacing between rows/columns of related items.

### 4. Limit Colours
Stick to the theme palette. Use accent colour sparingly — if everything is accented, nothing is.

### 5. Readable at a Glance
The diagram should communicate its main idea within 3 seconds. If it requires study to understand the basic concept, simplify.

### 6. One Concept Per Drawing
Like all visuals in the course, each drawing illustrates ONE concept. Don't combine unrelated ideas.

## Workflow

When asked to create an Excalidraw drawing:

1. **Understand the concept** — what should the learner take away?
2. **Plan the layout** — what shapes, where, what connections?
3. **Create the `.excalidraw` JSON** — full Excalidraw format with all elements, using the app theme colours.
4. **Create the `.svg`** — hand-crafted SVG matching the drawing, using the app theme and clean styling.
5. **Save both files** — to `content/module-{name}/images/`
6. **Output the markdown reference** — `![Alt text](/content/module-{name}/images/{name}.svg)`

## Reference Example

### Concept: "Mindset Shift" — two contrasting approaches

**`mindset-shift.excalidraw`** (abbreviated):
```json
{
  "type": "excalidraw",
  "version": 2,
  "source": "https://excalidraw.com",
  "elements": [
    {
      "id": "box-old",
      "type": "rectangle",
      "x": 40, "y": 60,
      "width": 280, "height": 200,
      "strokeColor": "#3f3f46",
      "backgroundColor": "#27272a",
      "fillStyle": "solid",
      "strokeWidth": 2,
      "roughness": 1,
      "roundness": { "type": 3 },
      "seed": 1001, "version": 1, "versionNonce": 1,
      "isDeleted": false, "groupIds": [], "boundElements": [
        { "id": "text-old-title", "type": "text" }
      ],
      "opacity": 100, "angle": 0, "updated": 1700000000000,
      "link": null, "locked": false
    },
    {
      "id": "text-old-title",
      "type": "text",
      "x": 100, "y": 80,
      "width": 160, "height": 25,
      "text": "Old Approach",
      "fontSize": 20,
      "fontFamily": 1,
      "textAlign": "center",
      "verticalAlign": "top",
      "strokeColor": "#a1a1aa",
      "backgroundColor": "transparent",
      "containerId": "box-old",
      "originalText": "Old Approach",
      "lineHeight": 1.25,
      "fillStyle": "solid", "strokeWidth": 1, "roughness": 1,
      "opacity": 100, "angle": 0, "seed": 1002, "version": 1,
      "versionNonce": 1, "isDeleted": false, "groupIds": [],
      "boundElements": null, "updated": 1700000000000,
      "link": null, "locked": false
    },
    {
      "id": "arrow-shift",
      "type": "arrow",
      "x": 340, "y": 160,
      "width": 120, "height": 0,
      "points": [[0, 0], [120, 0]],
      "strokeColor": "#00d9c0",
      "backgroundColor": "transparent",
      "strokeWidth": 3,
      "roughness": 1,
      "startBinding": { "elementId": "box-old", "focus": 0, "gap": 5 },
      "endBinding": { "elementId": "box-new", "focus": 0, "gap": 5 },
      "startArrowhead": null,
      "endArrowhead": "arrow",
      "fillStyle": "solid", "opacity": 100, "angle": 0,
      "seed": 1005, "version": 1, "versionNonce": 1,
      "isDeleted": false, "groupIds": [], "boundElements": null,
      "updated": 1700000000000, "link": null, "locked": false
    },
    {
      "id": "box-new",
      "type": "rectangle",
      "x": 480, "y": 60,
      "width": 280, "height": 200,
      "strokeColor": "#00d9c0",
      "backgroundColor": "#27272a",
      "fillStyle": "solid",
      "strokeWidth": 2,
      "roughness": 1,
      "roundness": { "type": 3 },
      "seed": 1003, "version": 1, "versionNonce": 1,
      "isDeleted": false, "groupIds": [], "boundElements": [
        { "id": "text-new-title", "type": "text" }
      ],
      "opacity": 100, "angle": 0, "updated": 1700000000000,
      "link": null, "locked": false
    },
    {
      "id": "text-new-title",
      "type": "text",
      "x": 540, "y": 80,
      "width": 160, "height": 25,
      "text": "New Approach",
      "fontSize": 20,
      "fontFamily": 1,
      "textAlign": "center",
      "verticalAlign": "top",
      "strokeColor": "#fafafa",
      "backgroundColor": "transparent",
      "containerId": "box-new",
      "originalText": "New Approach",
      "lineHeight": 1.25,
      "fillStyle": "solid", "strokeWidth": 1, "roughness": 1,
      "opacity": 100, "angle": 0, "seed": 1004, "version": 1,
      "versionNonce": 1, "isDeleted": false, "groupIds": [],
      "boundElements": null, "updated": 1700000000000,
      "link": null, "locked": false
    }
  ],
  "appState": {
    "viewBackgroundColor": "#18181b",
    "gridSize": null
  }
}
```

**Corresponding `mindset-shift.svg`:**
```svg
<svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 800 320"
     style="max-width: 800px; background-color: transparent;">
  <style>
    text { font-family: 'Plus Jakarta Sans', system-ui, sans-serif; }
    .title { font-size: 20px; font-weight: 600; }
    .item { font-size: 15px; }
    .dim { fill: #a1a1aa; }
    .bright { fill: #fafafa; }
    .accent { fill: #00d9c0; }
  </style>
  <defs>
    <marker id="arrowhead" markerWidth="10" markerHeight="7"
            refX="10" refY="3.5" orient="auto">
      <polygon points="0 0, 10 3.5, 0 7" fill="#00d9c0" />
    </marker>
  </defs>

  <!-- Old approach box -->
  <rect x="40" y="60" width="280" height="200" rx="8" ry="8"
        fill="#27272a" stroke="#3f3f46" stroke-width="2" />
  <text x="180" y="95" text-anchor="middle" class="title dim">Old Approach</text>
  <text x="70" y="135" class="item dim">Step-by-step instructions</text>
  <text x="70" y="165" class="item dim">Manual review</text>
  <text x="70" y="195" class="item dim">One task at a time</text>

  <!-- Arrow -->
  <line x1="340" y1="160" x2="460" y2="160"
        stroke="#00d9c0" stroke-width="3" marker-end="url(#arrowhead)" />

  <!-- New approach box (accent border) -->
  <rect x="480" y="60" width="280" height="200" rx="8" ry="8"
        fill="#27272a" stroke="#00d9c0" stroke-width="2" />
  <text x="620" y="95" text-anchor="middle" class="title bright">New Approach</text>
  <text x="510" y="135" class="item bright">Delegate the outcome</text>
  <text x="510" y="165" class="item bright">AI handles details</text>
  <text x="510" y="195" class="item bright">Parallel workstreams</text>
</svg>
```

## Common Mistakes to Avoid

- **Don't use bright backgrounds** — stick to zinc-900/800, not white
- **Don't forget the `.excalidraw` source** — always create both files
- **Don't use too many accent-coloured elements** — accent is for emphasis, not everything
- **Don't overcrowd** — whitespace is a feature, not wasted space
- **Don't make the SVG too small** — 600px minimum width for readability
- **Don't use Excalidraw for structured charts** — use `/mermaid` for flowcharts, sequences, etc.
