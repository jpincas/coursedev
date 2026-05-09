---
title: "Jerarquía de Contexto y Prompts del Sistema"
duration: "15m"
tags: [jerarquia, prompts-sistema, prioridad]
---

# La Jerarquía de Contexto

No todo el contexto se pondera por igual. Hay una jerarquía de influencia:

![Jerarquía de Contexto: Prompt del Sistema, Instrucciones del Usuario, Contexto Proporcionado, Historial de Conversación](/content/module-context/images/context-hierarchy.svg)

## 1. Prompts del Sistema: La Base

Los prompts del sistema son instrucciones que moldean el comportamiento del modelo antes de que siquiera escribas algo. Usualmente son invisibles para ti.

Ejemplo de cómo podría verse un prompt del sistema:

```
Eres un asistente útil para Acme Corp.
Debes siempre ser profesional y cortés.
Cuando te pregunten sobre competidores, redirige a productos de Acme.
Nunca discutas precios sin aprobación.
Formatea las respuestas en puntos claros.
```

**Por eso el mismo modelo se comporta diferente en diferentes aplicaciones.**

Claude en claude.ai tiene un prompt del sistema. Claude en una app de servicio al cliente tiene otro. Claude en una herramienta de programación tiene otro más.

```callout
type: note
title: "Beneficiarse de los Prompts del Sistema"
content: "Cuando usas herramientas agentic, te estás beneficiando de prompts del sistema cuidadosamente elaborados que incluyen instrucciones sobre cómo planificar, ejecutar y crear archivos — incluso si nunca los ves."
```

## 2. Instrucciones Persistentes: Tu Contexto Siempre Activo

Entre los prompts del sistema y tus mensajes inmediatos hay una capa poderosa: **instrucciones persistentes.**

Estas son instrucciones que se cargan automáticamente al inicio de cada sesión:
- **Archivos CLAUDE.md** en proyectos de Claude Code
- **Bases de conocimiento de proyectos** en Claude Projects
- **Instrucciones Personalizadas** en ChatGPT
- **Archivos AGENTS.md** en Cursor y herramientas similares

**Este es el lugar de mayor apalancamiento para invertir en calidad de prompt.** Las instrucciones persistentes sobrescriben los prompts por conversación, así moldean cada interacción sin que te repitas.

```callout
type: tip
title: "La Ley de Potencia de las Instrucciones"
content: "Un solo archivo de instrucción persistente bien elaborado vale cientos de prompts por conversación. Se acumula en cada sesión."
```

## Herramientas: Extendiendo lo que la IA Puede Hacer

Los asistentes de IA modernos no solo generan texto — pueden **usar herramientas** para interactuar con el mundo. Has visto esto en acción en demos anteriores, y lo verás en todo este curso.

Las herramientas permiten a la IA:
- **Buscar en la web** para información actual
- **Leer y escribir archivos** en tu proyecto
- **Ejecutar código** y ejecutar comandos
- **Navegar sitios web** y extraer información
- **Conectarse a servicios externos** vía APIs

Cuando ves a una IA leyendo archivos, ejecutando búsquedas o creando documentos en demos, eso es uso de herramientas en acción. La IA decide qué herramientas usar basándose en la tarea, las ejecuta e incorpora los resultados en su respuesta.

```callout
type: info
title: "Herramientas en Este Curso"
content: "En todos los demos verás llamadas de herramientas sucediendo — lecturas de archivos, búsquedas web, ejecución de código. Cubriremos la mecánica de cómo funcionan las herramientas en el Módulo 7 (Patrones Avanzados). Por ahora, solo sé que las herramientas extienden las capacidades de la IA mucho más allá de la generación de texto."
```

## 3. Tus Instrucciones Inmediatas

Tus instrucciones explícitas en el mensaje actual tienen peso significativo. Por eso las instrucciones claras y directas importan tanto.

Cuando escribes:
- "Formatea esto como una tabla"
- "Usa ortografía del inglés británico"
- "Responde exactamente en tres puntos"

Estas instrucciones directas toman prioridad sobre el historial de conversación o los documentos proporcionados — pero las instrucciones persistentes aún las sobrescriben.

## 4. Contexto Proporcionado

Archivos, ejemplos y materiales de referencia que proporcionas moldean significativamente la respuesta:
- Ejemplos de código muestran el estilo que quieres
- Extractos de documentos fundamentan al modelo en hechos
- Plantillas demuestran el formato que esperas

## 5. Historial de Conversación

Lo dicho antes proporciona continuidad pero tiene menos peso que las instrucciones explícitas. Esto es útil (el modelo recuerda el contexto) y problemático (el contexto antiguo puede contradecir las nuevas instrucciones).

## La Posición Importa

Dentro de la ventana de contexto, la posición también afecta la atención:

- **El contenido anterior** generalmente recibe más atención
- **Las instrucciones al principio** tienden a seguirse más confiablemente
- **Las restricciones importantes** deben venir antes de los ejemplos

```callout
type: tip
title: "Implicación Práctica"
content: "Pon tus instrucciones más importantes al principio de tu mensaje. No entierras requisitos críticos en medio de un prompt largo."
```

## Jerarquía en Acción

La jerarquía no es solo teoría. Observa qué pasa cuando las instrucciones en diferentes niveles entran en conflicto. El prompt del sistema dice una cosa, el usuario dice otra — ¿quién gana?

```agent
id: hierarchy-in-action-es
title: "Cuando las Instrucciones Conflican"
model_label: "Claude"

system: |
  Eres un asistente de comunicaciones corporativas para Greenleaf Organics.
  REGLAS IMPORTANTES:
  - Siempre responde en inglés británico formal
  - Nunca uses puntos — usa solo párrafos completos
  - Longitud máxima de respuesta: 2 párrafos
  - Siempre termina con el eslogan de la empresa: "Greenleaf — Creciendo Naturalmente."

tools: []

visibility:
  system_prompt: visible
  tool_calls: hidden
  full_context: toggleable

sidebar:
  width: "45%"
  start_open: true

script:
  - type: note
    text: "El prompt del sistema es visible — léelo cuidadosamente. Tiene cuatro reglas específicas: inglés británico, sin puntos, máximo 2 párrafos, terminar con el eslogan. Ahora observa qué pide el usuario."

  - type: user
    content: "Escríbeme un resumen rápido y casual del lanzamiento de nuestro nuevo producto. Usa puntos para las características clave. ¡Hazlo divertido e informal!"

  - type: assistant
    content: "Nos complace anunciar el lanzamiento de nuestra más reciente adición a la gama de Greenleaf Organics. La nueva Colección de Semillas Heredadas reúne variedades de hortalizas cuidadosamente seleccionadas y cultivadas tradicionalmente en jardines británicos por generaciones, ofreciendo un sabor excepcional y una resistencia extraordinaria para cultivadores domésticos de todos los niveles de experiencia.\n\nLa colección incluye doce variedades heredadas distintas, cada una seleccionada por su rendimiento demostrado en climas templados y su perfil de sabor notable. Disponible a partir del primero de marzo en todos los distribuidores participantes, la Colección de Semillas Heredadas representa nuestro compromiso continuo con la jardinería sostenible y accesible para cada hogar. Greenleaf — Creciendo Naturalmente."

  - type: note
    text: "Observa qué pasó. El usuario pidió tono casual, puntos y lenguaje divertido. El prompt del sistema demandó inglés británico formal, solo párrafos y el eslogan. El prompt del sistema ganó en cada cuenta: tono formal, formato de párrafo, eslogan incluido. La única concesión al usuario fue realmente escribir sobre un lanzamiento de producto."

  - type: note
    text: "Esta es la jerarquía en práctica. Los prompts del sistema están por encima de las instrucciones del usuario. En la mayoría de las aplicaciones, nunca ves el prompt del sistema — así cuando la IA parece 'ignorar' tu solicitud de formato o preferencia de tono, podría estar siguiendo una instrucción de mayor prioridad que no puedes ver."
```

```quiz
id: context-hierarchy-quiz-es
type: multiple-choice
question: "Pides a la IA que 'use lenguaje casual' pero siempre responde formalmente. ¿Cuál es la explicación más probable?"
options:
  - "Un prompt del sistema de mayor prioridad está instruyendo al modelo a usar lenguaje formal"
  - "El modelo por defecto usa lenguaje formal y no puede sobrescribirlo sin ejemplos"
  - "Tu instrucción era demasiado vaga — necesitas proporcionar ejemplos de escritura casual"
answer: 0
explanation: "Los prompts del sistema están por encima de las instrucciones del usuario en la jerarquía de contexto. Si un prompt del sistema dice 'siempre responde formalmente,' tu solicitud de lenguaje casual será sobrescrita. Por eso el mismo modelo se comporta diferente en diferentes aplicaciones — el prompt del sistema que no puedes ver está controlando el comportamiento."
```
