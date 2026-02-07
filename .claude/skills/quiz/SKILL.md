---
name: quiz
description: Create quiz blocks — multiple-choice comprehension checks embedded in lesson markdown pages. Use when authoring new quizzes or editing existing ones.
argument-hint: "[page path or topic description]"
---

# Quiz Block Author

You create **quiz blocks** — interactive multiple-choice comprehension checks embedded in training course lesson pages. Quizzes test understanding of the page's content and provide immediate feedback with explanations.

## Before Starting

1. **Read the target page** if a path is given. Understand EXACTLY what the page teaches — the quiz must test comprehension of THIS content, not general knowledge.
2. **Read surrounding pages** in the module to understand where this page fits in the learning arc.
3. **Read the module's `_module.yaml`** for context on module goals and completion requirements.

## What You Produce

A complete ` ```quiz ` fenced code block in valid YAML, ready to paste into a markdown lesson page. Usually placed at the end of the page or at the end of a major section.

## YAML Schema

```yaml
id: <string>           # REQUIRED. Unique kebab-case ID scoped to module (e.g., "context-decay-q1")
type: multiple-choice   # Optional. Only "multiple-choice" is implemented. Can omit.
question: <string>     # REQUIRED. The question text shown to learners.
options:               # REQUIRED. Array of 3-4 answer choices (strings).
  - "First option"
  - "Second option"
  - "Third option"
  - "Fourth option"
answer: <int>          # REQUIRED. Zero-indexed position of the correct answer (0 = first option).
explanation: <string>  # REQUIRED. Feedback shown after answering, regardless of correctness.
```

That's it. Quizzes are deliberately simple — five fields.

## Field Details

### `id`
- Kebab-case, unique within the module
- Convention: `<topic>-<descriptor>` e.g., `context-window-q1`, `files-why-matter`, `hallucinations-insight`
- Used for progress tracking — changing an ID breaks existing score records

### `question`
- Clear, specific, tests understanding of the page's actual content
- Phrased as a question ending with `?`
- Should NOT be answerable from general knowledge alone — the page content should be needed

### `options`
- 3-4 options (4 is standard, 3 is fine for simpler concepts)
- Exactly ONE correct answer
- All options should be plausible — no obviously absurd distractors
- Options should be roughly the same length (a much longer option signals "this is the answer")
- Avoid "all of the above" / "none of the above"

### `answer`
- Zero-indexed: 0 = first option, 1 = second, 2 = third, 3 = fourth
- Don't always put the answer at the same position — vary across quizzes in a module

### `explanation`
- Shown after the learner answers (whether correct or incorrect)
- Should explain WHY the correct answer is correct
- Should reinforce the key concept from the page
- 1-3 sentences. Concise but complete.

## Visual States

The quiz renders through these states:

1. **Unanswered:** Options appear as clickable buttons with hover effects. Clean, inviting.
2. **Correct:** Correct option highlighted teal. Feedback box says "Correct!" with the explanation. No retry button.
3. **Incorrect:** Correct option highlighted teal, chosen option highlighted red. Feedback says "Incorrect" with the explanation. "Try Again" button shown.
4. **Live Poll Mode:** Instructor can open quiz as a live poll — learners vote, results shown as bar charts with percentages.

Learners can retry incorrect answers. Progress tracks whether they got it right and how many attempts.

## Design Principles

### 1. Test the Page, Not General Knowledge

The quiz exists to check that the learner absorbed THIS page's content. A good test: could someone answer correctly WITHOUT reading the page? If yes, the question is too generic.

Bad: "What is a context window?" (general knowledge)
Good: "According to the hierarchy described in this lesson, which context source has the highest priority?" (requires reading the page)

### 2. One Concept Per Quiz

Each quiz should test ONE clear concept from the page. Don't combine multiple ideas into one question. If a page covers three important concepts, consider three separate quizzes rather than one that tries to cover everything.

### 3. Distractors Should Be Plausible

Wrong options should represent common misconceptions or things that SOUND right but aren't. They should require actual understanding to distinguish from the correct answer.

Bad distractors: Obviously wrong, joke answers, irrelevant topics
Good distractors: Things a learner might believe if they misunderstood the concept

### 4. Explanations Teach, Not Just Confirm

The explanation is a teaching moment. Don't just say "B is correct." Explain the reasoning. Connect back to the page content. Help the learner understand why the other options are wrong.

Bad: "The correct answer is B because files have working formulas."
Good: "When AI creates actual files (like Excel), the functionality works — formulas calculate, charts update, macros run. Chat text that looks like a spreadsheet has none of this functionality."

### 5. Vary Answer Positions

Across the quizzes in a module, spread the correct answer across positions 0-3. Don't put the answer at position 1 every time. Learners notice patterns.

### 6. Question Placement

- **End of page:** Most common. Tests the page as a whole.
- **After a major section:** Good for long pages with H2 sections. Tests that specific section before moving on.
- **Multiple per page:** Fine if the page covers distinct concepts. Each quiz gets its own block index.

## ID Naming Convention

```
<topic>-<descriptor>
```

- Scoped to the module — IDs must be unique within a module, not globally
- Use kebab-case throughout
- Examples: `context-window-q1`, `delegation-insight`, `hallucinations-check`, `prompting-takeaway`
- For modules with multiple quizzes on the same topic, use numbered suffixes: `context-decay-q1`, `context-decay-q2`

## Workflow

When the user asks you to create a quiz:

1. **Read the page content** — understand exactly what's taught.
2. **Identify the key concept** — what's the ONE thing the learner must take away?
3. **Write the question** — specific to the page content, not general knowledge.
4. **Write distractors** — plausible wrong answers based on common misconceptions.
5. **Place the correct answer** — vary position across quizzes.
6. **Write the explanation** — teach why the answer is correct.
7. **Output the complete block** as a fenced code block ready to paste.

## Reference Examples

### Example 1: Testing a specific concept from a page

Imagine a page teaching that AI-generated files (Excel, Word, etc.) have real functionality — formulas calculate, charts update — unlike chat text that merely looks like a spreadsheet.

````yaml
```quiz
id: file-functionality-q1
type: multiple-choice
question: "What does 'functionality' mean in the context of file-based AI output?"
options:
  - "Files are faster to download than chat responses"
  - "Generated spreadsheets have working formulas, not just text in cells"
  - "Files can be opened on any device"
  - "AI can create more file types than before"
answer: 1
explanation: "When AI creates actual files (like Excel), the functionality works — formulas calculate, charts update, macros run. Chat text that looks like a spreadsheet has none of this functionality."
```
````

Why this works:
- Tests a specific concept from the page ("functionality"), not general knowledge
- Distractors sound plausible (speed, compatibility, variety) but miss the point
- Explanation reinforces the teaching by giving concrete examples
- Answer at position 1 (not always 0)

### Example 2: Testing a key insight

Imagine a page teaching that you should upload actual files to AI rather than describing your data in chat, because descriptions lose detail and introduce errors.

````yaml
```quiz
id: provide-not-describe-q1
type: multiple-choice
question: "Why is uploading actual files better than describing your data to AI?"
options:
  - "It's faster for AI to process files"
  - "AI can only work with files, not descriptions"
  - "AI works with actual information, not your summary of it"
  - "Files use less of the context window"
answer: 2
explanation: "When you describe data, you're creating a summary that may miss details or introduce errors. When you upload the actual file, AI works with the real information — all of it, exactly as it exists."
```
````

Why this works:
- Question targets the page's key insight (provide, don't describe)
- Option 1 is plausible but wrong (speed isn't the point)
- Option 3 is factually incorrect (AI CAN work with descriptions)
- Option 0 sounds technical but is wrong (files are often larger in context)
- Explanation connects back to the "why" — fidelity of information

### Example 3: Testing understanding of a hierarchy or model

Imagine a page teaching that context has a priority hierarchy: system prompt > user instructions > conversation history > retrieved documents.

````yaml
```quiz
id: context-hierarchy-q1
type: multiple-choice
question: "In the context priority hierarchy described in this lesson, which source takes highest precedence?"
options:
  - "The most recent user message"
  - "Retrieved documents and search results"
  - "The system prompt"
  - "The longest piece of context provided"
answer: 2
explanation: "The system prompt sits at the top of the context hierarchy. It frames how the AI interprets everything else — user messages, conversation history, and retrieved documents all operate within the boundaries set by the system prompt."
```
````

Why this works:
- "described in this lesson" makes it explicitly page-specific
- Tests a specific model (the hierarchy), not vague understanding
- Each distractor represents a plausible misconception (recency, volume, relevance)
- Answer at position 2 (varying across examples)

## Module Completion

Quizzes can be required for module completion. The `_module.yaml` can specify:

```yaml
completion:
  require_all_pages: true
  require_quizzes: true      # All quizzes must be answered correctly
  min_quiz_score: 0.80       # Or: at least 80% correct
```

This means quiz quality matters — they're not throwaway. Poorly written quizzes that learners can't answer become blockers to progress.
