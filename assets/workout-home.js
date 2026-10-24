
// The dashboard offers the last interrupted workout and sends any result that
// was queued before this page (or the standalone app) was closed.
(function () {
  var config;
  try {
    config = $data("resume");
    if (typeof config === "string") config = JSON.parse(config);
  } catch (_) { return; }
  var state = window.MorningWorkoutState;
  if (!state || !config) return;
  var link = document.querySelector(".resume-workout");
  var status = document.querySelector(".home-save-status");
  var disposed = false;
  var reloading = false;
  var hasSaved = false;
  var reloadTimer = null;
  var openedOn = new Date().toDateString();
  var restartAt = Date.parse(config.restartAt);
  function reload() {
    if (disposed || reloading) return;
    reloading = true;
    window.location.reload();
  }
  function refreshIfDue() {
    if (disposed || document.visibilityState !== "visible") return;
    if (new Date().toDateString() !== openedOn || (Number.isFinite(restartAt) && Date.now() >= restartAt)) {
      reload();
    }
  }
  function update() {
    var resume = state.getResume(config.userId, config.position, config.programKey);
    if (link) {
      link.classList.toggle("hidden", !resume);
      if (resume) link.href = resume.url;
    }
    if (status) {
      var pending = state.pending(config.userId).length > 0;
      status.classList.toggle("hidden", !pending && !hasSaved);
      if (pending || hasSaved) status.textContent = status.getAttribute(pending ? "data-pending" : "data-saved");
    }
  }
  var unsubscribe = state.subscribe(function (event) {
    if (disposed || event.userId !== config.userId) return;
    if (event.status === "saved") hasSaved = true;
    update();
    if (event.status === "saved") {
      // Refresh the server-rendered calendar once its progress is out of date.
      if (event.receipt && (event.receipt.position !== config.position || event.receipt.streak !== config.streak) && reloadTimer === null) {
        reloadTimer = setTimeout(reload, 250);
      }
    }
  });
  update();
  state.start(config.userId);
  window.addEventListener("pageshow", refreshIfDue);
  document.addEventListener("visibilitychange", refreshIfDue);
  var dateTimer = setInterval(refreshIfDue, 60000);
  refreshIfDue();
  $sys.clean(function () {
    disposed = true;
    unsubscribe();
    if (reloadTimer !== null) clearTimeout(reloadTimer);
    clearInterval(dateTimer);
    window.removeEventListener("pageshow", refreshIfDue);
    document.removeEventListener("visibilitychange", refreshIfDue);
  });
})();
