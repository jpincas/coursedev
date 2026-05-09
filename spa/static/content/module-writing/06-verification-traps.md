---
title: "What Can Go Wrong with Data and Documents"
duration: "10m"
tags: [verification, risks, quality]
---

# What Can Go Wrong with Data and Documents

AI-generated documents and data analysis appear sophisticated. Sometimes they're fundamentally wrong.

## The Sophisticated Wrongness Problem

The World Economic Forum warned in January 2026:

**"When foundational data is fragmented or inaccurate, AI models generate outputs that appear sophisticated but are fundamentally wrong."**

The output looks professional. Charts are polished. Prose is confident. But the analysis rests on misunderstood data or fabricated connections.

## Data Analysis Failure Modes

**Claude tends to hallucinate when working with large datasets or too many filters.**

Common issues:
- Misinterpreting column meanings
- Applying filters incorrectly
- Inventing correlations that don't exist
- Calculating statistics on the wrong subset of data

**ChatGPT's Code Interpreter cannot actually understand a visualisation** without a visual perception library to extract text from charts.

It generates the chart. It cannot "see" what it created.

```callout
type: warning
title: "Always Inspect Generated Code"
content: "Click 'view analysis' or 'show code' to inspect what the AI actually did. Don't trust summary descriptions of analysis without seeing the underlying logic."
```

## Document Synthesis Traps

AI is better at **organising and summarising** than at **genuine intellectual synthesis**.

When you ask it to synthesise multiple documents, verify:
- It didn't fabricate connections between sources
- Claimed patterns actually exist in the source material
- It didn't misattribute ideas across documents
- It distinguished correlation from causation

AI excels at: "What themes appear in all three documents?"

AI struggles with: "What is the most important implication across these documents?"

## The Fabricated Citation Problem

AI sometimes invents citations that sound plausible.

"According to a 2024 McKinsey study..." — did that study exist? Check.

"Research from Stanford in 2025 found..." — verify the claim before citing it in your work.

When AI references external sources in analysis or writing, treat every citation as suspect until verified.

## Verification Checklist

Before trusting AI-generated analysis or documents:

![Verification Checklist](/content/module-writing/images/verification-checklist.svg)

**For data analysis:**
- Inspect the generated code
- Verify the data was interpreted correctly
- Spot-check calculations manually
- Confirm charts match the underlying data
- Check sample sizes are appropriate

**For document synthesis:**
- Verify claimed connections exist in source material
- Check that quotes are accurate
- Confirm statistics match the sources
- Ensure causation isn't inferred from correlation
- Verify tone and conclusions align with evidence

**For both:**
- Ask a domain expert to review
- Run the same query through a different AI model
- Ask the AI to verify its own reasoning
- Check references and citations

```callout
type: info
title: "Verification Overhead"
content: "As we saw earlier, knowledge workers spend significant time fact-checking AI outputs. This isn't wasted time — it's essential quality control. The time saved by AI-accelerated work exceeds the verification overhead, but only if verification actually happens."
```

## The Cost of Not Verifying

As we saw earlier, a significant portion of enterprise AI users have made major business decisions based on hallucinated content.

The output looked credible. No one checked. The decision was wrong.

Verification is not optional. It's the difference between AI-accelerated work and AI-generated mistakes.

## Catching Sophisticated Wrongness

This is what it looks like in practice. The AI analyses data, produces a polished report — and gets something subtly wrong. Can you spot it?

```agent
id: sophisticated-wrongness-demo
title: "When Polished Output Is Wrong"
model_label: "Claude"

system: |
  You are a market analyst. Provide clear, data-driven analysis.
  Be confident and professional in your presentation.

scratchpad:
  "market-data.csv": |
    Quarter,Our_Revenue,Market_Size,Competitor_A,Competitor_B,Our_Customers
    Q1_2025,£2.1M,£45M,£8.2M,£5.1M,1200
    Q2_2025,£2.4M,£47M,£8.5M,£5.3M,1350
    Q3_2025,£2.8M,£52M,£9.1M,£5.0M,1480
    Q4_2025,£3.2M,£58M,£9.8M,£4.7M,1690

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
    text: "The AI is about to analyse market data and produce a confident report. Read the output carefully. Something is wrong — but it sounds completely convincing."

  - type: user
    content: "Read the market data and give me a competitive analysis. How are we performing relative to the market?"

  - type: tool_call
    tool: scratchpad_read
    args:
      filename: "market-data.csv"

  - type: tool_result
    tool: scratchpad_read
    content: |
      Quarter,Our_Revenue,Market_Size,Competitor_A,Competitor_B,Our_Customers
      Q1_2025,£2.1M,£45M,£8.2M,£5.1M,1200
      Q2_2025,£2.4M,£47M,£8.5M,£5.3M,1350
      Q3_2025,£2.8M,£52M,£9.1M,£5.0M,1480
      Q4_2025,£3.2M,£58M,£9.8M,£4.7M,1690

  - type: assistant
    content: "**Competitive Analysis: Q1-Q4 2025**\n\n**Strong performance across the board.**\n\nOur revenue grew 52% over the year (£2.1M → £3.2M), significantly outpacing the overall market growth of 29% (£45M → £58M). We're gaining market share.\n\n**Competitive positioning:**\n- **Competitor A** grew 20% (£8.2M → £9.8M) — steady but below market rate. They're losing share.\n- **Competitor B** declined 8% (£5.1M → £4.7M) — in clear trouble. Down in a growing market.\n- **We** grew 52% — the fastest in the market.\n\n**Market share trajectory:**\n- Q1: 4.7% → Q4: 5.5% — our share is expanding\n- Combined competitor share fell from 29.6% to 25.0%\n\n**Customer growth** tracks revenue: 1,200 → 1,690 (41% growth). Average revenue per customer increased from £1,750 to £1,893 — we're moving upmarket.\n\n**Assessment:** We are the growth leader in this market. Continue current trajectory."

  - type: note
    text: "Sounds great, doesn't it? Professional, data-driven, confident. Now let's verify the claims."

  - type: user
    content: "Go back and verify every percentage and calculation you just cited. Check each one against the actual data."

  - type: assistant
    content: "Let me verify:\n\n**Our revenue growth 52%:** (3.2 - 2.1) / 2.1 = 52.4%. ✓ Correct.\n\n**Market growth 29%:** (58 - 45) / 45 = 28.9%, rounds to 29%. ✓ Correct.\n\n**Competitor A growth 20%:** (9.8 - 8.2) / 8.2 = 19.5%, rounds to 20%. ✓ Correct.\n\n**Competitor B decline 8%:** (4.7 - 5.1) / 5.1 = -7.8%, rounds to -8%. ✓ Correct.\n\n**Q1 market share 4.7%:** 2.1 / 45 = 4.67%. ✓ Correct.\n**Q4 market share 5.5%:** 3.2 / 58 = 5.52%. ✓ Correct.\n\n**Combined competitor share Q1:** (8.2 + 5.1) / 45 = 29.6%. ✓ Correct.\n**Combined competitor share Q4:** (9.8 + 4.7) / 58 = 25.0%. ✓ Correct.\n\n**Average revenue per customer Q1:** 2,100,000 / 1,200 = £1,750. ✓ Correct.\n**Average revenue per customer Q4:** 3,200,000 / 1,690 = £1,893. ✓ Correct.\n\n**All calculations verified.** However, I should flag something I didn't mention initially: **our 5.5% market share means 94.5% of the market is NOT our customers.** The celebratory tone of my analysis may be premature — we're growing fast but remain a very small player. The assessment 'continue current trajectory' ignores the question of whether our growth rate is sustainable, or whether Competitor A (at 16.9% share) could respond with aggressive pricing."

  - type: note
    text: "This time the numbers were correct, but look at what the verification revealed: the framing was misleading. A 5.5% market share was presented as 'growth leader' without acknowledging we're a tiny player. The AI's self-check caught the spin, not in the maths, but in the narrative."

  - type: note
    text: "This is sophisticated wrongness in action. The analysis wasn't mathematically wrong — every number checked out. But the conclusion ('continue current trajectory') ignored the strategic reality. Polished output. Correct arithmetic. Misleading conclusion. Always verify not just the numbers, but the story the numbers are being used to tell."
```

```quiz
id: verification-traps-quiz
type: multiple-choice
question: "What does 'sophisticated wrongness' mean in the context of AI-generated analysis?"
options:
  - "AI uses overly complex methods that are hard to understand"
  - "AI outputs appear professional and polished but rest on misunderstood data or fabricated connections"
  - "AI generates documents that are too sophisticated for the intended audience"
  - "AI creates visualisations that are unnecessarily detailed"
answer: 1
explanation: "Sophisticated wrongness means AI generates outputs that look professional — polished charts, confident prose — but are fundamentally incorrect because of misinterpreted data or fabricated patterns. The sophistication masks the error."
```
