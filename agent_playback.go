package main

import (
	"encoding/json"
	"fmt"
	"strings"
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
	displayName := displayToolName(event.Tool)
	agent.ContextMessages = append(agent.ContextMessages, ContextMessage{
		Role:    "tool_call",
		Content: fmt.Sprintf("%s(%s)", displayName, string(argsJSON)),
	})
	agent.TokenCount += estimateTokens(fmt.Sprintf("%s(%s)", displayName, string(argsJSON)))
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

	// Find the preceding tool_call to get the args
	findPrecedingCall := func(toolName string) map[string]string {
		for i := idx - 1; i >= 0; i-- {
			if script[i].Type == "tool_call" && script[i].Tool == toolName {
				return script[i].Args
			}
		}
		return nil
	}

	if agent.Scratchpad == nil {
		agent.Scratchpad = make(map[string]string)
	}

	switch event.Tool {
	case "scratchpad_write":
		if args := findPrecedingCall("scratchpad_write"); args != nil {
			agent.Scratchpad[args["filename"]] = args["content"]
			// Auto-expand parent folders
			autoExpandFolders(agent, args["filename"])
		}

	case "move_file":
		if args := findPrecedingCall("move_file"); args != nil {
			source := args["source"]
			destination := args["destination"]
			if content, ok := agent.Scratchpad[source]; ok {
				delete(agent.Scratchpad, source)
				agent.Scratchpad[destination] = content
				autoExpandFolders(agent, destination)
			}
		}

	case "create_folder":
		if args := findPrecedingCall("create_folder"); args != nil {
			path := strings.TrimSuffix(args["path"], "/")
			if agent.ExpandedFolders == nil {
				agent.ExpandedFolders = make(map[string]bool)
			}
			agent.ExpandedFolders[path] = true
			// Also expand all parent folders
			parts := strings.Split(path, "/")
			for i := 1; i < len(parts); i++ {
				agent.ExpandedFolders[strings.Join(parts[:i], "/")] = true
			}
		}

	case "delete_file":
		if args := findPrecedingCall("delete_file"); args != nil {
			delete(agent.Scratchpad, args["filename"])
			if agent.ViewingFile == args["filename"] {
				agent.ViewingFile = ""
			}
		}
	}
}

// autoExpandFolders ensures all parent folders of a file path are expanded
func autoExpandFolders(agent *AgentState, filePath string) {
	if agent.ExpandedFolders == nil {
		agent.ExpandedFolders = make(map[string]bool)
	}
	parts := strings.Split(filePath, "/")
	for i := 1; i < len(parts); i++ {
		agent.ExpandedFolders[strings.Join(parts[:i], "/")] = true
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
		// Re-expand folders for initial files
		agent.ExpandedFolders = make(map[string]bool)
		for path := range agent.InitialScratchpad {
			parts := strings.Split(path, "/")
			for i := 1; i < len(parts); i++ {
				agent.ExpandedFolders[strings.Join(parts[:i], "/")] = true
			}
		}
	}
}
