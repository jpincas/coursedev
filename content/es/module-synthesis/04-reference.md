---
title: "Referencia Rápida"
duration: "5m"
tags: [reference, glossary, summary]
---

# Referencia Rápida

Un resumen de conceptos clave para consulta fácil.

## Glosario

**LLM (Modelo de Lenguaje Grande)**
La tecnología de IA detrás de Claude, GPT y sistemas similares. Fundamentalmente un motor de predicción.

**Ventana de contexto**
Todo lo que el modelo puede "ver" al generar una respuesta. Incluye el sistema prompt, historial de conversación, archivos y tu mensaje.

**Token**
Una pieza de texto que el modelo procesa. Aproximadamente 3/4 de una palabra, o alrededor de 4 caracteres.

**Sistema prompt**
Instrucciones ocultas que moldean el comportamiento de la IA. Establecido por la aplicación, a menudo invisible para los usuarios.

**Alucinación**
Cuando la IA genera información plausible pero falsa. Intrínseco a cómo funciona la predicción.

**Aterrizaje (Grounding)**
Proporcionar documentos y datos específicos para que la IA los consulte. Reduce la alucinación, mejora la precisión.

**Skill**
Un procedimiento documentado que la IA puede seguir automáticamente. Instrucciones, plantillas, ejemplos agrupados juntos.

**Subagente**
Un asistente de IA generado por el agente principal para trabajo en paralelo.

**Memoria persistente**
Información sobre ti que la IA carga automáticamente en cada sesión. Preferencias, contexto, instrucciones vigentes.

**MCP (Protocolo de Contexto del Modelo)**
Estándar para conectar la IA a herramientas y sistemas externos.

## Las Tres Preguntas

Antes de cualquier tarea:

1. **¿Qué es "hecho"?**
   Define el estado final claramente.

2. **¿Qué contexto necesita?**
   ¿Qué información se requiere para el éxito?

3. **¿Cuáles son los límites?**
   ¿Qué restricciones o límites aplican?

## Estructura Básica del Prompt

**Rol** — ¿Quién debería ser la IA?
**Tarea** — ¿Qué debería hacer?
**Contexto** — ¿Qué trasfondo necesita?
**Formato** — ¿Cómo debería estructurarse la salida?
**Restricciones** — ¿Qué límites aplican?

## Lista de Verificación

- [ ] Contrastar 3-5 hechos específicos
- [ ] Comprobar coherencia interna
- [ ] Validar que la estructura coincide con los requisitos
- [ ] Probar contra el propósito previsto

## No Delegar

**Juicio** — Decisiones éticas, compensaciones de valor
**Relaciones** — Conversaciones sensibles, construcción de confianza
**Responsabilidad** — Aprobación final, firma de conformidad

## Los Cinco Temas

1. **Resultados sobre proceso** — Define hecho, no cómo
2. **El contexto es todo** — Entrada rica = salida rica
3. **IA primero, humano verifica** — Delegar, luego verificar
4. **La preparación es la nueva ejecución** — Pensar es valor
5. **Archivos, no chat** — Entregables, no conversaciones

## El Flujo de Trabajo

```
PREPARAR → DELEGAR → VERIFICAR → ENTREGAR
   ↑______________|
      (iterar)
```

## Comparación de Modelos: Cuándo Usar Cada Herramienta

**Claude (Anthropic)**
- Fortalezas: Contexto largo (hasta 1M tokens), seguimiento de instrucciones, programación, seguridad, salida estructurada
- Mejor para: Tareas multi-archivo complejas, análisis de codebase, escritura técnica, tareas que requieren contexto extendido

**GPT (OpenAI)**
- Fortalezas: Capacidades multimodales, integraciones del ecosistema, plugins, amplia base de conocimiento
- Mejor para: Tareas que mezclan texto/imágenes/audio, iteraciones rápidas, consultas de conocimiento general, flujos de trabajo con muchas integraciones

**Gemini (Google)**
- Fortalezas: Integración del ecosistema Google, ventana de contexto de 2M tokens, comprensión multimodal
- Mejor para: Tareas que involucran Google Workspace, documentos extremadamente largos, análisis multimodal

**Esto no es una comparación de productos.** Cada herramienta tiene fortalezas específicas. La pregunta no es "cuál es mejor" — es "cuál se ajusta a esta tarea". Usa la herramienta correcta para el trabajo.

```quiz
id: final-quiz
type: multiple-choice
question: "¿Cuál es la naturaleza fundamental de los LLMs que explica tanto sus capacidades como sus limitaciones?"
options:
  - "Son bases de datos de conocimiento que almacenan y recuperan hechos"
  - "Son motores de predicción que generan tokens probables siguientes"
  - "Son sistemas de razonamiento que resuelven problemas lógicamente"
answer: 1
explanation: "Los LLMs son motores de predicción — predicen qué texto sigue basado en patrones aprendidos de los datos de entrenamiento. Esto explica sus fortalezas (continuidad de patrones excelente, generación de lenguaje natural) y limitaciones (alucinaciones, sin acceso a verdad fundamental, sin conocimiento en tiempo real). Todo lo que has aprendido en este curso se basa en esta comprensión fundamental."
```

---

Ahora estás equipado para trabajar con IA como un socio capaz, no solo como un chatbot.

Ahora ve a construir algo.
