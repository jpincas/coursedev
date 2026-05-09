<script lang="ts">
	import { page } from '$app/stores';
	import { goto } from '$app/navigation';
	import { get } from 'svelte/store';
	import { course } from '$lib/stores/course.js';
	import { progress } from '$lib/stores/progress.js';
	import { language } from '$lib/stores/language.js';
	import { getPageMarkdown, getPagePath, getLocale } from '$lib/content/loader.js';
	import { parsePage } from '$lib/content/parser.js';
	import { isModuleUnlocked } from '$lib/stores/completion.js';
	import type { ParsedPage } from '$lib/content/types.js';
	import PageContent from '$lib/components/layout/PageContent.svelte';
	import PageNav from '$lib/components/layout/PageNav.svelte';

	let parsedPage = $state<ParsedPage | null>(null);
	let loading = $state(true);
	let error = $state('');
	let locked = $state(false);

	// Reactive: reload when route params change
	let moduleName = $derived($page.params.module ?? '');
	let pageSlug = $derived($page.params.page ?? '');

	$effect(() => {
		if (moduleName && pageSlug) {
			loadPage(moduleName, pageSlug);
		}
	});

	// Reload when language changes
	$effect(() => {
		// Access $language to trigger re-run on language change
		void $language;
		if (moduleName && pageSlug) {
			loadPage(moduleName, pageSlug);
		}
	});

	async function loadPage(mod: string, pg: string) {
		loading = true;
		error = '';
		locked = false;
		parsedPage = null;

		try {
			// Check module is unlocked
			const c = get(course);
			const p = get(progress);
			if (c.config && c.manifest) {
				const unlocked = isModuleUnlocked(
					mod,
					c.config.modules,
					c.moduleMeta,
					c.manifest,
					p,
					c.quizIdsByModule
				);
				if (!unlocked) {
					locked = true;
					loading = false;
					return;
				}
			}

			const pagePath = getPagePath(mod, pg, getLocale());
			if (!pagePath) {
				error = `Page not found: ${mod}/${pg}`;
				loading = false;
				return;
			}

			const markdown = await getPageMarkdown(pagePath);
			parsedPage = await parsePage(markdown);

			// Mark as viewed
			progress.markPageViewed(mod, pg);

			// Scroll to top
			window.scrollTo(0, 0);
		} catch (e) {
			error = e instanceof Error ? e.message : 'Failed to load page';
		} finally {
			loading = false;
		}
	}
</script>

<div class="mx-auto max-w-3xl px-8 py-12">
	{#if loading}
		<div class="flex items-center gap-3 py-20 justify-center">
			<div class="w-5 h-5 rounded-full border-2 border-accent border-t-transparent animate-spin"></div>
			<span class="text-slate-400 text-sm">Loading...</span>
		</div>
	{:else if locked}
		<div class="text-center py-24">
			<div class="w-16 h-16 rounded-2xl bg-slate-100 flex items-center justify-center mx-auto mb-5">
				<svg width="28" height="28" viewBox="0 0 16 16" fill="none" class="text-slate-400"><rect x="3" y="7" width="10" height="7" rx="2" stroke="currentColor" stroke-width="1.5"/><path d="M5 7V5a3 3 0 116 0v2" stroke="currentColor" stroke-width="1.5" stroke-linecap="round"/><circle cx="8" cy="10.5" r="1" fill="currentColor"/></svg>
			</div>
			<h2 class="text-xl font-bold text-slate-900 mb-2">Module Locked</h2>
			<p class="text-slate-500 mb-8">Complete the previous modules to unlock this one.</p>
			<a
				href="/training"
				class="rounded-full bg-accent px-6 py-2.5 text-sm font-semibold text-white transition-all hover:bg-accent-hover hover:shadow-md hover:shadow-accent/20 inline-block"
			>
				View Course Overview
			</a>
		</div>
	{:else if error}
		<div class="rounded-xl border border-red-200 bg-red-50 p-5 text-red-700 text-sm">
			{error}
		</div>
	{:else if parsedPage}
		{#if parsedPage.meta.title}
			<h1 class="text-3xl sm:text-4xl font-bold text-slate-900 tracking-tight mb-10 leading-tight">{parsedPage.meta.title}</h1>
		{/if}

		<PageContent page={parsedPage} {moduleName} />

		<PageNav {moduleName} {pageSlug} />
	{/if}
</div>
