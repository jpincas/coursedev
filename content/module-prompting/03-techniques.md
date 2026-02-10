---
title: "The Techniques That Work"
duration: "20m"
tags: [techniques, examples, multishot, xml, thinking, chaining]
---

# The Techniques That Actually Work

You've learned the framework. Now for the specific techniques that consistently deliver better results, ordered by effectiveness.

## 1. Be Specific and Direct

This remains the single highest-leverage technique. Modern models follow instructions very literally — vague prompts get vague results.

**The Colleague Test:** Show your prompt to a colleague. If they'd be confused about what you want, the AI will be too.

**Bad:**
```
Write about our product.
```

**Good:**
```
Write a 300-word feature description of our expense tracking product
for small business owners. Emphasise the automated receipt scanning
capability and mobile-first design. Professional but friendly tone.
```

Always specify: format, length, tone, scope, and audience.

```callout
type: tip
title: "Modern Models Are Literal"
content: "Claude Opus 4.6 and GPT-5.x take your instructions at face value. Vague input produces vague output. Specific input produces specific output."
```

## 2. Use Examples (Multishot Prompting)

Providing 2-5 examples of desired input-output pairs dramatically improves consistency and quality.

This is especially powerful when format or style matters.

**Without examples:**
```
Categorise these support tickets.
[List of tickets]
```

Results will be inconsistent. The AI has to guess what categories you want.

**With examples:**
```
Categorise these support tickets. Use these categories:

<examples>
"The app crashes when I upload photos" → Bug
"How do I export my data?" → Feature Request
"I can't log in" → Access Issue
</examples>

Now categorise these:
[List of tickets]
```

Results will match your desired pattern.

```callout
type: info
title: "The Power of Showing, Not Just Telling"
content: "Examples teach the model the pattern you want. Two to five examples are enough. More than that shows diminishing returns."
```

## 3. Enable Thinking for Hard Tasks

**Extended Thinking** is an internal reasoning phase before the model responds. Modern Claude models have this built in. OpenAI calls it "reasoning effort."

Use it for:
- Mathematical or logical problems
- Multi-step analysis
- Debugging complex issues
- Complex decisions requiring trade-off evaluation

**Counterintuitive insight:** Research shows that models often perform better with instructions to think deeply about a task rather than step-by-step prescriptive guidance.

**Instead of:**
```
First analyse X, then consider Y, then compare to Z, then conclude.
```

**Try:**
```
Think deeply about this problem before responding. Consider all angles.
```

The model's reasoning capabilities are sophisticated. Let it reason its own way.

## 4. Structure with XML Tags

Models trained on code are excellent at parsing structured formats like XML.

Tags separate prompt components, reduce misinterpretation, and make output easier to parse.

**Basic structure:**
```
<instructions>
Analyse the customer feedback and identify trends.
</instructions>

<context>
We launched a new mobile app two months ago.
This feedback comes from early adopters.
</context>

<examples>
"App is slow" → Performance Issue
"Love the dark mode" → Positive UI Feedback
</examples>

<output_format>
- List top 3 trends
- Include frequency counts
- Provide representative quotes
</output_format>

<constraints>
- Focus on actionable feedback
- Exclude feature requests we've already planned
</constraints>
```

This structure is much clearer than a wall of text with the same information.

**Note:** All major models benefit from structured formats — XML, JSON, or YAML. Use whichever feels most natural for your task.

```callout
type: tip
title: "Structure Reduces Ambiguity"
content: "Tags make it explicit where context ends and instructions begin. The model won't confuse examples with actual data, or constraints with tasks."
```

## 5. Chain Complex Prompts

For multi-step tasks, break the work into sequential subtask prompts where each output feeds the next.

This mirrors how you'd delegate work to a human: research first, then organise, then draft, then review.

**Single monolithic prompt:**
```
Research competitor pricing, analyse our positioning,
draft a pricing strategy, and write an exec summary.
```

This tries to do everything at once. Results are often shallow.

**Prompt chain:**
```
Step 1: Research competitor pricing for [products]. Create a comparison table.
[Review output, provide to next step]

Step 2: Given this competitor data, analyse our positioning. Where do we compete on price? Where on features?
[Review output, provide to next step]

Step 3: Draft a pricing strategy based on this analysis. Consider our cost structure and target margin.
[Review output, provide to next step]

Step 4: Summarise this strategy as a 1-page exec summary for the leadership team.
```

Each step has a clear input and output. You verify at each stage. The final result is much higher quality.

```callout
type: warning
title: "When to Chain vs When to Combine"
content: "Chain when each step requires verification or feeds the next. Combine when steps are independent. Don't chain unnecessarily — it takes longer."
```

## Before and After: Putting It All Together

**Before (vague, unstructured):**
```
Analyse our Q4 sales data and tell me what's important.
```

**After (specific, structured, with examples):**
```
<instructions>
Analyse Q4 2025 sales data. Identify the top 3 trends affecting revenue.
</instructions>

<context>
We sell B2B SaaS with annual and monthly plans.
Q4 is typically our strongest quarter due to year-end budget spending.
This year we launched a new enterprise tier in October.
</context>

<output_format>
For each trend:
- Trend name
- Supporting data (numbers, percentages)
- Business implication
- Recommended action
</output_format>

<examples>
"Enterprise tier adoption" → 15% of revenue, higher than forecast → Strong market fit → Invest in enterprise sales hiring
</examples>

<constraints>
- Focus on actionable insights, not just description
- Max 1 page
- Professional tone for exec team
</constraints>

[Attach: Q4 sales data CSV]
```

The second prompt gives the AI everything it needs to produce exactly what you want.

## Multishot in Practice

Descriptions of multishot prompting can feel abstract. Watch what it actually does — the same categorisation task, first without examples, then with.

```agent
id: multishot-demo
title: "The Power of Examples"
model_label: "Claude"

system: |
  You are a customer support analyst. Categorise tickets accurately
  based on the instructions and examples provided.

scratchpad:
  "tickets.txt": |
    Ticket #201: "Every time I try to upload a file over 10MB, the page goes white and I have to refresh."
    Ticket #202: "Is there a way to set up recurring tasks? I can't find it anywhere."
    Ticket #203: "My colleague shared a project with me but I can't see it in my dashboard."
    Ticket #204: "The mobile app logged me out and now my 2FA isn't working."
    Ticket #205: "Would love if you could add a dark mode option."
    Ticket #206: "Reports are showing last month's data even after I update the source spreadsheet."

tools:
  - scratchpad_read

visibility:
  system_prompt: hidden
  tool_calls: visible
  full_context: hidden

sidebar:
  width: "45%"
  start_open: true

script:
  - type: note
    text: "First attempt: no examples. Just a vague instruction to categorise. Watch how inconsistent the results are."

  - type: user
    content: "Read the tickets file and categorise each one."

  - type: tool_call
    tool: scratchpad_read
    args:
      filename: "tickets.txt"

  - type: tool_result
    tool: scratchpad_read
    content: |
      Ticket #201: "Every time I try to upload a file over 10MB, the page goes white and I have to refresh."
      Ticket #202: "Is there a way to set up recurring tasks? I can't find it anywhere."
      Ticket #203: "My colleague shared a project with me but I can't see it in my dashboard."
      Ticket #204: "The mobile app logged me out and now my 2FA isn't working."
      Ticket #205: "Would love if you could add a dark mode option."
      Ticket #206: "Reports are showing last month's data even after I update the source spreadsheet."

  - type: assistant
    content: "Here are the categorised tickets:\n\n- **#201** — Technical Issue / Upload Problem\n- **#202** — Question / Feature Inquiry\n- **#203** — Permissions / Sharing Issue\n- **#204** — Authentication / Login Problem\n- **#205** — Feature Suggestion\n- **#206** — Data/Reporting Bug"

  - type: note
    text: "Look at those categories. They're all different formats. Some have slashes, some don't. No consistent taxonomy. If you ran this on 500 tickets, you'd get 50 different category names. Now let's try with examples."

  - type: clear
    reset_scratchpad: false
    note: "Starting over with examples this time."

  - type: user
    content: "Read the tickets file and categorise each one.\n\nUse these exact categories. Here are examples:\n\n'App crashes on login' → BUG (P1)\n'Can I change my email address?' → QUESTION (P3)\n'Please add calendar integration' → FEATURE REQUEST (P3)\n'Can't access shared files' → ACCESS ISSUE (P2)\n\nFor each ticket give: Ticket # → CATEGORY (Priority)"

  - type: tool_call
    tool: scratchpad_read
    args:
      filename: "tickets.txt"

  - type: tool_result
    tool: scratchpad_read
    content: |
      Ticket #201: "Every time I try to upload a file over 10MB, the page goes white and I have to refresh."
      Ticket #202: "Is there a way to set up recurring tasks? I can't find it anywhere."
      Ticket #203: "My colleague shared a project with me but I can't see it in my dashboard."
      Ticket #204: "The mobile app logged me out and now my 2FA isn't working."
      Ticket #205: "Would love if you could add a dark mode option."
      Ticket #206: "Reports are showing last month's data even after I update the source spreadsheet."

  - type: assistant
    content: "#201 → BUG (P1)\n#202 → QUESTION (P3)\n#203 → ACCESS ISSUE (P2)\n#204 → ACCESS ISSUE (P1)\n#205 → FEATURE REQUEST (P3)\n#206 → BUG (P2)"

  - type: note
    text: "Consistent categories. Consistent format. Consistent priority assignment. Four examples were enough to teach the model your exact taxonomy and output format. Run this on 500 tickets and every one would follow the same pattern."

  - type: note
    text: "This is the power of showing over telling. The examples did what paragraphs of instructions couldn't: they defined the pattern unambiguously."
```

```quiz
id: techniques-multishot
type: multiple-choice
question: "You need AI to categorise 500 support tickets into exactly 6 categories. You describe the categories in detail. Results are inconsistent. What would most improve consistency?"
options:
  - "Add 3-4 example tickets with their correct categories so the model can learn the pattern"
  - "Use XML tags to separate the category definitions from the ticket data"
  - "Ask the model to think step-by-step before categorising each ticket"
answer: 0
explanation: "Multishot prompting (providing examples) is the most effective technique for categorisation and consistency tasks. Examples teach the exact pattern through demonstration -- the model matches your examples rather than interpreting your descriptions. Descriptions are ambiguous; examples are unambiguous."
```

## The Hierarchy of Techniques

![Techniques Hierarchy](/content/module-prompting/images/techniques-hierarchy.svg)

From most broadly effective to most specialised:

1. **Be specific and direct** — Always applicable. Highest leverage.
2. **Use examples** — When format, style, or categorisation matters.
3. **Enable thinking** — For genuinely hard reasoning tasks.
4. **Structure with tags** — When prompts are complex or multi-part.
5. **Chain prompts** — For multi-stage workflows requiring verification.

Master the first two and you'll see immediate results. Add the others as tasks get more complex.
