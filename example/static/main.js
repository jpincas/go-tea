(()=>{var e,t,n="undefined"==typeof document?void 0:document,a=!!n&&"content"in n.createElement("template"),i=!!n&&n.createRange&&"createContextualFragment"in n.createRange();function r(e,t){var n,a,i=e.nodeName,r=t.nodeName;return i===r||((n=i.charCodeAt(0),a=r.charCodeAt(0),n<=90&&a>=97)?i===r.toUpperCase():a<=90&&n>=97&&r===i.toUpperCase())}function o(e,t,n){e[n]!==t[n]&&(e[n]=t[n],e[n]?e.setAttribute(n,""):e.removeAttribute(n))}var d={OPTION:function(e,t){var n=e.parentNode;if(n){var a=n.nodeName.toUpperCase();"OPTGROUP"===a&&(a=(n=n.parentNode)&&n.nodeName.toUpperCase()),"SELECT"!==a||n.hasAttribute("multiple")||(e.hasAttribute("selected")&&!t.selected&&(e.setAttribute("selected","selected"),e.removeAttribute("selected")),n.selectedIndex=-1)}o(e,t,"selected")},INPUT:function(e,t){o(e,t,"checked"),o(e,t,"disabled"),e.value!==t.value&&(e.value=t.value),t.hasAttribute("value")||e.removeAttribute("value")},TEXTAREA:function(e,t){var n=t.value;e.value!==n&&(e.value=n);var a=e.firstChild;if(a){var i=a.nodeValue;if(i==n||!n&&i==e.placeholder)return;a.nodeValue=n}},SELECT:function(e,t){if(!t.hasAttribute("multiple")){for(var n,a,i=-1,r=0,o=e.firstChild;o;)if("OPTGROUP"===(a=o.nodeName&&o.nodeName.toUpperCase()))o=(n=o).firstChild;else{if("OPTION"===a){if(o.hasAttribute("selected")){i=r;break}r++}(o=o.nextSibling)||!n||(o=n.nextSibling,n=null)}e.selectedIndex=i}}};function l(){}function s(e){if(e)return e.getAttribute&&e.getAttribute("id")||e.id}var c=(t=function(e,t){var n,a,i,r,o=t.attributes;if(11!==t.nodeType&&11!==e.nodeType){for(var d=o.length-1;d>=0;d--)a=(n=o[d]).name,i=n.namespaceURI,r=n.value,i?(a=n.localName||a,e.getAttributeNS(i,a)!==r&&("xmlns"===n.prefix&&(a=n.name),e.setAttributeNS(i,a,r))):e.getAttribute(a)!==r&&e.setAttribute(a,r);for(var l=e.attributes,s=l.length-1;s>=0;s--)a=(n=l[s]).name,(i=n.namespaceURI)?(a=n.localName||a,t.hasAttributeNS(i,a)||e.removeAttributeNS(i,a)):t.hasAttribute(a)||e.removeAttribute(a)}},function(o,c,u){if(u||(u={}),"string"==typeof c){if("#document"===o.nodeName||"HTML"===o.nodeName||"BODY"===o.nodeName){var f,m,g,v,p,h,N,b,A=c;(c=n.createElement("html")).innerHTML=A}else f=(f=c).trim(),c=a?(m=f,(g=n.createElement("template")).innerHTML=m,g.content.childNodes[0]):i?(v=f,e||(e=n.createRange()).selectNode(n.body),e.createContextualFragment(v).childNodes[0]):(p=f,(h=n.createElement("body")).innerHTML=p,h.childNodes[0])}else 11===c.nodeType&&(c=c.firstElementChild);var E=u.getNodeKey||s,S=u.onBeforeNodeAdded||l,C=u.onNodeAdded||l,T=u.onBeforeElUpdated||l,y=u.onElUpdated||l,w=u.onBeforeNodeDiscarded||l,O=u.onNodeDiscarded||l,x=u.onBeforeElChildrenUpdated||l,k=u.skipFromChildren||l,U=u.addChild||function(e,t){return e.appendChild(t)},R=!0===u.childrenOnly,V=Object.create(null),I=[];function L(e){I.push(e)}function B(e,t,n){!1!==w(e)&&(t&&t.removeChild(e),O(e),function e(t,n){if(1===t.nodeType)for(var a=t.firstChild;a;){var i=void 0;n&&(i=E(a))?L(i):(O(a),a.firstChild&&e(a,n)),a=a.nextSibling}}(e,n))}function H(e){if(1===e.nodeType||11===e.nodeType)for(var t=e.firstChild;t;){var n=E(t);n&&(V[n]=t),H(t),t=t.nextSibling}}H(o);var M=o,P=M.nodeType,D=c.nodeType;if(!R){if(1===P)1===D?r(o,c)||(O(o),M=function(e,t){for(var n=e.firstChild;n;){var a=n.nextSibling;t.appendChild(n),n=a}return t}(o,(N=c.nodeName,(b=c.namespaceURI)&&"http://www.w3.org/1999/xhtml"!==b?n.createElementNS(b,N):n.createElement(N)))):M=c;else if(3===P||8===P){if(D===P)return M.nodeValue!==c.nodeValue&&(M.nodeValue=c.nodeValue),M;M=c}}if(M===c)O(o);else{if(c.isSameNode&&c.isSameNode(M))return;if(function e(a,i,o){var l=E(i);if(l&&delete V[l],!o){var s=T(a,i);if(!1===s||(s instanceof HTMLElement&&H(a=s),t(a,i),y(a),!1===x(a,i)))return}"TEXTAREA"!==a.nodeName?function(t,a){var i,o,l,s,c,u=k(t,a),f=a.firstChild,m=t.firstChild;e:for(;f;){for(s=f.nextSibling,i=E(f);!u&&m;){if(l=m.nextSibling,f.isSameNode&&f.isSameNode(m)){f=s,m=l;continue e}o=E(m);var g=m.nodeType,v=void 0;if(g===f.nodeType&&(1===g?(i?i!==o&&((c=V[i])?l===c?v=!1:(t.insertBefore(c,m),o?L(o):B(m,t,!0),o=E(m=c)):v=!1):o&&(v=!1),(v=!1!==v&&r(m,f))&&e(m,f)):(3===g||8==g)&&(v=!0,m.nodeValue!==f.nodeValue&&(m.nodeValue=f.nodeValue))),v){f=s,m=l;continue e}o?L(o):B(m,t,!0),m=l}if(i&&(c=V[i])&&r(c,f))u||U(t,c),e(c,f);else{var p=S(f);!1!==p&&(p&&(f=p),f.actualize&&(f=f.actualize(t.ownerDocument||n)),U(t,f),function t(n){C(n);for(var a=n.firstChild;a;){var i=a.nextSibling,o=E(a);if(o){var d=V[o];d&&r(a,d)?(a.parentNode.replaceChild(d,a),e(d,a)):t(a)}else t(a);a=i}}(f))}f=s,m=l}!function(e,t,n){for(;t;){var a=t.nextSibling;(n=E(t))?L(n):B(t,e,!0),t=a}}(t,m,o);var h=d[t.nodeName];h&&h(t,a)}(a,i):d.TEXTAREA(a,i)}(M,c,R),I)for(var J=0,z=I.length;J<z;J++){var F=V[I[J]];F&&B(F,F.parentNode,!1)}}return!R&&M!==o&&o.parentNode&&(M.actualize&&(M=M.actualize(o.ownerDocument||n)),o.parentNode.replaceChild(M,o)),M});let u=new WebSocket(("https:"===window.location.protocol?"wss://":"ws://")+window.location.host+"/server?whence="+document.location.pathname);u.onmessage=e=>{c(document.documentElement,e.data,{childrenOnly:!0})},window.gotea={sendMessage:(e,t)=>{let n={message:e,args:JSON.parse(t)};console.log("Sending websocket message: ",n),u.send(JSON.stringify(n))},submitForm:function(e,t){let n={message:e,args:f(t)};console.log("Sending websocket message: ",n),u.send(JSON.stringify(n))},sendMessageWithValue:function(e,t){let n={message:e,args:document.getElementById(t).value};console.log("Sending websocket message: ",n),u.send(JSON.stringify(n))}};let f=e=>{let t=[...document.getElementById(e).elements],n=e=>[...e.children].map(e=>e.selected?e.value:"").filter(e=>e.length>0),a=e=>e.multiple?n(e):e.value;return t.reduce((e,t)=>{switch(t.tagName){case"SELECT":e[t.name]=a(t);break;case"TEXTAREA":e[t.name]=t.value}switch(t.type){case"text":e[t.name]=t.value;break;case"checkbox":e[t.name]=t.checked;break;case"radio":t.checked&&(e[t.name]=t.value)}return e},{})};window.addEventListener("popstate",function(e){let t={message:"CHANGE_ROUTE",args:document.location.pathname};console.log("Sending websocket message: ",t),u.send(JSON.stringify(t))}),document.addEventListener("click",e=>{if("A"==e.target.tagName&&!1==/external/.test(e.target.className))return e.preventDefault(),function(e){history.pushState({},"",e);let t={message:"CHANGE_ROUTE",args:e};console.log("Sending websocket message: ",t),u.send(JSON.stringify(t))}(e.target.getAttribute("href")),!1},!1)})();
//# sourceMappingURL=main.js.map
