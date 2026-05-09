---
title: "Construir Sistemas, No Solo Prompts"
duration: "10m"
tags: [advanced, systems, workflows]
---

# Cómo Trabajan los Usuarios Más Efectivos

Los usuarios de IA más efectivos no solo escriben mejores prompts. Están construyendo sistemas.

## Más Allá de los Prompts Aislad

La mayoría de la gente usa la IA así:
1. Tiene una tarea
2. Escribe un prompt
3. Obtiene un resultado
4. Repite desde cero la próxima vez

Los usuarios efectivos hacen esto:
1. Tiene una tarea
2. Comprueba si tiene una solución reutilizable
3. Si no, la crea
4. La usa ahora y para siempre

## Los Tres Pilares

**Skills (Habilidades)**
Procedimientos reutilizables que capturan cómo haces las cosas. Escribes una vez, la IA los sigue para siempre.

**Subagents (Subagentes)**
IA que genera agentes auxiliares para trabajo en paralelo. Tareas complejas divididas y abordadas simultáneamente.

**Conectores (MCP)**
Conexiones que permiten a la IA interactuar directamente con tus herramientas y sistemas.

![Three Pillars: Skills, Subagents, and MCP Connectors leading to Consistent Results](/content/module-advanced/images/three-pillars.svg)

## El Efecto Compuesto

Cada pieza de infraestructura que construyes genera dividendos:
- Una skill que documentas ahorra tiempo en cada tarea futura de ese tipo
- Preferencias que estableces persisten en todas las conversaciones futuras
- Conexiones que estableces permanecen disponibles

Esta es la diferencia entre usar IA y construir con IA.

```callout
type: info
title: "La Inversión"
content: "Configurar skills, preferencias y conexiones requiere un esfuerzo inicial. Pero esa inversión se compuesta — cada tarea futura se beneficia de la infraestructura que has construido."
```

## Economía de Tokens

Los tokens son la nueva moneda. Entender la economía importa.

**El problema del consumo:**
Los agentes consumen **100x más tokens** que un chat simple. ¿Por qué? Porque todo el historial de conversación se reenvía con cada mensaje en APIs stateless.

**La trayectoria de precios:**
El precio de los tokens ha caído desde **$20 por millón de tokens** a finales de 2022 hasta aproximadamente **$0.40 por millón** para agosto de 2025.

Pero el consumo se ha disparado. Trabajar con agentes puede significar enviar decenas de millones de tokens al día.

**Estrategias de optimización de costes:**

**Caché de prompts** — Coloca contenido estático (prompts de sistema, documentos de referencia) al inicio. El caché ahorra **60-80%** en contenido repetido.

**Cascada de modelos** — Usa modelos baratos para tareas simples y modelos premium para los complejos. Reduce costes un **30-50%**.

**Procesamiento por lotes** — La mayoría de las plataformas ofrecen un **descuento del 50%** para peticiones de API por lotes que no necesitan respuestas inmediatas.

**Fine-tuning** — Para cargas de trabajo estables de alto volumen, los modelos ajustados pueden ser más rentables que la ingeniería de prompts.

```callout
type: note
title: "La Orientación de Deloitte"
content: "Los líderes empresariales deberían tratar la economía de la IA con el mismo rigor que la energía o la asignación de capital, reconociendo los tokens como la nueva moneda."
```

Para empresas que despliegan agentes a gran escala, los costes de tokens se convierten en un partida de gastos real. La solución más barata a menudo no es la mejor solución — pero entender los compromisos importa.

```quiz
id: advanced-systems
type: multiple-choice
question: "¿Qué separa a los usuarios efectivos de IA de los usuarios casuales?"
options:
  - "Escriben prompts más largos y detallados"
  - "Construyen sistemas reutilizables en lugar de prompts aislados"
  - "Usan modelos de IA más caros"
answer: 1
explanation: "Los usuarios efectivos de IA construyen infraestructura — skills, preferencias, conexiones — que facilita cada tarea futura. Invierten en sistemas reutilizables en lugar de empezar desde cero cada vez."
```

```quiz
id: token-economics
type: multiple-choice
question: "¿Por qué los flujos de trabajo agentic consumen 100x más tokens que un chat simple?"
options:
  - "Todo el historial de conversación se reenvía con cada mensaje en APIs stateless"
  - "Los agentes usan modelos de lenguaje más complejos"
  - "Los agentes usan múltiples modelos simultáneamente"
answer: 0
explanation: "Las APIs stateless significan que todo el historial de conversación se reenvía con cada mensaje. A medida que las conversaciones crecen más largas y los agentes realizan más acciones, el consumo de tokens se dispara — 100x mayor que las interacciones simples de chat."
```
