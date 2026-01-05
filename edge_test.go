package goraffe

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEdge_FromTo(t *testing.T) {
	t.Run("returns correct nodes", func(t *testing.T) {
		asrt := assert.New(t)

		n1 := NewNode("A")
		n2 := NewNode("B")
		g := NewGraph()

		e, err := g.AddEdge(n1, n2)
		asrt.NoError(err)

		asrt.Same(n1, e.From(), "expected From to return first node")
		asrt.Same(n2, e.To(), "expected To to return second node")
	})
}

func TestGraph_AddEdge(t *testing.T) {
	t.Run("creates edge when both nodes exist", func(t *testing.T) {
		asrt := assert.New(t)

		g := NewGraph()
		n1 := NewNode("A")
		n2 := NewNode("B")

		_ = g.AddNode(n1)
		_ = g.AddNode(n2)

		e, err := g.AddEdge(n1, n2)
		asrt.NoError(err)

		asrt.NotNil(e, "expected AddEdge to return an edge")
		asrt.Same(n1, e.From(), "expected edge from to be n1")
		asrt.Same(n2, e.To(), "expected edge to to be n2")
		asrt.Len(g.Edges(), 1, "expected graph to have 1 edge")
	})

	t.Run("implicitly adds nodes that don't exist", func(t *testing.T) {
		asrt := assert.New(t)

		g := NewGraph()
		n1 := NewNode("A")
		n2 := NewNode("B")

		// Don't add nodes to graph first
		e, err := g.AddEdge(n1, n2)
		asrt.NoError(err)

		asrt.NotNil(e, "expected AddEdge to return an edge")
		asrt.Len(g.Nodes(), 2, "expected graph to have 2 nodes after implicit add")
		asrt.Same(n1, g.GetNode("A"), "expected node A to be in graph")
		asrt.Same(n2, g.GetNode("B"), "expected node B to be in graph")
		asrt.Len(g.Edges(), 1, "expected graph to have 1 edge")
	})

	t.Run("partially adds nodes when one exists", func(t *testing.T) {
		asrt := assert.New(t)

		g := NewGraph()
		n1 := NewNode("A")
		n2 := NewNode("B")

		// Only add one node
		_ = g.AddNode(n1)

		e, err := g.AddEdge(n1, n2)
		asrt.NoError(err)

		asrt.NotNil(e, "expected AddEdge to return an edge")
		asrt.Len(g.Nodes(), 2, "expected graph to have 2 nodes after partial implicit add")
		asrt.Same(n1, g.GetNode("A"), "expected node A to be in graph")
		asrt.Same(n2, g.GetNode("B"), "expected node B to be in graph (implicitly added)")
		asrt.Len(g.Edges(), 1, "expected graph to have 1 edge")
	})

	t.Run("allows parallel edges between same nodes", func(t *testing.T) {
		asrt := assert.New(t)

		g := NewGraph()
		n1 := NewNode("A")
		n2 := NewNode("B")

		e1, err := g.AddEdge(n1, n2)
		asrt.NoError(err)
		e2, err := g.AddEdge(n1, n2) // Parallel edge
		asrt.NoError(err)

		asrt.Len(g.Edges(), 2, "expected graph to have 2 edges (parallel edges allowed)")
		asrt.NotSame(e1, e2, "expected parallel edges to be different instances")
		asrt.Same(n1, e1.From(), "expected both edges to have same from node")
		asrt.Same(n1, e2.From(), "expected both edges to have same from node")
		asrt.Same(n2, e1.To(), "expected both edges to have same to node")
		asrt.Same(n2, e2.To(), "expected both edges to have same to node")
	})

	t.Run("allows self-loop edges", func(t *testing.T) {
		asrt := assert.New(t)

		g := NewGraph()
		n1 := NewNode("A")

		e, err := g.AddEdge(n1, n1)
		asrt.NoError(err)

		asrt.NotNil(e, "expected AddEdge to create self-loop")
		asrt.Same(n1, e.From(), "expected from to be the same node")
		asrt.Same(n1, e.To(), "expected to to be the same node")
		asrt.Len(g.Edges(), 1, "expected graph to have 1 edge")
	})
}

func TestGraph_Edges(t *testing.T) {
	t.Run("returns all edges", func(t *testing.T) {
		asrt := assert.New(t)

		g := NewGraph()
		n1 := NewNode("A")
		n2 := NewNode("B")
		n3 := NewNode("C")

		e1, err := g.AddEdge(n1, n2)
		asrt.NoError(err)
		e2, err := g.AddEdge(n2, n3)
		asrt.NoError(err)
		e3, err := g.AddEdge(n3, n1)
		asrt.NoError(err)

		edges := g.Edges()
		asrt.Len(edges, 3, "expected graph to have 3 edges")
		asrt.Contains(edges, e1, "expected edges to contain e1")
		asrt.Contains(edges, e2, "expected edges to contain e2")
		asrt.Contains(edges, e3, "expected edges to contain e3")
	})

	t.Run("returns edges in insertion order", func(t *testing.T) {
		asrt := assert.New(t)

		g := NewGraph()
		n1 := NewNode("A")
		n2 := NewNode("B")
		n3 := NewNode("C")

		e1, err := g.AddEdge(n1, n2)
		asrt.NoError(err)
		e2, err := g.AddEdge(n2, n3)
		asrt.NoError(err)
		e3, err := g.AddEdge(n3, n1)
		asrt.NoError(err)

		edges := g.Edges()
		asrt.Equal(3, len(edges), "expected 3 edges")
		asrt.Same(e1, edges[0], "expected first edge to be e1")
		asrt.Same(e2, edges[1], "expected second edge to be e2")
		asrt.Same(e3, edges[2], "expected third edge to be e3")
	})
}

func TestEdge_Attrs_ReturnsAttributes(t *testing.T) {
	t.Run("returns non-nil attributes", func(t *testing.T) {
		asrt := assert.New(t)

		g := NewGraph()
		n1 := NewNode("A")
		n2 := NewNode("B")
		e, err := g.AddEdge(n1, n2)
		asrt.NoError(err)

		attrs := e.Attrs()
		asrt.NotNil(attrs, "expected Attrs() to return non-nil EdgeAttributes")
	})

	t.Run("returns same instance on multiple calls", func(t *testing.T) {
		asrt := assert.New(t)

		g := NewGraph()
		n1 := NewNode("A")
		n2 := NewNode("B")
		e, err := g.AddEdge(n1, n2)
		asrt.NoError(err)

		attrs1 := e.Attrs()
		attrs2 := e.Attrs()

		asrt.Same(attrs1, attrs2, "expected Attrs() to return the same EdgeAttributes instance")
	})

	t.Run("returns attributes with zero values for new edge", func(t *testing.T) {
		asrt := assert.New(t)

		g := NewGraph()
		n1 := NewNode("A")
		n2 := NewNode("B")
		e, err := g.AddEdge(n1, n2)
		asrt.NoError(err)

		attrs := e.Attrs()
		asrt.Empty(attrs.Label(), "expected Label to be empty for new edge")
		asrt.Empty(attrs.Color(), "expected Color to be empty for new edge")
		asrt.Empty(attrs.Style(), "expected Style to be empty for new edge")
		asrt.Empty(attrs.ArrowHead(), "expected ArrowHead to be empty for new edge")
		asrt.Empty(attrs.ArrowTail(), "expected ArrowTail to be empty for new edge")
		asrt.Equal(0.0, attrs.Weight(), "expected Weight to be zero for new edge")
	})
}

func TestGraph_AddEdge_WithOptions(t *testing.T) {
	t.Run("applies single EdgeLabel option", func(t *testing.T) {
		asrt := assert.New(t)

		g := NewGraph()
		n1 := NewNode("A")
		n2 := NewNode("B")

		e, err := g.AddEdge(n1, n2, WithEdgeLabel("connection"))
		asrt.NoError(err)

		asrt.Equal("connection", e.Attrs().Label(), "expected Label to be set")
	})

	t.Run("applies single EdgeColor option", func(t *testing.T) {
		asrt := assert.New(t)

		g := NewGraph()
		n1 := NewNode("A")
		n2 := NewNode("B")

		e, err := g.AddEdge(n1, n2, WithEdgeColor("red"))
		asrt.NoError(err)

		asrt.Equal("red", e.Attrs().Color(), "expected Color to be set")
	})

	t.Run("applies single EdgeStyle option", func(t *testing.T) {
		asrt := assert.New(t)

		g := NewGraph()
		n1 := NewNode("A")
		n2 := NewNode("B")

		e, err := g.AddEdge(n1, n2, WithEdgeStyle(EdgeStyleDashed))
		asrt.NoError(err)

		asrt.Equal(EdgeStyleDashed, e.Attrs().Style(), "expected Style to be set")
	})

	t.Run("applies single ArrowHead option", func(t *testing.T) {
		asrt := assert.New(t)

		g := NewGraph()
		n1 := NewNode("A")
		n2 := NewNode("B")

		e, err := g.AddEdge(n1, n2, WithArrowHead(ArrowNormal))
		asrt.NoError(err)

		asrt.Equal(ArrowNormal, e.Attrs().ArrowHead(), "expected ArrowHead to be set")
	})

	t.Run("applies single ArrowTail option", func(t *testing.T) {
		asrt := assert.New(t)

		g := NewGraph()
		n1 := NewNode("A")
		n2 := NewNode("B")

		e, err := g.AddEdge(n1, n2, WithArrowTail(ArrowDot))
		asrt.NoError(err)

		asrt.Equal(ArrowDot, e.Attrs().ArrowTail(), "expected ArrowTail to be set")
	})

	t.Run("applies single Weight option", func(t *testing.T) {
		asrt := assert.New(t)

		g := NewGraph()
		n1 := NewNode("A")
		n2 := NewNode("B")

		e, err := g.AddEdge(n1, n2, WithWeight(2.5))
		asrt.NoError(err)

		asrt.Equal(2.5, e.Attrs().Weight(), "expected Weight to be set")
	})
}

func TestGraph_AddEdge_WithMultipleOptions(t *testing.T) {
	t.Run("applies multiple options to same edge", func(t *testing.T) {
		asrt := assert.New(t)

		g := NewGraph()
		n1 := NewNode("A")
		n2 := NewNode("B")

		e, err := g.AddEdge(n1, n2,
			WithEdgeLabel("important"),
			WithEdgeColor("blue"),
			WithEdgeStyle(EdgeStyleBold),
			WithArrowHead(ArrowVee),
			WithArrowTail(ArrowNone),
			WithWeight(3.0),
		)
		asrt.NoError(err)

		attrs := e.Attrs()
		asrt.Equal("important", attrs.Label(), "expected Label to be set")
		asrt.Equal("blue", attrs.Color(), "expected Color to be set")
		asrt.Equal(EdgeStyleBold, attrs.Style(), "expected Style to be set")
		asrt.Equal(ArrowVee, attrs.ArrowHead(), "expected ArrowHead to be set")
		asrt.Equal(ArrowNone, attrs.ArrowTail(), "expected ArrowTail to be set")
		asrt.Equal(3.0, attrs.Weight(), "expected Weight to be set")
	})

	t.Run("does not affect other edges", func(t *testing.T) {
		asrt := assert.New(t)

		g := NewGraph()
		n1 := NewNode("A")
		n2 := NewNode("B")
		n3 := NewNode("C")

		e1, err := g.AddEdge(n1, n2, WithEdgeLabel("first"), WithEdgeColor("red"))
		asrt.NoError(err)
		e2, err := g.AddEdge(n2, n3, WithEdgeLabel("second"), WithEdgeColor("green"))
		asrt.NoError(err)

		asrt.Equal("first", e1.Attrs().Label(), "expected e1 to have its own label")
		asrt.Equal("red", e1.Attrs().Color(), "expected e1 to have its own color")
		asrt.Equal("second", e2.Attrs().Label(), "expected e2 to have its own label")
		asrt.Equal("green", e2.Attrs().Color(), "expected e2 to have its own color")
	})
}

func TestGraph_AddEdge_WithEdgeAttributesStruct(t *testing.T) {
	t.Run("uses EdgeAttributes as reusable template", func(t *testing.T) {
		asrt := assert.New(t)

		g := NewGraph()

		style := EdgeStyleDashed
		color := "gray"
		arrowHead := ArrowNormal
		weight := 1.5
		template := EdgeAttributes{
			style:     &style,
			color:     &color,
			arrowHead: &arrowHead,
			weight:    &weight,
		}

		n1 := NewNode("A")
		n2 := NewNode("B")

		e, err := g.AddEdge(n1, n2, template)
		asrt.NoError(err)

		attrs := e.Attrs()
		asrt.Equal(EdgeStyleDashed, attrs.Style(), "expected Style from template")
		asrt.Equal("gray", attrs.Color(), "expected Color from template")
		asrt.Equal(ArrowNormal, attrs.ArrowHead(), "expected ArrowHead from template")
		asrt.Equal(1.5, attrs.Weight(), "expected Weight from template")
	})

	t.Run("can combine EdgeAttributes template with additional options", func(t *testing.T) {
		asrt := assert.New(t)

		g := NewGraph()

		style := EdgeStyleDashed
		color := "gray"
		template := EdgeAttributes{
			style: &style,
			color: &color,
		}

		n1 := NewNode("A")
		n2 := NewNode("B")

		e, err := g.AddEdge(n1, n2, template, WithEdgeLabel("Custom"), WithWeight(2.0))
		asrt.NoError(err)

		attrs := e.Attrs()
		asrt.Equal(EdgeStyleDashed, attrs.Style(), "expected Style from template")
		asrt.Equal("gray", attrs.Color(), "expected Color from template")
		asrt.Equal("Custom", attrs.Label(), "expected Label from option")
		asrt.Equal(2.0, attrs.Weight(), "expected Weight from option")
	})

	t.Run("only copies non-nil fields from template", func(t *testing.T) {
		asrt := assert.New(t)

		g := NewGraph()

		// Template with only some fields set
		style := EdgeStyleBold
		color := "blue"
		template := EdgeAttributes{
			style: &style,
			color: &color,
			// label, arrowHead, arrowTail, weight are nil
		}

		n1 := NewNode("A")
		n2 := NewNode("B")

		// First set some attributes directly
		_, err := g.AddEdge(n1, n2, WithEdgeLabel("Important"), WithWeight(3.0))
		asrt.NoError(err)

		// Now apply template - should NOT overwrite existing non-nil values
		// (This test simulates the behavior, but in practice the template is applied during AddEdge)
		// Let me correct this to match the actual API usage

		// Actually, we need to create a new edge to test this properly
		n3 := NewNode("C")
		e2, err := g.AddEdge(n2, n3, WithEdgeLabel("PreSet"), WithWeight(5.0), template)
		asrt.NoError(err)

		attrs := e2.Attrs()
		asrt.Equal(EdgeStyleBold, attrs.Style(), "expected Style to be set from template")
		asrt.Equal("blue", attrs.Color(), "expected Color to be set from template")
		// Since template is applied first, then options, the options will override
		asrt.Equal("PreSet", attrs.Label(), "expected Label from option to override template's nil value")
		asrt.Equal(5.0, attrs.Weight(), "expected Weight from option to override template's nil value")
	})
}

func TestGraph_AddEdge_OptionsAppliedInOrder(t *testing.T) {
	t.Run("later options override earlier ones for same field", func(t *testing.T) {
		asrt := assert.New(t)

		g := NewGraph()
		n1 := NewNode("A")
		n2 := NewNode("B")

		e, err := g.AddEdge(n1, n2,
			WithEdgeLabel("first"),
			WithEdgeLabel("second"),
			WithEdgeLabel("third"),
		)
		asrt.NoError(err)

		asrt.Equal("third", e.Attrs().Label(), "expected last Label option to win")
	})

	t.Run("later options override earlier ones for different fields independently", func(t *testing.T) {
		asrt := assert.New(t)

		g := NewGraph()
		n1 := NewNode("A")
		n2 := NewNode("B")

		e, err := g.AddEdge(n1, n2,
			WithEdgeColor("red"),
			WithEdgeStyle(EdgeStyleSolid),
			WithEdgeColor("blue"),          // Override color
			WithWeight(1.0),                // New field
			WithEdgeStyle(EdgeStyleDashed), // Override style
		)
		asrt.NoError(err)

		attrs := e.Attrs()
		asrt.Equal("blue", attrs.Color(), "expected last Color option to win")
		asrt.Equal(EdgeStyleDashed, attrs.Style(), "expected last Style option to win")
		asrt.Equal(1.0, attrs.Weight(), "expected Weight to be set")
	})

	t.Run("options override template when template is first", func(t *testing.T) {
		asrt := assert.New(t)

		g := NewGraph()

		style := EdgeStyleDashed
		label := "Template Label"
		color := "gray"
		template := EdgeAttributes{
			style: &style,
			label: &label,
			color: &color,
		}

		n1 := NewNode("A")
		n2 := NewNode("B")

		e, err := g.AddEdge(n1, n2,
			template,
			WithEdgeLabel("Override"),
			WithEdgeColor("red"),
		)
		asrt.NoError(err)

		attrs := e.Attrs()
		asrt.Equal(EdgeStyleDashed, attrs.Style(), "expected Style from template (not overridden)")
		asrt.Equal("Override", attrs.Label(), "expected Label to be overridden by option")
		asrt.Equal("red", attrs.Color(), "expected Color to be overridden by option")
	})

	t.Run("template overrides options when template is last", func(t *testing.T) {
		asrt := assert.New(t)

		g := NewGraph()

		style := EdgeStyleBold
		label := "Template Label"
		template := EdgeAttributes{
			style: &style,
			label: &label,
		}

		n1 := NewNode("A")
		n2 := NewNode("B")

		e, err := g.AddEdge(n1, n2,
			WithEdgeLabel("First"),
			WithEdgeColor("blue"),
			template, // Applied last, will override Label but not Color
		)
		asrt.NoError(err)

		attrs := e.Attrs()
		asrt.Equal("Template Label", attrs.Label(), "expected Label from template (applied last)")
		asrt.Equal(EdgeStyleBold, attrs.Style(), "expected Style from template")
		asrt.Equal("blue", attrs.Color(), "expected Color from earlier option (template has nil value)")
	})
}
