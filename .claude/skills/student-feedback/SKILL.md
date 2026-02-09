---
name: student-feedback
description: Review the entire course through Chrome as a student would — navigate every page, interact with quizzes and demos, evaluate content quality, and compile a detailed feedback report.
argument-hint: "[optional: specific module to review, or 'full' for entire course]"
---

# Student Feedback Reviewer

You review the training course **as a student would**, navigating through it in Chrome using browser automation tools. You evaluate every aspect of the experience — content quality, interactive elements, navigation flow, visual design — and compile a comprehensive feedback report.

## Before Starting

1. **Get browser context** — call `tabs_context_mcp` to see current tabs.
2. **Create a new tab** — call `tabs_create_mcp` for a fresh session.
3. **Ensure the server is running** — navigate to `http://localhost:8080`. If it doesn't load, tell the user to run `./restart.sh`.
4. **Authenticate as admin** — enter the admin key (`admin-test-key`) at the cohort code screen. This gives full access to all modules without prerequisites blocking you.
5. **Read the course structure** — read `content/course.yaml` to know module order, then read each module's `_module.yaml` for metadata.

## What You Produce

A detailed markdown feedback report written **incrementally** to `course-feedback-report.md` in the project root.

## Incremental Writing & Resumption

**This is critical.** The review is long and may be interrupted by context limits. You MUST write to the report file as you go, not at the end.

### Writing Strategy

1. **Before starting**, check if `course-feedback-report.md` already exists. If it does, read it — a previous run was interrupted and you should **resume from where it left off**.
2. **Write the report header and skeleton immediately** after Phase 1 (module overview). This includes the frontmatter, executive summary placeholder, and all module section headers.
3. **After completing each module's review**, append that module's section to the report file. Do not wait until all modules are done.
4. **After completing all modules**, write the Cross-Module Observations, Priority Fixes, and Statistics sections. Update the Executive Summary with final insights.

### Progress Tracking

Include a `<!-- PROGRESS -->` comment at the top of the report that tracks where you are:

```markdown
<!-- PROGRESS: module-context page 3 of 5 -->
```

Update this comment each time you write to the file. When resuming:
- Read the existing report
- Parse the `<!-- PROGRESS -->` comment to find where you stopped
- Skip modules/pages already reviewed (their sections will be in the file)
- Continue from the next unreviewed page

If there is no PROGRESS comment but module sections exist, infer progress from which module headings are present in the file.

### Resume Checklist

When resuming an interrupted review:
1. Read `course-feedback-report.md`
2. Identify the last completed module and page from the PROGRESS comment
3. Get browser context (`tabs_context_mcp`) — you'll need a fresh tab and re-authentication
4. Navigate to the page after the last one reviewed
5. Continue the review, appending to the existing report

## Review Process

### Phase 1: Module Overview

Navigate to the module selector (the `/modules` route — click "View Progress" in the sidebar). For each module card, evaluate:

- **Card presentation** — title, description, difficulty badge, duration estimate
- **Visual consistency** — do all cards follow the same design pattern?
- **Ordering** — does the module sequence make logical sense?
- **Prerequisites** — are prerequisite relationships sensible?

Take a screenshot of the module selector for reference.

### Phase 2: Page-by-Page Review

Go through **every module in course order**, and within each module, **every page in order**. For each page:

#### A. First Impression (before reading)
- Take a screenshot
- Note: Does the page header (title, section indicator) look correct?
- Note: Is the layout clean? Any visual glitches?

#### B. Content Quality
Read the page text carefully (use `get_page_text` or `read_page`). Evaluate:

- **Clarity** — is the concept explained clearly? Would a non-technical professional understand it?
- **Accuracy** — are claims, statistics, and examples factually correct? Flag anything that seems wrong or outdated.
- **Conciseness** — is there unnecessary padding or filler? Every sentence should earn its place.
- **Engagement** — does the writing grab attention? Is it vivid and specific, or generic and bland?
- **Flow** — does the page connect logically to the previous one? Is there a clear narrative thread?
- **Specificity** — are there concrete examples, real data, named tools? Or vague generalities?
- **Tone** — is it professional but approachable? Not condescending, not overly academic?
- **British English** — check for American spellings (behavior→behaviour, organization→organisation, etc.)

#### C. Section/Slide Navigation (pages with H2s)
If the page has multiple sections (slides):

- Navigate through each slide using the Next button
- Check: Does each section feel like a complete thought?
- Check: Is the slide count shown correctly?
- Check: Are section titles shown in the sidebar?
- Check: Does the sidebar correctly highlight the current section?

#### D. Interactive Elements

**Callouts:**
- Are they used appropriately? (not overused, right type for content)
- Is the callout content concise and valuable?
- Does the type match the message? (tip for advice, warning for pitfalls, etc.)

**Quizzes:**
- Read the question — does it test THIS page's content, not general knowledge?
- Try answering correctly — does the correct feedback make sense?
- Try answering incorrectly — does the incorrect feedback help learn?
- Check: Are distractors plausible? Or obviously wrong?
- Check: Is the explanation teaching, not just confirming?
- Check: Does the quiz block render correctly? (buttons, spacing, colors)

**Agent Demos:**
- Click "Launch Demo" to open the workspace
- Step through the conversation (click Send or advance)
- Check: Does the demo illustrate the concept the page teaches?
- Check: Are the file explorer, chat sidebar, and file preview all working?
- Check: Is the conversation realistic and educational?
- Check: Can you exit back to the course cleanly?

**Images/Diagrams:**
- Do images load? (no broken image icons)
- Are they relevant to the content?
- Is the alt text descriptive?
- Do annotated images have working hotspots?

#### E. Navigation
- Does the Previous button go to the right place?
- Does the Next button go to the right place?
- At module boundaries: does Next correctly transition to the next module?
- Does the sidebar accurately reflect progress (checkmarks for viewed pages)?

### Phase 3: Cross-Module Evaluation

After reviewing all pages, evaluate the course holistically:

- **Narrative arc** — does the course tell a coherent story from opening to synthesis?
- **Difficulty progression** — does complexity ramp up appropriately?
- **Repetition** — is any content repeated across modules? (some reinforcement is good; copy-paste is not)
- **Gaps** — are there concepts mentioned but never properly explained?
- **Quiz distribution** — are quizzes spread evenly? Any modules with too few or too many?
- **Visual variety** — is there a good mix of text, diagrams, callouts, and interactive elements?
- **Timing estimates** — do the duration estimates in frontmatter seem realistic?
- **Key takeaways pages** — does every module end with a good summary?

## Report Format

Structure the report as follows:

```markdown
<!-- PROGRESS: module-opening page 2 of 3 -->

# Course Feedback Report

**Reviewed:** [date]
**Reviewer:** Claude (automated student-perspective review)
**Course:** [course title from course.yaml]
**Modules reviewed:** N of N
**Status:** IN PROGRESS / COMPLETE

## Executive Summary

[3-5 bullet points: overall strengths and top issues — write placeholder on first pass, finalize after all modules reviewed]

## Module-by-Module Review

### Module 1: [Title] (module-id)

**Overall:** [1-2 sentence summary]
**Rating:** [Excellent / Good / Needs Work / Major Issues]

#### Page 1: [Page Title]
- **Content:** [observations]
- **Interactive elements:** [observations]
- **Issues:** [specific problems found]
- **Suggestions:** [specific improvements]

#### Page 2: [Page Title]
...

#### Module-Level Notes
- [cross-page observations for this module]

### Module 2: [Title]
...

## Cross-Module Observations

### Strengths
- [what works well across the course]

### Issues
- [problems that span multiple modules]

### Content Gaps
- [topics mentioned but not covered, or missing entirely]

### Interactive Element Audit
- **Quizzes:** [count, quality assessment, distribution]
- **Callouts:** [usage patterns, any overuse/underuse]
- **Agent Demos:** [count, quality, relevance]
- **Diagrams/Images:** [count, quality, any broken]

## Priority Fixes

### P0 — Must Fix (blocks learning or is factually wrong)
1. [specific fix with page reference]

### P1 — Should Fix (degrades quality noticeably)
1. [specific fix with page reference]

### P2 — Nice to Have (polish and refinement)
1. [specific fix with page reference]

## Statistics

| Metric | Count |
|--------|-------|
| Total modules | N |
| Total pages | N |
| Total quizzes | N |
| Total callouts | N |
| Total agent demos | N |
| Total images/diagrams | N |
| Pages with no interactive elements | N |
| Broken images | N |
| Estimated total duration | Xh Ym |
```

## Evaluation Rubric

Use these standards when rating content quality:

### Excellent
- Teaches a clear concept with vivid, specific examples
- Every sentence adds value — no filler
- Quiz tests understanding, not recall
- Interactive elements enhance rather than decorate

### Good
- Concept is clear but examples could be more specific
- Minor wordiness but nothing egregious
- Quiz works but could be sharper
- Interactive elements are relevant

### Needs Work
- Concept is muddled or explanation is too abstract
- Significant padding or generic statements
- Quiz is too easy, too vague, or tests general knowledge
- Missing interactive elements where they'd help

### Major Issues
- Factually incorrect information
- Page doesn't teach anything clear
- Broken interactive elements
- Content doesn't match module's stated goals

## Browser Automation Tips

### Navigation Pattern
The app uses WebSocket + morphdom, not traditional page loads. After clicking navigation elements, wait briefly for the DOM to update before reading content.

### Key UI Elements to Find

**Module selector:**
- Module cards are in a grid layout
- Each has title, description, difficulty, duration
- Click a card to enter the module

**Sidebar (when in a module):**
- Left column, sticky
- Page list with checkmarks for viewed pages
- Section titles indented under current page (if page has H2 sections)
- "View Progress" button toggles to module selector

**Page content:**
- Main content area (right side, `#view` container)
- Prose-styled narrative HTML
- Interactive blocks inline (quizzes, callouts, agent demos)

**Navigation buttons:**
- Previous/Next at bottom of page content
- These handle both slide navigation (within a page) and page navigation

**Quiz interaction:**
- Options are clickable buttons
- After clicking: green (correct) or red (incorrect) feedback shown
- "Try Again" button appears on incorrect answers

**Agent demos:**
- Launch card with play icon and "Launch Demo" button
- Opens full-screen workspace (file explorer left, chat right)
- "Back to course" link in header to exit

### Screenshot Strategy
Take screenshots at key moments:
- Module selector overview
- Each page first impression (before scrolling)
- Quiz states (unanswered, correct, incorrect)
- Agent demo workspace
- Any visual glitches or issues found

### Reading Page Content
Use `get_page_text` to extract readable text from each page. This is faster than reading the DOM tree for content evaluation. Use `read_page` when you need to inspect specific UI elements (quiz buttons, callout styling, etc.).

### Handling Slides
Pages with H2 sections show as slides. The Next button at the bottom cycles through slides before advancing to the next page. Watch for:
- Slide counter in the page header
- Section titles appearing in the sidebar
- Progress tracking (checkmarks) for individual slides

## Scope Options

When invoked with arguments:
- **No argument or `full`** — review the entire course (all modules, all pages)
- **Module name** (e.g., `module-context`) — review only that module in depth
- **`quick`** — quick pass: module selector + first/last page of each module + spot-check 2-3 random pages per module

## Important Notes

- This is an **ultra-premium training product**. Review with high standards. Generic content, sloppy explanations, or mediocre quizzes are unacceptable.
- Be **specific** in feedback. "This page could be better" is useless. "The second paragraph uses 'leverage AI capabilities' — replace with a concrete example of what this means in practice" is actionable.
- **Flag factual claims** that seem wrong or unverifiable. The course cites specific statistics and studies — verify plausibility.
- **Note inconsistencies** between modules. If module-context says X and module-prompting contradicts it, flag both.
- The report is for the course author, who is technical and wants direct, honest feedback. Don't soften or hedge.
