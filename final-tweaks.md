# Final Tweaks — Action Plan

Based on human review of the complete course. Each section maps to a module-builder subagent workstream.

---

## Workstream 1: Modules 1, 2, 6, 9 — Audit Only

These modules have no structural changes, just demo + quiz quality checks.

### Demo Audit (all 4 modules)
Review every agent demo block. For each, check:
- [ ] First user prompt is NOT front-loaded with tonnes of detail — context should already be in scratchpad files
- [ ] Research happens first where relevant (e.g. web_search before writing)
- [ ] Content is built up via intermediate files, not one giant prompt→response
- [ ] Output goes to files, not dumped as huge assistant messages
- [ ] Files are organised in folders where it makes sense (not everything flat)
- Fix any demos that violate these principles by editing the YAML script

### Quiz Audit (all 4 modules)
Review every quiz block. For each, check:
- [ ] Exactly 3 options (not 4, not 2 — the course standard is 3)
- [ ] Options are genuinely difficult to distinguish — no obviously wrong answers
- [ ] Questions test understanding, not recall of definitions
- Fix any quizzes that don't meet these standards

**Modules:** module-opening, module-llms, module-writing, module-risks

---

## Workstream 2: Module 3 (Context) — Tools Introduction + Frustration Framing

### A1. Introduce tools early
Add a brief section to an appropriate page (probably `01-everything-is-context.md` or `03-context-hierarchy.md`) explaining that AI assistants can use **tools** — search the web, read/write files, run code, browse the web. Keep it short (1-2 paragraphs + a callout). The point is: by the time learners see tools in demos, they know what they are. Don't go deep — Module 7 covers the mechanics.

### A5. Frustration framing
Add to `04-context-decay.md`: when things go wrong with AI, it's almost always a context problem. Frame the tips in this module as ways to avoid the #1 source of frustration.

### Demo + Quiz Audit
Same checks as Workstream 1 for this module's demos and quizzes.

---

## Workstream 3: Module 4 (Prompting) — Interactive Speccing Demo + Frustration

### A8. Interactive speccing demo
Page `05-meta-prompting.md` explains interactive speccing well but has no demo. Add an agent demo showing:
- User wants to plan a company offsite event
- Instead of trying to spec everything upfront, user asks AI to interview them
- AI asks clarifying questions (budget? team size? goals? constraints?)
- User answers, AI builds the spec incrementally
- Final output: a structured spec document written to a file
This demonstrates the back-and-forth pattern that leads to better specs than trying to write everything yourself.

### A5. Frustration framing
Add to `06-iteration.md`: frame iteration tips as frustration-avoidance. When the AI gives you rubbish, it's because you didn't iterate — not because the AI is stupid.

### A10. Check "What do you think?" tip
Verify the tip in `06-iteration.md` is prominent. If it's buried in a paragraph, elevate it to a callout or its own H2 section.

### Demo + Quiz Audit
Same checks as Workstream 1.

---

## Workstream 4: Module 5 (Files) — Plan-First Pattern

### A7. Plan-first, write-to-file, follow-the-plan
Add content + agent demo showing the plan-first workflow. Could go in `05-productive-spiral.md` or a new section. The demo should show:
- User asks AI to reorganise a messy project
- AI's first move is to write a plan to `plan.md`
- User reviews and approves the plan
- AI executes step by step, referring back to the plan
- Final result is well-organised with the plan as documentation
This is THE key workflow pattern for non-trivial tasks.

### Demo + Quiz Audit
Same checks as Workstream 1.

---

## Workstream 5: Module 7 (Advanced) — AI Mistakes + Explicit Tool Use + Infographics

### A2. AI doesn't learn from its mistakes
Add content to `05-auto-memory.md` or `04-persistent-memory.md`. Key teaching: AI won't spontaneously learn from errors within a conversation. If it makes a mistake and you correct it, it won't remember that correction in the next conversation. YOU must proactively tell it to update its memory/skills/context files. This is a common source of frustration.

### A11. Explicit vs automatic tool/skill use
Add a brief section to `02-skills.md` or `03-subagents.md`. Sometimes AI automatically uses the right tool. Sometimes you have to explicitly say "use the web search tool" or "run the /review skill". Knowing when to be explicit is itself a skill. Rule of thumb: if AI isn't doing what you expect, try being explicit about which tool/skill to use.

### D. Infographics
Pages 01 (Building Systems), 02 (Skills), 04 (Persistent Memory), 05 (Auto Memory) have no visuals. Generate infographics for at least 2-3 of these using /image-generator.

### Demo + Quiz Audit
Same checks as Workstream 1.

---

## Workstream 6: Module 8 (Delegation) — Multiple Teaching Gaps + Infographic

### A4. Human as bottleneck
Add to `01-mindset-shift.md`. The AI can execute in seconds what used to take hours. The bottleneck has shifted from "can we build this?" to "can we specify this clearly enough?" The human is now the limiting factor — your ability to delegate, provide context, and verify is what determines output quality.

### A6. The failure of the AI is YOUR failure
Add to `02-delegation-skill.md`. Reframe: when AI produces rubbish, it's a delegation failure, not an AI failure. You gave unclear instructions, insufficient context, or didn't verify. Accept this and you'll improve much faster.

### A3. Practice is the only way
Add to `02-delegation-skill.md` or `03-preparation-value.md`. Reading about AI is necessary but nowhere near sufficient. The skill is in the doing — the muscle memory of structuring prompts, providing context, iterating. You can only build this through practice.

### A9. Validation as editorial review
Expand `04-quality-verification.md`. Currently covers fact-checking. Add: getting AI to re-read and critically evaluate its own output (like an editor, not just a fact-checker). The technique: "Now re-read what you just wrote. What would you improve?" Also: for visual outputs (websites, documents), use browser tools to actually look at what was produced.

### D. Infographic
Page `03-preparation-value.md` could use an infographic showing the preparation→delegation→verification cycle.

### Demo + Quiz Audit
Same checks as Workstream 1.

---

## Workstream 7: Module 10 (Synthesis) — Capstone Demo + Practice + Infographics

### E. THE CAPSTONE DEMO
This is the most important single deliverable. Replace or dramatically expand the demo on `03-next-steps.md` (currently a simple meeting summarisation). The new demo must be:

**Duration:** 5+ minutes of scripted interaction (the longest demo in the course)

**Scope:** Demonstrate EVERY major concept taught across all 10 modules:
- Start with context files (Module 3) — provide background files
- Use research tools (Module 5) — search for information
- Plan first, write to file (Module 5/8) — create a plan before executing
- Use structured prompting (Module 4) — clear, well-structured instructions
- Build up via intermediate files (Module 5) — don't try to do everything at once
- Use skills/tools explicitly (Module 7) — invoke specific capabilities
- Iterate on output (Module 4) — refine with specific feedback
- Get AI to self-review (Module 8) — "re-read and evaluate"
- Organise output in folders (Module 5) — clean file structure
- Produce a polished final deliverable (Module 6) — professional output

**Scenario suggestion:** A realistic business task that naturally requires all these skills. For example: "Your CEO has asked you to prepare a competitive analysis and strategy recommendation for entering a new market." This naturally involves:
1. Context files: company background, existing market data
2. Research: web search for competitor info, market trends
3. Planning: write analysis plan to file
4. Structured work: research → analysis → draft → review → final
5. Multiple output files: research notes, draft report, final presentation
6. Self-review: AI re-reads and improves its own analysis
7. Folder organisation: /research, /drafts, /final

The demo should leave people thinking "I want to go do this RIGHT NOW."

### A3. Practice emphasis
Add strong emphasis in `03-next-steps.md` that practice is the only path. Consider a callout: "You've learned the theory. Now the real learning begins."

### D. Infographics
Module 10 currently has ZERO images. Generate at least:
- A workflow diagram for page 01 (the AI-first workflow)
- A five-themes visual for page 02
- Something inspiring for page 03 (next steps)

### Demo + Quiz Audit
Same checks as Workstream 1 for existing content.

---

## Cross-Cutting Standards

### Demo Best Practices (applies to ALL workstreams)
Every agent demo in the course should model excellent AI usage:
1. **Context in files first** — scratchpad should be pre-loaded with relevant context files
2. **Research before writing** — use web_search or file reading before producing output
3. **Plan before executing** — for non-trivial tasks, write a plan first
4. **Intermediate files** — build up content through working files, not one mega-prompt
5. **Output to files** — final output written to files, not just returned as text
6. **Folder organisation** — use folders when there are 3+ related files
7. **Iteration** — show at least one round of user feedback → AI improvement

### Quiz Standards (applies to ALL workstreams)
1. **3 options** per question (not 4)
2. **Genuinely challenging** — all options should be plausible, correct answer requires understanding
3. **No obvious wrong answers** — if a quiz is easy to ace without reading the content, rewrite it
4. **Test comprehension** — "What would happen if..." not "What is the definition of..."
