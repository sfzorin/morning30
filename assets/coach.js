// Keeps the AI coach conversation scrolled to its newest message.
(function () {
  var thread = document.querySelector(".coach-thread");
  var last = thread && thread.lastElementChild;
  if (last && last.tagName === "SCRIPT") last = last.previousElementSibling;
  if (last) last.scrollIntoView({ block: "end" });
  window.scrollTo(0, document.documentElement.scrollHeight);
})();
