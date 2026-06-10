(function () {
  'use strict';

  var canvas = document.getElementById('bg-canvas');
  if (!canvas) return;
  var ctx = canvas.getContext('2d');

  var W = 0;
  var H = 0;
  var dpr = window.devicePixelRatio || 1;
  var particles = [];
  var COUNT = 120;

  function resize() {
    W = window.innerWidth;
    H = window.innerHeight;
    canvas.width = W * dpr;
    canvas.height = H * dpr;
    canvas.style.width = W + 'px';
    canvas.style.height = H + 'px';
    ctx.scale(dpr, dpr);
  }

  function rand(lo, hi) {
    return Math.random() * (hi - lo) + lo;
  }

  function makeParticle(offscreen) {
    var r = rand(1, 20);
    var x = offscreen ? -r - rand(0, W * 0.5) : rand(0, W);
    var cy = H * 0.925;
    var spread = H * 0.075;
    var y = cy + rand(-spread, spread);
    var a = rand(0.02, 0.3);
    return {
      x: x,
      y: y,
      baseY: y,
      r: r,
      a: a,
      vx: a * 0.8 + 0.3,
      amp: rand(30, 160),
      hue: rand(195, 220),
      sat: rand(50, 80)
    };
  }

  function init() {
    resize();
    particles = [];
    for (var i = 0; i < COUNT; i++) {
      particles.push(makeParticle(false));
    }
  }

  function draw() {
    ctx.clearRect(0, 0, W, H);

    for (var i = 0; i < particles.length; i++) {
      var p = particles[i];

      p.x += p.vx;
      p.y = p.baseY + p.amp * Math.sin(p.x * 0.003);

      if (p.x - p.r > W) {
        particles[i] = makeParticle(true);
        continue;
      }

      ctx.beginPath();
      ctx.arc(p.x, p.y, Math.max(p.r, 0.5), 0, Math.PI * 2);
      ctx.fillStyle = 'hsla(' + p.hue + ',' + p.sat + '%,55%,' + p.a + ')';
      ctx.fill();
    }

    requestAnimationFrame(draw);
  }

  window.addEventListener('resize', function () {
    resize();
  });

  init();
  draw();
})();
