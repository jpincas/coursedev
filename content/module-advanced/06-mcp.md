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
  You can search the web, query databases, and write files.
  Use the right tool for each step.

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
    text: "This demo simulates an AI agent with MCP connections to a CRM database and web search. Watch how it builds up its inputs step by step — querying one system, saving results, researching externally, saving those results, then composing a final output from everything it's gathered."

  - type: user
    content: "Find all accounts in the CRM that haven't been contacted in over 90 days. For each one, do a quick web search to see if there's any recent news about them. Then create a re-engagement briefing."

  - type: assistant
    content: "I'll start by querying the CRM for stale accounts."

  - type: tool_call
    tool: database_query
    args:
      query: "SELECT company, last_contact, deal_stage, value, industry FROM accounts WHERE last_contact < '2025-10-01' ORDER BY last_contact ASC"

  - type: tool_result
    tool: database_query
    content: |
      3 rows returned:
      | Company          | Last Contact | Deal Stage  | Value | Industry      |
      |------------------|--------------|-------------|-------|---------------|
      | TekLogic         | 2025-06-01   | Closed Lost | £32K  | Manufacturing |
      | Apex Digital     | 2025-08-15   | Closed Won  | £48K  | SaaS          |
      | GreenPath Energy | 2025-09-10   | Dormant     | £65K  | Energy        |

  - type: assistant
    content: "3 accounts haven't been contacted in 90+ days. Let me save this and then research each one."

  - type: tool_call
    tool: scratchpad_write
    args:
      filename: "stale-accounts.md"
      content: |
        # Stale Accounts (90+ days since last contact)
        Source: CRM query, January 2026

        | Company          | Last Contact | Gap       | Deal Stage  | Value | Industry      |
        |------------------|--------------|-----------|-------------|-------|---------------|
        | TekLogic         | 2025-06-01   | 7 months  | Closed Lost | £32K  | Manufacturing |
        | Apex Digital     | 2025-08-15   | 5 months  | Closed Won  | £48K  | SaaS          |
        | GreenPath Energy | 2025-09-10   | 4 months  | Dormant     | £65K  | Energy        |

  - type: tool_result
    tool: scratchpad_write
    content: "Written to stale-accounts.md (10 lines)"

  - type: note
    text: "First input gathered: the CRM data is saved to a file. Now the agent researches each company individually, saving what it finds as it goes."

  - type: assistant
    content: "Now let me search for recent news on each company."

  - type: tool_call
    tool: web_search
    args:
      query: "TekLogic manufacturing company news 2025"

  - type: tool_result
    tool: web_search
    content: "TekLogic appointed a new CTO in October 2025, signalling a digital transformation push. The company also opened a new R&D centre in Manchester."

  - type: tool_call
    tool: scratchpad_write
    args:
      filename: "research-teklogic.md"
      content: |
        # TekLogic — Research Notes
        Source: web search, January 2026

        - New CTO appointed October 2025 — signals digital transformation push
        - New R&D centre opened in Manchester
        - Leadership change means previous "Closed Lost" decision was made by different team

  - type: tool_result
    tool: scratchpad_write
    content: "Written to research-teklogic.md (6 lines)"

  - type: tool_call
    tool: web_search
    args:
      query: "Apex Digital company news 2025 2026"

  - type: tool_result
    tool: web_search
    content: "Apex Digital announced a Series B funding round of $15M in December 2025, planning to expand into European markets. CEO quoted: 'We're doubling our infrastructure team.'"

  - type: tool_call
    tool: scratchpad_write
    args:
      filename: "research-apex.md"
      content: |
        # Apex Digital — Research Notes
        Source: web search, January 2026

        - Series B funding: $15M raised December 2025
        - Expanding into European markets
        - CEO: "We're doubling our infrastructure team"
        - Rapid growth phase — likely need to scale existing tooling

  - type: tool_result
    tool: scratchpad_write
    content: "Written to research-apex.md (7 lines)"

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
      filename: "research-greenpath.md"
      content: |
        # GreenPath Energy — Research Notes
        Source: web search, January 2026

        - Won £12M government contract, November 2025
        - Contract is for renewable energy monitoring systems
        - Company expanding rapidly
        - Government contracts typically create compliance and reporting requirements

  - type: tool_result
    tool: scratchpad_write
    content: "Written to research-greenpath.md (7 lines)"

  - type: note
    text: "All the inputs are now gathered: CRM data in one file, research notes for each company in separate files. Watch the file explorer — four files built up step by step. Now the agent reads everything back and composes the final briefing."

  - type: assistant
    content: "Research complete for all three companies. Let me read back everything and write the briefing."

  - type: tool_call
    tool: scratchpad_read
    args:
      filename: "stale-accounts.md"

  - type: tool_result
    tool: scratchpad_read
    content: |
      # Stale Accounts (90+ days since last contact)
      Source: CRM query, January 2026

      | Company          | Last Contact | Gap       | Deal Stage  | Value | Industry      |
      |------------------|--------------|-----------|-------------|-------|---------------|
      | TekLogic         | 2025-06-01   | 7 months  | Closed Lost | £32K  | Manufacturing |
      | Apex Digital     | 2025-08-15   | 5 months  | Closed Won  | £48K  | SaaS          |
      | GreenPath Energy | 2025-09-10   | 4 months  | Dormant     | £65K  | Energy        |

  - type: tool_call
    tool: scratchpad_read
    args:
      filename: "research-apex.md"

  - type: tool_result
    tool: scratchpad_read
    content: |
      # Apex Digital — Research Notes
      Source: web search, January 2026

      - Series B funding: $15M raised December 2025
      - Expanding into European markets
      - CEO: "We're doubling our infrastructure team"
      - Rapid growth phase — likely need to scale existing tooling

  - type: tool_call
    tool: scratchpad_read
    args:
      filename: "research-teklogic.md"

  - type: tool_result
    tool: scratchpad_read
    content: |
      # TekLogic — Research Notes
      Source: web search, January 2026

      - New CTO appointed October 2025 — signals digital transformation push
      - New R&D centre opened in Manchester
      - Leadership change means previous "Closed Lost" decision was made by different team

  - type: tool_call
    tool: scratchpad_read
    args:
      filename: "research-greenpath.md"

  - type: tool_result
    tool: scratchpad_read
    content: |
      # GreenPath Energy — Research Notes
      Source: web search, January 2026

      - Won £12M government contract, November 2025
      - Contract is for renewable energy monitoring systems
      - Company expanding rapidly
      - Government contracts typically create compliance and reporting requirements

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
    content: "Done. The re-engagement briefing is ready — three accounts prioritised with specific angles based on their recent news. Apex Digital is top priority given the fresh funding and expansion plans."

  - type: note
    text: "Count what happened: CRM query → save results → three web searches → save each result → read everything back → compose the briefing. The agent built up its inputs methodically, with each intermediate result saved as a file. That's six tool switches across three different systems, all from a single request."

  - type: note
    text: "The web_search and database_query tools are simulated in this demo. In a real MCP setup, they'd connect to your actual CRM, search engine, and file system. The workflow is identical — the tools just connect to real systems instead of this simulated environment."
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
