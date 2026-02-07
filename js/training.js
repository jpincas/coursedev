// Training App Client-Side Initialization
// Handles Mermaid, KaTeX, asciinema player, and annotation canvas after morphdom patches

(function() {
  'use strict';

  // Cache for rendered Mermaid SVGs and KaTeX HTML, keyed by source text.
  // Morphdom patches replace rendered output with original source text;
  // the cache lets us restore instantly without re-calling the library.
  var renderCache = {};

  // Disable mermaid auto-rendering so we control all rendering ourselves.
  // This prevents the race where mermaid auto-renders but we can't cache
  // the result (because data-source isn't set yet), then morphdom wipes it.
  if (typeof mermaid !== 'undefined') {
    mermaid.initialize({ startOnLoad: false, theme: 'dark' });
  }

  // Track in-flight mermaid renders to avoid duplicates
  var mermaidRendering = {};
  var mermaidIdCounter = 0;

  // Initialize Mermaid diagrams
  function initMermaid() {
    if (typeof mermaid === 'undefined') return;

    document.querySelectorAll('.mermaid').forEach(function(el) {
      // Already rendered — cache and skip
      if (el.querySelector('svg')) {
        var src = el.getAttribute('data-source');
        if (src && !renderCache[src]) {
          renderCache[src] = el.innerHTML;
        }
        return;
      }

      var src = el.textContent.trim();
      if (!src) return;

      // Restore from cache (instant, survives morphdom patches)
      if (renderCache[src]) {
        el.innerHTML = renderCache[src];
        el.setAttribute('data-source', src);
        return;
      }

      // Already rendering this source — skip, the callback will handle it
      if (mermaidRendering[src]) return;

      // Render asynchronously using mermaid.render() which returns SVG as string.
      // This decouples rendering from the DOM element, so morphdom patches
      // can't interrupt the render. The SVG is cached and applied when ready.
      el.setAttribute('data-source', src);
      mermaidRendering[src] = true;
      var renderID = 'mermaid-render-' + (++mermaidIdCounter);
      mermaid.render(renderID, src).then(function(result) {
        delete mermaidRendering[src];
        renderCache[src] = result.svg;
        // Apply to all matching elements currently in DOM
        document.querySelectorAll('.mermaid').forEach(function(target) {
          if (target.getAttribute('data-source') === src || target.textContent.trim() === src) {
            target.innerHTML = result.svg;
            target.setAttribute('data-source', src);
          }
        });
      }).catch(function(err) {
        delete mermaidRendering[src];
        console.error('Mermaid render error:', err);
      });
    });
  }

  // Initialize KaTeX math blocks
  function initKaTeX() {
    if (typeof katex === 'undefined') return;

    document.querySelectorAll('.katex-block').forEach(function(el) {
      // Already rendered — cache and skip
      if (el.querySelector('.katex')) {
        var src = el.getAttribute('data-source');
        if (src && !renderCache[src]) {
          renderCache[src] = el.innerHTML;
        }
        return;
      }

      var src = el.textContent.trim();
      if (!src) return;

      // Restore from cache
      if (renderCache[src]) {
        el.innerHTML = renderCache[src];
        el.setAttribute('data-source', src);
        el.setAttribute('data-processed', 'true');
        return;
      }

      // First render
      el.setAttribute('data-source', src);
      try {
        katex.render(src, el, {
          throwOnError: false,
          displayMode: true
        });
        el.setAttribute('data-processed', 'true');
        renderCache[src] = el.innerHTML;
      } catch (e) {
        console.error('KaTeX render error:', e);
      }
    });
  }

  // Initialize asciinema players
  function initAsciinema() {
    if (typeof AsciinemaPlayer === 'undefined') return;

    document.querySelectorAll('.asciinema-player:not([data-processed])').forEach(el => {
      const src = el.dataset.src;
      const autoplay = el.dataset.autoplay === 'true';
      const speed = parseFloat(el.dataset.speed) || 1.0;

      if (src) {
        AsciinemaPlayer.create(src, el, {
          autoPlay: autoplay,
          speed: speed
        });
        el.setAttribute('data-processed', 'true');
      }
    });
  }

  // Scroll to top when page changes (but not on quiz answers or hotspot toggles)
  let lastPageID = null;
  function checkScrollToTop() {
    const article = document.querySelector('article[data-page-id]');
    if (!article) return;
    const id = article.getAttribute('data-page-id');
    if (id && id !== lastPageID) {
      lastPageID = id;
      window.scrollTo(0, 0);
    }
  }

  // ============================================================================
  // Annotation Canvas
  // ============================================================================

  // Module-level state that persists across morphdom updates
  let currentStroke = null;   // In-progress stroke {tool, points}
  let activeTool = null;      // 'pen', 'highlighter', or null
  let isDrawing = false;
  let canvasWithListeners = null; // Track which canvas element has listeners

  function initAnnotations() {
    const layer = document.getElementById('annotation-layer');

    if (!layer) {
      // No annotation layer — clean up if needed
      activeTool = null;
      isDrawing = false;
      canvasWithListeners = null;
      return;
    }

    const isPresenter = layer.dataset.isPresenter === 'true';
    let strokesData = [];
    try {
      strokesData = JSON.parse(layer.dataset.strokes || '[]');
    } catch (e) {
      strokesData = [];
    }

    // Find or create canvas inside the layer
    let canvas = layer.querySelector('canvas.annotation-canvas');
    if (!canvas) {
      canvas = document.createElement('canvas');
      canvas.className = 'annotation-canvas';
      canvas.style.position = 'absolute';
      canvas.style.top = '0';
      canvas.style.left = '0';
      layer.appendChild(canvas);
    }

    // Size canvas to match the #view container
    const container = document.getElementById('view');
    if (!container) return;

    const dpr = window.devicePixelRatio || 1;
    const w = container.clientWidth;
    const h = container.clientHeight;

    canvas.width = w * dpr;
    canvas.height = h * dpr;
    canvas.style.width = w + 'px';
    canvas.style.height = h + 'px';

    // Size the layer to match content
    layer.style.width = w + 'px';
    layer.style.height = h + 'px';

    const ctx = canvas.getContext('2d');
    ctx.setTransform(dpr, 0, 0, dpr, 0, 0);

    // Clear and redraw all strokes using absolute pixel coordinates
    ctx.clearRect(0, 0, w, h);
    drawStrokes(ctx, strokesData);

    // If mid-draw, redraw in-progress stroke too
    if (currentStroke && currentStroke.points.length > 1) {
      drawSingleStroke(ctx, currentStroke);
    }

    // Update pointer-events and text selection based on active tool
    const view = document.getElementById('view');
    if (isPresenter && activeTool) {
      layer.style.pointerEvents = 'auto';
      layer.style.cursor = 'crosshair';
      if (view) view.style.userSelect = 'none';
    } else {
      layer.style.pointerEvents = 'none';
      layer.style.cursor = '';
      if (view) view.style.userSelect = '';
    }

    // Attach drawing event listeners (presenter only, re-attach if canvas changed)
    if (isPresenter && canvas !== canvasWithListeners) {
      attachDrawingListeners(canvas, layer);
      canvasWithListeners = canvas;
    }

    // Update toolbar button states
    updateToolbarButtons();
  }

  function drawStrokes(ctx, strokes) {
    for (const stroke of strokes) {
      drawSingleStroke(ctx, stroke);
    }
  }

  // Strokes use absolute pixel coordinates — no scaling needed.
  // This works because #view has a fixed width (660px) so content
  // renders identically on all screens.
  function drawSingleStroke(ctx, stroke) {
    if (!stroke.points || stroke.points.length < 2) return;

    ctx.save();
    if (stroke.tool === 'highlighter') {
      ctx.strokeStyle = 'rgba(250, 204, 21, 0.4)';
      ctx.lineWidth = 20;
      ctx.lineCap = 'round';
      ctx.lineJoin = 'round';
    } else {
      // pen
      ctx.strokeStyle = '#ef4444';
      ctx.lineWidth = 3;
      ctx.lineCap = 'round';
      ctx.lineJoin = 'round';
    }

    ctx.beginPath();
    ctx.moveTo(stroke.points[0].x, stroke.points[0].y);
    for (let i = 1; i < stroke.points.length; i++) {
      ctx.lineTo(stroke.points[i].x, stroke.points[i].y);
    }
    ctx.stroke();
    ctx.restore();
  }

  function attachDrawingListeners(canvas, layer) {
    canvas.addEventListener('pointerdown', function(e) {
      if (!activeTool) return;
      e.preventDefault();

      isDrawing = true;
      const pt = eventToPixel(e);
      currentStroke = { tool: activeTool, points: [pt] };
    });

    canvas.addEventListener('pointermove', function(e) {
      if (!isDrawing || !currentStroke) return;
      e.preventDefault();

      const pt = eventToPixel(e);

      // Highlighter: lock Y to start position for straight horizontal lines
      if (currentStroke.tool === 'highlighter') {
        pt.y = currentStroke.points[0].y;
      }

      // Skip if too close to last point (reduces payload size)
      const last = currentStroke.points[currentStroke.points.length - 1];
      const dx = pt.x - last.x;
      const dy = pt.y - last.y;
      if (dx * dx + dy * dy < 9) return; // 3px threshold

      currentStroke.points.push(pt);

      // Draw the latest segment immediately for smooth feedback
      const ctx = canvas.getContext('2d');
      const dpr = window.devicePixelRatio || 1;
      ctx.setTransform(dpr, 0, 0, dpr, 0, 0);
      const points = currentStroke.points;
      const len = points.length;

      ctx.save();
      if (currentStroke.tool === 'highlighter') {
        ctx.strokeStyle = 'rgba(250, 204, 21, 0.4)';
        ctx.lineWidth = 20;
      } else {
        ctx.strokeStyle = '#ef4444';
        ctx.lineWidth = 3;
      }
      ctx.lineCap = 'round';
      ctx.lineJoin = 'round';
      ctx.beginPath();
      ctx.moveTo(points[len - 2].x, points[len - 2].y);
      ctx.lineTo(points[len - 1].x, points[len - 1].y);
      ctx.stroke();
      ctx.restore();
    });

    canvas.addEventListener('pointerup', function(e) {
      if (!isDrawing || !currentStroke) return;
      e.preventDefault();
      isDrawing = false;

      if (currentStroke.points.length >= 2) {
        // Send completed stroke to server
        if (typeof gotea !== 'undefined') {
          gotea.sendMessage({
            message: 'ADD_ANNOTATION',
            args: { tool: currentStroke.tool, points: currentStroke.points }
          });
        }
      }
      currentStroke = null;
    });

    canvas.addEventListener('pointerleave', function(e) {
      if (!isDrawing || !currentStroke) return;
      isDrawing = false;

      if (currentStroke.points.length >= 2) {
        if (typeof gotea !== 'undefined') {
          gotea.sendMessage({
            message: 'ADD_ANNOTATION',
            args: { tool: currentStroke.tool, points: currentStroke.points }
          });
        }
      }
      currentStroke = null;
    });
  }

  // Convert pointer event to absolute pixel coordinates on the canvas.
  // offsetX/offsetY are relative to the canvas's own top-left corner.
  // We use absolute pixels (not percentages) because #view has a fixed width,
  // so content renders at the same pixel positions on all screens.
  function eventToPixel(e) {
    return {
      x: Math.round(e.offsetX),
      y: Math.round(e.offsetY)
    };
  }

  function selectAnnotationTool(tool) {
    const layer = document.getElementById('annotation-layer');
    if (!layer) return;

    // Toggle: clicking same tool deselects
    const view = document.getElementById('view');
    if (activeTool === tool) {
      activeTool = null;
      layer.style.pointerEvents = 'none';
      layer.style.cursor = '';
      if (view) view.style.userSelect = '';
    } else {
      activeTool = tool;
      layer.style.pointerEvents = 'auto';
      layer.style.cursor = 'crosshair';
      if (view) view.style.userSelect = 'none';
      window.getSelection().removeAllRanges();
    }

    updateToolbarButtons();
  }

  function updateToolbarButtons() {
    const penBtn = document.getElementById('annotation-tool-pen');
    const hlBtn = document.getElementById('annotation-tool-highlighter');
    if (!penBtn || !hlBtn) return;

    const activeClass = 'py-2 px-4 text-sm rounded-full font-medium cursor-pointer transition-all duration-150 border-none';
    const inactiveStyle = activeClass + ' bg-zinc-800 text-zinc-300 hover:bg-zinc-700';
    const penActiveStyle = activeClass + ' bg-red-500 text-white';
    const hlActiveStyle = activeClass + ' bg-yellow-400 text-zinc-950';

    penBtn.className = activeTool === 'pen' ? penActiveStyle : inactiveStyle;
    hlBtn.className = activeTool === 'highlighter' ? hlActiveStyle : inactiveStyle;
  }

  // ============================================================================
  // Init All
  // ============================================================================

  function initAll() {
    checkScrollToTop();
    initMermaid();
    initKaTeX();
    initAsciinema();
    initAnnotations();
  }

  // Run initialization after DOM is ready
  if (document.readyState === 'loading') {
    document.addEventListener('DOMContentLoaded', initAll);
  } else {
    initAll();
  }

  // Re-initialize after morphdom patches
  const observer = new MutationObserver(function(mutations) {
    // Debounce to avoid multiple rapid calls
    clearTimeout(observer.timeout);
    observer.timeout = setTimeout(initAll, 50);
  });

  observer.observe(document.body, {
    childList: true,
    subtree: true
  });

  // Expose for manual calls if needed
  window.TrainingApp = {
    initMermaid: initMermaid,
    initKaTeX: initKaTeX,
    initAsciinema: initAsciinema,
    initAnnotations: initAnnotations,
    selectAnnotationTool: selectAnnotationTool,
    initAll: initAll
  };
})();
