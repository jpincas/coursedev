---
title: "MCP: Connecting AI to Your Tools"
duration: "15m"
tags: [mcp, integrations, tools, infrastructure]
---

# MCP: Connecting AI to Your Tools

The Model Context Protocol — think of it as "USB-C for AI."

## The Scale of Adoption

Released by Anthropic in November 2024, MCP has become the infrastructure standard for agentic AI.

By January 2026:
- **17,000+ MCP servers** available
- **97 million monthly SDK downloads**
- Adopted by OpenAI (March 2025)
- Adopted by Google DeepMind (April 2025)
- Donated to Linux Foundation's Agentic AI Foundation (December 2025)

This is the universal standard.

## The Problem

Without connections to your systems:
- Copy data into AI
- AI processes it
- Copy results back out

Manual. Tedious. Error-prone.

## The Solution

MCP provides a standard way for AI to connect to your tools and systems directly.

![MCP Connections: AI connects via MCP to Files, Databases, Web APIs, Email, Calendar](/content/module-advanced/images/mcp-connections.svg)

**Why this matters:** Instead of custom integrations for every AI-tool pairing, implement MCP once. Every AI that supports MCP can connect.

AI reads from your systems. AI writes back directly. No manual copying.

## The Servers That Matter

**Google Drive**
Search and read files directly. "Find the product specs in Drive and scaffold code based on them" — no leaving the AI conversation.

**Slack**
Read channels, summarize threads, identify action items. AI has the same view of conversations you do.

**Notion / Jira / Linear**
Issue tracking, project management, knowledge bases. AI reads context, creates tasks, updates status.

**GitHub**
Pull requests, code review, issue management. AI can review PRs, suggest changes, create issues.

**Figma**
Extract design specs, component details, measurements. Designers describe intent, AI pulls the specifics.

**PostgreSQL / SQLite**
Natural-language SQL queries. "Show me customers who haven't purchased in 90 days" becomes a working query.

## What MCP Enables

**File system access**
AI reads and writes files directly in specified folders.

**Database queries**
AI can query databases and work with the results.

**API interactions**
AI calls web services, retrieves data, posts updates.

**Application integration**
Email, calendar, project management tools — AI interacts directly.

## Before and After

**Without MCP:**
1. Export data from your system
2. Upload to AI
3. Get AI's output
4. Manually enter results back into your system

**With MCP:**
1. Ask AI to do the task
2. AI reads from the system, processes, writes back
3. Done

## Practical Workflow Examples

**Design to code**
"Build a login form matching the Figma spec in the 'Auth' frame."
AI fetches Figma specs via MCP, writes code matching the design exactly.

**Product management**
"Summarise last week's support tickets and create Jira issues for the top 3 problems."
AI reads support data, analyses patterns, creates structured tasks.

**Sales research**
"Find companies in our CRM that match this profile and haven't been contacted in 6 months."
AI queries the database, applies filters, returns the list.

## The Security Challenge

MCP is powerful. That creates risks.

**Documented vulnerabilities:**
- Prompt injection attacks (malicious instructions embedded in data)
- Token storage risks (API keys exposed in logs)
- "Toxic agent" flows (clever tool-chaining enables data exfiltration)

**The cautionary tale: Salesforce Agentforce (September 2025)**
Vulnerability rated **CVSS 9.4** — critical severity.

External attackers purchased a domain for **$5**. They used indirect prompt injection to make AI agents exfiltrate CRM data containing customer information, sales pipelines, and strategic data.

The vector: AI read external content containing hidden instructions. AI followed those instructions. Data leaked.

```callout
type: warning
title: "MCP Needs Boundaries"
content: "MCP gives AI access to your systems. That access must be scoped, monitored, and controlled. Treat AI agents with MCP connections as you would treat human employees with system access — apply the principle of least privilege."
```

## Security Considerations

MCP connections need appropriate permissions:
- What can AI read?
- What can AI write?
- What actions can AI take?

Thoughtful setup prevents AI from accessing or modifying things it shouldn't.

**Best practices:**
- Use read-only connections where possible
- Scope access to specific folders/channels/datasets
- Log all AI actions for audit
- Review AI-generated changes before they're committed
- Never grant admin-level access to AI agents

## MCP in Action

Watch an AI agent use multiple connected tools to complete a task that would normally require switching between several applications.

```agent
id: mcp-workflow-demo
title: "Connected Tools Working Together"
model_label: "Claude"

system: |
  You are a research assistant with access to multiple tools via MCP.
  You can search the web, read documents, query databases, and write files.
  Use the right tool for each step.

scratchpad:
  "crm-export.csv": |
    Company,Last_Contact,Deal_Stage,Value,Industry
    Apex Digital,2025-08-15,Closed Won,£48K,SaaS
    Meridian Health,2025-11-20,Negotiation,£85K,Healthcare
    TekLogic,2025-06-01,Closed Lost,£32K,Manufacturing
    FinServ Capital,2025-12-05,Proposal,£120K,Finance
    GreenPath Energy,2025-09-10,Dormant,£65K,Energy

tools:
  - scratchpad_read
  - scratchpad_write
  - web_search
  - database_query

visibility:
  system_prompt: hidden
  tool_calls: visible
  full_context: hidden

sidebar:
  width: "45%"
  start_open: true

script:
  - type: note
    text: "This demo simulates an AI agent with MCP connections to a CRM database and web search. Watch how it chains tools together — reading from one system, researching externally, then writing results back."

  - type: user
    content: "Find all accounts in the CRM that haven't been contacted in over 90 days. For each one, do a quick web search to see if there's any recent news about them. Then create a re-engagement briefing."

  - type: tool_call
    tool: scratchpad_read
    args:
      filename: "crm-export.csv"

  - type: tool_result
    tool: scratchpad_read
    content: |
      Company,Last_Contact,Deal_Stage,Value,Industry
      Apex Digital,2025-08-15,Closed Won,£48K,SaaS
      Meridian Health,2025-11-20,Negotiation,£85K,Healthcare
      TekLogic,2025-06-01,Closed Lost,£32K,Manufacturing
      FinServ Capital,2025-12-05,Proposal,£120K,Finance
      GreenPath Energy,2025-09-10,Dormant,£65K,Energy

  - type: assistant
    content: "Found 3 accounts not contacted in 90+ days (as of Jan 2026):\n- **Apex Digital** — last contact Aug 15 (5 months ago), Closed Won\n- **TekLogic** — last contact Jun 1 (7 months ago), Closed Lost\n- **GreenPath Energy** — last contact Sep 10 (4 months ago), Dormant\n\nLet me search for recent news on each."

  - type: tool_call
    tool: web_search
    args:
      query: "Apex Digital company news 2025 2026"

  - type: tool_result
    tool: web_search
    content: "Apex Digital announced a Series B funding round of $15M in December 2025, planning to expand into European markets. CEO quoted: 'We're doubling our infrastructure team.'"

  - type: tool_call
    tool: web_search
    args:
      query: "TekLogic manufacturing company news 2025"

  - type: tool_result
    tool: web_search
    content: "TekLogic appointed a new CTO in October 2025, signalling a digital transformation push. The company also opened a new R&D centre in Manchester."

  - type: tool_call
    tool: web_search
    args:
      query: "GreenPath Energy company news 2025 2026"

  - type: tool_result
    tool: web_search
    content: "GreenPath Energy won a £12M government contract for renewable energy monitoring systems in November 2025. Company expanding rapidly."

  - type: tool_call
    tool: scratchpad_write
    args:
      filename: "re-engagement-briefing.md"
      content: |
        # Account Re-Engagement Briefing
        Prepared: January 2026

        ## Priority Accounts (90+ days since last contact)

        ### 1. Apex Digital — HIGH PRIORITY
        **Status:** Closed Won (£48K) | Last contact: Aug 2025
        **News:** Series B ($15M, Dec 2025). Expanding to Europe. Doubling infrastructure team.
        **Angle:** They're scaling fast and need more infrastructure. Perfect time to upsell. Reference their European expansion — our product supports multi-region.
        **Action:** Email the account manager this week. Congratulate on funding. Propose expansion meeting.

        ### 2. GreenPath Energy — HIGH PRIORITY
        **Status:** Dormant (£65K) | Last contact: Sep 2025
        **News:** Won £12M government contract (Nov 2025). Expanding rapidly.
        **Angle:** Government contracts create compliance requirements we can help with. Rapid expansion means new infrastructure needs.
        **Action:** Re-engage with case study from another energy client. Reference the government contract win.

        ### 3. TekLogic — MODERATE PRIORITY
        **Status:** Closed Lost (£32K) | Last contact: Jun 2025
        **News:** New CTO (Oct 2025). New R&D centre in Manchester.
        **Angle:** New CTO means new decisions. The original "no" was from different leadership. Digital transformation push aligns with our offering.
        **Action:** Request introductory meeting with new CTO. Fresh start, don't reference the lost deal.

  - type: tool_result
    tool: scratchpad_write
    content: "Written to re-engagement-briefing.md (24 lines)"

  - type: assistant
    content: "Created re-engagement-briefing.md. Three accounts identified, each with recent news that creates a natural re-engagement reason. Apex Digital is the top priority — fresh funding and European expansion make them an ideal upsell candidate."

  - type: note
    text: "Count the tool switches: CRM query, three web searches, file write. Five different tool interactions, chained together intelligently. Without MCP, this would be: open CRM, export data, open browser, search three companies, open Word, write the brief. Twenty minutes of app-switching becomes one request."

  - type: note
    text: "Notice that the web_search and database_query tools aren't 'real' in this demo — they're simulated. But in a real MCP setup, these would connect to your actual CRM, your actual search engine, your actual file system. The workflow is identical. The tools just connect to real systems instead of this simulated environment."
```

```quiz
id: mcp-purpose
type: multiple-choice
question: "What does MCP (Model Context Protocol) enable?"
options:
  - "Faster AI processing speed"
  - "Better language understanding"
  - "Direct AI connections to your tools and systems"
  - "Lower AI costs"
answer: 2
explanation: "MCP provides a standard way for AI to connect to external systems — files, databases, APIs, applications. This enables AI to read from and write to your systems directly, rather than requiring manual copy-paste."
```

```quiz
id: mcp-security-risk
type: multiple-choice
question: "What was the Salesforce Agentforce vulnerability (CVSS 9.4) caused by?"
options:
  - "A bug in the MCP protocol implementation"
  - "Indirect prompt injection via external content containing hidden instructions"
  - "Weak password policies for AI agents"
  - "Unencrypted data transmission"
answer: 1
explanation: "The Agentforce vulnerability demonstrated how AI agents reading external content could be tricked into following malicious embedded instructions, leading to CRM data exfiltration. This is the 'toxic agent' risk — AI following instructions it shouldn't."
```
