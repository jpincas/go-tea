package test

import (
	"os"
	"testing"

	a "github.com/jpincas/go-tea/attributes"
	"github.com/jpincas/go-tea/css"
	h "github.com/jpincas/go-tea/html"
)

// The only test we have is to generate a sample HTML page that uses as much of the library as possible
// Obviously, the test serves to make sure the code compiles and runs, but you can also:
// - visually inspect the HTML to test anything you are working on
// - run it against an HTML validator. e.g., https://github.com/validator/validator/
// - run test, then: /usr/local/bin/vnu-runtime-image/bin/vnu index.html
// - inspect the generated HTML in a browser

// This page is also a good example of how to write HTML using the html package that can be presented to an AI
// It also serves as a style guide for how to format the code nicely following these guidelines:
// - The a.Attrs() function should be open on the same line as the HTML tag
// - The a.Attrs() function should be closed on the same line as the last attribute
// - Each attribute should be on its own line
// - Each HTML tag should be on its own line
// - Each HTML tag should be indented by one tab
// - Each HTML tag should be closed on its own line

func TestPageGen(t *testing.T) {
	// Header with navigation
	header :=
		h.Header(a.Attrs(
			a.Class("site-header"),
			a.Role("banner")),
			h.Nav(a.Attrs(
				a.Class("main-nav"),
				a.AriaLabel("Main navigation")),
				h.Ul(a.Attrs(
					a.Class("nav-list")),
					h.Li(a.Attrs(),
						h.A(a.Attrs(
							a.Href("#home"),
							a.AriaCurrent("page")),
							h.Text("Home"))),
					h.Li(a.Attrs(),
						h.A(a.Attrs(
							a.Href("#about")),
							h.Text("About"))),
					h.Li(a.Attrs(),
						h.A(a.Attrs(
							a.Href("#contact")),
							h.Text("Contact"))))))

	// Main content with various semantic elements
	mainContent :=
		h.Main(a.Attrs(
			a.Id("main-content"),
			a.Role("main")),
			// Article with semantic elements
			h.Article(a.Attrs(
				a.Class("article")),
				h.Hgroup(a.Attrs(),
					h.H1(a.Attrs(),
						h.Text("Article Title")),
					h.P(a.Attrs(
						a.Class("subtitle")),
						h.Text("A subtitle for the article"))),
				h.Time(a.Attrs(
					a.Datetime("2024-01-15")),
					h.Text("January 15, 2024")),
				h.P(a.Attrs(),
					h.Text("This is an article with "),
					h.Strong(a.Attrs(),
						h.Text("strong text")),
					h.Text(", "),
					h.Em(a.Attrs(),
						h.Text("emphasized text")),
					h.Text(", and "),
					h.Mark(a.Attrs(),
						h.Text("marked text")),
					h.Text(".")),
				h.Figure(a.Attrs(),
					h.Img(a.Attrs(
						a.Src("https://example.com/image.jpg"),
						a.Alt("Example image"),
						a.Loading("lazy"))),
					h.Figcaption(a.Attrs(),
						h.Text("Figure 1: An example image")))),

			// Code example
			h.Section(a.Attrs(
				a.Class("code-section")),
				h.H2(a.Attrs(),
					h.Text("Code Example")),
				h.Pre(a.Attrs(),
					h.Code(a.Attrs(
						a.Class("language-go")),
						h.Text("func main() {\n    fmt.Println(\"Hello, World!\")\n}")))),

			// Definition list
			h.Section(a.Attrs(),
				h.H2(a.Attrs(),
					h.Text("Glossary")),
				h.Dl(a.Attrs(),
					h.Dt(a.Attrs(),
						h.Text("HTML")),
					h.Dd(a.Attrs(),
						h.Text("HyperText Markup Language")),
					h.Dt(a.Attrs(),
						h.Text("CSS")),
					h.Dd(a.Attrs(),
						h.Text("Cascading Style Sheets")))),

			// Interactive elements
			h.Section(a.Attrs(
				a.Class("interactive")),
				h.H2(a.Attrs(),
					h.Text("Interactive Elements")),
				h.Details(a.Attrs(),
					h.Summary(a.Attrs(),
						h.Text("Click to expand")),
					h.P(a.Attrs(),
						h.Text("This is hidden content that appears when expanded."))),
				h.Dialog(a.Attrs(
					a.Id("demo-dialog")),
					h.P(a.Attrs(),
						h.Text("This is a dialog element.")))),

			// Form with fieldset
			h.Section(a.Attrs(),
				h.H2(a.Attrs(),
					h.Text("Contact Form")),
				h.Form(a.Attrs(
					a.Method("post"),
					a.Action("/submit")),
					h.Fieldset(a.Attrs(),
						h.Legend(a.Attrs(),
							h.Text("Personal Information")),
						h.Label(a.Attrs(
							a.For("name")),
							h.Text("Name:")),
						h.Input(a.Attrs(
							a.Type("text"),
							a.Id("name"),
							a.Name("name"),
							a.Required(true),
							a.Autocomplete("name"),
							a.Placeholder("Enter your name"))),
						h.Br(a.Attrs()),
						h.Label(a.Attrs(
							a.For("email")),
							h.Text("Email:")),
						h.Input(a.Attrs(
							a.Type("email"),
							a.Id("email"),
							a.Name("email"),
							a.Required(true),
							a.Autocomplete("email"),
							a.Inputmode("email"))),
						h.Br(a.Attrs()),
						h.Label(a.Attrs(
							a.For("browser")),
							h.Text("Browser:")),
						h.Input(a.Attrs(
							a.Type("text"),
							a.Id("browser"),
							a.Name("browser"),
							a.List("browsers"))),
						h.Datalist(a.Attrs(
							a.Id("browsers")),
							h.Option(a.Attrs(
								a.Value("Chrome"))),
							h.Option(a.Attrs(
								a.Value("Firefox"))),
							h.Option(a.Attrs(
								a.Value("Safari"))))),
					h.Fieldset(a.Attrs(),
						h.Legend(a.Attrs(),
							h.Text("Preferences")),
						h.Label(a.Attrs(
							a.For("country")),
							h.Text("Country:")),
						h.Select(a.Attrs(
							a.Id("country"),
							a.Name("country")),
							h.Optgroup(a.Attrs(
								a.Label("Europe")),
								h.Option(a.Attrs(
									a.Value("uk")),
									h.Text("United Kingdom")),
								h.Option(a.Attrs(
									a.Value("de")),
									h.Text("Germany"))),
							h.Optgroup(a.Attrs(
								a.Label("Americas")),
								h.Option(a.Attrs(
									a.Value("us")),
									h.Text("United States")),
								h.Option(a.Attrs(
									a.Value("ca")),
									h.Text("Canada"))))),
					h.Button(a.Attrs(
						a.Type("submit")),
						h.Text("Submit")))),

			// Media elements
			h.Section(a.Attrs(),
				h.H2(a.Attrs(),
					h.Text("Media")),
				h.Picture(a.Attrs(),
					h.Source(a.Attrs(
						a.Srcset("image-wide.jpg"),
						a.Media("(min-width: 800px)"))),
					h.Img(a.Attrs(
						a.Src("image.jpg"),
						a.Alt("Responsive image")))),
				h.Video(a.Attrs(
					a.Controls(true),
					a.Poster("poster.jpg"),
					a.Width("640"),
					a.Height("360")),
					h.Source(a.Attrs(
						a.Src("video.mp4"),
						a.Type("video/mp4"))),
					h.Track(a.Attrs(
						a.Src("subtitles.vtt"),
						a.Kind("subtitles"),
						a.Srclang("en"),
						a.Label("English"))),
					h.Text("Your browser does not support video.")),
				h.Audio(a.Attrs(
					a.Controls(true)),
					h.Source(a.Attrs(
						a.Src("audio.mp3"),
						a.Type("audio/mpeg"))),
					h.Text("Your browser does not support audio.")),
				h.Canvas(a.Attrs(
					a.Id("myCanvas"),
					a.Width("300"),
					a.Height("150")),
					h.Text("Canvas not supported"))),

			// Embedded content
			h.Section(a.Attrs(),
				h.H2(a.Attrs(),
					h.Text("Embedded Content")),
				h.Iframe(a.Attrs(
					a.Src("https://example.com"),
					a.Width("600"),
					a.Height("400"),
					a.Loading("lazy"),
					a.Sandbox("allow-scripts"),
					a.Title("Embedded frame")))),

			// Progress and meter
			h.Section(a.Attrs(),
				h.H2(a.Attrs(),
					h.Text("Progress and Meters")),
				h.Label(a.Attrs(
					a.For("task-progress")),
					h.Text("Task progress:")),
				h.Progress(a.Attrs(
					a.Id("task-progress"),
					a.Value("70"),
					a.Max("100")),
					h.Text("70%")),
				h.Br(a.Attrs()),
				h.Label(a.Attrs(
					a.For("disk-usage")),
					h.Text("Disk usage:")),
				h.Meter(a.Attrs(
					a.Id("disk-usage"),
					a.Value("0.7"),
					a.Min("0"),
					a.Max("1"),
					a.Low(0.3),
					a.High(0.8),
					a.Optimum(0.5)),
					h.Text("70%"))),

			// Table with full structure
			h.Section(a.Attrs(),
				h.H2(a.Attrs(),
					h.Text("Data Table")),
				h.Table(a.Attrs(
					a.Class("data-table")),
					h.Caption(a.Attrs(),
						h.Text("Monthly Sales Data")),
					h.Colgroup(a.Attrs(),
						h.Col(a.Attrs(
							a.Span(1),
							a.Class("month-col"))),
						h.Col(a.Attrs(
							a.Span(2),
							a.Class("data-col")))),
					h.THead(a.Attrs(),
						h.Tr(a.Attrs(),
							h.Th(a.Attrs(
								a.Scope("col")),
								h.Text("Month")),
							h.Th(a.Attrs(
								a.Scope("col")),
								h.Text("Sales")),
							h.Th(a.Attrs(
								a.Scope("col")),
								h.Text("Revenue")))),
					h.TBody(a.Attrs(),
						h.Tr(a.Attrs(),
							h.Th(a.Attrs(
								a.Scope("row")),
								h.Text("January")),
							h.Td(a.Attrs(),
								h.Text("100")),
							h.Td(a.Attrs(),
								h.Text("$10,000"))),
						h.Tr(a.Attrs(),
							h.Th(a.Attrs(
								a.Scope("row")),
								h.Text("February")),
							h.Td(a.Attrs(),
								h.Text("120")),
							h.Td(a.Attrs(),
								h.Text("$12,000")))),
					h.TFoot(a.Attrs(),
						h.Tr(a.Attrs(),
							h.Th(a.Attrs(
								a.Scope("row")),
								h.Text("Total")),
							h.Td(a.Attrs(),
								h.Text("220")),
							h.Td(a.Attrs(),
								h.Text("$22,000")))))),

			// SVG example
			h.Section(a.Attrs(),
				h.H2(a.Attrs(),
					h.Text("SVG Example")),
				h.SVG(a.Attrs(
					a.ViewBox("0 0 100 100"),
					a.Width("100"),
					a.Height("100"),
					a.AriaLabel("Simple shapes")),
					h.Defs(a.Attrs(),
						h.LinearGradient(a.Attrs(
							a.Id("gradient1"),
							a.X1(0),
							a.Y1(0),
							a.X2(1),
							a.Y2(1)),
							h.Stop(a.Attrs(
								a.Offset("0%"),
								a.StopColor("red"))),
							h.Stop(a.Attrs(
								a.Offset("100%"),
								a.StopColor("blue"))))),
					h.G(a.Attrs(
						a.Fill("url(#gradient1)")),
						h.Circle(a.Attrs(
							a.CX(50),
							a.CY(50),
							a.R(40))),
						h.Rect(a.Attrs(
							a.X(10),
							a.Y(10),
							a.Width("20"),
							a.Height("20"),
							a.RX(2),
							a.RY(2)))),
					h.Line(a.Attrs(
						a.X1(0),
						a.Y1(0),
						a.X2(100),
						a.Y2(100),
						a.Stroke("black"),
						a.StrokeWidth(2))),
					h.Path(a.Attrs(
						a.D("M10 80 Q 52.5 10, 95 80 T 180 80"),
						a.Fill("none"),
						a.Stroke("blue"),
						a.StrokeWidth(2))))))

	// Aside
	aside :=
		h.Aside(a.Attrs(
			a.Class("sidebar")),
			h.H3(a.Attrs(),
				h.Text("Related Links")),
			h.Ul(a.Attrs(),
				h.Li(a.Attrs(),
					h.A(a.Attrs(
						a.Href("#")),
						h.Text("Link 1"))),
				h.Li(a.Attrs(),
					h.A(a.Attrs(
						a.Href("#")),
						h.Text("Link 2")))))

	// Footer
	footer :=
		h.Footer(a.Attrs(
			a.Class("site-footer"),
			a.Role("contentinfo")),
			h.Address(a.Attrs(),
				h.Text("Contact us at "),
				h.A(a.Attrs(
					a.Href("mailto:info@example.com")),
					h.Text("info@example.com"))),
			h.Small(a.Attrs(),
				h.Text("© 2024 Example Corp. All rights reserved.")))

	// Head section
	head :=
		h.Head(a.Attrs(),
			h.Title(a.Attrs(),
				h.Text("Go-Tea HTML Test Document")),
			h.Meta(a.Attrs(
				a.Charset("UTF-8"))),
			h.Meta(a.Attrs(
				a.Name("viewport"),
				a.Content("width=device-width, initial-scale=1.0"))),
			h.Meta(a.Attrs(
				a.Name("description"),
				a.Content("A test page demonstrating go-tea html library features"))),
			h.Meta(a.Attrs(
				a.Property("og:title"),
				a.Content("Go-Tea HTML Test"))),
			h.Link(a.Attrs(
				a.Rel("stylesheet"),
				a.Href("styles.css"))),
			h.Style(a.Attrs(),
				h.Text(`
					body { font-family: system-ui, sans-serif; }
					.site-header { background: #f0f0f0; padding: 1rem; }
					.main-nav ul { list-style: none; display: flex; gap: 1rem; }
				`)))

	// Body with CSS styling
	body :=
		h.Body(a.Attrs(
			a.Class("page-body"),
			a.Style(
				css.Display(css.Flex),
				css.FlexDirection(css.Column),
				css.MinHeight("100vh"),
				css.FontFamily(css.FontFamilySystem))),
			header,
			h.Div(a.Attrs(
				a.Class("content-wrapper"),
				a.Style(
					css.Display(css.Flex),
					css.FlexGrow(1),
					css.Gap("2rem"),
					css.Padding("1rem"))),
				mainContent,
				aside),
			footer)

	// Full document
	doc :=
		h.Html(a.Attrs(
			a.Lang("en")),
			head,
			body)

	f, err := os.Create("index.html")
	if err != nil {
		t.Fatal("Could not create index.html")
	}
	defer f.Close()

	if err := doc.WriteDoc(f); err != nil {
		t.Fatal("Could not write to index.html: ", err)
	}
}
