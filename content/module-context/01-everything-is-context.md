---
title: "Everything Is Context"
duration: "15m"
tags: [context, fundamentals, core-concept]
---

# The Big Idea

If you remember one thing from this entire training, let it be this:

**Everything is context.**

Every word the model produces is a prediction based on what it's seen. You control what it sees.

```callout
type: tip
title: "The Central Insight"
content: "The quality of AI output is bounded by the quality of context you provide. Rich input produces rich output. Poor input produces poor output."
```

## What Does This Mean?

When you interact with an AI, the model generates its response based entirely on what's in its **context window**. This includes:

- **System prompt** — Hidden instructions about behavior
- **Conversation history** — Everything said so far in this chat
- **Files and documents** — Content you've uploaded or referenced
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

Understanding context explains most AI behavior:

**Good output?** The model had rich, relevant context.

**Poor output?** The model was missing something. Wrong examples. Missing background. Unclear instructions.

**Inconsistent output?** Context is changing between requests, or earlier context is contradicting current instructions.

Nine times out of ten, improving the context improves the output.

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
