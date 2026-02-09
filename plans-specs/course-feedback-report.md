<!-- PROGRESS: COMPLETE -->

# Course Feedback Report

**Reviewed:** 2026-02-09
**Reviewer:** Claude (automated student-perspective review)
**Course:** AI Training: From Novice to Expert in One Day
**Modules reviewed:** 10 of 10
**Status:** COMPLETE

## Executive Summary

- **Overall quality is Good.** The course has a strong conceptual framework (five themes), honest treatment of AI limitations, and practical takeaways. Most content is clear, well-written, and appropriately challenging for a non-technical professional audience.
- **One critical factual issue.** The Anthropic "£1.5 billion settlement" claim in the Legal and Copyright page (module-risks p5) cannot be verified and may be a hallucinated statistic. This must be confirmed or removed before delivery — a fabricated fact in a module about hallucination risks would be catastrophic for credibility.
- **Significant content duplication.** Nine pieces of content appear across 2-4 modules (sectional drafting, EchoWriting, METR study, skill atrophy data, security data, fact-checking stats). This is the largest structural issue. Consolidation needed.
- **Currency inconsistencies.** IBM breach cost in USD vs GBP, France fine in GBP instead of EUR, Anthropic settlement in GBP described as "US history." These undermine professionalism.
- **Light on interactive variety.** Only 1 agent demo in the entire course (Module 1). Modules 5-8 have zero diagrams or visual elements. Most pages follow the same pattern: text → callout → quiz. More visual and interactive variety would elevate the experience.

## Module-by-Module Review

### Module 1: Opening: The February 2026 Moment (module-opening)

**Overall:** Strong opening that immediately establishes the paradigm shift from chatbots to agentic AI. Confident, punchy writing with specific data points. The agent demo is an excellent "show, don't tell" moment.
**Rating:** Good

#### Page 1: The February 2026 Moment (4 slides)
- **Content:** Opens strong with a bold claim and backs it up. The before/after comparison is crisp. The 50x Reframe (Azeem Azhar) is a great pedagogical device. The AI timeline provides useful historical grounding. The "$285 billion market selloff" claim and "41% of all code written globally is now AI-generated" are powerful but need verification — these are the kind of statistics that could date badly or be disputed.
- **Interactive elements:** 1 quiz (50x Reframe concept — good, tests understanding not recall), 2 callouts (info + tip — appropriate use), 1 SVG timeline image (renders correctly, visually clear).
- **Slide navigation:** All 4 slides work correctly. Sidebar highlights current section. Slide counter updates properly.
- **Issues:**
  - The "Cowork Plugins" claim (30 January 2026, $285bn selloff) — verify this is accurate. If students Google this and find discrepancies, credibility suffers.
  - "41% of all code" stat needs a citation or source.
  - British English: "specialised" correctly used. No American spelling spotted.
- **Suggestions:** Add a source/citation for the 41% code stat, even if just "(Source: GitHub Copilot report, 2025)" inline.

#### Page 2: What You'll Learn Today (3 slides)
- **Content:** Clear learning objectives. The 6 numbered items are concrete and outcome-focused. Good that #6 is "Know when NOT to use AI" — sets a balanced tone early.
- **Interactive elements:** 1 quiz (paradigm shift — straightforward, correct answer is obvious), 1 callout (Demonstration First — effective teaser for the demo).
- **Issues:**
  - The quiz ("What is the fundamental shift in AI...") is too easy. The distractors ("slow to fast", "expensive to cheap", "text to image") are obviously wrong. A student who wasn't paying attention could still guess correctly.
  - Page 2 slide 2 is titled "The Demo: From Chaos to Report" — this is confusing because the actual demo is on page 3. This slide just describes what the demo will do, not the demo itself. It reads like intro text that belongs on the demo page.
- **Suggestions:** Strengthen the quiz distractors. Consider something like "From single-turn Q&A to multi-step autonomous execution" as a more plausible wrong answer.

#### Page 3: Demo: From Chaos to Report (1 page, no slides)
- **Content:** Brief intro paragraph that sets up the agent demo. The description references "the sidebar on the right" and "scratchpad" — but the actual workspace has a LEFT file explorer (not a right sidebar) and no panel labelled "scratchpad". This text is stale/inaccurate relative to the current UI.
- **Interactive elements:** 1 agent demo ("From Chaos to Report"). Launches into full-screen workspace. File explorer shows 5 files with coloured type icons. CSV file renders as a proper table. Conversation flow: notes → user delegation → AI reads all files → synthesises → writes report. New file appears in explorer. Excellent demonstration.
- **Navigation:** "NEXT MODULE: How LLMs Actually Work" correctly identifies the module boundary.
- **Issues:**
  - **P1: Stale copy.** "The sidebar on the right is an AI conversation walkthrough. The **scratchpad** shows five messy project files" — the scratchpad is now a file explorer on the left, and the chat sidebar title doesn't say "scratchpad". This will confuse students.
- **Suggestions:** Update the intro text to match the current workspace layout: "The file explorer on the left shows five messy project files... The conversation panel on the right walks through the AI interaction."

#### Module-Level Notes
- Good pacing: 30 min estimated, 3 pages, opening hook → objectives → demo. Classic intro structure.
- The demo is the star. It viscerally demonstrates the thesis before any theory.
- Module has no difficulty badge visible on the module selector card (just page count and duration) — consider adding difficulty indicators.
- No "key takeaways" page — acceptable for an opening module, but slightly inconsistent with later modules that all have one.

### Module 2: How LLMs Actually Work (module-llms)

**Overall:** Solid foundational module that demystifies LLMs without condescension. The "prediction engine" framing is the right pedagogical anchor. Hallucination coverage is excellent with specific, cited statistics. Good quiz distribution (1 per page).
**Rating:** Good

#### Page 1: LLMs Are Prediction Engines (4 slides)
- **Content:** Excellent core explanation. "It's a prediction engine" is the right one-sentence summary. The "capital of France" example with probability percentages makes next-token prediction tangible. The strengths/weaknesses framing ("the hard problems are easy and the easy problems are hard") from Carnegie Endowment is a memorable insight. The "41% of code" stat reappears (also cited in Module 1) — consistent but still needs a source.
- **Interactive elements:** 1 quiz (core mechanism — clear, distractors are weak but adequate), 2 callouts (info + tip). No diagrams — a visual showing token-by-token prediction would strengthen this page.
- **Issues:**
  - Carnegie Endowment attribution for "hard problems are easy" — verify this is correct and add a link if possible.
  - No visual/diagram on this page. The "Next-Token Prediction" concept begs for an animated or stepped visual.
- **Suggestions:** Add a simple diagram showing the token prediction chain visually (like a flowchart of the sequence building up).

#### Page 2: The Training Process (6 slides)
- **Content:** Clear three-phase explanation (pre-training → fine-tuning → RLHF). The "Think of it as:" analogies for each phase work well. Knowledge cutoff explanation is important and well-placed. The "Current Model Landscape" section lists Claude 4.x, GPT-5.x, and Gemini 3 Pro with specific details. "Extended thinking" concept introduced clearly.
- **Interactive elements:** 1 quiz (RLHF purpose — good), 2 callouts (note type), 1 SVG diagram (training process flow — renders correctly, clear visual).
- **Issues:**
  - "Gemini 3 Pro" scoring "1,501 Elo on LMArena" — verify this is accurate/current. LLM benchmarks change rapidly.
  - The callout about tools extending beyond training data mentions Claude, ChatGPT, and Gemini having web search. This is accurate but may become outdated.
- **Suggestions:** Consider noting these model details may change — perhaps a callout saying "Model landscape accurate as of February 2026."

#### Page 3: Hallucinations and Confident Wrongness (6 slides)
- **Content:** This is the strongest page in the module. Opens with a powerful reframe: "This Is Not a Bug" (danger callout). The statistics are specific and sourced: Gemini-2.0-Flash 0.7%, o3 at 33% on PersonQA, legal queries 58-82% (Stanford). The "47% of enterprise AI users made major decisions on hallucinated content" and "4.3 hours per week fact-checking" stats ground the problem in economic reality. Practical countermeasures (multi-model cross-validation, self-verification, web search, grounding) are actionable.
- **Interactive elements:** 1 quiz (why hallucinations happen — good), 3 callouts (2 danger, 1 warning — appropriate escalation of severity). Danger callouts have red left border and icon — visually distinct and attention-grabbing.
- **Issues:**
  - "GPT-4o achieved 90% accuracy when it could search" — check this stat's source.
  - "reduces errors by roughly 17%" for self-verification — needs a citation.
  - The statistics are very specific. If even one is wrong, it undermines the whole "verify everything" message (ironic).
- **Suggestions:** Add brief source citations for the key statistics, even if just "(Stanford, 2025)" or "(LMSys, 2026)".

#### Page 4: Key Takeaways: How LLMs Work (7 slides)
- **Content:** Good summary of the 4 key points. The "What About Intelligence?" section is thoughtful and balanced — doesn't oversell or undersell. "It's something new" is a good landing. The transition to the Context module is smooth.
- **Interactive elements:** 1 quiz (practical implication — good), 1 callout (tip).
- **Issues:**
  - Slide 1 is very sparse — just a heading and one line of text. This is a thin slide that feels like wasted space. The H1 "Key Takeaways" and "Let's summarize..." could be combined with the first substantive point.
  - 7 slides for a takeaways page feels like too many — each key point gets its own slide, which slows the pace for what should be a quick recap.
- **Suggestions:** Condense the takeaways to fewer slides (3-4 max). The intro slide adds no value — merge it with the first takeaway.

#### Module-Level Notes
- All 4 quizzes work correctly (tested incorrect and correct answers). Feedback is educational.
- Good progression: mechanism → training → failure modes → summary.
- The hallucination page is the highlight — the most valuable content in the module.
- No agent demo in this module (appropriate — it's conceptual).
- British English consistent: "optimises", "behaviour" not checked exhaustively but no American spellings spotted.
- Duration estimate 60m seems right for 4 pages with rich content.

### Module 3: Context — The Most Important Concept (module-context)

**Overall:** The conceptual anchor of the course. "Everything is context" is the right message and it's hammered home effectively. The Karpathy quote ("LLM is like the CPU, context window is like RAM") is perfect for this audience. Good progression from concept → mechanics → hierarchy → decay → summary. Slightly text-heavy — could use more visual variety.
**Rating:** Good

#### Page 1: Everything Is Context (4 slides)
- **Content:** Opens with the course's most important concept and frames it well. The Karpathy quote provides authority. "Context engineering" as a discipline (replacing "prompt engineering") is an important reframe. The diagnostic question ("What was the context?") is a practical takeaway. The enumerated list of what constitutes context (system prompts, persistent instructions, conversation history, files, tools, current message) is thorough.
- **Interactive elements:** 1 quiz (what determines output quality — good), 2 callouts (tip + warning).
- **Issues:**
  - "behavior" appears on line 56 of the source — should be "behaviour" for British English.
  - No diagram or visual on this page. The concept of "everything going into the context window" begs for a visual.
- **Suggestions:** Add a diagram showing different types of context flowing into the model. Fix "behavior" → "behaviour".

#### Page 2: The Context Window (4 slides)
- **Content:** The table showing context window growth from GPT-3 (4K) to 2026 (1M tokens) is excellent — makes the progress tangible. "Lost in the middle" effect introduced early. Token explanation callout is clear. "Infinite Chats" feature mentioned. The "What Goes Where" section provides a practical mental model.
- **Interactive elements:** 1 quiz (why larger windows matter — good), 1 callout (info about tokens), 1 SVG diagram (context window visualisation).
- **Issues:**
  - "Claude Opus 4.6 now supports 1 million tokens" — verify this is accurate.
  - "Gemini models support up to 2 million tokens" — check if this is Gemini 3 or earlier.
- **Suggestions:** None major. This is a well-structured page.

#### Page 3: Context Hierarchy and System Prompts (6 slides)
- **Content:** The 5-level hierarchy (system prompts → persistent instructions → immediate instructions → provided context → conversation history) is clear and actionable. System prompt examples make the invisible visible. The emphasis on persistent instructions as "highest-leverage" is an important practical insight. CLAUDE.md, Custom Instructions, AGENTS.md — specific tool references.
- **Interactive elements:** 1 quiz (why same model behaves differently — good), 3 callouts (note + 2 tips), 1 SVG diagram (context hierarchy).
- **Issues:**
  - The hierarchy claims persistent instructions "override per-conversation prompts" — this is somewhat model/platform-specific and may not be universally true. Could be slightly misleading.
- **Suggestions:** Soften the "override" language to "shape" or "strongly influence".

#### Page 4: Context Decay and Fresh Starts (7 slides)
- **Content:** This is the most practical page in the module. The degradation pattern (turns 1-10 fresh, 20-30 blurring, 50+ saturated) is a useful heuristic. The four causes (information density, contradictions, truncation, attention dilution) are well-explained. Claude Code commands (/clear, /compact, /rewind) are specific and actionable. The "Conversation as Refinement" pattern and Markdown section feel like they belong in the prompting module rather than here — they're about iteration, not context decay.
- **Interactive elements:** 1 quiz (degradation solution — good), 4 callouts (warning + tip + warning + tip). The callout density is high — 4 on one page is borderline overuse.
- **Issues:**
  - The Markdown section and "Conversation as Refinement Pattern" feel misplaced. They're about prompting technique, not context decay. This makes the page feel unfocused.
  - 7 slides is a lot for a "10m" estimated page. Students may lose patience.
  - 4 callouts is borderline overuse for one page.
- **Suggestions:** Move the Markdown and Conversation as Refinement sections to the prompting module. This page should focus tightly on decay → diagnosis → solution (fresh starts + tools).

#### Page 5: Key Takeaways: Context (6 slides)
- **Content:** Solid summary. The 4 key points distill the module well. The practical checklist (Before you prompt / When output is poor / For important tasks) is actionable. Smooth transition to prompting module.
- **Interactive elements:** 1 quiz (diagnostic question — good), 1 callout (tip).
- **Issues:**
  - Same sparse-first-slide issue as Module 2's takeaways. The intro slide is just "Let's summarize..." — merge with first point.
- **Suggestions:** Condense intro slide with first takeaway.

#### Module-Level Notes
- The strongest module conceptually — "Everything is context" is the course's core insight.
- Light on visuals. Only 2 SVGs across 5 pages. More diagrams would help.
- The context decay page tries to cover too much — the Markdown and iteration sections should move to prompting.
- British English mostly consistent but "behavior" slipped through on page 1.
- No agent demo (appropriate for a conceptual module).

### Module 4: The Art of Prompting (module-prompting)

**Overall:** The largest module (7 pages, 85m estimated) and the practical core of the course. Techniques are well-chosen and ordered by effectiveness. Good use of before/after comparisons. The "Three Questions" framework is the standout — simple enough to remember, specific enough to use. Some overlap with context module on iteration/refinement topics.
**Rating:** Good

#### Page 1: The Evolution of Prompting (2 slides)
- **Content:** Good framing of the evolution: prompt engineering → context engineering → delegation → agentic orchestration. The timeline (2022-2026) gives historical context. The message "you don't need prompting secrets, you need clear communication" sets the right expectations.
- **Interactive elements:** 1 quiz (evolution shift — good), 1 callout (info), 1 SVG diagram (prompting evolution).
- **Issues:** None significant. Short and effective intro.

#### Page 2: The Basic Framework (3 slides)
- **Content:** The 5-element framework (Role, Task, Context, Format, Constraints) is standard but well-presented. The bad/good prompt comparison is effective — the "bad" prompt is just "Write me a report about Q4" and the "good" prompt specifies everything. The "When to Use What" section manages expectations (not every request needs all five elements).
- **Interactive elements:** 1 quiz (which element = expertise — good), 1 callout (tip).
- **Issues:**
  - The framework is very similar to many existing prompt frameworks (RTCFC vs CRISPE vs others). It's well-executed but not novel. For a premium course, consider acknowledging alternatives or citing the source.
- **Suggestions:** Minor — could mention this is one of several frameworks but chosen for simplicity.

#### Page 3: The Techniques That Work (7 slides)
- **Content:** This is the highest-value page in the module. 5 techniques ranked by effectiveness: (1) Be specific and direct, (2) Use examples/multishot, (3) Enable thinking, (4) Structure with XML tags, (5) Chain complex prompts. Each has clear before/after examples. The "Colleague Test" is a memorable heuristic. The XML-structured prompt at the end is an excellent comprehensive example. The hierarchy of techniques at the bottom is a useful reference.
- **Interactive elements:** 1 quiz (multishot effectiveness — good), 4 callouts (tip + info + tip + warning).
- **Issues:**
  - "Anthropic's own research" and "Anthropic's signature technique" — slightly promotional. This is a training course, not marketing.
  - 7 slides for one page is heavy but justified by the density of techniques.
  - The "Before and After" comprehensive example is excellent but long — could overwhelm students in a live setting.
- **Suggestions:** Tone down the Anthropic-specific framing slightly. The techniques are good regardless of attribution.

#### Page 4: The Three Questions (4 slides)
- **Content:** The standout framework of the entire course. "What does done look like?", "What context does it need?", "What are the boundaries?" — simple, memorable, immediately actionable. The worked example (analysing customer feedback) shows the framework applied end-to-end, including the resulting prompt. The "Verification Test" callout is a good meta-skill.
- **Interactive elements:** 1 quiz (success criteria question — good), 1 callout (tip).
- **Issues:** None. This is a very strong page.
- **Suggestions:** This framework is so good it should be referenced more prominently in the opening module as a preview.

#### Page 5: Meta-Prompting (3 slides)
- **Content:** The "ask AI to help you write the prompt" technique is a genuinely useful insight. The worked example (customer feedback analysis) shows the conversational flow well. The transition from meta-prompting to "describe outcomes, not process" is smooth. "Don't describe the process, describe the outcome" — strong closing message.
- **Interactive elements:** 1 quiz (meta-prompting benefit — good), 1 callout (tip).
- **Issues:**
  - This is a relatively thin page for a 10m estimate. The concept could be explained in one slide.
- **Suggestions:** Could be merged with the Three Questions page or the Techniques page to reduce module page count.

#### Page 6: The Iteration Pattern (8 slides)
- **Content:** Thorough coverage of iterative refinement. The "make this better" vs specific feedback comparison is the key insight. The worked example (blog post → 3 iterations) is excellent and realistic. "When to Stop Iterating" and "Iteration vs Starting Fresh" sections add nuance. The "Colleague Test (Revisited)" connects back to earlier content.
- **Interactive elements:** 1 quiz (effective feedback — good), 3 callouts (info + tip + warning).
- **Issues:**
  - 8 slides is very heavy for what is essentially one concept (specific feedback beats generic). Some sections could be merged.
  - Overlap with Module 3 page 4 (Context Decay) which also covered the "Conversation as Refinement" pattern. This is the right place for it, but it should be removed from Module 3.
  - "Practice Exercise" at the end is good but won't be interactive in the current platform — it's just text.
- **Suggestions:** Trim to 5-6 slides max. Remove the iteration content from Module 3.

#### Page 7: Key Takeaways: Prompting (5 slides)
- **Content:** Good summary with 6 key points. The practical checklist with checkboxes is a nice touch (though not interactive). The transition to the files module is smooth.
- **Interactive elements:** 1 quiz (most important change — good), 1 callout (tip).
- **Issues:** The checkbox list renders as text, not interactive checkboxes. Not a bug but slightly disappointing.
- **Suggestions:** None major.

#### Module-Level Notes
- Largest module with good reason — prompting/delegation is the core practical skill.
- 85m estimated duration seems right but is long for a single module in a full-day course.
- Good quiz distribution (1 per page, 7 total).
- Some overlap with Module 3's context decay page on iteration/refinement.
- The Three Questions framework and Techniques page are the highlights.
- No agent demo — consider adding one showing the iteration pattern in practice.

### Module 5: Files — The Unit of Work (module-files)

**Overall:** Solid practical module that establishes the file-based workflow paradigm. Clean structure: paradigm shift → why files matter → inputs → outputs → grounding → summary. Content is clear and actionable. Some overlap with Module 6 (writing) on the sectional drafting method. No visual variety — entirely text-based with minimal interactive elements beyond quizzes.
**Rating:** Good

#### Page 1: The Paradigm Shift: Chat to Files (4 slides)
- **Content:** Effective opening that frames the "old way vs new way" clearly. The 6-step chat workflow vs 3-step file workflow contrast is immediately persuasive. The "Key Distinction" callout clarifies the boundary between chat and files.
- **Interactive elements:** 1 quiz (paradigm difference — good), 1 callout (tip), 1 SVG image (file workflow diagram — renders correctly).
- **Issues:** First slide is thin — just a heading and one sentence. Could be merged with slide 2.
- **Suggestions:** Merge the intro slide with the "Old Way vs New Way" slide.

#### Page 2: Why Files Matter (no H2 sections)
- **Content:** Six clear benefits (persistence, shareability, functionality, iteration, professionalism, completeness). Each is concise and specific. The "functionality" section with "Excel formulas work" is particularly concrete.
- **Interactive elements:** 1 quiz (functionality meaning — good), 1 callout (info).
- **Issues:** No visual variety. Six benefits as text blocks feels like a list that could use icons or a visual summary.
- **Suggestions:** Consider an infographic or comparison table to break up the wall of text.

#### Page 3: Input Files: Giving AI Your Data (no H2 sections after first)
- **Content:** Comprehensive coverage of file types AI can process. The "Provide, Don't Describe" callout is the core insight. Best practices section (machine-readable formats, token economics, file naming) is practical. Limitations section is honest.
- **Interactive elements:** 1 quiz (why upload beats describing — good), 2 callouts (info + tip).
- **Issues:** The "Current Capabilities" callout mentions ChatGPT limits (10 files, 512MB) and Claude features specifically — these will date quickly.
- **Suggestions:** Add a date qualifier to capability mentions or frame as "as of early 2026."

#### Page 4: Output Files: Receiving Deliverables (no H2 sections)
- **Content:** Clear enumeration of what AI can create. The "Sectional Drafting Method" appears here AND in Module 6 page 1 — this is duplicated content. The 3-step method (outline → expand → merge) is good but shouldn't appear twice.
- **Interactive elements:** 1 quiz (how to request files — good), 2 callouts (info + tip).
- **Issues:**
  - **P1: Duplicated content.** The Sectional Drafting Method is described identically here and in module-writing page 1. Should exist in only one place with a cross-reference.
- **Suggestions:** Remove the sectional drafting section from this page. It fits better in module-writing. Add a brief mention: "We'll cover the optimal method for creating long documents in the next module."

#### Page 5: Grounding: Your Specific Knowledge (no H2 sections)
- **Content:** Excellent page. The "grounding" concept — teaching AI your specific context — is well-explained with clear before/after examples. Style cards and EchoWriting techniques are practical and specific. The grounding materials table by task type is a useful reference. Multi-document synthesis covered with appropriate caveats.
- **Interactive elements:** 1 quiz (grounding purpose — good), 2 callouts (tip + warning).
- **Issues:** The EchoWriting technique is also covered in module-writing page 2. Another duplication.
- **Suggestions:** Consolidate EchoWriting coverage — keep the brief mention here and the detailed treatment in module-writing, or vice versa.

#### Page 6: Key Takeaways: Files (no H2 sections)
- **Content:** Good summary with 4 clear points matching the module's main concepts. The "Practical Test" checklist (4 diagnostic questions) is immediately usable. Smooth transition to delegation.
- **Interactive elements:** 1 quiz (generic output fix — good), 1 callout (tip).
- **Issues:** None significant. Clean summary.

#### Module-Level Notes
- 6 pages, 60m estimated — seems right.
- Good quiz distribution (1 per page, 6 total).
- Content overlap with module-writing on sectional drafting and EchoWriting needs resolving.
- Entirely text-based — no agent demos, no diagrams beyond the one SVG on page 1. Could use more visual variety.
- British English consistent throughout ("analysing", "organisation", "colour").

### Module 6: Document Creation and Data Analysis (module-writing)

**Overall:** A practical deep-dive into two important applications: writing and data analysis. Strongest on the data analysis page, which is genuinely useful for non-technical learners. The verification traps page is excellent — directly addresses a real risk. Some content duplicated from Module 5 (sectional drafting, EchoWriting, multi-document synthesis).
**Rating:** Good

#### Page 1: Writing with AI: Beyond "Write Me a Report" (no H2 sections)
- **Content:** The Nature survey stat (57% of scientists use AI writing help) grounds the discussion. The sectional drafting method is well-explained but duplicates Module 5 page 4. The Stanford HAI research reference adds authority. The 9-step human-AI workflow is clear and practical. The FoxPrint Editorial warning callout about AI stripping voice is powerful.
- **Interactive elements:** 1 quiz (why sectional drafting works — good), 1 callout (warning about voice loss).
- **Issues:** Sectional drafting method duplicated from module-files page 4. The "Where AI Fits" section (excels at / struggles with) is useful but basic.
- **Suggestions:** Remove the sectional drafting explanation from module-files and keep it here where it's more detailed.

#### Page 2: Making AI Write Like You (no H2 sections)
- **Content:** Style cards and EchoWriting technique covered in detail. The 3-step practical approach (provide samples → analyse style → write) is clear. Multi-document synthesis section duplicates Module 5 page 5. The "Style vs Substance" closing distinction is important.
- **Interactive elements:** 1 quiz (EchoWriting technique — good), 2 callouts (tip + warning).
- **Issues:** Multi-document synthesis content duplicated from Module 5. NotebookLM mentioned in both places.
- **Suggestions:** Consolidate multi-document synthesis into one location.

#### Page 3: Data Analysis Without a Data Team (no H2 sections)
- **Content:** The strongest page in this module. The "three tiers" framework (add-ins → chat-to-SQL → full-stack analysts) is a useful mental model. Specific tool names (GPTExcel, Numerous.ai, Julius AI, Anomaly AI, Quadratic) make it actionable. The upload-and-analyse workflow is step-by-step practical. Data cleaning section covers the most time-consuming part honestly. Context window trap section adds important nuance.
- **Interactive elements:** 1 quiz (why describe first — good), 1 callout (tip about cleaning).
- **Issues:** Some tool names may become outdated quickly. Consider noting the pace of change.
- **Suggestions:** Add a caveat: "Tool names and capabilities change rapidly; the patterns matter more than the specific products."

#### Page 4: What Can Go Wrong with Data and Documents (no H2 sections)
- **Content:** Excellent page. The WEF quote about "sophisticated wrongness" is a powerful framing. The distinction between AI generating charts and AI understanding charts is an insight many users miss. Document synthesis traps and fabricated citation warnings are directly actionable. The verification checklist (data analysis + document synthesis + both) is practical and comprehensive. The 47% stat about decisions based on hallucinated content is sobering.
- **Interactive elements:** 1 quiz (sophisticated wrongness meaning — good), 2 callouts (warning + info).
- **Issues:** The 4.3 hours/week fact-checking stat appears here AND in Module 2 (hallucinations page) AND in Module 10 (five themes page). Triple usage is borderline.
- **Suggestions:** Choose one location for the 4.3 hours stat and reference it briefly elsewhere.

#### Page 5: Key Takeaways (no H2 sections)
- **Content:** Good summary that connects forward to delegation. The four key points (sectional drafting, style transfer, data workflow, verification) are the right ones. Practical next steps for both writing and data analysis are immediately actionable.
- **Interactive elements:** 1 quiz (central insight — good), 1 callout (info).
- **Issues:** None significant.

#### Module-Level Notes
- 5 pages, 45m estimated — slightly tight for the data analysis content but reasonable.
- Good quiz distribution (1 per page, 5 total).
- Significant content overlap with Module 5: sectional drafting, EchoWriting, multi-document synthesis all appear in both modules.
- No agent demos, no diagrams, no images. Entirely text-based. Data analysis page especially could benefit from a screenshot or workflow diagram.
- British English consistent ("analysing", "summarising", "visualisations").

### Module 7: Delegation & The AI-First Philosophy (module-delegation)

**Overall:** The module that transforms the course from "how to use AI" to "how to work differently." The mindset shift from search engine to worker is compelling. Preparation-as-value is an important reframe. The "When NOT to Use AI" page is excellent and unusual — most AI training doesn't cover limits honestly. The skill atrophy and security content on page 5 feels misplaced — it belongs in Module 8 (Risks).
**Rating:** Good

#### Page 1: The Mindset Shift (3 slides)
- **Content:** Opens with a provocatively honest claim ("AI is already better than you at most routine knowledge tasks") then immediately provides the nuancing METR study (19% slower despite feeling 20% faster). This is excellent pedagogy — bold claim, then data-driven complexity. The Harvard entrepreneur study adds further nuance. The two mental models (search engine vs worker) are a clear dichotomy.
- **Interactive elements:** 1 quiz (search vs worker difference — good), 1 callout (info), 1 SVG image (mindset shift diagram — renders correctly).
- **Issues:**
  - The METR study claim ("16 experienced open-source developers") — verify this sample size and methodology. A study of 16 people making broad claims needs acknowledgment of its limitations.
  - The Harvard study of "640 entrepreneurs" — verify this study exists and is accurately summarised.
- **Suggestions:** Add brief methodological context for the METR study: "Though the sample was small (16 developers)..."

#### Page 2: Delegation as a Skill (no H2 sections)
- **Content:** The process-vs-outcome distinction is the core insight and it's well-illustrated with the sales data example. The "Signs You're Micromanaging" / "Signs You're Delegating Well" lists are practical self-diagnostics. The callout about division of labour is clear.
- **Interactive elements:** 1 quiz (problem with step-by-step — good), 1 callout (tip).
- **Issues:** None significant. This is a clean, focused page.

#### Page 3: Preparation Is Your Value (no H2 sections)
- **Content:** Strong reframe: "execution is cheap, preparation is where value lives." The before/after table showing how activities shift in importance is excellent. The McKinsey finding about "fifteen iterations in two days beats two iterations in five days" adds nuance — preparation matters but so does iteration speed. Good balance.
- **Interactive elements:** 1 quiz (why preparation matters — good), 1 callout (warning about skipping prep).
- **Issues:** The McKinsey "Technology delivers 20% of value; redesigning work delivers 80%" claim needs a citation.
- **Suggestions:** Add source for the McKinsey claim.

#### Page 4: Quality Verification (no H2 sections)
- **Content:** The 4-step verification framework (spot-check, coherence, structure, purpose) is practical and memorable. "You're not checking if AI is right; you're checking if the output is useful" is a good reframe. The "use AI to help verify" advice and the 17% error reduction stat are practical. The WEF quote reappears from Module 6.
- **Interactive elements:** 1 quiz (verification for non-experts — good), 1 callout (tip).
- **Issues:** The "reduces errors by roughly 17%" stat appears without citation. The WEF quote is used here and in Module 6 — some repetition.
- **Suggestions:** Add a source for the 17% figure.

#### Page 5: When NOT to Use AI (no H2 sections)
- **Content:** This is the most important page in the module. The three boundaries (judgment, relationships, accountability) are clearly defined with concrete examples. The complementary pattern table is a useful reference. HOWEVER: the second half of this page shifts to skill atrophy and security risks, which are Module 8's territory. The Microsoft/Carnegie Mellon study, MIT Media Lab research, Polish polyp study, and IBM breach cost data all properly belong in Module 8. This makes the page feel overloaded and unfocused.
- **Interactive elements:** 1 quiz (accountability — good), 2 callouts (warning + danger).
- **Issues:**
  - **P1: Content misplacement.** The skill atrophy section (Microsoft/Carnegie Mellon, MIT Media Lab, Polish polyp study) and digital insider risk section (38% sharing data, IBM breach cost, Gartner prediction) should be in Module 8, not here. Module 8 then re-covers this exact content, creating significant duplication.
  - **P1: Currency inconsistency.** IBM breach cost is "$4.63 million" here but "£4.63 million" in Module 8 page 4. Pick one currency and be consistent.
- **Suggestions:** Move the skill atrophy and security content to Module 8 where it properly belongs. Keep this page focused tightly on the three boundaries (judgment, relationships, accountability).

#### Page 6: Key Takeaways: Delegation (no H2 sections)
- **Content:** Good summary with 4 clear points and a practical delegation checklist. The checklist format (checkbox items) is a nice touch. Smooth transition to "Advanced Patterns."
- **Interactive elements:** 1 quiz (most important preparation — good), 1 callout (tip).
- **Issues:** The "Up Next" says "Advanced Patterns" but the actual next module in course.yaml is module-risks, not module-advanced. This suggests the module ordering was changed after the content was written, or the transition text is wrong.
- **Suggestions:** Fix the "Up Next" text to reference the actual next module (Risks).

#### Module-Level Notes
- 6 pages, 60m estimated — seems right.
- Good quiz distribution (1 per page, 6 total).
- The three boundaries page is the highlight — honest, practical, memorable.
- Significant content duplication with Module 8 on skill atrophy and security.
- No agent demos, 1 SVG diagram. Light on visual variety.
- British English consistent ("organisation", "behaviour").

### Module 8: Risks, Responsibility, and Realistic Expectations (module-risks)

**Overall:** The most important module for responsible AI use. Opens with honest data about AI underperformance (95% of organisations, METR study). Hallucination depth, skill atrophy, privacy, and legal coverage are thorough. BUT: contains serious content issues — a potentially fabricated Anthropic settlement claim, significant content duplication with Module 7, and a currency inconsistency on the IBM breach cost.
**Rating:** Needs Work

#### Page 1: The Honest Assessment (4 slides)
- **Content:** Bold opening with uncomfortable statistics (95% no returns, 1% maturity, 67% failure within 6 months). The METR study and Harvard entrepreneur study are reused from Module 7 page 1. The "What Actually Works" section provides balance — this isn't doom-and-gloom, it's calibration.
- **Interactive elements:** 1 quiz (cognitive bias — good), 1 callout (warning).
- **Issues:**
  - **Content duplication.** The METR study (19% slower) and Harvard study (no statistical difference) appear nearly verbatim from Module 7 page 1. Students will notice.
  - The MIT Media Lab "95% of organisations see no measurable returns" stat — verify this claim and source.
- **Suggestions:** In Module 7, introduce the studies briefly for the delegation argument. In Module 8, provide the full detailed treatment. Don't repeat the same paragraphs.

#### Page 2: Hallucination: The Numbers (no H2 sections)
- **Content:** Excellent deep-dive. Specific rates (Gemini-2.0-Flash 0.7%, o3 33%, o4-mini 48%) give students concrete benchmarks. Stanford legal query rates (58-82%) show domain-specific risk. The training incentive explanation ("rewards guessing over acknowledging uncertainty") is clear. Countermeasures (cross-validation, self-verification, web search, RAG) are practical.
- **Interactive elements:** 1 quiz (training incentives — good), 2 callouts (danger + tip).
- **Issues:** Much of this content appeared in Module 2 (Hallucinations page). The overlap is significant: PersonQA stats, 47% enterprise users, 4.3 hours fact-checking. Module 8 adds more detail, which is appropriate for a deeper treatment, but the redundancy should be acknowledged.
- **Suggestions:** In Module 2, cover hallucinations conceptually. In Module 8, cover the numbers and implications. Add a forward reference in Module 2: "We'll dig into the specific numbers in the Risks module."

#### Page 3: The Skill Atrophy Problem (no H2 sections)
- **Content:** Strong, evidence-based page. Microsoft/Carnegie Mellon, MIT Media Lab, and the Polish polyp study (detection rate fell from 28.4% to 22.4%) paint a clear picture. The automation paradox callout is well-framed. Countermeasures (AI-free zones, deliberate practice, baseline assessment) are practical.
- **Interactive elements:** 1 quiz (automation paradox — good), 2 callouts (warning + tip).
- **Issues:**
  - **P1: Content duplication with Module 7 page 5.** The Microsoft/Carnegie Mellon study, MIT Media Lab research, and Polish polyp study all appear nearly identically in Module 7's "When NOT to Use AI" page. This is clearly the right home for this content; it should be removed from Module 7.
  - The MIT Media Lab claim about "reduced brain activity, diminished memory retention, and less original thinking" — verify this specific study exists.
- **Suggestions:** Remove the skill atrophy content from Module 7 page 5. Keep this dedicated page as the authoritative treatment.

#### Page 4: Privacy, Security, and the Digital Insider (no H2 sections)
- **Content:** Important and well-structured. The 38% sharing data stat, IBM breach cost, and Gartner prediction ground the discussion in business reality. The agentic AI risk section (agents inherit file permissions) is an insight many students won't have considered. McKinsey's "digital insiders" framework is practical. Consumer vs enterprise tier comparison is useful.
- **Interactive elements:** 1 quiz (agentic vs chatbot risk — good), 2 callouts (danger + tip).
- **Issues:**
  - **P0: Currency inconsistency.** IBM breach cost is "£4.63 million" here but "$4.63 million" in Module 7 page 5. The original IBM report uses USD ($4.88M average, with AI-related breaches higher). Using GBP (£) without conversion context is misleading. Pick USD (the source currency) or convert consistently.
  - **Content duplication with Module 7 page 5.** The 38% stat, IBM cost, and Gartner prediction all appear in both modules.
- **Suggestions:** Use USD for IBM data (it's the source currency) or explicitly note the conversion. Remove the security content from Module 7.

#### Page 5: Legal and Copyright: What You Must Know (7 slides)
- **Content:** Important and practical. US Copyright Office ruling coverage is clear. International divergence section (Japan/Singapore permissive, France/EU restrictive) is useful for global organisations. Practical implications for content creation, reports, and code generation are actionable. Liability section is clear.
- **Interactive elements:** 1 quiz (copyright protection — good), 2 callouts (warning + tip).
- **Issues:**
  - **P0: Potentially fabricated Anthropic settlement.** "In June 2025, Anthropic paid a £1.5 billion settlement for using pirated books to train Claude" and "This is the largest copyright payout in US history." I cannot verify this claim. As of my knowledge, no such settlement has occurred. If this is fabricated, it is exactly the kind of hallucinated "confident wrongness" the course warns about — ironic and deeply damaging to credibility. **This MUST be verified before delivery.**
  - **P1: Currency confusion.** The settlement is stated in GBP (£1.5 billion) but described as "the largest copyright payout in US history." US legal settlements are denominated in USD. This inconsistency suggests the claim may not have been properly sourced.
  - France fine stated as "£250 million" — was this actually euros (€)?
- **Suggestions:** VERIFY the Anthropic settlement claim immediately. If it cannot be verified, REMOVE it entirely. Replace with a verifiable example (e.g., the New York Times v. OpenAI lawsuit, which is documented). Fix currency to match source denomination.

#### Page 6: Key Takeaways (no H2 sections)
- **Content:** Good summary with 5 critical realities. Each maps to a specific module page. The "What Changes on Monday" section connects back to practical skills learned throughout the course. Smooth transition to Advanced Patterns.
- **Interactive elements:** 1 quiz (most important factor — good), 1 callout (info).
- **Issues:** None significant beyond the module-level issues.

#### Module-Level Notes
- 6 pages, 45m estimated — may be tight given the density of the content.
- Good quiz distribution (1 per page, 6 total).
- Contains the course's most serious content issues: the Anthropic settlement claim and currency inconsistencies.
- Significant overlap with Module 7 on skill atrophy, security, METR/Harvard studies.
- No agent demos, no diagrams. Heavy text module that could benefit from a risk matrix or comparison visual.
- British English mostly consistent ("organisations", "authorisation") but GBP/USD confusion undermines professionalism.

### Module 9: Advanced Patterns (module-advanced)

**Overall:** The power-user module. Covers skills, subagents, persistent memory, code generation, and MCP — the five pillars of advanced AI use. The meta case study ("this course built itself") is a brilliant pedagogical device that demonstrates every concept simultaneously. The MCP page with the Salesforce Agentforce vulnerability is especially strong. Longest module alongside prompting (7 pages, 60m).
**Rating:** Good

#### Page 1: Building Systems, Not Just Prompts (4 slides)
- **Content:** Good framing: casual users → effective users. The three pillars (Skills, Subagents, Connectors/MCP) provide the module's structure. Token economics section is important for enterprise context — the 100x consumption for agents, pricing trajectory from $20 to $0.40 per million tokens. Cost optimisation strategies (caching, cascading, batching, fine-tuning) are practical.
- **Interactive elements:** 2 quizzes (systems vs prompts — good; token economics — good), 2 callouts (info + note), 1 SVG image (three pillars — renders correctly).
- **Issues:** None significant. Clean intro page.

#### Page 2: Skills: Reusable Procedures (no H2 sections)
- **Content:** Clear explanation of skills as "runbooks for AI." The anatomy breakdown (Purpose, Steps, Tone, Output) is practical. Good candidates list is helpful. The meta-tooling section (tools that make tools) is a powerful concept. Google Workspace Studio, Claude Code Skills, and MindStudio are specific and actionable examples.
- **Interactive elements:** 1 quiz (skill benefit — good), 2 callouts (tip + tip).
- **Issues:** None significant.

#### Page 3: Subagents: AI Spawns Helpers (very long page, no H2-based slides)
- **Content:** This is the standout page of the module. The subagent concept is explained simply, then the real-world patterns (ReAct, Plan-then-Execute, Hierarchical Task Decomposition, Generator-Evaluator) add depth. The "$4.35 billion market" stat grounds the discussion. BUT the real star is the meta case study: "This Course Built Itself." It walks through 5 phases of how the course was created using subagents, showing every concept from the training in action. The self-referential moment is genuinely powerful.
- **Interactive elements:** 2 quizzes (subagent benefit — good; meta case study concepts — good), 2 callouts (info + info), 2 images (subagent architecture SVG, course creation screenshot).
- **Issues:**
  - The "meta-course-creation-screenshot.png" — verify this image exists and renders. It's referenced from `/static/images/`.
  - The $4.35 billion autonomous agents market stat needs a source.
  - This page is very long for a single page with no slide breaks. Students in a live setting may lose focus.
- **Suggestions:** Consider splitting this page into two: "How Subagents Work" and "Case Study: This Course Built Itself." The case study deserves its own page given its importance.

#### Page 4: Persistent Memory (no H2 sections)
- **Content:** Clear problem/solution framing (every conversation starts fresh → persistent instructions fix this). The landscape section (CLAUDE.md, AGENTS.md, Claude Projects, ChatGPT Custom Instructions) is comprehensive. CLAUDE.md best practices (under 300 lines, progressive disclosure, commit to version control) are specific and actionable. The example persistent instructions block is realistic. Both developer and non-developer paths covered.
- **Interactive elements:** 2 quizzes (persistent memory benefit — good; highest leverage — good), 2 callouts (info + tip).
- **Issues:** The "60,000+ open-source projects" adopting AGENTS.md — verify this stat.
- **Suggestions:** None major. Well-structured page.

#### Page 5: Code Generation for Non-Coders (no H2 sections)
- **Content:** Good coverage of vibe coding. Karpathy attribution and etymology are interesting. Market stats (Y Combinator 25%, Gartner 75%, Citrix 4,500-6,000 apps) ground the trend. Practical applications section (data processing, file automation, visualisation) is actionable. The Stack Overflow realism check is honest. The "What Works / What Has Limitations" distinction is important.
- **Interactive elements:** 2 quizzes (code gen meaning — good; vibe coding scope — good), 2 callouts (info + info).
- **Issues:**
  - "Cursor: $500M ARR in June 2025, up from $1M twelve months earlier" — verify this dramatic claim.
  - "Collins Dictionary named it Word of the Year" — verify this is accurate (vibe coding as Word of the Year).
- **Suggestions:** Add brief citations for the Cursor ARR and Collins Dictionary claims.

#### Page 6: MCP: Connecting AI to Your Tools (no H2 sections)
- **Content:** Excellent page. The "USB-C for AI" analogy is perfect for this audience. Adoption stats (17,000+ servers, 97M SDK downloads, OpenAI/Google adoption, Linux Foundation donation) establish authority. The server examples (Drive, Slack, Notion, GitHub, Figma, PostgreSQL) make it concrete. The Salesforce Agentforce vulnerability (CVSS 9.4, $5 domain, CRM data exfiltration) is a compelling cautionary tale. Security best practices are practical.
- **Interactive elements:** 2 quizzes (MCP purpose — good; Agentforce vulnerability — good), 2 callouts (warning + info), 1 SVG image (MCP connections — renders correctly).
- **Issues:**
  - The 17,000+ MCP servers and 97M SDK downloads — verify these stats.
  - "Donated to Linux Foundation's Agentic AI Foundation (December 2025)" — verify this happened.
- **Suggestions:** None major. Strong page.

#### Page 7: Key Takeaways: Advanced Patterns (no H2 sections)
- **Content:** Clean summary with 5 points mapping to the 5 pages. The transition table (Beginner → Intermediate → Advanced) is a useful self-assessment. "Build incrementally" advice is practical. Smooth transition to synthesis.
- **Interactive elements:** 1 quiz (core principle — good), 1 callout (tip).
- **Issues:** None significant.

#### Module-Level Notes
- 7 pages, 60m estimated — may be tight given the density, especially the subagents case study.
- Good quiz distribution (10 total quizzes across 7 pages — some pages have 2).
- The meta case study is the course's most memorable moment — the self-referential proof of every concept.
- MCP and subagents pages are the strongest.
- 3 SVG diagrams — better visual variety than most modules.
- British English consistent ("parallelise", "specialisation", "optimisation").

### Module 10: Putting It Together (module-synthesis)

**Overall:** Effective synthesis that ties the course together. The five themes are the right distillation. The "quick wins" and "ten mistakes" sections are immediately actionable takeaways. The reference page with glossary and model comparison is a useful leave-behind. Feels slightly rushed at 4 pages / 45m compared to the depth of earlier modules.
**Rating:** Good

#### Page 1: Your AI-First Workflow (no H2 sections)
- **Content:** The core cycle (Prepare → Delegate → Verify → Deliver) is the right synthesis. McKinsey, PwC, and Bain citations lend authority. The 4-section breakdown with clear bullet points under each phase is practical. The "3x more likely to have redesigned workflows" stat is compelling.
- **Interactive elements:** 1 quiz (primary value phase — good), 2 callouts (info + tip), 1 SVG image (core cycle — renders correctly).
- **Issues:** None significant. Clean synthesis page.

#### Page 2: The Five Themes (no H2 sections)
- **Content:** The five themes (Outcomes over process, Context is everything, AI-first human-verified, Preparation is the new execution, Files not chat) are the right distillation of the entire course. Each is expressed as a principle with a practical implication. The "Verification Is Non-Negotiable" callout reuses the 4.3 hours and 47% stats for the third time.
- **Interactive elements:** 1 quiz (poor quality → which theme — good), 2 callouts (warning + info).
- **Issues:** The 4.3 hours/week and 47% enterprise stats now appear in Module 2, Module 6, and here. Triple usage dilutes impact.
- **Suggestions:** In synthesis, say "As we've seen..." rather than presenting these stats as if new.

#### Page 3: What To Do Next (no H2 sections)
- **Content:** The five quick wins (meeting summarisation, email drafting, document summarisation, status reports, research acceleration) are the right starting points. The ten mistakes list is comprehensive and practical. The action plan (this week / this month / ongoing) gives concrete next steps. Common objections section addresses real resistance. "Daily GenAI users save 4+ hours per week" stat is motivating.
- **Interactive elements:** 2 quizzes (time saved — good; common mistake — good), 2 callouts (info + warning + tip).
- **Issues:**
  - "92% productivity improvement" — this is a vague and hard-to-verify claim. What does "92% productivity improvement" mean precisely?
  - "Higher salaries" for daily GenAI users — correlation/causation issue. Early adopters may already be higher performers.
  - "4.3 hours per week fact-checking" stat appears AGAIN (fourth usage).
- **Suggestions:** Qualify the "92% productivity improvement" with a source and what was measured.

#### Page 4: Quick Reference (no H2 sections)
- **Content:** Useful leave-behind page. Glossary of 10 key terms is the right set. The Three Questions, Basic Prompt Structure, Verification Checklist, Don't Delegate list, Five Themes, and Workflow diagram are all summarised compactly. The model comparison (Claude, GPT, Gemini) is balanced and explicitly not a product comparison.
- **Interactive elements:** 1 quiz (fundamental nature of LLMs — good, full-circle callback to Module 2).
- **Issues:** "Go build something." as the closing line is effective but slightly informal for a premium training product. Consider whether this matches the desired tone.
- **Suggestions:** None major. This is a good reference page.

#### Module-Level Notes
- 4 pages, 45m estimated — feels right as a lighter wrap-up module.
- Good quiz distribution (5 total quizzes).
- The five themes and ten mistakes are the most immediately usable takeaways.
- Reference page is a valuable leave-behind.
- 1 SVG diagram. No agent demos.
- British English consistent ("summarisation", "organisation").

## Cross-Module Observations

### Strengths
- **Strong conceptual framework.** The five themes (outcomes, context, AI-first, preparation, files) form a coherent philosophy that threads through every module.
- **Honest about limitations.** The course doesn't oversell AI. METR study, hallucination rates, skill atrophy data — this is unusually balanced for AI training.
- **Practical over theoretical.** Every concept has concrete examples, specific tool names, and actionable techniques. The Three Questions framework and verification checklist are immediately usable.
- **The meta case study.** "This course built itself" in Module 9 is a brilliant pedagogical device — every concept demonstrated in one real example.
- **Quiz quality is generally high.** Most quizzes test understanding, not recall. Explanations teach rather than just confirm.
- **Agent demo in Module 1.** The opening demo viscerally demonstrates the thesis before any theory. Show, don't tell.

### Issues

#### Content Duplication (Systematic Problem)
The following content appears in multiple modules and needs consolidation:

| Content | Appears In | Should Live In |
|---------|-----------|---------------|
| Sectional Drafting Method | module-files p4, module-writing p1 | module-writing p1 only |
| EchoWriting technique | module-files p5, module-writing p2 | module-writing p2 (detailed), module-files p5 (brief mention) |
| Multi-document synthesis | module-files p5, module-writing p2 | module-writing p2 only |
| METR study (19% slower) | module-delegation p1, module-risks p1 | module-risks p1 (full), module-delegation p1 (brief reference) |
| Harvard entrepreneur study | module-delegation p1, module-risks p1 | module-risks p1 (full), module-delegation p1 (brief reference) |
| Skill atrophy data | module-delegation p5, module-risks p3 | module-risks p3 only |
| Security/insider risk data | module-delegation p5, module-risks p4 | module-risks p4 only |
| 4.3 hours/week fact-checking | module-llms p3, module-writing p4, module-delegation p4, module-synthesis p2-3 | module-risks p2 (primary), others brief reference |
| 47% enterprise decisions on hallucinated content | module-llms p3, module-writing p4, module-synthesis p2 | module-risks p2 (primary), others reference |

This is the course's biggest structural issue. Some reinforcement of key stats is fine. Verbatim duplication across 3-4 modules is not.

#### Currency Inconsistency
- IBM breach cost: "$4.63 million" (module-delegation p5) vs "£4.63 million" (module-risks p4). The source data is in USD.
- France Google fine: "£250 million" (module-risks p5) — likely should be euros (€).
- Anthropic settlement: "£1.5 billion" described as "the largest copyright payout in US history" — US settlements are in USD.

#### Factual Claims Requiring Verification
These claims could not be verified and may be fabricated or inaccurate:

| Claim | Location | Concern |
|-------|----------|---------|
| Anthropic £1.5B settlement (June 2025) | module-risks p5 | **CRITICAL.** Cannot verify. Likely hallucinated. |
| "95% of organisations see no measurable returns" (MIT Media Lab) | module-risks p1 | Verify source. MIT Media Lab is plausible but the stat is suspiciously round. |
| "Collins Dictionary named [vibe coding] Word of the Year" | module-advanced p5 | Verify. Collins did name AI-related terms in recent years. |
| "60,000+ open-source projects" adopting AGENTS.md | module-advanced p4 | Verify. Large number for a standard proposed in August 2025. |
| "Cursor: $500M ARR in June 2025, up from $1M twelve months earlier" | module-advanced p5 | Verify. 500x growth in 12 months is extraordinary. |
| "97 million monthly SDK downloads" for MCP | module-advanced p6 | Verify. |
| Polish polyp study (1,443 patients, 28.4% → 22.4%) | module-risks p3, module-delegation p5 | Verify study exists. |
| METR study (16 developers, 19% slower) | module-delegation p1, module-risks p1 | Study is real but verify sample size and methodology. |

### Content Gaps
- **No coverage of AI for presentation/visual work.** File module mentions PowerPoints but there's no guidance on AI + design.
- **Limited coverage of team collaboration with AI.** Most content is individual workflow. How do teams share skills, persistent instructions, or coordinate agent use?
- **No discussion of model selection in depth.** Module 10's reference page has a brief comparison but there's no page dedicated to "when to use which model."
- **No coverage of prompt libraries or community resources.** Where do students go to find good prompts and learn from others?

### Interactive Element Audit
- **Quizzes:** 52 total across 10 modules. Generally high quality. Most test understanding, not recall. A few have weak distractors (Module 1 page 2 quiz is too easy). Distribution is even — every page has at least one.
- **Callouts:** ~45 total. Generally appropriate use. Module 3 page 4 has 4 callouts (borderline overuse). Types are well-matched to content (danger for hallucinations, tips for techniques, warnings for risks).
- **Agent Demos:** 1 total (Module 1 page 3). This is LOW for a premium training product. The demo is excellent — consider adding more in other modules (prompting techniques, data analysis workflow, delegation pattern).
- **Images/Diagrams:** ~12 SVGs across the course. Most modules have 1-2. Modules 5, 6, 7, 8 are entirely text-based with no diagrams. More visual variety needed.
- **Pages with no interactive elements beyond quizzes:** Many. Most pages are just text + quiz + 1-2 callouts. No polls, no drag-and-drop, no interactive exercises.

## Priority Fixes

### P0 — Must Fix (blocks learning or is factually wrong)
1. **Verify/remove Anthropic £1.5B settlement claim** (module-risks p5). If this is fabricated, it is catastrophic for course credibility — a hallucinated fact in a module about hallucination risk. Replace with a verifiable legal case if unconfirmed.
2. **Fix currency inconsistency on IBM breach cost** (module-delegation p5 vs module-risks p4). Use USD ($4.63M) consistently since that's the source data currency, or convert and note the conversion.
3. **Fix France fine currency** (module-risks p5). Should be euros (€250 million), not pounds (£250 million).

### P1 — Should Fix (degrades quality noticeably)
1. **Consolidate duplicated content** across modules. The sectional drafting, EchoWriting, METR study, skill atrophy data, and security data all appear in 2-3 places. Keep the authoritative treatment in one module; use brief references elsewhere.
2. **Fix stale UI copy on Module 1 page 3** (agent demo intro). References "sidebar on the right" and "scratchpad" — the current UI has a left file explorer and no scratchpad label.
3. **Move skill atrophy and security content from Module 7 page 5 to Module 8** where it belongs. Module 7 page 5 should focus on the three boundaries (judgment, relationships, accountability).
4. **Fix Module 7 page 6 "Up Next" text** — says "Advanced Patterns" but the next module is Risks.
5. **Verify the ~15 factual claims flagged in the verification table above.** Any unverifiable claim should be removed or replaced with a citable source.
6. **Fix "behavior" → "behaviour"** on module-context page 1 (line 56 of source).

### P2 — Nice to Have (polish and refinement)
1. **Add more agent demos.** Currently only 1 in the entire course. Add demos for: prompting iteration (Module 4), data analysis workflow (Module 6), delegation pattern (Module 7).
2. **Add more diagrams/visuals** to text-heavy modules (5, 6, 7, 8). Every module should have at least 2 visual elements beyond callouts.
3. **Reduce slide count on heavy pages.** Module 2 page 4 (7 slides for takeaways), Module 3 page 4 (7 slides), Module 4 page 6 (8 slides) — condense to 4-5 each.
4. **Merge thin intro slides** on takeaway pages (Module 2, 3, 4 key takeaways all have a near-empty first slide).
5. **Strengthen weak quiz distractors** — Module 1 page 2 quiz is too easy (obvious wrong answers).
6. **Add date qualifiers to capability claims** (ChatGPT file limits, Claude features, model comparisons) — these will date quickly.
7. **Qualify the "92% productivity improvement" stat** (Module 10 page 3) with a source and measurement methodology.
8. **Consider adding a "Model Selection" page** — currently covered only briefly in the Module 10 reference page.

## Statistics

| Metric | Count |
|--------|-------|
| Total modules | 10 |
| Total pages | 53 |
| Total quizzes | ~52 |
| Total callouts | ~45 |
| Total agent demos | 1 |
| Total images/diagrams (SVG) | ~12 |
| Pages with no interactive elements beyond quiz | ~15 |
| Broken images | 0 confirmed (meta-course-creation-screenshot.png not verified) |
| Estimated total duration | 8h 40m (sum of module estimates: 30+60+60+85+60+45+60+45+60+45) |
| Content duplication instances | 9 significant |
| Unverified factual claims | ~15 |
| Currency inconsistencies | 3 |
| American English instances found | 1 ("behavior" in module-context) |
