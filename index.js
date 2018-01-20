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
  const message = element.srcElement.dataset.message;
  var tag = element.srcElement.dataset.tag;
  const data = {};
  if (tag) {
    data = JSON.parse(element.srcElement.dataset.tag);
  }

  const payload = { message, data };
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
