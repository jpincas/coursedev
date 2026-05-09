import { readFileSync, readdirSync, existsSync, writeFileSync, cpSync } from 'fs';
import { join, resolve } from 'path';
import yaml from 'js-yaml';
const parseYaml = yaml.load;

const ROOT = resolve(import.meta.dirname, '..');
const CONTENT_SRC = resolve(ROOT, '..', 'content');
const CONTENT_DEST = join(ROOT, 'static', 'content');
const MANIFEST_PATH = join(ROOT, 'static', 'content-manifest.json');

// Copy content directory to static/
console.log('Copying content to static/content/...');
cpSync(CONTENT_SRC, CONTENT_DEST, { recursive: true });

// Read course.yaml to get module order
const courseYaml = readFileSync(join(CONTENT_DEST, 'course.yaml'), 'utf-8');
const course = parseYaml(courseYaml);

const manifest = {
	modules: {}
};

for (const moduleName of course.modules) {
	const moduleDir = join(CONTENT_DEST, moduleName);
	if (!existsSync(moduleDir)) {
		console.warn(`Warning: module directory ${moduleName} not found`);
		continue;
	}

	const files = readdirSync(moduleDir);
	const pages = files
		.filter((f) => f.endsWith('.md') && !f.startsWith('_'))
		.sort();

	const images = existsSync(join(moduleDir, 'images'))
		? readdirSync(join(moduleDir, 'images')).map((f) => `content/${moduleName}/images/${f}`)
		: [];

	// Find any external YAML files (agent demos)
	const externalYaml = files
		.filter((f) => f.endsWith('.yaml') && f !== '_module.yaml')
		.map((f) => `content/${moduleName}/${f}`);

	manifest.modules[moduleName] = {
		moduleYaml: `content/${moduleName}/_module.yaml`,
		pages: pages.map((f) => `content/${moduleName}/${f}`),
		images,
		externalYaml
	};
}

writeFileSync(MANIFEST_PATH, JSON.stringify(manifest, null, 2));
console.log(`Manifest written to ${MANIFEST_PATH}`);
console.log(`  ${Object.keys(manifest.modules).length} modules`);
const totalPages = Object.values(manifest.modules).reduce((sum, m) => sum + m.pages.length, 0);
console.log(`  ${totalPages} pages total`);
