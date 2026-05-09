---
title: "Skills: Procedimientos Reutilizables"
duration: "15m"
tags: [skills, procedures, automation]
---

# Skills

Las skills son procedimientos documentados que la IA puede seguir. Piensa en ellas como manuales de operación para la IA.

## ¿Qué es una Skill?

Una skill agrupa todo lo necesario para hacer una tarea a tu manera:
- Instrucciones (cómo hacerla)
- Plantillas (qué formato usar)
- Ejemplos (qué aspecto tiene un buen resultado)
- Código (scripts que obtienen datos, llaman a APIs o automatizan la preparación)

Cuando le pides a la IA que haga algo, revisa tus skills disponibles. Si hay una relevante, sigue esos procedimientos automáticamente.

```callout
type: info
title: "Una Nota sobre la Nomenclatura"
content: "\"Skills\" es actualmente el nombre que Claude da a este concepto — una carpeta estructurada de instrucciones y recursos que la IA carga bajo demanda. Pero el patrón subyacente es universal. Todas las principales plataformas de IA convergen en ideas similares: procedimientos reutilizables definidos por el usuario que personalizan el comportamiento de la IA para tareas específicas. Sea cual sea el nombre que tu herramienta les dé, los principios de esta sección se aplican."
```

## Estructura de una Skill

Una carpeta de skills típica podría parecerse a:

```
Skills/
├── weekly-report/
│   ├── SKILL.md
│   ├── example-report.md
│   └── fetch_metrics.py
├── expense-summary/
│   ├── SKILL.md
│   └── format-template.xlsx
└── client-proposal/
    ├── SKILL.md
    ├── template.pptx
    └── lookup_client.py
```

Cada skill es una carpeta que contiene sus instrucciones (`SKILL.md`), material de referencia y opcionalmente **código** — pequeños scripts que se conectan a APIs, obtienen datos o automatizan pasos de preparación. La IA lee las instrucciones y usa los archivos de soporte para ejecutar la tarea.

## Anatomía de una Skill

Las instrucciones de una skill típicamente incluyen:

**Propósito** — ¿Para qué es esta skill?
"Generar el resumen de liderazgo del viernes."

**Pasos** — ¿Qué debería ocurrir?
1. Obtener las tareas completadas de esta semana
2. Identificar bloqueos y riesgos
3. Listar las prioridades de la próxima semana
4. Formatear usando la plantilla

**Tono** — ¿Cómo debería sonar?
"Profesional pero conciso. Sin jerga."

**Resultado** — ¿Qué se debe entregar exactamente?
"Documento de 1 página en Word siguiendo la plantilla."

## Las Skills Pueden Incluir Código

Las skills no se limitan a instrucciones y plantillas. Puedes incluir scripts que extiendan lo que la IA puede hacer:

**Conexiones a APIs** — Un script Python que obtiene métricas desde tu panel interno, para que la IA tenga datos frescos con los que trabajar.

**Procesamiento de datos** — Un script que limpia y formatea una exportación CSV antes de que la IA la analice.

**Generación de imágenes** — Código que llama a una API de imágenes para crear gráficos o ilustraciones como parte de un flujo de trabajo de informes.

**Integración con sistemas** — Scripts que publican resultados en Slack, actualizan una hoja de cálculo o activan un despliegue.

La IA lee el código, entiende qué hace y puede ejecutarlo como parte de la ejecución de la skill. No necesitas ser desarrollador — describe lo que necesitas y deja que la IA escriba el código por ti.

## Por Qué Importan las Skills

**Escribir una vez, usar para siempre**
Documenta el procedimiento una vez. Cada tarea futura de ese tipo lo seguirá automáticamente.

**Calidad consistente**
Sin variaciones por olvidar pasos o recordar mal las preferencias.

**Capturar conocimiento institucional**
Tus mejores prácticas se convierten en capacidades de la IA. Los nuevos miembros del equipo se benefician inmediatamente.

**Sin volver a explicar**
"Genera el informe semanal" basta — la IA sabe qué significa para tu organización.

```callout
type: tip
title: "Construyendo tu Biblioteca de Skills"
content: "Empieza con tareas que haces repetidamente. Cada vez que expliques algo a la IA, pregúntate: ¿debería esto ser una skill? Si lo harás de nuevo, la respuesta probablemente sea sí."
```

## Cuándo Ser Explícito sobre Skills y Herramientas

A veces la IA usa automáticamente la skill o herramienta correcta sin que se le pida. Dices "genera el informe semanal" y simplemente sabe usar la skill. Otras veces, necesitas ser explícito: "usa la skill /review" o "busca en la web sobre..."

**¿Por qué esta inconsistencia?** La IA decide según el contexto si una skill o herramienta es relevante. A veces esa decisión es obvia. A veces no.

Saber cuándo ser explícito es en sí una skill que desarrollas con la práctica.

**Regla general:** Si la IA no hace lo que esperas, intenta ser explícito sobre qué herramienta o skill usar.

En lugar de: "Busca información sobre esta empresa"
Prueba: "Busca en la web noticias recientes sobre esta empresa"

En lugar de: "Revisa este código"
Prueba: "Usa la skill /review para revisar este código"

Cuanto más específico seas sobre el mecanismo, menos tendrá que inferir la IA sobre tu intención.

```callout
type: tip
title: "Priorizar lo Implícito, Escalar a lo Explícito"
content: "Empieza con lenguaje natural. Si la IA no selecciona la herramienta o skill correcta, haz que tu próxima instrucción sea explícita sobre cuál usar. Este patrón — implícito primero, explícito al reintentar — es más rápido que ser excesivamente prescriptivo desde el inicio."
```

## Buenas Candidatas para Skills

- Informes regulares (semanales, mensuales, trimestrales)
- Tipos de documentos (propuestas, resúmenes, breves)
- Tareas de procesamiento de datos (limpieza, formateo, análisis)
- Plantillas de comunicación (emails, anuncios)
- Procesos de revisión (revisión de código, revisión de documentos)

Cualquier tarea con requisitos consistentes es candidata.

## Meta-herramientas: Herramientas que Crean Herramientas

El patrón más potente es usar la IA para crear herramientas personalizadas que luego usas repetidamente.

**El concepto:**
En lugar de hacer manualmente una tarea repetitiva o aprender a codificar una solución, describe la tarea a la IA una vez. La IA crea un script funcional. Lo implementas. Ahora es una herramienta.

**Ejemplos en la práctica:**

**Google Workspace Studio** — Construye agents para Gmail, Drive, Calendar usando lenguaje natural. "Resume las notas de la reunión y crea recordatorios en el calendario para los puntos de acción."

**Claude Code Skills** — Archivos de capacidad reutilizables que Claude carga bajo demanda. Crea una skill para ejecutar tu suite de tests, formatear código o desplegar en staging.

**MindStudio** — Constructor de agents sin código. Despliega como apps web, extensiones de navegador, activadores de email o automatizaciones programadas.

**El flujo de trabajo:**
1. Identifica una tarea repetitiva
2. Descríbesela a la IA — obtén un script funcional
3. Prueba e itera
4. Despliega como una herramienta reutilizable
5. Vuelve a la IA para mantenimiento cuando las necesidades cambien

Este es el patrón de usuario avanzado de mayor impacto. No solo estás usando IA — estás usando la IA para construir las herramientas que usarás a continuación.

```callout
type: tip
title: "Empezar Pequeño, Crecer Rápido"
content: "Tu primera herramienta personalizada podría ahorrar 10 minutos por semana. La décima podría ahorrar 2 horas. El efecto compuesto de construir herramientas es exponencial — cada herramienta que creas hace que la siguiente sea más fácil de construir y más valiosa para desplegar."
```

## Creando una Skill

No necesitas construir skills manualmente. Describe lo que necesitas y deja que la IA cree toda la skill por ti — estructura de carpetas, instrucciones, material de referencia, incluso el código.

Observa el proceso de convertir una tarea repetitiva en una carpeta de skill completa y reutilizable.

```agent
id: skill-creation-demo
title: "Construyendo una Skill Reutilizable"
model_label: "Claude"

system: |
  Eres un consultor de productividad ayudando a crear skills de IA reutilizables.
  Las skills deben ser claras, completas y accionables.
  Al crear carpetas de skill, usa la convención: Skills/<nombre-skill>/SKILL.md para instrucciones.

scratchpad:
  "weekly-report-example.md": |
    # Resumen de Liderazgo — Semana del 6 de Enero de 2026

    ## Resumen
    Semana sólida. Dos características principales entregadas. Pipeline saludable.

    ## Entregado esta Semana
    - Autenticación de usuarios v2 (proyecto de 3 semanas, completado en plazo)
    - Mejoras de rendimiento del panel (carga de página: 1.2s → 0.4s)

    ## En Progreso
    - Rediseño móvil: 60% completado, en camino para el 1 de febrero
    - API v3: Revisión de diseño completada, construcción empieza el lunes

    ## Bloqueos y Riesgos
    - Capacidad del equipo de diseño: 1 diseñador de baja hasta el 20 de enero
    - Backlog de QA creciendo: 12 tickets en cola (objetivo: menos de 5)

    ## Métricas Clave
    - Velocidad de sprint: 42 puntos (objetivo: 40) ✓
    - Tasa de fugas de bugs: 2.1% (objetivo: menos de 3%) ✓
    - Frecuencia de despliegue: 4 despliegues esta semana

    ## Prioridades la Próxima Semana
    1. Empezar construcción de API v3
    2. Limpiar backlog de QA
    3. Preparar presentación de revisión de hitos de febrero

tools:
  - scratchpad_read
  - scratchpad_write
  - create_folder
  - move_file

visibility:
  system_prompt: hidden
  tool_calls: visible
  full_context: hidden

sidebar:
  width: "45%"
  start_open: true

script:
  - type: note
    text: "Esta persona escribe un informe de liderazgo semanal cada viernes. Le toma 45 minutos. Quieren convertirlo en una skill propia — una carpeta con instrucciones, material de referencia y un script para obtener métricas automáticamente. Observa cómo la IA construye toda la skill para ellos."

  - type: user
    content: "Escribo este informe de liderazgo cada viernes y me lleva una eternidad. Tengo un ejemplo de informe aquí. ¿Puedes construirme una carpeta de skill propia para esto? Necesito instrucciones claras, mi ejemplo como referencia y un script Python que obtenga nuestras métricas de ingeniería desde la API del panel interno para no tener que copiar los números manualmente cada semana."

  - type: tool_call
    tool: scratchpad_read
    args:
      filename: "weekly-report-example.md"

  - type: tool_result
    tool: scratchpad_read
    content: |
      # Resumen de Liderazgo — Semana del 6 de Enero de 2026

      ## Resumen
      Semana sólida. Dos características principales entregadas. Pipeline saludable.

      ## Entregado esta Semana
      - Autenticación de usuarios v2 (proyecto de 3 semanas, completado en plazo)
      - Mejoras de rendimiento del panel (carga de página: 1.2s → 0.4s)

      ## En Progreso
      - Rediseño móvil: 60% completado, en camino para el 1 de febrero
      - API v3: Revisión de diseño completada, construcción empieza el lunes

      ## Bloqueos y Riesgos
      - Capacidad del equipo de diseño: 1 diseñador de baja hasta el 20 de enero
      - Backlog de QA creciendo: 12 tickets en cola (objetivo: menos de 5)

      ## Métricas Clave
      - Velocidad de sprint: 42 puntos (objetivo: 40) ✓
      - Tasa de fugas de bugs: 2.1% (objetivo: menos de 3%) ✓
      - Frecuencia de despliegue: 4 despliegues esta semana

      ## Prioridades la Próxima Semana
      1. Empezar construcción de API v3
      2. Limpiar backlog de QA
      3. Preparar presentación de hitos de febrero

  - type: tool_call
    tool: create_folder
    args:
      path: "Skills"

  - type: tool_result
    tool: create_folder
    content: "Carpeta creada: Skills"

  - type: tool_call
    tool: create_folder
    args:
      path: "Skills/weekly-report"

  - type: tool_result
    tool: create_folder
    content: "Carpeta creada: Skills/weekly-report"

  - type: tool_call
    tool: scratchpad_write
    args:
      filename: "Skills/weekly-report/SKILL.md"
      content: |
        # Skill: Informe de Liderazgo Semanal

        ## Propósito
        Generar el resumen de liderazgo del viernes para el equipo de ingeniería.

        ## Cuándo Usar
        Cada viernes por la tarde. Ejecuta fetch_metrics.py primero para obtener datos frescos, luego aporta cualquier contexto adicional.

        ## Preparación (la IA hace esto automáticamente)
        1. Ejecutar `python fetch_metrics.py` para obtener datos frescos del panel
        2. Usar los datos obtenidos junto con cualquier contexto adicional proporcionado

        ## Entradas Requeridas (del usuario)
        - Trabajo completado esta semana (desde Jira/Linear o notas de standup)
        - Elementos actuales en progreso con % de completado
        - Cualquier bloqueo o riesgo identificado esta semana
        - Top 3 prioridades de la próxima semana

        ## Formato de Salida
        Documento en Markdown con estas secciones exactas:

        1. **Resumen** — 1-2 frases. Tono general de la semana. Empieza con el titular.
        2. **Entregado esta Semana** — Lista con viñetas. Cada uno: nombre de característica, duración, estado de plazo.
        3. **En Progreso** — Lista con viñetas. Cada uno: nombre de característica, % completado, fecha esperada.
        4. **Bloqueos y Riesgos** — Lista con viñetas. Cada uno: problema, impacto, mitigación si se conoce.
        5. **Métricas Clave** — Lista con viñetas. Cada uno: nombre, valor, objetivo, aprobado/reprobado.
        6. **Prioridades la Próxima Semana** — Lista numerada. Solo top 3. Elementos accionables.

        ## Tono y Estilo
        - Directo y conciso. Sin palabras de relleno.
        - Confidente pero honesto sobre los riesgos.
        - Usar números específicos, no "mejorado" o "creciendo."
        - Viñetas, no párrafos.
        - Título: "Resumen de Liderazgo — Semana del [fecha]"

        ## Referencia
        Ver example-report.md para el formato y tono a imitar.

        ## Controles de Calidad
        - Cada elemento "En Progreso" tiene un porcentaje y fecha
        - Cada métrica tiene real vs objetivo
        - Los bloqueos incluyen impacto, no solo descripción
        - El Resumen refleja los datos generales, no solo las buenas noticias

  - type: tool_result
    tool: scratchpad_write
    content: "Escrito en Skills/weekly-report/SKILL.md (48 líneas)"

  - type: tool_call
    tool: move_file
    args:
      source: "weekly-report-example.md"
      destination: "Skills/weekly-report/example-report.md"

  - type: tool_result
    tool: move_file
    content: "Movido weekly-report-example.md → Skills/weekly-report/example-report.md"

  - type: tool_call
    tool: scratchpad_write
    args:
      filename: "Skills/weekly-report/fetch_metrics.py"
      content: |
        #!/usr/bin/env python3
        """Obtén métricas de ingeniería semanales desde la API del panel.

        Ejecuta antes de generar el informe semanal:
            python fetch_metrics.py > metrics.md

        Requiere variable de entorno DASHBOARD_API_KEY.
        """

        import os
        import requests
        from datetime import datetime, timedelta

        API = "https://dashboard.internal.company.com/api/v1"
        KEY = os.environ["DASHBOARD_API_KEY"]
        HEADERS = {"Authorization": f"Bearer {KEY}"}


        def get_sprint():
            r = requests.get(f"{API}/sprints/current", headers=HEADERS)
            r.raise_for_status()
            return r.json()


        def get_deploys():
            since = (datetime.now() - timedelta(days=7)).isoformat()
            r = requests.get(
                f"{API}/deploys",
                params={"since": since},
                headers=HEADERS,
            )
            r.raise_for_status()
            return r.json()["deploys"]


        def get_quality():
            r = requests.get(f"{API}/quality/summary", headers=HEADERS)
            r.raise_for_status()
            return r.json()


        if __name__ == "__main__":
            sprint = get_sprint()
            deploys = get_deploys()
            quality = get_quality()

            print("# Métricas de esta Semana\n")
            print(f"- Velocidad de sprint: {sprint['velocity']} pts "
                  f"(objetivo: {sprint['target']})")
            print(f"- Despliegues esta semana: {len(deploys)}")
            print(f"- Tasa de fugas de bugs: {quality['escape_rate']}% "
                  f"(objetivo: menos de {quality['target_rate']}%)")
            print(f"- Profundidad de cola de QA: {quality['queue_depth']} tickets")

  - type: tool_result
    tool: scratchpad_write
    content: "Escrito en Skills/weekly-report/fetch_metrics.py (42 líneas)"

  - type: assistant
    content: "Construí tu carpeta de skill completa. Esto es lo que tienes ahora:\n\n**Skills/weekly-report/**\n- `SKILL.md` — Instrucciones completas: propósito, entradas, formato de salida, tono, controles de calidad\n- `example-report.md` — Tu informe de referencia, para que la IA imite tu estilo\n- `fetch_metrics.py` — Obtén velocidad de sprint, conteo de despliegues, tasa de bugs y cola de QA desde tu API del panel\n\n**Cómo usarla:**\nSimplemente di *'Genera el resumen de liderazgo de esta semana'* y proporciona tus notas. La skill indica a la IA ejecutar el script de métricas automáticamente, obtener los números frescos y producir el informe en tu formato exacto. Tú nunca tocas el script tú mismo.\n\nEscríbela una vez, úsala cada viernes."

  - type: note
    text: "Observa la estructura de carpetas: Skills/weekly-report/ contiene todo — instrucciones, material de referencia y código. Esta es una skill autónoma y portátil. Podrías compartirla con un colega y tendría todo lo que necesita."

  - type: note
    text: "Idea clave: no escribiste nada de esto manualmente. Describiste lo que necesitabas y la IA construyó toda la skill — instrucciones, organización de archivos, incluso el script Python. Ese es el patrón de meta-herramientas: usa la IA para construir las herramientas que usarás con la IA."

  - type: note
    text: "El efecto compuesto: esta skill se usará 50+ veces al año. Cada uso ahorra 40 minutos. Eso son 33+ horas ahorradas a partir de una sola conversación. Y la calidad es consistente cada vez."
```

```quiz
id: skills-purpose
type: multiple-choice
question: "¿Cuál es el principal beneficio de crear una skill para una tarea recurrente?"
options:
  - "Documentar el procedimiento una vez; la IA lo sigue automáticamente en cada tarea futura"
  - "Las skills hacen que la IA sea más rápida al procesar"
  - "Las skills reducen el coste del uso de IA"
answer: 0
explanation: "Las skills capturan tus procedimientos para que los documentes una vez y la IA los siga automáticamente cada vez. Sin volver a explicar, resultados consistentes, conocimiento institucional preservado."
```
