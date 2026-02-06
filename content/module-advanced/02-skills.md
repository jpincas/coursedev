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
