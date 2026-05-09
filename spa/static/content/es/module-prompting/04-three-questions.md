---
title: "Las Tres Preguntas"
duration: "15m"
tags: [preguntas, especificación, claridad]
---

# Las Tres Preguntas

![El Framework de Las Tres Preguntas](/content/module-prompting/images/three-questions.svg)

Aquí hay un framework aún más simple. Antes de cualquier tarea con IA, hazte estas tres preguntas.

## 1. ¿Cómo se ve "terminado"?

Si no puedes describir el estado final, la IA no puede producirlo.

**Vago:**
"Ayúdame con mi presentación"

**Mejor:**
"Crea diapositivas para mi presentación"

**Claro:**
"Crea una presentación de 10 diapositivas sobre los resultados del Q4 para el equipo de ventas, con una métrica clave resaltada por diapositiva, terminando con las prioridades del Q1"

```callout
type: tip
title: "La prueba de verificación"
content: "¿Podrías verificar que el trabajo está hecho correctamente? Si no puedes describir cómo se ve lo 'correcto', tu especificación no es lo suficientemente clara."
```

## 2. ¿Qué contexto necesita?

¿Qué información se requiere para el éxito?

**Datos y hechos**
La información real con la que trabajar — hojas de cálculo, documentos, código.

**Antecedentes y situación**
Contexto de la empresa, historial del proyecto, decisiones previas. ¿Para qué existe esta tarea?

**Audiencia**
¿Quién usará esto? ¿Qué ya saben? ¿Ejecutivos? ¿Desarrolladores? ¿Clientes?

**Ejemplos**
Trabajo similar que te gustó. Plantillas. Referencias de estilo. Mostrar, no solo decir.

**Requisitos**
Must-haves explícitos. Necesidades de cumplimiento. Guías de marca.

Cuanto más contexto relevante proporciones, mejor será la salida.

## 3. ¿Cuáles son los límites?

¿Qué restricciones o límites aplican?

**Longitud**
Conteo de palabras, conteo de páginas, número de elementos. "Aproximadamente 500 palabras" es mejor que nada.

**Tono**
¿Profesional? ¿Casual? ¿Técnico? ¿Explicado como si fuera para un niño?

**Formato**
¿Viñetas? ¿Párrafos? ¿Tabla? ¿Código? Esto afecta dramáticamente la utilidad.

**Inclusiones**
¿Qué debe cubrirse? "Debe mencionar el nuevo modelo de precios."

**Exclusiones**
¿Qué evitar? "No discutas la fusión." "Salta los detalles de implementación."

**Restricciones**
Límites técnicos, requisitos de marca, necesidades de cumplimiento.

Los límites evitan que la IA se desvíe en direcciones que no quieres.

## Poniéndolo todo junto

**Tarea:** Analizar comentarios de clientes

**Pregunta 1 — Terminado:**
Un resumen de las 5 principales categorías de quejas con citas representativas

**Pregunta 2 — Contexto:**
- Archivo CSV con 500 entradas de comentarios
- Es para el equipo de producto
- Estamos enfocados en la experiencia de nuestra app móvil

**Pregunta 3 — Límites:**
- Máximo 1 página
- Incluir conteos de frecuencia
- Tono profesional pero conciso

**Prompt resultante:**
```
Analiza este CSV de comentarios de clientes (adjunto).
Identifica las 5 principales categorías de quejas con conteos de frecuencia.
Para cada categoría, proporciona 2-3 citas representativas.
Enfócate en problemas de la experiencia de la app móvil.
Formatea como un breve reporte para el equipo de producto, máximo 1 página.
```

## Las Tres Preguntas aplicadas

Observa a alguien aplicar las tres preguntas a una tarea real. Observa cómo cada pregunta añade precisión a lo que empezó como una idea vaga.

```agent
id: three-questions-demo
title: "De idea vaga a petición clara"
model_label: "Claude"

system: |
  Eres un consultor de comunicaciones ayudando a redactar
  comunicaciones internas. Escribe de forma clara y concisa
  para una audiencia profesional.

scratchpad:
  "policy-changes.txt": |
    Cambios efectivos desde 1 de marzo de 2026:
    - Trabajo remoto: Ahora 3 días oficina, 2 días remoto (era 2/3)
    - Viernes sin reuniones: Sin reuniones internas antes de las 2pm
    - Política de gastos: Pre-aprobación requerida para compras de más de £200 (era £500)
    - Presupuesto de formación: Presupuesto individual aumentado a £2,000/año (era £1,200)
    - Licencia parental: Extendida a 26 semanas con pago completo (era 16 semanas)
  "previous-comms.txt": |
    Asunto: Actualización de políticas de diciembre
    Equipo,
    Algunos cambios a tener en cuenta para el nuevo año.
    [Tres párrafos de texto denso, sin formato, los puntos clave enterrados]
    Resultado: El 40% del personal no lo leyó. RRHH atendió 200+ preguntas.

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
    text: "El usuario tiene cambios de política que comunicar al personal. En lugar de simplemente decir 'escribe un email sobre estos cambios,' observa cómo aplica las tres preguntas para construir una petición precisa."

  - type: user
    content: "Necesito comunicar algunos cambios de política al equipo. Lee ambos archivos primero — los cambios en sí y las comunicaciones previas que no funcionaron."

  - type: tool_call
    tool: scratchpad_read
    args:
      filename: "policy-changes.txt"

  - type: tool_result
    tool: scratchpad_read
    content: |
      Cambios efectivos desde 1 de marzo de 2026:
      - Trabajo remoto: Ahora 3 días oficina, 2 días remoto (era 2/3)
      - Viernes sin reuniones: Sin reuniones internas antes de las 2pm
      - Política de gastos: Pre-aprobación requerida para compras de más de £200 (era £500)
      - Presupuesto de formación: Presupuesto individual aumentado a £2,000/año (era £1,200)
      - Licencia parental: Extendida a 26 semanas con pago completo (era 16 semanas)

  - type: tool_call
    tool: scratchpad_read
    args:
      filename: "previous-comms.txt"

  - type: tool_result
    tool: scratchpad_read
    content: |
      Asunto: Actualización de políticas de diciembre
      Equipo,
      Algunos cambios a tener en cuenta para el nuevo año.
      [Tres párrafos de texto denso, sin formato, los puntos clave enterrados]
      Resultado: El 40% del personal no lo leyó. RRHH atendió 200+ preguntas.

  - type: assistant
    content: "He leído ambos archivos. La comunicación previa fue densa y sin formato, lo que llevó a un bajo engagement. Puedo ver cinco cambios de política, algunos positivos (presupuesto de formación, licencia parental) y otros que pueden ser menos bienvenidos (días de oficina, umbral de gastos). ¿Cómo te gustaría que aborde la comunicación?"

  - type: user
    content: "Esto es lo que necesito. He pensado en las tres preguntas:\n\n**Terminado:** Un email para todo el personal que todos realmente lean y entiendan. Sin preguntas de seguimiento a RRHH.\n\n**Contexto:** Empresa de 200 personas. Mix de personal en oficina y remoto. Las últimas comunicaciones fallaron porque fueron demasiado densas. Algunos cambios son populares (formación, licencia parental), otros no lo son (más días de oficina, menor umbral de gastos).\n\n**Límites:** Menos de 300 palabras. Empieza con las buenas noticias. Cada cambio tiene su propia viñeta con el antes y después claramente mostrado. Tono cálido pero directo — no endulces los cambios menos populares. La línea de asunto debe captar atención."

  - type: tool_call
    tool: scratchpad_write
    args:
      filename: "policy-email.md"
      content: |
        Asunto: 5 cambios de políticas a partir de marzo — incluyendo mayor presupuesto de formación y licencia parental extendida

        Hola equipo,

        Cinco cambios de políticas entran en vigor el 1 de marzo. Esto es lo que cambia y por qué.

        **Lo nuevo (y positivo):**

        - **Presupuesto de formación: £1,200 → £2,000/año.** Estamos invirtiendo más en vuestro desarrollo. Usadlo para cursos, conferencias, certificaciones — vuestra elección.
        - **Licencia parental: 16 → 26 semanas con pago completo.** Esto nos pone en línea con lo que creemos es correcto para nuestro equipo.
        - **Mañanas de viernes sin reuniones.** Sin reuniones internas antes de las 2pm los viernes. Proteged vuestro tiempo de trabajo profundo.

        **Lo que también cambia:**

        - **Días de oficina: 2 → 3 por semana.** Hemos visto que la calidad de la colaboración cae por debajo de tres días juntos. El nuevo default es lunes, martes, jueves en oficina. Hablad con vuestro gestor si necesitáis flexibilidad.
        - **Pre-aprobación de gastos: £500 → £200.** Finanzas necesita mejor visibilidad sobre los gastos. La pre-aprobación es rápida — enviad vía el portal, esperad respuesta el mismo día.

        Cada cambio tiene un FAQ detallado en la intranet. Si algo no está claro, checked ahí primero.

        — El Equipo Directivo

  - type: tool_result
    tool: scratchpad_write
    content: "Escrito a policy-email.md (23 líneas)"

  - type: assistant
    content: "Listo — policy-email.md está preparado. Empecé con los cambios más populares (aumento de presupuesto, licencia parental, viernes sin reuniones), luego abordé los menos bienvenidos directamente sin rodeos. Cada viñeta muestra el antes → después claramente. 220 palabras. La línea de asunto nombra el conteo y destaca las buenas noticias para impulsar aperturas."

  - type: note
    text: "Compara esto con una instrucción como 'escribe un email sobre nuestros cambios de política.' Las tres preguntas convirtieron una idea vaga en una especificación precisa: cómo se ve terminado (sin preguntas de RRHH), qué contexto importa (las últimas comunicaciones fallaron, se esperan reacciones mixtas), cuáles son los límites (300 palabras, empezar con las buenas noticias, tono directo)."
```

```quiz
id: three-questions-test
type: multiple-choice
question: "Le pides a la IA que 'ayude con la presentación trimestral.' Produce algo mediocre. ¿De las Tres Preguntas, cuál probablemente omitiste?"
options:
  - "'¿Cómo se ve terminado?' — no especificaste el entregable, la audiencia o el formato"
  - "'¿Qué contexto necesita?' — no proporcionaste los datos del Q4, la plantilla de diapositivas o la guía de marca"
  - "'¿Cuáles son los límites?' — no estableciste restricciones sobre longitud, tono o temas a cubrir"
answer: 0
explanation: "'Ayudar con' no es un resultado definido. Sin saber cómo se ve 'terminado' — ¿10 diapositivas? ¿Resumen ejecutivo? ¿Puntos para hablar? — la IA no puede producir el entregable correcto. Las otras preguntas también importan, pero 'terminado' es el fundamento que hace que el contexto y los límites tengan sentido."
```
