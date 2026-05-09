import { writable } from 'svelte/store';
import type { CourseConfig, CourseManifest, ModuleMeta, Locale } from '$lib/content/types.js';
import { getCourseConfig, getManifest, getModuleMeta, getPageMarkdown, getLocale } from '$lib/content/loader.js';

export interface CourseState {
	loaded: boolean;
	manifest: CourseManifest | null;
	config: CourseConfig | null;
	moduleMeta: Record<string, ModuleMeta>;
	quizIdsByModule: Record<string, string[]>;
}

// Lightweight regex to extract quiz IDs without full parsing
const QUIZ_ID_RE = /```quiz\n[\s\S]*?id:\s*([^\n]+)[\s\S]*?```/g;

async function extractQuizIds(manifest: CourseManifest, locale: Locale): Promise<Record<string, string[]>> {
	const result: Record<string, string[]> = {};
	const modules = manifest.locales[locale]?.modules;
	if (!modules) return result;

	const entries = Object.entries(modules);
	await Promise.all(
		entries.map(async ([moduleName, mod]) => {
			const ids: string[] = [];
			const pages = await Promise.all(mod.pages.map((p) => getPageMarkdown(p)));
			for (const pageContent of pages) {
				let match;
				QUIZ_ID_RE.lastIndex = 0;
				const re = new RegExp(QUIZ_ID_RE.source, 'g');
				while ((match = re.exec(pageContent)) !== null) {
					ids.push(match[1].trim());
				}
			}
			result[moduleName] = ids;
		})
	);

	return result;
}

function createCourseStore() {
	const { subscribe, set } = writable<CourseState>({
		loaded: false,
		manifest: null,
		config: null,
		moduleMeta: {},
		quizIdsByModule: {}
	});

	return {
		subscribe,

		async load() {
			const locale = getLocale();
			const [manifest, config] = await Promise.all([getManifest(), getCourseConfig(locale)]);

			// Load all module metadata in parallel
			const metaEntries = await Promise.all(
				config.modules.map(async (name) => {
					const meta = await getModuleMeta(name, locale);
					return [name, meta] as const;
				})
			);

			const moduleMeta = Object.fromEntries(metaEntries);

			// Extract quiz IDs from all pages
			const quizIdsByModule = await extractQuizIds(manifest, locale);

			set({
				loaded: true,
				manifest,
				config,
				moduleMeta,
				quizIdsByModule
			});
		}
	};
}

export const course = createCourseStore();
