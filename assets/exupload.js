// Reads a picked .json file into the nearest import textarea so the form
// submits text (no multipart). Direct paste still works without this script.
(function () {
  document.querySelectorAll(".ex-import-form").forEach(function (form) {
    var file = form.querySelector("input[type=file]");
    var ta = form.querySelector("textarea");
    if (!file || !ta) return;
    file.addEventListener("change", function () {
      var f = file.files && file.files[0];
      if (!f) return;
      var rd = new FileReader();
      rd.onload = function () { ta.value = String(rd.result || ""); };
      rd.readAsText(f);
    });
  });
})();
