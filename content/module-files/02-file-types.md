---
title: "Understanding File Types"
duration: "15m"
tags: [files, types, plain-text, markdown, csv]
---

# Understanding File Types

Before diving into how AI works with files, you need to understand a fundamental distinction that shapes everything: **plain text vs binary formats.**

## Plain Text: AI-Native Formats

Plain text files are exactly what they sound like -- files made of readable characters. You can open them in any text editor and read the contents directly.

**Common plain text formats:**

| Format | Extension | What It Is | Why It Matters |
|--------|-----------|------------|----------------|
| Plain text | .txt | Simple unformatted text | Universal, no conversion needed |
| CSV | .csv | Comma-separated values (data) | The universal data exchange format |
| Markdown | .md | Formatted text with simple syntax | The lingua franca of AI |
| JSON | .json | Structured data | How APIs and configs communicate |
| HTML/CSS/JS | .html, .css, .js | Web content | "Code" that is really just text |
| Source code | .py, .go, .java | Programming languages | Text files with special syntax |

**Why plain text matters for AI:** Models process text natively. When you give AI a plain text file, it reads the content directly -- every word, every data point, every line. No conversion, no loss, no interpretation layer.

## Binary Formats: Conversion Required

Binary files are encoded in ways that require specific programs to read. You cannot open a .docx file in a text editor and read it -- you would see gibberish.

**Common binary formats:**
- Word documents (.docx)
- Excel spreadsheets (.xlsx)
- PowerPoint presentations (.pptx)
- PDFs (.pdf)
- Images (.png, .jpg)

When you give AI a binary file, it must first **convert** the content into text it can process. This conversion sometimes loses formatting, embedded objects, or structural information.

```callout
type: info
title: "The Practical Implication"
content: "AI works best with plain text because it reads it directly. Binary formats work too, but there is always a conversion step that can lose information. When you have a choice, plain text formats are more reliable."
```

## CSV: The Universal Data Format

**CSV (Comma-Separated Values)** is the simplest and most portable data format. It is just text with commas separating columns and line breaks separating rows.

```
Name,Department,Salary,Start_Date
Sarah Chen,Engineering,85000,2023-03-15
James Wilson,Marketing,72000,2024-01-10
Maya Patel,Design,78000,2023-09-01
```

Every spreadsheet tool can export to CSV. Every programming language can read CSV. Every AI model can process CSV perfectly.

When sharing data with AI, CSV is almost always the best choice: no formatting to lose, no formulas to break, just clean data.

## Markdown: The Lingua Franca of AI

**Markdown** is a simple way to format text using plain characters. It is the format that almost every AI tool uses internally.

**Basic syntax:**

```
# Heading 1
## Heading 2
### Heading 3

**Bold text** and *italic text*

- Bullet point
- Another bullet point

1. Numbered item
2. Second item

| Column 1 | Column 2 |
|----------|----------|
| Data     | More data|

> Blockquote for emphasis

`inline code` and code blocks with triple backticks
```

```callout
type: tip
title: "The 5-Minute Investment"
content: "Learning basic Markdown is the single highest-ROI skill investment you can make for AI work. Every major AI tool -- Claude, ChatGPT, Gemini, Notion AI -- uses Markdown internally. When you write in Markdown, AI understands your structure perfectly."
```

**Why Markdown matters:**
- AI reads Markdown structure perfectly (headings, lists, tables)
- AI outputs in Markdown by default
- It is the standard format for CLAUDE.md files, project knowledge, and persistent instructions
- It renders beautifully in most tools but is still readable as plain text

## Images and AI

Images are binary files, but modern AI models can process them visually. You can upload screenshots, diagrams, photos, and charts.

However, AI cannot read text embedded in images as reliably as actual text files. If you have a choice between a screenshot of a spreadsheet and the actual CSV data, always use the CSV.

```callout
type: warning
title: "Screenshots vs Source Files"
content: "A screenshot of a spreadsheet gives AI an image to interpret. The actual CSV gives AI perfect data to work with. Always prefer source files over screenshots when the data matters."
```

## Explore the Difference

Click through the files in the explorer. Notice how the plain text files (.txt, .csv, .md, .json) are immediately readable -- you can see exactly what is in them. Now click on the image file and the PowerPoint. That is what binary looks like when you try to read it as text.

```agent
id: file-types-explorer
title: "What's Inside Your Files?"
model_label: "Claude"

system: |
  You are a helpful assistant that examines files and explains
  what type they are and how AI processes them.

scratchpad:
  "meeting-notes.txt": |
    Product team sync - Feb 3rd
    Attendees: Sarah, James, Priya, Tom
    Key decisions:
    - Ship v2.1 by end of month
    - Hire two more engineers in Q2
    - Move weekly sync to Tuesdays
    Action items:
    - Sarah: draft hiring brief by Friday
    - James: finalise release checklist
    - Priya: update roadmap in Notion
  "quarterly-sales.csv": |
    Region,Q1_Revenue,Q2_Revenue,Q3_Revenue,Q4_Revenue
    North,245000,268000,251000,312000
    South,189000,195000,203000,224000
    East,312000,298000,335000,358000
    West,156000,172000,168000,191000
  "project-brief.md": |
    # Project Phoenix: Website Redesign

    ## Objective
    Redesign the company website to improve **conversion rates** and reduce **bounce rate** by 30%.

    ## Timeline
    - **Phase 1:** Research and wireframes (Feb)
    - **Phase 2:** Design and prototyping (Mar)
    - **Phase 3:** Build and test (Apr-May)

    ## Key Requirements
    1. Mobile-first responsive design
    2. Faster page load times (< 2 seconds)
    3. Integrated blog with *SEO optimisation*

    > "The current site converts at 1.2%. Industry average is 2.8%." — Marketing team
  "app-config.json": |
    {
      "app_name": "Phoenix Dashboard",
      "version": "2.1.0",
      "database": {
        "host": "db.internal.company.com",
        "port": 5432,
        "name": "phoenix_prod"
      },
      "features": {
        "dark_mode": true,
        "notifications": true,
        "export_csv": true
      }
    }
  "team-photo.png": |
    âPNG

    IHDR╠É¢wôÜsRGBÇ╠î
    gAMA╠ ╠±ÅüIDATx^ì¢ÿ¦Ö§ÐÑ
    ÒÜàßÞ×ÖÏÎÇÆ¿½¶µ®¬¥¤£¢
    ¡ÿþýüûúùø÷öõôóòñðïîíìëê
    éèçæåäãâáàÞÝÜÛÚÙØ×ÖÕÔÓ
    zxvtsrqponmlkjihgfedcba`_
    ^]\[ZYXWVUTSRQPONMLKJIHGF
    EDCBA@?>=<;:9876543210/.-
    ╠ûIEND®B`é
  "presentation.pptx": |
    PK!╠Ú[Content_Types].xml
    ¡╠ÎN├0►Eÿ!╠½JÕ*Ç►
    ²ä╠╠·Ø╬╠╠╠╠Î╠ñ╠╠ì
    PPTx/slides/slide1.xml
    ½╠╠½╠╠╠╠╠ÿÉ½╠ñ╠╠ì
    ╠╠╠å╠╠╠╠Î╠╠╠╠╠╠╠╠
    ╠╠¥╠╠é╠╠╠╠Ð╠╠╠╠╠╠
    PK╠╠╠╠╠╠╠╠╠╠╠╠╠╠╠

tools:
  - scratchpad_read
  - list_files

visibility:
  system_prompt: hidden
  tool_calls: visible
  full_context: hidden

sidebar:
  width: "50%"
  start_open: true

script:
  - type: note
    text: "Click through the files in the explorer on the left. The first four files — .txt, .csv, .md, .json — are all plain text. You can read every word. Now click on 'team-photo.png' and 'presentation.pptx'. That garbled mess is what binary files look like as raw data. AI has the same experience: plain text is crystal clear, binary needs conversion."

  - type: user
    content: "I have a mix of files here. Can you read through them and tell me which ones you can work with directly?"

  - type: tool_call
    tool: list_files
    args: {}

  - type: tool_result
    tool: list_files
    content: |
      6 files:
      - app-config.json
      - meeting-notes.txt
      - presentation.pptx
      - project-brief.md
      - quarterly-sales.csv
      - team-photo.png

  - type: tool_call
    tool: scratchpad_read
    args:
      filename: "meeting-notes.txt"

  - type: tool_result
    tool: scratchpad_read
    content: |
      Product team sync - Feb 3rd
      Attendees: Sarah, James, Priya, Tom
      Key decisions:
      - Ship v2.1 by end of month
      - Hire two more engineers in Q2
      - Move weekly sync to Tuesdays

  - type: tool_call
    tool: scratchpad_read
    args:
      filename: "quarterly-sales.csv"

  - type: tool_result
    tool: scratchpad_read
    content: |
      Region,Q1_Revenue,Q2_Revenue,Q3_Revenue,Q4_Revenue
      North,245000,268000,251000,312000
      South,189000,195000,203000,224000

  - type: tool_call
    tool: scratchpad_read
    args:
      filename: "project-brief.md"

  - type: tool_result
    tool: scratchpad_read
    content: |
      # Project Phoenix: Website Redesign
      ## Objective
      Redesign the company website to improve **conversion rates**...

  - type: tool_call
    tool: scratchpad_read
    args:
      filename: "team-photo.png"

  - type: tool_result
    tool: scratchpad_read
    content: |
      âPNG
      IHDR╠É¢wôÜsRGBÇ╠î
      gAMA╠ ╠±ÅüIDATx^ì¢ÿ¦Ö§ÐÑ
      ÒÜàßÞ×ÖÏÎÇÆ¿½¶µ®¬¥¤£¢

  - type: assistant
    content: "Here's what I found:\n\n**Plain text — I can read these perfectly:**\n- **meeting-notes.txt** — Product team sync notes with decisions and action items\n- **quarterly-sales.csv** — Revenue data by region across 4 quarters\n- **project-brief.md** — Website redesign brief with timeline and requirements\n- **app-config.json** — Application configuration with database and feature settings\n\n**Binary — I cannot read the raw data:**\n- **team-photo.png** — This is image data. I can see it's a PNG file from the header, but the contents are encoded binary. I would need image processing to interpret this.\n- **presentation.pptx** — This is a PowerPoint file (actually a ZIP archive). The raw bytes are not readable as text.\n\nThe four plain text files give me perfect access to their contents. The binary files would need to be converted or processed through specialised tools before I can work with them meaningfully."

  - type: note
    text: "This is the fundamental distinction. Four files AI reads perfectly, two files it cannot. Every text format — .txt, .csv, .md, .json, .html, .py — gives AI direct access. Binary formats require an extra conversion step that can lose information."
```

```quiz
id: file-types-plain-text
type: multiple-choice
question: "You need AI to analyse financial data from a spreadsheet. Which approach gives AI the cleanest, most reliable access to the data?"
options:
  - "Upload the original .xlsx file so AI can see the formatting and formulas"
  - "Export to CSV and upload the .csv file for direct text processing"
  - "Take a screenshot of the key data and upload the image"
answer: 1
explanation: "CSV is plain text that AI reads directly -- every number, every column, perfectly. Excel files require conversion (which can lose structure), and screenshots require visual interpretation (which can misread numbers). For data analysis, CSV is the most reliable format."
```
