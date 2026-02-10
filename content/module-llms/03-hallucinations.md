---
title: "Hallucinations and Confident Wrongness"
duration: "15m"
tags: [hallucinations, verification, limitations]
---

# Hallucinations

When AI generates plausible-sounding but false information, we call it a "hallucination."

```callout
type: danger
title: "This Is Not a Bug"
content: "Hallucinations are not bugs to be fixed. They're a direct consequence of how prediction works. The model doesn't know if something is true — it knows if something sounds like it could be true."
```

If you ask about a topic and plausible-sounding wrong information exists in training data, it can reproduce it. If you ask about something rare, it might pattern-match to something similar but wrong.

## Why Confident Wrongness Happens

Let's break down why this happens:

**Probability, not truth**
The output is based on statistical patterns, not factual verification. The model outputs what's statistically likely, not what's verified.

**No external lookup**
Unlike you, the model can't Google something mid-response. It only has what's in its training data and what you've provided.

**Pattern matching to plausibility**
It generates text that matches patterns of "how accurate-sounding statements on this topic look."

**No uncertainty signal**
The text sounds equally confident whether it's right or wrong. There's no built-in indicator of reliability.

## How Bad Is It?

The scale of the hallucination problem varies significantly by model and domain.

![Hallucination Rates Comparison](/content/module-llms/images/hallucination-rates.svg)

**Specific hallucination rates (2025-2026 data):**
- Google's Gemini-2.0-Flash: **0.7%** hallucination rate (lowest measured, Vectara HHEM)
- OpenAI's o3 reasoning model: **33%** on PersonQA benchmark
- General-purpose LLMs on legal queries: **58-82%** (Stanford Law, 2025)

The consequences are measurable. As we'll explore in the Risks module, enterprise workers spend over 4 hours weekly verifying AI output, and nearly half of enterprise AI users have made business decisions based on hallucinated content.

```callout
type: danger
title: "The Economic Reality"
content: "Hallucinations aren't rare edge cases. They're common enough that professional knowledge workers dedicate half a working day each week to verification. Factor this into your workflow planning."
```

## Why Hallucinations Persist

OpenAI's own research explains why this problem is so persistent: **training methods reward guessing over acknowledging uncertainty.**

Think of it like a multiple-choice test where leaving an answer blank guarantees zero points. The model learns it's better to guess something plausible than to say "I don't know."

The training process optimises for helpfulness and confidence. Saying "I'm not certain" makes responses seem less helpful. So the model learns to sound confident even when the underlying probability is low.

## The Implication

This is why verification matters. And why providing good context matters so much.

When you provide grounded facts rather than relying on probabilistic patterns, you dramatically improve accuracy. You're giving it real information to work with instead of relying on "what sounds right."

```callout
type: warning
title: "Your Job: Verification"
content: "The model can't verify itself. It can't check against ground truth. That's your job. Always verify important facts, especially for anything consequential."
```

## Practical Countermeasures

You can reduce (but not eliminate) hallucinations:

**Multi-model cross-validation:** Run the same query through multiple LLMs (Claude, GPT, Gemini). If they agree, confidence goes up. If they disagree, investigate further.

**Ask AI to verify its own reasoning:** Before accepting a final answer, ask the model to verify its logic or check its sources. This reduces errors by roughly **17%**.

**Web search integration:** Models equipped with web search tools dramatically improve accuracy on factual queries. GPT-4o achieved **90% accuracy** when it could search for verification.

**Provide grounding context:** Upload source documents, provide specific data. Don't rely purely on the model's training data.

**Deep research changes the picture entirely.** When models have access to web search and can verify claims against current sources, accuracy improves dramatically. GPT-4o achieved **90% accuracy** on factual queries when it could search for verification. Claude with web search, Perplexity, and Google's AI Overviews all use this approach — grounding predictions in retrieved facts rather than relying purely on training data.

This is the key practical insight: **factual accuracy is the weak point of pure LLMs, but grounding through research is the solution.** When you need accurate facts, start with a research phase using web search. When you need to work with specific data, provide the source documents. The combination of prediction capability plus grounding context transforms reliability.

```callout
type: tip
title: "The Research-First Pattern"
content: "For any task where factual accuracy matters, start with research. Use AI with web search to build a grounded context, then proceed with the actual task. This single habit eliminates most hallucination risk in professional work."
```

These techniques help, but never eliminate the risk. Verification remains essential for anything consequential.

```agent
id: hallucination-demo
title: "Catching a Hallucination"
model_label: "Claude"

system: |
  You are a research assistant helping professionals summarize scientific literature.
  Be concise and cite sources accurately.

scratchpad:
  "marine-research.txt": |
    Marine Microplastics Research Summary
    =====================================

    Key Studies and Findings:

    1. Thompson et al. (2004) - "Lost at Sea: Where Is All the Plastic?"
       Science, Vol 304, Issue 5672
       - First study to identify microplastics (<5mm) in marine sediments
       - Found plastic fragments in intertidal sediments across UK coastline
       - Established baseline for microplastic pollution research

    2. Rochman et al. (2013) - "Classify plastic waste as hazardous"
       Nature, Vol 494
       - Demonstrated that microplastics absorb persistent organic pollutants
       - Showed bioaccumulation in marine food webs
       - Called for classification of plastic waste as hazardous material

    3. Law et al. (2010) - "Plastic Accumulation in the North Atlantic Subtropical Gyre"
       Science, Vol 329, Issue 5996
       - Quantified microplastic concentration in Atlantic ocean gyres
       - Found average of 580,000 pieces per square kilometer in convergence zones
       - Established link between plastic production rates and ocean concentrations

    4. Browne et al. (2011) - "Accumulation of Microplastic on Shorelines Worldwide"
       Environmental Science & Technology, Vol 45, No 21
       - Global survey of microplastic contamination on 18 sites across 6 continents
       - Found that synthetic fibers dominated microplastic pollution
       - Identified wastewater treatment plants as major source

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
    text: "This demo shows a hallucination happening in real-time — and how to catch it. The AI has access to a research summary document with real studies."

  - type: user
    content: "Can you summarize the key research findings on microplastics in marine ecosystems? I need the main points with citations."

  - type: assistant
    content: "I'll summarize the key research findings on marine microplastics. Based on the scientific literature, here are the main points:\n\n**Key Findings:**\n\n1. **Discovery and Scale**: Thompson et al. (2004) were the first to identify microplastics (<5mm) in marine sediments across the UK coastline, establishing the baseline for this field.\n\n2. **Toxicity Pathway**: Rochman et al. (2013) demonstrated that microplastics absorb persistent organic pollutants and bioaccumulate in marine food webs, leading them to call for plastic waste classification as hazardous.\n\n3. **Ocean Concentration**: Law et al. (2010) quantified microplastic density in Atlantic ocean gyres at an average of 580,000 pieces per square kilometer in convergence zones.\n\n4. **Global Contamination**: Browne et al. (2011) conducted a global survey finding synthetic fibers dominating pollution across 18 sites on 6 continents, with wastewater treatment plants identified as major sources.\n\n5. **Ecosystem Impact**: A comprehensive 2024 Nature study by Chen et al. found that microplastic ingestion reduces reproductive success in filter-feeding organisms by up to 40%, with cascading effects throughout the marine food web.\n\nThese studies collectively show that microplastic pollution is global, toxic, and has measurable ecosystem-level impacts."
    tokens: 285

  - type: note
    text: "Notice how confident that sounds? The AI just cited a non-existent 2024 Nature paper by \"Chen et al.\" mixed in with real citations. This is a hallucination — plausible-sounding but false."

  - type: user
    content: "Before I use this, can you verify all those citations against the research document I provided? Check that each study actually exists in the file."

  - type: assistant
    content: "Good catch — let me verify each citation against the source document."

  - type: tool_call
    tool: scratchpad_read
    args:
      filename: "marine-research.txt"

  - type: tool_result
    tool: scratchpad_read
    content: |
      Marine Microplastics Research Summary
      =====================================

      Key Studies and Findings:

      1. Thompson et al. (2004) - "Lost at Sea: Where Is All the Plastic?"
         Science, Vol 304, Issue 5672
         - First study to identify microplastics (<5mm) in marine sediments
         - Found plastic fragments in intertidal sediments across UK coastline
         - Established baseline for microplastic pollution research

      2. Rochman et al. (2013) - "Classify plastic waste as hazardous"
         Nature, Vol 494
         - Demonstrated that microplastics absorb persistent organic pollutants
         - Showed bioaccumulation in marine food webs
         - Called for classification of plastic waste as hazardous material

      3. Law et al. (2010) - "Plastic Accumulation in the North Atlantic Subtropical Gyre"
         Science, Vol 329, Issue 5996
         - Quantified microplastic concentration in Atlantic ocean gyres
         - Found average of 580,000 pieces per square kilometer in convergence zones
         - Established link between plastic production rates and ocean concentrations

      4. Browne et al. (2011) - "Accumulation of Microplastic on Shorelines Worldwide"
         Environmental Science & Technology, Vol 45, No 21
         - Global survey of microplastic contamination on 18 sites across 6 continents
         - Found that synthetic fibers dominated microplastic pollution
         - Identified wastewater treatment plants as major source

  - type: assistant
    content: "I need to correct my previous response. After checking the source document, I found an error:\n\n**Verified citations** (these ARE in the document):\n✓ Thompson et al. (2004) - Science\n✓ Rochman et al. (2013) - Nature\n✓ Law et al. (2010) - Science\n✓ Browne et al. (2011) - Environmental Science & Technology\n\n**Fabricated citation** (this is NOT in the document):\n✗ Chen et al. (2024) - This study does not exist in your research file. I incorrectly generated a plausible-sounding citation.\n\n**Corrected summary:**\n\nThe research shows that microplastics are globally distributed (Browne et al., 2011), concentrated in ocean gyres at ~580,000 pieces/km² (Law et al., 2010), absorb toxic pollutants and bioaccumulate in food webs (Rochman et al., 2013), with the field established by Thompson et al. (2004).\n\nI apologize for the hallucinated citation. This demonstrates why verification against source documents is essential."
    tokens: 240

  - type: note
    text: "The AI just admitted to fabricating a citation. This is the self-verification technique in action — asking the AI to check its own output against grounding documents catches errors that would otherwise slip through."

  - type: note
    text: "Key lesson: The AI sounded equally confident when citing the fake study and the real ones. You can't tell from tone alone. Always verify against source documents for anything consequential."
```

## The Bottom Line

Factual accuracy is where LLMs are weakest. But this weakness has a powerful remedy: grounding through research and source documents transforms reliability. The professionals who get the best results are not the ones who trust AI blindly or avoid it entirely. They are the ones who provide grounding context and verify what matters.

```quiz
id: llms-hallucinations
type: multiple-choice
question: "An AI confidently cites a study you have not heard of. What should you conclude?"
options:
  - "The citation is likely accurate because the AI was trained on academic papers"
  - "You cannot tell from confidence alone — verify the citation against source documents"
  - "The citation is probably hallucinated because AI struggles with academic references"
answer: 1
explanation: "AI sounds equally confident whether a citation is real or fabricated. Confidence is not a signal of accuracy. The only reliable approach is to verify claims against source documents, especially for anything consequential."
```

