package test

import (
	"testing"

	a "github.com/jpincas/go-tea/attributes"
	h "github.com/jpincas/go-tea/html"
)

func TestRenderDoc(t *testing.T) {
	el := h.Div(
		a.Attrs(
			a.Class("test"),
		),
		h.Text("TEST"),
	)

	output := el.DocString()
	expected := `<!DOCTYPE html>
<div class="test">
  TEST
</div>
`

	if output != expected {
		t.Errorf("Output was not as expected.  Output: \n%s\n Expected: \n%s\n", output, expected)
	}
}
