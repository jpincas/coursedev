# Meta Example: Include in Course Plan

## The Idea

Include a self-referential case study in the course: the actual process of creating/updating this course should be featured AS content in the course itself.

## What Happened

1. The instructor wrote a deep research report on AI training needs (`ai-training-report.md`)
2. Used Claude Code to launch a **course-director subagent**, which reviews the report and existing course, then produces a comprehensive update plan
3. The course-director can launch **module-builder subagents** (one per module, running in parallel)
4. Module-builders use **skills** (`/quiz`, `/agent-demo`, `/mermaid`, `/excalidraw`, `/image-generator`) for interactive content creation

## What This Demonstrates

- Multi-agent hierarchies (course-director → module-builders)
- Skill composition (agents invoking specialized skills)
- Background agents running in parallel
- Deep research workflows feeding into structured planning
- Human-in-the-loop oversight of autonomous agent work

## Screenshot

`static/images/meta-course-creation-screenshot.png` — shows the user instructing Claude Code to launch the course-director agent to review the training report and plan course updates.

## Where to Include

This should be a standout case study in whichever module covers agent architectures, multi-agent systems, or practical AI workflows. It's an "eating our own dog food" moment — attendees see the very tool they're learning about being used to build the course they're attending.
