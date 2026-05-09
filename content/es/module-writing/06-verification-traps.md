---
title: "Qué Puede Salir Mal con Datos y Documentos"
duration: "10m"
tags: [verification, risks, quality]
---

# Qué Puede Salir Mal con Datos y Documentos

Los documentos generados por IA y el análisis de datos aparecen sofisticados. A veces son fundamentalmente incorrectos.

## El Problema de la Falsa Sofisticación

El Foro Económico Mundial advirtió en enero de 2026:

**"Cuando los datos fundamentales son fragmentados o inexactos, los modelos de IA generan salidas que parecen sofisticadas pero son fundamentalmente incorrectas."**

La salida parece profesional. Los gráficos están pulidos. La prosa es segura. Pero el análisis se basa en datos mal interpretados o conexiones fabricadas.

## Modos de Fallo en el Análisis de Datos

**Claude tiende a alucinar cuando trabaja con datasets grandes o demasiados filtros.**

Problemas comunes:
- Malinterpretar el significado de las columnas
- Aplicar filtros incorrectamente
- Inventar correlaciones que no existen
- Calcular estadísticas sobre el subconjunto incorrecto de datos

**El Interpretador de Código de ChatGPT no puede realmente entender una visualización** sin una biblioteca de percepción visual para extraer texto de los gráficos.

Genera el gráfico. No puede "ver" lo que creó.

```callout
type: warning
title: "Inspecciona Siempre el Código Generado"
content: "Haz clic en 'ver análisis' o 'mostrar código' para inspeccionar lo que la IA hizo realmente. No confíes en descripciones resumidas del análisis sin ver la lógica subyacente."
```

## Trampas en la Síntesis de Documentos

La IA es mejor en **organizar y resumir** que en **hacer una síntesis intelectual genuina**.

Cuando le pides que sintetice múltiples documentos, verifica:
- Que no fabricó conexiones entre fuentes
- Que los patrones reclamados realmente existen en el material fuente
- Que no atribuyó ideas erróneamente entre documentos
- Que distinguió correlación de causalidad

La IA destaca en: "¿Qué temas aparecen en los tres documentos?"

La IA tiene dificultades con: "¿Cuál es la implicación más importante a través de estos documentos?"

## El Problema de las Citas Fabricadas

La IA a veces inventa citas que suenan plausibles.

"Según un estudio de McKinsey de 2024..." — ¿existió ese estudio? Compruébalo.

"Investigación de Stanford en 2025 encontró..." — verifica la afirmación antes de citarla en tu trabajo.

Cuando la IA referencia fuentes externas en análisis o escritura, trata cada cita como sospechosa hasta que se verifique.

## Lista de Verificación

Antes de confiar en análisis o documentos generados por IA:

![Lista de Verificación](/content/module-writing/images/verification-checklist.svg)

**Para análisis de datos:**
- Inspecciona el código generado
- Verifica que los datos se interpretaron correctamente
- Comprueba cálculos manualmente con muestras
- Confirma que los gráficos coinciden con los datos subyacentes
- Verifica que los tamaños de muestra son apropiados

**Para síntesis de documentos:**
- Verifica que las conexiones reclamadas existen en el material fuente
- Comprueba que las citas son exactas
- Confirma que las estadísticas coinciden con las fuentes
- Asegúrate de que no se infiere causalidad de correlación
- Verifica que el tono y las conclusiones se alinean con la evidencia

**Para ambos:**
- Pide a un experto del dominio que revise
- Ejecuta la misma consulta a través de un modelo de IA diferente
- Pide a la IA que verifique su propio razonamiento
- Comprueba referencias y citas

```callout
type: info
title: "Coste de la Verificación"
content: "Como vimos antes, los trabajadores del conocimiento invierten tiempo significativo verificando las salidas de la IA. Esto no es tiempo perdido — es control de calidad esencial. El tiempo ahorrado por el trabajo acelerado con IA supera el coste de verificación, pero solo si la verificación realmente ocurre."
```

## El Coste de No Verificar

Como vimos antes, una parte significativa de los usuarios empresariales de IA han tomado decisiones importantes de negocio basadas en contenido alucinado.

La salida parecía creíble. Nadie comprobó. La decisión fue incorrecta.

La verificación no es opcional. Es la diferencia entre trabajo acelerado por IA y errores generados por IA.

## Detectando la Falsa Sofisticación

Así es como se ve en la práctica. La IA analiza datos, produce un informe pulido — y se equivoca en algo sutil. ¿Puedes detectarlo?

```agent
id: sophisticated-wrongness-demo
title: "Cuando la Salida Pulida es Incorrecta"
model_label: "Claude"

system: |
  Eres un analista de mercado. Proporciona análisis claro y basado en datos.
  Sé seguro y profesional en tu presentación.

scratchpad:
  "market-data.csv": |
    Trimestre,Nuestros_Ingresos,Tamaño_Mercado,Competidor_A,Competidor_B,Nuestros_Clientes
    T1_2025,£2,1M,£45M,£8,2M,£5,1M,1200
    T2_2025,£2,4M,£47M,£8,5M,£5,3M,1350
    T3_2025,£2,8M,£52M,£9,1M,£5,0M,1480
    T4_2025,£3,2M,£58M,£9,8M,£4,7M,1690

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
    text: "La IA está a punto de analizar datos de mercado y producir un informe seguro. Lee la salida con atención. Algo está mal — pero suena completamente convincente."

  - type: user
    content: "Lee los datos del mercado y dame un análisis competitivo. ¿Cómo estamos rindiéndonos respecto al mercado?"

  - type: tool_call
    tool: scratchpad_read
    args:
      filename: "market-data.csv"

  - type: tool_result
    tool: scratchpad_read
    content: |
      Trimestre,Nuestros_Ingresos,Tamaño_Mercado,Competidor_A,Competidor_B,Nuestros_Clientes
      T1_2025,£2,1M,£45M,£8,2M,£5,1M,1200
      T2_2025,£2,4M,£47M,£8,5M,£5,3M,1350
      T3_2025,£2,8M,£52M,£9,1M,£5,0M,1480
      T4_2025,£3,2M,£58M,£9,8M,£4,7M,1690

  - type: assistant
    content: "**Análisis Competitivo: T1-T4 2025**\n\n**Fuerte rendimiento en todos frentes.**\n\nNuestros ingresos crecieron un 52% durante el año (£2,1M → £3,2M), superando significativamente el crecimiento general del mercado del 29% (£45M → £58M). Estamos ganando cuota de mercado.\n\n**Posicionamiento competitivo:**\n- **Competidor A** creció un 20% (£8,2M → £9,8M) — estable pero por debajo del ritmo del mercado. Está perdiendo cuota.\n- **Competidor B** cayó un 8% (£5,1M → £4,7M) — claramente en problemas. A la baja en un mercado en crecimiento.\n- **Nosotros** crecimos un 52% — los más rápidos del mercado.\n\n**Trayectoria de cuota de mercado:**\n- T1: 4,7% → T4: 5,5% — nuestra cuota se está expandiendo\n- La cuota combinada de los competidores cayó de 29,6% a 25,0%\n\n**Crecimiento de clientes** sigue los ingresos: 1.200 → 1.690 (crecimiento del 41%). Los ingresos medios por cliente aumentaron de £1.750 a £1.893 — estamos subiendo en el mercado.\n\n**Evaluación:** Somos el líder de crecimiento en este mercado. Continuar la trayectoria actual."

  - type: note
    text: "Suena bien, ¿verdad? Profesional, basado en datos, seguro. Ahora verifiquemos las afirmaciones."

  - type: user
    content: "Vuelve atrás y verifica cada porcentaje y cálculo que citaste. Comprueba cada uno contra los datos reales."

  - type: assistant
    content: "Permiso que verifique:\n\n**Nuestro crecimiento de ingresos 52%:** (3,2 - 2,1) / 2,1 = 52,4%. ✓ Correcto.\n\n**Crecimiento del mercado 29%:** (58 - 45) / 45 = 28,9%, redondeado a 29%. ✓ Correcto.\n\n**Crecimiento del Competidor A 20%:** (9,8 - 8,2) / 8,2 = 19,5%, redondeado a 20%. ✓ Correcto.\n\n**Caída del Competidor B 8%:** (4,7 - 5,1) / 5,1 = -7,8%, redondeado a -8%. ✓ Correcto.\n\n**Cuota de mercado T1 4,7%:** 2,1 / 45 = 4,67%. ✓ Correcto.\n**Cuota de mercado T4 5,5%:** 3,2 / 58 = 5,52%. ✓ Correcto.\n\n**Cuota combinada competidores T1:** (8,2 + 5,1) / 45 = 29,6%. ✓ Correcto.\n**Cuota combinada competidores T4:** (9,8 + 4,7) / 58 = 25,0%. ✓ Correcto.\n\n**Ingresos medios por cliente T1:** 2.100.000 / 1.200 = £1.750. ✓ Correcto.\n**Ingresos medios por cliente T4:** 3.200.000 / 1.690 = £1.893. ✓ Correcto.\n\n**Todos los cálculos verificados.** Sin embargo, debo señalar algo que no mencioné inicialmente: **nuestra cuota de mercado del 5,5% significa que el 94,5% del mercado NO son nuestros clientes.** El tono celebratorio de mi análisis puede ser prematuro — crecemos rápido pero seguimos siendo un jugador muy pequeño. La evaluación \"continuar la trayectoria actual\" ignora la pregunta de si nuestra tasa de crecimiento es sostenible, o si el Competidor A (con 16,9% de cuota) podría responder con precios agresivos."

  - type: note
    text: "Esta vez los números fueron correctos, pero fijáos en lo que reveló la verificación: el marco era engañoso. Una cuota de mercado del 5,5% se presentó como \"líder de crecimiento\" sin reconocer que somos un jugador minúsculo. La auto-verificación de la IA detectó el sesgo, no en las matemáticas, sino en la narrativa."

  - type: note
    text: "Esto es falsa sofisticación en acción. El análisis no era matemáticamente incorrecto — cada número se verificó. Pero la conclusión (\"continuar la trayectoria actual\") ignoró la realidad estratégica. Salida pulida. Aritmética correcta. Conclusión engañosa. Verifica siempre no solo los números, sino la historia que se está contando con ellos."
```

```quiz
id: verification-traps-quiz
type: multiple-choice
question: "¿Qué significa 'falsa sofisticación' en el contexto del análisis generado por IA?"
options:
  - "La IA usa métodos demasiado complejos que son difíciles de entender"
  - "Las salidas de la IA aparecen profesionales y pulidas pero se basan en datos mal interpretados o conexiones fabricadas"
  - "La IA genera documentos demasiado sofisticados para la audiencia pretendida"
  - "La IA crea visualizaciones innecesariamente detalladas"
answer: 1
explanation: "La falsa sofisticación significa que la IA genera salidas que parecen profesionales — gráficos pulidos, prosa segura — pero son fundamentalmente incorrectas debido a datos mal interpretados o patrones fabricados. La sofisticación oculta el error."
```
