import type { CourseProgress, ModuleMeta, CourseManifest } from '$lib/content/types.js';
import { getPageSlug } from '$lib/content/loader.js';

/**
 * Check if a module is complete based on its completion criteria.
 * A module is complete when:
 * - All pages have been viewed (if require_all_pages)
 * - All quizzes have been passed (if require_quizzes)
 */
export function isModuleComplete(
	moduleName: string,
	meta: ModuleMeta,
	manifest: CourseManifest,
	progress: CourseProgress,
	quizIdsByModule: Record<string, string[]>
): boolean {
	const modProgress = progress.modules[moduleName];
	if (!modProgress) return false;

	const modManifest = manifest.modules[moduleName];
	if (!modManifest) return false;

	// Check all pages viewed
	if (meta.completion?.require_all_pages) {
		const allPageSlugs = modManifest.pages.map(getPageSlug);
		const allViewed = allPageSlugs.every((slug) => modProgress.pagesViewed.includes(slug));
		if (!allViewed) return false;
	}

	// Check quizzes passed
	if (meta.completion?.require_quizzes) {
		const quizIds = quizIdsByModule[moduleName] || [];
		if (quizIds.length > 0) {
			const passedCount = quizIds.filter((id) => modProgress.quizScores[id]?.correct).length;
			const score = passedCount / quizIds.length;
			if (score < (meta.completion.min_quiz_score || 0)) return false;
		}
	}

	return true;
}

/**
 * Check if a module is unlocked (all prior modules are complete).
 * The first module is always unlocked.
 */
export function isModuleUnlocked(
	moduleName: string,
	moduleOrder: string[],
	moduleMeta: Record<string, ModuleMeta>,
	manifest: CourseManifest,
	progress: CourseProgress,
	quizIdsByModule: Record<string, string[]>
): boolean {
	const idx = moduleOrder.indexOf(moduleName);
	if (idx <= 0) return true; // First module always unlocked

	// All prior modules must be complete
	for (let i = 0; i < idx; i++) {
		const priorModule = moduleOrder[i];
		const priorMeta = moduleMeta[priorModule];
		if (priorMeta && !isModuleComplete(priorModule, priorMeta, manifest, progress, quizIdsByModule)) {
			return false;
		}
	}
	return true;
}

/**
 * Get progress stats for a module.
 */
export function getModuleStats(
	moduleName: string,
	manifest: CourseManifest,
	progress: CourseProgress
): { pagesViewed: number; totalPages: number; percent: number } {
	const modManifest = manifest.modules[moduleName];
	if (!modManifest) return { pagesViewed: 0, totalPages: 0, percent: 0 };

	const totalPages = modManifest.pages.length;
	const modProgress = progress.modules[moduleName];
	const pagesViewed = modProgress?.pagesViewed.length || 0;
	const percent = totalPages > 0 ? Math.round((pagesViewed / totalPages) * 100) : 0;

	return { pagesViewed, totalPages, percent };
}
