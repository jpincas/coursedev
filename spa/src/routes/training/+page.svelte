<script lang="ts">
	import { course } from '$lib/stores/course.js';
	import { progress } from '$lib/stores/progress.js';
	import { getPageSlug } from '$lib/content/loader.js';
	import { isModuleComplete, isModuleUnlocked, getModuleStats } from '$lib/stores/completion.js';
</script>

<div class="mx-auto max-w-3xl px-8 py-12">
	<div class="mb-10">
		<h1 class="text-3xl font-bold text-slate-900 tracking-tight">Your Progress</h1>
		<p class="mt-2 text-slate-500">Complete each module to unlock the next.</p>
	</div>

	{#if $course.config && $course.manifest}
		<!-- Overall progress -->
		{@const totalModules = $course.config.modules.length}
		{@const manifest = $course.manifest}
		{@const completedModules = $course.config.modules.filter((m) => {
			const meta = $course.moduleMeta[m];
			return meta && manifest && isModuleComplete(m, meta, manifest, $progress, $course.quizIdsByModule);
		}).length}

		<div class="rounded-2xl bg-gradient-to-br from-slate-900 to-slate-800 p-6 mb-8 text-white">
			<div class="flex items-center justify-between mb-4">
				<div>
					<div class="text-sm text-slate-400">Overall progress</div>
					<div class="text-2xl font-bold">{completedModules} of {totalModules} modules</div>
				</div>
				<div class="text-3xl font-bold text-accent">{Math.round((completedModules / totalModules) * 100)}%</div>
			</div>
			<div class="h-2 rounded-full bg-white/10 overflow-hidden">
				<div
					class="h-full rounded-full bg-gradient-to-r from-accent to-cyan-400 transition-all duration-500"
					style="width: {(completedModules / totalModules) * 100}%"
				></div>
			</div>
		</div>

		<!-- Module cards -->
		<div class="space-y-3">
			{#each $course.config.modules as moduleName, i}
				{@const meta = $course.moduleMeta[moduleName]}
				{@const unlocked = isModuleUnlocked(moduleName, $course.config.modules, $course.moduleMeta, $course.manifest, $progress, $course.quizIdsByModule)}
				{@const completed = meta && isModuleComplete(moduleName, meta, $course.manifest, $progress, $course.quizIdsByModule)}
				{@const stats = getModuleStats(moduleName, $course.manifest, $progress)}
				{@const firstPage = $course.manifest.modules[moduleName]?.pages[0]}
				{@const firstSlug = firstPage ? getPageSlug(firstPage) : ''}

				<div class="group rounded-xl border transition-all
					{completed
						? 'border-emerald-200 bg-emerald-50/30'
						: unlocked
							? 'border-slate-200 bg-white hover:border-accent/30 hover:shadow-sm'
							: 'border-slate-100 bg-slate-50/50'}">
					<div class="p-5 flex items-center gap-4">
						<!-- Module indicator -->
						<div class="flex-shrink-0 w-10 h-10 rounded-xl flex items-center justify-center text-sm font-bold
							{completed
								? 'bg-emerald-100 text-emerald-600'
								: unlocked
									? 'bg-gradient-to-br from-accent/10 to-cyan-400/10 text-accent border border-accent/20'
									: 'bg-slate-100 text-slate-400'}">
							{#if completed}
								<svg width="18" height="18" viewBox="0 0 16 16" fill="none"><path d="M3 8.5l3.5 3.5L13 4" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"/></svg>
							{:else}
								{i + 1}
							{/if}
						</div>

						<!-- Info -->
						<div class="flex-1 min-w-0">
							<div class="flex items-center gap-2">
								<h3 class="font-semibold text-slate-900 {!unlocked ? 'text-slate-400' : ''}">{meta?.title || moduleName}</h3>
								{#if meta?.estimated_duration}
									<span class="text-[11px] text-slate-400 font-mono bg-slate-100 px-1.5 py-0.5 rounded">{meta.estimated_duration}</span>
								{/if}
							</div>
							<p class="mt-0.5 text-sm text-slate-500 leading-relaxed {!unlocked ? 'text-slate-400' : ''}">{meta?.description}</p>

							<!-- Progress bar -->
							{#if stats.totalPages > 0 && unlocked && !completed}
								<div class="mt-3 flex items-center gap-3">
									<div class="flex-1 h-1 rounded-full bg-slate-100 overflow-hidden">
										<div
											class="h-full rounded-full bg-accent transition-all duration-300"
											style="width: {stats.percent}%"
										></div>
									</div>
									<span class="text-[11px] text-slate-400 font-mono whitespace-nowrap">
										{stats.pagesViewed}/{stats.totalPages}
									</span>
								</div>
							{/if}
						</div>

						<!-- Action -->
						<div class="flex-shrink-0">
							{#if !unlocked}
								<span class="inline-flex items-center gap-1.5 rounded-lg bg-slate-100 px-3 py-2 text-xs text-slate-400">
									<svg width="12" height="12" viewBox="0 0 16 16" fill="none"><rect x="4" y="6" width="8" height="7" rx="1.5" stroke="currentColor" stroke-width="1.5"/><path d="M6 6V4a2 2 0 114 0v2" stroke="currentColor" stroke-width="1.5" stroke-linecap="round"/></svg>
									Locked
								</span>
							{:else if completed}
								<a
									href="/training/{moduleName}/{firstSlug}"
									class="inline-flex items-center gap-1.5 rounded-lg border border-emerald-200 bg-emerald-50 px-3 py-2 text-xs font-medium text-emerald-600 hover:bg-emerald-100 transition-all"
								>
									Review
								</a>
							{:else if stats.pagesViewed > 0}
								{@const lastPage = $progress.modules[moduleName]?.pagesViewed.slice(-1)[0] || firstSlug}
								<a
									href="/training/{moduleName}/{lastPage}"
									class="inline-flex items-center gap-1.5 rounded-lg bg-accent px-4 py-2 text-xs font-semibold text-white hover:bg-accent-hover transition-all hover:shadow-md hover:shadow-accent/20"
								>
									Continue
									<svg width="12" height="12" viewBox="0 0 16 16" fill="none"><path d="M6 4l4 4-4 4" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"/></svg>
								</a>
							{:else}
								<a
									href="/training/{moduleName}/{firstSlug}"
									class="inline-flex items-center gap-1.5 rounded-lg bg-accent px-4 py-2 text-xs font-semibold text-white hover:bg-accent-hover transition-all hover:shadow-md hover:shadow-accent/20"
								>
									Start
									<svg width="12" height="12" viewBox="0 0 16 16" fill="none"><path d="M6 4l4 4-4 4" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"/></svg>
								</a>
							{/if}
						</div>
					</div>
				</div>
			{/each}
		</div>
	{/if}
</div>
