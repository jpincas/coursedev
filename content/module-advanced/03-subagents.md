---
title: "Subagents: AI Spawns Helpers"
duration: "15m"
tags: [subagents, parallel, architecture]
---

# Subagents

When one AI agent isn't enough, it can spawn more to work in parallel.

## How Subagents Work

![Subagent Architecture: Main Agent delegates to parallel Subagents](/content/module-advanced/images/subagents.svg)

1. You give one complex task
2. The main agent breaks it into pieces
3. Each piece goes to a subagent
4. They work simultaneously
5. Results come back and get combined

## Why This Matters

**Speed**
Multiple research threads run at once. What would take one agent an hour takes the team minutes.

**Thoroughness**
Different angles explored simultaneously. Nothing waiting in a queue.

**Specialisation**
Each subagent can focus on its piece. The main agent handles coordination.

## Good Use Cases

**Research tasks**
Research competitors, analyse the market, review academic literature — all at once.

**Multi-faceted analysis**
Financial analysis, operational review, customer feedback — parallel streams.

**Content creation**
Research, outlining, drafting different sections — simultaneous.

**Complex reports**
Each section tackled by a different subagent, main agent assembles.

```callout
type: info
title: "Behind the Scenes"
content: "You don't usually need to manage subagents directly. The main agent decides when to parallelise and handles coordination. You just see the task complete faster."
```

## The Mental Model

Think of it like this:
- **Without subagents**: One person doing everything sequentially
- **With subagents**: A team tackling different aspects simultaneously

You delegate to one "manager" agent. It staffs the project appropriately.

## Real-World Patterns

The autonomous agents market reached **$4.35 billion** in 2025, projected to exceed **$100 billion** by 2034. Practitioners have identified specific patterns that work.

**ReAct (Reasoning + Action)**
Interleave reasoning with tool calls. Agent thinks, acts, observes result, thinks again.

**Plan-then-Execute**
Large model does planning, smaller model does execution. Reduces cost while maintaining quality.

**Hierarchical Task Decomposition**
Manager agent breaks tasks into subtasks for specialist agents. Like a project manager distributing work.

**Generator-Evaluator Loop**
One agent generates solutions, another evaluates and suggests improvements. Continuous refinement.

## The Meta Case Study

This is where theory meets reality.

### This Course Built Itself

Every concept you've learned in this course — context engineering, delegation, skills, subagents, files, verification — was used to build the course you are sitting in.

![Course creation in action](/static/images/meta-course-creation-screenshot.png)

## The Multi-Agent Hierarchy

Here's how the work was structured:

![Multi-agent hierarchy: User delegates to course-director, which spawns parallel module-builders, which invoke skills](/content/module-advanced/images/meta-hierarchy.svg)

## The Creation Process

**Phase 1: Deep Research**
The instructor wrote an 8,000-word research report synthesising findings from Anthropic, OpenAI, Google, McKinsey, Bain, Deloitte, academic studies, and practitioner blogs.

This is **context engineering** in practice. The quality of downstream work depends on the quality of the context you provide.

**Phase 2: Course-Director Agent**
A subagent was launched with a detailed brief: read the research report, read every existing module, compare coverage against needs, and produce a comprehensive update plan.

The agent:
1. Read `ai-training-report.md` (8,000 words)
2. Read `plan.md` and every module file in `content/`
3. Cross-referenced the course against report findings
4. Identified gaps, outdated content, and missing concepts
5. Produced a prioritised 720-line update plan in `course-update-plan.md`

This is **Hierarchical Task Decomposition**. The user delegates to a "manager" agent that does analysis and planning.

## Parallel Execution

**Phase 3: Module-Builder Subagents**
The course-director can launch independent module-builder agents — one per module, running **in parallel**.

Each module-builder:
- Creates the directory structure (`content/module-{name}/`)
- Writes `_module.yaml` metadata
- Creates numbered markdown pages (`01-name.md`, `02-name.md`)
- Embeds interactive blocks (quizzes, callouts, diagrams)

Multiple modules can be built simultaneously. What would take one agent hours takes a team minutes.

**Phase 4: Skill Composition**
Module-builders don't work alone. They invoke **skills** for specialised work:

- `/quiz` — generates quiz blocks testing specific concepts
- `/agent-demo` — creates scripted AI conversation walkthroughs
- `/mermaid` — produces structured diagrams (flowcharts, sequences)
- `/excalidraw` — creates freeform drawings (sketches, concept maps)
- `/image-generator` — generates infographic-style visuals

This is the "tools that make tools" pattern. Agents use skills to handle specific content types.

## Human Oversight

**Phase 5: Human-in-the-Loop**
At every level, the human reviews and directs:

- The research report was human-authored
- The course-director's plan was reviewed before module-builders launched
- Agent output was verified before inclusion
- The human maintains editorial control

This demonstrates the **"AI-first, human-verified"** theme taught throughout the course.

### What This Demonstrates

Every major concept from the course appears in this single real example:

- **Multi-agent hierarchies** — user → course-director → module-builders
- **Skill composition** — agents invoking specialised skills on demand
- **Parallel agents** — multiple module-builders running simultaneously
- **Deep research workflows** — report → gap analysis → prioritised plan
- **Context engineering** — the research report IS the context that makes agent work high-quality
- **Persistent instructions** — CLAUDE.md and agent instruction files defining agent behaviour
- **File-based workflows** — everything is files (report, plan, module YAML, markdown pages)
- **Preparation-is-value** — the quality of research determines quality of everything downstream
- **Verification** — human reviews all agent output before accepting it

```callout
type: info
title: "The Self-Referential Moment"
content: "Every concept in this course — context engineering, delegation, skills, subagents, files, verification — was used to build the course you are sitting in. This is not a theoretical exercise. This is how work gets done in February 2026."
```

```quiz
id: subagents-benefit
type: multiple-choice
question: "What's the primary benefit of AI using subagents?"
options:
  - "Subagents are more accurate than single agents"
  - "Complex tasks get tackled in parallel, completing faster and more thoroughly"
  - "Subagents are cheaper to run"
  - "Subagents have access to more tools"
answer: 1
explanation: "Subagents allow parallel work — multiple aspects of a complex task being tackled simultaneously. This means faster completion and more thorough coverage than sequential single-agent work."
```

```quiz
id: meta-case-study-concepts
type: multiple-choice
question: "In the course creation case study, what role did the 8,000-word research report play?"
options:
  - "It was the final deliverable"
  - "It provided the context that shaped all downstream agent work"
  - "It was used to train the AI models"
  - "It replaced the need for human oversight"
answer: 1
explanation: "The research report was context engineering in action. Its quality determined the quality of the course-director's analysis, which determined the quality of the module-builders' output. Context is everything."
```
