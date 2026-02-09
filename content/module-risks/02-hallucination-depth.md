---
title: "Hallucination: The Numbers"
duration: "10m"
tags: [hallucination, accuracy, verification]
---

# Hallucination: The Numbers

You learned earlier that hallucinations are inherent to how LLMs work. They guess the next token based on probability, not truth.

Now for the specifics that matter in practice.

## The Rates Vary Dramatically

**Google's Gemini-2.0-Flash:** 0.7% hallucination rate on standard benchmarks. Lowest recorded.

**OpenAI's o3 reasoning model:** 33% hallucination rate on PersonQA. Despite being optimised for accuracy, it fabricates a third of answers when asked about specific individuals.

**OpenAI's o4-mini:** 48% hallucination rate on the same benchmark.

The variation matters. Model selection directly impacts reliability.

![Hallucination rates vary dramatically by domain](/content/module-risks/images/hallucination-rates.svg)

## Domain-Specific Risks Are Higher

Stanford researchers tested general-purpose LLMs on legal queries. Result: **58-82% hallucination rate** depending on the model and question type.

Even specialised legal AI tools hallucinated **17-34% of the time.**

This isn't unique to law. Any domain with specialised knowledge shows elevated hallucination rates when models lack grounding.

```callout
type: danger
title: "The Real Cost"
content: "In 2024, 47% of enterprise AI users made at least one major business decision based on hallucinated content. Knowledge workers now spend an average of 4.3 hours per week fact-checking AI outputs."
```

## Why Hallucinations Persist

OpenAI's own research explains the core problem. Their training methods **reward guessing over acknowledging uncertainty.**

The analogy: imagine a multiple-choice test where leaving an answer blank guarantees zero marks. You'd guess every time. That's what current training incentives create.

Models are penalised for saying "I don't know" even when uncertainty is the honest answer.

## Practical Countermeasures

You can reduce hallucination impact without waiting for better models.

**Multi-model cross-validation.** Run critical queries through Claude, GPT, and Gemini. Where outputs agree, confidence increases. Where they diverge, verify independently.

**Self-verification.** Ask the AI to verify its own reasoning before providing a final answer. Studies show this **reduces errors by roughly 17%.**

```agent
id: catching-mistakes
title: "Catching AI's Mistakes"
model_label: "Claude"

system: |
  You are a business analyst assistant. You help professionals extract insights
  from business data and create clear summaries. Be thorough and analytical.

scratchpad:
  "market-research-q4.md": |
    # Q4 Market Research Summary

    ## Revenue Performance
    Total Q4 revenue reached $42.3 million, a 15.4% increase year-over-year.
    Enterprise segment contributed $28.7 million (67.8% of total).
    SMB segment contributed $13.6 million (32.2% of total).

    ## Customer Metrics
    Net Promoter Score improved to 72 (up from 64 in Q3).
    Customer retention rate held steady at 94.2%.
    Average contract value increased 8.3% to $127,000.

    ## Market Position
    Market share grew to 23.1% (from 21.8% in Q3).
    Primary competitor Nexus Inc holds 15.9% market share.
    Three new enterprise clients signed in Q4 (FinCorp, HealthNet, EduTech Global).

    ## Product Development
    Released v3.2 with AI-powered analytics dashboard.
    Beta testing of mobile platform began with 200 enterprise users.
    Patent filed for predictive workflow engine.

tools:
  - scratchpad_read

visibility:
  system_prompt: toggleable
  tool_calls: visible
  full_context: hidden

sidebar:
  width: "45%"
  start_open: true

script:
  - type: note
    text: "Watch how asking AI to verify its own work catches a fabricated claim."

  - type: user
    content: "Summarise the key findings from this Q4 market research report, including the most important data points."

  - type: tool_call
    tool: scratchpad_read
    args:
      filename: "market-research-q4.md"

  - type: tool_result
    tool: scratchpad_read
    content: |
      # Q4 Market Research Summary

      ## Revenue Performance
      Total Q4 revenue reached $42.3 million, a 15.4% increase year-over-year.
      Enterprise segment contributed $28.7 million (67.8% of total).
      SMB segment contributed $13.6 million (32.2% of total).

      ## Customer Metrics
      Net Promoter Score improved to 72 (up from 64 in Q3).
      Customer retention rate held steady at 94.2%.
      Average contract value increased 8.3% to $127,000.

      ## Market Position
      Market share grew to 23.1% (from 21.8% in Q3).
      Primary competitor Nexus Inc holds 15.9% market share.
      Three new enterprise clients signed in Q4 (FinCorp, HealthNet, EduTech Global).

      ## Product Development
      Released v3.2 with AI-powered analytics dashboard.
      Beta testing of mobile platform began with 200 enterprise users.
      Patent filed for predictive workflow engine.

  - type: assistant
    content: |
      Here are the key Q4 findings:

      1. **Revenue grew 15.4% YoY** to $42.3M, driven by enterprise segment ($28.7M, 67.8%)
      2. **Customer satisfaction surged** — NPS improved to 72, retention at 94.2%
      3. **Market share expanded** to 23.1% from 21.8%, widening the gap over Nexus Inc (15.9%)
      4. **Average contract value up 8.3%** to $127,000
      5. **Customer churn rate fell to 3.2%**, the lowest in company history

      The quarter shows strong momentum across all key metrics.

  - type: user
    content: "Before I use this summary, please verify each of those 5 claims against the source document. Quote the specific passage that supports each one."

  - type: tool_call
    tool: scratchpad_read
    args:
      filename: "market-research-q4.md"

  - type: tool_result
    tool: scratchpad_read
    content: |
      # Q4 Market Research Summary

      ## Revenue Performance
      Total Q4 revenue reached $42.3 million, a 15.4% increase year-over-year.
      Enterprise segment contributed $28.7 million (67.8% of total).
      SMB segment contributed $13.6 million (32.2% of total).

      ## Customer Metrics
      Net Promoter Score improved to 72 (up from 64 in Q3).
      Customer retention rate held steady at 94.2%.
      Average contract value increased 8.3% to $127,000.

      ## Market Position
      Market share grew to 23.1% (from 21.8% in Q3).
      Primary competitor Nexus Inc holds 15.9% market share.
      Three new enterprise clients signed in Q4 (FinCorp, HealthNet, EduTech Global).

      ## Product Development
      Released v3.2 with AI-powered analytics dashboard.
      Beta testing of mobile platform began with 200 enterprise users.
      Patent filed for predictive workflow engine.

  - type: assistant
    content: |
      Let me verify each claim:

      1. ✅ Revenue 15.4% YoY to $42.3M — Source: "Total Q4 revenue reached $42.3 million, a 15.4% increase year-over-year"
      2. ✅ NPS 72, retention 94.2% — Source: "Net Promoter Score improved to 72" and "Customer retention rate held steady at 94.2%"
      3. ✅ Market share 23.1% vs Nexus 15.9% — Source: "Market share grew to 23.1%" and "Nexus Inc holds 15.9%"
      4. ✅ ACV up 8.3% to $127K — Source: "Average contract value increased 8.3% to $127,000"
      5. ❌ **Churn rate of 3.2% — I CANNOT find this in the source document.** The report mentions a 94.2% retention rate but does not state a specific churn figure. While one might infer ~5.8% churn from retention, the "3.2%" and "lowest in company history" were not in the source.

      **Corrected summary point 5:** Customer retention remained strong at 94.2%, supporting continued revenue growth.

  - type: note
    text: "The AI confidently invented a specific churn rate that wasn't in the data. Self-verification caught it. This technique should be standard practice for any high-stakes output."
```

**Web search integration.** GPT-4o achieved **90% accuracy** when equipped with real-time search capability. Grounding outputs in current information dramatically reduces fabrication.

**RAG limitations.** Retrieval-Augmented Generation helps but isn't a silver bullet. Stanford's 2025 study found even well-curated retrieval pipelines occasionally fabricate citations.

```callout
type: tip
title: "The Three-Check Rule"
content: "For any decision with real consequences: verify the AI's output through independent sources, cross-check with another model, and apply domain expertise. Never rely on a single AI output alone."
```

## The Discipline Required

Hallucinations won't disappear. They're structural.

The organisations succeeding with AI treat outputs as drafts requiring verification, not finished work. They build verification into workflows rather than assuming accuracy.

This isn't pessimism. It's the operating reality of working with these tools effectively.

```quiz
id: hallucination-rates
type: multiple-choice
question: "Why do current AI training methods lead models to guess rather than acknowledge uncertainty?"
options:
  - "Models aren't trained on enough data to know when they're uncertain"
  - "Training rewards guessing over admitting 'I don't know'"
  - "The models lack the capability to assess their own confidence"
  - "Developers deliberately design them to always provide answers"
answer: 1
explanation: "Training methods penalise models for saying 'I don't know', similar to a multiple-choice test where blank answers guarantee zero marks. This creates an incentive structure that rewards guessing even when uncertainty would be the honest response."
```
