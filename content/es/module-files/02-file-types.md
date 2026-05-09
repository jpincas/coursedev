---
title: "Comprender los tipos de archivo"
duration: "15m"
tags: [archivos, tipos, texto-plano, markdown, csv]
---

# Comprender los tipos de archivo

Antes de profundizar en cómo la IA trabaja con archivos, necesitas comprender una distinción fundamental que lo define todo: **texto plano frente a formatos binarios.**

## Texto plano: formatos nativos para la IA

Los archivos de texto plano son exactamente lo que su nombre indica: archivos compuestos de caracteres legibles. Puedes abrirlos en cualquier editor de texto y leer su contenido directamente.

**Formatos comunes de texto plano:**

| Formato | Extensión | Qué es | Por qué importa |
|---------|-----------|--------|-----------------|
| Texto plano | .txt | Texto sencillo sin formato | Universal, sin necesidad de conversión |
| CSV | .csv | Valores separados por comas (datos) | El formato universal de intercambio de datos |
| Markdown | .md | Texto formateado con sintaxis sencilla | La lingua franca de la IA |
| JSON | .json | Datos estructurados | Cómo se comunican las APIs y las configuraciones |
| HTML/CSS/JS | .html, .css, .js | Contenido web | "Código" que es realmente solo texto |
| Código fuente | .py, .go, .java | Lenguajes de programación | Archivos de texto con sintaxis especial |

**Por qué el texto plano importa para la IA:** Los modelos procesan texto de forma nativa. Cuando le das a la IA un archivo de texto plano, lee el contenido directamente: cada palabra, cada dato, cada línea. Sin conversión, sin pérdidas, sin capa de interpretación.

## Formatos binarios: se requiere conversión

Los archivos binarios están codificados de manera que requieren programas específicos para leerlos. No puedes abrir un archivo .docx en un editor de texto y leerlo: verías una serie de caracteres sin sentido.

**Formatos binarios comunes:**
- Documentos de Word (.docx)
- Hojas de cálculo de Excel (.xlsx)
- Presentaciones de PowerPoint (.pptx)
- PDFs (.pdf)
- Imágenes (.png, .jpg)

Cuando le das a la IA un archivo binario, primero debe **convertir** el contenido a texto que pueda procesar. Esta conversión a veces pierde formato, objetos incrustados o información estructural.

```callout
type: info
title: "La implicación práctica"
content: "La IA funciona mejor con texto plano porque lo lee directamente. Los formatos binarios también funcionan, pero siempre hay un paso de conversión que puede perder información. Cuando tienes una opción, los formatos de texto plano son más fiables."
```

## CSV: el formato universal de datos

**CSV (Valores separados por comas)** es el formato de datos más sencillo y portátil. Es simplemente texto con comas que separan columnas y saltos de línea que separan filas.

```
Nombre,Departamento,Salario,Fecha_Inicio
Sarah Chen,Ingeniería,85000,2023-03-15
James Wilson,Marketing,72000,2024-01-10
Maya Patel,Diseño,78000,2023-09-01
```

Cada herramienta de hojas de cálculo puede exportar a CSV. Cada lenguaje de programación puede leer CSV. Cada modelo de IA puede procesar CSV perfectamente.

Al compartir datos con la IA, el CSV es casi siempre la mejor opción: sin formato que perder, sin fórmulas que romper, solo datos limpios.

## Markdown: la lingua franca de la IA

**Markdown** es una forma sencilla de formatear texto usando caracteres sencillos. Es el formato que casi todas las herramientas de IA usan internamente.

**Sintaxis básica:**

```
# Título 1
## Título 2
### Título 3

**Texto en negrita** y *texto en cursiva*

- Punto con viñeta
- Otro punto con viñeta

1. Elemento numerado
2. Segundo elemento

| Columna 1 | Columna 2 |
|-----------|-----------|
| Datos      | Más datos  |

> Cita para énfasis

`código en línea` y bloques de código con triple backtick
```

```callout
type: tip
title: "La inversión de 5 minutos"
content: "Aprender Markdown básico es la inversión de habilidad con mayor retorno que puedes hacer para el trabajo con IA. Cada herramienta importante de IA — Claude, ChatGPT, Gemini, Notion AI — usa Markdown internamente. Cuando escribes en Markdown, la IA comprende tu estructura perfectamente."
```

**Por qué Markdown importa:**
- La IA lee la estructura de Markdown perfectamente (títulos, listas, tablas)
- La IA produce salida en Markdown por defecto
- Es el formato estándar para archivos CLAUDE.md, conocimiento de proyectos e instrucciones persistentes
- Se renderiza bien en la mayoría de herramientas pero sigue siendo legible como texto plano

## Imágenes e IA

Las imágenes son archivos binarios, pero los modelos modernos de IA pueden procesarlos visualmente. Puedes subir capturas de pantalla, diagramas, fotos y gráficos.

Sin embargo, la IA no puede leer el texto incrustado en imágenes con tanta fiabilidad como los archivos de texto reales. Si tienes una opción entre una captura de pantalla de una hoja de cálculo y los datos CSV reales, usa siempre el CSV.

```callout
type: warning
title: "Capturas de pantalla frente a archivos de origen"
content: "Una captura de pantalla de una hoja de cálculo le da a la IA una imagen para interpretar. El CSV real le da a la IA datos perfectos con los que trabajar. Prefiere siempre los archivos de origen sobre las capturas de pantalla cuando los datos son importantes."
```

## Explorar la diferencia

Haz clic en los archivos del explorador. Fíjate cómo los archivos de texto plano (.txt, .csv, .md, .json) son inmediatamente legibles: puedes ver exactamente lo que contienen. Ahora haz clic en el archivo de imagen y en el PowerPoint. Eso es lo que parece un archivo binario cuando intentas leerlo como texto.

```agent
id: file-types-explorer
title: "¿Qué hay dentro de tus archivos?"
model_label: "Claude"

system: |
  Eres un asistente útil que examina archivos y explica
  de qué tipo son y cómo los procesa la IA.

scratchpad:
  "notas-reunion.txt": |
    Reunión de equipo de producto — 3 de febrero
    Asistentes: Sarah, James, Priya, Tom
    Decisiones clave:
    - Lanzar v2.1 antes de fin de mes
    - Contratar dos ingenieros más en el Q2
    - Mover la reunión semanal a los martes
    Acciones:
    - Sarah: redactar brief de contratación antes del viernes
    - James: finalizar lista de verificación de lanzamiento
    - Priya: actualizar roadmap en Notion
  "ventas-trimestrales.csv": |
    Region,Ingresos_T1,Ingresos_T2,Ingresos_T3,Ingresos_T4
    Norte,245000,268000,251000,312000
    Sur,189000,195000,203000,224000
    Este,312000,298000,335000,358000
    Oeste,156000,172000,168000,191000
  "brief-proyecto.md": |
    # Proyecto Phoenix: Rediseño del sitio web

    ## Objetivo
    Rediseñar el sitio web de la empresa para mejorar **las tasas de conversión** y reducir **la tasa de rebote** un 30%.

    ## Cronograma
    - **Fase 1:** Investigación y wireframes (febrero)
    - **Fase 2:** Diseño y prototipado (marzo)
    - **Fase 3:** Construcción y pruebas (abril-mayo)

    ## Requisitos clave
    1. Diseño responsive mobile-first
    2. Tiempos de carga más rápidos (< 2 segundos)
    3. Blog integrado con *optimización SEO*

    > "El sitio actual convierte al 1,2%. El promedio del sector es del 2,8%." — Equipo de marketing
  "config-app.json": |
    {
      "nombre_app": "Phoenix Dashboard",
      "version": "2.1.0",
      "base_datos": {
        "host": "db.internal.company.com",
        "port": 5432,
        "nombre": "phoenix_prod"
      },
      "funciones": {
        "modo_oscuro": true,
        "notificaciones": true,
        "exportar_csv": true
      }
    }
  "foto-equipo.png": |
    âPNG

    IHDR╠É¢wôÜsRGBÇ╠î
    gAMA╠ ╠±ÅüIDATx^ì¢ÿ¦Ö§ÐÑ
    ÒÜàßÞ×ÖÏÎÇÆ¿½¶µ®¬¥¤£¢
    ¡ÿþýüûúùø÷öõôóòñðïîíìëê
    éèçæåäãâáàÞÝÜÛÚÙØ×ÖÕÔÓ
    zxvtsrqponmlkjihgfedcba`_
    ^]\[ZYXWVUTSRQPONMLKJIHGF
    EDCBA@?>=<;:9876543210/.-
    ╠ûIEND®B`é
  "presentacion.pptx": |
    PK!╠Ú[Content_Types].xml
    ¡╠ÎN├0►Eÿ!╠½JÕ*Ç►
    ²ä╠╠·Ø╬╠╠╠╠Î╠ñ╠╠ì
    PPTx/slides/slide1.xml
    ½╠╠½╠╠╠╠╠ÿÉ½╠ñ╠╠ì
    ╠╠╠å╠╠╠╠Î╠╠╠╠╠╠╠╠
    ╠╠¥╠╠é╠╠╠╠Ð╠╠╠╠╠╠
    PK╠╠╠╠╠╠╠╠╠╠╠╠╠╠╠

tools:
  - scratchpad_read
  - list_files

visibility:
  system_prompt: hidden
  tool_calls: visible
  full_context: hidden

sidebar:
  width: "50%"
  start_open: true

script:
  - type: note
    text: "Haz clic en los archivos del explorador de la izquierda. Los cuatro primeros archivos — .txt, .csv, .md, .json — son todos texto plano. Puedes leer cada palabra. Ahora haz clic en 'foto-equipo.png' y 'presentacion.pptx'. Ese enredo es lo que parecen los archivos binarios como datos crudos. La IA tiene la misma experiencia: el texto plano es cristalino, los binarios necesitan conversión."

  - type: user
    content: "Tengo una mezcla de archivos aquí. ¿Puedes leerlos y decirme con cuáles puedes trabajar directamente?"

  - type: tool_call
    tool: list_files
    args: {}

  - type: tool_result
    tool: list_files
    content: |
      6 archivos:
      - config-app.json
      - notas-reunion.txt
      - presentacion.pptx
      - brief-proyecto.md
      - ventas-trimestrales.csv
      - foto-equipo.png

  - type: tool_call
    tool: scratchpad_read
    args:
      filename: "notas-reunion.txt"

  - type: tool_result
    tool: scratchpad_read
    content: |
      Reunión de equipo de producto — 3 de febrero
      Asistentes: Sarah, James, Priya, Tom
      Decisiones clave:
      - Lanzar v2.1 antes de fin de mes
      - Contratar dos ingenieros más en el Q2
      - Mover la reunión semanal a los martes

  - type: tool_call
    tool: scratchpad_read
    args:
      filename: "ventas-trimestrales.csv"

  - type: tool_result
    tool: scratchpad_read
    content: |
      Region,Ingresos_T1,Ingresos_T2,Ingresos_T3,Ingresos_T4
      Norte,245000,268000,251000,312000
      Sur,189000,195000,203000,224000

  - type: tool_call
    tool: scratchpad_read
    args:
      filename: "brief-proyecto.md"

  - type: tool_result
    tool: scratchpad_read
    content: |
      # Proyecto Phoenix: Rediseño del sitio web
      ## Objetivo
      Rediseñar el sitio web de la empresa para mejorar **las tasas de conversión**...

  - type: tool_call
    tool: scratchpad_read
    args:
      filename: "foto-equipo.png"

  - type: tool_result
    tool: scratchpad_read
    content: |
      âPNG
      IHDR╠É¢wôÜsRGBÇ╠î
      gAMA╠ ╠±ÅüIDATx^ì¢ÿ¦Ö§ÐÑ
      ÒÜàßÞ×ÖÏÎÇÆ¿½¶µ®¬¥¤£¢

  - type: assistant
    content: "Esto es lo que encontré:\n\n**Texto plano — los leo perfectamente:**\n- **notas-reunion.txt** — Notas de la reunión de equipo de producto con decisiones y acciones\n- **ventas-trimestrales.csv** — Datos de ingresos por región a lo largo de 4 trimestres\n- **brief-proyecto.md** — Brief de rediseño del sitio web con cronograma y requisitos\n- **config-app.json** — Configuración de la aplicación con base de datos y ajustes de funciones\n\n**Binarios — no puedo leer los datos crudos:**\n- **foto-equipo.png** — Son datos de imagen. Veo que es un archivo PNG por la cabecera, pero el contenido está codificado en binario. Necesitaría procesamiento de imagen para interpretarlo.\n- **presentacion.pptx** — Este es un archivo de PowerPoint (de hecho, un archivo ZIP). Los bytes crudos no se pueden leer como texto.\n\nLos cuatro archivos de texto plano me dan acceso perfecto a su contenido. Los archivos binarios necesitarían ser convertidos o procesados a través de herramientas especializadas antes de que pueda trabajar con ellos de forma significativa."

  - type: note
    text: "Esta es la distinción fundamental. Cuatro archivos que la IA lee perfectamente, dos que no puede. Cada formato de texto — .txt, .csv, .md, .json, .html, .py — le da a la IA acceso directo. Los formatos binarios requieren un paso de conversión adicional que puede perder información."
```

```quiz
id: file-types-plain-text
type: multiple-choice
question: "Necesitas que la IA analice datos financieros de una hoja de cálculo. ¿Qué enfoque le da a la IA un acceso más limpio y fiable a los datos?"
options:
  - "Subir el archivo .xlsx original para que la IA pueda ver el formato y las fórmulas"
  - "Exportar a CSV y subir el archivo .csv para procesamiento de texto directo"
  - "Hacer una captura de pantalla de los datos clave y subir la imagen"
answer: 1
explanation: "CSV es texto plano que la IA lee directamente: cada número, cada columna, perfectamente. Los archivos de Excel requieren conversión (que puede perder estructura), y las capturas de pantalla requieren interpretación visual (que puede malinterpretar números). Para análisis de datos, el CSV es el formato más fiable."
```
