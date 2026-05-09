<script lang="ts">
	import type { AgentState } from '$lib/stores/agent.js';

	let { state }: { state: AgentState } = $props();

	function htmlEscape(s: string): string {
		return s.replace(/&/g, '&amp;').replace(/</g, '&lt;').replace(/>/g, '&gt;');
	}

	function roleColor(role: string): string {
		switch (role) {
			case 'system': return 'text-purple-400';
			case 'user': return 'text-accent';
			case 'assistant': return 'text-stone-600';
			case 'tool_call': return 'text-amber-600';
			case 'tool_result': return 'text-stone-400';
			case 'compaction_summary': return 'text-blue-600';
			default: return 'text-stone-400';
		}
	}
</script>

<div class="flex-1 overflow-y-auto px-4 py-3 bg-stone-50/50">
	{#if state.contextMessages.length === 0}
		<div class="flex items-center justify-center h-full text-stone-400 text-sm">
			No context yet
		</div>
	{:else}
		{#each state.contextMessages as cm}
			{@const content = cm.content.length > 500 ? cm.content.substring(0, 500) + '\n...' : cm.content}
			<div class="py-2 border-b border-stone-200/50 last:border-b-0">
				<div class="text-xs font-medium mb-1 {roleColor(cm.role)}">{cm.role}</div>
				<pre style="font-size:11px;line-height:1.5;color:#78716c;font-family:var(--font-mono);white-space:pre-wrap;margin:0;padding:0;background:transparent;border:none">{htmlEscape(content)}</pre>
			</div>
		{/each}
	{/if}
</div>
