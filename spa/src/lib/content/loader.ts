import yaml from 'js-yaml';
import type {
	CourseManifest,
	CourseConfig,
	ModuleMeta,
	ModuleManifest
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

export async function getCourseConfig(): Promise<CourseConfig> {
	if (courseConfig) return courseConfig;
	const text = await fetchText('content/course.yaml');
	courseConfig = yaml.load(text) as CourseConfig;
	return courseConfig;
}

export async function getModuleMeta(moduleName: string): Promise<ModuleMeta> {
	if (moduleMetaCache[moduleName]) return moduleMetaCache[moduleName];
	const m = await getManifest();
	const mod = m.modules[moduleName];
	if (!mod) throw new Error(`Module not found: ${moduleName}`);
	const text = await fetchText(mod.moduleYaml);
	const meta = yaml.load(text) as ModuleMeta;
	meta.prerequisites = meta.prerequisites || [];
	moduleMetaCache[moduleName] = meta;
	return meta;
}

export async function getPageMarkdown(pagePath: string): Promise<string> {
	if (pageCache[pagePath]) return pageCache[pagePath];
	const text = await fetchText(pagePath);
	pageCache[pagePath] = text;
	return text;
}

export function getModuleManifest(moduleName: string): ModuleManifest | undefined {
	return manifest?.modules[moduleName];
}

export function getPageSlug(pagePath: string): string {
	// "content/module-opening/01-welcome.md" -> "01-welcome"
	const filename = pagePath.split('/').pop() || '';
	return filename.replace('.md', '');
}

export function getPagePath(moduleName: string, pageSlug: string): string | undefined {
	const mod = manifest?.modules[moduleName];
	if (!mod) return undefined;
	return mod.pages.find((p) => p.endsWith(`/${pageSlug}.md`));
}
