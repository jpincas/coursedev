<script lang="ts">
	import type { AgentState } from '$lib/stores/agent.js';

	let { state, handlers }: { state: AgentState; handlers: Record<string, Function> } = $props();

	let atEnd = $derived(state.scriptIndex >= state.config.script.length);
	let nextEvent = $derived(!atEnd ? state.config.script[state.scriptIndex] : null);
</script>

<div class="px-3 py-3 border-t border-stone-200 shrink-0">
	<!-- Current note (if any) -->
	{#if state.currentNote}
		<div class="flex items-start gap-2.5 px-3 py-2.5 mb-2 rounded-xl bg-amber-50/80 border border-amber-200/50">
			<span class="text-base leading-5 shrink-0 mt-0.5">🧑‍🏫</span>
			<span class="text-sm text-amber-900/70 leading-relaxed">{state.currentNote}</span>
		</div>
	{/if}

	{#if atEnd}
		<!-- Demo complete -->
		<div class="flex items-center justify-between px-4 py-3 rounded-2xl bg-stone-50 border border-stone-200">
			<span class="text-sm text-accent font-medium">✓ Demo complete</span>
			<button
				onclick={() => handlers.reset()}
				class="py-1.5 px-3 text-sm rounded-lg bg-stone-200 text-stone-600 font-medium cursor-pointer transition-all duration-150 hover:bg-stone-300 border-none"
			>Done</button>
		</div>
	{:else if nextEvent?.type === 'user'}
		<!-- User message to send -->
		<div
			onclick={() => handlers.advance()}
			class="flex items-end gap-2 rounded-2xl border border-stone-300 bg-white px-4 py-3 cursor-pointer transition-all duration-150 hover:border-accent/50"
			role="button"
			tabindex="0"
			onkeydown={(e) => { if (e.key === 'Enter') handlers.advance() }}
		>
			<div class="flex-1 text-sm text-stone-800 leading-relaxed min-h-[20px]">
				{nextEvent.content}
			</div>
			<!-- Send button -->
			<div class="w-8 h-8 rounded-full bg-accent flex items-center justify-center shrink-0">
				<svg width="16" height="16" viewBox="0 0 16 16" fill="none"><path d="M8 12V4M8 4L4 8M8 4L12 8" stroke="white" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"/></svg>
			</div>
		</div>
	{:else}
		<!-- Continue button for non-user events -->
		{@const label = nextEvent?.type === 'compaction' ? 'Compact context' : nextEvent?.type === 'clear' ? 'Clear & continue' : 'Continue'}
		<div
			onclick={() => handlers.advance()}
			class="flex items-center justify-center rounded-xl border border-stone-200 bg-stone-50 px-3 py-2 cursor-pointer transition-all duration-150 hover:bg-stone-100 hover:border-stone-300"
			role="button"
			tabindex="0"
			onkeydown={(e) => { if (e.key === 'Enter') handlers.advance() }}
		>
			<span class="text-xs text-stone-500 font-medium">{label} →</span>
		</div>
	{/if}
</div>
