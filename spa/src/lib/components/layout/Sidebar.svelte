<script lang="ts">
	import { get } from 'svelte/store';
	import { page } from '$app/stores';
	import { course } from '$lib/stores/course.js';
	import { progress } from '$lib/stores/progress.js';
	import { language, setLanguage } from '$lib/stores/language.js';
	import { getPageSlug, getPagePath, getPageMarkdown } from '$lib/content/loader.js';
	import { isModuleComplete, isModuleUnlocked, getModuleStats } from '$lib/stores/completion.js';

	let currentModule = $derived($page.params.module || '');
	let currentPage = $derived($page.params.page || '');
	let expandedModule = $state('');
	let isOverviewPage = $derived($page.url.pathname === '/training');

	// Cache for page titles: { locale: { module: { slug: title } } }
	let pageTitleCache: Record<string, Record<string, Record<string, string>>> = {};

	$effect(() => {
		if (currentModule) {
			expandedModule = currentModule;
		}
	});

	// Fetch frontmatter titles for a module's pages
	async function fetchPageTitles(moduleName: string, locale: string) {
		const cacheKey = `${locale}:${moduleName}`;
		if (pageTitleCache[locale]?.[moduleName]) return pageTitleCache[locale][moduleName];

		const mod = get(course)?.manifest?.locales[locale]?.modules[moduleName];
		if (!mod) return {};

		const titles: Record<string, string> = {};
		const promises = mod.pages.map(async (pagePath: string) => {
			const slug = getPageSlug(pagePath);
			try {
				const markdown = await getPageMarkdown(pagePath);
				const match = markdown.match(/^---\n([\s\S]*?)\n---\n/);
				if (match) {
					const fm = match[1];
					const titleMatch = fm.match(/^title:\s*["']?(.+?)["']?$/m);
					if (titleMatch) {
						titles[slug] = titleMatch[1].trim();
						return;
					}
				}
			} catch {
				// ignore
			}
			// Fallback to slug
			titles[slug] = slug.replace(/^\d+-/, '').replace(/-/g, ' ');
		});
		await Promise.all(promises);
		pageTitleCache[locale] = pageTitleCache[locale] || {};
		pageTitleCache[locale][moduleName] = titles;
		return titles;
	}
</script>

<aside class="w-72 flex-shrink-0 bg-slate-950 border-r border-slate-800/50 overflow-y-auto h-full flex flex-col scrollbar-dark">
	<!-- Logo -->
	<div class="p-5 border-b border-slate-800/50">
		<a href="/training" class="flex items-center gap-2.5 group">
			<img src="/images/jon.png" alt="Jon" class="w-7 h-7 rounded-full object-cover ring-2 ring-accent/30 transition-transform group-hover:scale-105" />
			<span class="text-sm font-bold text-white tracking-tight">Learn AI with Jon</span>
		</a>
	</div>

	<!-- Language switcher -->
	<div class="px-3 py-2">
		<div class="flex items-center gap-1.5 bg-white/5 rounded-lg p-1">
			<button
				onclick={() => setLanguage('en')}
				class="flex-1 flex items-center justify-center gap-1 px-2 py-1.5 rounded-md text-xs font-medium transition-all
					{$language === 'en'
						? 'bg-white/10 text-white shadow-sm'
						: 'text-slate-400 hover:text-white hover:bg-white/5'}"
				title="English"
			>
				<span class="text-base leading-none">🇬🇧</span>
			</button>
			<button
				onclick={() => setLanguage('es')}
				class="flex-1 flex items-center justify-center gap-1 px-2 py-1.5 rounded-md text-xs font-medium transition-all
					{$language === 'es'
						? 'bg-white/10 text-white shadow-sm'
						: 'text-slate-400 hover:text-white hover:bg-white/5'}"
				title="Español"
			>
				<span class="text-base leading-none">🇪🇸</span>
			</button>
		</div>
	</div>

	<!-- Course overview link -->
	<div class="px-3 pt-1 pb-1">
		<a
			href="/training"
			class="flex items-center gap-2.5 px-3 py-2 rounded-lg text-sm font-medium transition-all
				{isOverviewPage
					? 'bg-accent/10 text-accent'
					: 'text-slate-400 hover:text-white hover:bg-white/5'}"
		>
			<svg width="16" height="16" viewBox="0 0 16 16" fill="none"><rect x="1" y="1" width="6" height="6" rx="1" stroke="currentColor" stroke-width="1.5"/><rect x="9" y="1" width="6" height="6" rx="1" stroke="currentColor" stroke-width="1.5"/><rect x="1" y="9" width="6" height="6" rx="1" stroke="currentColor" stroke-width="1.5"/><rect x="9" y="9" width="6" height="6" rx="1" stroke="currentColor" stroke-width="1.5"/></svg>
			Overview
		</a>
	</div>

	<!-- Module list -->
	<nav class="flex-1 px-3 pt-1 pb-3">
		{#if $course.config && $course.manifest}
			{#each $course.config.modules as moduleName, i}
				{@const meta = $course.moduleMeta[moduleName]}
				{@const isExpanded = expandedModule === moduleName}
				{@const isCurrent = currentModule === moduleName}
				{@const pages = $course.manifest.locales[$language]?.modules[moduleName]?.pages || []}
				{@const moduleProgress = $progress.modules[moduleName]}
				{@const unlocked = isModuleUnlocked(moduleName, $course.config.modules, $course.moduleMeta, $course.manifest, $progress, $course.quizIdsByModule)}
				{@const completed = meta && isModuleComplete(moduleName, meta, $course.manifest, $progress, $course.quizIdsByModule)}
				{@const stats = getModuleStats(moduleName, $course.manifest, $progress)}

				<div class="mb-0.5">
					{#if unlocked}
						<button
							onclick={() => { expandedModule = isExpanded ? '' : moduleName; fetchPageTitles(moduleName, $language); }}
							class="w-full text-left px-3 py-2 rounded-lg text-[13px] font-medium transition-all flex items-center gap-2
								{isCurrent
									? 'bg-white/10 text-white'
									: 'text-slate-400 hover:text-white hover:bg-white/5'}"
						>
							<span class="w-5 h-5 rounded-md flex items-center justify-center text-[10px] font-bold shrink-0
								{completed
									? 'bg-emerald-500/20 text-emerald-400'
									: isCurrent
										? 'bg-accent/20 text-accent'
										: 'bg-white/5 text-slate-500'}">
								{#if completed}
									<svg width="12" height="12" viewBox="0 0 16 16" fill="none"><path d="M3 8.5l3.5 3.5L13 4" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"/></svg>
								{:else}
									{i + 1}
								{/if}
							</span>
							<span class="flex-1 truncate">{meta?.title || moduleName}</span>
							{#if stats.totalPages > 0 && stats.pagesViewed > 0 && !completed}
								<span class="text-[10px] text-slate-600 font-mono">{stats.percent}%</span>
							{/if}
						</button>
					{:else}
						<div class="w-full text-left px-3 py-2 rounded-lg text-[13px] font-medium text-slate-600 cursor-default flex items-center gap-2">
							<span class="w-5 h-5 rounded-md flex items-center justify-center text-[10px] font-bold bg-white/[0.03] text-slate-700 shrink-0">
								{i + 1}
							</span>
							<span class="flex-1 truncate">{meta?.title || moduleName}</span>
							<svg width="12" height="12" viewBox="0 0 16 16" fill="none" class="text-slate-700 shrink-0"><rect x="4" y="1" width="8" height="10" rx="2" stroke="currentColor" stroke-width="1.5"/><path d="M6 5V4a2 2 0 114 0v1" stroke="currentColor" stroke-width="1.5" stroke-linecap="round"/><circle cx="8" cy="8" r="1" fill="currentColor"/></svg>
						</div>
					{/if}

					<!-- Page list -->
					{#if isExpanded && unlocked}
						<div class="ml-4 mt-0.5 mb-1 border-l border-slate-800/80 pl-3 space-y-px">
							{#await fetchPageTitles(moduleName, $language) then titles}
								{#each pages as pagePath}
									{@const slug = getPageSlug(pagePath)}
									{@const pageTitle = titles?.[slug] || slug.replace(/^\d+-/, '').replace(/-/g, ' ')}
									{@const isCurrentPage = isCurrent && currentPage === slug}
									{@const viewed = moduleProgress?.pagesViewed?.includes(slug)}
									<a
										href="/training/{moduleName}/{slug}"
										class="flex items-center gap-2 px-2.5 py-1.5 rounded-md text-[13px] transition-all truncate
											{isCurrentPage
												? 'bg-accent/15 text-accent font-medium'
												: viewed
													? 'text-slate-500 hover:text-slate-300 hover:bg-white/5'
													: 'text-slate-400 hover:text-white hover:bg-white/5'}"
									>
										{#if viewed && !isCurrentPage}
											<svg width="10" height="10" viewBox="0 0 16 16" fill="none" class="shrink-0 text-slate-600"><path d="M3 8.5l3.5 3.5L13 4" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"/></svg>
										{:else if isCurrentPage}
											<span class="w-1.5 h-1.5 rounded-full bg-accent shrink-0"></span>
										{:else}
											<span class="w-1.5 h-1.5 rounded-full bg-slate-700 shrink-0"></span>
										{/if}
										<span class="truncate">{pageTitle}</span>
									</a>
								{/each}
							{/await}
						</div>
					{/if}
				</div>
			{/each}
		{/if}
	</nav>

	<!-- Back to site -->
	<div class="p-3 border-t border-slate-800/50 shrink-0">
		<a
			href="/"
			class="flex items-center gap-2 px-3 py-2 rounded-lg text-sm text-slate-500 hover:text-slate-300 hover:bg-white/5 transition-all"
		>
			<svg width="14" height="14" viewBox="0 0 16 16" fill="none"><path d="M10 12L6 8l4-4" stroke="currentColor" stroke-width="1.5" stroke-linecap="round" stroke-linejoin="round"/></svg>
			Back to site
		</a>
	</div>
</aside>
