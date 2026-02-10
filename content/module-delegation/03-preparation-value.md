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

Real preparation isn't writing a better prompt — it's a process. Gather your data, commission research, review the findings, then commission the deliverable. Watch the full workflow.

```agent
id: deep-research-demo
title: "Commissioning Real Research"
model_label: "Claude"

system: |
  You are a business research analyst. Read all provided materials
  carefully before responding. Cross-reference internal data with
  your analysis. Be specific and evidence-based.

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
  "lost-deals.txt": |
    Enterprise Deal Post-Mortems (Last 12 months)

    1. Barclays (2,400 employees) — LOST
       Reason: "No API integration with our HR platform (Workday)"
       They liked: engagement approach, pricing competitive
       Went with: Virgin Pulse

    2. Deloitte UK (4,200 employees) — LOST
       Reason: "Need real-time dashboards for 12 office locations"
       They liked: mental health programme design
       Went with: Built in-house

    3. Tesco HQ (800 employees) — LOST
       Reason: "Can't demonstrate ROI at board level without data platform"
       They liked: personalisation, came back for SME subsidiary
       Went with: Wellhub (formerly Gympass)

    4. HSBC Innovation Team (350 employees) — LOST
       Reason: "No SSO, no compliance reporting, no data residency guarantees"
       They liked: pilot results (96% engagement in 4-week trial)
       Went with: Delayed decision

    5. Arup (600 employees) — WON (only enterprise win this year)
       Reason: "Champion in HR bypassed normal procurement"
       Risk: No platform means we may not survive annual review

tools:
  - scratchpad_read
  - scratchpad_write

visibility:
  system_prompt: hidden
  tool_calls: visible
  full_context: hidden

sidebar:
  width: "50%"
  start_open: true

script:
  - type: note
    text: "Phase 1 — Commission the research. This person has already done their preparation: they pulled five enterprise deal post-mortems from the CRM and compiled their company data. Now they commission the AI to do the research legwork."

  - type: user
    content: "Read our company profile in client-data.txt and our enterprise deal post-mortems in lost-deals.txt. I need a competitive landscape analysis: for each platform our prospects chose instead of us, what do they offer and why did they win? What requirements keep coming up? Write your findings to research-notes.md."

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
    tool: scratchpad_read
    args:
      filename: "lost-deals.txt"

  - type: tool_result
    tool: scratchpad_read
    content: |
      Enterprise Deal Post-Mortems (Last 12 months)

      1. Barclays (2,400 employees) — LOST
         Reason: "No API integration with our HR platform (Workday)"
         They liked: engagement approach, pricing competitive
         Went with: Virgin Pulse

      2. Deloitte UK (4,200 employees) — LOST
         Reason: "Need real-time dashboards for 12 office locations"
         They liked: mental health programme design
         Went with: Built in-house

      3. Tesco HQ (800 employees) — LOST
         Reason: "Can't demonstrate ROI at board level without data platform"
         They liked: personalisation, came back for SME subsidiary
         Went with: Wellhub (formerly Gympass)

      4. HSBC Innovation Team (350 employees) — LOST
         Reason: "No SSO, no compliance reporting, no data residency guarantees"
         They liked: pilot results (96% engagement in 4-week trial)
         Went with: Delayed decision

      5. Arup (600 employees) — WON (only enterprise win this year)
         Reason: "Champion in HR bypassed normal procurement"
         Risk: No platform means we may not survive annual review

  - type: tool_call
    tool: scratchpad_write
    args:
      filename: "research-notes.md"
      content: |
        # Competitive Landscape Research

        ## Competitors That Won Our Deals

        **Virgin Pulse** (won Barclays)
        - Full enterprise wellness platform, established market leader
        - 200+ HR system integrations including Workday, SAP, BambooHR
        - Real-time analytics dashboards, multi-site reporting
        - Enterprise SSO, ISO 27001 certified, GDPR compliant
        - Why they won: API-first architecture plugs into existing HR stack

        **Wellhub / formerly Gympass** (won Tesco)
        - Marketplace model — aggregates wellness providers under one platform
        - Corporate clients get single dashboard across all providers
        - Built-in ROI reporting and utilisation analytics
        - Note: HealthBridge could potentially join as a provider rather than compete
        - Why they won: Board-level reporting capabilities we couldn't match

        **Deloitte built in-house**
        - Had engineering budget to solve it themselves
        - Wanted multi-site dashboards across 12 locations
        - Why: Specific reporting needs for 4,200 employees across sites

        **HSBC delayed (no winner yet)**
        - Ran 4-week pilot with us — 96% engagement
        - Couldn't proceed without SSO, compliance reporting, data residency
        - Product worked; procurement process blocked it

        ## Recurring Enterprise Requirements

        | Requirement | Deals Affected | Category |
        |------------|---------------|----------|
        | HR platform integration (Workday, SAP) | Barclays, HSBC | Infrastructure |
        | Real-time dashboards & reporting | Deloitte, Tesco | Reporting |
        | SSO / single sign-on | HSBC | Security |
        | Compliance reporting | HSBC | Security |
        | Data residency guarantees | HSBC | Security |
        | ROI demonstration at board level | Tesco | Reporting |
        | Multi-site support | Deloitte | Infrastructure |

        ## Key Observation
        Every deal was lost on technology/infrastructure grounds. No prospect rejected our core service — engagement, programme design, and personalisation were praised in every post-mortem. Our only enterprise win (Arup) happened because someone bypassed procurement entirely.

  - type: tool_result
    tool: scratchpad_write
    content: "Written to research-notes.md (42 lines)"

  - type: assistant
    content: "Research written to research-notes.md. The key finding across all five deals: every enterprise loss was on technology and infrastructure grounds — integrations, dashboards, SSO, compliance. No prospect rejected the core wellness programmes. Your Arup win confirms it — they only converted because someone bypassed standard procurement."

  - type: note
    text: "Phase 2 — Review the research. Open research-notes.md in the sidebar to read through it. In a real workflow, this is where you'd spend time thinking: Does this match my experience? What's surprising? What pattern stands out? The human reads the research and spots something important."

  - type: note
    text: "Phase 3 — Commission the deliverable. Armed with the research AND their own insight from reviewing it, the human now commissions the board document. Watch the prompt — the human feeds back a specific strategic insight that shapes the entire output."

  - type: user
    content: "Good research. Here's the pattern I see: three of four losses are infrastructure gaps — integrations, SSO, compliance — not product quality. HSBC even ran a successful pilot and couldn't buy. That's our key insight for the board.\n\nNow write a board-ready analysis in competitive-analysis.md. The decision: should we spend £500K building our own platform, or partner with an existing one like CoreHealth or Wellhub? Lead with the recommendation. Use the lost deal evidence to make the case. Include a phased timeline. No consultant-speak — our board wants specifics."

  - type: tool_call
    tool: scratchpad_read
    args:
      filename: "research-notes.md"

  - type: tool_result
    tool: scratchpad_read
    content: |
      # Competitive Landscape Research

      ## Competitors That Won Our Deals

      **Virgin Pulse** (won Barclays)
      - Full enterprise wellness platform, established market leader
      - 200+ HR system integrations including Workday, SAP, BambooHR
      - Real-time analytics dashboards, multi-site reporting
      - Enterprise SSO, ISO 27001 certified, GDPR compliant
      - Why they won: API-first architecture plugs into existing HR stack

      **Wellhub / formerly Gympass** (won Tesco)
      - Marketplace model — aggregates wellness providers under one platform
      - Corporate clients get single dashboard across all providers
      - Built-in ROI reporting and utilisation analytics
      - Note: HealthBridge could potentially join as a provider rather than compete
      - Why they won: Board-level reporting capabilities we couldn't match

      **Deloitte built in-house**
      - Had engineering budget to solve it themselves
      - Wanted multi-site dashboards across 12 locations
      - Why: Specific reporting needs for 4,200 employees across sites

      **HSBC delayed (no winner yet)**
      - Ran 4-week pilot with us — 96% engagement
      - Couldn't proceed without SSO, compliance reporting, data residency
      - Product worked; procurement process blocked it

      ## Recurring Enterprise Requirements

      | Requirement | Deals Affected | Category |
      |------------|---------------|----------|
      | HR platform integration (Workday, SAP) | Barclays, HSBC | Infrastructure |
      | Real-time dashboards & reporting | Deloitte, Tesco | Reporting |
      | SSO / single sign-on | HSBC | Security |
      | Compliance reporting | HSBC | Security |
      | Data residency guarantees | HSBC | Security |
      | ROI demonstration at board level | Tesco | Reporting |
      | Multi-site support | Deloitte | Infrastructure |

      ## Key Observation
      Every deal was lost on technology/infrastructure grounds. No prospect rejected our core service — engagement, programme design, and personalisation were praised in every post-mortem. Our only enterprise win (Arup) happened because someone bypassed procurement entirely.

  - type: tool_call
    tool: scratchpad_write
    args:
      filename: "competitive-analysis.md"
      content: |
        # Board Briefing: Build vs. Partner

        **Recommendation: Partner first. Build later if needed.**

        ## The Evidence

        We lost four enterprise deals in 12 months. Not one rejected our programmes — they rejected our infrastructure:

        | Deal | What They Wanted | What We Lacked |
        |------|-----------------|---------------|
        | Barclays (2,400 emp) | Workday API integration | Any HR system integration |
        | Deloitte (4,200 emp) | Multi-site dashboards | Reporting platform |
        | Tesco (800 emp) | Board-level ROI data | Data export & analytics |
        | HSBC (350 emp) | SSO, compliance, data residency | Enterprise IT security |

        HSBC is the clearest case: they ran a pilot, got 96% engagement, and still couldn't buy. Our product works. Their procurement process needs infrastructure we don't have.

        ## Why Building Doesn't Work

        A £500K build gets us dashboards and basic reporting in 9-12 months. It does not get us:
        - Workday/SAP integrations (lost Barclays)
        - ISO 27001 certification (6+ months on its own)
        - SSO and enterprise security (stalled HSBC)

        Realistic cost for enterprise-grade: £800K-1.2M. We'd overshoot budget and still not solve two of the four blockers.

        ## Why Partnership Works

        White-label platforms like CoreHealth already have the enterprise wrapper: HR integrations, SSO, compliance certifications, multi-site dashboards. We plug our programmes into their infrastructure.

        - Timeline: 3-4 months to integrate (vs. 9-12 to build)
        - Cost: Revenue share (15-25%) instead of £500K capital
        - Solves: All four deal-breakers from our lost deals
        - We keep: Programme design, engagement methodology, client relationships

        Our 92% engagement rate makes us attractive to platforms — we have what they lack (programme quality that works). They have what we lack (enterprise infrastructure).

        ## Proposed Timeline

        1. **Months 1-2:** Evaluate CoreHealth (white-label) and Wellhub (marketplace) partnerships
        2. **Month 3:** Begin integration with chosen partner
        3. **Months 4-5:** Re-approach HSBC (they delayed — no competitor won) and Tesco subsidiary
        4. **Month 6+:** First enterprise revenue. Assess whether custom build still needed.

        ## Risk Factors

        - Partnership terms reduce margins (15-25% revenue share)
        - Integration timeline depends on partner's capacity
        - We become partially dependent on partner's roadmap
        - Mitigation: partnership revenue funds eventual independence if needed

  - type: tool_result
    tool: scratchpad_write
    content: "Written to competitive-analysis.md (48 lines)"

  - type: assistant
    content: "Board briefing written to competitive-analysis.md.\n\nThe case: £500K on a build gets us halfway in 12 months — dashboards but not the HR integrations or security infrastructure that lost us Barclays and HSBC. Partnership gets us everything in 3 months for revenue share. Our 92% engagement rate makes us the attractive partner — we bring programme quality they lack.\n\nHSBC is the strongest board argument: 96% pilot engagement, still couldn't buy. The gap is infrastructure, not product."

  - type: note
    text: "Three phases, three different types of human value. Phase 1: the human gathered proprietary data the AI couldn't access (CRM post-mortems). Phase 2: the human reviewed the research and spotted the strategic pattern (infrastructure vs product gaps). Phase 3: the human framed the deliverable around that insight for a specific audience (the board). The AI did the research and the writing — but the thinking that made it useful was human."
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
