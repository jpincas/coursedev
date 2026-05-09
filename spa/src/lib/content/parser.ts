import { Marked } from 'marked';
import hljs from 'highlight.js';
import yaml from 'js-yaml';
import type { Block, PageMeta, ParsedPage, ScriptEvent } from './types.js';

const FRONTMATTER_RE = /^---\n([\s\S]*?)\n---\n/;
const FENCED_BLOCK_RE = /```(quiz|callout|annotated-image|exercise|agent-demo|agent)\n([\s\S]*?)```/g;

const marked = new Marked({
	gfm: true,
	breaks: false,
	renderer: {
		code(this: unknown, ...args: unknown[]) {
			// marked v12+ passes an object; older versions pass (code, lang, escaped)
			let text: string;
			let lang: string | undefined;
			if (typeof args[0] === 'object' && args[0] !== null) {
				const obj = args[0] as { text: string; lang?: string };
				text = obj.text;
				lang = obj.lang;
			} else {
				text = args[0] as string;
				lang = args[1] as string | undefined;
			}
			if (lang && hljs.getLanguage(lang)) {
				const highlighted = hljs.highlight(text, { language: lang }).value;
				return `<pre><code class="hljs language-${lang}">${highlighted}</code></pre>`;
			}
			const escaped = text.replace(/&/g, '&amp;').replace(/</g, '&lt;').replace(/>/g, '&gt;');
			return `<pre><code>${escaped}</code></pre>`;
		}
	}
});

function extractFrontmatter(markdown: string): { meta: PageMeta; body: string } {
	const match = markdown.match(FRONTMATTER_RE);
	if (match) {
		const meta = yaml.load(match[1]) as PageMeta;
		return { meta, body: markdown.slice(match[0].length) };
	}
	return { meta: { title: '' }, body: markdown };
}

function parseBlock(type: string, yamlContent: string): Block | { blockType: 'agent-path'; path: string } {
	const data = yaml.load(yamlContent) as Record<string, unknown>;

	// Handle external path references for agent blocks
	if ((type === 'agent' || type === 'agent-demo') && data.path && !data.script) {
		return { blockType: 'agent-path' as const, path: data.path as string };
	}

	switch (type) {
		case 'quiz':
			return {
				blockType: 'quiz',
				id: data.id as string || '',
				type: data.type as string || 'multiple-choice',
				question: data.question as string || '',
				options: data.options as string[] || [],
				answer: data.answer as number ?? 0,
				answers: data.answers as number[] | undefined,
				explanation: data.explanation as string || ''
			};
		case 'callout':
			return {
				blockType: 'callout',
				type: data.type as string || 'info',
				title: data.title as string || '',
				content: data.content as string || ''
			};
		case 'annotated-image':
			return {
				blockType: 'annotated-image',
				id: data.id as string || '',
				src: data.src as string || '',
				alt: data.alt as string || '',
				hotspots: (data.hotspots as Array<{ x: string; y: string; label: string; detail: string }>) || []
			};
		case 'exercise':
			return {
				blockType: 'exercise',
				id: data.id as string || '',
				language: data.language as string || '',
				prompt: data.prompt as string || '',
				starter: data.starter as string || '',
				validation: (data.validation as Block & { blockType: 'exercise' })?.validation || { type: 'manual' }
			};
		case 'agent-demo':
		case 'agent':
			return parseAgentBlock(data);
		default:
			throw new Error(`Unknown block type: ${type}`);
	}
}

function parseAgentBlock(data: Record<string, unknown>): Block {
	const visibility = data.visibility as Record<string, string> | undefined;
	const sidebar = data.sidebar as Record<string, unknown> | undefined;
	const rawScript = data.script as Array<Record<string, unknown>> || [];

	return {
		blockType: 'agent',
		id: data.id as string || '',
		title: data.title as string || '',
		modelLabel: (data.model_label || data.modelLabel) as string || 'Claude',
		system: data.system as string || '',
		scratchpad: data.scratchpad as Record<string, string> || {},
		tools: data.tools as string[] || [],
		visibility: {
			systemPrompt: visibility?.system_prompt || visibility?.systemPrompt || 'hidden',
			toolCalls: visibility?.tool_calls || visibility?.toolCalls || 'visible',
			fullContext: visibility?.full_context || visibility?.fullContext || 'hidden',
			tokenCount: visibility?.token_count || visibility?.tokenCount || 'hidden',
			modelName: visibility?.model_name || visibility?.modelName || 'hidden'
		},
		sidebar: {
			width: sidebar?.width as string || '480px',
			startOpen: sidebar?.start_open as boolean ?? sidebar?.startOpen as boolean ?? true
		},
		script: rawScript.map((event) => ({
			type: event.type as ScriptEvent['type'],
			text: event.text as string | undefined,
			content: event.content as string | undefined,
			tokens: event.tokens as number | undefined,
			tool: event.tool as string | undefined,
			args: event.args as Record<string, string> | undefined,
			summary: event.summary as string | undefined,
			resetScratchpad: (event.reset_scratchpad ?? event.resetScratchpad) as boolean | undefined,
			note: event.note as string | undefined
		}))
	};
}

function extractBlocksSync(body: string): { html: string; blocks: Array<Block | { blockType: 'agent-path'; path: string }> } {
	const blocks: Array<Block | { blockType: 'agent-path'; path: string }> = [];
	let blockIndex = 0;

	const withPlaceholders = body.replace(FENCED_BLOCK_RE, (_match, type, yamlContent) => {
		const block = parseBlock(type, yamlContent);
		blocks.push(block);
		const placeholder = `<div data-block-index="${blockIndex}"></div>`;
		blockIndex++;
		return placeholder;
	});

	return { html: withPlaceholders, blocks };
}

async function fetchText(path: string): Promise<string> {
	const cleanPath = path.startsWith('/') ? path.slice(1) : path;
	const res = await fetch(`/${cleanPath}`);
	if (!res.ok) throw new Error(`Failed to fetch ${path}: ${res.status}`);
	return res.text();
}

async function resolveBlocks(blocks: Array<Block | { blockType: 'agent-path'; path: string }>): Promise<Block[]> {
	return Promise.all(
		blocks.map(async (block) => {
			if (block.blockType === 'agent-path') {
				const yamlContent = await fetchText(block.path);
				const data = yaml.load(yamlContent) as Record<string, unknown>;
				return parseAgentBlock(data);
			}
			return block as Block;
		})
	);
}

export async function parsePage(markdown: string): Promise<ParsedPage> {
	const { meta, body } = extractFrontmatter(markdown);
	const { html: markdownWithPlaceholders, blocks: rawBlocks } = extractBlocksSync(body);
	const blocks = await resolveBlocks(rawBlocks);
	const narrativeHTML = await marked.parse(markdownWithPlaceholders);

	// Fix image paths: convert /content/{module}/images/{file} → /content/en/{module}/images/{file}
	// Images are locale-neutral (same diagrams for all languages) and only exist under en/
	const localizedHTML = narrativeHTML.replace(
		/\/content\/(module-\w+)\/images\/([\w.-]+)/gi,
		`/content/en/$1/images/$2`
	);

	return { meta, narrativeHTML: localizedHTML, blocks };
}

export function parseAgentYaml(yamlContent: string): Block {
	const data = yaml.load(yamlContent) as Record<string, unknown>;
	return parseAgentBlock(data);
}
