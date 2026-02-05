package attributes

import (
	"fmt"
	"strconv"
	"strings"
)

// Attribue

type Attribute struct {
	Name, Val              string
	IsBool                 bool
	RenderWithSingleQuotes bool
}

func (attr Attribute) Output() string {
	// An attr without a name is not rendered
	if attr.Name == "" {
		return ""
	}

	if attr.IsBool {
		return attr.Name
	}

	quote := `"`
	if attr.RenderWithSingleQuotes {
		quote = `'`
	}

	return fmt.Sprintf(`%s=%s%s%s`, attr.Name, quote, attr.Val, quote)
}

func (attr Attribute) String() string {
	return attr.Output()
}

// Attributes

type Attributes []Attribute

func (attrs Attributes) Output() string {
	attrs = attrs.MergeDuplicates()

	var renderedAttrs []string

	for _, attr := range attrs {
		r := attr.Output()
		// Don't include blank attrs, to avoid extraneous spaces
		if r != "" {
			renderedAttrs = append(renderedAttrs, r)
		}
	}

	if len(renderedAttrs) == 0 {
		return ""
	}

	return " " + strings.Join(renderedAttrs, " ")
}

func (attrs Attributes) String() string {
	return attrs.String()
}

// Helpers

func (a Attribute) RenderIf(doRender bool) Attribute {
	if doRender {
		return a
	}

	return Attribute{}
}

func Attrs(attrs ...Attribute) Attributes {
	return attrs
}

// Merge helps deal with repition of certain attributes, merging them into one.
// At the moment, we deal with 'class' and 'style'
func (attrs Attributes) MergeDuplicates() Attributes {
	var classes []string
	var styles []string
	var otherAttrs Attributes

	for _, attr := range attrs {
		switch attr.Name {
		case class:
			classes = append(classes, attr.Val)
		case style:
			styles = append(styles, attr.Val)
		default:
			otherAttrs = append(otherAttrs, attr)
		}
	}

	if len(classes) > 0 {
		otherAttrs = append(otherAttrs, Classes(classes))
	}

	if len(styles) > 0 {
		otherAttrs = append(otherAttrs, regularAttribute(style, strings.Join(styles, ";")))
	}

	return otherAttrs
}

func regularAttribute(k, v string) Attribute {
	return Attribute{
		Name:   k,
		Val:    v,
		IsBool: false,
	}
}

func regularAttributeWithSingleQuotes(k, v string) Attribute {
	return Attribute{
		Name:                   k,
		Val:                    v,
		IsBool:                 false,
		RenderWithSingleQuotes: true,
	}
}

func booleanAttribute(k string, b bool) Attribute {
	return Attribute{
		Name:   k,
		IsBool: true,
	}.RenderIf(b)
}

func intAttribute(k string, v int) Attribute {
	return Attribute{
		Name:   k,
		Val:    strconv.Itoa(v),
		IsBool: false,
	}
}

func floatAttribute(k string, v float64) Attribute {
	return Attribute{
		Name:   k,
		Val:    fmt.Sprintf("%v", v),
		IsBool: false,
	}
}
