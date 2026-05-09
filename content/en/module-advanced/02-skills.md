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
- Code (scripts that fetch data, call APIs, or automate preparation)

When you ask AI to do something, it looks at your available skills. If there's a relevant one, it follows those procedures automatically.

```callout
type: info
title: "A Note on Naming"
content: "\"Skills\" is currently Claude's name for this concept — a structured folder of instructions and resources that the AI loads on demand. But the underlying pattern is universal. Every major AI platform is converging on similar ideas: reusable, user-defined procedures that customise AI behaviour for specific tasks. Whatever your tool calls them, the principles in this section apply."
```

## Skill Structure

A typical skills folder might look like:

```
Skills/
├── weekly-report/
│   ├── SKILL.md
│   ├── example-report.md
│   └── fetch_metrics.py
├── expense-summary/
│   ├── SKILL.md
│   └── format-template.xlsx
└── client-proposal/
    ├── SKILL.md
    ├── template.pptx
    └── lookup_client.py
```

Each skill is a folder containing its instructions (`SKILL.md`), reference material, and optionally **code** — small scripts that connect to APIs, fetch data, or automate preparation steps. The AI reads the instructions and uses the supporting files to execute the task.

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

## Skills Can Include Code

Skills aren't limited to instructions and templates. You can include scripts that extend what the AI can do:

**API connections** — A Python script that pulls metrics from your internal dashboard, so the AI has fresh data to work with.

**Data processing** — A script that cleans and formats a CSV export before the AI analyses it.

**Image generation** — Code that calls an image API to create charts or graphics as part of a reporting workflow.

**System integration** — Scripts that post results to Slack, update a spreadsheet, or trigger a deployment.

The AI reads the code, understands what it does, and can run it as part of executing the skill. You don't need to be a developer — describe what you need and let the AI write the code for you.

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

## When to Be Explicit About Skills and Tools

Sometimes AI automatically uses the right skill or tool without being asked. You say "generate the weekly report" and it just knows to use the skill. Other times, you need to be explicit: "use the /review skill" or "search the web for..."

**Why this inconsistency?** AI decides based on context whether a skill or tool is relevant. Sometimes that decision is obvious. Sometimes it's not.

Knowing when to be explicit is itself a skill you develop with practice.

**Rule of thumb:** If the AI isn't doing what you expect, try being explicit about which tool or skill to use.

Instead of: "Find information about this company"
Try: "Search the web for recent news about this company"

Instead of: "Check this code"
Try: "Use the /review skill to check this code"

The more specific you are about the mechanism, the less AI has to infer about your intent.

```callout
type: tip
title: "Default to Implicit, Escalate to Explicit"
content: "Start with natural language. If AI doesn't pick the right tool or skill, make your next instruction explicit about which one to use. This pattern — implicit first, explicit on retry — is faster than being overly prescriptive from the start."
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

You don't need to build skills by hand. Describe what you need and let AI create the entire skill for you — folder structure, instructions, reference material, even the code.

Watch the process of turning a repeating task into a complete, reusable skill folder.

```agent
id: skill-creation-demo
title: "Building a Reusable Skill"
model_label: "Claude"

system: |
  You are a productivity consultant helping create reusable AI skills.
  Skills should be clear, complete, and actionable.
  When creating skill folders, use the convention: Skills/<skill-name>/SKILL.md for instructions.

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
  - create_folder
  - move_file

visibility:
  system_prompt: hidden
  tool_calls: visible
  full_context: hidden

sidebar:
  width: "45%"
  start_open: true

script:
  - type: note
    text: "This person writes a weekly leadership report every Friday. It takes 45 minutes. They want to turn it into a proper skill — a folder with instructions, reference material, and a script to pull metrics automatically. Watch how AI builds the entire skill for them."

  - type: user
    content: "I write this leadership report every Friday and it takes forever. I've got an example report here. Can you build me a proper skill folder for this? I need clear instructions, my example as reference, and a Python script that pulls our engineering metrics from the internal dashboard API so I don't have to copy numbers manually each week."

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
    tool: create_folder
    args:
      path: "Skills"

  - type: tool_result
    tool: create_folder
    content: "Created folder: Skills"

  - type: tool_call
    tool: create_folder
    args:
      path: "Skills/weekly-report"

  - type: tool_result
    tool: create_folder
    content: "Created folder: Skills/weekly-report"

  - type: tool_call
    tool: scratchpad_write
    args:
      filename: "Skills/weekly-report/SKILL.md"
      content: |
        # Skill: Weekly Leadership Report

        ## Purpose
        Generate the Friday leadership update for the engineering team.

        ## When to Use
        Every Friday afternoon. Run fetch_metrics.py first to get fresh data, then provide any additional context.

        ## Preparation (AI does this automatically)
        1. Run `python fetch_metrics.py` to pull fresh dashboard data
        2. Use the metrics output alongside any additional context provided

        ## Required Inputs (from the user)
        - This week's completed work (from Jira/Linear or standup notes)
        - Current in-progress items with % completion
        - Any blockers or risks identified this week
        - Next week's top 3 priorities

        ## Output Format
        Markdown document with these exact sections:

        1. **Summary** — 1-2 sentences. Overall tone of the week. Lead with the headline.
        2. **Shipped This Week** — Bullet list. Each: feature name, duration, schedule status.
        3. **In Progress** — Bullet list. Each: feature name, % complete, expected date.
        4. **Blockers & Risks** — Bullet list. Each: issue, impact, mitigation if known.
        5. **Key Metrics** — Bullet list. Each: name, value, target, pass/fail.
        6. **Next Week Priorities** — Numbered list. Top 3 only. Actionable items.

        ## Tone & Style
        - Direct and concise. No filler words.
        - Confident but honest about risks.
        - Use specific numbers, not "improved" or "growing."
        - Bullet points, not paragraphs.
        - Title: "Leadership Update — Week of [date]"

        ## Reference
        See example-report.md for the format and tone to match.

        ## Quality Checks
        - Every "In Progress" item has a percentage and date
        - Every metric has actual vs target
        - Blockers include impact, not just description
        - Summary reflects the overall data, not just the good news

  - type: tool_result
    tool: scratchpad_write
    content: "Written to Skills/weekly-report/SKILL.md (48 lines)"

  - type: tool_call
    tool: move_file
    args:
      source: "weekly-report-example.md"
      destination: "Skills/weekly-report/example-report.md"

  - type: tool_result
    tool: move_file
    content: "Moved weekly-report-example.md → Skills/weekly-report/example-report.md"

  - type: tool_call
    tool: scratchpad_write
    args:
      filename: "Skills/weekly-report/fetch_metrics.py"
      content: |
        #!/usr/bin/env python3
        """Fetch weekly engineering metrics from the dashboard API.

        Run before generating the weekly report:
            python fetch_metrics.py > metrics.md

        Requires DASHBOARD_API_KEY environment variable.
        """

        import os
        import requests
        from datetime import datetime, timedelta

        API = "https://dashboard.internal.company.com/api/v1"
        KEY = os.environ["DASHBOARD_API_KEY"]
        HEADERS = {"Authorization": f"Bearer {KEY}"}


        def get_sprint():
            r = requests.get(f"{API}/sprints/current", headers=HEADERS)
            r.raise_for_status()
            return r.json()


        def get_deploys():
            since = (datetime.now() - timedelta(days=7)).isoformat()
            r = requests.get(
                f"{API}/deploys",
                params={"since": since},
                headers=HEADERS,
            )
            r.raise_for_status()
            return r.json()["deploys"]


        def get_quality():
            r = requests.get(f"{API}/quality/summary", headers=HEADERS)
            r.raise_for_status()
            return r.json()


        if __name__ == "__main__":
            sprint = get_sprint()
            deploys = get_deploys()
            quality = get_quality()

            print("# This Week's Metrics\n")
            print(f"- Sprint velocity: {sprint['velocity']} pts "
                  f"(target: {sprint['target']})")
            print(f"- Deploys this week: {len(deploys)}")
            print(f"- Bug escape rate: {quality['escape_rate']}% "
                  f"(target: under {quality['target_rate']}%)")
            print(f"- QA queue depth: {quality['queue_depth']} tickets")

  - type: tool_result
    tool: scratchpad_write
    content: "Written to Skills/weekly-report/fetch_metrics.py (42 lines)"

  - type: assistant
    content: "Built your complete skill folder. Here's what you have now:\n\n**Skills/weekly-report/**\n- `SKILL.md` — Full instructions: purpose, inputs, output format, tone, quality checks\n- `example-report.md` — Your reference report, so AI matches your style\n- `fetch_metrics.py` — Pulls sprint velocity, deploy count, bug rate, and QA queue from your dashboard API\n\n**How to use it:**\nJust say *'Generate this week's leadership update'* and provide your notes. The skill tells AI to run the metrics script automatically, pull the fresh numbers, and produce the report in your exact format. You never touch the script yourself.\n\nWrite it once, use it every Friday."

  - type: note
    text: "Notice the folder structure: Skills/weekly-report/ contains everything — instructions, reference material, and code. This is a self-contained, portable skill. You could share it with a colleague and they'd have everything they need."

  - type: note
    text: "Key insight: you didn't write any of this by hand. You described what you needed and AI built the entire skill — instructions, file organisation, even the Python script. That's the meta-tooling pattern: use AI to build the tools you'll use with AI."

  - type: note
    text: "The compound effect: this skill will be used 50+ times a year. Each use saves 40 minutes. That's 33+ hours saved from a single conversation. And the quality is consistent every time."
```

```quiz
id: skills-purpose
type: multiple-choice
question: "What's the main benefit of creating a skill for a recurring task?"
options:
  - "Document the procedure once; AI follows it automatically on every future task"
  - "Skills make AI faster at processing"
  - "Skills reduce the cost of AI usage"
answer: 0
explanation: "Skills capture your procedures so you document them once and AI follows them automatically every time. No re-explaining, consistent results, institutional knowledge preserved."
```
