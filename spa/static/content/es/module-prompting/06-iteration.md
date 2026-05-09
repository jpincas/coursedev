---
title: "El Patrón de Iteración"
duration: "15m"
tags: [iteración, refinamiento, feedback, conversación]
---

# La conversación como refinamiento

Las interacciones más poderosas con la IA rara vez son de un solo intento. Son conversaciones donde iteras hasta obtener exactamente lo que necesitas.

Este es el patrón de "conversación como refinamiento." Así es como los expertos obtienen resultados 10x mejores que los principiantes.

```callout
type: warning
title: "La mayor fuente de frustración"
content: "Cuando la IA te da una salida decepcionante, casi nunca es porque la IA es tonta. Es porque no iteraste. Esperar una salida perfecta desde el primer intento es como esperar que un colega te lea la mente. La frustración viene de expectativas poco realistas, no de limitaciones de la IA."
```

## El patrón central y los errores comunes

![Feedback Vago vs Específico](/content/module-prompting/images/feedback-comparison.svg)

El patrón de iteración:

1. **Empieza con una petición inicial clara** con restricciones
2. **Revisa la salida de forma crítica** identificando problemas específicos
3. **Da feedback dirigido** referenciando partes específicas
4. **Itera con precisión**

Repite hasta tener lo que necesitas.

```callout
type: info
title: "Por qué funciona"
content: "Los modelos de IA son excelentes refinando. Pueden tomar feedback y ajustarse. La primera salida es el borrador rough — tu trabajo es dirigirlo hacia la versión final."
```

### El error más común de principiantes

**Feedback mal:**
```
Hazlo mejor.
```

**¿Qué significa "mejor"?**
- ¿Más detallado? ¿Menos detallado?
- ¿Diferente tono?
- ¿Diferente estructura?
- ¿Diferente enfoque de contenido?

La IA tiene que adivinar. Los resultados son impredecibles.

**Feedback bien:**
```
El segundo párrafo es demasiado técnico para esta audiencia.
Simplifica la explicación y elimina el jargon.
La conclusión es demasiado abrupta — añade una llamada a la acción específica.
```

Ahora la IA sabe exactamente qué ajustar.

### Lo específico gana sobre lo genérico

**Genérico:**
```
Este reporte necesita mejoras.
```

**Específico:**
```
Este reporte:
- Carece de análisis de competidores en la sección 2
- Usa demasiado jargon para una audiencia no técnica
- Necesita ejemplos concretos en la sección de recomendaciones
- Debe ser 30% más corto en general
```

El feedback específico obtiene mejoras específicas.

## Cómo dar feedback efectivo

Referencia partes específicas de la salida.

**Vago:**
```
El tono está mal.
```

**Específico:**
```
El párrafo de apertura tiene el tono profesional correcto.
La sección intermedia ("Nuestro análisis muestra...") se vuelve demasiado casual.
Igual todo lo demás al tono del opening.
```

**Aún mejor:**
```
"Sin embargo, totalmente necesitamos pensar en..." — elimina "totalmente" y "necesitamos pensar en."
Cámbialo a: "Sin embargo, debemos considerar..."
```

Cita el texto que estás referenciando. Sé quirúrgico.

```callout
type: tip
title: "El bisturí del cirujano, no el mazo"
content: "El feedback de precisión produce resultados de precisión. Identifica la exacta frase, párrafo o sección que necesita cambio. No digas solo 'corrígelo.'"
```

### La prueba del colega

Muestra tu feedback a un colega. Si él no sabría qué cambiar, la IA tampoco.

**El colega se confundiría:** "El estilo necesita trabajo."

**El colega sabría qué hacer:** "Reemplaza las viñetas con pasos numerados. Cambia los títulos a preguntas. Añade un resumen de una frase al inicio de cada sección."

Instrucciones claras para humanos son instrucciones claras para la IA.

## Obsérvalo en acción

Observa cómo el patrón de iteración mejora una propuesta real a través de tres rondas de feedback específico:

```agent-demo
path: /content/module-prompting/agent-iteration-demo.yaml
```

## Dos movimientos potentes para la iteración

Más allá del feedback específico, dos técnicas mejoran consistentemente los resultados de la iteración.

### "¿Qué te parece?"

Termina tus instrucciones o seguimientos con **"¿Qué te parece?"** o **"¿Te parece correcto?"**

Esto activa a la IA para evaluar críticamente su propio trabajo. En lugar de simplemente entregar la salida, revisa lo que produjo y marca potenciales problemas — vacíos que notó, suposiciones que hizo, áreas donde está menos seguro.

```callout
type: tip
title: "Ejemplo práctico"
content: "En lugar de: 'Escribe una propuesta de proyecto para la nueva función.' Intenta: 'Escribe una propuesta de proyecto para la nueva función. Luego dime — ¿qué te parece? ¿Qué falta o es débil?' La IA frecuentemente detecta problemas que de otro modo requerirían tu revisión para encontrar."
```

Esto funciona porque los modelos son mejores evaluando texto que generándolo desde cero. Pedir auto-evaluación aprovecha esta asimetría.

### "Hazlo prometer"

Cuando la IA sigue ignorando una instrucción específica — usa viñetas cuando le pediste párrafos, o sigue siendo formal cuando le pediste casual — **hazla reconocer explícitamente la restricción antes de proceder.**

**La técnica:**
```
Antes de escribir la siguiente versión, confirma que entiendes
estas reglas:
1. Sin viñetas — solo párrafos
2. Tono casual, como si escribieras a un amigo
3. Menos de 200 palabras

¿Cuáles son las reglas que seguirás?
```

Hacer que la IA repita la restricción en sus propias palabras mejora dramáticamente el cumplimiento. Es equivalente a pedirle a un colega "¿Puedes repetir lo que acabo de pedir?" — fuerza la atención a la instrucción específica.

```callout
type: tip
title: "Cuándo usar esto"
content: "Reserva 'Hazlo prometer' para problemas persistentes donde la IA sigue ignorando una instrucción específica a pesar del feedback claro. Para la mayoría de tareas, el feedback específico es suficiente. Esta es la técnica de escalación."
```

## Cuándo dejar de iterar

Detente cuando:
- La salida cumple tus criterios de éxito (de "¿Cómo se ve terminado?")
- Los cambios adicionales serían preferencia personal, no mejoras
- Estarías satisfecho enviando esto a tu audiencia real

No iteres eternamente. Los rendimientos decrecientes aparecen después de 2-4 rondas para la mayoría de tareas.

```callout
type: warning
title: "La sobre-iteración es real"
content: "Después de 4-5 rondas, a menudo lo estás haciendo diferente, no mejor. Si no puedes articular una mejora clara, ya terminaste."
```

### Iteración vs empezar de nuevo

**Continúa iterando cuando:**
- La base es buena, solo necesita refinamiento
- Cada iteración se acerca más a lo que quieres
- El contexto de rondas previas es valioso

**Empieza de nuevo cuando:**
- La salida está fundamentalmente en dirección equivocada
- Has iterado 5+ veces sin convergencia
- El alcance de la tarea ha cambiado significativamente

Un inicio fresco con una mejor instrucción a menudo vence a la iteración interminable sobre una mala base.

### La mentalidad de iteración

Piensa en la IA como un colega hábil pero junior que produce buenos primeros borradores rápidamente. Tú das dirección. Ellos refinan. Tú verificas. Repite hasta terminar.

**Tu trabajo:** Dirección clara, feedback específico, verificación de calidad

**Trabajo de la IA:** Ejecución rápida, reconocimiento de patrones, refinamiento

Esta división del trabajo es de donde viene la productividad de la IA.

```quiz
id: iteration-feedback
type: multiple-choice
question: "Revisaste el primer borrador de la IA y necesita mejoras. ¿Qué feedback producirá el mejor segundo borrador?"
options:
  - "'El tono necesita trabajo' — dándole a la IA libertad para interpretar qué quieres decir"
  - "'El párrafo 3 usa jargon que nuestros clientes no entienden. Reemplaza los términos técnicos con lenguaje llano y añade un ejemplo concreto después de la primera frase.'"
  - "'Hazlo mejor y más profesional' — manteniendo el feedback amplio para que la IA pueda mejorar todo"
answer: 1
explanation: "La precisión quirúrgica gana sobre la dirección amplia. Referenciar párrafos exactos, citar texto específico y dar instrucciones concretas le da a la IA exactamente qué cambiar. El feedback vago fuerza adivinanzas, lo que a menudo mejora algunas cosas y empeora otras."
```
