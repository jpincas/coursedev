---
title: "Context Decay and Fresh Starts"
duration: "10m"
tags: [context-decay, conversations, best-practices]
---

# Why Long Conversations Degrade

Here's a common frustration: AI seems to get worse the longer you talk to it. This isn't your imagination. It's context decay.

## The Degradation Pattern

![Context Decay Over Conversation Turns](/content/module-context/images/context-decay.svg)

**Turns 1-10: Fresh and Clear**
- Context is fresh
- Instructions are clear and uncontradicted
- Output quality is high

**Turns 20-30: Starting to Blur**
- Context is filling up
- Earlier details are competing with newer ones
- Occasional confusion or inconsistency

**Turns 50+: Saturated**
- Context is cluttered or truncated
- Contradictions have accumulated
- Quality noticeably drops

## Why This Happens

Several factors contribute:

**Information density increases**
More messages mean more information competing for attention. Important early instructions get diluted.

**Contradictions accumulate**
You might have said "use formal tone" early on, then casually said "yeah just make it casual" later. Both are in context.

**Context truncation**
Very long conversations may get truncated — older messages dropped to fit newer ones. Your original instructions might disappear.

**Attention dilution**
The model has limited attention. With more in context, less attention goes to any single piece.

```callout
type: warning
title: "The Symptom"
content: "When the model starts 'forgetting' things you told it earlier, or contradicting previous outputs, you're experiencing context decay."
```

## The Solution: Fresh Starts

Don't fight context decay. Work with it.

**New task? New conversation.**

When you're shifting to a different task or topic, start a fresh conversation. You'll get:
- Clean context
- Full attention on your new instructions
- No contradictions from previous exchanges

```callout
type: tip
title: "Scope Each Conversation"
content: "Treat conversations like work sessions. One project or feature per conversation. Use external state files (progress notes, test results, git logs) rather than relying purely on conversation memory."
```

## Compaction: What Actually Happens

When context gets too long, the system can **compact** the conversation -- replacing the full history with a condensed summary.

Here is what happens mechanically:

1. The system takes your entire conversation history (potentially hundreds of messages)
2. It generates a summary capturing the key decisions, context, and current state
3. The full history is replaced with this summary as a single message at the top
4. Your persistent instructions (CLAUDE.md, project knowledge) are preserved -- they reload fresh
5. The conversation continues with clean context but essential knowledge retained

The effect: you get the benefits of a fresh start (clean attention, no contradictions) while keeping the critical context from your previous work.

```callout
type: info
title: "Automatic vs Manual Compaction"
content: "Claude Cowork auto-compacts when conversations get long -- you may never notice it happening. Claude Code gives you manual control with the /compact command, where you can specify what to preserve. Both achieve the same thing: fresh context with key information retained."
```

## Conversation Management Tools

Claude Code offers specific conversation management commands:

**`/clear`** -- Start completely fresh
- Clears conversation history
- Reloads CLAUDE.md (your persistent instructions remain)
- Use when switching to unrelated work

**`/compact`** -- Summarise and continue
- Compresses conversation history into a summary
- You specify what to keep focus on
- Use when context is cluttered but you need continuity

**`/rewind`** -- Selectively roll back
- Remove recent messages
- Restore to an earlier state
- Use when the conversation went off-track

**Principle:** Scope each conversation to one project or feature so context stays relevant.

## When to Preserve Context

Sometimes you do want conversation continuity:
- Iterating on the same document
- Building on previous work in the same session
- Asking follow-up questions about the same topic

In these cases, keep going. But be aware that quality may degrade over very long sessions.

## Watching Context Decay (and the Fix)

This demo simulates a long conversation that degrades, then shows the compaction technique that fixes it.

```agent
id: context-decay-demo
title: "Decay and Recovery"
model_label: "Claude"

system: |
  You are a writing assistant helping with a company blog post.
  Follow the user's style preferences exactly.

scratchpad:
  "brand-voice.txt": |
    Pinnacle Consulting — Brand Voice Guide
    Tone: Authoritative but warm. Think "trusted advisor."
    Avoid: Jargon, buzzwords, passive voice.
    Always: Use concrete examples. Short paragraphs.
    Audience: Senior executives, non-technical.

tools:
  - scratchpad_read
  - scratchpad_write

visibility:
  system_prompt: hidden
  tool_calls: visible
  full_context: toggleable
  token_count: visible

sidebar:
  width: "45%"
  start_open: true

script:
  - type: note
    text: "This simulates what happens over a long conversation. We'll start with clear instructions, build up context, then see what happens when things get cluttered."

  - type: user
    content: "Read our brand voice guide, then write the opening paragraph of a blog post about AI adoption in professional services. Keep it short and punchy — no buzzwords."

  - type: tool_call
    tool: scratchpad_read
    args:
      filename: "brand-voice.txt"

  - type: tool_result
    tool: scratchpad_read
    content: |
      Pinnacle Consulting — Brand Voice Guide
      Tone: Authoritative but warm. Think "trusted advisor."
      Avoid: Jargon, buzzwords, passive voice.
      Always: Use concrete examples. Short paragraphs.
      Audience: Senior executives, non-technical.

  - type: assistant
    content: "Most consulting firms talk about AI. Few have changed how they actually work. At Pinnacle, we started small — one partner used Claude to draft client proposals. Within three months, the entire team had adopted it. Not because we mandated it. Because the results were impossible to ignore."

  - type: note
    text: "Good output. Clear, punchy, no buzzwords, matches the brand voice. The context is clean. Now imagine 40 more messages — topic changes, contradictory edits, tangential discussions. We'll simulate what that accumulated context looks like."

  - type: user
    content: "Actually, make it more formal. Wait — no, keep it conversational but add more data. Also, can we mention the McKinsey study? And maybe reference our Q3 results? Let's make it longer. Actually shorter is better. Focus on the human angle."

  - type: assistant
    content: "The professional services industry stands at a pivotal inflection point in its digital transformation journey. According to a recent McKinsey study, firms that strategically leverage AI-powered solutions are seeing unprecedented efficiency gains across their operational workflows. At Pinnacle Consulting, our Q3 results demonstrate the tangible ROI of embracing cutting-edge AI capabilities, with our human-centered approach driving meaningful synergies between technology adoption and workforce empowerment."

  - type: note
    text: "Look what happened. The contradictory instructions ('more formal' / 'conversational' / 'add data' / 'shorter' / 'longer' / 'human angle') confused everything. The output is full of the exact buzzwords the brand guide prohibits: 'inflection point,' 'digital transformation,' 'leverage,' 'synergies,' 'empowerment.' The original clear instructions got buried under accumulated noise."

  - type: note
    text: "This is context decay. Now watch the fix — we compact the context, keeping only what matters."

  - type: compaction
    summary: |
      [Earlier: User asked for a blog post opening about AI adoption in professional services.
       Brand voice guide loaded: authoritative but warm, no jargon, short paragraphs, concrete examples.
       First draft was good but user gave contradictory revision instructions that degraded quality.
       Current need: restart the paragraph with clean, focused instructions.]

  - type: note
    text: "Context was just compacted. All the contradictory back-and-forth is gone — replaced with a clean summary of what matters. The token counter dropped. Now watch the difference."

  - type: user
    content: "Let's try again. Write the opening paragraph. Follow the brand voice guide: short, punchy, concrete examples, no buzzwords. Authoritative but warm."

  - type: assistant
    content: "Last year, 73% of consulting firms said they were 'exploring AI.' This year, the ones who actually adopted it are pulling ahead. We've seen it firsthand at Pinnacle — teams that use AI for first-draft proposals, research synthesis, and client analysis are delivering work 40% faster. Not because the technology is magic. Because it handles the tedious parts, so consultants can focus on what they're actually good at: judgment, relationships, and strategy."

  - type: note
    text: "Night and day. Compaction cleared the noise, and clean instructions restored quality. This is exactly what /compact does in Claude Code — it summarises the conversation history so you can continue with fresh context without starting completely over."
```

```quiz
id: context-decay-solution
type: multiple-choice
question: "After 30 messages, AI starts contradicting its earlier analysis. You still need continuity for this task. What is the best approach?"
options:
  - "Repeat your original instructions in a new message to override the contradictions"
  - "Compact the conversation to preserve key context while clearing the accumulated noise"
  - "Start a completely fresh conversation and re-provide all the necessary context"
answer: 1
explanation: "When you need continuity but context has decayed, compaction is the best tool. It preserves the essential context while clearing contradictions and noise. Starting fresh loses valuable context. Repeating instructions adds to the clutter without removing the contradictions."
```
