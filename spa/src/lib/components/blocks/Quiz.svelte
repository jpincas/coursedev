<script lang="ts">
	import type { QuizBlock } from '$lib/content/types.js';
	import { progress } from '$lib/stores/progress.js';

	let { quiz, moduleName }: { quiz: QuizBlock; moduleName: string } = $props();

	let answered = $state(false);
	let chosenIndex = $state(-1);
	let correct = $state(false);

	// Restore state from progress on mount
	$effect(() => {
		const unsub = progress.subscribe((p) => {
			const mod = p.modules[moduleName];
			const score = mod?.quizScores[quiz.id];
			if (score?.correct) {
				answered = true;
				chosenIndex = quiz.answer;
				correct = true;
			}
		});
		unsub();
	});

	function selectAnswer(index: number) {
		if (answered) return;
		chosenIndex = index;
		correct = index === quiz.answer;
		answered = true;
		progress.setQuizScore(quiz.id, moduleName, correct);
	}

	function retry() {
		answered = false;
		chosenIndex = -1;
		correct = false;
	}
</script>

<div class="rounded-2xl border border-slate-200 bg-white shadow-sm overflow-hidden">
	<!-- Header -->
	<div class="px-6 py-4 bg-slate-50 border-b border-slate-100 flex items-center gap-2">
		<svg width="16" height="16" viewBox="0 0 16 16" fill="none" class="text-accent"><circle cx="8" cy="8" r="6.5" stroke="currentColor" stroke-width="1.5"/><path d="M6 6.5a2 2 0 113.5 1.5c-.5.4-1 .8-1 1.5M8.5 12h-.01" stroke="currentColor" stroke-width="1.5" stroke-linecap="round"/></svg>
		<span class="text-xs font-semibold text-slate-500 uppercase tracking-wider">Check your understanding</span>
	</div>

	<div class="p-6">
		<p class="font-medium text-slate-900 mb-5 text-[15px] leading-relaxed">{quiz.question}</p>

		<div class="space-y-2.5">
			{#each quiz.options as option, i}
				{@const isCorrectAnswer = answered && i === quiz.answer}
				{@const isWrongChoice = answered && i === chosenIndex && !correct}
				{@const isInactive = answered && !isCorrectAnswer && !isWrongChoice}
				<button
					onclick={() => selectAnswer(i)}
					disabled={answered}
					class="w-full text-left px-4 py-3.5 rounded-xl border text-sm transition-all flex items-center gap-3
						{isCorrectAnswer
							? 'border-emerald-300 bg-emerald-50 text-emerald-800'
							: isWrongChoice
								? 'border-red-300 bg-red-50 text-red-800'
								: isInactive
									? 'border-slate-100 text-slate-400 cursor-default'
									: 'border-slate-200 hover:border-accent/40 hover:bg-accent/5 text-slate-700 cursor-pointer'}"
				>
					<span class="w-6 h-6 rounded-full border flex items-center justify-center text-xs font-medium shrink-0
						{isCorrectAnswer
							? 'border-emerald-400 bg-emerald-100 text-emerald-700'
							: isWrongChoice
								? 'border-red-400 bg-red-100 text-red-700'
								: isInactive
									? 'border-slate-200 text-slate-300'
									: 'border-slate-300 text-slate-500'}">
						{#if isCorrectAnswer}
							<svg width="12" height="12" viewBox="0 0 16 16" fill="none"><path d="M3 8.5l3.5 3.5L13 4" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"/></svg>
						{:else if isWrongChoice}
							<svg width="10" height="10" viewBox="0 0 16 16" fill="none"><path d="M4 4l8 8M12 4l-8 8" stroke="currentColor" stroke-width="2" stroke-linecap="round"/></svg>
						{:else}
							{String.fromCharCode(65 + i)}
						{/if}
					</span>
					<span class="leading-relaxed">{option}</span>
				</button>
			{/each}
		</div>

		{#if answered}
			<div class="mt-5 rounded-xl p-4 text-sm leading-relaxed
				{correct ? 'bg-emerald-50 border border-emerald-100' : 'bg-red-50 border border-red-100'}">
				<p class="font-semibold mb-1 {correct ? 'text-emerald-700' : 'text-red-700'}">
					{correct ? 'Correct!' : 'Not quite.'}
				</p>
				{#if quiz.explanation}
					<p class="{correct ? 'text-emerald-600' : 'text-red-600'}">{quiz.explanation}</p>
				{/if}
			</div>
			{#if !correct}
				<button
					onclick={retry}
					class="mt-4 inline-flex items-center gap-1.5 text-sm font-medium text-accent hover:text-accent-hover transition-colors"
				>
					<svg width="14" height="14" viewBox="0 0 16 16" fill="none"><path d="M2 8a6 6 0 1011.5-2.5M13.5 2v3.5H10" stroke="currentColor" stroke-width="1.5" stroke-linecap="round" stroke-linejoin="round"/></svg>
					Try again
				</button>
			{/if}
		{/if}
	</div>
</div>
