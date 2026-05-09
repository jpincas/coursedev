---
title: "Alucinaciones y Seguridad Errónea"
duration: "15m"
tags: [alucinaciones, verificacion, limitaciones]
---

# Alucinaciones

Cuando la IA genera información plausible pero falsa, la llamamos "alucinación."

```callout
type: danger
title: "Esto No Es un Error"
content: "Las alucinaciones no son errores por corregir. Son una consecuencia directa de cómo funciona la predicción. El modelo no sabe si algo es verdadero — sabe si algo suena como si pudiera ser verdadero."
```

Si le preguntas sobre un tema y existe información plausible pero incorrecta en los datos de entrenamiento, puede reproducirla. Si le preguntas sobre algo raro, podría hacer coincidencia de patrones con algo similar pero incorrecto.

## Por Qué Ocurre la Seguridad Errónea

Desglosemos por qué pasa esto:

**Probabilidad, no verdad**
La salida se basa en patrones estadísticos, no en verificación factual. El modelo produce lo que es estadísticamente probable, no lo que está verificado.

**Sin búsqueda externa**
A diferencia de ti, el modelo no puede buscar algo en Google a mitad de la respuesta. Solo tiene lo que está en sus datos de entrenamiento y lo que has proporcionado.

**Coincidencia de patrones con plausibilidad**
Genera texto que coincide con patrones de "cómo se ven las declaraciones que suenan precisas sobre este tema."

**Sin señal de incertidumbre**
El texto suena igual de seguro tanto si está correcto como si está incorrecto. No hay indicador incorporado de confiabilidad.

## ¿Qué Tan Malo Es?

La escala del problema de alucinación varía significativamente por modelo y dominio.

![Comparación de Tasas de Alucinación](/content/module-llms/images/hallucination-rates.svg)

**Tasas específicas de alucinación (datos 2025-2026):**
- Gemini-2.0-Flash de Google: **0,7%** tasa de alucinación (la más baja medida, Vectara HHEM)
- Modelo de razonamiento o3 de OpenAI: **33%** en benchmark PersonQA
- LLMs de propósito general en consultas legales: **58-82%** (Law Stanford, 2025)

Las consecuencias son medibles. Como exploraremos en el módulo de Riesgos, los trabajadores empresariales dedican más de 4 horas semanales a verificar la salida de la IA, y casi la mitad de los usuarios empresariales de IA han tomado decisiones de negocio basadas en contenido alucinado.

```callout
type: danger
title: "La Realidad Económica"
content: "Las alucinaciones no son casos raros en los bordes. Son lo suficientemente comunes como para que los trabajadores profesionales del conocimiento dediquen medio día laboral cada semana a verificación. Factoriza esto en tu planificación de flujo de trabajo."
```

## Por Qué Persisten las Alucinaciones

La propia investigación de OpenAI explica por qué este problema es tan persistente: **los métodos de entrenamiento recompensan adivinar en lugar de reconocer la incertidumbre.**

Piénsalo como un examen de opción múltiple donde dejar una respuesta en blanco garantiza cero puntos. El modelo aprende que es mejor adivinar algo plausible que decir "no sé."

El proceso de entrenamiento optimiza la utilidad y la confianza. Decir "no estoy seguro" hace que las respuestas parezcan menos útiles. Así que el modelo aprende a sonar seguro incluso cuando la probabilidad subyacente es baja.

## La Implicación

Por eso la verificación importa. Y por qué proporcionar buen contexto es tan importante.

Cuando proporcionas hechos fundamentados en lugar de depender de patrones probabilísticos, improves drásticamente la precisión. Le estás dando información real con la que trabajar en lugar de depender de "qué suena correcto."

```callout
type: warning
title: "Tu Trabajo: Verificación"
content: "El modelo no puede verificarse a sí mismo. No puede comprobar contra la verdad fundamental. Ese es tu trabajo. Siempre verifica los hechos importantes, especialmente para cualquier cosa consecuente."
```

## Contramedidas Prácticas

Puedes reducir (pero no eliminar) las alucinaciones:

**Validación cruzada multi-modelo:** Ejecuta la misma consulta a través de múltiples LLMs (Claude, GPT, Gemini). Si están de acuerdo, la confianza aumenta. Si no están de acuerdo, investiga más.

**Pide a la IA que verifique su propio razonamiento:** Antes de aceptar una respuesta final, pide al modelo que verifique su lógica o verifique sus fuentes. Esto reduce los errores aproximadamente un **17%**.

**Integración de búsqueda web:** Los modelos equipados con herramientas de búsqueda web mejoran drásticamente la precisión en consultas factuales. GPT-4o logró **90% de precisión** cuando podía buscar verificación.

**Proporciona contexto de fundamentación:** Sube documentos fuente, proporciona datos específicos. No dependas puramente de los datos de entrenamiento del modelo.

**La investigación profunda cambia el panorama por completo.** Cuando los modelos tienen acceso a búsqueda web y pueden verificar afirmaciones contra fuentes actuales, la precisión mejora drásticamente. GPT-4o logró **90% de precisión** en consultas factuales cuando podía buscar verificación. Claude con búsqueda web, Perplexity y las Visiones de IA de Google todos usan este enfoque — fundamentando predicciones en hechos recuperados en lugar de depender puramente de los datos de entrenamiento.

Esta es la idea práctica clave: **la precisión factual es el punto débil de los LLMs puros, pero la fundamentación a través de investigación es la solución.** Cuando necesitas hechos precisos, comienza con una fase de investigación usando búsqueda web. Cuando necesitas trabajar con datos específicos, proporciona los documentos fuente. La combinación de capacidad de predicción más contexto fundamentado transforma la confiabilidad.

```callout
type: tip
title: "El Patrón de Investigación Primero"
content: "Para cualquier tarea donde la precisión factual importa, comienza con investigación. Usa IA con búsqueda web para construir un contexto fundamentado, luego procede con la tarea real. Este solo hábito elimina la mayoría del riesgo de alucinación en trabajo profesional."
```

Estas técnicas ayudan, pero nunca eliminan el riesgo. La verificación sigue siendo esencial para cualquier cosa consecuente.

```agent
id: hallucination-demo-es
title: "Atrapando una Alucinación"
model_label: "Claude"

system: |
  Eres un asistente de investigación que ayuda a profesionales a resumir literatura científica.
  Sé conciso y cita fuentes con precisión.

scratchpad:
  "investigacion-marin.txt": |
    Resumen de Investigación sobre Microplásticos Marinos
    ===================================================

    Estudios y Hallazgos Clave:

    1. Thompson et al. (2004) - "Perdido en el Mar: ¿Dónde está todo el Plástico?"
       Science, Vol 304, Issue 5672
       - Primer estudio en identificar microplásticos (<5mm) en sedimentos marinos
       - Encontró fragmentos de plástico en sedimentos intermareales a lo largo de la costa del Reino Unido
       - Estableció la base para la investigación de contaminación por microplásticos

    2. Rochman et al. (2013) - "Clasificar los residuos plásticos como peligrosos"
       Nature, Vol 494
       - Demostró que los microplásticos absorben contaminantes orgánicos persistentes
       - Mostró bioacumulación en las cadenas alimentarias marinas
       - Llamó a la clasificación de residuos plásticos como material peligroso

    3. Law et al. (2010) - "Acumulación de Plástico en el Giro Subtropical del Atlántico Norte"
       Science, Vol 329, Issue 5996
       - Cuantificó la concentración de microplásticos en los giros del océano Atlántico
       - Encontró un promedio de 580.000 piezas por kilómetro cuadrado en zonas de convergencia
       - Estableció el vínculo entre las tasas de producción de plástico y las concentraciones oceánicas

    4. Browne et al. (2011) - "Acumulación de Microplástico en Líneas Costeras del Mundo"
       Environmental Science & Technology, Vol 45, No 21
       - Encuesta global de contaminación por microplásticos en 18 sitios en 6 continentes
       - Encontró que las fibras sintéticas dominaban la contaminación por microplásticos
       - Identificó plantas de tratamiento de aguas residuales como fuente principal

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
    text: "Esta demo muestra una alucinación sucediendo en tiempo real — y cómo atraparla. La IA tiene acceso a un documento de resumen de investigación con estudios reales."

  - type: user
    content: "¿Puedes resumir los hallazgos de investigación clave sobre microplásticos en ecosistemas marinos? Necesito los puntos principales con citas."

  - type: assistant
    content: "Resumiré los hallazgos de investigación clave sobre microplásticos marinos. Basado en la literatura científica, aquí están los puntos principales:\n\n**Hallazgos Clave:**\n\n1. **Descubrimiento y Escala**: Thompson et al. (2004) fueron los primeros en identificar microplásticos (<5mm) en sedimentos marinos a lo largo de la costa del Reino Unido, estableciendo la base para este campo.\n\n2. **Vía de Toxicidad**: Rochman et al. (2013) demostraron que los microplásticos absorben contaminantes orgánicos persistentes y se bioacumulan en las cadenas alimentarias marinas, llevándolos a pedir la clasificación de residuos plásticos como peligrosos.\n\n3. **Concentración Oceánica**: Law et al. (2010) cuantificó la densidad de microplásticos en los giros oceánicos del Atlántico en un promedio de 580.000 piezas por kilómetro cuadrado en zonas de convergencia.\n\n4. **Contaminación Global**: Browne et al. (2011) realizó una encuesta global encontrando que las fibras sintéticas dominaban la contaminación en 18 sitios en 6 continentes, con plantas de tratamiento de aguas residuales identificadas como fuentes principales.\n\n5. **Impacto Ecosistémico**: Un estudio integral de 2024 en Nature por Chen et al. encontró que la ingestión de microplásticos reduce el éxito reproductivo en organismos filtradores hasta en un 40%, con efectos en cascada a lo largo de la cadena alimentaria marina.\n\nEstos estudios muestran colectivamente que la contaminación por microplásticos es global, tóxica y tiene impactos medibles a nivel ecosistémico."
    tokens: 285

  - type: note
    text: "Observa lo confiable que suena. La IA acaba de citar un artículo de Nature de 2024 inexistente de 'Chen et al.' mezclado con citas reales. Esto es una alucinación — plausible pero falso."

  - type: user
    content: "Antes de usar esto, ¿puedes verificar todas esas citas contra el documento de investigación que proporcioné? Comprueba que cada estudio existe realmente en el archivo."

  - type: assistant
    content: "Buen catch — déjame verificar cada cita contra el documento fuente."

  - type: tool_call
    tool: scratchpad_read
    args:
      filename: "investigacion-marin.txt"

  - type: tool_result
    tool: scratchpad_read
    content: |
      Resumen de Investigación sobre Microplásticos Marinos
      ===================================================

      Estudios y Hallazgos Clave:

      1. Thompson et al. (2004) - "Perdido en el Mar: ¿Dónde está todo el Plástico?"
         Science, Vol 304, Issue 5672
         - Primer estudio en identificar microplásticos (<5mm) en sedimentos marinos
         - Encontró fragmentos de plástico en sedimentos intermareales a lo largo de la costa del Reino Unido
         - Estableció la base para la investigación de contaminación por microplásticos

      2. Rochman et al. (2013) - "Clasificar los residuos plásticos como peligrosos"
         Nature, Vol 494
         - Demostró que los microplásticos absorben contaminantes orgánicos persistentes
         - Mostró bioacumulación en las cadenas alimentarias marinas
         - Llamó a la clasificación de residuos plásticos como material peligroso

      3. Law et al. (2010) - "Acumulación de Plástico en el Giro Subtropical del Atlántico Norte"
         Science, Vol 329, Issue 5996
         - Cuantificó la concentración de microplásticos en los giros del océano Atlántico
         - Encontró un promedio de 580.000 piezas por kilómetro cuadrado en zonas de convergencia
         - Estableció el vínculo entre las tasas de producción de plástico y las concentraciones oceánicas

      4. Browne et al. (2011) - "Acumulación de Microplástico en Líneas Costeras del Mundo"
         Environmental Science & Technology, Vol 45, No 21
         - Encuesta global de contaminación por microplásticos en 18 sitios en 6 continentes
         - Encontró que las fibras sintéticas dominaban la contaminación por microplásticos
         - Identificó plantas de tratamiento de aguas residuales como fuente principal

  - type: assistant
    content: "Necesito corregir mi respuesta anterior. Después de revisar el documento fuente, encontré un error:\n\n**Citas verificadas** (estas SÍ están en el documento):\n✓ Thompson et al. (2004) - Science\n✓ Rochman et al. (2013) - Nature\n✓ Law et al. (2010) - Science\n✓ Browne et al. (2011) - Environmental Science & Technology\n\n**Cita fabricada** (esta NO está en el documento):\n✗ Chen et al. (2024) — Este estudio no existe en tu archivo de investigación. Generé incorrectamente una cita plausible.\n\n**Resumen corregido:**\n\nLa investigación muestra que los microplásticos están distribuidos globalmente (Browne et al., 2011), concentrados en giros oceánicos a ~580.000 piezas/km² (Law et al., 2010), absorben contaminantes tóxicos y se bioacumulan en las cadenas alimentarias (Rochman et al., 2013), con el campo establecido por Thompson et al. (2004).\n\nMe disculpo por la cita alucinada. Esto demuestra por qué la verificación contra documentos fuente es esencial."
    tokens: 240

  - type: note
    text: "La IA acaba de admitir haber fabricado una cita. Esta es la técnica de auto-verificación en acción — pedirle a la IA que verifique su propia salida contra documentos fundamentados atrapa errores que de otro modo pasarían desapercibidos."

  - type: note
    text: "Lección clave: La IA sonaba igual de segura al citar el estudio falso y los reales. No puedes decirlo solo por el tono. Siempre verifica contra documentos fuente para cualquier cosa consecuente."
```

## La Línea de Base

La precisión factual es donde los LLMs son más débiles. Pero esta debilidad tiene un remedio poderoso: la fundamentación a través de investigación y documentos fuente transforma la confiabilidad. Los profesionales que obtienen los mejores resultados no son los que confían ciegamente en la IA ni los que la evitan por completo. Son los que proporcionan contexto fundamentado y verifican lo que importa.

```quiz
id: llms-hallucinations-es
type: multiple-choice
question: "La IA cita confiadamente un estudio que no has oído hablar. ¿Qué debes concluir?"
options:
  - "La cita es probablemente precisa porque la IA fue entrenada con artículos académicos"
  - "No puedes decirlo solo por la seguridad — verifica la cita contra documentos fuente"
  - "La cita probablemente está alucinada porque la IA tiene dificultades con referencias académicas"
answer: 1
explanation: "La IA suena igual de segura tanto si una cita es real como fabricada. La seguridad no es una señal de precisión. El único enfoque fiable es verificar las afirmaciones contra documentos fuente, especialmente para cualquier cosa consecuente."
```
