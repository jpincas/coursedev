---
title: "Preparation Is Your Value"
duration: "10m"
tags: [preparation, value, planning]
---

# Preparation Matters More Now

Because execution is cheap.

## The Shift in Bottlenecks

**Before AI:**
When creating a report took 4 hours, you might skip proper planning and just start writing. The bottleneck was execution time.

**After AI:**
When creating a report takes 5 minutes, the bottleneck shifts. The quality of the output is determined by the quality of the input — your specification, your thinking, your planning.

## Where Value Lives Now

![Value Shift in AI-Assisted Work](/content/module-delegation/images/value-shift.svg)

| Activity | Before AI | After AI |
|----------|-----------|----------|
| Thinking | Important | Critical |
| Planning | Often skipped | Essential |
| Specification | Rough was fine | Precision matters |
| Execution | Main time sink | Near-instant |
| Review | Quick pass | Where you add value |

## Your New Value Proposition

Your value shifts from "I can do this" to "I know what should be done."

**Research matters more**
Understanding the problem space, knowing what's been tried, identifying the real requirements.

**Planning matters more**
Defining what success looks like, anticipating edge cases, setting constraints.

**Specification matters more**
Clear articulation of requirements, explicit success criteria, complete context.

**Execution?**
That's the cheap part now.

```callout
type: warning
title: "The Risk"
content: "If you skip preparation because 'AI will figure it out,' you'll get fast production of the wrong thing. Speed without direction is just efficient waste."
```

## Practical Implications

**Invest time upfront**
Five minutes of clear specification beats five iterations of vague requests.

**Think before delegating**
What exactly do you need? Who's the audience? What constraints apply?

**Define success criteria**
How will you know if the output is good? What would make it wrong?

**But don't over-plan**
Preparation matters, but iteration speed is also a superpower. McKinsey research found that **fifteen iterations in two days beats two iterations in five days**.

The sweet spot: enough preparation that the first output is in the ballpark, then rapid iteration to refine. Technology delivers 20% of value; redesigning work delivers 80%.

This preparation is what separates effective AI users from frustrated ones.

## Preparation Quality = Output Quality

Same task. Same AI. Same data. The only variable is how well the request is prepared. Watch the output quality track the input quality precisely.

```agent
id: preparation-demo
title: "The Preparation Difference"
model_label: "Claude"

system: |
  You are a business consultant. Produce output based on the
  information and instructions provided.

scratchpad:
  "client-data.txt": |
    HealthBridge Wellness — Client Summary
    Industry: Corporate wellness programmes
    Size: 45 employees, £3.2M revenue
    Founded: 2021
    Services: On-site fitness, mental health workshops, nutrition coaching
    Clients: 28 corporate accounts
    Growth: 40% YoY revenue growth
    Problem: Losing enterprise bids to larger competitors
    Win rate: 85% for SMEs, 12% for enterprise (100+ employees)
    Differentiator: Personalised programmes, high participant engagement (92%)
    Weakness: No technology platform, manual reporting

tools:
  - scratchpad_read
  - scratchpad_write

visibility:
  system_prompt: hidden
  tool_calls: visible
  full_context: hidden

sidebar:
  width: "45%"
  start_open: true

script:
  - type: note
    text: "Round 1: minimal preparation. Vague request with no specification."

  - type: user
    content: "Read the client data and write something about their strategy."

  - type: tool_call
    tool: scratchpad_read
    args:
      filename: "client-data.txt"

  - type: tool_result
    tool: scratchpad_read
    content: |
      HealthBridge Wellness — Client Summary
      Industry: Corporate wellness programmes
      Size: 45 employees, £3.2M revenue
      Founded: 2021
      Services: On-site fitness, mental health workshops, nutrition coaching
      Clients: 28 corporate accounts
      Growth: 40% YoY revenue growth
      Problem: Losing enterprise bids to larger competitors
      Win rate: 85% for SMEs, 12% for enterprise (100+ employees)
      Differentiator: Personalised programmes, high participant engagement (92%)
      Weakness: No technology platform, manual reporting

  - type: assistant
    content: "HealthBridge Wellness has experienced impressive growth since its founding in 2021, achieving 40% year-over-year revenue growth. The company excels in the SME market with an 85% win rate, demonstrating strong product-market fit. However, their enterprise win rate of just 12% represents a significant growth opportunity. To capitalise on this opportunity, HealthBridge should consider investing in technology infrastructure, as their current manual reporting process may be hindering enterprise sales. By developing a technology platform and leveraging their high engagement metrics, HealthBridge can position itself for success in the enterprise market."

  - type: note
    text: "Generic strategy-speak. Restates the data you already have. 'Capitalise on this opportunity.' 'Leveraging their metrics.' No actual insight. Now watch what happens with five minutes of preparation."

  - type: clear
    reset_scratchpad: false
    note: "Same data. Prepared request this time."

  - type: user
    content: "Read the client data. Then write a 1-page strategic recommendation for HealthBridge's CEO.\n\nContext: She's deciding whether to invest £500K in a technology platform or hire 3 more enterprise sales people. She needs to decide by end of Q1.\n\nFocus on: Which investment has the higher expected return? Use their existing data to make the case.\n\nFormat: Lead with the recommendation. Support with numbers. End with risk factors. No buzzwords — she hates consultant-speak.\n\nConstraints: The CEO has 5 minutes. Make every sentence count."

  - type: tool_call
    tool: scratchpad_read
    args:
      filename: "client-data.txt"

  - type: tool_result
    tool: scratchpad_read
    content: |
      HealthBridge Wellness — Client Summary
      Industry: Corporate wellness programmes
      Size: 45 employees, £3.2M revenue
      Founded: 2021
      Services: On-site fitness, mental health workshops, nutrition coaching
      Clients: 28 corporate accounts
      Growth: 40% YoY revenue growth
      Problem: Losing enterprise bids to larger competitors
      Win rate: 85% for SMEs, 12% for enterprise (100+ employees)
      Differentiator: Personalised programmes, high participant engagement (92%)
      Weakness: No technology platform, manual reporting

  - type: tool_call
    tool: scratchpad_write
    args:
      filename: "strategic-recommendation.md"
      content: |
        # Strategic Recommendation: Technology Before Sales Headcount

        **Invest in the platform. Hire later.**

        Your enterprise win rate is 12%. Hiring 3 more salespeople to lose 88% of enterprise bids is expensive failure at scale. The platform solves the root cause; more salespeople don't.

        ## The Numbers

        **Option A — 3 enterprise sales hires (£500K):**
        Assuming each targets 20 enterprise prospects/year at your current 12% win rate = ~7 new enterprise clients per year. If average enterprise contract is 3x your SME average (~£340K total), that's ~£2.4M potential revenue. But: your win rate stays at 12% because the gap isn't sales effort — it's credibility.

        **Option B — Technology platform (£500K):**
        Enterprise clients need dashboards, automated reporting, data integration. That's table stakes — not having it is why you lose. If the platform lifts enterprise win rate from 12% to even 25%, your existing pipeline produces more revenue than Option A — without adding headcount.

        ## The Real Argument

        Your 92% participant engagement is best-in-class. Enterprise buyers know it. They still don't buy because you can't prove ROI at their scale. A reporting platform converts your differentiator (engagement) into their language (data, dashboards, measurable outcomes).

        Build the platform. Then hire the sales team to sell it.

        ## Risk Factors
        - Platform build could exceed budget or timeline
        - Enterprise win rate improvement is estimated, not guaranteed
        - 6-9 month build time means delayed revenue impact
        - Consider a phased approach: MVP platform in 3 months, full build in 6

  - type: tool_result
    tool: scratchpad_write
    content: "Written to strategic-recommendation.md (25 lines)"

  - type: assistant
    content: "Created strategic-recommendation.md. The recommendation is platform-first: hiring salespeople to lose 88% of bids is expensive failure at scale. The platform addresses the root cause (credibility gap), not the symptom (win rate). Risk factors included."

  - type: note
    text: "Same AI. Same data. The first output was consultant filler you'd throw away. The second is a document you could hand to a CEO. The difference: five minutes of preparation defining the decision, the audience, the constraints, and the format."

  - type: note
    text: "Notice the preparation didn't take long — it's not a 20-page brief. It's a clear decision context (platform vs hires), a specific audience (CEO, 5 minutes, hates buzzwords), and focused constraints (lead with recommendation, use their numbers). That's where your value lives now."
```

```quiz
id: preparation-value
type: multiple-choice
question: "Why does preparation matter more in an AI-assisted workflow?"
options:
  - "AI needs detailed instructions to work at all"
  - "Execution is now cheap; input quality determines output quality"
  - "AI can't do planning on its own"
  - "Preparation reduces AI costs"
answer: 1
explanation: "When execution is nearly instant, the quality bottleneck shifts to input quality. Your thinking, planning, and specification determine whether AI produces something valuable or fast garbage."
```
