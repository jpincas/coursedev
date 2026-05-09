---
title: "Cuando la IA Escribe Código para Pensar"
duration: "15m"
tags: [analysis, code, data, computation]
---

# Cuando la IA Escribe Código para Pensar

La página anterior te mostró el flujo de trabajo de subida y análisis y las herramientas disponibles. Ahora veamos lo que realmente ocurre por dentro — porque hay dos cosas fundamentalmente diferentes que la IA puede hacer con tus datos, y entender la distinción cambiará la forma en que trabajas.

## Dos Modos de Análisis

**Modo 1: Leer y Razonar**

La IA lee tus datos en su ventana de contexto y razona sobre ellos usando lenguaje. Escanea los números, detecta patrones y describe lo que ve — muy parecido a una persona que echa un vistazo a una mesa y resume las tendencias.

Esto es lo que pasó en la demo de descripción primero de la página anterior. La IA leyó el CSV de la encuesta de empleados y razonó sobre él lingüísticamente. Detectó que Soporte tenía bajo compromiso y que los puntajes de manager seguían el compromiso — todo "leyendo" los datos.

**Modo 2: Escribir y Ejecutar Código**

La IA escribe un script en Python para procesar tus datos, lo ejecuta en un entorno aislado y devuelve los resultados computados. El código hace el cálculo real — calcular promedios, ejecutar correlaciones, proyectar tendencias.

Esto es lo que ocurre cuando ChatGPT muestra "Analizando..." o Claude muestra una llamada de herramienta "Análisis".

## Por Qué Importa

Aquí está la idea crítica: **los modelos de lenguaje no pueden hacer aritmética de forma fiable.** Predicen el siguiente token — no tienen una calculadora incorporada. Pide a un modelo que sume 500 números "de cabeza" y se equivocará. Pídele que escriba Python para sumar 500 números, ejecute el código y devuelva el resultado — y será exacto cada vez.

| | Leer y Razonar | Escribir y Ejecutar Código |
|---|---|---|
| **Precisión** | Aproximada — puede calcular mal | Exacta — cálculo real |
| **Escala** | Le cuesta con datasets grandes | Maneja millones de filas |
| **Mejor para** | Temas, resúmenes, cualitativo | Estadísticas, tendencias, correlaciones, proyecciones |
| **Salida** | Descripciones de texto | Números, tablas, gráficos, archivos |

Piensa en ello así: el Modo 1 es como pedirle a un compañero que eche un vistazo a una hoja de cálculo y te diga qué nota. El Modo 2 es como entregarle la hoja de cálculo a un analista que abre Python y ejecuta los cálculos. Ambos son útiles. Pero nunca confiarías el Modo 1 para tu previsión trimestral.

```callout
type: warning
title: "El Error Invisible"
content: "Cuando la IA analiza datos solo leyendo, a menudo suena segura sobre números que son incorrectos. Podría decir \"los ingresos crecieron aproximadamente un 23%\" cuando la cifra real es 18%. No puedes saber por la respuesta si calculó o adivinó. Empuja siempre a la IA a usar código para afirmaciones cuantitativas."
```

## El Lenguaje es el Puente

Aquí es donde se vuelve potente. **No necesitas saber Python.** describes lo que quieres en inglés llano, y la IA traduce tu pregunta a código preciso, lo ejecuta y te da la respuesta.

Esto es genuinamente nuevo. Anteriormente, obtener una respuesta precisa de los datos requería o saber fórmulas de Excel, saber un lenguaje de programación, o contratar a alguien que supiera. Ahora la barrera ha desaparecido. describes el análisis que quieres, y la IA escribe el código para hacerlo.

El lenguaje es lo que hace esto posible. Porque la IA habla tanto inglés como Python con fluidez, puede traducir entre tu intención y la capacidad de la máquina para calcular. Esta es una de las cosas más poderosas de trabajar con lenguaje — es la interfaz universal entre lo que quieres saber y lo que la máquina puede computar.

## Observa la Diferencia

Mismos datos, distinto enfoque. Esta vez, en lugar de solo leer y razonar sobre los números, observa cómo la IA escribe código Python para calcular cifras exactas de crecimiento y proyecciones.

```agent
id: code-analysis-demo
title: "IA Escribiendo Código para Analizar Datos"
model_label: "Claude"

system: |
  Eres un asistente analista de datos con acceso a un entorno de ejecución
  de código Python. Cuando los usuarios hagan preguntas cuantitativas sobre datos,
  escribe código Python para calcular respuestas precisas en lugar de estimar.

scratchpad:
  "data/regional-sales.csv": |
    Mes,Norte,Sur,Este,Oeste
    Ene,42300,38100,29400,51200
    Feb,44100,37800,31200,52800
    Mar,48700,39200,33100,54100
    Abr,46200,41500,35800,53400
    May,51800,40200,37200,56700
    Jun,53400,42800,39500,58200
    Jul,49800,44100,41200,55800
    Ago,55200,43500,43800,59100
    Sep,58100,45200,42100,61400
    Oct,61400,46800,44500,63200
    Nov,59800,48200,47200,65800
    Dic,64200,49500,49800,67100

tools:
  - list_files
  - scratchpad_read
  - scratchpad_write
  - run_code

visibility:
  system_prompt: hidden
  tool_calls: visible
  full_context: hidden

sidebar:
  width: "50%"
  start_open: true

script:
  - type: note
    text: "Un archivo CSV con 12 meses de datos de ventas en 4 regiones. El usuario quiere saber qué región crece más rápido. Observa cómo la IA escribe código Python real para calcular esto con precisión, en lugar de estimar visualmente los números."

  - type: user
    content: "Tengo datos de ventas regionales en data/regional-sales.csv. ¿Qué región está creciendo más rápido? ¿Y sobre la base de la tendencia, ¿cuánto deberíamos presupuestar de cada región en el T1 del próximo año?"

  - type: tool_call
    tool: scratchpad_read
    args:
      filename: "data/regional-sales.csv"

  - type: tool_result
    tool: scratchpad_read
    content: |
      Mes,Norte,Sur,Este,Oeste
      Ene,42300,38100,29400,51200
      Feb,44100,37800,31200,52800
      Mar,48700,39200,33100,54100
      Abr,46200,41500,35800,53400
      May,51800,40200,37200,56700
      Jun,53400,42800,39500,58200
      Jul,49800,44100,41200,55800
      Ago,55200,43500,43800,59100
      Sep,58100,45200,42100,61400
      Oct,61400,46800,44500,63200
      Nov,59800,48200,47200,65800
      Dic,64200,49500,49800,67100

  - type: assistant
    content: "Veo los datos. Voy a escribir Python para calcular las tasas exactas de crecimiento y las proyecciones en lugar de estimar a partir de los números en crudo."

  - type: tool_call
    tool: scratchpad_write
    args:
      filename: "analysis.py"
      content: |
        import csv
        from io import StringIO

        data = open("data/regional-sales.csv").read()
        reader = csv.DictReader(StringIO(data))
        regions = {"Norte": [], "Sur": [], "Este": [], "Oeste": []}

        for row in reader:
            for region in regions:
                regions[region].append(int(row[region]))

        print("=== Análisis de Crecimiento Regional ===\n")
        for region, values in regions.items():
            total_growth = (values[-1] - values[0]) / values[0] * 100
            monthly_increase = (values[-1] - values[0]) / 11
            projected = [values[-1] + monthly_increase * (i+1)
                        for i in range(3)]
            print(f"{region}:")
            print(f"  Ene: £{values[0]:,} → Dic: £{values[-1]:,}")
            print(f"  Crecimiento anual: {total_growth:.1f}%")
            print(f"  Proyección T1: £{sum(projected)/3:,.0f}/mes")
            print()

        growths = {r: (v[-1]-v[0])/v[0]*100
                  for r, v in regions.items()}
        fastest = max(growths, key=growths.get)
        print(f"Mayor crecimiento: {fastest} ({growths[fastest]:.1f}%)")

  - type: tool_result
    tool: scratchpad_write
    content: "Escrito en analysis.py (26 líneas)"

  - type: note
    text: "La IA ha escrito un script en Python. Ahora lo ejecuta. En herramientas como ChatGPT y Claude, esto se ejecuta en un entorno aislado seguro — el código realmente se ejecuta y produce resultados computados reales."

  - type: tool_call
    tool: run_code
    args:
      filename: "analysis.py"

  - type: tool_result
    tool: run_code
    content: |
      === Análisis de Crecimiento Regional ===

      Norte:
        Ene: £42.300 → Dic: £64.200
        Crecimiento anual: 51,8%
        Proyección T1: £66.191/mes

      Sur:
        Ene: £38.100 → Dic: £49.500
        Crecimiento anual: 29,9%
        Proyección T1: £51.573/mes

      Este:
        Ene: £29.400 → Dic: £49.800
        Crecimiento anual: 69,4%
        Proyección T1: £53.564/mes

      Oeste:
        Ene: £51.200 → Dic: £67.100
        Crecimiento anual: 31,1%
        Proyección T1: £69.518/mes

      Mayor crecimiento: Este (69,4%)

  - type: tool_call
    tool: scratchpad_write
    args:
      filename: "reports/growth-analysis.md"
      content: |
        # Análisis de Crecimiento de Ventas Regionales

        ## Hallazgo Clave
        **Este es la región de mayor crecimiento con un 69,4% de crecimiento anual** — más del doble de la tasa del Sur (29,9%) o del Oeste (31,1%).

        ## Desglose Completo

        | Región | Ene | Dic | Crecimiento | Proyección T1 |
        |--------|-----|-----|-------------|----------------|
        | Este | £29.400 | £49.800 | **69,4%** | £53.564/mes |
        | Norte | £42.300 | £64.200 | 51,8% | £66.191/mes |
        | Oeste | £51.200 | £67.100 | 31,1% | £69.518/mes |
        | Sur | £38.100 | £49.500 | 29,9% | £51.573/mes |

        ## Patrones Notables
        - Este empezó siendo el más pequeño pero está cerrando la brecha rápidamente
        - Norte tiene fuertes ingresos Y fuerte crecimiento — proyectado a superar al Oeste para el T2
        - Oeste lidera en ingresos absolutos pero crece lentamente
        - Sur es el más lento en crecimiento y en términos absolutos

  - type: tool_result
    tool: scratchpad_write
    content: "Escrito en reports/growth-analysis.md (19 líneas)"

  - type: assistant
    content: |
      Aquí está el análisis. **Este es tu región de mayor crecimiento con un 69,4% de crecimiento** — más del doble de la tasa del Sur o del Oeste.

      La historia interesante es Norte vs Oeste. Oeste lidera actualmente en ingresos (£67.100/mes en diciembre), pero Norte está creciendo más rápido (51,8% vs 31,1%). A los ritmos actuales, Norte supera a Oeste para el T2.

      Para el presupuesto del T1: Oeste ~£69,5K/mes, Norte ~£66,2K/mes, Este ~£53,6K/mes, Sur ~£51,6K/mes. Informe completo guardado en reports/growth-analysis.md.

  - type: note
    text: "La IA no estimó estos números. Escribió Python para calcular tasas exactas de crecimiento, promedios y proyecciones. La cifra del 69,4% para Este está calculada con precisión a partir de los datos — no es una aproximación grosera. Esta es la diferencia entre la IA 'leyendo' tus datos y la IA 'computando' con tus datos."

  - type: note
    text: "¿Podrías haber hecho esto en Excel? Sí, si sabes las fórmulas correctas. ¿En Python? Solo si sabes programar. Con IA, describiste lo que querías en inglés y hizo el cálculo por ti. El lenguaje fue el puente."
```

## Reconociendo los Dos Modos

La mayoría de herramientas de IA ahora indican claramente cuándo usan ejecución de código:

- **ChatGPT** muestra "Analizando..." y muestra el código Python que escribe
- **Claude** muestra una llamada de herramienta "Análisis" con código expandible
- **Google Gemini** muestra "Ejecutando código" con bloques expandibles

Si no ves estos indicadores, la IA probablemente está razonando sobre tus datos lingüísticamente en lugar de computar. Para preguntas cuantitativas, esto importa enormemente.

```callout
type: tip
title: "Forzar la Ejecución de Código"
content: "Si quieres asegurar que la IA use código en lugar de estimar, sé explícito: 'Escribe código Python para calcular esto' o 'Usa análisis de código para computar las cifras exactas.' Esto empuja a la herramienta a su modo de ejecución de código."
```

## Cuándo Brilla Cada Modo

**Leer y razonar es perfecto para:**
- Resumir informes y documentos
- Identificar temas en datos cualitativos (reseñas, feedback, entrevistas)
- Comparar contenido entre documentos
- Explicar qué significan los datos en contexto

**Busca ejecución de código cuando:**
- Necesitas cálculos exactos (promedios, tasas de crecimiento, totales)
- Trabajas con más de unas pocas decenas de números
- Necesitas análisis estadístico (correlaciones, distribuciones)
- Quieres gráficos o visualizaciones
- Necesitas transformar o limpiar datos
- La respuesta implica algún tipo de predicción o proyección

```quiz
id: code-vs-reading
type: multiple-choice
question: "Subes un CSV con 2.000 filas de datos de clientes y preguntas '¿Cuál es el valor medio del pedido para los clientes que se registraron en 2024?' La IA responde: 'Según los datos, el valor medio parece estar alrededor de £45-50.' ¿Qué debería preocuparte?"
options:
  - "Nada — la IA leyó los datos y dio una respuesta razonable"
  - "El rango vago sugiere que la IA estimó en lugar de calcular — deberías pedirle que escriba código para computar la cifra exacta"
  - "2.000 filas es demasiada data para que la IA procese con precisión"
answer: 1
explanation: "Un rango vago como '£45-50' es una señal de que la IA leyó los datos y estimó en lugar de computar la respuesta exacta. Con 2.000 filas, siempre debes empujar hacia la ejecución de código. Pide a la IA que 'escriba Python para calcular el valor medio exacto del pedido' y obtendrás un número preciso como £47,32 — no una aproximación."
```
