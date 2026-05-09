---
title: "Building Websites: Just More Files"
duration: "15m"
tags: [websites, html, output, deployment]
---

# Building Websites: Just More Files

Many people assume building a website requires a developer, a hosting provider, and months of work. That assumption is now wrong. A website is just a collection of text files -- HTML, CSS, and JavaScript -- and AI can create all of them from a conversation.

## What Is a Website, Really?

Strip away the mystique and a website is:

- **HTML files** — the content and structure (like a Word document's text and headings)
- **A CSS file** — the styling (like formatting rules: fonts, colours, layout)
- **A JavaScript file** — the behaviour (like a bit of interactivity: menus opening, buttons responding)

That is it. These are plain text files. You could open them in Notepad. When a browser opens the HTML file, it reads the CSS for how things should look and the JavaScript for how things should behave. The result is a web page.

```callout
type: tip
title: "The Key Insight"
content: "A website is a folder of text files. If AI can write a report, it can write a website. The only difference is the file format."
```

## Why This Matters for You

You do not need to understand HTML, CSS, or JavaScript. You describe what you want in plain English -- just like asking for a report or a spreadsheet -- and AI writes the code files. You can then:

- **Open them locally** — double-click the HTML file and it opens in your browser
- **Share them** — send the folder to someone and they can open it too
- **Publish them** — drag the folder to a free service like Netlify and you have a live website with a real URL, in seconds

This is not a theoretical exercise. Thousands of small businesses, freelancers, and project teams are building real websites this way right now.

## Watch It Happen

In this demo, someone who runs a small coffee shop builds a complete website through conversation. No code knowledge required -- just describing what they want and iterating on the result.

```agent
id: website-build-demo
title: "Building a Website from Scratch"
model_label: "Claude"

system: |
  You are a web developer assistant. You build clean, modern websites
  using plain HTML, CSS, and JavaScript. No frameworks or build tools.
  Create well-structured, accessible pages that work on mobile and desktop.

scratchpad:
  "brief.md": |
    # The Daily Grind — Website Brief

    ## About
    Independent coffee shop in Manchester's Northern Quarter.
    Opened in 2019 by two friends who wanted a place that cared
    as much about the beans as the community. Single-origin beans,
    locally roasted.

    ## What We Need
    A simple website: home page, menu page, about page with our
    story and opening hours. Clean, warm, modern. Not corporate.

    ## Menu
    **Coffee:** Espresso £2.80, Flat White £3.40, Americano £3.00,
    Cappuccino £3.40, Latte £3.40, Pour Over £4.00

    **Food:** Sourdough Toast & Jam £3.50, Granola Bowl £5.50,
    Bacon Roll £4.80, Avocado on Toast £6.00

    ## Details
    - Address: 47 Tib Street, Northern Quarter, Manchester M4 1LS
    - Hours: Mon–Fri 7am–6pm, Sat 8am–6pm, Sun 9am–4pm

tools:
  - list_files
  - scratchpad_read
  - scratchpad_write
  - create_folder

visibility:
  system_prompt: hidden
  tool_calls: visible
  full_context: hidden

sidebar:
  width: "50%"
  start_open: true

script:
  - type: note
    text: "Notice: all the shop details are already in a brief file — not pasted into the chat. The user points the AI at the brief and lets it read the source material. Good practice, just as we've been teaching."

  - type: user
    content: "I've put all the info about my coffee shop in brief.md. Can you read it and build me a simple website? Put everything in a public folder."

  - type: tool_call
    tool: scratchpad_read
    args:
      filename: "brief.md"

  - type: tool_result
    tool: scratchpad_read
    content: |
      # The Daily Grind — Website Brief

      ## About
      Independent coffee shop in Manchester's Northern Quarter.
      Opened in 2019 by two friends who wanted a place that cared
      as much about the beans as the community. Single-origin beans,
      locally roasted.

      ## What We Need
      A simple website: home page, menu page, about page with our
      story and opening hours. Clean, warm, modern. Not corporate.

      ## Menu
      **Coffee:** Espresso £2.80, Flat White £3.40, Americano £3.00,
      Cappuccino £3.40, Latte £3.40, Pour Over £4.00

      **Food:** Sourdough Toast & Jam £3.50, Granola Bowl £5.50,
      Bacon Roll £4.80, Avocado on Toast £6.00

      ## Details
      - Address: 47 Tib Street, Northern Quarter, Manchester M4 1LS
      - Hours: Mon–Fri 7am–6pm, Sat 8am–6pm, Sun 9am–4pm

  - type: assistant
    content: |
      Great brief. I'll build a 3-page site with a shared design:

      **Structure:**
      - `public/index.html` — home page with hero banner and welcome
      - `public/menu.html` — full menu with prices
      - `public/about.html` — your story, opening hours, location
      - `public/main.css` — warm brown/cream styling, mobile-friendly
      - `public/main.js` — mobile navigation toggle

      **Design direction:** Warm browns and creams, serif typography, feels independent not corporate.

      Let me build it.

  - type: tool_call
    tool: create_folder
    args:
      path: "public"

  - type: tool_result
    tool: create_folder
    content: "Created folder: public"

  - type: tool_call
    tool: scratchpad_write
    args:
      filename: "public/index.html"
      content: |
        <!DOCTYPE html>
        <html lang="en">
        <head>
            <meta charset="UTF-8">
            <meta name="viewport" content="width=device-width, initial-scale=1.0">
            <title>The Daily Grind — Manchester Coffee Shop</title>
            <link rel="stylesheet" href="main.css">
        </head>
        <body>
            <nav>
                <div class="brand">The Daily Grind</div>
                <button class="nav-toggle" aria-label="Menu">&#9776;</button>
                <ul class="nav-links">
                    <li><a href="index.html" class="active">Home</a></li>
                    <li><a href="menu.html">Menu</a></li>
                    <li><a href="about.html">About</a></li>
                </ul>
            </nav>
            <section class="hero">
                <h1>The Daily Grind</h1>
                <p>Independent coffee, brewed with care, in the Northern Quarter.</p>
                <a href="menu.html" class="btn">View Our Menu</a>
            </section>
            <section class="intro">
                <h2>Coffee With Intention</h2>
                <p>We source single-origin beans from small farms and roast
                   them locally. Every cup is made with care — not just
                   caffeine, but craft.</p>
                <p>Drop in, grab a seat, and stay a while.</p>
            </section>
            <footer>
                <p>47 Tib Street, Northern Quarter, Manchester M4 1LS</p>
                <p>&copy; 2026 The Daily Grind</p>
            </footer>
            <script src="main.js"></script>
        </body>
        </html>

  - type: tool_result
    tool: scratchpad_write
    content: "Written to public/index.html (36 lines)"

  - type: tool_call
    tool: scratchpad_write
    args:
      filename: "public/menu.html"
      content: |
        <!DOCTYPE html>
        <html lang="en">
        <head>
            <meta charset="UTF-8">
            <meta name="viewport" content="width=device-width, initial-scale=1.0">
            <title>Menu — The Daily Grind</title>
            <link rel="stylesheet" href="main.css">
        </head>
        <body>
            <nav>
                <div class="brand">The Daily Grind</div>
                <button class="nav-toggle" aria-label="Menu">&#9776;</button>
                <ul class="nav-links">
                    <li><a href="index.html">Home</a></li>
                    <li><a href="menu.html" class="active">Menu</a></li>
                    <li><a href="about.html">About</a></li>
                </ul>
            </nav>
            <main>
                <h1>Our Menu</h1>
                <section class="menu-category">
                    <h2>Coffee</h2>
                    <div class="item"><span>Espresso</span><span>£2.80</span></div>
                    <div class="item"><span>Flat White</span><span>£3.40</span></div>
                    <div class="item"><span>Americano</span><span>£3.00</span></div>
                    <div class="item"><span>Cappuccino</span><span>£3.40</span></div>
                    <div class="item"><span>Latte</span><span>£3.40</span></div>
                    <div class="item"><span>Pour Over</span><span>£4.00</span></div>
                </section>
                <section class="menu-category">
                    <h2>Food</h2>
                    <div class="item"><span>Sourdough Toast &amp; Jam</span><span>£3.50</span></div>
                    <div class="item"><span>Granola Bowl</span><span>£5.50</span></div>
                    <div class="item"><span>Bacon Roll</span><span>£4.80</span></div>
                    <div class="item"><span>Avocado on Toast</span><span>£6.00</span></div>
                </section>
            </main>
            <footer>
                <p>47 Tib Street, Northern Quarter, Manchester M4 1LS</p>
                <p>&copy; 2026 The Daily Grind</p>
            </footer>
            <script src="main.js"></script>
        </body>
        </html>

  - type: tool_result
    tool: scratchpad_write
    content: "Written to public/menu.html (44 lines)"

  - type: tool_call
    tool: scratchpad_write
    args:
      filename: "public/about.html"
      content: |
        <!DOCTYPE html>
        <html lang="en">
        <head>
            <meta charset="UTF-8">
            <meta name="viewport" content="width=device-width, initial-scale=1.0">
            <title>About — The Daily Grind</title>
            <link rel="stylesheet" href="main.css">
        </head>
        <body>
            <nav>
                <div class="brand">The Daily Grind</div>
                <button class="nav-toggle" aria-label="Menu">&#9776;</button>
                <ul class="nav-links">
                    <li><a href="index.html">Home</a></li>
                    <li><a href="menu.html">Menu</a></li>
                    <li><a href="about.html" class="active">About</a></li>
                </ul>
            </nav>
            <main>
                <h1>Our Story</h1>
                <p>The Daily Grind started in 2019 when two friends decided
                   Manchester needed a coffee shop that cared as much about
                   the beans as the community.</p>
                <p>We source single-origin beans from small farms and roast
                   them locally. Every cup is brewed with intention.</p>
                <h2>Opening Hours</h2>
                <div class="hours">
                    <div class="item"><span>Monday – Friday</span><span>7:00 – 18:00</span></div>
                    <div class="item"><span>Saturday</span><span>8:00 – 18:00</span></div>
                    <div class="item"><span>Sunday</span><span>9:00 – 16:00</span></div>
                </div>
                <h2>Find Us</h2>
                <p>47 Tib Street, Northern Quarter, Manchester, M4 1LS</p>
            </main>
            <footer>
                <p>&copy; 2026 The Daily Grind</p>
            </footer>
            <script src="main.js"></script>
        </body>
        </html>

  - type: tool_result
    tool: scratchpad_write
    content: "Written to public/about.html (40 lines)"

  - type: tool_call
    tool: scratchpad_write
    args:
      filename: "public/main.css"
      content: |
        * { margin: 0; padding: 0; box-sizing: border-box; }

        body {
            font-family: Georgia, "Times New Roman", serif;
            color: #2c1810;
            background: #faf6f1;
            line-height: 1.6;
        }

        nav {
            display: flex;
            justify-content: space-between;
            align-items: center;
            padding: 1rem 2rem;
            background: #2c1810;
            color: #faf6f1;
        }
        .brand { font-size: 1.4rem; font-weight: bold; }
        .nav-links {
            display: flex;
            list-style: none;
            gap: 1.5rem;
        }
        .nav-links a {
            color: #faf6f1;
            text-decoration: none;
        }
        .nav-links a.active {
            border-bottom: 2px solid #d4a574;
        }
        .nav-toggle {
            display: none;
            background: none;
            border: none;
            color: #faf6f1;
            font-size: 1.5rem;
            cursor: pointer;
        }

        .hero {
            text-align: center;
            padding: 5rem 2rem;
            background: #2c1810;
            color: #faf6f1;
        }
        .hero h1 { font-size: 2.5rem; color: #d4a574; }
        .hero p { margin-top: 0.5rem; font-size: 1.2rem; opacity: 0.9; }

        .btn {
            display: inline-block;
            margin-top: 1.5rem;
            padding: 0.75rem 2rem;
            background: #d4a574;
            color: #2c1810;
            text-decoration: none;
            border-radius: 4px;
            font-weight: bold;
        }

        main, .intro { max-width: 700px; margin: 2rem auto; padding: 0 1.5rem; }
        h1 { font-size: 2rem; margin-bottom: 1rem; }
        h2 { font-size: 1.3rem; margin: 1.5rem 0 0.5rem; color: #5c3a28; }
        p + p { margin-top: 0.8rem; }

        .menu-category { margin-bottom: 2rem; }
        .item {
            display: flex;
            justify-content: space-between;
            padding: 0.5rem 0;
            border-bottom: 1px solid #e0d5c8;
        }

        footer {
            text-align: center;
            padding: 2rem;
            color: #8c7b6b;
            font-size: 0.9rem;
        }

        @media (max-width: 600px) {
            .nav-links {
                display: none;
                flex-direction: column;
                position: absolute;
                top: 56px;
                right: 0;
                background: #2c1810;
                padding: 1rem 2rem;
            }
            .nav-links.open { display: flex; }
            .nav-toggle { display: block; }
        }

  - type: tool_result
    tool: scratchpad_write
    content: "Written to public/main.css (84 lines)"

  - type: tool_call
    tool: scratchpad_write
    args:
      filename: "public/main.js"
      content: |
        document.querySelector('.nav-toggle')
          .addEventListener('click', function () {
            document.querySelector('.nav-links')
              .classList.toggle('open');
          });

  - type: tool_result
    tool: scratchpad_write
    content: "Written to public/main.js (5 lines)"

  - type: assistant
    content: |
      Your website is ready. Five files in the `public` folder:

      - **index.html** — home page with hero banner and welcome text
      - **menu.html** — full coffee and food menu with prices
      - **about.html** — your story, opening hours, and address
      - **main.css** — warm brown/cream styling, mobile-friendly
      - **main.js** — mobile navigation menu toggle

      Open `index.html` in any browser to preview. Want me to change anything?

  - type: note
    text: "Five text files in a folder. That IS a website. You could open index.html right now in Chrome and navigate between all three pages. Every file is plain text — you could read it in Notepad."

  - type: user
    content: "I love it. Two things: can you add our phone number (0161 234 5678) to the about page under the address? And on the home page, can you make the hero tagline say 'Proper coffee. No pretence.' instead?"

  - type: tool_call
    tool: scratchpad_write
    args:
      filename: "public/index.html"
      content: |
        <!DOCTYPE html>
        <html lang="en">
        <head>
            <meta charset="UTF-8">
            <meta name="viewport" content="width=device-width, initial-scale=1.0">
            <title>The Daily Grind — Manchester Coffee Shop</title>
            <link rel="stylesheet" href="main.css">
        </head>
        <body>
            <nav>
                <div class="brand">The Daily Grind</div>
                <button class="nav-toggle" aria-label="Menu">&#9776;</button>
                <ul class="nav-links">
                    <li><a href="index.html" class="active">Home</a></li>
                    <li><a href="menu.html">Menu</a></li>
                    <li><a href="about.html">About</a></li>
                </ul>
            </nav>
            <section class="hero">
                <h1>The Daily Grind</h1>
                <p>Proper coffee. No pretence.</p>
                <a href="menu.html" class="btn">View Our Menu</a>
            </section>
            <section class="intro">
                <h2>Coffee With Intention</h2>
                <p>We source single-origin beans from small farms and roast
                   them locally. Every cup is made with care — not just
                   caffeine, but craft.</p>
                <p>Drop in, grab a seat, and stay a while.</p>
            </section>
            <footer>
                <p>47 Tib Street, Northern Quarter, Manchester M4 1LS</p>
                <p>&copy; 2026 The Daily Grind</p>
            </footer>
            <script src="main.js"></script>
        </body>
        </html>

  - type: tool_result
    tool: scratchpad_write
    content: "Written to public/index.html (36 lines)"

  - type: tool_call
    tool: scratchpad_write
    args:
      filename: "public/about.html"
      content: |
        <!DOCTYPE html>
        <html lang="en">
        <head>
            <meta charset="UTF-8">
            <meta name="viewport" content="width=device-width, initial-scale=1.0">
            <title>About — The Daily Grind</title>
            <link rel="stylesheet" href="main.css">
        </head>
        <body>
            <nav>
                <div class="brand">The Daily Grind</div>
                <button class="nav-toggle" aria-label="Menu">&#9776;</button>
                <ul class="nav-links">
                    <li><a href="index.html">Home</a></li>
                    <li><a href="menu.html">Menu</a></li>
                    <li><a href="about.html" class="active">About</a></li>
                </ul>
            </nav>
            <main>
                <h1>Our Story</h1>
                <p>The Daily Grind started in 2019 when two friends decided
                   Manchester needed a coffee shop that cared as much about
                   the beans as the community.</p>
                <p>We source single-origin beans from small farms and roast
                   them locally. Every cup is brewed with intention.</p>
                <h2>Opening Hours</h2>
                <div class="hours">
                    <div class="item"><span>Monday – Friday</span><span>7:00 – 18:00</span></div>
                    <div class="item"><span>Saturday</span><span>8:00 – 18:00</span></div>
                    <div class="item"><span>Sunday</span><span>9:00 – 16:00</span></div>
                </div>
                <h2>Find Us</h2>
                <p>47 Tib Street, Northern Quarter, Manchester, M4 1LS</p>
                <p>Phone: 0161 234 5678</p>
            </main>
            <footer>
                <p>&copy; 2026 The Daily Grind</p>
            </footer>
            <script src="main.js"></script>
        </body>
        </html>

  - type: tool_result
    tool: scratchpad_write
    content: "Written to public/about.html (42 lines)"

  - type: assistant
    content: "Done. Updated the hero tagline to 'Proper coffee. No pretence.' and added the phone number under the address on the about page. Everything else is unchanged."

  - type: note
    text: "Iteration works exactly like editing a document. Describe the change in English, AI updates the file. No code knowledge needed."

  - type: note
    text: "To publish this to the internet: go to netlify.com/drop and drag the public folder onto the page. Within seconds you have a live URL you can share. No server, no technical setup, no hosting fees for a basic site. The files you just watched being created ARE the website."
```

## Publishing Is the Easy Part

The files are the hard part -- and AI just did that for you. Publishing is trivial:

1. Go to **netlify.com/drop** (free, no account needed for a basic deploy)
2. Drag your `public` folder onto the page
3. Wait about 10 seconds
4. You have a live website with a real URL

That is genuinely all there is to it. The folder of text files becomes a website that anyone in the world can visit.

For something more permanent, you can connect a custom domain (like `thedailygrind.co.uk`) for about £10/year. But even the free Netlify URL works perfectly for sharing.

```callout
type: info
title: "Other Free Hosting Options"
content: "Netlify is not the only option. GitHub Pages, Vercel, and Cloudflare Pages all offer free hosting for static sites. They all work the same way: upload your folder of files and get a URL."
```

## What AI Can Build

Simple static sites like the coffee shop example are the starting point. AI can also create:

- **Portfolio sites** — showcase your work with project pages and a contact form
- **Event pages** — conference schedules, registration info, speaker bios
- **Documentation sites** — guides, FAQs, knowledge bases
- **Landing pages** — product launches, campaign pages, signup forms
- **Internal dashboards** — data visualisations for your team (no public hosting needed)

The pattern is always the same: describe what you want, AI creates the files, you open them in a browser or publish them.

```callout
type: warning
title: "Know the Limits"
content: "AI-built static sites are perfect for informational pages that change infrequently. They are not suited for user accounts, databases, payment processing, or content that needs to update in real time. For those features, you still need a developer — but a static site covers a surprising number of real-world needs."
```

## The Bigger Point

This page is not really about websites. It is about a mental shift.

A website is just files. A spreadsheet is just a file. A report is just a file. Code is just a file. When you stop thinking of these as different categories of "technical" and "non-technical" work, and start seeing them all as **things AI can write for you**, the range of what you can accomplish expands dramatically.

The person in that demo did not learn HTML. They described a coffee shop and had a conversation. The output happened to be a website instead of a Word document. That is the point.

```quiz
id: website-files-quiz
type: multiple-choice
question: "After AI creates a website as a folder of HTML, CSS, and JS files, what do you need to do to make it a live website that anyone can visit?"
options:
  - "Install a web server, configure DNS, set up SSL certificates, and deploy the code"
  - "Drag the folder to a free service like Netlify Drop and wait a few seconds"
  - "Email the files to a web developer to set up hosting"
  - "Convert the HTML files to a WordPress theme first"
answer: 1
explanation: "Static site hosting services like Netlify Drop let you drag a folder and get a live URL in seconds. No server setup, no developer, no conversion needed. The files AI created ARE the website — they just need somewhere to be served from."
```
