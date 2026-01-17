package goraffe

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewNode_SetsID(t *testing.T) {
	asrt := assert.New(t)

	n := NewNode("node1")

	asrt.Equal("node1", n.id, "expected node ID to be set to 'node1'")
}

func TestNewNode_EmptyID(t *testing.T) {
	asrt := assert.New(t)

	n := NewNode("")

	asrt.NotNil(n, "expected NewNode to return a non-nil node even with empty ID")
	asrt.Empty(n.id, "expected node ID to be empty string")
}

func TestNode_ID_ReturnsCorrectValue(t *testing.T) {
	t.Run("returns simple ID", func(t *testing.T) {
		asrt := assert.New(t)

		n := NewNode("simple")

		asrt.Equal("simple", n.ID(), "expected ID() to return 'simple'")
	})

	t.Run("returns ID with special characters", func(t *testing.T) {
		asrt := assert.New(t)

		n := NewNode("node-with-dashes")

		asrt.Equal("node-with-dashes", n.ID(), "expected ID() to return exact ID with dashes")
	})

	t.Run("returns ID with spaces", func(t *testing.T) {
		asrt := assert.New(t)

		n := NewNode("node with spaces")

		asrt.Equal("node with spaces", n.ID(), "expected ID() to return exact ID with spaces")
	})

	t.Run("returns ID with numbers", func(t *testing.T) {
		asrt := assert.New(t)

		n := NewNode("node123")

		asrt.Equal("node123", n.ID(), "expected ID() to return exact ID with numbers")
	})

	t.Run("returns empty ID unchanged", func(t *testing.T) {
		asrt := assert.New(t)

		n := NewNode("")

		asrt.Empty(n.ID(), "expected ID() to return empty string")
	})
}

func TestNode_Attrs_ReturnsAttributes(t *testing.T) {
	t.Run("returns non-nil attributes", func(t *testing.T) {
		asrt := assert.New(t)

		n := NewNode("test")

		attrs := n.Attrs()
		asrt.NotNil(attrs, "expected Attrs() to return non-nil NodeAttributes")
	})

	t.Run("returns same instance on multiple calls", func(t *testing.T) {
		asrt := assert.New(t)

		n := NewNode("test")

		attrs1 := n.Attrs()
		attrs2 := n.Attrs()

		asrt.Same(attrs1, attrs2, "expected Attrs() to return the same NodeAttributes instance")
	})

	t.Run("returns attributes with zero values for new node", func(t *testing.T) {
		asrt := assert.New(t)

		n := NewNode("test")

		attrs := n.Attrs()
		asrt.Empty(attrs.Label(), "expected Label to be empty for new node")
		asrt.Empty(attrs.Shape(), "expected Shape to be empty for new node")
		asrt.Empty(attrs.Color(), "expected Color to be empty for new node")
		asrt.Empty(attrs.FillColor(), "expected FillColor to be empty for new node")
		asrt.Empty(attrs.FontName(), "expected FontName to be empty for new node")
		asrt.Equal(0.0, attrs.FontSize(), "expected FontSize to be zero for new node")
	})
}

func TestNewNode_WithOptions(t *testing.T) {
	t.Run("applies single option", func(t *testing.T) {
		asrt := assert.New(t)

		n := NewNode("test", WithBoxShape())

		asrt.Equal(ShapeBox, n.Attrs().Shape(), "expected Shape to be set to ShapeBox")
	})

	t.Run("applies WithLabel option", func(t *testing.T) {
		asrt := assert.New(t)

		n := NewNode("test", WithLabel("My Label"))

		asrt.Equal("My Label", n.Attrs().Label(), "expected Label to be set")
	})

	t.Run("applies WithColor option", func(t *testing.T) {
		asrt := assert.New(t)

		n := NewNode("test", WithColor("red"))

		asrt.Equal("red", n.Attrs().Color(), "expected Color to be set to red")
	})

	t.Run("applies WithFillColor option", func(t *testing.T) {
		asrt := assert.New(t)

		n := NewNode("test", WithFillColor("blue"))

		asrt.Equal("blue", n.Attrs().FillColor(), "expected FillColor to be set to blue")
	})

	t.Run("applies WithFontName option", func(t *testing.T) {
		asrt := assert.New(t)

		n := NewNode("test", WithFontName("Arial"))

		asrt.Equal("Arial", n.Attrs().FontName(), "expected FontName to be set to Arial")
	})

	t.Run("applies WithFontSize option", func(t *testing.T) {
		asrt := assert.New(t)

		n := NewNode("test", WithFontSize(14.0))

		asrt.Equal(14.0, n.Attrs().FontSize(), "expected FontSize to be set to 14.0")
	})

	t.Run("works with no options", func(t *testing.T) {
		asrt := assert.New(t)

		n := NewNode("test")

		asrt.NotNil(n, "expected NewNode to work with no options")
		asrt.Equal("test", n.ID(), "expected ID to be set correctly")
	})
}

func TestNewNode_WithMultipleOptions(t *testing.T) {
	t.Run("applies multiple different options", func(t *testing.T) {
		asrt := assert.New(t)

		n := NewNode("test",
			WithBoxShape(),
			WithLabel("My Node"),
			WithColor("red"),
		)

		asrt.Equal(ShapeBox, n.Attrs().Shape(), "expected Shape to be set")
		asrt.Equal("My Node", n.Attrs().Label(), "expected Label to be set")
		asrt.Equal("red", n.Attrs().Color(), "expected Color to be set")
	})

	t.Run("applies all available option types", func(t *testing.T) {
		asrt := assert.New(t)

		n := NewNode("test",
			WithCircleShape(),
			WithLabel("All Options"),
			WithColor("blue"),
			WithFillColor("lightblue"),
			WithFontName("Helvetica"),
			WithFontSize(16.0),
		)

		asrt.Equal(ShapeCircle, n.Attrs().Shape(), "expected Shape to be set")
		asrt.Equal("All Options", n.Attrs().Label(), "expected Label to be set")
		asrt.Equal("blue", n.Attrs().Color(), "expected Color to be set")
		asrt.Equal("lightblue", n.Attrs().FillColor(), "expected FillColor to be set")
		asrt.Equal("Helvetica", n.Attrs().FontName(), "expected FontName to be set")
		asrt.Equal(16.0, n.Attrs().FontSize(), "expected FontSize to be set")
	})
}

func TestNewNode_WithNodeAttributesStruct(t *testing.T) {
	t.Run("applies NodeAttributes as an option", func(t *testing.T) {
		asrt := assert.New(t)

		shape := ShapeBox
		fontName := "Arial"
		fontSize := 12.0
		commonAttrs := NodeAttributes{
			shape:    &shape,
			fontName: &fontName,
			fontSize: &fontSize,
		}

		n := NewNode("test", commonAttrs)

		asrt.Equal(ShapeBox, n.Attrs().Shape(), "expected Shape to be set from NodeAttributes")
		asrt.Equal("Arial", n.Attrs().FontName(), "expected FontName to be set from NodeAttributes")
		asrt.Equal(12.0, n.Attrs().FontSize(), "expected FontSize to be set from NodeAttributes")
	})

	t.Run("applies NodeAttributes with additional options", func(t *testing.T) {
		asrt := assert.New(t)

		shape := ShapeBox
		fontName := "Arial"
		commonAttrs := NodeAttributes{
			shape:    &shape,
			fontName: &fontName,
		}

		n := NewNode("test", commonAttrs, WithLabel("Custom Label"))

		asrt.Equal(ShapeBox, n.Attrs().Shape(), "expected Shape from NodeAttributes")
		asrt.Equal("Arial", n.Attrs().FontName(), "expected FontName from NodeAttributes")
		asrt.Equal("Custom Label", n.Attrs().Label(), "expected Label from WithLabel option")
	})

	t.Run("NodeAttributes only copies non-nil fields", func(t *testing.T) {
		asrt := assert.New(t)

		// Create attrs with only some fields set
		shape := ShapeBox
		commonAttrs := NodeAttributes{
			shape: &shape,
			// Label, Color, etc. are nil
		}

		n := NewNode("test", commonAttrs, WithLabel("Test"))

		asrt.Equal(ShapeBox, n.Attrs().Shape(), "expected Shape to be set")
		asrt.Equal("Test", n.Attrs().Label(), "expected Label to be set")
		asrt.Empty(n.Attrs().Color(), "expected Color to remain empty (nil value not copied)")
	})
}

func TestNewNode_OptionsAppliedInOrder(t *testing.T) {
	t.Run("later options override earlier ones", func(t *testing.T) {
		asrt := assert.New(t)

		n := NewNode("test",
			WithCircleShape(),
			WithBoxShape(),
		)

		asrt.Equal(ShapeBox, n.Attrs().Shape(), "expected later shape option to override earlier one")
	})

	t.Run("later label overrides earlier label", func(t *testing.T) {
		asrt := assert.New(t)

		n := NewNode("test",
			WithLabel("First"),
			WithLabel("Second"),
		)

		asrt.Equal("Second", n.Attrs().Label(), "expected later label to override earlier one")
	})

	t.Run("NodeAttributes then individual options", func(t *testing.T) {
		asrt := assert.New(t)

		shape := ShapeCircle
		label := "From Attrs"
		commonAttrs := NodeAttributes{
			shape: &shape,
			label: &label,
		}

		n := NewNode("test", commonAttrs, WithLabel("Overridden"))

		asrt.Equal(ShapeCircle, n.Attrs().Shape(), "expected Shape from NodeAttributes")
		asrt.Equal("Overridden", n.Attrs().Label(), "expected Label to be overridden by later option")
	})

	t.Run("individual options then NodeAttributes", func(t *testing.T) {
		asrt := assert.New(t)

		shape := ShapeBox
		label := "From Attrs"
		commonAttrs := NodeAttributes{
			shape: &shape,
			label: &label,
		}

		n := NewNode("test", WithCircleShape(), commonAttrs)

		asrt.Equal(ShapeBox, n.Attrs().Shape(), "expected Shape from later NodeAttributes to override")
		asrt.Equal("From Attrs", n.Attrs().Label(), "expected Label from NodeAttributes")
	})

	t.Run("NodeAttributes with nil values should not override non-nil values", func(t *testing.T) {
		asrt := assert.New(t)

		// This test demonstrates the selective copying requirement
		// A reusable NodeAttributes template with only some fields set
		shape := ShapeBox
		fontName := "Arial"
		template := NodeAttributes{
			shape:    &shape,
			fontName: &fontName,
			// Label, Color, FillColor, FontSize are nil
		}

		// Apply options first, then the template
		// The template should only override Shape and FontName, not Label
		n := NewNode("test",
			WithLabel("Important Label"),
			WithColor("red"),
			template, // This should NOT override Label and Color to nil
		)

		asrt.Equal(ShapeBox, n.Attrs().Shape(), "expected Shape from template")
		asrt.Equal("Arial", n.Attrs().FontName(), "expected FontName from template")
		asrt.Equal("Important Label", n.Attrs().Label(), "expected Label to NOT be overridden by template's nil value")
		asrt.Equal("red", n.Attrs().Color(), "expected Color to NOT be overridden by template's nil value")
		asrt.Empty(n.Attrs().FillColor(), "expected FillColor to remain empty")
		asrt.Equal(0.0, n.Attrs().FontSize(), "expected FontSize to remain zero")
	})
}

func TestWithHTMLLabel_SetsLabel(t *testing.T) {
	t.Run("sets HTML label on node", func(t *testing.T) {
		asrt := assert.New(t)

		label := HTMLTable(
			Row(Cell(Text("Test Cell"))),
		)
		n := NewNode("A", WithHTMLLabel(label))

		asrt.NotNil(n.Attrs().HTMLLabel(), "expected HTMLLabel to be set")
		asrt.Same(label, n.Attrs().HTMLLabel(), "expected HTMLLabel to be the same instance")
	})
}

func TestWithRawHTMLLabel_SetsLabel(t *testing.T) {
	t.Run("sets raw HTML label on node", func(t *testing.T) {
		asrt := assert.New(t)

		rawHTML := "<<table><tr><td>Cell</td></tr></table>>"
		n := NewNode("A", WithRawHTMLLabel(rawHTML))

		asrt.Equal(rawHTML, n.Attrs().RawHTMLLabel(), "expected RawHTMLLabel to be set")
	})
}

func TestDOT_Node_WithHTMLLabel(t *testing.T) {
	t.Run("outputs node with HTML label", func(t *testing.T) {
		asrt := assert.New(t)

		label := HTMLTable(
			Row(Cell(Text("cell"))),
		)
		n := NewNode("A", WithHTMLLabel(label))

		output := n.String()
		asrt.Contains(output, "label=<", "expected label to start with angle bracket")
		asrt.Contains(output, "<table>", "expected table tag in output")
		asrt.Contains(output, "<tr>", "expected tr tag in output")
		asrt.Contains(output, "<td>", "expected td tag in output")
		asrt.Contains(output, "cell", "expected cell text in output")
	})
}

func TestDOT_Node_WithHTMLLabel_Ports(t *testing.T) {
	t.Run("wires port node context automatically", func(t *testing.T) {
		asrt := assert.New(t)

		cell := Cell(Text("output")).Port("out")
		label := HTMLTable(Row(cell))
		_ = NewNode("A", WithHTMLLabel(label))

		port := cell.GetPort()
		asrt.NotNil(port, "expected port to exist")
		asrt.Equal("A", port.NodeID(), "expected port node context to be set to node ID")
		asrt.Equal("out", port.ID(), "expected port ID to be 'out'")
	})

	t.Run("outputs HTML label with port", func(t *testing.T) {
		asrt := assert.New(t)

		label := HTMLTable(
			Row(Cell(Text("output")).Port("out")),
		)
		n := NewNode("A", WithHTMLLabel(label))

		output := n.String()
		asrt.Contains(output, "port=\"out\"", "expected port attribute in td tag")
	})
}

func TestDOT_Node_WithRawHTMLLabel(t *testing.T) {
	t.Run("outputs node with raw HTML label", func(t *testing.T) {
		asrt := assert.New(t)

		rawHTML := "<<table><tr><td>Cell</td></tr></table>>"
		n := NewNode("A", WithRawHTMLLabel(rawHTML))

		output := n.String()
		asrt.Contains(output, "label=<<table><tr><td>Cell</td></tr></table>>", "expected raw HTML in output")
	})
}

func TestDOT_HTMLLabel_NotDoubleEscaped(t *testing.T) {
	t.Run("HTML labels are not escaped", func(t *testing.T) {
		asrt := assert.New(t)

		label := HTMLTable(
			Row(Cell(Text("<special>"))),
		)
		n := NewNode("A", WithHTMLLabel(label))

		output := n.String()
		// The HTML label should contain angle brackets as-is, not escaped
		asrt.Contains(output, "label=<", "expected label to start with unescaped angle bracket")
		asrt.Contains(output, "<table>", "expected unescaped table tag")
	})

	t.Run("HTML labels take precedence over regular label", func(t *testing.T) {
		asrt := assert.New(t)

		label := HTMLTable(
			Row(Cell(Text("HTML"))),
		)
		n := NewNode("A", WithLabel("Regular"), WithHTMLLabel(label))

		output := n.String()
		asrt.Contains(output, "label=<", "expected HTML label format")
		asrt.Contains(output, "HTML", "expected HTML label text")
		asrt.NotContains(output, "Regular", "expected regular label to be overridden")
	})

	t.Run("raw HTML label takes precedence over HTML label", func(t *testing.T) {
		asrt := assert.New(t)

		label := HTMLTable(
			Row(Cell(Text("HTML"))),
		)
		rawHTML := "<<table><tr><td>Raw</td></tr></table>>"
		n := NewNode("A", WithHTMLLabel(label), WithRawHTMLLabel(rawHTML))

		output := n.String()
		asrt.Contains(output, "Raw", "expected raw HTML label text")
		asrt.NotContains(output, "HTML", "expected HTML label to be overridden by raw")
	})
}

func TestWithRecordLabel_SetsLabel(t *testing.T) {
	t.Run("sets record label on node", func(t *testing.T) {
		asrt := assert.New(t)

		label := Record(Field("a"), Field("b"))
		n := NewNode("A", WithRecordLabel(label))

		asrt.NotNil(n.Attrs().RecordLabel(), "expected RecordLabel to be set")
		asrt.Same(label, n.Attrs().RecordLabel(), "expected RecordLabel to be the same instance")
	})
}

func TestWithRecordLabel_SetsShape(t *testing.T) {
	t.Run("sets shape to record when record label is used", func(t *testing.T) {
		asrt := assert.New(t)

		label := Record(Field("a"))
		n := NewNode("A", WithRecordLabel(label))

		asrt.Equal(ShapeRecord, n.Attrs().Shape(), "expected shape to be set to record")
	})

	t.Run("does not override explicitly set shape", func(t *testing.T) {
		asrt := assert.New(t)

		label := Record(Field("a"))
		n := NewNode("A", WithBoxShape(), WithRecordLabel(label))

		asrt.Equal(ShapeRecord, n.Attrs().Shape(), "expected record label to set shape to record")
	})
}

func TestDOT_Node_WithRecordLabel_Simple(t *testing.T) {
	t.Run("outputs node with simple record label", func(t *testing.T) {
		asrt := assert.New(t)

		label := Record(Field("a"), Field("b"))
		n := NewNode("A", WithRecordLabel(label))

		output := n.String()
		asrt.Contains(output, `label="a | b"`, "expected quoted record label")
		asrt.Contains(output, `shape="record"`, "expected shape to be record")
	})
}

func TestDOT_Node_WithRecordLabel_WithPorts(t *testing.T) {
	t.Run("outputs record label with ports", func(t *testing.T) {
		asrt := assert.New(t)

		label := Record(
			Field("input").Port("in"),
			Field("output").Port("out"),
		)
		n := NewNode("A", WithRecordLabel(label))

		output := n.String()
		asrt.Contains(output, `label="<in> input | <out> output"`, "expected record label with ports")
	})

	t.Run("wires port node context automatically", func(t *testing.T) {
		asrt := assert.New(t)

		field := Field("output").Port("out")
		label := Record(field)
		_ = NewNode("A", WithRecordLabel(label))

		port := field.GetPort()
		asrt.NotNil(port, "expected port to exist")
		asrt.Equal("A", port.NodeID(), "expected port node context to be set to node ID")
		asrt.Equal("out", port.ID(), "expected port ID to be 'out'")
	})
}

func TestDOT_Node_WithRecordLabel_Nested(t *testing.T) {
	t.Run("outputs nested record label", func(t *testing.T) {
		asrt := assert.New(t)

		label := Record(
			Field("header"),
			FieldGroup(Field("left"), Field("right")),
			Field("footer"),
		)
		n := NewNode("A", WithRecordLabel(label))

		output := n.String()
		asrt.Contains(output, `label="header | { left | right } | footer"`, "expected nested record label")
	})
}

func TestDOT_Edge_ToRecordPort(t *testing.T) {
	t.Run("creates edge from record port", func(t *testing.T) {
		asrt := assert.New(t)

		field := Field("output").Port("out")
		label := Record(field)
		nodeA := NewNode("A", WithRecordLabel(label))
		nodeB := NewNode("B")

		g := NewGraph()
		port := field.GetPort()
		edge, err := g.AddEdge(nodeA, nodeB, FromPort(port))
		asrt.NoError(err)

		output := edge.ToString(true) // true for directed
		asrt.Contains(output, `"A":"out"`, "expected edge from record port")
	})
}
