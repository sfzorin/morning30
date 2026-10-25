// Morning workout player — runs entirely client-side from a server-provided
// payload. It steps through exercises (reps + timed), rest periods, voice cues
// and the celebratory finish, then durably reports completion through the stable HTTP API.
//
// Loaded as a managed inline script: $data / $sys are injected by doors.

(function () {
  var raw;
  try { raw = $data("payload"); } catch (e) { raw = null; }
  var payload = {};
  try {
    if (typeof raw === "string") payload = JSON.parse(raw);
    else if (raw && typeof raw === "object") payload = raw;
  } catch (e) { payload = {}; }
  var items = (payload.items || []).slice();
  var originalItems = items.slice();
  var slotRounds = Object.create(null);
  items.forEach(function (it) { slotRounds[it.slot] = Math.max(slotRounds[it.slot] || 1, it.round || 1); });
  var store = window.MorningWorkoutState;
  if (store) store.start(payload.userId);
  var labels = payload.labels || {};
  var cues = payload.cues || {};
  var enc = payload.enc || [];
  var voiceMode = payload.voiceMode || "normal";   // off | min | normal | detailed
  var voiceOn = voiceMode !== "off";
  var voiceLang = payload.bcp47 || "en-US";        // e.g. "tr-TR"
  var voicePrefix = (payload.lang || "en").toLowerCase(); // e.g. "tr"
  var fixedDifficulty = !!payload.fixedDifficulty;
  var selfManagedSides = !!payload.selfManagedSides;
  var level = fixedDifficulty ? 0 : (payload.level || 0); // difficulty −3..+3

  // jsScale mirrors the server's ScaleValue: ±10% per level on reps/seconds,
  // breaths fixed, with min clamps. Lets us re-scale main sets on the fly.
  function jsScale(base, unit, lvl) {
    if (lvl === 0 || unit === "breaths") return base;
    var out = Math.floor((base * (10 + lvl) + 5) / 10);
    if (unit === "reps" && out < 1) out = 1;
    if (unit === "seconds" && out < 5) out = 5;
    return out;
  }
  function val(it) { return it && it.scale && !fixedDifficulty ? jsScale(it.value, it.unit, level) : (it ? it.value : 0); }

  // Voice selection. Browsers otherwise default to an English voice even when
  // utterance.lang is "tr-TR", so we explicitly pick a voice whose lang matches
  // the target language (this is the Turkish-falls-back-to-English fix).
  var voices = [];
  var chosenVoice = null;
  function loadVoices() {
    if (!("speechSynthesis" in window)) return;
    try { voices = window.speechSynthesis.getVoices() || []; } catch (e) { voices = []; }
    chosenVoice = null;
    var i;
    for (i = 0; i < voices.length; i++) { // exact region match (tr-TR)
      if (voices[i].lang && voices[i].lang.toLowerCase() === voiceLang.toLowerCase()) { chosenVoice = voices[i]; return; }
    }
    for (i = 0; i < voices.length; i++) { // any voice of the language (tr-*)
      var lng = (voices[i].lang || "").toLowerCase().replace("_", "-");
      if (lng === voicePrefix || lng.indexOf(voicePrefix + "-") === 0) { chosenVoice = voices[i]; return; }
    }
  }
  if ("speechSynthesis" in window) {
    loadVoices();
    window.speechSynthesis.onvoiceschanged = loadVoices;
  }

  var root = document.querySelector(".player") || (document.currentScript && document.currentScript.parentElement);
  // Call sites pass CSS selectors (".pbar", ".ex-svg"). Do not prepend another
  // "." — querySelector("..pbar") throws and leaves the player blank.
  var $ = function (sel) {
    if (!root) return null;
    return root.querySelector(sel);
  };

  var el = {
    pbar: $(".pbar"),
    phase: $(".phase"),
    index: $(".ex-index"),
    svg: $(".ex-svg"),
    visual: $(".ex-visual"),
    demoCaption: $(".demo-caption"),
    name: $(".ex-name"),
    value: $(".ex-value"),
    stage: $(".stage"),
    infoBtn: $(".info-btn"),
    infoOverlay: $(".info-overlay"),
    infoMedia: $(".info-media"),
    infoVisual: $(".info-visual"),
    demoControls: $(".demo-controls"),
    demoToggle: $(".demo-toggle"),
    demoPrev: $(".demo-prev"),
    demoNext: $(".demo-next"),
    demoFrame: $(".demo-frame"),
    infoName: $(".info-name"),
    infoDesc: $(".info-desc"),
    infoHow: $(".info-how"),
    infoHowSec: $(".info-how-sec"),
    infoCorrect: $(".info-correct"),
    infoCorrectSec: $(".info-correct-sec"),
    infoWrong: $(".info-wrong"),
    infoWrongSec: $(".info-wrong-sec"),
    infoBreathing: $(".info-breathing"),
    infoBreathingSec: $(".info-breathing-sec"),
    infoWarn: $(".info-warn"),
    infoWarnSec: $(".info-warn-sec"),
    infoReplace: $(".info-replace"),
    infoClose: $(".info-close"),
    controls: $(".controls"),
    prev: $(".ctl-prev"),
    skip: $(".ctl-skip"),
    done: $(".ctl-done"),
    rest: $(".rest-overlay"),
    restCount: $(".rest-count"),
    restNext: $(".rest-next"),
    restSkip: $(".rest-skip"),
    restBack: $(".rest-back"),
    restTitle: $(".rest-title"),
    restIndex: $(".rest-index"),
    doneOverlay: $(".done-overlay"),
    applause: $(".applause"),
    doneTitle: $(".done-title"),
    doneEnc: $(".done-enc"),
    doneStreak: $(".done-streak"),
    backHome: $(".back-home"),
    diff: $(".lvl-row"),
    diffUp: $(".lvl-up"),
    diffDown: $(".lvl-down"),
    diffLabel: $(".lvl-label")
  };
  var vis = function (e) { return e && !e.classList.contains("hidden"); };

  var i = 0;          // current item index
  var paused = false;
  var ticker = null;
  var frameTimer = null;
  var frameIdx = 0;
  var frameStep = 0;
  var failedAnimations = Object.create(null);
  var reducedMotion = !!(window.matchMedia && window.matchMedia("(prefers-reduced-motion: reduce)").matches);
  var disposed = false;
  var resumeNeeded = false;
  var redirecting = false;
  var restoring = false;
  var activeClock = null;
  var state = {
    index: 0, side: 1, phase: "ready", remainingMs: 3000,
    nextIndex: 0, nextSide: 1, level: level, swaps: {},
    eventId: store ? store.newEventId() : String(Date.now()) + Math.random().toString(16).slice(2).repeat(3),
    completion: null
  };
  function checkpoint() {
    state.index = i;
    state.side = curSide || 1;
    state.level = level;
    if (store && !disposed) store.saveCheckpoint(payload, state);
  }
  function running() { return !disposed && !finished && !paused && document.visibilityState === "visible"; }


  function framesOf(it) {
    if (it && it.frames && it.frames.length) return it.frames;
    if (it && it.svg) return [it.svg];
    return [];
  }
  function stopFrames() {
    if (frameTimer) { clearInterval(frameTimer); frameTimer = null; }
    if (el.demoToggle) el.demoToggle.textContent = t("demo_play");
  }
  function animationOf(it) {
    var a = it && it.animation;
    var poseURL = /^\/poses\/[a-z0-9_-]+(?:\?sex=[mf])?$/;
    return a && el.visual && el.infoVisual && typeof a.src === "string" &&
      poseURL.test(a.src) && !failedAnimations[a.src] &&
      (!a.otherSrc || (typeof a.otherSrc === "string" && poseURL.test(a.otherSrc) && !failedAnimations[a.otherSrc])) &&
      (a.frames === 2 || a.frames === 3) &&
      (a.kind === "movement" || a.kind === "sequence" || (a.kind === "alternating" && a.frames === 3)) ? a : null;
  }
  function frameCount(it) {
    var a = animationOf(it);
    return a ? a.frames * (a.otherSrc ? 2 : 1) : framesOf(it).length;
  }
  function paintFrame(it, idx) {
    var a = animationOf(it), frames = framesOf(it), count = frameCount(it);
    if (!count) return;
    frameIdx = ((idx % count) + count) % count;
    if (a) {
      [el.visual, el.infoVisual].forEach(function (visual) {
        visual.style.backgroundImage = 'url("' + (frameIdx < a.frames ? a.src : a.otherSrc) + '")';
        visual.style.backgroundPosition = ((frameIdx % a.frames) * 100 / (a.frames - 1)) + "% 50%";
      });
    } else {
      var src = frames[frameIdx];
      if (el.svg) el.svg.src = src;
      if (el.infoMedia) el.infoMedia.src = src;
    }
    if (el.demoFrame) el.demoFrame.textContent = t("demo_frame") + " " + (frameIdx + 1) + "/" + count;
  }
  function playFrames(it, inInfo) {
    stopFrames();
    var a = animationOf(it), count = frameCount(it);
    if (count < 2 || disposed || document.visibilityState !== "visible" || (!inInfo && !running())) return;
    // Return through the middle for a movement, or through neutral between sides.
    var order = [];
    if (a && a.kind === "sequence") {
      for (var n = 0; n < count; n++) order.push(n);
    } else if (a && a.kind === "alternating") order = a.otherSrc ? [0, 1, 0, 2, 3, 4, 3, 5] : [0, 1, 0, 2];
    else {
      var stripSize = a ? a.frames : count;
      for (var base = 0; base < count; base += stripSize) {
        for (var f = 0; f < stripSize; f++) order.push(base + f);
        for (var back = stripSize - 2; back >= (a && a.otherSrc ? 0 : 1); back--) order.push(base + back);
      }
    }
    frameStep = Math.max(0, order.indexOf(frameIdx));
    if (el.demoToggle) el.demoToggle.textContent = t("demo_pause");
    frameTimer = setInterval(function () {
      if (disposed || document.visibilityState !== "visible" || (inInfo ? !vis(el.infoOverlay) : !running())) { stopFrames(); return; }
      frameStep = (frameStep + 1) % order.length;
      paintFrame(it, order[frameStep]);
    }, 1000);
  }
  function showFrames(it) {
    stopFrames();
    var a = animationOf(it), frames = framesOf(it), count = frameCount(it);
    [el.visual, el.infoVisual].forEach(function (visual) {
      if (!visual) return;
      visual.classList.toggle("hidden", !count);
      visual.style.backgroundImage = a ? 'url("' + a.src + '")' : "none";
      visual.style.backgroundSize = a ? (a.frames * 100) + "% 100%" : "";
      visual.setAttribute("aria-label", it.name || "");
    });
    [el.svg, el.infoMedia].forEach(function (img) {
      if (!img) return;
      img.classList.toggle("hidden", !!a || !frames.length);
      if (a || !frames.length) img.removeAttribute("src");
      img.alt = it.name || "";
    });
    if (el.demoCaption) {
      el.demoCaption.classList.toggle("hidden", !a);
      el.demoCaption.textContent = a && (a.kind === "alternating" || (a.otherSrc && it.perSide)) ? t("demo_both") : t("demo_play");
    }
    [el.demoControls, el.demoFrame].forEach(function (node) { if (node) node.classList.toggle("hidden", count < 2); });
    paintFrame(it, 0);
    if (a) {
      // Load only this exercise. A missing strip must retain the original still.
      [a.src, a.otherSrc].filter(Boolean).forEach(function (src) {
        var preload = new Image();
        preload.onerror = function () {
          failedAnimations[src] = true;
          if (!disposed && items[i] === it) {
            showFrames(it);
            if (vis(el.infoOverlay) && !reducedMotion) playFrames(it, true);
          }
        };
        preload.src = src;
      });
    }
    if (!reducedMotion) playFrames(it, false);
  }
  var finished = false;
  var wakeLock = null;

  // ---- helpers ----------------------------------------------------------
  function t(key) { return labels[key] || key; }

  function slotLabel(slot) {
    if (slot === "warmup") return t("warmup");
    if (slot === "cooldown") return t("cooldown");
    return t("main");
  }

  // ---- speech coordinator ----------------------------------------------
  // Two priorities, so the spoken countdown always lands on its second and
  // never queues behind technique narration:
  //   sayNow(text) — time-critical (countdown digits, Go, Rest, next-exercise,
  //                  half-time): cancels whatever is speaking, then speaks.
  //   utter(text)  — plain queued speech (narration, the rest announcement,
  //                  finish messages); any sayNow cuts the current line off.
  function utter(text) {
    if (!voiceOn || !text || disposed || paused || document.visibilityState !== "visible" || !("speechSynthesis" in window)) return;
    try {
      if (!voices.length || !chosenVoice) loadVoices();
      var u = new SpeechSynthesisUtterance(text);
      if (chosenVoice) {
        u.voice = chosenVoice;
        u.lang = chosenVoice.lang;
      } else {
        u.lang = voiceLang; // no matching voice installed on this device
      }
      u.rate = 1.0;
      window.speechSynthesis.speak(u);
    } catch (e) {}
  }
  var speak = utter; // queued (finish / celebration messages)
  function sayNow(text) {
    if (!voiceOn || !text) return;
    shutUp();      // drop any in-progress narration so the cue is on time
    utter(text);
  }
  function shutUp() {
    if ("speechSynthesis" in window) { try { window.speechSynthesis.cancel(); } catch (e) {} }
  }

  function clearTicker() {
    if (ticker) { clearInterval(ticker); ticker = null; }
  }

  // talkative covers the two spoken levels: "detailed" (chatty) and "normal".
  function talkative() { return voiceMode === "normal" || voiceMode === "detailed"; }

  // ---- scheduled speech: technique narration + the rest announcement -------
  // narrTimer drives the short technique cues; announceTimer holds the
  // delayed "Rest / next exercise" line. clearNarr cancels both — they are the
  // only pre-scheduled speech, and both must die on every transition.
  var narrTimer = null, narrQ = [], narrI = 0, announceTimer = null;
  function clearNarr() {
    if (narrTimer) { clearTimeout(narrTimer); narrTimer = null; }
    if (announceTimer) { clearTimeout(announceTimer); announceTimer = null; }
    narrQ = []; narrI = 0;
  }
  // startNarr begins the narration a few seconds into the exercise (after the Go
  // cue), spaced out and low priority. runSide stops it before the closing
  // countdown; any sayNow cue cuts the current line off. detailed = all cues.
  function startNarr(it) {
    if (!voiceOn || !talkative()) return;
    var c = it.correct || [], w = it.wrong || [], k;
    var nc = voiceMode === "detailed" ? c.length : Math.min(2, c.length);
    var nw = voiceMode === "detailed" ? w.length : Math.min(1, w.length);
    narrQ = []; narrI = 0;
    for (k = 0; k < nc; k++) narrQ.push(c[k]); // how to do it (phrased as a do)
    for (k = 0; k < nw; k++) narrQ.push(w[k]); // speak the cue without an introduction
    if (!narrQ.length) return;
    var gap = voiceMode === "detailed" ? 6500 : 11000;
    function step() {
      if (paused) { narrTimer = setTimeout(step, 1000); return; }
      if (narrI >= narrQ.length) { narrTimer = null; return; }
      utter(narrQ[narrI]); narrI++;
      narrTimer = setTimeout(step, gap);
    }
    narrTimer = setTimeout(step, 3000); // small pause after "Go" before instructions
  }

  // flashHalf speaks "half the time" at the midpoint of a timed set, clear of
  // both the start and the closing count. Voice only — nothing on screen.
  function flashHalf() {
    if (talkative()) sayNow(cues.halfway);
  }

  // Keep the remaining duration outside the timer closure so every phase can
  // survive page recreation. Hidden time never counts as exercise time.
  function countdown(seconds, onTick, onDone, remainingMs) {
    clearTicker();
    state.remainingMs = remainingMs == null ? seconds * 1000 : remainingMs;
    activeClock = { lastAt: Date.now(), onTick: onTick, onDone: onDone, shown: -1 };
    onTick(Math.ceil(state.remainingMs / 1000));
    activeClock.shown = Math.ceil(state.remainingMs / 1000);
    checkpoint();
    if (paused) return;
    if (state.remainingMs <= 0) { activeClock = null; onDone(); return; }
    ticker = setInterval(function () {
      if (!running()) return;
      var clock = activeClock;
      if (!clock) return;
      state.remainingMs = Math.max(0, state.remainingMs - Math.max(0, Date.now() - clock.lastAt));
      clock.lastAt = Date.now();
      var remaining = Math.ceil(state.remainingMs / 1000);
      if (remaining !== clock.shown) { clock.shown = remaining; clock.onTick(remaining); }
      checkpoint();
      if (state.remainingMs <= 0) {
        clearTicker(); activeClock = null;
        clock.onDone();
      }
    }, 1000);
  }

  function freezeClock() {
    if (activeClock && !paused) {
      state.remainingMs = Math.max(0, state.remainingMs - Math.max(0, Date.now() - activeClock.lastAt));
      activeClock.lastAt = Date.now();
    }
    clearTicker();
  }

  var wakeRequest = null;
  async function requestWake() {
    if (!running() || !("wakeLock" in navigator) || wakeRequest || (wakeLock && !wakeLock.released)) return;
    var request;
    try {
      request = navigator.wakeLock.request("screen");
      wakeRequest = request;
      var sentinel = await request;
      if (!running()) { await sentinel.release(); return; }
      wakeLock = sentinel;
      sentinel.addEventListener("release", function () { if (wakeLock === sentinel) wakeLock = null; });
    } catch (e) {
      // The Continue tap retries when a browser requires a user gesture.
    } finally {
      if (wakeRequest === request) wakeRequest = null;
    }
  }
  function releaseWake() {
    var sentinel = wakeLock;
    wakeLock = null;
    try { if (sentinel) sentinel.release(); } catch (e) {}
  }
  function pauseForReturn() {
    if (disposed || finished) return;
    freezeClock();
    paused = true;
    resumeNeeded = true;
    clearNarr(); stopFrames(); shutUp(); releaseWake();
    if (el.infoOverlay) el.infoOverlay.classList.add("hidden");
    root.classList.add("frozen");
    checkpoint();
    if (el.resumeOverlay) el.resumeOverlay.classList.remove("hidden");
  }
  // A scheduled cycle restart changes which day this page may run. Let the
  // server resolve the current day instead of resuming an overnight snapshot.
  function redirectAfterRestart() {
    if (disposed || finished || document.visibilityState !== "visible") return false;
    if (redirecting) return true;
    var deadline = typeof payload.restartAt === "string" ? Date.parse(payload.restartAt) : NaN;
    if (!Number.isFinite(deadline) || Date.now() < deadline) return false;
    redirecting = true;
    freezeClock(); paused = true;
    clearNarr(); stopFrames(); shutUp(); releaseWake();
    location.assign("/start");
    return true;
  }
  function onVisibility() {
    if (document.visibilityState !== "visible") pauseForReturn();
    else if (!redirectAfterRestart() && resumeNeeded && el.resumeOverlay) el.resumeOverlay.classList.remove("hidden");
  }
  function onPageShow() {
    if (redirectAfterRestart()) return;
    if (resumeNeeded && el.resumeOverlay) el.resumeOverlay.classList.remove("hidden");
    else if (running()) requestWake();
  }
  document.addEventListener("visibilitychange", onVisibility);
  window.addEventListener("pagehide", pauseForReturn);
  window.addEventListener("pageshow", onPageShow);
  window.addEventListener("focus", onPageShow);

  // ---- rendering --------------------------------------------------------
  // The top bar is split into one segment per phase (warm-up / main / cool-down),
  // each sized to its share of the exercises and tinted its phase colour. A
  // segment starts greyish-dim, brightens while it's the active phase, and fills
  // left→right as its exercises complete. barSegs is built once from the items.
  var barSegs = [];
  function buildBar() {
    barSegs = [];
    if (!el.pbar) return;
    el.pbar.textContent = "";
    var start = 0;
    for (var k = 1; k <= items.length; k++) {
      if (k === items.length || items[k].slot !== items[start].slot) {
        var slot = items[start].slot;
        var seg = document.createElement("div");
        seg.className = "pbar-seg slot-" + slot;
        seg.style.flexGrow = (k - start);          // width ∝ exercises in the phase
        var fill = document.createElement("div");
        fill.className = "pbar-seg-fill";
        seg.appendChild(fill);
        el.pbar.appendChild(seg);
        barSegs.push({ slot: slot, start: start, end: k, seg: seg, fill: fill });
        start = k;
      }
    }
  }

  // renderProgress fills each phase segment by how many of its exercises are done
  // and marks activeSlot's segment active. During a rest, donePos/activeSlot point
  // at the UPCOMING exercise, so the next phase lights up the moment its rest
  // begins (e.g. the main set "starts" during the pause that leads into it).
  function renderProgress(donePos, activeSlot) {
    for (var k = 0; k < barSegs.length; k++) {
      var s = barSegs[k];
      var n = s.end - s.start;
      var f = Math.max(0, Math.min(n, donePos - s.start));
      s.fill.style.width = (f / n * 100) + "%";
      s.seg.classList.toggle("active", s.slot === activeSlot);
    }
  }

  // setSlotTheme paints the whole player in the active phase colour
  // (warm-up = blue, main = amber, cool-down = mint) so counters, rest
  // overlay and primary buttons match the status-bar segment.
  function setSlotTheme(slot) {
    root.classList.remove("slot-warmup", "slot-main", "slot-cooldown");
    if (slot) root.classList.add("slot-" + slot);
  }

  // setPhaseLabel writes the phase + round into the top label, slot-coloured.
  function setPhaseLabel(it) {
    var totalRounds = slotRounds[it.slot] || 1;
    el.phase.textContent = slotLabel(it.slot) + " · " + t("round") + " " + it.round + (totalRounds > 1 ? "/" + totalRounds : "");
    el.phase.classList.remove("slot-warmup", "slot-main", "slot-cooldown");
    el.phase.classList.add("slot-" + it.slot);
    setSlotTheme(it.slot);
  }

  var pendingStart = null; // start-fn the rest/ready countdown will run at 0
  var curSide = 1;

  // Self-managed sides share one step; timed work still includes both sides.
  // Other plans retain separate timed sides and any prescribed switching pause.
  function sides(it) { return !selfManagedSides && it.perSide && (it.unit === "seconds" || it.sideRest > 0) ? 2 : 1; }
  function timedDuration(it) { return val(it) * (selfManagedSides && it.perSide ? 2 : 1); }

  function startItem(idx, side, remainingMs) {
    i = idx;
    var it = items[i];
    if (!it || !el.svg) return;
    if (root) root.classList.toggle("frozen", paused);
    if (el.rest) el.rest.classList.add("hidden");
    if (el.stage) el.stage.classList.remove("hidden");
    if (el.controls) el.controls.classList.remove("hidden");

    if (el.name) el.name.textContent = it.name;
    if (el.index) el.index.textContent = (i + 1) + " / " + items.length;
    setPhaseLabel(it);
    renderProgress(i, it.slot);
    runSide(it, side || 1, remainingMs);
  }

  // runSide runs one side. Each explicit side gets its own visible position.
  function runSide(it, side, remainingMs) {
    showFrames(it);
    curSide = side;
    state.phase = "exercise";
    state.nextIndex = i;
    state.nextSide = side;
    state.remainingMs = it.unit === "seconds" ? (remainingMs == null ? timedDuration(it) * 1000 : remainingMs) : null;
    if (el.rest) el.rest.classList.add("hidden");
    if (el.stage) el.stage.classList.remove("hidden");
    if (el.controls) el.controls.classList.remove("hidden");
    setPhaseLabel(it);
    if (sides(it) === 2) el.phase.textContent += " · " + t("side") + " " + side + "/2";
    if (selfManagedSides && it.perSide && it.unit === "seconds") el.phase.textContent += " · " + valueDisplay(it);
    clearNarr();
    if (side === 1 && !restoring) {
      sayNow(it.vStart);   // "Go" / "Старт"
      startNarr(it);       // first instruction ~3 s later
    }

    if (it.unit === "seconds") {
      el.done.classList.add("hidden");
      el.value.classList.add("timer");
      el.value.classList.remove("prep");
      var D = timedDuration(it);
      var halfAt = Math.round(D * 0.5);
      var lastSide = side >= sides(it);
      countdown(D, function (r) {
        el.value.textContent = r + "″";
        if (D >= 12 && r === halfAt && r > 5) flashHalf(); // mid-set, clear of the count
        if (r === 4) clearNarr();                // silence narration before the count
        if (r <= 3 && r >= 1 && lastSide) sayNow(String(r));
      }, function () { nextSide(it, side); }, remainingMs);
    } else {
      // reps / breaths: wait for the user to tap Done.
      clearTicker(); activeClock = null;
      el.value.classList.remove("timer", "prep");
      el.value.textContent = activeValueDisplay(it);
      el.done.classList.remove("hidden");
      checkpoint();
    }
  }

  function nextSide(it, side) {
    clearNarr();
    if (side < sides(it)) {
      sayNow(cues.switch_side);
      if (it.sideRest > 0) {
        doRest(it.sideRest, it, true, function () {
          sayNow(cues.go);
          runSide(it, side + 1);
        }, side + 1);
        el.restTitle.textContent = cues.switch_side || t("side");
        el.restNext.textContent = it.name + " · " + t("side") + " " + (side + 1) + "/2";
      } else {
        runSide(it, side + 1);
      }
    } else {
      endItem();
    }
  }

  function activeValueDisplay(it) {
    if (sides(it) !== 2) return valueDisplay(it);
    var displayItem = Object.assign({}, it, { perSide: false });
    return valueDisplay(displayItem);
  }

  // valueDisplay is the on-screen count for reps/breaths and the rest preview.
  function valueDisplay(it) {
    var v = val(it);
    if (it.unit === "seconds") {
      return it.perSide ? v + "″ " + t("per_side") : v + "″";
    }
    if (it.unit === "breaths") {
      return it.perSide ? v + " " + t("breaths") + " " + t("per_side") : v + " " + t("breaths");
    }
    return it.perSide ? "× " + v + " " + t("per_side") : "× " + v;
  }

  // unitFor returns the spoken unit word for an item.
  function unitWord(it) {
    if (it.unit === "seconds") {
      var w = cues.seconds || t("sec");
      return it.perSide ? w + " " + t("per_side") : w;
    }
    if (it.unit === "breaths") {
      var w = t("breaths");
      return it.perSide ? w + " " + t("per_side") : w;
    }
    var w = cues.reps || t("reps");
    return it.perSide ? w + " " + t("per_side") : w;
  }
  function nextDisplay(it) { // on-screen "Name · ×12 per side" / "Name · 30″ per side"
    return it.name + " · " + valueDisplay(it);
  }

  // doRest shows the rest/get-ready overlay and counts down. The big number
  // counts the seconds and the final 3 are spoken (3-2-1, exactly on the tick).
  // For a real rest it then speaks "Rest" + the next exercise with its count —
  // after a short beat so it never lands on top of the previous set's "one".
  // The get-ready before an exercise is just the spoken 3-2-1, leading into "Go".
  function doRest(seconds, upItem, isReady, onDone, nextSide, remainingMs) {
    if (!el.rest) { if (onDone) onDone(); return; }
    clearNarr(); stopFrames();
    state.phase = nextSide > 1 ? "side-rest" : (isReady ? "ready" : "rest");
    state.nextIndex = items.indexOf(upItem);
    state.nextSide = nextSide || 1;
    state.remainingMs = remainingMs == null ? seconds * 1000 : remainingMs;
    el.stage.classList.add("hidden");
    el.controls.classList.add("hidden");
    el.done.classList.add("hidden"); // Done lives on its own row, outside .controls
    el.rest.classList.remove("hidden");
    el.restTitle.textContent = isReady ? t("ready") : t("rest");
    el.restNext.textContent = upItem ? t("next") + ": " + nextDisplay(upItem) : "";
    // Keep the counter, phase label and bar aligned with the exercise the rest
    // leads into — the upcoming phase lights up during its lead-in pause.
    if (upItem) {
      var upIdx = items.indexOf(upItem);
      el.restIndex.textContent = (upIdx + 1) + " / " + items.length;
      setPhaseLabel(upItem);
      renderProgress(upIdx, upItem.slot);
    }
    if (state.phase === "side-rest") {
      el.restTitle.textContent = cues.switch_side || t("side");
      el.restNext.textContent = upItem.name + " · " + t("side") + " " + state.nextSide + "/2";
    }
    pendingStart = onDone;
    if (upItem && !isReady && seconds >= 6 && !restoring && !paused) {
      announceTimer = setTimeout(function () {
        announceTimer = null;
        if (paused || el.rest.classList.contains("hidden")) return;
        shutUp();
        if (cues.rest) utter(cues.rest); // "Rest" (queued)
        utter((cues.next_ex || t("next")) + ": " + upItem.name + ", " + val(upItem) + " " + unitWord(upItem));
      }, 700);
    }
    countdown(seconds, function (r) {
      el.restCount.textContent = r;
      if (r <= 3 && r >= 1) sayNow(String(r)); // digits land exactly on the second
    }, function () {
      var f = pendingStart; pendingStart = null;
      if (f) f();
    }, remainingMs);
  }

  // prev returns to the previous exercise (or restarts the current one when
  // resting), in case the user skipped or advanced by accident.
  function prev() {
    if (paused || finished) return;
    shutUp();
    clearTicker();
    paused = false;
    var target = vis(el.rest) ? i : Math.max(0, i - 1);
    doRest(3, items[target], true, function () { startItem(target); });
  }

  function endItem() {
    clearTicker();
    clearNarr();
    if (i >= items.length - 1) { finish(); return; }
    var rest = items[i].rest || 0;
    var next = i + 1;
    if (rest <= 0) { startItem(next); return; }
    doRest(rest, items[next], false, function () { startItem(next); });
  }

  function localDate(date) {
    return date.getFullYear() + "-" + String(date.getMonth() + 1).padStart(2, "0") + "-" + String(date.getDate()).padStart(2, "0");
  }
  var celebrated = false;
  function showCompletionStatus(status, receipt, durable) {
    if (disposed || !finished) return;
    el.saveStatus.textContent = t(status === "saved" ? "saved" : status === "saving" ? "saving" : (durable === false ? "save_unavailable" : "save_pending"));
    el.saveRetry.classList.toggle("hidden", status !== "pending");
    if (status !== "saved") {
      el.doneTitle.textContent = t("day_done");
      return;
    }
    if (el.applause) el.applause.classList.remove("hidden");
    el.doneStreak.classList.remove("hidden");
    var name = (payload.name || "").trim();
    el.doneTitle.textContent = name ? name + ", " + t("bravo") : t("done_title");
    if (receipt && typeof receipt.streak === "number") el.doneStreak.textContent = "🔥 " + receipt.streak + " · " + t("streak");
    if (!celebrated) {
      celebrated = true;
      var message = enc.length ? enc[Math.floor(Math.random() * enc.length)] : "";
      el.doneEnc.textContent = message;
      if (document.visibilityState === "visible") {
        celebrate();
        speak(name ? name + ", " + cues.well_done : cues.workout_end);
      }
    }
  }
  function sendCompletion() {
    if (!state.completion || disposed) return;
    if (!store) { showCompletionStatus("pending", null, false); return; }
    var queued = store.queueCompletion(state.completion);
    showCompletionStatus("saving", null, queued.durable);
    store.sendCompletion(state.completion).then(function (saved) {
      showCompletionStatus("saved", saved);
    }).catch(function () { showCompletionStatus("pending", null, queued.durable); });
  }
  // The event is stored before any network call. Closing the page while offline
  // leaves the same event for the authenticated home page to retry.
  function finish() {
    if (finished) return;
    finished = true;
    paused = false;
    state.phase = "complete";
    state.remainingMs = null;
    if (!state.completion) {
      var date = new Date();
      state.completion = { eventId: state.eventId, userId: payload.userId, day: payload.day,
        position: payload.position, completedAt: date.toISOString(), completedOn: localDate(date) };
    }
    checkpoint();
    renderProgress(items.length, null);
    clearTicker(); activeClock = null;
    stopFrames(); clearNarr(); shutUp(); releaseWake();
    el.resumeOverlay.classList.add("hidden");
    root.classList.remove("frozen");
    el.stage.classList.add("hidden"); el.controls.classList.add("hidden");
    el.done.classList.add("hidden");
    if (el.done.parentElement) el.done.parentElement.classList.add("hidden");
    el.rest.classList.add("hidden"); el.doneOverlay.classList.remove("hidden");
    if (el.applause) el.applause.classList.add("hidden");
    el.doneStreak.classList.add("hidden");
    sendCompletion();
  }

  // ---- celebration: confetti + a roaring crowd (applause, cheers, whistles) ----
  function celebrate() {
    confetti();
    try {
      var Ctx = window.AudioContext || window.webkitAudioContext;
      if (!Ctx) return;
      var ac = new Ctx();
      var t0 = ac.currentTime;
      var master = ac.createGain();
      master.gain.value = 0.9;
      master.connect(ac.destination);

      // Applause bed: a crowd of ~900 random claps over ~3 s, with a swell.
      var dur = 3.0;
      var buffer = ac.createBuffer(1, Math.floor(ac.sampleRate * dur), ac.sampleRate);
      var d = buffer.getChannelData(0), i, k;
      for (i = 0; i < 900; i++) {
        var at = Math.floor(Math.random() * d.length);
        var len = 40 + Math.floor(Math.random() * 80);
        var amp = 0.5 + Math.random() * 0.5;
        for (k = 0; k < len && at + k < d.length; k++) {
          var e = 1 - k / len;
          d[at + k] += (Math.random() * 2 - 1) * e * e * amp;
        }
      }
      var bed = ac.createBufferSource(); bed.buffer = buffer;
      var bp = ac.createBiquadFilter(); bp.type = "bandpass"; bp.frequency.value = 1900; bp.Q.value = 0.7;
      var bedG = ac.createGain();
      bedG.gain.setValueAtTime(0.0001, t0);
      bedG.gain.exponentialRampToValueAtTime(0.6, t0 + 0.25);
      bedG.gain.setValueAtTime(0.6, t0 + 1.7);
      bedG.gain.exponentialRampToValueAtTime(0.0001, t0 + dur);
      bed.connect(bp); bp.connect(bedG); bedG.connect(master);
      bed.start(t0);

      // Crowd roar: a throaty "ahhh" — brown-ish noise shaped by two vowel
      // formants, with a big swell and a slow undulation, like an approving
      // crowd. (Pitched oscillators sound fake; a real roar is filtered noise.)
      var roarBuf = ac.createBuffer(1, Math.floor(ac.sampleRate * dur), ac.sampleRate);
      var rd = roarBuf.getChannelData(0), lastp = 0;
      for (k = 0; k < rd.length; k++) {
        var wn = Math.random() * 2 - 1;
        lastp = (lastp + 0.02 * wn) / 1.02; // leaky integrator → brown-ish noise
        rd[k] = lastp * 3.0 + wn * 0.3;
      }
      var roar = ac.createBufferSource(); roar.buffer = roarBuf;
      var f1 = ac.createBiquadFilter(); f1.type = "bandpass"; f1.frequency.value = 620; f1.Q.value = 2.2;
      var f2 = ac.createBiquadFilter(); f2.type = "bandpass"; f2.frequency.value = 1300; f2.Q.value = 2.2;
      var roarG = ac.createGain();
      roarG.gain.setValueAtTime(0.0001, t0);
      roarG.gain.exponentialRampToValueAtTime(0.55, t0 + 0.4); // crowd swells up
      roarG.gain.setValueAtTime(0.55, t0 + 1.9);
      roarG.gain.exponentialRampToValueAtTime(0.0001, t0 + dur);
      var trem = ac.createOscillator(); trem.type = "sine"; trem.frequency.value = 4;
      var tremG = ac.createGain(); tremG.gain.value = 0.16; // undulation
      trem.connect(tremG); tremG.connect(roarG.gain);
      roar.connect(f1); roar.connect(f2);
      f1.connect(roarG); f2.connect(roarG); roarG.connect(master);
      roar.start(t0); roar.stop(t0 + dur); trem.start(t0); trem.stop(t0 + dur);

      // A couple of celebratory whistles.
      for (i = 0; i < 2; i++) {
        var ws = t0 + 0.3 + Math.random() * 1.2;
        var wo = ac.createOscillator(); wo.type = "sine";
        var wf = 2200 + Math.random() * 600;
        wo.frequency.setValueAtTime(wf, ws);
        wo.frequency.linearRampToValueAtTime(wf + 300, ws + 0.3);
        var wg = ac.createGain();
        wg.gain.setValueAtTime(0.0001, ws);
        wg.gain.exponentialRampToValueAtTime(0.06, ws + 0.05);
        wg.gain.exponentialRampToValueAtTime(0.0001, ws + 0.5);
        wo.connect(wg); wg.connect(master);
        wo.start(ws); wo.stop(ws + 0.55);
      }

      // Bright chime arpeggio (the firework "salute").
      [523, 659, 784, 1046].forEach(function (f, n) {
        var o = ac.createOscillator(), g = ac.createGain();
        o.type = "triangle"; o.frequency.value = f;
        o.connect(g); g.connect(master);
        var s = t0 + n * 0.12;
        g.gain.setValueAtTime(0.0001, s);
        g.gain.exponentialRampToValueAtTime(0.22, s + 0.03);
        g.gain.exponentialRampToValueAtTime(0.0001, s + 0.5);
        o.start(s); o.stop(s + 0.52);
      });
    } catch (e) {}
  }

  function confetti() {
    var colors = ["#fbbf24", "#f472b6", "#34d399", "#60a5fa", "#f87171"];
    for (var n = 0; n < 60; n++) {
      var c = document.createElement("div");
      c.className = "confetti-piece";
      c.style.left = Math.random() * 100 + "%";
      c.style.background = colors[n % colors.length];
      c.style.animationDelay = (Math.random() * 0.6) + "s";
      c.style.transform = "rotate(" + (Math.random() * 360) + "deg)";
      el.doneOverlay.appendChild(c);
    }
  }

  // ---- controls ---------------------------------------------------------
  if (el.done) el.done.addEventListener("click", function () {
    var it = items[i];
    if (!paused && !finished && it && it.unit !== "seconds") { nextSide(it, curSide); }
  });
  // Full page load — Doors ALink can pushState to "/" before the server
  // reroutes, after which further taps no-op and the overlay never closes.
  if (el.backHome) {
    el.backHome.addEventListener("click", function (e) {
      e.preventDefault();
      location.assign("/");
    });
  }
  if (el.skip) el.skip.addEventListener("click", function () { if (paused || finished) return; shutUp(); clearNarr(); endItem(); });
  if (el.prev) el.prev.addEventListener("click", prev);
  if (el.restBack) el.restBack.addEventListener("click", prev);
  // Skip jumps straight to whatever the rest/ready countdown was about to start.
  if (el.restSkip) el.restSkip.addEventListener("click", function () {
    if (paused || finished) return;
    clearTicker();
    clearNarr();
    shutUp();
    var f = pendingStart; pendingStart = null;
    if (f) f();
  });

  // ---- in-workout difficulty nudge (this session only) ----
  function clampLevel(n) { return n < -3 ? -3 : (n > 3 ? 3 : n); }
  function updateDiff() {
    if (el.diff) el.diff.classList.toggle("hidden", fixedDifficulty);
    if (!el.diffLabel) return;
    var s = level > 0 ? "+" + level : (level < 0 ? "−" + (-level) : "0");
    el.diffLabel.textContent = (labels.difficulty || "") + " " + s;
    if (el.diffDown) el.diffDown.disabled = fixedDifficulty || level <= -3;
    if (el.diffUp) el.diffUp.disabled = fixedDifficulty || level >= 3;
  }
  function changeLevel(delta) {
    if (fixedDifficulty || paused || finished) return;
    var nl = clampLevel(level + delta);
    if (nl === level) return;
    level = nl;
    updateDiff();
    // Refresh the current rep target live; a running timer keeps its duration and
    // the new level applies from the next set on.
    var it = items[i];
    if (it && vis(el.stage) && it.unit !== "seconds") {
      el.value.textContent = activeValueDisplay(it);
    }
    checkpoint();
  }
  if (el.diffUp) el.diffUp.addEventListener("click", function () { changeLevel(1); });
  if (el.diffDown) el.diffDown.addEventListener("click", function () { changeLevel(-1); });
  updateDiff();

  // Exercise detail card: description, how-to, correct, mistakes, breathing,
  // warning, and a manual replace button. Opening it pauses the timer (reading
  // is a pull action, not part of the workout flow).
  var infoWasPaused = false;
  function setSec(sec, p, text) {
    if (!sec) return;
    if (text) { p.textContent = text; sec.classList.remove("hidden"); }
    else { sec.classList.add("hidden"); }
  }
  function cautionText(it) {
    var warning = (it.warning || "").trim();
    var safety = (it.safety || "").trim();
    if (warning && safety && warning !== safety) return warning + " " + safety;
    return warning || safety;
  }
  function fillList(list, arr) {
    if (!list) return;
    var sec = list.parentElement;
    list.innerHTML = "";
    if (!arr || !arr.length) { sec.classList.add("hidden"); return; }
    sec.classList.remove("hidden");
    for (var k = 0; k < arr.length; k++) {
      var li = document.createElement("li");
      li.textContent = arr[k];
      list.appendChild(li);
    }
  }
  function openInfo() {
    if (resumeNeeded || finished) return;
    var it = items[i];
    if (!it) return;
    infoWasPaused = paused;
    freezeClock();
    paused = true;
    clearNarr(); stopFrames(); shutUp();
    checkpoint();
    paintFrame(it, frameIdx);
    el.infoName.textContent = it.name;
    el.infoDesc.textContent = it.desc || "";
    fillList(el.infoHow, it.howTo);
    fillList(el.infoCorrect, it.correct);
    fillList(el.infoWrong, it.wrong);
    setSec(el.infoBreathingSec, el.infoBreathing, it.breathing);
    setSec(el.infoWarnSec, el.infoWarn, cautionText(it));
    if (el.infoReplace) {
      if (it.repl) el.infoReplace.classList.remove("hidden");
      else el.infoReplace.classList.add("hidden");
    }
    el.infoOverlay.classList.remove("hidden");
    if (!reducedMotion) playFrames(it, true);
  }
  function closeInfo() {
    stopFrames();
    el.infoOverlay.classList.add("hidden");
    if (resumeNeeded) return;
    paused = infoWasPaused;
    if (!paused) restorePhase();
    root.classList.toggle("frozen", paused);
  }
  // doReplace swaps the current exercise for its replacement and restarts it.
  function replacementItem(it) {
    var r = it.repl;
    return {
      id: r.id, unit: r.unit, value: r.value, scale: r.scale, slot: it.slot, perSide: r.perSide,
      round: it.round, rest: it.rest, sideRest: it.sideRest, name: r.name, hint: r.hint, warning: r.warning,
      svg: r.svg, frames: r.frames || (r.svg ? [r.svg] : []), animation: r.animation, desc: "", howTo: [], correct: [], wrong: [],
      breathing: "", safety: "", vStart: cues.go, vMid: [], vLast: cues.five || "",
      vFinish: cues.well_done, repl: null
    };
  }
  function doReplace() {
    if (resumeNeeded || finished) return;
    var it = items[i];
    if (!it || !it.repl) return;
    state.swaps[i] = it.repl.id;
    items[i] = replacementItem(it);
    el.infoOverlay.classList.add("hidden");
    paused = false;
    shutUp();
    startItem(i);
  }
  if (el.infoBtn) el.infoBtn.addEventListener("click", openInfo);
  if (el.visual || el.svg) (el.visual || el.svg).addEventListener("click", openInfo);
  if (el.infoClose) el.infoClose.addEventListener("click", closeInfo);
  if (el.infoReplace) el.infoReplace.addEventListener("click", doReplace);
  if (el.demoToggle) el.demoToggle.addEventListener("click", function () {
    if (!vis(el.infoOverlay)) return;
    if (frameTimer) stopFrames(); else playFrames(items[i], true);
  });
  function stepFrame(delta) {
    if (!vis(el.infoOverlay)) return;
    stopFrames();
    paintFrame(items[i], frameIdx + delta);
  }
  if (el.demoPrev) el.demoPrev.addEventListener("click", function () { stepFrame(-1); });
  if (el.demoNext) el.demoNext.addEventListener("click", function () { stepFrame(1); });

  // Static control labels
  if (el.prev) el.prev.textContent = t("prev");
  if (el.skip) el.skip.textContent = t("skip");
  if (el.done) el.done.textContent = t("done");
  if (el.restSkip) el.restSkip.textContent = t("skip");
  if (el.restBack) el.restBack.textContent = t("prev");

  // Controls are created here so older cached page markup can also recover.
  function addNode(parent, tag, cls, text) {
    var node = document.createElement(tag);
    node.className = cls;
    if (text) node.textContent = text;
    parent.appendChild(node);
    return node;
  }
  el.resumeOverlay = addNode(root, "div", "resume-overlay rest-overlay hidden");
  el.resumeOverlay.style.zIndex = "30";
  el.resumeOverlay.setAttribute("role", "dialog");
  el.resumeOverlay.setAttribute("aria-modal", "true");
  el.resumeOverlay.setAttribute("aria-label", t("resume_title"));
  addNode(el.resumeOverlay, "h2", "resume-title", t("resume_title"));
  el.resume = addNode(el.resumeOverlay, "button", "resume-continue btn primary", t("resume"));
  el.restart = addNode(el.resumeOverlay, "button", "resume-restart btn", t("restart"));
  el.resume.type = el.restart.type = "button";
  el.saveStatus = addNode(el.doneOverlay, "p", "save-status");
  el.saveStatus.setAttribute("role", "status");
  el.saveRetry = addNode(el.doneOverlay, "button", "save-retry btn hidden", t("save_retry"));
  el.saveRetry.type = "button";
  el.saveRetry.addEventListener("click", sendCompletion);

  function restorePhase() {
    var snapshot = Object.assign({}, state);
    i = snapshot.index; curSide = snapshot.side;
    restoring = true;
    if (snapshot.phase === "complete") { restoring = false; finish(); return; }
    if (snapshot.phase === "exercise") {
      startItem(i, curSide, snapshot.remainingMs);
    } else {
      var target = snapshot.nextIndex, side = snapshot.nextSide;
      doRest(Math.ceil(snapshot.remainingMs / 1000), items[target], snapshot.phase !== "rest", function () {
        startItem(target, side);
      }, snapshot.phase === "side-rest" ? side : 1, snapshot.remainingMs);
    }
    restoring = false;
    checkpoint();
  }
  function validSnapshot(s) {
    if (!s || !Number.isInteger(s.index) || s.index < 0 || s.index >= items.length ||
        !Number.isInteger(s.side) || s.side < 1 || s.side > 2 ||
        ["exercise", "ready", "rest", "side-rest", "complete"].indexOf(s.phase) < 0 ||
        typeof s.eventId !== "string" || s.eventId.length < 32 || s.eventId.length > 128 ||
        !Number.isInteger(s.level) || s.level < -3 || s.level > 3 ||
        !s.swaps || typeof s.swaps !== "object" || Array.isArray(s.swaps)) return false;
    var validSwaps = Object.keys(s.swaps).every(function (key) {
      var idx = Number(key), it = originalItems[idx];
      return String(idx) === key && Number.isInteger(idx) && it && it.repl && it.repl.id === s.swaps[key];
    });
    if (!validSwaps) return false;
    if (s.phase === "complete") return s.completion && s.completion.eventId === s.eventId &&
      s.completion.userId === payload.userId && s.completion.day === payload.day && s.completion.position === payload.position &&
      typeof s.completion.completedAt === "string" && Number.isFinite(Date.parse(s.completion.completedAt)) &&
      /^\d{4}-\d{2}-\d{2}$/.test(s.completion.completedOn);
    if (!Number.isInteger(s.nextIndex) || s.nextIndex < 0 || s.nextIndex >= items.length ||
        !Number.isInteger(s.nextSide) || s.nextSide < 1 || s.nextSide > 2) return false;
    var it = s.swaps[s.index] ? replacementItem(originalItems[s.index]) : originalItems[s.index];
    if (s.side > sides(it)) return false;
    if (s.phase === "side-rest" && (sides(it) !== 2 || s.nextIndex !== s.index || s.nextSide !== 2)) return false;
    if (s.phase !== "exercise" || it.unit === "seconds") {
      if (!Number.isFinite(s.remainingMs) || s.remainingMs < 0 || s.remainingMs > 86400000) return false;
    }
    return true;
  }
  el.resume.addEventListener("click", function () {
    if (disposed || finished || document.visibilityState !== "visible" || redirectAfterRestart()) return;
    resumeNeeded = false; paused = false;
    el.resumeOverlay.classList.add("hidden");
    root.classList.remove("frozen");
    restorePhase();
    requestWake();
  });
  el.restart.addEventListener("click", function () {
    if (disposed || finished || redirectAfterRestart()) return;
    if (store) store.clearCheckpoint(payload);
    items = originalItems.slice();
    i = 0; curSide = 1;
    level = fixedDifficulty ? 0 : (payload.level || 0);
    state = { index: 0, side: 1, phase: "ready", remainingMs: 3000, nextIndex: 0, nextSide: 1,
      level: level, swaps: {}, eventId: store ? store.newEventId() : state.eventId, completion: null };
    paused = false; resumeNeeded = false;
    updateDiff();
    el.resumeOverlay.classList.add("hidden");
    root.classList.remove("frozen");
    doRest(3, items[0], true, function () { startItem(0); });
    requestWake();
  });
  var unsubscribe = store ? store.subscribe(function (event) {
    if (event.userId === payload.userId && event.eventId === state.eventId) showCompletionStatus(event.status, event.receipt, event.durable);
  }) : function () {};

  // ---- go ---------------------------------------------------------------
  $sys.clean(function () {
    if (!finished) { freezeClock(); checkpoint(); }
    disposed = true;
    clearTicker(); stopFrames(); clearNarr(); shutUp(); releaseWake();
    unsubscribe();
    document.removeEventListener("visibilitychange", onVisibility);
    window.removeEventListener("pagehide", pauseForReturn);
    window.removeEventListener("pageshow", onPageShow);
    window.removeEventListener("focus", onPageShow);
    if (window.speechSynthesis && window.speechSynthesis.onvoiceschanged === loadVoices) window.speechSynthesis.onvoiceschanged = null;
  });

  if (redirectAfterRestart()) return;
  buildBar();
  var saved = store ? store.loadCheckpoint(payload) : null;
  if (saved && validSnapshot(saved)) {
    state = saved;
    Object.keys(state.swaps || {}).forEach(function (key) { items[key] = replacementItem(originalItems[key]); });
    i = state.index; curSide = state.side;
    level = fixedDifficulty ? 0 : state.level;
    updateDiff();
    if (state.phase === "complete") finish();
    else {
      paused = true; resumeNeeded = true;
      restorePhase();
      root.classList.add("frozen");
      el.resumeOverlay.classList.remove("hidden");
    }
  } else if (items.length === 0) {
    finish();
  } else {
    if (saved && store) store.clearCheckpoint(payload);
    doRest(3, items[0], true, function () { startItem(0); });
    if (document.visibilityState !== "visible") pauseForReturn();
    else requestWake();
  }
})();
