---
title: "El cambio de paradigma: de chat a archivos"
duration: "10m"
tags: [archivos, paradigma, entregables]
---

# El cambio de paradigma: de chat a archivos

Aquí es donde pasamos de pensar en la IA como chat a pensar en la IA como una herramienta que produce entregables.

## Forma antigua vs nueva forma

**Forma antigua: flujo de trabajo por chat**
1. Chatear con la IA
2. Copiar el texto de la respuesta
3. Pegarlo en un documento
4. Darle formato
5. Corregir problemas
6. Repetir

**Forma nueva: flujo de trabajo por archivos**
1. Proporcionar archivos de entrada a la IA
2. Recibir archivos de salida
3. Listo

```callout
type: tip
title: "La distinción clave"
content: "El chat es para comunicación: aclarar lo que necesitas, revisar el progreso. Los archivos son para el trabajo real: documentos, hojas de cálculo, presentaciones, código, análisis."
```

## Por qué esto lo cambia todo

Cuando piensas en archivos en lugar de chat:

**Trabajas con entregables reales**
No texto que tienes que ensamblar para que sea útil.

**Mantiene un formato adecuado**
Las fórmulas de Excel funcionan. Las presentaciones de PowerPoint se presentan correctamente. El código se ejecuta.

**Puedes iterar correctamente**
Editas el documento real, no reinicias el chat.

**Produces resultados profesionales**
Los archivos tienen aspecto profesional. Las exportaciones de chat, no.

![Flujo de archivos: archivos de entrada al procesamiento de IA para archivos de salida](/content/module-files/images/file-workflow.svg)

## La implicación práctica

Si estás copiando y pegando desde el chat a Word, Excel o PowerPoint, lo estás haciendo mal.

Deberías recibir el archivo directamente. Pide el formato de archivo que necesitas y la IA lo creará.

## Ver la diferencia

Misma tarea, dos enfoques. Primero la forma por chat, luego la forma por archivos. La diferencia es notable.

```agent
id: chat-vs-files-demo
title: "Flujo de chat vs flujo de archivos"
model_label: "Claude"

system: |
  Eres un analista de negocios que ayuda a preparar materiales para reuniones.
  Trabaja con archivos cuando se te proporcionen.

scratchpad:
  "actualizaciones-equipo.txt": |
    Ingeniería: Lanzó autenticación v2, comenzó refactorización de API. 2 errores P1 abiertos.
    Diseño: Completó mockups del rediseño móvil. Pruebas de usuario la próxima semana.
    Marketing: Campaña lanzada el lunes. 12K impresiones el día 1. Por debajo del objetivo (15K).
    Ventas: Cerró 3 acuerdos empresariales (180K dólares en total). Pipeline de 420K.
    Soporte: Volumen de tickets subió un 22%. Tiempo medio de respuesta 4,2 horas (objetivo: 3 horas).

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
    text: "Primero, el enfoque por chat. Sin archivos. Solo escribir una petición y recibir texto que tendrías que copiar y pegar en un documento."

  - type: user
    content: "Tengo una reunión de equipo en 30 minutos y necesito puntos de conversación. ¿Puedes resumir lo que está pasando en todos los departamentos?"

  - type: assistant
    content: "Puedo ayudarte, pero necesitaría saber qué está pasando en cada departamento. ¿Podrías compartir las actualizaciones de tu equipo?"

  - type: user
    content: "Ingeniería lanzó autenticación v2 y comenzó una refactorización de API con 2 errores P1. Diseño terminó los mockups móviles, pruebas de usuario la próxima semana. La campaña de marketing obtuvo 12K impresiones el día 1 pero el objetivo era 15K. Ventas cerró 3 acuerdos empresariales por 180K en total con un pipeline de 420K. Soporte tiene un aumento del 22% en tickets y el tiempo de respuesta es de 4,2 horas frente al objetivo de 3 horas."

  - type: assistant
    content: "Aquí tienes un resumen para tu reunión:\n\n**Ingeniería:** Autenticación v2 lanzada. Refactorización de API iniciada. Nota: 2 errores P1 necesitan atención.\n\n**Diseño:** Mockups del rediseño móvil completados. Pruebas de usuario programadas para la próxima semana.\n\n**Marketing:** Campaña activa desde el lunes. 12K impresiones el día 1, por debajo del objetivo de 15K. Puede necesitar ajustes.\n\n**Ventas:** Buen trimestre. 3 acuerdos empresariales cerrados (180K dólares). Pipeline sólido con 420K.\n\n**Soporte:** Volumen aumentado un 22%. Tiempo de respuesta de 4,2 horas frente al objetivo de 3 horas, esto necesita discusión.\n\n¿Te gustaría que lo formatee de otra manera?"

  - type: note
    text: "Tuviste que escribir manualmente todos los datos en el chat. La IA te dio texto en una ventana de chat. Ahora necesitarías copiarlo, pegarlo en un documento, darle formato, quizás añadir un título, ajustar el diseño. Ese es el flujo de trabajo por chat. Ahora observa el enfoque por archivos."

  - type: clear
    reset_scratchpad: false
    note: "Ahora el enfoque por archivos. Mismos datos, pero en un archivo."

  - type: user
    content: "Lee el archivo de actualizaciones del equipo y crea un documento de orden del día para la reunión con formato. Incluye indicadores de estado: verde para en camino, ámbar para atención, rojo para necesita acción."

  - type: tool_call
    tool: scratchpad_read
    args:
      filename: "actualizaciones-equipo.txt"

  - type: tool_result
    tool: scratchpad_read
    content: |
      Ingeniería: Lanzó autenticación v2, comenzó refactorización de API. 2 errores P1 abiertos.
      Diseño: Completó mockups del rediseño móvil. Pruebas de usuario la próxima semana.
      Marketing: Campaña lanzada el lunes. 12K impresiones el día 1. Por debajo del objetivo (15K).
      Ventas: Cerró 3 acuerdos empresariales (180K dólares en total). Pipeline de 420K.
      Soporte: Volumen de tickets subió un 22%. Tiempo medio de respuesta 4,2 horas (objetivo: 3 horas).

  - type: tool_call
    tool: scratchpad_write
    args:
      filename: "orden-del-dia-reunion-equipo.md"
      content: |
        # Orden del día — Reunión semanal de equipo
        Fecha: 15 de enero de 2026

        ## Estado por departamento

        ### Ingeniería 🟡 ÁMBAR
        - **Lanzado:** Autenticación v2 (completo)
        - **En progreso:** Refactorización de API (iniciado)
        - **Acción necesaria:** 2 errores P1 abiertos — asignar responsables y fechas objetivo

        ### Diseño 🟢 VERDE
        - **Completado:** Mockups del rediseño móvil
        - **Siguiente:** Pruebas de usuario (próxima semana)

        ### Marketing 🟡 ÁMBAR
        - **Campaña activa:** Lanzamiento el lunes
        - **Impresiones día 1:** 12.000 (objetivo: 15.000) — 20% por debajo del objetivo
        - **Discusión:** ¿Ajustar segmentación o creatividades?

        ### Ventas 🟢 VERDE
        - **Cerrado:** 3 acuerdos empresariales (180K dólares en total)
        - **Pipeline:** 420K en oportunidades activas
        - **Nota:** Buen rendimiento, en camino

        ### Soporte 🔴 ROJO
        - **Volumen de tickets:** Aumento del 22% — investigar causa
        - **Tiempo de respuesta:** 4,2 horas (objetivo: 3 horas)
        - **Acción necesaria:** Revisión de plantilla o mejora de procesos

        ## Temas de discusión
        1. Errores P1 de Ingeniería — responsables y cronograma
        2. Ajuste de campaña de marketing — ¿aprobar segmentación revisada?
        3. Capacidad de Soporte — solución a corto plazo y causa raíz

  - type: tool_result
    tool: scratchpad_write
    content: "Escrito en orden-del-dia-reunion-equipo.md (32 líneas)"

  - type: assistant
    content: "He creado orden-del-dia-reunion-equipo.md — una orden del día para la reunión con indicadores de estado RAG. Tres elementos marcados para discusión: los errores P1, el bajo rendimiento de marketing y la capacidad de soporte. El archivo está listo para compartir o presentar."

  - type: note
    text: "Sin entrada manual de datos. Sin copiar y pegar. Sin reformatear. La IA leyó el archivo de origen, añadió estructura y análisis (los indicadores de estado, los temas de discusión) y creó un documento terminado. Ese es el cambio de paradigma: de hablar por chat sobre tu trabajo a archivos que SON tu trabajo."
```

```quiz
id: files-paradigm
type: multiple-choice
question: "¿Cuál es la diferencia clave entre los flujos de trabajo de IA basados en chat y los basados en archivos?"
options:
  - "Los flujos por chat producen texto que copias y pegas; los flujos por archivos producen entregables útiles listos para abrir en sus aplicaciones nativas"
  - "Los flujos basados en archivos solo funcionan con ciertos modelos de IA que soportan creación de archivos"
  - "El chat es mejor para tareas complejas porque puedes iterar en la conversación"
answer: 0
explanation: "El cambio fundamental es pasar de recibir texto de chat (que luego copias, pegas y formateas) a recibir archivos reales (documentos, hojas de cálculo, presentaciones) listos para usar o editar directamente. Esto elimina el trabajo de formateo y ensamblaje."
```
