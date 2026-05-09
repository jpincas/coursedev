---
title: "Archivos de salida: recibir entregables"
duration: "10m"
tags: [salida, entregables, formatos]
---

# Archivos de salida

Los archivos de salida son los entregables que la IA crea para ti. No texto de chat que copias: archivos reales que puedes descargar, abrir en aplicaciones nativas, editar, compartir y usar.

## Lo que la IA puede crear

![La IA generando diferentes tipos de archivo de salida: documentos, hojas de cálculo, presentaciones, código y visualizaciones](/content/module-files/images/output-file-types.svg)

**Documentos**
- Informes y documentos de posición
- Resúmenes de reuniones
- Documentos de política
- Cartas y correspondencia

**Hojas de cálculo**
- Análisis con fórmulas funcionales
- Modelos de datos
- Presupuestos y previsiones
- Tablas comparativas

**Presentaciones**
- Diapositivas
- Presentaciones para inversores
- Materiales de formación

**Código**
- Scripts y programas
- Archivos de configuración
- Herramientas de automatización

**Visualizaciones**
- Gráficos y diagramas
- Diagramas
- Infografías

## Archivos reales y utilizables

Estos son archivos reales. Cuando la IA crea una hoja de cálculo de Excel:
- Las fórmulas realmente calculan
- Los gráficos se actualizan cuando los datos cambian
- Puedes añadir tus propias modificaciones
- Se abre en Excel como cualquier otra hoja de cálculo

Cuando la IA crea un PowerPoint:
- Tiene diapositivas reales
- Se pueden añadir animaciones y transiciones
- Puedes presentarlo directamente
- O editarlo más en PowerPoint

```callout
type: info
title: "Calidad profesional"
content: "Los archivos generados por la IA están listos para uso profesional. Puedes enviar ese documento a tu jefe, compartir esa hoja de cálculo con tu equipo, presentar esas diapositivas a clientes."
```

## La IA operando en tu sistema de archivos

Más allá de crear archivos individuales, la IA puede trabajar con todo tu sistema de archivos: organizar, mover y limpiar archivos tal como lo haría un asistente humano.

```agent
path: /content/module-files/agent-desktop-cleanup.yaml
```

## El código es solo otra salida

Esto sorprende a los usuarios no técnicos: **el código son simplemente archivos de texto.** HTML, CSS y JavaScript son texto plano — y la IA puede crearlos de forma trivial.

Esto significa que puedes pedir a la IA que cree:
- **Webs sencillas** — páginas de aterrizaje, sitios de portfolio, páginas de proyecto
- **Paneles de datos** — gráficos interactivos y visualizaciones
- **Calculadoras y herramientas** — calculadoras de hipotecas, convertidores de unidades, estimadores de proyecto
- **Presentaciones interactivas** — diapositivas web con animaciones

No necesitas ser programador. Describes lo que quieres y la IA crea los archivos. Los abres en un navegador y funcionan.

**Ejemplo:**
```
Crea una página HTML sencilla que muestre nuestros datos de ventas del T4
como un gráfico de barras interactivo. Incluir los datos en línea.
Que tenga aspecto profesional con un tema oscuro.
```

La IA crea tres archivos (index.html, style.css, script.js), abres index.html en un navegador y tienes un panel interactivo.

```callout
type: tip
title: "El código como formato de documento"
content: "Piensa en el código no como 'programación' sino como otro formato de documento. Así como la IA puede crear un documento de Word o una hoja de cálculo, puede crear una página web. La salida son archivos de texto que sucede que hacen algo cuando se abren en un navegador."
```

## Consideraciones de formato

No todos los formatos de salida son iguales. La IA destaca en algunos y tiene dificultades con otros.

**La IA destaca en:**
- Documentos Markdown (su formato de salida nativo)
- Archivos de datos CSV
- Informes de texto plano
- Contenido web HTML/CSS/JS
- Código fuente en lenguajes populares
- JSON y datos estructurados

**La IA puede producir pero con limitaciones:**
- Hojas de cálculo de Excel (las fórmulas y el formato pueden necesitar ajustes)
- Presentaciones de PowerPoint (estructura básica, puede necesitar pulido de diseño)
- Documentos PDF (normalmente mediante conversión desde otro formato)

**La IA tiene dificultades o no puede producir:**
- Maquetaciones de InDesign
- Diseños de Figma
- Formatos de software especializado (AutoCAD, SPSS, etc.)
- Macros de Excel complejas y VBA

**El consejo práctico:** Pide primero texto plano o formatos estándar. Si necesitas un formato especializado, pide a la IA que cree el contenido en un formato en el que destaca (Markdown, CSV), y luego conviértelo tú mismo usando la herramienta apropiada.

## Cómo pedir archivos

Para recibir archivos en lugar de texto de chat, sé explícito:
- "Crea un archivo Excel con este análisis"
- "Genera un documento Word que contenga..."
- "Construye una presentación PowerPoint sobre..."
- "Guarda esto como un archivo CSV"

No pidas solo "un informe": pide "un documento Word que contenga el informe."

```callout
type: tip
title: "Creación de documentos extensos"
content: "Para informes, propuestas y otros documentos sustanciales, cubriremos el método óptimo — redacción por secciones — en el módulo de Creación de Documentos. Esta técnica mejora drásticamente la calidad para trabajos más largos."
```

## Iteración en archivos

Una vez que tienes el archivo:
1. Revísalo en la aplicación nativa
2. Anota lo que necesita cambiar
3. Ya sea editarlo directamente tú mismo, o
4. Pedir a la IA una versión revisada

Estás trabajando con entregables reales ahora, no con fragmentos de chat.

```quiz
id: output-files-explicit
type: multiple-choice
question: "Necesitas un panel interactivo que muestre tendencias de ventas. No eres programador. ¿Cuál es el enfoque más práctico?"
options:
  - "Pedir a la IA que cree un PowerPoint con gráficos incrustados, ya que es un formato que sabes usar"
  - "Pedir a la IA que cree archivos HTML/CSS/JS para un panel web interactivo, luego abrirlo en tu navegador"
  - "Pedir a la IA que describa cómo construir un panel, luego contratar a un desarrollador para implementarlo"
answer: 1
explanation: "El código son simplemente archivos de texto. La IA puede crear un panel web interactivo completo (HTML, CSS, JavaScript) que abres en un navegador, sin necesidad de conocimientos de programación. Esto produce un resultado más interactivo que los gráficos estáticos de PowerPoint, y no necesitas involucrar a un desarrollador."
```
