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
