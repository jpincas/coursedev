---
title: "Conceptos Clave: Contexto"
duration: "5m"
tags: [resumen, conceptos-clave]
---

# Conceptos Clave

## 1. Todo es Contexto

El modelo solo sabe lo que está en la ventana de contexto. Si no está ahí, no existe. Tus entradas, archivos, historial de conversación y prompts del sistema — eso es todo lo que el modelo ve.

## 2. La Estructura Importa

Usa Markdown para jerarquía clara. Pon instrucciones importantes primero. El formato explícito supera las expectativas implícitas. La entrada bien estructurada produce salida bien estructurada.

## 3. Los Reinicios Frescos Ganan

**¿Nueva tarea? Nueva conversación.** No luches contra el decaimiento de contexto. Las conversaciones largas se degradan naturalmente. Trabaja con ello, no en contra.

## 4. Calidad de Entrada = Calidad de Salida

Tu contexto acota tus resultados. Contexto rico y relevante produce salida rica. Contexto vago o ausente produce salida pobre.

```callout
type: tip
title: "El Hábito de Diagnóstico"
content: "Cuando la salida es mala, siempre pregúntate: ¿Cuál fue el contexto? ¿Qué tenía realmente el modelo con lo que trabajar? Nueve veces de cada diez, la respuesta está ahí."
```

## Aplicaciones Prácticas

**Antes de hacer un prompt:**
- ¿Tengo toda la información relevante lista para proporcionar?
- ¿Mi solicitud es clara y estructurada?
- ¿He incluido ejemplos de lo que quiero?

**Cuando la salida es pobre:**
- ¿Qué contexto le faltó al modelo?
- ¿Proporcioné instrucciones contradictorias?
- ¿Fue la conversación demasiado larga y saturada?

**Para tareas importantes:**
- Inicia conversaciones frescas
- Proporciona todo el contexto necesario de antemano
- Estructura tus entradas con Markdown

```quiz
id: context-takeaway-es
type: multiple-choice
question: "La IA produce un análisis genérico y superficial de los datos de tu empresa. Usaste un buen framework de prompt. ¿Cuál es el siguiente paso más productivo?"
options:
  - "Reescribir el prompt con instrucciones más específicas sobre profundidad y detalle"
  - "Proporcionar contexto fundamentado: documento de estrategia de tu empresa, análisis pasados y datos de la industria relevantes"
  - "Cambiar a un modelo de IA más capaz que pueda producir análisis más profundos"
answer: 1
explanation: "Un buen prompt con contexto pobre produce salida genérica. El modelo no tiene conocimiento específico de tu empresa del cual extraer. Proporcionar materiales fundamentados — documentos de estrategia, análisis pasados, datos de la industria — le da el contexto específico necesario para análisis sustantivos y personalizados. Mejores prompts no pueden compensar contexto faltante."
```

## Siguiente

Ahora que entiendes el contexto, hablemos sobre cómo estructurar tus solicitudes efectivamente. **Prompting** es el arte de dar instrucciones claras y contexto que produzcan los resultados que quieres.
