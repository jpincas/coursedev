---
title: "Output Files: Receiving Deliverables"
duration: "10m"
tags: [output, deliverables, formats]
---

# Output Files

Output files are the deliverables AI creates for you. Not chat text that you copy — actual files that you can download, open in native applications, edit, share, and use.

## What AI Can Create

![AI generating different output file types: documents, spreadsheets, presentations, code, and visualizations](/content/module-files/images/output-file-types.svg)

**Documents**
- Reports and white papers
- Meeting summaries
- Policy documents
- Letters and correspondence

**Spreadsheets**
- Analysis with working formulas
- Data models
- Budgets and forecasts
- Comparison tables

**Presentations**
- Slide decks
- Pitch presentations
- Training materials

**Code**
- Scripts and programs
- Configuration files
- Automation tools

**Visualisations**
- Charts and graphs
- Diagrams
- Infographics

## Real, Usable Files

These are real files. When AI creates an Excel spreadsheet:
- The formulas actually calculate
- Charts update when data changes
- You can add your own modifications
- It opens in Excel like any other spreadsheet

When AI creates a PowerPoint:
- It has real slides
- Animations and transitions can be added
- You can present it directly
- Or edit it further in PowerPoint

```callout
type: info
title: "Professional Quality"
content: "AI-generated files are ready for professional use. You can send that document to your boss, share that spreadsheet with your team, present those slides to clients."
```

## AI Operating in Your File System

Beyond creating individual files, AI can work with your entire file system — organising, moving, and cleaning up files just as a human assistant would.

```agent
path: /content/module-files/agent-desktop-cleanup.yaml
```

## Code Is Just Another Output

Here is something that surprises non-technical users: **code is just text files.** HTML, CSS, and JavaScript are plain text -- and AI can create them trivially.

This means you can ask AI to create:
- **Simple websites** -- landing pages, portfolio sites, project pages
- **Data dashboards** -- interactive charts and visualisations
- **Calculators and tools** -- mortgage calculators, unit converters, project estimators
- **Interactive presentations** -- web-based slides with animations

You do not need to be a programmer. You describe what you want, and AI creates the files. You open them in a browser and they work.

**Example:**
```
Create a simple HTML page that displays our Q4 sales data
as an interactive bar chart. Include the data inline.
Make it look professional with a dark theme.
```

AI creates three files (index.html, style.css, script.js), you open index.html in a browser, and you have an interactive dashboard.

```callout
type: tip
title: "Code as a Document Format"
content: "Think of code not as 'programming' but as another document format. Just as AI can create a Word document or a spreadsheet, it can create a web page. The output is text files that happen to do something when opened in a browser."
```

## Format Considerations

Not all output formats are equal. AI excels at some and struggles with others.

**AI excels at:**
- Markdown documents (its native output format)
- CSV data files
- Plain text reports
- HTML/CSS/JS web content
- Source code in popular languages
- JSON and structured data

**AI can produce but with limitations:**
- Excel spreadsheets (formulas and formatting may need adjustment)
- PowerPoint presentations (basic structure, may need design polish)
- PDF documents (usually via conversion from another format)

**AI struggles with or cannot produce:**
- InDesign layouts
- Figma designs
- Specialised software formats (AutoCAD, SPSS, etc.)
- Complex Excel macros and VBA

**The practical advice:** Ask for plain text or standard formats first. If you need a specialised format, have AI create the content in a format it does well (Markdown, CSV), then convert it yourself using the appropriate tool.

## Asking for Files

To receive files instead of chat text, be explicit:
- "Create an Excel file with this analysis"
- "Generate a Word document containing..."
- "Build a PowerPoint presentation about..."
- "Save this as a CSV file"

Don't just ask for "a report" — ask for "a Word document containing the report."

```callout
type: tip
title: "Creating Long Documents"
content: "For reports, proposals, and other substantial documents, we'll cover the optimal method — sectional drafting — in the Document Creation module. This technique dramatically improves quality for longer work."
```

## Iteration on Files

Once you have the file:
1. Review it in the native application
2. Note what needs changing
3. Either edit directly yourself, or
4. Ask AI for a revised version

You're working with real deliverables now, not chat snippets.

```quiz
id: output-files-explicit
type: multiple-choice
question: "You need an interactive dashboard showing sales trends. You are not a programmer. What is the most practical approach?"
options:
  - "Ask AI to create a PowerPoint with embedded charts, since that is a format you know how to use"
  - "Ask AI to create HTML/CSS/JS files for an interactive web dashboard, then open it in your browser"
  - "Ask AI to describe how to build a dashboard, then hire a developer to implement it"
answer: 1
explanation: "Code is just text files. AI can create a complete interactive web dashboard (HTML, CSS, JavaScript) that you open in a browser -- no programming knowledge needed. This produces a more interactive result than static PowerPoint charts, and you do not need to involve a developer."
```
