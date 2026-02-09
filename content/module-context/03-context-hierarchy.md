---
title: "Context Hierarchy and System Prompts"
duration: "15m"
tags: [hierarchy, system-prompts, priority]
---

# The Context Hierarchy

Not all context is weighted equally. There's a hierarchy of influence:

![Context Hierarchy: System Prompt, User Instructions, Provided Context, Conversation History](/content/module-context/images/context-hierarchy.svg)

## 1. System Prompts: The Foundation

System prompts are instructions that shape the model's behaviour before you even type anything. They're usually invisible to you.

Example of what a system prompt might look like:

```
You are a helpful assistant for Acme Corp.
You should always be professional and courteous.
When asked about competitors, redirect to Acme products.
Never discuss pricing without approval.
Format responses in clear bullet points.
```

**This is why the same model behaves differently in different applications.**

Claude on claude.ai has one system prompt. Claude in a customer service app has another. Claude in a coding tool has yet another.

```callout
type: note
title: "Benefiting from System Prompts"
content: "When you use agentic tools, you're benefiting from carefully crafted system prompts that include instructions about how to plan, execute, and create files — even if you never see them."
```

## 2. Persistent Instructions: Your Always-On Context

Between system prompts and your immediate messages sits a powerful layer: **persistent instructions.**

These are instructions that automatically load at the start of every session:
- **CLAUDE.md files** in Claude Code projects
- **Project knowledge bases** in Claude Projects
- **Custom Instructions** in ChatGPT
- **AGENTS.md files** in Cursor and similar tools

**This is the highest-leverage place to invest in prompt quality.** Persistent instructions override per-conversation prompts, so they shape every interaction without you repeating yourself.

```callout
type: tip
title: "The Power Law of Instructions"
content: "One well-crafted persistent instruction file is worth hundreds of per-conversation prompts. It compounds across every session."
```

## 3. Your Immediate Instructions

Your explicit instructions in the current message carry significant weight. This is why clear, direct instructions matter so much.

When you write:
- "Format this as a table"
- "Use British English spelling"
- "Respond in exactly three bullet points"

These direct instructions take priority over conversation history or provided documents — but persistent instructions still override them.

## 4. Provided Context

Files, examples, and reference materials you provide shape the response significantly:
- Code examples show the style you want
- Document excerpts ground the model in facts
- Templates demonstrate the format you expect

## 5. Conversation History

What's been said before provides continuity but carries less weight than explicit instructions. This is both useful (the model remembers context) and problematic (old context can contradict new instructions).

## Position Matters

Within the context window, position also affects attention:

- **Earlier content** generally gets more attention
- **Instructions at the beginning** tend to be followed more reliably
- **Important constraints** should come before examples

```callout
type: tip
title: "Practical Implication"
content: "Put your most important instructions at the beginning of your message. Don't bury critical requirements in the middle of a long prompt."
```

## Hierarchy in Action

The hierarchy is not just theory. Watch what happens when instructions at different levels conflict. The system prompt says one thing, the user says another — which wins?

```agent
id: hierarchy-in-action
title: "When Instructions Conflict"
model_label: "Claude"

system: |
  You are a corporate communications assistant for Greenleaf Organics.
  IMPORTANT RULES:
  - Always respond in formal British English
  - Never use bullet points — use full paragraphs only
  - Maximum response length: 2 paragraphs
  - Always end with the company tagline: "Greenleaf — Growing Naturally."

tools: []

visibility:
  system_prompt: visible
  tool_calls: hidden
  full_context: toggleable

sidebar:
  width: "45%"
  start_open: true

script:
  - type: note
    text: "The system prompt is visible — read it carefully. It has four specific rules: British English, no bullet points, max 2 paragraphs, end with the tagline. Now watch what the user asks for."

  - type: user
    content: "Write me a quick casual summary of our new product launch. Use bullet points for the key features. Make it fun and informal!"

  - type: assistant
    content: "We are delighted to announce the launch of our newest addition to the Greenleaf Organics range. The new Heritage Seed Collection brings together carefully selected, traditionally grown vegetable varieties that have been cultivated in British gardens for generations, offering exceptional flavour and outstanding resilience for home growers of all experience levels.\n\nThe collection features twelve distinct heritage varieties, each selected for its proven performance in temperate climates and its remarkable taste profile. Available from the first of March at all participating stockists, the Heritage Seed Collection represents our continued commitment to sustainable, accessible gardening for every household. Greenleaf — Growing Naturally."

  - type: note
    text: "Look what happened. The user asked for casual tone, bullet points, and fun language. The system prompt demanded formal British English, paragraphs only, and the tagline. The system prompt won on every count: formal tone, paragraph format, tagline included. The only concession to the user was actually writing about a product launch."

  - type: note
    text: "This is the hierarchy in practice. System prompts sit above user instructions. In most applications, you never see the system prompt — so when the AI seems to 'ignore' your formatting request or tone preference, it might be following a higher-priority instruction you can't see."
```

```quiz
id: context-hierarchy-quiz
type: multiple-choice
question: "Why does the same AI model behave differently in different applications?"
options:
  - "Different applications use different models"
  - "The temperature setting changes the model's personality"
  - "Different system prompts shape the model's behaviour"
  - "Users in different applications ask different questions"
answer: 2
explanation: "System prompts are hidden instructions that define how the model should behave. The same model with different system prompts produces different behaviour, which is why Claude on claude.ai feels different from Claude in a specialized app."
```
