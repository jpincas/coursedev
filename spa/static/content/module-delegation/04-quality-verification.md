---
title: "Quality Verification"
duration: "10m"
tags: [verification, quality, review]
---

# Quality Verification

You need to verify AI output. But how do you verify work when you're not an expert in that specific area?

## The Verification Framework

You don't need to be an expert in the content. You need a thoughtful verification strategy.

![Four-Step Verification Framework](/content/module-delegation/images/verification-framework.svg)

### 1. Spot-Check Facts

Pick 3-5 specific claims and verify them independently.
- Check a cited statistic against the source
- Verify a date or name mentioned
- Confirm a technical detail

If these spot-checks pass, confidence in the rest increases. If they fail, you know deeper review is needed.

**Advanced technique: Multi-model cross-validation**
Run the same query through multiple LLMs (Claude, GPT, Gemini). If they agree, confidence increases. If they diverge, investigate further.

### 2. Check for Coherence

Does it make internal sense?
- Are there contradictions?
- Does the logic flow?
- Do conclusions follow from evidence?
- Are there gaps in reasoning?

You don't need expertise to notice "this part contradicts that part."

### 3. Validate Structure

Does it have what you asked for?
- All required sections present?
- Format matches requirements?
- Length appropriate?
- Tone correct?

This is verification against your specification, not verification of content accuracy.

### 4. Test Against Purpose

Would this actually work?
- Does it achieve the goal?
- Would the audience find it useful?
- Can you use it for its intended purpose?

The ultimate test: Is this fit for purpose?

```callout
type: tip
title: "Verification Mindset"
content: "You're not checking if AI is 'right.' You're checking if the output is useful, accurate enough for your purposes, and fit for the intended use."
```

### 5. AI as Editor

One of the most effective verification techniques: **ask AI to review its own work.**

After AI produces content, don't just accept it. Commission a second pass:

> "Now re-read what you just wrote from the perspective of [target audience]. What would you change?"

This triggers a genuine second pass, not rubber-stamping. AI will catch:
- Jargon the audience won't understand
- Logic gaps in its own reasoning
- Missing context it assumed
- Tone mismatches

**Multi-pass review:** Have AI review its work through different lenses. Ask it to read once for accuracy, once for clarity, once for completeness, once for tone. Each pass catches different issues.

```callout
type: tip
title: "Different Lenses, Different Catches"
content: "Accuracy catches factual errors. Clarity catches confusing explanations. Completeness catches missing sections. Tone catches inappropriate language. One review pass rarely catches all four."
```

### 6. Visual Verification

For visual outputs — websites, formatted documents, presentations — **use browser tools to actually LOOK at what was produced.**

Don't just trust the code. Render it and check.

Common visual issues AI won't catch in code review:
- Misaligned elements
- Broken responsive layouts
- Color contrast problems
- Overflowing text
- Missing images or broken links

If AI built a webpage, open it in a browser. If AI formatted a document, export it and review the PDF. Code correctness ≠ visual quality.

## Practical Tips

**Don't verify everything**
That defeats the purpose of delegation. Verify enough to have confidence.

**Verify high-stakes claims more carefully**
A wrong statistic in a board presentation matters more than a typo in internal notes.

**Use AI to help verify**
Ask AI to fact-check its own claims, or use a different AI session to review. Research shows asking AI to verify its own reasoning before providing final answers reduces errors by roughly 17%.

**For data analysis, always inspect the code**
When AI generates charts or statistics, click "view analysis" to inspect the generated code. Claude tends to hallucinate when working with large datasets or too many filters. The World Economic Forum warned: "When foundational data is fragmented or inaccurate, AI models generate outputs that appear sophisticated but are fundamentally wrong."

**Web search as verification**
Use web search integration when available. AI with search access achieves much higher accuracy than AI working from training data alone.

**Trust but verify**
AI output is generally good. But "generally" isn't "always." Verification is your responsibility.

## The Four Steps in Practice

The verification framework works even when you're not an expert. Watch someone apply all four steps to an AI-generated competitive analysis.

```agent
id: verification-demo
title: "Verifying AI Output Step by Step"
model_label: "Claude"

system: |
  You are a market research analyst. Produce thorough analysis.
  When asked to verify, examine claims critically.

scratchpad:
  "competitor-intel.txt": |
    Competitor: NovaTech Solutions
    Last funding: Series B, £12M (March 2025)
    Headcount: ~120 (LinkedIn)
    Key product: AutoFlow — workflow automation for legal teams
    Pricing: £45/user/month (Pro), £89/user/month (Enterprise)
    Recent moves:
    - Launched AI contract review feature (Nov 2025)
    - Hired VP Sales from Salesforce (Jan 2026)
    - Opened Sydney office for APAC expansion
    Known weaknesses: slow onboarding (avg 6 weeks), no mobile app
    Customer base: ~200 firms, mostly mid-market

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
    text: "First, the AI produces a competitive briefing. Then we'll walk through all four verification steps: spot-check facts, check coherence, validate structure, test against purpose."

  - type: user
    content: "Read the competitor intel file and produce a competitive threat assessment for our leadership team. Include threat level, likely next moves, and our recommended response."

  - type: tool_call
    tool: scratchpad_read
    args:
      filename: "competitor-intel.txt"

  - type: tool_result
    tool: scratchpad_read
    content: |
      Competitor: NovaTech Solutions
      Last funding: Series B, £12M (March 2025)
      Headcount: ~120 (LinkedIn)
      Key product: AutoFlow — workflow automation for legal teams
      Pricing: £45/user/month (Pro), £89/user/month (Enterprise)
      Recent moves:
      - Launched AI contract review feature (Nov 2025)
      - Hired VP Sales from Salesforce (Jan 2026)
      - Opened Sydney office for APAC expansion
      Known weaknesses: slow onboarding (avg 6 weeks), no mobile app
      Customer base: ~200 firms, mostly mid-market

  - type: assistant
    content: "**NovaTech Solutions — Competitive Threat Assessment**\n\n**Threat Level: MODERATE-HIGH (trending upward)**\n\nNovaTech has shifted from a workflow tool to an AI-powered legal platform. Three signals indicate escalation:\n\n1. **AI pivot is real.** The contract review feature (Nov 2025) directly competes with our core offering. With £12M in funding, they have runway to iterate.\n\n2. **Sales muscle is increasing.** The Salesforce VP hire signals enterprise ambitions. Expect aggressive outbound and potential pricing pressure in Q2.\n\n3. **Geographic expansion.** The Sydney office targets APAC — a market where we have limited presence. First-mover advantage in the region could lock us out.\n\n**Likely next 6 months:**\n- Enterprise push with dedicated sales team\n- Mobile app launch (their biggest customer complaint)\n- Pricing aggression to win enterprise logos at scale\n- APAC market entry targeting Australian and Singapore law firms\n\n**Recommended response:**\n- Accelerate our AI contract features to maintain differentiation\n- Protect existing enterprise accounts with retention offers\n- Evaluate APAC expansion timeline — delay costs market share\n- Monitor their mobile launch for competitive positioning"

  - type: note
    text: "Looks professional and actionable. Now let's verify it using the four-step framework."

  - type: user
    content: "Good first draft. Now help me verify this. Let's go step by step.\n\nStep 1: Spot-check facts. Are there any claims in your analysis that go beyond what's in the source data?"

  - type: assistant
    content: "Good catch — let me audit each claim against the source:\n\n**Verified from source:**\n- £12M Series B ✓\n- AI contract review feature Nov 2025 ✓\n- VP Sales from Salesforce Jan 2026 ✓\n- Sydney office / APAC ✓\n- Slow onboarding and no mobile app ✓\n- ~200 firms, mid-market ✓\n\n**Claims I extrapolated (not directly in the data):**\n- 'Shifted from workflow tool to AI-powered legal platform' — my interpretation, not stated\n- 'Pricing aggression to win enterprise logos' — speculation, not evidenced\n- 'Mobile app launch' as a likely next move — reasonable inference from weakness, but not confirmed\n- 'Targeting Australian and Singapore law firms' — I added the specific countries; source only says 'APAC expansion'\n\n**I should flag these as inferences, not facts.** The analysis mixed sourced information with my predictions without distinguishing between them."

  - type: user
    content: "Step 2: Coherence check. Any internal contradictions?"

  - type: assistant
    content: "One tension:\n\nI described the threat as 'MODERATE-HIGH (trending upward)' but then listed four aggressive moves as 'likely next 6 months.' If all four happen, the threat level is HIGH, not moderate-high. The rating and the evidence don't fully align.\n\nAlso: I recommended 'protect existing enterprise accounts' but the source says NovaTech is 'mostly mid-market.' If they're mid-market, our enterprise accounts may not be their immediate target. The defensive recommendation may be premature."

  - type: note
    text: "Two verification steps complete and already caught: unsourced claims presented as facts, speculative predictions without labels, a threat rating that contradicts its own evidence, and a defensive recommendation based on a faulty assumption. The four-step framework works even though neither you nor the AI is a competitive intelligence expert."

  - type: note
    text: "Steps 3 (validate structure) and 4 (test against purpose) are straightforward: Does it have all the sections requested? Would the leadership team find it actionable? You can check those yourself in seconds. The hard work — spotting unsourced claims and internal contradictions — is where the framework earns its value."
```

```quiz
id: verification-method
type: multiple-choice
question: "How should you verify AI output when you're not an expert in the topic?"
options:
  - "Trust that AI is generally accurate"
  - "Use a verification strategy: spot-check, coherence, structure, purpose"
  - "Only use AI for topics you're expert in"
answer: 1
explanation: "You don't need topic expertise to verify effectively. A verification strategy — spot-checking facts, checking coherence, validating structure, testing against purpose — lets you assess quality without deep domain knowledge."
```
