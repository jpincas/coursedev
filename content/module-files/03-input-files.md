---
title: "Input Files: Giving AI Your Data"
duration: "10m"
tags: [input, data, context]
---

# Input Files

Input files are the data you give to AI. Instead of describing your information, you provide it directly.

## What AI Can Process

Modern AI systems can handle a wide variety of file types:

![Input file types flowing into AI: documents, data, visual, code, and mixed formats](/content/module-files/images/input-file-types.svg)

**Documents**
- PDF files (up to 100 pages per file recommended)
- Word documents
- Plain text files
- Markdown

**Data**
- Excel spreadsheets
- CSV files
- JSON data

**Visual**
- Images and screenshots
- Diagrams
- Charts

**Code**
- Source code files
- Configuration files
- Scripts

**Mixed**
- Folders containing multiple file types
- Email threads
- Chat logs

```callout
type: info
title: "Current Capabilities"
content: "As of early 2026: ChatGPT supports up to 10 files per conversation at 512MB per file. Claude's code execution generates downloadable spreadsheets, CSVs, reports, and interactive visualisations."
```

## The Key Insight

```callout
type: tip
title: "Provide, Don't Describe"
content: "Don't say 'I have a spreadsheet with sales figures from Q4.' Actually provide the spreadsheet. Let AI see the real data."
```

This eliminates the game of telephone. AI works with the actual information, not your description of it.

**Describing data:**
"The spreadsheet has columns for date, product, region, and revenue. There are about 500 rows. The regions are North, South, East, and West. Revenue ranges from..."

**Providing data:**
*[Upload the file]*

Which is more reliable? Which is less work for you?

## Multiple Files

You can provide multiple files at once:
- A folder of related documents
- Several spreadsheets covering different aspects
- Reference materials alongside the data to analyse

AI processes them together, understanding how they relate.

## Best Practices for File Preparation

**Use machine-readable formats.**
Markdown is preferred over complex formatting. Plain text beats heavily formatted documents.

**Break large documents into segments.**
Keep files under 100 pages each. Break large documents into logical chunks.

**Remove excessive formatting.**
Fonts, colours, and embedded images increase token usage without adding information.

**Understand token economics.**
For Claude Projects, Markdown/Word/PDF use similar space. HTML uses twice as much due to tag interpretation.

**Create focused projects.**
Organise by department or function rather than one massive knowledge base. This keeps context relevant.

**Maintain clear file naming.**
Descriptive names help AI (and you) understand what each file contains. Establish regular review cycles for uploaded content.

## File Limitations to Know

Not everything works perfectly:
- Very large files may need to be split
- Complex formatting sometimes gets lost
- Scanned PDFs need good quality to be readable
- Some proprietary formats may not be supported

When in doubt, try it. AI will tell you if it can't process something.

## Multiple Files Working Together

One file is useful. Multiple files together are powerful. Watch how AI synthesises information across three different source files to produce something none of them could alone.

```agent
id: multiple-inputs-demo
title: "Cross-File Synthesis"
model_label: "Claude"

system: |
  You are a business intelligence analyst. Synthesise information
  across multiple sources. Be specific and data-driven.

scratchpad:
  "data/sales-data.csv": |
    Product,Q3_Units,Q4_Units,Q3_Revenue,Q4_Revenue
    Pro Plan,1240,1580,£124K,£158K
    Team Plan,890,720,£44.5K,£36K
    Enterprise,45,62,£225K,£310K
    Free Tier,8500,9200,£0,£0
    Total,-,-,£393.5K,£504K
  "data/support-tickets.csv": |
    Product,Total_Tickets,Avg_Resolution_Hrs,Top_Issue
    Pro Plan,342,4.2,"Billing confusion on annual vs monthly"
    Team Plan,521,8.1,"Permission settings too complex"
    Enterprise,28,1.8,"Dedicated support — all resolved quickly"
    Free Tier,1205,24+,"No priority — self-service only"
  "data/churn-analysis.txt": |
    Q4 Churn Analysis
    Pro Plan: 3.1% monthly churn (down from 4.2% in Q3)
    Team Plan: 8.7% monthly churn (UP from 5.1% in Q3)
    Enterprise: 0% churn (no losses in Q4)
    Free→Paid conversion: 2.3% (up from 1.8%)

    Exit survey highlights (Team Plan):
    - "Too complicated for small teams"
    - "We only needed basic features, paying for stuff we don't use"
    - "Switched to competitor — simpler pricing"

tools:
  - list_files
  - scratchpad_read
  - scratchpad_write
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
    text: "Three data files in a folder: sales, support tickets, and churn analysis. Each tells part of the story. Watch the AI navigate the folder, connect the dots, and write an insight report."

  - type: user
    content: "I have three data files in the data folder — sales, support tickets, and churn analysis. Read them all, then write an insight report connecting the patterns. What story are they telling together?"

  - type: tool_call
    tool: list_files
    args:
      path: "data"

  - type: tool_result
    tool: list_files
    content: |
      3 files in data/:
      - data/churn-analysis.txt
      - data/sales-data.csv
      - data/support-tickets.csv

  - type: tool_call
    tool: scratchpad_read
    args:
      filename: "data/sales-data.csv"

  - type: tool_result
    tool: scratchpad_read
    content: |
      Product,Q3_Units,Q4_Units,Q3_Revenue,Q4_Revenue
      Pro Plan,1240,1580,£124K,£158K
      Team Plan,890,720,£44.5K,£36K
      Enterprise,45,62,£225K,£310K
      Free Tier,8500,9200,£0,£0
      Total,-,-,£393.5K,£504K

  - type: tool_call
    tool: scratchpad_read
    args:
      filename: "data/support-tickets.csv"

  - type: tool_result
    tool: scratchpad_read
    content: |
      Product,Total_Tickets,Avg_Resolution_Hrs,Top_Issue
      Pro Plan,342,4.2,"Billing confusion on annual vs monthly"
      Team Plan,521,8.1,"Permission settings too complex"
      Enterprise,28,1.8,"Dedicated support — all resolved quickly"
      Free Tier,1205,24+,"No priority — self-service only"

  - type: tool_call
    tool: scratchpad_read
    args:
      filename: "data/churn-analysis.txt"

  - type: tool_result
    tool: scratchpad_read
    content: |
      Q4 Churn Analysis
      Pro Plan: 3.1% monthly churn (down from 4.2% in Q3)
      Team Plan: 8.7% monthly churn (UP from 5.1% in Q3)
      Enterprise: 0% churn (no losses in Q4)
      Free→Paid conversion: 2.3% (up from 1.8%)

      Exit survey highlights (Team Plan):
      - "Too complicated for small teams"
      - "We only needed basic features, paying for stuff we don't use"
      - "Switched to competitor — simpler pricing"

  - type: tool_call
    tool: create_folder
    args:
      path: "reports"

  - type: tool_result
    tool: create_folder
    content: "Created folder: reports"

  - type: tool_call
    tool: scratchpad_write
    args:
      filename: "reports/cross-file-insights.md"
      content: |
        # Q4 Product Intelligence: The Team Plan Problem

        ## The Headlines
        Overall revenue grew 28% (£393.5K → £504K). But that growth is masking a crisis in one tier.

        ## The Pattern Across All Three Sources

        **Enterprise and Pro are thriving.** Revenue up, churn down, support manageable. These tiers are working.

        **Team Plan is in trouble — and all three datasets confirm it:**
        - **Sales:** Units dropped 19% (890 → 720) while every other tier grew
        - **Support:** Highest ticket volume (521) with longest resolution time (8.1 hrs). Top issue: "permission settings too complex"
        - **Churn:** Nearly doubled from 5.1% to 8.7%. Exit surveys cite complexity and feature bloat

        The story is consistent: Team Plan customers find the product too complicated for their needs. They're leaving for simpler alternatives.

        ## The Hidden Risk
        Revenue growth (driven by Enterprise) is hiding this problem. If Team Plan churn continues at 8.7%, the tier loses ~40% of customers annually — making it unsustainable within 3 quarters.

        ## Recommended Actions
        1. **Simplify Team Plan** — strip features, simplify permissions, reduce pricing to match value
        2. **Create a bridge** — consider a "Small Team" tier between Free and current Team Plan
        3. **Fix support bottleneck** — 8.1hr resolution for a paid tier is unacceptable

  - type: tool_result
    tool: scratchpad_write
    content: "Written to reports/cross-file-insights.md (23 lines)"

  - type: assistant
    content: "Created reports/cross-file-insights.md. The key finding: your overall growth is hiding a Team Plan crisis. All three data sources — sales, support, and churn — tell the same story independently. Revenue up overall, but Team Plan units down 19%, churn nearly doubled, and support tickets point to complexity as the cause."

  - type: note
    text: "No single file contained this insight. The sales data showed a unit decline. The support data showed complexity complaints. The churn data showed customers leaving. Only by reading all three together could the AI connect these into a coherent story: the Team Plan is too complex and customers are leaving because of it."

  - type: note
    text: "This is where file-based workflows genuinely shine. Three data sources, synthesised into one actionable insight document, in a single request. Try doing that by describing the data in chat."
```

```quiz
id: input-files-insight
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
