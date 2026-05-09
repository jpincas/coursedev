---
title: "The Basic Framework"
duration: "15m"
tags: [framework, structure, components]
---

# A Framework for Structuring Requests

Here's a basic framework for structuring AI requests. You don't need all five elements every time, but thinking through them helps.

## The Five Elements

![The Five Framework Components](/content/module-prompting/images/framework-components.svg)

Modern models follow instructions very literally — vague prompts get vague results. **Being specific and direct remains the single highest-leverage technique** even with advanced models like Claude Opus 4.6 and GPT-5.x.

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

## The Framework in Action

Words on a page can only explain so much. Watch the difference between a bad prompt and a good one — using the same AI, the same task, the same data.

```agent
id: framework-demo
title: "Bad Prompt vs Good Prompt"
model_label: "Claude"

system: |
  You are a helpful business assistant. Respond based on the
  information and instructions provided.

scratchpad:
  "q4-data.csv": |
    Month,Revenue,New_Customers,Churn_Rate,NPS
    Oct,1.82M,145,3.2%,42
    Nov,2.01M,168,2.8%,45
    Dec,2.34M,203,2.1%,51
    Q4 Total,6.17M,516,2.7% avg,46 avg
    vs Q3,+18%,+24%,-0.8pp,+6

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
    text: "First, the bad prompt. Vague, no context, no format specified. Watch what comes out."

  - type: user
    content: "Write me a report about Q4."

  - type: assistant
    content: "I'd be happy to help write a Q4 report! However, I need some more information:\n\n- Which company or organisation is this for?\n- What year's Q4?\n- What type of report (financial, operational, marketing)?\n- Who is the audience?\n- What data should I include?\n\nCould you provide these details so I can create a useful report?"

  - type: note
    text: "The AI has to ask for clarification because the prompt gave it nothing to work with. Every back-and-forth like this wastes time. Now let's clear this and try with the full framework."

  - type: clear
    reset_scratchpad: false
    note: "Starting fresh with a better prompt."

  - type: user
    content: "You are a senior analyst preparing a board summary.\n\nRead the Q4 data file, then write a one-page executive summary for the leadership team. Include: revenue trends, customer growth, churn improvement, and NPS trajectory. Format as 3-4 bullet points with supporting numbers. Professional tone."

  - type: tool_call
    tool: scratchpad_read
    args:
      filename: "q4-data.csv"

  - type: tool_result
    tool: scratchpad_read
    content: |
      Month,Revenue,New_Customers,Churn_Rate,NPS
      Oct,1.82M,145,3.2%,42
      Nov,2.01M,168,2.8%,45
      Dec,2.34M,203,2.1%,51
      Q4 Total,6.17M,516,2.7% avg,46 avg
      vs Q3,+18%,+24%,-0.8pp,+6

  - type: assistant
    content: "**Q4 2025 Executive Summary**\n\n**Revenue acceleration exceeded targets.** Q4 revenue reached $6.17M, an 18% increase over Q3. Monthly revenue grew consistently from $1.82M in October to $2.34M in December, demonstrating sustained momentum rather than a one-off spike.\n\n**Customer acquisition is scaling.** 516 new customers joined in Q4, a 24% increase over Q3. December alone (203 new customers) was our strongest acquisition month on record, suggesting that recent marketing investments are yielding returns.\n\n**Retention is improving materially.** Average churn fell to 2.7%, down 0.8 percentage points from Q3. The December figure of 2.1% is approaching our annual target of sub-2%. This trend directly strengthens our recurring revenue base.\n\n**Customer satisfaction is trending upward.** NPS improved from 42 in October to 51 in December (average 46 for the quarter, +6 over Q3). Crossing the 50 threshold in December is significant — scores above 50 are considered 'excellent' in B2B SaaS."

  - type: note
    text: "Same AI. Same data. Completely different outcome. The first prompt got a request for clarification. The second got a finished executive summary. The difference is the framework: Role (senior analyst), Task (executive summary), Context (Q4 data file), Format (bullet points with numbers), Constraints (one page, professional tone)."
```

```quiz
id: framework-elements
type: multiple-choice
question: "Your prompt says 'Analyse this data' and you attach a CSV. AI gives a superficial overview. What is the most impactful element to add?"
options:
  - "Role: 'You are a senior data analyst' so it brings deeper analytical expertise"
  - "Format: 'Present as a table with trend arrows' so the output is more structured"
  - "Constraints: 'Maximum 500 words' so it focuses on what matters"
answer: 0
explanation: "Setting a role changes the depth and approach of analysis. A 'senior data analyst' brings expertise-level pattern recognition and identifies insights that a generic assistant misses. Format and constraints improve presentation but do not change analytical depth."
```
