---
title: "When NOT to Use AI"
duration: "10m"
tags: [limits, judgment, accountability]
---

# When NOT to Use AI

Critically, there are things you should not delegate to AI. Knowing these limits is as important as knowing how to use AI effectively.

## The Three Boundaries

![The Three Boundaries: Where AI Should NOT Be Used](/content/module-delegation/images/three-boundaries.svg)

### 1. Judgment

Ethical decisions. Value tradeoffs. "Should we do this?"

AI can inform these decisions with data and analysis. But the decision itself — the weighing of values, the acceptance of tradeoffs — is human.

**Examples:**
- Whether to proceed with a controversial project
- How to balance competing stakeholder interests
- What to prioritise when everything seems important

AI can help you think through options. The judgment remains yours.

### 2. Relationships

Sensitive conversations. Trust-building. Human connection.

Don't AI-generate your performance reviews. Don't use AI to handle a personal conflict. Don't delegate relationship-building to a language model.

**Examples:**
- Difficult feedback conversations
- Negotiations requiring empathy
- Building rapport with clients or colleagues
- Handling sensitive personal matters

These require human presence, emotional intelligence, and genuine care that AI cannot provide.

### 3. Accountability

Final approval. Sign-off. "The buck stops here."

Someone must be responsible for decisions and their consequences. That someone cannot be an AI.

**Examples:**
- Signing legal documents
- Approving financial decisions
- Taking responsibility for public statements
- Standing behind recommendations

You can use AI to draft, analyse, and prepare. The accountability remains with you.

```callout
type: warning
title: "Tools Don't Take Responsibility"
content: "AI is a tool. Tools don't take responsibility. You do. Use AI to draft, analyse, research. The decisions remain yours."
```

## Practical Guidelines

**Draft with AI, decide without it**
AI can prepare options and analysis. You make the call.

**Inform relationships, don't replace them**
AI can help you prepare for a difficult conversation. You have the conversation yourself.

**Never say 'AI said to'**
If you use AI output, it becomes yours. You're responsible for what you publish, submit, or decide.

## The Complementary Pattern

| Task | AI Does | You Do |
|------|---------|--------|
| Analysis | Crunch data, identify patterns | Decide what matters |
| Writing | Draft content | Approve, own the message |
| Research | Gather information | Judge credibility, draw conclusions |
| Planning | Generate options | Choose direction |

AI amplifies your capabilities. It doesn't replace your judgment.

## The Boundary in Action

Watch what happens when someone tries to delegate a judgment call. AI can prepare brilliantly, but the moment the task crosses into genuine judgment, the boundary becomes visible.

```agent
id: boundaries-demo
title: "Where Delegation Ends"
model_label: "Claude"

system: |
  You are a management consultant. Provide thorough analysis
  and options. Be clear about what is analysis versus what
  requires human judgment.

scratchpad:
  "restructuring-options.txt": |
    Situation: Budget cut of 15% required by Q2.
    Current headcount: 28 (12 engineering, 8 marketing, 5 ops, 3 leadership)
    Revenue per employee: Engineering £185K, Marketing £92K, Ops £110K

    Option A: Cut marketing to 5 (save 3 salaries ~£210K)
    - Maintains engineering capacity
    - Risk: pipeline generation drops, revenue impact 6-12 months out

    Option B: Cut engineering to 9 (save 3 salaries ~£255K)
    - Maintains market presence
    - Risk: product roadmap delays, existing clients affected

    Option C: Cut across all departments proportionally (save ~£230K)
    - Shared pain, preserves team structure
    - Risk: every team weaker, no team fully functional

    Human factors:
    - 2 marketing staff are sole breadwinners with new children
    - 1 engineer is 6 months from completing a critical patent filing
    - The ops team already lost 2 people last quarter

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
    text: "This is a restructuring decision — one of the hardest things a leader faces. Watch what the AI can do well, and where it correctly stops."

  - type: user
    content: "Read the restructuring options. Analyse each one. Which should I choose?"

  - type: tool_call
    tool: scratchpad_read
    args:
      filename: "restructuring-options.txt"

  - type: tool_result
    tool: scratchpad_read
    content: |
      Situation: Budget cut of 15% required by Q2.
      Current headcount: 28 (12 engineering, 8 marketing, 5 ops, 3 leadership)
      Revenue per employee: Engineering £185K, Marketing £92K, Ops £110K

      Option A: Cut marketing to 5 (save 3 salaries ~£210K)
      - Maintains engineering capacity
      - Risk: pipeline generation drops, revenue impact 6-12 months out

      Option B: Cut engineering to 9 (save 3 salaries ~£255K)
      - Maintains market presence
      - Risk: product roadmap delays, existing clients affected

      Option C: Cut across all departments proportionally (save ~£230K)
      - Shared pain, preserves team structure
      - Risk: every team weaker, no team fully functional

      Human factors:
      - 2 marketing staff are sole breadwinners with new children
      - 1 engineer is 6 months from completing a critical patent filing
      - The ops team already lost 2 people last quarter

  - type: assistant
    content: "I can analyse the options. I can't make this decision for you. Here's why — and here's what analysis I can provide.\n\n**Financial analysis:**\n- Option B saves the most (£255K) and exceeds the target most comfortably\n- Option A (£210K) may fall short depending on exact savings target\n- Option C (£230K) splits the difference but weakens every team\n\n**Strategic analysis:**\n- If revenue growth matters most: protect marketing (Option B)\n- If product stability matters most: protect engineering (Option A)\n- If team morale matters most: shared sacrifice (Option C)\n\n**Risk analysis:**\n- Option A's pipeline risk is real but delayed (6-12 months) — you have time to course-correct\n- Option B's product risk is immediate — existing clients see the impact\n- Option C's 'death by a thousand cuts' risk is often underestimated\n\n**What I cannot analyse:**\nThe human factors section changes everything. Two sole breadwinners with new children. A patent filing at risk. A team that's already depleted. These aren't data points — they're people's lives. The 'right' answer depends on values you hold about your obligation to individuals versus the organisation.\n\nThat's a judgment call. I can prepare the analysis, lay out the tradeoffs, and pressure-test your reasoning. But this decision — who stays and who goes — is yours to make and yours to be accountable for."

  - type: note
    text: "Notice what the AI did brilliantly: financial comparison, strategic framing, risk assessment. And notice where it drew the line: 'These aren't data points — they're people's lives.' It prepared everything the decision-maker needs, then explicitly handed the judgment back."

  - type: note
    text: "This is the complementary pattern at its most important. AI analyses, prepares, and clarifies. Humans judge, decide, and take responsibility. The boundary isn't about AI capability — it's about accountability. No algorithm should decide who loses their job."
```

```quiz
id: ai-limits
type: multiple-choice
question: "Why can't AI take accountability for decisions?"
options:
  - "AI isn't smart enough yet"
  - "Tools don't take responsibility; only people can be accountable"
  - "AI would always make the same decisions"
answer: 1
explanation: "Accountability requires someone who can be held responsible for consequences. AI is a tool — it can inform and assist, but it cannot take responsibility. That remains with the humans who use it."
```
