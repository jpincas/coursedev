---
title: "The Three Questions"
duration: "15m"
tags: [questions, specification, clarity]
---

# The Three Questions

Here's an even simpler framework. Before any AI task, ask yourself these three questions.

## 1. What Does "Done" Look Like?

If you can't describe the end state, the AI can't produce it.

**Vague:**
"Help me with my presentation"

**Better:**
"Create slides for my presentation"

**Clear:**
"Create a 10-slide presentation on Q4 results for the sales team, with one key metric highlighted per slide, ending with Q1 priorities"

```callout
type: tip
title: "The Verification Test"
content: "Could you verify that the work is done correctly? If you can't describe what 'correct' looks like, your specification isn't clear enough."
```

## 2. What Context Does It Need?

What information is required for success?

**Data and facts**
The actual information to work with — spreadsheets, documents, code.

**Background and situation**
Company context, project history, prior decisions. Why does this task exist?

**Audience**
Who will use this? What do they already know? Executive? Developer? Customer?

**Examples**
Similar work you liked. Templates. Style references. Show, don't just tell.

**Requirements**
Explicit must-haves. Compliance needs. Brand guidelines.

The more relevant context you provide, the better the output.

## 3. What Are the Boundaries?

What constraints or limits apply?

**Length**
Word count, page count, number of items. "About 500 words" is better than nothing.

**Tone**
Professional? Casual? Technical? Explain-like-I'm-five?

**Format**
Bullets? Paragraphs? Table? Code? This dramatically affects usability.

**Inclusions**
What must be covered? "Must mention the new pricing model."

**Exclusions**
What to avoid? "Don't discuss the merger." "Skip implementation details."

**Constraints**
Technical limits, brand requirements, compliance needs.

Boundaries prevent the AI from going off in directions you don't want.

## Putting It Together

**Task:** Analyse customer feedback

**Question 1 — Done:**
A summary of the top 5 complaint categories with representative quotes

**Question 2 — Context:**
- CSV file with 500 feedback entries
- This is for the product team
- We're focused on our mobile app experience

**Question 3 — Boundaries:**
- Max 1 page
- Include frequency counts
- Professional but concise tone

**Result prompt:**
```
Analyse this customer feedback CSV (attached).
Identify the top 5 complaint categories with frequency counts.
For each category, provide 2-3 representative quotes.
Focus on mobile app experience issues.
Format as a brief report for the product team, max 1 page.
```

```quiz
id: three-questions-test
type: multiple-choice
question: "Which question helps you define success criteria for an AI task?"
options:
  - "What context does it need?"
  - "What are the boundaries?"
  - "What does 'done' look like?"
  - "What role should AI play?"
answer: 2
explanation: "The question 'What does done look like?' forces you to define the end state clearly. If you can describe what correct completion looks like, you've defined success criteria."
```
