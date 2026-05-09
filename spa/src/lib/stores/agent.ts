import type { AgentBlock, ScriptEvent } from '$lib/content/types.js';

export interface ChatMessage {
	type: 'user' | 'assistant' | 'tool_call' | 'tool_result' | 'compaction_divider' | 'clear_divider';
	content: string;
	toolName: string;
	toolArgs: Record<string, string>;
	groupIdx: number; // -1 = already displayed (no animation)
}

export interface ContextMessage {
	role: 'system' | 'user' | 'assistant' | 'tool_call' | 'tool_result' | 'compaction_summary';
	content: string;
}

export interface AgentState {
	config: AgentBlock;
	scriptIndex: number;
	chatMessages: ChatMessage[];
	contextMessages: ContextMessage[];
	currentNote: string;
	scratchpad: Record<string, string>;
	initialScratchpad: Record<string, string>;
	tokenCount: number;
	workspaceOpen: boolean;
	sidebarOpen: boolean;
	scratchpadOpen: boolean;
	showSystem: boolean;
	showToolCalls: boolean;
	showFullContext: boolean;
	viewingFile: string;
	expandedFolders: Record<string, boolean>;
}

function estimateTokens(text: string): number {
	return Math.ceil(text.length / 4);
}

function autoExpandFolders(state: AgentState, filePath: string) {
	const parts = filePath.split('/');
	for (let i = 1; i < parts.length; i++) {
		state.expandedFolders[parts.slice(0, i).join('/')] = true;
	}
}

export function initAgentState(block: AgentBlock): AgentState {
	const initial: Record<string, string> = { ...block.scratchpad };
	const current: Record<string, string> = { ...block.scratchpad };
	const expandedFolders: Record<string, boolean> = {};

	// Auto-expand all folders in initial scratchpad
	for (const path of Object.keys(current)) {
		const parts = path.split('/');
		for (let i = 1; i < parts.length; i++) {
			expandedFolders[parts.slice(0, i).join('/')] = true;
		}
	}

	const state: AgentState = {
		config: block,
		scriptIndex: 0,
		chatMessages: [],
		contextMessages: [],
		currentNote: '',
		scratchpad: current,
		initialScratchpad: initial,
		tokenCount: 0,
		workspaceOpen: false,
		sidebarOpen: block.sidebar.startOpen,
		scratchpadOpen: false,
		showSystem: false,
		showToolCalls: block.visibility.toolCalls === 'visible',
		showFullContext: false,
		viewingFile: '',
		expandedFolders
	};

	// Seed context with system prompt
	if (block.system) {
		state.contextMessages = [{ role: 'system', content: block.system }];
		state.tokenCount = estimateTokens(block.system);
	}

	return state;
}

function displayToolName(name: string): string {
	const map: Record<string, string> = {
		scratchpad_read: 'read_file',
		scratchpad_write: 'write_file',
		list_files: 'list_files',
		move_file: 'move_file',
		create_folder: 'create_folder',
		delete_file: 'delete_file',
		run_code: 'run_code',
		web_search: 'web_search',
		fetch_url: 'fetch_url'
	};
	return map[name] || name;
}

function appendUserEvent(state: AgentState, event: ScriptEvent, groupIdx: number) {
	state.chatMessages.push({
		type: 'user',
		content: event.content || '',
		toolName: '',
		toolArgs: {},
		groupIdx
	});
	state.contextMessages.push({ role: 'user', content: event.content || '' });
	state.tokenCount += estimateTokens(event.content || '');
}

function appendAssistantEvent(state: AgentState, event: ScriptEvent, groupIdx: number) {
	state.chatMessages.push({
		type: 'assistant',
		content: event.content || '',
		toolName: '',
		toolArgs: {},
		groupIdx
	});
	state.contextMessages.push({ role: 'assistant', content: event.content || '' });
	const tokens = event.tokens || estimateTokens(event.content || '');
	state.tokenCount += tokens;
}

function appendToolCallEvent(state: AgentState, event: ScriptEvent, groupIdx: number) {
	state.chatMessages.push({
		type: 'tool_call',
		content: '',
		toolName: event.tool || '',
		toolArgs: event.args || {},
		groupIdx
	});
	const name = displayToolName(event.tool || '');
	const argsStr = JSON.stringify(event.args || {});
	state.contextMessages.push({ role: 'tool_call', content: `${name}(${argsStr})` });
	state.tokenCount += estimateTokens(`${name}(${argsStr})`);
}

function appendToolResultEvent(state: AgentState, event: ScriptEvent, groupIdx: number) {
	state.chatMessages.push({
		type: 'tool_result',
		content: event.content || '',
		toolName: event.tool || '',
		toolArgs: {},
		groupIdx
	});
	state.contextMessages.push({ role: 'tool_result', content: event.content || '' });
	state.tokenCount += estimateTokens(event.content || '');
}

function findPrecedingCall(script: ScriptEvent[], idx: number, toolName: string): Record<string, string> | null {
	for (let i = idx - 1; i >= 0; i--) {
		if (script[i].type === 'tool_call' && script[i].tool === toolName) {
			return script[i].args || null;
		}
	}
	return null;
}

function applyToolSideEffects(state: AgentState, script: ScriptEvent[], idx: number) {
	const event = script[idx];

	switch (event.tool) {
		case 'scratchpad_write': {
			const args = findPrecedingCall(script, idx, 'scratchpad_write');
			if (args) {
				state.scratchpad[args.filename] = args.content;
				autoExpandFolders(state, args.filename);
			}
			break;
		}
		case 'move_file': {
			const args = findPrecedingCall(script, idx, 'move_file');
			if (args && state.scratchpad[args.source] !== undefined) {
				const content = state.scratchpad[args.source];
				delete state.scratchpad[args.source];
				state.scratchpad[args.destination] = content;
				autoExpandFolders(state, args.destination);
			}
			break;
		}
		case 'create_folder': {
			const args = findPrecedingCall(script, idx, 'create_folder');
			if (args) {
				const path = args.path.replace(/\/$/, '');
				state.expandedFolders[path] = true;
				const parts = path.split('/');
				for (let i = 1; i < parts.length; i++) {
					state.expandedFolders[parts.slice(0, i).join('/')] = true;
				}
			}
			break;
		}
		case 'delete_file': {
			const args = findPrecedingCall(script, idx, 'delete_file');
			if (args) {
				delete state.scratchpad[args.filename];
				if (state.viewingFile === args.filename) {
					state.viewingFile = '';
				}
			}
			break;
		}
	}
}

function processEventGroup(state: AgentState, startIdx: number): number {
	const script = state.config.script;
	let idx = startIdx;
	let groupPosition = 0;

	// Reset GroupIdx on existing messages so they don't re-animate
	for (const msg of state.chatMessages) {
		msg.groupIdx = -1;
	}

	while (idx < script.length) {
		const event = script[idx];

		switch (event.type) {
			case 'user':
				if (idx > startIdx) return idx; // Next user = new group
				appendUserEvent(state, event, groupPosition);
				groupPosition++;
				idx++;
				break;
			case 'assistant':
				appendAssistantEvent(state, event, groupPosition);
				groupPosition++;
				idx++;
				break;
			case 'tool_call':
				appendToolCallEvent(state, event, groupPosition);
				groupPosition++;
				idx++;
				break;
			case 'tool_result':
				appendToolResultEvent(state, event, groupPosition);
				applyToolSideEffects(state, script, idx);
				groupPosition++;
				idx++;
				break;
			default:
				return idx; // Boundary: note, compaction, clear
		}
	}

	return idx;
}

function processCompaction(state: AgentState, event: ScriptEvent) {
	const systemMsg = state.contextMessages.length > 0
		? state.contextMessages[0]
		: { role: 'system' as const, content: state.config.system };

	state.contextMessages = [
		systemMsg,
		{ role: 'compaction_summary', content: event.summary || '' }
	];
	state.chatMessages.push({
		type: 'compaction_divider',
		content: event.summary || '',
		toolName: '',
		toolArgs: {},
		groupIdx: -1
	});
	state.tokenCount = estimateTokens(state.config.system) + estimateTokens(event.summary || '');
}

function processClear(state: AgentState, event: ScriptEvent) {
	state.chatMessages = [];
	state.contextMessages = [{ role: 'system', content: state.config.system }];
	state.tokenCount = estimateTokens(state.config.system);

	if (event.note) {
		state.chatMessages.push({
			type: 'clear_divider',
			content: event.note,
			toolName: '',
			toolArgs: {},
			groupIdx: -1
		});
	}

	if (event.resetScratchpad) {
		state.scratchpad = { ...state.initialScratchpad };
		state.expandedFolders = {};
		for (const path of Object.keys(state.initialScratchpad)) {
			const parts = path.split('/');
			for (let i = 1; i < parts.length; i++) {
				state.expandedFolders[parts.slice(0, i).join('/')] = true;
			}
		}
	}
}

export function advance(state: AgentState): AgentState {
	const script = state.config.script;
	if (state.scriptIndex >= script.length) return state;

	const event = script[state.scriptIndex];

	switch (event.type) {
		case 'note':
			state.currentNote = event.text || '';
			state.scriptIndex++;
			break;
		case 'user':
			state.currentNote = '';
			state.scriptIndex = processEventGroup(state, state.scriptIndex);
			break;
		case 'compaction':
			state.currentNote = '';
			processCompaction(state, event);
			state.scriptIndex++;
			break;
		case 'clear':
			state.currentNote = '';
			processClear(state, event);
			state.scriptIndex++;
			break;
		case 'tool_call':
		case 'tool_result':
		case 'assistant':
			state.currentNote = '';
			state.scriptIndex = processEventGroup(state, state.scriptIndex);
			break;
	}

	return state;
}

// File tree helpers

export interface FileTreeNode {
	name: string;
	fullPath: string;
	isDir: boolean;
	children: FileTreeNode[];
}

export function buildFileTree(scratchpad: Record<string, string>): FileTreeNode[] {
	const root: FileTreeNode = { name: '', fullPath: '', isDir: true, children: [] };
	const dirNodes: Record<string, FileTreeNode> = { '': root };

	function getOrCreateDir(path: string): FileTreeNode {
		if (dirNodes[path]) return dirNodes[path];
		const parts = path.split('/');
		let current = root;
		for (let i = 0; i < parts.length; i++) {
			const dirPath = parts.slice(0, i + 1).join('/');
			if (dirNodes[dirPath]) {
				current = dirNodes[dirPath];
				continue;
			}
			const newDir: FileTreeNode = { name: parts[i], fullPath: dirPath, isDir: true, children: [] };
			current.children.push(newDir);
			dirNodes[dirPath] = newDir;
			current = newDir;
		}
		return current;
	}

	const filenames = Object.keys(scratchpad).sort();

	for (const path of filenames) {
		const lastSlash = path.lastIndexOf('/');
		if (lastSlash < 0) {
			root.children.push({ name: path, fullPath: path, isDir: false, children: [] });
		} else {
			const dirPath = path.substring(0, lastSlash);
			const parent = getOrCreateDir(dirPath);
			const fileName = path.substring(lastSlash + 1);
			parent.children.push({ name: fileName, fullPath: path, isDir: false, children: [] });
		}
	}

	// Sort: folders first, then alphabetically
	function sortChildren(node: FileTreeNode) {
		node.children.sort((a, b) => {
			if (a.isDir !== b.isDir) return a.isDir ? -1 : 1;
			return a.name.localeCompare(b.name);
		});
		for (const child of node.children) {
			if (child.isDir) sortChildren(child);
		}
	}
	sortChildren(root);

	return root.children;
}

export function shouldShowPanel(configValue: string, stateToggle: boolean): boolean {
	switch (configValue) {
		case 'visible': return true;
		case 'hidden': return false;
		case 'toggleable': return stateToggle;
		default: return false;
	}
}

export function chatAnimStyle(groupIdx: number): string {
	if (groupIdx < 0) return '';
	const delay = groupIdx * 400;
	return `opacity: 0; animation: agentFadeSlideIn 300ms ease ${delay}ms forwards`;
}

export { displayToolName };
