---
title: "Code Generation for Non-Coders"
duration: "20m"
tags: [code, automation, scripts, apis, json, vibe-coding]
---

# Code Generation for Non-Coders

You don't need to become a programmer. But understanding what scripts, APIs, and data formats actually are — even at a high level — turns you from someone who asks for help into someone who can direct AI to build real tools. This page gives you that understanding.

## Scripts: Your New Power Tool

When we say AI can "write code for you," we usually mean it writes **scripts**. But what is a script?

A **script** is a small, self-contained set of instructions that tells a computer to do something specific. Read these files. Calculate these totals. Rename these documents. Fetch this data. Scripts are the simplest kind of code — they run from top to bottom, do their job, and stop.

This is different from a full **program** or **application** — like Microsoft Word, or your company's CRM system, or a mobile app. Programs are large, complex, and take teams of developers months or years to build. They have user interfaces, databases, error handling for thousands of edge cases, and ongoing maintenance.

A script, by contrast, might be twenty lines long and solve one specific problem. That is what makes them so well-suited to AI generation.

**How scripts run: interpretation vs compilation**

There are two ways computers execute code. **Compiled** languages (like C or Java) require a separate step to translate the human-readable code into machine code before it can run. This is how most major software is built — it is faster but more complex.

**Interpreted** languages run directly — the computer reads the script line by line and executes it on the spot. No build step. No compilation. You write it, you run it.

This matters because almost all AI-generated code is in **interpreted scripting languages**, particularly **Python**. Python reads like English, has enormous libraries of pre-built tools, and AI models generate it more reliably than any other language. When you ask ChatGPT or Claude to "write a script," Python is almost always what you get.

```callout
type: tip
title: "Why Python?"
content: "Python was designed to be readable. A line like for file in folder: process(file) is close to plain English. This is also why AI models are so good at writing it — the gap between natural language and Python is smaller than for any other language. You do not need to learn Python. But when you see it in AI output, you will find you can often follow what it is doing."
```

## Data Formats: How Machines Read Data

Scripts need to read and write data. But computers cannot just read a Word document or a PDF the way you do — they need data in **structured formats** where every piece of information has a predictable place.

You already know one: **CSV** (comma-separated values). A CSV file is just rows of data separated by commas. Simple, flat, tabular — perfect for spreadsheet-style data.

But the format that dominates modern scripting and automation is **JSON** (JavaScript Object Notation, pronounced "jason"). JSON is how most software systems store and exchange structured data. Here is what it looks like:

```json
{
  "client": "Müller GmbH",
  "currency": "EUR",
  "amount": 4200.00,
  "items": [
    { "description": "Q4 consulting", "hours": 40 },
    { "description": "Travel expenses", "hours": 0 }
  ]
}
```

The rules are simple:

- **Curly braces** `{ }` hold an **object** — a collection of named fields
- **Square brackets** `[ ]` hold a **list** — multiple items of the same kind
- Every field has a **name** (in quotes) and a **value** (text, number, true/false, another object, or a list)
- Objects can contain other objects, and lists can contain other lists — data nests naturally

You do not need to memorise the syntax. But recognising JSON when you see it is genuinely useful. When AI shows you what it is doing — reading an API response, processing a configuration file, handling structured data — you will see JSON everywhere. Understanding its shape means you can check whether the AI is reading your data correctly.

```callout
type: info
title: "JSON Is Everywhere"
content: "Nearly every web service, every app on your phone, and every cloud platform uses JSON behind the scenes. When you fill in a web form and click submit, your data is almost certainly converted to JSON before it is sent to the server. When an app shows you a weather forecast, that data arrived as JSON. It is the common language of modern software."
```

## APIs: How Software Talks to Software

An **API** (Application Programming Interface) is a structured way for one piece of software to request something from another. Think of it as a service counter: you go to the window, ask a specific question in a specific format, and get a structured answer back.

Every API has **endpoints** — specific addresses for specific things. An exchange rate service might have:

- `/latest` — get today's rates
- `/convert?from=EUR&to=GBP&amount=4200` — convert a specific amount
- `/history?date=2025-01-15` — get rates for a past date

You make a **request** to an endpoint. The API sends back a **response** — almost always as JSON.

For example, asking an exchange rate API for today's rates might return:

```json
{
  "base": "GBP",
  "date": "2026-01-15",
  "rates": {
    "EUR": 1.18,
    "USD": 1.27,
    "JPY": 189.42
  }
}
```

Structured, predictable, machine-readable. A script can fetch this, extract the numbers it needs, and use them in calculations — all automatically.

**Why this matters for you:** Thousands of services offer APIs. Your company's CRM, your helpdesk, your project management tools, government data portals, financial data providers, weather services — all of them have APIs that scripts can talk to. When you ask AI to "pull data from our ticketing system" or "check today's exchange rates," this is what is happening under the hood: a script calling an API, receiving JSON, and processing the result.

```callout
type: info
title: "The Pattern"
content: "Script makes request → API returns JSON → Script reads JSON → Script produces output. This is the fundamental loop of modern automation. AI handles the technical details, but understanding this pattern means you can describe what you want with precision."
```

## Bringing It Together

A script is a small set of instructions. JSON is the data format. APIs are how software fetches data from other systems. Put them together and you have the building blocks of automation:

1. **You** describe what you want in plain English
2. **AI** writes a Python script
3. **The script** calls an API, reads the JSON response, processes the data
4. **You** get the result — a report, a spreadsheet, a summary

This is what people mean by "vibe coding." The term was coined by Andrej Karpathy in February 2025 to describe building software by stating your intent and letting AI handle the implementation. It became Collins Dictionary's Word of the Year that same year. Y Combinator reported 25% of startups in its Winter 2025 batch had codebases that were 95% AI-generated.

But for non-developers, the real revolution is not building startups. It is the fact that **scripts are no longer out of reach.** A thirty-line Python script that connects to an API, processes JSON, and produces a report used to require a developer. Now it requires a clear description of what you want.

## Scripts, APIs, and JSON in Action

Watch all three concepts come together. A non-programmer needs to convert overseas invoices to pounds using live exchange rates. No coding knowledge — just a clear description of the outcome.

```agent
id: api-scripting-demo
title: "Scripts, APIs, and JSON in Action"
model_label: "Claude"

system: |
  You are a coding assistant helping non-programmers automate tasks.
  Write Python scripts that are simple and well-commented. Explain
  what you are doing in plain English at each step.

scratchpad:
  "overseas-invoices.json": |
    {
      "invoices": [
        {
          "id": "INV-2024-089",
          "client": "Müller GmbH",
          "currency": "EUR",
          "amount": 4200.00,
          "description": "Q4 consulting"
        },
        {
          "id": "INV-2024-091",
          "client": "Tanaka Industries",
          "currency": "JPY",
          "amount": 850000,
          "description": "Software licence"
        },
        {
          "id": "INV-2024-093",
          "client": "Brightline Inc",
          "currency": "USD",
          "amount": 7500.00,
          "description": "Annual support contract"
        },
        {
          "id": "INV-2024-094",
          "client": "Dubois & Associés",
          "currency": "EUR",
          "amount": 3100.00,
          "description": "Training delivery"
        }
      ]
    }

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
    text: "The user has a JSON file containing invoices in different currencies — euros, yen, and dollars. They need everything converted to GBP for their finance report. Watch how a script, an API, and JSON work together."

  - type: user
    content: "I've got overseas invoices in a JSON file — different currencies. Can you pull today's exchange rates and convert everything to pounds for our finance report?"

  - type: tool_call
    tool: scratchpad_read
    args:
      filename: "overseas-invoices.json"

  - type: tool_result
    tool: scratchpad_read
    content: |
      {
        "invoices": [
          {
            "id": "INV-2024-089",
            "client": "Müller GmbH",
            "currency": "EUR",
            "amount": 4200.00,
            "description": "Q4 consulting"
          },
          {
            "id": "INV-2024-091",
            "client": "Tanaka Industries",
            "currency": "JPY",
            "amount": 850000,
            "description": "Software licence"
          },
          {
            "id": "INV-2024-093",
            "client": "Brightline Inc",
            "currency": "USD",
            "amount": 7500.00,
            "description": "Annual support contract"
          },
          {
            "id": "INV-2024-094",
            "client": "Dubois & Associés",
            "currency": "EUR",
            "amount": 3100.00,
            "description": "Training delivery"
          }
        ]
      }

  - type: note
    text: "This is JSON — the structured data format. Notice the curly braces holding objects, square brackets holding the list, and every field clearly named. A script can read this programmatically because every piece of data has a predictable location."

  - type: assistant
    content: "I can see four invoices in three different currencies: EUR, JPY, and USD. I'll write a Python script that reads this JSON, fetches today's exchange rates from an API, and converts everything to GBP."

  - type: tool_call
    tool: scratchpad_write
    args:
      filename: "convert_invoices.py"
      content: |
        import json
        import urllib.request

        # Step 1: Read our invoice data (JSON file)
        with open("overseas-invoices.json") as f:
            data = json.load(f)

        # Step 2: Call the exchange rate API
        # This fetches live rates from a public API endpoint
        url = "https://api.exchangerate-host.com/latest?base=GBP"
        response = urllib.request.urlopen(url)
        rates = json.loads(response.read())

        print("Exchange rates from API:")
        print(f"  1 GBP = {rates['rates']['EUR']} EUR")
        print(f"  1 GBP = {rates['rates']['USD']} USD")
        print(f"  1 GBP = {rates['rates']['JPY']} JPY")
        print()

        # Step 3: Convert each invoice to GBP
        print("Invoice Conversion Report")
        print("=" * 55)
        total_gbp = 0

        for inv in data["invoices"]:
            rate = rates["rates"][inv["currency"]]
            gbp = inv["amount"] / rate
            total_gbp += gbp
            print(f"{inv['id']}  {inv['client']:<22} "
                  f"{inv['currency']} {inv['amount']:>10,.0f}"
                  f"  →  GBP {gbp:>9,.2f}")

        print("=" * 55)
        print(f"{'TOTAL':>40}  GBP {total_gbp:>9,.2f}")

  - type: tool_result
    tool: scratchpad_write
    content: "Written to convert_invoices.py (27 lines)"

  - type: note
    text: "Look at what the script does. Line by line: it reads the JSON file, calls an exchange rate API (making an HTTP request to a specific endpoint), reads the JSON that the API sends back, then loops through each invoice converting the amounts. Three concepts — script, API, JSON — working together in 27 lines."

  - type: tool_call
    tool: run_code
    args:
      filename: "convert_invoices.py"

  - type: tool_result
    tool: run_code
    content: |
      Exchange rates from API:
        1 GBP = 1.18 EUR
        1 GBP = 1.27 USD
        1 GBP = 189.42 JPY

      Invoice Conversion Report
      =======================================================
      INV-2024-089  Müller GmbH            EUR      4,200  →  GBP  3,559.32
      INV-2024-091  Tanaka Industries      JPY    850,000  →  GBP  4,487.12
      INV-2024-093  Brightline Inc         USD      7,500  →  GBP  5,905.51
      INV-2024-094  Dubois & Associés      EUR      3,100  →  GBP  2,627.12
      =======================================================
                                         TOTAL  GBP 16,579.07

  - type: note
    text: "The script fetched live exchange rates from an API (the rates section at the top), then used those rates to convert each invoice. The API returned JSON — the script read it, extracted the rates it needed, and did the maths. All from a plain English request."

  - type: tool_call
    tool: scratchpad_write
    args:
      filename: "finance-report.md"
      content: |
        # Overseas Invoice Summary — January 2026

        **Exchange rates** (source: exchangerate-host API, 15 Jan 2026)
        - 1 GBP = 1.18 EUR | 1.27 USD | 189.42 JPY

        ## Invoices

        | Invoice | Client | Original | GBP Equivalent |
        |---------|--------|----------|----------------|
        | INV-2024-089 | Müller GmbH | €4,200 | £3,559.32 |
        | INV-2024-091 | Tanaka Industries | ¥850,000 | £4,487.12 |
        | INV-2024-093 | Brightline Inc | $7,500 | £5,905.51 |
        | INV-2024-094 | Dubois & Associés | €3,100 | £2,627.12 |

        **Total receivable: £16,579.07**

        ---
        *Rates sourced from live API at time of generation.
        Re-run script for updated figures.*

  - type: tool_result
    tool: scratchpad_write
    content: "Written to finance-report.md (17 lines)"

  - type: assistant
    content: |
      Done. Your four overseas invoices total **£16,579.07** at today's rates.

      The biggest item is Brightline Inc at £5,905.51 (originally $7,500). The Japanese invoice looks large in yen but converts to £4,487.12.

      I've saved a formatted report to finance-report.md. The script is reusable — next month, just run it again with updated invoices and it will fetch fresh exchange rates automatically.

  - type: note
    text: "The user described what they wanted in one sentence. The AI wrote a script that: read a JSON file, called an exchange rate API, parsed the JSON response, did the currency conversions, and produced both a console report and a formatted document. Script + API + JSON = automation."

  - type: user
    content: "This is brilliant. Can you turn this into a proper skill? I want to just say 'convert the overseas invoices' next month and have it all happen automatically."

  - type: note
    text: "Now watch the meta-tooling pattern from the Skills page in action. The user is asking AI to turn a one-off task into a reusable, self-contained skill — complete with instructions, the script, and a template for the output."

  - type: tool_call
    tool: create_folder
    args:
      path: "Skills"

  - type: tool_result
    tool: create_folder
    content: "Created folder: Skills"

  - type: tool_call
    tool: create_folder
    args:
      path: "Skills/invoice-converter"

  - type: tool_result
    tool: create_folder
    content: "Created folder: Skills/invoice-converter"

  - type: tool_call
    tool: scratchpad_write
    args:
      filename: "Skills/invoice-converter/SKILL.md"
      content: |
        # Skill: Overseas Invoice Converter

        ## Purpose
        Convert overseas invoices from multiple currencies to GBP
        using live exchange rates. Produce a formatted finance report.

        ## When to Use
        Monthly, or whenever new overseas invoices arrive.

        ## Steps
        1. Read the invoice JSON file (user provides the path)
        2. Run `convert_invoices.py` to fetch live exchange rates
           from the API and compute GBP equivalents
        3. Generate a formatted report in Markdown

        ## Required Input
        - A JSON file of invoices. Each invoice must have:
          id, client, currency (3-letter code), amount, description
        - See overseas-invoices.json for the exact format
        - Supported currencies: any available on exchangerate-host
          (EUR, USD, JPY, CHF, AUD, CAD, etc.)

        ## Output
        Markdown report following the template in example-report.md:
        - Exchange rates with source and date
        - Table: invoice ID, client, original amount, GBP equivalent
        - Total receivable in GBP

        ## Tone
        Professional. Exact figures only — no rounding or estimates.

  - type: tool_result
    tool: scratchpad_write
    content: "Written to Skills/invoice-converter/SKILL.md (38 lines)"

  - type: tool_call
    tool: move_file
    args:
      source: "convert_invoices.py"
      destination: "Skills/invoice-converter/convert_invoices.py"

  - type: tool_result
    tool: move_file
    content: "Moved convert_invoices.py → Skills/invoice-converter/convert_invoices.py"

  - type: tool_call
    tool: move_file
    args:
      source: "finance-report.md"
      destination: "Skills/invoice-converter/example-report.md"

  - type: tool_result
    tool: move_file
    content: "Moved finance-report.md → Skills/invoice-converter/example-report.md"

  - type: assistant
    content: |
      Done. Your skill folder is ready:

      **Skills/invoice-converter/**
      - `SKILL.md` — Instructions: purpose, input format, steps, output spec
      - `convert_invoices.py` — The script that calls the exchange rate API
      - `example-report.md` — Today's report as a reference template

      Next month, just say *"Convert the overseas invoices"* and drop in your updated JSON file. AI reads the skill, runs the script, fetches fresh rates, and gives you the report. Same format, same quality, zero effort.

  - type: note
    text: "One conversation produced a working script. One follow-up question turned it into a permanent, reusable skill. This is the meta-tooling pattern: use AI to build the tools you'll use with AI. The skill folder is self-contained — instructions, code, and reference material all in one place."

  - type: note
    text: "The compound effect: this skill will be used every month. Each use saves time and produces consistent output. And if the requirements change — new currencies, different report format, a different API — just describe the change and AI updates the skill. You never touch the code yourself."
```

## What's Realistic

**Python is the recommended language** for AI-assisted automation. Models generate and debug it most effectively.

**What works well:**
- Personal automation scripts — file processing, data transformation, report generation
- Connecting services via APIs — pulling data from one system into another
- Data analysis and visualisation
- Format conversion and batch processing
- Departmental dashboards and internal tools

**What has limitations:**
Production software. Anything customer-facing. Anything requiring security hardening or scale.

The common pattern for success: **clear inputs, clear outputs, limited edge cases.**

```callout
type: info
title: "The Scale Paradox"
content: "Building one automation script takes minutes. Building ten takes an hour. Building a hundred that are maintainable, documented, and don't break? That still requires engineering expertise. Scripts are tools for solving specific problems, not a replacement for software engineering."
```

```quiz
id: what-is-a-script
type: multiple-choice
question: "What distinguishes a script from a full application like Microsoft Word or a CRM system?"
options:
  - "Scripts are always written in Python; applications use other languages"
  - "Scripts are small, focused sets of instructions for specific tasks; applications are large, complex systems built by teams over months"
  - "Scripts run faster than applications"
  - "Scripts don't need a computer to run"
answer: 1
explanation: "A script is a small, self-contained set of instructions — often just tens of lines — that solves a specific problem. Applications like Word or Salesforce are massive, complex systems with user interfaces, databases, and thousands of edge cases. The simplicity of scripts is exactly what makes them well-suited to AI generation."
```

```quiz
id: json-api-pattern
type: multiple-choice
question: "When a script 'calls an API and processes the JSON response,' what is actually happening?"
options:
  - "The script is downloading a website and reading the HTML"
  - "The script is sending a structured request to another system's endpoint, receiving structured data (JSON) back, and extracting the fields it needs"
  - "The script is converting a PDF to a spreadsheet"
  - "The script is running a database query on your local machine"
answer: 1
explanation: "An API is a structured interface that one piece of software uses to request data from another. The script sends a request to a specific endpoint (like /latest for exchange rates), receives a JSON response containing structured data, and then extracts and processes the fields it needs. This script → API → JSON pattern is the foundation of modern automation."
```

```quiz
id: vibe-coding-scope
type: multiple-choice
question: "What's the realistic scope for AI-generated scripts?"
options:
  - "Full production applications for customers"
  - "Enterprise-scale systems with complex security requirements"
  - "Personal automation, API integrations, data processing, and departmental tools with clear inputs and outputs"
  - "Any software project regardless of complexity"
answer: 2
explanation: "AI-generated scripts work well for personal automation, connecting services via APIs, data processing, and departmental tools — tasks with clear inputs and outputs. They have real limitations for production software, customer-facing applications, and systems requiring security hardening at scale."
```
