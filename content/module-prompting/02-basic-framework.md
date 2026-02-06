---
title: "The Basic Framework"
duration: "15m"
tags: [framework, structure, components]
---

# A Framework for Structuring Requests

Here's a basic framework for structuring AI requests. You don't need all five elements every time, but thinking through them helps.

## The Five Elements

**Role** — Who should the AI be? What expertise should it bring?
- "You are a senior financial analyst"
- "Act as a technical writer for developer documentation"

**Task** — What specific thing should it do?
- Be specific about the deliverable
- "Analyse this quarterly report and summarise key trends"

**Context** — What background information does it need?
- Company context, prior decisions, relevant history
- Data, documents, examples

**Format** — How should the output be structured?
- Bullet points, paragraphs, tables, code
- Word count, number of items

**Constraints** — What limits or requirements apply?
- Topics to avoid
- Tone requirements
- Compliance needs

## Bad Prompt vs Good Prompt

Let's compare two approaches to the same task.

**Bad prompt:**
```
Write me a report about Q4.
```

Problems:
- What Q4? Which year?
- Which company?
- What kind of report?
- For what audience?
- What format? How long?

The AI has to guess at everything.

**Good prompt:**
```
Write a Q4 2025 executive summary for Acme Corp's board meeting.

Include: revenue trends, key wins, risks.

Format: 1 page, bullet points, professional tone.

Context: [attached Q4 data]
```

This prompt specifies:
- Specific timeframe and company
- Clear deliverable type (executive summary)
- Defined audience (board)
- Explicit format
- Required content

```callout
type: tip
title: "The Key Insight"
content: "This isn't about clever wording. It's about complete information. The AI can't guess what you mean — tell it explicitly."
```

## When to Use What

Not every request needs all five elements.

**Quick questions:** Just ask. "What's the capital of France?"

**Simple tasks:** Task + Format is often enough. "List the top 5 features of Python 3.12 as bullet points."

**Complex work:** Use the full framework. The more consequential the output, the more complete your specification should be.

```quiz
id: framework-elements
type: multiple-choice
question: "Which element of the framework tells the AI what expertise to bring?"
options:
  - "Task"
  - "Context"
  - "Role"
  - "Constraints"
answer: 2
explanation: "The Role element defines who the AI should be and what expertise it should apply. 'You are a senior financial analyst' or 'Act as a technical writer' sets the frame for how the AI approaches the task."
```
