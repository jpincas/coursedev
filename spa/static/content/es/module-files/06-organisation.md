---
title: "Tu sistema de archivos listo para IA"
duration: "10m"
tags: [organizacion, nombres, compartir, corporativo]
---

# Tu sistema de archivos listo para IA

Un sistema de archivos bien organizado no es solo limpio. Es la base para cada interacción con IA que tendrás.

Cada vez que le pides a la IA que 'lea estos archivos' o 'trabaje con estos datos', la calidad de lo que produce la IA depende en parte de qué tan bien están nombrados, organizados y formateados esos archivos.

## Convenciones de Nombres de Archivos

Los malos nombres de archivos desperdician contexto y confunden tanto a ti como a la IA.

**Nombres malos:**
```
datos (3).xlsx
Nuevo Documento.docx
Captura de Pantalla 2026-02-07 a las 10.23.45.png
definitivo_definitivo_v2_ACTUAL_DEFINITIVO.docx
notas.txt
```

**Nombres buenos:**
```
ventas-t4-2025-por-region.csv
brief-cliente-proyecto-henderson.md
notas-reunion-2026-02-07-equipo-producto.md
directrices-de-marca-v3.md
analisis-competencia-acme-2026.md
```

**Reglas:**
- Usar minúsculas con guiones (no espacios ni guiones bajos)
- Incluir el tema, no solo el tipo
- Incluir fechas en formato AAAA-MM-DD cuando sea relevante
- Incluir números de versión cuando se itera
- Ser lo suficientemente descriptivo como para que el nombre por sí solo te diga qué hay dentro

```callout
type: tip
title: "El nombre es contexto"
content: "La IA lee los nombres de los archivos. Un archivo llamado 'ventas-t4-2025-por-region.csv' le da contexto a la IA incluso antes de abrir el archivo. Un archivo llamado 'datos.csv' no le dice nada."
```

## Estructura de Carpetas

Organiza los archivos por proyecto o dominio, no por tipo de archivo.

**Mala estructura:**
```
Documentos/
  hojas-de-calculo/
    archivo1.xlsx
    archivo2.xlsx
  presentaciones/
    presentacion1.pptx
```

**Buena estructura:**
```
proyectos/
  proyecto-henderson/
    entrada/
      brief-cliente.md
      requisitos.csv
      directrices-marca.md
    salida/
      documento-estrategia.md
      borrador-propuesta.md
      presentacion-final.pptx
    referencia/
      analisis-competencia.md
      propuestas-anteriores/
```

**Principios:**
- Agrupar por proyecto o tema
- Separar archivos de entrada de archivos de salida
- Mantener materiales de referencia accesibles
- Usar una distinción `trabajo/` vs `final/` para borradores

## Limpieza de Archivos Impulsada por IA

Uno de los usos más satisfactorios de la IA: tomar un montón desordenado de archivos y organizarlos automáticamente.

```agent
id: file-cleanup-demo
title: "La IA Organiza tus Archivos"
model_label: "Claude"

system: |
  Eres un asistente de organización de archivos. Ayudas a profesionales
  a limpiar y organizar sus archivos con convenciones de nombres claras
  y estructuras lógicas de carpetas.

scratchpad:
  "datos (2).csv": |
    Mes,Ingreses,Gastos
    Ene,45000,32000
    Feb,52000,35000
    Mar,48000,33000
  "IMG_4521.txt": |
    Notas de la reunión de sincronización de producto
    - Fecha de lanzamiento confirmada: 15 de marzo
    - Presupuesto aprobado: 120K £
    - Necesito revisión de diseño antes del 20 de febrero
  "informe final.txt": |
    Resumen del Rendimiento T4
    Ingresos: 2,1M £ (subió un 18%)
    Nuevos clientes: 340
    Abandono: 2,3%
    NPS: 52
  "cosas.txt": |
    Investigación de precios de competidores
    Acme Pro: 49£/usuario/mes
    Beta Corp: 39£/usuario/mes
    Gamma Inc: 59£/usuario/mes (incluye soporte)
    Nuestro precio: 45£/usuario/mes
  "doc.txt": |
    Directrices de tono de marca
    Tono: Profesional pero cálido
    Evitar: Jerga, voz pasiva
    Siempre: Usar ejemplos concretos
    Audiencia: Líderes de empresas de mercado medio
  "notas feb.txt": |
    Retroalimentación de equipo 7 de febrero
    Qué salió bien: despliegues más rápidos, mejor QA
    Qué no salió bien: retrasos en comunicación con clientes
    Acción: actualizaciones semanales a clientes a partir del próximo sprint
  "cosa presupuesto.csv": |
    Categoria,Presupuesto_T1,Real_T1,Varianza
    Ingenieria,180000,175000,-5000
    Marketing,95000,102000,+7000
    Operaciones,45000,43000,-2000
  "nuevo doc.txt": |
    Proyecto Phoenix — Cronograma
    Fase 1: Investigación (febrero)
    Fase 2: Diseño (marzo)
    Fase 3: Construcción (abril-mayo)
    Fase 4: Lanzamiento (junio)
  "IMG_3847.png": |
    [Datos de imagen binaria — foto del equipo de producto del offsite del T4]
  "Captura 2026-01-15 a las 14.32.07.png": |
    [Datos de imagen binaria — captura de pantalla de la página principal del sitio web del competidor]
  "foto (1).jpg": |
    [Datos de imagen binaria — diagrama de pizarra blanca de la sesión de estrategia]
  "Copia de números T3.csv": |
    Producto,U_T3,Ingresos_T3
    Dashboard Pro,1250,187500
    Suite de Analíticas,890,133500
    Acceso API,2100,105000
    Soporte Plus,650,97500

tools:
  - scratchpad_read
  - scratchpad_write
  - list_files
  - move_file
  - create_folder

visibility:
  system_prompt: hidden
  tool_calls: visible
  full_context: hidden

sidebar:
  width: "45%"
  start_open: true

script:
  - type: note
    text: "Mira este desastre. Doce archivos con nombres como 'cosas.txt', 'datos (2).csv', 'IMG_3847.png', 'Copia de números T3.csv'. Una mezcla de texto, datos e imágenes. No puedes saber qué contienen ninguno de ellos solo por el nombre. Esto es lo que parece una carpeta de trabajo real después de seis meses de guardar archivos sin pensar."

  - type: user
    content: "Tengo un desastre de archivos — documentos de texto, hojas de cálculo, imágenes, todos con nombres terribles. ¿Puedes leer todo, renombrarlos con nombres descriptivos claros, organizarlos en carpetas lógicas y crear un índice de resumen?"

  - type: tool_call
    tool: list_files
    args: {}

  - type: tool_result
    tool: list_files
    content: |
      12 archivos:
      - cosa presupuesto.csv
      - Copia de números T3.csv
      - datos (2).csv
      - doc.txt
      - informe final.txt
      - IMG_3847.png
      - IMG_4521.txt
      - nuevo doc.txt
      - notas feb.txt
      - foto (1).jpg
      - Captura 2026-01-15 a las 14.32.07.png
      - cosas.txt

  - type: tool_call
    tool: scratchpad_read
    args:
      filename: "datos (2).csv"

  - type: tool_result
    tool: scratchpad_read
    content: |
      Mes,Ingresos,Gastos
      Ene,45000,32000
      Feb,52000,35000
      Mar,48000,33000

  - type: tool_call
    tool: scratchpad_read
    args:
      filename: "Copia de números T3.csv"

  - type: tool_result
    tool: scratchpad_read
    content: |
      Producto,U_T3,Ingresos_T3
      Dashboard Pro,1250,187500
      Suite de Analíticas,890,133500
      Acceso API,2100,105000
      Soporte Plus,650,97500

  - type: tool_call
    tool: scratchpad_read
    args:
      filename: "IMG_4521.txt"

  - type: tool_result
    tool: scratchpad_read
    content: |
      Notas de la reunión de sincronización de producto
      - Fecha de lanzamiento confirmada: 15 de marzo
      - Presupuesto aprobado: 120K £
      - Necesito revisión de diseño antes del 20 de febrero

  - type: tool_call
    tool: scratchpad_read
    args:
      filename: "cosas.txt"

  - type: tool_result
    tool: scratchpad_read
    content: |
      Investigación de precios de competidores
      Acme Pro: 49£/usuario/mes
      Beta Corp: 39£/usuario/mes
      Gamma Inc: 59£/usuario/mes (incluye soporte)
      Nuestro precio: 45£/usuario/mes

  - type: tool_call
    tool: scratchpad_read
    args:
      filename: "doc.txt"

  - type: tool_result
    tool: scratchpad_read
    content: |
      Directrices de tono de marca
      Tono: Profesional pero cálido
      Evitar: Jerga, voz pasiva
      Siempre: Usar ejemplos concretos
      Audiencia: Líderes de empresas de mercado medio

  - type: tool_call
    tool: scratchpad_read
    args:
      filename: "informe final.txt"

  - type: tool_result
    tool: scratchpad_read
    content: |
      Resumen del Rendimiento T4
      Ingresos: 2,1M £ (subió un 18%)
      Nuevos clientes: 340
      Abandono: 2,3%
      NPS: 52

  - type: tool_call
    tool: scratchpad_read
    args:
      filename: "notas feb.txt"

  - type: tool_result
    tool: scratchpad_read
    content: |
      Retroalimentación de equipo 7 de febrero
      Qué salió bien: despliegues más rápidos, mejor QA
      Qué no salió bien: retrasos en comunicación con clientes
      Acción: actualizaciones semanales a clientes a partir del próximo sprint

  - type: tool_call
    tool: scratchpad_read
    args:
      filename: "cosa presupuesto.csv"

  - type: tool_result
    tool: scratchpad_read
    content: |
      Categoria,Presupuesto_T1,Real_T1,Varianza
      Ingenieria,180000,175000,-5000
      Marketing,95000,102000,+7000
      Operaciones,45000,43000,-2000

  - type: tool_call
    tool: scratchpad_read
    args:
      filename: "nuevo doc.txt"

  - type: tool_result
    tool: scratchpad_read
    content: |
      Proyecto Phoenix — Cronograma
      Fase 1: Investigación (febrero)
      Fase 2: Diseño (marzo)
      Fase 3: Construcción (abril-mayo)
      Fase 4: Lanzamiento (junio)

  - type: tool_call
    tool: scratchpad_read
    args:
      filename: "IMG_3847.png"

  - type: tool_result
    tool: scratchpad_read
    content: "[Datos de imagen binaria — foto del equipo de producto del offsite del T4]"

  - type: tool_call
    tool: scratchpad_read
    args:
      filename: "Captura 2026-01-15 a las 14.32.07.png"

  - type: tool_result
    tool: scratchpad_read
    content: "[Datos de imagen binaria — captura de pantalla de la página principal del sitio web del competidor]"

  - type: tool_call
    tool: scratchpad_read
    args:
      filename: "foto (1).jpg"

  - type: tool_result
    tool: scratchpad_read
    content: "[Datos de imagen binaria — diagrama de pizarra blanca de la sesión de estrategia]"

  - type: assistant
    content: "He leído los 12 archivos. Esto es con lo que trabajamos:\n\n**Documentos de texto (7):**\n- Notas de reuniones, investigación de competidores, directrices de marca, informe T4, retroalimentación de equipo, cronograma del proyecto, más uno mal etiquetado como IMG_4521.txt\n\n**Archivos de datos (3):**\n- Ingresos/gastos T1, presupuesto T1 vs real, ventas de productos T3\n\n**Imágenes (3):**\n- Foto del equipo, captura de pantalla del sitio del competidor, diagrama de pizarra blanca\n\nNota: 'IMG_4521.txt' es en realidad notas de reuniones, no una imagen — alguien lo guardó con la convención incorrecta. Organizaré todo en carpetas lógicas ahora."

  - type: tool_call
    tool: create_folder
    args:
      path: "datos"

  - type: tool_result
    tool: create_folder
    content: "Carpeta creada: datos"

  - type: tool_call
    tool: create_folder
    args:
      path: "reuniones"

  - type: tool_result
    tool: create_folder
    content: "Carpeta creada: reuniones"

  - type: tool_call
    tool: create_folder
    args:
      path: "referencia"

  - type: tool_result
    tool: create_folder
    content: "Carpeta creada: referencia"

  - type: tool_call
    tool: create_folder
    args:
      path: "imagenes"

  - type: tool_result
    tool: create_folder
    content: "Carpeta creada: imagenes"

  - type: tool_call
    tool: move_file
    args:
      source: "datos (2).csv"
      destination: "datos/ingresos-y-gastos-t1.csv"

  - type: tool_result
    tool: move_file
    content: "Movido: datos (2).csv -> datos/ingresos-y-gastos-t1.csv"

  - type: tool_call
    tool: move_file
    args:
      source: "cosa presupuesto.csv"
      destination: "datos/presupuesto-t1-vs-real.csv"

  - type: tool_result
    tool: move_file
    content: "Movido: cosa presupuesto.csv -> datos/presupuesto-t1-vs-real.csv"

  - type: tool_call
    tool: move_file
    args:
      source: "Copia de números T3.csv"
      destination: "datos/ventas-productos-t3.csv"

  - type: tool_result
    tool: move_file
    content: "Movido: Copia de números T3.csv -> datos/ventas-productos-t3.csv"

  - type: tool_call
    tool: move_file
    args:
      source: "informe final.txt"
      destination: "datos/resumen-rendimiento-t4.md"

  - type: tool_result
    tool: move_file
    content: "Movido: informe final.txt -> datos/resumen-rendimiento-t4.md"

  - type: tool_call
    tool: move_file
    args:
      source: "IMG_4521.txt"
      destination: "reuniones/sincronizacion-producto-planificacion-lanzamiento.md"

  - type: tool_result
    tool: move_file
    content: "Movido: IMG_4521.txt -> reuniones/sincronizacion-producto-planificacion-lanzamiento.md"

  - type: tool_call
    tool: move_file
    args:
      source: "notas feb.txt"
      destination: "reuniones/retroalimentacion-2026-02-07.md"

  - type: tool_result
    tool: move_file
    content: "Movido: notas feb.txt -> reuniones/retroalimentacion-2026-02-07.md"

  - type: tool_call
    tool: move_file
    args:
      source: "cosas.txt"
      destination: "referencia/analisis-precios-competidores.md"

  - type: tool_result
    tool: move_file
    content: "Movido: cosas.txt -> referencia/analisis-precios-competidores.md"

  - type: tool_call
    tool: move_file
    args:
      source: "doc.txt"
      destination: "referencia/directrices-tono-marca.md"

  - type: tool_result
    tool: move_file
    content: "Movido: doc.txt -> referencia/directrices-tono-marca.md"

  - type: tool_call
    tool: move_file
    args:
      source: "nuevo doc.txt"
      destination: "referencia/cronograma-proyecto-phoenix.md"

  - type: tool_result
    tool: move_file
    content: "Movido: nuevo doc.txt -> referencia/cronograma-proyecto-phoenix.md"

  - type: tool_call
    tool: move_file
    args:
      source: "IMG_3847.png"
      destination: "imagenes/foto-equipo-offsite-t4.png"

  - type: tool_result
    tool: move_file
    content: "Movido: IMG_3847.png -> imagenes/foto-equipo-offsite-t4.png"

  - type: tool_call
    tool: move_file
    args:
      source: "Captura 2026-01-15 a las 14.32.07.png"
      destination: "imagenes/sitio-web-competidor-portada.png"

  - type: tool_result
    tool: move_file
    content: "Movido: Captura 2026-01-15 a las 14.32.07.png -> imagenes/sitio-web-competidor-portada.png"

  - type: tool_call
    tool: move_file
    args:
      source: "foto (1).jpg"
      destination: "imagenes/sesion-estrategia-pizarra-blanca.jpg"

  - type: tool_result
    tool: move_file
    content: "Movido: foto (1).jpg -> imagenes/sesion-estrategia-pizarra-blanca.jpg"

  - type: tool_call
    tool: scratchpad_write
    args:
      filename: "indice-archivos.md"
      content: |
        # Índice de Archivos

        ## datos/
        - **ingresos-y-gastos-t1.csv** — Ingresos y gastos mensuales para el T1 (ene-mar)
        - **presupuesto-t1-vs-real.csv** — Presupuesto T1 vs gasto real por departamento (Ingeniería, Marketing, Operaciones)
        - **ventas-productos-t3.csv** — Ventas de unidades e ingresos por línea de producto en T3
        - **resumen-rendimiento-t4.md** — Resultados destacados T4: 2,1M £ de ingresos, 340 nuevos clientes, 2,3% de abandono, NPS 52

        ## reuniones/
        - **sincronizacion-producto-planificacion-lanzamiento.md** — Sincronización de producto: lanzamiento 15 de marzo confirmado, 120K £ de presupuesto, revisión de diseño necesaria antes del 20 de febrero
        - **retroalimentacion-2026-02-07.md** — Retroalimentación de equipo: despliegues más rápidos y mejor QA positivo; la comunicación con clientes necesita actualizaciones semanales

        ## referencia/
        - **analisis-precios-competidores.md** — Comparación de precios: Acme 49£, Beta 39£, Gamma 59£ (incl. soporte), Nosotros 45£
        - **directrices-tono-marca.md** — Tono: profesional pero cálido. Evitar jerga. Usar ejemplos concretos. Audiencia: líderes de mercado medio
        - **cronograma-proyecto-phoenix.md** — 4 fases: Investigación (feb), Diseño (mar), Construcción (abril-mayo), Lanzamiento (junio)

        ## imagenes/
        - **foto-equipo-offsite-t4.png** — Foto del equipo de producto del evento offsite del T4
        - **sitio-web-competidor-portada.png** — Captura de pantalla del sitio del competidor para referencia
        - **sesion-estrategia-pizarra-blanca.jpg** — Diagrama de pizarra blanca de la sesión de planificación estratégica

  - type: tool_result
    tool: scratchpad_write
    content: "Escrito en indice-archivos.md (20 líneas)"

  - type: assistant
    content: "Hecho. Doce archivos desordenos están ahora en cuatro carpetas lógicas con nombres claros:\n\n- **datos/** — datos financieros e informes de rendimiento (4 archivos)\n- **reuniones/** — notas de reuniones y retroalimentación (2 archivos)\n- **referencia/** — directrices, investigación y planes de proyecto (3 archivos)\n- **imagenes/** — fotos y capturas de pantalla con nombres descriptivos (3 archivos)\n\nMás un indice-archivos.md que resume todo. Fíjate que 'IMG_4521.txt' era en realidad notas de reuniones — la IA lo identificó a partir del contenido, no del nombre engañoso."

  - type: note
    text: "Doce archivos pasaron del caos a una estructura limpia y navegable en menos de un minuto. La IA leyó cada archivo para entender su contenido real — detectó que 'IMG_4521.txt' era notas de reuniones a pesar del nombre estilo imagen. Las imágenes también recibieron nombres descriptivos: 'foto-equipo-offsite-t4.png' te dice exactamente qué estás mirando."

  - type: note
    text: "Esto no es solo ordenar: es construir una biblioteca de contexto lista para IA. 'referencia/directrices-tono-marca.md' es instantáneamente utilizable como contexto de anclaje para tareas futuras de escritura. 'doc.txt' no lo era. Los cinco minutos que la IA tarda en organizar ahorran horas de búsqueda después."
```

## Compartir y sincronización en el mundo corporativo

En la mayoría de las organizaciones, los archivos no viven solo en tu portátil. Viven en sistemas compartidos.

**Almacenamiento en la nube:** OneDrive, Google Drive, Dropbox, SharePoint. Aquí es donde viven los archivos del equipo. Cuando organizas archivos para trabajo con IA, se aplican los mismos principios — pero con una consideración adicional: tus colegas también necesitan encontrar las cosas.

**Consejos prácticos para entornos compartidos:**
- Acordar convenciones de nombres con tu equipo
- Crear una estructura de carpetas compartida que todos sigan
- Mantener archivos de trabajo con IA separados de los entregables compartidos hasta que estén listos
- Usar control de versiones o nombres de archivo (v1, v2, final) para rastrear iteraciones

**Un flujo de trabajo común:**
1. Extraer archivos fuente desde el drive compartido de tu equipo
2. Trabajar con la IA localmente o en tu herramienta de IA
3. Revisar y pulir la salida
4. Compartir el entregable terminado de vuelta en el drive del equipo

```callout
type: info
title: "La realidad corporativa"
content: "La mayoría de las organizaciones ya tienen compartición de archivos implementada. La clave es usarla deliberadamente: extraer archivos para contexto de IA, producir salidas, luego compartir resultados a través de tus canales normales. La IA no reemplaza tus herramientas de colaboración: produce mejor contenido para poner en ellas."
```

```quiz
id: file-organisation-naming
type: multiple-choice
question: "Tienes archivos llamados 'informe.docx', 'datos.csv' y 'notas.txt' para un proyecto. ¿Por qué esto es un problema para el trabajo con IA?"
options:
  - "Las extensiones de archivo son incorrectas — la IA prefiere archivos .md"
  - "Los nombres no proporcionan contexto — la IA no puede decir qué contienen estos archivos sin abrir cada uno, y tampoco puedes tú dentro de tres meses"
  - "Tener tres formatos de archivo diferentes hace más difícil para la IA procesarlos juntos"
answer: 1
explanation: "Los nombres descriptivos de archivos son contexto. 'ventas-t4-por-region.csv' le dice tanto a ti como a la IA qué contiene el archivo incluso antes de abrirlo. Los nombres genéricos como 'datos.csv' desperdician la oportunidad de proporcionar contexto a través del nombrado y hacen tu sistema de archivos más difícil de navegar."
```
