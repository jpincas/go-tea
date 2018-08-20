// var serialize = require("form-serialize");
import morphdom from "morphdom";
import serialize from "form-serialize";

// Websockets

const socket = new WebSocket(
  window.location.protocol === "https:"
    ? "wss://"
    : "ws://" + window.location.host + "/server"
);

socket.onopen = () => {
  const path = window.location.pathname;
  if (path != "/") {
    changeRoute(path);
  }
};

socket.onmessage = event => {
  swapDOM(event.data, "view");
};

// DOM Swap with Morphdom
// with the option 'childrenOnly' Morphom swaps children of the containers,
// therefore we place the incoming HTML in a div
const swapDOM = (incomingHTML, containerID) => {
  const el1 = document.createElement("div");
  el1.innerHTML = incomingHTML;
  morphdom(document.getElementById(containerID), el1, {
    childrenOnly: true
  });
};

const sendMessage = (msgString, args) => {
  const msg = {
    message: msgString,
    args: JSON.parse(args)
  };
  console.log("Sending websocket message: ", msg);
  socket.send(JSON.stringify(msg));
};

function submitForm(message, formID) {
  const msg = {
    message,
    args: serializeForm(formID)
  };

  console.log("Sending websocket message: ", msg);
  socket.send(JSON.stringify(msg));
}

window.gotea = {
  sendMessage: debounce(sendMessage, 200),
  submitForm: debounce(submitForm, 200)
};

function debounce(func, wait, immediate) {
  let timeout;
  return function() {
    const later = () => {
      timeout = null;
      if (!immediate) func.apply(this, arguments);
    };
    const callNow = immediate && !timeout;
    clearTimeout(timeout);
    timeout = setTimeout(later, wait);
    if (callNow) func.apply(this, arguments);
  };
}

const serializeForm = formID => {
  const formElements = [...document.getElementById(formID).elements];
  const TEXT_INPUT = "text";
  const CHECKBOX = "checkbox";
  const RADIO = "radio";
  const SELECT = "SELECT"; // for some reason tagName returns uppercase name
  const TEXTAREA = "TEXTAREA";

  // map over array of options (elements children). If option is selected, return it's value
  // otherwise return empty string and then filter out empty strings from the array so we end up
  // with array of selected options.
  const buildSelectArray = select =>
    [...select.children]
      .map(option => (option.selected ? option.value : ""))
      .filter(value => value.length > 0);

  // if select has multiple attribute, build array of selected options,
  // otherwise, for dropdown for example, return select's value
  const handleSelect = select =>
    select.multiple ? buildSelectArray(select) : select.value;

  // reduce over form elements and check against their type (for inputs) and tagName for other tags
  // (textarea/select etc). Return data shaped specifically for structs on server-side, so they are easy
  // to decode
  return formElements.reduce((acc, el) => {
    switch (el.tagName) {
      case SELECT:
        // select may have multiple values, check if it is a case and return proper value (array or string)
        acc[el.name] = handleSelect(el);
        break;
      case TEXTAREA:
        // simple string to field name assignment
        acc[el.name] = el.value;
        break;
    }
    switch (el.type) {
      // simple string to field name assignment
      case TEXT_INPUT:
        alert("czapa");
        acc[el.name] = el.value;
        break;
      case CHECKBOX:
        alert("tiu");
        // see if checkbox is checked and assign bool to it's name.
        acc[el.name] = el.checked;
        break;
      case RADIO:
        // if radio is checked, assign it's value (string) to it's name
        if (el.checked) {
          acc[el.name] = el.value;
        }
        break;
    }
    return acc;
  }, {});
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

  const msg = {
    message: "CHANGE_ROUTE",
    args: route
  };

  console.log("Sending websocket message: ", msg);
  socket.send(JSON.stringify(msg));
}

document.addEventListener(
  "click",
  e => {
    if (/gotea-link/.test(e.target.className)) {
      e.preventDefault();
      changeRoute(e.target.getAttribute("href"));
      return false;
    }
  },
  false
);
