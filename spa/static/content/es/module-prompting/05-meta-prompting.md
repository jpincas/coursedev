---
title: "Meta-Prompting"
duration: "10m"
tags: [meta-prompting, iteración, técnicas]
---

# Meta-Prompting

Aquí hay una técnica poderosa: en lugar de escribir tú la instrucción perfecta, pídele ayuda a la IA para escribirla.

## El patrón básico

```
Quiero [objetivo]. Ayúdame a escribir una instrucción detallada
que obtenga el mejor resultado.
Hazme preguntas de aclaración primero.
```

La IA te preguntará lo que necesita saber. Tus respuestas se convierten en la instrucción.

Esto es especialmente útil cuando no estás seguro de lo que necesitas.

## Meta-Prompting en acción

**Tú:** "Necesito analizar comentarios de clientes. Ayúdame a escribir una instrucción para eso."

**IA:** "Permíteme hacer unas preguntas:
- ¿En qué formato están los comentarios?
- ¿Qué insights específicos estás buscando?
- ¿Cómo debería estructurarse el análisis?
- ¿Quién leerá los resultados?"

**Tú:** "Están en un CSV. Busco quejas comunes. Necesito un resumen para el equipo de producto."

**IA:** "Aquí está tu instrucción: 'Analiza este CSV de comentarios de clientes. Identifica las 5 principales categorías de quejas con conteos de frecuencia. Para cada una, proporciona 2-3 citas representativas. Formatea como un breve reporte para el equipo de producto, máximo 1 página.'"

```callout
type: tip
title: "Por qué funciona"
content: "Las preguntas de la IA revelan qué información se necesita. Cada respuesta añade especificidad. La instrucción final suele ser mejor que la que habrías escrito desde cero."
```

## Cuándo usar meta-prompting

**No estás seguro del alcance**
Deja que las preguntas de la IA te ayuden a descubrir qué necesitas realmente.

**Tareas complejas o multifacéticas**
El proceso de preguntas saca a la luz aspectos que quizás no habías considerado.

**Dominios desconocidos**
La IA puede hacer preguntas de nivel experto sobre temas donde tú no eres el experto.

**Instrucciones reutilizables**
Si harás tareas similares repetidamente, invierte en obtener una gran instrucción.

## La idea clave

No describas el proceso. Describe el resultado.

**Mal:**
"Primero analiza los datos, luego identifica tendencias, luego resúmelos, luego formatea como viñetas, luego..."

**Bien:**
"Produce un análisis de tendencias con las 5 principales insights como viñetas."

La IA es capaz de figuring out el proceso. Ese es su trabajo. Tu trabajo es definir cómo se ve terminado.

Este es el cambio de prompting a delegación.

## Especificación interactiva

El meta-prompting llevado a su conclusión lógica para entregables complejos es la **especificación interactiva** — usar la IA para ayudarte a definir los detalles de un proyecto o documento mediante conversación iterativa antes de ejecutar.

**El patrón:**

1. Describe lo que quieres a alto nivel: "Necesito un plan de formación para nuevos managers"
2. La IA hace preguntas de aclaración sobre alcance, audiencia, restricciones
3. Tú respondes, añadiendo detalle y refinando
4. La IA propone una especificación u outline estructurada en base a tus respuestas
5. Revisas y ajustas la especificación: "Mueve la sección 3 antes de la sección 2, añade un módulo sobre conversaciones difíciles"
6. Una vez acordada la especificación, le pides a la IA ejecutar contra ella

**Por qué funciona:** Para entregables complejos, saltar directamente a la ejecución produce resultados mediocres porque los requisitos estaban sub-especificados. La especificación interactiva te obliga a pensar en los requisitos antes de que cualquier trabajo ocurra. Las preguntas de la IA sacan a la luz aspectos que no habrías considerado.

**Ejemplo de flujo:**
```
Tú: Necesito una propuesta de cliente de 30 páginas para un proyecto de transformación digital.
IA: ¿Qué industria? ¿Cuál es el estado actual del cliente? ¿Rango de presupuesto? ¿Plazo?
Tú: Servicios financieros, sistemas legacy, ~$2M, 18 meses.
IA: Aquí hay una estructura propuesta: [10 secciones con descripciones]
Tú: Bien, pero añade una sección sobre cumplimiento regulatorio y elimina el genérico 'sobre nosotros.'
IA: Estructura actualizada: [revisada]. ¿Debería redactar la propuesta completa contra esta especificación?
Tú: Sí, adelante.
```

La especificación se convierte en un contrato entre tú y la IA. El resultado es dramáticamente mejor que una sola instrucción monolítica.

```callout
type: tip
title: "Cuándo especificar vs cuándo preguntar directamente"
content: "Usa especificación interactiva para cualquier cosa que tome más de una página: propuestas, reportes, planes de formación, planes de proyecto. Para tareas rápidas (emails, resúmenes, análisis cortos), simplemente pregunta directamente — especificar sería excesivo."
```

## Especificación interactiva en acción

Observa a alguien usar la IA para definir los detalles de un proyecto complejo mediante conversación. El usuario tiene un objetivo vago y contexto mínimo — las preguntas de la IA sacan los detalles.

```agent
id: interactive-speccing-demo
title: "Especificación interactiva: Retreat de empresa"
model_label: "Claude"

system: |
  Eres un consultor de planificación de eventos. Cuando ayudas a clientes a planificar
  retreats empresariales, haz preguntas de aclaración para entender sus objetivos,
  restricciones y experiencias previas. Crea especificaciones detalladas y
  estructuradas que cubran todos los aspectos del evento.

scratchpad:
  "context/company-info.md": |
    # Vortex Technologies

    Empresa tecnológica, 50 empleados, totalmente remote-first desde 2020.
    Sede: Londres (oficina registrada solo, mayormente vacía)
    Distribución del equipo: UK (30), EU (15), Costa Este de EE.UU. (5)

    Cultura orientada a ingeniería. Mix de personal junior y senior.
    Edad promedio: 32. Varias familias con niños pequeños.

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
    text: "El usuario tiene un objetivo vago: planificar un retreat de empresa. En lugar de intentar especificar todo de antemano, observa cómo le piden a la IA que ayude a construir una especificación adecuada mediante conversación."

  - type: user
    content: "Necesito planificar un retreat de empresa. En lugar de que yo intente pensar en todo, ¿puedes entrevistarme para construir una especificación adecuada? Hay algo de información básica de la empresa en la carpeta de contexto."

  - type: tool_call
    tool: scratchpad_read
    args:
      filename: "context/company-info.md"

  - type: tool_result
    tool: scratchpad_read
    content: |
      # Vortex Technologies

      Empresa tecnológica, 50 empleados, totalmente remote-first desde 2020.
      Sede: Londres (oficina registrada solo, mayormente vacía)
      Distribución del equipo: UK (30), EU (15), Costa Este de EE.UU. (5)

      Cultura orientada a ingeniería. Mix de personal junior y senior.
      Edad promedio: 32. Varias familias con niños pequeños.

  - type: assistant
    content: "Buen enfoque — haré preguntas para construir una especificación completa. Empezaré con lo fundamental:\n\n**Objetivos y Propósito:**\n1. ¿Cuáles son los objetivos principales para este retreat? (¿Cohesión de equipo? ¿Planificación estratégica? ¿Celebrar hitos? ¿Desarrollo de habilidades?)\n2. ¿Hay algo específico que quieras lograr que los retreats previos no hayan conseguido?\n\n**Presupuesto y Restricciones:**\n3. ¿Cuál es tu rango de presupuesto por persona?\n4. ¿Alguna restricción fija sobre fechas o ubicación?\n5. ¿Cuántos días estás pensando?\n\n**Experiencia previa:**\n6. ¿Han hecho un retreat antes? ¿Qué funcionó bien? ¿Qué no?"

  - type: user
    content: "El objetivo principal es la cohesión de equipo — al ser remote-first, la mayoría de la gente nunca se ha conocido en persona. También queremos hacer alguna planificación estratégica para el H2. El presupuesto es de alrededor de £1,200 por persona. Fechas: idealmente finales de abril o principios de mayo, antes de que empiecen las vacaciones de verano. Pensando en 2-3 días. Intentamos una cosa de un día el año pasado en Londres pero se sintió apresurada y al equipo de EE.UU. no le justificó volar solo por un día."

  - type: assistant
    content: "Eso es un contexto útil. Que el evento de un día en Londres fuera demasiado corto es una lección clave. Preguntas de seguimiento:\n\n**Ubicación y Logística:**\n7. Dada la distribución de tu equipo, ¿dónde tiene sentido? (¿Ubicación europea central para más fácil acceso desde las tres regiones?)\n8. Las 50 personas, o solo ciertos equipos/niveles de senioridad?\n9. ¿Cubrirás los costes de viaje, o eso es parte del presupuesto de £1,200/persona?\n\n**Actividades y Estructura:**\n10. ¿Qué balance quieres entre sesiones de trabajo y actividades sociales?\n11. ¿Alguna actividad que definitivamente deba o no deba incluirse?\n12. ¿Quieres facilitadores externos para las porciones de planificación estratégica?"

  - type: user
    content: "Probablemente en algún lugar de Europa fácil de alcanzar — ¿tal vez Barcelona, Ámsterdam o Lisboa? Deben venir las 50 personas. Las £1,200 son para alojamiento, comida y actividades — el viaje es aparte. Balance: diría 40% sesiones de trabajo, 60% social y cohesión de equipo. Nada de trust falls ni diversión corporacional forzada. Quizá algo opcional y aventurero para quienes quieran. No estoy seguro sobre facilitadores — ¿qué piensas tú?"

  - type: assistant
    content: "Para planificación estratégica con 50 personas, un facilitador externo suele ayudar — os saca del modo de ser gestores de reuniones y os convierte en participantes en lugar de eso. Recomiendo presupuestar para uno.\n\nUnos detalles más para definir la especificación:\n\n**Alojamiento y Comidas:**\n13. ¿Hotel o algo más informal como una villa/centro de retiro?\n14. ¿Alguna restricción alimentaria o preferencia que planificar?\n\n**Familias e Inclusión:**\n15. Tus notas mencionan que varias personas tienen niños pequeños — ¿permitir parejas/familias, o solo empleados?\n\n**Resultados:**\n16. ¿Quieres un resultado tangible de las sesiones de planificación estratégica? (¿Roadmap del H2 documentado? ¿OKRs? ¿O solo alineación?)"

  - type: user
    content: "Un hotel tiene sentido para 50 personas — logística más fácil. Tenemos vegetarianos y un par de personas con alergias, así que el venue necesita catering flexible. Solo empleados para este — queremos tiempo de trabajo enfocado. Y sí, debemos salir con un roadmap del H2 documentado y OKRs acordados. Ese es el producto de trabajo."

  - type: assistant
    content: "Perfecto — tengo lo que necesito. Dejadme redactar una especificación completa del retreat basada en todo lo que me habéis dicho."

  - type: tool_call
    tool: scratchpad_write
    args:
      filename: "offsite-spec.md"
      content: |
        # Retreat Primavera 2026 de Vortex Technologies
        ## Especificación del Evento

        ### Propósito y Objetivos
        **Objetivo principal:** Construir cohesión de equipo en un equipo totalmente remoto (la mayoría nunca se ha conocido en persona)

        **Objetivo secundario:** Desarrollar el roadmap estratégico y los OKRs del H2 de 2026

        **Lección clave de 2025:** El evento de un día en Londres fue demasiado corto y no justificó el viaje del equipo de EE.UU. Este evento debe ofrecer valor suficiente para un viaje internacional.

        ---

        ### Logística
        **Asistentes:** Las 50 empleados (UK: 30, EU: 15, Este de EE.UU.: 5)

        **Fechas:** Finales de abril o principios de mayo de 2026 (previo a vacaciones de verano)
        - Sugerido: 28-30 de abril (mar-jue) o 5-7 de mayo

        **Duración:** 3 días, 2 noches
        - Día 1: Llegadas PM, cena de bienvenida
        - Día 2: Día completo (sesiones estratégicas + actividades de equipo)
        - Día 3: Medio día, salidas PM

        **Ubicación:** Ciudad europea con buenas conexiones de vuelo internacionales
        - Lista corta: Barcelona, Ámsterdam, Lisboa
        - Requisitos: Vuelos directos desde Londres, principales ciudades EU, Costa Este de EE.UU.

        ---

        ### Presupuesto
        **Por persona:** £1,200 (alojamiento, comidas, actividades, facilitador)

        **Presupuesto total:** £60,000

        **Viaje:** Presupuesto separado (no incluido en las £1,200)

        ---

        ### Alojamiento
        **Tipo:** Hotel (4 estrellas, céntrico)

        **Requisitos:**
        - 50 habitaciones individuales (2 noches)
        - Espacio de conferencia/reunión para 50 personas
        - Salas de breakout para sesiones de grupos pequeños
        - Catering flexible (opciones vegetarianas, acomodaciones para alergias)
        - Equipamiento AV (proyector, pantallas, buen wifi)

        ---

        ### Estructura del Programa
        **Balance:** 40% sesiones de trabajo estratégico, 60% social/cohesión de equipo

        **Día 1 (Llegadas):**
        - 18:00: Recepción de bienvenida
        - 19:30: Cena (informal, toma de contacto)

        **Día 2 (Día completo):**
        - 09:00 - 12:30: Sesión de planificación estratégica 1 (prioridades H2)
          - Liderada por facilitador externo
          - Grupos de breakout por función
        - 12:30: Almuerzo
        - 14:00 - 16:00: Sesión de planificación estratégica 2 (definición de OKRs)
        - 16:30 - 19:00: Actividad de equipo (ver opciones abajo)
        - 20:00: Cena grupal

        **Día 3 (Medio día):**
        - 09:00 - 11:30: Sesión de planificación estratégica 3 (finalización del roadmap)
        - 11:30: Clausura y próximos pasos
        - 12:30: Almuerzo, luego salidas

        ---

        ### Actividades
        **Principio rector:** Ningún team-building corporacional forzado. Experiencias genuinas.

        **Opciones para la tarde del Día 2:**
        - Opción A (Aventurera): Experiencia de coasteering o via ferrata (opcional, máx. 20 personas)
        - Opción B (Cultural): Tour guiado de arquitectura/comida de la ciudad
        - Opción C (Relajada): Clase de cocina o cata de vinos

        Todas las opciones deben permitir conversación natural e intermezcla.

        ---

        ### Resultado de Planificación Estratégica
        **Entregables:**
        1. Roadmap del H2 de 2026 documentado (iniciativas clave, responsables, plazos)
        2. OKRs de la empresa para el H2 (objetivos y resultados clave)
        3. Mapa de dependencias cross-funcionales

        **Brief del facilitador:**
        - Empresa tech de 50 personas, remote-first
        - Mix de IC junior y senior + gestión
        - Necesita proceso estructurado pero no enfoque corporacional rígido
        - Debe producir documentos accionables, no solo discusión

        ---

        ### Próximos pasos
        1. Elegir entre Barcelona, Ámsterdam, Lisboa (investigar acceso de vuelos + opciones de venue)
        2. Obtener 2-3 presupuestos de hotel (espacio de reunión + habitaciones + catering)
        3. Contratar facilitador externo (pedir recomendaciones de la red)
        4. Definir fechas exactas (verificar calendarios del equipo por conflictos)
        5. Investigar proveedores de actividades para las opciones del Día 2

        ---

        ### Criterios de éxito
        - Asistencia 90%+ (la logística de viajes funciona para todas las regiones)
        - El equipo reporta sentirse más conectado post-evento
        - Roadmap del H2 claro acordado y documentado
        - Feedback positivo sobre el balance trabajo/social
        - El equipo de EE.UU. siente que el viaje valió la pena

  - type: tool_result
    tool: scratchpad_write
    content: "Escrito a offsite-spec.md (104 líneas)"

  - type: assistant
    content: "Listo — offsite-spec.md está preparado. Es una especificación completa que cubre propósito, logística, desglose de presupuesto, requisitos de alojamiento, estructura del programa hora por hora, opciones de actividades, resultados de planificación estratégica y criterios de éxito. Ahora puedes compartir esto con venues, facilitadores y tu equipo para input."

  - type: note
    text: "Fíjate en lo que acaba de pasar. El usuario empezó con 'necesito planificar un retreat de empresa.' A través de 16 preguntas en cuatro rondas, la IA sacó detalles que el usuario no había pensado especificar de antemano: necesidades de facilitador, restricciones alimentarias, política de inclusión familiar, resultados tangibles de las sesiones de planificación."

  - type: note
    text: "La especificación ahora es lo suficientemente completa para hacer briefing a un venue, contratar un facilitador y obtener la aprobación del equipo. Compara esto con intentar escribir todo eso en una sola instrucción de antemano. El enfoque interactivo produjo una especificación mejor más rápido, porque las preguntas de la IA revelaron lo que necesitaba especificarse."
```

## Meta-Prompting en directo

En lugar de solo describir el meta-prompting, observa qué pasa. El usuario no sabe exactamente lo que necesita — así que le pide a la IA que le ayude a descubrirlo.

```agent
id: meta-prompting-demo
title: "La IA ayuda a escribir la instrucción"
model_label: "Claude"

system: |
  Eres un asistente de ingeniería de instrucciones. Cuando te piden ayudar a escribir
  una instrucción, haz primero preguntas de aclaración para entender exactamente
  qué se necesita. Luego produce una instrucción completa y bien estructurada.

tools: []

visibility:
  system_prompt: hidden
  tool_calls: hidden
  full_context: hidden

sidebar:
  width: "45%"
  start_open: true

script:
  - type: note
    text: "El usuario tiene un objetivo vago: quiere hacer algo útil con los datos de comentarios de clientes. En lugar de escribir una instrucción mediocre, le pide a la IA que le ayude a escribir una buena."

  - type: user
    content: "Quiero hacer algo útil con nuestros datos de comentarios de clientes. Ayúdame a escribir una instrucción que obtenga el mejor análisis. Hazme preguntas primero."

  - type: assistant
    content: "Buen enfoque — permíteme hacer unas preguntas para construir la instrucción correcta:\n\n1. **¿En qué formato están los comentarios?** (¿CSV, respuestas de encuesta, reseñas de texto libre, tickets de soporte?)\n2. **¿Cuántos datos?** (¿Decenas, centenas, miles de entradas?)\n3. **¿Qué decisiones informará esto?** (¿Roadmap de producto? ¿Marketing? ¿Staff de soporte?)\n4. **¿Quién leerá el análisis?** (¿Equipo ejecutivo? ¿Gestores de producto? ¿Toda la empresa?)\n5. **¿Qué has intentado antes?** (¿Intentos previos de análisis, o empezando desde cero?)\n6. **¿Alguna preocupación específica?** (¿Problemas conocidos que quieres validar? ¿Hipótesis que probar?)"

  - type: user
    content: "Es un export CSV de nuestra encuesta de NPS — unas 800 respuestas del Q4. Mix de puntajes y comentarios de texto libre. Lo necesito para la sesión de planificación del Q1 del equipo de producto la próxima semana. El VP de Producto lo presentará. Creemos que la experiencia móvil es un problema pero no estamos seguros."

  - type: assistant
    content: "Aquí está tu instrucción:\n\n---\n\n**Analiza estos datos de encuesta NPS (800 respuestas del Q4, CSV adjunto).**\n\n**Objetivos:**\n1. Identifica las 5 principales temáticas en los comentarios de texto libre, ordenadas por frecuencia\n2. Para cada temática, desglosa por banda de puntaje NPS (Detractores 0-6, Pasivos 7-8, Promotores 9-10)\n3. Investiga específicamente las menciones de experiencia móvil — frecuencia, sentimiento y correlación con puntajes bajos\n4. Marca cualquier temática emergente que aparezca en el 3%+ de las respuestas\n\n**Formato de salida:**\n- Resumen ejecutivo (3-4 viñetas, adecuado para presentación del VP)\n- Desglose detallado de temáticas como tabla\n- Sección de análisis profundo de experiencia móvil\n- 5 acciones recomendadas para la planificación de producto del Q1, ordenadas por impacto potencial en NPS\n\n**Restricciones:**\n- Usa citas directas de las respuestas como evidencia (3-5 por temática)\n- Distingue entre correlación y causalidad\n- Marca si el tamaño de muestra para cualquier segmento es demasiado pequeño para conclusiones confiables\n- Tono profesional adecuado para presentación de liderazgo\n\n---\n\nEsta instrucción le da a la IA tu contexto de datos, una estructura de análisis clara, una hipótesis específica que probar (móvil) y un formato de salida adaptado a tu audiencia. ¿Quieres que ajuste algo?"

  - type: note
    text: "Fíjate en lo que pasó. El usuario empezó con 'quiero hacer algo útil con nuestros comentarios de clientes.' Después de seis preguntas y un conjunto de respuestas, la IA produjo una instrucción con objetivos claros, formato de salida estructurado, una hipótesis específica que probar y restricciones. Esa instrucción habría tomado 15 minutos escribir desde cero."

  - type: note
    text: "Este es el meta-prompting: usar la IA para escribir la instrucción que la IA luego ejecutará. Las preguntas sacan a la luz lo que sabes pero aún no has articulado."
```

```quiz
id: meta-prompting-purpose
type: multiple-choice
question: "¿Cuándo es el meta-prompting más valioso?"
options:
  - "Cuando sabes exactamente lo que quieres pero no puedes articularlo de forma concisa"
  - "Cuando tienes un objetivo vago y no estás seguro de qué especificidad necesita la IA para producir una buena salida"
  - "Cuando necesitas producir múltiples variaciones del mismo documento"
answer: 1
explanation: "El meta-prompting es más valioso cuando no estás seguro del alcance o los requisitos. Las preguntas de aclaración de la IA sacan a la luz qué información y especificidad se necesita, convirtiendo un objetivo vago en una especificación precisa. Si ya sabes exactamente lo que quieres, simplemente pregunta directamente."
```

```quiz
id: interactive-speccing-when
type: multiple-choice
question: "Necesitas crear un informe anual de 20 páginas para tu empresa. ¿Qué enfoque producirá el mejor resultado?"
options:
  - "Escribir una instrucción detallada única especificando cada sección, formato y punto de datos"
  - "Pedirle a la IA que lo redacte, luego iterar a través de 5-6 rondas de feedback"
  - "Especificárselo interactivamente primero — acordar estructura y requisitos, luego ejecutar contra la especificación acordada"
answer: 2
explanation: "Para entregables complejos, la especificación interactiva produce mejores resultados que una instrucción monolítica (que inevitablemente pasa por alto requisitos) o una iteración pura (que carece de un objetivo claro). La especificación se convierte en un contrato que asegura que tanto tú como la IA están alineados sobre cómo se ve 'terminado' antes de que ocurra cualquier trabajo."
```
