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

Sometimes the AI orchestrates this automatically. But the real power comes when **you define the subagent templates** — instruction files that tell each subagent exactly how to behave, what tools to use, and what output to produce. You're not just delegating; you're designing the team.

```callout
type: info
title: "A Note on Naming"
content: "\"Subagents\" is Claude's terminology, but the pattern is universal. OpenAI calls them \"sub-tasks\", Google uses \"agent delegation\", and frameworks like LangGraph and CrewAI implement the same idea. The concept — a parent agent spawning specialised child agents — is converging across every major platform."
```

## Why This Matters

**Speed**
Multiple research threads run at once. What would take one agent an hour takes the team minutes.

**Thoroughness**
Different angles explored simultaneously. Nothing waiting in a queue.

**Context isolation**
Each subagent gets its own context window. The module-builder working on "data analysis" doesn't see the "prompting techniques" content. This prevents cross-contamination and keeps each agent focused on its specific task.

**Different models for different jobs**
Not every subtask needs the most expensive model. Your orchestrator might use the best model for planning, while subagents use faster, cheaper models for straightforward execution. A plan-then-execute pattern: the big model thinks, the small models do.

**Specialisation**
Each subagent can have different instructions, different tools, and different system prompts. One might be a researcher, another a writer, another a reviewer.

## Defining Subagent Templates

You don't have to build subagent templates by hand — describe what you need and let AI create them. But you should know what goes into one:

**Instructions** — What this subagent does, how it should behave, what conventions to follow.

**Tools** — Which capabilities the subagent has access to (file writing, code execution, image generation, API calls).

**Constraints** — What the subagent should NOT do. Boundaries keep agents focused.

**Output format** — Where to write results and in what structure.

A template is just a file. When the orchestrator needs that type of agent, it loads the template and launches a new instance. You create the template once; it gets reused every time that type of work is needed.

```callout
type: tip
title: "Let AI Build Your Templates"
content: "Just like skills, you don't need to write subagent templates from scratch. Describe the role — 'I need an agent that reviews documents for compliance issues and produces a checklist' — and let AI create the template file for you."
```

## Good Use Cases

**Research tasks**
Research competitors, analyse the market, review academic literature — all at once.

**Multi-faceted analysis**
Financial analysis, operational review, customer feedback — parallel streams.

**Content creation**
Research, outlining, drafting different sections — simultaneous.

**Complex reports**
Each section tackled by a different subagent, main agent assembles.

## Real-World Patterns

**ReAct (Reasoning + Action)**
Interleave reasoning with tool calls. Agent thinks, acts, observes result, thinks again.

**Plan-then-Execute**
Large model does planning, smaller model does execution. Reduces cost while maintaining quality.

**Hierarchical Task Decomposition**
Manager agent breaks tasks into subtasks for specialist agents. Like a project manager distributing work.

**Generator-Evaluator Loop**
One agent generates solutions, another evaluates and suggests improvements. Continuous refinement.

## Subagents in Action

Watch an orchestrator agent plan a training course and delegate to parallel subagents — each building a different module in its own folder simultaneously. This is directly modelled on how this course was built.

```agent
id: subagent-course-demo
title: "Course Building with Parallel Subagents"
model_label: "Course Director"

system: |
  You are a course-director agent. You design training courses by:
  1. Analysing a research brief
  2. Creating a structured course plan with model assignments
  3. Delegating module creation to independent subagent builders
  Each module-builder works in its own folder with its own context.

scratchpad:
  "research-brief.md": |
    # Research Brief: AI for Data Analysis

    ## Key Findings
    - 73% of analysts spend most time on data cleaning — AI automates 80% of this
    - Gap is methodology not tools: analysts need frameworks for working with AI
    - Three skill tiers identified: foundations, techniques, workflows
    - Practical exercises dramatically improve retention vs theory-only

    ## Target Audience
    Mid-career data analysts. Familiar with Excel and SQL. No AI experience.

    ## Scope
    3 modules, 45 min each. Practical focus — every concept needs an exercise.

tools:
  - scratchpad_read
  - scratchpad_write
  - create_folder
  - list_files

visibility:
  system_prompt: hidden
  tool_calls: visible
  full_context: hidden

sidebar:
  width: "50%"
  start_open: true

script:
  - type: note
    text: "This demonstrates the subagent pattern. An orchestrator reads a brief, creates a plan, then delegates to independent subagents — each building a different module in its own isolated folder. This is how this very course was built."

  - type: user
    content: "I need a 3-module training course on 'AI for Data Analysis'. Here's the research brief. Create a course plan and then build all three modules."

  - type: tool_call
    tool: scratchpad_read
    args:
      filename: "research-brief.md"

  - type: tool_result
    tool: scratchpad_read
    content: |
      # Research Brief: AI for Data Analysis

      ## Key Findings
      - 73% of analysts spend most time on data cleaning — AI automates 80% of this
      - Gap is methodology not tools: analysts need frameworks for working with AI
      - Three skill tiers identified: foundations, techniques, workflows
      - Practical exercises dramatically improve retention vs theory-only

      ## Target Audience
      Mid-career data analysts. Familiar with Excel and SQL. No AI experience.

      ## Scope
      3 modules, 45 min each. Practical focus — every concept needs an exercise.

  - type: tool_call
    tool: scratchpad_write
    args:
      filename: "course-plan.md"
      content: |
        # Course Plan: AI for Data Analysis

        ## Module 1: Data Foundations (45 min)
        Brief: What AI sees when it reads data. Types it handles well
        vs poorly. Preparing data for AI analysis.
        Exercise: Clean a messy dataset with AI.
        Assigned model: Haiku (fast, cheap — straightforward content)

        ## Module 2: Analysis Techniques (45 min)
        Brief: Exploratory analysis, pattern recognition, statistical
        analysis without code.
        Exercise: Analyse sales data to find insights.
        Assigned model: Sonnet (capable — needs nuanced examples)

        ## Module 3: Automated Workflows (45 min)
        Brief: Repeatable pipelines, connecting data sources,
        scheduling automated reports.
        Exercise: Build an automated weekly analysis.
        Assigned model: Opus (best — complex integration patterns)

  - type: tool_result
    tool: scratchpad_write
    content: "Written to course-plan.md (20 lines)"

  - type: note
    text: "The orchestrator has analysed the brief and produced a plan — including which AI model each subagent should use. Cheaper models for simple content, expensive models for complex material. Now it creates the workspace and delegates."

  - type: tool_call
    tool: create_folder
    args:
      path: "course/01-data-foundations"

  - type: tool_result
    tool: create_folder
    content: "Created folder: course/01-data-foundations"

  - type: tool_call
    tool: create_folder
    args:
      path: "course/02-analysis-techniques"

  - type: tool_result
    tool: create_folder
    content: "Created folder: course/02-analysis-techniques"

  - type: tool_call
    tool: create_folder
    args:
      path: "course/03-automated-workflows"

  - type: tool_result
    tool: create_folder
    content: "Created folder: course/03-automated-workflows"

  - type: note
    text: "Three subagents now launch in parallel. Each gets its own folder, its own context window, and only the plan section relevant to its module. This is context isolation — the module-1 builder doesn't see module-3 content, so it stays focused and its context window isn't wasted on irrelevant material."

  - type: tool_call
    tool: scratchpad_write
    args:
      filename: "course/01-data-foundations/01-what-ai-sees.md"
      content: |
        ---
        title: "What AI Sees in Your Data"
        duration: "15m"
        ---

        # What AI Sees in Your Data

        When you open a spreadsheet, you see rows and patterns
        you've learned to recognise. When AI reads the same file,
        it processes every cell but understands none of the context
        you take for granted.

        ## Structure vs Meaning

        AI excels at structured data — clean columns, consistent
        formats. It struggles with the messy reality: merged cells,
        inconsistent dates, notes jammed into number fields.

        ## The Preparation Principle

        The single most important step before AI analysis: clean
        your data. Not because AI can't handle mess — but because
        clean input produces dramatically better output.

  - type: tool_result
    tool: scratchpad_write
    content: "Written to course/01-data-foundations/01-what-ai-sees.md (22 lines)"

  - type: tool_call
    tool: scratchpad_write
    args:
      filename: "course/02-analysis-techniques/01-finding-patterns.md"
      content: |
        ---
        title: "Finding Patterns AI Can See"
        duration: "15m"
        ---

        # Finding Patterns AI Can See

        The most powerful analysis technique isn't asking "what do
        you see?" — it's telling AI what to look for.

        ## Directed Exploration

        Start with a hypothesis: "I think sales drop on Wednesdays."
        AI validates or refutes this in seconds across millions of
        rows. Undirected prompts produce generic observations.

        ## Anomaly Detection

        Give AI your baseline: "Sales are usually 10-15k per day."
        It flags every day outside that range with context about
        what else happened.

  - type: tool_result
    tool: scratchpad_write
    content: "Written to course/02-analysis-techniques/01-finding-patterns.md (20 lines)"

  - type: tool_call
    tool: scratchpad_write
    args:
      filename: "course/03-automated-workflows/01-from-analysis-to-automation.md"
      content: |
        ---
        title: "From Analysis to Automation"
        duration: "15m"
        ---

        # From Analysis to Automation

        The first time you analyse data with AI, it takes an hour.
        The second time, five minutes. The third time, it should
        happen automatically.

        ## The Automation Ladder

        1. Manual — you do the analysis with AI help
        2. Repeatable — you've saved the prompts and process
        3. Scheduled — it runs weekly without you
        4. Monitored — it alerts you when results are unusual

        ## Building Your First Pipeline

        Start with your most repeated analysis. Document the steps
        as a skill file. Add a script to fetch the data. Schedule it.

  - type: tool_result
    tool: scratchpad_write
    content: "Written to course/03-automated-workflows/01-from-analysis-to-automation.md (21 lines)"

  - type: tool_call
    tool: list_files
    args: {}

  - type: tool_result
    tool: list_files
    content: |
      course-plan.md
      research-brief.md
      course/
        01-data-foundations/
          01-what-ai-sees.md
        02-analysis-techniques/
          01-finding-patterns.md
        03-automated-workflows/
          01-from-analysis-to-automation.md

  - type: note
    text: "All three modules are now built. In a real system these run simultaneously — what takes one agent 3 hours takes three subagents 1 hour. Notice the folder structure in the explorer: each subagent wrote exclusively to its own directory. No risk of cross-contamination."

  - type: note
    text: "The model assignments matter. Module 1 (basic concepts) used Haiku — fast and cheap. Module 3 (complex integration) used Opus — the best model. Matching model capability to task complexity can cut costs by 60-80% without sacrificing quality where it matters."

  - type: note
    text: "Now watch the orchestrator come back. The subagents are done — the course-director reads their output and assembles the final course overview. This is the full pattern: plan → delegate → assemble."

  - type: tool_call
    tool: scratchpad_read
    args:
      filename: "course/01-data-foundations/01-what-ai-sees.md"

  - type: tool_result
    tool: scratchpad_read
    content: |
      ---
      title: "What AI Sees in Your Data"
      duration: "15m"
      ---

      # What AI Sees in Your Data

      When you open a spreadsheet, you see rows and patterns
      you've learned to recognise. When AI reads the same file,
      it processes every cell but understands none of the context
      you take for granted.

  - type: tool_call
    tool: scratchpad_read
    args:
      filename: "course/02-analysis-techniques/01-finding-patterns.md"

  - type: tool_result
    tool: scratchpad_read
    content: |
      ---
      title: "Finding Patterns AI Can See"
      duration: "15m"
      ---

      # Finding Patterns AI Can See

      The most powerful analysis technique isn't asking "what do
      you see?" — it's telling AI what to look for.

  - type: tool_call
    tool: scratchpad_read
    args:
      filename: "course/03-automated-workflows/01-from-analysis-to-automation.md"

  - type: tool_result
    tool: scratchpad_read
    content: |
      ---
      title: "From Analysis to Automation"
      duration: "15m"
      ---

      # From Analysis to Automation

      The first time you analyse data with AI, it takes an hour.
      The second time, five minutes. The third time, it should
      happen automatically.

  - type: tool_call
    tool: scratchpad_write
    args:
      filename: "course/README.md"
      content: |
        # AI for Data Analysis — Course Overview

        A 3-module practical course for data analysts.

        ## Modules

        ### 1. Data Foundations
        What AI sees when it reads your data. Structured vs messy,
        preparation principles, and your first AI-assisted cleanup.

        ### 2. Analysis Techniques
        Directed exploration, pattern recognition, and anomaly
        detection. Statistical analysis without writing code.

        ### 3. Automated Workflows
        From manual analysis to scheduled pipelines. Build
        repeatable, monitored data workflows.

        ## Prerequisites
        Familiarity with Excel and basic SQL. No AI experience needed.

        ## Duration
        ~2.5 hours total (45 min per module + exercises)

  - type: tool_result
    tool: scratchpad_write
    content: "Written to course/README.md (22 lines)"

  - type: assistant
    content: "Course complete. Three modules built in parallel, assembled into a final overview. The README pulls together the key themes from each module into a coherent programme.\n\nYour course is ready for review in the course/ folder."

  - type: note
    text: "That last step is the orchestrator pattern in full. The course-director read back each module's output, then wrote a README that synthesises them into a coherent whole. Plan → delegate → assemble. The subagents never see each other's work — only the orchestrator has the full picture."

  - type: note
    text: "This is exactly how this training course was built. A course-director agent analysed a research report, created a plan, then launched parallel module-builder subagents — each with its own instructions, its own folder, and access to skills like /quiz and /mermaid for specialised content. The orchestrator assembled the final result."
```

## This Course Built Itself

Every concept you've learned in this course — context engineering, delegation, skills, subagents, files, verification — was used to build the course you are sitting in.

![Course creation in action](/static/images/meta-course-creation-screenshot.png)

### The Creation Process

**Phase 1: Deep Research**
The instructor used Claude's Deep Research to produce an 8,000-word report synthesising findings from Anthropic, OpenAI, Google, McKinsey, Bain, Deloitte, academic studies, and practitioner blogs. This is **context engineering** in practice — the quality of everything downstream depends on the quality of the context you start with.

**Phase 2: Course-Director Agent**
A subagent was launched with a detailed brief: read the research report, read every existing module, compare coverage against needs, and produce a comprehensive update plan.

The agent read 8,000 words of research, cross-referenced it against the existing course, identified gaps, and produced a prioritised 720-line update plan.

This is **Hierarchical Task Decomposition** — the user delegates to a "manager" agent that does analysis and planning.

**Phase 3: Parallel Module-Builders**
The course-director launched independent module-builder agents — one per module, running in parallel. Each module-builder created its directory, wrote markdown pages, and invoked skills for specialised content: `/quiz` for comprehension checks, `/mermaid` for diagrams, `/agent-demo` for interactive walkthroughs.

**Phase 4: Human-in-the-Loop**
At every level, the human reviewed and directed. The research report was curated. The plan was reviewed before modules launched. Agent output was verified before inclusion. AI-first, human-verified.

### The Hierarchy

![Multi-agent hierarchy: User delegates to course-director, which spawns parallel module-builders, each using skills](/content/module-advanced/images/meta-hierarchy.svg)

### What This Demonstrates

- **Multi-agent hierarchies** — user → course-director → module-builders
- **Skill composition** — agents invoking specialised skills on demand
- **Context isolation** — each module-builder works in its own directory
- **Model selection** — different tasks, different models
- **Deep research as context** — quality of research determines quality of everything downstream
- **File-based workflows** — everything is files (report, plan, YAML, markdown)
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
id: subagent-context-isolation
type: multiple-choice
question: "Why do subagents each get their own context window?"
options:
  - "To increase the total amount of text the system can process"
  - "Each agent stays focused on its task without irrelevant material filling its context"
  - "It's a technical requirement of the AI platform"
  - "To prevent agents from communicating with each other"
answer: 1
explanation: "Context isolation keeps each subagent focused. A module-builder working on 'data analysis' doesn't need to see 'prompting techniques' content — that would waste context space and potentially cause confusion. Each agent gets exactly the context it needs."
```
