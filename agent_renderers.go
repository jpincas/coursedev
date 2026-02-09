package main

import (
	"bytes"
	"fmt"
	"sort"
	"strings"

	gt "github.com/jpincas/go-tea"
	a "github.com/jpincas/go-tea/attributes"
	h "github.com/jpincas/go-tea/html"
	"github.com/yuin/goldmark"
)

// renderAgentSidebar renders the complete agent conversation sidebar
func renderAgentSidebar(state *AgentState, isFollowing bool) h.Element {
	if state == nil {
		return h.Span(a.Attrs())
	}

	if !state.SidebarOpen {
		return h.Aside(a.Attrs(
			a.Class("w-14 h-screen sticky top-0 bg-white border-l border-stone-200 flex flex-col items-center py-4 shrink-0"),
		),
			h.Button(a.Attrs(
				a.Class("py-2 px-1 text-xs text-stone-500 cursor-pointer border-none bg-transparent hover:text-accent writing-mode-vertical"),
				a.OnClick(gt.SendBasicMessageNoArgs("AGENT_TOGGLE_SIDEBAR")),
				a.Custom("style", "writing-mode: vertical-rl; text-orientation: mixed"),
			), h.Text(state.Config.Title)),
		)
	}

	children := []h.Element{
		renderAgentTitleBar(state),
	}

	// Main area: either chat messages or full context view
	if state.ShowFullContext {
		children = append(children, renderFullContextView(state))
	} else {
		children = append(children, renderAgentChatMessages(state))
	}

	// Tool calls toggle
	if state.Config.Visibility.ToolCalls == "toggleable" {
		children = append(children, renderVisibilityToggle("AGENT_TOGGLE_TOOLS", "Tool Calls", state.ShowToolCalls))
	}

	// Token count (if visible)
	children = append(children, renderAgentStatusBar(state))

	// Chat input area (or read-only follow-mode version)
	if isFollowing {
		children = append(children, renderFollowModeChatArea(state))
	} else {
		children = append(children, renderChatInputArea(state))
	}

	// System prompt modal (overlays sidebar when open)
	if state.ShowSystem && state.Config.System != "" {
		children = append(children, renderSystemPromptModal(state))
	}

	return h.Aside(a.Attrs(
		a.Class("w-[480px] h-screen sticky top-0 bg-white border-l border-stone-200 flex flex-col overflow-hidden shrink-0 relative"),
	),
		children...,
	)
}

func renderAgentTitleBar(state *AgentState) h.Element {
	rightButtons := []h.Element{}

	// Context view toggle (always available)
	contextIconColor := "text-stone-400 hover:text-blue-500"
	if state.ShowFullContext {
		contextIconColor = "text-blue-500 hover:text-blue-600"
	}
	rightButtons = append(rightButtons,
		h.Button(a.Attrs(
			a.Class(contextIconColor+" cursor-pointer border-none bg-transparent transition-colors duration-150"),
			a.OnClick(gt.SendBasicMessageNoArgs("AGENT_TOGGLE_CONTEXT")),
			a.Custom("title", "Full context"),
		),
			h.UnsafeRaw(`<svg width="18" height="18" viewBox="0 0 18 18" fill="none"><path d="M6.75 4.5L3 9l3.75 4.5M11.25 4.5L15 9l-3.75 4.5" stroke="currentColor" stroke-width="1.5" stroke-linecap="round" stroke-linejoin="round"/></svg>`),
		),
	)

	// System prompt icon (if not hidden)
	if state.Config.Visibility.SystemPrompt != "hidden" && state.Config.System != "" {
		rightButtons = append(rightButtons,
			h.Button(a.Attrs(
				a.Class("text-stone-400 hover:text-purple-500 cursor-pointer border-none bg-transparent transition-colors duration-150"),
				a.OnClick(gt.SendBasicMessageNoArgs("AGENT_TOGGLE_SYSTEM")),
				a.Custom("title", "System prompt"),
			),
				h.UnsafeRaw(`<svg width="18" height="18" viewBox="0 0 18 18" fill="none"><path d="M9 2.25A6.75 6.75 0 1015.75 9 6.75 6.75 0 009 2.25zm0 12A5.25 5.25 0 1114.25 9 5.25 5.25 0 019 14.25z" fill="currentColor"/><path d="M9 5.25a.75.75 0 00-.75.75v3a.75.75 0 001.5 0V6A.75.75 0 009 5.25zM9 11.25a.75.75 0 100 1.5.75.75 0 000-1.5z" fill="currentColor"/></svg>`),
			),
		)
	}

	// Close sidebar button
	rightButtons = append(rightButtons,
		h.Button(a.Attrs(
			a.Class("text-stone-400 hover:text-stone-600 cursor-pointer border-none bg-transparent text-sm"),
			a.OnClick(gt.SendBasicMessageNoArgs("AGENT_TOGGLE_SIDEBAR")),
		), h.Text("\u2715")),
	)

	return h.Div(a.Attrs(a.Class("flex items-center justify-between px-4 py-3 border-b border-stone-200 shrink-0")),
		h.Div(a.Attrs(a.Class("flex items-center gap-2")),
			h.Span(a.Attrs(a.Class("text-sm font-semibold text-stone-800")), h.Text(state.Config.Title)),
		),
		h.Div(a.Attrs(a.Class("flex items-center gap-2")),
			rightButtons...,
		),
	)
}

// renderNarrationFooter renders a narration message in the footer area with a continue button
func renderNarrationFooter(note string) h.Element {
	return h.Div(a.Attrs(a.Class("px-3 py-3 border-t border-stone-200 shrink-0")),
		h.Div(a.Attrs(a.Class("flex items-start gap-2.5 px-3 py-2.5 mb-2 rounded-xl bg-amber-50/80 border border-amber-200/50")),
			h.Span(a.Attrs(a.Class("text-base leading-5 shrink-0 mt-0.5")), h.Text("\U0001F9D1\u200D\U0001F3EB")),
			h.Span(a.Attrs(a.Class("text-sm text-amber-900/70 leading-relaxed")), h.Text(note)),
		),
		h.Div(a.Attrs(
			a.Class("flex items-center justify-center rounded-xl border border-stone-200 bg-stone-50 px-3 py-2 cursor-pointer transition-all duration-150 hover:bg-stone-100 hover:border-stone-300"),
			a.OnClick(gt.SendBasicMessageNoArgs("AGENT_ADVANCE")),
		),
			h.Span(a.Attrs(a.Class("text-xs text-stone-500 font-medium")), h.Text("Continue \u2192")),
		),
	)
}

func renderAgentChatMessages(state *AgentState) h.Element {
	var msgs []h.Element

	for _, m := range state.ChatMessages {
		switch m.Type {
		case "user":
			msgs = append(msgs, renderUserBubble(m))
		case "assistant":
			msgs = append(msgs, renderAssistantBubble(m))
		case "tool_call":
			if shouldShowPanel(state.Config.Visibility.ToolCalls, state.ShowToolCalls) {
				msgs = append(msgs, renderToolCallBubble(m))
			}
		case "tool_result":
			if shouldShowPanel(state.Config.Visibility.ToolCalls, state.ShowToolCalls) {
				msgs = append(msgs, renderToolResultBubble(m))
			}
		case "compaction_divider":
			msgs = append(msgs, renderCompactionDivider())
		case "clear_divider":
			msgs = append(msgs, renderClearDivider(m.Content))
		}
	}

	if len(msgs) == 0 {
		msgs = append(msgs, h.Div(a.Attrs(a.Class("flex items-center justify-center h-full text-stone-400 text-sm")),
			h.Text("Send the first message to begin"),
		))
	}

	return h.Div(a.Attrs(
		a.Id("agent-messages"),
		a.Class("flex-1 overflow-y-auto px-3 py-3 flex flex-col gap-2"),
	),
		msgs...,
	)
}

func renderUserBubble(m ChatMessage) h.Element {
	animStyle := chatAnimStyle(m.GroupIdx)
	return h.Div(a.Attrs(
		a.Class("flex items-end justify-end gap-2"),
		a.Custom("style", animStyle),
	),
		h.Div(a.Attrs(a.Class("max-w-[80%] py-2.5 px-3.5 rounded-2xl rounded-br-sm bg-accent text-white text-sm leading-relaxed")),
			h.Text(m.Content),
		),
		renderChatAvatar("\U0001F464", "bg-accent/15"),
	)
}

func renderAssistantBubble(m ChatMessage) h.Element {
	animStyle := chatAnimStyle(m.GroupIdx)
	rendered := renderChatMarkdown(m.Content)
	return h.Div(a.Attrs(
		a.Class("flex items-end justify-start gap-2"),
		a.Custom("style", animStyle),
	),
		renderChatAvatar("\u2728", "bg-purple-50"),
		h.Div(a.Attrs(a.Class("max-w-[80%] py-2.5 px-3.5 rounded-2xl rounded-bl-sm bg-stone-100 text-stone-700 text-sm leading-relaxed agent-chat-md")),
			h.UnsafeRaw(rendered),
		),
	)
}

// renderChatAvatar renders a small circular avatar with an emoji
func renderChatAvatar(emoji, bgClass string) h.Element {
	return h.Div(a.Attrs(a.Class("w-7 h-7 rounded-full flex items-center justify-center shrink-0 text-sm " + bgClass)),
		h.Span(a.Attrs(a.Class("text-xs")), h.Text(emoji)),
	)
}

func renderToolCallBubble(m ChatMessage) h.Element {
	animStyle := chatAnimStyle(m.GroupIdx)
	toolName := displayToolName(m.ToolName)
	icon := toolCallIcon(m.ToolName)

	// Build args as raw HTML to avoid pretty-printer issues in mono block
	var argsHTML strings.Builder
	keys := make([]string, 0, len(m.ToolArgs))
	for k := range m.ToolArgs {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for _, k := range keys {
		v := m.ToolArgs[k]
		if len(v) > 120 {
			v = v[:120] + "..."
		}
		argsHTML.WriteString(`<div style="font-size:11px;line-height:1.5"><span style="color:#a1a1aa">` + htmlEscape(k) + `: </span><span style="color:#57534e">` + htmlEscape(v) + `</span></div>`)
	}

	return h.Div(a.Attrs(
		a.Class("mx-1"),
		a.Custom("style", animStyle),
	),
		h.Div(a.Attrs(a.Class("rounded-lg overflow-hidden border border-stone-300")),
			// Header bar — bold dark
			h.Div(a.Attrs(a.Class("flex items-center gap-1.5 px-3 py-1.5 bg-stone-800")),
				h.UnsafeRaw(icon),
				h.Span(a.Attrs(a.Class("text-xs font-semibold text-white")), h.Text(toolName)),
			),
			// Args body
			h.Div(a.Attrs(a.Class("px-3 py-2 bg-stone-50 font-mono")),
				h.UnsafeRaw(argsHTML.String()),
			),
		),
	)
}

func renderToolResultBubble(m ChatMessage) h.Element {
	animStyle := chatAnimStyle(m.GroupIdx)
	toolName := displayToolName(m.ToolName)
	icon := toolResultIcon(m.ToolName)

	content := m.Content
	if len(content) > 300 {
		content = content[:300] + "\n..."
	}

	// Use raw HTML for <pre> to avoid go-tea pretty-printer adding whitespace
	preHTML := `<pre style="font-size:11px;line-height:1.5;color:#a1a1aa;font-family:var(--font-mono);white-space:pre-wrap;margin:0;padding:0;background:transparent;border:none">` + htmlEscape(content) + `</pre>`

	return h.Div(a.Attrs(
		a.Class("mx-1"),
		a.Custom("style", animStyle),
	),
		h.Div(a.Attrs(a.Class("rounded-lg overflow-hidden border border-stone-200/80")),
			// Header bar — subtle grey
			h.Div(a.Attrs(a.Class("flex items-center gap-1.5 px-3 py-1 bg-stone-100/80 border-b border-stone-200/60")),
				h.UnsafeRaw(icon),
				h.Span(a.Attrs(a.Class("text-[11px] text-stone-400")), h.Text(toolName)),
			),
			// Result body
			h.Div(a.Attrs(a.Class("px-3 py-2 bg-stone-50/50")),
				h.UnsafeRaw(preHTML),
			),
		),
	)
}

// displayToolName cleans up internal tool names for display
func displayToolName(name string) string {
	switch name {
	case "scratchpad_read":
		return "read_file"
	case "scratchpad_write":
		return "write_file"
	case "list_files":
		return "list_files"
	case "move_file":
		return "move_file"
	case "create_folder":
		return "create_folder"
	case "delete_file":
		return "delete_file"
	case "web_search":
		return "web_search"
	case "fetch_url":
		return "fetch_url"
	default:
		return name
	}
}

// toolCallIcon returns an SVG icon string for the tool call header
func toolCallIcon(toolName string) string {
	switch toolName {
	case "web_search":
		// Magnifying glass
		return `<svg width="12" height="12" viewBox="0 0 16 16" fill="none"><circle cx="7" cy="7" r="4.5" stroke="#fbbf24" stroke-width="1.5"/><path d="M10.5 10.5L14 14" stroke="#fbbf24" stroke-width="1.5" stroke-linecap="round"/></svg>`
	case "fetch_url":
		// Globe
		return `<svg width="12" height="12" viewBox="0 0 16 16" fill="none"><circle cx="8" cy="8" r="6" stroke="#fbbf24" stroke-width="1.2"/><ellipse cx="8" cy="8" rx="3" ry="6" stroke="#fbbf24" stroke-width="1.2"/><path d="M2 8h12M3 4.5h10M3 11.5h10" stroke="#fbbf24" stroke-width="1"/></svg>`
	case "list_files":
		// List/directory icon
		return `<svg width="12" height="12" viewBox="0 0 16 16" fill="none"><path d="M2 4h12M2 8h12M2 12h8" stroke="#fbbf24" stroke-width="1.5" stroke-linecap="round"/></svg>`
	case "move_file":
		// Arrow moving right
		return `<svg width="12" height="12" viewBox="0 0 16 16" fill="none"><path d="M3 8h10M9 4l4 4-4 4" stroke="#fbbf24" stroke-width="1.5" stroke-linecap="round" stroke-linejoin="round"/></svg>`
	case "create_folder":
		// Folder with plus
		return `<svg width="12" height="12" viewBox="0 0 16 16" fill="none"><path d="M1.5 3A1.5 1.5 0 013 1.5h3.293a1 1 0 01.707.293L8.414 3.207a1 1 0 00.707.293H13A1.5 1.5 0 0114.5 5v8a1.5 1.5 0 01-1.5 1.5H3A1.5 1.5 0 011.5 13V3z" fill="#fbbf24" fill-opacity="0.3" stroke="#fbbf24" stroke-width="0.8"/><path d="M8 7v4M6 9h4" stroke="#fbbf24" stroke-width="1.2" stroke-linecap="round"/></svg>`
	case "delete_file":
		// Trash/X
		return `<svg width="12" height="12" viewBox="0 0 16 16" fill="none"><path d="M5.5 2h5M2 4h12M4 4l1 10h6l1-10M6.5 7v4M9.5 7v4" stroke="#fbbf24" stroke-width="1.2" stroke-linecap="round" stroke-linejoin="round"/></svg>`
	default:
		// Wrench (existing)
		return `<svg width="12" height="12" viewBox="0 0 16 16" fill="none"><path d="M14.25 6.14L13.07 5l.41-1.66a.38.38 0 00-.11-.36.37.37 0 00-.36-.1L11.35 3.3 10.2 2.11a.37.37 0 00-.53 0L8.54 3.24 7.47 2.87a.38.38 0 00-.42.09L1.17 8.84a.38.38 0 000 .53l2.12 2.12-1.72 1.72a.75.75 0 001.06 1.06l1.72-1.72 2.12 2.12a.38.38 0 00.53 0l5.88-5.88a.38.38 0 00.09-.42l-.37-1.07 1.13-1.13a.37.37 0 000-.53z" fill="#fbbf24"/></svg>`
	}
}

// toolResultIcon returns an SVG icon string for the tool result header
func toolResultIcon(toolName string) string {
	switch toolName {
	case "web_search":
		// Magnifying glass (muted)
		return `<svg width="11" height="11" viewBox="0 0 16 16" fill="none"><circle cx="7" cy="7" r="4.5" stroke="#d6d3d1" stroke-width="1.5"/><path d="M10.5 10.5L14 14" stroke="#d6d3d1" stroke-width="1.5" stroke-linecap="round"/></svg>`
	case "fetch_url":
		// Globe (muted)
		return `<svg width="11" height="11" viewBox="0 0 16 16" fill="none"><circle cx="8" cy="8" r="6" stroke="#d6d3d1" stroke-width="1.2"/><ellipse cx="8" cy="8" rx="3" ry="6" stroke="#d6d3d1" stroke-width="1.2"/><path d="M2 8h12M3 4.5h10M3 11.5h10" stroke="#d6d3d1" stroke-width="1"/></svg>`
	case "list_files":
		// List (muted)
		return `<svg width="11" height="11" viewBox="0 0 16 16" fill="none"><path d="M2 4h12M2 8h12M2 12h8" stroke="#d6d3d1" stroke-width="1.5" stroke-linecap="round"/></svg>`
	case "move_file":
		// Arrow (muted)
		return `<svg width="11" height="11" viewBox="0 0 16 16" fill="none"><path d="M3 8h10M9 4l4 4-4 4" stroke="#d6d3d1" stroke-width="1.5" stroke-linecap="round" stroke-linejoin="round"/></svg>`
	case "create_folder":
		// Folder (muted)
		return `<svg width="11" height="11" viewBox="0 0 16 16" fill="none"><path d="M1.5 3A1.5 1.5 0 013 1.5h3.293a1 1 0 01.707.293L8.414 3.207a1 1 0 00.707.293H13A1.5 1.5 0 0114.5 5v8a1.5 1.5 0 01-1.5 1.5H3A1.5 1.5 0 011.5 13V3z" fill="#d6d3d1" fill-opacity="0.2" stroke="#d6d3d1" stroke-width="0.8"/></svg>`
	case "delete_file":
		// Trash (muted)
		return `<svg width="11" height="11" viewBox="0 0 16 16" fill="none"><path d="M5.5 2h5M2 4h12M4 4l1 10h6l1-10M6.5 7v4M9.5 7v4" stroke="#d6d3d1" stroke-width="1.2" stroke-linecap="round" stroke-linejoin="round"/></svg>`
	default:
		// File (existing)
		return `<svg width="11" height="11" viewBox="0 0 16 16" fill="none"><path d="M14 4.5V14a1 1 0 01-1 1H3a1 1 0 01-1-1V2a1 1 0 011-1h6.5L14 4.5z" fill="#d6d3d1" fill-opacity="0.3" stroke="#d6d3d1" stroke-width="1"/><path d="M9.5 1v4H14" stroke="#d6d3d1" stroke-width="1" fill="none"/></svg>`
	}
}

func renderCompactionDivider() h.Element {
	return h.Div(a.Attrs(a.Class("flex items-center gap-2 py-2 mx-1")),
		h.Div(a.Attrs(a.Class("flex-1 h-px bg-stone-200"))),
		h.Span(a.Attrs(a.Class("text-xs text-stone-400 font-medium whitespace-nowrap")), h.Text("Context compacted")),
		h.Div(a.Attrs(a.Class("flex-1 h-px bg-stone-200"))),
	)
}

func renderClearDivider(note string) h.Element {
	text := "Conversation cleared"
	if note != "" {
		text = note
	}
	return h.Div(a.Attrs(a.Class("flex items-center gap-2 py-2 mx-1")),
		h.Div(a.Attrs(a.Class("flex-1 h-px bg-stone-200"))),
		h.Span(a.Attrs(a.Class("text-xs text-stone-400 font-medium whitespace-nowrap")), h.Text(text)),
		h.Div(a.Attrs(a.Class("flex-1 h-px bg-stone-200"))),
	)
}

// renderFullContextView renders the full context as the main scrollable area (replaces chat messages when toggled)
func renderFullContextView(state *AgentState) h.Element {
	var entries []h.Element

	for _, cm := range state.ContextMessages {
		var roleColor string
		switch cm.Role {
		case "system":
			roleColor = "text-purple-400"
		case "user":
			roleColor = "text-accent"
		case "assistant":
			roleColor = "text-stone-600"
		case "tool_call":
			roleColor = "text-amber-600"
		case "tool_result":
			roleColor = "text-stone-400"
		case "compaction_summary":
			roleColor = "text-blue-600"
		default:
			roleColor = "text-stone-400"
		}

		content := cm.Content
		if len(content) > 500 {
			content = content[:500] + "\n..."
		}

		// Use raw HTML for pre to avoid go-tea pretty-printer adding phantom indentation
		preHTML := `<pre style="font-size:11px;line-height:1.5;color:#78716c;font-family:var(--font-mono);white-space:pre-wrap;margin:0;padding:0;background:transparent;border:none">` + htmlEscape(content) + `</pre>`

		entries = append(entries,
			h.Div(a.Attrs(a.Class("py-2 border-b border-stone-200/50 last:border-b-0")),
				h.Div(a.Attrs(a.Class("text-xs font-medium mb-1 "+roleColor)),
					h.Text(cm.Role)),
				h.UnsafeRaw(preHTML),
			),
		)
	}

	if len(entries) == 0 {
		entries = append(entries, h.Div(a.Attrs(a.Class("flex items-center justify-center h-full text-stone-400 text-sm")),
			h.Text("No context yet"),
		))
	}

	return h.Div(a.Attrs(
		a.Id("agent-context"),
		a.Class("flex-1 overflow-y-auto px-4 py-3 bg-stone-50/50"),
	),
		entries...,
	)
}

func renderSystemPromptModal(state *AgentState) h.Element {
	return h.Div(a.Attrs(
		a.Class("absolute inset-0 z-50 flex flex-col"),
	),
		// Backdrop
		h.Div(a.Attrs(
			a.Class("absolute inset-0 bg-black/30"),
			a.OnClick(gt.SendBasicMessageNoArgs("AGENT_TOGGLE_SYSTEM")),
		)),
		// Modal card
		h.Div(a.Attrs(
			a.Class("relative mx-5 mt-14 mb-5 flex flex-col rounded-xl bg-white border border-stone-200 overflow-hidden max-h-[70%]"),
			a.Custom("style", "box-shadow: 0 8px 30px rgba(0,0,0,0.15)"),
		),
			// Header
			h.Div(a.Attrs(a.Class("flex items-center justify-between px-4 py-3 border-b border-stone-200 shrink-0")),
				h.Div(a.Attrs(a.Class("flex items-center gap-2")),
					h.UnsafeRaw(`<svg width="14" height="14" viewBox="0 0 18 18" fill="none"><path d="M9 2.25A6.75 6.75 0 1015.75 9 6.75 6.75 0 009 2.25zm0 12A5.25 5.25 0 1114.25 9 5.25 5.25 0 019 14.25z" fill="#a855f7"/><path d="M9 5.25a.75.75 0 00-.75.75v3a.75.75 0 001.5 0V6A.75.75 0 009 5.25zM9 11.25a.75.75 0 100 1.5.75.75 0 000-1.5z" fill="#a855f7"/></svg>`),
					h.Span(a.Attrs(a.Class("text-xs font-semibold text-purple-500 uppercase tracking-wider")), h.Text("System Prompt")),
				),
				h.Button(a.Attrs(
					a.Class("text-stone-400 hover:text-stone-600 cursor-pointer border-none bg-transparent text-sm"),
					a.OnClick(gt.SendBasicMessageNoArgs("AGENT_TOGGLE_SYSTEM")),
				), h.Text("\u2715")),
			),
			// Content
			h.Div(a.Attrs(a.Class("flex-1 overflow-y-auto px-4 py-3")),
				h.Pre(a.Attrs(a.Class("text-xs text-stone-600 font-mono whitespace-pre-wrap m-0 bg-transparent border-none p-0 leading-relaxed")),
					h.Text(state.Config.System)),
			),
		),
	)
}

func renderScratchpadPanel(state *AgentState) h.Element {
	fileCount := len(state.Scratchpad)

	if !state.ScratchpadOpen {
		return h.Div(a.Attrs(a.Class("mx-3 my-2 shrink-0")),
			h.Button(a.Attrs(
				a.Class("w-full py-2 px-3 rounded-lg border border-stone-200 bg-stone-50 text-left text-xs text-stone-500 font-medium cursor-pointer hover:border-stone-200 hover:text-stone-600 transition-all duration-150"),
				a.OnClick(gt.SendBasicMessageNoArgs("AGENT_TOGGLE_SCRATCHPAD")),
			), h.Text(fmt.Sprintf("Scratchpad (%d files) \u25B8", fileCount))),
		)
	}

	// Sort filenames for deterministic rendering
	filenames := make([]string, 0, fileCount)
	for f := range state.Scratchpad {
		filenames = append(filenames, f)
	}
	sort.Strings(filenames)

	var files []h.Element
	for _, filename := range filenames {
		content := state.Scratchpad[filename]
		isViewing := state.ViewingFile == filename

		viewLabel := "view"
		if isViewing {
			viewLabel = "hide"
		}

		// Clicking the same file toggles it off; clicking a different file opens it
		clickArg := filename
		if isViewing {
			clickArg = ""
		}

		fileEl := h.Div(a.Attrs(a.Class("py-1")),
			h.Div(a.Attrs(a.Class("flex items-center justify-between")),
				h.Span(a.Attrs(a.Class("text-xs font-mono text-stone-600")), h.Text(filename)),
				h.Button(a.Attrs(
					a.Class("text-xs text-accent cursor-pointer border-none bg-transparent hover:underline"),
					a.OnClick(gt.SendBasicMessage("AGENT_VIEW_FILE", clickArg)),
				), h.Text(viewLabel)),
			),
		)

		if isViewing {
			fileEl = h.Div(a.Attrs(a.Class("py-1")),
				h.Div(a.Attrs(a.Class("flex items-center justify-between")),
					h.Span(a.Attrs(a.Class("text-xs font-mono text-stone-600")), h.Text(filename)),
					h.Button(a.Attrs(
						a.Class("text-xs text-accent cursor-pointer border-none bg-transparent hover:underline"),
						a.OnClick(gt.SendBasicMessage("AGENT_VIEW_FILE", "")),
					), h.Text("hide")),
				),
				h.Pre(a.Attrs(a.Class("mt-1 text-xs text-stone-500 font-mono whitespace-pre-wrap m-0 bg-stone-50 border border-stone-200 rounded p-2 max-h-48 overflow-y-auto")),
					h.Text(content)),
			)
		}

		files = append(files, fileEl)
	}

	return h.Div(a.Attrs(a.Class("mx-3 my-2 rounded-lg border border-stone-200 bg-stone-50 shrink-0")),
		h.Div(a.Attrs(a.Class("flex items-center justify-between px-3 py-2 border-b border-stone-200")),
			h.Span(a.Attrs(a.Class("text-xs font-semibold text-stone-500 uppercase tracking-wider")), h.Text("Scratchpad")),
			h.Button(a.Attrs(
				a.Class("text-xs text-stone-400 cursor-pointer border-none bg-transparent hover:text-stone-600"),
				a.OnClick(gt.SendBasicMessageNoArgs("AGENT_TOGGLE_SCRATCHPAD")),
			), h.Text("\u25BE")),
		),
		h.Div(a.Attrs(a.Class("px-3 py-2")),
			files...,
		),
	)
}

func renderChatInputArea(state *AgentState) h.Element {
	atEnd := state.ScriptIndex >= len(state.Config.Script)

	// Demo complete
	if atEnd {
		var children []h.Element
		if state.CurrentNote != "" {
			children = append(children,
				h.Div(a.Attrs(a.Class("flex items-start gap-2.5 px-3 py-2.5 mb-2 rounded-xl bg-amber-50/80 border border-amber-200/50")),
					h.Span(a.Attrs(a.Class("text-base leading-5 shrink-0 mt-0.5")), h.Text("\U0001F9D1\u200D\U0001F3EB")),
					h.Span(a.Attrs(a.Class("text-sm text-amber-900/70 leading-relaxed")), h.Text(state.CurrentNote)),
				),
			)
		}
		children = append(children,
			h.Div(a.Attrs(a.Class("flex items-center justify-between px-4 py-3 rounded-2xl bg-stone-50 border border-stone-200")),
				h.Span(a.Attrs(a.Class("text-sm text-accent font-medium")), h.Text("\u2713 Demo complete")),
				h.Button(a.Attrs(
					a.Class("py-1.5 px-3 text-sm rounded-lg bg-stone-200 text-stone-600 font-medium cursor-pointer transition-all duration-150 hover:bg-stone-300 border-none"),
					a.OnClick(gt.SendBasicMessageNoArgs("AGENT_RESET")),
				), h.Text("Done")),
			),
		)
		return h.Div(a.Attrs(a.Class("px-3 py-3 border-t border-stone-200 shrink-0")),
			children...,
		)
	}

	next := state.Config.Script[state.ScriptIndex]

	// For user events: show the message in a fake chat input with send button
	if next.Type == "user" {
		var children []h.Element
		if state.CurrentNote != "" {
			children = append(children,
				h.Div(a.Attrs(a.Class("flex items-start gap-2.5 px-3 py-2.5 mb-2 rounded-xl bg-amber-50/80 border border-amber-200/50")),
					h.Span(a.Attrs(a.Class("text-base leading-5 shrink-0 mt-0.5")), h.Text("\U0001F9D1\u200D\U0001F3EB")),
					h.Span(a.Attrs(a.Class("text-sm text-amber-900/70 leading-relaxed")), h.Text(state.CurrentNote)),
				),
			)
		}
		children = append(children,
			h.Div(a.Attrs(
				a.Class("flex items-end gap-2 rounded-2xl border border-stone-300 bg-white px-4 py-3 cursor-pointer transition-all duration-150 hover:border-accent/50"),
				a.OnClick(gt.SendBasicMessageNoArgs("AGENT_ADVANCE")),
			),
				h.Div(a.Attrs(a.Class("flex-1 text-sm text-stone-800 leading-relaxed min-h-[20px]")),
					h.Text(next.Content),
				),
				renderSendButton(),
			),
		)
		return h.Div(a.Attrs(a.Class("px-3 py-3 border-t border-stone-200 shrink-0")),
			children...,
		)
	}

	// For non-user events: narration with continue, or just a continue pill
	if state.CurrentNote != "" {
		return renderNarrationFooter(state.CurrentNote)
	}

	label := "Continue"
	switch next.Type {
	case "compaction":
		label = "Compact context"
	case "clear":
		label = "Clear & continue"
	}

	return h.Div(a.Attrs(a.Class("px-3 py-3 border-t border-stone-200 shrink-0")),
		h.Div(a.Attrs(
			a.Class("flex items-center justify-center rounded-xl border border-stone-200 bg-stone-50 px-3 py-2 cursor-pointer transition-all duration-150 hover:bg-stone-100 hover:border-stone-300"),
			a.OnClick(gt.SendBasicMessageNoArgs("AGENT_ADVANCE")),
		),
			h.Span(a.Attrs(a.Class("text-xs text-stone-500 font-medium")), h.Text(label)),
		),
	)
}

// renderSendButton renders a circular send button like modern AI chat UIs
func renderSendButton() h.Element {
	return h.Div(a.Attrs(
		a.Class("w-8 h-8 rounded-full bg-accent flex items-center justify-center shrink-0"),
	),
		h.UnsafeRaw(`<svg width="16" height="16" viewBox="0 0 16 16" fill="none"><path d="M8 12V4M8 4L4 8M8 4L12 8" stroke="white" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"/></svg>`),
	)
}

// renderFollowModeChatArea renders a read-only version of the chat input area for followers.
// Shows narration, the upcoming user message preview, and demo complete state — but no click handlers.
func renderFollowModeChatArea(state *AgentState) h.Element {
	var children []h.Element

	atEnd := state.ScriptIndex >= len(state.Config.Script)

	// Narration
	if state.CurrentNote != "" {
		children = append(children,
			h.Div(a.Attrs(a.Class("flex items-start gap-2.5 px-3 py-2.5 mb-2 rounded-xl bg-amber-50/80 border border-amber-200/50")),
				h.Span(a.Attrs(a.Class("text-base leading-5 shrink-0 mt-0.5")), h.Text("\U0001F9D1\u200D\U0001F3EB")),
				h.Span(a.Attrs(a.Class("text-sm text-amber-900/70 leading-relaxed")), h.Text(state.CurrentNote)),
			),
		)
	}

	if atEnd {
		children = append(children,
			h.Div(a.Attrs(a.Class("flex items-center justify-center px-4 py-3 rounded-2xl bg-stone-50 border border-stone-200")),
				h.Span(a.Attrs(a.Class("text-sm text-accent font-medium")), h.Text("\u2713 Demo complete")),
			),
		)
	} else {
		next := state.Config.Script[state.ScriptIndex]
		if next.Type == "user" {
			// Show the upcoming user message as a read-only preview
			children = append(children,
				h.Div(a.Attrs(a.Class("flex items-end gap-2 rounded-2xl border border-stone-200 bg-stone-50/50 px-4 py-3")),
					h.Div(a.Attrs(a.Class("flex-1 text-sm text-stone-400 leading-relaxed min-h-[20px]")),
						h.Text(next.Content),
					),
				),
			)
		}
	}

	// Follow mode banner
	children = append(children,
		h.Div(a.Attrs(a.Class("flex items-center justify-center py-2")),
			h.Span(a.Attrs(a.Class("text-xs text-blue-500 font-medium")), h.Text("Your instructor is leading this demo")),
		),
	)

	return h.Div(a.Attrs(a.Class("px-3 py-3 border-t border-stone-200 shrink-0")),
		children...,
	)
}

func renderAgentStatusBar(state *AgentState) h.Element {
	if state.Config.Visibility.TokenCount != "visible" {
		return h.Span(a.Attrs())
	}

	return h.Div(a.Attrs(a.Class("px-3 pb-1 flex justify-end text-xs shrink-0")),
		h.Span(a.Attrs(a.Class("text-stone-400")),
			h.Text(fmt.Sprintf("~%d tokens", state.TokenCount))),
	)
}

func renderVisibilityToggle(messageName, label string, isOpen bool) h.Element {
	icon := "\u25B8"
	textColor := "text-stone-400"
	if isOpen {
		icon = "\u25BE"
		textColor = "text-stone-500"
	}
	return h.Div(a.Attrs(a.Class("mx-3 shrink-0")),
		h.Button(a.Attrs(
			a.Class(fmt.Sprintf("w-full py-1 px-2 text-xs %s font-medium cursor-pointer border-none bg-transparent hover:text-stone-600 text-left transition-colors duration-150", textColor)),
			a.OnClick(gt.SendBasicMessageNoArgs(messageName)),
		), h.Text(icon+" "+label)),
	)
}

// shouldShowPanel returns true if a panel should be visible based on config and toggle state
func shouldShowPanel(configValue string, stateToggle bool) bool {
	switch configValue {
	case "visible":
		return true
	case "hidden":
		return false
	case "toggleable":
		return stateToggle
	default:
		return false
	}
}

// containsTool checks if a tool name is in the tools list
func containsTool(tools []string, name string) bool {
	for _, t := range tools {
		if strings.Contains(t, name) {
			return true
		}
	}
	return false
}

// renderChatMarkdown converts markdown text to HTML for chat bubbles
func renderChatMarkdown(text string) string {
	var buf bytes.Buffer
	if err := goldmark.Convert([]byte(text), &buf); err != nil {
		return htmlEscape(text)
	}
	return buf.String()
}

// chatAnimStyle returns an inline style for staggered animation
func chatAnimStyle(groupIdx int) string {
	if groupIdx < 0 {
		return "" // Already displayed, no animation
	}
	delay := groupIdx * 400
	return fmt.Sprintf("opacity: 0; animation: agentFadeSlideIn 300ms ease %dms forwards", delay)
}

// renderAgentLaunchCard renders the inline card shown on the page before workspace is opened
func renderAgentLaunchCard(block *AgentBlock) h.Element {
	fileCount := len(block.Scratchpad)
	subtitle := fmt.Sprintf("%d files · interactive walkthrough", fileCount)

	return h.Div(a.Attrs(a.Class("not-prose my-8 rounded-xl border border-zinc-200 bg-gradient-to-br from-zinc-50 to-stone-100 p-8 text-center")),
		h.Div(a.Attrs(a.Class("text-3xl mb-3")), h.Text("\u25B6\uFE0F")),
		h.H3(a.Attrs(a.Class("text-lg font-semibold text-zinc-900 mb-1")), h.Text(block.Title)),
		h.P(a.Attrs(a.Class("text-sm text-zinc-500 mb-5")), h.Text(subtitle)),
		h.Button(a.Attrs(
			a.Class("px-6 py-2.5 rounded-lg bg-accent text-zinc-900 font-semibold text-sm cursor-pointer border-none hover:brightness-110 transition-all duration-150"),
			a.OnClick(gt.SendBasicMessageNoArgs("AGENT_OPEN_WORKSPACE")),
		), h.Text("Launch Demo")),
	)
}

// ============================================================================
// Agent Workspace — Full-screen file explorer + chat layout
// ============================================================================

// renderAgentWorkspace renders the full-screen agent demo workspace
func renderAgentWorkspace(m *Model, state *AgentState, isFollowing bool) h.Element {
	return h.Div(a.Attrs(a.Class("flex h-screen w-screen overflow-hidden")),
		renderAgentFileWorkspace(m, state),
		renderAgentSidebar(state, isFollowing),
	)
}

// renderAgentFileWorkspace renders the file explorer + preview panes
func renderAgentFileWorkspace(m *Model, state *AgentState) h.Element {
	return h.Div(a.Attrs(a.Class("flex-1 flex flex-col h-screen overflow-hidden bg-zinc-900")),
		renderAgentWorkspaceHeader(m),
		renderAgentFileExplorer(state),
		renderAgentFilePreview(state),
	)
}

// renderAgentWorkspaceHeader renders a thin header bar for the workspace
func renderAgentWorkspaceHeader(m *Model) h.Element {
	pageTitle := ""
	if page := m.currentPage(); page != nil {
		pageTitle = page.Meta.Title
	}

	return h.Div(a.Attrs(a.Class("flex items-center justify-between px-4 py-2 bg-zinc-900 border-b border-zinc-800 shrink-0")),
		h.Div(a.Attrs(a.Class("flex items-center gap-3")),
			h.Span(a.Attrs(a.Class("text-[11px] text-zinc-500 uppercase tracking-widest font-semibold")), h.Text("Demo")),
			h.Span(a.Attrs(a.Class("text-zinc-700")), h.Text("/")),
			h.Span(a.Attrs(a.Class("text-sm text-zinc-400")), h.Text(pageTitle)),
		),
		h.Button(a.Attrs(
			a.Class("py-1 px-3 text-xs rounded bg-zinc-800 text-zinc-400 font-medium cursor-pointer border border-zinc-700 hover:bg-zinc-700 hover:text-zinc-300 transition-colors duration-150"),
			a.OnClick(gt.SendBasicMessageNoArgs("AGENT_CLOSE_WORKSPACE")),
		), h.Text("← Back to course")),
	)
}

// fileTreeNode represents a node in the file explorer tree
type fileTreeNode struct {
	Name     string          // Just the filename or folder name (no path)
	FullPath string          // Full path for files, folder prefix for directories
	IsDir    bool
	Children []*fileTreeNode
}

// buildFileTree constructs a tree from the flat scratchpad map
func buildFileTree(scratchpad map[string]string) []*fileTreeNode {
	root := &fileTreeNode{IsDir: true}
	dirNodes := map[string]*fileTreeNode{"": root}

	// Ensure parent directories exist
	getOrCreateDir := func(path string) *fileTreeNode {
		if node, ok := dirNodes[path]; ok {
			return node
		}
		parts := strings.Split(path, "/")
		current := root
		for i, part := range parts {
			dirPath := strings.Join(parts[:i+1], "/")
			if existing, ok := dirNodes[dirPath]; ok {
				current = existing
				continue
			}
			newDir := &fileTreeNode{Name: part, FullPath: dirPath, IsDir: true}
			current.Children = append(current.Children, newDir)
			dirNodes[dirPath] = newDir
			current = newDir
		}
		return current
	}

	// Add all files
	filenames := make([]string, 0, len(scratchpad))
	for f := range scratchpad {
		filenames = append(filenames, f)
	}
	sort.Strings(filenames)

	for _, path := range filenames {
		lastSlash := strings.LastIndex(path, "/")
		if lastSlash < 0 {
			// Root-level file
			root.Children = append(root.Children, &fileTreeNode{
				Name:     path,
				FullPath: path,
			})
		} else {
			dirPath := path[:lastSlash]
			parent := getOrCreateDir(dirPath)
			fileName := path[lastSlash+1:]
			parent.Children = append(parent.Children, &fileTreeNode{
				Name:     fileName,
				FullPath: path,
			})
		}
	}

	// Sort children: folders first, then files, alphabetically within each group
	var sortChildren func(node *fileTreeNode)
	sortChildren = func(node *fileTreeNode) {
		sort.Slice(node.Children, func(i, j int) bool {
			if node.Children[i].IsDir != node.Children[j].IsDir {
				return node.Children[i].IsDir // dirs first
			}
			return node.Children[i].Name < node.Children[j].Name
		})
		for _, child := range node.Children {
			if child.IsDir {
				sortChildren(child)
			}
		}
	}
	sortChildren(root)

	return root.Children
}

// renderAgentFileExplorer renders the file list panel with tree view
func renderAgentFileExplorer(state *AgentState) h.Element {
	tree := buildFileTree(state.Scratchpad)

	var items []h.Element
	var renderNodes func(nodes []*fileTreeNode, depth int)
	renderNodes = func(nodes []*fileTreeNode, depth int) {
		for _, node := range nodes {
			if node.IsDir {
				items = append(items, renderFolderRow(state, node, depth))
				if state.ExpandedFolders[node.FullPath] {
					renderNodes(node.Children, depth+1)
				}
			} else {
				items = append(items, renderFileRow(state, node, depth))
			}
		}
	}
	renderNodes(tree, 0)

	if len(items) == 0 {
		items = append(items,
			h.Div(a.Attrs(a.Class("px-4 py-3 text-zinc-600 text-sm italic")),
				h.Text("No files yet"),
			),
		)
	}

	return h.Div(a.Attrs(a.Class("shrink-0 border-b border-zinc-800")),
		h.Div(a.Attrs(a.Class("px-4 py-2 text-[11px] text-zinc-500 font-semibold uppercase tracking-widest")),
			h.Text(fmt.Sprintf("Explorer \u00B7 %d files", len(state.Scratchpad))),
		),
		h.Div(a.Attrs(a.Class("pb-2")),
			items...,
		),
	)
}

func renderFolderRow(state *AgentState, node *fileTreeNode, depth int) h.Element {
	isExpanded := state.ExpandedFolders[node.FullPath]
	chevron := "\u25B8" // right-pointing triangle
	if isExpanded {
		chevron = "\u25BE" // down-pointing triangle
	}
	indent := 16 * depth

	return h.Div(a.Attrs(
		a.Class("flex items-center gap-1.5 py-1.5 cursor-pointer transition-colors duration-100 hover:bg-zinc-800/50"),
		a.Custom("style", fmt.Sprintf("padding-left:%dpx;padding-right:16px", 16+indent)),
		a.OnClick(gt.SendBasicMessage("AGENT_TOGGLE_FOLDER", node.FullPath)),
	),
		h.Span(a.Attrs(a.Class("text-[10px] text-zinc-500 w-3 text-center shrink-0")), h.Text(chevron)),
		renderFolderIcon(isExpanded),
		h.Span(a.Attrs(a.Class("text-sm text-zinc-300 font-mono")), h.Text(node.Name)),
	)
}

func renderFileRow(state *AgentState, node *fileTreeNode, depth int) h.Element {
	isSelected := state.ViewingFile == node.FullPath
	ext := getFileExt(node.Name)
	indent := 16 * depth

	rowClass := "flex items-center gap-2.5 py-1.5 cursor-pointer transition-colors duration-100 hover:bg-zinc-800/50"
	textClass := "text-sm text-zinc-400 font-mono"
	borderStyle := ""
	if isSelected {
		rowClass = "flex items-center gap-2.5 py-1.5 cursor-pointer bg-zinc-800 border-l-2 border-accent"
		textClass = "text-sm text-zinc-200 font-mono"
		borderStyle = fmt.Sprintf("padding-left:%dpx;padding-right:16px", 14+indent+16) // 14 = 16 - 2px border
	} else {
		borderStyle = fmt.Sprintf("padding-left:%dpx;padding-right:16px", 16+indent+16) // extra 16 for chevron space
	}

	return h.Div(a.Attrs(
		a.Class(rowClass),
		a.Custom("style", borderStyle),
		a.OnClick(gt.SendBasicMessage("AGENT_VIEW_FILE", node.FullPath)),
	),
		renderFileIcon(ext),
		h.Span(a.Attrs(a.Class(textClass)), h.Text(node.Name)),
	)
}

// renderFolderIcon returns an SVG folder icon
func renderFolderIcon(isOpen bool) h.Element {
	if isOpen {
		return h.UnsafeRaw(
			`<svg width="16" height="16" viewBox="0 0 16 16" fill="none" class="shrink-0">` +
				`<path d="M1.5 3A1.5 1.5 0 013 1.5h3.293a1 1 0 01.707.293L8.414 3.207a1 1 0 00.707.293H13A1.5 1.5 0 0114.5 5v1H1.5V3z" fill="#fbbf24" fill-opacity="0.3" stroke="#fbbf24" stroke-width="0.8"/>` +
				`<path d="M1 6.5h14l-1.5 8H2.5L1 6.5z" fill="#fbbf24" fill-opacity="0.25" stroke="#fbbf24" stroke-width="0.8"/>` +
				`</svg>`)
	}
	return h.UnsafeRaw(
		`<svg width="16" height="16" viewBox="0 0 16 16" fill="none" class="shrink-0">` +
			`<path d="M1.5 3A1.5 1.5 0 013 1.5h3.293a1 1 0 01.707.293L8.414 3.207a1 1 0 00.707.293H13A1.5 1.5 0 0114.5 5v8a1.5 1.5 0 01-1.5 1.5H3A1.5 1.5 0 011.5 13V3z" fill="#fbbf24" fill-opacity="0.25" stroke="#fbbf24" stroke-width="0.8"/>` +
			`</svg>`)
}

// renderAgentFilePreview renders the file content viewer
func renderAgentFilePreview(state *AgentState) h.Element {
	if state.ViewingFile == "" {
		return h.Div(a.Attrs(a.Class("flex-1 flex items-center justify-center bg-zinc-800/40")),
			h.Div(a.Attrs(a.Class("text-center")),
				renderFileIconLarge(),
				h.P(a.Attrs(a.Class("text-zinc-600 text-sm mt-2")), h.Text("Select a file to preview")),
			),
		)
	}

	content, ok := state.Scratchpad[state.ViewingFile]
	if !ok {
		return h.Div(a.Attrs(a.Class("flex-1 flex items-center justify-center bg-zinc-800/40")),
			h.Span(a.Attrs(a.Class("text-zinc-600 text-sm")), h.Text("File not found")),
		)
	}

	ext := getFileExt(state.ViewingFile)

	// CSV files rendered as a table inside the editor chrome
	if ext == "csv" {
		return renderEditorWindow(state.ViewingFile, ext, renderCSVTableHTML(content))
	}

	// Text/md files
	return renderEditorWindow(state.ViewingFile, ext, renderFileContentHTML(content, ext))
}

// renderEditorWindow wraps file content in a macOS-style editor window
func renderEditorWindow(filename, ext string, contentHTML string) h.Element {
	return h.Div(a.Attrs(a.Class("flex-1 flex flex-col bg-zinc-800/40 overflow-hidden min-h-0 p-5")),
		// Editor window with lift
		h.Div(a.Attrs(
			a.Class("flex-1 flex flex-col rounded-lg overflow-hidden min-h-0"),
			a.Custom("style", "box-shadow: 0 8px 30px rgba(0,0,0,0.4), 0 2px 8px rgba(0,0,0,0.3)"),
		),
			// Title bar
			renderEditorTitleBar(filename, ext),
			// Content area
			h.UnsafeRaw(contentHTML),
		),
	)
}

// renderEditorTitleBar renders a macOS-style window title bar
func renderEditorTitleBar(filename, ext string) h.Element {
	// Show just the basename in the title, full path if it differs
	displayName := filename
	if idx := strings.LastIndex(filename, "/"); idx >= 0 {
		displayName = filename[idx+1:]
	}

	titleElems := []h.Element{
		renderFileIcon(ext),
		h.Span(a.Attrs(a.Class("text-sm text-zinc-600 font-mono")), h.Text(displayName)),
	}
	// Show full path as subdued text if file is in a folder
	if displayName != filename {
		titleElems = append(titleElems,
			h.Span(a.Attrs(a.Class("text-xs text-zinc-400 font-mono ml-1")), h.Text(filename)),
		)
	}

	return h.Div(a.Attrs(a.Class("flex items-center px-4 py-2 bg-zinc-200 border-b border-zinc-300 shrink-0")),
		// Traffic lights
		h.Div(a.Attrs(a.Class("flex items-center gap-2 mr-4")),
			h.UnsafeRaw(`<span style="width:12px;height:12px;border-radius:50%;background:#ff5f57;display:inline-block"></span>`),
			h.UnsafeRaw(`<span style="width:12px;height:12px;border-radius:50%;background:#febc2e;display:inline-block"></span>`),
			h.UnsafeRaw(`<span style="width:12px;height:12px;border-radius:50%;background:#28c840;display:inline-block"></span>`),
		),
		// Filename
		h.Div(a.Attrs(a.Class("flex items-center gap-2")),
			titleElems...,
		),
	)
}

// renderFileContentHTML builds a raw HTML string for the file preview.
// We bypass go-tea's element tree because its pretty-printer adds tabs and
// newlines between children, which <pre> preserves and displays.
func renderFileContentHTML(content, ext string) string {
	content = strings.TrimRight(content, "\n")
	escaped := htmlEscape(content)

	var inner string
	if ext == "md" {
		inner = highlightMarkdown(content)
	} else {
		inner = escaped
	}

	return `<pre class="flex-1 overflow-auto m-0 px-4 py-3 font-mono text-sm whitespace-pre-wrap" style="background:white;border:none;color:#1c1917">` + inner + `</pre>`
}

func highlightMarkdown(content string) string {
	lines := strings.Split(content, "\n")
	var b strings.Builder
	for i, line := range lines {
		if i > 0 {
			b.WriteByte('\n')
		}
		trimmed := strings.TrimSpace(line)
		esc := htmlEscape(line)

		switch {
		case strings.HasPrefix(trimmed, "### "):
			b.WriteString(`<span style="color:#7c3aed">### </span><span style="font-weight:500">` + htmlEscape(trimmed[4:]) + `</span>`)
		case strings.HasPrefix(trimmed, "## "):
			b.WriteString(`<span style="color:#7c3aed">## </span><span style="font-weight:600">` + htmlEscape(trimmed[3:]) + `</span>`)
		case strings.HasPrefix(trimmed, "# "):
			b.WriteString(`<span style="color:#7c3aed"># </span><span style="font-weight:700">` + htmlEscape(trimmed[2:]) + `</span>`)
		case strings.HasPrefix(trimmed, "```"):
			b.WriteString(`<span style="color:#059669">` + esc + `</span>`)
		case strings.HasPrefix(trimmed, "> "):
			b.WriteString(`<span style="color:#2563eb">&gt; </span><span style="color:#3b82f6;font-style:italic">` + htmlEscape(trimmed[2:]) + `</span>`)
		case trimmed == "---" || trimmed == "***":
			b.WriteString(`<span style="color:#a1a1aa">` + esc + `</span>`)
		case strings.HasPrefix(trimmed, "- ") || strings.HasPrefix(trimmed, "* "):
			indent := htmlEscape(line[:len(line)-len(trimmed)])
			b.WriteString(indent + `<span style="color:#0d9488">` + htmlEscape(string(trimmed[0])+" ") + `</span>` + htmlEscape(trimmed[2:]))
		default:
			b.WriteString(esc)
		}
	}
	return b.String()
}

func htmlEscape(s string) string {
	s = strings.ReplaceAll(s, "&", "&amp;")
	s = strings.ReplaceAll(s, "<", "&lt;")
	s = strings.ReplaceAll(s, ">", "&gt;")
	return s
}

// renderCSVTableHTML renders CSV content as an HTML table string
func renderCSVTableHTML(content string) string {
	lines := strings.Split(strings.TrimSpace(content), "\n")
	if len(lines) == 0 {
		return `<div class="flex-1 flex items-center justify-center"><span class="text-zinc-600 text-sm">Empty file</span></div>`
	}

	var rows [][]string
	for _, line := range lines {
		rows = append(rows, parseCSVLine(line))
	}

	var b strings.Builder
	b.WriteString(`<div class="flex-1 overflow-auto p-4" style="background:white"><table class="w-full border-collapse">`)

	// Header
	if len(rows) > 0 {
		b.WriteString("<thead><tr>")
		for _, cell := range rows[0] {
			b.WriteString(`<th style="padding:8px 16px;text-align:left;font-size:12px;font-weight:600;color:#52525b;text-transform:uppercase;letter-spacing:0.05em;border-bottom:2px solid #e4e4e7;background:#f4f4f5">`)
			b.WriteString(htmlEscape(strings.TrimSpace(cell)))
			b.WriteString("</th>")
		}
		b.WriteString("</tr></thead>")
	}

	// Body
	b.WriteString("<tbody>")
	for i, row := range rows[1:] {
		rowClass := "border-b border-zinc-200"
		if i%2 == 1 {
			rowClass = "border-b border-zinc-200 bg-zinc-50"
		}
		b.WriteString(`<tr class="` + rowClass + `">`)
		for _, cell := range row {
			cellText := strings.TrimSpace(cell)
			cellClass := "px-4 py-1.5 text-sm text-zinc-700 font-mono"
			if isNumericCell(cellText) {
				cellClass = "px-4 py-1.5 text-sm text-zinc-900 font-mono text-right tabular-nums"
			}
			b.WriteString(`<td class="` + cellClass + `">` + htmlEscape(cellText) + `</td>`)
		}
		b.WriteString("</tr>")
	}
	b.WriteString("</tbody></table></div>")
	return b.String()
}

// parseCSVLine does simple comma splitting (no quoting support needed for demo data)
func parseCSVLine(line string) []string {
	return strings.Split(line, ",")
}

// isNumericCell checks if a cell value looks like a number (for right-alignment)
func isNumericCell(s string) bool {
	if s == "" {
		return false
	}
	for i, c := range s {
		if c >= '0' && c <= '9' {
			continue
		}
		if c == '.' || c == ',' {
			continue
		}
		if (c == '-' || c == '+') && i == 0 {
			continue
		}
		return false
	}
	return true
}

// renderFileIcon returns an SVG file icon colored by extension type
func renderFileIcon(ext string) h.Element {
	color := "#71717a" // zinc-500 default
	switch ext {
	case "txt":
		color = "#60a5fa" // blue-400
	case "csv":
		color = "#4ade80" // green-400
	case "md":
		color = "#c084fc" // purple-400
	case "json":
		color = "#facc15" // yellow-400
	case "yaml", "yml":
		color = "#fb923c" // orange-400
	}
	return h.UnsafeRaw(fmt.Sprintf(
		`<svg width="16" height="16" viewBox="0 0 16 16" fill="none" class="shrink-0">`+
			`<path d="M3 1.5A.5.5 0 013.5 1H9l4 4v9.5a.5.5 0 01-.5.5h-9a.5.5 0 01-.5-.5v-13z" fill="%s" fill-opacity="0.15" stroke="%s" stroke-width="1"/>`+
			`<path d="M9 1v4h4" stroke="%s" stroke-width="1" fill="none"/>`+
			`</svg>`,
		color, color, color))
}

// renderFileIconLarge returns a large file icon for the empty state
func renderFileIconLarge() h.Element {
	return h.UnsafeRaw(
		`<svg width="48" height="48" viewBox="0 0 16 16" fill="none" class="mx-auto opacity-20">` +
			`<path d="M3 1.5A.5.5 0 013.5 1H9l4 4v9.5a.5.5 0 01-.5.5h-9a.5.5 0 01-.5-.5v-13z" fill="#71717a" fill-opacity="0.15" stroke="#71717a" stroke-width="1"/>` +
			`<path d="M9 1v4h4" stroke="#71717a" stroke-width="1" fill="none"/>` +
			`</svg>`)
}

// getFileExt returns the extension without the dot
func getFileExt(filename string) string {
	if idx := strings.LastIndex(filename, "."); idx >= 0 {
		return filename[idx+1:]
	}
	return ""
}
