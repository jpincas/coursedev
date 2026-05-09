---
title: "LLMs Are Prediction Engines"
duration: "20m"
tags: [llms, prediction, next-token]
---

# What Is a Large Language Model?

Let's start with the fundamental question: What is an LLM?

**It's a prediction engine.**

That's it. That's the core of it. Everything else — the apparent intelligence, the creativity, the understanding — emerges from one capability: predicting what comes next.

## Next-Token Prediction

Here's how it works at the most basic level.

The model sees a sequence of text and predicts what comes next — not just once, but token by token, building up the response one step at a time.

![Next-Token Prediction Chain](/content/module-llms/images/token-prediction-chain.svg)

At each step, the model evaluates all possible next tokens and assigns probabilities. It picks the highest probability option and adds it to the sequence. Then it does it again with the new, longer sequence. And again. Word by word, token by token, building up a response.

```callout
type: info
title: "Token by Token"
content: "Watch it build up: 'The' → 'capital' → 'of' → 'France' → 'is' → 'Paris' → '.' Each step, it looks at everything that came before and predicts what comes next."
```

Scale this up — billions of parameters, trained on trillions of words — and you get something that can write essays, code, analysis, creative fiction.

But it's still, fundamentally, next-token prediction.

## Why This Matters

Understanding the mechanism explains everything.

**It explains the strengths:**
- Excellent at pattern continuation
- Great with formats it's seen before
- Good at "what usually follows"
- Excellent at style matching

**It explains the weaknesses:**
- No access to external facts
- Can't verify its own claims
- Confident about wrong things
- Pattern-matches to plausible wrong answers

If you give it something that looks like patterns it's seen before, it excels. Professional email? Seen millions. Legal document format? Seen millions. Code in Python? Seen millions.

But it has no connection to ground truth. It doesn't "know" things — it predicts what text about things looks like.

## What AI Does Well (and What It Doesn't)

This prediction mechanism creates a fascinating pattern: **the hard problems are easy and the easy problems are hard.**

That insight comes from the Carnegie Endowment for International Peace, and it captures something fundamental about AI capability in 2026.

**AI excels at:**
- **Code generation** — 41% of all code written globally is now AI-generated
- **Pattern recognition** — document review, contract analysis, spotting anomalies
- **Synthesis** — combining information from multiple sources into coherent summaries
- **First drafts** — getting something usable down fast across any format

**AI struggles with:**
- **Hallucination drift** — accuracy degrades across very long chains of reasoning
- **Cultural nuance** — context-dependent meaning, subtle implications
- **Phrasing sensitivity** — minor changes to how you ask can produce wildly different results
- **Novel common-sense reasoning** — genuinely new situations where pattern-matching fails

The paradox: AI can write production-ready code for complex algorithms, but might confidently claim a fictitious historical event occurred because the description "sounds right."

```callout
type: tip
title: "Calibrating Your Expectations"
content: "AI is extraordinarily good at tasks with clear patterns and structures it has seen millions of times. It is less reliable when genuine novelty, cultural context, or factual grounding matters. Plan your work accordingly."
```

## Prediction in Action

Watch how changing the context changes what the model predicts. The same opening, two different contexts, two completely different outputs.

```agent
id: prediction-in-action
title: "Prediction Changes with Context"
model_label: "Claude"

system: |
  You are a helpful writing assistant. You complete tasks based on the
  context and instructions provided. Be concise and direct.

scratchpad:
  "company-overview.txt": |
    Meridian Technologies — Annual Report 2025

    Founded: 2019
    Industry: Cloud Infrastructure
    Employees: 340
    Revenue: $42M (up 28% YoY)
    Key Product: CloudBridge — automated migration platform

    Q4 Highlights:
    - Signed 3 Fortune 500 clients
    - Launched CloudBridge v3 with AI-assisted migration
    - Opened London office for European expansion
    - Named "Rising Star" by Gartner in cloud migration category

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
    text: "Watch how the same request produces completely different outputs depending on what context the model has seen. First, we'll ask for a summary without providing any information — to show what pure prediction looks like."

  - type: user
    content: "Write a one-paragraph summary of Meridian Technologies."

  - type: assistant
    content: "I don't have any information about Meridian Technologies in my training data or files. I'd need you to provide context — perhaps a company overview document or annual report — before I could write an accurate summary."

  - type: note
    text: "The model correctly indicated it lacks context. But watch what happens if we simulate a common mistake: asking without checking whether the company exists in training data. Let's reset and show the contrast by providing the context properly."

  - type: user
    content: "Let me start again. I have our company overview in a file. Read company-overview.txt and then write a one-paragraph summary."

  - type: tool_call
    tool: scratchpad_read
    args:
      filename: "company-overview.txt"

  - type: tool_result
    tool: scratchpad_read
    content: |
      Meridian Technologies — Annual Report 2025

      Founded: 2019
      Industry: Cloud Infrastructure
      Employees: 340
      Revenue: $42M (up 28% YoY)
      Key Product: CloudBridge — automated migration platform

      Q4 Highlights:
      - Signed 3 Fortune 500 clients
      - Launched CloudBridge v3 with AI-assisted migration
      - Opened London office for European expansion
      - Named "Rising Star" by Gartner in cloud migration category

  - type: assistant
    content: "Meridian Technologies is a cloud infrastructure company founded in 2019 that has grown to 340 employees and $42M in revenue, up 28% year-over-year. The company's flagship product, CloudBridge, is an automated migration platform that recently launched its third version with AI-assisted capabilities. In Q4 2025, Meridian signed three Fortune 500 clients, expanded into Europe with a London office, and earned Gartner's 'Rising Star' recognition in cloud migration."

  - type: note
    text: "Same model. Same request. Dramatically different output. The first response was the model predicting 'what does a generic company description look like?' The second was predicting 'what does a summary of THIS company look like, given these facts?' Context changed the prediction."

  - type: note
    text: "This is the core mechanism at work. The model doesn't 'know' about Meridian Technologies. It predicts what good text looks like given the patterns it has seen — and the context you provide shapes those predictions entirely."
```

```quiz
id: llms-core-mechanism
type: multiple-choice
question: "An LLM writes a factually incorrect but perfectly formatted legal brief. Which explanation best accounts for this?"
options:
  - "The model has seen millions of legal briefs and predicts plausible legal text, regardless of factual accuracy"
  - "The model's legal training data was outdated, causing it to reference superseded case law"
  - "The legal domain requires fact-checking capabilities that weren't included in the training process"
answer: 0
explanation: "LLMs are prediction engines. They predict what text looks like based on patterns, not what is true. A model that has seen millions of legal briefs will produce perfectly formatted legal text -- but it has no mechanism for verifying whether the cited cases or facts are real."
```
