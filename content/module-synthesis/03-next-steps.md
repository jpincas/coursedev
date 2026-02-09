---
title: "What To Do Next"
duration: "15m"
tags: [action, next-steps, practice]
---

# What To Do Next

The tools will keep evolving. Here's how to stay effective.

## What Stays Constant

These principles apply regardless of which tools you use:
- Clear outcome specification
- Rich context provision
- Verification mindset
- Preparation discipline

Master the principles. Adapt to the tools.

## The Five Quick Wins

**Start here.** Research shows these tasks deliver the highest return on time invested.

**1. Meeting summarisation**
Record meetings, let AI transcribe, extract action items, and suggest follow-ups. Ten minutes of AI time replaces hours of note-taking and follow-up coordination.

**2. Email drafting and triage**
Draft responses to routine emails. Summarise long threads. Categorise incoming mail. AI handles the mechanics while you handle judgment.

**3. Document summarisation**
Drop in a 50-page report, receive a two-page executive summary. AI reads everything; you read what matters.

**4. Status report generation**
Pull data from project tools, generate weekly status updates automatically. Stop spending Friday afternoons writing what the systems already know.

**5. Research acceleration**
Multi-source synthesis: upload five articles, receive comparative analysis. AI does the reading and organising; you do the thinking.

```callout
type: info
title: "The Data Backs This"
content: "Daily GenAI users save 4+ hours per week and report significant productivity improvements. They also report higher job security and higher salaries compared to infrequent users. Starting matters."
```

## The Ten Mistakes Beginners Make

**Avoid these and you'll skip the frustration phase.**

**1. Being too vague**
"Make this better" tells AI nothing. Specify audience, tone, format, length, purpose.

**2. Overloading single prompts**
Don't ask for a complete 20-page report in one go. Break it into sections: outline, then expand each part.

**3. Skipping role assignment**
"You are a senior UX designer explaining to developers..." sets context. Generic AI gets generic results.

**4. Not iterating**
The first output is a draft. Review it. Give specific feedback. Iterate. This is conversation, not magic.

**5. Ignoring limitations**
AI hallucinates. It doesn't know today's date. It can't browse the web unless connected. Work within reality.

**6. Not providing examples**
Show AI what "good" looks like. Two examples improve output more than 200 words of description.

**7. Sharing sensitive data**
Check your tool's data policy. Understand what gets logged. Default to excluding confidential information.

**8. Using the wrong model**
Fast models for simple tasks. Reasoning models for complex logic. General models for broad work. Match tool to task.

**9. Failing to use meta-prompting**
Ask AI to help you write better prompts. "What information would help you produce a better analysis of this data?"

**10. Not verifying outputs**
Never treat AI as source of truth. Check facts. Verify citations. Test logic. You own the result.

```callout
type: warning
title: "These Mistakes Are Expensive"
content: "As we discussed in the delegation module, the verification overhead is real. That time drops dramatically when you avoid these ten mistakes upfront."
```

## Your Action Plan

**This Week**
Pick ONE of the five quick wins. Do it properly. Experience the full workflow: prepare context, delegate clearly, verify output, use the result.

**This Month**
Build your first complete AI-first workflow for a recurring deliverable. Document what works. Refine what doesn't.

**Ongoing**
Each time something works well, save the prompt. Build your personal toolkit. Track what saves time.

```callout
type: tip
title: "Start Small"
content: "Don't try to transform everything at once. Pick one task. Get it working. Then expand. Sustainable change beats dramatic change."
```

## Common Objections Addressed

**"What about confidential data?"**
Understand your tools' data policies. Many enterprise tools now offer data isolation. Start with non-sensitive tasks.

**"My work is too specialised."**
AI is often better at specialised tasks — if you provide the context. Your expertise combined with AI execution is powerful.

**"I tried it and it was bad."**
Likely a context problem. What did AI have to work with? The investment in setup pays off.

**"My company doesn't allow it."**
Policies are evolving rapidly. Document successful experiments. Build the case for approved tools.

## The Long Game

Six months from now:
- You'll have a library of skills
- Persistent memory will know how you work
- Connections to your tools will be established
- Tasks that used to take hours will take minutes

This doesn't happen overnight. It compounds over time. Start building now.

## Quick Win: Meeting Summarisation

The first quick win in action. Watch how a messy set of meeting notes becomes a structured action document — the kind of task you could start doing tomorrow.

```agent
id: quick-win-demo
title: "Your First Quick Win"
model_label: "Claude"

system: |
  You are a meeting assistant. Transform raw meeting notes into
  clear, structured summaries with action items. Be concise and
  focus on what matters.

scratchpad:
  "meeting-notes-raw.txt": |
    Product sync - 14 Jan
    Sarah, Tom, Dev (dial in late), Priya

    Started with Q4 review - Tom said numbers look good, will send
    final report by Friday. Sarah asked about the mobile numbers
    specifically, Tom said downloads up 40% but retention is flat.
    Dev joined around 10:15.

    Priya presented the new onboarding flow mockups. Everyone liked
    the simplified signup but Tom worried about losing the team size
    selector - says sales uses that data. Priya said she'd look at
    keeping it but making it optional. Sarah wants to see revised
    mockups by Wednesday.

    Dev raised the API issue from last week - still blocking 3 customers.
    Fix is in testing, should ship tomorrow (Wednesday). Sarah asked
    if we need to notify affected customers. Dev said yes, Tom offered
    to draft the email. Sarah approved, wants it sent today.

    Discussed Q1 planning briefly. Sarah wants each team lead to bring
    their top 3 priorities to Friday's session. No more than 3 each.
    She also mentioned the board wants a product vision doc by end of
    Jan - she'll take first pass and share for input next week.

    Meeting ran over by 10 mins. Need to keep these tighter.

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
    text: "These are real-world meeting notes: messy, conversational, with buried action items. This is the kind of task that takes 20 minutes to process manually and that most people put off (or never do). Watch it become structured in seconds."

  - type: user
    content: "Read these meeting notes and create two things: a clean summary for people who weren't there, and a separate action item tracker with owners and due dates."

  - type: tool_call
    tool: scratchpad_read
    args:
      filename: "meeting-notes-raw.txt"

  - type: tool_result
    tool: scratchpad_read
    content: |
      Product sync - 14 Jan
      Sarah, Tom, Dev (dial in late), Priya

      Started with Q4 review - Tom said numbers look good, will send
      final report by Friday. Sarah asked about the mobile numbers
      specifically, Tom said downloads up 40% but retention is flat.
      Dev joined around 10:15.

      Priya presented the new onboarding flow mockups. Everyone liked
      the simplified signup but Tom worried about losing the team size
      selector - says sales uses that data. Priya said she'd look at
      keeping it but making it optional. Sarah wants to see revised
      mockups by Wednesday.

      Dev raised the API issue from last week - still blocking 3 customers.
      Fix is in testing, should ship tomorrow (Wednesday). Dev said yes,
      Tom offered to draft the email. Sarah approved, wants it sent today.

      Discussed Q1 planning briefly. Sarah wants each team lead to bring
      their top 3 priorities to Friday's session. No more than 3 each.
      She also mentioned the board wants a product vision doc by end of
      Jan - she'll take first pass and share for input next week.

      Meeting ran over by 10 mins. Need to keep these tighter.

  - type: tool_call
    tool: scratchpad_write
    args:
      filename: "meeting-summary.md"
      content: |
        # Product Sync — 14 January 2026
        **Attendees:** Sarah, Tom, Dev, Priya

        ## Key Decisions
        - New onboarding flow approved in principle; team size selector to be made optional (not removed)
        - API fix to ship Wednesday; affected customers to be notified today
        - Q1 planning: each lead brings top 3 priorities to Friday session

        ## Discussion Highlights

        **Q4 Review:** Mobile downloads up 40%, but retention flat — needs investigation. Tom sending final Q4 report by Friday.

        **Onboarding Redesign:** Simplified signup well-received. One concern: sales relies on team size data. Priya to revise mockups keeping the field as optional.

        **API Issue:** Blocking 3 customers since last week. Fix in testing, shipping Wednesday. Customer notification approved.

        **Q1 Planning:** Board wants product vision document by end of January. Sarah taking first pass, team input next week.

        ---

        ## Action Items

        | # | Action | Owner | Due |
        |---|--------|-------|-----|
        | 1 | Send final Q4 report | Tom | Fri 17 Jan |
        | 2 | Revised onboarding mockups (optional team size selector) | Priya | Wed 15 Jan |
        | 3 | Ship API fix | Dev | Wed 15 Jan |
        | 4 | Draft and send customer notification about API issue | Tom | TODAY (14 Jan) |
        | 5 | Prepare top 3 Q1 priorities | All leads | Fri 17 Jan |
        | 6 | First draft of product vision doc | Sarah | w/c 20 Jan |

  - type: tool_result
    tool: scratchpad_write
    content: "Written to meeting-summary.md (30 lines)"

  - type: assistant
    content: "Created meeting-summary.md with both the clean summary and action tracker. Six action items extracted, each with an owner and due date. The urgent one (customer notification) is flagged as TODAY."

  - type: note
    text: "Messy notes became a structured summary in seconds. Six action items with owners and dates — extracted from conversational text where they were buried in narrative. This is Quick Win #1: meeting summarisation. It works with any meeting tool that produces a transcript."

  - type: note
    text: "Now imagine doing this after every meeting. Automatically. No more lost action items. No more 'I thought you were handling that.' This is what daily AI use looks like: not dramatic transformation, but consistent time savings on tasks you already do."
```

```quiz
id: next-steps-quick-wins
type: multiple-choice
question: "According to research, daily GenAI users save how much time per week compared to non-users?"
options:
  - "30 minutes or less"
  - "1-2 hours"
  - "4+ hours"
  - "10+ hours"
answer: 2
explanation: "Daily GenAI users save 4+ hours per week on average, and they report 92% productivity improvement, higher job security, and higher salaries. The key is consistent use focused on high-ROI tasks like the five quick wins."
```

```quiz
id: next-steps-mistakes
type: multiple-choice
question: "What's the most common mistake that causes poor AI output quality?"
options:
  - "Using the wrong AI model"
  - "Not providing enough examples"
  - "Being too vague in instructions"
  - "Not iterating on the first output"
answer: 2
explanation: "Being too vague is the #1 mistake. Modern models follow instructions literally — vague prompts get vague results. Specify audience, tone, format, length, and purpose. The 'Colleague Test': if a colleague would be confused by your instruction, AI will be too."
```
