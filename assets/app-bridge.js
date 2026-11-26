// Talks to the morning30 apps for iPhone and Android. The site is the same
// everywhere; inside an app it shows the app-only rows (the morning reminder)
// and hands their values to the phone, which schedules the notification.
//
// iOS: the app injects window.morning30App = {platform, reminder} before the
// page loads and listens on webkit.messageHandlers.morning30.
// Android: the app opens the site with ?app=android&reminder=HH:MM; the values
// are kept in localStorage, and changes go back through an intent:// link.
(function () {
  var KEY = "m30.app";
  var data = {};
  try {
    data = $data("reminder") || {};
    if (typeof data === "string") data = JSON.parse(data);
  } catch (_) {}

  function load() {
    var q = new URLSearchParams(window.location.search);
    if (q.get("app")) {
      save({ platform: q.get("app"), reminder: q.get("reminder") || "" });
    }
    if (window.morning30App && window.morning30App.platform) return window.morning30App;
    try { return JSON.parse(window.localStorage.getItem(KEY) || "null"); } catch (_) { return null; }
  }
  function save(app) {
    try { window.localStorage.setItem(KEY, JSON.stringify(app)); } catch (_) {}
  }
  // ---- Finished workouts → Apple Health (iOS) / Health Connect (Android) ----
  // player.js fires "m30:workout" when a workout ends. Each workout is sent once.
  var SENT = "m30.health.sent", PENDING = "m30.health.pending";
  function readList(key) {
    try { var v = JSON.parse(window.localStorage.getItem(key) || "[]"); return Array.isArray(v) ? v : []; } catch (_) { return []; }
  }
  function writeList(key, list) {
    try { window.localStorage.setItem(key, JSON.stringify(list.slice(-20))); } catch (_) {}
  }
  function workoutOf(d) {
    var end = Number(d && d.end);
    if (!Number.isFinite(end)) end = Date.now();
    var start = Number(d && d.start);
    // An interrupted and resumed workout can span hours; fall back to its planned length.
    if (!Number.isFinite(start) || start >= end || end - start > 3 * 3600 * 1000) {
      start = end - Math.max(60, Number(d && d.estSec) || 900) * 1000;
    }
    return { id: String(d && d.eventId || end), start: Math.round(start), end: Math.round(end),
      title: String(d && d.title || "").slice(0, 80), kind: d && d.kind === "stretch" ? "stretch" : "strength" };
  }
  function healthURL(w) {
    return "intent://health?start=" + w.start + "&end=" + w.end + "&kind=" + w.kind + "&title=" + encodeURIComponent(w.title) +
      "#Intent;scheme=morning30;package=" + (data.androidPackage || "com.morning30.app") + ";end";
  }
  function listenForWorkouts(app) {
    if (window.__m30Health) return; // the bridge can load more than once
    window.__m30Health = true;
    // Android opens the app through an intent:// link, which Chrome allows only
    // right after a tap: send now if the workout ended with a tap, else on the next one.
    function flush(ev) {
      var pending = readList(PENDING);
      if (!pending.length) return;
      writeList(PENDING, []);
      var link = ev && ev.target && ev.target.closest ? ev.target.closest("a[href]") : null;
      if (link && ev.type === "click") {
        ev.preventDefault();
        setTimeout(function () { window.location.href = link.href; }, 700);
      }
      window.location.href = healthURL(pending[pending.length - 1]);
    }
    if (app.platform === "android") document.addEventListener("click", flush, true);
    document.addEventListener("m30:workout", function (e) {
      var w = workoutOf(e.detail);
      if (w.end - w.start < 60000) return; // skipped through: not a workout worth recording
      var sent = readList(SENT);
      if (sent.indexOf(w.id) >= 0) return;
      sent.push(w.id);
      writeList(SENT, sent);
      if (app.platform === "ios" && window.webkit && window.webkit.messageHandlers && window.webkit.messageHandlers.morning30) {
        window.webkit.messageHandlers.morning30.postMessage({ type: "workout", start: w.start, end: w.end, title: w.title, kind: w.kind });
        return;
      }
      if (app.platform === "android") {
        var pending = readList(PENDING);
        pending.push(w);
        writeList(PENDING, pending);
        if (navigator.userActivation && navigator.userActivation.isActive) flush(null);
      }
    });
  }

  function send(app, time) {
    var msg = { type: "reminder", time: time, title: data.title || "", body: data.body || "" };
    if (app.platform === "ios" && window.webkit && window.webkit.messageHandlers && window.webkit.messageHandlers.morning30) {
      window.webkit.messageHandlers.morning30.postMessage(msg);
      return;
    }
    if (app.platform === "android") {
      var query = "time=" + encodeURIComponent(time) + "&title=" + encodeURIComponent(msg.title) + "&body=" + encodeURIComponent(msg.body);
      window.location.href = "intent://reminder?" + query + "#Intent;scheme=morning30;package=" + (data.androidPackage || "com.morning30.app") + ";end";
    }
  }

  var app = load();
  if (!app) return;
  listenForWorkouts(app);
  var row = document.querySelector("[data-app-reminder]");
  if (!row) return;
  var select = row.querySelector("select");
  row.classList.remove("hidden");
  select.value = app.reminder || "";
  select.addEventListener("change", function () {
    app.reminder = select.value;
    if (window.morning30App) window.morning30App.reminder = select.value;
    save(app);
    send(app, select.value);
  });
})();
