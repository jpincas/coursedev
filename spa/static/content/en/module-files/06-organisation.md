---
title: "Your AI-Ready File System"
duration: "10m"
tags: [organisation, naming, sharing, corporate]
---

# Your AI-Ready File System

A well-organised file system is not just tidy. It is the foundation for every AI interaction you will ever have.

Every time you ask AI to "read these files" or "work with this data," the quality of what AI produces depends partly on how well those files are named, organised, and formatted.

## File Naming Conventions

Bad file names waste context and confuse both you and AI.

**Bad names:**
```
data (3).xlsx
New Document.docx
Screenshot 2026-02-07 at 10.23.45.png
final_final_v2_ACTUAL_FINAL.docx
notes.txt
```

**Good names:**
```
q4-2025-sales-by-region.csv
client-brief-henderson-project.md
meeting-notes-2026-02-07-product-team.md
brand-guidelines-v3.md
competitor-analysis-acme-2026.md
```

**Rules:**
- Use lowercase with dashes (not spaces or underscores)
- Include the topic, not just the type
- Include dates in YYYY-MM-DD format when relevant
- Include version numbers when iterating
- Be descriptive enough that the name alone tells you what is inside

```callout
type: tip
title: "The Name Is Context"
content: "AI reads file names. A file named 'q4-2025-sales-by-region.csv' gives AI context before it even opens the file. A file named 'data.csv' tells it nothing."
```

## Folder Structure

Organise files by project or domain, not by file type.

**Bad structure:**
```
Documents/
  spreadsheets/
    file1.xlsx
    file2.xlsx
  presentations/
    deck1.pptx
```

**Good structure:**
```
projects/
  henderson-project/
    input/
      client-brief.md
      requirements.csv
      brand-guidelines.md
    output/
      strategy-document.md
      proposal-draft.md
      final-presentation.pptx
    reference/
      competitor-analysis.md
      past-proposals/
```

**Principles:**
- Group by project or topic
- Separate input files from output files
- Keep reference materials accessible
- Use a `working/` vs `final/` distinction for drafts

## AI-Powered File Cleanup

One of the most satisfying uses of AI: taking a messy pile of files and organising them automatically.

```agent
id: file-cleanup-demo
title: "AI Organises Your Files"
model_label: "Claude"

system: |
  You are a file organisation assistant. You help professionals
  clean up and organise their files with clear naming conventions
  and logical folder structures.

scratchpad:
  "data (2).csv": |
    Month,Revenue,Costs
    Jan,45000,32000
    Feb,52000,35000
    Mar,48000,33000
  "IMG_4521.txt": |
    Meeting notes from product sync
    - Launch date confirmed: March 15
    - Budget approved: $120K
    - Need design review by Feb 20
  "final report.txt": |
    Q4 Performance Summary
    Revenue: $2.1M (up 18%)
    New customers: 340
    Churn: 2.3%
    NPS: 52
  "stuff.txt": |
    Competitor pricing research
    Acme Pro: $49/user/month
    Beta Corp: $39/user/month
    Gamma Inc: $59/user/month (includes support)
    Our price: $45/user/month
  "doc.txt": |
    Brand voice guidelines
    Tone: Professional but warm
    Avoid: Jargon, passive voice
    Always: Use concrete examples
    Audience: Mid-market business leaders
  "notes feb.txt": |
    Team retrospective Feb 7
    What went well: faster deploys, better QA
    What didn't: client communication delays
    Action: weekly client updates starting next sprint
  "budget thing.csv": |
    Category,Q1_Budget,Q1_Actual,Variance
    Engineering,180000,175000,-5000
    Marketing,95000,102000,+7000
    Operations,45000,43000,-2000
  "new doc.txt": |
    Project Phoenix - Timeline
    Phase 1: Research (Feb)
    Phase 2: Design (Mar)
    Phase 3: Build (Apr-May)
    Phase 4: Launch (Jun)
  "IMG_3847.png": |
    [Binary image data — product team photo from Q4 offsite]
  "Screenshot 2026-01-15 at 14.32.07.png": |
    [Binary image data — screenshot of competitor website homepage]
  "photo (1).jpg": |
    [Binary image data — whiteboard diagram from strategy session]
  "Copy of Q3 numbers.csv": |
    Product,Q3_Units,Q3_Revenue
    Dashboard Pro,1250,187500
    Analytics Suite,890,133500
    API Access,2100,105000
    Support Plus,650,97500

tools:
  - scratchpad_read
  - scratchpad_write
  - list_files
  - move_file
  - create_folder

visibility:
  system_prompt: hidden
  tool_calls: visible
  full_context: hidden

sidebar:
  width: "45%"
  start_open: true

script:
  - type: note
    text: "Look at this mess. Twelve files with names like 'stuff.txt', 'data (2).csv', 'IMG_3847.png', 'Copy of Q3 numbers.csv'. A mix of text, data, and images. You cannot tell what any of them contain from the name alone. This is what a real working folder looks like after six months of saving files without thinking."

  - type: user
    content: "I have a mess of files — text documents, spreadsheets, images, all with terrible names. Can you read through everything, rename them with clear descriptive names, organise them into logical folders, and create a summary index?"

  - type: tool_call
    tool: list_files
    args: {}

  - type: tool_result
    tool: list_files
    content: |
      12 files:
      - budget thing.csv
      - Copy of Q3 numbers.csv
      - data (2).csv
      - doc.txt
      - final report.txt
      - IMG_3847.png
      - IMG_4521.txt
      - new doc.txt
      - notes feb.txt
      - photo (1).jpg
      - Screenshot 2026-01-15 at 14.32.07.png
      - stuff.txt

  - type: tool_call
    tool: scratchpad_read
    args:
      filename: "data (2).csv"

  - type: tool_result
    tool: scratchpad_read
    content: |
      Month,Revenue,Costs
      Jan,45000,32000
      Feb,52000,35000
      Mar,48000,33000

  - type: tool_call
    tool: scratchpad_read
    args:
      filename: "Copy of Q3 numbers.csv"

  - type: tool_result
    tool: scratchpad_read
    content: |
      Product,Q3_Units,Q3_Revenue
      Dashboard Pro,1250,187500
      Analytics Suite,890,133500
      API Access,2100,105000
      Support Plus,650,97500

  - type: tool_call
    tool: scratchpad_read
    args:
      filename: "IMG_4521.txt"

  - type: tool_result
    tool: scratchpad_read
    content: |
      Meeting notes from product sync
      - Launch date confirmed: March 15
      - Budget approved: $120K
      - Need design review by Feb 20

  - type: tool_call
    tool: scratchpad_read
    args:
      filename: "stuff.txt"

  - type: tool_result
    tool: scratchpad_read
    content: |
      Competitor pricing research
      Acme Pro: $49/user/month
      Beta Corp: $39/user/month
      Gamma Inc: $59/user/month (includes support)
      Our price: $45/user/month

  - type: tool_call
    tool: scratchpad_read
    args:
      filename: "doc.txt"

  - type: tool_result
    tool: scratchpad_read
    content: |
      Brand voice guidelines
      Tone: Professional but warm
      Avoid: Jargon, passive voice
      Always: Use concrete examples
      Audience: Mid-market business leaders

  - type: tool_call
    tool: scratchpad_read
    args:
      filename: "final report.txt"

  - type: tool_result
    tool: scratchpad_read
    content: |
      Q4 Performance Summary
      Revenue: $2.1M (up 18%)
      New customers: 340
      Churn: 2.3%
      NPS: 52

  - type: tool_call
    tool: scratchpad_read
    args:
      filename: "notes feb.txt"

  - type: tool_result
    tool: scratchpad_read
    content: |
      Team retrospective Feb 7
      What went well: faster deploys, better QA
      What didn't: client communication delays
      Action: weekly client updates starting next sprint

  - type: tool_call
    tool: scratchpad_read
    args:
      filename: "budget thing.csv"

  - type: tool_result
    tool: scratchpad_read
    content: |
      Category,Q1_Budget,Q1_Actual,Variance
      Engineering,180000,175000,-5000
      Marketing,95000,102000,+7000
      Operations,45000,43000,-2000

  - type: tool_call
    tool: scratchpad_read
    args:
      filename: "new doc.txt"

  - type: tool_result
    tool: scratchpad_read
    content: |
      Project Phoenix - Timeline
      Phase 1: Research (Feb)
      Phase 2: Design (Mar)
      Phase 3: Build (Apr-May)
      Phase 4: Launch (Jun)

  - type: tool_call
    tool: scratchpad_read
    args:
      filename: "IMG_3847.png"

  - type: tool_result
    tool: scratchpad_read
    content: "[Binary image data — product team photo from Q4 offsite]"

  - type: tool_call
    tool: scratchpad_read
    args:
      filename: "Screenshot 2026-01-15 at 14.32.07.png"

  - type: tool_result
    tool: scratchpad_read
    content: "[Binary image data — screenshot of competitor website homepage]"

  - type: tool_call
    tool: scratchpad_read
    args:
      filename: "photo (1).jpg"

  - type: tool_result
    tool: scratchpad_read
    content: "[Binary image data — whiteboard diagram from strategy session]"

  - type: assistant
    content: "I've read all 12 files. Here's what we're working with:\n\n**Text documents (7):**\n- Meeting notes, competitor research, brand guidelines, Q4 report, team retro, project timeline, plus one mislabelled as IMG_4521.txt\n\n**Data files (3):**\n- Q1 revenue/costs, Q1 budget vs actual, Q3 product sales\n\n**Images (3):**\n- Team photo, competitor website screenshot, whiteboard diagram\n\nNote: 'IMG_4521.txt' is actually meeting notes, not an image — someone saved it with the wrong convention. I'll organise everything into logical folders now."

  - type: tool_call
    tool: create_folder
    args:
      path: "data"

  - type: tool_result
    tool: create_folder
    content: "Created folder: data"

  - type: tool_call
    tool: create_folder
    args:
      path: "meetings"

  - type: tool_result
    tool: create_folder
    content: "Created folder: meetings"

  - type: tool_call
    tool: create_folder
    args:
      path: "reference"

  - type: tool_result
    tool: create_folder
    content: "Created folder: reference"

  - type: tool_call
    tool: create_folder
    args:
      path: "images"

  - type: tool_result
    tool: create_folder
    content: "Created folder: images"

  - type: tool_call
    tool: move_file
    args:
      source: "data (2).csv"
      destination: "data/q1-revenue-and-costs.csv"

  - type: tool_result
    tool: move_file
    content: "Moved: data (2).csv -> data/q1-revenue-and-costs.csv"

  - type: tool_call
    tool: move_file
    args:
      source: "budget thing.csv"
      destination: "data/q1-budget-vs-actual.csv"

  - type: tool_result
    tool: move_file
    content: "Moved: budget thing.csv -> data/q1-budget-vs-actual.csv"

  - type: tool_call
    tool: move_file
    args:
      source: "Copy of Q3 numbers.csv"
      destination: "data/q3-product-sales.csv"

  - type: tool_result
    tool: move_file
    content: "Moved: Copy of Q3 numbers.csv -> data/q3-product-sales.csv"

  - type: tool_call
    tool: move_file
    args:
      source: "final report.txt"
      destination: "data/q4-performance-summary.md"

  - type: tool_result
    tool: move_file
    content: "Moved: final report.txt -> data/q4-performance-summary.md"

  - type: tool_call
    tool: move_file
    args:
      source: "IMG_4521.txt"
      destination: "meetings/product-sync-launch-planning.md"

  - type: tool_result
    tool: move_file
    content: "Moved: IMG_4521.txt -> meetings/product-sync-launch-planning.md"

  - type: tool_call
    tool: move_file
    args:
      source: "notes feb.txt"
      destination: "meetings/retrospective-2026-02-07.md"

  - type: tool_result
    tool: move_file
    content: "Moved: notes feb.txt -> meetings/retrospective-2026-02-07.md"

  - type: tool_call
    tool: move_file
    args:
      source: "stuff.txt"
      destination: "reference/competitor-pricing-analysis.md"

  - type: tool_result
    tool: move_file
    content: "Moved: stuff.txt -> reference/competitor-pricing-analysis.md"

  - type: tool_call
    tool: move_file
    args:
      source: "doc.txt"
      destination: "reference/brand-voice-guidelines.md"

  - type: tool_result
    tool: move_file
    content: "Moved: doc.txt -> reference/brand-voice-guidelines.md"

  - type: tool_call
    tool: move_file
    args:
      source: "new doc.txt"
      destination: "reference/project-phoenix-timeline.md"

  - type: tool_result
    tool: move_file
    content: "Moved: new doc.txt -> reference/project-phoenix-timeline.md"

  - type: tool_call
    tool: move_file
    args:
      source: "IMG_3847.png"
      destination: "images/team-photo-q4-offsite.png"

  - type: tool_result
    tool: move_file
    content: "Moved: IMG_3847.png -> images/team-photo-q4-offsite.png"

  - type: tool_call
    tool: move_file
    args:
      source: "Screenshot 2026-01-15 at 14.32.07.png"
      destination: "images/competitor-website-homepage.png"

  - type: tool_result
    tool: move_file
    content: "Moved: Screenshot 2026-01-15 at 14.32.07.png -> images/competitor-website-homepage.png"

  - type: tool_call
    tool: move_file
    args:
      source: "photo (1).jpg"
      destination: "images/strategy-session-whiteboard.jpg"

  - type: tool_result
    tool: move_file
    content: "Moved: photo (1).jpg -> images/strategy-session-whiteboard.jpg"

  - type: tool_call
    tool: scratchpad_write
    args:
      filename: "file-index.md"
      content: |
        # File Index

        ## data/
        - **q1-revenue-and-costs.csv** — Monthly revenue and costs for Q1 (Jan-Mar)
        - **q1-budget-vs-actual.csv** — Q1 budget vs actual spend by department (Engineering, Marketing, Operations)
        - **q3-product-sales.csv** — Q3 unit sales and revenue by product line
        - **q4-performance-summary.md** — Q4 headline results: $2.1M revenue, 340 new customers, 2.3% churn, NPS 52

        ## meetings/
        - **product-sync-launch-planning.md** — Product sync: March 15 launch confirmed, $120K budget, design review needed by Feb 20
        - **retrospective-2026-02-07.md** — Team retro: faster deploys and better QA positive; client communication needs weekly updates

        ## reference/
        - **competitor-pricing-analysis.md** — Pricing comparison: Acme $49, Beta $39, Gamma $59 (incl. support), Us $45
        - **brand-voice-guidelines.md** — Tone: professional but warm. Avoid jargon. Use concrete examples. Audience: mid-market leaders
        - **project-phoenix-timeline.md** — 4 phases: Research (Feb), Design (Mar), Build (Apr-May), Launch (Jun)

        ## images/
        - **team-photo-q4-offsite.png** — Product team photo from Q4 offsite event
        - **competitor-website-homepage.png** — Screenshot of competitor's website for reference
        - **strategy-session-whiteboard.jpg** — Whiteboard diagram from strategy planning session

  - type: tool_result
    tool: scratchpad_write
    content: "Written to file-index.md (20 lines)"

  - type: assistant
    content: "Done. Twelve messy files are now in four logical folders with clear names:\n\n- **data/** — financial data and performance reports (4 files)\n- **meetings/** — meeting notes and retrospectives (2 files)\n- **reference/** — guidelines, research, and project plans (3 files)\n- **images/** — photos and screenshots with descriptive names (3 files)\n\nPlus a file-index.md summarising everything. Notice that 'IMG_4521.txt' was actually meeting notes — the AI identified this from the content, not the misleading filename."

  - type: note
    text: "Twelve files went from chaos to a clean, navigable structure in under a minute. The AI read every file to understand its actual contents — it caught that 'IMG_4521.txt' was meeting notes despite the image-style name. The images got descriptive names too: 'team-photo-q4-offsite.png' tells you exactly what you are looking at."

  - type: note
    text: "This is not just tidying — it is building an AI-ready context library. 'reference/brand-voice-guidelines.md' is instantly usable as grounding context for future writing tasks. 'doc.txt' was not. The five minutes AI spends organising saves hours of searching later."
```

## Sharing and Sync in the Corporate World

In most organisations, files do not live on your laptop alone. They live in shared systems.

**Cloud storage:** OneDrive, Google Drive, Dropbox, SharePoint. These are where team files live. When you organise files for AI work, the same principles apply -- but with an additional consideration: your colleagues need to find things too.

**Practical advice for shared environments:**
- Agree on naming conventions with your team
- Create a shared folder structure that everyone follows
- Keep AI working files separate from shared deliverables until they are ready
- Use version control or file naming (v1, v2, final) to track iterations

**A common workflow:**
1. Pull source files from your team's shared drive
2. Work with AI locally or in your AI tool
3. Review and polish the output
4. Share the finished deliverable back to the team drive

```callout
type: info
title: "The Corporate Reality"
content: "Most organisations have file sharing in place already. The key is to use it deliberately: pull files for AI context, produce outputs, then share results through your normal channels. AI does not replace your collaboration tools — it produces better content to put into them."
```

```quiz
id: file-organisation-naming
type: multiple-choice
question: "You have files named 'report.docx', 'data.csv', and 'notes.txt' for a project. Why is this a problem for AI work?"
options:
  - "The file extensions are wrong — AI prefers .md files"
  - "The names provide no context — AI cannot tell what these files contain without opening each one, and neither can you in three months"
  - "Having three different file formats makes it harder for AI to process them together"
answer: 1
explanation: "Descriptive file names are context. 'q4-sales-by-region.csv' tells both you and AI what the file contains before it is even opened. Generic names like 'data.csv' waste the opportunity to provide context through naming and make your file system harder to navigate."
```
