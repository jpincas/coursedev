<script lang="ts">
	import type { AgentState } from '$lib/stores/agent.js';
	import { buildFileTree, type FileTreeNode } from '$lib/stores/agent.js';
	import AgentSidebar from './AgentSidebar.svelte';

	let { state, handlers }: { state: AgentState; handlers: Record<string, Function> } = $props();

	let tree = $derived(buildFileTree(state.scratchpad));

	function htmlEscape(s: string): string {
		return s.replace(/&/g, '&amp;').replace(/</g, '&lt;').replace(/>/g, '&gt;');
	}

	function getFileExt(filename: string): string {
		const idx = filename.lastIndexOf('.');
		return idx >= 0 ? filename.substring(idx + 1) : '';
	}

	function fileIconColor(ext: string): string {
		switch (ext) {
			case 'txt': return '#60a5fa';
			case 'csv': return '#4ade80';
			case 'md': return '#c084fc';
			case 'json': return '#facc15';
			case 'yaml': case 'yml': return '#fb923c';
			default: return '#71717a';
		}
	}

	function isNumericCell(s: string): boolean {
		if (!s) return false;
		return /^[+-]?[\d.,]+$/.test(s);
	}

	function highlightMarkdown(content: string): string {
		return content.split('\n').map(line => {
			const trimmed = line.trimStart();
			const esc = htmlEscape(line);
			if (trimmed.startsWith('### ')) return `<span style="color:#7c3aed">### </span><span style="font-weight:500">${htmlEscape(trimmed.slice(4))}</span>`;
			if (trimmed.startsWith('## ')) return `<span style="color:#7c3aed">## </span><span style="font-weight:600">${htmlEscape(trimmed.slice(3))}</span>`;
			if (trimmed.startsWith('# ')) return `<span style="color:#7c3aed"># </span><span style="font-weight:700">${htmlEscape(trimmed.slice(2))}</span>`;
			if (trimmed.startsWith('```')) return `<span style="color:#059669">${esc}</span>`;
			if (trimmed.startsWith('> ')) return `<span style="color:#2563eb">&gt; </span><span style="color:#3b82f6;font-style:italic">${htmlEscape(trimmed.slice(2))}</span>`;
			if (trimmed === '---' || trimmed === '***') return `<span style="color:#a1a1aa">${esc}</span>`;
			if (trimmed.startsWith('- ') || trimmed.startsWith('* ')) {
				const indent = htmlEscape(line.slice(0, line.length - trimmed.length));
				return `${indent}<span style="color:#0d9488">${htmlEscape(trimmed[0] + ' ')}</span>${htmlEscape(trimmed.slice(2))}`;
			}
			return esc;
		}).join('\n');
	}

	function renderCSVTable(content: string): string {
		const lines = content.trim().split('\n');
		if (lines.length === 0) return '<div class="flex-1 flex items-center justify-center"><span class="text-zinc-600 text-sm">Empty file</span></div>';

		const rows = lines.map(l => l.split(','));
		let html = '<div class="flex-1 overflow-auto p-4" style="background:white"><table class="w-full border-collapse">';

		if (rows.length > 0) {
			html += '<thead><tr>';
			for (const cell of rows[0]) {
				html += `<th style="padding:8px 16px;text-align:left;font-size:12px;font-weight:600;color:#52525b;text-transform:uppercase;letter-spacing:0.05em;border-bottom:2px solid #e4e4e7;background:#f4f4f5">${htmlEscape(cell.trim())}</th>`;
			}
			html += '</tr></thead>';
		}

		html += '<tbody>';
		for (let i = 1; i < rows.length; i++) {
			const rowClass = (i - 1) % 2 === 1 ? 'border-b border-zinc-200 bg-zinc-50' : 'border-b border-zinc-200';
			html += `<tr class="${rowClass}">`;
			for (const cell of rows[i]) {
				const text = cell.trim();
				const cls = isNumericCell(text)
					? 'px-4 py-1.5 text-sm text-zinc-900 font-mono text-right tabular-nums'
					: 'px-4 py-1.5 text-sm text-zinc-700 font-mono';
				html += `<td class="${cls}">${htmlEscape(text)}</td>`;
			}
			html += '</tr>';
		}
		html += '</tbody></table></div>';
		return html;
	}

	function renderFileContent(content: string, ext: string): string {
		const trimmed = content.replace(/\n+$/, '');
		if (ext === 'md') {
			return `<pre class="flex-1 overflow-auto m-0 px-4 py-3 font-mono text-sm whitespace-pre-wrap" style="background:white;border:none;color:#1c1917">${highlightMarkdown(trimmed)}</pre>`;
		}
		return `<pre class="flex-1 overflow-auto m-0 px-4 py-3 font-mono text-sm whitespace-pre-wrap" style="background:white;border:none;color:#1c1917">${htmlEscape(trimmed)}</pre>`;
	}

	function flattenTree(nodes: FileTreeNode[], depth: number): Array<{ node: FileTreeNode; depth: number }> {
		const result: Array<{ node: FileTreeNode; depth: number }> = [];
		for (const node of nodes) {
			result.push({ node, depth });
			if (node.isDir && state.expandedFolders[node.fullPath]) {
				result.push(...flattenTree(node.children, depth + 1));
			}
		}
		return result;
	}

	let flatItems = $derived(flattenTree(tree, 0));
</script>

<div class="flex h-screen w-screen overflow-hidden">
	<!-- File workspace (left side) -->
	<div class="flex-1 flex flex-col h-screen overflow-hidden bg-zinc-900">
		<!-- Header -->
		<div class="flex items-center justify-between px-4 py-2 bg-zinc-900 border-b border-zinc-800 shrink-0">
			<div class="flex items-center gap-3">
				<span class="text-[11px] text-zinc-500 uppercase tracking-widest font-semibold">Demo</span>
				<span class="text-zinc-700">/</span>
				<span class="text-sm text-zinc-400">{state.config.title}</span>
			</div>
			<button
				onclick={() => handlers.closeWorkspace()}
				class="py-1 px-3 text-xs rounded bg-zinc-800 text-zinc-400 font-medium cursor-pointer border border-zinc-700 hover:bg-zinc-700 hover:text-zinc-300 transition-colors duration-150"
			>← Back to course</button>
		</div>

		<!-- File explorer -->
		<div class="shrink-0 border-b border-zinc-800 max-h-[50vh] flex flex-col">
			<div class="px-4 py-2 text-[11px] text-zinc-500 font-semibold uppercase tracking-widest shrink-0">
				Explorer · {Object.keys(state.scratchpad).length} files
			</div>
			<div class="pb-2 overflow-y-auto scrollbar-dark">
				{#if flatItems.length === 0}
					<div class="px-4 py-3 text-zinc-600 text-sm italic">No files yet</div>
				{:else}
					{#each flatItems as { node, depth }}
						{#if node.isDir}
							{@const isExpanded = state.expandedFolders[node.fullPath]}
							<div
								class="flex items-center gap-1.5 py-1.5 cursor-pointer transition-colors duration-100 hover:bg-zinc-800/50"
								style="padding-left:{16 + depth * 16}px;padding-right:16px"
								onclick={() => handlers.toggleFolder(node.fullPath)}
								role="button"
								tabindex="0"
								onkeydown={(e) => { if (e.key === 'Enter') handlers.toggleFolder(node.fullPath) }}
							>
								<span class="text-[10px] text-zinc-500 w-3 text-center shrink-0">{isExpanded ? '▾' : '▸'}</span>
								{#if isExpanded}
									<svg width="16" height="16" viewBox="0 0 16 16" fill="none" class="shrink-0"><path d="M1.5 3A1.5 1.5 0 013 1.5h3.293a1 1 0 01.707.293L8.414 3.207a1 1 0 00.707.293H13A1.5 1.5 0 0114.5 5v1H1.5V3z" fill="#fbbf24" fill-opacity="0.3" stroke="#fbbf24" stroke-width="0.8"/><path d="M1 6.5h14l-1.5 8H2.5L1 6.5z" fill="#fbbf24" fill-opacity="0.25" stroke="#fbbf24" stroke-width="0.8"/></svg>
								{:else}
									<svg width="16" height="16" viewBox="0 0 16 16" fill="none" class="shrink-0"><path d="M1.5 3A1.5 1.5 0 013 1.5h3.293a1 1 0 01.707.293L8.414 3.207a1 1 0 00.707.293H13A1.5 1.5 0 0114.5 5v8a1.5 1.5 0 01-1.5 1.5H3A1.5 1.5 0 011.5 13V3z" fill="#fbbf24" fill-opacity="0.25" stroke="#fbbf24" stroke-width="0.8"/></svg>
								{/if}
								<span class="text-sm text-zinc-300 font-mono">{node.name}</span>
							</div>
						{:else}
							{@const isSelected = state.viewingFile === node.fullPath}
							{@const ext = getFileExt(node.name)}
							{@const color = fileIconColor(ext)}
							<div
								class="flex items-center gap-2.5 py-1.5 cursor-pointer transition-colors duration-100
									{isSelected ? 'bg-zinc-800 border-l-2 border-accent' : 'hover:bg-zinc-800/50'}"
								style="padding-left:{(isSelected ? 14 : 16) + depth * 16 + 16}px;padding-right:16px"
								onclick={() => handlers.viewFile(node.fullPath)}
								role="button"
								tabindex="0"
								onkeydown={(e) => { if (e.key === 'Enter') handlers.viewFile(node.fullPath) }}
							>
								<svg width="16" height="16" viewBox="0 0 16 16" fill="none" class="shrink-0"><path d="M3 1.5A.5.5 0 013.5 1H9l4 4v9.5a.5.5 0 01-.5.5h-9a.5.5 0 01-.5-.5v-13z" fill="{color}" fill-opacity="0.15" stroke="{color}" stroke-width="1"/><path d="M9 1v4h4" stroke="{color}" stroke-width="1" fill="none"/></svg>
								<span class="text-sm font-mono {isSelected ? 'text-zinc-200' : 'text-zinc-400'}">{node.name}</span>
							</div>
						{/if}
					{/each}
				{/if}
			</div>
		</div>

		<!-- File preview -->
		{#if state.viewingFile && state.scratchpad[state.viewingFile] !== undefined}
			{@const ext = getFileExt(state.viewingFile)}
			{@const displayName = state.viewingFile.includes('/') ? state.viewingFile.split('/').pop() : state.viewingFile}
			{@const color = fileIconColor(ext)}
			<div class="flex-1 flex flex-col bg-zinc-800/40 overflow-hidden min-h-0 p-5">
				<div class="flex-1 flex flex-col rounded-lg overflow-hidden min-h-0" style="box-shadow: 0 8px 30px rgba(0,0,0,0.4), 0 2px 8px rgba(0,0,0,0.3)">
					<!-- Title bar -->
					<div class="flex items-center px-4 py-2 bg-zinc-200 border-b border-zinc-300 shrink-0">
						<div class="flex items-center gap-2 mr-4">
							<span style="width:12px;height:12px;border-radius:50%;background:#ff5f57;display:inline-block"></span>
							<span style="width:12px;height:12px;border-radius:50%;background:#febc2e;display:inline-block"></span>
							<span style="width:12px;height:12px;border-radius:50%;background:#28c840;display:inline-block"></span>
						</div>
						<div class="flex items-center gap-2">
							<svg width="16" height="16" viewBox="0 0 16 16" fill="none" class="shrink-0"><path d="M3 1.5A.5.5 0 013.5 1H9l4 4v9.5a.5.5 0 01-.5.5h-9a.5.5 0 01-.5-.5v-13z" fill="{color}" fill-opacity="0.15" stroke="{color}" stroke-width="1"/><path d="M9 1v4h4" stroke="{color}" stroke-width="1" fill="none"/></svg>
							<span class="text-sm text-zinc-600 font-mono">{displayName}</span>
							{#if displayName !== state.viewingFile}
								<span class="text-xs text-zinc-400 font-mono ml-1">{state.viewingFile}</span>
							{/if}
						</div>
					</div>
					<!-- Content -->
					{#if ext === 'csv'}
						{@html renderCSVTable(state.scratchpad[state.viewingFile])}
					{:else}
						{@html renderFileContent(state.scratchpad[state.viewingFile], ext)}
					{/if}
				</div>
			</div>
		{:else}
			<div class="flex-1 flex items-center justify-center bg-zinc-800/40">
				<div class="text-center">
					<svg width="48" height="48" viewBox="0 0 16 16" fill="none" class="mx-auto opacity-20"><path d="M3 1.5A.5.5 0 013.5 1H9l4 4v9.5a.5.5 0 01-.5.5h-9a.5.5 0 01-.5-.5v-13z" fill="#71717a" fill-opacity="0.15" stroke="#71717a" stroke-width="1"/><path d="M9 1v4h4" stroke="#71717a" stroke-width="1" fill="none"/></svg>
					<p class="text-zinc-600 text-sm mt-2">Select a file to preview</p>
				</div>
			</div>
		{/if}
	</div>

	<!-- Chat sidebar (right side) -->
	<AgentSidebar {state} {handlers} />
</div>
