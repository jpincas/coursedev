---
title: "Demo: From Chaos to Report"
duration: "5m"
tags: [demo, delegation]
---

# The Demo: From Chaos to Report

Imagine a disorganised folder with mixed file types — meeting notes, spreadsheets, emails, timelines — the kind of chaos that accumulates on real projects.

With one delegation instruction, AI can:
- Plan and execute the task autonomously
- Process multiple file types
- Deliver a polished, finished document

Notice: No step-by-step instructions. Just describing the outcome you want. AI figures out how to get there.

**This is what we mean by "AI that does work."**

## How It Works

The **file explorer on the left** shows five messy project files — meeting notes, budget data, client emails, a project timeline, and team feedback.

The **conversation panel on the right** walks through the AI interaction. Step through using the **Next** button and watch one instruction turn this mess into a polished executive report.

```agent
id: opening-demo
title: "From Chaos to Report"
model_label: "Claude"

system: |
  You are a project management assistant. You help professionals create clear, actionable reports from project materials. Be thorough but concise. Write in a professional tone suitable for executive audiences.

scratchpad:
  "meeting-notes.txt": |
    Project Phoenix - Team Standup Jan 15
    Attendees: Sarah (PM), Dev team, Lisa (Design)
    - API migration 60% complete, targeting Feb 1
    - Design review blocked on brand guidelines
    - Client wants mobile dashboard added to Phase 1
    - Contractor rates up 15% — budget concern
    - Sarah to escalate timeline risk to leadership
    - Next milestone review: Jan 22
  "budget-q1.csv": |
    Category,Planned,Actual,Variance
    Development,45000,38000,-7000
    Design,15000,12500,-2500
    Infrastructure,8000,9200,+1200
    Contractors,20000,27500,+7500
    Testing,7000,4200,-2800
    Total,95000,91400,-3600
  "client-emails.txt": |
    From: David Chen <d.chen@acmecorp.com>
    Date: Jan 14
    Subject: RE: Project Phoenix Timeline

    Sarah — The board wants mobile dashboard in Phase 1.
    Revised timeline? We need this by March 15 latest.
    ---
    From: Sarah Lopez <s.lopez@team.com>
    Date: Jan 12

    David — Assessing mobile addition now. Adds ~3 weeks.
    Revised plan by end of next week.
  "project-timeline.md": |
    # Project Phoenix Timeline (v2.1)
    Phase            | Target | Status
    API Migration    | Feb 1  | In Progress (60%)
    Frontend Rebuild | Feb 15 | Not Started
    Design System    | Feb 8  | Blocked
    Testing          | Feb 22 | Not Started
    Launch           | Mar 1  | At Risk
    *Does not include mobile dashboard scope change*
  "team-feedback.txt": |
    Anonymous Team Pulse Survey - January (5 responses)
    "Scope creep is a problem. Adding features without adjusting timeline."
    "Good API progress but team stretched thin with contractor transition."
    "Design blocked — creating a domino effect downstream."
    "Morale OK but worried about March deadline with mobile addition."
    "Client communication improved since Sarah took over PM."

tools:
  - scratchpad_read
  - scratchpad_write

visibility:
  system_prompt: toggleable
  tool_calls: visible
  full_context: hidden
  token_count: hidden
  model_name: visible

sidebar:
  width: "45%"

script:
  - type: note
    text: "This is a live walkthrough of AI doing real work. The file explorer shows five project files — meeting notes, budget data, client emails, a timeline, and team survey results. Messy, disorganised, typical."

  - type: note
    text: "Now watch. One instruction. No templates, no step-by-step breakdown. Just describe the outcome you need."

  - type: user
    content: "I've got a project folder with meeting notes, budget data, client emails, a timeline, and team feedback. I need an executive status report for leadership by end of day. Pull together the key points, flag the risks, and make it professional."

  - type: assistant
    content: "I'll review all the project materials and create a professional executive status report. Let me start by reading through each file."

  - type: tool_call
    tool: scratchpad_read
    args:
      filename: "meeting-notes.txt"

  - type: tool_result
    tool: scratchpad_read
    content: |
      Project Phoenix - Team Standup Jan 15
      Attendees: Sarah (PM), Dev team, Lisa (Design)
      - API migration 60% complete, targeting Feb 1
      - Design review blocked on brand guidelines
      - Client wants mobile dashboard added to Phase 1
      - Contractor rates up 15% — budget concern
      - Sarah to escalate timeline risk to leadership
      - Next milestone review: Jan 22

  - type: tool_call
    tool: scratchpad_read
    args:
      filename: "budget-q1.csv"

  - type: tool_result
    tool: scratchpad_read
    content: |
      Category,Planned,Actual,Variance
      Development,45000,38000,-7000
      Design,15000,12500,-2500
      Infrastructure,8000,9200,+1200
      Contractors,20000,27500,+7500
      Testing,7000,4200,-2800
      Total,95000,91400,-3600

  - type: tool_call
    tool: scratchpad_read
    args:
      filename: "client-emails.txt"

  - type: tool_result
    tool: scratchpad_read
    content: |
      From: David Chen <d.chen@acmecorp.com>
      Date: Jan 14
      Subject: RE: Project Phoenix Timeline

      Sarah — The board wants mobile dashboard in Phase 1.
      Revised timeline? We need this by March 15 latest.
      ---
      From: Sarah Lopez <s.lopez@team.com>
      Date: Jan 12

      David — Assessing mobile addition now. Adds ~3 weeks.
      Revised plan by end of next week.

  - type: tool_call
    tool: scratchpad_read
    args:
      filename: "project-timeline.md"

  - type: tool_result
    tool: scratchpad_read
    content: |
      # Project Phoenix Timeline (v2.1)
      Phase            | Target | Status
      API Migration    | Feb 1  | In Progress (60%)
      Frontend Rebuild | Feb 15 | Not Started
      Design System    | Feb 8  | Blocked
      Testing          | Feb 22 | Not Started
      Launch           | Mar 1  | At Risk
      *Does not include mobile dashboard scope change*

  - type: tool_call
    tool: scratchpad_read
    args:
      filename: "team-feedback.txt"

  - type: tool_result
    tool: scratchpad_read
    content: |
      Anonymous Team Pulse Survey - January (5 responses)
      "Scope creep is a problem. Adding features without adjusting timeline."
      "Good API progress but team stretched thin with contractor transition."
      "Design blocked — creating a domino effect downstream."
      "Morale OK but worried about March deadline with mobile addition."
      "Client communication improved since Sarah took over PM."

  - type: assistant
    content: "I've reviewed all five documents. Several clear themes: a scope change risk from the mobile dashboard request, a design blockage creating downstream delays, and contractor cost pressure. Let me synthesise this into a structured executive report."

  - type: tool_call
    tool: scratchpad_write
    args:
      filename: "exec-status-report.md"
      content: |
        # Project Phoenix — Executive Status Report
        Prepared: January 15, 2026

        ## Summary
        Project Phoenix is 60% through API migration but faces timeline risk from a client-requested scope change (mobile dashboard). Current spend is under budget but contractor overruns signal pressure.

        ## Key Metrics
        - Budget: 91,400 of 95,000 spent (3.8% under)
        - Timeline: Original launch Mar 1 — AT RISK
        - Scope: Mobile dashboard added to Phase 1, not yet planned

        ## Critical Risks

        1. SCOPE CHANGE: Mobile dashboard adds ~3 weeks. Client deadline March 15. Recommend presenting revised timeline with phased options.

        2. DESIGN BLOCKAGE: Blocked on brand guidelines. Creates downstream delays to frontend and testing. Escalate to client this week.

        3. CONTRACTOR COSTS: 7,500 over plan (37.5% variance). Will accelerate with expanded scope. Review rates and consider fixed-price for mobile work.

        ## Team Health
        Stable but concerned about scope creep and timeline pressure. PM communication with client rated positively.

        ## Recommended Next Steps
        1. Present revised timeline to client by Jan 22
        2. Escalate brand guidelines blocker
        3. Review contractor cost structure
        4. Schedule mobile dashboard scoping session

  - type: tool_result
    tool: scratchpad_write
    content: "Written to exec-status-report.md (26 lines)"

  - type: assistant
    content: "Done. I've created exec-status-report.md — a structured executive status report covering project summary, key metrics, three critical risks with recommendations, team health, and prioritised next steps. Ready for leadership review."

  - type: note
    text: "Five scattered files became one professional report. No copy-paste. No back-and-forth. No templates. The AI read every file, identified the themes, and delivered a finished document."

  - type: note
    text: "Notice what you didn't do: you didn't specify which files to read first, how to structure the report, or what sections to include. You described the outcome. AI figured out the rest. That's delegation, not prompting. That's what today's training will teach you."
```
