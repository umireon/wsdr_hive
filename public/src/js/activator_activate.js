const jQuery = require('jquery');
const toastr = require('toastr');
require('toastr/toastr.less');
var submit = document.querySelector('input[type=submit]');
submit.addEventListener('click', function(e) {
  e.preventDefault();
  fetch(e.target.form.action, {
    method: 'POST',
    body: new FormData(e.target.form),
  }).then(function(res) {
    return res.json();
  }).then(function(json) {
    toastr.info("Activated!");
  });
});
