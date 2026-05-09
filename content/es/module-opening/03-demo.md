---
title: "Demo: Del Caos al Informe"
duration: "5m"
tags: [demo, delegacion]
---

# La Demo: Del Caos al Informe

Imagina una carpeta desorganizada con tipos de archivos mezclados — notas de reuniones, hojas de cálculo, correos electrónicos, líneas temporales — el tipo de caos que se acumula en proyectos reales.

Con una sola instrucción de delegación, la IA puede:
- Planificar y ejecutar la tarea de forma autónoma
- Procesar múltiples tipos de archivos
- Entregar un documento pulido y terminado

Nota: Sin instrucciones paso a paso. Solo describes el resultado que quieres. La IA descubre cómo llegar allí.

**Así es como entendemos "IA que hace trabajo."**

## Cómo Funciona

El **explorador de archivos a la izquierda** muestra cinco archivos de proyecto desordenados — notas de reuniones, datos presupuestarios, correos de clientes, una línea temporal y retroalimentación del equipo.

El **panel de conversación a la derecha** muestra la interacción con la IA. Avanza usando el botón **Siguiente** y observa cómo una instrucción convierte este caos en un informe ejecutivo pulido.

```agent
id: opening-demo-es
title: "Del Caos al Informe"
model_label: "Claude"

system: |
  Eres un asistente de gestión de proyectos. Ayudas a profesionales a crear informes claros y accionables a partir de materiales de proyecto. Sé minucioso pero conciso. Escribe en un tono profesional adecuado para audiencias ejecutivas.

scratchpad:
  "notas-reunion.txt": |
    Proyecto Phoenix - Reunión de Equipo 15 de Enero
    Asistentes: Sarah (PM), Equipo de Desarrollo, Lisa (Diseño)
    - Migración de API 60% completa, objetivo 1 de febrero
    - Revisión de diseño bloqueada por directrices de marca
    - El cliente quiere añadir dashboard móvil a Fase 1
    - Tarifas de contratistas aumentaron 15% — preocupación presupuestaria
    - Sarah debe escalar riesgo de línea temporal a liderazgo
    - Próxima revisión de hito: 22 de enero
  "presupuesto-t1.csv": |
    Categoria,Planificado,Real,Variacion
    Desarrollo,45000,38000,-7000
    Diseño,15000,12500,-2500
    Infraestructura,8000,9200,+1200
    Contratistas,20000,27500,+7500
    Testing,7000,4200,-2800
    Total,95000,91400,-3600
  "correos-cliente.txt": |
    De: David Chen <d.chen@acmecorp.com>
    Fecha: 14 de enero
    Asunto: RE: Línea Temporal Proyecto Phoenix

    Sarah — La junta directiva quiere dashboard móvil en Fase 1.
    ¿Línea temporal revisada? Lo necesitamos para el 15 de marzo como máximo.
    ---
    De: Sarah Lopez <s.lopez@team.com>
    Fecha: 12 de enero

    David — Evaluando añadido móvil ahora. Añade ~3 semanas.
    Plan revisado para fin de la próxima semana.
  "linea-temporal.md": |
    # Línea Temporal Proyecto Phoenix (v2.1)
    Fase              | Objetivo | Estado
    Migración API     | 1 feb    | En Progreso (60%)
    Reconstrucción FE | 15 feb   | No Iniciada
    Sistema Diseño    | 8 feb    | Bloqueado
    Testing           | 22 feb   | No Iniciada
    Lanzamiento       | 1 mar    | En Riesgo
    *No incluye cambio de alcance de dashboard móvil*
  "retroalimentacion-equipo.txt": |
    Encuesta de Pulso del Equipo Anónima - Enero (5 respuestas)
    "El creep de alcance es un problema. Añadir características sin ajustar la línea temporal."
    "Buen progreso en API pero equipo muy estirado con transición de contratistas."
    "Diseño bloqueado — creando un efecto domino aguas abajo."
    "Moral OK pero preocupados por la fecha de marzo con añadido móvil."
    "Comunicación con cliente mejoró desde que Sarah tomó el rol de PM."

tools:
  - scratchpad_read
  - scratchpad_write

visibility:
  system_prompt: toggleable
  tool_calls: visible
  full_context: hidden
  token_count: hidden
  model_name: visible

sidebar:
  width: "45%"

script:
  - type: note
    text: "Este es un recorrido en vivo de la IA haciendo trabajo real. El explorador de archivos muestra cinco archivos de proyecto — notas de reunión, datos presupuestarios, correos de clientes, una línea temporal y resultados de encuesta del equipo. Desordenado, desorganizado, típico."

  - type: note
    text: "Ahora observa. Una instrucción. Sin plantillas, sin desglose paso a paso. Solo describe el resultado que necesitas."

  - type: user
    content: "Tengo una carpeta de proyecto con notas de reuniones, datos presupuestarios, correos de clientes, una línea temporal y retroalimentación del equipo. Necesito un informe ejecutivo de estado para el liderazgo para fin de día. Recopila los puntos clave, marca los riesgos y hazlo profesional."

  - type: assistant
    content: "Revisaré todos los materiales del proyecto y crearé un informe ejecutivo de estado profesional. Comenzaré leyendo cada archivo."

  - type: tool_call
    tool: scratchpad_read
    args:
      filename: "notas-reunion.txt"

  - type: tool_result
    tool: scratchpad_read
    content: |
      Proyecto Phoenix - Reunión de Equipo 15 de Enero
      Asistentes: Sarah (PM), Equipo de Desarrollo, Lisa (Diseño)
      - Migración de API 60% completa, objetivo 1 de febrero
      - Revisión de diseño bloqueada por directrices de marca
      - El cliente quiere añadir dashboard móvil a Fase 1
      - Tarifas de contratistas aumentaron 15% — preocupación presupuestaria
      - Sarah debe escalar riesgo de línea temporal a liderazgo
      - Próxima revisión de hito: 22 de enero

  - type: tool_call
    tool: scratchpad_read
    args:
      filename: "presupuesto-t1.csv"

  - type: tool_result
    tool: scratchpad_read
    content: |
      Categoria,Planificado,Real,Variacion
      Desarrollo,45000,38000,-7000
      Diseño,15000,12500,-2500
      Infraestructura,8000,9200,+1200
      Contratistas,20000,27500,+7500
      Testing,7000,4200,-2800
      Total,95000,91400,-3600

  - type: tool_call
    tool: scratchpad_read
    args:
      filename: "correos-cliente.txt"

  - type: tool_result
    tool: scratchpad_read
    content: |
      De: David Chen <d.chen@acmecorp.com>
      Fecha: 14 de enero
      Asunto: RE: Línea Temporal Proyecto Phoenix

      Sarah — La junta directiva quiere dashboard móvil en Fase 1.
      ¿Línea temporal revisada? Lo necesitamos para el 15 de marzo como máximo.
      ---
      De: Sarah Lopez <s.lopez@team.com>
      Fecha: 12 de enero

      David — Evaluando añadido móvil ahora. Añade ~3 semanas.
      Plan revisado para fin de la próxima semana.

  - type: tool_call
    tool: scratchpad_read
    args:
      filename: "linea-temporal.md"

  - type: tool_result
    tool: scratchpad_read
    content: |
      # Línea Temporal Proyecto Phoenix (v2.1)
      Fase              | Objetivo | Estado
      Migración API     | 1 feb    | En Progreso (60%)
      Reconstrucción FE | 15 feb   | No Iniciada
      Sistema Diseño    | 8 feb    | Bloqueado
      Testing           | 22 feb   | No Iniciada
      Lanzamiento       | 1 mar    | En Riesgo
      *No incluye cambio de alcance de dashboard móvil*

  - type: tool_call
    tool: scratchpad_read
    args:
      filename: "retroalimentacion-equipo.txt"

  - type: tool_result
    tool: scratchpad_read
    content: |
      Encuesta de Pulso del Equipo Anónima - Enero (5 respuestas)
      "El creep de alcance es un problema. Añadir características sin ajustar la línea temporal."
      "Buen progreso en API pero equipo muy estirado con transición de contratistas."
      "Diseño bloqueado — creando un efecto domino aguas abajo."
      "Moral OK pero preocupados por la fecha de marzo con añadido móvil."
      "Comunicación con cliente mejoró desde que Sarah tomó el rol de PM."

  - type: assistant
    content: "He revisado los cinco documentos. Varias claras tendencias: un riesgo de cambio de alcance por la solicitud del dashboard móvil, un bloqueo de diseño causando retrasos aguas abajo, y presión de costos de contratistas. Déjame sintetizar esto en un informe ejecutivo estructurado."

  - type: tool_call
    tool: scratchpad_write
    args:
      filename: "informe-estado-exec.md"
      content: |
        # Proyecto Phoenix — Informe de Estado Ejecutivo
        Preparado: 15 de enero de 2026

        ## Resumen
        El Proyecto Phoenix está 60% en la migración de API pero enfrenta riesgo de línea temporal por un cambio de alcance solicitado por el cliente (dashboard móvil). El gasto actual está debajo del presupuesto pero los sobrecostos de contratistas señalan presión.

        ## Métricas Clave
        - Presupuesto: 91.400 de 95.000 gastado (3,8% debajo)
        - Línea temporal: Lanzamiento original 1 de marzo — EN RIESGO
        - Alcance: Dashboard móvil añadido a Fase 1, aún no planificado

        ## Riesgos Críticos

        1. CAMBIO DE ALCANCE: Dashboard móvil añade ~3 semanas. Fecha límite del cliente 15 de marzo. Recomendar presentar línea temporal revisada con opciones escalonadas.

        2. BLOQUEO DE DISEÑO: Bloqueado por directrices de marca. Crea retrasos aguas abajo en frontend y testing. Escalar al cliente esta semana.

        3. COSTOS DE CONTRATISTAS: 7.500 sobre lo planificado (variación 37,5%). Se acelerará con alcance expandido. Revisar tarifas y considerar precio fijo para trabajo móvil.

        ## Salud del Equipo
        Estable pero preocupado por el creep de alcance y la presión de la línea temporal. Comunicación del PM con el cliente valorada positivamente.

        ## Próximos Pasos Recomendados
        1. Presentar línea temporal revisada al cliente para el 22 de enero
        2. Escalar bloqueo de directrices de marca
        3. Revisar estructura de costos de contratistas
        4. Programar sesión de escopado de dashboard móvil

  - type: tool_result
    tool: scratchpad_write
    content: "Escrito a informe-estado-exec.md (26 líneas)"

  - type: assistant
    content: "Listo. He creado informe-estado-exec.md — un informe de estado ejecutivo estructurado que cubre resumen del proyecto, métricas clave, tres riesgos críticos con recomendaciones, salud del equipo y próximos pasos priorizados. Listo para revisión del liderazgo."

  - type: note
    text: "Cinco archivos dispersos se convirtieron en un informe profesional. Sin copiar y pegar. Sin ida y vuelta. Sin plantillas. La IA leyó cada archivo, identificó las tendencias y entregó un documento terminado."

  - type: note
    text: "Nota lo que no hiciste: no especificaste qué archivos leer primero, cómo estructurar el informe, ni qué secciones incluir. Describiste el resultado. La IA descubrió el resto. Eso es delegación, no prompting. Eso es lo que la formación de hoy te enseñará."
```
