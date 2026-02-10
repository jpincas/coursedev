---
title: "The Three Questions"
duration: "15m"
tags: [questions, specification, clarity]
---

# The Three Questions

![The Three Questions Framework](/content/module-prompting/images/three-questions.svg)

Here's an even simpler framework. Before any AI task, ask yourself these three questions.

## 1. What Does "Done" Look Like?

If you can't describe the end state, the AI can't produce it.

**Vague:**
"Help me with my presentation"

**Better:**
"Create slides for my presentation"

**Clear:**
"Create a 10-slide presentation on Q4 results for the sales team, with one key metric highlighted per slide, ending with Q1 priorities"

```callout
type: tip
title: "The Verification Test"
content: "Could you verify that the work is done correctly? If you can't describe what 'correct' looks like, your specification isn't clear enough."
```

## 2. What Context Does It Need?

What information is required for success?

**Data and facts**
The actual information to work with — spreadsheets, documents, code.

**Background and situation**
Company context, project history, prior decisions. Why does this task exist?

**Audience**
Who will use this? What do they already know? Executive? Developer? Customer?

**Examples**
Similar work you liked. Templates. Style references. Show, don't just tell.

**Requirements**
Explicit must-haves. Compliance needs. Brand guidelines.

The more relevant context you provide, the better the output.

## 3. What Are the Boundaries?

What constraints or limits apply?

**Length**
Word count, page count, number of items. "About 500 words" is better than nothing.

**Tone**
Professional? Casual? Technical? Explain-like-I'm-five?

**Format**
Bullets? Paragraphs? Table? Code? This dramatically affects usability.

**Inclusions**
What must be covered? "Must mention the new pricing model."

**Exclusions**
What to avoid? "Don't discuss the merger." "Skip implementation details."

**Constraints**
Technical limits, brand requirements, compliance needs.

Boundaries prevent the AI from going off in directions you don't want.

## Putting It Together

**Task:** Analyse customer feedback

**Question 1 — Done:**
A summary of the top 5 complaint categories with representative quotes

**Question 2 — Context:**
- CSV file with 500 feedback entries
- This is for the product team
- We're focused on our mobile app experience

**Question 3 — Boundaries:**
- Max 1 page
- Include frequency counts
- Professional but concise tone

**Result prompt:**
```
Analyse this customer feedback CSV (attached).
Identify the top 5 complaint categories with frequency counts.
For each category, provide 2-3 representative quotes.
Focus on mobile app experience issues.
Format as a brief report for the product team, max 1 page.
```

## The Three Questions Applied

Watch someone apply the three questions to a real task. Notice how each question adds precision to what started as a vague idea.

```agent
id: three-questions-demo
title: "From Vague Idea to Clear Request"
model_label: "Claude"

system: |
  You are a communications consultant helping craft internal
  communications. Write clearly and concisely for a professional
  audience.

scratchpad:
  "policy-changes.txt": |
    Changes effective 1 March 2026:
    - Remote work: Now 3 days office, 2 days remote (was 2/3)
    - Meeting-free Fridays: No internal meetings before 2pm
    - Expense policy: Pre-approval required for purchases over £200 (was £500)
    - Training budget: Individual budget increased to £2,000/year (was £1,200)
    - Parental leave: Extended to 26 weeks full pay (was 16 weeks)
  "previous-comms.txt": |
    Subject: December Policy Update
    Team,
    A few changes to note for the new year.
    [Three paragraphs of dense text, no formatting, buried the key points]
    Result: 40% of staff didn't read it. HR fielded 200+ questions.

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
    text: "The user has policy changes to communicate to staff. Instead of just saying 'write an email about these changes,' watch them apply the three questions to build a precise request."

  - type: user
    content: "I need to communicate some policy changes to the team. Read both files first — the changes themselves and the previous comms that didn't work."

  - type: tool_call
    tool: scratchpad_read
    args:
      filename: "policy-changes.txt"

  - type: tool_result
    tool: scratchpad_read
    content: |
      Changes effective 1 March 2026:
      - Remote work: Now 3 days office, 2 days remote (was 2/3)
      - Meeting-free Fridays: No internal meetings before 2pm
      - Expense policy: Pre-approval required for purchases over £200 (was £500)
      - Training budget: Individual budget increased to £2,000/year (was £1,200)
      - Parental leave: Extended to 26 weeks full pay (was 16 weeks)

  - type: tool_call
    tool: scratchpad_read
    args:
      filename: "previous-comms.txt"

  - type: tool_result
    tool: scratchpad_read
    content: |
      Subject: December Policy Update
      Team,
      A few changes to note for the new year.
      [Three paragraphs of dense text, no formatting, buried the key points]
      Result: 40% of staff didn't read it. HR fielded 200+ questions.

  - type: assistant
    content: "I've read both files. The previous communication was dense and unformatted — which led to poor engagement. I can see five policy changes, some positive (training budget, parental leave) and some that may be less welcome (office days, expense threshold). How would you like me to approach the communication?"

  - type: user
    content: "Here's what I need. I've thought through the three questions:\n\n**Done:** An all-staff email that everyone actually reads and understands. No follow-up questions to HR.\n\n**Context:** 200-person company. Mix of office and remote staff. Last comms failed because it was too dense. Some changes are popular (training, parental leave), some are unpopular (more office days, lower expense threshold).\n\n**Boundaries:** Under 300 words. Lead with the good news. Each change gets its own bullet with old vs new clearly shown. Warm but direct tone — don't sugarcoat the less popular changes. Subject line must grab attention."

  - type: tool_call
    tool: scratchpad_write
    args:
      filename: "policy-email.md"
      content: |
        Subject: 5 Policy Changes Starting March — Including More Training Budget & Extended Parental Leave

        Hi team,

        Five policy changes take effect on 1 March. Here's what's changing and why.

        **What's new (and good):**

        - **Training budget: £1,200 → £2,000/year.** We're investing more in your development. Use it for courses, conferences, certifications — your choice.
        - **Parental leave: 16 → 26 weeks full pay.** This brings us in line with what we believe is right for our team.
        - **Meeting-free Friday mornings.** No internal meetings before 2pm on Fridays. Protect your deep work time.

        **What's also changing:**

        - **Office days: 2 → 3 per week.** We've seen that collaboration quality drops below three days together. The new default is Monday, Tuesday, Thursday in-office. Talk to your manager if you need flexibility.
        - **Expense pre-approval: £500 → £200.** Finance needs better visibility on spending. Pre-approval is quick — submit via the portal, expect same-day response.

        Each change has a detailed FAQ on the intranet. If something's unclear, check there first.

        — The Leadership Team

  - type: tool_result
    tool: scratchpad_write
    content: "Written to policy-email.md (23 lines)"

  - type: assistant
    content: "Done — policy-email.md is ready. I led with the popular changes (budget increase, parental leave, meeting-free Fridays), then addressed the less welcome ones directly without spin. Each bullet shows old → new clearly. 220 words. The subject line names the count and highlights the good news to drive opens."

  - type: note
    text: "Compare this to a prompt like 'write an email about our policy changes.' The three questions turned a vague idea into a precise specification: what done looks like (no HR follow-ups), what context matters (last comms failed, mixed reactions expected), what the boundaries are (300 words, lead with good news, direct tone)."
```

```quiz
id: three-questions-test
type: multiple-choice
question: "You ask AI to 'help with the quarterly presentation.' It produces something mediocre. Which of the Three Questions did you most likely skip?"
options:
  - "'What does done look like?' -- you did not specify the deliverable, audience, or format"
  - "'What context does it need?' -- you did not provide the Q4 data, slides template, or brand guide"
  - "'What are the boundaries?' -- you did not set constraints on length, tone, or topics to cover"
answer: 0
explanation: "'Help with' is not a defined outcome. Without knowing what 'done' looks like -- 10 slides? Executive summary? Talking points? -- the AI cannot produce the right deliverable. The other questions matter too, but 'done' is the foundation that makes context and boundaries meaningful."
```
