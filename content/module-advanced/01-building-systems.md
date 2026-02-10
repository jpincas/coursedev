---
title: "Building Systems, Not Just Prompts"
duration: "10m"
tags: [advanced, systems, workflows]
---

# How the Most Effective Users Work

The most effective AI users aren't just writing better prompts. They're building systems.

## Beyond One-Off Prompts

Most people use AI like this:
1. Have a task
2. Write a prompt
3. Get a result
4. Repeat from scratch next time

Effective users do this:
1. Have a task
2. Check if they have a reusable solution
3. If not, create one
4. Use it now and forever after

## The Three Pillars

**Skills**
Reusable procedures that capture how you do things. Write once, AI follows forever.

**Subagents**
AI that spawns helper agents for parallel work. Complex tasks broken down and tackled simultaneously.

**Connectors (MCP)**
Connections that let AI interact directly with your tools and systems.

![Three Pillars: Skills, Subagents, and MCP Connectors leading to Consistent Results](/content/module-advanced/images/three-pillars.svg)

## The Compounding Effect

Each piece of infrastructure you build pays dividends:
- A skill you document saves time on every future task of that type
- Preferences you set persist across all future conversations
- Connections you establish remain available

This is the difference between using AI and building with AI.

```callout
type: info
title: "The Investment"
content: "Setting up skills, preferences, and connections takes initial effort. But that investment compounds — every future task benefits from the infrastructure you've built."
```

## Token Economics

Tokens are the new currency. Understanding the economics matters.

**The consumption problem:**
Agents consume **100x more tokens** than simple chat. Why? Because the entire conversation history is resent with every message in stateless APIs.

**The pricing trajectory:**
Token pricing has dropped from **$20 per million tokens** in late 2022 to roughly **$0.40 per million** by August 2025.

But consumption has exploded. Working with agents can mean sending tens of millions of tokens per day.

**Cost optimisation strategies:**

**Prompt caching** — Place static content (system prompts, reference docs) at the start. Caching saves **60-80%** on repeated content.

**Model cascading** — Use cheap models for simple tasks, premium models for complex ones. Reduces costs **30-50%**.

**Batch processing** — Most platforms offer **50% discount** for batch API requests that don't need immediate responses.

**Fine-tuning** — For high-volume stable workloads, fine-tuned models can be more cost-effective than prompt engineering.

```callout
type: note
title: "Deloitte's Guidance"
content: "Business leaders should treat AI economics with the same rigour as energy or capital allocation, recognising tokens as the new currency."
```

For enterprises deploying agents at scale, token costs become a genuine line item. The cheapest solution is often not the best solution — but understanding the tradeoffs matters.

```quiz
id: advanced-systems
type: multiple-choice
question: "What separates effective AI users from casual users?"
options:
  - "They write longer, more detailed prompts"
  - "They build reusable systems instead of one-off prompts"
  - "They use more expensive AI models"
answer: 1
explanation: "Effective AI users build infrastructure — skills, preferences, connections — that makes every future task easier. They invest in reusable systems rather than starting from scratch each time."
```

```quiz
id: token-economics
type: multiple-choice
question: "Why do agentic workflows consume 100x more tokens than simple chat?"
options:
  - "The entire conversation history is resent with every message in stateless APIs"
  - "Agents use more complex language models"
  - "Agents use multiple models simultaneously"
answer: 0
explanation: "Stateless APIs mean the entire conversation history gets resent with every message. As conversations grow longer and agents take more actions, token consumption explodes — 100x higher than simple chat interactions."
```
