package goraffe

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGraph_DefaultNodeAttrs(t *testing.T) {
	t.Run("returns non-nil default node attributes", func(t *testing.T) {
		asrt := assert.New(t)

		g := NewGraph()

		attrs := g.DefaultNodeAttrs()
		asrt.NotNil(attrs, "expected DefaultNodeAttrs to return non-nil")
	})

	t.Run("returns empty attributes when no defaults set", func(t *testing.T) {
		asrt := assert.New(t)

		g := NewGraph()

		attrs := g.DefaultNodeAttrs()
		asrt.Empty(attrs.Label, "expected default Label to be empty")
		asrt.Empty(attrs.Shape, "expected default Shape to be empty")
		asrt.Empty(attrs.Color, "expected default Color to be empty")
		asrt.Empty(attrs.FillColor, "expected default FillColor to be empty")
		asrt.Empty(attrs.FontName, "expected default FontName to be empty")
		asrt.Zero(attrs.FontSize, "expected default FontSize to be zero")
	})

	t.Run("returns configured default attributes", func(t *testing.T) {
		asrt := assert.New(t)

		g := NewGraph(WithDefaultNodeAttrs(WithShape(ShapeBox), WithFontName("Arial")))

		attrs := g.DefaultNodeAttrs()
		asrt.Equal(ShapeBox, attrs.Shape, "expected default Shape to be box")
		asrt.Equal("Arial", attrs.FontName, "expected default FontName to be Arial")
	})
}

func TestGraph_DefaultEdgeAttrs(t *testing.T) {
	t.Run("returns non-nil default edge attributes", func(t *testing.T) {
		asrt := assert.New(t)

		g := NewGraph()

		attrs := g.DefaultEdgeAttrs()
		asrt.NotNil(attrs, "expected DefaultEdgeAttrs to return non-nil")
	})

	t.Run("returns empty attributes when no defaults set", func(t *testing.T) {
		asrt := assert.New(t)

		g := NewGraph()

		attrs := g.DefaultEdgeAttrs()
		asrt.Empty(attrs.Label, "expected default Label to be empty")
		asrt.Empty(attrs.Color, "expected default Color to be empty")
		asrt.Empty(attrs.Style, "expected default Style to be empty")
		asrt.Empty(attrs.ArrowHead, "expected default ArrowHead to be empty")
		asrt.Empty(attrs.ArrowTail, "expected default ArrowTail to be empty")
		asrt.Zero(attrs.Weight, "expected default Weight to be zero")
	})

	t.Run("returns configured default attributes", func(t *testing.T) {
		asrt := assert.New(t)

		g := NewGraph(WithDefaultEdgeAttrs(WithEdgeColor("gray"), WithEdgeStyle(EdgeStyleDashed)))

		attrs := g.DefaultEdgeAttrs()
		asrt.Equal("gray", attrs.Color, "expected default Color to be gray")
		asrt.Equal(EdgeStyleDashed, attrs.Style, "expected default Style to be dashed")
	})
}

func TestWithDefaultNodeAttrs(t *testing.T) {
	t.Run("sets default node attributes with single option", func(t *testing.T) {
		asrt := assert.New(t)

		g := NewGraph(WithDefaultNodeAttrs(WithShape(ShapeBox)))

		attrs := g.DefaultNodeAttrs()
		asrt.Equal(ShapeBox, attrs.Shape, "expected default Shape to be box")
	})

	t.Run("sets default node attributes with multiple options", func(t *testing.T) {
		asrt := assert.New(t)

		g := NewGraph(WithDefaultNodeAttrs(
			WithShape(ShapeBox),
			WithFontName("Arial"),
			WithFontSize(12.0),
			WithColor("blue"),
		))

		attrs := g.DefaultNodeAttrs()
		asrt.Equal(ShapeBox, attrs.Shape, "expected default Shape to be box")
		asrt.Equal("Arial", attrs.FontName, "expected default FontName to be Arial")
		asrt.Equal(12.0, attrs.FontSize, "expected default FontSize to be 12.0")
		asrt.Equal("blue", attrs.Color, "expected default Color to be blue")
	})

	t.Run("options are applied in order", func(t *testing.T) {
		asrt := assert.New(t)

		g := NewGraph(WithDefaultNodeAttrs(
			WithShape(ShapeBox),
			WithShape(ShapeCircle), // Later option should override
		))

		attrs := g.DefaultNodeAttrs()
		asrt.Equal(ShapeCircle, attrs.Shape, "expected later option to override (Shape should be circle)")
	})

	t.Run("can be combined with other graph options", func(t *testing.T) {
		asrt := assert.New(t)

		g := NewGraph(
			Directed,
			WithDefaultNodeAttrs(WithShape(ShapeBox)),
			WithGraphLabel("Test Graph"),
		)

		asrt.True(g.IsDirected(), "expected graph to be directed")
		asrt.Equal(ShapeBox, g.DefaultNodeAttrs().Shape, "expected default node shape to be box")
	})

	t.Run("with no options creates empty defaults", func(t *testing.T) {
		asrt := assert.New(t)

		g := NewGraph(WithDefaultNodeAttrs())

		attrs := g.DefaultNodeAttrs()
		asrt.Empty(attrs.Shape, "expected Shape to be empty when no options provided")
		asrt.Empty(attrs.Label, "expected Label to be empty when no options provided")
	})
}

func TestWithDefaultEdgeAttrs(t *testing.T) {
	t.Run("sets default edge attributes with single option", func(t *testing.T) {
		asrt := assert.New(t)

		g := NewGraph(WithDefaultEdgeAttrs(WithEdgeColor("gray")))

		attrs := g.DefaultEdgeAttrs()
		asrt.Equal("gray", attrs.Color, "expected default Color to be gray")
	})

	t.Run("sets default edge attributes with multiple options", func(t *testing.T) {
		asrt := assert.New(t)

		g := NewGraph(WithDefaultEdgeAttrs(
			WithEdgeColor("gray"),
			WithEdgeStyle(EdgeStyleDashed),
			WithArrowHead(ArrowVee),
			WithWeight(2.0),
		))

		attrs := g.DefaultEdgeAttrs()
		asrt.Equal("gray", attrs.Color, "expected default Color to be gray")
		asrt.Equal(EdgeStyleDashed, attrs.Style, "expected default Style to be dashed")
		asrt.Equal(ArrowVee, attrs.ArrowHead, "expected default ArrowHead to be vee")
		asrt.Equal(2.0, attrs.Weight, "expected default Weight to be 2.0")
	})

	t.Run("options are applied in order", func(t *testing.T) {
		asrt := assert.New(t)

		g := NewGraph(WithDefaultEdgeAttrs(
			WithEdgeStyle(EdgeStyleDashed),
			WithEdgeStyle(EdgeStyleDotted), // Later option should override
		))

		attrs := g.DefaultEdgeAttrs()
		asrt.Equal(EdgeStyleDotted, attrs.Style, "expected later option to override (Style should be dotted)")
	})

	t.Run("can be combined with other graph options", func(t *testing.T) {
		asrt := assert.New(t)

		g := NewGraph(
			Directed,
			WithDefaultEdgeAttrs(WithEdgeColor("gray")),
			WithGraphLabel("Test Graph"),
		)

		asrt.True(g.IsDirected(), "expected graph to be directed")
		asrt.Equal("gray", g.DefaultEdgeAttrs().Color, "expected default edge color to be gray")
	})

	t.Run("with no options creates empty defaults", func(t *testing.T) {
		asrt := assert.New(t)

		g := NewGraph(WithDefaultEdgeAttrs())

		attrs := g.DefaultEdgeAttrs()
		asrt.Empty(attrs.Style, "expected Style to be empty when no options provided")
		asrt.Empty(attrs.Color, "expected Color to be empty when no options provided")
	})
}

func TestGraph_DefaultAttrs_BothNodeAndEdge(t *testing.T) {
	t.Run("can set both default node and edge attributes", func(t *testing.T) {
		asrt := assert.New(t)

		g := NewGraph(
			Directed,
			WithDefaultNodeAttrs(WithShape(ShapeBox), WithFontName("Arial")),
			WithDefaultEdgeAttrs(WithEdgeColor("gray"), WithEdgeStyle(EdgeStyleDashed)),
		)

		nodeAttrs := g.DefaultNodeAttrs()
		asrt.Equal(ShapeBox, nodeAttrs.Shape, "expected default node shape to be box")
		asrt.Equal("Arial", nodeAttrs.FontName, "expected default node font to be Arial")

		edgeAttrs := g.DefaultEdgeAttrs()
		asrt.Equal("gray", edgeAttrs.Color, "expected default edge color to be gray")
		asrt.Equal(EdgeStyleDashed, edgeAttrs.Style, "expected default edge style to be dashed")
	})

	t.Run("default attributes do not affect each other", func(t *testing.T) {
		asrt := assert.New(t)

		g := NewGraph(
			WithDefaultNodeAttrs(WithShape(ShapeBox)),
		)

		nodeAttrs := g.DefaultNodeAttrs()
		edgeAttrs := g.DefaultEdgeAttrs()

		asrt.Equal(ShapeBox, nodeAttrs.Shape, "expected node defaults to be set")
		asrt.Empty(edgeAttrs.Color, "expected edge defaults to be unaffected")
	})

	t.Run("multiple calls to WithDefaultNodeAttrs only applies last one", func(t *testing.T) {
		asrt := assert.New(t)

		g := NewGraph(
			WithDefaultNodeAttrs(WithShape(ShapeBox)),
			WithDefaultNodeAttrs(WithShape(ShapeCircle)),
		)

		attrs := g.DefaultNodeAttrs()
		asrt.Equal(ShapeCircle, attrs.Shape, "expected second WithDefaultNodeAttrs to override first")
	})

	t.Run("multiple calls to WithDefaultEdgeAttrs only applies last one", func(t *testing.T) {
		asrt := assert.New(t)

		g := NewGraph(
			WithDefaultEdgeAttrs(WithEdgeColor("red")),
			WithDefaultEdgeAttrs(WithEdgeColor("blue")),
		)

		attrs := g.DefaultEdgeAttrs()
		asrt.Equal("blue", attrs.Color, "expected second WithDefaultEdgeAttrs to override first")
	})
}

func TestGraph_DefaultAttrs_IndependentFromInstanceAttrs(t *testing.T) {
	t.Run("default node attrs are independent from individual node attrs", func(t *testing.T) {
		asrt := assert.New(t)

		g := NewGraph(WithDefaultNodeAttrs(WithShape(ShapeBox)))
		n := NewNode("A", WithShape(ShapeCircle))

		g.AddNode(n)

		// Default should not be affected by individual node
		asrt.Equal(ShapeBox, g.DefaultNodeAttrs().Shape, "expected default to remain ShapeBox")
		asrt.Equal(ShapeCircle, n.Attrs().Shape, "expected node to have ShapeCircle")
	})

	t.Run("default edge attrs are independent from individual edge attrs", func(t *testing.T) {
		asrt := assert.New(t)

		g := NewGraph(WithDefaultEdgeAttrs(WithEdgeColor("gray")))
		n1 := NewNode("A")
		n2 := NewNode("B")
		e := g.AddEdge(n1, n2, WithEdgeColor("red"))

		// Default should not be affected by individual edge
		asrt.Equal("gray", g.DefaultEdgeAttrs().Color, "expected default to remain gray")
		asrt.Equal("red", e.Attrs().Color, "expected edge to have red color")
	})
}

func TestWithDefaultNodeAttrs_UsingNodeAttributesStruct(t *testing.T) {
	t.Run("accepts NodeAttributes struct as option", func(t *testing.T) {
		asrt := assert.New(t)

		template := NodeAttributes{
			Shape:     ShapeBox,
			FontName:  "Arial",
			FontSize:  12.0,
			Color:     "blue",
			FillColor: "lightblue",
		}

		g := NewGraph(WithDefaultNodeAttrs(template))

		attrs := g.DefaultNodeAttrs()
		asrt.Equal(ShapeBox, attrs.Shape, "expected Shape from NodeAttributes struct")
		asrt.Equal("Arial", attrs.FontName, "expected FontName from NodeAttributes struct")
		asrt.Equal(12.0, attrs.FontSize, "expected FontSize from NodeAttributes struct")
		asrt.Equal("blue", attrs.Color, "expected Color from NodeAttributes struct")
		asrt.Equal("lightblue", attrs.FillColor, "expected FillColor from NodeAttributes struct")
	})

	t.Run("combines NodeAttributes struct with additional options", func(t *testing.T) {
		asrt := assert.New(t)

		template := NodeAttributes{
			Shape:    ShapeBox,
			FontName: "Arial",
		}

		g := NewGraph(WithDefaultNodeAttrs(template, WithLabel("Default Label"), WithFontSize(14.0)))

		attrs := g.DefaultNodeAttrs()
		asrt.Equal(ShapeBox, attrs.Shape, "expected Shape from NodeAttributes struct")
		asrt.Equal("Arial", attrs.FontName, "expected FontName from NodeAttributes struct")
		asrt.Equal("Default Label", attrs.Label, "expected Label from WithLabel option")
		asrt.Equal(14.0, attrs.FontSize, "expected FontSize from WithFontSize option")
	})

	t.Run("later options override NodeAttributes struct", func(t *testing.T) {
		asrt := assert.New(t)

		template := NodeAttributes{
			Shape:    ShapeBox,
			FontName: "Arial",
			Color:    "red",
		}

		// Override Color with a later option
		g := NewGraph(WithDefaultNodeAttrs(template, WithColor("blue")))

		attrs := g.DefaultNodeAttrs()
		asrt.Equal(ShapeBox, attrs.Shape, "expected Shape from NodeAttributes struct")
		asrt.Equal("Arial", attrs.FontName, "expected FontName from NodeAttributes struct")
		asrt.Equal("blue", attrs.Color, "expected Color from later option to override struct")
	})

	t.Run("only copies non-zero fields from NodeAttributes struct", func(t *testing.T) {
		asrt := assert.New(t)

		// Template with only some fields set
		template := NodeAttributes{
			Shape: ShapeBox,
			// Label, Color, FillColor, FontName are zero values (empty strings)
			// FontSize is zero value (0.0)
		}

		g := NewGraph(WithDefaultNodeAttrs(template, WithLabel("Test Label")))

		attrs := g.DefaultNodeAttrs()
		asrt.Equal(ShapeBox, attrs.Shape, "expected Shape to be set from struct")
		asrt.Equal("Test Label", attrs.Label, "expected Label to be set from option")
		asrt.Empty(attrs.Color, "expected Color to remain empty (zero value not copied)")
		asrt.Empty(attrs.FillColor, "expected FillColor to remain empty (zero value not copied)")
		asrt.Empty(attrs.FontName, "expected FontName to remain empty (zero value not copied)")
		asrt.Zero(attrs.FontSize, "expected FontSize to remain zero (zero value not copied)")
	})

	t.Run("can use same NodeAttributes struct for defaults and individual nodes", func(t *testing.T) {
		asrt := assert.New(t)

		template := NodeAttributes{
			Shape:    ShapeBox,
			FontName: "Arial",
			FontSize: 12.0,
		}

		g := NewGraph(WithDefaultNodeAttrs(template))
		n := NewNode("A", template, WithLabel("Node A"))

		// Both should have the template values
		asrt.Equal(ShapeBox, g.DefaultNodeAttrs().Shape, "expected default Shape from template")
		asrt.Equal("Arial", g.DefaultNodeAttrs().FontName, "expected default FontName from template")
		asrt.Equal(ShapeBox, n.Attrs().Shape, "expected node Shape from template")
		asrt.Equal("Arial", n.Attrs().FontName, "expected node FontName from template")
		asrt.Equal("Node A", n.Attrs().Label, "expected node Label from option")
	})
}

func TestWithDefaultEdgeAttrs_UsingEdgeAttributesStruct(t *testing.T) {
	t.Run("accepts EdgeAttributes struct as option", func(t *testing.T) {
		asrt := assert.New(t)

		template := EdgeAttributes{
			Style:     EdgeStyleDashed,
			Color:     "gray",
			ArrowHead: ArrowVee,
			ArrowTail: ArrowNormal,
			Weight:    2.5,
		}

		g := NewGraph(WithDefaultEdgeAttrs(template))

		attrs := g.DefaultEdgeAttrs()
		asrt.Equal(EdgeStyleDashed, attrs.Style, "expected Style from EdgeAttributes struct")
		asrt.Equal("gray", attrs.Color, "expected Color from EdgeAttributes struct")
		asrt.Equal(ArrowVee, attrs.ArrowHead, "expected ArrowHead from EdgeAttributes struct")
		asrt.Equal(ArrowNormal, attrs.ArrowTail, "expected ArrowTail from EdgeAttributes struct")
		asrt.Equal(2.5, attrs.Weight, "expected Weight from EdgeAttributes struct")
	})

	t.Run("combines EdgeAttributes struct with additional options", func(t *testing.T) {
		asrt := assert.New(t)

		template := EdgeAttributes{
			Style: EdgeStyleDashed,
			Color: "gray",
		}

		g := NewGraph(WithDefaultEdgeAttrs(template, WithEdgeLabel("Default"), WithWeight(1.0)))

		attrs := g.DefaultEdgeAttrs()
		asrt.Equal(EdgeStyleDashed, attrs.Style, "expected Style from EdgeAttributes struct")
		asrt.Equal("gray", attrs.Color, "expected Color from EdgeAttributes struct")
		asrt.Equal("Default", attrs.Label, "expected Label from WithEdgeLabel option")
		asrt.Equal(1.0, attrs.Weight, "expected Weight from WithWeight option")
	})

	t.Run("later options override EdgeAttributes struct", func(t *testing.T) {
		asrt := assert.New(t)

		template := EdgeAttributes{
			Style: EdgeStyleDashed,
			Color: "red",
			Weight: 1.0,
		}

		// Override Color and Weight with later options
		g := NewGraph(WithDefaultEdgeAttrs(template, WithEdgeColor("blue"), WithWeight(3.0)))

		attrs := g.DefaultEdgeAttrs()
		asrt.Equal(EdgeStyleDashed, attrs.Style, "expected Style from EdgeAttributes struct")
		asrt.Equal("blue", attrs.Color, "expected Color from later option to override struct")
		asrt.Equal(3.0, attrs.Weight, "expected Weight from later option to override struct")
	})

	t.Run("only copies non-zero fields from EdgeAttributes struct", func(t *testing.T) {
		asrt := assert.New(t)

		// Template with only some fields set
		template := EdgeAttributes{
			Style: EdgeStyleDashed,
			// Label, Color are zero values (empty strings)
			// ArrowHead, ArrowTail are zero values (empty strings)
			// Weight is zero value (0.0)
		}

		g := NewGraph(WithDefaultEdgeAttrs(template, WithEdgeLabel("Test Label")))

		attrs := g.DefaultEdgeAttrs()
		asrt.Equal(EdgeStyleDashed, attrs.Style, "expected Style to be set from struct")
		asrt.Equal("Test Label", attrs.Label, "expected Label to be set from option")
		asrt.Empty(attrs.Color, "expected Color to remain empty (zero value not copied)")
		asrt.Empty(attrs.ArrowHead, "expected ArrowHead to remain empty (zero value not copied)")
		asrt.Empty(attrs.ArrowTail, "expected ArrowTail to remain empty (zero value not copied)")
		asrt.Zero(attrs.Weight, "expected Weight to remain zero (zero value not copied)")
	})

	t.Run("can use same EdgeAttributes struct for defaults and individual edges", func(t *testing.T) {
		asrt := assert.New(t)

		template := EdgeAttributes{
			Style:  EdgeStyleDashed,
			Color:  "gray",
			Weight: 1.5,
		}

		g := NewGraph(WithDefaultEdgeAttrs(template))
		n1 := NewNode("A")
		n2 := NewNode("B")
		e := g.AddEdge(n1, n2, template, WithEdgeLabel("A to B"))

		// Both should have the template values
		asrt.Equal(EdgeStyleDashed, g.DefaultEdgeAttrs().Style, "expected default Style from template")
		asrt.Equal("gray", g.DefaultEdgeAttrs().Color, "expected default Color from template")
		asrt.Equal(EdgeStyleDashed, e.Attrs().Style, "expected edge Style from template")
		asrt.Equal("gray", e.Attrs().Color, "expected edge Color from template")
		asrt.Equal("A to B", e.Attrs().Label, "expected edge Label from option")
	})
}
