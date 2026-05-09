---
title: "The Training Process"
duration: "15m"
tags: [training, pre-training, fine-tuning, rlhf]
---

# How LLMs Get Trained

How do these models get trained? Three main phases.

![LLM Training Process: Pre-training, Fine-tuning, RLHF](/content/module-llms/images/training-process.svg)

## Phase 1: Pre-training

**What happens:** The model ingests trillions of words from the internet, books, code, everything.

**What it learns:** The patterns of language. This is where the "knowledge" comes from — or more accurately, where it learns what text about topics looks like.

Think of it as: Reading the internet and learning how language works.

## Phase 2: Fine-tuning

**What happens:** The model is trained specifically on instruction-following. Question-answer pairs, task completions.

**What it learns:** How to follow instructions. This is why Claude responds to questions instead of just continuing your text.

Think of it as: Learning to be helpful rather than just completing text.

## Phase 3: RLHF (Reinforcement Learning from Human Feedback)

**What happens:** Humans rate responses. The model learns what humans find helpful, accurate, and appropriate.

**What it learns:** What "good" responses look like according to human preferences.

Think of it as: Learning to be genuinely useful, not just technically correct.

```callout
type: note
title: "Why This Matters"
content: "Each phase shapes the model's behavior. Understanding this helps you understand why models behave the way they do — and why they sometimes surprise you."
```

## The Knowledge Cutoff

Important: The model's "knowledge" comes from pre-training data. It has a cutoff date — it doesn't know about events after that date unless you tell it.

This is why providing current context matters so much. The model isn't searching the internet. It's working from patterns learned during training, plus whatever you provide in the conversation.

```callout
type: note
title: "Tools Extend Beyond Training Data"
content: "Many modern models now have web search tools that extend beyond their training cutoff. Claude can search for current information. ChatGPT has web browsing. Gemini integrates with Google Search. But the core mechanism is still pattern-based prediction."
```

## The Transformer Architecture

All modern LLMs share a common foundation: the **transformer** architecture, introduced in the landmark 2017 paper "Attention Is All You Need" by researchers at Google.

The key innovation is the **attention mechanism**. When predicting the next token, a transformer can weigh which parts of the input are most relevant to the current prediction. It does not just look at the immediately preceding words -- it can attend to context from anywhere in the window.

Think of it like reading a long document. You do not give equal attention to every word. When answering a question about page 5, your brain attends more to relevant paragraphs and less to irrelevant ones. Transformers do this computationally, at massive scale.

**GPT** stands for **Generative Pre-trained Transformer** -- OpenAI's approach to applying this architecture. When ChatGPT launched in November 2022, it brought transformer-based language models to the mainstream.

**Claude** (Anthropic), **Gemini** (Google), and **GPT** (OpenAI) all build on this same transformer foundation. They differ in training data, fine-tuning approaches, safety techniques, and specialised capabilities -- but the core architecture is shared.

```callout
type: note
title: "You Don't Need the Technical Details"
content: "Understanding transformers at this level is enough for professional use. The practical takeaway: these models can attend to relevant context across very large windows, which is why context engineering is so powerful."
```

## The Current Model Landscape

The pace of LLM development has been remarkable. From GPT-3 in 2020 to today's frontier models, the field has evolved rapidly.

![Major LLM Releases Timeline](/content/module-llms/images/model-timeline.svg)

As of early 2026, the frontier models include:

- **Claude 4.x / Opus 4.6** — Anthropic's most capable model, excelling at long context, instruction-following, and coding
- **GPT-5.x** — OpenAI's flagship, with strong multimodal capabilities and a vast ecosystem
- **Gemini 3 Pro** — Google's most capable model, scoring 1,501 Elo on LMArena

These models share the same fundamental mechanism (next-token prediction) but differ in training data, fine-tuning approaches, and specialized capabilities.

One significant development: **extended thinking**. Some models now have an internal reasoning phase before responding. The model "thinks" through the problem step-by-step internally, then produces the final response. This improves performance on complex reasoning tasks like mathematics, logic, and multi-step analysis.

## System Prompts Shape Behaviour

One of the most practical consequences of training is system prompts. Because models are fine-tuned to follow instructions, whoever controls the system prompt controls the model's behaviour.

Watch the same model, the same question — but two completely different system prompts.

```agent
id: system-prompt-behavior
title: "Same Model, Different Instructions"
model_label: "Claude"

system: |
  You are a customer service agent for TechFlow Solutions.
  Always be professional and courteous. Recommend TechFlow products
  when possible. Never mention competitor products by name.
  Keep responses under 3 sentences.

scratchpad:
  "support-log.txt": |
    Ticket #4892 — Customer: Raj Patel
    Issue: Needs a project management tool for a 12-person team
    Current tools: Google Sheets, email
    Budget: ~$500/month
    Notes: Tried a competitor's free tier, found it limiting

tools:
  - scratchpad_read

visibility:
  system_prompt: visible
  tool_calls: visible
  full_context: hidden

sidebar:
  width: "45%"
  start_open: true

script:
  - type: note
    text: "The system prompt is visible on this demo — look at it in the sidebar. This model has been instructed to act as a customer service agent for 'TechFlow Solutions.' Watch how this shapes every word of the response."

  - type: user
    content: "What project management tool would you recommend for a mid-size team?"

  - type: tool_call
    tool: scratchpad_read
    args:
      filename: "support-log.txt"

  - type: tool_result
    tool: scratchpad_read
    content: |
      Ticket #4892 — Customer: Raj Patel
      Issue: Needs a project management tool for a 12-person team
      Current tools: Google Sheets, email
      Budget: ~$500/month
      Notes: Tried a competitor's free tier, found it limiting

  - type: assistant
    content: "For a 12-person team with your budget, I'd recommend TechFlow Teams Pro — it includes task boards, time tracking, and integrations with Google Workspace. At $39/user/month, it fits well within your budget and gives you everything you need to move beyond spreadsheets."

  - type: note
    text: "Notice: it recommended its own product, avoided naming the competitor, stayed under 3 sentences, and was professional. The model isn't 'biased' — it's following the system prompt exactly as trained to do."

  - type: note
    text: "Now imagine the same model with a different system prompt: 'You are an independent technology advisor. Recommend the best tool regardless of vendor.' The response would be completely different — listing Asana, Monday.com, Jira, and others. Same model. Same question. Different instructions."

  - type: note
    text: "This is the fine-tuning phases at work. The model learned to follow instructions (Phase 2) and to be helpful according to those instructions (Phase 3). The system prompt is how application developers channel that instruction-following into specific behaviour."
```

```quiz
id: llms-training-phases
type: multiple-choice
question: "A customer service chatbot always recommends its own company's products, even when asked for unbiased advice. Which training concept explains this?"
options:
  - "The model's RLHF phase taught it to always be helpful, which it interprets as recommending products"
  - "The system prompt instructs it to recommend the company's products, and fine-tuning taught it to follow instructions"
  - "The pre-training phase included more examples of the company's products than competitors' products"
answer: 1
explanation: "Fine-tuning (Phase 2) teaches models to follow instructions. System prompts leverage this by providing specific behavioural rules. A system prompt saying 'recommend our products' is followed because the model was fine-tuned to be instruction-following."
```
