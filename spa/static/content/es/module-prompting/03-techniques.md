---
title: "Las técnicas que funcionan"
duration: "20m"
tags: [técnicas, ejemplos, multishot, xml, pensamiento, encadenamiento]
---

# Las técnicas que realmente funcionan

Ya aprendiste el framework. Ahora están las técnicas específicas que consistentemente producen mejores resultados, ordenadas por efectividad.

## 1. Sé específico y directo

Esta sigue siendo la técnica de mayor impacto. Los modelos modernos siguen las instrucciones de forma muy literal: las instrucciones vagas producen resultados vagos.

**La prueba del colega:** Muestra tu instrucción a un colega. Si él se confunde sobre lo que quieres, la IA también lo hará.

**Mal:**
```
Escribe sobre nuestro producto.
```

**Bien:**
```
Escribe una descripción de 300 palabras de nuestra producto para el seguimiento
de gastos para dueños de pequeños negocios. Enfatiza la capacidad de
escaneo automático de recibos y el diseño mobile-first. Tono profesional pero amigable.
```

Especifica siempre: formato, longitud, tono, alcance y audiencia.

```callout
type: tip
title: "Los modelos modernos son literales"
content: "Claude Opus 4.6 y GPT-5.x toman tus instrucciones de forma literal. Entrada vaga produce salida vaga. Entrada específica produce salida específica."
```

## 2. Usa ejemplos (Multishot Prompting)

Proporcionar 2-5 ejemplos de pares entrada-salida deseados mejora dramáticamente la consistencia y la calidad.

Esto es especialmente poderoso cuando importa el formato o el estilo.

**Sin ejemplos:**
```
Categoriza estos tickets de soporte.
[Lista de tickets]
```

Los resultados serán inconsistentes. La IA tiene que adivinar qué categorías quieres.

**Con ejemplos:**
```
Categoriza estos tickets de soporte. Usa estas categorías:

<examples>
"La app se cierra cuando intento subir fotos" → Bug
"¿Cómo exporto mis datos?" → Petición de función
"No puedo iniciar sesión" → Problema de acceso
</examples>

Ahora categoriza estos:
[Lista de tickets]
```

Los resultados coincidirán con el patrón deseado.

```callout
type: info
title: "El poder de mostrar, no solo decir"
content: "Los ejemplos enseñan al modelo el patrón que quieres. Dos a cinco ejemplos son suficientes. Más de eso muestra rendimientos decrecientes."
```

## 3. Activa el pensamiento para tareas difíciles

**Pensamiento extendido** es una fase interna de razonamiento antes de que el modelo responda. Los modelos Claude modernos lo tienen integrado. OpenAI lo llama "esfuerzo de razonamiento."

Úsalo para:
- Problemas matemáticos o lógicos
- Análisis multi-paso
- Depurar problemas complejos
- Decisiones complejas que requieren evaluar compensaciones

**Perspectiva contraintuitiva:** Las investigaciones muestran que los modelos suelen tener un mejor rendimiento con instrucciones para pensar profundamente sobre una tarea en lugar de guía prescriptiva paso a paso.

**En lugar de:**
```
Primero analiza X, luego considera Y, luego compara con Z, luego concluye.
```

**Intenta:**
```
Piensa profundamente sobre este problema antes de responder. Considera todos los ángulos.
```

Las capacidades de razonamiento del modelo son sofisticadas. Déjalo razonar a su manera.

## 4. Estructura con etiquetas XML

Los modelos entrenados en código son excelentes analizando formatos estructurados como XML.

Las etiquetas separan los componentes de la instrucción, reducen malinterpretaciones y hacen la salida más fácil de analizar.

**Estructura básica:**
```
<instructions>
Analiza los comentarios de los clientes e identifica tendencias.
</instructions>

<context>
Lanzamos una nueva app móvil hace dos meses.
Estos comentarios vienen de early adopters.
</context>

<examples>
"La app es lenta" → Problema de rendimiento
"Me encanta el modo oscuro" → Comentario positivo de UI
</examples>

<output_format>
- Lista las 3 principales tendencias
- Incluye conteos de frecuencia
- Proporciona citas representativas
</output_format>

<constraints>
- Céntrate en comentarios accionables
- Excluye peticiones de función ya planificadas
</constraints>
```

Esta estructura es mucho más clara que un muro de texto con la misma información.

**Nota:** Todos los modelos principales se benefician de formatos estructurados: XML, JSON o YAML. Usa el que te resulte más natural para tu tarea.

```callout
type: tip
title: "La estructura reduce la ambigüedad"
content: "Las etiquetas hacen explícito dónde termina el contexto y empiezan las instrucciones. El modelo no confundirá ejemplos con datos reales, ni restricciones con tareas."
```

## 5. Encadena instrucciones complejas

Para tareas multi-paso, divide el trabajo en instrucciones de subtarea secuencial donde cada salida alimenta la siguiente.

Esto refleja cómo delegarías trabajo a un humano: investigar primero, luego organizar, luego redactar, luego revisar.

**Instrucción monolítica única:**
```
Investiga precios de competidores, analiza nuestro posicionamiento,
redacta una estrategia de precios y escribe un resumen ejecutivo.
```

Esto intenta hacer todo a la vez. Los resultados suelen ser superficiales.

**Encadenamiento de instrucciones:**
```
Paso 1: Investiga precios de competidores para [productos]. Crea una tabla comparativa.
[Revisa la salida, entrégala al siguiente paso]

Paso 2: Con estos datos de competidores, analiza nuestro posicionamiento. ¿Dónde competimos en precio? ¿Dónde en características?
[Revisa la salida, entrégala al siguiente paso]

Paso 3: Redacta una estrategia de precios basada en este análisis. Considera nuestra estructura de costes y margen objetivo.
[Revisa la salida, entrégala al siguiente paso]

Paso 4: Resume esta estrategia como un resumen ejecutivo de 1 página para el equipo directivo.
```

Cada paso tiene una entrada y una salida claras. Verificas en cada etapa. El resultado final es mucho más cualitativo.

```callout
type: warning
title: "Cuándo encadenar vs cuándo combinar"
content: "Encadena cuando cada paso requiere verificación o alimenta el siguiente. Combina cuando los pasos son independientes. No encadenes innecesariamente — toma más tiempo."
```

## Antes y después: poniéndolo todo junto

**Antes (vago, no estructurado):**
```
Analiza nuestros datos de ventas del Q4 y dime qué es importante.
```

**Después (específico, estructurado, con ejemplos):**
```
<instructions>
Analiza los datos de ventas del Q4 de 2025. Identifica las 3 principales tendencias que afectan los ingresos.
</instructions>

<context>
Vendemos SaaS B2B con planes anuales y mensuales.
El Q4 suele ser nuestro trimestre más fuerte por los gastos de presupuesto de fin de año.
Este año lanzamos un nuevo plan enterprise en octubre.
</context>

<output_format>
Para cada tendencia:
- Nombre de la tendencia
- Datos de respaldo (números, porcentajes)
- Implicación empresarial
- Acción recomendada
</output_format>

<examples>
"Adopción del plan enterprise" → 15% de los ingresos, mayor que lo pronosticado → Fuerte ajuste al mercado → Invertir en contratación de ventas enterprise
</examples>

<constraints>
- Céntrate en insights accionables, no solo descripción
- Máximo 1 página
- Tono profesional para el equipo ejecutivo
</constraints>

[Adjuntar: CSV de datos de ventas del Q4]
```

La segunda instrucción le da a la IA todo lo que necesita para producir exactamente lo que quieres.

## Multishot en práctica

Las descripciones del multishot prompting pueden parecer abstractas. Observa lo que realmente hace — la misma tarea de categorización, primero sin ejemplos, luego con ellos.

```agent
id: multishot-demo
title: "El poder de los ejemplos"
model_label: "Claude"

system: |
  Eres un analista de soporte al cliente. Categoriza los tickets de forma
  precisa en base a las instrucciones y ejemplos proporcionados.

scratchpad:
  "tickets.txt": |
    Ticket #201: "Cada vez que intento subir un archivo de más de 10MB, la página se pone blanca y tengo que refrescar."
    Ticket #202: "¿Hay alguna forma de configurar tareas recurrentes? No lo encuentro en ningún lado."
    Ticket #203: "Mi compañero me compartió un proyecto pero no puedo verlo en mi panel."
    Ticket #204: "La app móvil me desconectó y ahora mi 2FA no funciona."
    Ticket #205: "Me encantaría que pudieran añadir una opción de modo oscuro."
    Ticket #206: "Los reportes están mostrando datos del mes anterior incluso después de actualizar la hoja de cálculo fuente."

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
    text: "Primer intento: sin ejemplos. Solo una instrucción vaga para categorizar. Observa qué tan inconsistentes son los resultados."

  - type: user
    content: "Lee el archivo de tickets y categoriza cada uno."

  - type: tool_call
    tool: scratchpad_read
    args:
      filename: "tickets.txt"

  - type: tool_result
    tool: scratchpad_read
    content: |
      Ticket #201: "Cada vez que intento subir un archivo de más de 10MB, la página se pone blanca y tengo que refrescar."
      Ticket #202: "¿Hay alguna forma de configurar tareas recurrentes? No lo encuentro en ningún lado."
      Ticket #203: "Mi compañero me compartió un proyecto pero no puedo verlo en mi panel."
      Ticket #204: "La app móvil me desconectó y ahora mi 2FA no funciona."
      Ticket #205: "Me encantaría que pudieran añadir una opción de modo oscuro."
      Ticket #206: "Los reportes están mostrando datos del mes anterior incluso después de actualizar la hoja de cálculo fuente."

  - type: assistant
    content: "Aquí están los tickets categorizados:\n\n- **#201** — Problema técnico / Problema de subida\n- **#202** — Pregunta / Consulta de función\n- **#203** — Permisos / Problema de compartir\n- **#204** — Autenticación / Problema de inicio de sesión\n- **#205** — Sugerencia de función\n- **#206** — Bug de datos/reportes"

  - type: note
    text: "Fíjate en esas categorías. Todas tienen formatos distintos. Algunas con barras, otras sin. Sin taxonomía consistente. Si ejecutabas esto con 500 tickets, obtendrías 50 nombres de categoría distintos. Ahora intentémoslo con ejemplos."

  - type: clear
    reset_scratchpad: false
    note: "Empezando de nuevo esta vez con ejemplos."

  - type: user
    content: "Lee el archivo de tickets y categoriza cada uno.\n\nUsa estas categorías exactas. Aquí hay ejemplos:\n\n'La app se cierra al iniciar sesión' → BUG (P1)\n'¿Puedo cambiar mi dirección de email?' → PREGUNTA (P3)\n'Por favor añaden integración con calendario' → PETICIÓN DE FUNCIÓN (P3)\n'No puedo acceder a archivos compartidos' → PROBLEMA DE ACCESO (P2)\n\nPara cada ticket da: Ticket # → CATEGORÍA (Prioridad)"

  - type: tool_call
    tool: scratchpad_read
    args:
      filename: "tickets.txt"

  - type: tool_result
    tool: scratchpad_read
    content: |
      Ticket #201: "Cada vez que intento subir un archivo de más de 10MB, la página se pone blanca y tengo que refrescar."
      Ticket #202: "¿Hay alguna forma de configurar tareas recurrentes? No lo encuentro en ningún lado."
      Ticket #203: "Mi compañero me compartió un proyecto pero no puedo verlo en mi panel."
      Ticket #204: "La app móvil me desconectó y ahora mi 2FA no funciona."
      Ticket #205: "Me encantaría que pudieran añadir una opción de modo oscuro."
      Ticket #206: "Los reportes están mostrando datos del mes anterior incluso después de actualizar la hoja de cálculo fuente."

  - type: assistant
    content: "#201 → BUG (P1)\n#202 → PREGUNTA (P3)\n#203 → PROBLEMA DE ACCESO (P2)\n#204 → PROBLEMA DE ACCESO (P1)\n#205 → PETICIÓN DE FUNCIÓN (P3)\n#206 → BUG (P2)"

  - type: note
    text: "Categorías consistentes. Formato consistente. Asignación de prioridad consistente. Cuatro ejemplos fueron suficientes para enseñarle al modelo tu taxonomía exacta y formato de salida. Ejecuta esto con 500 tickets y cada uno seguirá el mismo patrón."

  - type: note
    text: "Este es el poder de mostrar sobre decir. Los ejemplos hicieron lo que párrafos de instrucciones no pudieron: definieron el patrón sin ambigüedad."
```

```quiz
id: techniques-multishot
type: multiple-choice
question: "Necesitas que la IA categorice 500 tickets de soporte en exactamente 6 categorías. Describes las categorías en detalle. Los resultados son inconsistentes. ¿Qué mejoraría más la consistencia?"
options:
  - "Añadir 3-4 tickets de ejemplo con sus categorías correctas para que el modelo aprenda el patrón"
  - "Usar etiquetas XML para separar las definiciones de categoría de los datos de los tickets"
  - "Pedirle al modelo que piense paso a paso antes de categorizar cada ticket"
answer: 0
explanation: "El multishot prompting (proporcionar ejemplos) es la técnica más efectiva para tareas de categorización y consistencia. Los ejemplos enseñan el patrón exacto mediante demostración — el modelo coincide con tus ejemplos en lugar de interpretar tus descripciones. Las descripciones son ambiguas; los ejemplos no lo son."
```

## La jerarquía de técnicas

![Jerarquía de Técnicas](/content/module-prompting/images/techniques-hierarchy.svg)

De más ampliamente efectiva a más especializada:

1. **Sé específico y directo** — Siempre aplicable. Mayor impacto.
2. **Usa ejemplos** — Cuando importa el formato, estilo o categorización.
3. **Activa el pensamiento** — Para tareas de razonamiento genuinamente difíciles.
4. **Estructura con etiquetas** — Cuando las instrucciones son complejas o multi-parte.
5. **Encadena instrucciones** — Para flujos de trabajo multi-etapa que requieren verificación.

Domina las dos primeras y verás resultados inmediatos. Añade las demás conforme las tareas se vuelvan más complejas.
