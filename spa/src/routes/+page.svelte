<script lang="ts">
	import { onMount } from 'svelte';
	import { get } from 'svelte/store';
	import { progress } from '$lib/stores/progress.js';
	import { course } from '$lib/stores/course.js';

	let loaded = $state(false);
	let hasProgress = $state(false);
	let resumeUrl = $state('/training');

	onMount(async () => {
		await course.load();
		const p = get(progress);
		if (p.lastPosition.module && p.lastPosition.page) {
			hasProgress = true;
			resumeUrl = `/training/${p.lastPosition.module}/${p.lastPosition.page}`;
		}
		loaded = true;
	});

	const features = [
		{ icon: '🧠', title: 'How LLMs Actually Work', desc: 'Understand the prediction engine behind every AI tool — no PhD required.' },
		{ icon: '🎯', title: 'Prompting That Delivers', desc: 'Move beyond trial and error. Learn frameworks that produce consistent results.' },
		{ icon: '🤖', title: 'Interactive AI Demos', desc: 'Watch scripted AI conversations unfold step by step. See exactly what works and why.' },
		{ icon: '📁', title: 'Real-World Workflows', desc: 'From document creation to data analysis — practical patterns you can use today.' },
		{ icon: '🔗', title: 'Advanced Patterns', desc: 'Multi-agent delegation, tool use, and automation strategies for power users.' },
		{ icon: '⚖️', title: 'Risks & Responsibility', desc: 'Navigate the real challenges: hallucinations, bias, privacy, and professional judgement.' },
	];
</script>

<div class="min-h-screen bg-surface-darker text-white">
	<!-- Navigation -->
	<nav class="fixed top-0 left-0 right-0 z-50 border-b border-white/5 bg-surface-darker/80 backdrop-blur-xl">
		<div class="mx-auto max-w-6xl px-6 py-4 flex items-center justify-between">
			<a href="/" class="flex items-center gap-3">
				<img src="/images/jon.png" alt="Jon" class="w-8 h-8 rounded-full object-cover ring-2 ring-accent/30" />
				<span class="text-sm font-bold tracking-tight">Learn AI with Jon</span>
			</a>
			<div class="flex items-center gap-5">
				<a href="/contact" class="text-sm text-slate-400 hover:text-white transition-colors">Contact</a>
				<a
					href={hasProgress ? resumeUrl : '/training'}
					class="rounded-full bg-accent px-5 py-2 text-sm font-semibold text-white transition-all hover:bg-accent-hover hover:shadow-lg hover:shadow-accent/20"
				>
					{hasProgress ? 'Continue Learning' : 'Start Free'}
				</a>
			</div>
		</div>
	</nav>

	<!-- Hero -->
	<section class="relative overflow-hidden pt-32 pb-24">
		<!-- Background effects -->
		<div class="absolute inset-0 hero-grid"></div>
		<div class="absolute top-1/4 left-1/4 w-96 h-96 bg-accent/5 rounded-full blur-3xl"></div>
		<div class="absolute bottom-1/4 right-1/4 w-96 h-96 bg-violet-500/5 rounded-full blur-3xl"></div>

		<div class="relative mx-auto max-w-4xl px-6 text-center">
			<div class="inline-flex items-center gap-2 rounded-full border border-white/10 bg-white/5 px-4 py-1.5 text-sm text-slate-300 mb-8 backdrop-blur-sm">
				<span class="w-2 h-2 rounded-full bg-accent animate-pulse-soft"></span>
				Free & self-paced
			</div>

			<h1 class="text-5xl sm:text-7xl font-extrabold tracking-tight leading-[1.1]">
				Master AI.
				<br />
				<span class="text-gradient">Transform your work.</span>
			</h1>

			<p class="mt-6 max-w-2xl mx-auto text-lg sm:text-xl text-slate-400 leading-relaxed">
				A comprehensive, hands-on training for professionals who want to go beyond chatting with AI.
				Learn to think in context, delegate effectively, and build workflows
				that multiply your output.
			</p>

			<div class="mt-10 flex flex-col sm:flex-row items-center justify-center gap-4">
				<a
					href={hasProgress ? resumeUrl : '/training'}
					class="rounded-full bg-accent px-8 py-3.5 text-lg font-semibold text-white transition-all hover:bg-accent-hover hover:shadow-xl hover:shadow-accent/25 hover:-translate-y-0.5"
				>
					{hasProgress ? 'Continue Where You Left Off' : 'Start Training — It\'s Free'}
				</a>
				<a
					href="#modules"
					class="rounded-full border border-white/10 px-6 py-3.5 text-sm font-medium text-slate-300 transition-all hover:bg-white/5 hover:border-white/20"
				>
					View curriculum
				</a>
			</div>

			<!-- Stats -->
			<div class="mt-16 flex items-center justify-center gap-8 sm:gap-16 text-center">
				<div>
					<div class="text-3xl font-bold text-white">10</div>
					<div class="text-sm text-slate-500 mt-1">Modules</div>
				</div>
				<div class="w-px h-10 bg-white/10"></div>
				<div>
					<div class="text-3xl font-bold text-white">40+</div>
					<div class="text-sm text-slate-500 mt-1">Interactive demos</div>
				</div>
				<div class="w-px h-10 bg-white/10"></div>
				<div>
					<div class="text-3xl font-bold text-white">66</div>
					<div class="text-sm text-slate-500 mt-1">Quizzes</div>
				</div>
			</div>
		</div>
	</section>

	<!-- About / Personal -->
	<section class="py-24 border-t border-white/5">
		<div class="mx-auto max-w-4xl px-6">
			<div class="flex flex-col md:flex-row items-center gap-10 md:gap-16">
				<div class="shrink-0">
					<img
						src="/images/jon.png"
						alt="Jon"
						class="w-40 h-40 rounded-2xl object-cover ring-4 ring-white/5 shadow-2xl"
					/>
				</div>
				<div>
					<h2 class="text-2xl sm:text-3xl font-bold tracking-tight mb-4">
						Hi, I'm Jon.
					</h2>
					<p class="text-slate-400 leading-relaxed mb-4">
						I'm a developer, AI experimenter, and teacher. I've spent the last two years deep in the weeds with large language models — building tools, breaking things, and figuring out what actually works in practice.
					</p>
					<p class="text-slate-400 leading-relaxed mb-4">
						This course is everything I wish someone had shown me when I started. Not the hype, not the theory — the real, practical knowledge that transforms how you work with AI every day.
					</p>
					<p class="text-slate-400 leading-relaxed mb-6">
						I built this training to be free because I believe everyone deserves to understand these tools properly. No gatekeeping, no upsell — just honest, thorough teaching from someone who uses AI to build real things.
					</p>

					<!-- Project links -->
					<div class="flex flex-wrap gap-3 mb-6">
						<a href="https://jonathanpincas.com" target="_blank" rel="noopener" class="inline-flex items-center gap-2 rounded-lg bg-white/5 border border-white/10 px-3.5 py-2 text-xs text-slate-300 hover:bg-white/10 hover:border-white/20 transition-all">
							<svg width="14" height="14" viewBox="0 0 16 16" fill="none"><path d="M2 3h12v10H2z" stroke="currentColor" stroke-width="1.5"/><path d="M2 6h12" stroke="currentColor" stroke-width="1.5"/></svg>
							Blog
						</a>
						<a href="https://sitestakk.com" target="_blank" rel="noopener" class="inline-flex items-center gap-2 rounded-lg bg-white/5 border border-white/10 px-3.5 py-2 text-xs text-slate-300 hover:bg-white/10 hover:border-white/20 transition-all">
							<svg width="14" height="14" viewBox="0 0 16 16" fill="none"><path d="M8 1l7 4v6l-7 4-7-4V5l7-4z" stroke="currentColor" stroke-width="1.5" stroke-linejoin="round"/></svg>
							SiteStakk
						</a>
						<a href="https://yagni.co.uk" target="_blank" rel="noopener" class="inline-flex items-center gap-2 rounded-lg bg-white/5 border border-white/10 px-3.5 py-2 text-xs text-slate-300 hover:bg-white/10 hover:border-white/20 transition-all">
							<svg width="14" height="14" viewBox="0 0 16 16" fill="none"><path d="M4 2l4 6v6M12 2L8 8" stroke="currentColor" stroke-width="1.5" stroke-linecap="round" stroke-linejoin="round"/></svg>
							Yagni
						</a>
						<a href="https://agenticresearchlab.substack.com" target="_blank" rel="noopener" class="inline-flex items-center gap-2 rounded-lg bg-white/5 border border-white/10 px-3.5 py-2 text-xs text-slate-300 hover:bg-white/10 hover:border-white/20 transition-all">
							<svg width="14" height="14" viewBox="0 0 16 16" fill="none"><circle cx="8" cy="8" r="6" stroke="currentColor" stroke-width="1.5"/><path d="M8 5v3l2 1.5" stroke="currentColor" stroke-width="1.5" stroke-linecap="round" stroke-linejoin="round"/></svg>
							Agentic Research Lab
						</a>
					</div>

					<!-- Contact CTA -->
					<a
						href="/contact"
						class="inline-flex items-center gap-2 rounded-full bg-white/5 border border-white/10 px-5 py-2.5 text-sm font-medium text-slate-300 hover:bg-white/10 hover:border-white/20 transition-all"
					>
						<svg width="16" height="16" viewBox="0 0 16 16" fill="none"><path d="M2 4l6 4 6-4M2 4v8h12V4H2z" stroke="currentColor" stroke-width="1.5" stroke-linecap="round" stroke-linejoin="round"/></svg>
						Get in touch — consulting, training & speaking
					</a>
				</div>
			</div>
		</div>
	</section>

	<!-- Features grid -->
	<section class="py-24 border-t border-white/5">
		<div class="mx-auto max-w-6xl px-6">
			<div class="text-center mb-16">
				<h2 class="text-3xl sm:text-4xl font-bold tracking-tight">Everything you need to become<br /><span class="text-gradient">AI-fluent</span></h2>
				<p class="mt-4 text-slate-400 max-w-xl mx-auto">Not just tips and tricks. A structured programme that builds real understanding from the ground up.</p>
			</div>

			<div class="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-3 gap-6">
				{#each features as feature}
					<div class="group rounded-2xl border border-white/5 bg-white/[0.02] p-6 transition-all hover:bg-white/[0.04] hover:border-white/10">
						<div class="text-3xl mb-4">{feature.icon}</div>
						<h3 class="text-lg font-semibold text-white mb-2">{feature.title}</h3>
						<p class="text-sm text-slate-400 leading-relaxed">{feature.desc}</p>
					</div>
				{/each}
			</div>
		</div>
	</section>

	<!-- Curriculum / Module list -->
	<section id="modules" class="py-24 border-t border-white/5">
		<div class="mx-auto max-w-4xl px-6">
			<div class="text-center mb-16">
				<h2 class="text-3xl sm:text-4xl font-bold tracking-tight">The curriculum</h2>
				<p class="mt-4 text-slate-400">10 modules. One day. A complete transformation in how you work with AI.</p>
			</div>

			{#if loaded}
				{@const courseState = $course}
				{#if courseState.config}
					<div class="space-y-3">
						{#each courseState.config.modules as moduleName, i}
							{@const meta = courseState.moduleMeta[moduleName]}
							{#if meta}
								<div class="group flex items-center gap-5 rounded-xl border border-white/5 bg-white/[0.02] p-5 transition-all hover:bg-white/[0.04] hover:border-white/10">
									<div class="flex-shrink-0 w-10 h-10 rounded-xl bg-gradient-to-br from-accent/20 to-cyan-400/10 border border-accent/20 flex items-center justify-center text-sm font-bold text-accent">
										{i + 1}
									</div>
									<div class="flex-1 min-w-0">
										<h3 class="font-semibold text-white group-hover:text-accent transition-colors">{meta.title}</h3>
										<p class="mt-0.5 text-sm text-slate-500 leading-relaxed">{meta.description}</p>
									</div>
									{#if meta.estimated_duration}
										<span class="flex-shrink-0 text-xs text-slate-600 font-mono">{meta.estimated_duration}</span>
									{/if}
								</div>
							{/if}
						{/each}
					</div>
				{/if}
			{/if}
		</div>
	</section>

	<!-- CTA -->
	<section class="py-24 border-t border-white/5">
		<div class="mx-auto max-w-3xl px-6 text-center">
			<h2 class="text-3xl sm:text-4xl font-bold tracking-tight">Ready to level up?</h2>
			<p class="mt-4 text-slate-400 text-lg">No signup. No paywall. Just start learning.</p>
			<div class="mt-8">
				<a
					href={hasProgress ? resumeUrl : '/training'}
					class="rounded-full bg-accent px-10 py-4 text-lg font-semibold text-white transition-all hover:bg-accent-hover hover:shadow-xl hover:shadow-accent/25 hover:-translate-y-0.5 inline-block"
				>
					{hasProgress ? 'Continue Learning' : 'Start Training'}
				</a>
			</div>
		</div>
	</section>

	<!-- Footer -->
	<footer class="border-t border-white/5 py-10">
		<div class="mx-auto max-w-6xl px-6 flex items-center justify-between">
			<div class="flex items-center gap-3">
				<img src="/images/jon.png" alt="Jon" class="w-6 h-6 rounded-full object-cover" />
				<span class="text-sm text-slate-600">Learn AI with Jon</span>
			</div>
			<div class="flex items-center gap-6">
				<a href="/contact" class="text-sm text-slate-600 hover:text-slate-400 transition-colors">Contact</a>
				<a href="https://jonathanpincas.com" target="_blank" rel="noopener" class="text-sm text-slate-600 hover:text-slate-400 transition-colors">Blog</a>
				<span class="text-sm text-slate-700">&copy; {new Date().getFullYear()}</span>
			</div>
		</div>
	</footer>
</div>
