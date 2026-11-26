// The voice-guided fitness self-test: push-ups, squats in one minute, a plank
// hold and pull-ups. Each result is entered on the spot; at the end the results
// go to /api/fitness/test and the result screen (/fitness) opens.
//
// Loaded as a managed inline script: $data is injected by doors.
(function () {
  var data = $data("payload") || {};
  if (typeof data === "string") { try { data = JSON.parse(data); } catch (_) { data = {}; } }
  var steps = data.steps || [];
  var L = data.labels || {};
  var root = document.querySelector(".fit-test");
  if (!root || !steps.length) return;
  function $(sel) { return root.querySelector(sel); }
  var el = {
    bar: $(".fit-bar"), img: $(".ex-svg"), name: $(".ex-name"), hint: $(".fit-hint"),
    count: $(".fit-count"), value: $(".fit-value"), minus: $(".fit-minus"), plus: $(".fit-plus"),
    clock: $(".fit-clock"), main: $(".fit-main"), back: $(".fit-back"), skip: $(".fit-skip-step"), err: $(".fit-err")
  };
  var results = {};       // step id -> number (-1 = skipped)
  var idx = 0;
  var phase = "";         // count | ready | running | enter
  var value = 0;
  var startedAt = 0, ticker = null, frameTimer = null, sending = false;

  // ---- voice ----
  function speak(text) {
    if (!data.voice || !text || !window.speechSynthesis) return;
    try {
      window.speechSynthesis.cancel();
      var u = new SpeechSynthesisUtterance(text);
      u.lang = data.bcp47 || "en-US";
      window.speechSynthesis.speak(u);
    } catch (_) {}
  }
  function beep() {
    try {
      var Ctx = window.AudioContext || window.webkitAudioContext;
      var ac = new Ctx(), o = ac.createOscillator(), g = ac.createGain();
      o.frequency.value = 880; g.gain.value = 0.25;
      o.connect(g); g.connect(ac.destination);
      o.start(); o.stop(ac.currentTime + 0.35);
    } catch (_) {}
  }

  // ---- drawing ----
  function show(node, on) { if (node) node.classList.toggle("hidden", !on); }
  function clock(sec) { sec = Math.max(0, Math.floor(sec)); return Math.floor(sec / 60) + ":" + ("0" + sec % 60).slice(-2); }
  function drawBar() {
    el.bar.textContent = "";
    for (var i = 0; i < steps.length; i++) {
      var seg = document.createElement("div");
      seg.className = "pbar-seg slot-warmup" + (i === idx ? " active" : "");
      var fill = document.createElement("div");
      fill.className = "pbar-seg-fill";
      fill.style.width = i < idx ? "100%" : "0";
      seg.appendChild(fill);
      seg.style.flex = "1";
      el.bar.appendChild(seg);
    }
  }
  function drawFrames(step) {
    clearInterval(frameTimer);
    var frames = step.frames || [], f = 0;
    if (!frames.length) { el.img.removeAttribute("src"); return; }
    el.img.src = frames[0];
    if (frames.length > 1) {
      frameTimer = setInterval(function () { f = (f + 1) % frames.length; el.img.src = frames[f]; }, 900);
    }
  }
  function draw() {
    var step = steps[idx];
    el.value.textContent = String(value);
    var counting = phase === "count" || phase === "enter";
    show(el.count, counting);
    show(el.clock, !counting);
    if (step.kind === "minute" && !counting) el.clock.textContent = clock(phase === "running" ? 60 - (Date.now() - startedAt) / 1000 : 60);
    if (step.kind === "hold") el.clock.textContent = clock(phase === "running" ? (Date.now() - startedAt) / 1000 : 0);
    el.main.textContent = counting ? L.done : phase === "running" ? L.stop : L.start;
    show(el.main.parentElement, !(step.kind === "minute" && phase === "running"));
    el.back.disabled = idx === 0;
  }
  function open(i) {
    stopClock();
    idx = i;
    var step = steps[idx];
    phase = step.kind === "count" ? "count" : "ready";
    value = results[step.id] > 0 ? results[step.id] : 0;
    el.name.textContent = step.name;
    el.hint.textContent = step.hint;
    show(el.err, false);
    drawBar();
    drawFrames(step);
    draw();
    speak(step.name + ". " + step.hint);
  }

  // ---- clocks ----
  function stopClock() { clearInterval(ticker); ticker = null; }
  function startClock() {
    var step = steps[idx];
    phase = "running";
    startedAt = Date.now();
    speak(L.go);
    ticker = setInterval(function () {
      if (step.kind === "minute" && Date.now() - startedAt >= 60000) {
        stopClock();
        beep();
        speak(L.stop);
        phase = "enter";
      }
      draw();
    }, 250);
    draw();
  }

  // ---- flow ----
  function record(v) {
    results[steps[idx].id] = v;
    if (idx + 1 < steps.length) { open(idx + 1); return; }
    finish();
  }
  function finish() {
    if (sending) return;
    stopClock(); clearInterval(frameTimer);
    var body = { pushUps: pick("push"), squats: pick("squat"), plankSec: pick("plank"), pullUps: pick("pull") };
    if (body.pushUps < 0 && body.squats < 0 && body.plankSec < 0 && body.pullUps < 0) { window.location.href = "/"; return; }
    sending = true;
    window.fetch("/api/fitness/test", {
      method: "POST", credentials: "same-origin",
      headers: { "Content-Type": "application/json" }, body: JSON.stringify(body)
    }).then(function (r) {
      if (!r.ok) throw new Error("status " + r.status);
      window.location.href = "/fitness";
    }).catch(function () {
      sending = false;
      el.err.textContent = L.error || "";
      show(el.err, true);
    });
  }
  function pick(id) { return typeof results[id] === "number" ? results[id] : -1; }

  el.minus.addEventListener("click", function () { value = Math.max(0, value - 1); draw(); });
  el.plus.addEventListener("click", function () { value = Math.min(200, value + 1); draw(); });
  el.main.addEventListener("click", function () {
    if (sending) return;
    var step = steps[idx];
    if (phase === "count" || phase === "enter") { record(value); return; }
    if (phase === "ready") { startClock(); return; }
    if (phase === "running" && step.kind === "hold") {
      var held = Math.round((Date.now() - startedAt) / 1000);
      stopClock();
      record(held);
    }
  });
  el.back.addEventListener("click", function () { if (idx > 0) open(idx - 1); });
  el.skip.addEventListener("click", function () { stopClock(); record(-1); });
  el.back.textContent = L.back || "";
  el.skip.textContent = L.skip || "";
  open(0);
})();
