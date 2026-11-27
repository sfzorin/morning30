// Keeps the AI coach conversation scrolled to its newest message, and empties
// the message field once the sent message shows up in the conversation. The
// field lives outside the part of the page that re-renders, and the framework
// handles the form's submit itself, so the thread is watched instead.
(function () {
  function scrollDown() {
    var thread = document.querySelector(".coach-thread");
    var last = thread && thread.lastElementChild;
    if (last && last.tagName === "SCRIPT") last = last.previousElementSibling;
    if (last) last.scrollIntoView({ block: "end" });
    window.scrollTo(0, document.documentElement.scrollHeight);
  }
  var seen = document.querySelectorAll(".coach-thread .bubble.me").length;
  // Clears the field only when one more own message has just appeared.
  function clearIfSent() {
    var mine = document.querySelectorAll(".coach-thread .bubble.me");
    var grew = mine.length > seen;
    seen = mine.length;
    if (!grew) return;
    var input = document.querySelector(".coach-input input[name=text]");
    var last = mine[mine.length - 1].textContent.trim();
    if (input && last && input.value.trim() === last) input.value = "";
  }
  scrollDown();
  var page = document.querySelector("main.coach");
  if (!page || page.dataset.watch) return;
  page.dataset.watch = "1";
  new MutationObserver(function () {
    clearIfSent();
    scrollDown();
  }).observe(page, { childList: true, subtree: true });
})();
