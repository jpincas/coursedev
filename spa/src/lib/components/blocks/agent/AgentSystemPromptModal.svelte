<script lang="ts">
	import type { AgentState } from '$lib/stores/agent.js';

	let { state, handlers }: { state: AgentState; handlers: Record<string, Function> } = $props();
</script>

<div class="absolute inset-0 z-50 flex flex-col">
	<!-- Backdrop -->
	<div
		class="absolute inset-0 bg-black/30"
		onclick={() => handlers.toggleSystem()}
		role="button"
		tabindex="-1"
		onkeydown={(e) => { if (e.key === 'Escape') handlers.toggleSystem() }}
	></div>
	<!-- Modal card -->
	<div
		class="relative mx-5 mt-14 mb-5 flex flex-col rounded-xl bg-white border border-stone-200 overflow-hidden max-h-[70%]"
		style="box-shadow: 0 8px 30px rgba(0,0,0,0.15)"
	>
		<!-- Header -->
		<div class="flex items-center justify-between px-4 py-3 border-b border-stone-200 shrink-0">
			<div class="flex items-center gap-2">
				<svg width="14" height="14" viewBox="0 0 18 18" fill="none"><path d="M9 2.25A6.75 6.75 0 1015.75 9 6.75 6.75 0 009 2.25zm0 12A5.25 5.25 0 1114.25 9 5.25 5.25 0 019 14.25z" fill="#a855f7"/><path d="M9 5.25a.75.75 0 00-.75.75v3a.75.75 0 001.5 0V6A.75.75 0 009 5.25zM9 11.25a.75.75 0 100 1.5.75.75 0 000-1.5z" fill="#a855f7"/></svg>
				<span class="text-xs font-semibold text-purple-500 uppercase tracking-wider">System Prompt</span>
			</div>
			<button
				onclick={() => handlers.toggleSystem()}
				class="text-stone-400 hover:text-stone-600 cursor-pointer border-none bg-transparent text-sm"
			>✕</button>
		</div>
		<!-- Content -->
		<div class="flex-1 overflow-y-auto px-4 py-3">
			<pre class="text-xs text-stone-600 font-mono whitespace-pre-wrap m-0 bg-transparent border-none p-0 leading-relaxed">{state.config.system}</pre>
		</div>
	</div>
</div>
