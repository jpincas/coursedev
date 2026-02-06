---
title: "Context Decay and Fresh Starts"
duration: "10m"
tags: [context-decay, conversations, best-practices]
---

# Why Long Conversations Degrade

Here's a common frustration: AI seems to get worse the longer you talk to it. This isn't your imagination. It's context decay.

## The Degradation Pattern

**Turns 1-10: Fresh and Clear**
- Context is fresh
- Instructions are clear and uncontradicted
- Output quality is high

**Turns 20-30: Starting to Blur**
- Context is filling up
- Earlier details are competing with newer ones
- Occasional confusion or inconsistency

**Turns 50+: Saturated**
- Context is cluttered or truncated
- Contradictions have accumulated
- Quality noticeably drops

## Why This Happens

Several factors contribute:

**Information density increases**
More messages mean more information competing for attention. Important early instructions get diluted.

**Contradictions accumulate**
You might have said "use formal tone" early on, then casually said "yeah just make it casual" later. Both are in context.

**Context truncation**
Very long conversations may get truncated — older messages dropped to fit newer ones. Your original instructions might disappear.

**Attention dilution**
The model has limited attention. With more in context, less attention goes to any single piece.

```callout
type: warning
title: "The Symptom"
content: "When the model starts 'forgetting' things you told it earlier, or contradicting previous outputs, you're experiencing context decay."
```

## The Solution: Fresh Starts

Don't fight context decay. Work with it.

**New task? New conversation.**

When you're shifting to a different task or topic, start a fresh conversation. You'll get:
- Clean context
- Full attention on your new instructions
- No contradictions from previous exchanges

```callout
type: tip
title: "Think in Sessions"
content: "Treat conversations like work sessions. One focused task per conversation. When you're done with that task, start fresh for the next one."
```

## When to Preserve Context

Sometimes you do want conversation continuity:
- Iterating on the same document
- Building on previous work in the same session
- Asking follow-up questions about the same topic

In these cases, keep going. But be aware that quality may degrade over very long sessions.

## Markdown: The Universal Language

One tool that helps combat context decay is structured formatting. Markdown is the universal language of AI context.

**Why Markdown?**
- Most AI training data uses it
- Provides clear structure the model understands well
- Easy for humans to read and write too

**The pattern:** Structure your inputs in Markdown. Get structured outputs back.

Use:
- `# Headings` for hierarchy
- `- Bullets` for lists
- `**Bold**` for emphasis
- ` ``` ` for code blocks

Well-structured context remains clearer even as conversations grow.

```quiz
id: context-decay-solution
type: multiple-choice
question: "What's the recommended approach when you notice AI output quality degrading in a long conversation?"
options:
  - "Add more detailed instructions to fix it"
  - "Start a fresh conversation for your next task"
  - "Ask the AI to remember your earlier instructions"
  - "Use shorter messages to save context space"
answer: 1
explanation: "Context decay is natural and hard to fight. The most effective solution is to start fresh conversations for new tasks, giving you clean context and full attention on current instructions."
```
