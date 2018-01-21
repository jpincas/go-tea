var snabbdom = require("snabbdom");
var patch = snabbdom.init([
  // Init patch function with chosen modules
  require("snabbdom/modules/class").default, // makes it easy to toggle classes
  require("snabbdom/modules/props").default, // for setting properties on DOM elements
  require("snabbdom/modules/style").default, // handles styling on elements with support for animations
  //require("snabbdom/modules/eventlisteners").default, // attaches event listeners
  require("snabbdom/modules/attributes").default // for dataset attributes
]);
var h = require("snabbdom/h").default; // helper function for creating vnodes
var toVNode = require("snabbdom/tovnode").default;

var socket = new WebSocket(
  (window.location.protocol === "https:" ? "wss://" : "ws://") +
    window.location.host +
    "/server"
);

var container;
var oldNode;

socket.onmessage = function(event) {
  var el = document.createElement("div");
  el.innerHTML = event.data;
  el.setAttribute("id", "view");
  newNode = toVNode(el);

  // For the first render
  // patch the empty container
  // for subsequent renders, use the old node
  if (!container) {
    container = document.getElementById("view");
    patch(container, newNode);
    oldNode = newNode;
  } else {
    patch(oldNode, newNode);
    oldNode = newNode;
  }
};

function sendMessage(element) {
  const message = element.target.dataset.message;
  var _tag = element.target.dataset.tag;
  let tag = {};
  if (_tag) {
    tag = JSON.parse(_tag);
  }

  const payload = { message, tag };
  console.log("Sending websocket message: ", payload);
  socket.send(JSON.stringify(payload));
}

document.addEventListener(
  "click",
  function(e) {
    if (/gotea-click/.test(e.target.className)) {
      sendMessage(e);
    }
  },
  false
);

/*
 * This is obviously similar to the above, however I think we have a real
 * problem here, presumably we might have some elements to which we wish to
 * attach multiple messages for different kinds of events. Eg. we might respond
 * different to a focus event as to an input event, or a click event as to a
 * hover event. Fortunately I think we can parameterise this somewhat by just
 * including the name of the event in the name of the attribute.
 */

function onInputSendMessage(event) {
  const message = event.target.dataset.message;
  let tag = { "input": event.target.value };
  console.log ("The value is: ", event.target.value);
  const payload = { message, tag };
  console.log("Sending websocket message in onInputSendMessage: ", payload);
  socket.send(JSON.stringify(payload));
}

document.addEventListener(
  "input",
  function(e){
    if(/gotea-input/.test(e.target.className)){
      onInputSendMessage(e);
    }
  },
  false
);
