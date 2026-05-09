<script lang="ts">
	import type { ParsedPage } from '$lib/content/types.js';
	import Quiz from '$lib/components/blocks/Quiz.svelte';
	import Callout from '$lib/components/blocks/Callout.svelte';
	import AnnotatedImage from '$lib/components/blocks/AnnotatedImage.svelte';
	import Exercise from '$lib/components/blocks/Exercise.svelte';
	import AgentDemo from '$lib/components/blocks/AgentDemo.svelte';

	let { page, moduleName }: { page: ParsedPage; moduleName: string } = $props();

	// Split narrative HTML at block placeholders and interleave with block components
	interface Segment {
		type: 'html' | 'block';
		html?: string;
		blockIndex?: number;
	}

	let segments = $derived.by(() => {
		const result: Segment[] = [];
		const html = page.narrativeHTML;
		const placeholderRe = /<div data-block-index="(\d+)"><\/div>/g;
		let lastIndex = 0;
		let match;

		while ((match = placeholderRe.exec(html)) !== null) {
			// Add HTML before this placeholder
			if (match.index > lastIndex) {
				result.push({ type: 'html', html: html.slice(lastIndex, match.index) });
			}
			result.push({ type: 'block', blockIndex: parseInt(match[1]) });
			lastIndex = match.index + match[0].length;
		}

		// Add remaining HTML
		if (lastIndex < html.length) {
			result.push({ type: 'html', html: html.slice(lastIndex) });
		}

		return result;
	});
</script>

<div class="content-prose prose prose-slate prose-lg max-w-none
	prose-headings:tracking-tight prose-headings:text-slate-900
	prose-p:text-slate-600 prose-p:leading-relaxed
	prose-li:text-slate-600
	prose-strong:text-slate-900
	prose-h2:text-2xl prose-h2:font-bold prose-h2:mt-12 prose-h2:mb-4
	prose-h3:text-xl prose-h3:font-semibold">
	{#each segments as segment}
		{#if segment.type === 'html' && segment.html}
			{@html segment.html}
		{:else if segment.type === 'block' && segment.blockIndex !== undefined}
			{@const block = page.blocks[segment.blockIndex]}
			{#if block}
				<div class="not-prose my-8">
					{#if block.blockType === 'quiz'}
						<Quiz quiz={block} {moduleName} />
					{:else if block.blockType === 'callout'}
						<Callout callout={block} />
					{:else if block.blockType === 'annotated-image'}
						<AnnotatedImage image={block} />
					{:else if block.blockType === 'exercise'}
						<Exercise exercise={block} />
					{:else if block.blockType === 'agent'}
						<AgentDemo agent={block} />
					{/if}
				</div>
			{/if}
		{/if}
	{/each}
</div>
