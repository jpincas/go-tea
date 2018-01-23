package gotea

var index = `
<html lang="en"> <head> <meta charset="UTF-8"> <meta name="viewport" content="width=device-width, initial-scale=1.0"> <meta http-equiv="X-UA-Compatible" content="ie=edge"> <title>gotea App Framework</title> <script>
require=function(r,e,n){function t(n,o){function i(r){return t(i.resolve(r))}function f(e){return r[n][1][e]||e}if(!e[n]){if(!r[n]){var c="function"==typeof require&&require;if(!o&&c)return c(n,!0);if(u)return u(n,!0);var l=new Error("Cannot find module '"+n+"'");throw l.code="MODULE_NOT_FOUND",l}i.resolve=f;var a=e[n]=new t.Module;r[n][0].call(a.exports,i,a,a.exports)}return e[n].exports}function o(){this.bundle=t,this.exports={}}var u="function"==typeof require&&require;t.Module=o,t.modules=r,t.cache=e,t.parent=u;for(var i=0;i<n.length;i++)t(n[i]);return t}({15:[function(require,module,exports) {
	"use strict";function e(e,t,o,r,d){return{sel:e,data:t,children:o,text:r,elm:d,key:void 0===t?void 0:t.key}}Object.defineProperty(exports,"__esModule",{value:!0}),exports.vnode=e,exports.default=e;
	},{}],17:[function(require,module,exports) {
	"use strict";Object.defineProperty(exports,"__esModule",{value:!0}),exports.primitive=e;var r=exports.array=Array.isArray;function e(r){return"string"==typeof r||"number"==typeof r}
	},{}],16:[function(require,module,exports) {
	"use strict";function e(e){return document.createElement(e)}function t(e,t){return document.createElementNS(e,t)}function n(e){return document.createTextNode(e)}function r(e){return document.createComment(e)}function o(e,t,n){e.insertBefore(t,n)}function u(e,t){e.removeChild(t)}function i(e,t){e.appendChild(t)}function c(e){return e.parentNode}function m(e){return e.nextSibling}function d(e){return e.tagName}function a(e,t){e.textContent=t}function f(e){return e.textContent}function l(e){return 1===e.nodeType}function p(e){return 3===e.nodeType}function s(e){return 8===e.nodeType}Object.defineProperty(exports,"__esModule",{value:!0});var x=exports.htmlDomApi={createElement:e,createElementNS:t,createTextNode:n,createComment:r,insertBefore:o,removeChild:u,appendChild:i,parentNode:c,nextSibling:m,tagName:d,setTextContent:a,getTextContent:f,isElement:l,isText:p,isComment:s};exports.default=x;
	},{}],18:[function(require,module,exports) {
	"use strict";Object.defineProperty(exports,"__esModule",{value:!0}),exports.h=v;var e=require("./vnode"),r=require("./is"),i=o(r);function o(e){if(e&&e.__esModule)return e;var r={};if(null!=e)for(var i in e)Object.prototype.hasOwnProperty.call(e,i)&&(r[i]=e[i]);return r.default=e,r}function t(e,r,i){if(e.ns="http://www.w3.org/2000/svg","foreignObject"!==i&&void 0!==r)for(var o=0;o<r.length;++o){var v=r[o].data;void 0!==v&&t(v,r[o].children,r[o].sel)}}function v(r,o,v){var n,a,d,l={};if(void 0!==v?(l=o,i.array(v)?n=v:i.primitive(v)?a=v:v&&v.sel&&(n=[v])):void 0!==o&&(i.array(o)?n=o:i.primitive(o)?a=o:o&&o.sel?n=[o]:l=o),i.array(n))for(d=0;d<n.length;++d)i.primitive(n[d])&&(n[d]=(0,e.vnode)(void 0,void 0,void 0,n[d],void 0));return"s"!==r[0]||"v"!==r[1]||"g"!==r[2]||3!==r.length&&"."!==r[3]&&"#"!==r[3]||t(l,n,r),(0,e.vnode)(r,l,n,a,void 0)}exports.default=v;
	},{"./vnode":15,"./is":17}],19:[function(require,module,exports) {
	"use strict";Object.defineProperty(exports,"__esModule",{value:!0}),exports.thunk=void 0;var t=require("./h");function a(t,a){a.elm=t.elm,t.data.fn=a.data.fn,t.data.args=a.data.args,a.data=t.data,a.children=t.children,a.text=t.text,a.elm=t.elm}function e(t){var e=t.data;a(e.fn.apply(void 0,e.args),t)}function n(t,e){var n,r=t.data,d=e.data,i=r.args,o=d.args;if(r.fn===d.fn&&i.length===o.length){for(n=0;n<o.length;++n)if(i[n]!==o[n])return void a(d.fn.apply(void 0,o),e);a(t,e)}else a(d.fn.apply(void 0,o),e)}var r=exports.thunk=function(a,r,d,i){return void 0===i&&(i=d,d=r,r=void 0),(0,t.h)(a,{key:r,hook:{init:e,prepatch:n},fn:d,args:i})};exports.default=r;
	},{"./h":18}],11:[function(require,module,exports) {
	"use strict";Object.defineProperty(exports,"__esModule",{value:!0}),exports.thunk=exports.h=void 0;var e=require("./h");Object.defineProperty(exports,"h",{enumerable:!0,get:function(){return e.h}});var t=require("./thunk");Object.defineProperty(exports,"thunk",{enumerable:!0,get:function(){return t.thunk}}),exports.init=x;var r=require("./vnode"),n=d(r),o=require("./is"),l=u(o),i=require("./htmldomapi"),a=d(i);function u(e){if(e&&e.__esModule)return e;var t={};if(null!=e)for(var r in e)Object.prototype.hasOwnProperty.call(e,r)&&(t[r]=e[r]);return t.default=e,t}function d(e){return e&&e.__esModule?e:{default:e}}function f(e){return void 0===e}function s(e){return void 0!==e}var v=(0,n.default)("",{},[],void 0,void 0);function c(e,t){return e.key===t.key&&e.sel===t.sel}function h(e){return void 0!==e.sel}function m(e,t,r){var n,o,l,i={};for(n=t;n<=r;++n)null!=(l=e[n])&&void 0!==(o=l.key)&&(i[o]=n);return i}var p=["create","update","remove","destroy","pre","post"];function x(e,t){var r,o,i={},u=void 0!==t?t:a.default;for(r=0;r<p.length;++r)for(i[p[r]]=[],o=0;o<e.length;++o){var d=e[o][p[r]];void 0!==d&&i[p[r]].push(d)}function x(e,t){return function(){if(0==--t){var r=u.parentNode(e);u.removeChild(r,e)}}}function g(e,t){var r,n=e.data;void 0!==n&&s(r=n.hook)&&s(r=r.init)&&(r(e),n=e.data);var o=e.children,a=e.sel;if("!"===a)f(e.text)&&(e.text=""),e.elm=u.createComment(e.text);else if(void 0!==a){var d=a.indexOf("#"),c=a.indexOf(".",d),h=d>0?d:a.length,m=c>0?c:a.length,p=-1!==d||-1!==c?a.slice(0,Math.min(h,m)):a,x=e.elm=s(n)&&s(r=n.ns)?u.createElementNS(r,p):u.createElement(p);for(h<m&&x.setAttribute("id",a.slice(h+1,m)),c>0&&x.setAttribute("class",a.slice(m+1).replace(/\./g," ")),r=0;r<i.create.length;++r)i.create[r](v,e);if(l.array(o))for(r=0;r<o.length;++r){var k=o[r];null!=k&&u.appendChild(x,g(k,t))}else l.primitive(e.text)&&u.appendChild(x,u.createTextNode(e.text));s(r=e.data.hook)&&(r.create&&r.create(v,e),r.insert&&t.push(e))}else e.elm=u.createTextNode(e.text);return e.elm}function k(e,t,r,n,o,l){for(;n<=o;++n){var i=r[n];null!=i&&u.insertBefore(e,g(i,l),t)}}function y(e){var t,r,n=e.data;if(void 0!==n){for(s(t=n.hook)&&s(t=t.destroy)&&t(e),t=0;t<i.destroy.length;++t)i.destroy[t](e);if(void 0!==e.children)for(r=0;r<e.children.length;++r)null!=(t=e.children[r])&&"string"!=typeof t&&y(t)}}function b(e,t,r,n){for(;r<=n;++r){var o=void 0,l=void 0,a=void 0,d=t[r];if(null!=d)if(s(d.sel)){for(y(d),l=i.remove.length+1,a=x(d.elm,l),o=0;o<i.remove.length;++o)i.remove[o](d,a);s(o=d.data)&&s(o=o.hook)&&s(o=o.remove)?o(d,a):a()}else u.removeChild(e,d.elm)}}function C(e,t,r){var n,o;s(n=t.data)&&s(o=n.hook)&&s(n=o.prepatch)&&n(e,t);var l=t.elm=e.elm,a=e.children,d=t.children;if(e!==t){if(void 0!==t.data){for(n=0;n<i.update.length;++n)i.update[n](e,t);s(n=t.data.hook)&&s(n=n.update)&&n(e,t)}f(t.text)?s(a)&&s(d)?a!==d&&function(e,t,r,n){for(var o,l,i,a=0,d=0,s=t.length-1,v=t[0],h=t[s],p=r.length-1,x=r[0],y=r[p];a<=s&&d<=p;)null==v?v=t[++a]:null==h?h=t[--s]:null==x?x=r[++d]:null==y?y=r[--p]:c(v,x)?(C(v,x,n),v=t[++a],x=r[++d]):c(h,y)?(C(h,y,n),h=t[--s],y=r[--p]):c(v,y)?(C(v,y,n),u.insertBefore(e,v.elm,u.nextSibling(h.elm)),v=t[++a],y=r[--p]):c(h,x)?(C(h,x,n),u.insertBefore(e,h.elm,v.elm),h=t[--s],x=r[++d]):(void 0===o&&(o=m(t,a,s)),f(l=o[x.key])?(u.insertBefore(e,g(x,n),v.elm),x=r[++d]):((i=t[l]).sel!==x.sel?u.insertBefore(e,g(x,n),v.elm):(C(i,x,n),t[l]=void 0,u.insertBefore(e,i.elm,v.elm)),x=r[++d]));(a<=s||d<=p)&&(a>s?k(e,null==r[p+1]?null:r[p+1].elm,r,d,p,n):b(e,t,a,s))}(l,a,d,r):s(d)?(s(e.text)&&u.setTextContent(l,""),k(l,null,d,0,d.length-1,r)):s(a)?b(l,a,0,a.length-1):s(e.text)&&u.setTextContent(l,""):e.text!==t.text&&u.setTextContent(l,t.text),s(o)&&s(n=o.postpatch)&&n(e,t)}}return function(e,t){var r,o,l,a,d,f,s=[];for(r=0;r<i.pre.length;++r)i.pre[r]();for(h(e)||(d=(a=e).id?"#"+a.id:"",f=a.className?"."+a.className.split(" ").join("."):"",e=(0,n.default)(u.tagName(a).toLowerCase()+d+f,{},[],void 0,a)),c(e,t)?C(e,t,s):(o=e.elm,l=u.parentNode(o),g(t,s),null!==l&&(u.insertBefore(l,t.elm,u.nextSibling(o)),b(l,[e],0,0))),r=0;r<s.length;++r)s[r].data.hook.insert(s[r]);for(r=0;r<i.post.length;++r)i.post[r]();return t}}
	},{"./vnode":15,"./is":17,"./htmldomapi":16,"./h":18,"./thunk":19}],8:[function(require,module,exports) {
	"use strict";function e(e,s){var a,t,o=s.elm,r=e.data.class,l=s.data.class;if((r||l)&&r!==l){r=r||{},l=l||{};for(t in r)l[t]||o.classList.remove(t);for(t in l)(a=l[t])!==r[t]&&o.classList[a?"add":"remove"](t)}}Object.defineProperty(exports,"__esModule",{value:!0}),exports.classModule={create:e,update:e},exports.default=exports.classModule;
	},{}],7:[function(require,module,exports) {
	"use strict";function e(e,o){var r,t,p=o.elm,s=e.data.props,a=o.data.props;if((s||a)&&s!==a){s=s||{},a=a||{};for(r in s)a[r]||delete p[r];for(r in a)t=a[r],s[r]===t||"value"===r&&p[r]===t||(p[r]=t)}}Object.defineProperty(exports,"__esModule",{value:!0}),exports.propsModule={create:e,update:e},exports.default=exports.propsModule;
	},{}],10:[function(require,module,exports) {
	"use strict";Object.defineProperty(exports,"__esModule",{value:!0});var e="undefined"!=typeof window&&window.requestAnimationFrame||setTimeout,t=function(t){e(function(){e(t)})};function r(e,r,o){t(function(){e[r]=o})}function o(e,t){var o,n,d=t.elm,i=e.data.style,a=t.data.style;if((i||a)&&i!==a){i=i||{},a=a||{};var l="delayed"in i;for(n in i)a[n]||("-"===n[0]&&"-"===n[1]?d.style.removeProperty(n):d.style[n]="");for(n in a)if(o=a[n],"delayed"===n&&a.delayed)for(var s in a.delayed)o=a.delayed[s],l&&o===i.delayed[s]||r(d.style,s,o);else"remove"!==n&&o!==i[n]&&("-"===n[0]&&"-"===n[1]?d.style.setProperty(n,o):d.style[n]=o)}}function n(e){var t,r,o=e.elm,n=e.data.style;if(n&&(t=n.destroy))for(r in t)o.style[r]=t[r]}function d(e,t){var r=e.data.style;if(r&&r.remove){var o,n=e.elm,d=0,i=r.remove,a=0,l=[];for(o in i)l.push(o),n.style[o]=i[o];for(var s=getComputedStyle(n)["transition-property"].split(", ");d<s.length;++d)-1!==l.indexOf(s[d])&&a++;n.addEventListener("transitionend",function(e){e.target===n&&--a,0===a&&t()})}else t()}exports.styleModule={create:o,update:o,destroy:n,remove:d},exports.default=exports.styleModule;
	},{}],9:[function(require,module,exports) {
	"use strict";Object.defineProperty(exports,"__esModule",{value:!0});var t="http://www.w3.org/1999/xlink",e="http://www.w3.org/XML/1998/namespace",r=58,a=120;function o(o,i){var s,u=i.elm,d=o.data.attrs,b=i.data.attrs;if((d||b)&&d!==b){d=d||{},b=b||{};for(s in b){var A=b[s];d[s]!==A&&(!0===A?u.setAttribute(s,""):!1===A?u.removeAttribute(s):s.charCodeAt(0)!==a?u.setAttribute(s,A):s.charCodeAt(3)===r?u.setAttributeNS(e,s,A):s.charCodeAt(5)===r?u.setAttributeNS(t,s,A):u.setAttribute(s,A))}for(s in d)s in b||u.removeAttribute(s)}}exports.attributesModule={create:o,update:o},exports.default=exports.attributesModule;
	},{}],12:[function(require,module,exports) {
	"use strict";function r(r){return"string"==typeof r||"number"==typeof r}Object.defineProperty(exports,"__esModule",{value:!0}),exports.array=Array.isArray,exports.primitive=r;
	},{}],5:[function(require,module,exports) {
	"use strict";Object.defineProperty(exports,"__esModule",{value:!0});var e=require("./vnode"),r=require("./is");function i(e,r,o){if(e.ns="http://www.w3.org/2000/svg","foreignObject"!==o&&void 0!==r)for(var v=0;v<r.length;++v){var t=r[v].data;void 0!==t&&i(t,r[v].children,r[v].sel)}}function o(o,v,t){var d,n,a,s={};if(void 0!==t?(s=v,r.array(t)?d=t:r.primitive(t)?n=t:t&&t.sel&&(d=[t])):void 0!==v&&(r.array(v)?d=v:r.primitive(v)?n=v:v&&v.sel?d=[v]:s=v),r.array(d))for(a=0;a<d.length;++a)r.primitive(d[a])&&(d[a]=e.vnode(void 0,void 0,void 0,d[a],void 0));return"s"!==o[0]||"v"!==o[1]||"g"!==o[2]||3!==o.length&&"."!==o[3]&&"#"!==o[3]||i(s,d,o),e.vnode(o,s,d,n,void 0)}exports.h=o,exports.default=o;
	},{"./vnode":15,"./is":12}],14:[function(require,module,exports) {
	"use strict";function e(e){return document.createElement(e)}function t(e,t){return document.createElementNS(e,t)}function n(e){return document.createTextNode(e)}function o(e){return document.createComment(e)}function r(e,t,n){e.insertBefore(t,n)}function u(e,t){e.removeChild(t)}function i(e,t){e.appendChild(t)}function c(e){return e.parentNode}function m(e){return e.nextSibling}function d(e){return e.tagName}function f(e,t){e.textContent=t}function a(e){return e.textContent}function l(e){return 1===e.nodeType}function p(e){return 3===e.nodeType}function s(e){return 8===e.nodeType}Object.defineProperty(exports,"__esModule",{value:!0}),exports.htmlDomApi={createElement:e,createElementNS:t,createTextNode:n,createComment:o,insertBefore:r,removeChild:u,appendChild:i,parentNode:c,nextSibling:m,tagName:d,setTextContent:f,getTextContent:a,isElement:l,isText:p,isComment:s},exports.default=exports.htmlDomApi;
	},{}],6:[function(require,module,exports) {
	"use strict";Object.defineProperty(exports,"__esModule",{value:!0});var e=require("./vnode"),t=require("./htmldomapi");function o(d,i){var r,a=void 0!==i?i:t.default;if(a.isElement(d)){var s,u=d.id?"#"+d.id:"",l=d.getAttribute("class"),n=l?"."+l.split(" ").join("."):"",v=a.tagName(d).toLowerCase()+u+n,f={},m=[],p=void 0,c=void 0,g=d.attributes,x=d.childNodes;for(p=0,c=g.length;p<c;p++)"id"!==(s=g[p].nodeName)&&"class"!==s&&(f[s]=g[p].nodeValue);for(p=0,c=x.length;p<c;p++)m.push(o(x[p]));return e.default(v,{attrs:f},m,void 0,d)}return a.isText(d)?(r=a.getTextContent(d),e.default(void 0,void 0,void 0,r,d)):a.isComment(d)?(r=a.getTextContent(d),e.default("!",{},[],r,d)):e.default("",{},[],void 0,d)}exports.toVNode=o,exports.default=o;
	},{"./vnode":15,"./htmldomapi":14}],4:[function(require,module,exports) {
	var e,t,a=require("snabbdom"),o=a.init([require("snabbdom/modules/class").default,require("snabbdom/modules/props").default,require("snabbdom/modules/style").default,require("snabbdom/modules/attributes").default]),s=require("snabbdom/h").default,n=require("snabbdom/tovnode").default,d=new WebSocket(("https:"===window.location.protocol?"wss://":"ws://")+window.location.host+"/server");function r(e){const t=e.target.dataset.message;var a=e.target.dataset.tag;let o={};a&&(o=JSON.parse(a));const s={message:t,tag:o};d.send(JSON.stringify(s))}d.onmessage=function(a){var s=document.createElement("div");s.innerHTML=a.data,s.setAttribute("id","view"),newNode=n(s),e?(o(t,newNode),t=newNode):(e=document.getElementById("view"),o(e,newNode),t=newNode)},document.addEventListener("click",function(e){/gotea-click/.test(e.target.className)&&r(e)},!1);
	},{"snabbdom":11,"snabbdom/modules/class":8,"snabbdom/modules/props":7,"snabbdom/modules/style":10,"snabbdom/modules/attributes":9,"snabbdom/h":5,"snabbdom/tovnode":6}]},{},[4])
</script> </head> <body> <div id="view"></div> </body> </html>
`