---
title: "Meta-Prompting"
duration: "10m"
tags: [meta-prompting, iteration, techniques]
---

# Meta-Prompting

Here's a powerful technique: instead of writing the perfect prompt yourself, ask AI to help you write it.

## The Basic Pattern

```
I want to [goal]. Help me write a detailed prompt
that will get the best result.
Ask me clarifying questions first.
```

The AI will ask what it needs to know. Your answers become the prompt.

This is especially useful when you're not sure what you need.

## Meta-Prompting in Action

**You:** "I need to analyse customer feedback. Help me write a prompt for that."

**AI:** "Let me ask a few questions:
- What format is the feedback in?
- What specific insights are you looking for?
- How should the analysis be structured?
- Who will read the results?"

**You:** "It's in a CSV. Looking for common complaints. Need a summary for the product team."

**AI:** "Here's your prompt: 'Analyse this customer feedback CSV. Identify the top 5 complaint categories with frequency counts. For each, provide 2-3 representative quotes. Format as a brief report for the product team, max 1 page.'"

```callout
type: tip
title: "Why This Works"
content: "The AI's questions reveal what information is needed. Each answer adds specificity. The final prompt is often better than what you would have written from scratch."
```

## When to Use Meta-Prompting

**You're unsure about scope**
Let the AI's questions help you figure out what you actually need.

**Complex or multi-faceted tasks**
The questioning process surfaces aspects you might not have considered.

**Unfamiliar domains**
AI can ask expert-level questions about topics where you're not the expert.

**Reusable prompts**
If you'll do similar tasks repeatedly, invest in getting a great prompt.

## The Key Insight

Don't describe the process. Describe the outcome.

**Bad:**
"First analyse the data, then identify trends, then summarise them, then format as bullet points, then..."

**Good:**
"Produce a trend analysis with the top 5 insights as bullet points."

AI is capable of figuring out the process. That's its job. Your job is defining what done looks like.

This is the shift from prompting to delegation.

## Interactive Speccing

Meta-prompting taken to its logical conclusion for complex deliverables is **interactive speccing** -- using AI to help you spec out a project or document through iterative conversation before executing.

**The pattern:**

1. Describe what you want at a high level: "I need a training curriculum for new managers"
2. AI asks clarifying questions about scope, audience, constraints
3. You answer, adding detail and refining
4. AI proposes a structured spec or outline based on your answers
5. You review and adjust the spec: "Move section 3 before section 2, add a module on difficult conversations"
6. Once the spec is agreed, you ask AI to execute against it

**Why this works:** For complex deliverables, jumping straight to execution produces mediocre results because requirements were underspecified. Interactive speccing forces you to think through the requirements before any work happens. AI's questions surface aspects you would not have considered.

**Example flow:**
```
You: I need a 30-page client proposal for a digital transformation project.
AI: What industry? What's the client's current state? Budget range? Timeline?
You: Financial services, legacy systems, ~$2M, 18 months.
AI: Here's a proposed structure: [10 sections with descriptions]
You: Good, but add a section on regulatory compliance and remove the generic 'about us.'
AI: Updated structure: [revised]. Shall I draft the full proposal against this spec?
You: Yes, go ahead.
```

The spec becomes a contract between you and the AI. The result is dramatically better than a single monolithic prompt.

```callout
type: tip
title: "When to Spec vs When to Just Ask"
content: "Use interactive speccing for anything that would take more than a page: proposals, reports, curricula, project plans. For quick tasks (emails, summaries, short analyses), just ask directly -- speccing would be overkill."
```

## Interactive Speccing in Action

Watch someone use AI to spec out a complex project through conversation. The user has a vague goal and minimal context — the AI's questions draw out the details.

```agent
id: interactive-speccing-demo
title: "Interactive Speccing: Company Offsite"
model_label: "Claude"

system: |
  You are an event planning consultant. When helping clients plan
  company offsites, ask clarifying questions to understand their goals,
  constraints, and past experiences. Create detailed, structured specs
  that cover all aspects of the event.

scratchpad:
  "context/company-info.md": |
    # Vortex Technologies

    Tech company, 50 employees, fully remote-first since 2020.
    Headquarters: London (registered office only, mostly empty)
    Team distribution: UK (30), EU (15), US East Coast (5)

    Engineering-heavy culture. Mix of junior and senior staff.
    Average age: 32. Several families with young children.

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
    text: "The user has a vague goal: plan a company offsite. Instead of trying to specify everything upfront, watch them ask the AI to help build a proper spec through conversation."

  - type: user
    content: "I need to plan a company offsite. Rather than me trying to think of everything, can you interview me to build a proper spec? There's some basic company info in the context folder."

  - type: tool_call
    tool: scratchpad_read
    args:
      filename: "context/company-info.md"

  - type: tool_result
    tool: scratchpad_read
    content: |
      # Vortex Technologies

      Tech company, 50 employees, fully remote-first since 2020.
      Headquarters: London (registered office only, mostly empty)
      Team distribution: UK (30), EU (15), US East Coast (5)

      Engineering-heavy culture. Mix of junior and senior staff.
      Average age: 32. Several families with young children.

  - type: assistant
    content: "Good approach — I'll ask questions to build a complete spec. Let me start with the fundamentals:\n\n**Goals & Purpose:**\n1. What are the primary goals for this offsite? (Team building? Strategic planning? Celebrating milestones? Skills development?)\n2. Is there anything specific you want to achieve that previous offsites haven't?\n\n**Budget & Constraints:**\n3. What's your budget range per person?\n4. Any hard constraints on dates or location?\n5. How many days are you thinking?\n\n**Past Experience:**\n6. Have you done an offsite before? What worked well? What didn't?"

  - type: user
    content: "Main goal is team cohesion — being remote-first, most people have never met in person. We also want to do some strategic planning for H2. Budget is around £1,200 per person. Dates: ideally late April or early May, before summer holidays kick in. Thinking 2-3 days. We tried a one-day thing last year in London but it felt rushed and the US team couldn't justify flying over for just one day."

  - type: assistant
    content: "That's helpful context. The one-day London event being too short is a key lesson. Follow-up questions:\n\n**Location & Logistics:**\n7. Given your team distribution, where makes sense? (Central European location for easier access from all three regions?)\n8. All 50 people, or just certain teams/seniority levels?\n9. Will you cover travel costs, or is that part of the £1,200/person budget?\n\n**Activities & Structure:**\n10. What balance do you want between work sessions and social activities?\n11. Any activities that definitely should or shouldn't be included?\n12. Do you want external facilitators for the strategic planning portions?"

  - type: user
    content: "Probably somewhere in Europe that's easy to reach — maybe Barcelona, Amsterdam, or Lisbon? All 50 people should come. The £1,200 is for accommodation, food, and activities — travel is separate. Balance: I'd say 40% work sessions, 60% social and team building. No trust falls or forced corporate fun. Maybe something optional and adventurous for people who want it. Not sure about facilitators — what do you think?"

  - type: assistant
    content: "For strategic planning with 50 people, an external facilitator often helps — keeps you out of meeting-runner mode and participating instead. I'd recommend budgeting for one.\n\nA few more details to nail down the spec:\n\n**Accommodation & Meals:**\n13. Hotel or something more informal like a villa/retreat center?\n14. Any dietary requirements or preferences to plan for?\n\n**Families & Inclusion:**\n15. Your notes mention several people have young children — allow partners/families, or employees only?\n\n**Outputs:**\n16. Do you want a tangible outcome from the strategic planning sessions? (Documented H2 roadmap? OKRs? Or just alignment?)"

  - type: user
    content: "Hotel makes sense for 50 people — easier logistics. We've got vegetarians and a couple of people with allergies, so venue needs flexible catering. Employees only for this one — we want focused work time. And yes, we should come away with a documented H2 roadmap and agreed OKRs. That's the work product."

  - type: assistant
    content: "Perfect — I've got what I need. Let me draft a complete offsite spec based on everything you've told me."

  - type: tool_call
    tool: scratchpad_write
    args:
      filename: "offsite-spec.md"
      content: |
        # Vortex Technologies Spring 2026 Offsite
        ## Event Specification

        ### Purpose & Goals
        **Primary objective:** Build team cohesion across a fully remote team (most people have never met in person)

        **Secondary objective:** Develop H2 2026 strategic roadmap and OKRs

        **Key lesson from 2025:** One-day London event was too short and didn't justify US team travel. This event must deliver sufficient value for international travel.

        ---

        ### Logistics
        **Attendees:** All 50 employees (UK: 30, EU: 15, US East: 5)

        **Dates:** Late April or early May 2026 (pre-summer holidays)
        - Suggested: April 28-30 (Tue-Thu) or May 5-7

        **Duration:** 3 days, 2 nights
        - Day 1: Arrivals PM, welcome dinner
        - Day 2: Full day (strategic sessions + team activities)
        - Day 3: Half day, departures PM

        **Location:** European city with good international flight connections
        - Shortlist: Barcelona, Amsterdam, Lisbon
        - Requirements: Direct flights from London, major EU cities, US East Coast

        ---

        ### Budget
        **Per person:** £1,200 (accommodation, meals, activities, facilitator)

        **Total budget:** £60,000

        **Travel:** Separate budget (not included in £1,200)

        ---

        ### Accommodation
        **Type:** Hotel (4-star, centrally located)

        **Requirements:**
        - 50 single rooms (2 nights)
        - Conference/meeting space for 50 people
        - Breakout rooms for smaller group sessions
        - Flexible catering (vegetarian options, allergy accommodations)
        - AV equipment (projector, screens, good wifi)

        ---

        ### Program Structure
        **Balance:** 40% strategic work sessions, 60% social/team building

        **Day 1 (Arrivals):**
        - 6:00 PM: Welcome reception
        - 7:30 PM: Dinner (informal, getting-to-know-you)

        **Day 2 (Full Day):**
        - 9:00 AM - 12:30 PM: Strategic planning session 1 (H2 priorities)
          - External facilitator-led
          - Breakout groups by function
        - 12:30 PM: Lunch
        - 2:00 PM - 4:00 PM: Strategic planning session 2 (OKR setting)
        - 4:30 PM - 7:00 PM: Team activity (see options below)
        - 8:00 PM: Group dinner

        **Day 3 (Half Day):**
        - 9:00 AM - 11:30 AM: Strategic planning session 3 (roadmap finalization)
        - 11:30 AM: Closing & next steps
        - 12:30 PM: Lunch, then departures

        ---

        ### Activities
        **Guiding principle:** No forced corporate team-building. Offer genuine experiences.

        **Day 2 afternoon options:**
        - Option A (Adventurous): Coasteering or via ferrata experience (optional, 20-person max)
        - Option B (Cultural): Guided city architecture/food tour
        - Option C (Relaxed): Cooking class or wine tasting

        All options should allow natural conversation and mixing.

        ---

        ### Strategic Planning Output
        **Deliverables:**
        1. Documented H2 2026 roadmap (key initiatives, owners, timelines)
        2. Company OKRs for H2 (objectives and key results)
        3. Cross-functional dependencies map

        **Facilitator brief:**
        - 50-person tech company, remote-first
        - Mix of junior and senior IC + management
        - Need structured process but not rigid corporate approach
        - Should produce actionable documents, not just discussion

        ---

        ### Next Steps
        1. Choose between Barcelona, Amsterdam, Lisbon (research flight access + venue options)
        2. Get 2-3 hotel quotes (meeting space + rooms + catering)
        3. Book external facilitator (ask for recommendations from network)
        4. Finalize exact dates (check team calendars for conflicts)
        5. Research activity providers for Day 2 options

        ---

        ### Success Criteria
        - 90%+ attendance (travel logistics work for all regions)
        - Team reports feeling more connected post-event
        - Clear H2 roadmap agreed and documented
        - Positive feedback on work/social balance
        - US team feels the travel was worthwhile

  - type: tool_result
    tool: scratchpad_write
    content: "Written to offsite-spec.md (104 lines)"

  - type: assistant
    content: "Done — offsite-spec.md is ready. It's a complete specification covering purpose, logistics, budget breakdown, accommodation requirements, hour-by-hour program structure, activity options, strategic planning outputs, and success criteria. You can now share this with venues, facilitators, and your team for input."

  - type: note
    text: "Look at what just happened. The user started with 'I need to plan a company offsite.' Through 16 questions across four rounds, the AI drew out details the user hadn't thought to specify upfront: facilitator needs, dietary requirements, family inclusion policy, tangible outputs from planning sessions."

  - type: note
    text: "The spec is now comprehensive enough to brief a venue, book a facilitator, and get team buy-in. Compare this to trying to write all of that in a single upfront prompt. The interactive approach produced a better spec faster, because the AI's questions revealed what needed to be specified."
```

## Meta-Prompting Live

Instead of just describing meta-prompting, watch it happen. The user doesn't know exactly what they need — so they ask the AI to help figure it out.

```agent
id: meta-prompting-demo
title: "AI Helps Write the Prompt"
model_label: "Claude"

system: |
  You are a prompt engineering assistant. When asked to help write
  a prompt, ask clarifying questions first to understand exactly
  what's needed. Then produce a complete, well-structured prompt.

tools: []

visibility:
  system_prompt: hidden
  tool_calls: hidden
  full_context: hidden

sidebar:
  width: "45%"
  start_open: true

script:
  - type: note
    text: "The user has a vague goal: they want to do something with customer feedback. Instead of writing a mediocre prompt, they ask the AI to help write a good one."

  - type: user
    content: "I want to do something useful with our customer feedback data. Help me write a prompt that will get the best analysis. Ask me questions first."

  - type: assistant
    content: "Good approach — let me ask a few questions to build the right prompt:\n\n1. **What format is the feedback in?** (CSV, survey responses, free-text reviews, support tickets?)\n2. **How much data?** (Dozens, hundreds, thousands of entries?)\n3. **What decisions will this inform?** (Product roadmap? Marketing? Support staffing?)\n4. **Who will read the analysis?** (Exec team? Product managers? The whole company?)\n5. **What have you tried before?** (Any previous analysis attempts, or starting fresh?)\n6. **Any specific concerns?** (Known issues you want validated? Hypotheses to test?)"

  - type: user
    content: "It's a CSV export from our NPS survey — about 800 responses from Q4. Mix of scores and free-text comments. I need this for the product team's Q1 planning session next week. The VP of Product will present it. We think mobile experience is a problem but we're not sure."

  - type: assistant
    content: "Here's your prompt:\n\n---\n\n**Analyse this NPS survey data (800 Q4 responses, CSV attached).**\n\n**Objectives:**\n1. Identify the top 5 themes in the free-text comments, ranked by frequency\n2. For each theme, break down by NPS score band (Detractors 0-6, Passives 7-8, Promoters 9-10)\n3. Specifically investigate mobile experience mentions — frequency, sentiment, and correlation with low scores\n4. Flag any emerging themes that appear in 3%+ of responses\n\n**Output format:**\n- Executive summary (3-4 bullet points, suitable for VP presentation)\n- Detailed theme breakdown as a table\n- Mobile experience deep-dive section\n- 5 recommended actions for Q1 product planning, ranked by potential NPS impact\n\n**Constraints:**\n- Use direct quotes from responses as evidence (3-5 per theme)\n- Distinguish between correlation and causation\n- Flag if sample size for any segment is too small for reliable conclusions\n- Professional tone suitable for leadership presentation\n\n---\n\nThis prompt gives the AI your data context, a clear analysis structure, a specific hypothesis to test (mobile), and an output format matched to your audience. Want me to adjust anything?"

  - type: note
    text: "Look at what happened. The user started with 'I want to do something useful with our customer feedback.' After six questions and one set of answers, the AI produced a prompt with clear objectives, structured output format, a specific hypothesis to test, and constraints. That prompt would have taken 15 minutes to write from scratch."

  - type: note
    text: "This is meta-prompting: using AI to write the prompt that AI will then execute. The questions surface what you know but haven't articulated yet."
```

```quiz
id: meta-prompting-purpose
type: multiple-choice
question: "When is meta-prompting most valuable?"
options:
  - "When you know exactly what you want but cannot articulate it concisely"
  - "When you have a vague goal and are unsure what specificity the AI needs to produce good output"
  - "When you need to produce multiple variations of the same document"
answer: 1
explanation: "Meta-prompting is most valuable when you are unsure about scope or requirements. The AI's clarifying questions surface what information and specificity is needed, turning a vague goal into a precise specification. If you already know exactly what you want, just ask directly."
```

```quiz
id: interactive-speccing-when
type: multiple-choice
question: "You need to create a 20-page annual report for your company. What approach will produce the best result?"
options:
  - "Write a detailed single prompt specifying every section, format, and data point"
  - "Ask AI to draft it, then iterate through 5-6 rounds of feedback"
  - "Spec it out interactively first -- agree on structure and requirements, then execute against the agreed spec"
answer: 2
explanation: "For complex deliverables, interactive speccing produces better results than either a monolithic prompt (which inevitably misses requirements) or pure iteration (which lacks a clear target). The spec becomes a contract that ensures both you and the AI are aligned on what 'done' looks like before any work begins."
```
