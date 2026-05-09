<script lang="ts">
	import { course } from '$lib/stores/course.js';
	import { language } from '$lib/stores/language.js';
	import { getPageSlug } from '$lib/content/loader.js';

	let { moduleName, pageSlug }: { moduleName: string; pageSlug: string } = $props();

	interface NavTarget {
		module: string;
		page: string;
		label: string;
		moduleTitle?: string;
	}

	let currentLocale = $derived($language);

	let prev = $derived.by((): NavTarget | null => {
		const c = $course;
		if (!c.config || !c.manifest) return null;
		const modules = c.config.modules;
		const moduleIndex = modules.indexOf(moduleName);
		const pages = c.manifest.locales[currentLocale]?.modules[moduleName]?.pages || [];
		const pageIndex = pages.findIndex((p: string) => getPageSlug(p) === pageSlug);

		if (pageIndex > 0) {
			return { module: moduleName, page: getPageSlug(pages[pageIndex - 1]), label: 'Previous' };
		}
		if (moduleIndex > 0) {
			const prevModule = modules[moduleIndex - 1];
			const prevPages = c.manifest.locales[currentLocale]?.modules[prevModule]?.pages || [];
			if (prevPages.length > 0) {
				return { module: prevModule, page: getPageSlug(prevPages[prevPages.length - 1]), label: 'Previous Module', moduleTitle: c.moduleMeta[prevModule]?.title };
			}
		}
		return null;
	});

	let next = $derived.by((): NavTarget | null => {
		const c = $course;
		if (!c.config || !c.manifest) return null;
		const modules = c.config.modules;
		const moduleIndex = modules.indexOf(moduleName);
		const pages = c.manifest.locales[currentLocale]?.modules[moduleName]?.pages || [];
		const pageIndex = pages.findIndex((p: string) => getPageSlug(p) === pageSlug);

		if (pageIndex < pages.length - 1) {
			return { module: moduleName, page: getPageSlug(pages[pageIndex + 1]), label: 'Next' };
		}
		if (moduleIndex < modules.length - 1) {
			const nextModule = modules[moduleIndex + 1];
			const nextPages = c.manifest.locales[currentLocale]?.modules[nextModule]?.pages || [];
			if (nextPages.length > 0) {
				return { module: nextModule, page: getPageSlug(nextPages[0]), label: 'Next Module', moduleTitle: c.moduleMeta[nextModule]?.title };
			}
		}
		return null;
	});
</script>

<div class="flex items-stretch justify-between mt-16 pt-8 border-t border-slate-100 gap-4">
	{#if prev}
		<a
			href="/training/{prev.module}/{prev.page}"
			class="group flex-1 flex items-center gap-3 rounded-xl border border-slate-200 px-5 py-4 text-sm transition-all hover:border-slate-300 hover:shadow-sm"
		>
			<svg width="16" height="16" viewBox="0 0 16 16" fill="none" class="text-slate-400 group-hover:text-accent transition-colors shrink-0"><path d="M10 12L6 8l4-4" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"/></svg>
			<div>
				<div class="text-xs text-slate-400">{prev.label}</div>
				{#if prev.moduleTitle}
					<div class="font-medium text-slate-700 group-hover:text-accent transition-colors">{prev.moduleTitle}</div>
				{/if}
			</div>
		</a>
	{:else}
		<div></div>
	{/if}

	{#if next}
		<a
			href="/training/{next.module}/{next.page}"
			class="group flex-1 flex items-center justify-end gap-3 rounded-xl border border-accent/20 bg-accent/5 px-5 py-4 text-sm transition-all hover:bg-accent/10 hover:border-accent/30 hover:shadow-sm"
		>
			<div class="text-right">
				<div class="text-xs text-accent/60">{next.label}</div>
				{#if next.moduleTitle}
					<div class="font-medium text-accent">{next.moduleTitle}</div>
				{/if}
			</div>
			<svg width="16" height="16" viewBox="0 0 16 16" fill="none" class="text-accent shrink-0"><path d="M6 4l4 4-4 4" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"/></svg>
		</a>
	{/if}
</div>
