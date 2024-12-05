(()=>{var e,t,n="undefined"==typeof document?void 0:document,a=!!n&&"content"in n.createElement("template"),r=!!n&&n.createRange&&"createContextualFragment"in n.createRange();function i(e,t){var n,a,r=e.nodeName,i=t.nodeName;return r===i||((n=r.charCodeAt(0),a=i.charCodeAt(0),n<=90&&a>=97)?r===i.toUpperCase():a<=90&&n>=97&&i===r.toUpperCase())}function o(e,t,n){e[n]!==t[n]&&(e[n]=t[n],e[n]?e.setAttribute(n,""):e.removeAttribute(n))}var l={OPTION:function(e,t){var n=e.parentNode;if(n){var a=n.nodeName.toUpperCase();"OPTGROUP"===a&&(a=(n=n.parentNode)&&n.nodeName.toUpperCase()),"SELECT"!==a||n.hasAttribute("multiple")||(e.hasAttribute("selected")&&!t.selected&&(e.setAttribute("selected","selected"),e.removeAttribute("selected")),n.selectedIndex=-1)}o(e,t,"selected")},INPUT:function(e,t){o(e,t,"checked"),o(e,t,"disabled"),e.value!==t.value&&(e.value=t.value),t.hasAttribute("value")||e.removeAttribute("value")},TEXTAREA:function(e,t){var n=t.value;e.value!==n&&(e.value=n);var a=e.firstChild;if(a){var r=a.nodeValue;if(r==n||!n&&r==e.placeholder)return;a.nodeValue=n}},SELECT:function(e,t){if(!t.hasAttribute("multiple")){for(var n,a,r=-1,i=0,o=e.firstChild;o;)if("OPTGROUP"===(a=o.nodeName&&o.nodeName.toUpperCase()))o=(n=o).firstChild;else{if("OPTION"===a){if(o.hasAttribute("selected")){r=i;break}i++}(o=o.nextSibling)||!n||(o=n.nextSibling,n=null)}e.selectedIndex=r}}};function d(){}function s(e){if(e)return e.getAttribute&&e.getAttribute("id")||e.id}var c=(t=function(e,t){var n,a,r,i,o=t.attributes;if(11!==t.nodeType&&11!==e.nodeType){for(var l=o.length-1;l>=0;l--)a=(n=o[l]).name,r=n.namespaceURI,i=n.value,r?(a=n.localName||a,e.getAttributeNS(r,a)!==i&&("xmlns"===n.prefix&&(a=n.name),e.setAttributeNS(r,a,i))):e.getAttribute(a)!==i&&e.setAttribute(a,i);for(var d=e.attributes,s=d.length-1;s>=0;s--)a=(n=d[s]).name,(r=n.namespaceURI)?(a=n.localName||a,t.hasAttributeNS(r,a)||e.removeAttributeNS(r,a)):t.hasAttribute(a)||e.removeAttribute(a)}},function(o,c,u){if(u||(u={}),"string"==typeof c){if("#document"===o.nodeName||"HTML"===o.nodeName||"BODY"===o.nodeName){var f,m,v,p,g,h,N,b,A=c;(c=n.createElement("html")).innerHTML=A}else f=(f=c).trim(),c=a?(m=f,(v=n.createElement("template")).innerHTML=m,v.content.childNodes[0]):r?(p=f,e||(e=n.createRange()).selectNode(n.body),e.createContextualFragment(p).childNodes[0]):(g=f,(h=n.createElement("body")).innerHTML=g,h.childNodes[0])}else 11===c.nodeType&&(c=c.firstElementChild);var E=u.getNodeKey||s,C=u.onBeforeNodeAdded||d,T=u.onNodeAdded||d,S=u.onBeforeElUpdated||d,y=u.onElUpdated||d,w=u.onBeforeNodeDiscarded||d,O=u.onNodeDiscarded||d,x=u.onBeforeElChildrenUpdated||d,U=u.skipFromChildren||d,R=u.addChild||function(e,t){return e.appendChild(t)},k=!0===u.childrenOnly,V=Object.create(null),I=[];function L(e){I.push(e)}function B(e,t,n){!1!==w(e)&&(t&&t.removeChild(e),O(e),function e(t,n){if(1===t.nodeType)for(var a=t.firstChild;a;){var r=void 0;n&&(r=E(a))?L(r):(O(a),a.firstChild&&e(a,n)),a=a.nextSibling}}(e,n))}function M(e){if(1===e.nodeType||11===e.nodeType)for(var t=e.firstChild;t;){var n=E(t);n&&(V[n]=t),M(t),t=t.nextSibling}}M(o);var $=o,H=$.nodeType,P=c.nodeType;if(!k){if(1===H)1===P?i(o,c)||(O(o),$=function(e,t){for(var n=e.firstChild;n;){var a=n.nextSibling;t.appendChild(n),n=a}return t}(o,(N=c.nodeName,(b=c.namespaceURI)&&"http://www.w3.org/1999/xhtml"!==b?n.createElementNS(b,N):n.createElement(N)))):$=c;else if(3===H||8===H){if(P===H)return $.nodeValue!==c.nodeValue&&($.nodeValue=c.nodeValue),$;$=c}}if($===c)O(o);else{if(c.isSameNode&&c.isSameNode($))return;if(function e(a,r,o){var d=E(r);if(d&&delete V[d],!o){var s=S(a,r);if(!1===s||(s instanceof HTMLElement&&M(a=s),t(a,r),y(a),!1===x(a,r)))return}"TEXTAREA"!==a.nodeName?function(t,a){var r,o,d,s,c,u=U(t,a),f=a.firstChild,m=t.firstChild;e:for(;f;){for(s=f.nextSibling,r=E(f);!u&&m;){if(d=m.nextSibling,f.isSameNode&&f.isSameNode(m)){f=s,m=d;continue e}o=E(m);var v=m.nodeType,p=void 0;if(v===f.nodeType&&(1===v?(r?r!==o&&((c=V[r])?d===c?p=!1:(t.insertBefore(c,m),o?L(o):B(m,t,!0),o=E(m=c)):p=!1):o&&(p=!1),(p=!1!==p&&i(m,f))&&e(m,f)):(3===v||8==v)&&(p=!0,m.nodeValue!==f.nodeValue&&(m.nodeValue=f.nodeValue))),p){f=s,m=d;continue e}o?L(o):B(m,t,!0),m=d}if(r&&(c=V[r])&&i(c,f))u||R(t,c),e(c,f);else{var g=C(f);!1!==g&&(g&&(f=g),f.actualize&&(f=f.actualize(t.ownerDocument||n)),R(t,f),function t(n){T(n);for(var a=n.firstChild;a;){var r=a.nextSibling,o=E(a);if(o){var l=V[o];l&&i(a,l)?(a.parentNode.replaceChild(l,a),e(l,a)):t(a)}else t(a);a=r}}(f))}f=s,m=d}!function(e,t,n){for(;t;){var a=t.nextSibling;(n=E(t))?L(n):B(t,e,!0),t=a}}(t,m,o);var h=l[t.nodeName];h&&h(t,a)}(a,r):l.TEXTAREA(a,r)}($,c,k),I)for(var D=0,J=I.length;D<J;D++){var z=V[I[D]];z&&B(z,z.parentNode,!1)}}return!k&&$!==o&&o.parentNode&&($.actualize&&($=$.actualize(o.ownerDocument||n)),o.parentNode.replaceChild($,o)),$});let u="Sent Message:",f=new WebSocket(`${"https:"===window.location.protocol?"wss://":"ws://"}${window.location.host}/server?whence=${document.location.pathname}`);f.onmessage=e=>{c(document.documentElement,e.data,{childrenOnly:!0})};let m=e=>{let t=[...document.getElementById(e).elements],n=e=>[...e.children].map(e=>e.selected?e.value:"").filter(e=>e.length>0),a=e=>e.multiple?n(e):e.value;return t.reduce((e,t)=>{switch(t.tagName){case"SELECT":e[t.name]=a(t);break;case"TEXTAREA":e[t.name]=t.value}switch(t.type){case"text":e[t.name]=t.value;break;case"checkbox":e[t.name]=t.checked;break;case"radio":t.checked&&(e[t.name]=t.value)}return e},{})},v=e=>{history.pushState({},"",e);let t={message:"CHANGE_ROUTE",args:e};console.log(`${u}`,t),f.send(JSON.stringify(t))};window.gotea={sendMessage:(e,t)=>{let n={message:e,args:JSON.parse(t)};console.log(`${u}`,n),f.send(JSON.stringify(n))},submitForm:(e,t)=>{let n={message:e,args:m(t)};console.log(`${u}`,n),f.send(JSON.stringify(n))},sendMessageWithValue:(e,t)=>{let n={message:e,args:document.getElementById(t).value};console.log(`${u}`,n),f.send(JSON.stringify(n))}},window.addEventListener("popstate",e=>{let t={message:"CHANGE_ROUTE",args:document.location.pathname};console.log(`${u}`,t),f.send(JSON.stringify(t))}),document.addEventListener("click",e=>{if("A"===e.target.tagName&&!/external/.test(e.target.className))return e.preventDefault(),v(e.target.getAttribute("href")),!1},!1)})();
//# sourceMappingURL=main.js.map
