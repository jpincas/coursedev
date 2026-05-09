<script lang="ts">
	import type { CalloutBlock } from '$lib/content/types.js';

	let { callout }: { callout: CalloutBlock } = $props();

	const styles: Record<string, { border: string; bg: string; iconBg: string; titleColor: string; textColor: string; icon: string }> = {
		info: {
			border: 'border-blue-200',
			bg: 'bg-blue-50/50',
			iconBg: 'bg-blue-100 text-blue-600',
			titleColor: 'text-blue-900',
			textColor: 'text-blue-700',
			icon: '<svg width="16" height="16" viewBox="0 0 16 16" fill="none"><circle cx="8" cy="8" r="6.5" stroke="currentColor" stroke-width="1.5"/><path d="M8 7v4M8 5h.01" stroke="currentColor" stroke-width="1.5" stroke-linecap="round"/></svg>'
		},
		warning: {
			border: 'border-amber-200',
			bg: 'bg-amber-50/50',
			iconBg: 'bg-amber-100 text-amber-600',
			titleColor: 'text-amber-900',
			textColor: 'text-amber-700',
			icon: '<svg width="16" height="16" viewBox="0 0 16 16" fill="none"><path d="M7.134 2.5a1 1 0 011.732 0l5.196 9A1 1 0 0113.196 13H2.804a1 1 0 01-.866-1.5l5.196-9z" stroke="currentColor" stroke-width="1.5"/><path d="M8 6v3M8 11h.01" stroke="currentColor" stroke-width="1.5" stroke-linecap="round"/></svg>'
		},
		tip: {
			border: 'border-teal-200',
			bg: 'bg-teal-50/50',
			iconBg: 'bg-teal-100 text-teal-600',
			titleColor: 'text-teal-900',
			textColor: 'text-teal-700',
			icon: '<svg width="16" height="16" viewBox="0 0 16 16" fill="none"><path d="M6 14h4M5.5 11c-1.5-1-2.5-2.5-2.5-4.5a5 5 0 1110 0c0 2-1 3.5-2.5 4.5" stroke="currentColor" stroke-width="1.5" stroke-linecap="round"/></svg>'
		},
		danger: {
			border: 'border-red-200',
			bg: 'bg-red-50/50',
			iconBg: 'bg-red-100 text-red-600',
			titleColor: 'text-red-900',
			textColor: 'text-red-700',
			icon: '<svg width="16" height="16" viewBox="0 0 16 16" fill="none"><circle cx="8" cy="8" r="6.5" stroke="currentColor" stroke-width="1.5"/><path d="M5.5 5.5l5 5M10.5 5.5l-5 5" stroke="currentColor" stroke-width="1.5" stroke-linecap="round"/></svg>'
		},
		note: {
			border: 'border-slate-200',
			bg: 'bg-slate-50/50',
			iconBg: 'bg-slate-100 text-slate-600',
			titleColor: 'text-slate-900',
			textColor: 'text-slate-600',
			icon: '<svg width="16" height="16" viewBox="0 0 16 16" fill="none"><path d="M3 2.5h10a.5.5 0 01.5.5v10a.5.5 0 01-.5.5H3a.5.5 0 01-.5-.5V3a.5.5 0 01.5-.5z" stroke="currentColor" stroke-width="1.5"/><path d="M5.5 5.5h5M5.5 8h5M5.5 10.5h3" stroke="currentColor" stroke-width="1.5" stroke-linecap="round"/></svg>'
		}
	};

	let style = $derived(styles[callout.type] || styles.info);
</script>

<div class="rounded-xl border {style.border} {style.bg} p-5 flex gap-4">
	<div class="w-8 h-8 rounded-lg {style.iconBg} flex items-center justify-center shrink-0">
		{@html style.icon}
	</div>
	<div class="flex-1 min-w-0">
		{#if callout.title}
			<h4 class="font-semibold {style.titleColor} mb-1">{callout.title}</h4>
		{/if}
		<p class="text-sm {style.textColor} leading-relaxed">{callout.content}</p>
	</div>
</div>
