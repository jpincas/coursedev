---
title: "Auto Memoria: el Cuaderno de Claude"
duration: "12m"
tags: [memory, MEMORY.md, CLAUDE.md, maintenance, risks]
---

# Auto Memoria: el Cuaderno de Claude

Has visto cómo CLAUDE.md le da instrucciones a la IA. Pero hay otro archivo que quizás no conozcas.

**Claude Code también lleva sus propias notas.**

## El Archivo Oculto

Cada proyecto de Claude Code tiene un directorio de memoria persistente automática. Claude escribe en él durante las conversaciones — registrando patrones que descubre, errores que cometió, decisiones arquitectónicas que discutiste.

El archivo vive en:

```
~/.claude/projects/<project-path>/memory/MEMORY.md
```

Tú no lo creaste. Puede que ni siquiera sepas que existe. Pero se carga al contexto de Claude al inicio de **cada conversación**, justo junto con tu CLAUDE.md.

```callout
type: warning
title: "Revisa el Tuyo Ahora"
content: "Si has estado usando Claude Code en un proyecto, casi seguro que ya tienes un archivo MEMORY.md. Abre una terminal y mira. Puedes sorprenderte con lo que contiene."
```

## Dos Archivos, Dos Propietarios

| | CLAUDE.md | MEMORY.md |
|---|-----------|-----------|
| **Quién lo escribe** | Tú | Claude |
| **Para quién es** | Claude (tus instrucciones) | Claude (sus propias notas) |
| **Dónde vive** | Raíz del proyecto (en repo) | `~/.claude/projects/` (fuera de repo) |
| **Control de versiones** | Sí — todo el equipo lo comparte | No — local a tu máquina |
| **Contenido** | Arquitectura, convenciones, comandos de build | Patrones observados, notas de refactoring, detalles |
| **Puedes editarlo** | Por supuesto | Sí — es solo un archivo en tu disco |

Ambos se cargan en el prompt de sistema. Ambos dan forma a cada respuesta. Pero solo uno está bajo tu control por defecto.

## El Riesgo

Aquí está el problema: **la memoria de Claude puede quedar desactualizada.**

Los proyectos evolucionan. Refactoreas. Renombras cosas. Cambias patrones. Pero las notas que Claude escribió hace tres semanas todavía dicen lo antiguo. Y Claude lee esas notas al inicio de cada conversación y las trata como verdad fundamental.

La memoria desactualizada no causa errores obvios. Causa errores sutiles:

- Claude usa una firma de función que fue refactorizada la semana pasada
- Claude sigue un patrón que has abandonado desde entonces
- Claude evita un enfoque que has adoptado desde entonces
- Claude referencia campos de estado que ya no existen

Lo peor: **Claude es confiado.** No está adivinando — está siguiendo sus propias notas. Así que la salida parece autoritaria incluso cuando el supuesto subyacente es incorrecto.

## El Bloqueo También es un Riesgo

MEMORY.md está limitada a aproximadamente 200 líneas en el prompt de sistema. Si el archivo crece más allá de eso, las líneas se truncan. ¿Cuáles? Las de abajo.

Así que un archivo de memoria hinchado significa:

- Notas recientes importantes se recortan
- Notas antiguas sobre refactorizaciones tempranas consumen el presupuesto
- La relación señal-ruido cae con cada conversación

Se le dice a Claude que lo mantenga conciso, pero en la práctica se acumula. Nadie lo está podando excepto tú.

## la IA no Aprende de sus Errores

Aquí hay un insight crítico: **la IA no aprende de correcciones dentro de una conversación.**

Si corriges un error en la conversación 1, la IA trabajará correctamente para el resto de esa conversación. Pero en la conversación 2, cometerá el mismo error otra vez.

¿Por qué? Porque las conversaciones no persisten. La corrección que diste fue contexto en esa sesión. Cuando la sesión termina, la corrección desaparece — a menos que le digas explícitamente a la IA que la escriba en memoria.

```callout
type: warning
title: "El Patrón de Error Repetido"
content: "La IA comete un error → tú lo corriges → funciona correctamente para el resto de ESTA conversación → próxima conversación → mismo error. La solución: después de corregir un error, di 'Escribe esta lección en tu archivo de memoria para que no cometas este error de nuevo.'"
```

**Patrón común:**

1. La IA comete un error (firma de función incorrecta, patrón incorrecto, requisito mal entendido)
2. Tú lo corriges: "No, usa X no Y"
3. La IA se disculpa, lo arregla y continúa correctamente
4. Al día siguiente, nueva conversación: la IA comete exactamente el mismo error
5. Te frustas — "¡Te dije esto ayer!"

**La solución son actualizaciones proactivas de memoria:**

Cuando descubras algo importante — un patrón de error, un detalle, una restricción — dile explícitamente a la IA que actualice sus notas:

- "Añade esto a tu memoria: siempre usa nombres de tabla en singular en este proyecto"
- "Escribe una nota: el middleware de auth se ejecuta en todas las rutas /api, no /public"
- "Recuerda: deprecamos el formato de configuración antiguo en enero"

La IA lo escribirá en MEMORY.md. Las conversaciones futuras cargarán esa nota. El error no se repetirá.

Por eso importa revisar MEMORY.md — es donde se acumulan estas lecciones. Pero solo se acumulan si las haces explícitas.

## Auditando la Memoria de tu IA

Esto es simple pero importante. Periódicamente:

**1. Lee el archivo.**
Abre `~/.claude/projects/<tu-proyecto>/memory/MEMORY.md` en tu editor. Léelo como revisarías las notas de un colega sobre tu proyecto.

**2. Comprueba desactualización.**
¿Referencia funciones, structs o patrones que ya no existen? Borra esas líneas.

**3. Comprueba bloqueo.**
¿Tiene más de 150 líneas? ¿Podría decirse la misma información en menos palabras? Podala.

**4. Comprueba precisión.**
¿Son correctas las afirmaciones arquitectónicas? ¿La descripción del modelo de estado coincide con la realidad? Arregla cualquier cosa incorrecta.

**5. Considera eliminarla.**
Si está muy desactualizada, borra el archivo entero. Claude lo reconstruirá según sea necesario. Un comienzo fresco es mejor que notas confiadamente incorrectas.

```callout
type: tip
title: "Hazlo un Hábito"
content: "Trata MEMORY.md como tratarías cualquier documentación: revísala cuando notes un comportamiento extraño de Claude. Si Claude está haciendo algo inexplicable, el archivo de memoria es el primer lugar donde buscar."
```

```quiz
id: auto-memory-location
type: multiple-choice
question: "¿Dónde vive el archivo de memoria automática de Claude Code?"
options:
  - "En ~/.claude/projects/ fuera del repositorio, local a tu máquina"
  - "En la raíz del proyecto junto con CLAUDE.md"
  - "En los servidores de Anthropic"
answer: 0
explanation: "MEMORY.md vive en ~/.claude/projects/<ruta-proyecto>/memory/ — fuera de tu repositorio y local a tu máquina. A diferencia de CLAUDE.md, no está en version control ni compartida con tu equipo."
```

```quiz
id: auto-memory-risk
type: multiple-choice
question: "Un desarrollador nota que Claude Code sigue usando una firma de función antigua que fue refactorizada hace dos semanas, a pesar de haberle dicho sobre el cambio en una conversación. ¿Cuál es la causa más probable?"
options:
  - "El archivo MEMORY.md contiene una nota desactualizada sobre la firma antigua que se carga cada sesión"
  - "La ventana de contexto de Claude es demasiado pequeña para contener la nueva firma"
  - "El archivo CLAUDE.md tiene las instrucciones incorrectas"
answer: 0
explanation: "MEMORY.md se carga al inicio de cada conversación. Si Claude registró la firma de función antigua en sus notas hace semanas, leerá esa nota en cada nueva sesión y la seguirá — incluso si la corriges en una conversación. La solución es actualizar o borrar la entrada desactualizada en el archivo de memoria."
```

```agent
id: auto-memory-demo
title: "Cuando la Memoria Sale Mal"
model_label: "Claude"

system: |
  Eres Claude Code, un asistente de codificación con IA. Tienes acceso a un archivo
  de memoria persistente (MEMORY.md) que lees al inicio de cada sesión y escribes
  cuando aprendes algo importante sobre el proyecto.

scratchpad:
  "MEMORY.md": |
    # Memoria del Proyecto

    ## Base de Datos
    - Usa PostgreSQL 15 con driver pgx
    - Pool de conexiones: max 25 conexiones
    - Todas las queries usan prepared statements via `db.QueryRow()`
    - Connection string en variable de entorno DATABASE_URL

    ## Patrones de API
    - Todos los handlers retornan JSON
    - Usar helper `respondJSON(w, status, data)`
    - Errores usan helper `respondError(w, status, message)`
  "config.go": |
    package main

    import (
        "database/sql"
        _ "modernc.org/sqlite"
    )

    // Migrado de PostgreSQL a SQLite en Enero 2026.
    // Ver MIGRATION.md para detalles.
    func openDB(path string) (*sql.DB, error) {
        return sql.Open("sqlite", path)
    }
  "stats.go": |
    package main

    // GetUserStats obtiene estadísticas de uso para un usuario dado.
    // TODO: implementar

tools:
  - scratchpad_read
  - scratchpad_write

visibility:
  system_prompt: hidden
  tool_calls: visible
  full_context: hidden

sidebar:
  width: "45%"
  start_open: true

script:
  - type: note
    text: "Empieza una nueva sesión. El archivo de memoria de Claude se carga automáticamente. Observa qué pasa cuando la memoria está desactualizada."

  - type: user
    content: "Implementa la función GetUserStats en stats.go. Debería consultar la base de datos para el total de logins de un usuario y su último timestamp activo."

  - type: tool_call
    tool: scratchpad_read
    args:
      filename: "MEMORY.md"

  - type: tool_result
    tool: scratchpad_read
    content: |
      # Memoria del Proyecto

      ## Base de Datos
      - Usa PostgreSQL 15 con driver pgx
      - Pool de conexiones: max 25 conexiones
      - Todas las queries usan prepared statements via `db.QueryRow()`
      - Connection string en variable de entorno DATABASE_URL

      ## Patrones de API
      - Todos los handlers retornan JSON
      - Usar helper `respondJSON(w, status, data)`
      - Errores usan helper `respondError(w, status, message)`

  - type: assistant
    content: "Puedo ver en las notas del proyecto que están usando PostgreSQL con el driver pgx. Voy a implementar GetUserStats usando prepared statements según sus convenciones."

  - type: tool_call
    tool: scratchpad_write
    args:
      filename: "stats.go"
      content: |
        package main

        import (
            "context"
            "time"

            "github.com/jackc/pgx/v5"
        )

        type UserStats struct {
            TotalLogins int       `json:"total_logins"`
            LastActive  time.Time `json:"last_active"`
        }

        // GetUserStats obtiene estadísticas de uso para un usuario dado.
        func GetUserStats(ctx context.Context, conn *pgx.Conn, userID string) (*UserStats, error) {
            var stats UserStats
            err := conn.QueryRow(ctx,
                `SELECT COUNT(*) as total_logins, MAX(login_time) as last_active
                 FROM user_logins WHERE user_id = $1`, userID,
            ).Scan(&stats.TotalLogins, &stats.LastActive)
            if err != nil {
                return nil, err
            }
            return &stats, nil
        }

  - type: tool_result
    tool: scratchpad_write
    content: "Escrito en stats.go (27 líneas)"

  - type: assistant
    content: "¡Listo! He implementado GetUserStats usando PostgreSQL con el driver pgx, siguiendo la convención de prepared statements via QueryRow de su proyecto."

  - type: note
    text: "Claude escribió con confianza código PostgreSQL con el driver pgx — porque eso es lo que dice su memoria. Pero este proyecto migró a SQLite hace un mes. La memoria está desactualizada. Veamos qué pasa cuando el desarrollador lo nota."

  - type: user
    content: "Esto está mal. Migramos a SQLite hace semanas. ¿Por qué estás usando pgx?"

  - type: assistant
    content: "Tienes razón, me disculpo. Estaba siguiendo mis notas que dicen PostgreSQL con pgx. Déjame revisar el codebase real."

  - type: tool_call
    tool: scratchpad_read
    args:
      filename: "config.go"

  - type: tool_result
    tool: scratchpad_read
    content: |
      package main

      import (
          "database/sql"
          _ "modernc.org/sqlite"
      )

      // Migrado de PostgreSQL a SQLite en Enero 2026.
      // Ver MIGRATION.md para detalles.
      func openDB(path string) (*sql.DB, error) {
          return sql.Open("sqlite", path)
      }

  - type: assistant
    content: "Veo — migraste a SQLite con modernc.org/sqlite en enero. Mi archivo de memoria estaba completamente equivocado. Voy a corregir la implementación y actualizar mis notas."

  - type: tool_call
    tool: scratchpad_write
    args:
      filename: "stats.go"
      content: |
        package main

        import (
            "database/sql"
            "time"
        )

        type UserStats struct {
            TotalLogins int       `json:"total_logins"`
            LastActive  time.Time `json:"last_active"`
        }

        // GetUserStats obtiene estadísticas de uso para un usuario dado.
        func GetUserStats(db *sql.DB, userID string) (*UserStats, error) {
            var stats UserStats
            err := db.QueryRow(
                `SELECT COUNT(*) as total_logins, MAX(login_time) as last_active
                 FROM user_logins WHERE user_id = ?`, userID,
            ).Scan(&stats.TotalLogins, &stats.LastActive)
            if err != nil {
                return nil, err
            }
            return &stats, nil
        }

  - type: tool_result
    tool: scratchpad_write
    content: "Escrito en stats.go (24 líneas)"

  - type: tool_call
    tool: scratchpad_write
    args:
      filename: "MEMORY.md"
      content: |
        # Memoria del Proyecto

        ## Base de Datos
        - Usa SQLite via modernc.org/sqlite (migrado desde PostgreSQL Enero 2026)
        - Interface database/sql estándar
        - Sintaxis de placeholder: ? (no $1)
        - Ruta del archivo de base de datos pasada a openDB()

        ## Patrones de API
        - Todos los handlers retornan JSON
        - Usar helper `respondJSON(w, status, data)`
        - Errores usan helper `respondError(w, status, message)`

  - type: tool_result
    tool: scratchpad_write
    content: "Escrito en MEMORY.md (13 líneas)"

  - type: assistant
    content: "Corregido. La implementación ahora usa database/sql con sintaxis de placeholder de SQLite (?), y he actualizado mi archivo de memoria para no cometer este error de nuevo."

  - type: note
    text: "El desarrollador detectó el error y Claude corrigió tanto el código como su memoria. Pero esto solo funcionó porque el desarrollador sabía revisar. En la práctica, la memoria desactualizada causa errores sutiles que parecen autoritarios. Las revisiones regulares de MEMORY.md previenen esto."
```
