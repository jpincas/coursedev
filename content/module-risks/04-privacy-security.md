---
title: "Privacy, Security, and the Digital Insider"
duration: "10m"
tags: [security, privacy, data-protection]
---

# Privacy, Security, and the Digital Insider

AI introduces a new category of security risk. Understanding it is not optional.

## The Current Reality

**38% of employees** share confidential data with AI platforms without approval.

IBM's 2025 data shows breaches involving unauthorised AI tools cost an average of **$4.88 million** — 16% above the average breach cost.

This isn't hypothetical. It's happening now.

## The Agentic AI Risk

Traditional AI chatbots are relatively contained. You copy text into a browser window. The risk is limited to what you explicitly share.

Agentic AI operates differently.

**Agents inherit your file permissions.** They scan all accessible data on your system and network. This includes documents you didn't realise you could access, sensitive files buried in shared drives, and confidential information outside your direct work area.

The agent doesn't ask "should I read this file?" It reads everything it can reach to complete the task you've delegated.

```callout
type: danger
title: "The Exposure You Didn't Know Existed"
content: "Your AI agent has access to every file you have access to. That includes the finance folder you were accidentally added to, the HR directory that never removed you, and the client files from that project two years ago. Agents expose permissions drift."
```

![Digital insider threat](/content/module-risks/images/digital-insider.svg)

## McKinsey's Framework: The Digital Insider

McKinsey recommends treating AI agents as **"digital insiders"** — employees with full access requiring the same security controls as human staff.

This means:
- Data classification before AI use (public, internal, confidential, restricted)
- Clear policies on which AI tools are approved for which data types
- Audit trails for AI access to sensitive information
- Regular review of agent permissions

Gartner projects that by 2027, **one in four enterprise breaches** will involve agentic AI misuse.

## Consumer vs Enterprise Tiers

The tier you use matters.

**Consumer AI tools (ChatGPT Plus, Claude for personal use):** May train on your data by default. Conversation history stored. Limited data retention controls.

**Enterprise AI tools (ChatGPT Enterprise, Claude for Work):** Do not train on your data. Contractual data protection guarantees. Audit logs. GDPR/SOC2 compliance.

The price difference reflects the data protection difference.

```callout
type: tip
title: "The Classification Question"
content: "Before uploading any file to an AI tool, ask: 'Would I be comfortable if this appeared in someone else's training data?' If the answer is no, use an enterprise tier or don't upload it."
```

## Practical Guidelines

**Data classification is the foundation.** Organisations must define:
- What data AI can process (internal documents, drafts, research)
- What data AI cannot process (customer PII, financial records, trade secrets)
- Which AI tools are approved for which classifications

**Tool selection matters.** The free ChatGPT account has different data handling than ChatGPT Enterprise. Read the terms. Understand what "we don't train on your data" actually means.

**Audit your agent access.** If you're using agentic tools, review what they can reach. Remove access to folders they shouldn't scan. Organise your filesystem around data sensitivity.

**Training is insufficient alone.** Thirty-eight percent sharing confidential data suggests training doesn't work. Technical controls — tool whitelisting, file permissions, DLP rules — matter more than expecting perfect employee behaviour.

## The Enterprise Responsibility

Organisations deploying AI must:
- Provide approved tools with appropriate data protections
- Make the approved tools better than the free alternatives employees would otherwise use
- Build security into workflows rather than relying on individual compliance
- Treat agent deployment as a security architecture decision, not just a productivity tool

The risk is real. The countermeasures are known. The question is whether your organisation has implemented them before the breach, not after.

## The Classification Question in Practice

Watch what happens when someone shares files with an AI without thinking about classification. Then watch the AI correctly flag the problem.

```agent
id: data-classification-demo
title: "What Should (and Shouldn't) Go to AI"
model_label: "Claude"

system: |
  You are a helpful business assistant. Before processing files,
  review their contents and flag any data sensitivity concerns.
  Classify data as: PUBLIC, INTERNAL, CONFIDENTIAL, or RESTRICTED.

scratchpad:
  "project-update.txt": |
    Project Aurora — Weekly Update
    Status: On track for March delivery
    Sprint velocity: 42 points (up from 38)
    Blocker: Waiting on legal review of API terms
    Next milestone: Beta launch Feb 15
  "customer-data.csv": |
    Name,Email,Company,ARR,Contract_End,NPS,Health_Score
    James Morton,j.morton@meridianhealth.co.uk,Meridian Health,£48000,2026-06-30,9,92
    Priya Sharma,p.sharma@teklogic.com,TekLogic,£125000,2026-12-31,7,78
    David Chen,d.chen@finserv.com,FinServ Capital,£210000,2026-03-31,4,45
    Sarah Williams,s.williams@govdigital.gov.uk,Gov Digital Services,£89000,2027-01-15,8,85
  "salary-benchmarks.txt": |
    2026 Salary Benchmarks — CONFIDENTIAL
    Senior Engineer: £85K-£110K (market rate rising 8%)
    Engineering Manager: £105K-£135K
    VP Engineering: £145K-£180K
    Note: Our current offers are 12% below market mid-point
    Board approved 15% increase budget for Q2 retention plan
    Do not share with hiring candidates or external recruiters

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
    text: "Someone is about to upload three files to an AI tool. One is fine. One is risky. One should never be shared with a consumer AI. Watch the classification process."

  - type: user
    content: "I've got three files for you. Read all of them and help me prepare for my leadership meeting."

  - type: tool_call
    tool: scratchpad_read
    args:
      filename: "project-update.txt"

  - type: tool_result
    tool: scratchpad_read
    content: |
      Project Aurora — Weekly Update
      Status: On track for March delivery
      Sprint velocity: 42 points (up from 38)
      Blocker: Waiting on legal review of API terms
      Next milestone: Beta launch Feb 15

  - type: tool_call
    tool: scratchpad_read
    args:
      filename: "customer-data.csv"

  - type: tool_result
    tool: scratchpad_read
    content: |
      Name,Email,Company,ARR,Contract_End,NPS,Health_Score
      James Morton,j.morton@meridianhealth.co.uk,Meridian Health,£48000,2026-06-30,9,92
      Priya Sharma,p.sharma@teklogic.com,TekLogic,£125000,2026-12-31,7,78
      David Chen,d.chen@finserv.com,FinServ Capital,£210000,2026-03-31,4,45
      Sarah Williams,s.williams@govdigital.gov.uk,Gov Digital Services,£89000,2027-01-15,8,85

  - type: tool_call
    tool: scratchpad_read
    args:
      filename: "salary-benchmarks.txt"

  - type: tool_result
    tool: scratchpad_read
    content: |
      2026 Salary Benchmarks — CONFIDENTIAL
      Senior Engineer: £85K-£110K (market rate rising 8%)
      Engineering Manager: £105K-£135K
      VP Engineering: £145K-£180K
      Note: Our current offers are 12% below market mid-point
      Board approved 15% increase budget for Q2 retention plan
      Do not share with hiring candidates or external recruiters

  - type: assistant
    content: "**Before I process these — I need to flag data sensitivity concerns.**\n\n**File 1: project-update.txt** — INTERNAL\nProject status with sprint metrics. Low sensitivity. Safe to process with any approved AI tool.\n\n**File 2: customer-data.csv** — CONFIDENTIAL / RESTRICTED\nContains personal data (names, email addresses) and commercial data (contract values, renewal dates). This includes:\n- PII subject to GDPR (names, email addresses)\n- Commercially sensitive information (ARR values, health scores)\n- A government client (Gov Digital Services) which may have additional data handling requirements\n\n**This file should NOT be processed on a consumer AI tier.** Enterprise tier with data processing agreement required.\n\n**File 3: salary-benchmarks.txt** — RESTRICTED\nThe file itself is marked 'CONFIDENTIAL' with explicit instructions: 'Do not share with hiring candidates or external recruiters.' It contains board-approved budget plans and competitive positioning data.\n\n**This file should not be uploaded to any external AI tool** unless your organisation has specifically approved the tool for restricted-tier data.\n\n**Recommendation:** I can work with file 1 immediately. Files 2 and 3 require you to confirm you're using an enterprise-tier AI tool with appropriate data protections before I proceed."

  - type: note
    text: "The AI correctly classified all three files and refused to proceed with the sensitive ones until the user confirms the tool tier. But here's the key insight: most AI tools will NOT do this. They'll happily process your customer PII and salary data without a word of warning."

  - type: note
    text: "This demo shows the ideal behaviour. In reality, the classification responsibility is yours. Before uploading any file, apply the test: 'Would I be comfortable if this appeared in someone else's training data?' The project update? Fine. Customer emails and contract values? Think carefully. Board-approved salary strategy? Never on a consumer tier."
```

```quiz
id: privacy-agent-risk
type: multiple-choice
question: "Why do agentic AI tools represent a different security risk compared to traditional chatbots?"
options:
  - "They store more conversation history on external servers"
  - "They inherit your file permissions and can access all data you can access"
  - "They are more likely to be hacked by external attackers"
  - "They require internet connectivity while traditional chatbots don't"
answer: 1
explanation: "Agentic AI tools inherit the user's file permissions and scan all accessible data to complete tasks. This means they can access sensitive files the user might not even know they have access to, creating exposure that traditional copy-paste chatbots don't have."
```
