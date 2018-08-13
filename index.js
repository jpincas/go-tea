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

var serialize = require("form-serialize");

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
  let _message = element.target.dataset.msg;
  let message = {};
  if (_message) {
    message = JSON.parse(_message);
  }
  console.log("Sending websocket message: ", message);
  socket.send(JSON.stringify(message));
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

function sendFormSubmitMessage(element) {
  let _message = element.target.dataset.msg;
  let message = {};

  let _form = element.target.parentNode;
  let form = {};
  if (_form) {
    form = serialize(_form, { hash: true });
  }

  if (_message) {
    message = JSON.parse(_message);
    message.arguments = form;
  }

  console.log("Sending websocket message: ", message);
  socket.send(JSON.stringify(message));
}

document.addEventListener(
  "click",
  function(e) {
    if (/gotea-form-submit/.test(e.target.className)) {
      sendFormSubmitMessage(e);
    }
  },
  false
);

function sendRouteChangeMessage(element) {
  let rel = element.target.getAttribute("href");
  // update the browser url
  history.pushState({}, "", rel);

  let message = { func: "ChangeRoute", arguments: rel };

  console.log("Sending websocket message: ", message);
  socket.send(JSON.stringify(message));
}

document.addEventListener(
  "click",
  function(e) {
    if (/gotea-link/.test(e.target.className)) {
      sendRouteChangeMessage(e);
      e.preventDefault();
      return false;
    }
  },
  false
);
