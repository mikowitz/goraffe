package goraffe

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseString_SimpleGraph(t *testing.T) {
	asrt := assert.New(t)

	input := "digraph G { A -> B; }"
	g, err := ParseString(input)

	asrt.NoError(err, "Should parse simple graph without error")
	asrt.NotNil(g, "Graph should not be nil")
	asrt.Equal("G", g.Name(), "Graph name should be G")
	asrt.True(g.IsDirected(), "Graph should be directed")
	asrt.Equal(2, len(g.Nodes()), "Graph should have 2 nodes")
	asrt.Equal(1, len(g.Edges()), "Graph should have 1 edge")
}

func TestParseString_ComplexGraph(t *testing.T) {
	asrt := assert.New(t)

	input := `digraph G {
		node [shape=box];
		A [label="Start"];
		B [label="Middle"];
		C [label="End"];

		A -> B [label="first"];
		B -> C [label="second"];

		subgraph cluster_0 {
			D;
			E;
		}
	}`

	g, err := ParseString(input)

	asrt.NoError(err, "Should parse complex graph without error")
	asrt.NotNil(g, "Graph should not be nil")
	asrt.Equal(5, len(g.Nodes()), "Graph should have 5 nodes")
	asrt.Equal(2, len(g.Edges()), "Graph should have 2 edges")
	asrt.Equal(1, len(g.Subgraphs()), "Graph should have 1 subgraph")
}

func TestParseString_SyntaxError(t *testing.T) {
	asrt := assert.New(t)

	input := "digraph { A -> }"

	g, err := ParseString(input)

	asrt.Error(err, "Should return error for invalid syntax")
	asrt.Nil(g, "Graph should be nil on error")

	// Check that error is a ParseError with location info
	perr, ok := err.(*ParseError)
	asrt.True(ok, "Error should be a ParseError")
	asrt.Greater(perr.Line, 0, "ParseError should have line number")
	asrt.Greater(perr.Col, 0, "ParseError should have column number")
}

func TestParse_FromReader(t *testing.T) {
	asrt := assert.New(t)

	input := "digraph { A -> B -> C; }"
	reader := strings.NewReader(input)

	g, err := Parse(reader)

	asrt.NoError(err, "Should parse from reader without error")
	asrt.NotNil(g, "Graph should not be nil")
	asrt.Equal(3, len(g.Nodes()), "Graph should have 3 nodes")
	asrt.Equal(2, len(g.Edges()), "Graph should have 2 edges")
}

func TestParseFile_ValidFile(t *testing.T) {
	asrt := assert.New(t)

	// Create temporary test file
	tmpDir := t.TempDir()
	testFile := filepath.Join(tmpDir, "test.dot")

	content := "digraph TestGraph { A -> B; B -> C; }"
	err := os.WriteFile(testFile, []byte(content), 0644)
	asrt.NoError(err, "Should write test file without error")

	// Parse the file
	g, err := ParseFile(testFile)

	asrt.NoError(err, "Should parse file without error")
	asrt.NotNil(g, "Graph should not be nil")
	asrt.Equal("TestGraph", g.Name(), "Graph name should match")
	asrt.Equal(3, len(g.Nodes()), "Graph should have 3 nodes")
	asrt.Equal(2, len(g.Edges()), "Graph should have 2 edges")
}

func TestParseFile_NotFound(t *testing.T) {
	asrt := assert.New(t)

	g, err := ParseFile("/nonexistent/file.dot")

	asrt.Error(err, "Should return error for nonexistent file")
	asrt.Nil(g, "Graph should be nil on error")

	// Check that error is a ParseError
	perr, ok := err.(*ParseError)
	asrt.True(ok, "Error should be a ParseError")
	asrt.Contains(perr.Message, "failed to open file", "Error message should mention file open failure")
}

func TestParseError_ErrorMessage(t *testing.T) {
	asrt := assert.New(t)

	tests := []struct {
		name     string
		perr     *ParseError
		expected string
	}{
		{
			name: "with location",
			perr: &ParseError{
				Message: "unexpected token",
				Line:    5,
				Col:     10,
			},
			expected: "parse error at 5:10: unexpected token",
		},
		{
			name: "without location",
			perr: &ParseError{
				Message: "unexpected token",
			},
			expected: "parse error: unexpected token",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			asrt.Equal(tt.expected, tt.perr.Error(), "Error message should match expected")
		})
	}
}

func TestParseString_EmptyInput(t *testing.T) {
	asrt := assert.New(t)

	g, err := ParseString("")

	asrt.Error(err, "Should return error for empty input")
	asrt.Nil(g, "Graph should be nil on error")
}

func TestParseString_WithSubgraphEndpoints(t *testing.T) {
	asrt := assert.New(t)

	input := `digraph {
		subgraph cluster_0 {
			A;
			B;
		} -> C;
	}`

	g, err := ParseString(input)

	asrt.NoError(err, "Should parse subgraph as edge endpoint without error")
	asrt.NotNil(g, "Graph should not be nil")
	asrt.Equal(3, len(g.Nodes()), "Graph should have 3 nodes")
	asrt.Equal(2, len(g.Edges()), "Graph should have 2 edges (A->C, B->C)")
}

func TestParse_RoundTrip_SimpleGraph(t *testing.T) {
	asrt := assert.New(t)

	// Create a simple graph
	g1 := NewGraph(Directed, WithName("TestGraph"))
	nodeA := NewNode("A", WithLabel("Start"))
	nodeB := NewNode("B", WithLabel("End"))
	g1.AddNode(nodeA)
	g1.AddNode(nodeB)
	g1.AddEdge(nodeA, nodeB, WithEdgeLabel("connects"))

	// Convert to DOT
	dot := g1.String()

	// Parse the DOT back
	g2, err := ParseString(dot)

	asrt.NoError(err, "Should parse generated DOT without error")
	asrt.NotNil(g2, "Parsed graph should not be nil")

	// Verify semantic equivalence
	asrt.Equal(g1.Name(), g2.Name(), "Graph names should match")
	asrt.Equal(g1.IsDirected(), g2.IsDirected(), "Directedness should match")
	asrt.Equal(len(g1.Nodes()), len(g2.Nodes()), "Node count should match")
	asrt.Equal(len(g1.Edges()), len(g2.Edges()), "Edge count should match")

	// Verify nodes exist
	asrt.NotNil(g2.GetNode("A"), "Node A should exist")
	asrt.NotNil(g2.GetNode("B"), "Node B should exist")

	// Verify node labels
	asrt.Equal("Start", g2.GetNode("A").Attrs().Label(), "Node A label should match")
	asrt.Equal("End", g2.GetNode("B").Attrs().Label(), "Node B label should match")

	// Verify edges
	edges := g2.Edges()
	asrt.Equal(1, len(edges), "Should have 1 edge")
	asrt.Equal("A", edges[0].From().ID(), "Edge should be from A")
	asrt.Equal("B", edges[0].To().ID(), "Edge should be to B")
	asrt.Equal("connects", edges[0].Attrs().Label(), "Edge label should match")
}

func TestParse_RoundTrip_ComplexGraph(t *testing.T) {
	asrt := assert.New(t)

	// Create a complex graph with subgraphs and attributes
	g1 := NewGraph(Directed, WithName("ComplexGraph"))

	// Add nodes
	nodeA := NewNode("A")
	nodeB := NewNode("B")
	nodeC := NewNode("C")
	g1.AddNode(nodeA)
	g1.AddNode(nodeB)
	g1.AddNode(nodeC)

	// Add edges
	g1.AddEdge(nodeA, nodeB)
	g1.AddEdge(nodeB, nodeC)

	// Add a subgraph
	g1.Subgraph("cluster_0", func(sg *Subgraph) {
		nodeD := NewNode("D")
		nodeE := NewNode("E")
		sg.AddNode(nodeD)
		sg.AddNode(nodeE)
		sg.AddEdge(nodeD, nodeE)
	})

	// Convert to DOT
	dot := g1.String()

	// Parse the DOT back
	g2, err := ParseString(dot)

	asrt.NoError(err, "Should parse generated DOT without error")
	asrt.NotNil(g2, "Parsed graph should not be nil")

	// Verify semantic equivalence
	asrt.Equal(g1.Name(), g2.Name(), "Graph names should match")
	asrt.Equal(g1.IsDirected(), g2.IsDirected(), "Directedness should match")
	asrt.Equal(len(g1.Nodes()), len(g2.Nodes()), "Node count should match")
	asrt.Equal(len(g1.Edges()), len(g2.Edges()), "Edge count should match")
	asrt.Equal(len(g1.Subgraphs()), len(g2.Subgraphs()), "Subgraph count should match")

	// Verify all nodes exist
	for _, node := range []string{"A", "B", "C", "D", "E"} {
		asrt.NotNil(g2.GetNode(node), "Node %s should exist", node)
	}

	// Verify subgraph
	subs := g2.Subgraphs()
	asrt.Equal(1, len(subs), "Should have 1 subgraph")
	asrt.Equal("cluster_0", subs[0].Name(), "Subgraph name should match")
	asrt.Equal(2, len(subs[0].Nodes()), "Subgraph should have 2 nodes")
}

func TestParse_RoundTrip_UndirectedGraph(t *testing.T) {
	asrt := assert.New(t)

	// Create an undirected graph
	g1 := NewGraph(Undirected, WithName("UndirectedGraph"))
	nodeA := NewNode("A")
	nodeB := NewNode("B")
	g1.AddNode(nodeA)
	g1.AddNode(nodeB)
	g1.AddEdge(nodeA, nodeB)

	// Convert to DOT
	dot := g1.String()

	// Parse the DOT back
	g2, err := ParseString(dot)

	asrt.NoError(err, "Should parse generated DOT without error")
	asrt.NotNil(g2, "Parsed graph should not be nil")

	// Verify it's still undirected
	asrt.False(g2.IsDirected(), "Graph should be undirected")
	asrt.Equal("UndirectedGraph", g2.Name(), "Graph name should match")
	asrt.Equal(2, len(g2.Nodes()), "Node count should match")
	asrt.Equal(1, len(g2.Edges()), "Edge count should match")
}

func TestParse_RoundTrip_StrictGraph(t *testing.T) {
	asrt := assert.New(t)

	// Create a strict graph
	g1 := NewGraph(Directed, Strict, WithName("StrictGraph"))
	nodeA := NewNode("A")
	nodeB := NewNode("B")
	g1.AddNode(nodeA)
	g1.AddNode(nodeB)
	g1.AddEdge(nodeA, nodeB)

	// Convert to DOT
	dot := g1.String()

	// Parse the DOT back
	g2, err := ParseString(dot)

	asrt.NoError(err, "Should parse generated DOT without error")
	asrt.NotNil(g2, "Parsed graph should not be nil")

	// Verify it's still strict
	asrt.True(g2.IsStrict(), "Graph should be strict")
	asrt.True(g2.IsDirected(), "Graph should be directed")
	asrt.Equal("StrictGraph", g2.Name(), "Graph name should match")
}

func TestParseFile_SimpleFixture(t *testing.T) {
	asrt := assert.New(t)

	g, err := ParseFile("testdata/simple.dot")

	asrt.NoError(err, "Should parse simple.dot without error")
	asrt.NotNil(g, "Graph should not be nil")
	asrt.Equal("SimpleGraph", g.Name(), "Graph name should be SimpleGraph")
	asrt.Equal(3, len(g.Nodes()), "Graph should have 3 nodes")
	asrt.Equal(3, len(g.Edges()), "Graph should have 3 edges")
}

func TestParseFile_ComplexFixture(t *testing.T) {
	asrt := assert.New(t)

	g, err := ParseFile("testdata/complex.dot")

	asrt.NoError(err, "Should parse complex.dot without error")
	asrt.NotNil(g, "Graph should not be nil")
	asrt.Equal("ComplexGraph", g.Name(), "Graph name should be ComplexGraph")
	asrt.Equal(5, len(g.Nodes()), "Graph should have 5 nodes (A, B, C, D, E)")
	asrt.Equal(4, len(g.Edges()), "Graph should have 4 edges")

	// Verify some node attributes
	nodeA := g.GetNode("A")
	asrt.NotNil(nodeA, "Node A should exist")
	asrt.Equal("Start Node", nodeA.Attrs().Label(), "Node A should have correct label")

	// Verify some edge attributes
	edges := g.Edges()
	hasLabeledEdge := false
	for _, edge := range edges {
		if edge.Attrs().Label() == "first step" {
			hasLabeledEdge = true
			break
		}
	}
	asrt.True(hasLabeledEdge, "Should have edge with label 'first step'")
}

func TestParseFile_ClusterFixture(t *testing.T) {
	asrt := assert.New(t)

	g, err := ParseFile("testdata/cluster.dot")

	asrt.NoError(err, "Should parse cluster.dot without error")
	asrt.NotNil(g, "Graph should not be nil")
	asrt.Equal("ClusterGraph", g.Name(), "Graph name should be ClusterGraph")
	asrt.Equal(6, len(g.Nodes()), "Graph should have 6 nodes")
	asrt.Equal(5, len(g.Edges()), "Graph should have 5 edges")
	asrt.Equal(2, len(g.Subgraphs()), "Graph should have 2 subgraphs")

	// Verify subgraphs
	for _, sg := range g.Subgraphs() {
		asrt.True(sg.IsCluster(), "Subgraph should be a cluster")
		asrt.True(len(sg.Nodes()) > 0, "Subgraph should have nodes")
	}
}
