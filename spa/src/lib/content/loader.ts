import yaml from 'js-yaml';
import type {
	CourseManifest,
	CourseConfig,
	ModuleMeta,
	ModuleManifest,
	Locale
} from './types.js';

let manifest: CourseManifest | null = null;
let courseConfig: CourseConfig | null = null;
const moduleMetaCache: Record<string, ModuleMeta> = {};
const pageCache: Record<string, string> = {};

async function fetchText(path: string): Promise<string> {
	const res = await fetch(`/${path}`);
	if (!res.ok) throw new Error(`Failed to fetch ${path}: ${res.status}`);
	return res.text();
}

export async function getManifest(): Promise<CourseManifest> {
	if (manifest) return manifest;
	const text = await fetchText('content-manifest.json');
	manifest = JSON.parse(text) as CourseManifest;
	return manifest;
}

export async function getCourseConfig(locale: Locale = 'en'): Promise<CourseConfig> {
	if (courseConfig && getLocale() === locale) return courseConfig;
	const text = await fetchText(`content/${locale}/course.yaml`);
	courseConfig = yaml.load(text) as CourseConfig;
	return courseConfig;
}

export async function getModuleMeta(moduleName: string, locale: Locale = 'en'): Promise<ModuleMeta> {
	const cacheKey = `${locale}:${moduleName}`;
	if (moduleMetaCache[cacheKey]) return moduleMetaCache[cacheKey];
	const m = await getManifest();
	const mod = m.locales[locale]?.modules[moduleName];
	if (!mod) throw new Error(`Module not found in ${locale}: ${moduleName}`);
	const text = await fetchText(mod.moduleYaml);
	const meta = yaml.load(text) as ModuleMeta;
	meta.prerequisites = meta.prerequisites || [];
	moduleMetaCache[cacheKey] = meta;
	return meta;
}

export async function getPageMarkdown(pagePath: string): Promise<string> {
	if (pageCache[pagePath]) return pageCache[pagePath];
	const text = await fetchText(pagePath);
	pageCache[pagePath] = text;
	return text;
}

export function getModuleManifest(moduleName: string, locale: Locale = 'en'): ModuleManifest | undefined {
	return manifest?.locales[locale]?.modules[moduleName];
}

export function getPageSlug(pagePath: string): string {
	// "content/en/module-opening/01-welcome.md" -> "01-welcome"
	const filename = pagePath.split('/').pop() || '';
	return filename.replace('.md', '');
}

export function getPagePath(moduleName: string, pageSlug: string, locale: Locale = 'en'): string | undefined {
	const mod = manifest?.locales[locale]?.modules[moduleName];
	if (!mod) return undefined;
	return mod.pages.find((p) => p.endsWith(`/${pageSlug}.md`));
}

// Current locale (default to 'en')
let currentLocale: Locale = 'en';

export function setLocale(locale: Locale) {
	currentLocale = locale;
	// Reset caches on locale change
	courseConfig = null;
	Object.keys(moduleMetaCache).forEach((k) => {
		if (k.startsWith(`${locale}:`) || !k.startsWith('en:') && !k.startsWith('es:')) return;
		delete moduleMetaCache[k];
	});
	Object.keys(pageCache).forEach((k) => delete pageCache[k]);
}

export function getLocale(): Locale {
	return currentLocale;
}
