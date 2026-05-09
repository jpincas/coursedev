---
title: "Hacer que la IA Escriba Como Tú"
duration: "12m"
tags: [style, voice, writing]
---

# Hacer que la IA Escriba Como Tú

La IA escribe en una voz genérica por defecto. Puedes entrenarla para que escriba como tú.

## Tarjetas de Estilo

Una **tarjeta de estilo** es un prompt reutilizable que codifica tus preferencias.

Define:
- **Tono** — formal, conversacional, directo, cálido
- **Vocabulario** — términos técnicos, lenguaje llano, jerga del sector
- **Estructura de las frases** — frases cortas e impactantes, prosa fluida, listas con viñetas
- **Audiencia** — expertos, público general, ejecutivos, estudiantes

Guárdala. Reutilízala. Refínala con el tiempo.

## Fundamentación con Ejemplos

La forma más rápida de enseñar a la IA tu estilo:

**Proporciona 2-3 párrafos de tu mejor escritura y pide a la IA que analice tu estilo.**

Dale un trabajo reciente del que estés orgulloso. Pide:
- "Analiza el tono y la estructura de esta escritura."
- "¿Qué patrones de vocabulario detectas?"
- "¿Cómo describirías el ritmo de las frases?"

Y luego: "Escribe la siguiente sección usando este estilo."

```callout
type: tip
title: "Enfoque Práctico"
content: "Empieza cada proyecto de escritura dando a la IA 2-3 muestras de tu mejor trabajo y pidiéndole que analice tu estilo antes de redactar. Esta inversión de 2 minutos transforma la calidad de la salida."
```

## La Técnica de EchoWriting

Para transferencia de estilo persistente:

![Proceso de EchoWriting](/content/module-writing/images/echowriting-process.svg)

1. **Introduce 15-20 muestras** de tu escritura (correos, informes, artículos)
2. **Pídele que analice tus patrones de estilo** — lo que hace que tu escritura sea reconocible como tuya
3. **Crea un prompt de estilo persistente** que reutilices en todas las sesiones

Guarda este prompt en un Proyecto de Claude, Instrucciones Personalizadas de ChatGPT o tu archivo de instrucciones persistente.

Cada conversación empieza con tu estilo ya cargado.

```agent
id: echowriting-demo
title: "Escribir en tu Tono"
model_label: "Claude"

system: |
  Eres un asistente de escritura. Cuando analices el estilo de escritura, identifica
  patrones concretos en tono, estructura de frases, vocabulario y formato. Al escribir
  en un estilo específico, coincide con esos patrones con precisión.

scratchpad:
  "sample1.txt": |
    Asunto: Acceso al Panel Q3

    El equipo de analíticas necesita acceso al panel para el viernes.

    Bloqueos actuales:
    - IT no ha provisionado las cuentas
    - Los documentos de formación no están listos

    ¿Puedes escalar la solicitud a IT? Yo terminaré los documentos el miércoles.

    Avísame.

    Sarah
  "sample2.txt": |
    ## Estado del Proyecto — Semana 12

    Cumplimos el hito. La integración de la API está operativa. Dos problemas aparecieron en las pruebas.

    El tiempo de espera de autenticación es demasiado corto. Los usuarios se desconectan a mitad de sesión.
    Solucionado en el siguiente sprint.

    La función de exportación falla con datasets grandes. El equipo de desarrollo lo está investigando.
    Solución alternativa: dividir las exportaciones en fragmentos más pequeños.

    Próxima revisión: 15 de marzo.
  "sample3.txt": |
    Gracias por la propuesta.

    El plazo parece ajustado. ¿Podemos retrasar el lanzamiento a mayo? Eso nos da
    tres semanas adicionales para pruebas.

    El presupuesto está aprobado. La contingencia de £15K debería cubrir los costes del contratista.

    Una pregunta: ¿quién asume el soporte post-lanzamiento? Necesitamos eso definido antes
    de dar el visto bueno.

    Hablemos el viernes.
  "brief.txt": |
    CONTENIDO A REDACTAR:
    Correo de actualización del proyecto al equipo ejecutivo

    PUNTOS CLAVE:
    - Fase 1 completada según lo previsto (migración API, autenticación de usuarios)
    - Fase 2 retrasada 2 semanas por problemas inesperados con el esquema de la base de datos
    - Presupuesto actualmente £8.000 por debajo del plan pero podría necesitar contingencia para la Fase 2
    - Cliente satisfecho con el progreso, pequeñas solicitudes de funcionalidades para la Fase 3
    - Próximo hito: entrega de la Fase 2, ahora prevista para el 30 de abril

    CONTEXTO:
    Esta es una actualización mensual. Los ejecutivos quieren hechos, no rodeos. Les interesan
    el plazo, el presupuesto y los riesgos. Mantén un tono profesional pero sin lenguaje corporativo vacío.

tools:
  - scratchpad_read
  - scratchpad_write

visibility:
  system_prompt: hidden
  tool_calls: visible
  full_context: hidden

sidebar:
  width: "45%"
  start_open: true

script:
  - type: note
    text: "Esto es EchoWriting en acción. Tres muestras de escritura están en el explorador de archivos — todas escritas por la misma persona. Fijáos en los patrones consistentes: párrafos cortos, tono directo, inglés británico, sin jerga."

  - type: note
    text: "Observa cómo la IA analiza el estilo, y luego usa ese análisis para escribir nuevo contenido que suena como si la misma persona lo hubiera escrito."

  - type: user
    content: "Por favor, lee estas tres muestras de escritura y analiza mi estilo. ¿Qué patrones detectas?"

  - type: assistant
    content: "Revisaré las tres muestras e identificaré los patrones que definen tu estilo de escritura."

  - type: tool_call
    tool: scratchpad_read
    args:
      filename: "sample1.txt"

  - type: tool_result
    tool: scratchpad_read
    content: |
      Asunto: Acceso al Panel Q3

      El equipo de analíticas necesita acceso al panel para el viernes.

      Bloqueos actuales:
      - IT no ha provisionado las cuentas
      - Los documentos de formación no están listos

      ¿Puedes escalar la solicitud a IT? Yo terminaré los documentos el miércoles.

      Avísame.

      Sarah

  - type: tool_call
    tool: scratchpad_read
    args:
      filename: "sample2.txt"

  - type: tool_result
    tool: scratchpad_read
    content: |
      ## Estado del Proyecto — Semana 12

      Cumplimos el hito. La integración de la API está operativa. Dos problemas aparecieron en las pruebas.

      El tiempo de espera de autenticación es demasiado corto. Los usuarios se desconectan a mitad de sesión.
      Solucionado en el siguiente sprint.

      La función de exportación falla con datasets grandes. El equipo de desarrollo lo está investigando.
      Solución alternativa: dividir las exportaciones en fragmentos más pequeños.

      Próxima revisión: 15 de marzo.

  - type: tool_call
    tool: scratchpad_read
    args:
      filename: "sample3.txt"

  - type: tool_result
    tool: scratchpad_read
    content: |
      Gracias por la propuesta.

      El plazo parece ajustado. ¿Podemos retrasar el lanzamiento a mayo? Eso nos da
      tres semanas adicionales para pruebas.

      El presupuesto está aprobado. La contingencia de £15K debería cubrir los costes del contratista.

      Una pregunta: ¿quién asume el soporte post-lanzamiento? Necesitamos eso definido antes
      de dar el visto bueno.

      Hablemos el viernes.

  - type: assistant
    content: |
      Patrones claros en las tres muestras:

      **Estructura de las frases:** Frases cortas e impactantes. Rara vez más de 15 palabras. Párrafos de una sola frase frecuentes para énfasis.

      **Tono:** Directo y sin rodeos. Expones hechos, identificas problemas, propones soluciones. Sin lenguaje dudoso ("podría", "quizás", "posiblemente").

      **Vocabulario:** Lenguaje llano en todo momento. Inglés británico ("realised" en vez de "realized", "£" en vez de "$"). Voz activa dominante. Jerga mínima.

      **Formato:** Saltos de párrafo frecuentes para legibilidad. Puntos con viñetas para listas. Encabezados en mayúscula solo en la primera palabra. Sin bloques largos de texto.

      **Voz:** Segura y orientada a la acción. Las preguntas son solicitudes directas, no sugerencias. Asumes la información.

  - type: user
    content: "Perfecto. Ahora redacta un correo de actualización del proyecto usando el brief en brief.txt, pero coincide exactamente con mi estilo de escritura."

  - type: assistant
    content: "Leeré el brief y redactaré la actualización en tu estilo — párrafos cortos, tono directo, inglés británico, voz activa."

  - type: tool_call
    tool: scratchpad_read
    args:
      filename: "brief.txt"

  - type: tool_result
    tool: scratchpad_read
    content: |
      CONTENIDO A REDACTAR:
      Correo de actualización del proyecto al equipo ejecutivo

      PUNTOS CLAVE:
      - Fase 1 completada según lo previsto (migración API, autenticación de usuarios)
      - Fase 2 retrasada 2 semanas por problemas inesperados con el esquema de la base de datos
      - Presupuesto actualmente £8.000 por debajo del plan pero podría necesitar contingencia para la Fase 2
      - Cliente satisfecho con el progreso, pequeñas solicitudes de funcionalidades para la Fase 3
      - Próximo hito: entrega de la Fase 2, ahora prevista para el 30 de abril

      CONTEXTO:
      Esta es una actualización mensual. Los ejecutivos quieren hechos, no rodeos. Les interesan
      el plazo, el presupuesto y los riesgos. Mantén un tono profesional pero sin lenguaje corporativo vacío.

  - type: tool_call
    tool: scratchpad_write
    args:
      filename: "project-update-email.txt"
      content: |
        Asunto: Actualización del Proyecto — Marzo

        La Fase 1 está terminada. Migración de la API y autenticación de usuarios entregadas según lo previsto.

        La Fase 2 está retrasada dos semanas. Los problemas con el esquema de la base de datos
        tardaron más de lo previsto en resolverse. Nueva fecha de entrega: 30 de abril.

        Estado del presupuesto: £8.000 por debajo del plan. Podríamos necesitar contingencia para la Fase 2.
        Confirmaré la próxima semana.

        La feedback del cliente es positiva. Han solicitado pequeñas funcionalidades para la Fase 3.
        Nada que afecte al plazo actual.

        Riesgos:
        - El retraso de la Fase 2 empuja el lanzamiento final a mediados de mayo
        - El trabajo con la base de datos podría descubrir más problemas

        Próximo hito: entrega de la Fase 2, 30 de abril.

        ¿Preguntas?

  - type: tool_result
    tool: scratchpad_write
    content: "Escrito en project-update-email.txt (17 líneas)"

  - type: assistant
    content: "Listo. He redactado la actualización coincidiendo con tu estilo — párrafos cortos, afirmaciones directas, inglés británico, voz activa, estructura sin viñetas con claros saltos de párrafo para legibilidad. Mismo tono seguro y sin rodeos que tus muestras."

  - type: note
    text: "Compara la salida de la IA con las muestras originales. Mismo ritmo. Misma directo. Misma estructura de párrafos. La IA aprendió tu voz de tres ejemplos."

  - type: note
    text: "Esto es EchoWriting. Introduce muestras de tu escritura en la IA, pídele que analice los patrones, y entonces cada documento que cree sonará como si lo hubieras escrito tú. Sin la voz genérica de la IA. Tu voz."
```

## Síntesis Multi-Documento

El caso de uso de mayor valor de la IA para el trabajo de conocimiento: sintetizar entre múltiples fuentes.

Sube múltiples documentos a Claude o ChatGPT. Pídele que:
- Identifique temas comunes
- Detecte contradicciones
- Encuentre lagunas en la cobertura
- Extraiga argumentos clave

Google's NotebookLM se especializa en esto — analiza información a través de documentos, sitios web y tipos de medios simultáneamente.

```callout
type: warning
title: "Limitación de la Síntesis"
content: "La IA es mejor organizando y resumiendo que haciendo una síntesis intelectual genuina. Verifica siempre que no haya fabricado conexiones entre fuentes. Comprueba que los patrones reclamados realmente existan en el material fuente."
```

## Estilo vs Esencia

La IA puede imitar tu estilo. No puede replicar tu juicio.

**Tu estilo:** estructura de frases, elección de palabras, ritmo de párrafos. La IA puede aprender esto.

**Tu juicio:** qué decir, qué enfatizar, qué argumento hacer. La IA no puede replicar esto.

Usa la IA para escribir en tu voz. Nunca delegues la decisión de lo que esa voz debe decir.

```quiz
id: style-transfer-quiz
type: multiple-choice
question: "¿Qué es la técnica EchoWriting?"
options:
  - "Redactar borradores completos tú mismo y pedir a la IA que los convierta en diferentes formatos"
  - "Introducir 15-20 muestras de tu escritura en la IA para que analice los patrones, y luego crear un prompt de estilo persistente"
  - "Pedir a la IA que repita tus exactas palabras para verificar la comprensión"
answer: 1
explanation: "EchoWriting consiste en dar a la IA muchas muestras de tu escritura, pedirle que analice tus patrones de estilo, y luego crear un prompt de estilo reutilizable que haga que la IA escriba como tú en todas las sesiones futuras."
```
