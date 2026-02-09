---
name: image-generator
description: Generate infographic images and visual assets for training course modules using OpenAI's image generation API. Use this when module content needs conceptual visualizations, process graphics, comparison visuals, or illustrative infographics. Requires OPENAI_API_KEY environment variable.
---

# Training Course Image Generator

This skill generates infographic-style images for training course modules using OpenAI's image generation API.

## When to Use

- Creating conceptual visualizations for teaching AI concepts
- Producing infographic-style process or comparison graphics
- Generating illustrative visuals that complement narrative content
- Any time a module page needs a richer visual than a Mermaid or Excalidraw diagram can provide

**When NOT to use** (use these skills instead):
- Structured diagrams with nodes/edges — use `/mermaid`
- Freeform sketches and whiteboard-style drawings — use `/excalidraw`

## Requirements

The `OPENAI_API_KEY` environment variable must be set:

```bash
export OPENAI_API_KEY="sk-..."
```

## Python Dependencies

```bash
pip install openai requests pillow
```

## Usage

When generating images for modules:

1. **Understand the teaching context** — Read the page content to understand what concept the image should reinforce
2. **Craft an infographic prompt** — Create a detailed description specifying the visual style, content elements, and layout
3. **Generate the image** — Use the helper script to create the image
4. **Save to the module's images directory** — Store in `content/module-{name}/images/` with a descriptive filename
5. **Reference from markdown** — Use standard image syntax: `![Alt text](/content/module-{name}/images/filename.png)`

## Image Guidelines

- **Style**: Clean infographic aesthetic — flat design, clear typography, structured layout. NOT photographic.
- **Palette**: Dark background (`#18181b` zinc-900) with accent colour `#00d9c0` (teal). Use white/light text. This matches the app's dark theme.
- **Content**: Conceptual visualizations, process flows, comparison graphics, data-style layouts, labelled diagrams
- **Quality**: High-resolution, suitable for full-width display (1536x1024 landscape is default)
- **Text in images**: Keep text minimal and large enough to read. Labels and short phrases, not paragraphs.
- **Simplicity**: One concept per image. 3-7 visual elements is ideal.

## gpt-image-1.5 Parameters

- **Size options**: 1024x1024, 1024x1536 (portrait), 1536x1024 (landscape), auto
- **Quality options**: low, medium, high, auto
- **Default**: 1536x1024 at high quality
- **Response format**: Images are returned as base64 data (not URLs)

## Example Prompts

For a module about context windows:
```
"Clean infographic on dark background (#18181b) showing a context window as a container being filled with different types of content: system prompt (teal #00d9c0), conversation history (grey), user instructions (white), attached files (light blue). Flat design, minimal text labels, modern data visualization style."
```

For a module about AI delegation patterns:
```
"Infographic on dark background (#18181b) comparing two workflows side by side: left shows a person doing 5 sequential tasks manually, right shows a person delegating to 3 AI agents working in parallel. Flat design, teal accent (#00d9c0), clean arrows and icons, minimal labels."
```

For a module about prompt engineering:
```
"Clean infographic on dark background (#18181b) showing the anatomy of an effective prompt: role, context, task, constraints, output format — each as a labelled layer in a stacked diagram. Flat design, teal (#00d9c0) highlights, white text, modern minimal style."
```

## Helper Script

The skill includes `scripts/generate_image.py` which:
- Takes a text prompt as input
- Generates an image using the OpenAI API
- Downloads and saves the image locally
- Returns the file path

## Example Usage in Workflow

```bash
# Generate infographic for a module page
python .claude/skills/image-generator/scripts/generate_image.py \
  "Clean infographic on dark background showing context window layers..." \
  --output content/module-context/images/context-layers.png \
  --size 1536x1024 \
  --quality high
```
