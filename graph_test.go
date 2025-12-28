package goraffe

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewGraph_DefaultValues(t *testing.T) {
	asrt := assert.New(t)

	g := NewGraph()

	asrt.False(g.directed, "expected directed to be false")
	asrt.False(g.strict, "expected strict to be false")
	asrt.Empty(g.name, "expected name to be empty")
}

func TestIsDirected(t *testing.T) {
	t.Run("returns false for undirected graph", func(t *testing.T) {
		asrt := assert.New(t)

		g := NewGraph()
		g.directed = false

		asrt.False(g.IsDirected(), "expected IsDirected to return false")
	})

	t.Run("returns true for directed graph", func(t *testing.T) {
		asrt := assert.New(t)

		g := NewGraph()
		g.directed = true

		asrt.True(g.IsDirected(), "expected IsDirected to return true")
	})
}

func TestIsStrict(t *testing.T) {
	t.Run("returns false for non-strict graph", func(t *testing.T) {
		asrt := assert.New(t)

		g := NewGraph()
		g.strict = false

		asrt.False(g.IsStrict(), "expected IsStrict to return false")
	})

	t.Run("returns true for strict graph", func(t *testing.T) {
		asrt := assert.New(t)

		g := NewGraph()
		g.strict = true

		asrt.True(g.IsStrict(), "expected IsStrict to return true")
	})
}

func TestName(t *testing.T) {
	t.Run("returns empty string for unnamed graph", func(t *testing.T) {
		asrt := assert.New(t)

		g := NewGraph()
		g.name = ""

		asrt.Empty(g.Name(), "expected Name to return empty string")
	})

	t.Run("returns name for named graph", func(t *testing.T) {
		asrt := assert.New(t)

		g := NewGraph()
		g.name = "TestGraph"

		asrt.Equal("TestGraph", g.Name(), "expected Name to return 'TestGraph'")
	})
}

func TestGraph_AddNode(t *testing.T) {
	t.Run("adds single node", func(t *testing.T) {
		asrt := assert.New(t)

		g := NewGraph()
		n := NewNode("A")

		g.AddNode(n)

		asrt.Equal(1, len(g.Nodes()), "expected graph to have 1 node")
		asrt.Equal(n, g.GetNode("A"), "expected GetNode to return the added node")
	})

	t.Run("adds multiple nodes", func(t *testing.T) {
		asrt := assert.New(t)

		g := NewGraph()
		n1 := NewNode("A")
		n2 := NewNode("B")
		n3 := NewNode("C")

		g.AddNode(n1)
		g.AddNode(n2)
		g.AddNode(n3)

		asrt.Equal(3, len(g.Nodes()), "expected graph to have 3 nodes")
		asrt.Equal(n1, g.GetNode("A"), "expected GetNode to return node A")
		asrt.Equal(n2, g.GetNode("B"), "expected GetNode to return node B")
		asrt.Equal(n3, g.GetNode("C"), "expected GetNode to return node C")
	})

	t.Run("replaces node with duplicate ID", func(t *testing.T) {
		asrt := assert.New(t)

		g := NewGraph()
		n1 := NewNode("A")
		n2 := NewNode("A") // same ID, different node

		g.AddNode(n1)
		g.AddNode(n2)

		asrt.Equal(1, len(g.Nodes()), "expected graph to have 1 node after adding duplicate")
		asrt.Same(n2, g.GetNode("A"), "expected GetNode to return the replacement node")
		asrt.NotSame(n1, g.GetNode("A"), "expected original node to be replaced")
	})

	t.Run("preserves insertion order", func(t *testing.T) {
		asrt := assert.New(t)

		g := NewGraph()
		n1 := NewNode("A")
		n2 := NewNode("B")
		n3 := NewNode("C")

		g.AddNode(n1)
		g.AddNode(n2)
		g.AddNode(n3)

		nodes := g.Nodes()
		asrt.Equal(n1, nodes[0], "expected first node to be A")
		asrt.Equal(n2, nodes[1], "expected second node to be B")
		asrt.Equal(n3, nodes[2], "expected third node to be C")
	})
}

func TestGraph_GetNode(t *testing.T) {
	t.Run("returns node when it exists", func(t *testing.T) {
		asrt := assert.New(t)

		g := NewGraph()
		n := NewNode("A")

		g.AddNode(n)

		retrieved := g.GetNode("A")
		asrt.NotNil(retrieved, "expected GetNode to return non-nil for existing node")
		asrt.Equal(n, retrieved, "expected GetNode to return the same node instance")
	})

	t.Run("returns nil when node not found", func(t *testing.T) {
		asrt := assert.New(t)

		g := NewGraph()

		retrieved := g.GetNode("NonExistent")
		asrt.Nil(retrieved, "expected GetNode to return nil for non-existent node")
	})
}

func TestGraph_Nodes(t *testing.T) {
	t.Run("returns all nodes", func(t *testing.T) {
		asrt := assert.New(t)

		g := NewGraph()
		n1 := NewNode("A")
		n2 := NewNode("B")
		n3 := NewNode("C")

		g.AddNode(n1)
		g.AddNode(n2)
		g.AddNode(n3)

		nodes := g.Nodes()
		asrt.Len(nodes, 3, "expected Nodes to return all 3 nodes")
		asrt.Contains(nodes, n1, "expected nodes to contain n1")
		asrt.Contains(nodes, n2, "expected nodes to contain n2")
		asrt.Contains(nodes, n3, "expected nodes to contain n3")
	})

	t.Run("returns nodes in insertion order", func(t *testing.T) {
		asrt := assert.New(t)

		g := NewGraph()
		n1 := NewNode("Z")
		n2 := NewNode("A")
		n3 := NewNode("M")

		// Add in a specific order (not alphabetical)
		g.AddNode(n1)
		g.AddNode(n2)
		g.AddNode(n3)

		nodes := g.Nodes()
		asrt.Equal(3, len(nodes), "expected 3 nodes")
		asrt.Equal("Z", nodes[0].ID(), "expected first node to be Z (insertion order)")
		asrt.Equal("A", nodes[1].ID(), "expected second node to be A (insertion order)")
		asrt.Equal("M", nodes[2].ID(), "expected third node to be M (insertion order)")
	})

	t.Run("preserves position when node replaced", func(t *testing.T) {
		asrt := assert.New(t)

		g := NewGraph()
		n1 := NewNode("Z")
		n2 := NewNode("A")
		n3 := NewNode("M")

		g.AddNode(n1)
		g.AddNode(n2)
		g.AddNode(n3)

		// Replace node at position 1
		n4 := NewNode("A")
		g.AddNode(n4)

		nodes := g.Nodes()
		asrt.Equal(3, len(nodes), "expected still 3 nodes after replace")
		asrt.Equal("Z", nodes[0].ID(), "expected first node still Z")
		asrt.Equal("A", nodes[1].ID(), "expected second node still A (same position)")
		asrt.Equal(n4, nodes[1], "expected second node to be replaced instance")
		asrt.Equal("M", nodes[2].ID(), "expected third node still M")
	})
}
