---
title: "Hallucinations and Confident Wrongness"
duration: "15m"
tags: [hallucinations, verification, limitations]
---

# Hallucinations

When AI generates plausible-sounding but false information, we call it a "hallucination."

```callout
type: danger
title: "This Is Not a Bug"
content: "Hallucinations are not bugs to be fixed. They're a direct consequence of how prediction works. The model doesn't know if something is true — it knows if something sounds like it could be true."
```

If you ask about a topic and plausible-sounding wrong information exists in training data, it can reproduce it. If you ask about something rare, it might pattern-match to something similar but wrong.

## Why Confident Wrongness Happens

Let's break down why this happens:

**Probability, not truth**
The output is based on statistical patterns, not factual verification. The model outputs what's statistically likely, not what's verified.

**No external lookup**
Unlike you, the model can't Google something mid-response. It only has what's in its training data and what you've provided.

**Pattern matching to plausibility**
It generates text that matches patterns of "how accurate-sounding statements on this topic look."

**No uncertainty signal**
The text sounds equally confident whether it's right or wrong. There's no built-in indicator of reliability.

## The Implication

This is why verification matters. And why providing good context matters so much.

When you provide grounded facts rather than relying on probabilistic patterns, you dramatically improve accuracy. You're giving it real information to work with instead of relying on "what sounds right."

```callout
type: warning
title: "Your Job: Verification"
content: "The model can't verify itself. It can't check against ground truth. That's your job. Always verify important facts, especially for anything consequential."
```

```quiz
id: llms-hallucinations
type: multiple-choice
question: "Why do LLMs sometimes generate false but plausible-sounding information?"
options:
  - "Because they're programmed to lie"
  - "Because they optimize for plausibility, not truth"
  - "Because they only have access to false information"
  - "Because they run out of memory"
answer: 1
explanation: "LLMs optimize for what sounds plausible based on training patterns, not for what is factually true. They can't verify their own claims against external reality."
```
