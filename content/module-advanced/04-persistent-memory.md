---
title: "Persistent Memory"
duration: "10m"
tags: [memory, preferences, context]
---

# Persistent Memory

Tell AI about yourself once. It remembers forever.

## The Problem

Every new conversation starts fresh. AI doesn't know:
- Who you are
- What you're working on
- How you like things done
- Your role and context

So you explain it. Again. And again.

## The Solution

Persistent memory files (sometimes called Claude.md or similar) store information that AI loads automatically at the start of every conversation.

Write it once. Every session starts with that context already loaded.

## What Goes in Persistent Memory

**About you**
Your role, responsibilities, expertise area.

**Your preferences**
- Formatting preferences (bullet points vs paragraphs)
- Language preferences (UK vs US English)
- Tone preferences (formal vs casual)

**Your context**
Current projects, team information, company background.

**Standing instructions**
Things you always want done a certain way.

## Example Content

```
# About Me
Senior Product Manager at TechCorp.
Working on the mobile app team.
Reports to: VP of Product

# My Preferences
- Bullet points over paragraphs
- UK English spelling
- Concise, direct tone
- Include action items at the end of summaries

# Current Focus
Q1 launch of payment features.
Key stakeholders: Engineering, Design, Finance.
Main risk: Third-party payment provider integration.

# Standing Instructions
Always format dates as DD/MM/YYYY.
When drafting emails, keep under 200 words.
```

## Benefits

**No re-explaining**
"I'm a product manager working on..." is already known.

**Consistent output**
Your preferences are applied automatically.

**Accumulated context**
Add information over time. Context grows.

**Project awareness**
AI knows what you're working on without being told each time.

```callout
type: tip
title: "Build Incrementally"
content: "You don't need to write everything at once. Start with basics. Each time you find yourself explaining something repeatedly, add it to your persistent memory."
```

```quiz
id: persistent-memory-benefit
type: multiple-choice
question: "What problem does persistent memory solve?"
options:
  - "AI running out of context window space"
  - "Having to re-explain who you are and your preferences in every conversation"
  - "AI forgetting what it said earlier in the same conversation"
  - "Slow AI response times"
answer: 1
explanation: "Persistent memory solves the problem of starting every conversation fresh. Your context, preferences, and standing instructions load automatically, so you don't re-explain them every time."
```
