---
title: "Los LLMs Son Motores de Predicción"
duration: "20m"
tags: [llms, prediccion, siguiente-token]
---

# ¿Qué Es un Modelo de Lenguaje Grande?

Comencemos con la pregunta fundamental: ¿Qué es un LLM?

**Es un motor de predicción.**

Eso es todo. Eso es lo fundamental. Todo lo demás — la aparente inteligencia, la creatividad, la comprensión — emerge de una capacidad: predecir qué sigue.

## Predicción del Siguiente Token

Así es como funciona a nivel más básico.

El modelo ve una secuencia de texto y predice qué sigue — no solo una vez, sino token por token, construyendo la respuesta un paso a la vez.

![Cadena de Predicción del Siguiente Token](/content/module-llms/images/token-prediction-chain.svg)

En cada paso, el modelo evalúa todos los posibles tokens siguientes y asigna probabilidades. Elige la opción con mayor probabilidad y la añade a la secuencia. Luego lo hace de nuevo con la secuencia nueva y más larga. Y otra vez. Palabra por palabra, token por token, construyendo una respuesta.

```callout
type: info
title: "Token por Token"
content: "Observa cómo construye: 'La' → 'capital' → 'de' → 'Francia' → 'es' → 'París' → '.' En cada paso, mira todo lo que vino antes y predice qué sigue."
```

Escalá esto — miles de millones de parámetros, entrenado con billones de palabras — y obtienes algo que puede escribir ensayos, código, análisis, ficción creativa.

Pero sigue siendo, fundamentalmente, predicción del siguiente token.

## Por Qué Esto Importa

Entender el mecanismo lo explica todo.

**Explica las fortalezas:**
- Excelente en continuación de patrones
- Muy bueno con formatos que ha visto antes
- Bueno en "qué suele seguir"
- Excelente en coincidencia de estilo

**Explica las debilidades:**
- Sin acceso a hechos externos
- No puede verificar sus propias afirmaciones
- Seguro de cosas incorrectas
- Hace coincidencia de patrones con respuestas plausibles pero incorrectas

Si le das algo que parece patrones que ha visto antes, destaca. ¿Correo profesional? Ha visto millones. ¿Formato de documento legal? Ha visto millones. ¿Código en Python? Ha visto millones.

Pero no tiene conexión con la verdad. No "sabe" cosas — predice cómo se ve el texto sobre cosas.

## Qué Hace Bien la IA (y Qué No)

Este mecanismo de predicción crea un patrón fascinante: **los problemas difíciles son fáciles y los problemas fáciles son difíciles.**

Esta idea proviene de la Fundación Carnegie para la Paz Internacional, y captura algo fundamental sobre la capacidad de la IA en 2026.

**La IA destaca en:**
- **Generación de código** — El 41% de todo el código escrito globalmente ahora es generado por IA
- **Reconocimiento de patrones** — Revisión de documentos, análisis de contratos, detección de anomalías
- **Síntesis** — Combinar información de múltiples fuentes en resúmenes coherentes
- **Primeros borradores** — Obtener algo utilizable rápidamente en cualquier formato

**La IA tiene dificultades con:**
- **Desvío por alucinación** — La precisión se degrada en cadenas muy largas de razonamiento
- **Matiz cultural** — Significado dependiente del contexto, implicaciones sutiles
- **Sensibilidad al enunciado** — Cambios menores en cómo preguntas pueden producir resultados muy diferentes
- **Razonamiento de sentido común novedoso** — Situaciones genuinamente nuevas donde la coincidencia de patrones falla

La paradoja: La IA puede escribir código listo para producción para algoritmos complejos, pero podría afirmar con seguridad que un evento histórico ficticio ocurrió porque la descripción "suena correcta."

```callout
type: tip
title: "Calibrando Tus Expectativas"
content: "La IA es extraordinariamente buena en tareas con patrones claros y estructuras que ha visto millones de veces. Es menos fiable cuando la novedad genuina, el contexto cultural o la fundamentación factual importan. Planifica tu trabajo en consecuencia."
```

## Predicción en Acción

Observa cómo cambiar el contexto cambia lo que el modelo predice. La misma apertura, dos contextos diferentes, dos resultados completamente diferentes.

```agent
id: prediction-in-action-es
title: "La Predicción Cambia con el Contexto"
model_label: "Claude"

system: |
  Eres un asistente de escritura útil. Completas tareas basándote en el
  contexto e instrucciones proporcionadas. Sé conciso y directo.

scratchpad:
  "resumen-empresa.txt": |
    Meridian Technologies — Informe Anual 2025

    Fundada: 2019
    Industria: Infraestructura en la Nube
    Empleados: 340
    Ingresos: $42M (28% interanual)
    Producto Principal: CloudBridge — plataforma de migración automatizada

    Destacados Q4:
    - Firmó 3 clientes Fortune 500
    - Lanzó CloudBridge v3 con migración asistida por IA
    - Abrió oficina en Londres para expansión europea
    - Nombrada "Estrella Emergente" por Gartner en categoría de migración en la nube

tools:
  - scratchpad_read

visibility:
  system_prompt: hidden
  tool_calls: visible
  full_context: hidden

sidebar:
  width: "45%"
  start_open: true

script:
  - type: note
    text: "Observa cómo la misma petición produce resultados completamente diferentes dependiendo del contexto que el modelo ha visto. Primero, pediremos un resumen sin proporcionar información — para mostrar cómo se ve la predicción pura."

  - type: user
    content: "Escribe un párrafo de resumen sobre Meridian Technologies."

  - type: assistant
    content: "No tengo ninguna información sobre Meridian Technologies en mis datos de entrenamiento ni en archivos. Necesitaría que proporcionaras contexto — tal vez un documento de resumen de la empresa o un informe anual — antes de que pueda escribir un resumen preciso."

  - type: note
    text: "El modelo indicó correctamente que le falta contexto. Pero observa qué pasa si simulamos un error común: pedir sin verificar si la empresa existe en los datos de entrenamiento. Reiniciemos y mostremos el contraste proporcionando el contexto correctamente."

  - type: user
    content: "Déjame empezar de nuevo. Tengo el resumen de nuestra empresa en un archivo. Lee resumen-empresa.txt y luego escribe un párrafo de resumen."

  - type: tool_call
    tool: scratchpad_read
    args:
      filename: "resumen-empresa.txt"

  - type: tool_result
    tool: scratchpad_read
    content: |
      Meridian Technologies — Informe Anual 2025

      Fundada: 2019
      Industria: Infraestructura en la Nube
      Empleados: 340
      Ingresos: $42M (28% interanual)
      Producto Principal: CloudBridge — plataforma de migración automatizada

      Destacados Q4:
      - Firmó 3 clientes Fortune 500
      - Lanzó CloudBridge v3 con migración asistida por IA
      - Abrió oficina en Londres para expansión europea
      - Nombrada "Estrella Emergente" por Gartner en categoría de migración en la nube

  - type: assistant
    content: "Meridian Technologies es una empresa de infraestructura en la nube fundada en 2019 que ha crecido a 340 empleados e ingresos de $42M, un 28% interanual. El producto estrella de la empresa, CloudBridge, es una plataforma de migración automatizada que lanzó recientemente su tercera versión con capacidades asistidas por IA. En el Q4 de 2025, Meridian firmó tres clientes Fortune 500, se expandió a Europa con una oficina en Londres y obtuvo el reconocimiento de 'Estrella Emergente' de Gartner en migración en la nube."

  - type: note
    text: "Mismo modelo. Misma petición. Resultado dramáticamente diferente. La primera respuesta fue el modelo prediciendo '¿cómo se ve una descripción genérica de una empresa?' La segunda fue prediciendo '¿cómo se ve el resumen de ESTA empresa, dados estos hechos?' El contexto cambió la predicción."

  - type: note
    text: "Este es el mecanismo fundamental en acción. El modelo no 'sabe' sobre Meridian Technologies. Predice cómo se ve un buen texto dado los patrones que ha visto — y el contexto que proporcionas moldea esas predicciones por completo."
```

```quiz
id: llms-core-mechanism-es
type: multiple-choice
question: "Un LLM escribe un memorial legal perfectamente formateado pero factualmente incorrecto. ¿Qué explicación mejor rinde cuenta de esto?"
options:
  - "El modelo ha visto millones de memorials legales y predice texto legal plausible, independientemente de la precisión factual"
  - "Los datos de entrenamiento legal del modelo estaban desactualizados, causando que referenciara jurisprudencia obsoleta"
  - "El dominio legal requiere capacidades de verificación que no se incluyeron en el proceso de entrenamiento"
answer: 0
explanation: "Los LLMs son motores de predicción. Predicen cómo se ve el texto basado en patrones, no en lo que es verdadero. Un modelo que ha visto millones de memorials legales producirá texto legal perfectamente formateado — pero no tiene mecanismo para verificar si los casos citados o los hechos son reales."
```
