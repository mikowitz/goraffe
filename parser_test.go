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
