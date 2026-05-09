---
title: "Verificación de calidad"
duration: "10m"
tags: [verificación, calidad, revisión]
---

# Verificación de calidad

Necesitas verificar los resultados de la IA. Pero ¿cómo verificas un trabajo cuando no eres experto en esa área específica?

## El marco de verificación

No necesitas ser experto en el contenido. Necesitas una estrategia de verificación reflexiva.

![Marco de verificación de cuatro pasos](/content/module-delegation/images/verification-framework.svg)

### 1. Verificar datos concretos

Elige 3-5 afirmaciones concretas y verifícalas de forma independiente.
- Compara una estadística citada con la fuente original
- Verifica una fecha o nombre mencionado
- Confirma un detalle técnico

Si estas verificaciones puntuales superan la prueba, aumenta la confianza en el resto. Si fallan, sabes que se necesita una revisión más profunda.

**Técnica avanzada: Validación cruzada multi-modelo**
Ejecuta la misma consulta a través de múltiples LLMs (Claude, GPT, Gemini). Si coinciden, aumenta la confianza. Si divergen, investiga más a fondo.

### 2. Comprobar la coherencia

¿Tiene sentido interno?
- ¿Hay contradicciones?
- ¿El lógica se sigue?
- ¿Las conclusiones se derivan de la evidencia?
- ¿Hay brechas en el razonamiento?

No necesitas ser experto para notar "esta parte contradice a esa parte."

### 3. Validar la estructura

¿Tiene lo que pediste?
- ¿Todas las secciones requeridas están presentes?
- ¿El formato coincide con los requisitos?
- ¿La longitud es apropiada?
- ¿El tono es correcto?

Esta es una verificación contra tu especificación, no una verificación de la precisión del contenido.

### 4. Comprobar contra el propósito

¿Funcionaría realmente?
- ¿Alcanza el objetivo?
- ¿La audiencia lo encontraría útil?
- ¿Puedes usarlo para su propósito intended?

La prueba definitiva: ¿es apto para el propósito?

```callout
type: tip
title: "Mentalidad de verificación"
content: "No estás comprobando si la IA tiene 'razón.' Estás comprobando si el resultado es útil, suficientemente preciso para tus propósitos, y apto para el uso intendido."
```

### 5. La IA como editor

Una de las técnicas de verificación más efectivas: **pedirle a la IA que revise su propio trabajo.**

Después de que la IA produzca contenido, no lo aceptes simplemente. Encarga una segunda pasada:

> "Ahora relee lo que acabas de escribir desde la perspectiva de [audiencia objetivo]. ¿Qué cambiarías?"

Esto activa una verdadera segunda pasada, no un mero sellado de aprobación. La IA detectará:
- Jerga que la audiencia no entenderá
- Brechas lógicas en su propio razonamiento
- Contexto faltante que asumió
- Desajustes de tono

**Revisión multi-pasada:** Pídele a la IA que revise su trabajo a través de diferentes filtros. Pídele que lo lea una vez por precisión, otra por claridad, otra por completitud, otra por tono. Cada pasada detecta problemas diferentes.

```callout
type: tip
title: "Diferentes filtros, diferentes detecciones"
content: "La precisión detecta errores factuales. La claridad detecta explicaciones confusas. La completitud detecta secciones faltantes. El tono detecta lenguaje inapropiado. Una sola pasada de revisión rara vez detecta las cuatro cosas."
```

### 6. Verificación visual

Para resultados visuales — sitios web, documentos formatados, presentaciones — **usa herramientas del navegador para realmente MIRAR lo que se produjo.**

No confíes simplemente en el código. Renderízalo y comprueba.

Problemas visuales comunes que la IA no detecta en la revisión del código:
- Elementos desalineados
- Diseños responsivos rotos
- Problemas de contraste de color
- Texto desbordado
- Imágenes faltantes o enlaces rotos

Si la IA construyó una página web, ábrela en un navegador. Si la IA formató un documento, exportalo y revisa el PDF. La corrección del código ≠ calidad visual.

## Consejos prácticos

**No verifiques todo**
Eso anula el propósito de la delegación. Verifica lo suficiente para tener confianza.

**Verifica las afirmaciones de alto riesgo con más cuidado**
Una estadística incorrecta en una presentación al consejo importa más que un error tipográfico en notas internas.

**Usa la IA para ayudar a verificar**
Pídele a la IA que verifique sus propias afirmaciones, o usa una sesión diferente de IA para revisar. La investigación muestra que pedirle a la IA que verifique su propio razonamiento antes de proporcionar respuestas finales reduce los errores aproximadamente un 17%.

**Para análisis de datos, inspecciona siempre el código**
Cuando la IA genera gráficos o estadísticas, haz clic en "ver análisis" para inspeccionar el código generado. Claude tiende a alucinar cuando trabaja con conjuntos de datos grandes o demasiados filtros. El Foro Económico Mundial advirtió: "Cuando los datos fundamentales son fragmentados o inexactos, los modelos de IA generan resultados que parecen sofisticados pero son fundamentalmente erróneos."

**Búsqueda web como verificación**
Usa la integración de búsqueda web cuando esté disponible. La IA con acceso a búsqueda consigue mucha mayor precisión que la IA que trabaja solo con datos de entrenamiento.

**Confía pero verifica**
Los resultados de la IA son generalmente buenos. Pero "generalmente" no es "siempre." La verificación es tu responsabilidad.

## Los cuatro pasos en la práctica

El marco de verificación funciona incluso cuando no eres experto. Observa cómo alguien aplica los cuatro pasos a un análisis competitivo generado por IA.

```agent
id: verification-demo
title: "Verificando resultados de la IA paso a paso"
model_label: "Claude"

system: |
  Eres un analista de investigación de mercado. Produce un análisis exhaustivo.
  Cuando se te pida verificar, examina las afirmaciones de forma crítica.

scratchpad:
  "competitor-intel.txt": |
    Competidor: NovaTech Solutions
    Última financiación: Serie B, 12M £ (marzo de 2025)
    Plantilla: ~120 (LinkedIn)
    Producto clave: AutoFlow — automatización de flujos de trabajo para equipos legales
    Precios: 45 £/usuario/mes (Pro), 89 £/usuario/mes (Enterprise)
    Movimientos recientes:
    - Lanó función de revisión de contratos con IA (noviembre de 2025)
    - Contrató al VP de Ventas de Salesforce (enero de 2026)
    - Abrió oficina en Sídney para expansión APAC
    Debilidades conocidas: onboarding lento (promedio 6 semanas), sin aplicación móvil
    Base de clientes: ~200 firmas, mayoritariamente mercado medio

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
    text: "Primero, la IA produce un informe competitivo. Luego pasaremos por los cuatro pasos de verificación: verificar datos concretos, comprobar coherencia, validar estructura, comprobar contra el propósito."

  - type: user
    content: "Lee el archivo de inteligencia del competidor y produce una evaluación de amenaza competitiva para nuestro equipo directivo. Incluye nivel de amenaza, próximos movimientos probables y nuestra respuesta recomendada."

  - type: tool_call
    tool: scratchpad_read
    args:
      filename: "competitor-intel.txt"

  - type: tool_result
    tool: scratchpad_read
    content: |
      Competidor: NovaTech Solutions
      Última financiación: Serie B, 12M £ (marzo de 2025)
      Plantilla: ~120 (LinkedIn)
      Producto clave: AutoFlow — automatización de flujos de trabajo para equipos legales
      Precios: 45 £/usuario/mes (Pro), 89 £/usuario/mes (Enterprise)
      Movimientos recientes:
      - Lanó función de revisión de contratos con IA (noviembre de 2025)
      - Contrató al VP de Ventas de Salesforce (enero de 2026)
      - Abrió oficina en Sídney para expansión APAC
      Debilidades conocidas: onboarding lento (promedio 6 semanas), sin aplicación móvil
      Base de clientes: ~200 firmas, mayoritariamente mercado medio

  - type: assistant
    content: "**NovaTech Solutions — Evaluación de amenaza competitiva**\n\n**Nivel de amenaza: MODERADA-ALTA (en tendencia ascendente)**\n\nNovaTech ha pasado de una herramienta de flujos de trabajo a una plataforma legal potenciada con IA. Tres señales indican escalada:\n\n1. **El giro hacia IA es real.** La función de revisión de contratos (noviembre de 2025) compite directamente con nuestra oferta principal. Con 12M £ en financiación, tienen margen para iterar.\n\n2. **La capacidad comercial está aumentando.** La contratación del VP de Salesforce indica ambiciones empresariales. Espera ventas outbound agresivas y potencial presión de precios en Q2.\n\n3. **Expansión geográfica.** La oficina de Sídney apunta a APAC — un mercado donde tenemos presencia limitada. La ventaja de primer entrante en la región podría bloquearnos.\n\n**Próximos 6 meses probables:**\n- Impulso empresarial con equipo de ventas dedicado\n- Lanzamiento de aplicación móvil (la queja más grande de sus clientes)\n- Agresividad de precios para ganar logos empresariales a escala\n- Entrada al mercado APAC apuntando a firmas legales australianas y singapurense\n\n**Respuesta recomendada:**\n- Acelerar nuestras funciones de IA de contratos para mantener la diferenciación\n- Proteger cuentas empresariales existentes con ofertas de retención\n- Evaluar cronograma de expansión APAC — retrasar cuesta participación de mercado\n- Vigilar su lanzamiento de aplicación para posicionamiento competitivo"

  - type: note
    text: "Parece profesional y accionable. Ahora verifiquemos usando el marco de cuatro pasos."

  - type: user
    content: "Buen primer borrador. Ahora ayúdame a verificarlo. Vamos paso a paso.\n\nPaso 1: Verificar datos concretos. ¿Hay alguna afirmación en tu análisis que vaya más allá de lo que está en los datos de origen?"

  - type: assistant
    content: "Buena observación — déjame auditar cada afirmación contra la fuente:\n\n**Verificado desde la fuente:**\n- Serie B de 12M £ ✓\n- Función de revisión de contratos con IA noviembre 2025 ✓\n- VP de Ventas de Salesforce enero 2026 ✓\n- Oficina de Sídney / APAC ✓\n- Onboarding lento y sin aplicación móvil ✓\n- ~200 firmas, mercado medio ✓\n\n**Afirmaciones que extrapolé (no directamente en los datos):**\n- 'Pasó de herramienta de flujos de trabajo a plataforma legal potenciada con IA' — mi interpretación, no declarado\n- 'Agresividad de precios para ganar logos empresariales' — especulación, no evidenciado\n- 'Lanzamiento de aplicación móvil' como próximo movimiento probable — inferencia razonable a partir de la debilidad, pero no confirmado\n- 'Apuntando a firmas legales australianas y singapurense' — añadí los países específicos; el origen solo dice 'expansión APAC'\n\n**Debería marcar estas como inferencias, no hechos.** El análisis mezcló información con fuentes con mis predicciones sin distinguir entre ellas."

  - type: user
    content: "Paso 2: Comprobación de coherencia. ¿Alguna contradicción interna?"

  - type: assistant
    content: "Una tensión:\n\nDescribí la amenaza como 'MODERADA-ALTA (en tendencia ascendente)' pero luego listé cuatro movimientos agresivos como 'próximos 6 meses probables.' Si los cuatro ocurren, el nivel de amenaza es ALTO, no moderado-alto. La calificación y la evidencia no se alinean completamente.\n\nAdemás: recomendé 'proteger cuentas empresariales existentes' pero el origen dice que NovaTech es 'mayormente mercado medio.' Si es mercado medio, nuestras cuentas empresariales pueden no ser su objetivo inmediato. La recomendación defensiva podría ser prematura."

  - type: note
    text: "Dos pasos de verificación completados y ya se detectó: afirmaciones sin fuente presentadas como hechos, predicciones especulativas sin etiquetas, una calificación de amenaza que contradice su propia evidencia, y una recomendación defensiva basada en un supuesto incorrecto. El marco de cuatro pasos funciona incluso aunque ni tú ni la IA sean expertos en inteligencia competitiva."

  - type: note
    text: "Los pasos 3 (validar estructura) y 4 (comprobar contra el propósito) son sencillos: ¿Tiene todas las secciones solicitadas? ¿El equipo directivo lo encontraría accionable? Puedes comprobarlos tú mismo en segundos. El trabajo difícil — detectar afirmaciones sin fuente y contradicciones internas — es donde el marco gana su valor."
```

```quiz
id: verification-method
type: multiple-choice
question: "¿Cómo debes verificar los resultados de la IA cuando no eres experto en el tema?"
options:
  - "Confiar en que la IA es generalmente precisa"
  - "Usar una estrategia de verificación: datos concretos, coherencia, estructura, propósito"
  - "Usar la IA solo para temas en los que seas experto"
answer: 1
explanation: "No necesitas conocimiento del tema para verificar de forma efectiva. Una estrategia de verificación — comprobar datos concretos, verificar coherencia, validar estructura, comprobar contra el propósito — te permite evaluar la calidad sin conocimiento profundo del dominio."
```
