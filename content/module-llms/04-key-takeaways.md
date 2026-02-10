---
title: "Key Takeaways: How LLMs Work"
duration: "10m"
tags: [summary, takeaways]
---

# Key Takeaways: How LLMs Work

These three mental models will shape every interaction you have with AI for the rest of this training and beyond. They explain why AI succeeds, why it fails, and what you can do about it.

## 1. Prediction Engines, Not Knowledge Bases

LLMs are prediction engines. Not knowledge bases. Not reasoners. Prediction machines.

They predict what token comes next based on patterns learned from training data. Everything else emerges from this core capability.

## 2. Probability, Not Truth

The output is optimized for plausibility, not accuracy.

When you ask a question, the model generates what a good answer would look like, based on patterns. It doesn't "look up" the answer or verify it against ground truth.

## 3. Hallucinations Are Inherent — Always Verify

Hallucinations aren't bugs to be fixed. They're a direct consequence of the mechanism. The rates vary (from 0.7% to 80%+ depending on model and domain), but they're always present.

The model can't verify itself. That's your job. For anything consequential, use multi-model cross-validation, ask the AI to verify its own reasoning, and provide grounding context. These techniques reduce errors but don't eliminate them. Plan verification time into your workflow.

![Key Takeaways Summary](/content/module-llms/images/key-takeaways.svg)

```callout
type: tip
title: "The Practical Implication"
content: "Understanding this foundation will inform everything else today. When AI fails, ask: Was this a prediction problem? Did it lack the right context? Was it pattern-matching to the wrong thing?"
```

## What About "Intelligence"?

Is it intelligent? The honest answer: It depends on what you mean by intelligence.

**It IS:**
- Remarkable pattern recognition
- Capable of human-level task performance
- Able to reason through problems
- Creative and flexible

**It ISN'T:**
- Understanding in the way you understand
- Conscious or self-aware
- Reliable factual knowledge
- Consistent reasoning (it can solve a hard problem and fail an easy one depending on framing)

For practical purposes, what matters is: It can do useful work. Treat it as a very capable tool with specific limitations, not as a thinking being or as a simple text generator.

It's something new.

## Verification in Practice

The page says "the model can't verify itself" — but actually, there's a useful trick. You can ask the AI to check its own reasoning. It won't catch everything, but it catches more than you'd expect. Watch.

```agent
id: cross-validation-demo
title: "Self-Verification in Action"
model_label: "Claude"

system: |
  You are a research assistant. Help users analyse information
  accurately. When asked to verify, critically examine your own
  previous statements for accuracy and flag anything uncertain.

scratchpad:
  "quarterly-data.csv": |
    Region,Q1_Revenue,Q2_Revenue,Q3_Revenue,Q4_Revenue
    North,2.4M,2.8M,3.1M,3.6M
    South,1.8M,1.9M,2.0M,2.1M
    East,3.2M,2.9M,3.4M,3.8M
    West,1.1M,1.3M,1.2M,1.5M
    Total,8.5M,8.9M,9.7M,11.0M

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
    text: "This demo shows a practical verification technique. The AI will analyse some data, then we'll ask it to verify its own claims. Watch what happens."

  - type: user
    content: "Read the quarterly data file and give me a quick executive summary with key trends."

  - type: tool_call
    tool: scratchpad_read
    args:
      filename: "quarterly-data.csv"

  - type: tool_result
    tool: scratchpad_read
    content: |
      Region,Q1_Revenue,Q2_Revenue,Q3_Revenue,Q4_Revenue
      North,2.4M,2.8M,3.1M,3.6M
      South,1.8M,1.9M,2.0M,2.1M
      East,3.2M,2.9M,3.4M,3.8M
      West,1.1M,1.3M,1.2M,1.5M
      Total,8.5M,8.9M,9.7M,11.0M

  - type: assistant
    content: "Here's the executive summary:\n\n**Overall:** Revenue grew 29% from Q1 ($8.5M) to Q4 ($11.0M), showing strong upward momentum.\n\n**Key trends:**\n- **North** is the growth leader — 50% increase Q1 to Q4, accelerating each quarter\n- **East** is the largest region by revenue ($3.8M in Q4) but had a Q2 dip before recovering\n- **South** shows steady but slow growth (17% over the year) — the weakest performer\n- **West** is smallest but grew 36% — punching above its weight\n\n**Concern:** The East region's Q2 decline ($3.2M → $2.9M) warrants investigation — was this seasonal or a lost client?"

  - type: note
    text: "Looks thorough and confident. But is it all accurate? Let's ask it to verify."

  - type: user
    content: "Now go back and verify every number and percentage you just cited. Check each one against the source data."

  - type: assistant
    content: "Let me verify each claim:\n\n**Overall 29% growth:** Q1 $8.5M → Q4 $11.0M. That's ($11.0 - $8.5) / $8.5 = 29.4%. ✓ Correct.\n\n**North 50% increase:** Q1 $2.4M → Q4 $3.6M. That's ($3.6 - $2.4) / $2.4 = 50%. ✓ Correct.\n\n**East largest region at $3.8M Q4:** Checking — North $3.6M, South $2.1M, East $3.8M, West $1.5M. ✓ East is largest.\n\n**East Q2 dip $3.2M → $2.9M:** Q1 was $3.2M, Q2 was $2.9M. ✓ Correct, that is a decline.\n\n**South 17% growth:** Q1 $1.8M → Q4 $2.1M. That's ($2.1 - $1.8) / $1.8 = 16.7%, rounds to 17%. ✓ Correct.\n\n**West 36% growth:** Q1 $1.1M → Q4 $1.5M. That's ($1.5 - $1.1) / $1.1 = 36.4%. ✓ Correct.\n\n**All figures verified.** In this case everything checks out, but this process is worth doing — I could have miscalculated a percentage or misread a column."

  - type: note
    text: "This time everything was accurate. But notice the technique: asking the AI to re-examine its own output against the source data forces it to do a second pass. In real work — especially with more complex data — this second pass frequently catches errors the first pass missed."

  - type: note
    text: "This isn't foolproof. The same blindspots can persist across passes. For truly critical work, use a second model (GPT to check Claude, or vice versa) or verify the numbers yourself. But self-verification is a useful first line of defence."
```

```quiz
id: llms-takeaway
type: multiple-choice
question: "You ask AI to summarise a 50-page contract. The summary is well-written and plausible. What is the most appropriate next step?"
options:
  - "Verify key claims against the source document, since plausibility does not equal accuracy"
  - "Accept the summary because AI excels at document synthesis tasks"
  - "Run the same query through a second AI model and accept whichever summary is longer"
answer: 0
explanation: "Well-written and plausible does not mean accurate. The summary could contain hallucinated clauses, misattributed terms, or invented details that sound right. Verifying key claims against the source document is essential, especially for consequential work like contracts."
```

## Up Next

Now that you understand the mechanism, let's talk about the most important concept for getting good results: **Context**.

Context is everything. The quality of what comes out is bounded by the quality of what goes in.
