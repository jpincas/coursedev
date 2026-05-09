---
title: "La Ventana de Contexto"
duration: "15m"
tags: [ventana-contexto, tokens, capacidad]
---

# ¿Qué Es la Ventana de Contexto?

La ventana de contexto es la "memoria de trabajo" del modelo. Es todo lo que el modelo puede ver al generar una respuesta.

![Ventana de Contexto: Prompt del Sistema, Historial de Conversación, Archivos y Mensaje Actual fluyen hacia el Procesamiento del Modelo](/content/module-context/images/context-window.svg)

## Tamaños de Ventana de Contexto

Las ventanas de contexto han crecido drásticamente con el tiempo:

| Era | Tamaño | Equivalente |
|-----|--------|-------------|
| GPT-3 (2020) | ~4K tokens | ~3.000 palabras |
| GPT-4 (inicial) | ~8K tokens | ~6.000 palabras |
| Claude 2 (2023) | ~100K tokens | ~75.000 palabras |
| 2025 (típico) | ~200K tokens | ~150.000 palabras |
| **Actual (2026)** | **~1M tokens** | **~750.000 palabras** |

**Claude Opus 4.6 ahora soporta 1 millón de tokens — aproximadamente 750.000 palabras, o unas 2.500 páginas de texto.** Los modelos Gemini soportan hasta 2 millones de tokens.

Pero hay un inconveniente: el efecto **"perdido en el medio."** Los modelos tienen dificultades con información enterrada profundamente en contextos grandes. La posición importa incluso más a esta escala.

```callout
type: info
title: "¿Qué es un Token?"
content: "Los tokens son fragmentos de texto que el modelo procesa. Aproximadamente 1 token = 0,75 palabras, o sobre 4 caracteres en inglés. 'Hola' es 1 token. 'Entendiendo' es 2 tokens."
```

## El Cambio en las Restricciones

Este crecimiento lo cambia todo.

**Restricción antigua:** ¿Qué cabe en el contexto? (Muy limitado. Había que ser selectivo.)

**Nueva restricción:** ¿Qué es relevante? (Puede caber libros enteros, pero la relevancia aún importa.)

Ahora puedes caber:
- Bases de código de aplicaciones enteras con dependencias
- Bibliotecas completas de documentación técnica
- Múltiples libros de longitud completa simultáneamente
- Días de historial de conversación detallado

**La función "Conversaciones Infinitas" de Anthropic** va aún más allá: la resumización del lado del servidor extiende las conversaciones indefinidamente sin perder contexto crítico.

Pero más grande no es automáticamente mejor. Los modelos aún necesitan encontrar lo relevante en toda esa información — y el efecto "perdido en el medio" significa que la estrategia de colocación importa más que nunca.

## Qué va Dónde

No todo el contexto es igual. Esto es lo que típicamente llena la ventana:

**Prompt del Sistema** (usualmente invisible para ti)
- Lo establece la aplicación que estás usando
- Define personalidad, capacidades, restricciones
- Por qué Claude en claude.ai se comporta diferente que Claude en una app personalizada

**Historial de Conversación**
- Todo lo que tú y el modelo han intercambiado
- Crece con cada mensaje
- Proporciona continuidad pero también puede saturar

**Archivos y Documentos**
- PDFs, archivos de código, imágenes, datos
- Proporcionados explícitamente por ti
- La forma más directa de dar información fundamentada

**Tu Mensaje Actual**
- Lo que acabas de escribir
- Tu instrucción o pregunta inmediata

## Viendo Cómo se Llena el Contexto

Esta demo hace visible lo invisible. Observa el contador de tokens mientras se cargan archivos en el contexto — puedes ver la ventana llenarse en tiempo real.

```agent
id: token-counting-demo-es
title: "La Ventana de Contexto Llenándose"
model_label: "Claude"

system: |
  Eres un analista de documentos. Lee los archivos proporcionados y responde preguntas
  sobre ellos. Sé conciso.

scratchpad:
  "notas-reunion.txt": |
    Equipo de Producto Semanal — 14 Ene 2026
    Asistentes: Maya (PM), Equipo de Desarrollo, Equipo de UX

    Decisiones:
    - Lanzar v2.3 para fin de mes
    - Posponer modo oscuro a v2.4
    - Contratar QA contractual para sprint de lanzamiento

    Elementos de acción:
    - Maya: Finalizar checklist de lanzamiento para viernes
    - Dev: Corregir bug crítico de autenticación (#4521)
    - UX: Actualizar mockups del flujo de incorporación
  "documento-estrategia.txt": |
    Estrategia de Producto Q1 2026

    Visión: Convertirnos en la herramienta de colaboración por defecto para equipos de mercado medio.

    Tres Pilares:
    1. Velocidad — reducir tiempos de carga de página a menos de 200ms
    2. Integración — lanzar conectores para Slack, Teams y Jira
    3. Auto-servicio — habilitar administración de equipo sin involucrar a TI

    Métricas objetivo:
    - Aumento del 40% en usuarios activos semanales
    - Net Promoter Score por encima de 50
    - Rotación por debajo del 3% mensual

    Panorama competitivo:
    - Notion expandiéndose en gestión de proyectos
    - Asana añadiendo funciones de IA
    - Monday.com agresivo en precios

    Nuestra ventaja: simplicidad y velocidad para equipos no técnicos.
  "retroalimentacion-cliente.txt": |
    Respuestas Recientes de NPS (Ene 2026)

    Puntuación 9: "Finalmente una herramienta que todo mi equipo realmente usa. La simplicidad es clave."
    Puntuación 8: "Me encanta pero necesito mejor app móvil. Difícil de usar en movimiento."
    Puntuación 3: "La integración con nuestras herramientas existentes es dolorosa. Tuvimos que construir soluciones alternativas."
    Puntuación 10: "Reemplacé 3 herramientas con esta una. Ahorro horas por semana."
    Puntuación 4: "La búsqueda está rota para espacios de trabajo grandes. No encuentro nada."
    Puntuación 9: "Mejor experiencia de incorporación que he visto. El equipo estuvo operativo en 20 minutos."
    Puntuación 2: "Sin soporte SSO es un punto decisivo para enterprise. Tuve que cancelar."
    Puntuación 7: "Buena pero no gran cosa. Le faltan gráficos Gantt para seguimiento de proyectos."

tools:
  - scratchpad_read

visibility:
  system_prompt: hidden
  tool_calls: visible
  full_context: hidden
  token_count: visible

sidebar:
  width: "45%"
  start_open: true

script:
  - type: note
    text: "Mira el contador de tokens en la barra lateral. Ahora el contexto está casi vacío — solo el prompt del sistema. Observa cómo crece mientras cargamos archivos."

  - type: user
    content: "Lee el archivo de notas de la reunión."

  - type: tool_call
    tool: scratchpad_read
    args:
      filename: "notas-reunion.txt"

  - type: tool_result
    tool: scratchpad_read
    content: |
      Equipo de Producto Semanal — 14 Ene 2026
      Asistentes: Maya (PM), Equipo de Desarrollo, Equipo de UX

      Decisiones:
      - Lanzar v2.3 para fin de mes
      - Posponer modo oscuro a v2.4
      - Contratar QA contractual para sprint de lanzamiento

      Elementos de acción:
      - Maya: Finalizar checklist de lanzamiento para viernes
      - Dev: Corregir bug crítico de autenticación (#4521)
      - UX: Actualizar mockups del flujo de incorporación

  - type: assistant
    content: "Leí las notas de la reunión. Puntos clave: v2.3 se lanza a fin de mes, modo oscuro pospuesto, y hay un bug crítico de autenticación por corregir. ¿Qué te gustaría saber?"
    tokens: 38

  - type: note
    text: "Revisa el contador de tokens — saltó. Ese solo archivo más la conversación añadieron tokens. Ahora carguemos los otros dos archivos y veamos cómo se acumula el contexto."

  - type: user
    content: "Ahora lee el documento de estrategia y la retroalimentación del cliente también. Luego dime: ¿está alineado el equipo de producto con la estrategia?"

  - type: tool_call
    tool: scratchpad_read
    args:
      filename: "documento-estrategia.txt"

  - type: tool_result
    tool: scratchpad_read
    content: |
      Estrategia de Producto Q1 2026

      Visión: Convertirnos en la herramienta de colaboración por defecto para equipos de mercado medio.

      Tres Pilares:
      1. Velocidad — reducir tiempos de carga de página a menos de 200ms
      2. Integración — lanzar conectores para Slack, Teams y Jira
      3. Auto-servicio — habilitar administración de equipo sin involucrar a TI

      Métricas objetivo:
      - Aumento del 40% en usuarios activos semanales
      - Net Promoter Score por encima de 50
      - Rotación por debajo del 3% mensual

      Panorama competitivo:
      - Notion expandiéndose en gestión de proyectos
      - Asana añadiendo funciones de IA
      - Monday.com agresivo en precios

      Nuestra ventaja: simplicidad y velocidad para equipos no técnicos.

  - type: tool_call
    tool: scratchpad_read
    args:
      filename: "retroalimentacion-cliente.txt"

  - type: tool_result
    tool: scratchpad_read
    content: |
      Respuestas Recientes de NPS (Ene 2026)

      Puntuación 9: "Finalmente una herramienta que todo mi equipo realmente usa. La simplicidad es clave."
      Puntuación 8: "Me encanta pero necesito mejor app móvil. Difícil de usar en movimiento."
      Puntuación 3: "La integración con nuestras herramientas existentes es dolorosa. Tuvimos que construir soluciones alternativas."
      Puntuación 10: "Reemplacé 3 herramientas con esta una. Ahorro horas por semana."
      Puntuación 4: "La búsqueda está rota para espacios de trabajo grandes. No encuentro nada."
      Puntuación 9: "Mejor experiencia de incorporación que he visto. El equipo estuvo operativo en 20 minutos."
      Puntuación 2: "Sin soporte SSO es un punto decisivo para enterprise. Tuve que cancelar."
      Puntuación 7: "Buena pero no gran cosa. Le faltan gráficos Gantt para seguimiento de proyectos."

  - type: assistant
    content: "**Parcialmente alineado, con una brecha importante.**\n\nLa estrategia tiene tres pilares: Velocidad, Integración y Auto-servicio.\n\n**Alineado:** La reunión se centra en lanzar v2.3 (velocidad/características) y corregir el bug de autenticación (estabilidad). La retroalimentación del cliente confirma que la simplicidad es la fortaleza — puntuaciones de 9 y 10 elogian la facilidad de uso.\n\n**Desalineado:** El segundo pilar de la estrategia es Integración, pero la reunión no lo menciona en absoluto. Mientras tanto, la retroalimentación del cliente señala esto explícitamente: 'La integración con herramientas existentes es dolorosa' (Puntuación 3) y 'Sin soporte SSO es un punto decisivo' (Puntuación 2). Estas son las puntuaciones más bajas en el conjunto de datos.\n\n**Recomendación:** El equipo debería priorizar el pilar de integración en el próximo sprint. El dolor del cliente y la prioridad estratégica están alineados, pero la ejecución semanal no lo refleja."
    tokens: 165

  - type: note
    text: "Mira el contador de tokens ahora versus cuando empezamos. Tres archivos más la conversación han consumido un trozo significativo del contexto. En un flujo de trabajo real con documentos más grandes, esto se acumula rápido — por eso la relevancia importa más que el volumen."
```

```quiz
id: context-window-understanding-es
type: multiple-choice
question: "Subes un manual técnico de 500 páginas a la IA y haces una pregunta específica. La respuesta es vaga y genérica. ¿Cuál es la causa más probable?"
options:
  - "El modelo no pudo procesar tantas páginas y truncó silenciosamente el documento"
  - "La información relevante estaba enterrada en el medio donde la atención del modelo es más débil"
  - "El modelo necesita que le especifiques qué números de página enfocarse antes de poder dar una buena respuesta"
answer: 1
explanation: "El efecto 'perdido en el medio' significa que los modelos tienen dificultades con información enterrada profundamente en contextos grandes. Aunque la ventana es lo suficientemente grande para caber el documento, la atención del modelo es más fuerte al principio y al final. Colocar información clave estratégicamente importa más que meterlo todo."
```
