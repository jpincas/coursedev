---
title: "Escritura con IA: Más allá de 'Escríbeme un informe'"
duration: "10m"
tags: [writing, documents, workflow]
---

# Escritura con IA: Más allá de "Escríbeme un informe"

"Escríbeme un informe sobre X" produce resultados mediocres. La escritura con IA ha madurado más allá de la generación con un solo prompt.

Una encuesta de Nature de 2025 reveló que el **57% de los científicos** ahora utilizan ayuda de IA para escribir. La técnica importa.

## El Método de Redacción con Investigación Previa

**No saltes directamente a escribir. Y nunca pidas un documento largo completo en un solo prompt.**

El enfoque más fiable:

1. **Empieza con una investigación profunda.** Antes de delinear, antes de escribir, pide a la IA que investigue el tema a fondo. Proporciónale tu contexto y pídele que cree un breve de investigación — escrito a un archivo. Esto fundamenta todo lo que sigue en hallazgos específicos en lugar de conocimiento genérico.
2. **Crea un esquema detallado.** Sobre la base de la investigación, enumera secciones, subsecciones, puntos clave. Escribe el esquema también en un archivo. Obtén sugerencias de la IA pero mantén tú la estructura.
3. **Expande cada sección individualmente.** Un prompt por sección, con restricciones específicas: número de palabras, estilo, requisitos de datos.
4. **Une y suaviza las transiciones.** Conecta las secciones, ajusta el flujo, asegura la coherencia.

El paso de investigación es el que la mayoría de la gente salta — y es el que marca la mayor diferencia. Una propuesta fundamentada en un breve de investigación exhaustivo se lee completamente diferente a una donde la IA solo genera texto que suena plausible.

## El Flujo de Trabajo Humano-IA Óptimo

La investigación de Stanford HAI sobre 1.440 historias encontró que la colaboración con LLM aumentó la productividad y redujo errores. Pero una colaboración extensa con IA redujo el sentido de propiedad de los escritores.

![Flujo de Trabajo de Escritura Humano-IA](/content/module-writing/images/writing-workflow.svg)

El patrón que funciona:

- **El humano define** el tema y el objetivo
- **La IA genera** ideas y esquemas
- **El humano selecciona** el enfoque
- **La IA crea** borradores sección por sección
- **El humano identifica** secciones débiles con feedback específico
- **La IA revisa** secciones concretas
- **La IA revisa** gramática y estilo
- **El humano asegura** la voz y la precisión
- **El humano verifica** todas las afirmaciones

La IA acelera. Los humanos son los responsables.

```callout
type: warning
title: "El Problema de la Voz"
content: "Un editor de FoxPrint Editorial advirtió que la revisión de un escritor talentoso quedó 'desprovista de voz, una sopa insípida de frases bonitas y excesivamente familiares' después del pulido con IA. Usa la IA para la mecánica, no para la voz."
```

## Dónde encaja la IA en la Escritura

**La IA destaca en:**
- Generar alternativas de redacción
- Expandir puntos en párrafos
- Revisar gramática y claridad
- Reorganizar contenido
- Crear primeros borradores a partir de esquemas

**La IA tiene dificultades con:**
- La voz personal auténtica
- El contexto cultural sutil
- La argumentación original
- El pensamiento estratégico sobre la audiencia
- Saber qué enfatizar

El escritor que usa bien la IA sabe distinguir cuáles son cuáles.

## Redacción con Investigación Previa en la Práctica

Observa el proceso completo desde el inicio hasta el final: investigación, esquema, redacción de cada sección y finalmente la reunión del documento. Presta atención al explorador de archivos — observa cómo el espacio de trabajo crece desde un único breve hasta un proyecto organizado con un producto final pulido.

```agent
id: sectional-drafting-demo
title: "Creación de Documentos con Investigación Previa"
model_label: "Claude"

system: |
  Eres un escritor e investigador de negocios. Cuando se te pida investigar,
  sé exhaustivo y específico. Crea documentos profesionales sección
  por sección. Sigue el esquema y las restricciones con precisión.

scratchpad:
  "project-brief.txt": |
    Proyecto: Rediseño del Portal de Clientes
    Plazo: T2 2026 (abril-junio)
    Presupuesto: £85.000
    Equipo: 2 diseñadores, 3 desarrolladores, 1 PM
    Objetivo: Reducir las incidencias de soporte un 30% mediante mejoras de autoatención
    Problemas principales:
    - El portal actual tiene 5 años y no es apto para móviles
    - El 60% de las incidencias son preguntas de "cómo encuentro X"
    - Puntuación de satisfacción del cliente del portal: 2,8/5
    Interlocutores: VP de Producto, Jefe de Soporte, CTO

tools:
  - scratchpad_read
  - scratchpad_write
  - create_folder
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
    text: "Paso 1: Investigación profunda. Antes de delinear, antes de escribir una sola palabra, le damos a la IA nuestro contexto y le pedimos que investigue el problema. La investigación se escribe a un archivo — no se queda en el chat."

  - type: user
    content: "Lee el breve del proyecto. Antes de escribir nada, necesito que investigues esto a fondo. ¿Qué hace que los rediseños de portales tengan éxito o fracasen? ¿Qué realmente reduce las incidencias de soporte? ¿Qué deberíamos priorizar dados nuestros problemas específicos? Crea una carpeta de investigación y escribe un breve de investigación exhaustivo allí."

  - type: tool_call
    tool: scratchpad_read
    args:
      filename: "project-brief.txt"

  - type: tool_result
    tool: scratchpad_read
    content: |
      Proyecto: Rediseño del Portal de Clientes
      Plazo: T2 2026 (abril-junio)
      Presupuesto: £85.000
      Equipo: 2 diseñadores, 3 desarrolladores, 1 PM
      Objetivo: Reducir las incidencias de soporte un 30% mediante mejoras de autoatención
      Problemas principales:
      - El portal actual tiene 5 años y no es apto para móviles
      - El 60% de las incidencias son preguntas de "cómo encuentro X"
      - Puntuación de satisfacción del cliente del portal: 2,8/5
      Interlocutores: VP de Producto, Jefe de Soporte, CTO

  - type: tool_call
    tool: create_folder
    args:
      path: "research"

  - type: tool_result
    tool: create_folder
    content: "Carpeta creada: research"

  - type: tool_call
    tool: scratchpad_write
    args:
      filename: "research/portal-research.md"
      content: |
        # Rediseño del Portal de Clientes: Brief de Investigación

        ## Qué Reduce las Incidencias de Soporte

        El mayor palanca para las incidencias de "cómo encuentro X" (vuestro 60%) es **el rediseño de búsqueda y navegación**. Los datos del sector muestran que implementar una búsqueda inteligente con sugerencias automáticas y resultados contextuales normalmente desvía el 35-50% de las incidencias de localización. Solo esto podría reducir el volumen total de incidencias entre un 20-30%.

        Segundo mayor impacto: **flujos de autoatención guiados** para tareas habituales (restablecimiento de contraseñas, cambios de cuenta, consultas de facturación). Las empresas que añaden asistentes paso a paso para sus 5 solicitudes más frecuentes ven una reducción total de incidencias del 15-25%.

        Combinados, estos dos cambios alcanzan realísticamente vuestro objetivo del 30% sin necesitar una integración de IA compleja.

        ## Por Qué Fracasan los Rediseños de Portales

        El modo de fracaso más común es **rediseñar todo a la vez**. Los portales que se lanzan como renovaciones completas tienen un 60% de probabilidad de incumplir los plazos y un alto riesgo de romper flujos de trabajo existentes que los clientes usan.

        El segundo modo de fracaso es **diseñar primero para escritorio** y luego adaptar para móvil. Dado que la mayoría del tráfico de portales B2B ahora es móvil, esto garantiza una experiencia deficiente para la mayoría de usuarios.

        ## Prioridades para Vuestra Situación

        **Mayor impacto (hacer primero):**
        - Renovación de la búsqueda con sugerencias automáticas — ataca directamente el problema del 60% de "cómo encuentro X"
        - Rediseño responsive mobile-first — aborda la brecha móvil de hace 5 años

        **Impacto medio (hacer segundo):**
        - Flujos de autoatención para las 3 categorías principales de incidencias
        - Ayudas contextuales (tooltips) en páginas complejas

        **Menor prioridad (Fase 2):**
        - Renovación visual y actualización de marca
        - Panel de analíticas avanzado
        - Integración de chatbot IA (rendimientos decrecientes una vez arreglada la búsqueda)

        ## Consideraciones de Presupuesto
        Con £85K, un equipo de 6 personas y 3 meses, el presupuesto es ajustado pero viable si el alcance se mantiene disciplinado. La renovación de búsqueda y el rediseño mobile juntos son alcanzables en T2. Los flujos de autoatención podrían necesitar extenderse a principios del T3 dependiendo de la complejidad.

        ## Idea Clave
        Vuestras incidencias de "cómo encuentro X" al 60% son una oportunidad — apuntan directamente a la solución. La mayoría de rediseños de portales malgastan meses debatiendo funcionalidades. Ya sabéis que el problema es la localización. Liderad con la búsqueda.

  - type: tool_result
    tool: scratchpad_write
    content: "Escrito en research/portal-research.md (38 líneas)"

  - type: assistant
    content: "Brief de investigación escrito en research/portal-research.md. Hallazgo clave: vuestro 60% de incidencias de \"cómo encuentro X\" apuntan directamente a la solución — el rediseño de búsqueda y navegación por sí solo podría alcanzar vuestro objetivo de reducción del 30%. Recomiendo liderar con eso en lugar de una renovación completa, que es el modo de fracaso más común en los rediseños de portales."

  - type: note
    text: "Mirad el explorador de archivos: ahora hay una carpeta de investigación con un brief exhaustivo fundamentado en datos específicos del proyecto. No consejos genéricos sobre portales — hallazgos reales sobre qué reduce incidencias PARA vuestro problema. Este archivo de investigación se convierte en la base de todo lo que construyamos a continuación."

  - type: user
    content: "Buena investigación. Ahora crea un esquema para la propuesta basado en esos hallazgos. Es para el VP de Producto, el Jefe de Soporte y el CTO — les importa el impacto y el riesgo. Crea una carpeta de borradores y escribe el esquema allí."

  - type: tool_call
    tool: scratchpad_read
    args:
      filename: "research/portal-research.md"

  - type: tool_result
    tool: scratchpad_read
    content: |
      # Rediseño del Portal de Clientes: Brief de Investigación
      ## Qué Reduce las Incidencias de Soporte
      El mayor palanca para las incidencias de "cómo encuentro X" (vuestro 60%) es el rediseño de búsqueda y navegación...

  - type: tool_call
    tool: create_folder
    args:
      path: "drafts"

  - type: tool_result
    tool: create_folder
    content: "Carpeta creada: drafts"

  - type: tool_call
    tool: scratchpad_write
    args:
      filename: "drafts/proposal-outline.md"
      content: |
        # Rediseño del Portal de Clientes — Esquema de la Propuesta

        ## 1. Resumen Ejecutivo
        Un párrafo: el problema de localización, la solución centrada en búsqueda, la solicitud.

        ## 2. El Problema
        Estado actual: portal de 5 años, el 60% de incidencias son "cómo encuentro X", satisfacción de 2,8/5, £180K/año en costes de soporte evitables. Liderad con el coste de no actuar.

        ## 3. Qué Muestra la Investigación
        El rediseño de búsqueda y navegación es la palanca de mayor impacto. Los flujos de autoatención guiados, segundo. Las renovaciones completas son el modo de fracaso más común — argumentad a favor de un alcance focalizado.

        ## 4. Enfoque Propuesto
        Fase 1 (abril-mayo): Renovación de búsqueda + rediseño mobile-first — ataca el 60% directamente.
        Fase 2 (junio): Flujos de autoatención para las 3 categorías principales de incidencias.
        No una reconstrucción completa — soluciones focalizadas en las áreas de mayor impacto.

        ## 5. Impacto Esperado
        Reducción del 30% en incidencias (realista según la investigación). Mejora de CSAT de 2,8 a 3,5+. Ahorro de £54K/año en costes de soporte.

        ## 6. Plazo, Presupuesto y Riesgo
        £85K en T2. Mitigación de riesgos: enfoque por fases, no todo de golpe. El mayor riesgo es la deriva de alcance — detallad lo que NO estamos haciendo.

        ## 7. Próximos Pasos
        Aprobación para proceder. Inicio del diseño sprint la primera semana de abril.

  - type: tool_result
    tool: scratchpad_write
    content: "Escrito en drafts/proposal-outline.md (22 líneas)"

  - type: assistant
    content: "Esquema escrito en drafts/proposal-outline.md. Siete secciones, estructuradas alrededor de los hallazgos de la investigación. Fijáos en la sección 3 — \"Qué Muestra la Investigación\" — que da a la propuesta una base de evidencias que la mayoría de propuestas internas no tienen. Listo para empezar a redactar las secciones."

  - type: note
    text: "Paso 2 completo. El esquema es un archivo en la carpeta de borradores, fundamentado en la investigación. La estructura de la propuesta — enfoque por fases, no renovación completa — proviene directamente del hallazgo de que las renovaciones completas son el modo de fracaso más común. Ahora redactamos cada sección, una a una."

  - type: user
    content: "Empieza a redactar. Escribe las secciones 1 y 2 — el Resumen Ejecutivo y El Problema. Mantened cada una concisa y específica. Usad los datos de la investigación. Escribidlas en la carpeta de borradores."

  - type: tool_call
    tool: scratchpad_write
    args:
      filename: "drafts/01-executive-summary.md"
      content: |
        ## 1. Resumen Ejecutivo

        Nuestro portal de clientes nos cuesta £180K al año en incidencias de soporte evitables. El 60% de todas las incidencias son preguntas de \"cómo encuentro X\" que un portal bien diseñado debería gestionar sin intervención humana. Proponemos un rediseño focalizado en dos fases — renovación de búsqueda y reconstrucción mobile-first en la Fase 1, flujos de autoatención en la Fase 2 — entregable dentro del T2 2026 por £85K. Resultado objetivo: reducción del 30% en incidencias y satisfacción del cliente por encima de 3,5/5.

  - type: tool_result
    tool: scratchpad_write
    content: "Escrito en drafts/01-executive-summary.md (4 líneas)"

  - type: tool_call
    tool: scratchpad_write
    args:
      filename: "drafts/02-the-problem.md"
      content: |
        ## 2. El Problema

        Nuestro portal de clientes tiene cinco años y no fue diseñado para móviles. La mayoría de nuestros usuarios ahora acceden desde teléfonos, y la experiencia lo demuestra.

        Los números cuentan la historia. El 60% de nuestras incidencias de soporte son variaciones de \"cómo encuentro X\" — preguntas que el portal debería responder sin intervención humana. La satisfacción del portal se sitúa en 2,8 de 5, la más baja de cualquier punto de contacto con el cliente.

        Cada incidencia cuesta aproximadamente £12 en tiempo de personal. Con el volumen actual, los problemas de localización por sí solos cuestan unos £180K al año. El portal no solo está obsoleto — está activamente aumentando los costes de soporte y reduciendo la satisfacción del cliente.

  - type: tool_result
    tool: scratchpad_write
    content: "Escrito en drafts/02-the-problem.md (8 líneas)"

  - type: assistant
    content: "Secciones 1 y 2 escritas. El resumen ejecutivo es un párrafo contundente — problema, solución, solicitud — con la cifra de costes de £180K al inicio. La sección 2 expande con los tres puntos de datos y termina con el argumento de costes. Listo para las siguientes secciones."

  - type: user
    content: "Secciones 3 y 4 — Qué Muestra la Investigación y Enfoque Propuesto. La sección 3 debe dar confianza a la dirección de que hay evidencia detrás del plan. La sección 4 debe ser concreta."

  - type: tool_call
    tool: scratchpad_write
    args:
      filename: "drafts/03-what-research-shows.md"
      content: |
        ## 3. Qué Muestra la Investigación

        Investigué qué reduce las incidencias en rediseños de portales. Los hallazgos han moldeado esta propuesta.

        **La búsqueda y la navegación son la palanca de mayor impacto.** Las empresas que implementan búsqueda inteligente con sugerencias automáticas y resultados contextuales normalmente desvían el 35-50% de las incidencias de localización. Dado que la localización es nuestro problema dominante, ahí es donde debemos liderar.

        **Los flujos de autoatención son la siguiente palanca.** Los asistentes paso a paso guiados para tareas habituales (restablecimiento de contraseñas, cambios de cuenta, facturación) reducen el volumen total de incidencias un 15-25% adicional.

        **Las renovaciones completas son el modo de fracaso más común.** Los proyectos de portal que intentan rediseñar todo a la vez tienen un 60% de probabilidad de incumplir plazos y corren riesgo de romper flujos de trabajo en los que los clientes confían. En cambio, proponemos un enfoque focalizado y por fases.

  - type: tool_result
    tool: scratchpad_write
    content: "Escrito en drafts/03-what-research-shows.md (10 líneas)"

  - type: tool_call
    tool: scratchpad_write
    args:
      filename: "drafts/04-proposed-approach.md"
      content: |
        ## 4. Enfoque Propuesto

        **Fase 1 (abril-mayo): Búsqueda y Móvil**
        Renovar la búsqueda del portal con sugerencias automáticas, resultados contextuales y navegación mejorada. Simultáneamente reconstruir el portal con diseño responsive mobile-first. Estos dos cambios atacan directamente el problema de localización del 60%.

        **Fase 2 (junio): Flujos de Autoatención**
        Añadir asistentes guiados para las tres categorías principales de incidencias de soporte. Flujos paso a paso que resuelven problemas habituales sin intervención de agentes.

        **Lo que NO estamos haciendo:** No es una reconstrucción completa del portal. No estamos rediseñando la identidad visual, construyendo un panel de analíticas ni añadiendo un chatbot IA. Esos podrían venir después, pero la investigación muestra que la localización y el móvil son donde está el impacto. Nos mantenemos focalizados.

  - type: tool_result
    tool: scratchpad_write
    content: "Escrito en drafts/04-proposed-approach.md (10 líneas)"

  - type: assistant
    content: "Secciones 3 y 4 escritas. La sección 3 presenta las evidencias — tres hallazgos que justifican el enfoque. La sección 4 expone las dos fases e, importantemente, establece lo que no estamos haciendo. Ese párrafo de \"no hacerlo\" anticipa la conversación sobre deriva de alcance."

  - type: note
    text: "Fijáos en cómo el usuario da diferentes restricciones para cada sección. La sección 2 fue \"persuasiva para un CTO\" — así que lideró con costes. La sección 3 fue \"dar confianza a la dirección\" — así que lideró con evidencia. La sección 4 fue \"sé concreto\" — así que tiene fases específicas y un límite de alcance explícito. Por eso funciona la redacción por secciones: cada sección recibe instrucciones adaptadas a su propósito."

  - type: user
    content: "Secciones 5, 6 y 7 — Impacto Esperado, Plazo y Presupuesto, y Próximos Pasos. Mantenédlos concisos. Deben cerrar la propuesta."

  - type: tool_call
    tool: scratchpad_write
    args:
      filename: "drafts/05-expected-impact.md"
      content: |
        ## 5. Impacto Esperado

        Basado en la investigación y nuestros datos específicos de incidencias:

        - **Volumen de incidencias:** 3.400 → ~2.400/mes (reducción del 30%)
        - **Tiempo de primera respuesta:** La cola reducida significa respuesta más rápida en las incidencias restantes
        - **Satisfacción del portal:** 2,8 → 3,5+ de 5
        - **Ahorro anual en costes de soporte:** Aproximadamente £54K
        - **Capacidad de los agentes:** Liberados de consultas repetitivas de localización para centrarse en interacciones complejas y de alto valor con el cliente

  - type: tool_result
    tool: scratchpad_write
    content: "Escrito en drafts/05-expected-impact.md (9 líneas)"

  - type: tool_call
    tool: scratchpad_write
    args:
      filename: "drafts/06-timeline-budget-risk.md"
      content: |
        ## 6. Plazo, Presupuesto y Riesgo

        **Presupuesto:** £85.000 | **Plazo:** T2 2026 | **Equipo:** 2 diseñadores, 3 desarrolladores, 1 PM

        - **Abril:** Diseño sprint, arquitectura de búsqueda, marco responsive mobile
        - **Mayo:** Construcción, pruebas de integración, piloto interno
        - **Junio:** Flujos de autoatención, implementación por fases, monitoreo

        **Mitigación de riesgos:** El enfoque por fases es la mitigación de riesgos. La Fase 1 entrega mejoras de búsqueda independientemente del plazo de la Fase 2. Si los flujos de autoatención requieren más iteración, se extienden a principios del T3 sin bloquear la entrega principal. El mayor riesgo es la deriva de alcance — esta propuesta define explícitamente lo que no estamos construyendo.

  - type: tool_result
    tool: scratchpad_write
    content: "Escrito en drafts/06-timeline-budget-risk.md (9 líneas)"

  - type: tool_call
    tool: scratchpad_write
    args:
      filename: "drafts/07-next-steps.md"
      content: |
        ## 7. Próximos Pasos

        1. **Aprobar esta propuesta** para proceder con el plan de dos fases
        2. **Inicio del diseño sprint:** Primera semana de abril
        3. **Revisión de interlocutores de la Fase 1:** Final de mayo — demo de mejoras de búsqueda y rediseño móvil
        4. **Implementación completa:** Final de junio
        5. **Medición del impacto:** Revisión a 90 días comparando volumen de incidencias, tiempos de respuesta y CSAT frente a las líneas base actuales

  - type: tool_result
    tool: scratchpad_write
    content: "Escrito en drafts/07-next-steps.md (7 líneas)"

  - type: assistant
    content: "Las siete secciones redactadas. La sección de impacto lidera con los números que interesan a la dirección. El plazo da un desglose claro mes a mes. Los Próximos Pasos terminan con una solicitud concreta y un plan de medición — para que la dirección sepa exactamente qué está aprobando y cómo se medirá el éxito."

  - type: note
    text: "Mirad la carpeta de borradores en el explorador. Siete archivos de sección, cada uno escrito con restricciones específicas para su propósito. Cada sección se basa en la investigación — los números, el enfoque por fases, los límites de alcance. Ahora el paso final: montaje."

  - type: user
    content: "Ahora reúne todo. Lee todos los borradores de sección y monta la propuesta final — un documento limpio con transiciones suaves entre secciones. Crea una carpeta final y escríbela allí."

  - type: tool_call
    tool: list_files
    args: {}

  - type: tool_result
    tool: list_files
    content: |
      3 carpetas, 9 archivos:
      - project-brief.txt
      - drafts/proposal-outline.md
      - drafts/01-executive-summary.md
      - drafts/02-the-problem.md
      - drafts/03-what-research-shows.md
      - drafts/04-proposed-approach.md
      - drafts/05-expected-impact.md
      - drafts/06-timeline-budget-risk.md
      - drafts/07-next-steps.md
      - research/portal-research.md

  - type: tool_call
    tool: scratchpad_read
    args:
      filename: "drafts/01-executive-summary.md"

  - type: tool_result
    tool: scratchpad_read
    content: |
      ## 1. Resumen Ejecutivo
      Nuestro portal de clientes nos cuesta £180K al año en incidencias de soporte evitables...

  - type: tool_call
    tool: scratchpad_read
    args:
      filename: "drafts/02-the-problem.md"

  - type: tool_result
    tool: scratchpad_read
    content: |
      ## 2. El Problema
      Nuestro portal de clientes tiene cinco años y no fue diseñado para móviles...

  - type: tool_call
    tool: scratchpad_read
    args:
      filename: "drafts/03-what-research-shows.md"

  - type: tool_result
    tool: scratchpad_read
    content: |
      ## 3. Qué Muestra la Investigación
      Investigué qué reduce las incidencias en rediseños de portales...

  - type: tool_call
    tool: scratchpad_read
    args:
      filename: "drafts/04-proposed-approach.md"

  - type: tool_result
    tool: scratchpad_read
    content: |
      ## 4. Enfoque Propuesto
      Fase 1 (abril-mayo): Búsqueda y Móvil...

  - type: tool_call
    tool: scratchpad_read
    args:
      filename: "drafts/05-expected-impact.md"

  - type: tool_result
    tool: scratchpad_read
    content: |
      ## 5. Impacto Esperado
      Basado en la investigación y nuestros datos específicos de incidencias...

  - type: tool_call
    tool: scratchpad_read
    args:
      filename: "drafts/06-timeline-budget-risk.md"

  - type: tool_result
    tool: scratchpad_read
    content: |
      ## 6. Plazo, Presupuesto y Riesgo
      Presupuesto: £85.000 | Plazo: T2 2026...

  - type: tool_call
    tool: scratchpad_read
    args:
      filename: "drafts/07-next-steps.md"

  - type: tool_result
    tool: scratchpad_read
    content: |
      ## 7. Próximos Pasos
      1. Aprobar esta propuesta para proceder...

  - type: tool_call
    tool: create_folder
    args:
      path: "final"

  - type: tool_result
    tool: create_folder
    content: "Carpeta creada: final"

  - type: tool_call
    tool: scratchpad_write
    args:
      filename: "final/customer-portal-proposal.md"
      content: |
        # Rediseño del Portal de Clientes — Propuesta
        **Preparado para:** VP de Producto, Jefe de Soporte, CTO
        **Fecha:** Febrero 2026

        ---

        ## Resumen Ejecutivo

        Nuestro portal de clientes nos cuesta £180K al año en incidencias de soporte evitables. El 60% de todas las incidencias son preguntas de "cómo encuentro X" que un portal bien diseñado debería gestionar sin intervención humana. Proponemos un rediseño focalizado en dos fases — renovación de búsqueda y reconstrucción mobile-first en la Fase 1, flujos de autoatención en la Fase 2 — entregable dentro del T2 2026 por £85K. Resultado objetivo: reducción del 30% en incidencias y satisfacción del cliente por encima de 3,5/5.

        ## El Problema

        Nuestro portal de clientes tiene cinco años y no fue diseñado para móviles. La mayoría de nuestros usuarios ahora acceden desde teléfonos, y la experiencia lo demuestra.

        Los números cuentan la historia. El 60% de nuestras incidencias de soporte son variaciones de "cómo encuentro X" — preguntas que el portal debería responder sin intervención humana. La satisfacción del portal se sitúa en 2,8 de 5, la más baja de cualquier punto de contacto con el cliente.

        Cada incidencia cuesta aproximadamente £12 en tiempo de personal. Con el volumen actual, los problemas de localización por sí solos cuestan unos £180K al año. El portal no solo está obsoleto — está activamente aumentando los costes de soporte y reduciendo la satisfacción del cliente.

        ## Qué Muestra la Investigación

        Investigué qué reduce las incidencias en rediseños de portales. Tres hallazgos han moldeado esta propuesta.

        Primero, **la búsqueda y la navegación son la palanca de mayor impacto.** Las empresas que implementan búsqueda inteligente con sugerencias automáticas y resultados contextuales normalmente desvían el 35-50% de las incidencias de localización. Dado que la localización es nuestro problema dominante, ahí es donde lideramos.

        Segundo, **los flujos de autoatención ofrecen la siguiente capa de impacto.** Los asistentes paso a paso guiados para tareas habituales reducen el volumen total de incidencias un 15-25% adicional.

        Tercero, **las renovaciones completas son el modo de fracaso más común.** Los proyectos de portal que intentan rediseñar todo a la vez tienen un 60% de probabilidad de incumplir plazos. En cambio, proponemos un enfoque focalizado y por fases.

        ## Enfoque Propuesto

        **Fase 1 (abril-mayo): Búsqueda y Móvil**
        Renovar la búsqueda del portal con sugerencias automáticas, resultados contextuales y navegación mejorada. Simultáneamente reconstruir el portal con diseño responsive mobile-first. Estos dos cambios atacan directamente el problema de localización del 60%.

        **Fase 2 (junio): Flujos de Autoatención**
        Añadir asistentes guiados para las tres categorías principales de incidencias de soporte. Flujos paso a paso que resuelven problemas habituales sin intervención de agentes.

        **Lo que no estamos haciendo:** No es una reconstrucción completa del portal. No estamos rediseñando la identidad visual, construyendo un panel de analíticas ni añadiendo un chatbot IA. La investigación muestra que la localización y el móvil son donde está el impacto. Nos mantenemos focalizados.

        ## Impacto Esperado

        Basado en la investigación y nuestros datos específicos de incidencias:

        - **Volumen de incidencias:** 3.400 → ~2.400/mes (reducción del 30%)
        - **Tiempo de primera respuesta:** La cola reducida significa respuesta más rápida en las incidencias restantes
        - **Satisfacción del portal:** 2,8 → 3,5+ de 5
        - **Ahorro anual en costes de soporte:** Aproximadamente £54K
        - **Capacidad de los agentes:** Liberados de consultas repetitivas de localización para centrarse en interacciones complejas y de alto valor

        ## Plazo, Presupuesto y Riesgo

        **Presupuesto:** £85.000 | **Plazo:** T2 2026 | **Equipo:** 2 diseñadores, 3 desarrolladores, 1 PM

        - **Abril:** Diseño sprint, arquitectura de búsqueda, marco responsive mobile
        - **Mayo:** Construcción, pruebas de integración, piloto interno
        - **Junio:** Flujos de autoatención, implementación por fases, monitoreo

        El enfoque por fases es la mitigación de riesgos. La Fase 1 entrega mejoras de búsqueda independientemente del plazo de la Fase 2. Si los flujos de autoatención requieren más iteración, se extienden a principios del T3 sin bloquear la entrega principal. El mayor riesgo es la deriva de alcance — esta propuesta define explícitamente lo que no estamos construyendo.

        ## Próximos Pasos

        1. Aprobar esta propuesta para proceder con el plan de dos fases
        2. Inicio del diseño sprint: primera semana de abril
        3. Revisión de interlocutores de la Fase 1: final de mayo
        4. Implementación completa: final de junio
        5. Medición del impacto: revisión a 90 días comparando volumen de incidencias, tiempos de respuesta y CSAT frente a las líneas base actuales

  - type: tool_result
    tool: scratchpad_write
    content: "Escrito en final/customer-portal-proposal.md (62 líneas)"

  - type: assistant
    content: "Propuesta final montada y escrita en final/customer-portal-proposal.md. Las siete secciones pulidas en un único documento coherente con página de título, transiciones suaves entre secciones y formato consistente en todo el documento. Lista para compartir con los interlocutores."

  - type: note
    text: "Mirad el explorador de archivos ahora. Desde un único breve del proyecto, construimos un espacio de trabajo estructurado:\n\n• research/ — la investigación profunda que fundamentó todo\n• drafts/ — el esquema y los siete borradores de sección individuales\n• final/ — la propuesta pulida y montada\n\nEl documento final se lee como un todo coherente porque cada sección fue redactada con restricciones específicas y cada afirmación se remonta a la investigación. El propio espacio de trabajo es reutilizable — si los interlocutores quieren cambios, editáis el borrador de la sección relevante y reasamáis."

  - type: note
    text: "Este es el método completo de redacción con investigación previa. Investigación → esquema → redacción sección por sección → montaje. Cada paso produjo un archivo. Cada archivo se construyó sobre el anterior. La propuesta final tiene evidencias, especificidad y un argumento claro — porque empezamos con investigación, no con \"escríbeme una propuesta.\""
```

```quiz
id: sectional-drafting-quiz
type: multiple-choice
question: "¿Cuál es el mayor beneficio de empezar con un paso de investigación profunda antes de escribir un documento?"
options:
  - "Hace que la IA genere más palabras para cumplir los requisitos de longitud"
  - "Da a cada paso siguiente — esquema, redacción, revisión — una base de evidencia específica en lugar de conocimiento genérico"
  - "Elimina la necesidad de revisar la salida de la IA ya que la investigación valida todo"
answer: 1
explanation: "El paso de investigación fundamenta todo lo que sigue. El esquema tiene una base de evidencias. Los borradores de sección se basan en hallazgos específicos. Sin investigación, la IA rellena los huecos con texto plausible pero genérico. Con investigación, cada párrafo se ancla en algo específico."
```
