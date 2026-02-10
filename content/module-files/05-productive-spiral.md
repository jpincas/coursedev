---
title: "The Productive Spiral"
duration: "15m"
tags: [spiral, outputs-as-inputs, grounding, iteration]
---

# The Productive Spiral

Here is the concept that transforms AI from a useful tool into a compounding productivity system.

**Output files become input files for the next round.**

This creates a positive spiral where quality compounds with each iteration:

1. You give AI messy source files
2. AI produces a clean summary document
3. That summary becomes context for the next task
4. That next output enriches the round after
5. Each cycle builds on the last

![The Productive Spiral: outputs feed back as inputs, quality compounds](/content/module-files/images/productive-spiral.svg)

## The Spiral in Practice

**Example: Preparing a Client Pitch**

**Round 1 -- Research:** Give AI your client's website, recent press releases, and industry reports. AI produces a comprehensive client research brief.

**Round 2 -- Strategy:** Feed that research brief plus your service offerings. AI produces a tailored strategy document with specific recommendations.

**Round 3 -- Proposal:** Feed the strategy document plus your proposal template and past successful proposals. AI produces a polished proposal that builds on the research and strategy.

**Round 4 -- Presentation:** Feed the proposal plus your slide template. AI creates presentation slides that summarise the proposal's key points.

Each round used the output of the previous round as input context. The final presentation contains insights from every stage -- research, strategy, and proposal -- because they all accumulated as context.

```callout
type: info
title: "The Compounding Effect"
content: "Each round does not start from zero. It starts from the accumulated context of all previous rounds. This is why the fifth deliverable in a spiral is dramatically better than five separate one-shot requests."
```

## The Spiral in Action

Watch outputs become inputs across three rounds. Pay attention to the file explorer -- notice how the workspace grows richer with each round.

```agent
id: productive-spiral-demo
title: "Outputs Become Inputs"
model_label: "Claude"

system: |
  You are a business analyst and writer. You create clear,
  well-structured documents from raw data and materials.

scratchpad:
  "raw-meeting-notes.txt": |
    Q4 all hands - Dec 15
    Revenue hit $2.1M (target was $1.9M)
    New customers: 340 (up from 210 last quarter)
    Churn dropped to 2.3% (was 3.1%)
    NPS: 52 (up from 44)
    Enterprise deals: closed Meridian ($180K ARR),
    close to signing DataFlow ($240K ARR)
    Hiring: 12 new hires in Q4, 8 more approved for Q1
    Product: v3.0 shipped on time, usage up 40%
    Support tickets down 25% since v3.0
    CEO quote: "Best quarter in company history"
  "customer-feedback.csv": |
    Customer,Segment,NPS_Score,Comment
    Meridian Corp,Enterprise,9,Love the new dashboard
    TechStart Ltd,SMB,7,Good but onboarding was slow
    DataFlow Inc,Enterprise,10,Best tool we have used
    GreenLeaf Co,SMB,5,Missing Excel export feature
    Atlas Group,Enterprise,9,Support team is excellent
    BrightPath,SMB,8,Great value for money
    CoreSync,Mid-market,6,API documentation needs work
    Pinnacle Ltd,Enterprise,8,Solid product reliable team
    Nova Digital,SMB,9,Transformed our workflow
    Summit Inc,Mid-market,7,Would like more integrations
  "previous-quarterly.md": |
    # Q3 Quarterly Summary

    ## Highlights
    - Revenue: $1.7M (5% above target)
    - New customers: 210
    - Product v2.5 shipped with mixed reception

    ## Concerns
    - Churn at 3.1% — above 2.5% target
    - Enterprise pipeline thin
    - Support ticket volume increasing

    ## Q4 Priorities
    - Ship v3.0 on schedule
    - Close 2+ enterprise deals
    - Reduce churn below 2.5%

tools:
  - scratchpad_read
  - scratchpad_write
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
    text: "Round 1: We start with raw materials — meeting notes, customer data, and last quarter's summary. Watch what the AI produces, and then watch that output become the input for round 2."

  - type: user
    content: "I have my Q4 all-hands notes, customer feedback data, and last quarter's summary. Please read everything and create a Q4 quarterly summary in the same format as the Q3 one."

  - type: tool_call
    tool: scratchpad_read
    args:
      filename: "raw-meeting-notes.txt"

  - type: tool_result
    tool: scratchpad_read
    content: |
      Q4 all hands - Dec 15
      Revenue hit $2.1M (target was $1.9M)
      New customers: 340 (up from 210 last quarter)
      Churn dropped to 2.3% (was 3.1%)
      NPS: 52 (up from 44)
      Enterprise deals: closed Meridian ($180K ARR)...

  - type: tool_call
    tool: scratchpad_read
    args:
      filename: "customer-feedback.csv"

  - type: tool_result
    tool: scratchpad_read
    content: |
      Customer,Segment,NPS_Score,Comment
      Meridian Corp,Enterprise,9,Love the new dashboard
      TechStart Ltd,SMB,7,Good but onboarding was slow
      DataFlow Inc,Enterprise,10,Best tool we have used...

  - type: tool_call
    tool: scratchpad_read
    args:
      filename: "previous-quarterly.md"

  - type: tool_result
    tool: scratchpad_read
    content: |
      # Q3 Quarterly Summary
      ## Highlights
      - Revenue: $1.7M (5% above target)
      - New customers: 210...

  - type: tool_call
    tool: scratchpad_write
    args:
      filename: "q4-quarterly-summary.md"
      content: |
        # Q4 Quarterly Summary

        ## Highlights
        - Revenue: $2.1M (11% above $1.9M target, up 24% from Q3)
        - New customers: 340 (62% increase from Q3's 210)
        - Churn: 2.3% — down from 3.1%, now below 2.5% target
        - NPS: 52 (up from 44 in Q3)
        - Product v3.0 shipped on time — usage up 40%, support tickets down 25%

        ## Enterprise Momentum
        - Closed Meridian Corp ($180K ARR)
        - DataFlow Inc ($240K ARR) near close
        - Enterprise NPS averaging 9.0 — highest segment

        ## Customer Sentiment
        - Overall NPS: 52 across 10 surveyed customers
        - Enterprise segment strongest (avg 9.0)
        - SMB segment solid (avg 7.3) but onboarding flagged as slow
        - Improvement areas: Excel export, API docs, integrations

        ## Q3 Priorities — Status
        - ✅ Ship v3.0 on schedule — delivered, strong adoption
        - ✅ Close 2+ enterprise deals — Meridian closed, DataFlow pending
        - ✅ Reduce churn below 2.5% — achieved at 2.3%

        ## Team
        - 12 new hires in Q4, 8 more approved for Q1
        - CEO: "Best quarter in company history"

  - type: tool_result
    tool: scratchpad_write
    content: "Written to q4-quarterly-summary.md (28 lines)"

  - type: assistant
    content: "I've created the Q4 quarterly summary. Key story: every Q3 priority was hit. Revenue up 24%, customers up 62%, churn down, and v3.0 landed well. The customer feedback data shows enterprise is your strongest segment with SMB onboarding as the main improvement area."

  - type: note
    text: "Round 1 complete. Look at the file explorer — there's now a new file: q4-quarterly-summary.md. That file was created from three raw inputs. Now watch it become the input for round 2."

  - type: clear
    reset_scratchpad: false
    note: "Round 2 — the summary from round 1 is still in the workspace. Now we build on it."

  - type: user
    content: "Now I need to write an email to the board. Read the Q4 summary you just created and the customer feedback, and draft a board update email — 3 paragraphs max, confident but not hype-y."

  - type: tool_call
    tool: scratchpad_read
    args:
      filename: "q4-quarterly-summary.md"

  - type: tool_result
    tool: scratchpad_read
    content: |
      # Q4 Quarterly Summary
      ## Highlights
      - Revenue: $2.1M (11% above $1.9M target, up 24% from Q3)
      - New customers: 340 (62% increase from Q3's 210)
      - Churn: 2.3% — down from 3.1%, now below 2.5% target...

  - type: tool_call
    tool: scratchpad_read
    args:
      filename: "customer-feedback.csv"

  - type: tool_result
    tool: scratchpad_read
    content: |
      Customer,Segment,NPS_Score,Comment
      Meridian Corp,Enterprise,9,Love the new dashboard...

  - type: tool_call
    tool: scratchpad_write
    args:
      filename: "board-update-email.md"
      content: |
        Subject: Q4 Results — Strongest Quarter to Date

        Q4 closed at $2.1M revenue, 11% above target and up 24% from Q3. We added 340 new customers, closed our largest enterprise deal to date (Meridian Corp, $180K ARR), and have a $240K deal in late stages. Churn dropped from 3.1% to 2.3%, hitting our sub-2.5% target a quarter early.

        Product v3.0 shipped on schedule and is performing well — usage up 40% and support volume down 25%. Customer NPS reached 52, with enterprise customers rating us 9.0 out of 10 on average. The main feedback themes are requests for more integrations and faster SMB onboarding, both of which are on the Q1 roadmap.

        We are entering Q1 with strong momentum, 8 approved hires, and a healthy enterprise pipeline. The focus will be converting the DataFlow opportunity, improving SMB onboarding, and maintaining the retention gains we made in Q4.

  - type: tool_result
    tool: scratchpad_write
    content: "Written to board-update-email.md (7 lines)"

  - type: assistant
    content: "Done — three paragraphs covering performance, product, and outlook. It draws from both the quarterly summary and the raw customer feedback for the NPS and sentiment details. Confident but grounded in specific numbers."

  - type: note
    text: "Look at the workspace now. We started with 3 raw files. Round 1 produced a structured summary. Round 2 used that summary to produce a board email. The board email is better BECAUSE it built on the summary, which built on the raw data. Each round inherited the context of all previous rounds."

  - type: note
    text: "This is the productive spiral. The quarterly summary is not throwaway — it is now a reusable asset. Next month you could feed it into a strategy document, a hiring plan, or a sales forecast. Every output enriches your context library for future work."
```

## Why This Changes Your Workflow

The productive spiral means you should think about AI work not as isolated tasks but as **building a library of context.**

Every document AI creates is a future input file. Every analysis is grounding context for the next task. Every summary is a building block.

This is why file system organisation matters so much. You are not just keeping things tidy -- you are building the context library that makes every future AI interaction better.

## Grounding: Your Specific Knowledge

One of the most powerful applications of the productive spiral is **grounding** -- giving AI access to your specific knowledge so it produces tailored output instead of generic text.

AI knows general things. It does not know your company's specific policies, your project's history, your industry's terminology, or your team's preferred formats.

When you provide grounding materials -- your style guide, past reports, templates, examples -- AI produces output that fits your actual context.

**Without grounding:**
"Write a client email about the project delay."
Result: Generic corporate-speak that does not match your voice.

**With grounding:**
"Write a client email about the project delay. Here are three previous client emails I've sent for tone and style." *[attach emails]*
Result: Email matches your actual communication style.

```callout
type: tip
title: "Show, Don't Tell"
content: "Rather than describing your preferred style, provide an example of work you liked. AI learns from the actual artifact more reliably than from your description of it."
```

## Grounding in Action

Watch the difference grounding makes -- same request, with and without your specific materials.

```agent
id: grounding-demo
title: "Generic vs Grounded Output"
model_label: "Claude"

system: |
  You are a communications writer. Match the tone and style of
  any reference materials provided.

scratchpad:
  "previous-emails.txt": |
    Subject: Quick update on the Henderson project
    Hey team — wanted to flag something before it becomes a thing.
    Henderson's timeline just shifted left by two weeks. I know, I know.
    Here's what I need from each of you by Thursday:
    - Dev: revised sprint plan
    - Design: can we fast-track the prototype?
    - QA: what's the minimum viable test coverage?
    No need to panic. We've dealt with worse. Let's chat at standup tomorrow.
    — Sarah

    Subject: Re: Q3 planning
    Quick thought on Q3 — I think we're overcomplicating the roadmap.
    Last quarter we shipped 3 features well. This quarter we're planning 7.
    I'd rather ship 4 things brilliantly than 7 things badly.
    Can we trim before the planning session? Bring your "cut list" on Monday.
    — Sarah

tools:
  - scratchpad_read

visibility:
  system_prompt: hidden
  tool_calls: visible
  full_context: hidden

sidebar:
  width: "45%"
  start_open: true

script:
  - type: note
    text: "First: no grounding. Just a bare request. Watch how generic the output is."

  - type: user
    content: "Write an email to my team about pushing back the product launch by one week due to QA findings."

  - type: assistant
    content: "Subject: Product Launch Date Update\n\nDear Team,\n\nI hope this message finds you well. I am writing to inform you of an important update regarding our upcoming product launch timeline.\n\nAfter a thorough review of our Quality Assurance findings, we have made the decision to postpone the launch by one week. This adjustment will allow us to address the identified issues and ensure we deliver a product that meets our high standards of quality.\n\nPlease adjust your schedules accordingly. If you have any questions or concerns about this change, do not hesitate to reach out.\n\nBest regards,\n[Your Name]"

  - type: note
    text: "Corporate boilerplate. 'I hope this message finds you well.' 'Do not hesitate to reach out.' This could have been written by anyone at any company. Now let's ground it with Sarah's actual writing."

  - type: clear
    reset_scratchpad: false
    note: "Same request, but grounded with real examples this time."

  - type: user
    content: "Read my previous emails first — that's how I actually write. Then write an email to my team about pushing back the product launch by one week because QA found issues we need to fix."

  - type: tool_call
    tool: scratchpad_read
    args:
      filename: "previous-emails.txt"

  - type: tool_result
    tool: scratchpad_read
    content: |
      Subject: Quick update on the Henderson project
      Hey team — wanted to flag something before it becomes a thing.
      Henderson's timeline just shifted left by two weeks. I know, I know.
      Here's what I need from each of you by Thursday:
      - Dev: revised sprint plan
      - Design: can we fast-track the prototype?
      - QA: what's the minimum viable test coverage?
      No need to panic. We've dealt with worse. Let's chat at standup tomorrow.
      — Sarah

      Subject: Re: Q3 planning
      Quick thought on Q3 — I think we're overcomplicating the roadmap.
      Last quarter we shipped 3 features well. This quarter we're planning 7.
      I'd rather ship 4 things brilliantly than 7 things badly.
      Can we trim before the planning session? Bring your "cut list" on Monday.
      — Sarah

  - type: assistant
    content: "Subject: Launch is sliding a week — here's why and what I need\n\nHey team — QA turned up some things we can't ship with. Nothing catastrophic, but enough that I'd rather push a week than rush a fix.\n\nNew launch date: [date + 1 week].\n\nWhat I need from you:\n- Dev: triage the QA findings by severity. What's fix-now vs fix-in-patch?\n- QA: once fixes land, how fast can we re-validate?\n- Marketing: adjust the launch comms timeline. One week shift, same sequence.\n\nI know nobody loves a delay. But shipping something broken is worse than shipping something late. We've got this.\n\n— Sarah"

  - type: note
    text: "Read both outputs side by side. The first: 'Dear Team, I hope this message finds you well.' The second: 'Hey team — QA turned up some things we can't ship with.' Same AI. Same task. The difference is two example emails. That's grounding."

  - type: note
    text: "Notice what the AI picked up from the examples: the casual 'Hey team' opener, the dash-separated asides, the direct action items per person, the reassuring closer, the signature style. It learned Sarah's voice from her own writing — far better than any description of 'casual but professional' could achieve."
```

## Building Your Context Library

Practical steps to start your productive spiral:

**Save AI outputs deliberately.** When AI creates a good document, save it. It becomes grounding context for future tasks.

**Build style references.** Collect examples of writing, reports, and communications you like. These become your grounding materials.

**Create templates.** Have AI create templates based on your best work. Use those templates as input for future tasks.

**Maintain a project knowledge base.** For ongoing projects, keep a running document of decisions, context, and progress. Feed it to AI at the start of each session.

```callout
type: tip
title: "The Mindset"
content: "Every AI interaction is not just producing today's deliverable. It is building context for tomorrow's work. Treat your file system as a compounding asset."
```

```quiz
id: productive-spiral-concept
type: multiple-choice
question: "After AI creates a research brief, you use that brief as context when asking AI to write a strategy document. Why does this produce better results than writing the strategy from scratch?"
options:
  - "The model remembers the research brief from the previous conversation"
  - "The strategy document inherits the research context, so the model has both your requirements and comprehensive background information"
  - "Using two separate prompts is always better than one because it splits the workload"
answer: 1
explanation: "The productive spiral works because each round's output enriches the next round's context. The strategy document is better because the model has both your strategy requirements AND the comprehensive research to draw on. It is not about splitting workload -- it is about compounding context quality."
```
