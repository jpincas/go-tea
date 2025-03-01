import morphdom from "morphdom";

// Constants
const SOCKET_MESSAGE = "Sent Message:";

// Websockets
console.log("Attempting to establish WebSocket connection");
const socket = new WebSocket(
  `${window.location.protocol === "https:" ? "wss://" : "ws://"}${window.location.host}/server?whence=${document.location.pathname}`
);

// Handle incoming messages from the server
socket.onmessage = event => {
  console.log("Received rerender from server");
  morphdom(document.documentElement, event.data, {
    childrenOnly: true
  });
};

// Handle WebSocket open event
socket.onopen = () => {
  console.log("WebSocket connection established.");
};

// Handle WebSocket error event
socket.onerror = error => {
  console.error("WebSocket error:", error);
};

// Handle WebSocket close event
socket.onclose = event => {
  if (event.wasClean) {
    console.log(`WebSocket connection closed cleanly, code=${event.code}, reason=${event.reason}`);
  } else {
    console.error("WebSocket connection closed unexpectedly, code=", event.code, "reason=", event.reason);
  }
};

// Send a message through the websocket
const sendMessage = (msg) => {
  const msgJsonString = JSON.stringify(msg);
  console.log(`${SOCKET_MESSAGE}`, msgJsonString);
  socket.send(msgJsonString);
};

// Send a message with a value from an input field
const sendMessageWithValueFromInput = (msg, inputID) => {
  msg.args = document.getElementById(inputID).value;

  const msgJsonString = JSON.stringify(msg);
  console.log(`${SOCKET_MESSAGE}`, msgJsonString);
  socket.send(msgJsonString);
};

const sendMessageWithValueFromThisInput = (msg) => {
  msg.args = document.activeElement.value;

  const msgJsonString = JSON.stringify(msg);
  console.log(`${SOCKET_MESSAGE}`, msgJsonString);
  socket.send(msgJsonString);
};

// Submit a form through the websocket
const updateFormState = (msg, formID) => {
  msg.args = serializeForm(formID);

  console.log(`${SOCKET_MESSAGE}`, msg);
  socket.send(JSON.stringify(msg));
};


// Serialize form data into an object
const serializeForm = formID => {
  const formElements = [...document.getElementById(formID).elements];
  const TEXT_INPUT = "text";
  const CHECKBOX = "checkbox";
  const RADIO = "radio";
  const SELECT = "SELECT";
  const TEXTAREA = "TEXTAREA";

  const buildSelectArray = select => [...select.children]
    .map(option => (option.selected ? option.value : ""))
    .filter(value => value.length > 0);

  const handleSelect = select =>
    select.multiple ? buildSelectArray(select) : select.value;

  return formElements.reduce((acc, el) => {
    switch (el.tagName) {
      case SELECT:
        acc[el.name] = handleSelect(el);
        break;
      case TEXTAREA:
        acc[el.name] = el.value;
        break;
    }
    switch (el.type) {
      case TEXT_INPUT:
        acc[el.name] = el.value;
        break;
      case CHECKBOX:
        acc[el.name] = el.checked;
        break;
      case RADIO:
        if (el.checked) {
          acc[el.name] = el.value;
        }
        break;
    }
    return acc;
  }, {});
};

// Change the route and notify the server
const changeRoute = route => {
  history.pushState({}, "", route);
  const msg = {
    message: "CHANGE_ROUTE",
    args: route
  };
  console.log(`${SOCKET_MESSAGE}`, msg);
  socket.send(JSON.stringify(msg));
};

// Expose functions to the global window object
window.gotea = {
  sendMessage,
  updateFormState,
  sendMessageWithValueFromInput,
  sendMessageWithValueFromThisInput
};

// Handle browser back/forward navigation
window.addEventListener('popstate', event => {
  const msg = {
    message: "CHANGE_ROUTE",
    args: document.location.pathname,
  };
  console.log(`${SOCKET_MESSAGE}`, msg);
  socket.send(JSON.stringify(msg));
});

// Intercept link clicks and handle routing
document.addEventListener(
  "click",
  e => {
    let target = e.target;
    while (target && target.tagName !== 'A') {
      target = target.parentElement;
    }
    if (target && !/external/.test(target.className)) {
      e.preventDefault();
      changeRoute(target.getAttribute("href"));
      return false;
    }
  },
  false
);
