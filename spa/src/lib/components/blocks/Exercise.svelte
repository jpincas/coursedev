<script lang="ts">
	import type { ExerciseBlock } from '$lib/content/types.js';

	let { exercise }: { exercise: ExerciseBlock } = $props();
</script>

<div class="rounded-xl border border-stone-200 bg-white p-6">
	<div class="flex items-center gap-2 mb-3">
		<span class="text-sm font-semibold text-accent uppercase tracking-wide">Exercise</span>
		{#if exercise.language}
			<span class="rounded bg-stone-100 px-2 py-0.5 text-xs text-stone-500 font-mono">{exercise.language}</span>
		{/if}
	</div>

	<p class="text-sm text-stone-700 mb-4">{exercise.prompt}</p>

	{#if exercise.starter}
		<pre class="rounded-lg bg-stone-50 border border-stone-200 p-4 text-sm font-mono text-stone-800 overflow-x-auto">{exercise.starter}</pre>
	{/if}
</div>
