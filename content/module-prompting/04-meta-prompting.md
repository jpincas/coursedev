---
title: "Meta-Prompting"
duration: "10m"
tags: [meta-prompting, iteration, techniques]
---

# Meta-Prompting

Here's a powerful technique: instead of writing the perfect prompt yourself, ask AI to help you write it.

## The Basic Pattern

```
I want to [goal]. Help me write a detailed prompt
that will get the best result.
Ask me clarifying questions first.
```

The AI will ask what it needs to know. Your answers become the prompt.

This is especially useful when you're not sure what you need.

## Meta-Prompting in Action

**You:** "I need to analyse customer feedback. Help me write a prompt for that."

**AI:** "Let me ask a few questions:
- What format is the feedback in?
- What specific insights are you looking for?
- How should the analysis be structured?
- Who will read the results?"

**You:** "It's in a CSV. Looking for common complaints. Need a summary for the product team."

**AI:** "Here's your prompt: 'Analyse this customer feedback CSV. Identify the top 5 complaint categories with frequency counts. For each, provide 2-3 representative quotes. Format as a brief report for the product team, max 1 page.'"

```callout
type: tip
title: "Why This Works"
content: "The AI's questions reveal what information is needed. Each answer adds specificity. The final prompt is often better than what you would have written from scratch."
```

## When to Use Meta-Prompting

**You're unsure about scope**
Let the AI's questions help you figure out what you actually need.

**Complex or multi-faceted tasks**
The questioning process surfaces aspects you might not have considered.

**Unfamiliar domains**
AI can ask expert-level questions about topics where you're not the expert.

**Reusable prompts**
If you'll do similar tasks repeatedly, invest in getting a great prompt.

## The Key Insight

Don't describe the process. Describe the outcome.

**Bad:**
"First analyse the data, then identify trends, then summarise them, then format as bullet points, then..."

**Good:**
"Produce a trend analysis with the top 5 insights as bullet points."

AI is capable of figuring out the process. That's its job. Your job is defining what done looks like.

This is the shift from prompting to delegation.

```quiz
id: meta-prompting-purpose
type: multiple-choice
question: "What's the main benefit of meta-prompting?"
options:
  - "It makes the AI faster"
  - "It helps you discover what information the AI needs"
  - "It reduces the cost of API calls"
  - "It makes prompts shorter"
answer: 1
explanation: "Meta-prompting helps you discover what information and specificity the AI needs to produce good output. The AI's questions reveal gaps in your initial request, and your answers build a more complete prompt."
```
