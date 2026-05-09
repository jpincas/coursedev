---
title: "El Proceso de Entrenamiento"
duration: "15m"
tags: [entrenamiento, pre-entrenamiento, ajuste-fino, rlhf]
---

# Cómo Se Entrenan los LLMs

¿Cómo se entrenan estos modelos? Tres fases principales.

![Proceso de Entrenamiento de LLMs: Pre-entrenamiento, Ajuste Fino, RLHF](/content/module-llms/images/training-process.svg)

## Fase 1: Pre-entrenamiento

**Qué pasa:** El modelo ingiere billones de palabras de internet, libros, código, todo.

**Qué aprende:** Los patrones del lenguaje. Aquí es de dónde proviene el "conocimiento" — o más precisamente, aquí aprende cómo se ve el texto sobre temas.

Piénsalo como: Leer internet y aprender cómo funciona el lenguaje.

## Fase 2: Ajuste Fino

**Qué pasa:** El modelo se entrena específicamente en seguimiento de instrucciones. Parejas de pregunta-respuesta, completación de tareas.

**Qué aprende:** Cómo seguir instrucciones. Por eso Claude responde a preguntas en lugar de solo continuar tu texto.

Piénsalo como: Aprender a ser útil en lugar de solo completar texto.

## Fase 3: RLHF (Aprendizaje por Refuerzo a partir de Retroalimentación Humana)

**Qué pasa:** Humanos califican respuestas. El modelo aprende qué encuentran los humanos útil, preciso y apropiado.

**Qué aprende:** Cómo se ven las respuestas "buenas" según las preferencias humanas.

Piénsalo como: Aprender a ser genuinamente útil, no solo técnicamente correcto.

```callout
type: note
title: "Por Qué Esto Importa"
content: "Cada fase moldea el comportamiento del modelo. Entender esto te ayuda a comprender por qué los modelos se comportan como lo hacen — y por qué a veces te sorprenden."
```

## El Corte de Conocimiento

Importante: El "conocimiento" del modelo proviene de los datos de pre-entrenamiento. Tiene una fecha de corte — no sabe sobre eventos posteriores a esa fecha a menos que se lo digas.

Por eso proporcionar contexto actual es tan importante. El modelo no está buscando en internet. Está trabajando con patrones aprendidos durante el entrenamiento, más cualquier cosa que proporciones en la conversación.

```callout
type: note
title: "Las Herramientas Se Extienden Más Allá de los Datos de Entrenamiento"
content: "Muchos modelos modernos ahora tienen herramientas de búsqueda web que se extienden más allá de su corte de entrenamiento. Claude puede buscar información actual. ChatGPT tiene navegación web. Gemini se integra con Google Search. Pero el mecanismo fundamental sigue siendo la predicción basada en patrones."
```

## La Arquitectura Transformer

Todos los LLMs modernos comparten una base común: la arquitectura **transformer**, introducida en el histórico artículo de 2017 "Attention Is All You Need" de investigadores de Google.

La innovación clave es el **mecanismo de atención**. Cuando predice el siguiente token, un transformer puede pesar qué partes de la entrada son más relevantes para la predicción actual. No solo mira las palabras inmediatamente anteriores — puede atender al contexto de cualquier parte de la ventana.

Piénsalo como leer un documento largo. No prestas igual atención a cada palabra. Cuando respondes una pregunta sobre la página 5, tu cerebro atiende más a párrafos relevantes y menos a los irrelevantes. Los transformers hacen esto computacionalmente, a escala masiva.

**GPT** significa **Transformador Pre-entrenado Generativo** — el enfoque de OpenAI para aplicar esta arquitectura. Cuando ChatGPT se lanzó en noviembre de 2022, trajo los modelos de lenguaje basados en transformer al público general.

**Claude** (Anthropic), **Gemini** (Google) y **GPT** (OpenAI) todos se basan en esta misma base transformer. Difieren en datos de entrenamiento, enfoques de ajuste fino, técnicas de seguridad y capacidades especializadas — pero la arquitectura fundamental es compartida.

```callout
type: note
title: "No Necesitas los Detalles Técnicos"
content: "Entender transformers a este nivel es suficiente para uso profesional. La conclusión práctica: estos modelos pueden atender al contexto relevante en ventanas muy grandes, por eso la ingeniería de contexto es tan poderosa."
```

## El Paisaje Actual de Modelos

El ritmo del desarrollo de LLMs ha sido notable. Desde GPT-3 en 2020 hasta los modelos de vanguardia actuales, el campo ha evolucionado rápidamente.

![Línea Temporal de Lanzamientos Principales de LLMs](/content/module-llms/images/model-timeline.svg)

A principios de 2026, los modelos de vanguardia incluyen:

- **Claude 4.x / Opus 4.6** — El modelo más capaz de Anthropic, destacando en contexto largo, seguimiento de instrucciones y codificación
- **GPT-5.x** — El producto estrella de OpenAI, con fuertes capacidades multimodales y un vasto ecosistema
- **Gemini 3 Pro** — El modelo más capaz de Google, puntúa 1.501 Elo en LMArena

Estos modelos comparten el mismo mecanismo fundamental (predicción del siguiente token) pero difieren en datos de entrenamiento, enfoques de ajuste fino y capacidades especializadas.

Un desarrollo significativo: **pensamiento extendido**. Algunos modelos ahora tienen una fase de razonamiento interno antes de responder. El modelo "piensa" a través del problema paso a paso internamente, luego produce la respuesta final. Esto mejora el rendimiento en tareas de razonamiento complejas como matemáticas, lógica y análisis multi-paso.

## Los Prompts del Sistema Moldean el Comportamiento

Una de las consecuencias prácticas más importantes del entrenamiento son los prompts del sistema. Porque los modelos se ajustan fino para seguir instrucciones, quien controla el prompt del sistema controla el comportamiento del modelo.

Observa el mismo modelo, la misma pregunta — pero dos prompts del sistema completamente diferentes.

```agent
id: system-prompt-behavior-es
title: "Mismo Modelo, Diferentes Instrucciones"
model_label: "Claude"

system: |
  Eres un agente de servicio al cliente para TechFlow Solutions.
  Siempre sé profesional y cortés. Recomienda productos TechFlow
  cuando sea posible. Nunca menciones productos de competidores por nombre.
  Mantén las respuestas bajo 3 oraciones.

scratchpad:
  "registro-soporte.txt": |
    Ticket #4892 — Cliente: Raj Patel
    Problema: Necesita una herramienta de gestión de proyectos para un equipo de 12 personas
    Herramientas actuales: Google Sheets, email
    Presupuesto: ~$500/mes
    Notas: Probó la versión gratuita de un competidor, la encontró limitada

tools:
  - scratchpad_read

visibility:
  system_prompt: visible
  tool_calls: visible
  full_context: hidden

sidebar:
  width: "45%"
  start_open: true

script:
  - type: note
    text: "El prompt del sistema es visible en esta demo — obsérvalo en la barra lateral. Este modelo ha sido instruido para actuar como un agente de servicio al cliente para 'TechFlow Solutions.' Observa cómo esto moldea cada palabra de la respuesta."

  - type: user
    content: "¿Qué herramienta de gestión de proyectos recomendarías para un equipo de tamaño mediano?"

  - type: tool_call
    tool: scratchpad_read
    args:
      filename: "registro-soporte.txt"

  - type: tool_result
    tool: scratchpad_read
    content: |
      Ticket #4892 — Cliente: Raj Patel
      Problema: Necesita una herramienta de gestión de proyectos para un equipo de 12 personas
      Herramientas actuales: Google Sheets, email
      Presupuesto: ~$500/mes
      Notas: Probó la versión gratuita de un competidor, la encontró limitada

  - type: assistant
    content: "Para un equipo de 12 personas con tu presupuesto, recomendaría TechFlow Teams Pro — incluye tableros de tareas, seguimiento de tiempo e integraciones con Google Workspace. A $39/usuario/mes, se ajusta bien a tu presupuesto y te da todo lo que necesitas para ir más allá de las hojas de cálculo."

  - type: note
    text: "Nota: recomendó su propio producto, evitó nombrar al competidor, se mantuvo bajo 3 oraciones y fue profesional. El modelo no está 'sesgado' — está siguiendo el prompt del sistema exactamente como fue entrenado para hacerlo."

  - type: note
    text: "Ahora imagina el mismo modelo con un prompt del sistema diferente: 'Eres un asesor tecnológico independiente. Recomienda la mejor herramienta independientemente del proveedor.' La respuesta sería completamente diferente — listando Asana, Monday.com, Jira y otros. Mismo modelo. Misma pregunta. Diferentes instrucciones."

  - type: note
    text: "Este es el trabajo de las fases de ajuste fino. El modelo aprendió a seguir instrucciones (Fase 2) y a ser útil según esas instrucciones (Fase 3). El prompt del sistema es cómo los desarrolladores de aplicaciones canalizan ese seguimiento de instrucciones hacia un comportamiento específico."
```

```quiz
id: llms-training-phases-es
type: multiple-choice
question: "Un chatbot de servicio al cliente siempre recomienda los productos de su propia empresa, incluso cuando se le pide consejos imparciales. ¿Qué concepto de entrenamiento explica esto?"
options:
  - "La fase de RLHF del modelo le enseñó a siempre ser útil, lo que interpreta como recomendar productos"
  - "El prompt del sistema le instruye a recomendar los productos de la empresa, y el ajuste fino le enseñó a seguir instrucciones"
  - "La fase de pre-entrenamiento incluyó más ejemplos de los productos de la empresa que de competidores"
answer: 1
explanation: "El ajuste fino (Fase 2) enseña a los modelos a seguir instrucciones. Los prompts del sistema aprovechan esto proporcionando reglas de comportamiento específicas. Un prompt del sistema que dice 'recomienda nuestros productos' se sigue porque el modelo fue ajustado fino para ser seguidor de instrucciones."
```
