import morphdom from "morphdom";

// Constants
const SOCKET_MESSAGE = "Sent Message:";

// Websockets
const socket = new WebSocket(
  `${window.location.protocol === "https:" ? "wss://" : "ws://"}${window.location.host}/server?whence=${document.location.pathname}`
);

// Handle incoming messages from the server
socket.onmessage = event => {
  morphdom(document.documentElement, event.data, {
    childrenOnly: true
  });
};

// Send a message through the websocket
const sendMessage = (message, args) => {
  const msg = {
    message,
    args: JSON.parse(args)
  };
  console.log(`${SOCKET_MESSAGE}`, msg);
  socket.send(JSON.stringify(msg));
};

// Submit a form through the websocket
const submitForm = (message, formID) => {
  const msg = {
    message,
    args: serializeForm(formID)
  };
  console.log(`${SOCKET_MESSAGE}`, msg);
  socket.send(JSON.stringify(msg));
};

// Send a message with a value from an input field
const sendMessageWithValue = (message, inputID) => {
  const value = document.getElementById(inputID).value;
  const msg = {
    message,
    args: value
  };
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
  submitForm,
  sendMessageWithValue
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
    if (e.target.tagName === 'A' && !/external/.test(e.target.className)) {
      e.preventDefault();
      changeRoute(e.target.getAttribute("href"));
      return false;
    }
  },
  false
);
