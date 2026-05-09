# Learn AI with Jon — Course Source

Source code and content for [learnai.jonathanpincas.com](https://learnai.jonathanpincas.com) — a free, self-paced AI training course for professionals.

## About the Course

Ten modules taking you from "I've heard of ChatGPT" to building real AI-powered workflows. Topics include how LLMs actually work, context and prompting, working with files, document creation, delegation patterns, and honest coverage of risks and limitations. No coding required.

The course is completely free. If you want to do it, just head to the live site and start learning.

## Using This Repo

### Take the course

Go to [learnai.jonathanpincas.com](https://learnai.jonathanpincas.com). No account, no signup, no cost.

### Run it yourself

```bash
git clone https://github.com/jpincas/coursedev
cd coursedev/spa
npm install
npm run build
cd build && sitestakk dev --port 3001
```

Requires [Node.js](https://nodejs.org) and the [SiteStakk CLI](https://sitestakk.com).

### Fork it, adapt it, translate it

You're welcome to fork this repo and use it as the basis for your own course — whether that's:

- **Translating** the content into another language
- **Adapting** the material for a specific industry or audience
- **Extending** the modules with new content
- **Using the platform** (the SvelteKit SPA) for an entirely different course

The course content lives in `content/` as plain markdown files. The SPA that renders it lives in `spa/`. They're deliberately separate so you can swap out one without touching the other.

If you do build something on top of this, a mention or link back would be appreciated but isn't required.

## Structure

```
content/          Course material — markdown pages, YAML metadata, images
spa/              SvelteKit SPA that renders the course
  src/
    routes/       Pages (landing, training shell, content renderer)
    lib/          Content loading, parsing, stores, components
  static/         Static assets, SiteStakk config
  build/          Production output (deploy from here)
```

## Tech Stack

- **SvelteKit 2** (SPA mode, static output)
- **Svelte 5** with runes
- **Tailwind CSS v4**
- **marked** for client-side markdown rendering
- **SiteStakk** for hosting

## Module Map

| # | Title |
|---|-------|
| 1 | Opening: The February 2026 Moment |
| 2 | How LLMs Actually Work |
| 3 | Context — The Most Important Concept |
| 4 | The Art of Prompting |
| 5 | Files — The Unit of Work |
| 6 | Document Creation and Data Analysis |
| 7 | Advanced Patterns |
| 8 | Delegation & The AI-First Philosophy |
| 9 | Risks, Responsibility, and Realistic Expectations |
| 10 | Putting It Together |

## Contributing

Issues and PRs welcome, particularly for:

- Corrections or improvements to the course content
- Translations (open an issue first to coordinate)
- Bug fixes in the SPA

## Licence

Content and code are both MIT licensed. Do what you like with them.
