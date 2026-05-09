<script lang="ts">
	import type { AgentState } from '$lib/stores/agent.js';
	import AgentTitleBar from './AgentTitleBar.svelte';
	import AgentChatMessages from './AgentChatMessages.svelte';
	import AgentFullContextView from './AgentFullContextView.svelte';
	import AgentStatusBar from './AgentStatusBar.svelte';
	import AgentChatInput from './AgentChatInput.svelte';
	import AgentSystemPromptModal from './AgentSystemPromptModal.svelte';

	let { state, handlers }: { state: AgentState; handlers: Record<string, Function> } = $props();
</script>

{#if !state.sidebarOpen}
	<aside class="w-14 h-full bg-white border-l border-stone-200 flex flex-col items-center py-4 shrink-0">
		<button
			onclick={() => handlers.toggleSidebar()}
			class="py-2 px-1 text-xs text-stone-500 cursor-pointer border-none bg-transparent hover:text-accent"
			style="writing-mode: vertical-rl; text-orientation: mixed"
		>
			{state.config.title}
		</button>
	</aside>
{:else}
	<aside class="w-[480px] h-full bg-white border-l border-stone-200 flex flex-col overflow-hidden shrink-0 relative">
		<AgentTitleBar {state} {handlers} />

		{#if state.showFullContext}
			<AgentFullContextView {state} />
		{:else}
			<AgentChatMessages {state} />
		{/if}

		{#if state.config.visibility.toolCalls === 'toggleable'}
			<div class="mx-3 shrink-0">
				<button
					onclick={() => handlers.toggleTools()}
					class="w-full py-1 px-2 text-xs font-medium cursor-pointer border-none bg-transparent hover:text-stone-600 text-left transition-colors duration-150
						{state.showToolCalls ? 'text-stone-500' : 'text-stone-400'}"
				>
					{state.showToolCalls ? '▾' : '▸'} Tool Calls
				</button>
			</div>
		{/if}

		<AgentStatusBar {state} />
		<AgentChatInput {state} {handlers} />

		{#if state.showSystem && state.config.system}
			<AgentSystemPromptModal {state} {handlers} />
		{/if}
	</aside>
{/if}
