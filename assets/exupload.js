// Reads a picked .json file into the import textarea so the form submits text
// (no multipart). Direct paste into the textarea works without this script.
(function () {
  var file = document.querySelector('.ex-file');
  var ta = document.querySelector('.ex-import');
  if (!file || !ta) return;
  file.addEventListener('change', function () {
    var f = file.files && file.files[0];
    if (!f) return;
    var rd = new FileReader();
    rd.onload = function () { ta.value = String(rd.result || ''); };
    rd.readAsText(f);
  });
})();
