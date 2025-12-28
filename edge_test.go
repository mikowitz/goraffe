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

		e := g.AddEdge(n1, n2)

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

		g.AddNode(n1)
		g.AddNode(n2)

		e := g.AddEdge(n1, n2)

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
		e := g.AddEdge(n1, n2)

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
		g.AddNode(n1)

		e := g.AddEdge(n1, n2)

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

		e1 := g.AddEdge(n1, n2)
		e2 := g.AddEdge(n1, n2) // Parallel edge

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

		e := g.AddEdge(n1, n1)

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

		e1 := g.AddEdge(n1, n2)
		e2 := g.AddEdge(n2, n3)
		e3 := g.AddEdge(n3, n1)

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

		e1 := g.AddEdge(n1, n2)
		e2 := g.AddEdge(n2, n3)
		e3 := g.AddEdge(n3, n1)

		edges := g.Edges()
		asrt.Equal(3, len(edges), "expected 3 edges")
		asrt.Same(e1, edges[0], "expected first edge to be e1")
		asrt.Same(e2, edges[1], "expected second edge to be e2")
		asrt.Same(e3, edges[2], "expected third edge to be e3")
	})
}
