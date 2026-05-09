<script lang="ts">
	import type { AnnotatedImageBlock } from '$lib/content/types.js';

	let { image }: { image: AnnotatedImageBlock } = $props();
	let activeHotspot = $state(-1);

	function toggleHotspot(index: number) {
		activeHotspot = activeHotspot === index ? -1 : index;
	}
</script>

<div class="relative rounded-xl border border-stone-200 overflow-hidden bg-white">
	<img src={image.src} alt={image.alt} class="w-full" />

	{#each image.hotspots as hotspot, i}
		<button
			onclick={() => toggleHotspot(i)}
			class="absolute w-7 h-7 rounded-full bg-accent text-white text-xs font-bold flex items-center justify-center
				shadow-lg hover:scale-110 transition transform -translate-x-1/2 -translate-y-1/2"
			style="left: {hotspot.x}; top: {hotspot.y};"
		>
			+
		</button>

		{#if activeHotspot === i}
			<div
				class="absolute z-10 w-64 rounded-lg bg-zinc-900 text-white p-3 text-sm shadow-xl -translate-x-1/2"
				style="left: {hotspot.x}; top: calc({hotspot.y} + 20px);"
			>
				<p class="font-semibold mb-1">{hotspot.label}</p>
				<p class="text-zinc-300 text-xs leading-relaxed">{hotspot.detail}</p>
			</div>
		{/if}
	{/each}
</div>
