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

The model sees a sequence of text:

> "The capital of France is ▊"

And it predicts probabilities for what word comes next:

- **"Paris"** — 97.3% probability
- **"Lyon"** — 0.8%
- **"Berlin"** — 0.2%

It picks the most likely continuation and adds it. Then it does it again. And again. Word by word, token by token, building up a response.

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

```quiz
id: llms-core-mechanism
type: multiple-choice
question: "What is the core mechanism of a Large Language Model?"
options:
  - "Looking up answers in a database"
  - "Predicting the next token based on patterns"
  - "Running logical rules on input"
  - "Searching the internet in real-time"
answer: 1
explanation: "LLMs work by predicting what token (word or piece of text) comes next, based on patterns learned during training. They don't look things up or search — they predict plausible continuations."
```
