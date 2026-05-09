<script lang="ts">
	import type { AgentState } from '$lib/stores/agent.js';

	let { state, handlers }: { state: AgentState; handlers: Record<string, Function> } = $props();
</script>

<div class="flex items-center justify-between px-4 py-3 border-b border-stone-200 shrink-0">
	<div class="flex items-center gap-2">
		<span class="text-sm font-semibold text-stone-800">{state.config.title}</span>
	</div>
	<div class="flex items-center gap-2">
		<!-- Context view toggle -->
		<button
			onclick={() => handlers.toggleContext()}
			class="cursor-pointer border-none bg-transparent transition-colors duration-150
				{state.showFullContext ? 'text-blue-500 hover:text-blue-600' : 'text-stone-400 hover:text-blue-500'}"
			title="Full context"
		>
			<svg width="18" height="18" viewBox="0 0 18 18" fill="none"><path d="M6.75 4.5L3 9l3.75 4.5M11.25 4.5L15 9l-3.75 4.5" stroke="currentColor" stroke-width="1.5" stroke-linecap="round" stroke-linejoin="round"/></svg>
		</button>

		<!-- System prompt icon -->
		{#if state.config.visibility.systemPrompt !== 'hidden' && state.config.system}
			<button
				onclick={() => handlers.toggleSystem()}
				class="text-stone-400 hover:text-purple-500 cursor-pointer border-none bg-transparent transition-colors duration-150"
				title="System prompt"
			>
				<svg width="18" height="18" viewBox="0 0 18 18" fill="none"><path d="M9 2.25A6.75 6.75 0 1015.75 9 6.75 6.75 0 009 2.25zm0 12A5.25 5.25 0 1114.25 9 5.25 5.25 0 019 14.25z" fill="currentColor"/><path d="M9 5.25a.75.75 0 00-.75.75v3a.75.75 0 001.5 0V6A.75.75 0 009 5.25zM9 11.25a.75.75 0 100 1.5.75.75 0 000-1.5z" fill="currentColor"/></svg>
			</button>
		{/if}

		<!-- Close sidebar -->
		<button
			onclick={() => handlers.toggleSidebar()}
			class="text-stone-400 hover:text-stone-600 cursor-pointer border-none bg-transparent text-sm"
		>✕</button>
	</div>
</div>
