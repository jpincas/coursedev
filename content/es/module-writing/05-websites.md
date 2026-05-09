---
title: "Construir Sitios Web: Solo Más Archivos"
duration: "15m"
tags: [websites, html, output, deployment]
---

# Construir Sitios Web: Solo Más Archivos

Mucha gente asume que construir un sitio web requiere un desarrollador, un proveedor de alojamiento y meses de trabajo. Esa suposición ahora es incorrecta. Un sitio web es solo una colección de archivos de texto — HTML, CSS y JavaScript — y la IA puede crear todos ellos a partir de una conversación.

## ¿Qué es un Sitio Web, Realmente?

Quitémosle el misticismo y un sitio web es:

- **Archivos HTML** — el contenido y la estructura (como el texto y encabezados de un documento Word)
- **Un archivo CSS** — el estilo (como las reglas de formato: fuentes, colores, diseño)
- **Un archivo JavaScript** — el comportamiento (como un poco de interactividad: menús que se abren, botones que responden)

Eso es todo. Son archivos de texto plano. Podrías abrirlos en el Bloc de Notas. Cuando un navegador abre el archivo HTML, lee el CSS para saber cómo deben verse las cosas y el JavaScript para saber cómo deben comportarse. El resultado es una página web.

```callout
type: tip
title: "La Idea Clave"
content: "Un sitio web es una carpeta de archivos de texto. Si la IA puede escribir un informe, puede escribir un sitio web. La única diferencia es el formato del archivo."
```

## Por Qué Esto Te Importa

No necesitas entender HTML, CSS ni JavaScript. describes lo que quieres en inglés llano — exactamente igual que pedir un informe o una hoja de cálculo — y la IA escribe los archivos de código. Luego puedes:

- **Abrirlos localmente** — doble clic en el archivo HTML y se abre en tu navegador
- **Compartirlos** — enviar la carpeta a alguien y también podrá abrirla
- **Publicarlos** — arrastrar la carpeta a un servicio gratuito como Netlify y tendrás un sitio web vivo con una URL real, en segundos

Esto no es un ejercicio teórico. Miles de pequeñas empresas, freelancers y equipos de proyecto están construyendo sitios web reales de esta forma ahora mismo.

## Observa que Ocurra

En esta demo, alguien que dirige una pequeña cafetería construye un sitio web completo mediante conversación. No se necesita ningún conocimiento de código — solo describir lo que quiere e iterar sobre el resultado.

```agent
id: website-build-demo
title: "Construir un Sitio Web desde Cero"
model_label: "Claude"

system: |
  Eres un asistente desarrollador web. Construyes sitios web limpios y modernos
  usando HTML, CSS y JavaScript plain. Sin frameworks ni herramientas de build.
  Crea páginas bien estructuradas y accesibles que funcionen en móvil y escritorio.

scratchpad:
  "brief.md": |
    # The Daily Grind — Brief del Sitio Web

    ## Sobre Nosotros
    Cafetería independiente en el Northern Quarter de Mánchester.
    Abierta en 2019 por dos amigos que querían un lugar que cuidara
    tanto de los granos como de la comunidad. Granos de origen único,
    tostados localmente.

    ## Qué Necesitamos
    Un sitio web sencillo: página principal, página de menú, página de
    acerca de con nuestra historia y horarios. Limpio, cálido, moderno.
    No corporativo.

    ## Menú
    **Café:** Espresso £2,80, Flat White £3,40, Americano £3,00,
    Cappuccino £3,40, Latte £3,40, Pour Over £4,00

    **Comida:** Sourdough Toast & Jam £3,50, Granola Bowl £5,50,
    Bacon Roll £4,80, Avocado on Toast £6,00

    ## Detalles
    - Dirección: 47 Tib Street, Northern Quarter, Manchester M4 1LS
    - Horarios: Lun-Vie 7:00-18:00, Sáb 8:00-18:00, Dom 9:00-16:00

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
    text: "Fijáos: todos los datos de la cafetería ya están en un archivo brief — no pegados en el chat. El usuario señala a la IA hacia el brief y le deja leer el material fuente. Buena práctica, tal como hemos estado enseñando."

  - type: user
    content: "He puesto toda la información de mi cafetería en brief.md. ¿Puedes leerla y construirme un sitio web sencillo? Poned todo en una carpeta pública."

  - type: tool_call
    tool: scratchpad_read
    args:
      filename: "brief.md"

  - type: tool_result
    tool: scratchpad_read
    content: |
      # The Daily Grind — Brief del Sitio Web

      ## Sobre Nosotros
      Cafetería independiente en el Northern Quarter de Mánchester.
      Abierta en 2019 por dos amigos que querían un lugar que cuidara
      tanto de los granos como de la comunidad. Granos de origen único,
      tostados localmente.

      ## Qué Necesitamos
      Un sitio web sencillo: página principal, página de menú, página de
      acerca de con nuestra historia y horarios. Limpio, cálido, moderno.
      No corporativo.

      ## Menú
      **Café:** Espresso £2,80, Flat White £3,40, Americano £3,00,
      Cappuccino £3,40, Latte £3,40, Pour Over £4,00

      **Comida:** Sourdough Toast & Jam £3,50, Granola Bowl £5,50,
      Bacon Roll £4,80, Avocado on Toast £6,00

      ## Detalles
      - Dirección: 47 Tib Street, Northern Quarter, Manchester M4 1LS
      - Horarios: Lun-Vie 7:00-18:00, Sáb 8:00-18:00, Dom 9:00-16:00

  - type: assistant
    content: |
      Buen brief. Construiré un sitio de 3 páginas con un diseño compartido:

      **Estructura:**
      - `public/index.html` — página principal con banner hero y bienvenida
      - `public/menu.html` — menú completo con precios
      - `public/about.html` — nuestra historia, horarios y ubicación
      - `public/main.css` — estilo cálido marrón/crema, compatible con móvil
      - `public/main.js` — toggle de navegación móvil

      **Dirección del diseño:** Marrones cálidos y cremas, tipografía serif, aspecto independiente no corporativo.

      Voy a construirlo.

  - type: tool_call
    tool: create_folder
    args:
      path: "public"

  - type: tool_result
    tool: create_folder
    content: "Carpeta creada: public"

  - type: tool_call
    tool: scratchpad_write
    args:
      filename: "public/index.html"
      content: |
        <!DOCTYPE html>
        <html lang="es">
        <head>
            <meta charset="UTF-8">
            <meta name="viewport" content="width=device-width, initial-scale=1.0">
            <title>The Daily Grind — Cafetería en Mánchester</title>
            <link rel="stylesheet" href="main.css">
        </head>
        <body>
            <nav>
                <div class="brand">The Daily Grind</div>
                <button class="nav-toggle" aria-label="Menú">&#9776;</button>
                <ul class="nav-links">
                    <li><a href="index.html" class="active">Inicio</a></li>
                    <li><a href="menu.html">Menú</a></li>
                    <li><a href="about.html">Acerca de</a></li>
                </ul>
            </nav>
            <section class="hero">
                <h1>The Daily Grind</h1>
                <p>Café independiente, preparado con cuidado, en el Northern Quarter.</p>
                <a href="menu.html" class="btn">Ver Nuestro Menú</a>
            </section>
            <section class="intro">
                <h2>Café con Intención</h2>
                <p>Obtenemos granos de origen único de pequeñas fincas y los
                   tostamos localmente. Cada taza se prepara con cuidado — no solo
                   cafeína, sino oficio.</p>
                <p>Pasa, coge un asiento, y quédate un rato.</p>
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
    content: "Escrito en public/index.html (36 líneas)"

  - type: tool_call
    tool: scratchpad_write
    args:
      filename: "public/menu.html"
      content: |
        <!DOCTYPE html>
        <html lang="es">
        <head>
            <meta charset="UTF-8">
            <meta name="viewport" content="width=device-width, initial-scale=1.0">
            <title>Menú — The Daily Grind</title>
            <link rel="stylesheet" href="main.css">
        </head>
        <body>
            <nav>
                <div class="brand">The Daily Grind</div>
                <button class="nav-toggle" aria-label="Menú">&#9776;</button>
                <ul class="nav-links">
                    <li><a href="index.html">Inicio</a></li>
                    <li><a href="menu.html" class="active">Menú</a></li>
                    <li><a href="about.html">Acerca de</a></li>
                </ul>
            </nav>
            <main>
                <h1>Nuestro Menú</h1>
                <section class="menu-category">
                    <h2>Café</h2>
                    <div class="item"><span>Espresso</span><span>£2,80</span></div>
                    <div class="item"><span>Flat White</span><span>£3,40</span></div>
                    <div class="item"><span>Americano</span><span>£3,00</span></div>
                    <div class="item"><span>Cappuccino</span><span>£3,40</span></div>
                    <div class="item"><span>Latte</span><span>£3,40</span></div>
                    <div class="item"><span>Pour Over</span><span>£4,00</span></div>
                </section>
                <section class="menu-category">
                    <h2>Comida</h2>
                    <div class="item"><span>Sourdough Toast &amp; Mermelada</span><span>£3,50</span></div>
                    <div class="item"><span>Bowl de Granola</span><span>£5,50</span></div>
                    <div class="item"><span>Bacon Roll</span><span>£4,80</span></div>
                    <div class="item"><span>Aguacate en Tostada</span><span>£6,00</span></div>
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
    content: "Escrito en public/menu.html (44 líneas)"

  - type: tool_call
    tool: scratchpad_write
    args:
      filename: "public/about.html"
      content: |
        <!DOCTYPE html>
        <html lang="es">
        <head>
            <meta charset="UTF-8">
            <meta name="viewport" content="width=device-width, initial-scale=1.0">
            <title>Acerca de — The Daily Grind</title>
            <link rel="stylesheet" href="main.css">
        </head>
        <body>
            <nav>
                <div class="brand">The Daily Grind</div>
                <button class="nav-toggle" aria-label="Menú">&#9776;</button>
                <ul class="nav-links">
                    <li><a href="index.html">Inicio</a></li>
                    <li><a href="menu.html">Menú</a></li>
                    <li><a href="about.html" class="active">Acerca de</a></li>
                </ul>
            </nav>
            <main>
                <h1>Nuestra Historia</h1>
                <p>The Daily Grind empezó en 2019 cuando dos amigos decidieron
                   que Mánchester necesitaba una cafetería que cuidara tanto
                   de los granos como de la comunidad.</p>
                <p>Obtenemos granos de origen único de pequeñas fincas y los
                   tostamos localmente. Cada taza se prepara con intención.</p>
                <h2>Horarios</h2>
                <div class="hours">
                    <div class="item"><span>Lunes – Viernes</span><span>7:00 – 18:00</span></div>
                    <div class="item"><span>Sábado</span><span>8:00 – 18:00</span></div>
                    <div class="item"><span>Domingo</span><span>9:00 – 16:00</span></div>
                </div>
                <h2>Dónde Estamos</h2>
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
    content: "Escrito en public/about.html (40 líneas)"

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
    content: "Escrito en public/main.css (84 líneas)"

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
    content: "Escrito en public/main.js (5 líneas)"

  - type: assistant
    content: |
      Tu sitio web está listo. Cinco archivos en la carpeta `public`:

      - **index.html** — página principal con banner hero y texto de bienvenida
      - **menu.html** — menú completo de café y comida con precios
      - **about.html** — nuestra historia, horarios y dirección
      - **main.css** — estilo cálido marrón/crema, compatible con móvil
      - **main.js** — toggle de menú de navegación móvil

      Abre `index.html` en cualquier navegador para previsualizar. ¿Quieres que cambie algo?

  - type: note
    text: "Cinco archivos de texto en una carpeta. Eso ES un sitio web. Podrías abrir index.html ahora mismo en Chrome y navegar entre las tres páginas. Cada archivo es texto plano — podrías leerlo en el Bloc de Notas."

  - type: user
    content: "Me encanta. Dos cosas: ¿puedes añadir nuestro número de teléfono (0161 234 5678) a la página de acerca de bajo la dirección? Y en la página principal, ¿puedes hacer que el eslogan del hero diga 'Café de verdad. Sin pretensiones.' en lugar del actual?"

  - type: tool_call
    tool: scratchpad_write
    args:
      filename: "public/index.html"
      content: |
        <!DOCTYPE html>
        <html lang="es">
        <head>
            <meta charset="UTF-8">
            <meta name="viewport" content="width=device-width, initial-scale=1.0">
            <title>The Daily Grind — Cafetería en Mánchester</title>
            <link rel="stylesheet" href="main.css">
        </head>
        <body>
            <nav>
                <div class="brand">The Daily Grind</div>
                <button class="nav-toggle" aria-label="Menú">&#9776;</button>
                <ul class="nav-links">
                    <li><a href="index.html" class="active">Inicio</a></li>
                    <li><a href="menu.html">Menú</a></li>
                    <li><a href="about.html">Acerca de</a></li>
                </ul>
            </nav>
            <section class="hero">
                <h1>The Daily Grind</h1>
                <p>Café de verdad. Sin pretensiones.</p>
                <a href="menu.html" class="btn">Ver Nuestro Menú</a>
            </section>
            <section class="intro">
                <h2>Café con Intención</h2>
                <p>Obtenemos granos de origen único de pequeñas fincas y los
                   tostamos localmente. Cada taza se prepara con cuidado — no solo
                   cafeína, sino oficio.</p>
                <p>Pasa, coge un asiento, y quédate un rato.</p>
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
    content: "Escrito en public/index.html (36 líneas)"

  - type: tool_call
    tool: scratchpad_write
    args:
      filename: "public/about.html"
      content: |
        <!DOCTYPE html>
        <html lang="es">
        <head>
            <meta charset="UTF-8">
            <meta name="viewport" content="width=device-width, initial-scale=1.0">
            <title>Acerca de — The Daily Grind</title>
            <link rel="stylesheet" href="main.css">
        </head>
        <body>
            <nav>
                <div class="brand">The Daily Grind</div>
                <button class="nav-toggle" aria-label="Menú">&#9776;</button>
                <ul class="nav-links">
                    <li><a href="index.html">Inicio</a></li>
                    <li><a href="menu.html">Menú</a></li>
                    <li><a href="about.html" class="active">Acerca de</a></li>
                </ul>
            </nav>
            <main>
                <h1>Nuestra Historia</h1>
                <p>The Daily Grind empezó en 2019 cuando dos amigos decidieron
                   que Mánchester necesitaba una cafetería que cuidara tanto
                   de los granos como de la comunidad.</p>
                <p>Obtenemos granos de origen único de pequeñas fincas y los
                   tostamos localmente. Cada taza se prepara con intención.</p>
                <h2>Horarios</h2>
                <div class="hours">
                    <div class="item"><span>Lunes – Viernes</span><span>7:00 – 18:00</span></div>
                    <div class="item"><span>Sábado</span><span>8:00 – 18:00</span></div>
                    <div class="item"><span>Domingo</span><span>9:00 – 16:00</span></div>
                </div>
                <h2>Dónde Estamos</h2>
                <p>47 Tib Street, Northern Quarter, Manchester, M4 1LS</p>
                <p>Teléfono: 0161 234 5678</p>
            </main>
            <footer>
                <p>&copy; 2026 The Daily Grind</p>
            </footer>
            <script src="main.js"></script>
        </body>
        </html>

  - type: tool_result
    tool: scratchpad_write
    content: "Escrito en public/about.html (42 líneas)"

  - type: assistant
    content: "Listo. Actualicé el eslogan del hero a 'Café de verdad. Sin pretensiones.' y añadí el número de teléfono bajo la dirección en la página de acerca de. Todo lo demás sin cambios."

  - type: note
    text: "La iteración funciona exactamente igual que editar un documento. Describe el cambio en inglés, la IA actualiza el archivo. No se necesita conocimiento de código."

  - type: note
    text: "Para publicar esto en internet: ve a netlify.com/drop y arrastra la carpeta public sobre la página. En unos segundos tienes una URL viva que puedes compartir. Sin servidor, sin configuración técnica, sin costes de alojamiento para un sitio básico. Los archivos que acabas de ver crear SON el sitio web."
```

## Publicar es la Parte Fácil

Los archivos son la parte difícil — y la IA acaba de hacerlos por ti. Publicar es trivial:

1. Ve a **netlify.com/drop** (gratuito, no cuenta necesaria para un despliegue básico)
2. Arrastra tu carpeta `public` sobre la página
3. Espera unos 10 segundos
4. Tienes un sitio web vivo con una URL real

Esto es genuinamente todo lo que hay. La carpeta de archivos de texto se convierte en un sitio web que cualquier persona en el mundo puede visitar.

Para algo más permanente, puedes conectar un dominio personalizado (como `thedailygrind.co.uk`) por unos £10/año. Pero incluso la URL gratuita de Netlify funciona perfectamente para compartir.

```callout
type: info
title: "Otras Opciones de Hosting Gratuito"
content: "Netlify no es la única opción. GitHub Pages, Vercel y Cloudflare Pages ofrecen todos alojamiento gratuito para sitios estáticos. Funcionan todos igual: sube tu carpeta de archivos y obtén una URL."
```

## Qué Puede Construir la IA

Los sitios estáticos sencillos como el ejemplo de la cafetería son el punto de partida. La IA también puede crear:

- **Sitios de portafolio** — muestra tu trabajo con páginas de proyectos y un formulario de contacto
- **Páginas de eventos** — programas de conferencias, información de registro, biografías de ponentes
- **Sitios de documentación** — guías, FAQs, bases de conocimiento
- **Landing pages** — lanzamientos de productos, páginas de campañas, formularios de registro
- **Paneles internos** — visualizaciones de datos para tu equipo (no se necesita hosting público)

El patrón es siempre el mismo: describe lo que quieres, la IA crea los archivos, los abres en un navegador o los publicas.

```callout
type: warning
title: "Conoce los Límites"
content: "Los sitios estáticos creados por IA son perfectos para páginas informativas que cambian poco frecuentemente. No están diseñados para cuentas de usuario, bases de datos, procesamiento de pagos o contenido que necesite actualizarse en tiempo real. Para esas funcionalidades aún necesitas un desarrollador — pero un sitio estático cubre un número sorprendente de necesidades del mundo real."
```

## El Punto Más Grande

Esta página no trata realmente sobre sitios web. Trata sobre un cambio mental.

Un sitio web son solo archivos. Una hoja de cálculo es solo un archivo. Un informe es solo un archivo. El código es solo un archivo. Cuando dejas de pensar en esto como diferentes categorías de trabajo "técnico" y "no técnico", y empiezas a verlos todos como **cosas que la IA puede escribir por ti**, el alcance de lo que puedes lograr se expande dramáticamente.

La persona de esa demo no aprendió HTML. Describió una cafetería y tuvo una conversación. El resultado happenó ser un sitio web en lugar de un documento Word. Ese es el punto.

```quiz
id: website-files-quiz
type: multiple-choice
question: "Después de que la IA cree un sitio web como una carpeta de archivos HTML, CSS y JS, ¿qué necesitas hacer para convertirlo en un sitio web vivo que cualquiera pueda visitar?"
options:
  - "Instalar un servidor web, configurar DNS, configurar certificados SSL y desplegar el código"
  - "Arrastrar la carpeta a un servicio gratuito como Netlify Drop y esperar unos segundos"
  - "Enviar los archivos por correo a un desarrollador web para que configure el alojamiento"
  - "Convertir los archivos HTML a un tema de WordPress primero"
answer: 1
explanation: "Los servicios de hosting de sitios estáticos como Netlify Drop te permiten arrastrar una carpeta y obtener una URL viva en segundos. Sin configuración de servidor, sin desarrollador, sin conversión necesaria. Los archivos que creó la IA SON el sitio web — solo necesitan un lugar desde donde servirse."
```
