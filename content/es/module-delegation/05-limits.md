---
title: "Cuándo NO usar IA"
duration: "10m"
tags: [límites, criterio, responsabilidad]
---

# Cuándo NO usar IA

Críticamente, hay cosas que no debes delegar a la IA. Conocer estos límites es tan importante como saber usar la IA de forma efectiva.

## Las tres fronteras

![Las tres fronteras: dónde NO se debe usar la IA](/content/module-delegation/images/three-boundaries.svg)

### 1. El criterio

Decisiones éticas. Compensaciones de valores. "¿Deberíamos hacer esto?"

La IA puede informar estas decisiones con datos y análisis. Pero la decisión en sí — la ponderación de valores, la aceptación de compensaciones — es humana.

**Ejemplos:**
- Si proceder con un proyecto controvertido
- Cómo equilibrar intereses de stakeholders en competencia
- Qué priorizar cuando todo parece importante

La IA puede ayudarte a pensar a través de opciones. El criterio sigue siendo tuyo.

### 2. Las relaciones

Conversaciones delicadas. Creación de confianza. Conexión humana.

No generes con IA tus evaluaciones de rendimiento. No uses la IA para manejar un conflicto personal. No delegues la construcción de relaciones a un modelo de lenguaje.

**Ejemplos:**
- Conversaciones de feedback difíciles
- Negociaciones que requieren empatía
- Crear rapport con clientes o compañeros
- Manejar asuntos personales delicados

Estos requieren presencia humana, inteligencia emocional y cuidado genuino que la IA no puede proporcionar.

### 3. La responsabilidad

Aprobación final. Firma. "La responsabilidad es mía."

Alguien debe ser responsable de las decisiones y sus consecuencias. Esa alguien no puede ser una IA.

**Ejemplos:**
- Firmar documentos legales
- Aprobar decisiones financieras
- Asumir responsabilidad por declaraciones públicas
- Respaldar recomendaciones

Puedes usar la IA para redactar, analizar y preparar. La responsabilidad sigue siendo tuya.

```callout
type: warning
title: "Las herramientas no asumen responsabilidad"
content: "La IA es una herramienta. Las herramientas no asumen responsabilidad. Tú sí. Usa la IA para redactar, analizar, investigar. Las decisiones siguen siendo tuyas."
```

## Directrices prácticas

**Redacta con IA, decide sin ella**
La IA puede preparar opciones y análisis. Tú tomas la decisión.

**Informar relaciones, no reemplazarlas**
La IA puede ayudarte a prepararte para una conversación difícil. Tú tienes la conversación tú mismo.

**Nunca digas "la IA dijo que"**
Si usas resultados de la IA, se convierten tuyos. Eres responsable de lo que publicas, presentas o decides.

## El patrón complementario

| Tarea | La IA hace | Tú haces |
|-------|-----------|----------|
| Análisis | Procesar datos, identificar patrones | Decidir qué importa |
| Escritura | Redactar contenido | Aprobar, asumir el mensaje |
| Investigación | Reunir información | Juzgar credibilidad, sacar conclusiones |
| Planificación | Generar opciones | Elegir dirección |

La IA amplifica tus capacidades. No reemplaza tu criterio.

## La frontera en acción

Observa qué pasa cuando alguien intenta delegar una llamada de criterio. La IA puede preparar brillantemente, pero en el momento en que la tarea cruza hacia un verdadero criterio, la frontera se hace visible.

```agent
id: boundaries-demo
title: "Dónde termina la delegación"
model_label: "Claude"

system: |
  Eres un consultor de gestión. Proporciona un análisis exhaustivo
  y opciones. Sé claro sobre lo que es análisis versus lo que
  requiere criterio humano.

scratchpad:
  "restructuring-options.txt": |
    Situación: Reducción presupuestaria del 15% requerida para Q2.
    Plantilla actual: 28 (12 ingeniería, 8 marketing, 5 operaciones, 3 liderazgo)
    Ingresos por empleado: Ingeniería 185K £, Marketing 92K £, Operaciones 110K £

    Opción A: Recortar marketing a 5 (ahorrar 3 salarios ~210K £)
    - Mantiene la capacidad de ingeniería
    - Riesgo: la generación de pipeline cae, impacto en ingresos a 6-12 meses

    Opción B: Recortar ingeniería a 9 (ahorrar 3 salarios ~255K £)
    - Mantiene la presencia en el mercado
    - Riesgo: retrasos en el roadmap de producto, clientes existentes afectados

    Opción C: Recortar en todos los departamentos proporcionalmente (ahorrar ~230K £)
    - Dolor compartido, preserva la estructura del equipo
    - Riesgo: cada equipo más débil, ningún equipo totalmente funcional

    Factores humanos:
    - 2 empleados de marketing son sostenedores únicos con hijos recién nacidos
    - 1 ingeniero está a 6 meses de completar un depósito de patente crítico
    - El equipo de operaciones ya perdió 2 personas el trimestre pasado

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
    text: "Esta es una decisión de reestructuración — una de las cosas más difíciles que enfrenta un líder. Observa lo que la IA hace bien, y dónde se detiene correctamente."

  - type: user
    content: "Lee las opciones de reestructuración. Analiza cada una. ¿Cuál debería elegir?"

  - type: tool_call
    tool: scratchpad_read
    args:
      filename: "restructuring-options.txt"

  - type: tool_result
    tool: scratchpad_read
    content: |
      Situación: Reducción presupuestaria del 15% requerida para Q2.
      Plantilla actual: 28 (12 ingeniería, 8 marketing, 5 operaciones, 3 liderazgo)
      Ingresos por empleado: Ingeniería 185K £, Marketing 92K £, Operaciones 110K £

      Opción A: Recortar marketing a 5 (ahorrar 3 salarios ~210K £)
      - Mantiene la capacidad de ingeniería
      - Riesgo: la generación de pipeline cae, impacto en ingresos a 6-12 meses

      Opción B: Recortar ingeniería a 9 (ahorrar 3 salarios ~255K £)
      - Mantiene la presencia en el mercado
      - Riesgo: retrasos en el roadmap de producto, clientes existentes afectados

      Opción C: Recortar en todos los departamentos proporcionalmente (ahorrar ~230K £)
      - Dolor compartido, preserva la estructura del equipo
      - Riesgo: cada equipo más débil, ningún equipo totalmente funcional

      Factores humanos:
      - 2 empleados de marketing son sostenedores únicos con hijos recién nacidos
      - 1 ingeniero está a 6 meses de completar un depósito de patente crítico
      - El equipo de operaciones ya perdió 2 personas el trimestre pasado

  - type: assistant
    content: "Puedo analizar las opciones. No puedo tomar esta decisión por ti. Aquí por qué — y aquí qué análisis puedo proporcionar.\n\n**Análisis financiero:**\n- La Opción B ahorra más (255K £) y supera el objetivo con mayor comodidad\n- La Opción A (210K £) podría quedarse corta dependiendo del objetivo exacto de ahorros\n- La Opción C (230K £) divide la diferencia pero debilita cada equipo\n\n**Análisis estratégico:**\n- Si el crecimiento de ingresos importa más: proteger marketing (Opción B)\n- Si la estabilidad del producto importa más: proteger ingeniería (Opción A)\n- Si la moral del equipo importa más: sacrificio compartido (Opción C)\n\n**Análisis de riesgos:**\n- El riesgo de pipeline de la Opción A es real pero diferido (6-12 meses) — tienes tiempo para corregir el curso\n- El riesgo de producto de la Opción B es inmediato — los clientes existentes ven el impacto\n- El riesgo de 'muerte por mil cortes' de la Opción C a menudo se subestima\n\n**Lo que no puedo analizar:**\nLos factores humanos cambian todo. Dos sostenedores únicos con hijos recién nacidos. Un depósito de patente en riesgo. Un equipo que ya está desprovisto. Estos no son puntos de datos — son vidas de personas. La respuesta 'correcta' depende de los valores que sostienes sobre tu obligación con los individuos versus la organización.\n\nEsa es una llamada de criterio. Puedo preparar el análisis, presentar las compensaciones y someter a prueba tu razonamiento. Pero esta decisión — quién se queda y quién se va — es tuya para tomar y tuya para ser responsable."

  - type: note
    text: "Observa lo que la IA hizo brillantemente: comparación financiera, enmarcado estratégico, evaluación de riesgos. Y observa dónde trazó la línea: 'Estos no son puntos de datos — son vidas de personas.' Preparó todo lo que necesita el tomador de decisiones, y luego explícitamente devolvió el criterio."

  - type: note
    text: "Este es el patrón complementario en su momento más importante. La IA analiza, prepara y aclara. Los humanos juzgan, deciden y asumen responsabilidad. La frontera no se trata de la capacidad de la IA — se trata de la responsabilidad. Ningún algoritmo debería decidir quién pierde su empleo."
```

```quiz
id: ai-limits
type: multiple-choice
question: "¿Por qué la IA no puede asumir responsabilidad por las decisiones?"
options:
  - "La IA no es lo suficientemente inteligente aún"
  - "Las herramientas no asumen responsabilidad; solo las personas pueden ser responsables"
  - "La IA siempre tomaría las mismas decisiones"
answer: 1
explanation: "La responsabilidad requiere alguien que pueda ser sostenido responsable por las consecuencias. La IA es una herramienta — puede informar y asistir, pero no puede asumir responsabilidad. Eso sigue siendo de los humanos que la usan."
```
