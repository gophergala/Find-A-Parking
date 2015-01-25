$(document).on("ready",init);

var token = "";
var myMap = {};
var data = {};

var callback = function () {
    
    token = $("#token").text();
	openChannel();
    myMap.setMapStyle('APPLE');
    //

}

function openChannel () {
	channel = new goog.appengine.Channel(token);
    socket = channel.open();
    socket.onopen = socketOpening;
    socket.onmessage = socketMessage;
    socket.onerror = socketError;
    socket.onclose = socketClose;
}

function socketOpening (e) {
	console.log("opening");
}

function socketMessage (message) {
	data = message;

	parking = JSON.parse(data.data);
	console.log(parking);
	var marker = myMap.addMarker(parking.Location.Lat,parking.Location.Lng, parking.Owner);
    var contentString = '<h2>Owner: '+parking.Owner+'</h2><h2>Contact: '+parking.Mail+'</h2><h2>Price:'+parking.Price+'</h2>';
    marker.addInfo(contentString);
}

function socketError (e) {
	console.log("error");
}

function socketClose (e) {
	
}

function init () {
	myMap = new DMaps("map",25.670708,-100.308172,callback);
}

function getToken () {
	$.ajax({
	  type: "POST",
	  url: "getToken",
	  dataType: "json",
		success: function  (result) {
			token = result.token;
			openChannel();
		},
	});
}

