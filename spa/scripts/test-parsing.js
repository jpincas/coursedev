import { readFileSync, readdirSync } from 'fs';
import { join, resolve } from 'path';
import yaml from 'js-yaml';
import { Marked } from 'marked';

const ROOT = resolve(import.meta.dirname, '..');
const CONTENT = join(ROOT, 'static', 'content');

const FRONTMATTER_RE = /^---\n([\s\S]*?)\n---\n/;
const FENCED_BLOCK_RE = /```(quiz|callout|annotated-image|exercise|agent-demo|agent)\n([\s\S]*?)```/g;

const marked = new Marked({ gfm: true });

const courseYaml = readFileSync(join(CONTENT, 'course.yaml'), 'utf-8');
const course = yaml.load(courseYaml);

let totalPages = 0;
let totalBlocks = 0;
const blockCounts = {};
const errors = [];

for (const moduleName of course.modules) {
	const moduleDir = join(CONTENT, moduleName);
	const files = readdirSync(moduleDir).filter(f => f.endsWith('.md')).sort();

	for (const file of files) {
		totalPages++;
		const path = join(moduleDir, file);
		const content = readFileSync(path, 'utf-8');

		try {
			// Extract frontmatter
			let body = content;
			const fmMatch = content.match(FRONTMATTER_RE);
			if (fmMatch) {
				yaml.load(fmMatch[1]);
				body = content.slice(fmMatch[0].length);
			}

			// Extract blocks
			let blockIndex = 0;
			const withPlaceholders = body.replace(FENCED_BLOCK_RE, (match, type, yamlContent) => {
				totalBlocks++;
				blockCounts[type] = (blockCounts[type] || 0) + 1;
				try {
					yaml.load(yamlContent);
				} catch (e) {
					errors.push(`  YAML error in ${moduleName}/${file} block ${blockIndex} (${type}): ${e.message}`);
				}
				blockIndex++;
				return `<div data-block-index="${blockIndex - 1}"></div>`;
			});

			// Parse markdown
			await marked.parse(withPlaceholders);
		} catch (e) {
			errors.push(`  Parse error in ${moduleName}/${file}: ${e.message}`);
		}
	}
}

console.log(`Parsed ${totalPages} pages, ${totalBlocks} blocks`);
console.log('Block types:', blockCounts);
if (errors.length) {
	console.log(`\n${errors.length} errors:`);
	errors.forEach(e => console.log(e));
} else {
	console.log('All pages parsed successfully!');
}
