$(function() {
  $("#Activate").click(function() {
    $.getJSON("/event/activate", $(document.forms["activate"]).serialize())
    .then(function(data) {
      console.log("Command Controller", data.NumListener + " core(s) started.");
    });
  });
});
