---
title: "Skills: Reusable Procedures"
duration: "15m"
tags: [skills, procedures, automation]
---

# Skills

Skills are documented procedures that AI can follow. Think of them as runbooks for AI.

## What Is a Skill?

A skill bundles together everything needed to do a task your way:
- Instructions (how to do it)
- Templates (what format to use)
- Examples (what good output looks like)

When you ask AI to do something, it looks at your available skills. If there's a relevant one, it follows those procedures automatically.

## Skill Structure

A typical skills folder might look like:

```
Skills/
├── weekly-report/
│   ├── instructions.md
│   └── template.docx
├── expense-summary/
│   ├── instructions.md
│   └── format.xlsx
└── client-proposal/
    ├── instructions.md
    └── template.pptx
```

Each skill is a folder containing its instructions and supporting files.

## Anatomy of a Skill

A skill's instructions typically include:

**Purpose** — What is this skill for?
"Generate the Friday leadership update."

**Steps** — What should happen?
1. Pull this week's completed tasks
2. Identify blockers and risks
3. List next week's priorities
4. Format using the template

**Tone** — How should it sound?
"Professional but concise. No jargon."

**Output** — What exactly should be delivered?
"1-page Word document following the template."

## Why Skills Matter

**Write once, use forever**
Document the procedure once. Every future task of that type follows it automatically.

**Consistent quality**
No variation from forgetting steps or misremembering preferences.

**Capture institutional knowledge**
Your best practices become AI capabilities. New team members benefit immediately.

**No re-explaining**
"Generate the weekly report" works — AI knows what that means for your organisation.

```callout
type: tip
title: "Building Your Skill Library"
content: "Start with tasks you do repeatedly. Each time you explain something to AI, ask: should this be a skill? If you'll do it again, the answer is probably yes."
```

## Good Candidates for Skills

- Regular reports (weekly, monthly, quarterly)
- Document types (proposals, summaries, briefs)
- Data processing tasks (cleaning, formatting, analysis)
- Communication templates (emails, announcements)
- Review processes (code review, document review)

Any task with consistent requirements is a candidate.

## Meta-Tooling: Tools That Make Tools

The most powerful pattern is using AI to create custom tools you then use repeatedly.

**The concept:**
Instead of manually doing a repetitive task or learning to code a solution, describe the task to AI once. AI creates a working script. You deploy it. Now it's a tool.

**Examples in practice:**

**Google Workspace Studio** — Build agents for Gmail, Drive, Calendar using plain English. "Summarise meeting notes and create calendar reminders for action items."

**Claude Code Skills** — Reusable capability files that Claude loads on demand. Create a skill for running your test suite, formatting code, or deploying to staging.

**MindStudio** — No-code agent builder. Deploy as web apps, browser extensions, email triggers, or scheduled automations.

**The workflow:**
1. Identify a repetitive task
2. Describe it to AI — get a working script
3. Test and iterate
4. Deploy as a reusable tool
5. Return to AI for maintenance when needs change

This is the highest-leverage power user pattern. You're not just using AI — you're using AI to build the tools you'll use next.

```callout
type: tip
title: "Start Small, Compound Fast"
content: "Your first custom tool might save 10 minutes per week. Your tenth might save 2 hours. The compound effect of building tools is exponential — each tool you create makes the next one easier to build and more valuable to deploy."
```

## Creating a Skill

Watch the process of turning a repeating task into a reusable skill. This is the meta-tooling pattern: using AI to create the tool you'll use repeatedly.

```agent
id: skill-creation-demo
title: "Building a Reusable Skill"
model_label: "Claude"

system: |
  You are a productivity consultant helping create reusable AI skills.
  Skills should be clear, complete, and actionable.

scratchpad:
  "weekly-report-example.md": |
    # Leadership Update — Week of 6 Jan 2026

    ## Summary
    Strong week. Two major features shipped. Pipeline healthy.

    ## Shipped This Week
    - User authentication v2 (3-week project, completed on schedule)
    - Dashboard performance improvements (page load: 1.2s → 0.4s)

    ## In Progress
    - Mobile redesign: 60% complete, on track for Feb 1
    - API v3: Design review complete, build starting Monday

    ## Blockers & Risks
    - Design team capacity: 1 designer on leave until Jan 20
    - QA backlog growing: 12 tickets in queue (target: under 5)

    ## Key Metrics
    - Sprint velocity: 42 points (target: 40) ✓
    - Bug escape rate: 2.1% (target: under 3%) ✓
    - Deploy frequency: 4 deploys this week

    ## Next Week Priorities
    1. Start API v3 build
    2. Clear QA backlog
    3. Prepare Feb milestone review deck

tools:
  - scratchpad_read
  - scratchpad_write

visibility:
  system_prompt: hidden
  tool_calls: visible
  full_context: hidden

sidebar:
  width: "45%"
  start_open: true

script:
  - type: note
    text: "This person writes a weekly leadership report every Friday. It takes 45 minutes. They want to turn it into a skill so AI can draft it in 5 minutes. Watch the skill creation process."

  - type: user
    content: "I write this leadership report every Friday and it takes forever. Read my example report, then create a skill file — a reusable instruction document that I can use every week so AI produces this format automatically."

  - type: tool_call
    tool: scratchpad_read
    args:
      filename: "weekly-report-example.md"

  - type: tool_result
    tool: scratchpad_read
    content: |
      # Leadership Update — Week of 6 Jan 2026

      ## Summary
      Strong week. Two major features shipped. Pipeline healthy.

      ## Shipped This Week
      - User authentication v2 (3-week project, completed on schedule)
      - Dashboard performance improvements (page load: 1.2s → 0.4s)

      ## In Progress
      - Mobile redesign: 60% complete, on track for Feb 1
      - API v3: Design review complete, build starting Monday

      ## Blockers & Risks
      - Design team capacity: 1 designer on leave until Jan 20
      - QA backlog growing: 12 tickets in queue (target: under 5)

      ## Key Metrics
      - Sprint velocity: 42 points (target: 40) ✓
      - Bug escape rate: 2.1% (target: under 3%) ✓
      - Deploy frequency: 4 deploys this week

      ## Next Week Priorities
      1. Start API v3 build
      2. Clear QA backlog
      3. Prepare Feb milestone review deck

  - type: tool_call
    tool: scratchpad_write
    args:
      filename: "skill-weekly-report.md"
      content: |
        # Skill: Weekly Leadership Report

        ## Purpose
        Generate the Friday leadership update for the engineering team.

        ## When to Use
        Every Friday afternoon. Takes raw inputs and produces a formatted report.

        ## Required Inputs
        Before running this skill, gather:
        - This week's completed work (from Jira/Linear or standup notes)
        - Current in-progress items with % completion
        - Any blockers or risks identified this week
        - Key metrics from the dashboard (velocity, bug rate, deploy count)
        - Next week's top 3 priorities

        ## Output Format
        Markdown document with these exact sections:

        1. **Summary** — 1-2 sentences. Overall tone of the week. Lead with the headline.
        2. **Shipped This Week** — Bullet list. Each item: feature name, project duration, schedule status.
        3. **In Progress** — Bullet list. Each item: feature name, % complete, expected completion date.
        4. **Blockers & Risks** — Bullet list. Each item: issue, impact, mitigation if known.
        5. **Key Metrics** — Bullet list. Each metric: name, value, target, ✓ or ✗.
        6. **Next Week Priorities** — Numbered list. Top 3 only. Actionable items.

        ## Tone & Style
        - Direct and concise. No filler words.
        - Confident but honest about risks.
        - Use specific numbers, not "improved" or "growing."
        - Bullet points, not paragraphs.
        - Title format: "Leadership Update — Week of [date]"

        ## Example
        See weekly-report-example.md for the reference format.

        ## Quality Checks
        - Every "In Progress" item has a percentage and date
        - Every metric has actual vs target
        - Blockers include impact, not just description
        - Summary reflects the overall data, not just the good news

  - type: tool_result
    tool: scratchpad_write
    content: "Written to skill-weekly-report.md (42 lines)"

  - type: assistant
    content: "Created skill-weekly-report.md. This is your reusable skill file. Every Friday, just provide the raw data and reference this skill. AI will produce a report matching your exact format, tone, and quality standards.\n\nUsage: 'Using the weekly-report skill, generate this week's leadership update. Here are the inputs: [paste raw data]'\n\nThe skill captures: your section structure, your tone preferences, your metric format, and your quality checks. Write it once, use it every week."

  - type: note
    text: "45 minutes of writing became a 5-minute task. The skill file captures not just the format but the thinking behind it: what inputs are needed, what tone to use, what quality checks to apply. This is institutional knowledge in a reusable form."

  - type: note
    text: "The compound effect: this person will use this skill 50+ times a year. Each use saves 40 minutes. That's 33+ hours saved from a 10-minute investment in creating the skill. And the quality is consistent every single time."
```

```quiz
id: skills-purpose
type: multiple-choice
question: "What's the main benefit of creating a skill for a recurring task?"
options:
  - "Skills make AI faster at processing"
  - "Document the procedure once; AI follows it automatically on every future task"
  - "Skills are required for complex tasks"
  - "Skills reduce the cost of AI usage"
answer: 1
explanation: "Skills capture your procedures so you document them once and AI follows them automatically every time. No re-explaining, consistent results, institutional knowledge preserved."
```
