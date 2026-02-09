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

## Meta-Prompting Live

Instead of just describing meta-prompting, watch it happen. The user doesn't know exactly what they need — so they ask the AI to help figure it out.

```agent
id: meta-prompting-demo
title: "AI Helps Write the Prompt"
model_label: "Claude"

system: |
  You are a prompt engineering assistant. When asked to help write
  a prompt, ask clarifying questions first to understand exactly
  what's needed. Then produce a complete, well-structured prompt.

tools: []

visibility:
  system_prompt: hidden
  tool_calls: hidden
  full_context: hidden

sidebar:
  width: "45%"
  start_open: true

script:
  - type: note
    text: "The user has a vague goal: they want to do something with customer feedback. Instead of writing a mediocre prompt, they ask the AI to help write a good one."

  - type: user
    content: "I want to do something useful with our customer feedback data. Help me write a prompt that will get the best analysis. Ask me questions first."

  - type: assistant
    content: "Good approach — let me ask a few questions to build the right prompt:\n\n1. **What format is the feedback in?** (CSV, survey responses, free-text reviews, support tickets?)\n2. **How much data?** (Dozens, hundreds, thousands of entries?)\n3. **What decisions will this inform?** (Product roadmap? Marketing? Support staffing?)\n4. **Who will read the analysis?** (Exec team? Product managers? The whole company?)\n5. **What have you tried before?** (Any previous analysis attempts, or starting fresh?)\n6. **Any specific concerns?** (Known issues you want validated? Hypotheses to test?)"

  - type: user
    content: "It's a CSV export from our NPS survey — about 800 responses from Q4. Mix of scores and free-text comments. I need this for the product team's Q1 planning session next week. The VP of Product will present it. We think mobile experience is a problem but we're not sure."

  - type: assistant
    content: "Here's your prompt:\n\n---\n\n**Analyse this NPS survey data (800 Q4 responses, CSV attached).**\n\n**Objectives:**\n1. Identify the top 5 themes in the free-text comments, ranked by frequency\n2. For each theme, break down by NPS score band (Detractors 0-6, Passives 7-8, Promoters 9-10)\n3. Specifically investigate mobile experience mentions — frequency, sentiment, and correlation with low scores\n4. Flag any emerging themes that appear in 3%+ of responses\n\n**Output format:**\n- Executive summary (3-4 bullet points, suitable for VP presentation)\n- Detailed theme breakdown as a table\n- Mobile experience deep-dive section\n- 5 recommended actions for Q1 product planning, ranked by potential NPS impact\n\n**Constraints:**\n- Use direct quotes from responses as evidence (3-5 per theme)\n- Distinguish between correlation and causation\n- Flag if sample size for any segment is too small for reliable conclusions\n- Professional tone suitable for leadership presentation\n\n---\n\nThis prompt gives the AI your data context, a clear analysis structure, a specific hypothesis to test (mobile), and an output format matched to your audience. Want me to adjust anything?"

  - type: note
    text: "Look at what happened. The user started with 'I want to do something useful with our customer feedback.' After six questions and one set of answers, the AI produced a prompt with clear objectives, structured output format, a specific hypothesis to test, and constraints. That prompt would have taken 15 minutes to write from scratch."

  - type: note
    text: "This is meta-prompting: using AI to write the prompt that AI will then execute. The questions surface what you know but haven't articulated yet."
```

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
