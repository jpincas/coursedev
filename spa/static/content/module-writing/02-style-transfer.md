---
title: "Making AI Write Like You"
duration: "12m"
tags: [style, voice, writing]
---

# Making AI Write Like You

AI writes in a generic voice by default. You can train it to write like you.

## Style Cards

A **style card** is a reusable prompt encoding your preferences.

It defines:
- **Tone** — formal, conversational, direct, warm
- **Vocabulary** — technical terms, plain language, industry jargon
- **Sentence structure** — short punchy sentences, flowing prose, bullet lists
- **Audience** — experts, general public, executives, students

Save it. Reuse it. Refine it over time.

## Grounding with Examples

The fastest way to teach AI your style:

**Provide 2-3 paragraphs of your best writing and ask AI to analyse your style.**

Give it recent work you're proud of. Ask:
- "Analyse the tone and structure of this writing."
- "What vocabulary patterns do you notice?"
- "How would you describe the sentence rhythm?"

Then: "Write the next section using this style."

```callout
type: tip
title: "Practical Approach"
content: "Start every writing project by giving AI 2-3 samples of your best work and asking it to analyse your style before drafting. This 2-minute investment transforms output quality."
```

## The EchoWriting Technique

For persistent style transfer:

![EchoWriting Process](/content/module-writing/images/echowriting-process.svg)

1. **Feed AI 15-20 samples** of your writing (emails, reports, articles)
2. **Have it analyse your style patterns** — what makes your writing recognisably yours
3. **Create a persistent style prompt** you reuse across sessions

Store this prompt in a Claude Project, ChatGPT Custom Instructions, or your persistent instruction file.

Every conversation starts with your style already loaded.

```agent
id: echowriting-demo
title: "Writing in Your Voice"
model_label: "Claude"

system: |
  You are a writing assistant. When analysing writing style, identify concrete
  patterns in tone, sentence structure, vocabulary, and formatting. When writing
  in a specific style, match those patterns precisely.

scratchpad:
  "sample1.txt": |
    Subject: Q3 Dashboard Access

    The analytics team needs dashboard access by Friday.

    Current blockers:
    - IT hasn't provisioned accounts
    - Training docs aren't ready

    Can you escalate the IT request? I'll finish the docs by Wednesday.

    Let me know.

    Sarah
  "sample2.txt": |
    ## Project Status — Week 12

    We hit the milestone. API integration is live. Two issues surfaced in testing.

    The authentication timeout is too short. Users get logged out mid-session.
    Fixed in the next sprint.

    The export function crashes on large datasets. Dev team is investigating.
    Workaround: split exports into smaller chunks.

    Next review: March 15.
  "sample3.txt": |
    Thanks for the proposal.

    The timeline looks tight. Can we push the launch to May? That gives us
    three extra weeks for testing.

    Budget is approved. The £15K contingency should cover the contractor costs.

    One question: who owns the post-launch support? We need that defined before
    we sign off.

    Let's discuss Friday.
  "brief.txt": |
    CONTENT TO WRITE:
    Project update email to the executive team

    KEY POINTS:
    - Phase 1 completed on schedule (API migration, user authentication)
    - Phase 2 delayed by 2 weeks due to unexpected database schema issues
    - Budget currently under by £8,000 but may need contingency for Phase 2
    - Client satisfied with progress, minor feature requests for Phase 3
    - Next milestone: Phase 2 delivery, now scheduled for April 30

    CONTEXT:
    This is a monthly update. Executives want facts, not fluff. They care about
    timeline, budget, and risks. Keep it professional but not corporate-speak.

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
    text: "This is EchoWriting in action. Three writing samples sit in the file explorer — all written by the same person. Notice the consistent patterns: short paragraphs, direct tone, British English, no jargon."

  - type: note
    text: "Watch the AI analyse the style, then use that analysis to write new content that sounds like the same person wrote it."

  - type: user
    content: "Please read these three writing samples and analyse my writing style. What patterns do you notice?"

  - type: assistant
    content: "I'll review all three samples and identify the patterns that define your writing style."

  - type: tool_call
    tool: scratchpad_read
    args:
      filename: "sample1.txt"

  - type: tool_result
    tool: scratchpad_read
    content: |
      Subject: Q3 Dashboard Access

      The analytics team needs dashboard access by Friday.

      Current blockers:
      - IT hasn't provisioned accounts
      - Training docs aren't ready

      Can you escalate the IT request? I'll finish the docs by Wednesday.

      Let me know.

      Sarah

  - type: tool_call
    tool: scratchpad_read
    args:
      filename: "sample2.txt"

  - type: tool_result
    tool: scratchpad_read
    content: |
      ## Project Status — Week 12

      We hit the milestone. API integration is live. Two issues surfaced in testing.

      The authentication timeout is too short. Users get logged out mid-session.
      Fixed in the next sprint.

      The export function crashes on large datasets. Dev team is investigating.
      Workaround: split exports into smaller chunks.

      Next review: March 15.

  - type: tool_call
    tool: scratchpad_read
    args:
      filename: "sample3.txt"

  - type: tool_result
    tool: scratchpad_read
    content: |
      Thanks for the proposal.

      The timeline looks tight. Can we push the launch to May? That gives us
      three extra weeks for testing.

      Budget is approved. The £15K contingency should cover the contractor costs.

      One question: who owns the post-launch support? We need that defined before
      we sign off.

      Let's discuss Friday.

  - type: assistant
    content: |
      Clear patterns across all three samples:

      **Sentence structure:** Short, punchy sentences. Rarely more than 15 words. Frequent single-sentence paragraphs for emphasis.

      **Tone:** Direct and no-nonsense. You state facts, identify problems, propose solutions. No hedging language ("might", "perhaps", "possibly").

      **Vocabulary:** Plain language throughout. British English ("realised" not "realized", "£" not "$"). Active voice dominates. Minimal jargon.

      **Formatting:** Frequent paragraph breaks for readability. Bullet points for lists. Headers in sentence case. No long blocks of text.

      **Voice:** Confident and action-oriented. Questions are direct requests, not suggestions. You own the information.

  - type: user
    content: "Perfect. Now write a project update email using the brief in brief.txt, but match my writing style exactly."

  - type: assistant
    content: "I'll read the brief and write the update in your style — short paragraphs, direct tone, British English, active voice."

  - type: tool_call
    tool: scratchpad_read
    args:
      filename: "brief.txt"

  - type: tool_result
    tool: scratchpad_read
    content: |
      CONTENT TO WRITE:
      Project update email to the executive team

      KEY POINTS:
      - Phase 1 completed on schedule (API migration, user authentication)
      - Phase 2 delayed by 2 weeks due to unexpected database schema issues
      - Budget currently under by £8,000 but may need contingency for Phase 2
      - Client satisfied with progress, minor feature requests for Phase 3
      - Next milestone: Phase 2 delivery, now scheduled for April 30

      CONTEXT:
      This is a monthly update. Executives want facts, not fluff. They care about
      timeline, budget, and risks. Keep it professional but not corporate-speak.

  - type: tool_call
    tool: scratchpad_write
    args:
      filename: "project-update-email.txt"
      content: |
        Subject: Project Update — March

        Phase 1 is done. API migration and user authentication delivered on schedule.

        Phase 2 is delayed by two weeks. Database schema issues took longer to resolve
        than expected. New delivery date: April 30.

        Budget status: £8,000 under plan. We may need contingency for Phase 2.
        I'll confirm by next week.

        Client feedback is positive. They've requested minor features for Phase 3.
        Nothing that affects current timeline.

        Risks:
        - Phase 2 delay pushes final launch to mid-May
        - Database work may uncover more issues

        Next milestone: Phase 2 delivery, April 30.

        Questions?

  - type: tool_result
    tool: scratchpad_write
    content: "Written to project-update-email.txt (17 lines)"

  - type: assistant
    content: "Done. I've written the update matching your style — short paragraphs, direct statements, British English, active voice, bullet-free structure with clear paragraph breaks for readability. Same confident, no-nonsense tone as your samples."

  - type: note
    text: "Compare the AI's output to the original samples. Same rhythm. Same directness. Same paragraph structure. The AI learned your voice from three examples."

  - type: note
    text: "This is EchoWriting. Feed AI your writing samples, have it analyse the patterns, then every document it creates sounds like you wrote it. No generic AI voice. Your voice."
```

## Multi-Document Synthesis

AI's highest-value use case for knowledge work: synthesising across multiple sources.

Upload multiple documents to Claude or ChatGPT. Ask it to:
- Identify common themes
- Spot contradictions
- Find gaps in coverage
- Extract key arguments

Google's NotebookLM specialises in this — it analyses information across documents, websites, and media types simultaneously.

```callout
type: warning
title: "Synthesis Limitation"
content: "AI is better at organising and summarising than at genuine intellectual synthesis. Always verify it hasn't fabricated connections between sources. Check that claimed patterns actually exist in the source material."
```

## Style vs Substance

AI can mimic your style. It cannot replicate your judgment.

**Your style:** sentence structure, word choice, paragraph rhythm. AI can learn this.

**Your judgment:** what to say, what to emphasise, what argument to make. AI cannot replicate this.

Use AI to write in your voice. Never delegate deciding what that voice should say.

```quiz
id: style-transfer-quiz
type: multiple-choice
question: "What is the EchoWriting technique?"
options:
  - "Writing complete drafts yourself and having AI convert them into different formats"
  - "Feeding AI 15-20 samples of your writing to analyse patterns, then creating a persistent style prompt"
  - "Asking AI to repeat your exact words back to you to verify comprehension"
answer: 1
explanation: "EchoWriting involves giving AI many samples of your writing, having it analyse your style patterns, then creating a reusable style prompt that makes AI write like you across all future sessions."
```
