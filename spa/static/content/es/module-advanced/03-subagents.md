---
title: "Subagentes: la IA Genera Ayudantes"
duration: "15m"
tags: [subagents, parallel, architecture]
---

# Subagentes

Cuando un agente de IA no es suficiente, puede generar más para trabajar en paralelo.

## Cómo Funcionan los Subagentes

![Subagent Architecture: Main Agent delegates to parallel Subagents](/content/module-advanced/images/subagents.svg)

1. Le das una tarea compleja
2. El agente principal la divide en piezas
3. Cada pieza va a un subagente
4. Trabajan simultáneamente
5. Los resultados regresan y se combinan

A veces la IA orquesta esto automáticamente. Pero el verdadero poder viene cuando **tú defines las plantillas de subagente** — archivos de instrucción que le dicen a cada subagente exactamente cómo comportarse, qué herramientas usar y qué resultado producir. No solo delegas; estás diseñando el equipo.

```callout
type: info
title: "Una Nota sobre la Nomenclatura"
content: "\"Subagents\" es la terminología de Claude, pero el patrón es universal. OpenAI los llama \"sub-tasks\", Google usa \"agent delegation\", y frameworks como LangGraph y CrewAI implementan la misma idea. El concepto — un agente padre generando agentes hijos especializados — está convergiendo en todas las principales plataformas."
```

## Por Qué Importa

**Velocidad**
Múltiples hilos de investigación se ejecutan a la vez. Lo que le tomaría a un agente una hora le toma al equipo minutos.

**Minuciosidad**
Diferentes ángulos explorados simultáneamente. Nada esperando en una cola.

**Aislamiento de contexto**
Cada subagente recibe su propia ventana de contexto. El module-builder que trabaja en "análisis de datos" no ve el contenido de "técnicas de prompting". Esto previene la contaminación cruzada y mantiene a cada agente enfocado en su tarea específica.

**Diferentes modelos para diferentes trabajos**
No cada subtarea necesita el modelo más caro. Tu orquestador podría usar el mejor modelo para planificación, mientras los subagentes usan modelos más rápidos y baratos para ejecución sencilla. Un patrón de planificar-entonces-ejecutar: el modelo grande piensa, los modelos pequeños ejecutan.

**Especialización**
Cada subagente puede tener diferentes instrucciones, diferentes herramientas y diferentes prompts de sistema. Uno podría ser un investigador, otro un escritor, otro un revisor.

## Definiendo Plantillas de Subagente

No necesitas construir plantillas de subagente manualmente — describe lo que necesitas y deja que la IA las cree. Pero deberías saber qué va en una:

**Instrucciones** — Qué hace este subagente, cómo debería comportarse, qué convenciones seguir.

**Herramientas** — Qué capacidades tiene acceso el subagente (escritura de archivos, ejecución de código, generación de imágenes, llamadas a APIs).

**Restricciones** — Qué NO debería hacer el subagente. Los límites mantienen a los agentes enfocados.

**Formato de salida** — Dónde escribir los resultados y en qué estructura.

Una plantilla es solo un archivo. Cuando el orquestador necesita ese tipo de agente, carga la plantilla y lanza una nueva instancia. Creas la plantilla una vez; se reutiliza cada vez que se necesita ese tipo de trabajo.

```callout
type: tip
title: "Dejar que la IA Construya tus Plantillas"
content: "Al igual que con las skills, no necesitas escribir plantillas de subagente desde cero. Describe el rol — 'necesito un agente que revise documentos por problemas de cumplimiento y produzca una lista de verificación' — y deja que la IA cree el archivo de plantilla por ti."
```

## Buenos Casos de Uso

**Tareas de investigación**
Investigar competidores, analizar el mercado, revisar literatura académica — todo a la vez.

**Análisis multifacético**
Análisis financiero, revisión operativa, feedback del cliente — flujos en paralelo.

**Creación de contenido**
Investigación, estructuración, redactar diferentes secciones — simultáneo.

**Informes complejos**
Cada sección abordada por un subagente diferente, el agente principal ensambla.

## Patrones del Mundo Real

**ReAct (Reasoning + Action)**
Intercalar razonamiento con llamadas a herramientas. El agente piensa, actúa, observa el resultado, piensa de nuevo.

**Planificar-entonces-Ejecutar**
El modelo grande hace planificación, el modelo pequeño hace ejecución. Reduce costes manteniendo calidad.

**Descomposición Jerárquica de Tareas**
Un agente gerente divide tareas en subtareas para agentes especialistas. Como un gerente de proyecto distribuyendo trabajo.

**Ciclo Generador-Evaluador**
Un agente genera soluciones, otro evalúa y sugiere mejoras. Refinamiento continuo.

## Subagentes en Acción

Observa un agente orquestador planificar un curso de formación y delegar a subagentes en paralelo — cada uno construyendo un módulo diferente en su propia carpeta simultáneamente. Esto está directamente modelado sobre cómo se construyó este curso.

```agent
id: subagent-course-demo
title: "Construcción de Cursos con Subagentes en Paralelo"
model_label: "Course Director"

system: |
  Eres un agente course-director. Diseñas cursos de formación mediante:
  1. Analizar un brief de investigación
  2. Crear un plan de curso estructurado con asignaciones de modelos
  3. Delegar la creación de módulos a constructores de subagentes independientes
  Cada module-builder trabaja en su propia carpeta con su propio contexto.

scratchpad:
  "research-brief.md": |
    # Brief de Investigación: IA para Análisis de Datos

    ## Hallazgos Clave
    - El 73% de analistas pasa la mayor parte del tiempo limpiando datos — la IA automatiza el 80% de esto
    - La brecha es de metodología no de herramientas: los analistas necesitan frameworks para trabajar con IA
    - Tres niveles de habilidades identificados: fundamentos, técnicas, flujos de trabajo
    - Los ejercicios prácticos mejoran dramáticamente la retención vs solo teoría

    ## Público Objetivo
    Analistas de datos en carrera intermedia. Familiarizados con Excel y SQL. Sin experiencia en IA.

    ## Alcance
    3 módulos, 45 min cada uno. Enfoque práctico — cada concepto necesita un ejercicio.

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
    text: "Esto demuestra el patrón de subagente. Un orquestador lee un brief, crea un plan, luego delega a subagentes independientes — cada uno construyendo un módulo diferente en su propia carpeta aislada. Así es como se construyó este mismo curso."

  - type: user
    content: "Necesito un curso de formación de 3 módulos sobre 'IA para Análisis de Datos'. Aquí está el brief de investigación. Crea un plan de curso y luego construye los tres módulos."

  - type: tool_call
    tool: scratchpad_read
    args:
      filename: "research-brief.md"

  - type: tool_result
    tool: scratchpad_read
    content: |
      # Brief de Investigación: IA para Análisis de Datos

      ## Hallazgos Clave
      - El 73% de analistas pasa la mayor parte del tiempo limpiando datos — la IA automatiza el 80% de esto
      - La brecha es de metodología no de herramientas: los analistas necesitan frameworks para trabajar con IA
      - Tres niveles de habilidades identificados: fundamentos, técnicas, flujos de trabajo
      - Los ejercicios prácticos mejoran dramáticamente la retención vs solo teoría

      ## Público Objetivo
      Analistas de datos en carrera intermedia. Familiarizados con Excel y SQL. Sin experiencia en IA.

      ## Alcance
      3 módulos, 45 min cada uno. Enfoque práctico — cada concepto necesita un ejercicio.

  - type: tool_call
    tool: scratchpad_write
    args:
      filename: "course-plan.md"
      content: |
        # Plan de Curso: IA para Análisis de Datos

        ## Módulo 1: Fundamentos de Datos (45 min)
        Brief: Lo que la IA ve al leer datos. Tipos que maneja bien
        vs mal. Preparar datos para análisis con IA.
        Ejercicio: Limpiar un dataset desordenado con IA.
        Modelo asignado: Haiku (rápido, barato — contenido directo)

        ## Módulo 2: Técnicas de Análisis (45 min)
        Brief: Análisis exploratorio, reconocimiento de patrones, análisis
        estadístico sin código.
        Ejercicio: Analizar datos de ventas para encontrar insights.
        Modelo asignado: Sonnet (capaz — necesita ejemplos matizados)

        ## Módulo 3: Flujos de Trabajo Automatizados (45 min)
        Brief: Pipelines repetibles, conectar fuentes de datos,
        programar informes automáticos.
        Ejercicio: Construir un análisis semanal automatizado.
        Modelo asignado: Opus (mejor — patrones de integración complejos)

  - type: tool_result
    tool: scratchpad_write
    content: "Escrito en course-plan.md (20 líneas)"

  - type: note
    text: "El orquestador ha analizado el brief y producido un plan — incluyendo qué modelo de IA debe usar cada subagente. Modelos más baratos para contenido simple, modelos más caros para material complejo. Ahora crea el espacio de trabajo y delega."

  - type: tool_call
    tool: create_folder
    args:
      path: "course/01-data-foundations"

  - type: tool_result
    tool: create_folder
    content: "Carpeta creada: course/01-data-foundations"

  - type: tool_call
    tool: create_folder
    args:
      path: "course/02-analysis-techniques"

  - type: tool_result
    tool: create_folder
    content: "Carpeta creada: course/02-analysis-techniques"

  - type: tool_call
    tool: create_folder
    args:
      path: "course/03-automated-workflows"

  - type: tool_result
    tool: create_folder
    content: "Carpeta creada: course/03-automated-workflows"

  - type: note
    text: "Tres subagentes se lanzan ahora en paralelo. Cada uno recibe su propia carpeta, su propia ventana de contexto y solo la sección del plan relevante para su módulo. Este es el aislamiento de contexto — el module-1 builder no ve contenido de module-3, así que se mantiene enfocado y su ventana de contexto no se desperdicia en material irrelevante."

  - type: tool_call
    tool: scratchpad_write
    args:
      filename: "course/01-data-foundations/01-what-ai-sees.md"
      content: |
        ---
        title: "Lo que la IA Ve en tus Datos"
        duration: "15m"
        ---

        # Lo que la IA Ve en tus Datos

        Cuando abres una hoja de cálculo, ves filas y patrones
        que has aprendido a reconocer. Cuando la IA lee el mismo archivo,
        procesa cada celda pero no entiende ninguno del contexto
        que das por sentado.

        ## Estructura vs Significado

        La IA sobresale en datos estructurados — columnas limpias, formatos
        consistentes. Lucha con la realidad desordenada: celdas fusionadas,
        fechas inconsistentes, notas apretadas en campos numéricos.

        ## El Principio de Preparación

        El paso más importante antes del análisis con IA: limpia
        tus datos. No porque la IA no pueda manejar el desorden — sino que
        una entrada limpia produce una salida dramáticamente mejor.

  - type: tool_result
    tool: scratchpad_write
    content: "Escrito en course/01-data-foundations/01-what-ai-sees.md (22 líneas)"

  - type: tool_call
    tool: scratchpad_write
    args:
      filename: "course/02-analysis-techniques/01-finding-patterns.md"
      content: |
        ---
        title: "Encontrando Patrones que la IA Puede Ver"
        duration: "15m"
        ---

        # Encontrando Patrones que la IA Puede Ver

        La técnica de análisis más poderosa no es preguntar "qué ves?" —
        es decirle a la IA qué buscar.

        ## Exploración Dirigida

        Empieza con una hipótesis: "Creo que las ventas caen los miércoles."
        La IA valida o refuta esto en segundos a través de millones de
        filas. Los prompts no dirigidos producen observaciones genéricas.

        ## Detección de Anomalías

        Dale a la IA tu línea base: "Las ventas son normalmente de 10-15k por día."
        Marca cada día fuera de ese rango con contexto sobre
        qué más sucedió.

  - type: tool_result
    tool: scratchpad_write
    content: "Escrito en course/02-analysis-techniques/01-finding-patterns.md (20 líneas)"

  - type: tool_call
    tool: scratchpad_write
    args:
      filename: "course/03-automated-workflows/01-from-analysis-to-automation.md"
      content: |
        ---
        title: "Del Análisis a la Automatización"
        duration: "15m"
        ---

        # Del Análisis a la Automatización

        La primera vez que analizas datos con IA, toma una hora.
        La segunda vez, cinco minutos. La tercera vez, debería
        suceder automáticamente.

        ## La Escalera de Automatización

        1. Manual — tú haces el análisis con ayuda de IA
        2. Repetible — has guardado los prompts y el proceso
        3. Programado — se ejecuta semanalmente sin ti
        4. Monitoreado — te alerta cuando los resultados son inusuales

        ## Construyendo tu Primera Pipeline

        Empieza con tu análisis más repetido. Documenta los pasos
        como un archivo de skill. Añade un script para obtener los datos. Programa la ejecución.

  - type: tool_result
    tool: scratchpad_write
    content: "Escrito en course/03-automated-workflows/01-from-analysis-to-automation.md (21 líneas)"

  - type: tool_call
    tool: list_files
    args: {}

  - type: tool_result
    tool: list_files
    content: |
        course-plan.md
        research-brief.md
        course/
          01-data-foundations/
            01-what-ai-sees.md
          02-analysis-techniques/
            01-finding-patterns.md
          03-automated-workflows/
            01-from-analysis-to-automation.md

  - type: note
    text: "Los tres módulos están ahora construidos. En un sistema real estos se ejecutan simultáneamente — lo que le toma a un agente 3 horas le toma a tres subagentes 1 hora. Observa la estructura de carpetas en el explorador: cada subagente escribió exclusivamente a su propio directorio. Sin riesgo de contaminación cruzada."

  - type: note
    text: "Las asignaciones de modelos importan. El Módulo 1 (conceptos básicos) usó Haiku — rápido y barato. El Módulo 3 (integración compleja) usó Opus — el mejor modelo. Ajustar la capacidad del modelo a la complejidad de la tarea puede reducir costes un 60-80% sin sacrificar calidad donde importa."

  - type: note
    text: "Ahora observa al orquestador regresar. Los subagentes terminaron — el course-director lee su salida y ensambla la vista general final del curso. Este es el patrón completo: planificar → delegar → ensamblar."

  - type: tool_call
    tool: scratchpad_read
    args:
      filename: "course/01-data-foundations/01-what-ai-sees.md"

  - type: tool_result
    tool: scratchpad_read
    content: |
        ---
        title: "Lo que la IA Ve en tus Datos"
        duration: "15m"
        ---

        # Lo que la IA Ve en tus Datos

        Cuando abres una hoja de cálculo, ves filas y patrones
        que has aprendido a reconocer. Cuando la IA lee el mismo archivo,
        procesa cada celda pero no entiende ninguno del contexto
        que das por sentado.

  - type: tool_call
    tool: scratchpad_read
    args:
      filename: "course/02-analysis-techniques/01-finding-patterns.md"

  - type: tool_result
    tool: scratchpad_read
    content: |
        ---
        title: "Encontrando Patrones que la IA Puede Ver"
        duration: "15m"
        ---

        # Encontrando Patrones que la IA Puede Ver

        La técnica de análisis más poderosa no es preguntar "qué ves?" —
        es decirle a la IA qué buscar.

  - type: tool_call
    tool: scratchpad_read
    args:
      filename: "course/03-automated-workflows/01-from-analysis-to-automation.md"

  - type: tool_result
    tool: scratchpad_read
    content: |
        ---
        title: "Del Análisis a la Automatización"
        duration: "15m"
        ---

        # Del Análisis a la Automatización

        La primera vez que analizas datos con IA, toma una hora.
        La segunda vez, cinco minutos. La tercera vez, debería
        suceder automáticamente.

  - type: tool_call
    tool: scratchpad_write
    args:
      filename: "course/README.md"
      content: |
        # IA para Análisis de Datos — Resumen del Curso

        Un curso práctico de 3 módulos para analistas de datos.

        ## Módulos

        ### 1. Fundamentos de Datos
        Lo que la IA ve al leer tus datos. Estructurado vs desordenado,
        principios de preparación y tu primera limpieza asistida por IA.

        ### 2. Técnicas de Análisis
        Exploración dirigida, reconocimiento de patrones y detección
        de anomalías. Análisis estadístico sin escribir código.

        ### 3. Flujos de Trabajo Automatizados
        Del análisis manual a pipelines programadas. Construye
        flujos de trabajo de datos repetibles y monitoreados.

        ## Prerrequisitos
        Familiaridad con Excel y SQL básico. No se necesita experiencia en IA.

        ## Duración
        ~2.5 horas en total (45 min por módulo + ejercicios)

  - type: tool_result
    tool: scratchpad_write
    content: "Escrito en course/README.md (22 líneas)"

  - type: assistant
    content: "Curso completo. Tres módulos construidos en paralelo, ensamblados en un resumen final. El README reúne los temas clave de cada módulo en un programa coherente.\n\nTu curso está listo para revisión en la carpeta course/."

  - type: note
    text: "Ese último paso es el patrón orquestador en su totalidad. El course-director leyó la salida de cada módulo, luego escribió un README que los sintetiza en un todo coherente. Planificar → delegar → ensamblar. Los subagentes nunca ven el trabajo de los otros — solo el orquestador tiene el panorama completo."

  - type: note
    text: "Esto es exactamente cómo se construyó este curso de formación. Un agente course-director analizó un informe de investigación, creó un plan, luego lanzó subagentes module-builder en paralelo — cada uno con sus propias instrucciones, su propia carpeta y acceso a skills como /quiz y /mermaid para contenido especializado. El orquestador ensambló el resultado final."
```

## Este Curso se Construyó a Sí Mismo

Cada concepto que has aprendido en este curso — ingeniería de contexto, delegación, skills, subagentes, archivos, verificación — fue usado para construir el curso en el que estás sentado.

![Course creation in action](/static/images/meta-course-creation-screenshot.png)

### El Proceso de Creación

**Fase 1: Investigación Profunda**
El instructor usó la Deep Research de Claude para producir un informe de 8.000 palabras que sintetiza hallazgos de Anthropic, OpenAI, Google, McKinsey, Bain, Deloitte, estudios académicos y blogs de practicantes. Esto es **ingeniería de contexto** en práctica — la calidad de todo lo downstream depende de la calidad del contexto con el que empiezas.

**Fase 2: Agente Course-Director**
Se lanzó un subagente con un brief detallado: leer el informe de investigación, leer cada módulo existente, comparar cobertura contra necesidades y producir un plan de actualización integral.

El agente leyó 8.000 palabras de investigación, las cruzó contra el curso existente, identificó brechas y produjo un plan de actualización priorizado de 720 líneas.

Esto es **Descomposición Jerárquica de Tareas** — el usuario delega a un agente "manager" que hace análisis y planificación.

**Fase 3: Module-Builders en Paralelo**
El course-director lanzó agentes module-builder independientes — uno por módulo, ejecutándose en paralelo. Cada module-builder creó su directorio, escribió páginas en markdown e invocó skills para contenido especializado: `/quiz` para controles de comprensión, `/mermaid` para diagramas, `/agent-demo` para walkthroughs interactivos.

**Fase 4: Humano en el Bucle**
En cada nivel, el humano revisó y dirigió. El informe de investigación fue curado. El plan fue revisado antes de lanzar módulos. La salida del agente fue verificada antes de incluirse. IA-first, humano-verificado.

### La Jerarquía

![Multi-agent hierarchy: User delegates to course-director, which spawns parallel module-builders, each using skills](/content/module-advanced/images/meta-hierarchy.svg)

### Qué Esto Demuestra

- **Jerarquías multi-agente** — usuario → course-director → module-builders
- **Composición de skills** — agentes invocando skills especializadas bajo demanda
- **Aislamiento de contexto** — cada module-builder trabaja en su propio directorio
- **Selección de modelos** — diferentes tareas, diferentes modelos
- **Investigación profunda como contexto** — la calidad de la investigación determina la calidad de todo lo downstream
- **Flujos de trabajo basados en archivos** — todo son archivos (informe, plan, YAML, markdown)
- **Verificación** — el humano revisa toda la salida del agente antes de aceptarla

```callout
type: info
title: "El Momento Autorreferencial"
content: "Cada concepto en este curso — ingeniería de contexto, delegación, skills, subagentes, archivos, verificación — fue usado para construir el curso en el que estás sentado. Esto no es un ejercicio teórico. Así es como se hace el trabajo en febrero de 2026."
```

```quiz
id: subagents-benefit
type: multiple-choice
question: "¿Cuál es el beneficio principal de que la IA use subagentes?"
options:
  - "Las tareas complejas se abordan en paralelo, completándose más rápido y de forma más minuciosa"
  - "Los subagentes son más precisos que los agentes individuales"
  - "Los subagentes son más baratos de ejecutar"
answer: 0
explanation: "Los subagentes permiten trabajo en paralelo — múltiples aspectos de una tarea compleja siendo abordados simultáneamente. Esto significa una completación más rápida y cobertura más minuciosa que el trabajo secuencial de un solo agente."
```

```quiz
id: subagent-context-isolation
type: multiple-choice
question: "¿Por qué cada subagente recibe su propia ventana de contexto?"
options:
  - "Cada agente se mantiene enfocado en su tarea sin material irrelevante llenando su contexto"
  - "Para aumentar la cantidad total de texto que el sistema puede procesar"
  - "Para prevenir que los agentes se comuniquen entre sí"
answer: 0
explanation: "El aislamiento de contexto mantiene a cada subagente enfocado. Un module-builder trabajando en 'análisis de datos' no necesita ver contenido de 'técnicas de prompting' — eso desperdiciaría espacio de contexto y potencialmente causaría confusión. Cada agente recibe exactamente el contexto que necesita."
```
