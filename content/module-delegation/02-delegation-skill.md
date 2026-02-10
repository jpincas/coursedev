---
title: "Delegation as a Skill"
duration: "10m"
tags: [delegation, outcomes, specification]
---

# Delegation as a Skill

Delegation is a skill. And most people are bad at it.

## The Common Mistake

Most people describe processes, not outcomes.

![Micromanaging vs Delegating](/content/module-delegation/images/micromanaging-vs-delegating.svg)

**Process description (micromanaging):**
```
First look at the sales data, then identify the
top performers, then calculate growth rates,
then create a chart, then write a summary...
```

This is you doing the thinking and making AI do the typing. You're still doing the intellectual work.

**Outcome specification (delegation):**
```
Create a one-page sales performance summary with:
- Top 10 performers
- Growth trends
- Chart showing Q4 vs Q3
```

This is you defining what done looks like and letting AI figure out the process.

```callout
type: tip
title: "The Division of Labour"
content: "Your job: Define what done looks like. AI's job: Figure out how to get there. This is proper delegation."
```

## When AI Fails, You Failed

Here's the provocative truth: **when AI produces rubbish, that's a delegation failure, not an AI failure.**

The AI did what you asked. If the output is wrong, ask yourself:
- Did I give clear instructions?
- Did I provide sufficient context?
- Did I verify the output?
- Did I specify success criteria?

Most people blame the AI. "It didn't understand." "It got it wrong." "AI isn't ready yet."

That's a convenient deflection. It lets you avoid the uncomfortable reality: **you delegated badly.**

```callout
type: danger
title: "Accept Responsibility for AI Output"
content: "If you accept this framing — that AI failures are delegation failures — you'll improve much faster than people who blame the technology. Own the outcome."
```

When AI produces low-quality output, that's feedback on your delegation skill. Treat it as such.

## Why Process Description Fails

When you describe every step:
- You do the hard work (planning)
- AI just follows orders (easy part)
- You limit AI to your approach
- You miss better methods AI might use

When you describe outcomes:
- You focus on what matters (the result)
- AI brings its capabilities to the process
- You get potentially better approaches
- Work completes faster

```agent
id: delegation-vs-micromanagement
title: "Micromanaging vs Delegating"
model_label: "Claude"

system: |
  You are a business analyst assistant. You help create reports from sales
  data and business metrics. Follow the user's instructions precisely.

scratchpad:
  "sales-data.csv": |
    Quarter,Revenue,Growth
    Q1 2025,1200000,8.5%
    Q2 2025,1450000,20.8%
    Q3 2025,1380000,-4.8%
    Q4 2025,1920000,39.1%
  "report-template.md": |
    # Executive Summary Template

    ## Overview
    Brief statement of overall performance

    ## Key Trends
    What patterns emerged across the period

    ## Strategic Insights
    What this means for the business

    ## Recommendations
    What actions should be taken

tools:
  - scratchpad_read
  - scratchpad_write

visibility:
  system_prompt: toggleable
  tool_calls: visible
  full_context: hidden

sidebar:
  width: "45%"
  start_open: true

script:
  - type: note
    text: "Watch the difference between telling AI HOW to work vs WHAT you need. Same task, two approaches."

  - type: user
    content: "Step 1: Read the sales data. Step 2: Calculate the total for each quarter. Step 3: Find the quarter with highest growth. Step 4: Write one paragraph about each quarter. Step 5: Format it as markdown."

  - type: assistant
    content: "I'll follow those steps. Starting with Step 1 — reading the sales data."

  - type: tool_call
    tool: scratchpad_read
    args:
      filename: "sales-data.csv"

  - type: tool_result
    tool: scratchpad_read
    content: |
      Quarter,Revenue,Growth
      Q1 2025,1200000,8.5%
      Q2 2025,1450000,20.8%
      Q3 2025,1380000,-4.8%
      Q4 2025,1920000,39.1%

  - type: assistant
    content: "Now Step 2 — the totals are already shown. Step 3 — Q4 had the highest growth at 39.1%. Step 4 and 5 — writing the report now."

  - type: tool_call
    tool: scratchpad_write
    args:
      filename: "sales-report.md"
      content: |
        # Quarterly Sales Report

        ## Q1 2025
        Revenue was $1,200,000 with 8.5% growth.

        ## Q2 2025
        Revenue was $1,450,000 with 20.8% growth.

        ## Q3 2025
        Revenue was $1,380,000 with -4.8% growth (decline).

        ## Q4 2025
        Revenue was $1,920,000 with 39.1% growth. This was the highest growth quarter.

  - type: tool_result
    tool: scratchpad_write
    content: "Written to sales-report.md (11 lines)"

  - type: assistant
    content: "Done. I've completed all five steps and created sales-report.md with one paragraph for each quarter."

  - type: note
    text: "The AI followed orders perfectly. But look at the output — it just restates the numbers. No insight, no analysis, no strategic value. You did the thinking (planning the steps), AI did the typing."

  - type: clear
    reset_scratchpad: true
    note: "Now watch delegation in action. Same data, different approach."

  - type: user
    content: "Analyse our quarterly sales data and write an executive summary highlighting the key trends, risks, and recommended actions. Format it like the template. The audience is the board — they want strategic insights, not raw numbers."

  - type: assistant
    content: "I'll analyse the sales data and create a strategic executive summary. Let me review both the data and the template format."

  - type: tool_call
    tool: scratchpad_read
    args:
      filename: "sales-data.csv"

  - type: tool_result
    tool: scratchpad_read
    content: |
      Quarter,Revenue,Growth
      Q1 2025,1200000,8.5%
      Q2 2025,1450000,20.8%
      Q3 2025,1380000,-4.8%
      Q4 2025,1920000,39.1%

  - type: tool_call
    tool: scratchpad_read
    args:
      filename: "report-template.md"

  - type: tool_result
    tool: scratchpad_read
    content: |
      # Executive Summary Template

      ## Overview
      Brief statement of overall performance

      ## Key Trends
      What patterns emerged across the period

      ## Strategic Insights
      What this means for the business

      ## Recommendations
      What actions should be taken

  - type: assistant
    content: "I can see a volatile pattern — strong start, mid-year decline, exceptional Q4 recovery. Let me create a strategic analysis for the board."

  - type: tool_call
    tool: scratchpad_write
    args:
      filename: "executive-summary.md"
      content: |
        # 2025 Sales Performance — Executive Summary

        ## Overview
        Annual revenue reached $5.95M, representing 60% growth from Q1 to Q4. However, performance was highly volatile with a mid-year contraction followed by exceptional Q4 recovery.

        ## Key Trends

        **Volatile growth pattern:** After strong Q1-Q2 momentum (8.5% → 20.8%), Q3 contracted by 4.8% — our only decline of the year. Q4's 39.1% surge suggests successful recovery, but the volatility indicates underlying instability.

        **Accelerating momentum:** Q4's revenue ($1.92M) exceeded Q1 by 60% and represents 32% of annual revenue in a single quarter. This concentration raises both opportunity and risk questions.

        **Mid-year correction:** Q3's decline broke a growth trend. Understanding the root cause (seasonal, competitive, operational) is critical to prevent recurrence.

        ## Strategic Insights

        The data suggests we've found a winning formula by Q4, but we haven't sustained it consistently. Q3's decline indicates we're reactive, not proactive. The business is capable of exceptional performance (Q4) but hasn't institutionalised it.

        **Risk:** Q4 may represent unsustainable spike (one-time deals, discounting) rather than structural improvement. We need to verify whether Q4's success is repeatable.

        **Opportunity:** If Q4 represents genuine breakthrough, scaling those practices could drive sustained 30%+ growth.

        ## Recommendations

        1. **Root-cause analysis:** Conduct immediate post-mortem on Q3 decline and Q4 surge. Identify structural drivers vs one-time factors.

        2. **Sustain Q4 momentum:** Whatever drove Q4 success, codify it. If it's new sales practices, roll them out company-wide. If it's product-market fit breakthrough, double down.

        3. **De-risk concentration:** 32% of annual revenue in one quarter creates forecasting risk. Build pipeline visibility for Q1 2026 to ensure we're not facing another Q3-style gap.

        4. **Set realistic 2026 targets:** If Q4 is repeatable, 30%+ growth is achievable. If it's an outlier, plan for 15-20% as the sustainable baseline.

  - type: tool_result
    tool: scratchpad_write
    content: "Written to executive-summary.md (28 lines)"

  - type: assistant
    content: "I've created a board-level executive summary analysing the volatility pattern, identifying the Q3 decline and Q4 surge as the critical strategic questions, and providing four specific recommendations with rationale. The focus is on what the numbers mean for decision-making, not just what the numbers are."

  - type: note
    text: "Same data. Completely different output. The second version identifies trends, asks strategic questions, and recommends actions. Why? Because you delegated the outcome and let the AI apply its analytical capabilities."

  - type: note
    text: "In the first approach, you prescribed the process (five steps). The AI followed orders but added no value. In the second approach, you specified the outcome (executive summary for the board) and let the AI figure out HOW to analyse the data. That's delegation."
```

## Signs You're Micromanaging

- Your instructions have numbered steps
- You're telling AI "first... then... then..."
- You're prescribing specific tools or methods
- The instruction is longer than the output would be

## Signs You're Delegating Well

- You describe the finished state
- You specify requirements and constraints
- You let AI propose an approach
- You review the plan before execution

## The Only Way to Learn This

Reading about delegation is necessary. It is nowhere near sufficient.

**The skill of AI delegation is like driving.** You can read the manual cover to cover. You can watch videos. You can memorize the rules. But you only get good by doing it.

The muscle memory of structuring prompts, providing context, recognising when to iterate, knowing when output is good enough — that only comes through practice.

Every hour of practice is worth ten hours of reading.

You will fail. Your first attempts will produce garbage. You'll give vague instructions and wonder why the output is vague. You'll skip context and wonder why AI misunderstood.

**That's the learning process.** Each failure teaches you something about delegation. Each iteration improves your instinct for what works.

```callout
type: tip
title: "Start Practicing Now"
content: "Don't wait until you've finished this course. Pick a real task today and delegate it to AI. You'll learn more from one failed attempt than from three more modules of reading."
```

```quiz
id: delegation-skill
type: multiple-choice
question: "What's the problem with giving AI step-by-step instructions?"
options:
  - "AI can't follow complex instructions reliably"
  - "You're doing the intellectual work; AI just executes"
  - "Step-by-step instructions cost more tokens"
answer: 1
explanation: "When you provide step-by-step instructions, you're doing the planning and thinking — the hard part. AI just follows your orders. Proper delegation means defining the outcome and letting AI figure out the approach."
```
