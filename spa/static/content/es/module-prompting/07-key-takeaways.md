---
title: "Ideas clave: Prompting"
duration: "5m"
tags: [resumen, ideas clave]
---

# Ideas clave: Prompting

Ahora tienes un toolkit completo para comunicarte con la IA de forma efectiva. Estas técnicas no son teóricas — son la diferencia entre una salida genérica y exactamente lo que necesitas. Aquí está el resumen esencial.

## 1. Resultados sobre proceso

Describe cómo se ve terminado, no cómo llegar allí.

La IA figuring out el proceso. Tu trabajo es definir el destino. "Produce un análisis de tendencias con las 5 principales insights como viñetas" gana a "Primero analiza los datos, luego identifica tendencias, luego resúmelos."

## 2. Las Tres Preguntas

Antes de cualquier tarea, pregúntate:
- **¿Cómo se ve "terminado"?** — Define el estado final
- **¿Qué contexto necesita?** — Proporciona la información requerida
- **¿Cuáles son los límites?** — Establece las restricciones

Responde estas tres y tendrás una especificación sólida.

## 3. Las técnicas que entregan resultados

De mayor a menor impacto:
- **Sé específico y directo** — Siempre aplicable. Mayor impacto.
- **Usa ejemplos (multishot prompting)** — 2-5 ejemplos enseñan el patrón que quieres
- **Estructura con etiquetas XML** — Separa instrucciones, contexto, ejemplos, restricciones
- **Activa el pensamiento** — Para tareas de razonamiento difíciles, deja que el modelo piense profundamente
- **Encadena instrucciones** — Divide el trabajo multi-paso en subtareas secuenciales

## 4. La iteración es cómo trabajan los expertos

La primera salida es el borrador rough. Da feedback específico y dirigido referenciando partes exactas.

"Hazlo mejor" → la IA adivina
"El párrafo 2 es demasiado técnico — simplifícalo" → la IA sabe exactamente qué cambiar

## 5. Meta-prompting y especificación interactiva

Deja que la IA te ayude a escribir mejores instrucciones. Pídele que te haga preguntas de aclaración. Para entregables complejos, usa especificación interactiva: acuerda la estructura y los requisitos antes de ejecutar. La especificación se convierte en un contrato que asegura alineación.

Las preguntas de la IA revelan lo que se necesita. Tus respuestas construyen la instrucción. Para proyectos significativos, este enfoque de conversación antes de ejecución transforma la calidad.

## 6. La información completa gana sobre el lenguaje ingenioso

No se trata de palabras mágicas o frases secretas.

Se trata de darle a la IA la información que necesita para tener éxito. Tarea clara, contexto relevante, formato explícito.

```callout
type: tip
title: "El cambio de mentalidad"
content: "Piensa en la IA como un trabajador capaz pero literal. No puede leerte la mente. Sé explícito sobre lo que quieres, proporciona lo que necesita, y entregará."
```

## Checklist práctica

Antes de tu próxima tarea con IA, repasa esto:

- [ ] ¿Puedo describir exactamente cómo se ve "terminado"?
- [ ] ¿He proporcionado todo el contexto y datos relevantes?
- [ ] ¿He especificado formato y longitud?
- [ ] ¿He mencionado alguna restricción o exclusión?
- [ ] ¿Un humano competente tendría suficiente para completar esta tarea?
- [ ] ¿He proporcionado 2-3 ejemplos si el formato/estilo importa?
- [ ] ¿Estoy preparado para iterar con feedback específico?

Si respondes "no" a alguna de estas, tu petición probablemente necesita más detalle.

```quiz
id: prompting-takeaway
type: multiple-choice
question: "Necesitas un reporte de análisis de competidores. Tienes 15 minutos. ¿Qué combinación de técnicas producirá el mejor resultado?"
options:
  - "Usar las tres preguntas para definir terminado, proporcionar archivos de datos de competidores como contexto, e iterar una vez con feedback específico"
  - "Escribir una instrucción larga usando etiquetas XML, activar pensamiento extendido, y encadenarla en tres pasos"
  - "Usar meta-prompting para que la IA escriba la instrucción perfecta, luego ejecutar esa instrucción"
answer: 0
explanation: "La combinación de mayor impacto es: definir cómo se ve terminado (las tres preguntas), proporcionar contexto completo (archivos de datos), e iterar con precisión. Los otros enfoques son técnicas válidas pero sobre-engineeran una tarea de 15 minutos. Igual la complejidad de la técnica a la complejidad de la tarea."
```

## Siguiente

Ahora que sabes cómo comunicarte con la IA de forma efectiva, hablemos sobre trabajar con **archivos**. Pasar del chat a entregables reales es donde la IA se vuelve verdaderamente útil para el trabajo.
