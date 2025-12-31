package goraffe

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Test WithNodeAttribute option
func TestWithNodeAttribute_SetsCustom(t *testing.T) {
	t.Run("sets a single custom attribute on new node", func(t *testing.T) {
		asrt := assert.New(t)

		n := NewNode("test", WithNodeAttribute("peripheries", "2"))

		custom := n.Attrs().Custom()
		asrt.Len(custom, 1, "expected custom map to have 1 entry")
		asrt.Equal("2", custom["peripheries"], "expected peripheries to be set to '2'")
	})

	t.Run("sets multiple custom attributes", func(t *testing.T) {
		asrt := assert.New(t)

		n := NewNode("test",
			WithNodeAttribute("peripheries", "2"),
			WithNodeAttribute("tooltip", "Hover text"),
			WithNodeAttribute("URL", "http://example.com"),
		)

		custom := n.Attrs().Custom()
		asrt.Len(custom, 3, "expected custom map to have 3 entries")
		asrt.Equal("2", custom["peripheries"], "expected peripheries to be set")
		asrt.Equal("Hover text", custom["tooltip"], "expected tooltip to be set")
		asrt.Equal("http://example.com", custom["URL"], "expected URL to be set")
	})

	t.Run("combines with typed options", func(t *testing.T) {
		asrt := assert.New(t)

		n := NewNode("test",
			WithShape(ShapeBox),
			WithLabel("My Node"),
			WithNodeAttribute("peripheries", "2"),
		)

		asrt.Equal(ShapeBox, n.Attrs().Shape, "expected Shape to be set")
		asrt.Equal("My Node", n.Attrs().Label, "expected Label to be set")
		custom := n.Attrs().Custom()
		asrt.Equal("2", custom["peripheries"], "expected custom attribute to be set")
	})

	t.Run("custom attributes accumulate", func(t *testing.T) {
		asrt := assert.New(t)

		n := NewNode("test",
			WithNodeAttribute("attr1", "value1"),
			WithNodeAttribute("attr2", "value2"),
			WithNodeAttribute("attr3", "value3"),
		)

		custom := n.Attrs().Custom()
		asrt.Len(custom, 3, "expected all custom attributes to accumulate")
		asrt.Equal("value1", custom["attr1"], "expected attr1 to be set")
		asrt.Equal("value2", custom["attr2"], "expected attr2 to be set")
		asrt.Equal("value3", custom["attr3"], "expected attr3 to be set")
	})

	t.Run("later custom attribute overwrites earlier with same key", func(t *testing.T) {
		asrt := assert.New(t)

		n := NewNode("test",
			WithNodeAttribute("peripheries", "2"),
			WithNodeAttribute("peripheries", "3"),
		)

		custom := n.Attrs().Custom()
		asrt.Len(custom, 1, "expected only one entry for duplicated key")
		asrt.Equal("3", custom["peripheries"], "expected later value to overwrite earlier one")
	})
}

// Test WithEdgeAttribute option
func TestWithEdgeAttribute_SetsCustom(t *testing.T) {
	t.Run("sets a single custom attribute on new edge", func(t *testing.T) {
		asrt := assert.New(t)
		g := NewGraph()
		n1 := NewNode("A")
		n2 := NewNode("B")

		e := g.AddEdge(n1, n2, WithEdgeAttribute("penwidth", "2.0"))

		custom := e.Attrs().Custom()
		asrt.Len(custom, 1, "expected custom map to have 1 entry")
		asrt.Equal("2.0", custom["penwidth"], "expected penwidth to be set to '2.0'")
	})

	t.Run("sets multiple custom attributes", func(t *testing.T) {
		asrt := assert.New(t)
		g := NewGraph()
		n1 := NewNode("A")
		n2 := NewNode("B")

		e := g.AddEdge(n1, n2,
			WithEdgeAttribute("penwidth", "2.0"),
			WithEdgeAttribute("dir", "both"),
			WithEdgeAttribute("constraint", "false"),
		)

		custom := e.Attrs().Custom()
		asrt.Len(custom, 3, "expected custom map to have 3 entries")
		asrt.Equal("2.0", custom["penwidth"], "expected penwidth to be set")
		asrt.Equal("both", custom["dir"], "expected dir to be set")
		asrt.Equal("false", custom["constraint"], "expected constraint to be set")
	})

	t.Run("combines with typed options", func(t *testing.T) {
		asrt := assert.New(t)
		g := NewGraph()
		n1 := NewNode("A")
		n2 := NewNode("B")

		e := g.AddEdge(n1, n2,
			WithEdgeLabel("connects"),
			WithEdgeColor("red"),
			WithEdgeAttribute("penwidth", "2.0"),
		)

		asrt.Equal("connects", e.Attrs().Label, "expected Label to be set")
		asrt.Equal("red", e.Attrs().Color, "expected Color to be set")
		custom := e.Attrs().Custom()
		asrt.Equal("2.0", custom["penwidth"], "expected custom attribute to be set")
	})

	t.Run("custom attributes accumulate", func(t *testing.T) {
		asrt := assert.New(t)
		g := NewGraph()
		n1 := NewNode("A")
		n2 := NewNode("B")

		e := g.AddEdge(n1, n2,
			WithEdgeAttribute("attr1", "value1"),
			WithEdgeAttribute("attr2", "value2"),
			WithEdgeAttribute("attr3", "value3"),
		)

		custom := e.Attrs().Custom()
		asrt.Len(custom, 3, "expected all custom attributes to accumulate")
		asrt.Equal("value1", custom["attr1"], "expected attr1 to be set")
		asrt.Equal("value2", custom["attr2"], "expected attr2 to be set")
		asrt.Equal("value3", custom["attr3"], "expected attr3 to be set")
	})

	t.Run("later custom attribute overwrites earlier with same key", func(t *testing.T) {
		asrt := assert.New(t)
		g := NewGraph()
		n1 := NewNode("A")
		n2 := NewNode("B")

		e := g.AddEdge(n1, n2,
			WithEdgeAttribute("penwidth", "1.0"),
			WithEdgeAttribute("penwidth", "3.0"),
		)

		custom := e.Attrs().Custom()
		asrt.Len(custom, 1, "expected only one entry for duplicated key")
		asrt.Equal("3.0", custom["penwidth"], "expected later value to overwrite earlier one")
	})
}

// Test WithGraphAttribute option
func TestWithGraphAttribute_SetsCustom(t *testing.T) {
	t.Run("sets a single custom attribute on new graph", func(t *testing.T) {
		asrt := assert.New(t)

		g := NewGraph(WithGraphAttribute("ratio", "fill"))

		custom := g.Attrs().Custom()
		asrt.Len(custom, 1, "expected custom map to have 1 entry")
		asrt.Equal("fill", custom["ratio"], "expected ratio to be set to 'fill'")
	})

	t.Run("sets multiple custom attributes", func(t *testing.T) {
		asrt := assert.New(t)

		g := NewGraph(
			WithGraphAttribute("ratio", "fill"),
			WithGraphAttribute("concentrate", "true"),
			WithGraphAttribute("overlap", "false"),
		)

		custom := g.Attrs().Custom()
		asrt.Len(custom, 3, "expected custom map to have 3 entries")
		asrt.Equal("fill", custom["ratio"], "expected ratio to be set")
		asrt.Equal("true", custom["concentrate"], "expected concentrate to be set")
		asrt.Equal("false", custom["overlap"], "expected overlap to be set")
	})

	t.Run("combines with typed options", func(t *testing.T) {
		asrt := assert.New(t)

		g := NewGraph(
			WithGraphLabel("My Graph"),
			WithRankDir(RankDirLR),
			WithGraphAttribute("ratio", "fill"),
		)

		asrt.Equal("My Graph", g.Attrs().Label(), "expected Label to be set")
		asrt.Equal(RankDirLR, g.Attrs().RankDir(), "expected RankDir to be set")
		custom := g.Attrs().Custom()
		asrt.Equal("fill", custom["ratio"], "expected custom attribute to be set")
	})

	t.Run("custom attributes accumulate", func(t *testing.T) {
		asrt := assert.New(t)

		g := NewGraph(
			WithGraphAttribute("attr1", "value1"),
			WithGraphAttribute("attr2", "value2"),
			WithGraphAttribute("attr3", "value3"),
		)

		custom := g.Attrs().Custom()
		asrt.Len(custom, 3, "expected all custom attributes to accumulate")
		asrt.Equal("value1", custom["attr1"], "expected attr1 to be set")
		asrt.Equal("value2", custom["attr2"], "expected attr2 to be set")
		asrt.Equal("value3", custom["attr3"], "expected attr3 to be set")
	})

	t.Run("later custom attribute overwrites earlier with same key", func(t *testing.T) {
		asrt := assert.New(t)

		g := NewGraph(
			WithGraphAttribute("ratio", "fill"),
			WithGraphAttribute("ratio", "auto"),
		)

		custom := g.Attrs().Custom()
		asrt.Len(custom, 1, "expected only one entry for duplicated key")
		asrt.Equal("auto", custom["ratio"], "expected later value to overwrite earlier one")
	})

	t.Run("combines with graph structure options", func(t *testing.T) {
		asrt := assert.New(t)

		g := NewGraph(
			Directed,
			Strict,
			WithGraphAttribute("ratio", "fill"),
		)

		asrt.True(g.IsDirected(), "expected graph to be directed")
		asrt.True(g.IsStrict(), "expected graph to be strict")
		custom := g.Attrs().Custom()
		asrt.Equal("fill", custom["ratio"], "expected custom attribute to be set")
	})
}

// Test that custom attributes do not override typed attributes
func TestCustomAttributes_DoNotOverrideTyped(t *testing.T) {
	t.Run("custom 'shape' does not override typed Shape for nodes", func(t *testing.T) {
		asrt := assert.New(t)

		n := NewNode("test",
			WithShape(ShapeBox),
			WithNodeAttribute("shape", "circle"),
		)

		// The typed Shape should be preserved
		asrt.Equal(ShapeBox, n.Attrs().Shape, "expected typed Shape to be preserved")

		// The custom attribute should also be present (for DOT output purposes)
		custom := n.Attrs().Custom()
		asrt.Equal("circle", custom["shape"], "expected custom shape attribute to be stored")
	})

	t.Run("custom 'label' does not override typed Label for nodes", func(t *testing.T) {
		asrt := assert.New(t)

		n := NewNode("test",
			WithLabel("Typed Label"),
			WithNodeAttribute("label", "Custom Label"),
		)

		asrt.Equal("Typed Label", n.Attrs().Label, "expected typed Label to be preserved")
		custom := n.Attrs().Custom()
		asrt.Equal("Custom Label", custom["label"], "expected custom label attribute to be stored")
	})

	t.Run("custom 'color' does not override typed Color for nodes", func(t *testing.T) {
		asrt := assert.New(t)

		n := NewNode("test",
			WithColor("red"),
			WithNodeAttribute("color", "blue"),
		)

		asrt.Equal("red", n.Attrs().Color, "expected typed Color to be preserved")
		custom := n.Attrs().Custom()
		asrt.Equal("blue", custom["color"], "expected custom color attribute to be stored")
	})

	t.Run("custom 'label' does not override typed Label for edges", func(t *testing.T) {
		asrt := assert.New(t)
		g := NewGraph()
		n1 := NewNode("A")
		n2 := NewNode("B")

		e := g.AddEdge(n1, n2,
			WithEdgeLabel("Typed Label"),
			WithEdgeAttribute("label", "Custom Label"),
		)

		asrt.Equal("Typed Label", e.Attrs().Label, "expected typed Label to be preserved")
		custom := e.Attrs().Custom()
		asrt.Equal("Custom Label", custom["label"], "expected custom label attribute to be stored")
	})

	t.Run("custom 'style' does not override typed Style for edges", func(t *testing.T) {
		asrt := assert.New(t)
		g := NewGraph()
		n1 := NewNode("A")
		n2 := NewNode("B")

		e := g.AddEdge(n1, n2,
			WithEdgeStyle(EdgeStyleDashed),
			WithEdgeAttribute("style", "dotted"),
		)

		asrt.Equal(EdgeStyleDashed, e.Attrs().Style, "expected typed Style to be preserved")
		custom := e.Attrs().Custom()
		asrt.Equal("dotted", custom["style"], "expected custom style attribute to be stored")
	})

	t.Run("custom 'label' does not override typed Label for graphs", func(t *testing.T) {
		asrt := assert.New(t)

		g := NewGraph(
			WithGraphLabel("Typed Label"),
			WithGraphAttribute("label", "Custom Label"),
		)

		asrt.Equal("Typed Label", g.Attrs().Label(), "expected typed Label to be preserved")
		custom := g.Attrs().Custom()
		asrt.Equal("Custom Label", custom["label"], "expected custom label attribute to be stored")
	})

	t.Run("custom 'rankdir' does not override typed RankDir for graphs", func(t *testing.T) {
		asrt := assert.New(t)

		g := NewGraph(
			WithRankDir(RankDirLR),
			WithGraphAttribute("rankdir", "TB"),
		)

		asrt.Equal(RankDirLR, g.Attrs().RankDir(), "expected typed RankDir to be preserved")
		custom := g.Attrs().Custom()
		asrt.Equal("TB", custom["rankdir"], "expected custom rankdir attribute to be stored")
	})
}

// Test that multiple custom attribute calls accumulate
func TestCustomAttributes_MultipleCalls_Accumulate(t *testing.T) {
	t.Run("multiple WithNodeAttribute calls accumulate", func(t *testing.T) {
		asrt := assert.New(t)

		n := NewNode("test",
			WithNodeAttribute("peripheries", "2"),
			WithLabel("Node"),
			WithNodeAttribute("tooltip", "Hover text"),
			WithShape(ShapeBox),
			WithNodeAttribute("URL", "http://example.com"),
		)

		custom := n.Attrs().Custom()
		asrt.Len(custom, 3, "expected 3 custom attributes to accumulate")
		asrt.Equal("2", custom["peripheries"], "expected peripheries to be set")
		asrt.Equal("Hover text", custom["tooltip"], "expected tooltip to be set")
		asrt.Equal("http://example.com", custom["URL"], "expected URL to be set")

		// Verify typed attributes are also set
		asrt.Equal("Node", n.Attrs().Label, "expected typed Label to be set")
		asrt.Equal(ShapeBox, n.Attrs().Shape, "expected typed Shape to be set")
	})

	t.Run("multiple WithEdgeAttribute calls accumulate", func(t *testing.T) {
		asrt := assert.New(t)
		g := NewGraph()
		n1 := NewNode("A")
		n2 := NewNode("B")

		e := g.AddEdge(n1, n2,
			WithEdgeAttribute("penwidth", "2.0"),
			WithEdgeLabel("connects"),
			WithEdgeAttribute("dir", "both"),
			WithEdgeColor("red"),
			WithEdgeAttribute("constraint", "false"),
		)

		custom := e.Attrs().Custom()
		asrt.Len(custom, 3, "expected 3 custom attributes to accumulate")
		asrt.Equal("2.0", custom["penwidth"], "expected penwidth to be set")
		asrt.Equal("both", custom["dir"], "expected dir to be set")
		asrt.Equal("false", custom["constraint"], "expected constraint to be set")

		// Verify typed attributes are also set
		asrt.Equal("connects", e.Attrs().Label, "expected typed Label to be set")
		asrt.Equal("red", e.Attrs().Color, "expected typed Color to be set")
	})

	t.Run("multiple WithGraphAttribute calls accumulate", func(t *testing.T) {
		asrt := assert.New(t)

		g := NewGraph(
			WithGraphAttribute("ratio", "fill"),
			WithGraphLabel("My Graph"),
			WithGraphAttribute("concentrate", "true"),
			WithRankDir(RankDirLR),
			WithGraphAttribute("overlap", "false"),
		)

		custom := g.Attrs().Custom()
		asrt.Len(custom, 3, "expected 3 custom attributes to accumulate")
		asrt.Equal("fill", custom["ratio"], "expected ratio to be set")
		asrt.Equal("true", custom["concentrate"], "expected concentrate to be set")
		asrt.Equal("false", custom["overlap"], "expected overlap to be set")

		// Verify typed attributes are also set
		asrt.Equal("My Graph", g.Attrs().Label(), "expected typed Label to be set")
		asrt.Equal(RankDirLR, g.Attrs().RankDir(), "expected typed RankDir to be set")
	})

	t.Run("mixing custom and typed options in various orders", func(t *testing.T) {
		asrt := assert.New(t)

		n := NewNode("test",
			WithNodeAttribute("attr1", "value1"),
			WithShape(ShapeBox),
			WithNodeAttribute("attr2", "value2"),
			WithLabel("Label"),
			WithNodeAttribute("attr3", "value3"),
			WithColor("red"),
		)

		custom := n.Attrs().Custom()
		asrt.Len(custom, 3, "expected all custom attributes to accumulate regardless of order")
		asrt.Equal("value1", custom["attr1"], "expected attr1 to be set")
		asrt.Equal("value2", custom["attr2"], "expected attr2 to be set")
		asrt.Equal("value3", custom["attr3"], "expected attr3 to be set")

		// All typed attributes should also be set
		asrt.Equal(ShapeBox, n.Attrs().Shape, "expected typed Shape to be set")
		asrt.Equal("Label", n.Attrs().Label, "expected typed Label to be set")
		asrt.Equal("red", n.Attrs().Color, "expected typed Color to be set")
	})
}
