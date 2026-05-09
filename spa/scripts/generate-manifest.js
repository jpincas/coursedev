import { readFileSync, readdirSync, existsSync, writeFileSync, cpSync, rmSync, mkdirSync } from 'fs';
import { join, resolve } from 'path';
import yaml from 'js-yaml';
const parseYaml = yaml.load;

const ROOT = resolve(import.meta.dirname, '..');
const CONTENT_SRC = resolve(ROOT, '..', 'content');
const CONTENT_DEST = join(ROOT, 'static', 'content');
const MANIFEST_PATH = join(ROOT, 'static', 'content-manifest.json');

// Copy content directory to static/
console.log('Copying content to static/content/...');
if (existsSync(CONTENT_DEST)) {
	rmSync(CONTENT_DEST, { recursive: true });
}
mkdirSync(CONTENT_DEST, { recursive: true });
cpSync(CONTENT_SRC, CONTENT_DEST, { recursive: true });

// Build manifest for each locale
const LOCALES = ['en', 'es'];
const manifest = { locales: {}, defaultLocale: 'en' };

for (const locale of LOCALES) {
	const localeContentDest = join(CONTENT_DEST, locale);
	const courseYaml = readFileSync(join(localeContentDest, 'course.yaml'), 'utf-8');
	const course = parseYaml(courseYaml);

	const modules = {};

	for (const moduleName of course.modules) {
		const moduleDir = join(localeContentDest, moduleName);
		if (!existsSync(moduleDir)) {
			console.warn(`Warning: module directory ${moduleName} not found in ${locale}`);
			continue;
		}

		const files = readdirSync(moduleDir);
		const pages = files
			.filter((f) => f.endsWith('.md') && !f.startsWith('_'))
			.sort();

const images = existsSync(join(moduleDir, 'images'))
			? readdirSync(join(moduleDir, 'images')).map((f) => `content/${locale}/${moduleName}/images/${f}`)
			: [];

		// Find any external YAML files (agent demos)
		const externalYaml = files
			.filter((f) => f.endsWith('.yaml') && f !== '_module.yaml')
			.map((f) => `content/${locale}/${moduleName}/${f}`);

		modules[moduleName] = {
			moduleYaml: `content/${locale}/${moduleName}/_module.yaml`,
			pages: pages.map((f) => `content/${locale}/${moduleName}/${f}`),
			images,
			externalYaml
		};
	}

	manifest.locales[locale] = { modules };
}

writeFileSync(MANIFEST_PATH, JSON.stringify(manifest, null, 2));
console.log(`Manifest written to ${MANIFEST_PATH}`);
console.log(`  Locales: ${LOCALES.join(', ')}`);
for (const locale of LOCALES) {
	const modCount = Object.keys(manifest.locales[locale].modules).length;
	const pageCount = Object.values(manifest.locales[locale].modules).reduce((sum, m) => sum + m.pages.length, 0);
	console.log(`  ${locale}: ${modCount} modules, ${pageCount} pages`);
}
