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

		_ = g.AddNode(n)

		asrt.Equal(1, len(g.Nodes()), "expected graph to have 1 node")
		asrt.Equal(n, g.GetNode("A"), "expected GetNode to return the added node")
	})

	t.Run("adds multiple nodes", func(t *testing.T) {
		asrt := assert.New(t)

		g := NewGraph()
		n1 := NewNode("A")
		n2 := NewNode("B")
		n3 := NewNode("C")

		_ = g.AddNode(n1)
		_ = g.AddNode(n2)
		_ = g.AddNode(n3)

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

		_ = g.AddNode(n1)
		_ = g.AddNode(n2)

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

		_ = g.AddNode(n1)
		_ = g.AddNode(n2)
		_ = g.AddNode(n3)

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

		_ = g.AddNode(n)

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

		_ = g.AddNode(n1)
		_ = g.AddNode(n2)
		_ = g.AddNode(n3)

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
		_ = g.AddNode(n1)
		_ = g.AddNode(n2)
		_ = g.AddNode(n3)

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

		_ = g.AddNode(n1)
		_ = g.AddNode(n2)
		_ = g.AddNode(n3)

		// Replace node at position 1
		n4 := NewNode("A")
		_ = g.AddNode(n4)

		nodes := g.Nodes()
		asrt.Equal(3, len(nodes), "expected still 3 nodes after replace")
		asrt.Equal("Z", nodes[0].ID(), "expected first node still Z")
		asrt.Equal("A", nodes[1].ID(), "expected second node still A (same position)")
		asrt.Equal(n4, nodes[1], "expected second node to be replaced instance")
		asrt.Equal("M", nodes[2].ID(), "expected third node still M")
	})
}

func TestGraph_AddNode_ValidatesNilNode(t *testing.T) {
	t.Run("returns error when node is nil", func(t *testing.T) {
		asrt := assert.New(t)

		g := NewGraph()

		err := g.AddNode(nil)

		asrt.Error(err, "expected error when adding nil node")
		asrt.ErrorIs(err, ErrNilNode, "expected ErrNilNode sentinel error")
		asrt.Equal(0, len(g.Nodes()), "expected no nodes to be added when nil is passed")
	})

	t.Run("does not modify graph when nil node is passed", func(t *testing.T) {
		asrt := assert.New(t)

		g := NewGraph()
		n1 := NewNode("A")
		_ = g.AddNode(n1)

		err := g.AddNode(nil)

		asrt.Error(err, "expected error when adding nil node")
		asrt.ErrorIs(err, ErrNilNode, "expected ErrNilNode sentinel error")
		asrt.Equal(1, len(g.Nodes()), "expected node count to remain unchanged")
		asrt.Same(n1, g.Nodes()[0], "expected existing node to remain unchanged")
	})

	t.Run("succeeds when node is valid", func(t *testing.T) {
		asrt := assert.New(t)

		g := NewGraph()
		n := NewNode("A")

		err := g.AddNode(n)

		asrt.NoError(err, "expected no error when adding valid node")
		asrt.Equal(1, len(g.Nodes()), "expected node to be added")
		asrt.Same(n, g.GetNode("A"), "expected node to be retrievable")
	})
}

func TestGraph_AddEdge_ValidatesNilNodes(t *testing.T) {
	t.Run("returns error when from node is nil", func(t *testing.T) {
		asrt := assert.New(t)

		g := NewGraph()
		to := NewNode("B")

		edge, err := g.AddEdge(nil, to)

		asrt.Error(err, "expected error when from node is nil")
		asrt.Nil(edge, "expected nil edge when error occurs")
		asrt.ErrorIs(err, ErrNilNode, "expected ErrNilNode sentinel error")
		asrt.Equal(0, len(g.Edges()), "expected no edges to be added")
	})

	t.Run("returns error when to node is nil", func(t *testing.T) {
		asrt := assert.New(t)

		g := NewGraph()
		from := NewNode("A")

		edge, err := g.AddEdge(from, nil)

		asrt.Error(err, "expected error when to node is nil")
		asrt.Nil(edge, "expected nil edge when error occurs")
		asrt.ErrorIs(err, ErrNilNode, "expected ErrNilNode sentinel error")
		asrt.Equal(0, len(g.Edges()), "expected no edges to be added")
	})

	t.Run("returns error when both nodes are nil", func(t *testing.T) {
		asrt := assert.New(t)

		g := NewGraph()

		edge, err := g.AddEdge(nil, nil)

		asrt.Error(err, "expected error when both nodes are nil")
		asrt.Nil(edge, "expected nil edge when error occurs")
		asrt.ErrorIs(err, ErrNilNode, "expected ErrNilNode sentinel error")
		asrt.Equal(0, len(g.Edges()), "expected no edges to be added")
	})

	t.Run("does not modify graph when nil node is passed", func(t *testing.T) {
		asrt := assert.New(t)

		g := NewGraph()
		n1 := NewNode("A")
		n2 := NewNode("B")
		_, _ = g.AddEdge(n1, n2)

		edge, err := g.AddEdge(nil, n2)

		asrt.Error(err, "expected error when from node is nil")
		asrt.Nil(edge, "expected nil edge when error occurs")
		asrt.ErrorIs(err, ErrNilNode, "expected ErrNilNode sentinel error")
		asrt.Equal(1, len(g.Edges()), "expected edge count to remain unchanged")
		asrt.Equal(2, len(g.Nodes()), "expected node count to remain unchanged")
	})

	t.Run("succeeds when both nodes are valid", func(t *testing.T) {
		asrt := assert.New(t)

		g := NewGraph()
		from := NewNode("A")
		to := NewNode("B")

		edge, err := g.AddEdge(from, to)

		asrt.NoError(err, "expected no error when adding valid edge")
		asrt.NotNil(edge, "expected edge to be created")
		asrt.Equal(1, len(g.Edges()), "expected edge to be added")
		asrt.Same(from, edge.From(), "expected edge to have correct from node")
		asrt.Same(to, edge.To(), "expected edge to have correct to node")
	})

	t.Run("succeeds with options when nodes are valid", func(t *testing.T) {
		asrt := assert.New(t)

		g := NewGraph()
		from := NewNode("A")
		to := NewNode("B")

		edge, err := g.AddEdge(from, to, WithEdgeLabel("test"))

		asrt.NoError(err, "expected no error when adding valid edge with options")
		asrt.NotNil(edge, "expected edge to be created")
		asrt.Equal("test", edge.Attrs().Label(), "expected edge options to be applied")
	})

	t.Run("does not apply options when nil node causes error", func(t *testing.T) {
		asrt := assert.New(t)

		g := NewGraph()
		to := NewNode("B")

		edge, err := g.AddEdge(nil, to, WithEdgeLabel("should not be set"))

		asrt.Error(err, "expected error when from node is nil")
		asrt.Nil(edge, "expected nil edge when error occurs")
		asrt.ErrorIs(err, ErrNilNode, "expected ErrNilNode sentinel error")
		asrt.Equal(0, len(g.Edges()), "expected no edges to be added")
	})

	t.Run("does not implicitly add nil nodes to graph", func(t *testing.T) {
		asrt := assert.New(t)

		g := NewGraph()
		n := NewNode("A")

		edge, err := g.AddEdge(n, nil)

		asrt.Error(err, "expected error when to node is nil")
		asrt.Nil(edge, "expected nil edge")
		asrt.ErrorIs(err, ErrNilNode, "expected ErrNilNode sentinel error")
		// The valid node should not be added to the graph if the edge creation fails
		asrt.Equal(0, len(g.Nodes()), "expected no nodes when edge creation fails due to nil")
	})
}

func TestGraph_AddEdge_NilValidation_WithImplicitNodeAddition(t *testing.T) {
	t.Run("does not add nodes when from is nil", func(t *testing.T) {
		asrt := assert.New(t)

		g := NewGraph()
		to := NewNode("B")

		edge, err := g.AddEdge(nil, to)

		asrt.Error(err, "expected error")
		asrt.Nil(edge, "expected nil edge")
		asrt.ErrorIs(err, ErrNilNode, "expected ErrNilNode sentinel error")
		asrt.Equal(0, len(g.Nodes()), "expected no nodes added when from is nil")
		asrt.Nil(g.GetNode("B"), "expected to node not added when from is nil")
	})

	t.Run("does not add nodes when to is nil", func(t *testing.T) {
		asrt := assert.New(t)

		g := NewGraph()
		from := NewNode("A")

		edge, err := g.AddEdge(from, nil)

		asrt.Error(err, "expected error")
		asrt.Nil(edge, "expected nil edge")
		asrt.ErrorIs(err, ErrNilNode, "expected ErrNilNode sentinel error")
		asrt.Equal(0, len(g.Nodes()), "expected no nodes added when to is nil")
		asrt.Nil(g.GetNode("A"), "expected from node not added when to is nil")
	})

	t.Run("adds nodes normally when both are valid", func(t *testing.T) {
		asrt := assert.New(t)

		g := NewGraph()
		from := NewNode("A")
		to := NewNode("B")

		edge, err := g.AddEdge(from, to)

		asrt.NoError(err, "expected no error")
		asrt.NotNil(edge, "expected edge to be created")
		asrt.Equal(2, len(g.Nodes()), "expected both nodes to be added")
		asrt.NotNil(g.GetNode("A"), "expected from node to be added")
		asrt.NotNil(g.GetNode("B"), "expected to node to be added")
	})
}

func TestGraph_SameRank(t *testing.T) {
	asrt := assert.New(t)

	g := NewGraph()
	n1 := NewNode("A")
	n2 := NewNode("B")
	n3 := NewNode("C")

	sg, err := g.SameRank(n1, n2, n3)

	asrt.NoError(err, "expected no error")
	asrt.NotNil(sg, "expected subgraph to be created")
	asrt.Equal(RankSame, sg.Rank(), "expected rank to be 'same'")
	asrt.Len(sg.Nodes(), 3, "expected 3 nodes in subgraph")

	// Verify all nodes were added
	nodeIDs := make(map[string]bool)
	for _, n := range sg.Nodes() {
		nodeIDs[n.ID()] = true
	}
	asrt.True(nodeIDs["A"], "expected node A in subgraph")
	asrt.True(nodeIDs["B"], "expected node B in subgraph")
	asrt.True(nodeIDs["C"], "expected node C in subgraph")
}

func TestGraph_MinRank(t *testing.T) {
	asrt := assert.New(t)

	g := NewGraph()
	n1 := NewNode("A")
	n2 := NewNode("B")

	sg, err := g.MinRank(n1, n2)

	asrt.NoError(err, "expected no error")
	asrt.NotNil(sg, "expected subgraph to be created")
	asrt.Equal(RankMin, sg.Rank(), "expected rank to be 'min'")
	asrt.Len(sg.Nodes(), 2, "expected 2 nodes in subgraph")
}

func TestGraph_MaxRank(t *testing.T) {
	asrt := assert.New(t)

	g := NewGraph()
	n1 := NewNode("A")
	n2 := NewNode("B")

	sg, err := g.MaxRank(n1, n2)

	asrt.NoError(err, "expected no error")
	asrt.NotNil(sg, "expected subgraph to be created")
	asrt.Equal(RankMax, sg.Rank(), "expected rank to be 'max'")
	asrt.Len(sg.Nodes(), 2, "expected 2 nodes in subgraph")
}

func TestGraph_SourceRank(t *testing.T) {
	asrt := assert.New(t)

	g := NewGraph()
	n1 := NewNode("A")

	sg, err := g.SourceRank(n1)

	asrt.NoError(err, "expected no error")
	asrt.NotNil(sg, "expected subgraph to be created")
	asrt.Equal(RankSource, sg.Rank(), "expected rank to be 'source'")
	asrt.Len(sg.Nodes(), 1, "expected 1 node in subgraph")
}

func TestGraph_SinkRank(t *testing.T) {
	asrt := assert.New(t)

	g := NewGraph()
	n1 := NewNode("A")
	n2 := NewNode("B")
	n3 := NewNode("C")

	sg, err := g.SinkRank(n1, n2, n3)

	asrt.NoError(err, "expected no error")
	asrt.NotNil(sg, "expected subgraph to be created")
	asrt.Equal(RankSink, sg.Rank(), "expected rank to be 'sink'")
	asrt.Len(sg.Nodes(), 3, "expected 3 nodes in subgraph")
}

func TestGraph_RankMethods_AnonymousSubgraphs(t *testing.T) {
	asrt := assert.New(t)

	g := NewGraph()
	n1 := NewNode("A")

	sg, err := g.SameRank(n1)

	asrt.NoError(err, "expected no error")
	asrt.Equal("", sg.Name(), "expected anonymous subgraph to have empty name")
	asrt.Equal(RankSame, sg.Rank(), "expected rank to be set")
}

func TestGraph_RankMethods_NodesAddedToGraph(t *testing.T) {
	asrt := assert.New(t)

	g := NewGraph()
	n1 := NewNode("A")
	n2 := NewNode("B")

	_, err := g.SameRank(n1, n2)

	asrt.NoError(err, "expected no error")
	// Verify nodes were added to parent graph
	asrt.Len(g.Nodes(), 2, "expected nodes to be added to parent graph")
	asrt.NotNil(g.GetNode("A"), "expected node A in parent graph")
	asrt.NotNil(g.GetNode("B"), "expected node B in parent graph")
}

func TestGraph_RankMethods_ErrorOnNilNode(t *testing.T) {
	t.Run("SameRank returns error with nil node", func(t *testing.T) {
		asrt := assert.New(t)
		g := NewGraph()

		sg, err := g.SameRank(NewNode("A"), nil, NewNode("B"))
		asrt.Error(err, "expected error when nil node passed")
		asrt.Nil(sg, "expected nil subgraph when error occurs")
		asrt.ErrorIs(err, ErrNilNode, "expected ErrNilNode sentinel error")
	})

	t.Run("SameRank returns error with only nil node", func(t *testing.T) {
		asrt := assert.New(t)
		g := NewGraph()

		sg, err := g.SameRank(nil)
		asrt.Error(err, "expected error when nil node passed")
		asrt.Nil(sg, "expected nil subgraph when error occurs")
		asrt.ErrorIs(err, ErrNilNode, "expected ErrNilNode sentinel error")
	})

	t.Run("MinRank returns error with nil node", func(t *testing.T) {
		asrt := assert.New(t)
		g := NewGraph()

		sg, err := g.MinRank(nil)
		asrt.Error(err, "expected error when nil node passed")
		asrt.Nil(sg, "expected nil subgraph when error occurs")
		asrt.ErrorIs(err, ErrNilNode, "expected ErrNilNode sentinel error")
	})

	t.Run("MaxRank returns error with nil node", func(t *testing.T) {
		asrt := assert.New(t)
		g := NewGraph()

		sg, err := g.MaxRank(NewNode("A"), nil)
		asrt.Error(err, "expected error when nil node passed")
		asrt.Nil(sg, "expected nil subgraph when error occurs")
		asrt.ErrorIs(err, ErrNilNode, "expected ErrNilNode sentinel error")
	})

	t.Run("SourceRank returns error with nil node", func(t *testing.T) {
		asrt := assert.New(t)
		g := NewGraph()

		sg, err := g.SourceRank(nil, NewNode("A"))
		asrt.Error(err, "expected error when nil node passed")
		asrt.Nil(sg, "expected nil subgraph when error occurs")
		asrt.ErrorIs(err, ErrNilNode, "expected ErrNilNode sentinel error")
	})

	t.Run("SinkRank returns error with nil node", func(t *testing.T) {
		asrt := assert.New(t)
		g := NewGraph()

		sg, err := g.SinkRank(nil, nil, NewNode("A"))
		asrt.Error(err, "expected error when nil node passed")
		asrt.Nil(sg, "expected nil subgraph when error occurs")
		asrt.ErrorIs(err, ErrNilNode, "expected ErrNilNode sentinel error")
	})

	t.Run("rank methods do not add nodes when error occurs", func(t *testing.T) {
		asrt := assert.New(t)
		g := NewGraph()
		n1 := NewNode("A")

		sg, err := g.SameRank(n1, nil)
		asrt.Error(err, "expected error when nil node passed")
		asrt.Nil(sg, "expected nil subgraph when error occurs")
		// The valid node should not be added when the operation fails
		asrt.Len(g.Nodes(), 0, "expected no nodes added when rank method fails")
	})
}
