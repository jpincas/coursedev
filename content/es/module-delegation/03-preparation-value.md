---
title: "La preparación es tu valor"
duration: "10m"
tags: [preparación, valor, planificación]
---

# La preparación importa más ahora

Porque la ejecución es barata.

## El cambio en los cuellos de botella

**Antes de la IA:**
Cuando crear un informe tardaba 4 horas, podrías omitir la planificación adecuada y simplemente empezar a escribir. El cuello de botella era el tiempo de ejecución.

**Después de la IA:**
Cuando crear un informe tarda 5 minutos, el cuello de botella se traslada. La calidad del resultado está determinada por la calidad de la entrada — tu especificación, tu pensamiento, tu planificación.

## Dónde reside el valor ahora

![El ciclo de preparación-delegación-verificación](/content/module-delegation/images/preparation-cycle.png)

El flujo de trabajo asistido por IA es un ciclo: preparación → delegación → verificación. Pero estas tres etapas no son iguales.

**La preparación es donde aportas más valor.** Aquí es donde tu conocimiento del dominio, tu pensamiento estratégico y tu comprensión de los requisitos dan forma a lo que se construye.

**La delegación es barata.** La ejecución de la IA es casi instantánea. Esta es la parte fácil.

**La verificación es esencial.** Pero si la preparación fue exhaustiva, la verificación es rápida. Si la preparación fue débil, la verificación se convierte en trabajo de rescate.

![Cambio de valor en el trabajo asistido por IA](/content/module-delegation/images/value-shift.svg)

| Actividad | Antes de IA | Después de IA |
|----------|-------------|---------------|
| Pensar | Importante | Crítico |
| Planificar | A menudo omitido | Esencial |
| Especificación | Aproximado era suficiente | La precisión importa |
| Ejecución | Principal consumo de tiempo | Casi instantánea |
| Revisión | Pase rápido | Donde aportas valor |

## Tu nueva propuesta de valor

Tu valor se traslada de "puedo hacer esto" a "sé lo que debería hacerse."

**La investigación importa más**
Entender el espacio del problema, saber qué se ha intentado, identificar los requisitos reales.

**La planificación importa más**
Definir cómo se ve el éxito, anticipar casos extremos, establecer restricciones.

**La especificación importa más**
Articulación clara de requisitos, criterios de éxito explícitos, contexto completo.

**¿Ejecución?**
Esa es la parte barata ahora.

```callout
type: warning
title: "El riesgo"
content: "Si omites la preparación porque 'la IA lo averiguará', obtendrás una producción rápida de lo incorrecto. La velocidad sin dirección es simplemente un desperdicio eficiente."
```

## Implicaciones prácticas

**Invierte tiempo al inicio**
Cinco minutos de especificación clara superan cinco iteraciones de peticiones vagas.

**Piensa antes de delegar**
¿Qué necesitas exactamente? ¿Quién es la audiencia? ¿Qué restricciones se aplican?

**Define criterios de éxito**
¿Cómo sabrás si el resultado es bueno? ¿Qué lo haría incorrecto?

**Pero no planifiques en exceso**
La preparación importa, pero la velocidad de iteración también es una superpotencia. La investigación de McKinsey reveló que **quince iteraciones en dos días superan a dos iteraciones en cinco días**.

El punto óptimo: suficiente preparación para que el primer resultado esté en el rango correcto, luego iteración rápida para refinar. La tecnología aporta el 20% del valor; rediseñar el trabajo aporta el 80%.

Esta preparación es lo que separa a los usuarios efectivos de IA de los frustrados.

## Calidad de preparación = Calidad del resultado

La preparación real no es escribir un mejor prompt — es un proceso. Reúne tus datos, encarga una investigación, revisa los resultados y luego encarga el entregable. Observa el flujo de trabajo completo.

```agent
id: deep-research-demo
title: "Encargando una investigación real"
model_label: "Claude"

system: |
  Eres un analista de investigación comercial. Lee todos los materiales
  proporcionados con atención antes de responder. Cruza los datos internos
  con tu análisis. Sé específico y basado en evidencia.

scratchpad:
  "client-data.txt": |
    HealthBridge Wellness — Resumen del cliente
    Sector: Programas de bienestar corporativo
    Tamaño: 45 empleados, 3,2M £ de ingresos
    Fundada: 2021
    Servicios: Fitness in situ, talleres de salud mental, coaching nutricional
    Clientes: 28 cuentas corporativas
    Crecimiento: Crecimiento anual del 40% en ingresos
    Problema: Perdiendo licitaciones empresariales frente a competidores más grandes
    Tasa de éxito: 85% para PYMES, 12% para empresas (100+ empleados)
    Diferenciador: Programas personalizados, alta participación de los usuarios (92%)
    Debilidad: Sin plataforma tecnológica, informes manuales
  "lost-deals.txt": |
    Autopsias de acuerdos empresariales perdidos (últimos 12 meses)

    1. Barclays (2.400 empleados) — PERDIDO
       Motivo: "Sin integración API con nuestra plataforma de RRHH (Workday)"
       Les gustó: enfoque de participación, precios competitivos
       Se quedaron con: Virgin Pulse

    2. Deloitte UK (4.200 empleados) — PERDIDO
       Motivo: "Necesitan dashboards en tiempo real para 12 ubicaciones de oficinas"
       Les gustó: diseño del programa de salud mental
       Se quedaron con: Desarrollo interno

    3. Tesco HQ (800 empleados) — PERDIDO
       Motivo: "No pueden demostrar ROI a nivel de consejo sin plataforma de datos"
       Les gustó: personalización, regresaron para la filial PYME
       Se quedaron con: Wellhub (anteriormente Gympass)

    4. HSBC Equipo de Innovación (350 empleados) — PERDIDO
       Motivo: "Sin SSO, sin informes de cumplimiento, sin garantías de residencia de datos"
       Les gustó: resultados de la prueba piloto (96% de participación en prueba de 4 semanas)
       Se quedaron con: Decisión diferida

    5. Arup (600 empleados) — GANADO (único acuerdo empresarial ganado este año)
       Motivo: "Un defensor en RRHH eludió la adquisición normal"
       Riesgo: Sin plataforma significa que quizás no sobreviviremos a la revisión anual

tools:
  - scratchpad_read
  - scratchpad_write

visibility:
  system_prompt: hidden
  tool_calls: visible
  full_context: hidden

sidebar:
  width: "50%"
  start_open: true

script:
  - type: note
    text: "Fase 1 — Encargar la investigación. Esta persona ya ha hecho su preparación: ha extraído cinco autopsias de acuerdos empresariales del CRM y compilado los datos de su empresa. Ahora encarga a la IA que haga el trabajo de investigación."

  - type: user
    content: "Lee el perfil de nuestra empresa en client-data.txt y nuestras autopsias de acuerdos empresariales perdidos en lost-deals.txt. Necesito un análisis del panorama competitivo: para cada plataforma que nuestros prospectos eligieron en lugar de nosotros, ¿qué ofrecen y por qué ganaron? ¿Qué requisitos siguen apareciendo? Escribe tus hallazgos en research-notes.md."

  - type: tool_call
    tool: scratchpad_read
    args:
      filename: "client-data.txt"

  - type: tool_result
    tool: scratchpad_read
    content: |
      HealthBridge Wellness — Resumen del cliente
      Sector: Programas de bienestar corporativo
      Tamaño: 45 empleados, 3,2M £ de ingresos
      Fundada: 2021
      Servicios: Fitness in situ, talleres de salud mental, coaching nutricional
      Clientes: 28 cuentas corporativas
      Crecimiento: Crecimiento anual del 40% en ingresos
      Problema: Perdiendo licitaciones empresariales frente a competidores más grandes
      Tasa de éxito: 85% para PYMES, 12% para empresas (100+ empleados)
      Diferenciador: Programas personalizados, alta participación de los usuarios (92%)
      Debilidad: Sin plataforma tecnológica, informes manuales

  - type: tool_call
    tool: scratchpad_read
    args:
      filename: "lost-deals.txt"

  - type: tool_result
    tool: scratchpad_read
    content: |
      Autopsias de acuerdos empresariales perdidos (últimos 12 meses)

      1. Barclays (2.400 empleados) — PERDIDO
         Motivo: "Sin integración API con nuestra plataforma de RRHH (Workday)"
         Les gustó: enfoque de participación, precios competitivos
         Se quedaron con: Virgin Pulse

      2. Deloitte UK (4.200 empleados) — PERDIDO
         Motivo: "Necesitan dashboards en tiempo real para 12 ubicaciones de oficinas"
         Les gustó: diseño del programa de salud mental
         Se quedaron con: Desarrollo interno

      3. Tesco HQ (800 empleados) — PERDIDO
         Motivo: "No pueden demostrar ROI a nivel de consejo sin plataforma de datos"
         Les gustó: personalización, regresaron para la filial PYME
         Se quedaron con: Wellhub (anteriormente Gympass)

      4. HSBC Equipo de Innovación (350 empleados) — PERDIDO
         Motivo: "Sin SSO, sin informes de cumplimiento, sin garantías de residencia de datos"
         Les gustó: resultados de la prueba piloto (96% de participación en prueba de 4 semanas)
         Se quedaron con: Decisión diferida

      5. Arup (600 empleados) — GANADO (único acuerdo empresarial ganado este año)
         Motivo: "Un defensor en RRHH eludió la adquisición normal"
         Riesgo: Sin plataforma significa que quizás no sobreviviremos a la revisión anual

  - type: tool_call
    tool: scratchpad_write
    args:
      filename: "research-notes.md"
      content: |
        # Investigación del panorama competitivo

        ## Competidores que ganaron nuestros acuerdos

        **Virgin Pulse** (ganó a Barclays)
        - Plataforma de bienestar empresarial completa, líder de mercado establecida
        - Más de 200 integraciones con sistemas de RRHH incluyendo Workday, SAP, BambooHR
        - Dashboards de análisis en tiempo real, informes multi-sitio
        - SSO empresarial, certificado ISO 27001, conforme con GDPR
        - Por qué ganó: arquitectura API-first que se conecta al stack de RRHH existente

        **Wellhub / anteriormente Gympass** (ganó a Tesco)
        - Modelo de mercado — agrega proveedores de bienestar bajo una sola plataforma
        - Los clientes corporativos obtienen un dashboard único para todos los proveedores
        - Informes de ROI y análisis de utilización integrados
        - Nota: HealthBridge podría potencialmente unirse como proveedor en lugar de competir
        - Por qué ganó: capacidades de informes a nivel de consejo que no pudimos igualar

        **Deloitte desarrolló internamente**
        - Tenía presupuesto de ingeniería para resolverlo ellos mismos
        - Quería dashboards multi-sitio para 12 ubicaciones
        - Por qué: necesidades específicas de informes para 4.200 empleados en diferentes ubicaciones

        **HSBC diferido (sin ganador aún)**
        - Realizó una prueba piloto de 4 semanas con nosotros — 96% de participación
        - No pudo proceder sin SSO, informes de cumplimiento, residencia de datos
        - El producto funcionó; el proceso de adquisición lo bloqueó

        ## Requisitos empresariales recurrentes

        | Requisito | Acuerdos afectados | Categoría |
        |-----------|-------------------|-----------|
        | Integración con plataforma de RRHH (Workday, SAP) | Barclays, HSBC | Infraestructura |
        | Dashboards e informes en tiempo real | Deloitte, Tesco | Informes |
        | SSO / inicio de sesión único | HSBC | Seguridad |
        | Informes de cumplimiento | HSBC | Seguridad |
        | Garantías de residencia de datos | HSBC | Seguridad |
        | Demostración de ROI a nivel de consejo | Tesco | Informes |
        | Soporte multi-sitio | Deloitte | Infraestructura |

        ## Observación clave
        Todos los acuerdos se perdieron por motivos tecnológicos/de infraestructura. Ningún prospecto rechazó nuestro servicio principal — la participación, el diseño del programa y la personalización fueron elogiados en todas las autopsias. Nuestro único acuerdo empresarial ganado (Arup) ocurrió porque alguien eludió la adquisición por completo.

  - type: tool_result
    tool: scratchpad_write
    content: "Escrito en research-notes.md (42 líneas)"

  - type: assistant
    content: "Investigación escrita en research-notes.md. El hallazgo clave en los cinco acuerdos: toda pérdida empresarial fue por motivos tecnológicos y de infraestructura — integraciones, dashboards, SSO, cumplimiento. Ningún prospecto rechazó los programas principales de bienestar. Tu ganancia con Arup lo confirma — solo convirtieron porque alguien eludió la adquisición estándar."

  - type: note
    text: "Fase 2 — Revisar la investigación. Abre research-notes.md en la barra lateral para leerla. En un flujo de trabajo real, aquí es donde gastarías tiempo pensando: ¿Esto coincide con mi experiencia? ¿Qué es sorprendente? ¿Qué patrón destaca? El humano lee la investigación y detecta algo importante."

  - type: note
    text: "Fase 3 — Encargar el entregable. Armado con la investigación Y su propia reflexión al revisarla, el humano ahora encarga el documento para el consejo. Observa el prompt — el humano retroalimenta una reflexión estratégica específica que moldea todo el resultado."

  - type: user
    content: "Buena investigación. Este es el patrón que veo: tres de cuatro pérdidas son brechas de infraestructura — integraciones, SSO, cumplimiento — no calidad del producto. HSBC incluso realizó una prueba piloto exitosa y no pudo comprar. Esa es nuestra reflexión clave para el consejo.\n\nAhora escribe un análisis listo para el consejo en competitive-analysis.md. La decisión: ¿deberíamos gastar 500K £ construyendo nuestra propia plataforma, o asociarnos con una existente como CoreHealth o Wellhub? Lidera con la recomendación. Usa la evidencia de los acuerdos perdidos para respaldar el caso. Incluye un cronograma por fases. Sin lenguaje de consultor — nuestro consejo quiere especificidades."

  - type: tool_call
    tool: scratchpad_read
    args:
      filename: "research-notes.md"

  - type: tool_result
    tool: scratchpad_read
    content: |
      # Investigación del panorama competitivo

      ## Competidores que ganaron nuestros acuerdos

      **Virgin Pulse** (ganó a Barclays)
      - Plataforma de bienestar empresarial completa, líder de mercado establecida
      - Más de 200 integraciones con sistemas de RRHH incluyendo Workday, SAP, BambooHR
      - Dashboards de análisis en tiempo real, informes multi-sitio
      - SSO empresarial, certificado ISO 27001, conforme con GDPR
      - Por qué ganó: arquitectura API-first que se conecta al stack de RRHH existente

      **Wellhub / anteriormente Gympass** (ganó a Tesco)
      - Modelo de mercado — agrega proveedores de bienestar bajo una sola plataforma
      - Los clientes corporativos obtienen un dashboard único para todos los proveedores
      - Informes de ROI y análisis de utilización integrados
      - Nota: HealthBridge podría potencialmente unirse como proveedor en lugar de competir
      - Por qué ganó: capacidades de informes a nivel de consejo que no pudimos igualar

      **Deloitte desarrolló internamente**
      - Tenía presupuesto de ingeniería para resolverlo ellos mismos
      - Quería dashboards multi-sitio para 12 ubicaciones
      - Por qué: necesidades específicas de informes para 4.200 empleados en diferentes ubicaciones

      **HSBC diferido (sin ganador aún)**
      - Realizó una prueba piloto de 4 semanas con nosotros — 96% de participación
      - No pudo proceder sin SSO, informes de cumplimiento, residencia de datos
      - El producto funcionó; el proceso de adquisición lo bloqueó

      ## Requisitos empresariales recurrentes

      | Requisito | Acuerdos afectados | Categoría |
      |-----------|-------------------|-----------|
      | Integración con plataforma de RRHH (Workday, SAP) | Barclays, HSBC | Infraestructura |
      | Dashboards e informes en tiempo real | Deloitte, Tesco | Informes |
      | SSO / inicio de sesión único | HSBC | Seguridad |
      | Informes de cumplimiento | HSBC | Seguridad |
      | Garantías de residencia de datos | HSBC | Seguridad |
      | Demostración de ROI a nivel de consejo | Tesco | Informes |
      | Soporte multi-sitio | Deloitte | Infraestructura |

      ## Observación clave
      Todos los acuerdos se perdieron por motivos tecnológicos/de infraestructura. Ningún prospecto rechazó nuestro servicio principal — la participación, el diseño del programa y la personalización fueron elogiados en todas las autopsias. Nuestro único acuerdo empresarial ganado (Arup) ocurrió porque alguien eludió la adquisición por completo.

  - type: tool_call
    tool: scratchpad_write
    args:
      filename: "competitive-analysis.md"
      content: |
        # Informe para el consejo: Construir vs. Asociarse

        **Recomendación: Asociarse primero. Construir después si es necesario.**

        ## La evidencia

        Perdimos cuatro acuerdos empresariales en 12 meses. Ninguno rechazó nuestros programas — rechazaron nuestra infraestructura:

        | Acuerdo | Lo que querían | Lo que nos faltaba |
        |---------|---------------|-------------------|
        | Barclays (2.400 emp) | Integración API con Workday | Cualquier integración con sistemas de RRHH |
        | Deloitte (4.200 emp) | Dashboards multi-sitio | Plataforma de informes |
        | Tesco (800 emp) | Datos de ROI a nivel de consejo | Exportación de datos y análisis |
        | HSBC (350 emp) | SSO, cumplimiento, residencia de datos | Seguridad IT empresarial |

        HSBC es el caso más claro: realizaron una prueba piloto, obtuvieron 96% de participación y aún no pudieron comprar. Nuestro producto funciona. El proceso de adquisición de ellos necesita infraestructura que no tenemos.

        ## Por qué construir no funciona

        Una construcción de 500K £ nos lleva a dashboards e informes básicos en 9-12 meses. No nos consigue:
        - Integraciones con Workday/SAP (perdimos Barclays)
        - Certificación ISO 27001 (6+ meses por sí sola)
        - SSO y seguridad empresarial (estancó HSBC)

        Coste realista para nivel empresarial: 800K-1,2M £. Nos pasaríamos del presupuesto y aun así no resolveríamos dos de los cuatro bloqueos.

        ## Por qué la asociación funciona

        Las plataformas con marca blanca como CoreHealth ya tienen el envoltorio empresarial: integraciones de RRHH, SSO, certificaciones de cumplimiento, dashboards multi-sitio. Nosotros conectamos nuestros programas a su infraestructura.

        - Cronograma: 3-4 meses para integrar (vs. 9-12 para construir)
        - Coste: Compartición de ingresos (15-25%) en lugar de 500K £ de capital
        - Resuelve: Los cuatro puntos de ruptura de nuestros acuerdos perdidos
        - Mantenemos: Diseño del programa, metodología de participación, relaciones con clientes

        Nuestra tasa de participación del 92% nos hace atractivos para las plataformas — tenemos lo que les falta (calidad del programa que funciona). Ellos tienen lo que nos falta (infraestructura empresarial).

        ## Cronograma propuesto

        1. **Meses 1-2:** Evaluar asociaciones con CoreHealth (marca blanca) y Wellhub (mercado)
        2. **Mes 3:** Comenzar integración con el socio elegido
        3. **Meses 4-5:** Re-approach HSBC (diferido — sin competidor ganó) y filial de Tesco
        4. **Mes 6+:** Primeros ingresos empresariales. Evaluar si aún se necesita una construcción personalizada.

        ## Factores de riesgo

        - Los términos de asociación reducen márgenes (compartición de ingresos del 15-25%)
        - El cronograma de integración depende de la capacidad del socio
        - Nos volvemos parcialmente dependientes del roadmap del socio
        - Mitigación: los ingresos de la asociación financian la independencia eventual si es necesaria

  - type: tool_result
    tool: scratchpad_write
    content: "Escrito en competitive-analysis.md (48 líneas)"

  - type: assistant
    content: "Informe para el consejo escrito en competitive-analysis.md.\n\nEl caso: 500K £ en una construcción nos lleva a mitad de camino en 12 meses — dashboards pero no las integraciones de RRHH ni la infraestructura de seguridad que nos hizo perder Barclays y HSBC. La asociación nos consigue todo en 3 meses por compartición de ingresos. Nuestra tasa de participación del 92% nos hace el socio atractivo — aportamos calidad de programa que ellos no tienen.\n\nHSBC es el argumento más fuerte para el consejo: 96% de participación en la prueba piloto, aún no pudieron comprar. La brecha es de infraestructura, no de producto."

  - type: note
    text: "Tres fases, tres tipos diferentes de valor humano. Fase 1: el humano reunió datos propietarios que la IA no podía acceder (autopsias del CRM). Fase 2: el humano revisó la investigación y detectó el patrón estratégico (brechas de infraestructura vs producto). Fase 3: el humano enmarcó el entregable alrededor de esa reflexión para una audiencia específica (el consejo). La IA hizo la investigación y la escritura — pero el pensamiento que lo hizo útil fue humano."
```

```quiz
id: preparation-value
type: multiple-choice
question: "¿Por qué la preparación importa más en un flujo de trabajo asistido por IA?"
options:
  - "La IA necesita instrucciones detalladas para funcionar"
  - "La ejecución ahora es barata; la calidad de la entrada determina la calidad del resultado"
  - "La IA no puede hacer planificación por sí sola"
  - "La preparación reduce los costes de la IA"
answer: 1
explanation: "Cuando la ejecución es casi instantánea, el cuello de botella de la calidad se traslada a la calidad de la entrada. Tu pensamiento, planificación y especificación determinan si la IA produce algo valioso o basura rápida."
```
