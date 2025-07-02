package visitor

import "fmt"

// —— Elements ——

type Visitor interface {
	VisitText(n *TextNode)
	VisitLink(n *LinkNode)
}

type Node interface {
	Accept(v Visitor)
}

// TextNode holds plain text
type TextNode struct{ Text string }

func (t *TextNode) Accept(v Visitor) { v.VisitText(t) }

// LinkNode holds a URL and link text
type LinkNode struct {
	URL, Text string
}

func (l *LinkNode) Accept(v Visitor) { v.VisitLink(l) }

// A small document tree
func buildDoc() []Node {
	return []Node{
		&TextNode{Text: "Check out "},
		&LinkNode{URL: "https://golang.org", Text: "Go"},
		&TextNode{Text: " for more info."},
	}
}

// —— Concrete Visitors ——

type HTMLVisitor struct {
	Output string
}

func (h *HTMLVisitor) VisitText(n *TextNode) {
	h.Output += n.Text
}
func (h *HTMLVisitor) VisitLink(n *LinkNode) {
	h.Output += fmt.Sprintf(`<a href="%s">%s</a>`, n.URL, n.Text)
}

type PlainTextVisitor struct {
	Output string
}

func (p *PlainTextVisitor) VisitText(n *TextNode) {
	p.Output += n.Text
}
func (p *PlainTextVisitor) VisitLink(n *LinkNode) {
	p.Output += fmt.Sprintf("%s (%s)", n.Text, n.URL)
}

// —— Client ——

func RunDocumentExportVisitor() {
	doc := buildDoc()

	htmlV := &HTMLVisitor{}
	txtV := &PlainTextVisitor{}

	for _, node := range doc {
		node.Accept(htmlV)
		node.Accept(txtV)
	}

	fmt.Println("HTML Render:")
	fmt.Println(htmlV.Output)
	fmt.Println("\nPlain Text:")
	fmt.Println(txtV.Output)
}
