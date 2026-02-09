---
title: "Legal and Copyright: What You Must Know"
duration: "9m"
tags: [legal, copyright, liability]
---

# Legal and Copyright: What You Must Know

The legal landscape around AI is evolving rapidly. What you don't know can be expensive.

## High-Stakes Copyright Litigation

In December 2023, The New York Times filed a landmark lawsuit against OpenAI and Microsoft, alleging copyright infringement through the use of millions of Times articles to train AI models.

The case remains ongoing as of February 2026.

The lawsuit seeks billions in statutory and actual damages. It represents the first major legal challenge from a content publisher with the resources to pursue litigation to precedent-setting conclusion.

The case will likely establish critical legal principles: whether training on copyrighted material constitutes fair use, whether AI companies can claim transformative use, and what damages apply when copyrighted works appear in training data without authorisation.

## The Copyright Question: Who Owns AI Output?

The US Copyright Office ruled in January 2025. The decision matters globally because it reflects emerging consensus.

**AI-generated outputs receive copyright protection only where a human author determined sufficient expressive elements.**

What this means in practice:
- Providing a prompt alone is insufficient for copyright protection
- Iterative refinement with creative direction can establish authorship
- The more human creative input, the stronger the copyright claim
- Pure AI output with minimal human involvement gets no protection

**No jurisdiction recognises AI as a legal person capable of holding copyright.**

```callout
type: warning
title: "Never Say 'AI Said To'"
content: "If you use AI output in your work, it becomes yours — and so does the legal responsibility. You cannot claim 'the AI made the mistake' as a defence. You chose to use the output."
```

## International Divergence

The legal landscape varies dramatically by jurisdiction.

**Permissive jurisdictions:**
- **Japan:** Permits training on copyrighted material without consent
- **Singapore:** Similar permissive stance, positioning as AI development hub

**Restrictive jurisdictions:**
- **France:** Fined Google €250 million for training on news articles without authorisation
- **European Union:** Implementing stricter controls through the AI Act

**United States:** Case-by-case approach through courts. The New York Times lawsuit and similar ongoing cases will determine whether courts treat AI training as fair use or infringement.

This creates compliance complexity for global organisations. What's legal in Tokyo may be illegal in Paris.

![International legal landscape spectrum](/content/module-risks/images/legal-landscape.svg)

## Practical Implications

**For content creation:**
- AI-assisted writing is legally acceptable if you exercise creative control
- You must be able to demonstrate human authorship for copyright protection
- The more you iterate and refine, the stronger your authorship claim

**For report writing:**
- Using AI to draft sections is acceptable with human review and editing
- You remain responsible for accuracy, even if AI generated the text
- Attribution of AI use varies by organisation — check internal policies

**For code generation:**
- Generated code may include patterns from copyrighted training data
- Enterprise AI tools typically provide IP indemnification; consumer tools do not
- Code review remains essential for quality and to establish human authorship

```callout
type: tip
title: "The Emerging Consensus"
content: "Human oversight and meaningful creative control leads to protectable work. Minimal prompting with no refinement does not. Courts are converging on this principle globally."
```

## The Liability Question

When AI output causes harm, who is liable?

Current legal thinking places responsibility on the **person who chose to use and publish the output.**

If AI-generated analysis leads to a poor business decision, the organisation is liable, not the AI vendor. If AI-drafted content contains defamatory claims, the publisher is liable, not the model creator.

**The "intermediate product" framing:** Treat AI output as an intermediate product requiring quality control, not a finished deliverable. This positions your workflow correctly for both legal and quality purposes.

## What Your Organisation Needs

**Clear policies on:**
- When AI use must be disclosed to clients or stakeholders
- What review processes apply to AI-generated work
- Who has authority to approve AI output for external publication
- How to document human creative input for copyright purposes

**Training on:**
- The "never say AI said to" principle
- When to disclose AI use
- How to establish sufficient human authorship
- Verification requirements before publication

**Legal review of:**
- AI vendor contracts, particularly IP indemnification clauses
- Data sources used by your chosen AI tools
- Compliance requirements in your jurisdictions

The legal framework is stabilising around a clear principle: AI is a tool, humans are responsible for how they use it.

```quiz
id: legal-copyright-ownership
type: multiple-choice
question: "According to the US Copyright Office ruling, when does AI-generated output receive copyright protection?"
options:
  - "Always, because the AI created original content"
  - "Never, because AI cannot hold copyright"
  - "Only when a human determined sufficient expressive elements through creative control"
  - "Only when the AI's training data was all legally licensed"
answer: 2
explanation: "The Copyright Office ruled that AI output gets protection only where human authorship is demonstrated through determining sufficient expressive elements. This means meaningful human creative input and control, not just providing a prompt. The key is the human's creative contribution, not the AI's."
```
