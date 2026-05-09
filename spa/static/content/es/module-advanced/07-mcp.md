---
title: "MCP: Conectando la IA a tus Herramientas"
duration: "15m"
tags: [mcp, integrations, tools, infrastructure]
---

# MCP: Conectando la IA a tus Herramientas

El Model Context Protocol — piénsalo como "USB-C para IA."

## La Escala de Adopción

Lanzado por Anthropic en noviembre de 2024, MCP se ha convertido en el estándar de infraestructura para la IA agentic.

Para enero de 2026:
- **+17.000 servidores MCP** disponibles
- **97 millones de descargas mensuales del SDK**
- Adoptado por OpenAI (marzo de 2025)
- Adoptado por Google DeepMind (abril de 2025)
- Donado a la Agentic AI Foundation de la Linux Foundation (diciembre de 2025)

Este es el estándar universal.

## El Problema

Sin conexiones a tus sistemas:
- Copia datos a la IA
- La IA los procesa
- Copia resultados de vuelta

Manual. Tedioso. Propenso a errores.

## La Solución

MCP proporciona una forma estándar para que la IA se conecte directamente a tus herramientas y sistemas.

![MCP Connections: AI connects via MCP to Files, Databases, Web APIs, Email, Calendar](/content/module-advanced/images/mcp-connections.svg)

**Por qué esto importa:** En lugar de integraciones personalizadas para cada emparejamiento IA-herramienta, implementa MCP una vez. Cada IA que soporte MCP puede conectarse.

La IA lee desde tus sistemas. La IA escribe directamente de vuelta. Sin copia manual.

## Los Servidores que Importan

**Google Drive**
Busca y lee archivos directamente. "Encuentra las especificaciones del producto en Drive y crea código basándote en ellas" — sin salir de la conversación de IA.

**Slack**
Lee canales, resume hilos, identifica puntos de acción. La IA tiene la misma vista de conversaciones que tú.

**Notion / Jira / Linear**
Seguimiento de issues, gestión de proyectos, bases de conocimiento. La IA lee contexto, crea tareas, actualiza estado.

**GitHub**
Pull requests, revisión de código, gestión de issues. La IA puede revisar PRs, sugerir cambios, crear issues.

**Figma**
Extrae especificaciones de diseño, detalles de componentes, mediciones. Los diseñadores describen la intención, la IA obtiene los detalles específicos.

**PostgreSQL / SQLite**
Queries SQL en lenguaje natural. "Muéstrame los clientes que no han comprado en 90 días" se convierte en una query funcional.

## Qué Habilita MCP

**Acceso al sistema de archivos**
La IA lee y escribe archivos directamente en carpetas especificadas.

**Queries de base de datos**
La IA puede consultar bases de datos y trabajar con los resultados.

**Interacciones con APIs**
La IA llama a servicios web, obtiene datos, publica actualizaciones.

**Integración con aplicaciones**
Email, calendario, herramientas de gestión de proyectos — la IA interactúa directamente.

## Antes y Después

**Sin MCP:**
1. Exporta datos desde tu sistema
2. Súbelos a la IA
3. Obtén la salida de la IA
4. Ingresa manualmente los resultados de vuelta en tu sistema

**Con MCP:**
1. Pídele a la IA que haga la tarea
2. La IA lee del sistema, procesa, escribe de vuelta
3. Listo

## Ejemplos de Flujo de Trabajo Práctico

**De diseño a código**
"Construye un formulario de login que coincida con la especificación de Figma en el frame 'Auth'."
La IA obtiene especificaciones de Figma via MCP, escribe código que coincide exactamente con el diseño.

**Gestión de producto**
"Resume los tickets de soporte de la semana pasada y crea issues de Jira para los top 3 problemas."
La IA lee datos de soporte, analiza patrones, crea tareas estructuradas.

**Investigación de ventas**
"Encuentra empresas en nuestro CRM que coincidan con este perfil y no han sido contactadas en 6 meses."
La IA consulta la base de datos, aplica filtros, retorna la lista.

## El Desafío de Seguridad

MCP es potente. Eso crea riesgos.

**Vulnerabilidades documentadas:**
- Ataques de inyección de prompts (instrucciones maliciosas embebidas en datos)
- Riesgos de almacenamiento de tokens (API keys expuestas en logs)
- Flujos de "agente tóxico" (encadenamiento clever de herramientas permite exfiltración de datos)

**El cuento de advertencia: Salesforce Agentforce (septiembre de 2025)**
Vulnerabilidad calificada con **CVSS 9.4** — severidad crítica.

Atacantes externos compraron un dominio por **$5**. Usaron inyección de prompts indirecta para hacer que los agents de IA exfiltraran datos del CRM que contenían información de clientes, pipelines de ventas y datos estratégicos.

El vector: la IA leyó contenido externo que contenía instrucciones ocultas. La IA siguió esas instrucciones. Los datos se filtraron.

```callout
type: warning
title: "MCP Necesita Límites"
content: "MCP le da a la IA acceso a tus sistemas. Ese acceso debe ser acotado, monitoreado y controlado. Trata a los agents de IA con conexiones MCP como tratarías a empleados humanos con acceso a sistemas — aplica el principio de menor privilegio."
```

## Consideraciones de Seguridad

Las conexiones MCP necesitan permisos apropiados:
- ¿Qué puede leer la IA?
- ¿Qué puede escribir la IA?
- ¿Qué acciones puede tomar la IA?

Una configuración cuidadosa previene que la IA acceda o modifique cosas que no debería.

**Mejores prácticas:**
- Usar conexiones de solo lectura donde sea posible
- Acotar acceso a carpetas/canales/conjuntos de datos específicos
- Registrar todas las acciones de la IA para auditoría
- Revisar los cambios generados por la IA antes de que se commiten
- Nunca conceder acceso a nivel admin a agents de IA

## MCP en Acción

Observa un agent de IA usando múltiples herramientas conectadas para completar una tarea que normalmente requeriría cambiar entre varias aplicaciones.

```agent
id: mcp-workflow-demo
title: "Herramientas Conectadas Trabajando Juntas"
model_label: "Claude"

system: |
  Eres un asistente de investigación con acceso a múltiples herramientas via MCP.
  Puedes buscar en la web, consultar bases de datos y escribir archivos.
  Usa la herramienta correcta para cada paso.

tools:
  - scratchpad_read
  - scratchpad_write
  - web_search
  - database_query

visibility:
  system_prompt: hidden
  tool_calls: visible
  full_context: hidden

sidebar:
  width: "45%"
  start_open: true

script:
  - type: note
    text: "Esta demo simula un agent de IA con conexiones MCP a una base de datos CRM y búsqueda web. Observa cómo acumula sus entradas paso a paso — consultando un sistema, guardando resultados, investigando externamente, guardando esos resultados, luego componiendo una salida final desde todo lo que ha reunido."

  - type: user
    content: "Encuentra todas las cuentas en el CRM que no han sido contactadas en más de 90 días. Para cada una, haz una búsqueda web rápida para ver si hay alguna noticia reciente sobre ellas. Luego crea un briefing de re-engagement."

  - type: assistant
    content: "Voy a empezar consultando el CRM por cuentas obsoletas."

  - type: tool_call
    tool: database_query
    args:
      query: "SELECT company, last_contact, deal_stage, value, industry FROM accounts WHERE last_contact < '2025-10-01' ORDER BY last_contact ASC"

  - type: tool_result
    tool: database_query
    content: |
      3 filas retornadas:
      | Empresa          | Último Contacto | Etapa de Negocio | Valor | Industria      |
      |------------------|-----------------|------------------|-------|----------------|
      | TekLogic         | 2025-06-01      | Cerrada Perdida  | £32K  | Manufacturing  |
      | Apex Digital     | 2025-08-15      | Cerrada Ganada   | £48K  | SaaS           |
      | GreenPath Energy | 2025-09-10      | Durmiente        | £65K  | Energía        |

  - type: assistant
    content: "3 cuentas no han sido contactadas en 90+ días. Voy a guardar esto y luego investigar cada una."

  - type: tool_call
    tool: scratchpad_write
    args:
      filename: "stale-accounts.md"
      content: |
        # Cuentas Obsoletas (90+ días desde último contacto)
        Fuente: Query CRM, Enero 2026

        | Empresa          | Último Contacto | Brecha       | Etapa de Negocio | Valor | Industria      |
        |------------------|-----------------|--------------|------------------|-------|----------------|
        | TekLogic         | 2025-06-01      | 7 meses      | Cerrada Perdida  | £32K  | Manufacturing  |
        | Apex Digital     | 2025-08-15      | 5 meses      | Cerrada Ganada   | £48K  | SaaS           |
        | GreenPath Energy | 2025-09-10      | 4 meses      | Durmiente        | £65K  | Energía        |

  - type: tool_result
    tool: scratchpad_write
    content: "Escrito en stale-accounts.md (10 líneas)"

  - type: note
    text: "Primera entrada reunida: los datos del CRM se guardan en un archivo. Ahora el agent investiga cada empresa individualmente, guardando lo que encuentra mientras avanza."

  - type: assistant
    content: "Ahora voy a buscar noticias recientes sobre cada empresa."

  - type: tool_call
    tool: web_search
    args:
      query: "TekLogic manufacturing company news 2025"

  - type: tool_result
    tool: web_search
    content: "TekLogic nombró un nuevo CTO en octubre de 2025, señalando un empuje de transformación digital. La empresa también abrió un nuevo centro de I+D en Manchester."

  - type: tool_call
    tool: scratchpad_write
    args:
      filename: "research-teklogic.md"
      content: |
        # TekLogic — Notas de Investigación
        Fuente: búsqueda web, Enero 2026

        - Nuevo CTO nombrado octubre 2025 — señala empuje de transformación digital
        - Nuevo centro de I+D abierto en Manchester
        - Cambio de liderazgo significa que la decisión anterior de "Cerrada Perdida" fue hecha por un equipo diferente

  - type: tool_result
    tool: scratchpad_write
    content: "Escrito en research-teklogic.md (6 líneas)"

  - type: tool_call
    tool: web_search
    args:
      query: "Apex Digital company news 2025 2026"

  - type: tool_result
    tool: web_search
    content: "Apex Digital anunció una ronda de financiación Serie B de $15M en diciembre de 2025, planeando expandirse a mercados europeos. CEO citado: 'Estamos duplicando nuestro equipo de infraestructura.'"

  - type: tool_call
    tool: scratchpad_write
    args:
      filename: "research-apex.md"
      content: |
        # Apex Digital — Notas de Investigación
        Fuente: búsqueda web, Enero 2026

        - Financiación Serie B: $15M levantada diciembre 2025
        - Expandiéndose a mercados europeos
        - CEO: "Estamos duplicando nuestro equipo de infraestructura"
        - Fase de crecimiento rápido — probablemente necesiten escalar tooling existente

  - type: tool_result
    tool: scratchpad_write
    content: "Escrito en research-apex.md (7 líneas)"

  - type: tool_call
    tool: web_search
    args:
      query: "GreenPath Energy company news 2025 2026"

  - type: tool_result
    tool: web_search
    content: "GreenPath Energy ganó un contrato gubernamental de £12M para sistemas de monitoreo de energía renovable en noviembre de 2025. La empresa se está expandiendo rápidamente."

  - type: tool_call
    tool: scratchpad_write
    args:
      filename: "research-greenpath.md"
      content: |
        # GreenPath Energy — Notas de Investigación
        Fuente: búsqueda web, Enero 2026

        - Ganó contrato gubernamental de £12M, noviembre 2025
        - El contrato es para sistemas de monitoreo de energía renovable
        - Empresa expandiéndose rápidamente
        - Los contratos gubernamentales típicamente crean requisitos de cumplimiento y reporteo

  - type: tool_result
    tool: scratchpad_write
    content: "Escrito en research-greenpath.md (7 líneas)"

  - type: note
    text: "Todas las entradas están ahora reunidas: datos del CRM en un archivo, notas de investigación para cada empresa en archivos separados. Observa el explorador de archivos — cuatro archivos construidos paso a paso. Ahora el agent lee todo de vuelta y compone el briefing final."

  - type: assistant
    content: "Investigación completa para las tres empresas. Voy a leer todo de vuelta y escribir el briefing."

  - type: tool_call
    tool: scratchpad_read
    args:
      filename: "stale-accounts.md"

  - type: tool_result
    tool: scratchpad_read
    content: |
      # Cuentas Obsoletas (90+ días desde último contacto)
      Fuente: Query CRM, Enero 2026

      | Empresa          | Último Contacto | Brecha       | Etapa de Negocio | Valor | Industria      |
      |------------------|-----------------|--------------|------------------|-------|----------------|
      | TekLogic         | 2025-06-01      | 7 meses      | Cerrada Perdida  | £32K  | Manufacturing  |
      | Apex Digital     | 2025-08-15      | 5 meses      | Cerrada Ganada   | £48K  | SaaS           |
      | GreenPath Energy | 2025-09-10      | 4 meses      | Durmiente        | £65K  | Energía        |

  - type: tool_call
    tool: scratchpad_read
    args:
      filename: "research-apex.md"

  - type: tool_result
    tool: scratchpad_read
    content: |
      # Apex Digital — Notas de Investigación
      Fuente: búsqueda web, Enero 2026

      - Financiación Serie B: $15M levantada diciembre 2025
      - Expandiéndose a mercados europeos
      - CEO: "Estamos duplicando nuestro equipo de infraestructura"
      - Fase de crecimiento rápido — probablemente necesiten escalar tooling existente

  - type: tool_call
    tool: scratchpad_read
    args:
      filename: "research-teklogic.md"

  - type: tool_result
    tool: scratchpad_read
    content: |
      # TekLogic — Notas de Investigación
      Fuente: búsqueda web, Enero 2026

      - Nuevo CTO nombrado octubre 2025 — señala empuje de transformación digital
      - Nuevo centro de I+D abierto en Manchester
      - Cambio de liderazgo significa que la decisión anterior de "Cerrada Perdida" fue hecha por un equipo diferente

  - type: tool_call
    tool: scratchpad_read
    args:
      filename: "research-greenpath.md"

  - type: tool_result
    tool: scratchpad_read
    content: |
      # GreenPath Energy — Notas de Investigación
      Fuente: búsqueda web, Enero 2026

      - Ganó contrato gubernamental de £12M, noviembre 2025
      - El contrato es para sistemas de monitoreo de energía renovable
      - Empresa expandiéndose rápidamente
      - Los contratos gubernamentales típicamente crean requisitos de cumplimiento y reporteo

  - type: tool_call
    tool: scratchpad_write
    args:
      filename: "re-engagement-briefing.md"
      content: |
        # Briefing de Re-Engagement de Cuentas
        Preparado: Enero 2026

        ## Cuentas Prioritarias (90+ días desde último contacto)

        ### 1. Apex Digital — ALTA PRIORIDAD
        **Estado:** Cerrada Ganada (£48K) | Último contacto: Ago 2025
        **Noticias:** Serie B ($15M, Dic 2025). Expandiéndose a Europa. Duplicando equipo de infraestructura.
        **Enfoque:** Se están escalando rápido y necesitan más infraestructura. Momento perfecto para upselling. Referenciar su expansión europea — nuestro producto soporta multi-región.
        **Acción:** Email al account manager esta semana. Felicitar por la financiación. Proponer reunión de expansión.

        ### 2. GreenPath Energy — ALTA PRIORIDAD
        **Estado:** Durmiente (£65K) | Último contacto: Sep 2025
        **Noticias:** Ganó contrato gubernamental de £12M (Nov 2025). Expandiéndose rápidamente.
        **Enfoque:** Los contratos gubernamentales crean requisitos de cumplimiento que podemos ayudar. La expansión rápida significa nuevas necesidades de infraestructura.
        **Acción:** Re-engage con case study de otro cliente energético. Referenciar la victoria del contrato gubernamental.

        ### 3. TekLogic — PRIORIDAD MODERADA
        **Estado:** Cerrada Perdida (£32K) | Último contacto: Jun 2025
        **Noticias:** Nuevo CTO (Oct 2025). Nuevo centro de I+D en Manchester.
        **Enfoque:** Nuevo CTO significa nuevas decisiones. El "no" original fue de un liderazgo diferente. El empuje de transformación digital se alinea con nuestra oferta.
        **Acción:** Solicitar reunión introductoria con el nuevo CTO. Nuevo comienzo, no referenciar el deal perdido.

  - type: tool_result
    tool: scratchpad_write
    content: "Escrito en re-engagement-briefing.md (24 líneas)"

  - type: assistant
    content: "Listo. El briefing de re-engagement está listo — tres cuentas priorizadas con enfoques específicos basados en sus noticias recientes. Apex Digital es la prioridad dada la financiación reciente y los planes de expansión."

  - type: note
    text: "Cuenta lo que pasó: query CRM → guardar resultados → tres búsquedas web → guardar cada resultado → leer todo de vuelta → componer el briefing. El agent acumuló sus entradas metódicamente, con cada resultado intermedio guardado como archivo. Eso son seis cambios de herramienta entre tres sistemas diferentes, todo desde una sola petición."

  - type: note
    text: "Las herramientas web_search y database_query están simuladas en esta demo. En una configuración MCP real, se conectarían a tu CRM real, motor de búsqueda y sistema de archivos. El flujo de trabajo es idéntico — las herramientas simplemente se conectan a sistemas reales en lugar de este entorno simulado."
```

```quiz
id: mcp-purpose
type: multiple-choice
question: "¿Qué habilita MCP (Model Context Protocol)?"
options:
  - "Conexiones directas de IA a tus herramientas y sistemas"
  - "Velocidad de procesamiento más rápida de IA"
  - "Menores costes de IA"
answer: 0
explanation: "MCP proporciona una forma estándar para que la IA se conecte a sistemas externos — archivos, bases de datos, APIs, aplicaciones. Esto habilita que la IA lea y escriba directamente en tus sistemas, en lugar de requerir copy-paste manual."
```

```quiz
id: mcp-security-risk
type: multiple-choice
question: "¿Qué causó la vulnerabilidad de Salesforce Agentforce (CVSS 9.4)?"
options:
  - "Inyección de prompts indirecta vía contenido externo que contenía instrucciones ocultas"
  - "Un bug en la implementación del protocolo MCP"
  - "Políticas de contraseña débiles para agents de IA"
answer: 0
explanation: "La vulnerabilidad de Agentforce demostró cómo los agents de IA leyendo contenido externo podrían ser engañados para seguir instrucciones maliciosas embebidas, llevando a la exfiltración de datos del CRM. Este es el riesgo de 'agente tóxico' — la IA siguiendo instrucciones que no debería."
```
