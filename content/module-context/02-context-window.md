---
title: "The Context Window"
duration: "15m"
tags: [context-window, tokens, capacity]
---

# What Is the Context Window?

The context window is the model's "working memory." It's everything the model can see when generating a response.

![Context Window: System Prompt, Conversation History, Files, and Current Message flow into Model Processing](/content/module-context/images/context-window.svg)

## Context Window Sizes

Context windows have grown dramatically over time:

| Era | Size | Equivalent |
|-----|------|------------|
| GPT-3 (2020) | ~4K tokens | ~3,000 words |
| GPT-4 (initial) | ~8K tokens | ~6,000 words |
| Claude 2 (2023) | ~100K tokens | ~75,000 words |
| 2025 (typical) | ~200K tokens | ~150,000 words |
| **Current (2026)** | **~1M tokens** | **~750,000 words** |

**Claude Opus 4.6 now supports 1 million tokens — roughly 750,000 words, or about 2,500 pages of text.** Gemini models support up to 2 million tokens.

But there's a catch: **the "lost in the middle" effect.** Models struggle with information buried deep in large contexts. Position matters even more at this scale.

```callout
type: info
title: "What's a Token?"
content: "Tokens are chunks of text the model processes. Roughly 1 token = 0.75 words, or about 4 characters in English. 'Hello' is 1 token. 'Understanding' is 2 tokens."
```

## The Shift in Constraints

This growth changes everything.

**Old constraint:** What fits in context? (Very limited. Had to be selective.)

**New constraint:** What's relevant? (Can fit entire books, but relevance still matters.)

You can now fit:
- Entire application codebases with dependencies
- Complete technical documentation libraries
- Multiple full-length books simultaneously
- Days of detailed conversation history

**Anthropic's "Infinite Chats" feature** goes even further: server-side summarisation extends conversations indefinitely without losing critical context.

But bigger isn't automatically better. Models still need to find what's relevant in all that information — and the "lost in the middle" effect means placement strategy matters more than ever.

## What Goes Where

Not all context is equal. Here's what typically fills the window:

**System Prompt** (usually invisible to you)
- Set by the application you're using
- Defines personality, capabilities, constraints
- Why Claude on claude.ai behaves differently than Claude in a custom app

**Conversation History**
- Everything you and the model have exchanged
- Grows with each message
- Provides continuity but can also clutter

**Files and Documents**
- PDFs, code files, images, data
- Explicitly provided by you
- Most direct way to give grounded information

**Your Current Message**
- What you just typed
- Your immediate instruction or question

## Watching Context Fill Up

This demo makes the invisible visible. Watch the token counter as files are loaded into context — you can see the window filling in real time.

```agent
id: token-counting-demo
title: "The Context Window Filling Up"
model_label: "Claude"

system: |
  You are a document analyst. Read files provided and answer questions
  about them. Be concise.

scratchpad:
  "meeting-notes.txt": |
    Product Team Weekly — 14 Jan 2026
    Attendees: Maya (PM), Dev team, UX team

    Decisions:
    - Ship v2.3 by end of month
    - Postpone dark mode to v2.4
    - Hire contract QA for launch sprint

    Action items:
    - Maya: Finalise launch checklist by Friday
    - Dev: Fix critical auth bug (#4521)
    - UX: Update onboarding flow mockups
  "strategy-doc.txt": |
    Q1 2026 Product Strategy

    Vision: Become the default collaboration tool for mid-market teams.

    Three Pillars:
    1. Speed — reduce page load times to under 200ms
    2. Integration — ship Slack, Teams, and Jira connectors
    3. Self-serve — enable team admin without IT involvement

    Target metrics:
    - 40% increase in weekly active users
    - Net Promoter Score above 50
    - Churn below 3% monthly

    Competitive landscape:
    - Notion expanding into project management
    - Asana adding AI features
    - Monday.com aggressive on pricing

    Our advantage: simplicity and speed for non-technical teams.
  "customer-feedback.txt": |
    Recent NPS Responses (Jan 2026)

    Score 9: "Finally a tool my whole team actually uses. The simplicity is key."
    Score 8: "Love it but need better mobile app. Hard to use on the go."
    Score 3: "Integration with our existing tools is painful. Had to build workarounds."
    Score 10: "Replaced 3 tools with this one. Saving hours per week."
    Score 4: "Search is broken for large workspaces. Can't find anything."
    Score 9: "Best onboarding experience I've seen. Team was up in 20 minutes."
    Score 2: "No SSO support is a dealbreaker for enterprise. Had to cancel."
    Score 7: "Good but not great. Missing Gantt charts for project tracking."

tools:
  - scratchpad_read

visibility:
  system_prompt: hidden
  tool_calls: visible
  full_context: hidden
  token_count: visible

sidebar:
  width: "45%"
  start_open: true

script:
  - type: note
    text: "Look at the token counter in the sidebar. Right now the context is nearly empty — just the system prompt. Watch it grow as we load files."

  - type: user
    content: "Read the meeting notes file."

  - type: tool_call
    tool: scratchpad_read
    args:
      filename: "meeting-notes.txt"

  - type: tool_result
    tool: scratchpad_read
    content: |
      Product Team Weekly — 14 Jan 2026
      Attendees: Maya (PM), Dev team, UX team

      Decisions:
      - Ship v2.3 by end of month
      - Postpone dark mode to v2.4
      - Hire contract QA for launch sprint

      Action items:
      - Maya: Finalise launch checklist by Friday
      - Dev: Fix critical auth bug (#4521)
      - UX: Update onboarding flow mockups

  - type: assistant
    content: "Read the meeting notes. Key points: v2.3 shipping end of month, dark mode postponed, and there's a critical auth bug to fix. What would you like to know?"
    tokens: 38

  - type: note
    text: "Check the token counter — it jumped. That one file plus the conversation added tokens. Now let's load the other two files and see how context accumulates."

  - type: user
    content: "Now read the strategy doc and customer feedback too. Then tell me: is the product team aligned with the strategy?"

  - type: tool_call
    tool: scratchpad_read
    args:
      filename: "strategy-doc.txt"

  - type: tool_result
    tool: scratchpad_read
    content: |
      Q1 2026 Product Strategy

      Vision: Become the default collaboration tool for mid-market teams.

      Three Pillars:
      1. Speed — reduce page load times to under 200ms
      2. Integration — ship Slack, Teams, and Jira connectors
      3. Self-serve — enable team admin without IT involvement

      Target metrics:
      - 40% increase in weekly active users
      - Net Promoter Score above 50
      - Churn below 3% monthly

      Competitive landscape:
      - Notion expanding into project management
      - Asana adding AI features
      - Monday.com aggressive on pricing

      Our advantage: simplicity and speed for non-technical teams.

  - type: tool_call
    tool: scratchpad_read
    args:
      filename: "customer-feedback.txt"

  - type: tool_result
    tool: scratchpad_read
    content: |
      Recent NPS Responses (Jan 2026)

      Score 9: "Finally a tool my whole team actually uses. The simplicity is key."
      Score 8: "Love it but need better mobile app. Hard to use on the go."
      Score 3: "Integration with our existing tools is painful. Had to build workarounds."
      Score 10: "Replaced 3 tools with this one. Saving hours per week."
      Score 4: "Search is broken for large workspaces. Can't find anything."
      Score 9: "Best onboarding experience I've seen. Team was up in 20 minutes."
      Score 2: "No SSO support is a dealbreaker for enterprise. Had to cancel."
      Score 7: "Good but not great. Missing Gantt charts for project tracking."

  - type: assistant
    content: "**Partially aligned, with one major gap.**\n\nThe strategy has three pillars: Speed, Integration, and Self-serve.\n\n**Aligned:** The meeting focuses on shipping v2.3 (speed/features) and fixing the auth bug (stability). Customer feedback confirms simplicity is the strength — scores of 9 and 10 praise ease of use.\n\n**Misaligned:** The strategy's second pillar is Integration, but the meeting doesn't mention it at all. Meanwhile, customer feedback explicitly flags this: 'Integration with existing tools is painful' (Score 3) and 'No SSO support is a dealbreaker' (Score 2). These are the lowest scores in the dataset.\n\n**Recommendation:** The team should prioritise the integration pillar in the next sprint. Customer pain and strategic priority are aligned, but the weekly execution isn't reflecting it."
    tokens: 165

  - type: note
    text: "Look at the token counter now versus when we started. Three files plus the conversation has consumed a meaningful chunk of context. In a real workflow with larger documents, this adds up fast — which is why relevance matters more than volume."
```

```quiz
id: context-window-understanding
type: multiple-choice
question: "Why have larger context windows changed how we work with AI?"
options:
  - "They make the model smarter"
  - "They shift the constraint from 'what fits' to 'what's relevant'"
  - "They eliminate the need for good prompts"
  - "They allow the model to access the internet"
answer: 1
explanation: "Larger context windows mean we can now include much more information, but the challenge shifts to ensuring what we include is relevant and well-organized."
```
