<script lang="ts">
	import { tick } from 'svelte';
	import type { AgentState } from '$lib/stores/agent.js';
	import { shouldShowPanel, chatAnimStyle, displayToolName } from '$lib/stores/agent.js';
	import { Marked } from 'marked';

	let { state }: { state: AgentState } = $props();
	let messagesEl: HTMLDivElement;

	const chatMarked = new Marked({ gfm: true, breaks: false });

	function renderChatMarkdown(text: string): string {
		// marked.parse returns string|Promise<string>, but synchronously when no async extensions
		return chatMarked.parse(text) as string;
	}

	function htmlEscape(s: string): string {
		return s.replace(/&/g, '&amp;').replace(/</g, '&lt;').replace(/>/g, '&gt;');
	}

	function toolCallIcon(toolName: string): string {
		const icons: Record<string, string> = {
			web_search: `<svg width="12" height="12" viewBox="0 0 16 16" fill="none"><circle cx="7" cy="7" r="4.5" stroke="#fbbf24" stroke-width="1.5"/><path d="M10.5 10.5L14 14" stroke="#fbbf24" stroke-width="1.5" stroke-linecap="round"/></svg>`,
			fetch_url: `<svg width="12" height="12" viewBox="0 0 16 16" fill="none"><circle cx="8" cy="8" r="6" stroke="#fbbf24" stroke-width="1.2"/><ellipse cx="8" cy="8" rx="3" ry="6" stroke="#fbbf24" stroke-width="1.2"/><path d="M2 8h12M3 4.5h10M3 11.5h10" stroke="#fbbf24" stroke-width="1"/></svg>`,
			list_files: `<svg width="12" height="12" viewBox="0 0 16 16" fill="none"><path d="M2 4h12M2 8h12M2 12h8" stroke="#fbbf24" stroke-width="1.5" stroke-linecap="round"/></svg>`,
			move_file: `<svg width="12" height="12" viewBox="0 0 16 16" fill="none"><path d="M3 8h10M9 4l4 4-4 4" stroke="#fbbf24" stroke-width="1.5" stroke-linecap="round" stroke-linejoin="round"/></svg>`,
			create_folder: `<svg width="12" height="12" viewBox="0 0 16 16" fill="none"><path d="M1.5 3A1.5 1.5 0 013 1.5h3.293a1 1 0 01.707.293L8.414 3.207a1 1 0 00.707.293H13A1.5 1.5 0 0114.5 5v8a1.5 1.5 0 01-1.5 1.5H3A1.5 1.5 0 011.5 13V3z" fill="#fbbf24" fill-opacity="0.3" stroke="#fbbf24" stroke-width="0.8"/><path d="M8 7v4M6 9h4" stroke="#fbbf24" stroke-width="1.2" stroke-linecap="round"/></svg>`,
			delete_file: `<svg width="12" height="12" viewBox="0 0 16 16" fill="none"><path d="M5.5 2h5M2 4h12M4 4l1 10h6l1-10M6.5 7v4M9.5 7v4" stroke="#fbbf24" stroke-width="1.2" stroke-linecap="round" stroke-linejoin="round"/></svg>`,
			run_code: `<svg width="12" height="12" viewBox="0 0 16 16" fill="none"><rect x="1" y="2" width="14" height="12" rx="1.5" stroke="#fbbf24" stroke-width="1.2"/><path d="M4 6l3 2-3 2" stroke="#fbbf24" stroke-width="1.5" stroke-linecap="round" stroke-linejoin="round"/><path d="M9 10h3" stroke="#fbbf24" stroke-width="1.5" stroke-linecap="round"/></svg>`
		};
		return icons[toolName] || `<svg width="12" height="12" viewBox="0 0 16 16" fill="none"><path d="M14.25 6.14L13.07 5l.41-1.66a.38.38 0 00-.11-.36.37.37 0 00-.36-.1L11.35 3.3 10.2 2.11a.37.37 0 00-.53 0L8.54 3.24 7.47 2.87a.38.38 0 00-.42.09L1.17 8.84a.38.38 0 000 .53l2.12 2.12-1.72 1.72a.75.75 0 001.06 1.06l1.72-1.72 2.12 2.12a.38.38 0 00.53 0l5.88-5.88a.38.38 0 00.09-.42l-.37-1.07 1.13-1.13a.37.37 0 000-.53z" fill="#fbbf24"/></svg>`;
	}

	function toolResultIcon(toolName: string): string {
		const muted: Record<string, string> = {
			web_search: `<svg width="11" height="11" viewBox="0 0 16 16" fill="none"><circle cx="7" cy="7" r="4.5" stroke="#d6d3d1" stroke-width="1.5"/><path d="M10.5 10.5L14 14" stroke="#d6d3d1" stroke-width="1.5" stroke-linecap="round"/></svg>`,
			fetch_url: `<svg width="11" height="11" viewBox="0 0 16 16" fill="none"><circle cx="8" cy="8" r="6" stroke="#d6d3d1" stroke-width="1.2"/><ellipse cx="8" cy="8" rx="3" ry="6" stroke="#d6d3d1" stroke-width="1.2"/><path d="M2 8h12M3 4.5h10M3 11.5h10" stroke="#d6d3d1" stroke-width="1"/></svg>`,
			list_files: `<svg width="11" height="11" viewBox="0 0 16 16" fill="none"><path d="M2 4h12M2 8h12M2 12h8" stroke="#d6d3d1" stroke-width="1.5" stroke-linecap="round"/></svg>`,
			move_file: `<svg width="11" height="11" viewBox="0 0 16 16" fill="none"><path d="M3 8h10M9 4l4 4-4 4" stroke="#d6d3d1" stroke-width="1.5" stroke-linecap="round" stroke-linejoin="round"/></svg>`,
			create_folder: `<svg width="11" height="11" viewBox="0 0 16 16" fill="none"><path d="M1.5 3A1.5 1.5 0 013 1.5h3.293a1 1 0 01.707.293L8.414 3.207a1 1 0 00.707.293H13A1.5 1.5 0 0114.5 5v8a1.5 1.5 0 01-1.5 1.5H3A1.5 1.5 0 011.5 13V3z" fill="#d6d3d1" fill-opacity="0.2" stroke="#d6d3d1" stroke-width="0.8"/></svg>`,
			delete_file: `<svg width="11" height="11" viewBox="0 0 16 16" fill="none"><path d="M5.5 2h5M2 4h12M4 4l1 10h6l1-10M6.5 7v4M9.5 7v4" stroke="#d6d3d1" stroke-width="1.2" stroke-linecap="round" stroke-linejoin="round"/></svg>`,
			run_code: `<svg width="11" height="11" viewBox="0 0 16 16" fill="none"><rect x="1" y="2" width="14" height="12" rx="1.5" stroke="#d6d3d1" stroke-width="1.2"/><path d="M4 6l3 2-3 2" stroke="#d6d3d1" stroke-width="1.5" stroke-linecap="round" stroke-linejoin="round"/><path d="M9 10h3" stroke="#d6d3d1" stroke-width="1.5" stroke-linecap="round"/></svg>`
		};
		return muted[toolName] || `<svg width="11" height="11" viewBox="0 0 16 16" fill="none"><path d="M14 4.5V14a1 1 0 01-1 1H3a1 1 0 01-1-1V2a1 1 0 011-1h6.5L14 4.5z" fill="#d6d3d1" fill-opacity="0.3" stroke="#d6d3d1" stroke-width="1"/><path d="M9.5 1v4H14" stroke="#d6d3d1" stroke-width="1" fill="none"/></svg>`;
	}

	// Auto-scroll when messages change
	$effect(() => {
		// Track state.chatMessages.length to trigger on new messages
		const _len = state.chatMessages.length;
		tick().then(() => {
			if (messagesEl) {
				// Find max animation delay for smooth scroll timing
				const animated = state.chatMessages.filter(m => m.groupIdx >= 0);
				const maxDelay = animated.length > 0
					? Math.max(...animated.map(m => m.groupIdx)) * 400 + 400
					: 0;
				setTimeout(() => {
					messagesEl?.scrollTo({ top: messagesEl.scrollHeight, behavior: 'smooth' });
				}, maxDelay);
			}
		});
	});
</script>

<div bind:this={messagesEl} class="flex-1 overflow-y-auto px-3 py-3 flex flex-col gap-2">
	{#if state.chatMessages.length === 0}
		<div class="flex items-center justify-center h-full text-stone-400 text-sm">
			Send the first message to begin
		</div>
	{:else}
		{#each state.chatMessages as msg}
			{#if msg.type === 'user'}
				<div class="flex items-end justify-end gap-2" style={chatAnimStyle(msg.groupIdx)}>
					<div class="max-w-[80%] py-2.5 px-3.5 rounded-2xl rounded-br-sm bg-accent text-white text-sm leading-relaxed">
						{msg.content}
					</div>
					<div class="w-7 h-7 rounded-full flex items-center justify-center shrink-0 text-sm bg-accent/15">
						<span class="text-xs">👤</span>
					</div>
				</div>
			{:else if msg.type === 'assistant'}
				<div class="flex items-end justify-start gap-2" style={chatAnimStyle(msg.groupIdx)}>
					<div class="w-7 h-7 rounded-full flex items-center justify-center shrink-0 text-sm bg-purple-50">
						<span class="text-xs">✨</span>
					</div>
					<div class="max-w-[80%] py-2.5 px-3.5 rounded-2xl rounded-bl-sm bg-stone-100 text-stone-700 text-sm leading-relaxed agent-chat-md">
						{@html renderChatMarkdown(msg.content)}
					</div>
				</div>
			{:else if msg.type === 'tool_call' && shouldShowPanel(state.config.visibility.toolCalls, state.showToolCalls)}
				<div class="mx-1" style={chatAnimStyle(msg.groupIdx)}>
					<div class="rounded-lg overflow-hidden border border-stone-300">
						<div class="flex items-center gap-1.5 px-3 py-1.5 bg-stone-800">
							{@html toolCallIcon(msg.toolName)}
							<span class="text-xs font-semibold text-white">{displayToolName(msg.toolName)}</span>
						</div>
						<div class="px-3 py-2 bg-stone-50 font-mono">
							{#each Object.keys(msg.toolArgs).sort() as key}
								{@const val = msg.toolArgs[key].length > 120 ? msg.toolArgs[key].substring(0, 120) + '...' : msg.toolArgs[key]}
								<div style="font-size:11px;line-height:1.5">
									<span style="color:#a1a1aa">{key}: </span><span style="color:#57534e">{val}</span>
								</div>
							{/each}
						</div>
					</div>
				</div>
			{:else if msg.type === 'tool_result' && shouldShowPanel(state.config.visibility.toolCalls, state.showToolCalls)}
				{@const content = msg.content.length > 300 ? msg.content.substring(0, 300) + '\n...' : msg.content}
				<div class="mx-1" style={chatAnimStyle(msg.groupIdx)}>
					<div class="rounded-lg overflow-hidden border border-stone-200/80">
						<div class="flex items-center gap-1.5 px-3 py-1 bg-stone-100/80 border-b border-stone-200/60">
							{@html toolResultIcon(msg.toolName)}
							<span class="text-stone-400" style="font-size:11px">{displayToolName(msg.toolName)}</span>
						</div>
						<div class="px-3 py-2 bg-stone-50/50">
							<pre style="font-size:11px;line-height:1.5;color:#a1a1aa;font-family:var(--font-mono);white-space:pre-wrap;margin:0;padding:0;background:transparent;border:none">{htmlEscape(content)}</pre>
						</div>
					</div>
				</div>
			{:else if msg.type === 'compaction_divider'}
				<div class="flex items-center gap-2 py-2 mx-1">
					<div class="flex-1 h-px bg-stone-200"></div>
					<span class="text-xs text-stone-400 font-medium whitespace-nowrap">Context compacted</span>
					<div class="flex-1 h-px bg-stone-200"></div>
				</div>
			{:else if msg.type === 'clear_divider'}
				<div class="flex items-center gap-2 py-2 mx-1">
					<div class="flex-1 h-px bg-stone-200"></div>
					<span class="text-xs text-stone-400 font-medium whitespace-nowrap">{msg.content || 'Conversation cleared'}</span>
					<div class="flex-1 h-px bg-stone-200"></div>
				</div>
			{/if}
		{/each}
	{/if}
</div>
