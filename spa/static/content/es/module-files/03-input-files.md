---
title: "Archivos de entrada: proporcionar tus datos a la IA"
duration: "10m"
tags: [entrada, datos, contexto]
---

# Archivos de entrada

Los archivos de entrada son los datos que le das a la IA. En lugar de describir tu información, la proporcionas directamente.

## La preparación de archivos importa

Antes de dar archivos a la IA, el estado de esos archivos importa enormemente. Los archivos con nombres incorrectos dispersos por tu escritorio desperdician contexto y confunden al modelo. Los archivos limpios en una estructura lógica producen mejores resultados.

**Buena preparación de archivos:**
- **Nombres descriptivos:** `ventas-t4-2025-por-region.csv` en lugar de `datos (3).xlsx`
- **Carpetas lógicas:** Agrupar archivos relacionados
- **Formatos consistentes:** Escoger CSV para datos, Markdown para documentos, y mantenerlos
- **Eliminar desorden:** Borrar o archivar archivos que no son relevantes para la tarea actual

Piénsalo así: cada archivo que le das a la IA consume espacio en la ventana de contexto. Los archivos irrelevantes diluyen la atención de los archivos que importan. Un conjunto curado de 5 archivos relevantes supera a un vertido de 50 desordenados.

```callout
type: tip
title: "El principio de base"
content: "Un sistema de archivos organizado no es solo limpio: es una biblioteca de contexto para cada tarea futura con IA. Cada archivo que limpias y nombras correctamente es una inversión en mejor salida de IA."
```

## Lo que la IA puede procesar

Los sistemas modernos de IA pueden manejar una amplia variedad de tipos de archivo:

![Tipos de archivo de entrada fluyendo hacia la IA: documentos, datos, visuales, código y formatos mixtos](/content/module-files/images/input-file-types.svg)

**Documentos**
- Archivos PDF (se recomiendan hasta 100 páginas por archivo)
- Documentos de Word
- Archivos de texto plano
- Markdown

**Datos**
- Hojas de cálculo de Excel
- Archivos CSV
- Datos JSON

**Visuales**
- Imágenes y capturas de pantalla
- Diagramas
- Gráficos

**Código**
- Archivos de código fuente
- Archivos de configuración
- Scripts

**Mixto**
- Carpetas con múltiples tipos de archivo
- Hilos de correo electrónico
- Registros de chat

```callout
type: info
title: "Capacidades actuales"
content: "A principios de 2026: ChatGPT soporta hasta 10 archivos por conversación a 512MB por archivo. La ejecución de código de Claude genera hojas de cálculo descargables, CSVs, informes y visualizaciones interactivas."
```

## La idea clave

```callout
type: tip
title: "Proporcionar, no describir"
content: "No digas 'tengo una hoja de cálculo con cifras de ventas del T4'. Proporciona realmente la hoja de cálculo. Deja que la IA vea los datos reales."
```

Esto elimina el juego de 'teléfono sin teléfono'. La IA trabaja con la información real, no con tu descripción de ella.

**Describir datos:**
"La hoja de cálculo tiene columnas para fecha, producto, región e ingresos. Hay unas 500 filas. Las regiones son Norte, Sur, Este y Oeste. Los ingresos van desde..."

**Proporcionar datos:**
*[Subir el archivo]*

¿Cuál es más fiable? ¿Cuál requiere menos trabajo por tu parte?

## Múltiples archivos

Puedes proporcionar múltiples archivos a la vez:
- Una carpeta de documentos relacionados
- Varias hojas de cálculo que cubren diferentes aspectos
- Materiales de referencia junto con los datos a analizar

La IA los procesa juntos, entendiendo cómo se relacionan.

## Buenas prácticas para la preparación de archivos

**Usar formatos legibles por máquina.**
Se prefiere Markdown sobre un formato complejo. El texto plano supera a documentos muy formateados.

**Dividir documentos grandes en segmentos.**
Mantener archivos por debajo de las 100 páginas. Dividir documentos grandes en fragmentos lógicos.

**Eliminar formato excesivo.**
Fuentes, colores e imágenes incrustadas aumentan el uso de tokens sin añadir información.

**Comprender la economía de tokens.**
Para Claude Projects, Markdown/Word/PDF usan espacio similar. HTML usa el doble debido a la interpretación de etiquetas.

**Crear proyectos enfocados.**
Organizar por departamento o función en lugar de una única base de conocimiento masiva. Esto mantiene el contexto relevante.

**Mantener nombres de archivos claros.**
Los nombres descriptivos ayudan a la IA (y a ti) a entender qué contiene cada archivo. Establecer ciclos de revisión regulares para el contenido subido.

## Limitaciones de archivos que debes conocer

No todo funciona perfectamente:
- Los archivos muy grandes pueden necesitar dividirse
- El formato complejo a veces se pierde
- Los PDF escaneados necesitan buena calidad para ser legibles
- Algunos formatos propietarios pueden no ser compatibles

Si tienes dudas, pruébalo. La IA te dirá si no puede procesar algo.

## Múltiples archivos trabajando juntos

Un archivo es útil. Múltiples archivos juntos son potentes. Observa cómo la IA sintetiza información a través de tres archivos de origen diferentes para producir algo que ninguno de ellos podría hacer solo.

```agent
id: multiple-inputs-demo
title: "Síntesis entre archivos"
model_label: "Claude"

system: |
  Eres un analista de inteligencia de negocios. Sintetiza información
  a través de múltiples fuentes. Sé específico y orientado a datos.

scratchpad:
  "datos/ventas-datos.csv": |
    Producto,U_T3,U_T4,Ingresos_T3,Ingresos_T4
    Plan Pro,1240,1580,124K£,158K£
    Plan Equipo,890,720,44.5K£,36K£
    Empresarial,45,62,225K£,310K£
    Nivel Gratuito,8500,9200,£0,£0
    Total,-,-,393.5K£,504K£
  "datos/tickets-soporte.csv": |
    Producto,Tickets_Totales,Resolucion_Prom_Hrs,Problema_Principal
    Plan Pro,342,4.2,"Confusión de facturación anual vs mensual"
    Plan Equipo,521,8.1,"Configuración de permisos demasiado compleja"
    Empresarial,28,1.8,"Soporte dedicado — todos resueltos rápidamente"
    Nivel Gratuito,1205,24+,"Sin prioridad — solo autoservicio"
  "datos/analisis-churn.txt": |
    Análisis de abandono T4
    Plan Pro: abandono mensual del 3,1% (bajó del 4,2% en T3)
    Plan Equipo: abandono mensual del 8,7% (subió del 5,1% en T3)
    Empresarial: 0% de abandono (sin pérdidas en T4)
    Conversión Gratuito→Pago: 2,3% (subió del 1,8%)

    Destacados de encuestas de salida (Plan Equipo):
    - "Demasiado complicado para equipos pequeños"
    - "Solo necesitábamos funciones básicas, pagando por cosas que no usamos"
    - "Cambiado a la competencia — precios más sencillos"

tools:
  - list_files
  - scratchpad_read
  - scratchpad_write
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
    text: "Tres archivos de datos en una carpeta: ventas, tickets de soporte y análisis de abandono. Cada uno cuenta parte de la historia. Observa cómo la IA navega por la carpeta, conecta los puntos y escribe un informe de insights."

  - type: user
    content: "Tengo tres archivos de datos en la carpeta de datos — ventas, tickets de soporte y análisis de abandono. Léelos todos, luego escribe un informe de insights conectando los patrones. ¿Qué historia cuentan juntos?"

  - type: tool_call
    tool: list_files
    args:
      path: "datos"

  - type: tool_result
    tool: list_files
    content: |
      3 archivos en datos/:
      - datos/analisis-churn.txt
      - datos/ventas-datos.csv
      - datos/tickets-soporte.csv

  - type: tool_call
    tool: scratchpad_read
    args:
      filename: "datos/ventas-datos.csv"

  - type: tool_result
    tool: scratchpad_read
    content: |
      Producto,U_T3,U_T4,Ingresos_T3,Ingresos_T4
      Plan Pro,1240,1580,124K£,158K£
      Plan Equipo,890,720,44.5K£,36K£
      Empresarial,45,62,225K£,310K£
      Nivel Gratuito,8500,9200,£0,£0
      Total,-,-,393.5K£,504K£

  - type: tool_call
    tool: scratchpad_read
    args:
      filename: "datos/tickets-soporte.csv"

  - type: tool_result
    tool: scratchpad_read
    content: |
      Producto,Tickets_Totales,Resolucion_Prom_Hrs,Problema_Principal
      Plan Pro,342,4.2,"Confusión de facturación anual vs mensual"
      Plan Equipo,521,8.1,"Configuración de permisos demasiado compleja"
      Empresarial,28,1.8,"Soporte dedicado — todos resueltos rápidamente"
      Nivel Gratuito,1205,24+,"Sin prioridad — solo autoservicio"

  - type: tool_call
    tool: scratchpad_read
    args:
      filename: "datos/analisis-churn.txt"

  - type: tool_result
    tool: scratchpad_read
    content: |
      Análisis de abandono T4
      Plan Pro: abandono mensual del 3,1% (bajó del 4,2% en T3)
      Plan Equipo: abandono mensual del 8,7% (subió del 5,1% en T3)
      Empresarial: 0% de abandono (sin pérdidas en T4)
      Conversión Gratuito→Pago: 2,3% (subió del 1,8%)

      Destacados de encuestas de salida (Plan Equipo):
      - "Demasiado complicado para equipos pequeños"
      - "Solo necesitábamos funciones básicas, pagando por cosas que no usamos"
      - "Cambiado a la competencia — precios más sencillos"

  - type: tool_call
    tool: create_folder
    args:
      path: "informes"

  - type: tool_result
    tool: create_folder
    content: "Carpeta creada: informes"

  - type: tool_call
    tool: scratchpad_write
    args:
      filename: "informes/insights-entre-archivos.md"
      content: |
        # Inteligencia de Producto T4: El Problema del Plan Equipo

        ## Los titulares
        Los ingresos totales crecieron un 28% (393.5K£ → 504K£). Pero ese crecimiento está ocultando una crisis en una categoría.

        ## El patrón en las tres fuentes

        **Empresarial y Pro están prósperos.** Ingresos en aumento, abandono reducido, soporte manejable. Estas categorías están funcionando.

        **El Plan Equipo está en problemas — y las tres bases de datos lo confirman:**
        - **Ventas:** Las unidades cayeron un 19% (890 → 720) mientras cada otra categoría creció
        - **Soporte:** Mayor volumen de tickets (521) con el tiempo de resolución más largo (8,1 horas). Problema principal: "configuración de permisos demasiado compleja"
        - **Abandono:** Casi se duplicó de 5,1% a 8,7%. Las encuestas de salida citan complejidad e inflación de funciones

        La historia es consistente: los clientes del Plan Equipo consideran el producto demasiado complicado para sus necesidades. Están marchándose hacia alternativas más sencillas.

        ## El riesgo oculto
        El crecimiento de ingresos (impulsado por Empresarial) está ocultando este problema. Si el abandono del Plan Equipo continúa al 8,7%, la categoría perderá ~40% de clientes anualmente, haciéndola insostenible en 3 trimestres.

        ## Acciones recomendadas
        1. **Simplificar Plan Equipo** — reducir funciones, simplificar permisos, reducir precios para que coincidan con el valor
        2. **Crear un puente** — considerar una categoría "Equipo Pequeño" entre Gratuito y el actual Plan Equipo
        3. **Corregir el cuello de botella de soporte** — resolución de 8,1 horas para una categoría de pago es inaceptable

  - type: tool_result
    tool: scratchpad_write
    content: "Escrito en informes/insights-entre-archivos.md (23 líneas)"

  - type: assistant
    content: "He creado informes/insights-entre-archivos.md. El hallazgo clave: tu crecimiento global está ocultando una crisis del Plan Equipo. Las tres fuentes de datos — ventas, soporte y abandono — cuentan la misma historia de forma independiente. Ingresos en aumento en general, pero unidades del Plan Equipo un 19% menores, abandono casi duplicado, y los tickets de soporte apuntan a la complejidad como causa."

  - type: note
    text: "Ningún archivo individual contenía este insight. Los datos de ventas mostraron una disminución de unidades. Los datos de soporte mostraron quejas de complejidad. Los datos de abandono mostraron clientes marchándose. Solo leyendo los tres juntos la IA pudo conectarlos en una historia coherente: el Plan Equipo es demasiado complejo y los clientes se marchan por eso."

  - type: note
    text: "Aquí es donde los flujos de trabajo por archivos brillan genuinamente. Tres fuentes de datos, sintetizadas en un único documento de insights accionables, con una sola petición. Intenta hacer eso describiendo los datos en chat."
```

```quiz
id: input-files-insight
type: multiple-choice
question: "Tienes una carpeta con 50 archivos de proyecto. Necesitas que la IA encuentre patrones en 5 archivos de datos específicos. ¿Qué enfoque produce los mejores resultados?"
options:
  - "Subir los 50 archivos para que la IA tenga contexto completo del proyecto"
  - "Subir solo los 5 archivos de datos relevantes para mantener el contexto enfocado y evitar diluir la atención"
  - "Describir los 5 archivos en texto y pedir a la IA que trabaje a partir de tus descripciones"
answer: 1
explanation: "Más archivos significa más contexto consumido, y la atención se diluye entre contenido irrelevante. Subir solo los archivos relevantes mantiene el contexto enfocado. El efecto 'perdido en el medio' significa que enterrar datos clave entre 45 archivos irrelevantes puede reducir la calidad."
```
