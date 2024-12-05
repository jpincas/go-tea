import morphdom from "morphdom";


// Websockets
const socket = new WebSocket(
  (window.location.protocol === "https:" ? "wss://" : "ws://") + window.location.host + "/server?whence=" + document.location.pathname
);

// socket.onopen = () => {
//   const path = window.location.pathname;
//   if (path != "/") {
//     changeRoute(path);
//   }
// };

socket.onmessage = event => {
  morphdom(document.documentElement, event.data, {
    childrenOnly: true
  });
};

// DOM Swap with Morphdom
// with the option 'childrenOnly' Morphom swaps children of the containers,
// therefore we place the incoming HTML in a div
// const swapDOM = (incomingHTML) => {
//   // const el1 = document.createElement("div");
//   // el1.innerHTML = incomingHTML;
//   morphdom(document.getElementById(containerID), incomingHTML, {
//     childrenOnly: true
//   });
// };

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

function sendMessageWithValue(message, inputID) {
  let value = document.getElementById(inputID).value;

  let msg = {
    message,
    args: value
  };

  console.log("Sending websocket message: ", msg);
  socket.send(JSON.stringify(msg));
}

window.gotea = {
  sendMessage: sendMessage,
  submitForm: submitForm,
  sendMessageWithValue: sendMessageWithValue
};

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
  const buildSelectArray = select => [...select.children]
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
        acc[el.name] = el.value;
        break;
      case CHECKBOX:
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

function changeRoute(route) {
  history.pushState({}, "", route);

  const msg = {
    message: "CHANGE_ROUTE",
    args: route
  };

  console.log("Sending websocket message: ", msg);
  socket.send(JSON.stringify(msg));
}


// For some reason only work when place on Window
window.addEventListener('popstate', function (event) {
  const msg = {
    message: "CHANGE_ROUTE",
    args: document.location.pathname,
  };

  console.log("Sending websocket message: ", msg);
  socket.send(JSON.stringify(msg));
});

document.addEventListener(
  "click",
  e => {
    if (e.target.tagName == 'A' && /external/.test(e.target.className) == false) {
      e.preventDefault();
      changeRoute(e.target.getAttribute("href"));
      return false;
    }
  },
  false
);

export default {
  sendMessage,
  submitForm,
  sendMessageWithValue,
  serializeForm,
  changeRoute
};