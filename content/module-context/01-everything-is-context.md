---
title: "Everything Is Context"
duration: "15m"
tags: [context, fundamentals, core-concept]
---

# The Big Idea

If you remember one thing from this entire training, let it be this:

**Everything is context.**

Every word the model produces is a prediction based on what it's seen. You control what it sees.

This isn't just about writing better prompts. It's about **context engineering** — designing the entire information environment surrounding the model: memory, retrieved data, tools, state, metadata, and structured inputs.

Former OpenAI researcher Andrej Karpathy frames it perfectly: **"The LLM is like the CPU, and its context window is like RAM."** Context engineering decides what fills that working memory.

```callout
type: tip
title: "The Conceptual Shift of 2025"
content: "Context engineering has replaced prompt engineering as the core discipline. The smartest AI practitioners don't ask better questions — they build better conditions for answers to emerge."
```

## What Context Engineering Means

When you interact with an AI, the model generates its response based entirely on what's in its **context window**. Context engineering is the practice of deliberately shaping this environment. This includes:

- **System prompts** — Hidden instructions about behaviour (the foundation layer)
- **Persistent instructions** — CLAUDE.md files, project knowledge bases, custom instructions
- **Conversation history** — Everything said so far in this chat
- **Files and documents** — Content you've uploaded or referenced
- **Tools and capabilities** — What the model can access and execute
- **Your current message** — What you just typed

The model processes all of this together and predicts what response should come next.

## Nothing Else Exists

Here's the critical point: **Nothing outside the context window exists to the model.**

If it's not in context, it's not there. The model can't access:
- Information you haven't provided
- Details from other conversations
- External websites or databases (unless it has tools for that)
- Your intentions that you haven't stated

```callout
type: warning
title: "The Diagnostic Question"
content: "When AI output is bad, the first question is always: What was the context? What did the model actually see?"
```

## Why This Matters Practically

Understanding context explains most AI behaviour:

**Good output?** The model had rich, relevant context.

**Poor output?** The model was missing something. Wrong examples. Missing background. Unclear instructions.

**Inconsistent output?** Context is changing between requests, or earlier context is contradicting current instructions.

Nine times out of ten, improving the context improves the output.

```agent
id: context-changes-everything
title: "Context Changes Everything"
model_label: "Claude"

system: |
  You are a professional business writing assistant. You help create
  reports, summaries, and communications for corporate audiences.

scratchpad:
  "project-brief.md": |
    Project: Q4 Campaign Rollout for Acme Industries

    Client: Acme Industries Ltd.
    Budget: £450,000
    Timeline: October–December 2025

    Objectives:
    - Launch new product line across UK markets
    - Achieve 15% brand awareness lift
    - Generate 5,000 qualified leads

    Key deliverables:
    - Multi-channel campaign strategy
    - Content creation and distribution
    - Performance analytics and reporting
  "style-guide.md": |
    Acme Industries Writing Style Guide

    - Use British English spelling and conventions
    - Keep paragraphs concise (3-4 sentences maximum)
    - Avoid marketing jargon and buzzwords
    - Use active voice
    - Lead with data and specifics
    - Professional but approachable tone
  "quarterly-data.csv": |
    Quarter,Leads,Awareness,Budget_Spent
    Q1,892,8.2%,£95000
    Q2,1247,11.5%,£118000
    Q3,1653,13.8%,£142000
    Q4,2103,15.4%,£95000

tools:
  - scratchpad_read

visibility:
  system_prompt: toggleable
  tool_calls: visible
  full_context: hidden

sidebar:
  width: "45%"
  start_open: true

script:
  - type: note
    text: "Watch how the same question produces completely different results depending on context."

  - type: user
    content: "Write a quarterly summary for the board."

  - type: assistant
    content: |
      Here's a quarterly summary for your board meeting:

      Our team has made solid progress this quarter across key metrics. We've seen growth in lead generation and brand awareness continues to trend upward. The marketing initiatives have been performing well within budget parameters.

      Looking ahead, we're positioned to build on these gains and continue our momentum into the next quarter. The foundational work we've done should support our ongoing objectives.

  - type: note
    text: "Generic and vague. No specifics, no data, American English. Now let's provide context..."

  - type: user
    content: "Here are our project files. Now write the same quarterly summary using this context."

  - type: assistant
    content: "Let me review the project materials to create a data-driven summary."

  - type: tool_call
    tool: scratchpad_read
    args:
      filename: "project-brief.md"

  - type: tool_result
    tool: scratchpad_read
    content: |
      Project: Q4 Campaign Rollout for Acme Industries

      Client: Acme Industries Ltd.
      Budget: £450,000
      Timeline: October–December 2025

      Objectives:
      - Launch new product line across UK markets
      - Achieve 15% brand awareness lift
      - Generate 5,000 qualified leads

      Key deliverables:
      - Multi-channel campaign strategy
      - Content creation and distribution
      - Performance analytics and reporting

  - type: tool_call
    tool: scratchpad_read
    args:
      filename: "style-guide.md"

  - type: tool_result
    tool: scratchpad_read
    content: |
      Acme Industries Writing Style Guide

      - Use British English spelling and conventions
      - Keep paragraphs concise (3-4 sentences maximum)
      - Avoid marketing jargon and buzzwords
      - Use active voice
      - Lead with data and specifics
      - Professional but approachable tone

  - type: tool_call
    tool: scratchpad_read
    args:
      filename: "quarterly-data.csv"

  - type: tool_result
    tool: scratchpad_read
    content: |
      Quarter,Leads,Awareness,Budget_Spent
      Q1,892,8.2%,£95000
      Q2,1247,11.5%,£118000
      Q3,1653,13.8%,£142000
      Q4,2103,15.4%,£95000

  - type: assistant
    content: |
      Q4 Campaign Summary for Acme Industries Board

      We exceeded our objectives for the Q4 campaign rollout. Lead generation reached 2,103 qualified leads—42% above target. Brand awareness lifted to 15.4%, meeting our 15% goal.

      Budget performance was exceptional. We delivered full campaign objectives whilst spending only £95,000 of the allocated £450,000, creating significant cost savings without compromising results.

      The multi-channel strategy proved effective. Each quarter showed consistent growth: Q1 delivered 892 leads at 8.2% awareness, Q2 reached 1,247 leads at 11.5%, Q3 achieved 1,653 leads at 13.8%, culminating in Q4's strong finish.

  - type: note
    text: "Same question, same model, same user—dramatically different output. The only variable was context."
```

```quiz
id: context-core-concept
type: multiple-choice
question: "What determines the quality of AI output?"
options:
  - "The specific AI model being used"
  - "The quality and completeness of the context provided"
  - "The time of day the request is made"
  - "How politely you phrase the request"
answer: 1
explanation: "While model choice matters, the quality of context you provide is the primary factor determining output quality. The same model produces dramatically different results with different context."
```
