<script lang="ts">
	import { onMount } from 'svelte';
	import { onDestroy } from 'svelte';
	import { course } from '$lib/stores/course.js';
	import { language } from '$lib/stores/language.js';
	import Sidebar from '$lib/components/layout/Sidebar.svelte';

	let { children } = $props();
	let loaded = $state(false);

	async function reload() {
		await course.load();
	}

	onMount(async () => {
		await course.load();
		loaded = true;
	});

	let unsubscribe: (() => void) | null = null;
	onMount(() => {
		unsubscribe = language.subscribe(async () => {
			// Language changed - reload course data
			await course.load();
			loaded = true;
		});
	});

	onDestroy(() => {
		unsubscribe?.();
	});
</script>

{#if !loaded}
	<div class="flex h-screen items-center justify-center bg-white">
		<div class="flex items-center gap-3">
			<img src="/images/jon.png" alt="Jon" class="w-8 h-8 rounded-full object-cover animate-pulse-soft" />
			<span class="text-slate-400 text-sm">Loading course...</span>
		</div>
	</div>
{:else}
	<div class="flex h-screen">
		<Sidebar />
		<main class="flex-1 overflow-y-auto bg-white scrollbar-light">
			{@render children()}
		</main>
	</div>
{/if}
