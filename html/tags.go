package html

import (
	"fmt"

	"github.com/jpincas/go-tea/attributes"
	"github.com/jpincas/go-tea/css"
)

// =============================================================================
// Tag Constants
// =============================================================================

const (
	// Document structure
	html     = "html"
	head     = "head"
	body     = "body"
	title    = "title"
	meta     = "meta"
	link     = "link"
	script   = "script"
	noScript = "noscript"
	style    = "style"
	base     = "base"

	// Sectioning
	header  = "header"
	footer  = "footer"
	main    = "main"
	nav     = "nav"
	section = "section"
	article = "article"
	aside   = "aside"
	address = "address"
	hgroup  = "hgroup"

	// Headings
	h1 = "h1"
	h2 = "h2"
	h3 = "h3"
	h4 = "h4"
	h5 = "h5"
	h6 = "h6"

	// Content grouping
	div        = "div"
	p          = "p"
	hr         = "hr"
	pre        = "pre"
	blockquote = "blockquote"
	figure     = "figure"
	figcaption = "figcaption"

	// Text semantics
	a      = "a"
	em     = "em"
	strong = "strong"
	small  = "small"
	s      = "s"
	cite   = "cite"
	q      = "q"
	dfn    = "dfn"
	abbr   = "abbr"
	data   = "data"
	time_  = "time"
	code   = "code"
	var_   = "var"
	samp   = "samp"
	kbd    = "kbd"
	sub    = "sub"
	sup    = "sup"
	i      = "i"
	b      = "b"
	u      = "u"
	mark   = "mark"
	ruby   = "ruby"
	rt     = "rt"
	rp     = "rp"
	bdi    = "bdi"
	bdo    = "bdo"
	span   = "span"
	br     = "br"
	wbr    = "wbr"

	// Lists
	ul = "ul"
	ol = "ol"
	li = "li"
	dl = "dl"
	dt = "dt"
	dd = "dd"

	// Tables
	table    = "table"
	caption  = "caption"
	colgroup = "colgroup"
	col      = "col"
	thead    = "thead"
	tbody    = "tbody"
	tfoot    = "tfoot"
	tr       = "tr"
	th       = "th"
	td       = "td"

	// Forms
	form      = "form"
	fieldset  = "fieldset"
	legend    = "legend"
	label     = "label"
	input     = "input"
	button    = "button"
	selectTag = "select"
	datalist  = "datalist"
	optgroup  = "optgroup"
	option    = "option"
	textarea  = "textarea"
	output    = "output"
	progress  = "progress"
	meter     = "meter"

	// Interactive
	details = "details"
	summary = "summary"
	dialog  = "dialog"

	// Media
	img     = "img"
	picture = "picture"
	source  = "source"
	video   = "video"
	audio   = "audio"
	track   = "track"
	canvas  = "canvas"

	// Embedded
	iframe = "iframe"
	embed  = "embed"
	object = "object"
	param  = "param"

	// SVG
	svg            = "svg"
	g              = "g"
	defs           = "defs"
	symbol         = "symbol"
	use            = "use"
	path           = "path"
	rect           = "rect"
	circle         = "circle"
	ellipse        = "ellipse"
	line           = "line"
	polyline       = "polyline"
	polygon        = "polygon"
	text           = "text"
	tspan          = "tspan"
	textPath       = "textPath"
	image          = "image"
	clipPathTag    = "clipPath"
	maskTag        = "mask"
	linearGradient = "linearGradient"
	radialGradient = "radialGradient"
	stop           = "stop"
	pattern        = "pattern"
	filter         = "filter"
	feGaussianBlur = "feGaussianBlur"
	feOffset       = "feOffset"
	feMerge        = "feMerge"
	feMergeNode    = "feMergeNode"
	feBlend        = "feBlend"
	feColorMatrix  = "feColorMatrix"
	feDropShadow   = "feDropShadow"
	foreignObject  = "foreignObject"

	// Template
	templateTag = "template"
	slot        = "slot"
)

// Special tag for text content
const textTag = "text"

// =============================================================================
// Document Structure
// =============================================================================

func Html(attrs attributes.Attributes, elements ...Element) Element {
	return basicTag(html, attrs, elements)
}

func Head(attrs attributes.Attributes, elements ...Element) Element {
	return basicTag(head, attrs, elements)
}

func Body(attrs attributes.Attributes, elements ...Element) Element {
	return basicTag(body, attrs, elements)
}

func Title(attrs attributes.Attributes, elements ...Element) Element {
	return basicTag(title, attrs, elements)
}

func Meta(attrs attributes.Attributes) Element {
	return selfClosingTag(meta, attrs)
}

func Link(attrs attributes.Attributes) Element {
	return selfClosingTag(link, attrs)
}

func Script(attrs attributes.Attributes, elements ...Element) Element {
	return basicTag(script, attrs, elements)
}

func NoScript(attrs attributes.Attributes, elements ...Element) Element {
	return basicTag(noScript, attrs, elements)
}

func Style(attrs attributes.Attributes, elements ...Element) Element {
	return basicTag(style, attrs, elements)
}

func Base(attrs attributes.Attributes) Element {
	return selfClosingTag(base, attrs)
}

// =============================================================================
// Sectioning
// =============================================================================

func Header(attrs attributes.Attributes, elements ...Element) Element {
	return basicTag(header, attrs, elements)
}

func Footer(attrs attributes.Attributes, elements ...Element) Element {
	return basicTag(footer, attrs, elements)
}

func Main(attrs attributes.Attributes, elements ...Element) Element {
	return basicTag(main, attrs, elements)
}

func Nav(attrs attributes.Attributes, elements ...Element) Element {
	return basicTag(nav, attrs, elements)
}

func Section(attrs attributes.Attributes, elements ...Element) Element {
	return basicTag(section, attrs, elements)
}

func Article(attrs attributes.Attributes, elements ...Element) Element {
	return basicTag(article, attrs, elements)
}

func Aside(attrs attributes.Attributes, elements ...Element) Element {
	return basicTag(aside, attrs, elements)
}

func Address(attrs attributes.Attributes, elements ...Element) Element {
	return basicTag(address, attrs, elements)
}

func Hgroup(attrs attributes.Attributes, elements ...Element) Element {
	return basicTag(hgroup, attrs, elements)
}

// =============================================================================
// Headings
// =============================================================================

func H1(attrs attributes.Attributes, elements ...Element) Element {
	return basicTag(h1, attrs, elements)
}

func H2(attrs attributes.Attributes, elements ...Element) Element {
	return basicTag(h2, attrs, elements)
}

func H3(attrs attributes.Attributes, elements ...Element) Element {
	return basicTag(h3, attrs, elements)
}

func H4(attrs attributes.Attributes, elements ...Element) Element {
	return basicTag(h4, attrs, elements)
}

func H5(attrs attributes.Attributes, elements ...Element) Element {
	return basicTag(h5, attrs, elements)
}

func H6(attrs attributes.Attributes, elements ...Element) Element {
	return basicTag(h6, attrs, elements)
}

// =============================================================================
// Content Grouping
// =============================================================================

func Div(attrs attributes.Attributes, elements ...Element) Element {
	return basicTag(div, attrs, elements)
}

func P(attrs attributes.Attributes, elements ...Element) Element {
	return basicTag(p, attrs, elements)
}

func Hr(attrs attributes.Attributes) Element {
	return selfClosingTag(hr, attrs)
}

func Pre(attrs attributes.Attributes, elements ...Element) Element {
	return basicTag(pre, attrs, elements)
}

func Blockquote(attrs attributes.Attributes, elements ...Element) Element {
	return basicTag(blockquote, attrs, elements)
}

func Figure(attrs attributes.Attributes, elements ...Element) Element {
	return basicTag(figure, attrs, elements)
}

func Figcaption(attrs attributes.Attributes, elements ...Element) Element {
	return basicTag(figcaption, attrs, elements)
}

// =============================================================================
// Text Semantics
// =============================================================================

func A(attrs attributes.Attributes, elements ...Element) Element {
	return basicTag(a, attrs, elements)
}

func Em(attrs attributes.Attributes, elements ...Element) Element {
	return basicTag(em, attrs, elements)
}

func Strong(attrs attributes.Attributes, elements ...Element) Element {
	return basicTag(strong, attrs, elements)
}

func Small(attrs attributes.Attributes, elements ...Element) Element {
	return basicTag(small, attrs, elements)
}

func S(attrs attributes.Attributes, elements ...Element) Element {
	return basicTag(s, attrs, elements)
}

func Cite(attrs attributes.Attributes, elements ...Element) Element {
	return basicTag(cite, attrs, elements)
}

func Q(attrs attributes.Attributes, elements ...Element) Element {
	return basicTag(q, attrs, elements)
}

func Dfn(attrs attributes.Attributes, elements ...Element) Element {
	return basicTag(dfn, attrs, elements)
}

func Abbr(attrs attributes.Attributes, elements ...Element) Element {
	return basicTag(abbr, attrs, elements)
}

func Data(attrs attributes.Attributes, elements ...Element) Element {
	return basicTag(data, attrs, elements)
}

func Time(attrs attributes.Attributes, elements ...Element) Element {
	return basicTag(time_, attrs, elements)
}

func Code(attrs attributes.Attributes, elements ...Element) Element {
	return basicTag(code, attrs, elements)
}

func Var(attrs attributes.Attributes, elements ...Element) Element {
	return basicTag(var_, attrs, elements)
}

func Samp(attrs attributes.Attributes, elements ...Element) Element {
	return basicTag(samp, attrs, elements)
}

func Kbd(attrs attributes.Attributes, elements ...Element) Element {
	return basicTag(kbd, attrs, elements)
}

func Sub(attrs attributes.Attributes, elements ...Element) Element {
	return basicTag(sub, attrs, elements)
}

func Sup(attrs attributes.Attributes, elements ...Element) Element {
	return basicTag(sup, attrs, elements)
}

func I(attrs attributes.Attributes, elements ...Element) Element {
	return basicTag(i, attrs, elements)
}

func B(attrs attributes.Attributes, elements ...Element) Element {
	return basicTag(b, attrs, elements)
}

func U(attrs attributes.Attributes, elements ...Element) Element {
	return basicTag(u, attrs, elements)
}

func Mark(attrs attributes.Attributes, elements ...Element) Element {
	return basicTag(mark, attrs, elements)
}

func Ruby(attrs attributes.Attributes, elements ...Element) Element {
	return basicTag(ruby, attrs, elements)
}

func Rt(attrs attributes.Attributes, elements ...Element) Element {
	return basicTag(rt, attrs, elements)
}

func Rp(attrs attributes.Attributes, elements ...Element) Element {
	return basicTag(rp, attrs, elements)
}

func Bdi(attrs attributes.Attributes, elements ...Element) Element {
	return basicTag(bdi, attrs, elements)
}

func Bdo(attrs attributes.Attributes, elements ...Element) Element {
	return basicTag(bdo, attrs, elements)
}

func Span(attrs attributes.Attributes, elements ...Element) Element {
	return basicTag(span, attrs, elements)
}

func Br(attrs attributes.Attributes) Element {
	return selfClosingTag(br, attrs)
}

func Wbr(attrs attributes.Attributes) Element {
	return selfClosingTag(wbr, attrs)
}

// =============================================================================
// Lists
// =============================================================================

func Ul(attrs attributes.Attributes, elements ...Element) Element {
	return basicTag(ul, attrs, elements)
}

func Ol(attrs attributes.Attributes, elements ...Element) Element {
	return basicTag(ol, attrs, elements)
}

func Li(attrs attributes.Attributes, elements ...Element) Element {
	return basicTag(li, attrs, elements)
}

func Dl(attrs attributes.Attributes, elements ...Element) Element {
	return basicTag(dl, attrs, elements)
}

func Dt(attrs attributes.Attributes, elements ...Element) Element {
	return basicTag(dt, attrs, elements)
}

func Dd(attrs attributes.Attributes, elements ...Element) Element {
	return basicTag(dd, attrs, elements)
}

// =============================================================================
// Tables
// =============================================================================

func Table(attrs attributes.Attributes, elements ...Element) Element {
	return basicTag(table, attrs, elements)
}

func Caption(attrs attributes.Attributes, elements ...Element) Element {
	return basicTag(caption, attrs, elements)
}

func Colgroup(attrs attributes.Attributes, elements ...Element) Element {
	return basicTag(colgroup, attrs, elements)
}

func Col(attrs attributes.Attributes) Element {
	return selfClosingTag(col, attrs)
}

func THead(attrs attributes.Attributes, elements ...Element) Element {
	return basicTag(thead, attrs, elements)
}

func TBody(attrs attributes.Attributes, elements ...Element) Element {
	return basicTag(tbody, attrs, elements)
}

func TFoot(attrs attributes.Attributes, elements ...Element) Element {
	return basicTag(tfoot, attrs, elements)
}

func Tr(attrs attributes.Attributes, elements ...Element) Element {
	return basicTag(tr, attrs, elements)
}

func Th(attrs attributes.Attributes, elements ...Element) Element {
	return basicTag(th, attrs, elements)
}

func Td(attrs attributes.Attributes, elements ...Element) Element {
	return basicTag(td, attrs, elements)
}

// =============================================================================
// Forms
// =============================================================================

func Form(attrs attributes.Attributes, elements ...Element) Element {
	return basicTag(form, attrs, elements)
}

func Fieldset(attrs attributes.Attributes, elements ...Element) Element {
	return basicTag(fieldset, attrs, elements)
}

func Legend(attrs attributes.Attributes, elements ...Element) Element {
	return basicTag(legend, attrs, elements)
}

func Label(attrs attributes.Attributes, elements ...Element) Element {
	return basicTag(label, attrs, elements)
}

func Input(attrs attributes.Attributes) Element {
	return selfClosingTag(input, attrs)
}

func Button(attrs attributes.Attributes, elements ...Element) Element {
	return basicTag(button, attrs, elements)
}

func Select(attrs attributes.Attributes, elements ...Element) Element {
	return basicTag(selectTag, attrs, elements)
}

func Datalist(attrs attributes.Attributes, elements ...Element) Element {
	return basicTag(datalist, attrs, elements)
}

func Optgroup(attrs attributes.Attributes, elements ...Element) Element {
	return basicTag(optgroup, attrs, elements)
}

func Option(attrs attributes.Attributes, elements ...Element) Element {
	return basicTag(option, attrs, elements)
}

func TextArea(attrs attributes.Attributes, elements ...Element) Element {
	return basicTag(textarea, attrs, elements)
}

func Output(attrs attributes.Attributes, elements ...Element) Element {
	return basicTag(output, attrs, elements)
}

func Progress(attrs attributes.Attributes, elements ...Element) Element {
	return basicTag(progress, attrs, elements)
}

func Meter(attrs attributes.Attributes, elements ...Element) Element {
	return basicTag(meter, attrs, elements)
}

// =============================================================================
// Interactive
// =============================================================================

func Details(attrs attributes.Attributes, elements ...Element) Element {
	return basicTag(details, attrs, elements)
}

func Summary(attrs attributes.Attributes, elements ...Element) Element {
	return basicTag(summary, attrs, elements)
}

func Dialog(attrs attributes.Attributes, elements ...Element) Element {
	return basicTag(dialog, attrs, elements)
}

// =============================================================================
// Media
// =============================================================================

func Img(attrs attributes.Attributes) Element {
	return selfClosingTag(img, attrs)
}

func Picture(attrs attributes.Attributes, elements ...Element) Element {
	return basicTag(picture, attrs, elements)
}

func Source(attrs attributes.Attributes) Element {
	return selfClosingTag(source, attrs)
}

func Video(attrs attributes.Attributes, elements ...Element) Element {
	return basicTag(video, attrs, elements)
}

func Audio(attrs attributes.Attributes, elements ...Element) Element {
	return basicTag(audio, attrs, elements)
}

func Track(attrs attributes.Attributes) Element {
	return selfClosingTag(track, attrs)
}

func Canvas(attrs attributes.Attributes, elements ...Element) Element {
	return basicTag(canvas, attrs, elements)
}

// =============================================================================
// Embedded
// =============================================================================

func Iframe(attrs attributes.Attributes, elements ...Element) Element {
	return basicTag(iframe, attrs, elements)
}

func Embed(attrs attributes.Attributes) Element {
	return selfClosingTag(embed, attrs)
}

func Object(attrs attributes.Attributes, elements ...Element) Element {
	return basicTag(object, attrs, elements)
}

func Param(attrs attributes.Attributes) Element {
	return selfClosingTag(param, attrs)
}

// =============================================================================
// SVG
// =============================================================================

func SVG(attrs attributes.Attributes, elements ...Element) Element {
	return basicTag(svg, attrs, elements)
}

func G(attrs attributes.Attributes, elements ...Element) Element {
	return basicTag(g, attrs, elements)
}

func Defs(attrs attributes.Attributes, elements ...Element) Element {
	return basicTag(defs, attrs, elements)
}

func Symbol(attrs attributes.Attributes, elements ...Element) Element {
	return basicTag(symbol, attrs, elements)
}

func Use(attrs attributes.Attributes) Element {
	return basicTag(use, attrs, Elements{})
}

func Path(attrs attributes.Attributes) Element {
	return basicTag(path, attrs, Elements{})
}

func Rect(attrs attributes.Attributes) Element {
	return basicTag(rect, attrs, Elements{})
}

func Circle(attrs attributes.Attributes) Element {
	return basicTag(circle, attrs, Elements{})
}

func Ellipse(attrs attributes.Attributes) Element {
	return basicTag(ellipse, attrs, Elements{})
}

func Line(attrs attributes.Attributes) Element {
	return basicTag(line, attrs, Elements{})
}

func Polyline(attrs attributes.Attributes) Element {
	return basicTag(polyline, attrs, Elements{})
}

func Polygon(attrs attributes.Attributes) Element {
	return basicTag(polygon, attrs, Elements{})
}

// SVGText creates an SVG <text> element (named differently to avoid conflict with Text())
func SVGText(attrs attributes.Attributes, elements ...Element) Element {
	return basicTag(text, attrs, elements)
}

func TSpan(attrs attributes.Attributes, elements ...Element) Element {
	return basicTag(tspan, attrs, elements)
}

func TextPath(attrs attributes.Attributes, elements ...Element) Element {
	return basicTag(textPath, attrs, elements)
}

func Image(attrs attributes.Attributes) Element {
	return basicTag(image, attrs, Elements{})
}

func ClipPathEl(attrs attributes.Attributes, elements ...Element) Element {
	return basicTag(clipPathTag, attrs, elements)
}

func MaskEl(attrs attributes.Attributes, elements ...Element) Element {
	return basicTag(maskTag, attrs, elements)
}

func LinearGradient(attrs attributes.Attributes, elements ...Element) Element {
	return basicTag(linearGradient, attrs, elements)
}

func RadialGradient(attrs attributes.Attributes, elements ...Element) Element {
	return basicTag(radialGradient, attrs, elements)
}

func Stop(attrs attributes.Attributes) Element {
	return basicTag(stop, attrs, Elements{})
}

func Pattern(attrs attributes.Attributes, elements ...Element) Element {
	return basicTag(pattern, attrs, elements)
}

func Filter(attrs attributes.Attributes, elements ...Element) Element {
	return basicTag(filter, attrs, elements)
}

func FeGaussianBlur(attrs attributes.Attributes) Element {
	return basicTag(feGaussianBlur, attrs, Elements{})
}

func FeOffset(attrs attributes.Attributes) Element {
	return basicTag(feOffset, attrs, Elements{})
}

func FeMerge(attrs attributes.Attributes, elements ...Element) Element {
	return basicTag(feMerge, attrs, elements)
}

func FeMergeNode(attrs attributes.Attributes) Element {
	return basicTag(feMergeNode, attrs, Elements{})
}

func FeBlend(attrs attributes.Attributes) Element {
	return basicTag(feBlend, attrs, Elements{})
}

func FeColorMatrix(attrs attributes.Attributes) Element {
	return basicTag(feColorMatrix, attrs, Elements{})
}

func FeDropShadow(attrs attributes.Attributes) Element {
	return basicTag(feDropShadow, attrs, Elements{})
}

func ForeignObject(attrs attributes.Attributes, elements ...Element) Element {
	return basicTag(foreignObject, attrs, elements)
}

// =============================================================================
// Template
// =============================================================================

func Template(attrs attributes.Attributes, elements ...Element) Element {
	return basicTag(templateTag, attrs, elements)
}

func Slot(attrs attributes.Attributes, elements ...Element) Element {
	return basicTag(slot, attrs, elements)
}

// =============================================================================
// Text Content
// =============================================================================

func Text(i interface{}) Element {
	return Element{
		Tag:  textTag,
		Text: fmt.Sprintf("%v", i),
	}
}

func Textf(template string, i interface{}) Element {
	return Element{
		Tag:  textTag,
		Text: fmt.Sprintf(template, i),
	}
}

// =============================================================================
// Helper Functions
// =============================================================================

func BoldText(i interface{}) Element {
	return Span(
		attributes.Attrs(
			attributes.Style(
				css.FontWeight(css.Bold),
			),
		),
		Text(i),
	)
}

func ItalicText(i interface{}) Element {
	return Span(
		attributes.Attrs(
			attributes.Style(
				css.FontStyle(css.Italic),
			),
		),
		Text(i),
	)
}

// Nothing generates a blank element. The only reason we have the arguments
// is to make the function type signature the same as the other construction
// functions
func Nothing(_ attributes.Attributes, _ ...Element) Element {
	return Element{}
}
