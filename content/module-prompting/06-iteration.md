---
title: "The Iteration Pattern"
duration: "15m"
tags: [iteration, refinement, feedback, conversation]
---

# Conversation as Refinement

The most powerful AI interactions are rarely one-shot. They're conversations where you iterate toward exactly what you need.

This is the "conversation as refinement" pattern. It's how experts get 10x better results than beginners.

```callout
type: warning
title: "The Single Biggest Source of Frustration"
content: "When AI gives you disappointing output, it's almost never because the AI is stupid. It's because you didn't iterate. Expecting perfect output on the first try is like expecting a colleague to read your mind. The frustration comes from unrealistic expectations, not AI limitations."
```

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

## Two Power Moves for Iteration

Beyond specific feedback, two techniques consistently improve iteration results.

### "What Do You Think?"

End your prompts or follow-ups with **"What do you think?"** or **"Does this look right to you?"**

This triggers the AI to critically evaluate its own work. Instead of just delivering output, it reviews what it produced and flags potential issues -- gaps it noticed, assumptions it made, areas where it is less confident.

```callout
type: tip
title: "Practical Example"
content: "Instead of: 'Write a project proposal for the new feature.' Try: 'Write a project proposal for the new feature. Then tell me -- what do you think? What's missing or weak?' The AI will often catch issues that would otherwise require your review to find."
```

This works because models are better at evaluating text than generating it from scratch. Asking for self-evaluation leverages this asymmetry.

### "Make It Promise"

When AI keeps ignoring a specific instruction -- it uses bullet points when you asked for paragraphs, or it keeps being formal when you asked for casual -- **make it explicitly acknowledge the constraint before proceeding.**

**The technique:**
```
Before you write the next version, confirm that you understand
these rules:
1. No bullet points -- paragraphs only
2. Casual tone, as if writing to a friend
3. Under 200 words

What are the rules you will follow?
```

Getting the AI to restate the constraint in its own words dramatically improves compliance. It is the equivalent of asking a colleague "Can you repeat back what I just asked for?" -- it forces attention to the specific instruction.

```callout
type: tip
title: "When to Use This"
content: "Reserve 'Make It Promise' for persistent issues where the AI keeps ignoring a specific instruction despite clear feedback. For most tasks, specific feedback is enough. This is the escalation technique."
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
question: "You reviewed AI's first draft and it needs improvement. Which feedback will produce the best second draft?"
options:
  - "'The tone needs work' -- giving the AI freedom to interpret what you mean"
  - "'Paragraph 3 uses jargon our clients won't understand. Replace technical terms with plain language and add a concrete example after the first sentence.'"
  - "'Make it better and more professional' -- keeping the feedback broad so the AI can improve everything"
answer: 1
explanation: "Surgical precision beats broad direction. Referencing exact paragraphs, quoting specific text, and providing concrete instructions gives the AI exactly what to change. Vague feedback forces guessing, which often makes some things better and others worse."
```
