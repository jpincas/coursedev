---
title: "Memoria Persistente"
duration: "15m"
tags: [memory, preferences, context, CLAUDE.md]
---

# Memoria Persistente

Cuéntale a la IA sobre ti una vez. Lo recuerda para siempre.

## El Problema

Cada nueva conversación empieza fresca. La IA no sabe:
- Quién eres
- En qué estás trabajando
- Cómo te gustan las cosas hechas
- Tu rol y contexto

Así que lo explicas. Otra vez. Y otra vez.

## La Solución

**Archivos de instrucción persistentes** — archivos markdown que se cargan automáticamente al inicio de cada sesión. Diferentes plataformas tienen diferentes nombres, pero el concepto es universal.

Escríbelo una vez. Cada sesión empieza con ese contexto ya cargado.

## El Panorama

**CLAUDE.md (Claude Code)**
Archivos markdown que se cargan automáticamente al inicio de cada sesión de Claude Code. Viven en la raíz de tu proyecto y pueden importar archivos de subdirectorios para organización.

**AGENTS.md (Cursor, Zed y otros)**
Propuesto por OpenAI en agosto de 2025, AGENTS.md ha sido adoptado por **+60.000 proyectos open-source**. El estándar para herramientas de codificación con IA.

**Claude Projects (no-codificación)**
Para trabajo no relacionado con codificación, los Claude Projects ofrecen prompts de sistema persistentes con bases de conocimiento adjuntas. Documentos, guías de estilo y plantillas persisten entre conversaciones.

**Instrucciones Personalizadas de ChatGPT**
Instrucciones persistentes a nivel de usuario. Lo que GPT debería saber sobre ti y cómo debería responder.

```callout
type: info
title: "El Lugar de Mayor Impacto para Invertir"
content: "Las instrucciones persistentes sobrescriben los prompts por conversación. Dan forma a cada interacción. Esto las convierte en el lugar de mayor impacto para invertir en calidad de prompts."
```

## Lo que Funciona: Mejores Prácticas para CLAUDE.md

**Mantener bajo 300 líneas**
La investigación muestra que los modelos de vanguardia siguen aproximadamente **150-200 instrucciones** antes de que la adherencia caiga uniformemente. Más instrucciones significa efectividad diluida.

**Usar revelación progresiva**
Importar archivos de subdirectorios para diferentes contextos. `#import .claude/coding-style.md` mantiene el archivo principal enfocado.

**Incluir lo universalmente aplicable**
- Comandos de build (`npm run build:css`, `./restart.sh`)
- Esenciales de estilo de código (formatting, convenciones de naming)
- Patrones arquitectónicos clave (propiedad de estado, reglas de renderizado)
- Reglas de flujo de trabajo (cuándo ejecutar tests, patrones de git commit)

**Commit a version control**
Todo el equipo se beneficia. Los nuevos desarrolladores hacen onboarding más rápido. Todos reciben las mismas instrucciones.

## Para Desarrolladores: la Estructura de CLAUDE.md

```markdown
# Resumen del Proyecto
[Descripción de 2-3 frases de qué hace este codebase]

## Stack Tecnológico
- Framework/lenguaje
- Librerías clave
- Herramientas de build

## Arquitectura
[Patrones clave — modelo de estado, estructura de componentes, etc.]

## Desarrollo
Build: npm run build
Test: go test ./...
Run: ./restart.sh

## Convenciones
[Estilo de código, naming, organización de archivos]
```

Eso es todo. Corto, factual, universalmente relevante.

## Para No-Desarrolladores: Claude Projects

Sube documentos que la IA debería siempre consultar:
- Guías de estilo
- Documentos de voz de marca
- Ejemplos de plantillas
- Documentación de procesos
- Contexto de la organización

**El formato de archivo importa.**
Markdown, Word y PDF usan un espacio similar de base de conocimiento. HTML usa **el doble** debido a la interpretación de etiquetas. Prefiere formatos limpios.

**Organiza por función.**
Crea proyectos enfocados por departamento o caso de uso. No eches todo en una única base de conocimiento masiva.

## Qué Poner en Memoria Persistente

**Sobre ti**
Tu rol, responsabilidades, área de experiencia.

**Tus preferencias**
- Preferencias de formateo (viñetas vs párrafos)
- Preferencias de idioma (inglés UK vs US)
- Preferencias de tono (formal vs casual)

**Tu contexto**
Proyectos actuales, información del equipo, contexto de la empresa.

**Instrucciones permanentes**
Cosas que siempre quieres que se hagan de cierta manera.

## Ejemplo: Contexto de Asistente Personal

```
# Sobre Mí
Senior Product Manager en TechCorp.
Trabajando en el equipo de app móvil.
Reporta a: VP de Producto

# Mis Preferencias
- Viñetas sobre párrafos
- Ortografía en inglés UK
- Tono conciso y directo
- Incluir puntos de acción al final de los resúmenes

# Enfoque Actual
Lanzamiento de funcionalidades de pago en Q1.
Stakeholders clave: Ingeniería, Diseño, Finanzas.
Riesgo principal: Integración con proveedor de pagos de terceros.

# Instrucciones Permanentes
Formatear siempre las fechas como DD/MM/YYYY.
Al redactar emails, mantener bajo 200 palabras.
```

## Beneficios

**Sin volver a explicar**
"Soy un product manager trabajando en..." ya es conocido.

**Resultado consistente**
Tus preferencias se aplican automáticamente.

**Contexto acumulado**
Añade información con el tiempo. El contexto crece.

**Conciencia de proyecto**
La IA sabe en qué estás trabajando sin que se lo digas cada vez.

```callout
type: tip
title: "Construir Incrementalmente"
content: "No necesitas escribir todo a la vez. Empieza con lo básico. Cada vez que te encuentres explicando algo repetidamente, añádelo a tu memoria persistente."
```

```quiz
id: persistent-memory-benefit
type: multiple-choice
question: "¿Qué problema resuelve la memoria persistente?"
options:
  - "Tener que volver a explicar quién eres y tus preferencias en cada conversación"
  - "La IA quedarse sin espacio en la ventana de contexto"
  - "La IA olvidando lo que dijo antes en la misma conversación"
answer: 0
explanation: "La memoria persistente resuelve el problema de empezar cada conversación desde cero. Tu contexto, preferencias e instrucciones permanentes se cargan automáticamente, así que no necesitas volver a explicarlas cada vez."
```

```quiz
id: persistent-memory-adoption
type: multiple-choice
question: "¿Por qué las instrucciones persistentes son el lugar de mayor impacto para invertir en calidad de prompts?"
options:
  - "Sobrescriben los prompts por conversación y dan forma a cada interacción"
  - "Requieren menos mantenimiento que los prompts regulares"
  - "Funcionan con todos los modelos de IA"
answer: 0
explanation: "Las instrucciones persistentes sobrescriben los prompts por conversación, lo que significa que afectan a cada interacción individual. Esto las hace mucho más valiosas que redactar prompts individuales — la inversión se compuesta."
```

```agent
id: persistent-memory-demo
title: "El Poder de la Memoria Persistente"
model_label: "Claude"

system: |
  Eres un asistente de programación útil. Ayudas a desarrolladores a implementar
  funcionalidades en sus proyectos. Escribe código limpio y funcional.

scratchpad:
  "CLAUDE.md": |
    # Convenciones del Proyecto
    - Go con formateo estándar gofmt
    - Todos los handlers usan pointer receivers
    - Inglés británico en todos los comentarios y documentación
    - Tests basados en tablas requeridos para todas las nuevas funciones
    - Usar clases de utilidad de Tailwind, no CSS custom
    - Los mensajes de error deberían ser amigables para el usuario, no técnicos
  "feature-request.md": |
    Añadir un endpoint /health que retorne el estado del servidor incluyendo uptime,
    versión y estado de conexión a la base de datos.

tools:
  - scratchpad_read
  - scratchpad_write

visibility:
  system_prompt: toggleable
  tool_calls: visible
  full_context: hidden

sidebar:
  width: "45%"
  start_open: true

script:
  - type: note
    text: "Observa cómo un único archivo de configuración transforma cada interacción futura con IA."

  - type: user
    content: "Añadir un endpoint de health check al servidor"

  - type: assistant
    content: "Voy a crear un endpoint de health check para ti. Aquí tienes una implementación sencilla:"

  - type: tool_call
    tool: scratchpad_write
    args:
      filename: "health.go"
      content: |
        package main

        import (
            "encoding/json"
            "net/http"
            "time"
        )

        var startTime = time.Now()

        // HealthResponse contiene información de estado del servidor
        type HealthResponse struct {
            Status   string `json:"status"`
            Uptime   string `json:"uptime"`
            Version  string `json:"version"`
            Database string `json:"database"`
        }

        func healthHandler(w http.ResponseWriter, r *http.Request) {
            uptime := time.Since(startTime).String()

            // Comprobar si la base de datos está conectada
            dbStatus := "connected"
            if err := db.Ping(); err != nil {
                dbStatus = "disconnected"
            }

            response := HealthResponse{
                Status:   "ok",
                Uptime:   uptime,
                Version:  "1.0.0",
                Database: dbStatus,
            }

            w.Header().Set("Content-Type", "application/json")
            json.NewEncoder(w).Encode(response)
        }

  - type: tool_result
    tool: scratchpad_write
    content: "Escrito en health.go (36 líneas)"

  - type: assistant
    content: "¡Listo! He creado un endpoint /health que retorna estado del servidor, uptime, versión y estado de conexión a la base de datos."

  - type: note
    text: "La salida funciona, pero no coincide con nuestras convenciones del proyecto. Observa el inglés americano ('connected'), el patrón de value receiver y la ausencia de tests. Ahora carguemos nuestras instrucciones persistentes..."

  - type: user
    content: "He configurado las convenciones de nuestro proyecto. Por favor lee el archivo CLAUDE.md y luego implementa el mismo endpoint de health."

  - type: tool_call
    tool: scratchpad_read
    args:
      filename: "CLAUDE.md"

  - type: tool_result
    tool: scratchpad_read
    content: |
      # Convenciones del Proyecto
      - Go con formateo estándar gofmt
      - Todos los handlers usan pointer receivers
      - Inglés británico en todos los comentarios y documentación
      - Tests basados en tablas requeridos para todas las nuevas funciones
      - Usar clases de utilidad de Tailwind, no CSS custom
      - Los mensajes de error deberían ser amigables para el usuario, no técnicos

  - type: tool_call
    tool: scratchpad_read
    args:
      filename: "feature-request.md"

  - type: tool_result
    tool: scratchpad_read
    content: |
      Añadir un endpoint /health que retorne el estado del servidor incluyendo uptime,
      versión y estado de conexión a la base de datos.

  - type: assistant
    content: "Perfecto, puedo ver tus convenciones del proyecto. Voy a implementar el endpoint de health siguiendo tus estándares — inglés británico, pointer receivers y con tests comprehensivos."

  - type: tool_call
    tool: scratchpad_write
    args:
      filename: "health.go"
      content: |
        package main

        import (
            "encoding/json"
            "net/http"
            "time"
        )

        var startTime = time.Now()

        // HealthResponse contiene información de estado del servidor
        type HealthResponse struct {
            Status   string `json:"status"`
            Uptime   string `json:"uptime"`
            Version  string `json:"version"`
            Database string `json:"database"`
        }

        // HealthHandler retorna el estado actual de salud del servidor
        func (m *Model) HealthHandler(w http.ResponseWriter, r *http.Request) {
            uptime := time.Since(startTime).String()

            // Comprobar si la base de datos está conectada
            dbStatus := "connected"
            if err := db.Ping(); err != nil {
                dbStatus = "disconnected"
                // Error amigable para usuario registrado, pero sin exponer detalles técnicos
            }

            response := HealthResponse{
                Status:   "ok",
                Uptime:   uptime,
                Version:  "1.0.0",
                Database: dbStatus,
            }

            w.Header().Set("Content-Type", "application/json")
            json.NewEncoder(w).Encode(response)
        }

  - type: tool_result
    tool: scratchpad_write
    content: "Escrito en health.go (39 líneas)"

  - type: tool_call
    tool: scratchpad_write
    args:
      filename: "health_test.go"
      content: |
        package main

        import (
            "net/http"
            "net/http/httptest"
            "testing"
        )

        func TestHealthHandler(t *testing.T) {
            tests := []struct {
                name           string
                dbConnected    bool
                expectedStatus int
                checkDB        bool
            }{
                {
                    name:           "servidor saludable con base de datos",
                    dbConnected:    true,
                    expectedStatus: http.StatusOK,
                    checkDB:        true,
                },
                {
                    name:           "servidor saludable sin base de datos",
                    dbConnected:    false,
                    expectedStatus: http.StatusOK,
                    checkDB:        true,
                },
            }

            for _, tt := range tests {
                t.Run(tt.name, func(t *testing.T) {
                    req := httptest.NewRequest("GET", "/health", nil)
                    w := httptest.NewRecorder()

                    model := &Model{}
                    model.HealthHandler(w, req)

                    if w.Code != tt.expectedStatus {
                        t.Errorf("expected status %d, got %d", tt.expectedStatus, w.Code)
                    }

                    if contentType := w.Header().Get("Content-Type"); contentType != "application/json" {
                        t.Errorf("expected Content-Type application/json, got %s", contentType)
                    }
                })
            }
        }

  - type: tool_result
    tool: scratchpad_write
    content: "Escrito en health_test.go (46 líneas)"

  - type: assistant
    content: "¡Implementado! He creado el endpoint de health siguiendo tus convenciones: pointer receiver en el handler, inglés británico en los comentarios ('whether' en lugar de 'if'), manejo de errores amigable para el usuario y tests comprehensivos basados en tablas cubriendo ambos estados de la base de datos."

  - type: note
    text: "Misma tarea, resultado dramáticamente diferente. El archivo CLAUDE.md asegura código consistente que sigue las convenciones en cada conversación. Escríbelo una vez, benéficiale para siempre."
```
