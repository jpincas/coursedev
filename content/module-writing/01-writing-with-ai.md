---
title: "Writing with AI: Beyond 'Write Me a Report'"
duration: "10m"
tags: [writing, documents, workflow]
---

# Writing with AI: Beyond "Write Me a Report"

"Write me a report on X" produces mediocre output. AI writing has matured beyond single-prompt generation.

A 2025 Nature survey found **57% of scientists** now use AI writing help. The technique matters.

## The Sectional Drafting Method

**Never ask for an entire long document in one prompt.**

The most reliable approach:

1. **Create a detailed outline first.** List sections, subsections, key points. Get AI's suggestions but own the structure.
2. **Expand each section individually.** One prompt per section, with specific constraints: word count, style, data requirements.
3. **Merge and smooth transitions.** Connect the sections, adjust flow, ensure consistency.

This approach produces coherent, focused documents that don't meander.

## The Optimal Human-AI Workflow

Stanford HAI research on 1,440 stories found that LLM collaboration increased productivity and reduced errors. But extensive AI collaboration reduced writers' sense of ownership.

![Human-AI Writing Workflow](/content/module-writing/images/writing-workflow.svg)

The pattern that works:

- **Human defines** the topic and goal
- **AI generates** ideas and outlines
- **Human selects** the approach
- **AI creates** section-by-section drafts
- **Human identifies** weak sections with specific feedback
- **AI revises** targeted sections
- **AI checks** grammar and style
- **Human ensures** voice and accuracy
- **Human fact-checks** all claims

AI accelerates. Humans own.

```callout
type: warning
title: "The Voice Problem"
content: "An editor at FoxPrint Editorial warned that a talented writer's revision became 'stripped of voice, a bland soup of pretty and overfamiliar phrasing' after AI polishing. Use AI for mechanics, not voice."
```

## Where AI Fits in Writing

**AI excels at:**
- Generating alternative phrasings
- Expanding bullet points into paragraphs
- Checking grammar and clarity
- Reorganising content
- Creating first drafts from outlines

**AI struggles with:**
- Authentic personal voice
- Subtle cultural context
- Original argumentation
- Strategic thinking about audience
- Knowing what to emphasise

The writer who uses AI well knows which is which.

## Sectional Drafting in Practice

The sectional method sounds simple, but seeing it work makes the difference clear. Watch the process: outline first, then expand one section at a time.

```agent
id: sectional-drafting-demo
title: "Section-by-Section Document Creation"
model_label: "Claude"

system: |
  You are a business writer. Create professional documents
  section by section. Follow the outline and constraints for
  each section precisely.

scratchpad:
  "project-brief.txt": |
    Project: Customer Portal Redesign
    Timeline: Q2 2026 (April-June)
    Budget: £85,000
    Team: 2 designers, 3 developers, 1 PM
    Goal: Reduce support tickets by 30% through self-service improvements
    Key problems:
    - Current portal is 5 years old, mobile-unfriendly
    - 60% of support tickets are "how do I find X" questions
    - Customer satisfaction score for portal: 2.8/5
    Stakeholders: VP Product, Head of Support, CTO

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
    text: "Step 1 of sectional drafting: create the outline. The user controls the structure — AI suggests, human decides."

  - type: user
    content: "Read the project brief. I need a proposal document for the stakeholders. First, suggest an outline — section titles and what each should cover. Don't write the document yet."

  - type: tool_call
    tool: scratchpad_read
    args:
      filename: "project-brief.txt"

  - type: tool_result
    tool: scratchpad_read
    content: |
      Project: Customer Portal Redesign
      Timeline: Q2 2026 (April-June)
      Budget: £85,000
      Team: 2 designers, 3 developers, 1 PM
      Goal: Reduce support tickets by 30% through self-service improvements
      Key problems:
      - Current portal is 5 years old, mobile-unfriendly
      - 60% of support tickets are "how do I find X" questions
      - Customer satisfaction score for portal: 2.8/5
      Stakeholders: VP Product, Head of Support, CTO

  - type: assistant
    content: "Here's the proposed outline:\n\n**1. Executive Summary** — 1 paragraph. The problem, the solution, the ask.\n\n**2. The Problem** — Current state: 5-year-old portal, mobile issues, support ticket burden. Data: 60% 'how do I find X' tickets, 2.8/5 satisfaction. Cost of inaction.\n\n**3. Proposed Solution** — What we'll build. Mobile-first redesign, self-service features, search improvements. High-level approach.\n\n**4. Expected Impact** — 30% ticket reduction target. Projected satisfaction improvement. Support cost savings.\n\n**5. Timeline & Budget** — Q2 phasing (April design, May build, June launch). £85K breakdown by phase.\n\n**6. Team & Next Steps** — Who's involved, what we need to start, approval ask.\n\nShall I adjust the structure before I start writing?"

  - type: user
    content: "Good structure. Now write section 2 only — 'The Problem.' Make it compelling for a CTO and VP Product. Use the data. 150 words max. Don't write any other section."

  - type: tool_call
    tool: scratchpad_write
    args:
      filename: "proposal-draft.md"
      content: |
        ## The Problem

        Our customer portal is five years old and it shows. It wasn't designed for mobile — yet 58% of our users now access it on phones. The result: frustrated customers who can't find what they need.

        The numbers tell the story. 60% of our support tickets are variations of "how do I find X" — questions the portal should answer without human intervention. Our portal satisfaction score sits at 2.8 out of 5, the lowest of any customer touchpoint.

        Every ticket costs us roughly £12 in staff time. At current volume, that's approximately £180K per year spent answering questions a well-designed portal would handle automatically.

        The portal isn't just outdated. It's actively driving support costs up and customer satisfaction down.

  - type: tool_result
    tool: scratchpad_write
    content: "Written to proposal-draft.md (12 lines)"

  - type: assistant
    content: "Section 2 written — 130 words. It leads with the mobile problem, uses all three data points (60% tickets, 2.8/5 satisfaction, 5-year age), and ends with the cost argument that'll resonate with the CTO. Ready for Section 3?"

  - type: note
    text: "Notice: the user controls the structure, reviews each section, and gives specific constraints (audience, word count, data to include). The AI writes one focused section at a time. Each section is tight because it had clear, narrow instructions."

  - type: note
    text: "Compare this to 'Write me a proposal about the portal redesign' — which would produce a 1,000-word generic document that meanders. Section by section, you get a 1,000-word precise document where every paragraph earns its place."
```

```quiz
id: sectional-drafting-quiz
type: multiple-choice
question: "Why does sectional drafting outperform single-prompt generation for long documents?"
options:
  - "It produces more words per minute"
  - "It allows specific constraints per section and maintains focus"
  - "It uses less of the AI's context window"
  - "It requires less human oversight"
answer: 1
explanation: "Sectional drafting lets you give specific constraints for each section — word count, style, required data — and keeps each section focused. Single prompts for long documents produce unfocused, meandering output."
```
