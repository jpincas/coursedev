---
title: "Grounding: Your Specific Knowledge"
duration: "10m"
tags: [grounding, context, specifics]
---

# Grounding

Grounding means giving AI access to your specific knowledge — the information that makes your context unique.

![Grounding transformation: generic output becomes tailored output when you provide specific materials](/content/module-files/images/grounding-concept.png)

## The Problem

AI knows general things. It doesn't know:
- Your company's specific policies
- Your project's history
- Your industry's particular terminology
- Your team's preferred formats
- Your organisation's style guide

When AI lacks this specific knowledge, it produces generic output that doesn't quite fit your context.

## The Solution

Provide the specific documents, data, and context that contain your knowledge.

**Upload your style guide**
AI writes in your voice, uses your terminology, follows your conventions.

**Provide past reports**
AI matches your format, structure, and level of detail.

**Include policy documents**
AI cites correct procedures, references actual policies.

**Share templates**
AI follows your established patterns.

```callout
type: tip
title: "Show, Don't Tell"
content: "Rather than describing your preferred style, provide an example of work you liked. AI will learn from the actual artifact more reliably than from your description of it."
```

## Style Cards and Writing

**Style cards** are reusable prompts encoding your preferences for tone, vocabulary, sentence structure, and audience.

Create one once. Use it across every session. AI maintains your voice consistently.

```callout
type: tip
title: "EchoWriting Technique"
content: "A powerful technique called EchoWriting — feeding AI samples of your writing to learn your style — is covered in depth in the Document Creation module."
```

## Multi-Document Work

**One of AI's highest-value use cases** is working across multiple documents simultaneously.

Upload multiple documents. AI can identify common themes, contradictions between sources, and gaps in coverage.

```callout
type: tip
title: "Document Synthesis"
content: "Multi-document synthesis (uploading multiple documents for AI to compare and connect) is covered in the Document Creation module."
```

## Grounding in Practice

**Without grounding:**
"Write a client email about the project delay."
Result: Generic corporate-speak that doesn't match your voice.

**With grounding:**
"Write a client email about the project delay. Here are three previous client emails I've sent for tone and style." *[attach emails]*
Result: Email matches your actual communication style.

## What to Provide

Depending on the task, consider including:

| Task Type | Grounding Materials |
|-----------|---------------------|
| Writing | Style guide, example pieces, brand guidelines |
| Analysis | Past analyses, template formats, metric definitions |
| Presentations | Previous decks, company templates, brand assets |
| Code | Existing codebase, coding standards, architecture docs |
| Reports | Previous reports, formatting preferences, required sections |

## The Payoff

Grounded AI output:
- Fits your context
- Uses correct terminology
- Follows your conventions
- Requires less editing
- Feels like it came from your team

Ungrounded AI output:
- Feels generic
- Uses unfamiliar phrasing
- Misses context-specific details
- Requires significant rework

The investment in providing grounding materials pays off in output quality.

## Grounding in Action

Words about grounding are themselves generic. Watch what it actually does — same request, with and without grounding materials.

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

```quiz
id: grounding-purpose
type: multiple-choice
question: "What problem does grounding solve?"
options:
  - "AI running out of context window space"
  - "AI producing generic output that doesn't fit your specific context"
  - "AI taking too long to generate responses"
  - "AI hallucinating incorrect facts"
answer: 1
explanation: "Grounding solves the problem of generic output. AI knows general patterns but not your specific context — your style, terminology, policies, and preferences. Providing grounding materials lets AI produce output that fits your actual situation."
```
