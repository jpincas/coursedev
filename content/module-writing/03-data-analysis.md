---
title: "Data Analysis Without a Data Team"
duration: "12m"
tags: [data, analysis, spreadsheets]
---

# Data Analysis Without a Data Team

AI-powered data analysis has democratised work that previously required specialists.

## The Three Tiers

![Three Tiers of AI Data Analysis Tools](/content/module-writing/images/data-analysis-tiers.svg)

**AI add-ins** like GPTExcel and Numerous.ai handle formula generation, report summaries, and basic cleaning. You plus a spreadsheet plus better macros.

**Chat-to-SQL translators** like Julius AI let you type questions and receive charts and statistical output. "Show me sales trends by region" becomes a visualisation.

**Full-stack AI analysts** like Anomaly AI and Quadratic connect datasets, inspect schemas, clean data, propose metrics, and build dashboards. The SQL is visible behind every insight.

```callout
type: note
title: "Tool Names Change"
content: "The specific tools mentioned here are examples from early 2026. Tool names and capabilities evolve rapidly. Focus on understanding the three-tier pattern — add-ins, chat-to-SQL, and full-stack analysts — rather than memorising specific product names."
```

Choose based on complexity and frequency of use.

## The Upload-and-Analyse Workflow

Both ChatGPT (Advanced Data Analysis) and Claude accept CSV and Excel files.

**ChatGPT:** Up to 10 files per conversation at 512MB each. Writes and executes Python code in a secure sandbox. Uses pandas for analysis and Matplotlib for charts.

**Claude:** Code execution generates downloadable spreadsheets, CSVs, and reports. Creates interactive visualisations via Artifacts using Plotly.js and D3.js.

![Upload-and-Analyse Workflow](/content/module-writing/images/data-workflow.svg)

The workflow that works:

1. **Prepare clean data** with descriptive column headers
2. **Upload it**
3. **Ask AI to describe the dataset first** — this confirms it "understands" the structure
4. **Then ask specific analytical questions**
5. **Verify the results** — check code, spot-check calculations

Don't start with "analyse this data." Start with "describe what's in this dataset" to calibrate.

## Data Cleaning

Typically the most time-consuming part of any analysis. AI can automate:

- **Duplicate removal** — identify and merge identical records
- **Missing value detection** — flag incomplete rows, suggest fill strategies
- **Date format standardisation** — convert mixed date formats to consistent structure
- **Outlier identification** — spot values that fall outside expected ranges
- **Text normalisation** — fix capitalisation, trim whitespace, standardise categories

Numerous.ai's `/clean` command and similar tools handle this conversationally.

```callout
type: tip
title: "Clean Before You Upload"
content: "Spend 10 minutes preparing data before uploading. Remove obviously broken rows, standardise column names, convert dates to ISO format. Clean input produces reliable output."
```

## What AI Can Do With Your Data

Once uploaded:
- Calculate statistics (mean, median, distributions)
- Identify trends and patterns
- Generate visualisations (bar charts, line graphs, heatmaps)
- Perform correlation analysis
- Run statistical tests
- Create summary tables
- Export results as files

All of this conversationally, no code required.

## The Context Window Trap

Large datasets can overflow the context window. If your data has thousands of rows, AI may:
- Only analyse a sample
- Miss patterns in data it can't "see"
- Produce inconsistent results across requests

For very large datasets, consider:
- Aggregating data before upload (daily summaries instead of individual transactions)
- Using specialised tools (Julius AI, Quadratic) designed for large data
- Breaking analysis into smaller focused questions

## The Describe-First Workflow

Watch the correct approach to data analysis: describe first, then analyse. The describe step catches misunderstandings before they corrupt the analysis.

```agent
id: data-analysis-demo
title: "Describe, Then Analyse"
model_label: "Claude"

system: |
  You are a data analyst. When given data, describe its structure
  first before performing analysis. Be precise about what you see.

scratchpad:
  "employee-survey.csv": |
    dept,role_level,tenure_yrs,engagement_score,intent_to_stay,remote_pct,manager_rating
    Engineering,senior,7,4.2,yes,80,4.5
    Engineering,mid,3,3.1,unsure,60,3.2
    Engineering,junior,1,4.5,yes,40,4.0
    Marketing,senior,5,2.8,no,90,2.1
    Marketing,mid,2,3.9,yes,70,4.2
    Marketing,junior,1,4.1,yes,30,3.8
    Sales,senior,8,3.5,unsure,20,3.9
    Sales,mid,4,2.4,no,10,2.5
    Sales,junior,1,4.3,yes,15,3.7
    Product,senior,6,4.6,yes,75,4.8
    Product,mid,3,3.8,yes,65,4.1
    Product,junior,2,4.0,yes,50,3.6
    Support,senior,9,2.1,no,95,1.8
    Support,mid,5,2.6,no,80,2.3
    Support,junior,1,3.5,unsure,60,3.0

tools:
  - scratchpad_read

visibility:
  system_prompt: hidden
  tool_calls: visible
  full_context: hidden

sidebar:
  width: "45%"
  start_open: true

script:
  - type: note
    text: "Step 1 of the upload-and-analyse workflow: describe the dataset. This catches misunderstandings before they corrupt the analysis."

  - type: user
    content: "Read the employee survey data. Before analysing anything, describe what you see — columns, data types, ranges, any immediate observations."

  - type: tool_call
    tool: scratchpad_read
    args:
      filename: "employee-survey.csv"

  - type: tool_result
    tool: scratchpad_read
    content: |
      dept,role_level,tenure_yrs,engagement_score,intent_to_stay,remote_pct,manager_rating
      Engineering,senior,7,4.2,yes,80,4.5
      Engineering,mid,3,3.1,unsure,60,3.2
      Engineering,junior,1,4.5,yes,40,4.0
      Marketing,senior,5,2.8,no,90,2.1
      Marketing,mid,2,3.9,yes,70,4.2
      Marketing,junior,1,4.1,yes,30,3.8
      Sales,senior,8,3.5,unsure,20,3.9
      Sales,mid,4,2.4,no,10,2.5
      Sales,junior,1,4.3,yes,15,3.7
      Product,senior,6,4.6,yes,75,4.8
      Product,mid,3,3.8,yes,65,4.1
      Product,junior,2,4.0,yes,50,3.6
      Support,senior,9,2.1,no,95,1.8
      Support,mid,5,2.6,no,80,2.3
      Support,junior,1,3.5,unsure,60,3.0

  - type: assistant
    content: "**Dataset description:**\n\n**Structure:** 15 rows x 7 columns. One row per employee (sampled by department and role level).\n\n**Columns:**\n- `dept` — 5 departments: Engineering, Marketing, Sales, Product, Support\n- `role_level` — 3 levels: junior, mid, senior (one per department)\n- `tenure_yrs` — integer, range 1-9 years\n- `engagement_score` — decimal, range 2.1-4.6 (appears to be 1-5 scale)\n- `intent_to_stay` — categorical: yes/no/unsure\n- `remote_pct` — integer 10-95, percentage of time working remotely\n- `manager_rating` — decimal, range 1.8-4.8 (appears to be 1-5 scale)\n\n**Immediate observations:**\n- Small sample (15 rows, 3 per department) — statistical significance will be limited\n- Two score columns (engagement and manager rating) seem to track together\n- The 'no' responses for intent_to_stay cluster in Support and Marketing seniors\n\n**Caveat:** With only 3 data points per department, any department-level analysis should be treated as directional, not statistically conclusive."

  - type: note
    text: "Look at that caveat at the end: 'only 3 data points per department.' The AI flagged the sample size limitation upfront. If you'd skipped the describe step and gone straight to 'analyse engagement by department,' you'd get confident-looking averages based on just 3 people each. The describe step caught the limitation before it became a misleading conclusion."

  - type: user
    content: "Good. Now the analysis: what's the relationship between remote work percentage and engagement? And which departments should I worry about for retention?"

  - type: assistant
    content: "**Remote work and engagement:**\n\nWith the caveat that n=15 is small, the data suggests a complex picture. Remote percentage alone doesn't predict engagement:\n- Support has the highest remote % (avg 78%) but lowest engagement (avg 2.7)\n- Product has moderate remote % (avg 63%) and highest engagement (avg 4.1)\n- Sales has lowest remote % (avg 15%) with mixed engagement (avg 3.4)\n\nThe pattern suggests **manager rating is a stronger predictor than remote percentage.** Engagement and manager rating track closely across the dataset (correlation appears strong).\n\n**Retention risk — two departments to watch:**\n\n**Support (HIGH RISK):** Lowest engagement (2.7 avg), lowest manager ratings (2.4 avg), 2 of 3 respondents intend to leave. The senior member (9 years, engagement 2.1, manager rating 1.8) is a significant flight risk.\n\n**Marketing (MODERATE RISK):** The senior member (5 years, engagement 2.8, manager rating 2.1) intends to leave. Mid and junior levels are healthier. This looks like a leadership-specific issue.\n\n**Note:** These patterns are directional given the sample size. I'd recommend surveying the full departments before acting on these findings."

  - type: note
    text: "Notice: the AI consistently flags the sample size limitation. It gives directional findings, not false certainty. That's what the describe-first step enables — it calibrated the AI's confidence to match the data quality. Skip that step and you'd get the same analysis without the caveats."
```

```quiz
id: data-workflow-quiz
type: multiple-choice
question: "Why should you ask AI to 'describe the dataset' before asking analytical questions?"
options:
  - "It reduces token usage by loading the data into memory more efficiently"
  - "It confirms the AI correctly understands the data structure before analysis"
  - "It automatically cleans the data and removes outliers"
answer: 1
explanation: "Asking AI to describe the dataset first lets you verify it correctly understands column names, data types, and structure. This prevents analysis errors caused by misinterpreting the data."
```
