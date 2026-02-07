package main

import (
	"encoding/json"
	"fmt"
)

// processEventGroup processes a user-initiated group: the user message
// and all subsequent tool_call, tool_result, assistant events until
// hitting a boundary (note, user, compaction, clear, or end of script).
// Returns the new script index.
func processEventGroup(agent *AgentState, startIdx int) int {
	script := agent.Config.Script
	idx := startIdx
	groupPosition := 0

	// Reset GroupIdx on all existing messages so they don't re-animate
	for i := range agent.ChatMessages {
		agent.ChatMessages[i].GroupIdx = -1
	}

	for idx < len(script) {
		event := script[idx]

		switch event.Type {
		case "user":
			if idx > startIdx {
				return idx // Next user message = new group
			}
			appendUserEvent(agent, event, groupPosition)
			groupPosition++
			idx++

		case "assistant":
			appendAssistantEvent(agent, event, groupPosition)
			groupPosition++
			idx++

		case "tool_call":
			appendToolCallEvent(agent, event, groupPosition)
			groupPosition++
			idx++

		case "tool_result":
			appendToolResultEvent(agent, event, groupPosition)
			applyToolSideEffects(agent, script, idx)
			groupPosition++
			idx++

		default:
			return idx // Boundary: note, compaction, clear
		}
	}

	return idx
}

func appendUserEvent(agent *AgentState, event ScriptEvent, groupIdx int) {
	agent.ChatMessages = append(agent.ChatMessages, ChatMessage{
		Type:     "user",
		Content:  event.Content,
		GroupIdx: groupIdx,
	})
	agent.ContextMessages = append(agent.ContextMessages, ContextMessage{
		Role:    "user",
		Content: event.Content,
	})
	agent.TokenCount += estimateTokens(event.Content)
}

func appendAssistantEvent(agent *AgentState, event ScriptEvent, groupIdx int) {
	agent.ChatMessages = append(agent.ChatMessages, ChatMessage{
		Type:     "assistant",
		Content:  event.Content,
		GroupIdx: groupIdx,
	})
	agent.ContextMessages = append(agent.ContextMessages, ContextMessage{
		Role:    "assistant",
		Content: event.Content,
	})
	tokens := event.Tokens
	if tokens == 0 {
		tokens = estimateTokens(event.Content)
	}
	agent.TokenCount += tokens
}

func appendToolCallEvent(agent *AgentState, event ScriptEvent, groupIdx int) {
	agent.ChatMessages = append(agent.ChatMessages, ChatMessage{
		Type:     "tool_call",
		ToolName: event.Tool,
		ToolArgs: event.Args,
		GroupIdx: groupIdx,
	})
	argsJSON, _ := json.Marshal(event.Args)
	agent.ContextMessages = append(agent.ContextMessages, ContextMessage{
		Role:    "tool_call",
		Content: fmt.Sprintf("%s(%s)", event.Tool, string(argsJSON)),
	})
	agent.TokenCount += estimateTokens(fmt.Sprintf("%s(%s)", event.Tool, string(argsJSON)))
}

func appendToolResultEvent(agent *AgentState, event ScriptEvent, groupIdx int) {
	agent.ChatMessages = append(agent.ChatMessages, ChatMessage{
		Type:     "tool_result",
		ToolName: event.Tool,
		Content:  event.Content,
		GroupIdx: groupIdx,
	})
	agent.ContextMessages = append(agent.ContextMessages, ContextMessage{
		Role:    "tool_result",
		Content: event.Content,
	})
	agent.TokenCount += estimateTokens(event.Content)
}

func applyToolSideEffects(agent *AgentState, script []ScriptEvent, idx int) {
	event := script[idx]
	if event.Tool != "scratchpad_write" {
		return
	}
	// Find the preceding tool_call to get the write args
	for i := idx - 1; i >= 0; i-- {
		if script[i].Type == "tool_call" && script[i].Tool == "scratchpad_write" {
			filename := script[i].Args["filename"]
			content := script[i].Args["content"]
			if agent.Scratchpad == nil {
				agent.Scratchpad = make(map[string]string)
			}
			agent.Scratchpad[filename] = content
			return
		}
	}
}

func processCompaction(agent *AgentState, event ScriptEvent) {
	// Keep system prompt, replace everything else with summary
	var systemMsg ContextMessage
	if len(agent.ContextMessages) > 0 {
		systemMsg = agent.ContextMessages[0]
	} else {
		systemMsg = ContextMessage{Role: "system", Content: agent.Config.System}
	}

	agent.ContextMessages = []ContextMessage{
		systemMsg,
		{Role: "compaction_summary", Content: event.Summary},
	}
	agent.ChatMessages = append(agent.ChatMessages, ChatMessage{
		Type:     "compaction_divider",
		Content:  event.Summary,
		GroupIdx: -1,
	})
	agent.TokenCount = estimateTokens(agent.Config.System) + estimateTokens(event.Summary)
}

func processClear(agent *AgentState, event ScriptEvent) {
	agent.ChatMessages = nil
	agent.ContextMessages = []ContextMessage{
		{Role: "system", Content: agent.Config.System},
	}
	agent.TokenCount = estimateTokens(agent.Config.System)

	if event.Note != "" {
		agent.ChatMessages = append(agent.ChatMessages, ChatMessage{
			Type:     "clear_divider",
			Content:  event.Note,
			GroupIdx: -1,
		})
	}

	if event.ResetScratchpad {
		agent.Scratchpad = make(map[string]string)
		for k, v := range agent.InitialScratchpad {
			agent.Scratchpad[k] = v
		}
	}
}
