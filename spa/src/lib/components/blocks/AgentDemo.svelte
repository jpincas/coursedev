<script lang="ts">
	import type { AgentBlock } from '$lib/content/types.js';
	import { initAgentState, advance, type AgentState } from '$lib/stores/agent.js';
	import AgentSidebar from './agent/AgentSidebar.svelte';
	import AgentWorkspace from './agent/AgentWorkspace.svelte';

	let { agent }: { agent: AgentBlock } = $props();

	let agentConfig = $derived(agent);
	let state = $state<AgentState>(initAgentState(agentConfig));

	function handleAdvance() {
		state = advance({ ...state, chatMessages: [...state.chatMessages], contextMessages: [...state.contextMessages], scratchpad: { ...state.scratchpad }, expandedFolders: { ...state.expandedFolders } });
	}

	function handleReset() {
		state = initAgentState(agentConfig);
	}

	function handleToggleSystem() {
		state.showSystem = !state.showSystem;
	}

	function handleToggleTools() {
		state.showToolCalls = !state.showToolCalls;
	}

	function handleToggleContext() {
		state.showFullContext = !state.showFullContext;
	}

	function handleToggleSidebar() {
		state.sidebarOpen = !state.sidebarOpen;
	}

	function handleToggleScratchpad() {
		state.scratchpadOpen = !state.scratchpadOpen;
	}

	function handleViewFile(filename: string) {
		state.viewingFile = state.viewingFile === filename ? '' : filename;
	}

	function handleToggleFolder(path: string) {
		state.expandedFolders = { ...state.expandedFolders };
		if (state.expandedFolders[path]) {
			delete state.expandedFolders[path];
		} else {
			state.expandedFolders[path] = true;
		}
	}

	function handleOpenWorkspace() {
		state.workspaceOpen = true;
	}

	function handleCloseWorkspace() {
		state.workspaceOpen = false;
	}

	const handlers = {
		advance: handleAdvance,
		reset: handleReset,
		toggleSystem: handleToggleSystem,
		toggleTools: handleToggleTools,
		toggleContext: handleToggleContext,
		toggleSidebar: handleToggleSidebar,
		toggleScratchpad: handleToggleScratchpad,
		viewFile: handleViewFile,
		toggleFolder: handleToggleFolder,
		openWorkspace: handleOpenWorkspace,
		closeWorkspace: handleCloseWorkspace
	};

	// Determine if this agent has scratchpad files (workspace mode) or is chat-only
	let hasWorkspace = $derived(Object.keys(agent.scratchpad).length > 0);
</script>

{#if state.workspaceOpen}
	<!-- Full-screen workspace mode -->
	<div class="fixed inset-0 z-50 flex">
		<AgentWorkspace {state} {handlers} />
	</div>
{:else if hasWorkspace}
	<!-- Launch card for workspace demos -->
	<div class="rounded-2xl border border-slate-200 bg-gradient-to-br from-slate-900 via-slate-800 to-slate-900 p-10 text-center relative overflow-hidden">
		<div class="absolute inset-0 hero-grid opacity-50"></div>
		<div class="absolute top-1/2 left-1/2 -translate-x-1/2 -translate-y-1/2 w-64 h-64 bg-accent/10 rounded-full blur-3xl"></div>
		<div class="relative">
			<div class="w-14 h-14 rounded-2xl bg-white/10 backdrop-blur-sm border border-white/10 flex items-center justify-center mx-auto mb-5">
				<svg width="24" height="24" viewBox="0 0 16 16" fill="none"><path d="M5 3l8 5-8 5V3z" fill="#14b8a6" stroke="#14b8a6" stroke-width="1"/></svg>
			</div>
			<h3 class="text-lg font-bold text-white mb-1">{agent.title}</h3>
			<p class="text-sm text-slate-400 mb-6">{Object.keys(agent.scratchpad).length} files · interactive walkthrough</p>
			<button
				onclick={handleOpenWorkspace}
				class="px-8 py-3 rounded-full bg-accent text-white font-semibold text-sm cursor-pointer border-none hover:bg-accent-hover transition-all hover:shadow-lg hover:shadow-accent/25 hover:-translate-y-0.5"
			>
				Launch Demo
			</button>
		</div>
	</div>
{:else}
	<!-- Inline sidebar demo (no workspace) -->
	<div class="rounded-2xl border border-slate-200 overflow-hidden shadow-sm" style="height: 600px;">
		<AgentSidebar {state} {handlers} />
	</div>
{/if}
