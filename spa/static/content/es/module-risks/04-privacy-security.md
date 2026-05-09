---
title: "Privacidad, Seguridad y el Informático Digital"
duration: "10m"
tags: [security, privacy, data-protection]
---

# Privacidad, Seguridad y el Informático Digital

La IA introduce una nueva categoría de riesgo de seguridad. Entenderlo no es opcional.

## La Realidad Actual

**El 38% de los empleados** comparte datos confidenciales con plataformas de IA sin aprobación.

Los datos de IBM de 2025 muestran que las brechas que involucran herramientas de IA no autorizadas cuestan un promedio de **$4.88 millones** — un 16% por encima del coste promedio de una brecha.

Esto no es hipotético. Está sucediendo ahora.

## El Riesgo de la IA Agente

Los chatbots de IA tradicionales son relativamente contenidos. Copias texto en una ventana del navegador. El riesgo se limita a lo que compartes explícitamente.

La IA agente opera de manera diferente.

**Los agentes heredan tus permisos de archivos.** Escanean todos los datos accesibles en tu sistema y red. Esto incluye documentos que no sabías que podías acceder, archivos sensibles enterrados en unidades compartidas e información confidencial fuera de tu área de trabajo directa.

El agente no pregunta "¿debería leer este archivo?" Lee todo lo que puede alcanzar para completar la tarea que le has delegado.

```callout
type: danger
title: "La Exposición que No Sabías que Existía"
content: "Tu agente de IA tiene acceso a cada archivo al que tú tienes acceso. Eso incluye la carpeta de finanzas a la que fuiste añadido accidentalmente, el directorio de RRHH que nunca te eliminó y los archivos de clientes de ese proyecto de hace dos años. Los agentes exponen la deriva de permisos."
```

![Ruta de informático digital](/content/module-risks/images/digital-insider.svg)

## El Marco de McKinsey: El Informático Digital

McKinsey recomienda tratar a los agentes de IA como **"informáticos digitales"** — empleados con acceso completo que requieren los mismos controles de seguridad que el personal humano.

Esto significa:
- Clasificación de datos antes del uso de IA (público, interno, confidencial, restringido)
- Políticas claras sobre qué herramientas de IA están aprobadas para qué tipos de datos
- Registros de auditoría del acceso de IA a información sensible
- Revisión regular de los permisos del agente

Gartner proyecta que para 2027, **uno de cada cuatro brechas empresariales** involucrará mal uso de la IA agente.

## Capas de Consumo vs Empresarial

La capa que uses importa.

**Herramientas de IA de consumo (ChatGPT Plus, Claude para uso personal):** Pueden entrenar con tus datos por defecto. Historial de conversaciones almacenado. Controles limitados de retención de datos.

**Herramientas de IA empresarial (ChatGPT Enterprise, Claude for Work):** No entrenan con tus datos. Garantías contractuales de protección de datos. Registros de auditoría. Cumplimiento GDPR/SOC2.

La diferencia de precio refleja la diferencia en protección de datos.

```callout
type: tip
title: "La Pregunta de Clasificación"
content: "Antes de subir cualquier archivo a una herramienta de IA, pregúntate: '¿Estaría cómodo si esto apareciera en los datos de entrenamiento de otra persona?' Si la respuesta es no, usa una capa empresarial o no lo subas."
```

## Directrices Prácticas

**La clasificación de datos es la base.** Las organizaciones deben definir:
- Qué datos puede procesar la IA (documentos internos, borradores, investigación)
- Qué datos la IA no puede procesar (PII de clientes, registros financieros, secretos comerciales)
- Qué herramientas de IA están aprobadas para qué clasificaciones

**La selección de herramientas importa.** La cuenta gratuita de ChatGPT tiene un manejo de datos diferente al de ChatGPT Enterprise. Lee los términos. Entiende lo que "no entrenamos con tus datos" significa realmente.

**Audita el acceso de tu agente.** Si estás usando herramientas agentes, revisa lo que pueden alcanzar. Elimina el acceso a las carpetas que no deberían escanear. Organiza tu sistema de archivos alrededor de la sensibilidad de los datos.

**La formación por sí sola es insuficiente.** El treinta y ocho por ciento que comparte datos confidenciales sugiere que la formación no funciona. Los controles técnicos — lista blanca de herramientas, permisos de archivos, reglas DLP — importan más que esperar un comportamiento perfecto de los empleados.

## La Responsabilidad Empresarial

Las organizaciones que despliegan IA deben:
- Proporcionar herramientas aprobadas con las protecciones de datos apropiadas
- Hacer que las herramientas aprobadas sean mejores que las alternativas gratuitas que los empleados usarían de otro modo
- Construir la seguridad en los flujos de trabajo en lugar de depender del cumplimiento individual
- Tratar el despliegue de agentes como una decisión de arquitectura de seguridad, no solo como una herramienta de productividad

El riesgo es real. Las contramedidas son conocidas. La pregunta es si tu organización las ha implementado antes de la brecha, no después.

## La Pregunta de Clasificación en la Práctica

Observa lo que sucede cuando alguien comparte archivos con una IA sin pensar en la clasificación. Luego observa cómo la IA marca correctamente el problema.

```agent
id: data-classification-demo
title: "Lo Que Debería (y No Debería) Ir a la IA"
model_label: "Claude"

system: |
  Eres un asistente empresarial útil. Antes de procesar archivos,
  revisa sus contenidos y marca cualquier preocupación de sensibilidad
  de los datos. Clasifica los datos como: PÚBLICO, INTERNO,
  CONFIDENCIAL o RESTRINGIDO.

scratchpad:
  "project-update.txt": |
    Proyecto Aurora — Actualización Semanal
    Estado: En camino para la entrega de marzo
    Velocidad del sprint: 42 puntos (subido de 38)
    Bloqueante: Esperando revisión legal de los términos de la API
    Próximo hito: Lanzamiento beta el 15 de febrero
  "customer-data.csv": |
    Nombre,Email,Empresa,ARR,Fin_Contrato,NPS,Health_Score
    James Morton,j.morton@meridianhealth.co.uk,Meridian Health,£48000,2026-06-30,9,92
    Priya Sharma,p.sharma@teklogic.com,TekLogic,£125000,2026-12-31,7,78
    David Chen,d.chen@finserv.com,FinServ Capital,£210000,2026-03-31,4,45
    Sarah Williams,s.williams@govdigital.gov.uk,Gov Digital Services,£89000,2027-01-15,8,85
  "salary-benchmarks.txt": |
    Benchmark Salariales 2026 — CONFIDENCIAL
    Ingeniero Senior: £85K-£110K (tasa de mercado subiendo 8%)
    Director de Ingeniería: £105K-£135K
    VP de Ingeniería: £145K-£180K
    Nota: Nuestras ofertas actuales están un 12% por debajo del punto medio del mercado
    La Junta aprobó un presupuesto de aumento del 15% para el plan de retención del Q2
    No compartir con candidatos a contratación ni reclutadores externos

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
    text: "Alguien está a punto de subir tres archivos a una herramienta de IA. Uno está bien. Uno es arriesgado. Uno nunca debería compartirse con una IA de consumo. Observa el proceso de clasificación."

  - type: user
    content: "Tengo tres archivos para ti. Léelos todos y ayúdame a prepararme para mi reunión de liderazgo."

  - type: tool_call
    tool: scratchpad_read
    args:
      filename: "project-update.txt"

  - type: tool_result
    tool: scratchpad_read
    content: |
      Proyecto Aurora — Actualización Semanal
      Estado: En camino para la entrega de marzo
      Velocidad del sprint: 42 puntos (subido de 38)
      Bloqueante: Esperando revisión legal de los términos de la API
      Próximo hito: Lanzamiento beta el 15 de febrero

  - type: tool_call
    tool: scratchpad_read
    args:
      filename: "customer-data.csv"

  - type: tool_result
    tool: scratchpad_read
    content: |
      Nombre,Email,Empresa,ARR,Fin_Contrato,NPS,Health_Score
      James Morton,j.morton@meridianhealth.co.uk,Meridian Health,£48000,2026-06-30,9,92
      Priya Sharma,p.sharma@teklogic.com,TekLogic,£125000,2026-12-31,7,78
      David Chen,d.chen@finserv.com,FinServ Capital,£210000,2026-03-31,4,45
      Sarah Williams,s.williams@govdigital.gov.uk,Gov Digital Services,£89000,2027-01-15,8,85

  - type: tool_call
    tool: scratchpad_read
    args:
      filename: "salary-benchmarks.txt"

  - type: tool_result
    tool: scratchpad_read
    content: |
      Benchmark Salariales 2026 — CONFIDENCIAL
      Ingeniero Senior: £85K-£110K (tasa de mercado subiendo 8%)
      Director de Ingeniería: £105K-£135K
      VP de Ingeniería: £145K-£180K
      Nota: Nuestras ofertas actuales están un 12% por debajo del punto medio del mercado
      La Junta aprobó un presupuesto de aumento del 15% para el plan de retención del Q2
      No compartir con candidatos a contratación ni reclutadores externos

  - type: assistant
    content: "**Antes de procesar esto — necesito marcar preocupaciones de sensibilidad de datos.**\n\n**Archivo 1: project-update.txt** — INTERNO\nEstado del proyecto con métricas de sprint. Baja sensibilidad. Seguro para procesar con cualquier herramienta de IA aprobada.\n\n**Archivo 2: customer-data.csv** — CONFIDENCIAL / RESTRINGIDO\nContiene datos personales (nombres, direcciones de email) y datos comerciales (valores de contrato, fechas de renovación). Esto incluye:\n- PII sujeto al GDPR (nombres, direcciones de email)\n- Información comercialmente sensible (valores ARR, puntuaciones de salud)\n- Un cliente gubernamental (Gov Digital Services) que puede tener requisitos adicionales de manejo de datos\n\n**Este archivo NO debería procesarse en una capa de IA de consumo.** Se requiere la capa empresarial con acuerdo de procesamiento de datos.\n\n**Archivo 3: salary-benchmarks.txt** — RESTRINGIDO\nEl archivo en sí está marcado como 'CONFIDENCIAL' con instrucciones explícitas: 'No compartir con candidatos a contratación ni reclutadores externos.' Contiene planes de presupuesto aprobados por la Junta y datos de posicionamiento competitivo.\n\n**Este archivo no debería subirse a ninguna herramienta de IA externa** a menos que tu organización haya aprobado específicamente la herramienta para datos de capa restringida.\n\n**Recomendación:** Puedo trabajar con el archivo 1 inmediatamente. Los archivos 2 y 3 requieren que confirmes que estás usando una herramienta de IA de capa empresarial con protecciones de datos apropiadas antes de proceder."

  - type: note
    text: "La IA clasificó correctamente los tres archivos y se negó a proceder con los sensibles hasta que el usuario confirme la capa de la herramienta. Pero aquí está la idea clave: la mayoría de las herramientas de IA NO harán esto. Procesarán felizmente tu PII de clientes y datos salariales sin una palabra de advertencia."

  - type: note
    text: "Esta demo muestra el comportamiento ideal. En la realidad, la responsabilidad de clasificación es tuya. Antes de subir cualquier archivo, aplica la prueba: '¿Estaría cómodo si esto apareciera en los datos de entrenamiento de otra persona?' La actualización del proyecto ¿bien? Emails de clientes y valores de contrato ? Piensa detenidamente. Estrategia salarial aprobada por la Junta ? Nunca en una capa de consumo."
```

```quiz
id: privacy-agent-risk
type: multiple-choice
question: "¿Por qué las herramientas de IA agente representan un riesgo de seguridad diferente en comparación con los chatbots tradicionales?"
options:
  - "Procesan datos en servidores externos mientras los chatbots funcionan localmente"
  - "Heredan tus permisos de archivos y pueden acceder a todos los datos a los que tú puedes acceder"
  - "Comparten datos entre múltiples usuarios para mejorar el rendimiento"
answer: 1
explanation: "Las herramientas de IA agente heredan los permisos de archivos del usuario y escanean todos los datos accesibles para completar tareas. Esto significa que pueden acceder a archivos sensibles que el usuario podría ni siquiera saber que tiene acceso, creando una exposición que los chatbots tradicionales de copiar y pegar no tienen."
```
