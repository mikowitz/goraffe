package goraffe

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParse_EmptyDigraph(t *testing.T) {
	asrt := assert.New(t)

	input := "digraph {}"
	parser := newParser(input)
	g, err := parser.parseGraph()

	asrt.NoError(err, "Should parse empty digraph without error")
	asrt.NotNil(g, "Graph should not be nil")
	asrt.True(g.IsDirected(), "Graph should be directed")
	asrt.False(g.IsStrict(), "Graph should not be strict")
	asrt.Equal("", g.Name(), "Graph name should be empty")
}

func TestParse_EmptyGraph(t *testing.T) {
	asrt := assert.New(t)

	input := "graph {}"
	parser := newParser(input)
	g, err := parser.parseGraph()

	asrt.NoError(err, "Should parse empty graph without error")
	asrt.NotNil(g, "Graph should not be nil")
	asrt.False(g.IsDirected(), "Graph should be undirected")
	asrt.False(g.IsStrict(), "Graph should not be strict")
	asrt.Equal("", g.Name(), "Graph name should be empty")
}

func TestParse_StrictGraph(t *testing.T) {
	asrt := assert.New(t)

	tests := []struct {
		name     string
		input    string
		directed bool
	}{
		{"strict digraph", "strict digraph {}", true},
		{"strict graph", "strict graph {}", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			parser := newParser(tt.input)
			g, err := parser.parseGraph()

			asrt.NoError(err, "Should parse strict graph without error")
			asrt.NotNil(g, "Graph should not be nil")
			asrt.Equal(tt.directed, g.IsDirected(), "Graph directedness should match")
			asrt.True(g.IsStrict(), "Graph should be strict")
		})
	}
}

func TestParse_NamedGraph(t *testing.T) {
	asrt := assert.New(t)

	tests := []struct {
		name         string
		input        string
		expectedName string
	}{
		{"identifier name", "digraph G {}", "G"},
		{"multi-char name", "digraph MyGraph {}", "MyGraph"},
		{"quoted name", `digraph "My Graph" {}`, "My Graph"},
		{"undirected with name", "graph TestGraph {}", "TestGraph"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			parser := newParser(tt.input)
			g, err := parser.parseGraph()

			asrt.NoError(err, "Should parse named graph without error")
			asrt.NotNil(g, "Graph should not be nil")
			asrt.Equal(tt.expectedName, g.Name(), "Graph name should match")
		})
	}
}

func TestParse_InvalidSyntax_Error(t *testing.T) {
	asrt := assert.New(t)

	tests := []struct {
		name  string
		input string
	}{
		{"missing keyword", "{}"},
		{"invalid keyword", "graphviz {}"},
		{"missing brace", "digraph"},
		{"missing closing brace", "digraph {"},
		{"wrong closing token", "digraph }"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			parser := newParser(tt.input)
			g, err := parser.parseGraph()

			asrt.Error(err, "Should return error for invalid syntax")
			if g != nil {
				// In some error cases, graph might be partially constructed
				t.Logf("Partial graph constructed: %v", g)
			}
		})
	}
}

func TestParse_WithWhitespaceAndComments(t *testing.T) {
	asrt := assert.New(t)

	input := `
		// This is a comment
		digraph G {
			/* Multi-line
			   comment */
		}
	`

	parser := newParser(input)
	g, err := parser.parseGraph()

	asrt.NoError(err, "Should parse graph with comments without error")
	asrt.NotNil(g, "Graph should not be nil")
	asrt.True(g.IsDirected(), "Graph should be directed")
	asrt.Equal("G", g.Name(), "Graph name should be G")
}

func TestParse_StrictNamedGraph(t *testing.T) {
	asrt := assert.New(t)

	input := "strict digraph MyGraph {}"
	parser := newParser(input)
	g, err := parser.parseGraph()

	asrt.NoError(err, "Should parse strict named graph without error")
	asrt.NotNil(g, "Graph should not be nil")
	asrt.True(g.IsDirected(), "Graph should be directed")
	asrt.True(g.IsStrict(), "Graph should be strict")
	asrt.Equal("MyGraph", g.Name(), "Graph name should be MyGraph")
}

func TestParse_SkipsStatements(t *testing.T) {
	asrt := assert.New(t)

	// Parser should skip statements it doesn't fully understand yet
	input := `digraph {
		A;
		B;
		A -> B;
		node [shape=box];
		edge [color=red];
		graph [rankdir=LR];
	}`

	parser := newParser(input)
	g, err := parser.parseGraph()

	asrt.NoError(err, "Should parse graph with statements without error")
	asrt.NotNil(g, "Graph should not be nil")
	asrt.True(g.IsDirected(), "Graph should be directed")
}

func TestParse_SkipsSubgraphs(t *testing.T) {
	asrt := assert.New(t)

	tests := []struct {
		name  string
		input string
	}{
		{"named subgraph", `digraph { subgraph cluster_0 { A; B; } }`},
		{"anonymous subgraph", `digraph { { A; B; } }`},
		{"nested subgraphs", `digraph { subgraph { subgraph { A; } } }`},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			parser := newParser(tt.input)
			g, err := parser.parseGraph()

			asrt.NoError(err, "Should parse graph with subgraphs without error")
			asrt.NotNil(g, "Graph should not be nil")
		})
	}
}

func TestParse_MultipleStatements(t *testing.T) {
	asrt := assert.New(t)

	input := `digraph {
		A;
		B;
		C;
		D -> E;
		node [shape=box];
	}`

	parser := newParser(input)
	g, err := parser.parseGraph()

	asrt.NoError(err, "Should parse graph with multiple statements without error")
	asrt.NotNil(g, "Graph should not be nil")
	asrt.True(g.IsDirected(), "Graph should be directed")
}

func TestParse_SemicolonHandling(t *testing.T) {
	asrt := assert.New(t)

	tests := []struct {
		name  string
		input string
	}{
		{"with semicolons", "digraph { A; B; }"},
		{"without semicolons", "digraph { A B }"},
		{"mixed", "digraph { A; B C; D }"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			parser := newParser(tt.input)
			g, err := parser.parseGraph()

			asrt.NoError(err, "Should parse graph regardless of semicolon usage")
			asrt.NotNil(g, "Graph should not be nil")
		})
	}
}

func TestParse_SingleNode(t *testing.T) {
	asrt := assert.New(t)

	input := "digraph { A; }"
	parser := newParser(input)
	g, err := parser.parseGraph()

	asrt.NoError(err, "Should parse single node without error")
	asrt.NotNil(g, "Graph should not be nil")
	asrt.NotNil(g.GetNode("A"), "Node A should exist")
	asrt.Equal(1, len(g.Nodes()), "Graph should have 1 node")
}

func TestParse_NodeWithAttributes(t *testing.T) {
	asrt := assert.New(t)

	input := `digraph { A [shape=box, label="Node A", color=red]; }`
	parser := newParser(input)
	g, err := parser.parseGraph()

	asrt.NoError(err, "Should parse node with attributes without error")
	asrt.NotNil(g, "Graph should not be nil")

	node := g.GetNode("A")
	asrt.NotNil(node, "Node A should exist")
	asrt.Equal("box", string(node.Attrs().Shape()), "Node should have shape=box")
	asrt.Equal("Node A", node.Attrs().Label(), "Node should have label")
	asrt.Equal("red", node.Attrs().Color(), "Node should have color")
}

func TestParse_SingleEdge(t *testing.T) {
	asrt := assert.New(t)

	input := "digraph { A -> B; }"
	parser := newParser(input)
	g, err := parser.parseGraph()

	asrt.NoError(err, "Should parse single edge without error")
	asrt.NotNil(g, "Graph should not be nil")
	asrt.Equal(2, len(g.Nodes()), "Graph should have 2 nodes")
	asrt.Equal(1, len(g.Edges()), "Graph should have 1 edge")

	edge := g.Edges()[0]
	asrt.Equal("A", edge.From().ID(), "Edge should be from A")
	asrt.Equal("B", edge.To().ID(), "Edge should be to B")
}

func TestParse_EdgeWithAttributes(t *testing.T) {
	asrt := assert.New(t)

	input := `digraph { A -> B [label="edge label", color=blue]; }`
	parser := newParser(input)
	g, err := parser.parseGraph()

	asrt.NoError(err, "Should parse edge with attributes without error")
	asrt.NotNil(g, "Graph should not be nil")
	asrt.Equal(1, len(g.Edges()), "Graph should have 1 edge")

	edge := g.Edges()[0]
	asrt.Equal("edge label", edge.Attrs().Label(), "Edge should have label")
	asrt.Equal("blue", edge.Attrs().Color(), "Edge should have color")
}

func TestParse_EdgeChain(t *testing.T) {
	asrt := assert.New(t)

	input := "digraph { A -> B -> C; }"
	parser := newParser(input)
	g, err := parser.parseGraph()

	asrt.NoError(err, "Should parse edge chain without error")
	asrt.NotNil(g, "Graph should not be nil")
	asrt.Equal(3, len(g.Nodes()), "Graph should have 3 nodes")
	asrt.Equal(2, len(g.Edges()), "Graph should have 2 edges (A->B and B->C)")

	edges := g.Edges()
	asrt.Equal("A", edges[0].From().ID(), "First edge should be from A")
	asrt.Equal("B", edges[0].To().ID(), "First edge should be to B")
	asrt.Equal("B", edges[1].From().ID(), "Second edge should be from B")
	asrt.Equal("C", edges[1].To().ID(), "Second edge should be to C")
}

func TestParse_MixedNodesAndEdges(t *testing.T) {
	asrt := assert.New(t)

	input := `digraph {
		A [shape=box];
		B [label="Node B"];
		A -> B [label="edge"];
		C;
		B -> C;
	}`

	parser := newParser(input)
	g, err := parser.parseGraph()

	asrt.NoError(err, "Should parse mixed nodes and edges without error")
	asrt.NotNil(g, "Graph should not be nil")
	asrt.Equal(3, len(g.Nodes()), "Graph should have 3 nodes")
	asrt.Equal(2, len(g.Edges()), "Graph should have 2 edges")

	nodeA := g.GetNode("A")
	asrt.NotNil(nodeA, "Node A should exist")
	asrt.Equal("box", string(nodeA.Attrs().Shape()), "Node A should have shape=box")

	nodeB := g.GetNode("B")
	asrt.NotNil(nodeB, "Node B should exist")
	asrt.Equal("Node B", nodeB.Attrs().Label(), "Node B should have label")
}

func TestParse_DefaultNodeAttrs(t *testing.T) {
	asrt := assert.New(t)

	input := `digraph {
		node [shape=circle, color=red];
		A;
		B;
	}`

	parser := newParser(input)
	g, err := parser.parseGraph()

	asrt.NoError(err, "Should parse default node attributes without error")
	asrt.NotNil(g, "Graph should not be nil")

	defaultAttrs := g.DefaultNodeAttrs()
	asrt.Equal("circle", string(defaultAttrs.Shape()), "Default shape should be circle")
	asrt.Equal("red", defaultAttrs.Color(), "Default color should be red")
}

func TestParse_DefaultEdgeAttrs(t *testing.T) {
	asrt := assert.New(t)

	input := `digraph {
		edge [color=blue, style=dashed];
		A -> B;
	}`

	parser := newParser(input)
	g, err := parser.parseGraph()

	asrt.NoError(err, "Should parse default edge attributes without error")
	asrt.NotNil(g, "Graph should not be nil")

	defaultAttrs := g.DefaultEdgeAttrs()
	asrt.Equal("blue", defaultAttrs.Color(), "Default edge color should be blue")
	asrt.Equal("dashed", string(defaultAttrs.Style()), "Default edge style should be dashed")
}

func TestParse_QuotedStrings(t *testing.T) {
	asrt := assert.New(t)

	input := `digraph {
		"Node 1" -> "Node 2" [label="my edge"];
	}`

	parser := newParser(input)
	g, err := parser.parseGraph()

	asrt.NoError(err, "Should parse quoted node IDs without error")
	asrt.NotNil(g, "Graph should not be nil")
	asrt.NotNil(g.GetNode("Node 1"), "Node 'Node 1' should exist")
	asrt.NotNil(g.GetNode("Node 2"), "Node 'Node 2' should exist")
}

func TestParse_Numbers(t *testing.T) {
	asrt := assert.New(t)

	input := "digraph { 1 -> 2 -> 3; }"
	parser := newParser(input)
	g, err := parser.parseGraph()

	asrt.NoError(err, "Should parse numeric node IDs without error")
	asrt.NotNil(g, "Graph should not be nil")
	asrt.Equal(3, len(g.Nodes()), "Graph should have 3 nodes")
	asrt.NotNil(g.GetNode("1"), "Node 1 should exist")
	asrt.NotNil(g.GetNode("2"), "Node 2 should exist")
	asrt.NotNil(g.GetNode("3"), "Node 3 should exist")
}

func TestParse_CompleteExample(t *testing.T) {
	asrt := assert.New(t)

	input := `digraph G {
		// Set default attributes
		node [shape=box];
		edge [color=red];

		// Define nodes
		A [label="Start"];
		B [label="Middle", color=blue];
		C [label="End"];

		// Define edges
		A -> B [label="first"];
		B -> C [label="second"];
		A -> C [style=dashed];
	}`

	parser := newParser(input)
	g, err := parser.parseGraph()

	asrt.NoError(err, "Should parse complete example without error")
	asrt.NotNil(g, "Graph should not be nil")
	asrt.Equal("G", g.Name(), "Graph name should be G")
	asrt.Equal(3, len(g.Nodes()), "Graph should have 3 nodes")
	asrt.Equal(3, len(g.Edges()), "Graph should have 3 edges")

	// Check default attributes were applied
	asrt.Equal("box", string(g.DefaultNodeAttrs().Shape()), "Default node shape should be box")
	asrt.Equal("red", g.DefaultEdgeAttrs().Color(), "Default edge color should be red")
}
