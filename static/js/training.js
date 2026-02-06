// Training App Client-Side Initialization
// Handles Mermaid, KaTeX, and asciinema player initialization after morphdom patches

(function() {
  'use strict';

  // Initialize Mermaid diagrams
  function initMermaid() {
    if (typeof mermaid === 'undefined') return;

    const elements = document.querySelectorAll('.mermaid:not([data-processed])');
    if (elements.length === 0) return;

    mermaid.init(undefined, elements);
    elements.forEach(el => el.setAttribute('data-processed', 'true'));
  }

  // Initialize KaTeX math blocks
  function initKaTeX() {
    if (typeof katex === 'undefined') return;

    document.querySelectorAll('.katex-block:not([data-processed])').forEach(el => {
      try {
        katex.render(el.textContent, el, {
          throwOnError: false,
          displayMode: true
        });
        el.setAttribute('data-processed', 'true');
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

  // Initialize all client-side rendered elements
  function initAll() {
    initMermaid();
    initKaTeX();
    initAsciinema();
  }

  // Run initialization after DOM is ready
  if (document.readyState === 'loading') {
    document.addEventListener('DOMContentLoaded', initAll);
  } else {
    initAll();
  }

  // Re-initialize after morphdom patches
  // Gotea triggers a custom event after updating the DOM
  // We use a MutationObserver as fallback for detecting DOM changes
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
    initAll: initAll
  };
})();
