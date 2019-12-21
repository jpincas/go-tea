package gotea

const goteaJS = `parcelRequire=function(e,r,n,t){var i="function"==typeof parcelRequire&&parcelRequire,o="function"==typeof require&&require;function u(n,t){if(!r[n]){if(!e[n]){var f="function"==typeof parcelRequire&&parcelRequire;if(!t&&f)return f(n,!0);if(i)return i(n,!0);if(o&&"string"==typeof n)return o(n);var c=new Error("Cannot find module '"+n+"'");throw c.code="MODULE_NOT_FOUND",c}p.resolve=function(r){return e[n][1][r]||r};var l=r[n]=new u.Module(n);e[n][0].call(l.exports,p,l,l.exports,this)}return r[n].exports;function p(e){return u(p.resolve(e))}}u.isParcelRequire=!0,u.Module=function(e){this.id=e,this.bundle=u,this.exports={}},u.modules=e,u.cache=r,u.parent=i,u.register=function(r,n){e[r]=[function(e,r){r.exports=n},{}]};for(var f=0;f<n.length;f++)u(n[f]);if(n.length){var c=u(n[n.length-1]);"object"==typeof exports&&"undefined"!=typeof module?module.exports=c:"function"==typeof define&&define.amd?define(function(){return c}):t&&(this[t]=c)}return u}({"hAtP":[function(require,module,exports) {
	"use strict";var e,t,n="http://www.w3.org/1999/xhtml",r="undefined"==typeof document?void 0:document,i=r?r.body||r.createElement("div"):{},a=t=i.hasAttributeNS?function(e,t,n){return e.hasAttributeNS(t,n)}:i.hasAttribute?function(e,t,n){return e.hasAttribute(n)}:function(e,t,n){return null!=e.getAttributeNode(t,n)};function o(t){var n;return!e&&r.createRange&&(e=r.createRange()).selectNode(r.body),e&&e.createContextualFragment?n=e.createContextualFragment(t):(n=r.createElement("body")).innerHTML=t,n.childNodes[0]}function d(e,t){var n=e.nodeName,r=t.nodeName;return n===r||!!(t.actualize&&n.charCodeAt(0)<91&&r.charCodeAt(0)>90)&&n===r.toUpperCase()}function l(e,t){return t&&t!==n?r.createElementNS(t,e):r.createElement(e)}function u(e,t){for(var n=e.firstChild;n;){var r=n.nextSibling;t.appendChild(n),n=r}return t}function f(e,t){var n,r,i,o,d,l=t.attributes;for(n=l.length-1;n>=0;--n)i=(r=l[n]).name,o=r.namespaceURI,d=r.value,o?(i=r.localName||i,e.getAttributeNS(o,i)!==d&&e.setAttributeNS(o,i,d)):e.getAttribute(i)!==d&&e.setAttribute(i,d);for(n=(l=e.attributes).length-1;n>=0;--n)!1!==(r=l[n]).specified&&(i=r.name,(o=r.namespaceURI)?(i=r.localName||i,a(t,o,i)||e.removeAttributeNS(o,i)):a(t,null,i)||e.removeAttribute(i))}function c(e,t,n){e[n]!==t[n]&&(e[n]=t[n],e[n]?e.setAttribute(n,""):e.removeAttribute(n,""))}var v={OPTION:function(e,t){c(e,t,"selected")},INPUT:function(e,t){c(e,t,"checked"),c(e,t,"disabled"),e.value!==t.value&&(e.value=t.value),a(t,null,"value")||e.removeAttribute("value")},TEXTAREA:function(e,t){var n=t.value;e.value!==n&&(e.value=n);var r=e.firstChild;if(r){var i=r.nodeValue;if(i==n||!n&&i==e.placeholder)return;r.nodeValue=n}},SELECT:function(e,t){if(!a(t,null,"multiple")){for(var n=0,r=t.firstChild;r;){var i=r.nodeName;if(i&&"OPTION"===i.toUpperCase()){if(a(r,null,"selected")){n;break}n++}r=r.nextSibling}e.selectedIndex=n}}},s=1,m=3,p=8;function N(){}function h(e){return e.id}function b(e){return function(t,n,i){if(i||(i={}),"string"==typeof n)if("#document"===t.nodeName||"HTML"===t.nodeName){var a=n;(n=r.createElement("html")).innerHTML=a}else n=o(n);var f,c=i.getNodeKey||h,b=i.onBeforeNodeAdded||N,g=i.onNodeAdded||N,C=i.onBeforeElUpdated||N,A=i.onElUpdated||N,S=i.onBeforeNodeDiscarded||N,T=i.onNodeDiscarded||N,x=i.onBeforeElChildrenUpdated||N,E=!0===i.childrenOnly,y={};function V(e){f?f.push(e):f=[e]}function U(e,t,n){!1!==S(e)&&(t&&t.removeChild(e),T(e),function e(t,n){if(t.nodeType===s)for(var r=t.firstChild;r;){var i=void 0;n&&(i=c(r))?V(i):(T(r),r.firstChild&&e(r,n)),r=r.nextSibling}}(e,n))}function I(e){g(e);for(var t=e.firstChild;t;){var n=t.nextSibling,r=c(t);if(r){var i=y[r];i&&d(t,i)&&(t.parentNode.replaceChild(i,t),R(i,t))}I(t),t=n}}function R(i,a,o){var l,u=c(a);if(u&&delete y[u],!n.isSameNode||!n.isSameNode(t)){if(!o){if(!1===C(i,a))return;if(e(i,a),A(i),!1===x(i,a))return}if("TEXTAREA"!==i.nodeName){var f,N,h,g,S=a.firstChild,T=i.firstChild;e:for(;S;){for(h=S.nextSibling,f=c(S);T;){if(N=T.nextSibling,S.isSameNode&&S.isSameNode(T)){S=h,T=N;continue e}l=c(T);var E=T.nodeType,w=void 0;if(E===S.nodeType&&(E===s?(f?f!==l&&((g=y[f])?T.nextSibling===g?w=!1:(i.insertBefore(g,T),N=T.nextSibling,l?V(l):U(T,i,!0),T=g):w=!1):l&&(w=!1),(w=!1!==w&&d(T,S))&&R(T,S)):E!==m&&E!=p||(w=!0,T.nodeValue!==S.nodeValue&&(T.nodeValue=S.nodeValue))),w){S=h,T=N;continue e}l?V(l):U(T,i,!0),T=N}if(f&&(g=y[f])&&d(g,S))i.appendChild(g),R(g,S);else{var z=b(S);!1!==z&&(z&&(S=z),S.actualize&&(S=S.actualize(i.ownerDocument||r)),i.appendChild(S),I(S))}S=h,T=N}for(;T;)N=T.nextSibling,(l=c(T))?V(l):U(T,i,!0),T=N}var B=v[i.nodeName];B&&B(i,a)}}!function e(t){if(t.nodeType===s)for(var n=t.firstChild;n;){var r=c(n);r&&(y[r]=n),e(n),n=n.nextSibling}}(t);var w=t,z=w.nodeType,B=n.nodeType;if(!E)if(z===s)B===s?d(t,n)||(T(t),w=u(t,l(n.nodeName,n.namespaceURI))):w=n;else if(z===m||z===p){if(B===z)return w.nodeValue!==n.nodeValue&&(w.nodeValue=n.nodeValue),w;w=n}if(w===n)T(t);else if(R(w,n,E),f)for(var O=0,D=f.length;O<D;O++){var L=y[f[O]];L&&U(L,L.parentNode,!1)}return!E&&w!==t&&t.parentNode&&(w.actualize&&(w=w.actualize(t.ownerDocument||r)),t.parentNode.replaceChild(w,t)),w}}var g=b(f);module.exports=g;
	},{}],"PK39":[function(require,module,exports) {
	"use strict";var e=require("morphdom"),n=t(e);function t(e){return e&&e.__esModule?e:{default:e}}function a(e){if(Array.isArray(e)){for(var n=0,t=Array(e.length);n<e.length;n++)t[n]=e[n];return t}return Array.from(e)}var s=new WebSocket(("https:"===window.location.protocol?"wss://":"ws://")+window.location.host+"/server?whence="+document.location.pathname);s.onmessage=function(e){(0,n.default)(document.documentElement,e.data,{childrenOnly:!0})};var r=function(e,n){var t={message:e,args:JSON.parse(n)};console.log("Sending websocket message: ",t),s.send(JSON.stringify(t))};function o(e,n){var t={message:e,args:i(n)};console.log("Sending websocket message: ",t),s.send(JSON.stringify(t))}function c(e,n){var t={message:e,args:document.getElementById(n).value};console.log("Sending websocket message: ",t),s.send(JSON.stringify(t))}window.gotea={sendMessage:r,submitForm:o,sendMessageWithValue:c};var i=function(e){var n=[].concat(a(document.getElementById(e).elements)),t=function(e){return e.multiple?function(e){return[].concat(a(e.children)).map(function(e){return e.selected?e.value:""}).filter(function(e){return e.length>0})}(e):e.value};return n.reduce(function(e,n){switch(n.tagName){case"SELECT":e[n.name]=t(n);break;case"TEXTAREA":e[n.name]=n.value}switch(n.type){case"text":e[n.name]=n.value;break;case"checkbox":e[n.name]=n.checked;break;case"radio":n.checked&&(e[n.name]=n.value)}return e},{})};function u(e){history.pushState({},"",e);var n={message:"CHANGE_ROUTE",args:e};console.log("Sending websocket message: ",n),s.send(JSON.stringify(n))}window.addEventListener("popstate",function(e){var n={message:"CHANGE_ROUTE",args:document.location.pathname};console.log("Sending websocket message: ",n),s.send(JSON.stringify(n))}),document.addEventListener("click",function(e){if("A"==e.target.tagName&&0==/external/.test(e.target.className))return e.preventDefault(),u(e.target.getAttribute("href")),!1},!1);
	},{"morphdom":"hAtP"}]},{},["PK39"], null)
	//# sourceMappingURL=/gotea.map`
