// Training App Client-Side Initialization
// Handles annotation canvas

(function() {
  'use strict';

  // Scroll to top when page or slide changes (but not on quiz answers or hotspot toggles)
  let lastPageID = null;
  function checkScrollToTop() {
    const article = document.querySelector('article[data-page-id]');
    if (!article) return;
    const id = article.getAttribute('data-page-id');
    if (id && id !== lastPageID) {
      lastPageID = id;
      // Use requestAnimationFrame to ensure layout is complete before scrolling
      // (prevents intermittent failures on long pages)
      requestAnimationFrame(() => window.scrollTo(0, 0));
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

  function drawSingleStroke(ctx, stroke) {
    if (!stroke.points || stroke.points.length < 2) return;

    ctx.save();
    if (stroke.tool === 'highlighter') {
      ctx.strokeStyle = 'rgba(250, 204, 21, 0.4)';
      ctx.lineWidth = 20;
      ctx.lineCap = 'round';
      ctx.lineJoin = 'round';
    } else {
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

  function eventToPixel(e) {
    return {
      x: Math.round(e.offsetX),
      y: Math.round(e.offsetY)
    };
  }

  function selectAnnotationTool(tool) {
    const layer = document.getElementById('annotation-layer');
    if (!layer) return;

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
  // Init & afterRender hook
  // ============================================================================

  // Scroll agent messages to bottom after staggered animations complete
  function scrollAgentMessages() {
    var msgContainer = document.getElementById('agent-messages');
    if (!msgContainer) return;

    // Find max animation delay by counting animated messages
    var animatedCount = 0;
    var children = msgContainer.children;
    for (var i = 0; i < children.length; i++) {
      var style = children[i].getAttribute('style') || '';
      if (style.indexOf('agentFadeSlideIn') !== -1) {
        animatedCount++;
      }
    }

    // Delay scroll to after animations complete
    var delay = animatedCount > 0 ? (animatedCount * 400 + 400) : 100;
    setTimeout(function() {
      msgContainer.scrollTop = msgContainer.scrollHeight;
    }, delay);
  }

  function afterRender() {
    checkScrollToTop();
    initAnnotations();
    scrollAgentMessages();
  }

  // Run initialization after DOM is ready
  if (document.readyState === 'loading') {
    document.addEventListener('DOMContentLoaded', afterRender);
  } else {
    afterRender();
  }

  // Register afterRender hook with gotea — called after every morphdom patch.
  window.gotea = window.gotea || {};
  window.gotea._afterRender = afterRender;

  // Expose for manual calls
  window.TrainingApp = {
    initAnnotations: initAnnotations,
    selectAnnotationTool: selectAnnotationTool
  };
})();
