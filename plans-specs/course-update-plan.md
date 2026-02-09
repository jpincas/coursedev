# Course Update Plan: AI Training From Novice to Expert

**Date:** 9 February 2026
**Status:** Planning document — master reference for all content work
**Based on:** Gap analysis between `ai-training-report.md` and current course content

---

## Executive Summary

The current course is a strong foundation with excellent pedagogy. The learning arc (LLMs -> Context -> Prompting -> Files -> Delegation -> Advanced -> Synthesis) is sound. The writing is sharp, opinionated, and practical. The agent demos are a genuine differentiator.

However, the report reveals several critical gaps that must be addressed for this to be the premier product in the market:

1. **Context engineering is underserved.** The report calls it "the single most important conceptual shift of 2025" and the course mentions it briefly in the prompting module's evolution page but never teaches it as a discipline distinct from prompting.
2. **The agentic paradigm needs deeper treatment.** Cowork, Claude Code, ChatGPT Agent Mode, and Gemini Agent are mentioned in the opening but never taught in depth. Learners need to understand what agentic AI is, how it differs from chat, and how to work with it.
3. **Risk, verification, and responsible use are scattered and insufficient.** The report dedicates significant attention to hallucination rates, the METR study showing 19% slower performance, skill atrophy, privacy risks, and legal issues. The course has verification in delegation and hallucinations in LLMs but lacks a cohesive treatment of risks.
4. **Data analysis and document creation workflows are missing.** The report covers democratised data analysis (Part 6) and document creation techniques (Part 5). The course has nothing on either.
5. **The "50x Reframe" and other mental models from the report are absent.** These are powerful teaching tools that belong in the course.
6. **Multishot prompting, extended thinking, XML structure, and prompt chaining are not taught.** The report's Part 3 covers specific techniques that the current prompting module skips entirely.
7. **Token economics and cost awareness are missing.** The report discusses the 100x token consumption of agentic workflows and cost optimisation. A premium audience needs this.
8. **The competitive landscape and model selection are absent.** Learners need to understand Claude vs GPT vs Gemini — not as a product comparison, but as a tool selection framework.
9. **The context window data is outdated.** The course says "Current (2025): ~200K tokens" when Claude Opus 4.6 supports 1 million tokens.

---

## Gap Analysis: Report vs Course

### Covered Well

| Report Topic | Course Location | Assessment |
|---|---|---|
| Next-token prediction | module-llms/01 | Solid |
| Training process (pre-training, fine-tuning, RLHF) | module-llms/02 | Good |
| Hallucinations as inherent | module-llms/03 | Good |
| Context window concept | module-context/02 | Good but data outdated |
| Context hierarchy | module-context/03 | Good |
| Context decay and fresh starts | module-context/04 | Good |
| Outcomes over process | module-prompting/01, module-delegation/02 | Excellent |
| Meta-prompting | module-prompting/04 | Good |
| File-based workflows | module-files (entire module) | Very good |
| Delegation mindset | module-delegation (entire module) | Excellent |
| Quality verification | module-delegation/04 | Good |
| When NOT to use AI | module-delegation/05 | Good |
| Skills/reusable procedures | module-advanced/02 | Good |
| Subagents | module-advanced/03 | Good but too abstract |
| Persistent memory (CLAUDE.md) | module-advanced/04 | Good but too brief |
| MCP | module-advanced/06 | Surface-level |
| Code generation for non-coders | module-advanced/05 | Good |

### Critical Gaps (Not Covered)

| Report Topic | Report Section | Priority | Notes |
|---|---|---|---|
| Context engineering as a discipline | Part 3 | **P0** | The report calls this "the single most important conceptual shift of 2025". Currently just a bullet in the prompting evolution page. |
| Agentic AI paradigm (Cowork, Agent Mode) | Part 2 | **P0** | Mentioned in opening but never explained or taught. |
| Specific prompting techniques (multishot, XML, thinking, chaining) | Part 3 | **P0** | The current prompting module teaches frameworks but skips the specific high-value techniques. |
| Risk landscape (METR study, skill atrophy, privacy, legal) | Part 9 | **P1** | Currently scattered thin. Needs cohesive treatment. |
| Data analysis workflows | Part 6 | **P1** | Not covered at all. Report covers upload-and-analyse, AI-assisted data cleaning, verification requirements. |
| Document creation techniques (sectional drafting, style transfer) | Part 5 | **P1** | Not covered at all. Report covers the optimal human-AI writing workflow. |
| The "50x Reframe" mental model | Part 1 | **P1** | Powerful concept not in the course. |
| Token economics and cost optimisation | Part 8 | **P2** | Not mentioned. Premium audience will care. |
| Competitive landscape / model selection | Part 2 | **P2** | Course is Claude-centric. Needs broadening. |
| Vibe coding / citizen developer | Part 7 | **P2** | Code generation page exists but doesn't cover the broader citizen developer movement. |
| AI-first workflow redesign (McKinsey/PwC findings) | Part 10 | **P2** | The synthesis module touches this but without the research backing. |
| Quick wins and highest-ROI starting points | Part 10 | **P2** | The "next steps" page exists but lacks the specific research-backed recommendations. |
| Multi-model arbitrage | Part 8 | **P3** | Power-user technique worth including. |
| Prompt libraries as shared assets | Part 10 | **P3** | Not covered. |

### Outdated Content

| Item | Current State | Should Be | Location |
|---|---|---|---|
| Context window sizes | "Current (2025): ~200K tokens" | Claude Opus 4.6 = 1M tokens; mention Infinite Chats | module-context/02 |
| Opening timeline | "February 2026... 8 weeks ago" | Needs to reference specific products: Cowork (Jan 12), Codex, Gemini Agent | module-opening/01 |
| Prompting evolution | Lists 2022, 2024, 2025 | Should include 2026: context engineering as formalized discipline | module-prompting/01 |
| MCP coverage | "Evolving Standard" | 17,000+ servers, 97M monthly SDK downloads, Linux Foundation governance | module-advanced/06 |
| Code generation | Generic examples | Should reference vibe coding, specific tools (Cursor $500M ARR), citizen developer stats | module-advanced/05 |

---

## Module-by-Module Task Lists

### Module 1: module-opening (30m)

**Current state:** 3 pages. Sets up the "chatbot to worker" framing. Has a good agent demo. Solid opening.

**Tasks:**

1. **Page 01 (welcome) — Update the timeline and specific products** (HIGH)
   - Replace vague "8 weeks ago" with specific dates: Claude Cowork (Jan 12, 2026), ChatGPT Agent Mode (July 2025), Gemini Agent
   - Add Anthropic's "Chat -> Code -> Cowork" trajectory as the narrative arc (report Part 2)
   - Add the $285B market selloff detail — this is a powerful "this is real" moment
   - Update the "AI Timeline" graphic to include: Claude Code (Feb 2025, 5.5x revenue growth), Operator (Jan 2025), Cowork (Jan 2026), Cowork Plugins (Jan 30 2026)
   - Add mention of "41% of all code is now AI-generated" — a concrete stat that lands

2. **Page 01 — Add the "50x Reframe" mental model** (HIGH)
   - This belongs in the opening as a hook. It is perhaps the most powerful motivational concept in the report
   - Teach Azeem Azhar's framing: "What would I do if I had 50 people on this?"
   - Give the concrete example: instead of finding 5 podcast guests, evaluate the top 1,000
   - This reframe prevents anchoring to existing constraints and opens minds

3. **Page 02 (what you'll learn) — Add a "Digital Intern to Digital Executive" framing** (MEDIUM)
   - The report's evolution model (intern -> mid-tenure employee -> executive) is a powerful way to set expectations
   - Reference METR data: task length AI can handle doubles every 4 months, now ~2 hours autonomous
   - This gives learners a concrete sense of what is possible now

4. **Page 03 (demo) — No major changes needed** (LOW)
   - The Cowork-style demo is excellent and perfectly illustrates the opening's thesis
   - Minor: could add a note about how Cowork Plugins extend this to legal, financial, marketing workflows

5. **Consider adding a quiz to page 01** (MEDIUM)
   - Currently page 01 has no quiz. Add one testing the "50x Reframe" concept or the chat-to-worker distinction

### Module 2: module-llms (60m)

**Current state:** 4 pages. Covers prediction engines, training, hallucinations, takeaways. Solid foundation.

**Tasks:**

1. **Page 01 (prediction engines) — Add "what AI does well vs poorly" from the report** (HIGH)
   - The report has a crisp framing: "AI excels at code generation (41% of code), pattern recognition, synthesis, first drafts. It struggles with hallucination drift on long chains, cultural nuance, sensitivity to phrasing, novel common-sense reasoning."
   - The Carnegie Endowment line "the hard problems are easy and the easy problems are hard" is memorable and should be included
   - This belongs on this page because it flows directly from understanding the prediction mechanism

2. **Page 02 (training process) — Update with current model landscape** (MEDIUM)
   - Add a brief mention of the model ecosystem: Claude 4.x / Opus 4.6, GPT-5.x, Gemini 3 Pro
   - Update knowledge cutoff framing: models now have web search tools that extend beyond training data
   - Add mention of "extended thinking" as a training/capability development: some models now have an internal reasoning phase before responding

3. **Page 03 (hallucinations) — Add specific hallucination data from the report** (HIGH)
   - Add concrete rates: Gemini-2.0-Flash at 0.7%, while OpenAI's o3 hallucinates 33% on PersonQA
   - Add Stanford finding: LLMs hallucinate 58-82% of legal queries
   - Add the stat: "47% of enterprise AI users made at least one major business decision based on hallucinated content"
   - Add the "knowledge workers spend 4.3 hours per week fact-checking AI" stat
   - Add the OpenAI insight on why hallucinations persist: "training rewards guessing over acknowledging uncertainty"
   - Add practical countermeasures from the report: multi-model cross-validation, asking AI to verify its own reasoning (reduces errors ~17%), web search integration (GPT-4o achieved 90% accuracy with search)
   - This page currently explains WHY hallucinations happen but not HOW BAD they are or WHAT TO DO about them specifically

4. **Page 04 (takeaways) — Update to reflect new content** (LOW)
   - Adjust to summarise new additions from pages above

5. **Add a new page: "What AI Is Good At (and What It Is Not)" between pages 01 and 02** (MEDIUM)
   - Or incorporate into page 01 as a new section
   - Cover: the report's "what AI does well vs poorly" framework
   - Include the Computer Use capability mention: Claude success rate 60%+ on OSWorld, high 80s on standard office tasks
   - This gives learners practical calibration of current capabilities before diving into the mechanism

### Module 3: module-context (60m)

**Current state:** 5 pages. Strong conceptual foundation. The "everything is context" framing is excellent.

**Tasks:**

1. **Page 02 (context window) — CRITICAL UPDATE: Fix outdated data** (P0)
   - The table stops at "Current (2025): ~200K tokens". This is now wrong.
   - Update to include: Claude Opus 4.6 = 1M tokens (~750,000 words), Gemini = 2M tokens
   - Add mention of Anthropic's "Infinite Chats" feature (server-side summarization for indefinite conversations)
   - Add the "lost in the middle" effect — the report emphasizes that models struggle with information buried deep in large contexts, which is a critical practical concept
   - Update the "you can now fit" list to reflect the much larger windows

2. **Page 03 (context hierarchy) — Add persistent instructions layer** (HIGH)
   - The current hierarchy is: system prompts > user instructions > provided context > conversation history
   - Need to add: CLAUDE.md / persistent instructions sit between system prompts and user instructions
   - Reference the report's finding: "persistent instructions override per-conversation prompts, so they're the highest-leverage place to invest in prompt quality"
   - This connects to the advanced module's persistent memory page but should be introduced here conceptually

3. **Page 04 (context decay) — Add /clear, /compact, /rewind as concrete tools** (MEDIUM)
   - The page teaches fresh starts but doesn't mention the specific tools available
   - Add Claude Code's conversation management: /clear (fresh start, still loads CLAUDE.md), /compact (summarize and continue with custom focus), /rewind (selectively roll back)
   - Add the "conversation as refinement" pattern from the report: start with clear initial request, review output critically, give targeted feedback referencing specific parts, iterate with precision
   - Add the common mistake: "saying 'make this better' without specifying what 'better' means"

4. **Page 04 — Add "Scope each conversation to one project or feature" principle** (MEDIUM)
   - The report emphasizes: "Use external state files (progress notes, test results, git logs) rather than relying purely on conversation memory"
   - This is the bridge between context management and file-based workflows

5. **Page 01 (everything is context) — Elevate "context engineering" as a term** (HIGH)
   - The report positions context engineering as "the single most important conceptual shift of 2025"
   - This page should explicitly introduce the term and define it: "designing the entire information environment surrounding the model: memory, retrieved data, tools, state, metadata, and structured inputs"
   - Include Karpathy's metaphor: "The LLM is like the CPU, and its context window is like RAM"
   - This frames the entire module as teaching context engineering, not just "understanding context"

### Module 4: module-prompting (60m)

**Current state:** 5 pages. Covers evolution, basic framework, three questions, meta-prompting, takeaways. Good structure but missing the specific techniques that the report identifies as highest-value.

**Tasks:**

1. **Page 01 (evolution) — Update timeline to include 2026** (MEDIUM)
   - Currently stops at "2025: Delegation"
   - Add: "2026: Agentic orchestration — designing multi-step workflows where AI plans and executes autonomously"
   - Reference Gartner formally defining context engineering
   - Reference the adoption of the framing by every major AI company

2. **Add new page: "The Techniques That Work" between pages 02 and 03** (P0 — HIGH)
   - This is the biggest gap in the prompting module. The report lists specific techniques ordered by effectiveness (from Anthropic's own documentation). The current course has "Role, Task, Context, Format, Constraints" but misses:
   - **Multishot prompting (examples):** Providing 2-5 examples of desired input-output pairs. The report says this "dramatically improves consistency and quality." Show how to wrap examples in XML tags for Claude.
   - **Extended thinking:** Explain what it is (internal reasoning phase before responding), when to use it (math, logic, debugging, multi-step analysis), and the counterintuitive Anthropic guidance: "Claude often performs better with instructions to just think deeply rather than step-by-step prescriptive guidance."
   - **XML structure:** Claude models are specifically trained to parse XML tags. Show `<instructions>`, `<context>`, `<examples>`, `<output_format>`, `<constraints>`. This is Anthropic's signature technique.
   - **Prompt chaining:** Breaking multi-step tasks into sequential subtask prompts where each output feeds the next. "Research first, then organize, then draft, then review."
   - Each technique should have a before/after example
   - Include an agent demo showing multishot prompting or XML structure in action

3. **Add new page: "The Iteration Pattern" between meta-prompting and takeaways** (HIGH)
   - The report emphasizes the "conversation as refinement" pattern
   - Teach: start with a clear initial request with constraints, review output critically identifying specific issues, give targeted feedback referencing specific parts, iterate with precision
   - The most common beginner mistake: "make this better" without specificity
   - Include the "Colleague Test" from the report: show your prompt to a colleague — if they'd be confused, the AI will be too
   - Consider an agent demo showing iterative refinement

4. **Page 02 (basic framework) — Strengthen with "be specific and direct" emphasis** (MEDIUM)
   - The report says this "remains the single highest-leverage technique" even with modern models
   - Add: "Modern models follow instructions very literally — vague prompts get vague results"
   - The page already teaches this implicitly but should be more emphatic

5. **Page 05 (takeaways) — Update to incorporate new techniques** (LOW)
   - Adjust checklist and summary to include multishot, extended thinking, XML structure, chaining, iteration

### Module 5: module-files (60m)

**Current state:** 6 pages. Good coverage of the paradigm shift, input/output files, and grounding. This module is solid.

**Tasks:**

1. **Page 03 (input files) — Add file preparation best practices from the report** (HIGH)
   - The report has specific guidance: "Use machine-readable formats -- Markdown is preferred over complex formatting. Break large documents into logical segments under 100 pages per file. Remove excessive formatting. For Claude Projects, Markdown/Word/PDF use similar space but HTML uses twice as much."
   - Add: "Create focused projects by department or function rather than one massive knowledge base"
   - Add: "Maintain clear file naming and establish regular review cycles"

2. **Page 03 — Update file types list with current capabilities** (MEDIUM)
   - The report specifies: ChatGPT supports up to 10 files per conversation at 512MB per file
   - Claude's code execution generates downloadable spreadsheets, CSVs, reports, interactive visualizations
   - These specifics matter for a premium course

3. **Page 05 (grounding) — Add the "style card" and "EchoWriting" technique** (HIGH)
   - From the report Part 5: Create a "style card" — a reusable prompt encoding preferences for tone, vocabulary, sentence structure, audience
   - EchoWriting: feed AI 15-20 samples of your writing, have it analyze your style patterns, then create a persistent style prompt
   - This is a concrete, actionable technique that elevates the grounding concept

4. **Page 05 — Add multi-document synthesis as a key use case** (MEDIUM)
   - The report calls this "one of AI's highest-value use cases"
   - Upload multiple documents, ask AI to identify common themes, contradictions, gaps
   - Mention Google's NotebookLM as a tool for this
   - Add the limitation: "AI is better at organizing and summarizing than at genuine intellectual synthesis. Always verify it hasn't fabricated connections."

5. **Page 04 (output files) — Add the "sectional drafting method"** (HIGH)
   - From the report: "Never ask for an entire long document in one prompt. Instead: create a detailed outline first, expand each section individually with specific constraints, then merge and smooth transitions."
   - This is a practical technique that dramatically improves output quality for longer documents

6. **Page 02 (why files matter) — No major changes needed** (LOW)
   - Content is solid and relevant

### Module 6: module-delegation (60m)

**Current state:** 6 pages. Excellent module. The mindset shift, delegation skill, preparation value, verification, and limits pages are all strong.

**Tasks:**

1. **Page 04 (quality verification) — Add specific verification techniques from the report** (HIGH)
   - Add multi-model cross-validation: run the same query through multiple LLMs
   - Add asking AI to verify its own reasoning (reduces errors ~17%)
   - Add web search integration as a verification tool
   - Add: "Always click 'view analysis' to inspect generated code" for data analysis tasks
   - Add the WEF warning: "When foundational data is fragmented or inaccurate, AI models generate outputs that appear sophisticated but are fundamentally wrong"

2. **Page 01 (mindset shift) — Add the METR study as a reality check** (HIGH)
   - The METR study found experienced developers were 19% SLOWER with AI tools, despite believing they were 20% faster
   - This is a powerful teaching moment: "people overestimate AI's contribution because the work feels easier"
   - The Harvard study: no statistical difference in business performance between AI users and control groups
   - The key insight: "High-performing users benefited not because they got different advice but because they had better judgment about which AI advice to follow"
   - This belongs early in the delegation module to set a realistic tone

3. **Page 05 (limits) — Add skill atrophy and digital insider risk** (HIGH)
   - Skill atrophy is measurable: Microsoft/Carnegie Mellon study found less critical thinking with more AI use
   - MIT Media Lab: reduced brain activity, diminished memory retention, less original thinking
   - The Polish medical study: doctors' unassisted detection rate fell from 28.4% to 22.4% after AI exposure
   - Countermeasures: regularly practice without AI, "AI-free zones," "bicycle for the mind" model
   - Digital insider risk: agents inherit file permissions, 38% of employees share confidential data with AI
   - IBM data: breaches involving unauthorized AI tools cost $4.63M average (16% above average breach)

4. **Page 03 (preparation value) — Add the "15 iterations in 2 days beats 2 iterations in 5 days" principle** (MEDIUM)
   - From the report's discussion of McKinsey findings
   - Reinforce that preparation matters, but iteration speed is also a superpower
   - Add: "Technology delivers 20% of value; redesigning work delivers 80%"

5. **Consider adding an agent demo to this module** (MEDIUM)
   - Currently no agent demo. A demo showing the difference between micromanaging AI (step-by-step) vs delegating (outcome specification) would be powerful
   - Show the same task done both ways, with the delegation approach producing better results

### Module 7: module-advanced (60m)

**Current state:** 7 pages. Covers skills, subagents, persistent memory, code generation, MCP, and takeaways. This is the densest module and several pages need significant expansion.

**Tasks:**

1. **Page 04 (persistent memory) — Major expansion needed** (P0)
   - Currently very brief and generic. Needs to be specific about CLAUDE.md files:
     - They're markdown files automatically loaded at start of every Claude Code session
     - Best practices: keep under 300 lines (frontier models follow ~150-200 instructions before adherence drops)
     - Use progressive disclosure with imports and subdirectory files
     - Include build commands, code style, key architectural patterns, workflow rules
     - Commit to version control so the whole team benefits
   - Add AGENTS.md (Cursor, Zed, OpenAI standard, 60,000+ open-source projects)
   - Add Claude Projects (persistent system prompts with attached knowledge bases for non-coding work)
   - Add ChatGPT Custom Instructions
   - The key insight: "persistent instructions override per-conversation prompts, so they're the highest-leverage place to invest"
   - Consider an agent demo showing a CLAUDE.md file being created and its effect on subsequent interactions

2. **Page 06 (MCP) — Major expansion needed** (P0)
   - Currently surface-level. Needs:
   - Scale: 17,000+ MCP servers, 97 million monthly SDK downloads
   - Adoption: OpenAI (March 2025), Google DeepMind (April 2025), Linux Foundation governance (December 2025)
   - Specific useful servers: Google Drive, Slack, Notion, Jira, GitHub, Figma, PostgreSQL/SQLite
   - Security challenges: prompt injection vulnerabilities, token storage risks, "toxic agent" flows
   - The Salesforce Agentforce vulnerability (CVSS 9.4) as a cautionary tale
   - Practical workflow examples that are more specific than the current generic ones

3. **Page 05 (code generation) — Expand with vibe coding and citizen developer context** (HIGH)
   - Add "vibe coding" concept (Andrej Karpathy, Collins Dictionary Word of the Year)
   - Add stats: Y Combinator reported 25% of startups had 95% AI-generated codebases
   - Add realism: the Stack Overflow writer's experience ("almost immediately, error messages popped up")
   - Add Python as recommended language for AI-assisted non-developer coding
   - Add practical scope: "works for personal automation, departmental dashboards, web scraping, API integrations. Has real limitations for production use."
   - Add the Gartner projection: 75% of all apps built with low-code tools by 2026
   - The current page is good but lacks the research grounding and realistic expectations

4. **Page 03 (subagents) — Add concrete workflow patterns** (MEDIUM)
   - The report names specific patterns: ReAct (reasoning + tool calls), Plan-then-Execute (big model plans, small model executes), Hierarchical Task Decomposition, Generator-Evaluator Loop
   - Add the market context: $4.35 billion in 2025, projected >$100B by 2034
   - Make the subagent concept more concrete with a specific workflow example rather than abstract descriptions
   - Consider an agent demo showing a task being broken into subtasks

5. **Page 03 (subagents) — Add the "Meta Case Study: This Course Built Itself"** (P0 — HIGH)
   - This is a standout self-referential case study that should be a flagship teaching moment in the course. The actual process of creating and updating this course IS the content. Attendees will see the very tool they are learning about being used to build the course they are attending.
   - **The screenshot:** Embed the image at `/static/images/meta-course-creation-screenshot.png`, which shows the user instructing Claude Code to launch the course-director agent to review the training report and plan course updates. Use an annotated-image block or a standard image with callouts.
   - **What the case study should walk through:**
     1. **The deep research phase.** The instructor wrote an 8,000-word research report (`ai-training-report.md`) synthesising findings from Anthropic, OpenAI, Google, McKinsey, Bain, Deloitte, academic studies, and practitioner experience. This is context engineering in practice: assembling the right information before any AI work begins.
     2. **The course-director agent.** A subagent was launched with a detailed brief: read the research report, read every existing module, compare coverage against needs, and produce a comprehensive update plan. This is Hierarchical Task Decomposition — the user delegates to a "manager" agent, which does the analysis and planning.
     3. **The module-builder subagents.** The course-director can then launch independent module-builder agents (one per module, running in parallel) with specific briefs. Each module-builder creates the directory structure, `_module.yaml`, and markdown pages. This is parallelisation in action — multiple agents working on independent modules simultaneously.
     4. **Skill composition.** Module-builder agents invoke specialised skills for interactive content: `/quiz` for assessment blocks, `/agent-demo` for scripted AI conversation walkthroughs, `/mermaid` for structured diagrams, `/excalidraw` for freeform visuals, `/image-generator` for infographic-style graphics. This is the "tools that make tools" pattern from the report.
     5. **Human-in-the-loop oversight.** At every level, the human reviews and directs. The research report was human-authored. The course-director's plan was reviewed before module-builders were launched. Agent output was verified before inclusion. This demonstrates the "AI-first, human-verified" theme taught earlier in the course.
   - **What concepts this demonstrates in a single real example:**
     - Multi-agent hierarchies (user -> course-director -> module-builders)
     - Skill composition (agents invoking specialised skills on demand)
     - Background/parallel agents (multiple module-builders running simultaneously)
     - Deep research workflows feeding structured planning (report -> gap analysis -> prioritised plan)
     - Context engineering (the research report IS the context that makes agent work high-quality)
     - Persistent instructions (CLAUDE.md and agent instruction files defining how agents behave)
     - File-based workflows (everything is files: the report, the plan, the module YAML, the markdown pages)
     - The preparation-is-value principle (the quality of the research report determines the quality of everything downstream)
   - **Teaching approach:** This should NOT be a dry description. Structure it as a walkthrough with the screenshot, then peel back each layer. Start with "This course was updated by the system you are learning about." Then show the hierarchy diagram. Then walk through each stage, connecting it to concepts the learner has already encountered. End with a callout: "Every concept in this course — context engineering, delegation, skills, subagents, files, verification — was used to build the course you are sitting in."
   - **Include a quiz** testing whether learners can identify which concepts from the course are demonstrated in the case study.
   - **Priority justification:** This is rated P0 because it is a unique differentiator. No other AI training product can claim this level of self-referential demonstration. It transforms the subagents page from abstract theory into tangible proof. For enterprise buyers evaluating premium training products, this is the kind of "wow" moment that justifies the price point and makes the course memorable.

6. **Page 02 (skills) — Add meta-tooling concept** (MEDIUM)
   - From the report: "using AI to create custom tools you then use repeatedly"
   - Claude Code Skills as a specific example
   - Google Workspace Studio for Gmail, Drive, Calendar agents
   - The practical workflow: identify repetitive task -> describe to AI -> get working script -> test/iterate -> deploy -> return to AI for maintenance

7. **Page 01 (building systems) — Add token economics** (HIGH)
   - This is the right place for cost awareness
   - Agents consume 100x more tokens because entire conversation history resent with every message
   - Cost optimization: prompt caching (60-80% savings), model cascading (30-50% reduction), batch processing (50% discount)
   - Deloitte's framing: "treat AI economics with the same rigor as energy or capital allocation"
   - Token pricing trajectory: $20/M in 2022 to $0.40/M by mid-2025

### Module 8: module-synthesis (45m)

**Current state:** 4 pages. Workflow, five themes, next steps, reference. Serves as a good wrap-up.

**Tasks:**

1. **Page 03 (next steps) — Replace with research-backed quick wins** (HIGH)
   - The report identifies specific highest-ROI starting points: meeting summarization, email drafting/triage, document summarization, status report generation, research acceleration
   - Add the stat: daily GenAI users save 4+ hours/week, report 92% productivity improvement
   - Replace generic "tomorrow/this week/this month" with specific task recommendations tied to the quick wins research
   - Add the PwC finding: daily users report higher job security and higher salaries

2. **Page 03 — Add the "ten mistakes beginners make" from the report** (HIGH)
   - Being too vague, overloading single prompts, skipping role assignment, not iterating, ignoring limitations, not providing examples, sharing sensitive data, using wrong model for task, failing to use meta-prompting
   - This is a high-value practical checklist that learners will reference repeatedly

3. **Page 04 (reference) — Add model comparison reference** (MEDIUM)
   - Brief comparison of when to use which model/tool
   - Claude (strengths: long context, instruction following, coding, safety), GPT (strengths: multimodal, ecosystem, plugins), Gemini (strengths: Google integration, 2M context, multimodal)
   - Not a product comparison — a practical tool selection framework

4. **Page 01 (workflow) — Add McKinsey/PwC research backing** (MEDIUM)
   - McKinsey: high-performing organizations 3x more likely to have fundamentally redesigned workflows rather than bolting AI on
   - PwC: "Instead of cutting a few steps, rethink the workflow — an AI-first approach may turn into a single step"
   - Bain: AI leaders delivering 10-25% EBITDA gains
   - This gives the course's workflow model the research authority a premium product needs

5. **Page 02 (five themes) — Consider adding a sixth theme: "Verify and Own"** (MEDIUM)
   - The current five themes don't explicitly include verification/ownership as a standalone theme
   - Given the report's emphasis on risks (METR study, hallucinations, skill atrophy), this may deserve promotion to a core theme
   - Alternative: strengthen "AI-first, human-verified" theme with more specific verification techniques

---

## New Modules to Add

### NEW: module-risks — "Risks, Responsibility, and Realistic Expectations" (45-60m)

**Rationale:** The report devotes an entire section (Part 9) to risks, plus significant space in Parts 1 and 10 to realistic expectations. Currently this content is scattered thinly across multiple modules. For a premium product aimed at enterprises paying significant money, a dedicated risk module is essential. It demonstrates maturity and builds trust.

**Position in course:** After module-delegation, before module-advanced.
- Prerequisites: module-delegation
- This placement means learners understand how to use AI effectively before learning what can go wrong
- The advanced module then builds on risk-awareness to teach responsible advanced patterns

**Proposed pages:**

1. **01-reality-check.md** — "The Honest Assessment"
   - MIT Media Lab: 95% of organizations see no measurable returns from AI
   - McKinsey: only 1% of companies have reached AI maturity
   - 67% of AI workflow projects fail within six months
   - The METR study: 19% slower, believing 20% faster
   - Harvard study: no statistical difference in business performance
   - The point: AI works, but only with the right approach — which is what this course teaches
   - This is NOT a doom-and-gloom page. It is a calibration page that makes the rest of the training more credible.

2. **02-hallucination-depth.md** — "Hallucination: The Numbers"
   - Move deeper hallucination content here (currently light treatment in module-llms)
   - Specific rates by model and domain
   - Stanford legal hallucination data (58-82%)
   - The 4.3 hours/week fact-checking cost
   - Why hallucinations persist (OpenAI's "guessing over uncertainty" insight)
   - Practical countermeasures: multi-model validation, self-verification, search grounding, RAG limitations
   - Agent demo: asking AI to verify its own reasoning, showing how this catches errors

3. **03-skill-atrophy.md** — "The Skill Atrophy Problem"
   - Microsoft/Carnegie Mellon: less critical thinking with more AI use
   - MIT Media Lab: reduced brain activity, memory retention, original thinking
   - Polish medical study: detection rate fell from 28.4% to 22.4%
   - Countermeasures: AI-free zones, "bicycle for the mind" model, deliberate practice
   - This is a critical message for enterprise training: AI should enhance, not replace, human capability

4. **04-privacy-security.md** — "Privacy, Security, and the Digital Insider"
   - 38% of employees share confidential data without approval
   - IBM: AI-related breaches cost $4.63M average
   - Agentic AI risk: agents inherit file permissions, scan accessible data
   - McKinsey: treat AI agents as "digital insiders"
   - Gartner: 1 in 4 enterprise breaches will involve agentic AI
   - Practical guidelines: data classification, approved tools, enterprise vs consumer tiers

5. **05-legal-landscape.md** — "Legal and Copyright: What You Must Know"
   - Anthropic $1.5B settlement — largest copyright payout in US history
   - US Copyright Office: AI outputs get protection only where human author determined sufficient expressive elements
   - "Never say 'AI said to'" — if you use AI output, it becomes yours
   - International divergence: Japan/Singapore permit training without consent; France fined Google 250M
   - Practical implications for content creation, report writing, code generation

6. **06-key-takeaways.md** — Key takeaways page with quiz

**Module metadata:**
- difficulty: intermediate
- estimated_duration: 45m
- completion: require_quizzes: true, min_quiz_score: 0.6

### NEW: module-writing — "Document Creation and Data Analysis" (45-60m)

**Rationale:** The report devotes Parts 5 and 6 to these topics. They represent some of the highest-value use cases for knowledge workers. The current course teaches files and delegation but never shows the specific techniques for creating high-quality documents or analysing data with AI. For a premium product, this is a significant gap.

**Position in course:** After module-files, before module-delegation.
- Prerequisites: module-files
- This placement means learners understand file-based workflows before learning the specific techniques for document creation and data analysis
- The delegation module then builds on these practical skills

**Proposed pages:**

1. **01-writing-with-ai.md** — "Writing with AI: Beyond 'Write Me a Report'"
   - The sectional drafting method: outline first, expand sections individually, merge and smooth
   - The Stanford finding: AI collaboration increased productivity but reduced sense of ownership
   - The FoxPrint editorial warning: AI polishing can strip voice — use AI for mechanics, not voice
   - The optimal human-AI writing workflow (human defines -> AI generates ideas -> human selects -> AI drafts sections -> human identifies weak sections -> AI revises -> AI checks grammar -> human ensures voice and accuracy -> human fact-checks)

2. **02-style-transfer.md** — "Making AI Write Like You"
   - Style cards: encoding tone, vocabulary, sentence structure, audience as reusable prompts
   - EchoWriting: feeding AI 15-20 samples, analyzing style patterns, creating persistent style prompt
   - Grounding with examples: "provide 2-3 paragraphs of your best writing and ask AI to analyze your style"
   - Agent demo showing style analysis and then writing in that style

3. **03-data-analysis.md** — "Data Analysis Without a Data Team"
   - The three tiers: AI add-ins, chat-to-SQL, full-stack AI analysts
   - The practical upload-and-analyse workflow: clean data, upload, describe first, then ask specific questions
   - Data cleaning: duplicate removal, missing values, date standardization, outlier identification
   - ChatGPT capabilities: 10 files, 512MB each, Python sandbox, Matplotlib charts
   - Claude capabilities: code execution, downloadable outputs, interactive Plotly/D3 visualizations

4. **04-verification-traps.md** — "What Can Go Wrong with Data and Documents"
   - The WEF warning about sophisticated-looking but fundamentally wrong outputs
   - "Claude tends to hallucinate when working with large datasets or too many filters"
   - ChatGPT's Code Interpreter cannot actually understand a visualization
   - Always inspect generated code
   - AI is better at organizing/summarizing than genuine intellectual synthesis
   - Verify it hasn't fabricated connections between sources

5. **05-key-takeaways.md** — Key takeaways with quiz

**Module metadata:**
- difficulty: intermediate
- estimated_duration: 45m
- completion: require_quizzes: true, min_quiz_score: 0.6

---

## Updated Course Structure

### Current (8 modules, ~435 minutes)

```
module-opening (30m)
  -> module-llms (60m)
    -> module-context (60m)
      -> module-prompting (60m)
        -> module-files (60m)
          -> module-delegation (60m)
            -> module-advanced (60m)
              -> module-synthesis (45m)
```

### Proposed (10 modules, ~525 minutes / ~8.75 hours)

```
module-opening (30m) — The February 2026 Moment
  -> module-llms (60m) — How LLMs Actually Work
    -> module-context (60m) — Context Engineering
      -> module-prompting (75m) — The Art of Prompting [+15m for new pages]
        -> module-files (60m) — Files: The Unit of Work
          -> module-writing (45m) — Document Creation and Data Analysis [NEW]
            -> module-delegation (60m) — Delegation and the AI-First Philosophy
              -> module-risks (45m) — Risks, Responsibility, and Realistic Expectations [NEW]
                -> module-advanced (75m) — Advanced Patterns [+15m for expansions]
                  -> module-synthesis (45m) — Putting It Together
```

**Notes on structure:**
- Total duration increases from ~435m to ~555m (~9.25 hours). This is long for a single day. Options:
  - (a) Make it a 1.5-day course
  - (b) Trim some existing content to make room (the key-takeaways pages on modules 3-6 could be shortened)
  - (c) Make module-risks and/or module-writing optional "deep dive" tracks
  - (d) Split into Day 1 (foundations: opening through delegation) and Day 2 (advanced: risks through synthesis)
- The new modules slot naturally into the learning progression
- module-writing after module-files is logical (you learn files, then learn what to do with them)
- module-risks after module-delegation is logical (you learn to delegate, then learn what can go wrong)

### course.yaml update

```yaml
title: "AI Training: From Novice to Expert"
description: "A comprehensive training on working effectively with AI, from understanding LLMs to building AI-powered workflows."
modules:
  - module-opening
  - module-llms
  - module-context
  - module-prompting
  - module-files
  - module-writing        # NEW
  - module-delegation
  - module-risks          # NEW
  - module-advanced
  - module-synthesis
```

---

## Priority Ordering

### P0 — Must Do (Critical for Premium Positioning)

These changes address the most significant gaps. Without them, the course will feel outdated to any informed buyer.

1. **Update context window data** (module-context/02) — Currently wrong. Takes 15 minutes to fix.
2. **Add specific prompting techniques page** (module-prompting, new page) — The biggest content gap. The report's techniques hierarchy is the most practical teaching in the report.
3. **Expand persistent memory page** (module-advanced/04) — Currently too vague for a premium product. Needs CLAUDE.md specifics.
4. **Expand MCP page** (module-advanced/06) — Currently surface-level. Needs the scale/adoption/security data.
5. **Add "context engineering" framing** (module-context/01) — The report calls this the most important shift. Must be in the course.
6. **Update opening with specific products and timeline** (module-opening/01) — Premium buyers will notice if this is vague.
7. **Add meta case study to subagents page** (module-advanced/03) — Unique differentiator. The self-referential "this course built itself" case study with screenshot at `/static/images/meta-course-creation-screenshot.png`. No competing product can replicate this. Demonstrates every major concept in the course in a single real example.

### P1 — Should Do (Significant Quality Improvement)

These changes meaningfully improve the course's depth and authority.

8. **Create module-risks** (new module) — Enterprise buyers need this. Shows maturity.
9. **Create module-writing** (new module) — Fills the biggest use-case gap.
10. **Add hallucination data to LLMs module** (module-llms/03) — Specific numbers make the teaching concrete.
11. **Add METR study and skill atrophy to delegation** (module-delegation/01, 05) — Reality calibration.
12. **Add iteration pattern to prompting** (module-prompting, new page) — High-value practical skill.
13. **Add verification techniques from report** (module-delegation/04) — Multi-model validation, self-verification.
14. **Update synthesis quick wins with research** (module-synthesis/03) — Specific ROI data makes recommendations credible.
15. **Expand code generation with vibe coding** (module-advanced/05) — Missing major cultural moment.

### P2 — Nice to Have (Polish and Depth)

These changes add depth and authority but are not critical for launch.

16. **Add token economics** (module-advanced/01) — Cost awareness for sophisticated buyers.
17. **Add model comparison reference** (module-synthesis/04) — Practical tool selection.
18. **Add style card/EchoWriting to grounding** (module-files/05) — Advanced technique.
19. **Add file preparation best practices** (module-files/03) — From report's practical guidance.
20. **Add subagent workflow patterns** (module-advanced/03) — Named patterns from the report.
21. **Add ten beginner mistakes** (module-synthesis/03) — Practical checklist.
22. **Add McKinsey/PwC workflow research** (module-synthesis/01) — Authority backing.
23. **Add agent demo to delegation module** — Currently missing interactive element.

### P3 — Future Consideration

24. **Multi-model arbitrage teaching** — Power user technique.
25. **Prompt library management** — Shared asset approach.
26. **Agent demo showing vibe coding** — Would require new demo infrastructure.

---

## Interactive Element Audit

### Current Agent Demos
- module-opening/03: "From Chaos to Report" — Excellent. Keep as-is.

### Agent Demos to Add
- **module-prompting** (new techniques page): Demo showing multishot prompting or XML-structured prompting
- **module-writing** (style transfer page): Demo showing style analysis and writing in the user's style
- **module-delegation**: Demo showing micromanaging vs delegation approaches
- **module-risks** (hallucination page): Demo showing AI self-verification catching an error
- **module-advanced** (persistent memory page): Demo showing how CLAUDE.md affects responses
- **module-advanced** (subagents page, meta case study): Consider an agent demo showing the course-director launching a module-builder, or showing how the research report is fed into the planning process. This could be the most compelling demo in the entire course because it is real, not contrived.

### Quiz Coverage
Every module has quizzes. The new modules (module-risks, module-writing) will need 3-4 quizzes each. Ensure answer positions vary across all quizzes (currently many have answer index 1 or 2 — should be more varied).

### Diagrams and Visuals
Several new diagrams will be needed:
- Updated context window size comparison
- Context engineering conceptual diagram
- Prompting techniques hierarchy
- Data analysis workflow tiers
- Risk landscape overview
- Agentic workflow patterns
- **Meta case study hierarchy diagram** — a visual showing: User -> Claude Code -> course-director agent -> module-builder agents (x N, parallel) -> skills (/quiz, /agent-demo, /mermaid, /excalidraw, /image-generator). This should use the same visual language as the existing subagent diagram but with the real labels from this project. Mermaid or Excalidraw format.
- **Meta case study screenshot annotation** — the screenshot at `/static/images/meta-course-creation-screenshot.png` should be embedded with callouts identifying key elements visible in the interface

---

## Content Tone and Style Notes

The current course tone is excellent: direct, opinionated, no-nonsense, respects the learner's time. When adding new content, maintain:

- **Short paragraphs.** The current pages use 1-3 sentence paragraphs. Do not introduce walls of text.
- **Bold assertions followed by evidence.** The course leads with claims and backs them up. It does not hedge.
- **Practical framing.** Every concept connects to "what will you do differently on Monday?"
- **No filler.** If a concept can be a callout instead of a page, make it a callout.
- **British English.** The course uses "organise," "analyse," "programme," etc. Maintain this.
- **Specific over generic.** Replace "AI tools" with "Claude Cowork," "ChatGPT Agent Mode," "Cursor." Name names.

---

## Implementation Sequence

Given the priority ordering, the recommended implementation sequence is:

**Phase 1: Critical fixes (P0 items, can be done in parallel)**
1. Fix context window data (module-context/02)
2. Add context engineering framing (module-context/01)
3. Update opening with specifics (module-opening/01)
4. Add prompting techniques page (module-prompting, new page)
5. Expand persistent memory (module-advanced/04)
6. Expand MCP (module-advanced/06)
7. Add meta case study to subagents page (module-advanced/03) — with screenshot, walkthrough, and quiz

**Phase 2: New modules (P1 items, can be done in parallel with each other)**
8. Build module-risks (new module, 6 pages)
9. Build module-writing (new module, 5 pages)
10. Update course.yaml with new module order

**Phase 3: Enrichment (P1 continued, module-by-module)**
11. Enrich module-llms (hallucination data, capabilities framing)
12. Enrich module-delegation (METR study, skill atrophy, verification techniques)
13. Enrich module-prompting (iteration pattern page)
14. Enrich module-synthesis (quick wins, beginner mistakes, McKinsey research)

**Phase 4: Polish (P2 items)**
15. Token economics, model comparison, style cards, file preparation
16. Agent demos for new content
17. Diagram and visual updates
18. Quiz answer position audit

---

## Appendix: Key Statistics from the Report

These statistics should be incorporated into the course at the locations noted:

| Statistic | Source in Report | Course Location |
|---|---|---|
| 41% of code now AI-generated | Part 1 | module-opening/01, module-llms/01 |
| Claude Opus 4.6 = 1M token context | Part 1 | module-context/02 |
| METR: task length doubles every 4 months | Part 1 | module-opening/02 |
| Context engineering adopted by every major AI company | Part 3 | module-context/01 |
| CLAUDE.md: 150-200 instruction limit before adherence drops | Part 3 | module-advanced/04 |
| 60,000+ projects adopted AGENTS.md | Part 3 | module-advanced/04 |
| 17,000+ MCP servers, 97M monthly downloads | Part 4 | module-advanced/06 |
| Cursor: $500M ARR, up from $1M in 12 months | Part 4 | module-advanced/05 |
| 57% of scientists use AI writing help | Part 5 | module-writing/01 |
| Gemini-2.0-Flash hallucination rate: 0.7% | Part 9 | module-risks/02 |
| Stanford legal hallucination: 58-82% | Part 9 | module-risks/02 |
| 47% made business decisions on hallucinated content | Part 9 | module-risks/02 |
| METR: 19% slower despite believing 20% faster | Part 9 | module-delegation/01 |
| Doctor detection rate: 28.4% -> 22.4% after AI | Part 9 | module-risks/03 |
| IBM: AI breaches cost $4.63M average | Part 9 | module-risks/04 |
| 38% share confidential data without approval | Part 9 | module-risks/04 |
| Anthropic: $1.5B copyright settlement | Part 9 | module-risks/05 |
| 95% of organizations see no measurable AI returns | Part 10 | module-risks/01 |
| Daily GenAI users save 4+ hours/week | Part 10 | module-synthesis/03 |
| McKinsey: 3x more likely to redesign workflows | Part 10 | module-synthesis/01 |
| Y Combinator: 25% of startups 95% AI-generated code | Part 7 | module-advanced/05 |
| $285B market selloff from Cowork launch | Part 2 | module-opening/01 |
