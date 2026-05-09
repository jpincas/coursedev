// Course structure
export interface CourseManifest {
	modules: Record<string, ModuleManifest>;
}

export interface ModuleManifest {
	moduleYaml: string;
	pages: string[];
	images: string[];
	externalYaml: string[];
}

export interface CourseConfig {
	title: string;
	description: string;
	modules: string[];
}

export interface ModuleMeta {
	title: string;
	description: string;
	prerequisites: string[];
	difficulty: string;
	estimated_duration: string;
	completion: {
		require_all_pages: boolean;
		require_quizzes: boolean;
		min_quiz_score: number;
	};
}

export interface PageMeta {
	title: string;
	duration?: string;
	tags?: string[];
	notes?: string;
}

// Parsed page content
export interface ParsedPage {
	meta: PageMeta;
	narrativeHTML: string;
	blocks: Block[];
}

// Block types
export type Block =
	| QuizBlock
	| CalloutBlock
	| AnnotatedImageBlock
	| ExerciseBlock
	| AgentBlock;

export interface QuizBlock {
	blockType: 'quiz';
	id: string;
	type: string; // multiple-choice, multi-select, true-false, ordering, free-text
	question: string;
	options: string[];
	answer: number;
	answers?: number[];
	explanation: string;
}

export interface CalloutBlock {
	blockType: 'callout';
	type: string; // info, warning, tip, danger, note
	title: string;
	content: string;
}

export interface AnnotatedImageBlock {
	blockType: 'annotated-image';
	id: string;
	src: string;
	alt: string;
	hotspots: Hotspot[];
}

export interface Hotspot {
	x: string;
	y: string;
	label: string;
	detail: string;
}

export interface ExerciseBlock {
	blockType: 'exercise';
	id: string;
	language: string;
	prompt: string;
	starter: string;
	validation: {
		type: string;
		expected?: string;
		keywords?: string[];
	};
}

export interface AgentBlock {
	blockType: 'agent';
	id: string;
	title: string;
	modelLabel: string;
	system: string;
	scratchpad: Record<string, string>;
	tools: string[];
	visibility: AgentVisibility;
	sidebar: AgentSidebarConfig;
	script: ScriptEvent[];
}

export interface AgentVisibility {
	systemPrompt: string;
	toolCalls: string;
	fullContext: string;
	tokenCount: string;
	modelName: string;
}

export interface AgentSidebarConfig {
	width: string;
	startOpen: boolean;
}

export interface ScriptEvent {
	type: 'note' | 'user' | 'assistant' | 'tool_call' | 'tool_result' | 'clear' | 'compaction';
	text?: string;
	content?: string;
	tokens?: number;
	tool?: string;
	args?: Record<string, string>;
	summary?: string;
	resetScratchpad?: boolean;
	note?: string;
}

// Progress tracking
export interface CourseProgress {
	modules: Record<string, ModuleProgress>;
	lastPosition: { module: string; page: string };
	preferences: { fontSize: string };
}

export interface ModuleProgress {
	started: boolean;
	pagesViewed: string[];
	quizScores: Record<string, { correct: boolean; attempts: number }>;
	completedAt: string | null;
}
