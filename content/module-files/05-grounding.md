---
title: "Grounding: Your Specific Knowledge"
duration: "10m"
tags: [grounding, context, specifics]
---

# Grounding

Grounding means giving AI access to your specific knowledge — the information that makes your context unique.

## The Problem

AI knows general things. It doesn't know:
- Your company's specific policies
- Your project's history
- Your industry's particular terminology
- Your team's preferred formats
- Your organisation's style guide

When AI lacks this specific knowledge, it produces generic output that doesn't quite fit your context.

## The Solution

Provide the specific documents, data, and context that contain your knowledge.

**Upload your style guide**
AI writes in your voice, uses your terminology, follows your conventions.

**Provide past reports**
AI matches your format, structure, and level of detail.

**Include policy documents**
AI cites correct procedures, references actual policies.

**Share templates**
AI follows your established patterns.

```callout
type: tip
title: "Show, Don't Tell"
content: "Rather than describing your preferred style, provide an example of work you liked. AI will learn from the actual artifact more reliably than from your description of it."
```

## Grounding in Practice

**Without grounding:**
"Write a client email about the project delay."
Result: Generic corporate-speak that doesn't match your voice.

**With grounding:**
"Write a client email about the project delay. Here are three previous client emails I've sent for tone and style." *[attach emails]*
Result: Email matches your actual communication style.

## What to Provide

Depending on the task, consider including:

| Task Type | Grounding Materials |
|-----------|---------------------|
| Writing | Style guide, example pieces, brand guidelines |
| Analysis | Past analyses, template formats, metric definitions |
| Presentations | Previous decks, company templates, brand assets |
| Code | Existing codebase, coding standards, architecture docs |
| Reports | Previous reports, formatting preferences, required sections |

## The Payoff

Grounded AI output:
- Fits your context
- Uses correct terminology
- Follows your conventions
- Requires less editing
- Feels like it came from your team

Ungrounded AI output:
- Feels generic
- Uses unfamiliar phrasing
- Misses context-specific details
- Requires significant rework

The investment in providing grounding materials pays off in output quality.

```quiz
id: grounding-purpose
type: multiple-choice
question: "What problem does grounding solve?"
options:
  - "AI running out of context window space"
  - "AI producing generic output that doesn't fit your specific context"
  - "AI taking too long to generate responses"
  - "AI hallucinating incorrect facts"
answer: 1
explanation: "Grounding solves the problem of generic output. AI knows general patterns but not your specific context — your style, terminology, policies, and preferences. Providing grounding materials lets AI produce output that fits your actual situation."
```
