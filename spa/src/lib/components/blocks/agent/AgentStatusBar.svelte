<script lang="ts">
	import type { AgentState } from '$lib/stores/agent.js';

	let { state }: { state: AgentState } = $props();
</script>

{#if state.config.visibility.tokenCount === 'visible'}
	<div class="px-3 pb-1 flex justify-end text-xs shrink-0">
		<span class="text-stone-400">~{state.tokenCount} tokens</span>
	</div>
{/if}
