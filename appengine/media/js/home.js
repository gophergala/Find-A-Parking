$(document).on("ready",init);

var token = "";
var myMap = {};

var callback = function () {
    
    token = $("#token").text();
	channel = new goog.appengine.Channel(token);
    socket = channel.open();
    socket.onopen = socketOpening;
    socket.onmessage = socketMessage;
    socket.onerror = socketError;
    socket.onclose = socketClose;
    //marker = initMap.addMarker(25.670708,-100.308172, "Hello Map World");
    //var contentString = '<h1>Mi primer infoWindow con dmaps</h1>';
    //marker.addInfo(contentString);

}

function socketOpening (e) {
	console.log("opening");
}

function socketMessage (message) {
	console.log(message);
}

function socketError (e) {
	console.log("error");
}

function socketClose (e) {
	console.log("close");
}

function init () {
	myMap = new DMaps("map",25.670708,-100.308172,callback);
}

