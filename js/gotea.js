import morphdom from "morphdom";

// Constants
const SOCKET_MESSAGE = "Sent Message:";
const INITIAL_RECONNECT_DELAY = 1000;  // 1 second
const MAX_RECONNECT_DELAY = 30000;     // 30 seconds
const RECONNECT_BACKOFF_MULTIPLIER = 2;

// Helpers for state persistence
function getCookie(name) {
  const value = `; ${document.cookie}`;
  const parts = value.split(`; ${name}=`);
  if (parts.length === 2) return parts.pop().split(';').shift();
}

function storeState(sessionId, stateData) {
  try {
    localStorage.setItem(`gotea_state_${sessionId}`, stateData);
  } catch (e) {
    console.warn('Failed to store state:', e);
  }
}

function getStoredState(sessionId) {
  try {
    return localStorage.getItem(`gotea_state_${sessionId}`) || null;
  } catch (e) {
    console.warn('Failed to retrieve state:', e);
    return null;
  }
}

// WebSocket connection management
let socket = null;
let reconnectDelay = INITIAL_RECONNECT_DELAY;
let reconnectTimeout = null;
let intentionalClose = false;

function buildWebSocketUrl() {
  const sessionId = getCookie('session_id');
  const storedState = getStoredState(sessionId);
  const restoredStateParam = storedState ?
    `&restored_state=${encodeURIComponent(storedState)}` : '';

  return `${window.location.protocol === "https:" ? "wss://" : "ws://"}${window.location.host}/server?whence=${document.location.pathname}${restoredStateParam}`;
}

function connect() {
  console.log("Attempting to establish WebSocket connection");

  socket = new WebSocket(buildWebSocketUrl());

  socket.onmessage = event => {
    const data = event.data;

    // Try to parse as JSON to check for system messages
    try {
      const msg = JSON.parse(data);
      if (msg.type === 'STATE_SNAPSHOT') {
        const sessionId = getCookie('session_id');
        storeState(sessionId, msg.data);
        return; // Don't render system messages
      }
    } catch (e) {
      // Not JSON, treat as HTML
    }

    console.log("Received rerender from server");
    morphdom(document.documentElement, event.data, {
      childrenOnly: true
    });
  };

  socket.onopen = () => {
    console.log("WebSocket connection established.");
    // Reset reconnect delay on successful connection
    reconnectDelay = INITIAL_RECONNECT_DELAY;
  };

  socket.onerror = error => {
    console.error("WebSocket error:", error);
  };

  socket.onclose = event => {
    if (event.wasClean) {
      console.log(`WebSocket connection closed cleanly, code=${event.code}, reason=${event.reason}`);
    } else {
      console.error("WebSocket connection closed unexpectedly, code=", event.code, "reason=", event.reason);
    }

    // Attempt reconnection unless intentionally closed
    if (!intentionalClose) {
      scheduleReconnect();
    }
  };
}

function scheduleReconnect() {
  if (reconnectTimeout) {
    clearTimeout(reconnectTimeout);
  }

  console.log(`Scheduling reconnection in ${reconnectDelay}ms...`);

  reconnectTimeout = setTimeout(() => {
    console.log("Attempting to reconnect...");
    connect();

    // Increase delay for next attempt (exponential backoff)
    reconnectDelay = Math.min(reconnectDelay * RECONNECT_BACKOFF_MULTIPLIER, MAX_RECONNECT_DELAY);
  }, reconnectDelay);
}

// Initial connection
connect();

// Helper to safely send through websocket
function safeSend(data) {
  if (socket && socket.readyState === WebSocket.OPEN) {
    socket.send(data);
    return true;
  } else {
    console.warn("WebSocket not connected. Message queued for reconnection.");
    // Could implement message queuing here if needed
    return false;
  }
}

// Send a message through the websocket
const sendMessage = (msg) => {
  const msgJsonString = JSON.stringify(msg);
  console.log(`${SOCKET_MESSAGE}`, msgJsonString);
  safeSend(msgJsonString);
};

// Send a message with a value from an input field
const sendMessageWithValueFromInput = (msg, inputID) => {
  msg.args = document.getElementById(inputID).value;

  const msgJsonString = JSON.stringify(msg);
  console.log(`${SOCKET_MESSAGE}`, msgJsonString);
  safeSend(msgJsonString);
};

const sendMessageWithValueFromThisInput = (msg) => {
  msg.args = document.activeElement.value;

  const msgJsonString = JSON.stringify(msg);
  console.log(`${SOCKET_MESSAGE}`, msgJsonString);
  safeSend(msgJsonString);
};

// Submit a form through the websocket
const updateFormState = (msg, formID) => {
  msg.args = serializeForm(formID);

  console.log(`${SOCKET_MESSAGE}`, msg);
  safeSend(JSON.stringify(msg));
};


// Serialize form data into an object
const serializeForm = formID => {
  const formElements = [...document.getElementById(formID).elements];
  const TEXT_TYPES = ["text", "email", "number", "tel", "url", "password", "search", "date", "datetime-local", "time", "month", "week", "color", "range", "hidden"];
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
    if (TEXT_TYPES.includes(el.type)) {
      acc[el.name] = el.value;
    } else switch (el.type) {
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
  safeSend(JSON.stringify(msg));
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
  safeSend(JSON.stringify(msg));
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
