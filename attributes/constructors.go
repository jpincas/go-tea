package attributes

import (
	"fmt"
	"strings"

	"github.com/jpincas/go-tea/css"
)

// =============================================================================
// Global Attributes
// =============================================================================

const (
	id              = "id"
	class           = "class"
	style           = "style"
	title           = "title"
	lang            = "lang"
	dir             = "dir"
	hidden          = "hidden"
	tabindex        = "tabindex"
	accesskey       = "accesskey"
	contenteditable = "contenteditable"
	draggable       = "draggable"
	spellcheck      = "spellcheck"
	translate       = "translate"
)

// =============================================================================
// Form Attributes
// =============================================================================

const (
	name         = "name"
	value        = "value"
	type_        = "type"
	placeholder  = "placeholder"
	required     = "required"
	disabled     = "disabled"
	readonly     = "readonly"
	checked      = "checked"
	selected     = "selected"
	autofocus    = "autofocus"
	autocomplete = "autocomplete"
	form         = "form"
	formaction   = "formaction"
	formmethod   = "formmethod"
	formtarget   = "formtarget"
	accept       = "accept"
	inputmode    = "inputmode"
	min          = "min"
	max          = "max"
	minLength    = "minlength"
	maxLength    = "maxlength"
	pattern      = "pattern"
	step         = "step"
	multiple     = "multiple"
	list         = "list"
	rows         = "rows"
	cols         = "cols"
	wrap         = "wrap"
	size         = "size"
	for_         = "for"
	method       = "method"
	action       = "action"
)

// =============================================================================
// Link/Navigation Attributes
// =============================================================================

const (
	href     = "href"
	target   = "target"
	rel      = "rel"
	download = "download"
	hreflang = "hreflang"
	ping     = "ping"
)

// =============================================================================
// Media Attributes
// =============================================================================

const (
	src         = "src"
	alt         = "alt"
	width       = "width"
	height      = "height"
	srcset      = "srcset"
	sizes       = "sizes"
	media       = "media"
	autoplay    = "autoplay"
	controls    = "controls"
	loop        = "loop"
	muted       = "muted"
	poster      = "poster"
	preload     = "preload"
	playsinline = "playsinline"
)

// =============================================================================
// Script/Resource Attributes
// =============================================================================

const (
	defer_      = "defer"
	async       = "async"
	as          = "as"
	integrity   = "integrity"
	crossOrigin = "crossorigin"
	nonce       = "nonce"
)

// =============================================================================
// SVG Attributes
// =============================================================================

const (
	viewBox             = "viewBox" // yes, camel case is correct
	stroke              = "stroke"
	fill                = "fill"
	strokeWidth         = "stroke-width"
	strokeLineCap       = "stroke-linecap"
	strokeLineJoin      = "stroke-linejoin"
	strokeDasharray     = "stroke-dasharray"
	strokeDashoffset    = "stroke-dashoffset"
	strokeOpacity       = "stroke-opacity"
	fillOpacity         = "fill-opacity"
	fillRule            = "fill-rule"
	d                   = "d"
	points              = "points"
	x                   = "x"
	y                   = "y"
	x1                  = "x1"
	y1                  = "y1"
	x2                  = "x2"
	y2                  = "y2"
	cx                  = "cx"
	cy                  = "cy"
	r                   = "r"
	rx                  = "rx"
	ry                  = "ry"
	transform           = "transform"
	pathLength          = "pathLength"
	preserveAspectRatio = "preserveAspectRatio"
	clipPath            = "clip-path"
	mask                = "mask"
	gradientUnits       = "gradientUnits"
	gradientTransform   = "gradientTransform"
	offset              = "offset"
	stopColor           = "stop-color"
	stopOpacity         = "stop-opacity"
	xlinkHref           = "xlink:href"
	textAnchor          = "text-anchor"
	dominantBaseline    = "dominant-baseline"
)

// =============================================================================
// Event Handlers
// =============================================================================

const (
	// Mouse events
	onclick       = "onclick"
	ondblclick    = "ondblclick"
	onmousedown   = "onmousedown"
	onmouseup     = "onmouseup"
	onmouseover   = "onmouseover"
	onmouseout    = "onmouseout"
	onmousemove   = "onmousemove"
	onmouseenter  = "onmouseenter"
	onmouseleave  = "onmouseleave"
	oncontextmenu = "oncontextmenu"

	// Keyboard events
	onkeydown  = "onkeydown"
	onkeyup    = "onkeyup"
	onkeypress = "onkeypress"

	// Focus events
	onfocus = "onfocus"
	onblur  = "onblur"

	// Form events
	onchange  = "onchange"
	oninput   = "oninput"
	onsubmit  = "onsubmit"
	onreset   = "onreset"
	oninvalid = "oninvalid"

	// Document/Window events
	onload   = "onload"
	onerror  = "onerror"
	onscroll = "onscroll"
	onresize = "onresize"

	// Media events
	onplay       = "onplay"
	onpause      = "onpause"
	onended      = "onended"
	onloadeddata = "onloadeddata"
	oncanplay    = "oncanplay"
	ontimeupdate = "ontimeupdate"

	// Drag events
	ondrag      = "ondrag"
	ondragstart = "ondragstart"
	ondragend   = "ondragend"
	ondragover  = "ondragover"
	ondragenter = "ondragenter"
	ondragleave = "ondragleave"
	ondrop      = "ondrop"

	// Touch events
	ontouchstart  = "ontouchstart"
	ontouchend    = "ontouchend"
	ontouchmove   = "ontouchmove"
	ontouchcancel = "ontouchcancel"

	// Clipboard events
	oncopy  = "oncopy"
	oncut   = "oncut"
	onpaste = "onpaste"
)

// =============================================================================
// ARIA Attributes
// =============================================================================

const (
	role                 = "role"
	ariaLabel            = "aria-label"
	ariaLabelledby       = "aria-labelledby"
	ariaDescribedby      = "aria-describedby"
	ariaHidden           = "aria-hidden"
	ariaExpanded         = "aria-expanded"
	ariaControls         = "aria-controls"
	ariaLive             = "aria-live"
	ariaAtomic           = "aria-atomic"
	ariaBusy             = "aria-busy"
	ariaCurrent          = "aria-current"
	ariaDisabled         = "aria-disabled"
	ariaSelected         = "aria-selected"
	ariaPressed          = "aria-pressed"
	ariaChecked          = "aria-checked"
	ariaInvalid          = "aria-invalid"
	ariaRequired         = "aria-required"
	ariaReadonly         = "aria-readonly"
	ariaValuenow         = "aria-valuenow"
	ariaValuemin         = "aria-valuemin"
	ariaValuemax         = "aria-valuemax"
	ariaValuetext        = "aria-valuetext"
	ariaModal            = "aria-modal"
	ariaHaspopup         = "aria-haspopup"
	ariaOwns             = "aria-owns"
	ariaActivedescendant = "aria-activedescendant"
	ariaSort             = "aria-sort"
	ariaColcount         = "aria-colcount"
	ariaRowcount         = "aria-rowcount"
	ariaColindex         = "aria-colindex"
	ariaRowindex         = "aria-rowindex"
	ariaColspan          = "aria-colspan"
	ariaRowspan          = "aria-rowspan"
	ariaLevel            = "aria-level"
	ariaSetsize          = "aria-setsize"
	ariaPosinset         = "aria-posinset"
)

// =============================================================================
// Table Attributes
// =============================================================================

const (
	cellPadding = "cellpadding"
	cellSpacing = "cellspacing"
	border      = "border"
	colSpan     = "colspan"
	rowSpan     = "rowspan"
	scope       = "scope"
	headers     = "headers"
	align       = "align"
	vAlign      = "valign"
)

// =============================================================================
// Meta/Document Attributes
// =============================================================================

const (
	charset   = "charset"
	content   = "content"
	httpEquiv = "http-equiv"
	property  = "property"
	itemprop  = "itemprop"
	itemscope = "itemscope"
	itemtype  = "itemtype"
	xmlns     = "xmlns"
)

// =============================================================================
// Iframe Attributes
// =============================================================================

const (
	sandbox         = "sandbox"
	allow           = "allow"
	allowfullscreen = "allowfullscreen"
	loading         = "loading"
	referrerpolicy  = "referrerpolicy"
)

// =============================================================================
// Misc Attributes
// =============================================================================

const (
	datetime = "datetime"
	open_    = "open"
	cite     = "cite"
	label    = "label"
	high     = "high"
	low      = "low"
	optimum  = "optimum"
	span     = "span"
	start    = "start"
	reversed = "reversed"
	kind     = "kind"
	srclang  = "srclang"
)

// =============================================================================
// Common Values
// =============================================================================

const (
	Round = "round"
	Blank = "_blank"
)

// =============================================================================
// Attribute Constructors - Custom
// =============================================================================

func Custom(tagName, s string) Attribute {
	return regularAttribute(tagName, s)
}

func Style(styles ...css.KeyValuePair) Attribute {
	includedStyles := []css.KeyValuePair{}
	for _, style := range styles {
		if style.Include {
			includedStyles = append(includedStyles, style)
		}
	}

	if len(includedStyles) == 0 {
		return Attribute{}
	}

	return regularAttribute(style, css.PrintStyles(includedStyles))
}

func Style1(style css.KeyValuePair, styles ...css.KeyValuePair) Attribute {
	ss := append(styles, style)
	return Style(ss...)
}

// =============================================================================
// Attribute Constructors - Global
// =============================================================================

func Class(s string) Attribute {
	return regularAttribute(class, s)
}

// ClassIf is a shortcut for Class().RenderIf()
func ClassIf(s string, b bool) Attribute {
	return regularAttribute(class, s).RenderIf(b)
}

// ClassesIf takes a list of classes to apply according to a corresponding list
// of booleans. If more classes than appliers are provided, then extra
// classes are automatically applied, which is a convenient way to provide
// unconditional classes to the function.
func ClassesIf(classes []string, appliers []bool) Attribute {
	var classesToApply []string

	for i, class := range classes {
		if i < len(appliers) {
			if appliers[i] {
				classesToApply = append(classesToApply, class)
			}
		} else {
			classesToApply = append(classesToApply, class)
		}
	}

	return Class(strings.Join(classesToApply, " "))
}

func Classes(classes []string) Attribute {
	return Class(strings.Join(classes, " "))
}

func Id(s string) Attribute {
	return regularAttribute(id, s)
}

func Title(s string) Attribute {
	return regularAttribute(title, s)
}

func Lang(s string) Attribute {
	return regularAttribute(lang, s)
}

func Dir(s string) Attribute {
	return regularAttribute(dir, s)
}

func Hidden(b bool) Attribute {
	return booleanAttribute(hidden, b)
}

func Tabindex(n int) Attribute {
	return intAttribute(tabindex, n)
}

func Accesskey(s string) Attribute {
	return regularAttribute(accesskey, s)
}

func Contenteditable(b bool) Attribute {
	return booleanAttribute(contenteditable, b)
}

func Draggable(s string) Attribute {
	return regularAttribute(draggable, s)
}

func Spellcheck(s string) Attribute {
	return regularAttribute(spellcheck, s)
}

func Translate(s string) Attribute {
	return regularAttribute(translate, s)
}

func Data(suffix, s string) Attribute {
	return regularAttribute(fmt.Sprintf("data-%s", suffix), s)
}

// =============================================================================
// Attribute Constructors - Form
// =============================================================================

func Name(s string) Attribute {
	return regularAttribute(name, s)
}

func Value(s string) Attribute {
	return regularAttribute(value, s)
}

func Type(s string) Attribute {
	return regularAttribute(type_, s)
}

func Placeholder(s string) Attribute {
	return regularAttribute(placeholder, s)
}

func Required(b bool) Attribute {
	return booleanAttribute(required, b)
}

func Disabled(b bool) Attribute {
	return booleanAttribute(disabled, b)
}

func Readonly(b bool) Attribute {
	return booleanAttribute(readonly, b)
}

func Checked(b bool) Attribute {
	return booleanAttribute(checked, b)
}

func Selected(b bool) Attribute {
	return booleanAttribute(selected, b)
}

func Autofocus(b bool) Attribute {
	return booleanAttribute(autofocus, b)
}

func Autocomplete(s string) Attribute {
	return regularAttribute(autocomplete, s)
}

func Form(s string) Attribute {
	return regularAttribute(form, s)
}

func Formaction(s string) Attribute {
	return regularAttribute(formaction, s)
}

func Formmethod(s string) Attribute {
	return regularAttribute(formmethod, s)
}

func Formtarget(s string) Attribute {
	return regularAttribute(formtarget, s)
}

func Accept(s string) Attribute {
	return regularAttribute(accept, s)
}

func Inputmode(s string) Attribute {
	return regularAttribute(inputmode, s)
}

func Min(s string) Attribute {
	return regularAttribute(min, s)
}

func Max(s string) Attribute {
	return regularAttribute(max, s)
}

func MinLength(n int) Attribute {
	return intAttribute(minLength, n)
}

func MaxLength(n int) Attribute {
	return intAttribute(maxLength, n)
}

func Pattern(s string) Attribute {
	return regularAttribute(pattern, s)
}

func Step(s string) Attribute {
	return regularAttribute(step, s)
}

func Multiple(b bool) Attribute {
	return booleanAttribute(multiple, b)
}

func List(s string) Attribute {
	return regularAttribute(list, s)
}

func Rows(n int) Attribute {
	return intAttribute(rows, n)
}

func Cols(n int) Attribute {
	return intAttribute(cols, n)
}

func Wrap(s string) Attribute {
	return regularAttribute(wrap, s)
}

func Size(n int) Attribute {
	return intAttribute(size, n)
}

func For(s string) Attribute {
	return regularAttribute(for_, s)
}

func Method(s string) Attribute {
	return regularAttribute(method, s)
}

func Action(s string) Attribute {
	return regularAttribute(action, s)
}

// =============================================================================
// Attribute Constructors - Link/Navigation
// =============================================================================

func Href(s string) Attribute {
	return regularAttribute(href, s)
}

func Target(s string) Attribute {
	return regularAttribute(target, s)
}

func Rel(s string) Attribute {
	return regularAttribute(rel, s)
}

func Download(s string) Attribute {
	return regularAttribute(download, s)
}

func Hreflang(s string) Attribute {
	return regularAttribute(hreflang, s)
}

func Ping(s string) Attribute {
	return regularAttribute(ping, s)
}

// =============================================================================
// Attribute Constructors - Media
// =============================================================================

func Src(s string) Attribute {
	return regularAttribute(src, s)
}

func Alt(s string) Attribute {
	return regularAttribute(alt, s)
}

func Width(s string) Attribute {
	return regularAttribute(width, s)
}

func Height(s string) Attribute {
	return regularAttribute(height, s)
}

func Srcset(s string) Attribute {
	return regularAttribute(srcset, s)
}

func Sizes(s string) Attribute {
	return regularAttribute(sizes, s)
}

func Media(s string) Attribute {
	return regularAttribute(media, s)
}

func Autoplay(b bool) Attribute {
	return booleanAttribute(autoplay, b)
}

func Controls(b bool) Attribute {
	return booleanAttribute(controls, b)
}

func Loop(b bool) Attribute {
	return booleanAttribute(loop, b)
}

func Muted(b bool) Attribute {
	return booleanAttribute(muted, b)
}

func Poster(s string) Attribute {
	return regularAttribute(poster, s)
}

func Preload(s string) Attribute {
	return regularAttribute(preload, s)
}

func Playsinline(b bool) Attribute {
	return booleanAttribute(playsinline, b)
}

// =============================================================================
// Attribute Constructors - Script/Resource
// =============================================================================

func Defer(b bool) Attribute {
	return booleanAttribute(defer_, b)
}

func Async(b bool) Attribute {
	return booleanAttribute(async, b)
}

func As(s string) Attribute {
	return regularAttribute(as, s)
}

func Integrity(s string) Attribute {
	return regularAttribute(integrity, s)
}

func CrossOrigin(s string) Attribute {
	return regularAttribute(crossOrigin, s)
}

func Nonce(s string) Attribute {
	return regularAttribute(nonce, s)
}

// =============================================================================
// Attribute Constructors - SVG
// =============================================================================

func ViewBox(s string) Attribute {
	return regularAttribute(viewBox, s)
}

func Stroke(s string) Attribute {
	return regularAttribute(stroke, s)
}

func Fill(s string) Attribute {
	return regularAttribute(fill, s)
}

func StrokeWidth(f float64) Attribute {
	return floatAttribute(strokeWidth, f)
}

func StrokeLineCap(s string) Attribute {
	return regularAttribute(strokeLineCap, s)
}

func StrokeLineJoin(s string) Attribute {
	return regularAttribute(strokeLineJoin, s)
}

func StrokeDasharray(s string) Attribute {
	return regularAttribute(strokeDasharray, s)
}

func StrokeDashoffset(s string) Attribute {
	return regularAttribute(strokeDashoffset, s)
}

func StrokeOpacity(f float64) Attribute {
	return floatAttribute(strokeOpacity, f)
}

func FillOpacity(f float64) Attribute {
	return floatAttribute(fillOpacity, f)
}

func FillRule(s string) Attribute {
	return regularAttribute(fillRule, s)
}

func D(s string) Attribute {
	return regularAttribute(d, s)
}

func Points(s string) Attribute {
	return regularAttribute(points, s)
}

func X(f float64) Attribute {
	return floatAttribute(x, f)
}

func Y(f float64) Attribute {
	return floatAttribute(y, f)
}

func X1(f float64) Attribute {
	return floatAttribute(x1, f)
}

func Y1(f float64) Attribute {
	return floatAttribute(y1, f)
}

func X2(f float64) Attribute {
	return floatAttribute(x2, f)
}

func Y2(f float64) Attribute {
	return floatAttribute(y2, f)
}

func CX(f float64) Attribute {
	return floatAttribute(cx, f)
}

func CY(f float64) Attribute {
	return floatAttribute(cy, f)
}

func R(f float64) Attribute {
	return floatAttribute(r, f)
}

func RX(f float64) Attribute {
	return floatAttribute(rx, f)
}

func RY(f float64) Attribute {
	return floatAttribute(ry, f)
}

func Transform(s string) Attribute {
	return regularAttribute(transform, s)
}

func PathLength(f float64) Attribute {
	return floatAttribute(pathLength, f)
}

func PreserveAspectRatio(s string) Attribute {
	return regularAttribute(preserveAspectRatio, s)
}

func ClipPath(s string) Attribute {
	return regularAttribute(clipPath, s)
}

func Mask(s string) Attribute {
	return regularAttribute(mask, s)
}

func GradientUnits(s string) Attribute {
	return regularAttribute(gradientUnits, s)
}

func GradientTransform(s string) Attribute {
	return regularAttribute(gradientTransform, s)
}

func Offset(s string) Attribute {
	return regularAttribute(offset, s)
}

func StopColor(s string) Attribute {
	return regularAttribute(stopColor, s)
}

func StopOpacity(f float64) Attribute {
	return floatAttribute(stopOpacity, f)
}

func XlinkHref(s string) Attribute {
	return regularAttribute(xlinkHref, s)
}

func TextAnchor(s string) Attribute {
	return regularAttribute(textAnchor, s)
}

func DominantBaseline(s string) Attribute {
	return regularAttribute(dominantBaseline, s)
}

// =============================================================================
// Attribute Constructors - Event Handlers (Mouse)
// =============================================================================

func OnClick(s string) Attribute {
	return regularAttributeWithSingleQuotes(onclick, s)
}

func OnDblclick(s string) Attribute {
	return regularAttributeWithSingleQuotes(ondblclick, s)
}

func OnMousedown(s string) Attribute {
	return regularAttributeWithSingleQuotes(onmousedown, s)
}

func OnMouseup(s string) Attribute {
	return regularAttributeWithSingleQuotes(onmouseup, s)
}

func OnMouseover(s string) Attribute {
	return regularAttributeWithSingleQuotes(onmouseover, s)
}

func OnMouseout(s string) Attribute {
	return regularAttributeWithSingleQuotes(onmouseout, s)
}

func OnMousemove(s string) Attribute {
	return regularAttributeWithSingleQuotes(onmousemove, s)
}

func OnMouseenter(s string) Attribute {
	return regularAttributeWithSingleQuotes(onmouseenter, s)
}

func OnMouseleave(s string) Attribute {
	return regularAttributeWithSingleQuotes(onmouseleave, s)
}

func OnContextmenu(s string) Attribute {
	return regularAttributeWithSingleQuotes(oncontextmenu, s)
}

// =============================================================================
// Attribute Constructors - Event Handlers (Keyboard)
// =============================================================================

func OnKeydown(s string) Attribute {
	return regularAttributeWithSingleQuotes(onkeydown, s)
}

func OnKeyUp(s string) Attribute {
	return regularAttributeWithSingleQuotes(onkeyup, s)
}

func OnKeypress(s string) Attribute {
	return regularAttributeWithSingleQuotes(onkeypress, s)
}

// =============================================================================
// Attribute Constructors - Event Handlers (Focus)
// =============================================================================

func OnFocus(s string) Attribute {
	return regularAttributeWithSingleQuotes(onfocus, s)
}

func OnBlur(s string) Attribute {
	return regularAttributeWithSingleQuotes(onblur, s)
}

// =============================================================================
// Attribute Constructors - Event Handlers (Form)
// =============================================================================

func OnChange(s string) Attribute {
	return regularAttributeWithSingleQuotes(onchange, s)
}

func OnInput(s string) Attribute {
	return regularAttributeWithSingleQuotes(oninput, s)
}

func OnSubmit(s string) Attribute {
	return regularAttributeWithSingleQuotes(onsubmit, s)
}

func OnReset(s string) Attribute {
	return regularAttributeWithSingleQuotes(onreset, s)
}

func OnInvalid(s string) Attribute {
	return regularAttributeWithSingleQuotes(oninvalid, s)
}

// =============================================================================
// Attribute Constructors - Event Handlers (Document/Window)
// =============================================================================

func OnLoad(s string) Attribute {
	return regularAttributeWithSingleQuotes(onload, s)
}

func OnError(s string) Attribute {
	return regularAttributeWithSingleQuotes(onerror, s)
}

func OnScroll(s string) Attribute {
	return regularAttributeWithSingleQuotes(onscroll, s)
}

func OnResize(s string) Attribute {
	return regularAttributeWithSingleQuotes(onresize, s)
}

// =============================================================================
// Attribute Constructors - Event Handlers (Media)
// =============================================================================

func OnPlay(s string) Attribute {
	return regularAttributeWithSingleQuotes(onplay, s)
}

func OnPause(s string) Attribute {
	return regularAttributeWithSingleQuotes(onpause, s)
}

func OnEnded(s string) Attribute {
	return regularAttributeWithSingleQuotes(onended, s)
}

func OnLoadeddata(s string) Attribute {
	return regularAttributeWithSingleQuotes(onloadeddata, s)
}

func OnCanplay(s string) Attribute {
	return regularAttributeWithSingleQuotes(oncanplay, s)
}

func OnTimeupdate(s string) Attribute {
	return regularAttributeWithSingleQuotes(ontimeupdate, s)
}

// =============================================================================
// Attribute Constructors - Event Handlers (Drag)
// =============================================================================

func OnDrag(s string) Attribute {
	return regularAttributeWithSingleQuotes(ondrag, s)
}

func OnDragstart(s string) Attribute {
	return regularAttributeWithSingleQuotes(ondragstart, s)
}

func OnDragend(s string) Attribute {
	return regularAttributeWithSingleQuotes(ondragend, s)
}

func OnDragover(s string) Attribute {
	return regularAttributeWithSingleQuotes(ondragover, s)
}

func OnDragenter(s string) Attribute {
	return regularAttributeWithSingleQuotes(ondragenter, s)
}

func OnDragleave(s string) Attribute {
	return regularAttributeWithSingleQuotes(ondragleave, s)
}

func OnDrop(s string) Attribute {
	return regularAttributeWithSingleQuotes(ondrop, s)
}

// =============================================================================
// Attribute Constructors - Event Handlers (Touch)
// =============================================================================

func OnTouchstart(s string) Attribute {
	return regularAttributeWithSingleQuotes(ontouchstart, s)
}

func OnTouchend(s string) Attribute {
	return regularAttributeWithSingleQuotes(ontouchend, s)
}

func OnTouchmove(s string) Attribute {
	return regularAttributeWithSingleQuotes(ontouchmove, s)
}

func OnTouchcancel(s string) Attribute {
	return regularAttributeWithSingleQuotes(ontouchcancel, s)
}

// =============================================================================
// Attribute Constructors - Event Handlers (Clipboard)
// =============================================================================

func OnCopy(s string) Attribute {
	return regularAttributeWithSingleQuotes(oncopy, s)
}

func OnCut(s string) Attribute {
	return regularAttributeWithSingleQuotes(oncut, s)
}

func OnPaste(s string) Attribute {
	return regularAttributeWithSingleQuotes(onpaste, s)
}

// =============================================================================
// Attribute Constructors - ARIA
// =============================================================================

func Role(s string) Attribute {
	return regularAttribute(role, s)
}

func AriaLabel(s string) Attribute {
	return regularAttribute(ariaLabel, s)
}

func AriaLabelledby(s string) Attribute {
	return regularAttribute(ariaLabelledby, s)
}

func AriaDescribedby(s string) Attribute {
	return regularAttribute(ariaDescribedby, s)
}

func AriaHidden(b bool) Attribute {
	return regularAttribute(ariaHidden, fmt.Sprintf("%t", b))
}

func AriaExpanded(b bool) Attribute {
	return regularAttribute(ariaExpanded, fmt.Sprintf("%t", b))
}

func AriaControls(s string) Attribute {
	return regularAttribute(ariaControls, s)
}

func AriaLive(s string) Attribute {
	return regularAttribute(ariaLive, s)
}

func AriaAtomic(b bool) Attribute {
	return regularAttribute(ariaAtomic, fmt.Sprintf("%t", b))
}

func AriaBusy(b bool) Attribute {
	return regularAttribute(ariaBusy, fmt.Sprintf("%t", b))
}

func AriaCurrent(s string) Attribute {
	return regularAttribute(ariaCurrent, s)
}

func AriaDisabled(b bool) Attribute {
	return regularAttribute(ariaDisabled, fmt.Sprintf("%t", b))
}

func AriaSelected(b bool) Attribute {
	return regularAttribute(ariaSelected, fmt.Sprintf("%t", b))
}

func AriaPressed(s string) Attribute {
	return regularAttribute(ariaPressed, s)
}

func AriaChecked(s string) Attribute {
	return regularAttribute(ariaChecked, s)
}

func AriaInvalid(s string) Attribute {
	return regularAttribute(ariaInvalid, s)
}

func AriaRequired(b bool) Attribute {
	return regularAttribute(ariaRequired, fmt.Sprintf("%t", b))
}

func AriaReadonly(b bool) Attribute {
	return regularAttribute(ariaReadonly, fmt.Sprintf("%t", b))
}

func AriaValuenow(f float64) Attribute {
	return floatAttribute(ariaValuenow, f)
}

func AriaValuemin(f float64) Attribute {
	return floatAttribute(ariaValuemin, f)
}

func AriaValuemax(f float64) Attribute {
	return floatAttribute(ariaValuemax, f)
}

func AriaValuetext(s string) Attribute {
	return regularAttribute(ariaValuetext, s)
}

func AriaModal(b bool) Attribute {
	return regularAttribute(ariaModal, fmt.Sprintf("%t", b))
}

func AriaHaspopup(s string) Attribute {
	return regularAttribute(ariaHaspopup, s)
}

func AriaOwns(s string) Attribute {
	return regularAttribute(ariaOwns, s)
}

func AriaActivedescendant(s string) Attribute {
	return regularAttribute(ariaActivedescendant, s)
}

func AriaSort(s string) Attribute {
	return regularAttribute(ariaSort, s)
}

func AriaColcount(n int) Attribute {
	return intAttribute(ariaColcount, n)
}

func AriaRowcount(n int) Attribute {
	return intAttribute(ariaRowcount, n)
}

func AriaColindex(n int) Attribute {
	return intAttribute(ariaColindex, n)
}

func AriaRowindex(n int) Attribute {
	return intAttribute(ariaRowindex, n)
}

func AriaColspan(n int) Attribute {
	return intAttribute(ariaColspan, n)
}

func AriaRowspan(n int) Attribute {
	return intAttribute(ariaRowspan, n)
}

func AriaLevel(n int) Attribute {
	return intAttribute(ariaLevel, n)
}

func AriaSetsize(n int) Attribute {
	return intAttribute(ariaSetsize, n)
}

func AriaPosinset(n int) Attribute {
	return intAttribute(ariaPosinset, n)
}

// =============================================================================
// Attribute Constructors - Table
// =============================================================================

func CellPadding(pixels int) Attribute {
	return intAttribute(cellPadding, pixels)
}

func CellSpacing(pixels int) Attribute {
	return intAttribute(cellSpacing, pixels)
}

func Border(pixels int) Attribute {
	return intAttribute(border, pixels)
}

func ColSpan(i int) Attribute {
	return intAttribute(colSpan, i)
}

func RowSpan(i int) Attribute {
	return intAttribute(rowSpan, i)
}

func Scope(s string) Attribute {
	return regularAttribute(scope, s)
}

func Headers(s string) Attribute {
	return regularAttribute(headers, s)
}

func Align(s string) Attribute {
	return regularAttribute(align, s)
}

func VAlign(s string) Attribute {
	return regularAttribute(vAlign, s)
}

// =============================================================================
// Attribute Constructors - Meta/Document
// =============================================================================

func Charset(s string) Attribute {
	return regularAttribute(charset, s)
}

func Content(s string) Attribute {
	return regularAttribute(content, s)
}

func HttpEquiv(s string) Attribute {
	return regularAttribute(httpEquiv, s)
}

func Property(s string) Attribute {
	return regularAttribute(property, s)
}

func Itemprop(s string) Attribute {
	return regularAttribute(itemprop, s)
}

func Itemscope(b bool) Attribute {
	return booleanAttribute(itemscope, b)
}

func Itemtype(s string) Attribute {
	return regularAttribute(itemtype, s)
}

func Xmlns(s string) Attribute {
	return regularAttribute(xmlns, s)
}

// =============================================================================
// Attribute Constructors - Iframe
// =============================================================================

func Sandbox(s string) Attribute {
	return regularAttribute(sandbox, s)
}

func Allow(s string) Attribute {
	return regularAttribute(allow, s)
}

func Allowfullscreen(b bool) Attribute {
	return booleanAttribute(allowfullscreen, b)
}

func Loading(s string) Attribute {
	return regularAttribute(loading, s)
}

func Referrerpolicy(s string) Attribute {
	return regularAttribute(referrerpolicy, s)
}

// =============================================================================
// Attribute Constructors - Misc
// =============================================================================

func Datetime(s string) Attribute {
	return regularAttribute(datetime, s)
}

func Open(b bool) Attribute {
	return booleanAttribute(open_, b)
}

func Cite(s string) Attribute {
	return regularAttribute(cite, s)
}

func Label(s string) Attribute {
	return regularAttribute(label, s)
}

func High(f float64) Attribute {
	return floatAttribute(high, f)
}

func Low(f float64) Attribute {
	return floatAttribute(low, f)
}

func Optimum(f float64) Attribute {
	return floatAttribute(optimum, f)
}

func Span(n int) Attribute {
	return intAttribute(span, n)
}

func Start(n int) Attribute {
	return intAttribute(start, n)
}

func Reversed(b bool) Attribute {
	return booleanAttribute(reversed, b)
}

func Kind(s string) Attribute {
	return regularAttribute(kind, s)
}

func Srclang(s string) Attribute {
	return regularAttribute(srclang, s)
}
