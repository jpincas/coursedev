---
title: "Key Takeaways: Prompting"
duration: "5m"
tags: [summary, takeaways]
---

# Key Takeaways: Prompting

You now have a complete toolkit for communicating with AI effectively. These techniques are not theoretical -- they are the difference between generic output and exactly what you need. Here is the essential summary.

## 1. Outcomes Over Process

Describe what done looks like, not how to get there.

The AI figures out the process. Your job is defining the destination. "Produce a trend analysis with the top 5 insights as bullet points" beats "First analyse the data, then identify trends, then summarise them."

## 2. The Three Questions

Before any task, ask:
- **What does "done" look like?** — Define the end state
- **What context does it need?** — Provide the information required
- **What are the boundaries?** — Set the constraints

Answer these three and you've got a solid specification.

## 3. The Techniques That Deliver Results

From highest to lowest leverage:
- **Be specific and direct** — Always applicable. Highest impact.
- **Use examples (multishot prompting)** — 2-5 examples teach the pattern you want
- **Structure with XML tags** — Separate instructions, context, examples, constraints
- **Enable thinking** — For hard reasoning tasks, let the model think deeply
- **Chain prompts** — Break multi-step work into sequential subtasks

## 4. Iteration Is How Experts Work

The first output is the rough draft. Give specific, targeted feedback referencing exact parts.

"Make this better" → AI guesses
"Paragraph 2 is too technical — simplify it" → AI knows exactly what to change

## 5. Meta-Prompting and Interactive Speccing

Let AI help you write better prompts. Ask it to ask you clarifying questions. For complex deliverables, use interactive speccing: agree on the structure and requirements before executing. The spec becomes a contract that ensures alignment.

The AI's questions reveal what's needed. Your answers build the prompt. For significant projects, this conversation-before-execution approach transforms quality.

## 6. Complete Information Beats Clever Wording

It's not about magic words or secret phrases.

It's about giving AI the information it needs to succeed. Clear task, relevant context, explicit format.

```callout
type: tip
title: "The Mindset Shift"
content: "Think of AI as a capable but literal worker. It can't read your mind. Be explicit about what you want, provide what it needs, and it will deliver."
```

## Practical Checklist

Before your next AI task, run through this:

- [ ] Can I describe exactly what "done" looks like?
- [ ] Have I provided all relevant context and data?
- [ ] Have I specified format and length?
- [ ] Have I mentioned any constraints or exclusions?
- [ ] Would a capable human have enough to complete this task?
- [ ] Have I provided 2-3 examples if format/style matters?
- [ ] Am I prepared to iterate with specific feedback?

If you answer "no" to any of these, your request probably needs more detail.

```quiz
id: prompting-takeaway
type: multiple-choice
question: "You need a competitor analysis report. You have 15 minutes. What combination of techniques will produce the best result?"
options:
  - "Use the three questions to define done, provide competitor data files as context, and iterate once with specific feedback"
  - "Write a long prompt using XML tags, enable extended thinking, and chain it across three steps"
  - "Use meta-prompting to have AI write the perfect prompt, then execute that prompt"
answer: 0
explanation: "The highest-leverage combination is: define what done looks like (three questions), provide complete context (data files), and iterate with precision. The other approaches are valid techniques but over-engineer a 15-minute task. Match technique complexity to task complexity."
```

## Up Next

Now that you know how to communicate with AI effectively, let's talk about working with **files**. Moving from chat to real deliverables is where AI becomes truly useful for work.
