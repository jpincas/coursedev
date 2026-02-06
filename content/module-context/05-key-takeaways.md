---
title: "Key Takeaways: Context"
duration: "5m"
tags: [summary, takeaways]
---

# Key Takeaways

Let's summarize what we've learned about context.

## 1. Everything Is Context

The model only knows what's in the context window. If it's not there, it doesn't exist. Your inputs, files, conversation history, and system prompts — that's everything the model sees.

## 2. Structure Matters

Use Markdown for clear hierarchy. Put important instructions first. Explicit formatting beats implicit expectations. Well-structured input produces well-structured output.

## 3. Fresh Starts Win

New task? New conversation. Don't fight context decay. Long conversations degrade naturally. Work with it, not against it.

## 4. Quality In = Quality Out

Your context bounds your results. Rich, relevant context produces rich output. Vague or missing context produces poor output.

```callout
type: tip
title: "The Diagnostic Habit"
content: "When output is bad, always ask: What was the context? What did the model actually have to work with? Nine times out of ten, the answer is there."
```

## Practical Applications

**Before you prompt:**
- Do I have all the relevant information ready to provide?
- Is my request clear and structured?
- Have I included examples of what I want?

**When output is poor:**
- What context was the model missing?
- Did I provide contradictory instructions?
- Was the conversation too long and cluttered?

**For important tasks:**
- Start fresh conversations
- Provide all necessary context upfront
- Structure your inputs with Markdown

```quiz
id: context-takeaway
type: multiple-choice
question: "What should be your first diagnostic question when AI output is poor?"
options:
  - "Is the model smart enough for this task?"
  - "Should I use a different AI model?"
  - "What was the context the model was working with?"
  - "Did I ask politely enough?"
answer: 2
explanation: "When output is poor, the first question should always be about context. What did the model see? What was missing? What might have been contradictory? Context problems explain most quality issues."
```

## Up Next

Now that you understand context, let's talk about how to structure your requests effectively. **Prompting** is the art of giving clear instructions and context that produce the results you want.
