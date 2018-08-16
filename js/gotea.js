// var serialize = require("form-serialize");
import morphdom from "morphdom";
import serialize from "form-serialize";

// Websockets

var socket = new WebSocket(
  (window.location.protocol === "https:" ? "wss://" : "ws://") + window.location.host + "/server"
);

socket.onopen = function () {
  let path = window.location.pathname;
  if (path != "/") {
    changeRoute(path)
  }
}

socket.onmessage = function (event) {
  swapDOM(event.data, "view");
};

// DOM Swap with Morphdom
// with the option 'childrenOnly' Morphom swaps children of the containers,
// therefore we place the incoming HTML in a div
const swapDOM = (incomingHTML, containerID) => {
  var el1 = document.createElement('div');
  el1.innerHTML = incomingHTML;
  morphdom(document.getElementById(containerID), el1, {
    "childrenOnly": true
  });
}


const sendMessage = (msgString, args) => {
  let msg = {
    "message": msgString,
    "args": JSON.parse(args)
  }
  console.log("Sending websocket message: ", msg);
  socket.send(JSON.stringify(msg));
}

function submitForm(message, formID) {
  let form = document.getElementById(formID);
  let obj = serialize(form, {
    hash: true
  });

  let msg = {
    "message": message,
    "args": obj
  }

  console.log("Sending websocket message: ", msg);
  socket.send(JSON.stringify(msg));
}

const sendMessageDebounce = debounce(sendMessage, 200);
const submitFormDebounce = debounce(submitForm, 200);


window.gotea = {};
window.gotea.sendMessage = sendMessageDebounce;
window.gotea.submitForm = submitFormDebounce;



function debounce(func, wait, immediate) {
  var timeout;
  return function () {
    var context = this,
      args = arguments;
    var later = function () {
      timeout = null;
      if (!immediate) func.apply(context, args);
    };
    var callNow = immediate && !timeout;
    clearTimeout(timeout);
    timeout = setTimeout(later, wait);
    if (callNow) func.apply(context, args);
  };
};


// function sendMessage(element) {
//   let _message = element.target.dataset.msg;
//   let message = {};
//   if (_message) {
//     message = JSON.parse(_message);
//   }
//   console.log("Sending websocket message: ", message);
//   socket.send(JSON.stringify(message));
// }

// document.addEventListener(
//   "click",
//   function (e) {
//     if (/gotea-click/.test(e.target.className)) {
//       sendMessage(e);
//     }
//   },
//   false
// );

// function sendFormSubmitMessage(element) {
//   let _message = element.target.dataset.msg;
//   let message = {};

//   let _form = element.target.parentNode;
//   let form = {};
//   if (_form) {
//     form = serialize(_form, { hash: true });
//   }

//   if (_message) {
//     message = JSON.parse(_message);
//     message.arguments = form;
//   }

//   console.log("Sending websocket message: ", message);
//   socket.send(JSON.stringify(message));
// }

// document.addEventListener(
//   "click",
//   function (e) {
//     if (/gotea-form-submit/.test(e.target.className)) {
//       sendFormSubmitMessage(e);
//     }
//   },
//   false
// );

function changeRoute(route) {
  history.pushState({}, "", route);

  let msg = {
    "message": "CHANGE_ROUTE",
    "args": route
  }

  console.log("Sending websocket message: ", msg);
  socket.send(JSON.stringify(msg));
}

document.addEventListener(
  "click",
  function (e) {
    if (/gotea-link/.test(e.target.className)) {
      e.preventDefault();
      changeRoute(e.target.getAttribute("href"));
      return false;
    }
  },
  false
);