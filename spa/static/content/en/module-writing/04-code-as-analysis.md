---
title: "When AI Writes Code to Think"
duration: "15m"
tags: [analysis, code, data, computation]
---

# When AI Writes Code to Think

The previous page showed you the upload-and-analyse workflow and the tools available. Now let's look at what's actually happening under the hood -- because there are two fundamentally different things AI can do with your data, and understanding the distinction will change how you work.

## Two Modes of Analysis

**Mode 1: Reading and Reasoning**

The AI reads your data into its context window and reasons about it using language. It scans the numbers, spots patterns, and describes what it sees -- much like a person glancing over a table and summarising the trends.

This is what happened in the describe-first demo on the previous page. The AI read the employee survey CSV and reasoned about it linguistically. It spotted that Support had low engagement and that manager ratings tracked with engagement -- all by "reading" the data.

**Mode 2: Writing and Executing Code**

The AI writes a Python script to process your data, runs it in a sandboxed environment, and returns the computed results. The code does the actual arithmetic -- calculating averages, running correlations, projecting trends.

This is what happens when ChatGPT shows "Analysing..." or Claude shows an "Analysis" tool call.

## Why This Matters

Here is the critical insight: **language models cannot reliably do arithmetic.** They predict the next token -- they do not have a calculator built in. Ask a model to add up 500 numbers "in its head" and it will get it wrong. Ask it to write Python to add up 500 numbers, execute the code, and return the result -- and it will be exact every time.

| | Reading & Reasoning | Writing & Executing Code |
|---|---|---|
| **Precision** | Approximate -- may miscalculate | Exact -- real computation |
| **Scale** | Struggles with large datasets | Handles millions of rows |
| **Best for** | Themes, summaries, qualitative insight | Statistics, trends, correlations, projections |
| **Output** | Text descriptions | Numbers, tables, charts, files |

Think of it this way: Mode 1 is like asking a colleague to glance at a spreadsheet and tell you what they notice. Mode 2 is like handing the spreadsheet to an analyst who opens Python and runs the numbers. Both are useful. But you would never trust Mode 1 for your quarterly forecast.

```callout
type: warning
title: "The Invisible Mistake"
content: "When AI analyses data by reading alone, it often sounds confident about numbers that are wrong. It might say 'revenue grew approximately 23%' when the real figure is 18%. You cannot tell from the response whether it calculated or guessed. Always push AI to use code for quantitative claims."
```

## Language Is the Bridge

Here is where it gets powerful. **You do not need to know Python.** You describe what you want in plain English, and the AI translates your question into precise code, runs it, and gives you the answer.

This is genuinely new. Previously, getting a precise answer from data required either knowing Excel formulas, knowing a programming language, or hiring someone who did. Now the barrier is gone. You describe the analysis you want, and AI writes the code to do it.

Language is what makes this possible. Because AI speaks both English and Python fluently, it can translate between your intent and the computer's ability to calculate. This is one of the most powerful things about working with language -- it is the universal interface between what you want to know and what the machine can compute.

## Watch the Difference

Same data, different approach. This time, instead of just reading and reasoning about the numbers, watch the AI write Python code to compute exact growth figures and projections.

```agent
id: code-analysis-demo
title: "AI Writing Code to Analyse Data"
model_label: "Claude"

system: |
  You are a data analyst assistant with access to a Python code execution
  environment. When users ask quantitative questions about data, write Python
  code to compute precise answers rather than estimating.

scratchpad:
  "data/regional-sales.csv": |
    Month,North,South,East,West
    Jan,42300,38100,29400,51200
    Feb,44100,37800,31200,52800
    Mar,48700,39200,33100,54100
    Apr,46200,41500,35800,53400
    May,51800,40200,37200,56700
    Jun,53400,42800,39500,58200
    Jul,49800,44100,41200,55800
    Aug,55200,43500,43800,59100
    Sep,58100,45200,42100,61400
    Oct,61400,46800,44500,63200
    Nov,59800,48200,47200,65800
    Dec,64200,49500,49800,67100

tools:
  - list_files
  - scratchpad_read
  - scratchpad_write
  - run_code

visibility:
  system_prompt: hidden
  tool_calls: visible
  full_context: hidden

sidebar:
  width: "50%"
  start_open: true

script:
  - type: note
    text: "A CSV file with 12 months of sales data across 4 regions. The user wants to know which region is growing fastest. Watch how the AI writes actual Python code to calculate this precisely, rather than eyeballing the numbers."

  - type: user
    content: "I have regional sales data in data/regional-sales.csv. Which region is growing fastest? And based on the trend, what should we budget for from each region in Q1 next year?"

  - type: tool_call
    tool: scratchpad_read
    args:
      filename: "data/regional-sales.csv"

  - type: tool_result
    tool: scratchpad_read
    content: |
      Month,North,South,East,West
      Jan,42300,38100,29400,51200
      Feb,44100,37800,31200,52800
      Mar,48700,39200,33100,54100
      Apr,46200,41500,35800,53400
      May,51800,40200,37200,56700
      Jun,53400,42800,39500,58200
      Jul,49800,44100,41200,55800
      Aug,55200,43500,43800,59100
      Sep,58100,45200,42100,61400
      Oct,61400,46800,44500,63200
      Nov,59800,48200,47200,65800
      Dec,64200,49500,49800,67100

  - type: assistant
    content: "I can see the data. Let me write Python to calculate the exact growth rates and projections rather than estimating from the raw numbers."

  - type: tool_call
    tool: scratchpad_write
    args:
      filename: "analysis.py"
      content: |
        import csv
        from io import StringIO

        data = open("data/regional-sales.csv").read()
        reader = csv.DictReader(StringIO(data))
        regions = {"North": [], "South": [], "East": [], "West": []}

        for row in reader:
            for region in regions:
                regions[region].append(int(row[region]))

        print("=== Regional Growth Analysis ===\n")
        for region, values in regions.items():
            total_growth = (values[-1] - values[0]) / values[0] * 100
            monthly_increase = (values[-1] - values[0]) / 11
            projected = [values[-1] + monthly_increase * (i+1)
                        for i in range(3)]
            print(f"{region}:")
            print(f"  Jan: £{values[0]:,} → Dec: £{values[-1]:,}")
            print(f"  Annual growth: {total_growth:.1f}%")
            print(f"  Q1 projection: £{sum(projected)/3:,.0f}/month")
            print()

        growths = {r: (v[-1]-v[0])/v[0]*100
                  for r, v in regions.items()}
        fastest = max(growths, key=growths.get)
        print(f"Fastest growing: {fastest} ({growths[fastest]:.1f}%)")

  - type: tool_result
    tool: scratchpad_write
    content: "Written to analysis.py (26 lines)"

  - type: note
    text: "The AI has written a Python script. Now it executes it. In tools like ChatGPT and Claude, this runs in a secure sandbox — the code actually executes and produces real computed output."

  - type: tool_call
    tool: run_code
    args:
      filename: "analysis.py"

  - type: tool_result
    tool: run_code
    content: |
      === Regional Growth Analysis ===

      North:
        Jan: £42,300 → Dec: £64,200
        Annual growth: 51.8%
        Q1 projection: £66,191/month

      South:
        Jan: £38,100 → Dec: £49,500
        Annual growth: 29.9%
        Q1 projection: £51,573/month

      East:
        Jan: £29,400 → Dec: £49,800
        Annual growth: 69.4%
        Q1 projection: £53,564/month

      West:
        Jan: £51,200 → Dec: £67,100
        Annual growth: 31.1%
        Q1 projection: £69,518/month

      Fastest growing: East (69.4%)

  - type: tool_call
    tool: scratchpad_write
    args:
      filename: "reports/growth-analysis.md"
      content: |
        # Regional Sales Growth Analysis

        ## Key Finding
        **East is the fastest-growing region at 69.4% annual growth** -- more than double the rate of South (29.9%) or West (31.1%).

        ## Full Breakdown

        | Region | Jan | Dec | Growth | Q1 Projection |
        |--------|-----|-----|--------|----------------|
        | East | £29,400 | £49,800 | **69.4%** | £53,564/mo |
        | North | £42,300 | £64,200 | 51.8% | £66,191/mo |
        | West | £51,200 | £67,100 | 31.1% | £69,518/mo |
        | South | £38,100 | £49,500 | 29.9% | £51,573/mo |

        ## Notable Patterns
        - East started smallest but is closing the gap rapidly
        - North has strong revenue AND strong growth -- projected to overtake West by Q2
        - West leads in absolute revenue but growing slowly
        - South is slowest in both growth and absolute terms

  - type: tool_result
    tool: scratchpad_write
    content: "Written to reports/growth-analysis.md (19 lines)"

  - type: assistant
    content: |
      Here's the analysis. **East is your fastest-growing region at 69.4% growth** -- more than double the rate of South or West.

      The interesting story is North vs West. West currently leads in revenue (£67,100/mo in December), but North is growing faster (51.8% vs 31.1%). At current rates, North overtakes West by Q2.

      For Q1 budgeting: West ~£69.5K/mo, North ~£66.2K/mo, East ~£53.6K/mo, South ~£51.6K/mo. Full report saved to reports/growth-analysis.md.

  - type: note
    text: "The AI did not estimate these numbers. It wrote Python to calculate exact growth rates, averages, and projections. The 69.4% figure for East is precisely computed from the data — not a rough guess. This is the difference between AI 'reading' your data and AI 'computing' with your data."

  - type: note
    text: "Could you have done this in Excel? Yes, if you know the right formulas. In Python? Only if you can code. With AI, you described what you wanted in English and it did the computation for you. Language was the bridge."
```

## Recognising the Two Modes

Most AI tools now clearly indicate when they are using code execution:

- **ChatGPT** shows "Analysing..." and displays the Python code it writes
- **Claude** shows an "Analysis" tool call with expandable code
- **Google Gemini** shows "Running code" with expandable blocks

If you do not see these indicators, the AI is probably reasoning about your data linguistically rather than computing. For quantitative questions, this matters enormously.

```callout
type: tip
title: "Force Code Execution"
content: "If you want to ensure AI uses code rather than estimating, be explicit: 'Write Python code to calculate this' or 'Use code analysis to compute the exact figures.' This pushes the tool into its code execution mode."
```

## When Each Mode Shines

**Reading and reasoning is perfect for:**
- Summarising reports and documents
- Identifying themes in qualitative data (reviews, feedback, interviews)
- Comparing document content
- Explaining what data means in context

**Push for code execution when:**
- You need exact calculations (averages, growth rates, totals)
- Working with more than a few dozen numbers
- You need statistical analysis (correlations, distributions)
- You want charts or visualisations
- You need to transform or clean data
- The answer involves any kind of prediction or projection

```quiz
id: code-vs-reading
type: multiple-choice
question: "You upload a CSV with 2,000 rows of customer data and ask 'What is the average order value for customers who signed up in 2024?' The AI responds: 'Based on the data, the average order value appears to be around £45-50.' What should concern you?"
options:
  - "Nothing -- the AI read the data and gave a reasonable answer"
  - "The vague range suggests the AI estimated rather than calculated -- you should ask it to write code to compute the exact figure"
  - "2,000 rows is too much data for AI to process accurately"
answer: 1
explanation: "A vague range like '£45-50' is a telltale sign that the AI read the data and estimated rather than computing the exact answer. With 2,000 rows, you should always push for code execution. Ask the AI to 'write Python to calculate the exact average order value' and you will get a precise number like £47.32 -- not an approximation."
```
