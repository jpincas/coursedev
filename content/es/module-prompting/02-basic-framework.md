---
title: "El Framework Básico"
duration: "15m"
tags: [framework, estructura, componentes]
---

# Un framework para estructurar peticiones a la IA

Aquí tienes un framework básico para estructurar peticiones a la IA. No necesitas los cinco elementos cada vez, pero pensar en ellos te ayuda.

## Los Cinco Elementos

![Los Cinco Componentes del Framework](/content/module-prompting/images/framework-components.svg)

Los modelos modernos siguen las instrucciones de forma muy literal: las instrucciones vagas producen resultados vagos. **Ser específico y directo sigue siendo la técnica de mayor impacto** incluso con modelos avanzados como Claude Opus 4.6 y GPT-5.x.

**Rol** — ¿Quién debería ser la IA? ¿Qué experiencia debería aportar?
- "Eres un analista financiero senior"
- "Actúa como escritor técnico para documentación de desarrolladores"

**Tarea** — ¿Qué cosa específica debería hacer?
- Sé específico sobre el entregable
- "Analiza este informe trimestral y resume las tendencias clave"

**Contexto** — ¿Qué información de respaldo necesita?
- Contexto de la empresa, decisiones previas, historia relevante
- Datos, documentos, ejemplos

**Formato** — ¿Cómo debería estructurarse la salida?
- Viñetas, párrafos, tablas, código
- Conteo de palabras, número de elementos

**Restricciones** — ¿Qué límites o requisitos aplican?
- Temas a evitar
- Requisitos de tono
- Necesidades de cumplimiento

## Mala Instrucción vs Buena Instrucción

Vamos a comparar dos enfoques para la misma tarea.

**Mala instrucción:**
```
Escríbeme un informe sobre el Q4.
```

Problemas:
- ¿Qué Q4? ¿De qué año?
- ¿De qué empresa?
- ¿Qué tipo de informe?
- ¿Para qué audiencia?
- ¿Qué formato? ¿Qué longitud?

La IA tiene que adivinarlo todo.

**Buena instrucción:**
```
Escribe un resumen ejecutivo del Q4 de 2025 para la reunión de la junta directiva de Acme Corp.

Incluir: tendencias de ingresos, éxitos clave, riesgos.

Formato: 1 página, viñetas, tono profesional.

Contexto: [datos del Q4 adjuntos]
```

Esta instrucción especifica:
- Periodo y empresa concretos
- Tipo de entregable claro (resumen ejecutivo)
- Audiencia definida (junta directiva)
- Formato explícito
- Contenido requerido

```callout
type: tip
title: "La idea clave"
content: "No se trata de encontrar las palabras mágicas. Se trata de información completa. La IA no puede adivinar lo que quieres decir — dímelo explícitamente."
```

## Cuándo usar qué

No toda petición necesita los cinco elementos.

**Preguntas rápidas:** Simplemente pregunta. "¿Cuál es la capital de Francia?"

**Tareas simples:** Tarea + Formato suele ser suficiente. "Lista las 5 principales características de Python 3.12 en forma de viñetas."

**Trabajo complejo:** Usa el framework completo. Cuanto más relevante sea el resultado, más completa debe ser tu especificación.

## El framework en acción

Las palabras en una página solo pueden explicar tanto. Observa la diferencia entre una mala instrucción y una buena — usando la misma IA, la misma tarea, los mismos datos.

```agent
id: framework-demo
title: "Mala instrucción vs Buena instrucción"
model_label: "Claude"

system: |
  Eres un asistente empresarial útil. Responde en base a la
  información e instrucciones proporcionadas.

scratchpad:
  "q4-data.csv": |
    Mes,Ingresos,Nuevos_Clientes,Tasa_Churn,NPS
    Oct,1,82M,145,3,2%,42
    Nov,2,01M,168,2,8%,45
    Dic,2,34M,203,2,1%,51
    Total Q4,6,17M,516,2,7% promedio,46 promedio
    vs Q3,+18%,+24%,-0,8pp,+6

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
    text: "Primero, la mala instrucción. Vaga, sin contexto, sin formato especificado. Observa qué sale."

  - type: user
    content: "Escríbeme un informe sobre el Q4."

  - type: assistant
    content: "¡Con gusto te ayudo a escribir un informe del Q4! Sin embargo, necesito más información:\n\n- ¿Para qué empresa u organización es?\n- ¿Del Q4 de qué año?\n- ¿Qué tipo de informe (financiero, operativo, marketing)?\n- ¿Quién es la audiencia?\n- ¿Qué datos debo incluir?\n\n¿Podrías proporcionarme estos detalles para crear un informe útil?"

  - type: note
    text: "La IA tiene que pedir aclaraciones porque la instrucción no le dio nada con qué trabajar. Cada ida y vuelta así pierde tiempo. Ahora limpiemos esto e intentemos con el framework completo."

  - type: clear
    reset_scratchpad: false
    note: "Empezando de nuevo con una mejor instrucción."

  - type: user
    content: "Eres un analista senior preparando un resumen para la junta directiva.\n\nLee el archivo de datos del Q4, luego escribe un resumen ejecutivo de una página para el equipo directivo. Incluir: tendencias de ingresos, crecimiento de clientes, mejora de churn y trayectoria del NPS. Formato de 3-4 viñetas con números de respaldo. Tono profesional."

  - type: tool_call
    tool: scratchpad_read
    args:
      filename: "q4-data.csv"

  - type: tool_result
    tool: scratchpad_read
    content: |
      Mes,Ingresos,Nuevos_Clientes,Tasa_Churn,NPS
      Oct,1,82M,145,3,2%,42
      Nov,2,01M,168,2,8%,45
      Dic,2,34M,203,2,1%,51
      Total Q4,6,17M,516,2,7% promedio,46 promedio
      vs Q3,+18%,+24%,-0,8pp,+6

  - type: assistant
    content: "**Resumen Ejecutivo Q4 2025**\n\n**La aceleración de ingresos superó los objetivos.** Los ingresos del Q4 alcanzaron $6.17M, un aumento del 18% respecto al Q3. Los ingresos mensuales crearon consistentemente desde $1.82M en octubre hasta $2.34M en diciembre, demostrando un impulso sostenido en lugar de un pico aislado.\n\n**La adquisición de clientes está escalando.** 516 nuevos clientes se unieron en el Q4, un aumento del 24% respecto al Q3. Solo diciembre (203 nuevos clientes) fue nuestro mejor mes de adquisición en la historia, lo que sugiere que las recientes inversiones en marketing están rindiendo frutos.\n\n**La retención está mejorando significativamente.** El churn promedio bajó a 2.7%, 0.8 puntos porcentuales por debajo del Q3. La cifra de diciembre de 2.1% se acerca a nuestro objetivo anual de menos del 2%. Esta tendencia fortalece directamente nuestra base de ingresos recurrentes.\n\n**La satisfacción del cliente tiene tendencia alcista.** El NPS mejoró de 42 en octubre a 51 en diciembre (promedio 46 para el trimestre, +6 sobre el Q3). Cruzar el umbral de 50 en diciembre es significativo — los puntajes por encima de 50 se consideran 'excelentes' en B2B SaaS."

  - type: note
    text: "Misma IA. Mismos datos. Resultado completamente diferente. La primera instrucción obtuvo una petición de aclaración. La segunda obtuvo un resumen ejecutivo terminado. La diferencia es el framework: Rol (analista senior), Tarea (resumen ejecutivo), Contexto (archivo de datos del Q4), Formato (viñetas con números), Restricciones (una página, tono profesional)."
```

```quiz
id: framework-elements
type: multiple-choice
question: "Tu instrucción dice 'Analiza estos datos' y adjuntas un CSV. La IA da un resumen superficial. ¿Cuál es el elemento más impactante por añadir?"
options:
  - "Rol: 'Eres un analista de datos senior' para que aporte una experiencia analítica más profunda"
  - "Formato: 'Preséntalo como una tabla con flechas de tendencia' para que la salida sea más estructurada"
  - "Restricciones: 'Máximo 500 palabras' para que se centre en lo importante"
answer: 0
explanation: "Establecer un rol cambia la profundidad y el enfoque del análisis. Un 'analista de datos senior' aporta reconocimiento de patrones de nivel experto e identifica insights que un asistente genérico pasa por alto. El formato y las restricciones mejoran la presentación pero no cambian la profundidad analítica."
```
