package goraffe

import (
	"bytes"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Test Graph.String() for empty directed graph
func TestGraph_String_EmptyDirected(t *testing.T) {
	t.Run("outputs basic directed graph syntax", func(t *testing.T) {
		asrt := assert.New(t)
		g := NewGraph(Directed)

		output := g.String()

		asrt.Equal("digraph {\n}", output, "expected directed graph with empty body")
	})

	t.Run("default graph without Directed option is undirected", func(t *testing.T) {
		asrt := assert.New(t)
		g := NewGraph()

		output := g.String()

		asrt.NotContains(output, "digraph", "expected default graph to not use digraph keyword")
	})
}

// Test Graph.String() for empty undirected graph
func TestGraph_String_EmptyUndirected(t *testing.T) {
	t.Run("outputs basic undirected graph syntax", func(t *testing.T) {
		asrt := assert.New(t)
		g := NewGraph(Undirected)

		output := g.String()

		asrt.Equal("graph {\n}", output, "expected undirected graph with empty body")
	})

	t.Run("explicit Undirected option produces undirected graph", func(t *testing.T) {
		asrt := assert.New(t)
		g := NewGraph(Undirected)

		output := g.String()

		asrt.Contains(output, "graph {", "expected graph keyword for undirected graph")
		asrt.NotContains(output, "digraph", "expected no digraph keyword for undirected graph")
	})

	t.Run("default NewGraph is undirected", func(t *testing.T) {
		asrt := assert.New(t)
		g := NewGraph()

		output := g.String()

		asrt.Equal("graph {\n}", output, "expected default graph to be undirected")
	})
}

// Test Graph.String() with Strict option
func TestGraph_String_Strict(t *testing.T) {
	t.Run("strict directed graph", func(t *testing.T) {
		asrt := assert.New(t)
		g := NewGraph(Directed, Strict)

		output := g.String()

		asrt.Equal("strict digraph {\n}", output, "expected strict prefix for strict directed graph")
	})

	t.Run("strict undirected graph", func(t *testing.T) {
		asrt := assert.New(t)
		g := NewGraph(Undirected, Strict)

		output := g.String()

		asrt.Equal("strict graph {\n}", output, "expected strict prefix for strict undirected graph")
	})

	t.Run("strict keyword comes before graph type", func(t *testing.T) {
		asrt := assert.New(t)
		g := NewGraph(Strict, Directed)

		output := g.String()

		asrt.True(strings.HasPrefix(output, "strict digraph"), "expected 'strict' to come before 'digraph'")
	})

	t.Run("non-strict graph has no strict keyword", func(t *testing.T) {
		asrt := assert.New(t)
		g := NewGraph(Directed)

		output := g.String()

		asrt.NotContains(output, "strict", "expected no strict keyword for non-strict graph")
	})
}

// Test Graph.String() with graph name
func TestGraph_String_WithName(t *testing.T) {
	t.Run("directed graph with name", func(t *testing.T) {
		asrt := assert.New(t)
		g := NewGraph(Directed, WithName("G"))

		output := g.String()

		asrt.Equal("digraph G {\n}", output, "expected graph name between type and opening brace")
	})

	t.Run("undirected graph with name", func(t *testing.T) {
		asrt := assert.New(t)
		g := NewGraph(Undirected, WithName("MyGraph"))

		output := g.String()

		asrt.Equal("graph MyGraph {\n}", output, "expected graph name for undirected graph")
	})

	t.Run("strict directed graph with name", func(t *testing.T) {
		asrt := assert.New(t)
		g := NewGraph(Strict, Directed, WithName("StrictGraph"))

		output := g.String()

		asrt.Equal("strict digraph StrictGraph {\n}", output, "expected strict, digraph, name order")
	})

	t.Run("graph name with underscores", func(t *testing.T) {
		asrt := assert.New(t)
		g := NewGraph(Directed, WithName("my_graph_name"))

		output := g.String()

		asrt.Contains(output, "digraph my_graph_name", "expected name with underscores to be preserved")
	})

	t.Run("graph name with numbers", func(t *testing.T) {
		asrt := assert.New(t)
		g := NewGraph(Directed, WithName("Graph123"))

		output := g.String()

		asrt.Contains(output, "digraph Graph123", "expected name with numbers to be preserved")
	})

	t.Run("empty name is omitted", func(t *testing.T) {
		asrt := assert.New(t)
		g := NewGraph(Directed, WithName(""))

		output := g.String()

		asrt.Equal("digraph {\n}", output, "expected empty name to be omitted from output")
	})
}

// Test Graph.WriteDOT() method
func TestGraph_WriteDOT_WritesToWriter(t *testing.T) {
	t.Run("writes directed graph to writer", func(t *testing.T) {
		asrt := assert.New(t)
		g := NewGraph(Directed)
		var buf bytes.Buffer

		err := g.WriteDOT(&buf)

		asrt.NoError(err, "expected WriteDOT to succeed without error")
		asrt.Equal("digraph {\n}", buf.String(), "expected DOT output to be written to buffer")
	})

	t.Run("writes undirected graph to writer", func(t *testing.T) {
		asrt := assert.New(t)
		g := NewGraph(Undirected)
		var buf bytes.Buffer

		err := g.WriteDOT(&buf)

		asrt.NoError(err, "expected WriteDOT to succeed without error")
		asrt.Equal("graph {\n}", buf.String(), "expected undirected DOT output")
	})

	t.Run("writes strict graph to writer", func(t *testing.T) {
		asrt := assert.New(t)
		g := NewGraph(Strict, Directed)
		var buf bytes.Buffer

		err := g.WriteDOT(&buf)

		asrt.NoError(err, "expected WriteDOT to succeed without error")
		asrt.Equal("strict digraph {\n}", buf.String(), "expected strict graph output")
	})

	t.Run("writes named graph to writer", func(t *testing.T) {
		asrt := assert.New(t)
		g := NewGraph(Directed, WithName("TestGraph"))
		var buf bytes.Buffer

		err := g.WriteDOT(&buf)

		asrt.NoError(err, "expected WriteDOT to succeed without error")
		asrt.Equal("digraph TestGraph {\n}", buf.String(), "expected named graph output")
	})

	t.Run("WriteDOT and String produce same output", func(t *testing.T) {
		asrt := assert.New(t)
		g := NewGraph(Strict, Directed, WithName("G"))
		var buf bytes.Buffer

		err := g.WriteDOT(&buf)
		stringOutput := g.String()

		asrt.NoError(err, "expected WriteDOT to succeed")
		asrt.Equal(stringOutput, buf.String(), "expected WriteDOT and String to produce identical output")
	})

	t.Run("multiple writes to same writer append content", func(t *testing.T) {
		asrt := assert.New(t)
		g1 := NewGraph(Directed)
		g2 := NewGraph(Undirected)
		var buf bytes.Buffer

		err1 := g1.WriteDOT(&buf)
		err2 := g2.WriteDOT(&buf)

		asrt.NoError(err1, "expected first WriteDOT to succeed")
		asrt.NoError(err2, "expected second WriteDOT to succeed")
		expected := "digraph {\n}graph {\n}"
		asrt.Equal(expected, buf.String(), "expected both graphs written to buffer")
	})

	t.Run("writes to strings.Builder", func(t *testing.T) {
		asrt := assert.New(t)
		g := NewGraph(Directed, WithName("SB"))
		var sb strings.Builder

		err := g.WriteDOT(&sb)

		asrt.NoError(err, "expected WriteDOT to work with strings.Builder")
		asrt.Equal("digraph SB {\n}", sb.String(), "expected output written to strings.Builder")
	})
}

// Test output formatting
func TestGraph_String_OutputFormatting(t *testing.T) {
	t.Run("opening brace on same line as graph declaration", func(t *testing.T) {
		asrt := assert.New(t)
		g := NewGraph(Directed)

		output := g.String()

		asrt.Contains(output, "digraph {", "expected opening brace on same line")
		asrt.NotContains(output, "digraph\n{", "expected no newline before opening brace")
	})

	t.Run("closing brace on its own line", func(t *testing.T) {
		asrt := assert.New(t)
		g := NewGraph(Directed)

		output := g.String()

		lines := strings.Split(output, "\n")
		asrt.Equal("}", lines[1], "expected closing brace on second line")
	})

	t.Run("output ends with closing brace", func(t *testing.T) {
		asrt := assert.New(t)
		g := NewGraph(Directed)

		output := g.String()

		asrt.True(strings.HasSuffix(output, "}"), "expected output to end with closing brace")
		asrt.False(strings.HasSuffix(output, "\n"), "expected output to NOT end with trailing newline")
	})

	t.Run("exactly two lines for empty graph", func(t *testing.T) {
		asrt := assert.New(t)
		g := NewGraph(Directed)

		output := g.String()

		// Split by newline will give us ["digraph {", "}"]
		lines := strings.Split(output, "\n")
		asrt.Equal(2, len(lines), "expected two lines when split by newline")
		asrt.Equal("digraph {", lines[0], "expected first line to be graph declaration")
		asrt.Equal("}", lines[1], "expected second line to be closing brace")
	})

	t.Run("space between strict and digraph", func(t *testing.T) {
		asrt := assert.New(t)
		g := NewGraph(Strict, Directed)

		output := g.String()

		asrt.Contains(output, "strict digraph", "expected space between strict and digraph")
		asrt.NotContains(output, "strictdigraph", "expected no concatenation of keywords")
	})

	t.Run("space between digraph and name", func(t *testing.T) {
		asrt := assert.New(t)
		g := NewGraph(Directed, WithName("G"))

		output := g.String()

		asrt.Contains(output, "digraph G", "expected space between digraph and name")
		asrt.NotContains(output, "digraphG", "expected no concatenation of keyword and name")
	})

	t.Run("space between name and opening brace", func(t *testing.T) {
		asrt := assert.New(t)
		g := NewGraph(Directed, WithName("G"))

		output := g.String()

		asrt.Contains(output, "G {", "expected space between name and opening brace")
		asrt.NotContains(output, "G{", "expected no concatenation of name and brace")
	})
}

// Test option order independence
func TestGraph_String_OptionOrderIndependence(t *testing.T) {
	t.Run("Directed then Strict produces same output as Strict then Directed", func(t *testing.T) {
		asrt := assert.New(t)
		g1 := NewGraph(Directed, Strict)
		g2 := NewGraph(Strict, Directed)

		output1 := g1.String()
		output2 := g2.String()

		asrt.Equal(output1, output2, "expected option order not to affect output")
		asrt.Equal("strict digraph {\n}", output1, "expected consistent strict digraph output")
	})

	t.Run("options with name in different positions", func(t *testing.T) {
		asrt := assert.New(t)
		g1 := NewGraph(WithName("G"), Directed, Strict)
		g2 := NewGraph(Directed, WithName("G"), Strict)
		g3 := NewGraph(Strict, Directed, WithName("G"))

		output1 := g1.String()
		output2 := g2.String()
		output3 := g3.String()

		asrt.Equal(output1, output2, "expected same output regardless of option order")
		asrt.Equal(output2, output3, "expected same output regardless of option order")
		asrt.Equal("strict digraph G {\n}", output1, "expected consistent output")
	})
}

// Test edge cases
func TestGraph_String_EdgeCases(t *testing.T) {
	t.Run("graph with only Strict but no direction defaults to undirected", func(t *testing.T) {
		asrt := assert.New(t)
		g := NewGraph(Strict)

		output := g.String()

		asrt.Equal("strict graph {\n}", output, "expected strict undirected graph")
	})

	t.Run("multiple Directed options last one wins", func(t *testing.T) {
		asrt := assert.New(t)
		g := NewGraph(Directed, Undirected)

		output := g.String()

		asrt.Equal("graph {\n}", output, "expected last Undirected option to override Directed")
	})

	t.Run("Undirected then Directed produces directed graph", func(t *testing.T) {
		asrt := assert.New(t)
		g := NewGraph(Undirected, Directed)

		output := g.String()

		asrt.Equal("digraph {\n}", output, "expected last Directed option to override Undirected")
	})

	t.Run("multiple WithName options last one wins", func(t *testing.T) {
		asrt := assert.New(t)
		g := NewGraph(Directed, WithName("First"), WithName("Second"))

		output := g.String()

		asrt.Equal("digraph Second {\n}", output, "expected last name to be used")
		asrt.NotContains(output, "First", "expected earlier name to be overridden")
	})
}
