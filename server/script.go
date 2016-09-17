package server

import "net/http"

func script(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/javascript")

	w.Write(javaScript)
}

var javaScript = []byte(`// generate.name script

var generate = function(name, count, successHandler, errorHandler) {
  getJSON("//generate.name/"+ name + ".json?n="+count, successHandler, errorHandler);
}

var getJSON = function(url, successHandler, errorHandler) {
  var xhr = typeof XMLHttpRequest != 'undefined'
    ? new XMLHttpRequest()
    : new ActiveXObject('Microsoft.XMLHTTP');
  xhr.open('get', url, true);
  xhr.responseType = 'json';
  xhr.onreadystatechange = function() {
    var status;
    var data;

    if (xhr.readyState == 4) {
      status = xhr.status;
      if (status == 200) {
        successHandler && successHandler(xhr.response);
      } else {
        errorHandler && errorHandler(status);
      }
    }
  };
  xhr.send();
};`)
