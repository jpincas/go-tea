package css

import (
	"fmt"
	"strings"
)

// =============================================================================
// Box Model Constants
// =============================================================================

const (
	// Sizing
	width     = "width"
	height    = "height"
	minWidth  = "min-width"
	minHeight = "min-height"
	maxWidth  = "max-width"
	maxHeight = "max-height"
	boxSizing = "box-sizing"

	// Margin
	margin       = "margin"
	marginTop    = "margin-top"
	marginRight  = "margin-right"
	marginBottom = "margin-bottom"
	marginLeft   = "margin-left"

	// Padding
	padding       = "padding"
	paddingTop    = "padding-top"
	paddingRight  = "padding-right"
	paddingBottom = "padding-bottom"
	paddingLeft   = "padding-left"
)

// =============================================================================
// Positioning Constants
// =============================================================================

const (
	position = "position"
	top      = "top"
	right    = "right"
	bottom   = "bottom"
	left     = "left"
	zIndex   = "z-index"
	float_   = "float"
	clear    = "clear"
)

// =============================================================================
// Display & Visibility Constants
// =============================================================================

const (
	display    = "display"
	visibility = "visibility"
	opacity    = "opacity"
	overflow   = "overflow"
	overflowX  = "overflow-x"
	overflowY  = "overflow-y"
)

// =============================================================================
// Flexbox Constants
// =============================================================================

const (
	flexDirection  = "flex-direction"
	flexWrap       = "flex-wrap"
	flexFlow       = "flex-flow"
	justifyContent = "justify-content"
	alignItems     = "align-items"
	alignContent   = "align-content"
	alignSelf      = "align-self"
	flex           = "flex"
	flexGrow       = "flex-grow"
	flexShrink     = "flex-shrink"
	flexBasis      = "flex-basis"
	order          = "order"
	gap            = "gap"
	rowGap         = "row-gap"
	columnGap      = "column-gap"
)

// =============================================================================
// Grid Constants
// =============================================================================

const (
	gridTemplateColumns = "grid-template-columns"
	gridTemplateRows    = "grid-template-rows"
	gridTemplateAreas   = "grid-template-areas"
	gridTemplate        = "grid-template"
	gridColumn          = "grid-column"
	gridColumnStart     = "grid-column-start"
	gridColumnEnd       = "grid-column-end"
	gridRow             = "grid-row"
	gridRowStart        = "grid-row-start"
	gridRowEnd          = "grid-row-end"
	gridArea            = "grid-area"
	gridAutoFlow        = "grid-auto-flow"
	gridAutoColumns     = "grid-auto-columns"
	gridAutoRows        = "grid-auto-rows"
	justifyItems        = "justify-items"
	justifySelf         = "justify-self"
	placeContent        = "place-content"
	placeItems          = "place-items"
	placeSelf           = "place-self"
)

// =============================================================================
// Typography Constants
// =============================================================================

const (
	fontFamily     = "font-family"
	fontSize       = "font-size"
	fontWeight     = "font-weight"
	fontStyle      = "font-style"
	fontVariant    = "font-variant"
	lineHeight     = "line-height"
	letterSpacing  = "letter-spacing"
	wordSpacing    = "word-spacing"
	textAlign      = "text-align"
	textDecoration = "text-decoration"
	textTransform  = "text-transform"
	textIndent     = "text-indent"
	textShadow     = "text-shadow"
	textOverflow   = "text-overflow"
	whiteSpace     = "white-space"
	wordBreak      = "word-break"
	wordWrap       = "word-wrap"
	overflowWrap   = "overflow-wrap"
	verticalAlign  = "vertical-align"
	direction      = "direction"
	unicodeBidi    = "unicode-bidi"
)

// =============================================================================
// Color & Background Constants
// =============================================================================

const (
	color                = "color"
	backgroundColor      = "background-color"
	backgroundImage      = "background-image"
	backgroundRepeat     = "background-repeat"
	backgroundPosition   = "background-position"
	backgroundSize       = "background-size"
	backgroundAttachment = "background-attachment"
	backgroundClip       = "background-clip"
	backgroundOrigin     = "background-origin"
	background           = "background"
)

// =============================================================================
// Border Constants
// =============================================================================

const (
	border                  = "border"
	borderWidth             = "border-width"
	borderStyle             = "border-style"
	borderColor             = "border-color"
	borderTop               = "border-top"
	borderRight             = "border-right"
	borderBottom            = "border-bottom"
	borderLeft              = "border-left"
	borderRadius            = "border-radius"
	borderTopLeftRadius     = "border-top-left-radius"
	borderTopRightRadius    = "border-top-right-radius"
	borderBottomLeftRadius  = "border-bottom-left-radius"
	borderBottomRightRadius = "border-bottom-right-radius"
	borderCollapse          = "border-collapse"
	borderSpacing           = "border-spacing"
)

// =============================================================================
// Outline Constants
// =============================================================================

const (
	outline       = "outline"
	outlineWidth  = "outline-width"
	outlineStyle  = "outline-style"
	outlineColor  = "outline-color"
	outlineOffset = "outline-offset"
)

// =============================================================================
// Box Shadow & Effects Constants
// =============================================================================

const (
	boxShadow      = "box-shadow"
	filter_        = "filter"
	backdropFilter = "backdrop-filter"
	clipPath       = "clip-path"
	objectFit      = "object-fit"
	objectPosition = "object-position"
)

// =============================================================================
// Transforms Constants
// =============================================================================

const (
	transform          = "transform"
	transformOrigin    = "transform-origin"
	transformStyle     = "transform-style"
	perspective        = "perspective"
	perspectiveOrigin  = "perspective-origin"
	backfaceVisibility = "backface-visibility"
)

// =============================================================================
// Transitions & Animations Constants
// =============================================================================

const (
	transition               = "transition"
	transitionProperty       = "transition-property"
	transitionDuration       = "transition-duration"
	transitionTimingFunction = "transition-timing-function"
	transitionDelay          = "transition-delay"
	animation                = "animation"
	animationName            = "animation-name"
	animationDuration        = "animation-duration"
	animationTimingFunction  = "animation-timing-function"
	animationDelay           = "animation-delay"
	animationIterationCount  = "animation-iteration-count"
	animationDirection       = "animation-direction"
	animationFillMode        = "animation-fill-mode"
	animationPlayState       = "animation-play-state"
)

// =============================================================================
// List Constants
// =============================================================================

const (
	listStyle         = "list-style"
	listStyleType     = "list-style-type"
	listStylePosition = "list-style-position"
	listStyleImage    = "list-style-image"
)

// =============================================================================
// Table Constants
// =============================================================================

const (
	tableLayout = "table-layout"
	captionSide = "caption-side"
	emptyCells  = "empty-cells"
)

// =============================================================================
// Cursor & User Interface Constants
// =============================================================================

const (
	cursor        = "cursor"
	pointerEvents = "pointer-events"
	userSelect    = "user-select"
	resize        = "resize"
	appearance    = "appearance"
)

// =============================================================================
// Scroll Constants
// =============================================================================

const (
	scrollBehavior     = "scroll-behavior"
	scrollSnapType     = "scroll-snap-type"
	scrollSnapAlign    = "scroll-snap-align"
	scrollPadding      = "scroll-padding"
	scrollMargin       = "scroll-margin"
	overscrollBehavior = "overscroll-behavior"
)

// =============================================================================
// Content & Counters Constants
// =============================================================================

const (
	content          = "content"
	quotes           = "quotes"
	counterReset     = "counter-reset"
	counterIncrement = "counter-increment"
)

// =============================================================================
// Print Constants
// =============================================================================

const (
	pageBreakBefore = "page-break-before"
	pageBreakAfter  = "page-break-after"
	pageBreakInside = "page-break-inside"
)

// =============================================================================
// Aspect Ratio
// =============================================================================

const (
	aspectRatio = "aspect-ratio"
)

// =============================================================================
// Common CSS Values
// =============================================================================

const (
	Zero = "0"
	None = "none"

	// CSS Units
	Px      = "px"
	Pt      = "pt"
	Percent = "%"
	Em      = "em"
	Rem     = "rem"
	Cm      = "cm"
	Mm      = "mm"
	Vh      = "vh"
	Vw      = "vw"
	Vmin    = "vmin"
	Vmax    = "vmax"
	Ch      = "ch"
	Ex      = "ex"

	// Font weight
	Bold    = "bold"
	Normal  = "normal"
	Lighter = "lighter"
	Bolder  = "bolder"

	// Font style
	Italic  = "italic"
	Oblique = "oblique"

	// Text decoration
	Underline   = "underline"
	Overline    = "overline"
	LineThrough = "line-through"

	// Display values
	Block       = "block"
	Inline      = "inline"
	InlineBlock = "inline-block"
	Flex        = "flex"
	InlineFlex  = "inline-flex"
	Grid        = "grid"
	InlineGrid  = "inline-grid"

	// Position values
	Static   = "static"
	Relative = "relative"
	Absolute = "absolute"
	Fixed    = "fixed"
	Sticky   = "sticky"

	// Sizing values
	Auto       = "auto"
	FitContent = "fit-content"
	MinContent = "min-content"
	MaxContent = "max-content"
	Inherit    = "inherit"
	Initial    = "initial"
	Unset      = "unset"

	// Flexbox values
	FlexStart     = "flex-start"
	FlexEnd       = "flex-end"
	Center        = "center"
	SpaceBetween  = "space-between"
	SpaceAround   = "space-around"
	SpaceEvenly   = "space-evenly"
	Stretch       = "stretch"
	Baseline      = "baseline"
	Row           = "row"
	RowReverse    = "row-reverse"
	Column        = "column"
	ColumnReverse = "column-reverse"
	Wrap          = "wrap"
	NoWrap        = "nowrap"
	WrapReverse   = "wrap-reverse"

	// Alignment
	Start  = "start"
	End    = "end"
	Left   = "left"
	Right  = "right"
	Top    = "top"
	Bottom = "bottom"

	// Border styles
	Solid  = "solid"
	Dashed = "dashed"
	Dotted = "dotted"
	Double = "double"
	Groove = "groove"
	Ridge  = "ridge"
	Inset  = "inset"
	Outset = "outset"

	// Overflow values
	Visible = "visible"
	Hidden  = "hidden"
	Scroll  = "scroll"
	Clip    = "clip"

	// Backgrounds
	NoRepeat = "no-repeat"
	Repeat   = "repeat"
	RepeatX  = "repeat-x"
	RepeatY  = "repeat-y"
	Cover    = "cover"
	Contain  = "contain"

	// Border collapse
	Collapse = "collapse"
	Separate = "separate"

	// Vertical align
	Sub        = "sub"
	Super      = "super"
	TextTop    = "text-top"
	Middle     = "middle"
	TextBottom = "text-bottom"

	// Text transform
	Capitalize = "capitalize"
	Uppercase  = "uppercase"
	Lowercase  = "lowercase"

	// White space
	PreWrap     = "pre-wrap"
	PreLine     = "pre-line"
	Pre         = "pre"
	BreakSpaces = "break-spaces"

	// Page break
	Always = "always"
	Avoid  = "avoid"

	// Colors
	Transparent  = "transparent"
	CurrentColor = "currentColor"

	// Cursor values
	Pointer    = "pointer"
	Default    = "default"
	Text       = "text"
	Move       = "move"
	NotAllowed = "not-allowed"
	Grab       = "grab"
	Grabbing   = "grabbing"
	Crosshair  = "crosshair"
	ZoomIn     = "zoom-in"
	ZoomOut    = "zoom-out"

	// Object fit
	ObjectFill      = "fill"
	ObjectContain   = "contain"
	ObjectCover     = "cover"
	ObjectScaleDown = "scale-down"

	// Box sizing
	ContentBox = "content-box"
	BorderBox  = "border-box"
)

// =============================================================================
// Font Family Stacks
// =============================================================================

var (
	FontFamilyHelvetica  = Fonts([]string{"Helvetica Neue", "Helvetica", "Arial", "sans-serif"})
	FontFamilyCourierNew = Fonts([]string{"Courier New", "monospace"})
	FontFamilySystem     = Fonts([]string{"-apple-system", "BlinkMacSystemFont", "Segoe UI", "Roboto", "Helvetica Neue", "Arial", "sans-serif"})
	FontFamilyMonospace  = Fonts([]string{"SFMono-Regular", "Menlo", "Monaco", "Consolas", "Liberation Mono", "Courier New", "monospace"})
	FontFamilySerif      = Fonts([]string{"Georgia", "Cambria", "Times New Roman", "Times", "serif"})
)

// =============================================================================
// Types
// =============================================================================

type KeyValuePair struct {
	key     string
	value   string
	Include bool
}

type Styles []KeyValuePair

// =============================================================================
// Core Functions
// =============================================================================

func StyleList(kvs ...KeyValuePair) Styles {
	return kvs
}

func (kvp KeyValuePair) String() string {
	return fmt.Sprintf("%s:%s", kvp.key, kvp.value)
}

func PrintStyles(styles []KeyValuePair) string {
	components := []string{}

	for _, kvp := range styles {
		components = append(components, kvp.String())
	}

	return strings.Join(components, ";")
}

func constructKeyValuePair(key, value string, include ...bool) KeyValuePair {
	i := true
	if len(include) > 0 {
		i = include[0]
	}

	return KeyValuePair{key, value, i}
}

// =============================================================================
// Helper Functions
// =============================================================================

func WithUnits(value interface{}, unit string) string {
	return fmt.Sprintf("%v%s", value, unit)
}

func NoUnits(value interface{}) string {
	return fmt.Sprintf("%v", value)
}

func Fonts(fonts []string) string {
	return strings.Join(fonts, ", ")
}

func MultipleArgs(ss ...string) string {
	return strings.Join(ss, " ")
}

func MultipleValues(ss ...string) string {
	return strings.Join(ss, ",")
}

func RGBA(r, g, b, a float64) string {
	return fmt.Sprintf("rgba(%v,%v,%v,%v)", r, g, b, a)
}

func RGB(r, g, b float64) string {
	return fmt.Sprintf("rgb(%v,%v,%v)", r, g, b)
}

func HSL(h, s, l float64) string {
	return fmt.Sprintf("hsl(%v,%v%%,%v%%)", h, s, l)
}

func HSLA(h, s, l, a float64) string {
	return fmt.Sprintf("hsla(%v,%v%%,%v%%,%v)", h, s, l, a)
}

func URL(s string) string {
	return fmt.Sprintf(`url('%s')`, s)
}

func Var(name string) string {
	return fmt.Sprintf("var(--%s)", name)
}

func VarWithFallback(name, fallback string) string {
	return fmt.Sprintf("var(--%s, %s)", name, fallback)
}

func Calc(expression string) string {
	return fmt.Sprintf("calc(%s)", expression)
}

// =============================================================================
// Box Model Properties
// =============================================================================

func Width(s string, include ...bool) KeyValuePair {
	return constructKeyValuePair(width, s, include...)
}

func Height(s string, include ...bool) KeyValuePair {
	return constructKeyValuePair(height, s, include...)
}

func MinWidth(s string, include ...bool) KeyValuePair {
	return constructKeyValuePair(minWidth, s, include...)
}

func MinHeight(s string, include ...bool) KeyValuePair {
	return constructKeyValuePair(minHeight, s, include...)
}

func MaxWidth(s string, include ...bool) KeyValuePair {
	return constructKeyValuePair(maxWidth, s, include...)
}

func MaxHeight(s string, include ...bool) KeyValuePair {
	return constructKeyValuePair(maxHeight, s, include...)
}

func BoxSizing(s string, include ...bool) KeyValuePair {
	return constructKeyValuePair(boxSizing, s, include...)
}

func Margin(s string, include ...bool) KeyValuePair {
	return constructKeyValuePair(margin, s, include...)
}

func MarginVertHoriz(vertical, horizontal string, include ...bool) KeyValuePair {
	s := []string{vertical, horizontal}
	return constructKeyValuePair(margin, strings.Join(s, " "), include...)
}

func MarginTop(s string, include ...bool) KeyValuePair {
	return constructKeyValuePair(marginTop, s, include...)
}

func MarginRight(s string, include ...bool) KeyValuePair {
	return constructKeyValuePair(marginRight, s, include...)
}

func MarginBottom(s string, include ...bool) KeyValuePair {
	return constructKeyValuePair(marginBottom, s, include...)
}

func MarginLeft(s string, include ...bool) KeyValuePair {
	return constructKeyValuePair(marginLeft, s, include...)
}

func Padding(s string, include ...bool) KeyValuePair {
	return constructKeyValuePair(padding, s, include...)
}

func PaddingVertHoriz(vertical, horizontal string, include ...bool) KeyValuePair {
	s := []string{vertical, horizontal}
	return constructKeyValuePair(padding, strings.Join(s, " "), include...)
}

func PaddingTop(s string, include ...bool) KeyValuePair {
	return constructKeyValuePair(paddingTop, s, include...)
}

func PaddingRight(s string, include ...bool) KeyValuePair {
	return constructKeyValuePair(paddingRight, s, include...)
}

func PaddingBottom(s string, include ...bool) KeyValuePair {
	return constructKeyValuePair(paddingBottom, s, include...)
}

func PaddingLeft(s string, include ...bool) KeyValuePair {
	return constructKeyValuePair(paddingLeft, s, include...)
}

// =============================================================================
// Positioning Properties
// =============================================================================

func Position(s string, include ...bool) KeyValuePair {
	return constructKeyValuePair(position, s, include...)
}

func Top_(s string, include ...bool) KeyValuePair {
	return constructKeyValuePair(top, s, include...)
}

func Right_(s string, include ...bool) KeyValuePair {
	return constructKeyValuePair(right, s, include...)
}

func Bottom_(s string, include ...bool) KeyValuePair {
	return constructKeyValuePair(bottom, s, include...)
}

func Left_(s string, include ...bool) KeyValuePair {
	return constructKeyValuePair(left, s, include...)
}

func ZIndex(n int, include ...bool) KeyValuePair {
	return constructKeyValuePair(zIndex, fmt.Sprintf("%d", n), include...)
}

func Float(s string, include ...bool) KeyValuePair {
	return constructKeyValuePair(float_, s, include...)
}

func Clear(s string, include ...bool) KeyValuePair {
	return constructKeyValuePair(clear, s, include...)
}

// =============================================================================
// Display & Visibility Properties
// =============================================================================

func Display(s string, include ...bool) KeyValuePair {
	return constructKeyValuePair(display, s, include...)
}

func Visibility(s string, include ...bool) KeyValuePair {
	return constructKeyValuePair(visibility, s, include...)
}

func Opacity(s string, include ...bool) KeyValuePair {
	return constructKeyValuePair(opacity, s, include...)
}

func Overflow(s string, include ...bool) KeyValuePair {
	return constructKeyValuePair(overflow, s, include...)
}

func OverflowX(s string, include ...bool) KeyValuePair {
	return constructKeyValuePair(overflowX, s, include...)
}

func OverflowY(s string, include ...bool) KeyValuePair {
	return constructKeyValuePair(overflowY, s, include...)
}

// =============================================================================
// Flexbox Properties
// =============================================================================

func FlexDirection(s string, include ...bool) KeyValuePair {
	return constructKeyValuePair(flexDirection, s, include...)
}

func FlexWrap(s string, include ...bool) KeyValuePair {
	return constructKeyValuePair(flexWrap, s, include...)
}

func FlexFlow(s string, include ...bool) KeyValuePair {
	return constructKeyValuePair(flexFlow, s, include...)
}

func JustifyContent(s string, include ...bool) KeyValuePair {
	return constructKeyValuePair(justifyContent, s, include...)
}

func AlignItems(s string, include ...bool) KeyValuePair {
	return constructKeyValuePair(alignItems, s, include...)
}

func AlignContent(s string, include ...bool) KeyValuePair {
	return constructKeyValuePair(alignContent, s, include...)
}

func AlignSelf(s string, include ...bool) KeyValuePair {
	return constructKeyValuePair(alignSelf, s, include...)
}

func Flex_(s string, include ...bool) KeyValuePair {
	return constructKeyValuePair(flex, s, include...)
}

func FlexGrow(n float64, include ...bool) KeyValuePair {
	return constructKeyValuePair(flexGrow, fmt.Sprintf("%v", n), include...)
}

func FlexShrink(n float64, include ...bool) KeyValuePair {
	return constructKeyValuePair(flexShrink, fmt.Sprintf("%v", n), include...)
}

func FlexBasis(s string, include ...bool) KeyValuePair {
	return constructKeyValuePair(flexBasis, s, include...)
}

func Order(n int, include ...bool) KeyValuePair {
	return constructKeyValuePair(order, fmt.Sprintf("%d", n), include...)
}

func Gap(s string, include ...bool) KeyValuePair {
	return constructKeyValuePair(gap, s, include...)
}

func RowGap(s string, include ...bool) KeyValuePair {
	return constructKeyValuePair(rowGap, s, include...)
}

func ColumnGap(s string, include ...bool) KeyValuePair {
	return constructKeyValuePair(columnGap, s, include...)
}

// =============================================================================
// Grid Properties
// =============================================================================

func GridTemplateColumns(s string, include ...bool) KeyValuePair {
	return constructKeyValuePair(gridTemplateColumns, s, include...)
}

func GridTemplateRows(s string, include ...bool) KeyValuePair {
	return constructKeyValuePair(gridTemplateRows, s, include...)
}

func GridTemplateAreas(s string, include ...bool) KeyValuePair {
	return constructKeyValuePair(gridTemplateAreas, s, include...)
}

func GridTemplate(s string, include ...bool) KeyValuePair {
	return constructKeyValuePair(gridTemplate, s, include...)
}

func GridColumn(s string, include ...bool) KeyValuePair {
	return constructKeyValuePair(gridColumn, s, include...)
}

func GridColumnStart(n int, include ...bool) KeyValuePair {
	return constructKeyValuePair(gridColumnStart, fmt.Sprintf("%d", n), include...)
}

func GridColumnEnd(n int, include ...bool) KeyValuePair {
	return constructKeyValuePair(gridColumnEnd, fmt.Sprintf("%d", n), include...)
}

func GridRow(s string, include ...bool) KeyValuePair {
	return constructKeyValuePair(gridRow, s, include...)
}

func GridRowStart(n int, include ...bool) KeyValuePair {
	return constructKeyValuePair(gridRowStart, fmt.Sprintf("%d", n), include...)
}

func GridRowEnd(n int, include ...bool) KeyValuePair {
	return constructKeyValuePair(gridRowEnd, fmt.Sprintf("%d", n), include...)
}

func GridArea(s string, include ...bool) KeyValuePair {
	return constructKeyValuePair(gridArea, s, include...)
}

func GridAutoFlow(s string, include ...bool) KeyValuePair {
	return constructKeyValuePair(gridAutoFlow, s, include...)
}

func GridAutoColumns(s string, include ...bool) KeyValuePair {
	return constructKeyValuePair(gridAutoColumns, s, include...)
}

func GridAutoRows(s string, include ...bool) KeyValuePair {
	return constructKeyValuePair(gridAutoRows, s, include...)
}

func JustifyItems(s string, include ...bool) KeyValuePair {
	return constructKeyValuePair(justifyItems, s, include...)
}

func JustifySelf(s string, include ...bool) KeyValuePair {
	return constructKeyValuePair(justifySelf, s, include...)
}

func PlaceContent(s string, include ...bool) KeyValuePair {
	return constructKeyValuePair(placeContent, s, include...)
}

func PlaceItems(s string, include ...bool) KeyValuePair {
	return constructKeyValuePair(placeItems, s, include...)
}

func PlaceSelf(s string, include ...bool) KeyValuePair {
	return constructKeyValuePair(placeSelf, s, include...)
}

// =============================================================================
// Typography Properties
// =============================================================================

func FontFamily(s string, include ...bool) KeyValuePair {
	return constructKeyValuePair(fontFamily, s, include...)
}

func FontSize(s string, include ...bool) KeyValuePair {
	return constructKeyValuePair(fontSize, s, include...)
}

func FontWeight(s string, include ...bool) KeyValuePair {
	return constructKeyValuePair(fontWeight, s, include...)
}

func FontStyle(s string, include ...bool) KeyValuePair {
	return constructKeyValuePair(fontStyle, s, include...)
}

func FontVariant(s string, include ...bool) KeyValuePair {
	return constructKeyValuePair(fontVariant, s, include...)
}

func LineHeight(s string, include ...bool) KeyValuePair {
	return constructKeyValuePair(lineHeight, s, include...)
}

func LetterSpacing(s string, include ...bool) KeyValuePair {
	return constructKeyValuePair(letterSpacing, s, include...)
}

func WordSpacing(s string, include ...bool) KeyValuePair {
	return constructKeyValuePair(wordSpacing, s, include...)
}

func TextAlign(s string, include ...bool) KeyValuePair {
	return constructKeyValuePair(textAlign, s, include...)
}

func TextDecoration(s string, include ...bool) KeyValuePair {
	return constructKeyValuePair(textDecoration, s, include...)
}

func TextTransform(s string, include ...bool) KeyValuePair {
	return constructKeyValuePair(textTransform, s, include...)
}

func TextIndent(s string, include ...bool) KeyValuePair {
	return constructKeyValuePair(textIndent, s, include...)
}

func TextShadow(s string, include ...bool) KeyValuePair {
	return constructKeyValuePair(textShadow, s, include...)
}

func TextOverflow(s string, include ...bool) KeyValuePair {
	return constructKeyValuePair(textOverflow, s, include...)
}

func WhiteSpace(s string, include ...bool) KeyValuePair {
	return constructKeyValuePair(whiteSpace, s, include...)
}

func WordBreak(s string, include ...bool) KeyValuePair {
	return constructKeyValuePair(wordBreak, s, include...)
}

func WordWrap(s string, include ...bool) KeyValuePair {
	return constructKeyValuePair(wordWrap, s, include...)
}

func OverflowWrap(s string, include ...bool) KeyValuePair {
	return constructKeyValuePair(overflowWrap, s, include...)
}

func VerticalAlign(s string, include ...bool) KeyValuePair {
	return constructKeyValuePair(verticalAlign, s, include...)
}

func Direction(s string, include ...bool) KeyValuePair {
	return constructKeyValuePair(direction, s, include...)
}

func UnicodeBidi(s string, include ...bool) KeyValuePair {
	return constructKeyValuePair(unicodeBidi, s, include...)
}

// =============================================================================
// Color & Background Properties
// =============================================================================

func Color(s string, include ...bool) KeyValuePair {
	return constructKeyValuePair(color, s, include...)
}

func BackgroundColor(s string, include ...bool) KeyValuePair {
	return constructKeyValuePair(backgroundColor, s, include...)
}

func BackgroundImage(s string, include ...bool) KeyValuePair {
	return constructKeyValuePair(backgroundImage, s, include...)
}

func BackgroundRepeat(s string, include ...bool) KeyValuePair {
	return constructKeyValuePair(backgroundRepeat, s, include...)
}

func BackgroundPosition(s string, include ...bool) KeyValuePair {
	return constructKeyValuePair(backgroundPosition, s, include...)
}

func BackgroundSize(s string, include ...bool) KeyValuePair {
	return constructKeyValuePair(backgroundSize, s, include...)
}

func BackgroundAttachment(s string, include ...bool) KeyValuePair {
	return constructKeyValuePair(backgroundAttachment, s, include...)
}

func BackgroundClip(s string, include ...bool) KeyValuePair {
	return constructKeyValuePair(backgroundClip, s, include...)
}

func BackgroundOrigin(s string, include ...bool) KeyValuePair {
	return constructKeyValuePair(backgroundOrigin, s, include...)
}

func Background(s string, include ...bool) KeyValuePair {
	return constructKeyValuePair(background, s, include...)
}

// =============================================================================
// Border Properties
// =============================================================================

func Border(s string, include ...bool) KeyValuePair {
	return constructKeyValuePair(border, s, include...)
}

func BorderWidth(s string, include ...bool) KeyValuePair {
	return constructKeyValuePair(borderWidth, s, include...)
}

func BorderStyle(s string, include ...bool) KeyValuePair {
	return constructKeyValuePair(borderStyle, s, include...)
}

func BorderColor(s string, include ...bool) KeyValuePair {
	return constructKeyValuePair(borderColor, s, include...)
}

func BorderTop(s string, include ...bool) KeyValuePair {
	return constructKeyValuePair(borderTop, s, include...)
}

func BorderRight(s string, include ...bool) KeyValuePair {
	return constructKeyValuePair(borderRight, s, include...)
}

func BorderBottom(s string, include ...bool) KeyValuePair {
	return constructKeyValuePair(borderBottom, s, include...)
}

func BorderLeft(s string, include ...bool) KeyValuePair {
	return constructKeyValuePair(borderLeft, s, include...)
}

func BorderRadius(s string, include ...bool) KeyValuePair {
	return constructKeyValuePair(borderRadius, s, include...)
}

func BorderTopLeftRadius(s string, include ...bool) KeyValuePair {
	return constructKeyValuePair(borderTopLeftRadius, s, include...)
}

func BorderTopRightRadius(s string, include ...bool) KeyValuePair {
	return constructKeyValuePair(borderTopRightRadius, s, include...)
}

func BorderBottomLeftRadius(s string, include ...bool) KeyValuePair {
	return constructKeyValuePair(borderBottomLeftRadius, s, include...)
}

func BorderBottomRightRadius(s string, include ...bool) KeyValuePair {
	return constructKeyValuePair(borderBottomRightRadius, s, include...)
}

func BorderCollapse(s string, include ...bool) KeyValuePair {
	return constructKeyValuePair(borderCollapse, s, include...)
}

func BorderSpacing(s string, include ...bool) KeyValuePair {
	return constructKeyValuePair(borderSpacing, s, include...)
}

// =============================================================================
// Outline Properties
// =============================================================================

func Outline(s string, include ...bool) KeyValuePair {
	return constructKeyValuePair(outline, s, include...)
}

func OutlineWidth(s string, include ...bool) KeyValuePair {
	return constructKeyValuePair(outlineWidth, s, include...)
}

func OutlineStyle(s string, include ...bool) KeyValuePair {
	return constructKeyValuePair(outlineStyle, s, include...)
}

func OutlineColor(s string, include ...bool) KeyValuePair {
	return constructKeyValuePair(outlineColor, s, include...)
}

func OutlineOffset(s string, include ...bool) KeyValuePair {
	return constructKeyValuePair(outlineOffset, s, include...)
}

// =============================================================================
// Box Shadow & Effects Properties
// =============================================================================

func BoxShadow(s string, include ...bool) KeyValuePair {
	return constructKeyValuePair(boxShadow, s, include...)
}

func Filter(s string, include ...bool) KeyValuePair {
	return constructKeyValuePair(filter_, s, include...)
}

func BackdropFilter(s string, include ...bool) KeyValuePair {
	return constructKeyValuePair(backdropFilter, s, include...)
}

func ClipPath(s string, include ...bool) KeyValuePair {
	return constructKeyValuePair(clipPath, s, include...)
}

func ObjectFit(s string, include ...bool) KeyValuePair {
	return constructKeyValuePair(objectFit, s, include...)
}

func ObjectPosition(s string, include ...bool) KeyValuePair {
	return constructKeyValuePair(objectPosition, s, include...)
}

// =============================================================================
// Transform Properties
// =============================================================================

func Transform(s string, include ...bool) KeyValuePair {
	return constructKeyValuePair(transform, s, include...)
}

func TransformOrigin(s string, include ...bool) KeyValuePair {
	return constructKeyValuePair(transformOrigin, s, include...)
}

func TransformStyle(s string, include ...bool) KeyValuePair {
	return constructKeyValuePair(transformStyle, s, include...)
}

func Perspective(s string, include ...bool) KeyValuePair {
	return constructKeyValuePair(perspective, s, include...)
}

func PerspectiveOrigin(s string, include ...bool) KeyValuePair {
	return constructKeyValuePair(perspectiveOrigin, s, include...)
}

func BackfaceVisibility(s string, include ...bool) KeyValuePair {
	return constructKeyValuePair(backfaceVisibility, s, include...)
}

// =============================================================================
// Transition & Animation Properties
// =============================================================================

func Transition(s string, include ...bool) KeyValuePair {
	return constructKeyValuePair(transition, s, include...)
}

func TransitionProperty(s string, include ...bool) KeyValuePair {
	return constructKeyValuePair(transitionProperty, s, include...)
}

func TransitionDuration(s string, include ...bool) KeyValuePair {
	return constructKeyValuePair(transitionDuration, s, include...)
}

func TransitionTimingFunction(s string, include ...bool) KeyValuePair {
	return constructKeyValuePair(transitionTimingFunction, s, include...)
}

func TransitionDelay(s string, include ...bool) KeyValuePair {
	return constructKeyValuePair(transitionDelay, s, include...)
}

func Animation(s string, include ...bool) KeyValuePair {
	return constructKeyValuePair(animation, s, include...)
}

func AnimationName(s string, include ...bool) KeyValuePair {
	return constructKeyValuePair(animationName, s, include...)
}

func AnimationDuration(s string, include ...bool) KeyValuePair {
	return constructKeyValuePair(animationDuration, s, include...)
}

func AnimationTimingFunction(s string, include ...bool) KeyValuePair {
	return constructKeyValuePair(animationTimingFunction, s, include...)
}

func AnimationDelay(s string, include ...bool) KeyValuePair {
	return constructKeyValuePair(animationDelay, s, include...)
}

func AnimationIterationCount(s string, include ...bool) KeyValuePair {
	return constructKeyValuePair(animationIterationCount, s, include...)
}

func AnimationDirection(s string, include ...bool) KeyValuePair {
	return constructKeyValuePair(animationDirection, s, include...)
}

func AnimationFillMode(s string, include ...bool) KeyValuePair {
	return constructKeyValuePair(animationFillMode, s, include...)
}

func AnimationPlayState(s string, include ...bool) KeyValuePair {
	return constructKeyValuePair(animationPlayState, s, include...)
}

// =============================================================================
// List Properties
// =============================================================================

func ListStyle(s string, include ...bool) KeyValuePair {
	return constructKeyValuePair(listStyle, s, include...)
}

func ListStyleType(s string, include ...bool) KeyValuePair {
	return constructKeyValuePair(listStyleType, s, include...)
}

func ListStylePosition(s string, include ...bool) KeyValuePair {
	return constructKeyValuePair(listStylePosition, s, include...)
}

func ListStyleImage(s string, include ...bool) KeyValuePair {
	return constructKeyValuePair(listStyleImage, s, include...)
}

// =============================================================================
// Table Properties
// =============================================================================

func TableLayout(s string, include ...bool) KeyValuePair {
	return constructKeyValuePair(tableLayout, s, include...)
}

func CaptionSide(s string, include ...bool) KeyValuePair {
	return constructKeyValuePair(captionSide, s, include...)
}

func EmptyCells(s string, include ...bool) KeyValuePair {
	return constructKeyValuePair(emptyCells, s, include...)
}

// =============================================================================
// Cursor & User Interface Properties
// =============================================================================

func Cursor(s string, include ...bool) KeyValuePair {
	return constructKeyValuePair(cursor, s, include...)
}

func PointerEvents(s string, include ...bool) KeyValuePair {
	return constructKeyValuePair(pointerEvents, s, include...)
}

func UserSelect(s string, include ...bool) KeyValuePair {
	return constructKeyValuePair(userSelect, s, include...)
}

func Resize(s string, include ...bool) KeyValuePair {
	return constructKeyValuePair(resize, s, include...)
}

func Appearance(s string, include ...bool) KeyValuePair {
	return constructKeyValuePair(appearance, s, include...)
}

// =============================================================================
// Scroll Properties
// =============================================================================

func ScrollBehavior(s string, include ...bool) KeyValuePair {
	return constructKeyValuePair(scrollBehavior, s, include...)
}

func ScrollSnapType(s string, include ...bool) KeyValuePair {
	return constructKeyValuePair(scrollSnapType, s, include...)
}

func ScrollSnapAlign(s string, include ...bool) KeyValuePair {
	return constructKeyValuePair(scrollSnapAlign, s, include...)
}

func ScrollPadding(s string, include ...bool) KeyValuePair {
	return constructKeyValuePair(scrollPadding, s, include...)
}

func ScrollMargin(s string, include ...bool) KeyValuePair {
	return constructKeyValuePair(scrollMargin, s, include...)
}

func OverscrollBehavior(s string, include ...bool) KeyValuePair {
	return constructKeyValuePair(overscrollBehavior, s, include...)
}

// =============================================================================
// Content & Counter Properties
// =============================================================================

func Content(s string, include ...bool) KeyValuePair {
	return constructKeyValuePair(content, s, include...)
}

func Quotes(s string, include ...bool) KeyValuePair {
	return constructKeyValuePair(quotes, s, include...)
}

func CounterReset(s string, include ...bool) KeyValuePair {
	return constructKeyValuePair(counterReset, s, include...)
}

func CounterIncrement(s string, include ...bool) KeyValuePair {
	return constructKeyValuePair(counterIncrement, s, include...)
}

// =============================================================================
// Print Properties
// =============================================================================

func PageBreakBefore(s string, include ...bool) KeyValuePair {
	return constructKeyValuePair(pageBreakBefore, s, include...)
}

func PageBreakAfter(s string, include ...bool) KeyValuePair {
	return constructKeyValuePair(pageBreakAfter, s, include...)
}

func PageBreakInside(s string, include ...bool) KeyValuePair {
	return constructKeyValuePair(pageBreakInside, s, include...)
}

// =============================================================================
// Aspect Ratio
// =============================================================================

func AspectRatio(s string, include ...bool) KeyValuePair {
	return constructKeyValuePair(aspectRatio, s, include...)
}
