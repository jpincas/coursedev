---
title: "Quick Reference"
duration: "5m"
tags: [reference, glossary, summary]
---

# Quick Reference

A summary of key concepts for easy reference.

## Glossary

**LLM (Large Language Model)**
The AI technology behind Claude, GPT, and similar systems. Fundamentally a prediction engine.

**Context window**
Everything the model can "see" when generating a response. Includes system prompt, conversation history, files, and your message.

**Token**
A piece of text the model processes. Roughly 3/4 of a word, or about 4 characters.

**System prompt**
Hidden instructions that shape AI behaviour. Set by the application, often invisible to users.

**Hallucination**
When AI generates plausible-sounding but false information. Inherent to how prediction works.

**Grounding**
Providing specific documents and data for AI to reference. Reduces hallucination, improves accuracy.

**Skill**
A documented procedure AI can follow automatically. Instructions, templates, examples bundled together.

**Subagent**
A helper AI spawned by the main agent for parallel work.

**Persistent memory**
Information about you that AI loads automatically in every session. Preferences, context, standing instructions.

**MCP (Model Context Protocol)**
Standard for connecting AI to external tools and systems.

## The Three Questions

Before any task:

1. **What does "done" look like?**
   Define the end state clearly.

2. **What context does it need?**
   What information is required for success?

3. **What are the boundaries?**
   What constraints or limits apply?

## Basic Prompt Structure

**Role** — Who should AI be?
**Task** — What should it do?
**Context** — What background does it need?
**Format** — How should output be structured?
**Constraints** — What limits apply?

## Verification Checklist

- [ ] Spot-check 3-5 specific facts
- [ ] Check for internal coherence
- [ ] Validate structure matches requirements
- [ ] Test against intended purpose

## Don't Delegate

**Judgment** — Ethical decisions, value tradeoffs
**Relationships** — Sensitive conversations, trust-building
**Accountability** — Final approval, sign-off

## The Five Themes

1. **Outcomes over process** — Define done, not how
2. **Context is everything** — Rich input = rich output
3. **AI-first, human-verified** — Delegate, then verify
4. **Preparation is the new execution** — Thinking is value
5. **Files, not chat** — Deliverables, not conversations

## The Workflow

```
PREPARE → DELEGATE → VERIFY → DELIVER
   ↑______________|
      (iterate)
```

## Model Comparison: When To Use Which Tool

**Claude (Anthropic)**
- Strengths: Long context (up to 1M tokens), instruction following, coding, safety, structured output
- Best for: Complex multi-file tasks, codebase analysis, technical writing, tasks requiring extended context

**GPT (OpenAI)**
- Strengths: Multimodal capabilities, ecosystem integrations, plugins, broad knowledge base
- Best for: Tasks mixing text/images/audio, quick iterations, general knowledge queries, integration-heavy workflows

**Gemini (Google)**
- Strengths: Google ecosystem integration, 2M token context window, multimodal understanding
- Best for: Tasks involving Google Workspace, extremely long documents, multimodal analysis

**This is not a product comparison.** Each tool has specific strengths. The question isn't "which is best" — it's "which matches this task?" Use the right tool for the job.

```quiz
id: final-quiz
type: multiple-choice
question: "What's the fundamental nature of LLMs that explains both their capabilities and their limitations?"
options:
  - "They are knowledge databases that store and retrieve facts"
  - "They are prediction engines that generate probable next tokens"
  - "They are reasoning systems that solve problems logically"
answer: 1
explanation: "LLMs are prediction engines — they predict what text comes next based on patterns learned from training data. This explains their strengths (excellent pattern continuation, natural language generation) and limitations (hallucinations, no access to ground truth, no real-time knowledge). Everything you've learned in this course builds on this foundational understanding."
```

---

You're now equipped to work with AI as a capable partner, not just a chatbot.

Now go build something.
