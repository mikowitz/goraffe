package goraffe

// NodeOption is a functional option for configuring node attributes.
// Options can be passed to NewNode or used to build reusable attribute templates.
type NodeOption interface {
	applyNode(*NodeAttributes)
}

type nodeOptionFunc func(*NodeAttributes)

func (f nodeOptionFunc) applyNode(a *NodeAttributes) {
	f(a)
}

func newNodeOption(fn func(*NodeAttributes)) NodeOption {
	return nodeOptionFunc(fn)
}

func withShape(s Shape) NodeOption {
	return newNodeOption(func(a *NodeAttributes) {
		a.shape = &s
	})
}

// WithBoxShape sets the node shape to a rectangular box.
func WithBoxShape() NodeOption {
	return withShape(ShapeBox)
}

// WithCircleShape sets the node shape to a circle.
func WithCircleShape() NodeOption {
	return withShape(ShapeCircle)
}

// WithEllipseShape sets the node shape to an ellipse.
// This is the default shape in Graphviz.
func WithEllipseShape() NodeOption {
	return withShape(ShapeEllipse)
}

// WithDiamondShape sets the node shape to a diamond.
func WithDiamondShape() NodeOption {
	return withShape(ShapeDiamond)
}

// WithRecordShape sets the node shape to record-based, useful for structured nodes.
func WithRecordShape() NodeOption {
	return withShape(ShapeRecord)
}

// WithPlaintextShape sets the node to display as plain text with no surrounding shape.
func WithPlaintextShape() NodeOption {
	return withShape(ShapePlaintext)
}

// WithLabel sets the text label displayed on or near the node.
// If not set, the node ID is used as the label by default in Graphviz.
//
// Example:
//
//	n := goraffe.NewNode("node1", goraffe.WithLabel("Start"))
func WithLabel(l string) NodeOption {
	return newNodeOption(func(a *NodeAttributes) {
		a.label = &l
	})
}

// WithColor sets the color of the node border and text.
// Accepts color names (e.g., "red", "blue") or hex values (e.g., "#FF0000").
//
// Example:
//
//	n := goraffe.NewNode("node1", goraffe.WithColor("red"))
func WithColor(c string) NodeOption {
	return newNodeOption(func(a *NodeAttributes) {
		a.color = &c
	})
}

// WithFillColor sets the fill color for the node interior.
// Note: Currently automatically sets style="filled" in the output.
// Accepts color names (e.g., "lightblue") or hex values (e.g., "#E0E0E0").
//
// Example:
//
//	n := goraffe.NewNode("node1", goraffe.WithFillColor("lightblue"))
func WithFillColor(c string) NodeOption {
	return newNodeOption(func(a *NodeAttributes) {
		a.fillColor = &c
	})
}

// WithFontName sets the font family used for the node label text.
//
// Example:
//
//	n := goraffe.NewNode("node1", goraffe.WithFontName("Helvetica"))
func WithFontName(n string) NodeOption {
	return newNodeOption(func(a *NodeAttributes) {
		a.fontName = &n
	})
}

// WithFontSize sets the font size for the node label text in points.
//
// Example:
//
//	n := goraffe.NewNode("node1", goraffe.WithFontSize(14.0))
func WithFontSize(s float64) NodeOption {
	return newNodeOption(func(a *NodeAttributes) {
		a.fontSize = &s
	})
}

// WithNodeAttribute sets a custom attribute on a node.
// This is an escape hatch for Graphviz attributes that don't have typed options.
//
// Example:
//
//	n := NewNode("A",
//	    WithShape(ShapeBox),
//	    WithNodeAttribute("peripheries", "2"),
//	    WithNodeAttribute("tooltip", "Hover text"),
//	)
func WithNodeAttribute(k, v string) NodeOption {
	return newNodeOption(func(a *NodeAttributes) {
		a.setCustom(k, v)
	})
}

// WithHTMLLabel sets an HTML table label on the node.
// HTML labels allow for rich formatting and port-based edge connections.
// The node context is automatically set on all ports defined in the label.
//
// Example:
//
//	label := HTMLTable(
//	    Row(Cell(Text("Input")).Port("in")),
//	    Row(Cell(Text("Output")).Port("out")),
//	)
//	n := NewNode("A", WithHTMLLabel(label))
func WithHTMLLabel(label *HTMLLabel) NodeOption {
	return nodeOptionFunc(func(a *NodeAttributes) {
		a.htmlLabel = label
	})
}

// WithRawHTMLLabel sets a raw HTML label string on the node.
// This is an escape hatch for cases where you want to provide the HTML directly.
// The HTML should be in the format expected by Graphviz (angle bracket delimited).
//
// Example:
//
//	n := NewNode("A", WithRawHTMLLabel("<<table><tr><td>Cell</td></tr></table>>"))
func WithRawHTMLLabel(html string) NodeOption {
	return newNodeOption(func(a *NodeAttributes) {
		a.rawHTMLLabel = &html
	})
}

// WithRecordLabel sets a record label on the node.
// Record labels allow for structured, port-based nodes with fields and groups.
// The node context is automatically set on all ports defined in the label.
// Also sets the shape to record if not already set.
//
// Example:
//
//	label := Record(
//	    Field("input").Port("in"),
//	    Field("output").Port("out"),
//	)
//	n := NewNode("A", WithRecordLabel(label))
func WithRecordLabel(label *RecordLabel) NodeOption {
	return nodeOptionFunc(func(a *NodeAttributes) {
		a.recordLabel = label
		// Automatically set shape to record
		shape := ShapeRecord
		a.shape = &shape
	})
}
