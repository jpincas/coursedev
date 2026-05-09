import { writable } from 'svelte/store';
import type { CourseProgress, ModuleProgress } from '$lib/content/types.js';

const STORAGE_KEY = 'ai-training-progress';

function createEmptyProgress(): CourseProgress {
	return {
		modules: {},
		lastPosition: { module: '', page: '' },
		preferences: { fontSize: 'font-medium' }
	};
}

function loadProgress(): CourseProgress {
	if (typeof window === 'undefined') return createEmptyProgress();
	try {
		const stored = localStorage.getItem(STORAGE_KEY);
		if (stored) return JSON.parse(stored) as CourseProgress;
	} catch {
		// ignore
	}
	return createEmptyProgress();
}

function saveProgress(progress: CourseProgress) {
	if (typeof window === 'undefined') return;
	try {
		localStorage.setItem(STORAGE_KEY, JSON.stringify(progress));
	} catch {
		// ignore
	}
}

function createProgressStore() {
	const { subscribe, update, set } = writable<CourseProgress>(loadProgress());

	// Persist on every change
	subscribe((value) => saveProgress(value));

	function ensureModule(progress: CourseProgress, moduleName: string): ModuleProgress {
		if (!progress.modules[moduleName]) {
			progress.modules[moduleName] = {
				started: false,
				pagesViewed: [],
				quizScores: {},
				completedAt: null
			};
		}
		return progress.modules[moduleName];
	}

	return {
		subscribe,
		set,

		markPageViewed(moduleName: string, pageSlug: string) {
			update((p) => {
				const mod = ensureModule(p, moduleName);
				mod.started = true;
				if (!mod.pagesViewed.includes(pageSlug)) {
					mod.pagesViewed.push(pageSlug);
				}
				p.lastPosition = { module: moduleName, page: pageSlug };
				return p;
			});
		},

		setQuizScore(quizId: string, moduleName: string, correct: boolean) {
			update((p) => {
				const mod = ensureModule(p, moduleName);
				const existing = mod.quizScores[quizId];
				mod.quizScores[quizId] = {
					correct: correct || existing?.correct || false,
					attempts: (existing?.attempts || 0) + 1
				};
				return p;
			});
		},

		setFontSize(fontSize: string) {
			update((p) => {
				p.preferences.fontSize = fontSize;
				return p;
			});
		},

		reset() {
			set(createEmptyProgress());
		}
	};
}

export const progress = createProgressStore();
