---
title: "The Iteration Pattern"
duration: "15m"
tags: [iteration, refinement, feedback, conversation]
---

# Conversation as Refinement

The most powerful AI interactions are rarely one-shot. They're conversations where you iterate toward exactly what you need.

This is the "conversation as refinement" pattern. It's how experts get 10x better results than beginners.

## The Core Pattern and Common Mistakes

![Vague vs Specific Feedback](/content/module-prompting/images/feedback-comparison.svg)

The iteration pattern:
1. **Start with a clear initial request** with constraints
2. **Review the output critically** identifying specific issues
3. **Give targeted feedback** referencing specific parts
4. **Iterate with precision**

Repeat until you have what you need.

```callout
type: info
title: "Why This Works"
content: "AI models are excellent at refinement. They can take feedback and adjust. The first output is the rough draft — your job is directing it toward the final version."
```

### The Most Common Beginner Mistake

**Bad feedback:**
```
Make this better.
```

**What does "better" mean?**
- More detailed? Less detailed?
- Different tone?
- Different structure?
- Different content focus?

The AI has to guess. Results are unpredictable.

**Good feedback:**
```
The second paragraph is too technical for this audience.
Simplify the explanation and remove the jargon.
The conclusion is too abrupt — add a specific call to action.
```

Now the AI knows exactly what to adjust.

### Specific Beats Generic

**Generic:**
```
This report needs improvement.
```

**Specific:**
```
This report:
- Is missing competitor analysis in section 2
- Uses too much jargon for a non-technical audience
- Needs concrete examples in the recommendations section
- Should be 30% shorter overall
```

Specific feedback gets specific improvements.

## How to Give Effective Feedback

Reference specific parts of the output.

**Vague:**
```
The tone is wrong.
```

**Specific:**
```
The opening paragraph has the right professional tone.
The middle section ("Our analysis shows...") becomes too casual.
Match the whole thing to the opening tone.
```

**Even better:**
```
"However, we totally need to think about..." — remove "totally" and "need to think about."
Make it: "However, we must consider..."
```

Quote the text you're referencing. Be surgical.

```callout
type: tip
title: "The Surgeon's Scalpel, Not the Sledgehammer"
content: "Precision feedback gets precision results. Identify the exact sentence, paragraph, or section that needs change. Don't just say 'fix it.'"
```

### The Colleague Test

Show your feedback to a colleague. If they wouldn't know what to change, neither will the AI.

**Colleague would be confused:** "The style needs work."

**Colleague would know what to do:** "Replace bullet points with numbered steps. Change the headings to questions. Add a one-sentence summary at the start of each section."

Clear instructions to humans are clear instructions to AI.

## See It in Action

Watch the iteration pattern improve a real proposal through three rounds of specific feedback:

```agent-demo
path: /content/module-prompting/agent-iteration-demo.yaml
```

## When to Stop Iterating

Stop when:
- The output meets your success criteria (from "What does done look like?")
- Further changes would be personal preference, not improvements
- You'd be satisfied sending this to your actual audience

Don't iterate endlessly. Diminishing returns kick in after 2-4 rounds for most tasks.

```callout
type: warning
title: "Over-Iteration Is Real"
content: "After 4-5 rounds, you're often making it different, not better. If you can't articulate a clear improvement, you're done."
```

### Iteration vs Starting Fresh

**Continue iterating when:**
- The foundation is good, just needs refinement
- Each iteration is getting closer to what you want
- The context from previous rounds is valuable

**Start fresh when:**
- The output is fundamentally wrong direction
- You've iterated 5+ times with no convergence
- The task scope has changed significantly

A fresh start with a better prompt often beats endless iteration on a poor foundation.

### The Iteration Mindset

Think of AI as a skilled but junior colleague who produces good first drafts quickly. You provide direction. They refine. You verify. Repeat until done.

**Your job:** Clear direction, specific feedback, quality verification

**AI's job:** Fast execution, pattern recognition, refinement

This division of labour is where AI productivity comes from.

```quiz
id: iteration-feedback
type: multiple-choice
question: "What makes feedback most effective in the iteration pattern?"
options:
  - "Using technical terminology to sound professional"
  - "Being polite and encouraging to the AI"
  - "Providing specific, targeted corrections referencing exact parts"
  - "Asking the AI what it thinks needs improvement"
answer: 2
explanation: "Specific, targeted feedback that references exact parts of the output is most effective. Generic feedback like 'make this better' forces the AI to guess. Precise corrections like 'paragraph 2 is too technical — simplify it' give clear direction."
```
