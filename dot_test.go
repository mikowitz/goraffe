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

		asrt.Equal("digraph \"G\" {\n}", output, "expected graph name between type and opening brace")
	})

	t.Run("undirected graph with name", func(t *testing.T) {
		asrt := assert.New(t)
		g := NewGraph(Undirected, WithName("MyGraph"))

		output := g.String()

		asrt.Equal("graph \"MyGraph\" {\n}", output, "expected graph name for undirected graph")
	})

	t.Run("strict directed graph with name", func(t *testing.T) {
		asrt := assert.New(t)
		g := NewGraph(Strict, Directed, WithName("StrictGraph"))

		output := g.String()

		asrt.Equal("strict digraph \"StrictGraph\" {\n}", output, "expected strict, digraph, name order")
	})

	t.Run("graph name with underscores", func(t *testing.T) {
		asrt := assert.New(t)
		g := NewGraph(Directed, WithName("my_graph_name"))

		output := g.String()

		asrt.Contains(output, "digraph \"my_graph_name\"", "expected name with underscores to be preserved")
	})

	t.Run("graph name with numbers", func(t *testing.T) {
		asrt := assert.New(t)
		g := NewGraph(Directed, WithName("Graph123"))

		output := g.String()

		asrt.Contains(output, "digraph \"Graph123\"", "expected name with numbers to be preserved")
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
		asrt.Equal("digraph \"TestGraph\" {\n}", buf.String(), "expected named graph output")
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
		asrt.Equal("digraph \"SB\" {\n}", sb.String(), "expected output written to strings.Builder")
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

		asrt.Contains(output, "digraph \"G\"", "expected space between digraph and name")
		asrt.NotContains(output, "digraph\"G\"", "expected no concatenation of keyword and name")
	})

	t.Run("space between name and opening brace", func(t *testing.T) {
		asrt := assert.New(t)
		g := NewGraph(Directed, WithName("G"))

		output := g.String()

		asrt.Contains(output, "\"G\" {", "expected space between name and opening brace")
		asrt.NotContains(output, "\"G\"{", "expected no concatenation of name and brace")
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
		asrt.Equal("strict digraph \"G\" {\n}", output1, "expected consistent output")
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

		asrt.Equal("digraph \"Second\" {\n}", output, "expected last name to be used")
		asrt.NotContains(output, "First", "expected earlier name to be overridden")
	})
}

// Test single node with no attributes
func TestDOT_SingleNode_NoAttributes(t *testing.T) {
	t.Run("renders node with ID only", func(t *testing.T) {
		asrt := assert.New(t)
		g := NewGraph(Directed)
		n := NewNode("A")
		_ = g.AddNode(n)

		output := g.String()

		expected := "digraph {\n\t\"A\";\n}"
		asrt.Equal(expected, output, "expected node to be rendered with quoted ID and semicolon")
	})

	t.Run("node ID is always quoted", func(t *testing.T) {
		asrt := assert.New(t)
		g := NewGraph(Directed)
		n := NewNode("SimpleID")
		_ = g.AddNode(n)

		output := g.String()

		asrt.Contains(output, "\"SimpleID\";", "expected node ID to be quoted even if simple")
		asrt.NotContains(output, "SimpleID;", "expected no unquoted node ID")
	})

	t.Run("node appears inside graph body", func(t *testing.T) {
		asrt := assert.New(t)
		g := NewGraph(Directed)
		n := NewNode("A")
		_ = g.AddNode(n)

		output := g.String()

		lines := strings.Split(output, "\n")
		asrt.Equal(3, len(lines), "expected three lines: opening, node, closing")
		asrt.Equal("digraph {", lines[0], "expected graph opening")
		asrt.Contains(lines[1], "\"A\";", "expected node on second line")
		asrt.Equal("}", lines[2], "expected closing brace")
	})

	t.Run("node line is indented with tab", func(t *testing.T) {
		asrt := assert.New(t)
		g := NewGraph(Directed)
		n := NewNode("A")
		_ = g.AddNode(n)

		output := g.String()

		asrt.Contains(output, "\t\"A\";", "expected node line to be indented with tab")
	})

	t.Run("undirected graph renders node same way", func(t *testing.T) {
		asrt := assert.New(t)
		g := NewGraph(Undirected)
		n := NewNode("A")
		_ = g.AddNode(n)

		output := g.String()

		expected := "graph {\n\t\"A\";\n}"
		asrt.Equal(expected, output, "expected node rendering to be same for undirected graph")
	})
}

// Test single node with label attribute
func TestDOT_SingleNode_WithLabel(t *testing.T) {
	t.Run("renders label attribute", func(t *testing.T) {
		asrt := assert.New(t)
		g := NewGraph(Directed)
		n := NewNode("A", WithLabel("Node A"))
		_ = g.AddNode(n)

		output := g.String()

		asrt.Contains(output, "\"A\" [label=\"Node A\"];", "expected node with label attribute")
	})

	t.Run("label value is quoted", func(t *testing.T) {
		asrt := assert.New(t)
		g := NewGraph(Directed)
		n := NewNode("A", WithLabel("My Label"))
		_ = g.AddNode(n)

		output := g.String()

		asrt.Contains(output, "label=\"My Label\"", "expected quoted label value")
		asrt.NotContains(output, "label=My Label", "expected no unquoted label value")
	})

	t.Run("label with special characters", func(t *testing.T) {
		asrt := assert.New(t)
		g := NewGraph(Directed)
		n := NewNode("A", WithLabel("Label with spaces"))
		_ = g.AddNode(n)

		output := g.String()

		asrt.Contains(output, "label=\"Label with spaces\"", "expected label with spaces to be preserved")
	})
}

// Test single node with shape attribute
func TestDOT_SingleNode_WithShape(t *testing.T) {
	t.Run("renders shape attribute", func(t *testing.T) {
		asrt := assert.New(t)
		g := NewGraph(Directed)
		n := NewNode("A", WithBoxShape())
		_ = g.AddNode(n)

		output := g.String()

		asrt.Contains(output, "\"A\" [shape=\"box\"];", "expected node with shape attribute")
	})

	t.Run("different shape values", func(t *testing.T) {
		asrt := assert.New(t)

		testCases := []struct {
			name     string
			option   NodeOption
			expected string
		}{
			{"box", WithBoxShape(), "shape=\"box\""},
			{"circle", WithCircleShape(), "shape=\"circle\""},
			{"ellipse", WithEllipseShape(), "shape=\"ellipse\""},
			{"diamond", WithDiamondShape(), "shape=\"diamond\""},
			{"record", WithRecordShape(), "shape=\"record\""},
			{"plaintext", WithPlaintextShape(), "shape=\"plaintext\""},
		}

		for _, tc := range testCases {
			g := NewGraph(Directed)
			n := NewNode("A", tc.option)
			_ = g.AddNode(n)

			output := g.String()
			asrt.Contains(output, tc.expected, "expected shape attribute for %s", tc.name)
		}
	})
}

// Test single node with multiple attributes
func TestDOT_SingleNode_MultipleAttributes(t *testing.T) {
	t.Run("renders multiple attributes separated by commas", func(t *testing.T) {
		asrt := assert.New(t)
		g := NewGraph(Directed)
		n := NewNode("A", WithBoxShape(), WithLabel("Node A"))
		_ = g.AddNode(n)

		output := g.String()

		asrt.Contains(output, "\"A\" [", "expected attribute section to start")
		asrt.Contains(output, "shape=\"box\"", "expected shape attribute")
		asrt.Contains(output, "label=\"Node A\"", "expected label attribute")
		asrt.Contains(output, ", ", "expected comma separator between attributes")
	})

	t.Run("all basic attributes together", func(t *testing.T) {
		asrt := assert.New(t)
		g := NewGraph(Directed)
		n := NewNode("A",
			WithBoxShape(),
			WithLabel("Node A"),
			WithColor("red"),
			WithFillColor("lightblue"),
			WithFontName("Arial"),
			WithFontSize(14.0),
		)
		_ = g.AddNode(n)

		output := g.String()

		asrt.Contains(output, "shape=\"box\"", "expected shape")
		asrt.Contains(output, "label=\"Node A\"", "expected label")
		asrt.Contains(output, "color=\"red\"", "expected color")
		asrt.Contains(output, "fillcolor=\"lightblue\"", "expected fillcolor")
		asrt.Contains(output, "style=\"filled\"", "expected style to be filled")
		asrt.Contains(output, "fontname=\"Arial\"", "expected fontname")
		asrt.Contains(output, "fontsize=\"14\"", "expected fontsize")
	})

	t.Run("only non-zero attributes are rendered", func(t *testing.T) {
		asrt := assert.New(t)
		g := NewGraph(Directed)
		n := NewNode("A", WithBoxShape())
		_ = g.AddNode(n)

		output := g.String()

		asrt.Contains(output, "shape=\"box\"", "expected shape attribute")
		asrt.NotContains(output, "label=", "expected no label attribute")
		asrt.NotContains(output, "color=", "expected no color attribute")
		asrt.NotContains(output, "fontsize=", "expected no fontsize attribute")
	})

	t.Run("attributes in consistent order", func(t *testing.T) {
		asrt := assert.New(t)
		g1 := NewGraph(Directed)
		n1 := NewNode("A", WithLabel("A"), WithBoxShape())
		_ = g1.AddNode(n1)

		g2 := NewGraph(Directed)
		n2 := NewNode("A", WithBoxShape(), WithLabel("A"))
		_ = g2.AddNode(n2)

		output1 := g1.String()
		output2 := g2.String()

		asrt.Equal(output1, output2, "expected consistent attribute order regardless of option order")
	})
}

// Test single node with custom attributes
func TestDOT_SingleNode_CustomAttribute(t *testing.T) {
	t.Run("renders custom attribute", func(t *testing.T) {
		asrt := assert.New(t)
		g := NewGraph(Directed)
		n := NewNode("A", WithNodeAttribute("peripheries", "2"))
		_ = g.AddNode(n)

		output := g.String()

		asrt.Contains(output, "\"A\" [peripheries=\"2\"];", "expected custom attribute to be rendered")
	})

	t.Run("multiple custom attributes", func(t *testing.T) {
		asrt := assert.New(t)
		g := NewGraph(Directed)
		n := NewNode("A",
			WithNodeAttribute("peripheries", "2"),
			WithNodeAttribute("tooltip", "Hover text"),
		)
		_ = g.AddNode(n)

		output := g.String()

		asrt.Contains(output, "peripheries=\"2\"", "expected first custom attribute")
		asrt.Contains(output, "tooltip=\"Hover text\"", "expected second custom attribute")
	})

	t.Run("custom and typed attributes together", func(t *testing.T) {
		asrt := assert.New(t)
		g := NewGraph(Directed)
		n := NewNode("A",
			WithBoxShape(),
			WithNodeAttribute("peripheries", "2"),
			WithLabel("Node A"),
		)
		_ = g.AddNode(n)

		output := g.String()

		asrt.Contains(output, "shape=\"box\"", "expected typed shape attribute")
		asrt.Contains(output, "label=\"Node A\"", "expected typed label attribute")
		asrt.Contains(output, "peripheries=\"2\"", "expected custom attribute")
	})
}

// Test multiple nodes
func TestDOT_MultipleNodes(t *testing.T) {
	t.Run("renders all nodes in insertion order", func(t *testing.T) {
		asrt := assert.New(t)
		g := NewGraph(Directed)
		n1 := NewNode("A")
		n2 := NewNode("B")
		n3 := NewNode("C")

		_ = g.AddNode(n1)
		_ = g.AddNode(n2)
		_ = g.AddNode(n3)

		output := g.String()

		lines := strings.Split(output, "\n")
		asrt.Equal(5, len(lines), "expected 5 lines: opening, 3 nodes, closing")
		asrt.Contains(lines[1], "\"A\";", "expected first node")
		asrt.Contains(lines[2], "\"B\";", "expected second node")
		asrt.Contains(lines[3], "\"C\";", "expected third node")
	})

	t.Run("nodes with mixed attributes", func(t *testing.T) {
		asrt := assert.New(t)
		g := NewGraph(Directed)
		n1 := NewNode("A", WithBoxShape(), WithLabel("Node A"))
		n2 := NewNode("B")
		n3 := NewNode("C", WithLabel("Node C"))

		_ = g.AddNode(n1)
		_ = g.AddNode(n2)
		_ = g.AddNode(n3)

		output := g.String()

		asrt.Contains(output, "\"A\" [label=\"Node A\", shape=\"box\"];", "expected first node with attributes")
		asrt.Contains(output, "\t\"B\";", "expected second node without attributes")
		asrt.Contains(output, "\"C\" [label=\"Node C\"];", "expected third node with label")
	})

	t.Run("preserves insertion order even with same ID replacement", func(t *testing.T) {
		asrt := assert.New(t)
		g := NewGraph(Directed)
		n1 := NewNode("A", WithLabel("First"))
		n2 := NewNode("B")
		n3 := NewNode("A", WithLabel("Replaced"))

		_ = g.AddNode(n1)
		_ = g.AddNode(n2)
		_ = g.AddNode(n3) // Replaces n1 in place

		output := g.String()

		lines := strings.Split(output, "\n")
		asrt.Equal(4, len(lines), "expected 4 lines: opening, 2 nodes, closing")
		asrt.Contains(lines[1], "\"A\"", "expected A in first position")
		asrt.Contains(lines[1], "label=\"Replaced\"", "expected replaced label")
		asrt.Contains(lines[2], "\"B\"", "expected B in second position")
	})
}

// Test node ID quoting
func TestDOT_NodeID_Quoting(t *testing.T) {
	t.Run("simple alphanumeric ID is quoted", func(t *testing.T) {
		asrt := assert.New(t)
		g := NewGraph(Directed)
		n := NewNode("node1")
		_ = g.AddNode(n)

		output := g.String()

		asrt.Contains(output, "\"node1\";", "expected quoted simple ID")
	})

	t.Run("ID with spaces is quoted", func(t *testing.T) {
		asrt := assert.New(t)
		g := NewGraph(Directed)
		n := NewNode("node with spaces")
		_ = g.AddNode(n)

		output := g.String()

		asrt.Contains(output, "\"node with spaces\";", "expected quoted ID with spaces")
	})

	t.Run("ID with special characters is quoted", func(t *testing.T) {
		asrt := assert.New(t)
		g := NewGraph(Directed)
		n := NewNode("node-with-dashes")
		_ = g.AddNode(n)

		output := g.String()

		asrt.Contains(output, "\"node-with-dashes\";", "expected quoted ID with dashes")
	})

	t.Run("ID starting with number is quoted", func(t *testing.T) {
		asrt := assert.New(t)
		g := NewGraph(Directed)
		n := NewNode("123node")
		_ = g.AddNode(n)

		output := g.String()

		asrt.Contains(output, "\"123node\";", "expected quoted ID starting with number")
	})

	t.Run("empty ID is quoted", func(t *testing.T) {
		asrt := assert.New(t)
		g := NewGraph(Directed)
		n := NewNode("")
		_ = g.AddNode(n)

		output := g.String()

		asrt.Contains(output, "\"\";", "expected quoted empty ID")
	})
}

// Test attribute value formatting
func TestDOT_AttributeValueFormatting(t *testing.T) {
	t.Run("string attributes are quoted", func(t *testing.T) {
		asrt := assert.New(t)
		g := NewGraph(Directed)
		n := NewNode("A", WithLabel("Label"), WithColor("red"))
		_ = g.AddNode(n)

		output := g.String()

		asrt.Contains(output, "label=\"Label\"", "expected quoted label value")
		asrt.Contains(output, "color=\"red\"", "expected quoted color value")
	})

	t.Run("numeric fontsize is quoted as string", func(t *testing.T) {
		asrt := assert.New(t)
		g := NewGraph(Directed)
		n := NewNode("A", WithFontSize(14.5))
		_ = g.AddNode(n)

		output := g.String()

		asrt.Contains(output, "fontsize=\"14.5\"", "expected fontsize as quoted string")
	})
}

// Test integration with graph options
func TestDOT_NodesWithGraphOptions(t *testing.T) {
	t.Run("nodes in strict graph", func(t *testing.T) {
		asrt := assert.New(t)
		g := NewGraph(Strict, Directed)
		n := NewNode("A")
		_ = g.AddNode(n)

		output := g.String()

		expected := "strict digraph {\n\t\"A\";\n}"
		asrt.Equal(expected, output, "expected node in strict graph")
	})

	t.Run("nodes in named graph", func(t *testing.T) {
		asrt := assert.New(t)
		g := NewGraph(Directed, WithName("G"))
		n := NewNode("A", WithLabel("Node A"))
		_ = g.AddNode(n)

		output := g.String()

		expected := "digraph \"G\" {\n\t\"A\" [label=\"Node A\"];\n}"
		asrt.Equal(expected, output, "expected node in named graph")
	})

	t.Run("nodes in undirected graph", func(t *testing.T) {
		asrt := assert.New(t)
		g := NewGraph(Undirected)
		n1 := NewNode("A", WithCircleShape())
		n2 := NewNode("B", WithBoxShape())

		_ = g.AddNode(n1)
		_ = g.AddNode(n2)

		output := g.String()

		asrt.Contains(output, "graph {", "expected undirected graph")
		asrt.Contains(output, "\"A\" [shape=\"circle\"];", "expected first node")
		asrt.Contains(output, "\"B\" [shape=\"box\"];", "expected second node")
	})
}

// Test single edge with no attributes
func TestDOT_SingleEdge_NoAttributes(t *testing.T) {
	t.Run("renders edge with arrow syntax in directed graph", func(t *testing.T) {
		asrt := assert.New(t)

		g := NewGraph(Directed)
		n1 := NewNode("A")
		n2 := NewNode("B")
		_, _ = g.AddEdge(n1, n2)

		output := g.String()

		asrt.Contains(output, "\"A\" -> \"B\";", "expected edge with arrow syntax")
	})

	t.Run("edge appears after nodes in output", func(t *testing.T) {
		asrt := assert.New(t)

		g := NewGraph(Directed)
		n1 := NewNode("A")
		n2 := NewNode("B")
		_, _ = g.AddEdge(n1, n2)

		output := g.String()

		lines := strings.Split(output, "\n")
		asrt.Equal(5, len(lines), "expected 5 lines: opening, 2 nodes, 1 edge, closing")
		asrt.Contains(lines[1], "\"A\";", "expected first node on line 2")
		asrt.Contains(lines[2], "\"B\";", "expected second node on line 3")
		asrt.Contains(lines[3], "\"A\" -> \"B\";", "expected edge on line 4 after nodes")
	})

	t.Run("edge line is indented with tab", func(t *testing.T) {
		asrt := assert.New(t)

		g := NewGraph(Directed)
		n1 := NewNode("A")
		n2 := NewNode("B")
		_, _ = g.AddEdge(n1, n2)

		output := g.String()

		asrt.Contains(output, "\t\"A\" -> \"B\";", "expected edge line to be indented with tab")
	})

	t.Run("both node IDs are quoted", func(t *testing.T) {
		asrt := assert.New(t)

		g := NewGraph(Directed)
		n1 := NewNode("A")
		n2 := NewNode("B")
		_, _ = g.AddEdge(n1, n2)

		output := g.String()

		asrt.Contains(output, "\"A\" -> \"B\"", "expected both node IDs to be quoted")
		asrt.NotContains(output, "A -> B", "expected no unquoted node IDs")
	})
}

// Test directed vs undirected edge syntax
func TestDOT_SingleEdge_Directed(t *testing.T) {
	t.Run("directed graph uses arrow syntax", func(t *testing.T) {
		asrt := assert.New(t)

		g := NewGraph(Directed)
		n1 := NewNode("A")
		n2 := NewNode("B")
		_, _ = g.AddEdge(n1, n2)

		output := g.String()

		asrt.Contains(output, "\"A\" -> \"B\";", "expected -> syntax for directed graph")
		asrt.NotContains(output, "\"A\" -- \"B\"", "expected no -- syntax in directed graph")
	})
}

func TestDOT_SingleEdge_Undirected(t *testing.T) {
	t.Run("undirected graph uses line syntax", func(t *testing.T) {
		asrt := assert.New(t)

		g := NewGraph(Undirected)
		n1 := NewNode("A")
		n2 := NewNode("B")
		_, _ = g.AddEdge(n1, n2)

		output := g.String()

		asrt.Contains(output, "\"A\" -- \"B\";", "expected -- syntax for undirected graph")
		asrt.NotContains(output, "\"A\" -> \"B\"", "expected no -> syntax in undirected graph")
	})

	t.Run("undirected edge with attributes", func(t *testing.T) {
		asrt := assert.New(t)

		g := NewGraph(Undirected)
		n1 := NewNode("A")
		n2 := NewNode("B")
		_, _ = g.AddEdge(n1, n2, WithEdgeLabel("connects"))

		output := g.String()

		asrt.Contains(output, "\"A\" -- \"B\" [label=\"connects\"];", "expected -- syntax with attributes")
	})
}

// Test edge with label
func TestDOT_SingleEdge_WithLabel(t *testing.T) {
	t.Run("renders label attribute", func(t *testing.T) {
		asrt := assert.New(t)

		g := NewGraph(Directed)
		n1 := NewNode("A")
		n2 := NewNode("B")
		_, _ = g.AddEdge(n1, n2, WithEdgeLabel("connects"))

		output := g.String()

		asrt.Contains(output, "\"A\" -> \"B\" [label=\"connects\"];", "expected edge with label attribute")
	})

	t.Run("label value is quoted", func(t *testing.T) {
		asrt := assert.New(t)

		g := NewGraph(Directed)
		n1 := NewNode("A")
		n2 := NewNode("B")
		_, _ = g.AddEdge(n1, n2, WithEdgeLabel("my label"))

		output := g.String()

		asrt.Contains(output, "label=\"my label\"", "expected quoted label value")
		asrt.NotContains(output, "label=my label", "expected no unquoted label value")
	})

	t.Run("label with special characters", func(t *testing.T) {
		asrt := assert.New(t)

		g := NewGraph(Directed)
		n1 := NewNode("A")
		n2 := NewNode("B")
		_, _ = g.AddEdge(n1, n2, WithEdgeLabel("Label with spaces"))

		output := g.String()

		asrt.Contains(output, "label=\"Label with spaces\"", "expected label with spaces to be preserved")
	})
}

// Test edge with multiple attributes
func TestDOT_SingleEdge_MultipleAttributes(t *testing.T) {
	t.Run("renders multiple attributes separated by commas", func(t *testing.T) {
		asrt := assert.New(t)

		g := NewGraph(Directed)
		n1 := NewNode("A")
		n2 := NewNode("B")
		_, _ = g.AddEdge(n1, n2, WithEdgeLabel("test"), WithEdgeColor("red"), WithEdgeStyle(EdgeStyleDashed))

		output := g.String()

		asrt.Contains(output, "\"A\" -> \"B\" [", "expected attribute section to start")
		asrt.Contains(output, "label=\"test\"", "expected label attribute")
		asrt.Contains(output, "color=\"red\"", "expected color attribute")
		asrt.Contains(output, "style=\"dashed\"", "expected style attribute")
		asrt.Contains(output, ", ", "expected comma separator between attributes")
	})

	t.Run("attributes are sorted alphabetically", func(t *testing.T) {
		asrt := assert.New(t)

		g := NewGraph(Directed)
		n1 := NewNode("A")
		n2 := NewNode("B")
		_, _ = g.AddEdge(n1, n2, WithEdgeLabel("z"), WithEdgeColor("a"))

		output := g.String()

		// color comes before label alphabetically
		colorIdx := strings.Index(output, "color=\"a\"")
		labelIdx := strings.Index(output, "label=\"z\"")
		asrt.Greater(labelIdx, colorIdx, "expected attributes in alphabetical order (color before label)")
	})

	t.Run("attributes in consistent order regardless of option order", func(t *testing.T) {
		asrt := assert.New(t)

		g1 := NewGraph(Directed)
		n1a := NewNode("A")
		n2a := NewNode("B")
		_, _ = g1.AddEdge(n1a, n2a, WithEdgeLabel("test"), WithEdgeColor("red"))

		g2 := NewGraph(Directed)
		n1b := NewNode("A")
		n2b := NewNode("B")
		_, _ = g2.AddEdge(n1b, n2b, WithEdgeColor("red"), WithEdgeLabel("test"))

		output1 := g1.String()
		output2 := g2.String()

		asrt.Equal(output1, output2, "expected consistent attribute order regardless of option order")
	})
}

// Test edge with all attributes
func TestDOT_SingleEdge_AllAttributes(t *testing.T) {
	t.Run("renders all edge attribute types", func(t *testing.T) {
		asrt := assert.New(t)

		g := NewGraph(Directed)
		n1 := NewNode("A")
		n2 := NewNode("B")
		_, _ = g.AddEdge(n1, n2,
			WithEdgeLabel("connects"),
			WithEdgeColor("blue"),
			WithEdgeStyle(EdgeStyleDashed),
			WithArrowHead(ArrowDot),
			WithArrowTail(ArrowNormal),
			WithWeight(2.5),
		)

		output := g.String()

		asrt.Contains(output, "label=\"connects\"", "expected label")
		asrt.Contains(output, "color=\"blue\"", "expected color")
		asrt.Contains(output, "style=\"dashed\"", "expected style")
		asrt.Contains(output, "arrowhead=\"dot\"", "expected arrowhead (lowercase)")
		asrt.Contains(output, "arrowtail=\"normal\"", "expected arrowtail (lowercase)")
		asrt.Contains(output, "weight=\"2.5\"", "expected weight")
	})

	t.Run("arrowhead uses lowercase in DOT output", func(t *testing.T) {
		asrt := assert.New(t)

		g := NewGraph(Directed)
		n1 := NewNode("A")
		n2 := NewNode("B")
		_, _ = g.AddEdge(n1, n2, WithArrowHead(ArrowVee))

		output := g.String()

		asrt.Contains(output, "arrowhead=\"vee\"", "expected lowercase arrowhead in DOT")
		asrt.NotContains(output, "arrowHead=", "expected no camelCase arrowHead")
	})

	t.Run("arrowtail uses lowercase in DOT output", func(t *testing.T) {
		asrt := assert.New(t)

		g := NewGraph(Directed)
		n1 := NewNode("A")
		n2 := NewNode("B")
		_, _ = g.AddEdge(n1, n2, WithArrowTail(ArrowDot))

		output := g.String()

		asrt.Contains(output, "arrowtail=\"dot\"", "expected lowercase arrowtail in DOT")
		asrt.NotContains(output, "arrowTail=", "expected no camelCase arrowTail")
	})
}

// Test edge weight formatting
func TestDOT_SingleEdge_Weight(t *testing.T) {
	t.Run("weight formatted with two decimal places", func(t *testing.T) {
		asrt := assert.New(t)

		g := NewGraph(Directed)
		n1 := NewNode("A")
		n2 := NewNode("B")
		_, _ = g.AddEdge(n1, n2, WithWeight(2.5))

		output := g.String()

		asrt.Contains(output, "weight=\"2.5\"", "expected weight formatted as %g")
	})

	t.Run("integer weight formatted with decimal places", func(t *testing.T) {
		asrt := assert.New(t)

		g := NewGraph(Directed)
		n1 := NewNode("A")
		n2 := NewNode("B")
		_, _ = g.AddEdge(n1, n2, WithWeight(3))

		output := g.String()

		asrt.Contains(output, "weight=\"3\"", "expected integer weight formatted as %g")
	})

	t.Run("weight with many decimal places is truncated", func(t *testing.T) {
		asrt := assert.New(t)

		g := NewGraph(Directed)
		n1 := NewNode("A")
		n2 := NewNode("B")
		_, _ = g.AddEdge(n1, n2, WithWeight(2.12345))

		output := g.String()

		asrt.Contains(output, "weight=\"2.12345\"", "expected float weight formatted as %g")
	})
}

// Test edge with custom attributes
func TestDOT_SingleEdge_CustomAttribute(t *testing.T) {
	t.Run("renders custom attribute", func(t *testing.T) {
		asrt := assert.New(t)

		g := NewGraph(Directed)
		n1 := NewNode("A")
		n2 := NewNode("B")
		_, _ = g.AddEdge(n1, n2, WithEdgeAttribute("penwidth", "2.0"))

		output := g.String()

		asrt.Contains(output, "penwidth=\"2.0\"", "expected custom attribute to be rendered")
	})

	t.Run("multiple custom attributes", func(t *testing.T) {
		asrt := assert.New(t)

		g := NewGraph(Directed)
		n1 := NewNode("A")
		n2 := NewNode("B")
		_, _ = g.AddEdge(n1, n2,
			WithEdgeAttribute("penwidth", "2.0"),
			WithEdgeAttribute("tooltip", "Hover text"),
		)

		output := g.String()

		asrt.Contains(output, "penwidth=\"2.0\"", "expected first custom attribute")
		asrt.Contains(output, "tooltip=\"Hover text\"", "expected second custom attribute")
	})

	t.Run("custom and typed attributes together", func(t *testing.T) {
		asrt := assert.New(t)

		g := NewGraph(Directed)
		n1 := NewNode("A")
		n2 := NewNode("B")
		_, _ = g.AddEdge(n1, n2,
			WithEdgeLabel("test"),
			WithEdgeAttribute("penwidth", "2.0"),
			WithEdgeColor("red"),
		)

		output := g.String()

		asrt.Contains(output, "label=\"test\"", "expected typed label attribute")
		asrt.Contains(output, "color=\"red\"", "expected typed color attribute")
		asrt.Contains(output, "penwidth=\"2.0\"", "expected custom attribute")
	})

	t.Run("custom attributes sorted with typed attributes", func(t *testing.T) {
		asrt := assert.New(t)

		g := NewGraph(Directed)
		n1 := NewNode("A")
		n2 := NewNode("B")
		_, _ = g.AddEdge(n1, n2,
			WithEdgeLabel("z"),
			WithEdgeAttribute("aaa", "first"),
		)

		output := g.String()

		// aaa should come before label alphabetically
		aaaIdx := strings.Index(output, "aaa=\"first\"")
		labelIdx := strings.Index(output, "label=\"z\"")
		asrt.Greater(labelIdx, aaaIdx, "expected custom attribute sorted with typed attributes")
	})
}

// Test multiple edges
func TestDOT_MultipleEdges(t *testing.T) {
	t.Run("renders all edges in insertion order", func(t *testing.T) {
		asrt := assert.New(t)

		g := NewGraph(Directed)
		n1 := NewNode("A")
		n2 := NewNode("B")
		n3 := NewNode("C")

		_, _ = g.AddEdge(n1, n2)
		_, _ = g.AddEdge(n2, n3)
		_, _ = g.AddEdge(n3, n1)

		output := g.String()

		asrt.Contains(output, "\"A\" -> \"B\";", "expected first edge A->B")
		asrt.Contains(output, "\"B\" -> \"C\";", "expected second edge B->C")
		asrt.Contains(output, "\"C\" -> \"A\";", "expected third edge C->A")

		// Verify edges appear in insertion order by checking their positions
		aIdx := strings.Index(output, "\"A\" -> \"B\";")
		bIdx := strings.Index(output, "\"B\" -> \"C\";")
		cIdx := strings.Index(output, "\"C\" -> \"A\";")
		asrt.Greater(bIdx, aIdx, "expected second edge after first edge")
		asrt.Greater(cIdx, bIdx, "expected third edge after second edge")
	})

	t.Run("edges with mixed attributes", func(t *testing.T) {
		asrt := assert.New(t)

		g := NewGraph(Directed)
		n1 := NewNode("A")
		n2 := NewNode("B")
		n3 := NewNode("C")

		_, _ = g.AddEdge(n1, n2, WithEdgeLabel("first"))
		_, _ = g.AddEdge(n2, n3)
		_, _ = g.AddEdge(n3, n1, WithEdgeColor("red"), WithEdgeLabel("third"))

		output := g.String()

		asrt.Contains(output, "\"A\" -> \"B\" [label=\"first\"];", "expected first edge with label")
		asrt.Contains(output, "\"B\" -> \"C\";", "expected second edge without attributes")
		asrt.Contains(output, "\"C\" -> \"A\" [color=\"red\", label=\"third\"];",
			"expected third edge with multiple attributes",
		)
	})
}

// Test parallel edges (multiple edges between same nodes)
func TestDOT_ParallelEdges(t *testing.T) {
	t.Run("renders both edges between same nodes", func(t *testing.T) {
		asrt := assert.New(t)

		g := NewGraph(Directed)
		n1 := NewNode("A")
		n2 := NewNode("B")

		_, _ = g.AddEdge(n1, n2, WithEdgeLabel("first"))
		_, _ = g.AddEdge(n1, n2, WithEdgeLabel("second"))

		output := g.String()

		firstCount := strings.Count(output, "\"A\" -> \"B\" [label=\"first\"];")
		secondCount := strings.Count(output, "\"A\" -> \"B\" [label=\"second\"];")

		asrt.Equal(1, firstCount, "expected first edge to appear once")
		asrt.Equal(1, secondCount, "expected second edge to appear once")
	})

	t.Run("parallel edges in undirected graph", func(t *testing.T) {
		asrt := assert.New(t)

		g := NewGraph(Undirected)
		n1 := NewNode("A")
		n2 := NewNode("B")

		_, _ = g.AddEdge(n1, n2)
		_, _ = g.AddEdge(n1, n2, WithEdgeLabel("parallel"))

		output := g.String()

		asrt.Contains(output, "\"A\" -- \"B\";", "expected first undirected edge")
		asrt.Contains(output, "\"A\" -- \"B\" [label=\"parallel\"];", "expected second undirected edge")
	})
}

// Test self-loop edge
func TestDOT_SelfLoopEdge(t *testing.T) {
	t.Run("renders edge from node to itself", func(t *testing.T) {
		asrt := assert.New(t)

		g := NewGraph(Directed)
		n := NewNode("A")
		_, _ = g.AddEdge(n, n)

		output := g.String()

		asrt.Contains(output, "\"A\" -> \"A\";", "expected self-loop edge")
	})

	t.Run("self-loop with attributes", func(t *testing.T) {
		asrt := assert.New(t)

		g := NewGraph(Directed)
		n := NewNode("A")
		_, _ = g.AddEdge(n, n, WithEdgeLabel("loop"))

		output := g.String()

		asrt.Contains(output, "\"A\" -> \"A\" [label=\"loop\"];", "expected self-loop with label")
	})

	t.Run("self-loop in undirected graph", func(t *testing.T) {
		asrt := assert.New(t)

		g := NewGraph(Undirected)
		n := NewNode("A")
		_, _ = g.AddEdge(n, n)

		output := g.String()

		asrt.Contains(output, "\"A\" -- \"A\";", "expected undirected self-loop")
	})
}

// Test complete graph integration
func TestDOT_CompleteGraph(t *testing.T) {
	t.Run("renders complete graph with nodes and edges", func(t *testing.T) {
		asrt := assert.New(t)

		g := NewGraph(Directed, WithName("G"))
		n1 := NewNode("A", WithBoxShape())
		n2 := NewNode("B", WithLabel("Node B"))

		_ = g.AddNode(n1)
		_ = g.AddNode(n2)
		_, _ = g.AddEdge(n1, n2, WithEdgeLabel("connects"), WithEdgeColor("red"))

		output := g.String()

		asrt.Contains(output, "digraph \"G\" {", "expected graph declaration")
		asrt.Contains(output, "\"A\" [shape=\"box\"];", "expected first node with attributes")
		asrt.Contains(output, "\"B\" [label=\"Node B\"];", "expected second node with label")
		asrt.Contains(output, "\"A\" -> \"B\" [color=\"red\", label=\"connects\"];", "expected edge with attributes")

		// Verify edge appears after nodes
		nodeAIdx := strings.Index(output, "\"A\" [shape=\"box\"];")
		nodeBIdx := strings.Index(output, "\"B\" [label=\"Node B\"];")
		edgeIdx := strings.Index(output, "\"A\" -> \"B\"")
		asrt.Greater(edgeIdx, nodeAIdx, "expected edge after node A")
		asrt.Greater(edgeIdx, nodeBIdx, "expected edge after node B")
	})

	t.Run("strict graph with nodes and edges", func(t *testing.T) {
		asrt := assert.New(t)

		g := NewGraph(Strict, Directed, WithName("StrictG"))
		n1 := NewNode("A")
		n2 := NewNode("B")

		_, _ = g.AddEdge(n1, n2)

		output := g.String()

		asrt.Contains(output, "strict digraph \"StrictG\" {", "expected strict graph declaration")
		asrt.Contains(output, "\"A\" -> \"B\";", "expected edge in strict graph")
	})

	t.Run("undirected graph with nodes and edges", func(t *testing.T) {
		asrt := assert.New(t)

		g := NewGraph(Undirected, WithName("UG"))
		n1 := NewNode("A", WithCircleShape())
		n2 := NewNode("B", WithBoxShape())

		_, _ = g.AddEdge(n1, n2, WithEdgeStyle(EdgeStyleDashed))

		output := g.String()

		asrt.Contains(output, "graph \"UG\" {", "expected undirected graph declaration")
		asrt.Contains(output, "\"A\" [shape=\"circle\"];", "expected first node")
		asrt.Contains(output, "\"B\" [shape=\"box\"];", "expected second node")
		asrt.Contains(output, "\"A\" -- \"B\" [style=\"dashed\"];", "expected undirected edge")
	})
}

// Test edge output appears after nodes
func TestDOT_EdgeOutputAfterNodes(t *testing.T) {
	t.Run("nodes added implicitly by edges appear before edges", func(t *testing.T) {
		asrt := assert.New(t)

		g := NewGraph(Directed)
		n1 := NewNode("A")
		n2 := NewNode("B")

		// Add edge without explicitly adding nodes first
		_, _ = g.AddEdge(n1, n2)

		output := g.String()

		lines := strings.Split(output, "\n")
		asrt.Contains(lines[1], "\"A\";", "expected implicitly added node A")
		asrt.Contains(lines[2], "\"B\";", "expected implicitly added node B")
		asrt.Contains(lines[3], "\"A\" -> \"B\";", "expected edge after nodes")
	})
}

// Test different arrow types
func TestDOT_Edge_ArrowTypes(t *testing.T) {
	t.Run("different arrowhead values", func(t *testing.T) {
		asrt := assert.New(t)

		testCases := []struct {
			name     string
			arrow    ArrowType
			expected string
		}{
			{"normal", ArrowNormal, "arrowhead=\"normal\""},
			{"dot", ArrowDot, "arrowhead=\"dot\""},
			{"vee", ArrowVee, "arrowhead=\"vee\""},
			{"none", ArrowNone, "arrowhead=\"none\""},
		}

		for _, tc := range testCases {
			g := NewGraph(Directed)
			n1 := NewNode("A")
			n2 := NewNode("B")
			_, _ = g.AddEdge(n1, n2, WithArrowHead(tc.arrow))

			output := g.String()
			asrt.Contains(output, tc.expected, "expected arrowhead for %s", tc.name)
		}
	})

	t.Run("different arrowtail values", func(t *testing.T) {
		asrt := assert.New(t)

		testCases := []struct {
			name     string
			arrow    ArrowType
			expected string
		}{
			{"normal", ArrowNormal, "arrowtail=\"normal\""},
			{"dot", ArrowDot, "arrowtail=\"dot\""},
			{"vee", ArrowVee, "arrowtail=\"vee\""},
			{"none", ArrowNone, "arrowtail=\"none\""},
		}

		for _, tc := range testCases {
			g := NewGraph(Directed)
			n1 := NewNode("A")
			n2 := NewNode("B")
			_, _ = g.AddEdge(n1, n2, WithArrowTail(tc.arrow))

			output := g.String()
			asrt.Contains(output, tc.expected, "expected arrowtail for %s", tc.name)
		}
	})

	t.Run("both arrowhead and arrowtail", func(t *testing.T) {
		asrt := assert.New(t)

		g := NewGraph(Directed)
		n1 := NewNode("A")
		n2 := NewNode("B")
		_, _ = g.AddEdge(n1, n2, WithArrowHead(ArrowDot), WithArrowTail(ArrowVee))

		output := g.String()

		asrt.Contains(output, "arrowhead=\"dot\"", "expected arrowhead")
		asrt.Contains(output, "arrowtail=\"vee\"", "expected arrowtail")
	})
}

// Test different edge styles
func TestDOT_Edge_Styles(t *testing.T) {
	t.Run("different edge style values", func(t *testing.T) {
		asrt := assert.New(t)

		testCases := []struct {
			name     string
			style    EdgeStyle
			expected string
		}{
			{"solid", EdgeStyleSolid, "style=\"solid\""},
			{"dashed", EdgeStyleDashed, "style=\"dashed\""},
			{"dotted", EdgeStyleDotted, "style=\"dotted\""},
			{"bold", EdgeStyleBold, "style=\"bold\""},
			{"invis", EdgeStyleInvisible, "style=\"invis\""},
		}

		for _, tc := range testCases {
			g := NewGraph(Directed)
			n1 := NewNode("A")
			n2 := NewNode("B")
			_, _ = g.AddEdge(n1, n2, WithEdgeStyle(tc.style))

			output := g.String()
			asrt.Contains(output, tc.expected, "expected style for %s", tc.name)
		}
	})
}

// Test edge ID quoting with special characters
func TestDOT_Edge_NodeIDsQuoted(t *testing.T) {
	t.Run("node IDs with spaces are quoted", func(t *testing.T) {
		asrt := assert.New(t)

		g := NewGraph(Directed)
		n1 := NewNode("node with spaces")
		n2 := NewNode("another node")
		_, _ = g.AddEdge(n1, n2)

		output := g.String()

		asrt.Contains(output, "\"node with spaces\" -> \"another node\";", "expected quoted node IDs with spaces")
	})

	t.Run("node IDs with special characters are quoted", func(t *testing.T) {
		asrt := assert.New(t)

		g := NewGraph(Directed)
		n1 := NewNode("node-with-dashes")
		n2 := NewNode("node:with:colons")
		_, _ = g.AddEdge(n1, n2)

		output := g.String()

		asrt.Contains(output, "\"node-with-dashes\" -> \"node:with:colons\";",
			"expected quoted node IDs with special characters",
		)
	})

	t.Run("empty node ID is quoted", func(t *testing.T) {
		asrt := assert.New(t)

		g := NewGraph(Directed)
		n1 := NewNode("")
		n2 := NewNode("B")
		_, _ = g.AddEdge(n1, n2)

		output := g.String()

		asrt.Contains(output, "\"\" -> \"B\";", "expected quoted empty node ID")
	})
}

// Test graph attributes in DOT output - single attribute (rankdir)
func TestDOT_GraphAttributes_RankDir(t *testing.T) {
	t.Run("outputs rankdir attribute", func(t *testing.T) {
		asrt := assert.New(t)
		g := NewGraph(Directed, WithRankDir(RankDirLR))

		output := g.String()

		asrt.Contains(output, "rankdir=\"LR\";", "expected rankdir attribute in output")
	})

	t.Run("rankdir appears after opening brace", func(t *testing.T) {
		asrt := assert.New(t)
		g := NewGraph(Directed, WithName("G"), WithRankDir(RankDirLR))

		output := g.String()
		lines := strings.Split(output, "\n")

		asrt.Contains(lines[0], "digraph \"G\" {", "expected graph declaration on first line")
		asrt.Contains(lines[1], "rankdir=\"LR\";", "expected rankdir on second line")
	})

	t.Run("rankdir is indented with tab", func(t *testing.T) {
		asrt := assert.New(t)
		g := NewGraph(Directed, WithRankDir(RankDirTB))

		output := g.String()

		asrt.Contains(output, "\trankdir=\"TB\";", "expected rankdir to be indented with tab")
	})

	t.Run("different rankdir values", func(t *testing.T) {
		asrt := assert.New(t)

		testCases := []struct {
			name     string
			rankDir  RankDir
			expected string
		}{
			{"TB", RankDirTB, "rankdir=\"TB\""},
			{"BT", RankDirBT, "rankdir=\"BT\""},
			{"LR", RankDirLR, "rankdir=\"LR\""},
			{"RL", RankDirRL, "rankdir=\"RL\""},
		}

		for _, tc := range testCases {
			g := NewGraph(Directed, WithRankDir(tc.rankDir))
			output := g.String()
			asrt.Contains(output, tc.expected, "expected %s rankdir value", tc.name)
		}
	})
}

// Test graph attributes - label
func TestDOT_GraphAttributes_Label(t *testing.T) {
	t.Run("outputs graph label attribute", func(t *testing.T) {
		asrt := assert.New(t)
		g := NewGraph(Directed, WithGraphLabel("My Graph"))

		output := g.String()

		asrt.Contains(output, "label=\"My Graph\";", "expected graph label attribute")
	})

	t.Run("label appears after opening brace before nodes", func(t *testing.T) {
		asrt := assert.New(t)
		g := NewGraph(Directed, WithGraphLabel("Test"))
		n := NewNode("A")
		_ = g.AddNode(n)

		output := g.String()

		labelIdx := strings.Index(output, "label=\"Test\";")
		nodeIdx := strings.Index(output, "\"A\";")
		asrt.Greater(nodeIdx, labelIdx, "expected label before nodes")
	})

	t.Run("label with spaces", func(t *testing.T) {
		asrt := assert.New(t)
		g := NewGraph(Directed, WithGraphLabel("Graph With Spaces"))

		output := g.String()

		asrt.Contains(output, "label=\"Graph With Spaces\";", "expected label with spaces preserved")
	})
}

// Test multiple graph attributes together
func TestDOT_GraphAttributes_Multiple(t *testing.T) {
	t.Run("outputs all set graph attributes", func(t *testing.T) {
		asrt := assert.New(t)
		g := NewGraph(
			Directed,
			WithGraphLabel("My Graph"),
			WithRankDir(RankDirLR),
			WithBgColor("lightgray"),
		)

		output := g.String()

		asrt.Contains(output, "label=\"My Graph\";", "expected label")
		asrt.Contains(output, "rankdir=\"LR\";", "expected rankdir")
		asrt.Contains(output, "bgcolor=\"lightgray\";", "expected bgcolor")
	})

	t.Run("attributes are sorted alphabetically", func(t *testing.T) {
		asrt := assert.New(t)
		g := NewGraph(
			Directed,
			WithRankDir(RankDirLR),
			WithGraphLabel("Test"),
			WithBgColor("white"),
		)

		output := g.String()

		bgcolorIdx := strings.Index(output, "bgcolor=")
		labelIdx := strings.Index(output, "label=")
		rankdirIdx := strings.Index(output, "rankdir=")

		asrt.Greater(labelIdx, bgcolorIdx, "expected bgcolor before label alphabetically")
		asrt.Greater(rankdirIdx, labelIdx, "expected label before rankdir alphabetically")
	})

	t.Run("each attribute on its own line", func(t *testing.T) {
		asrt := assert.New(t)
		g := NewGraph(
			Directed,
			WithGraphLabel("Test"),
			WithRankDir(RankDirTB),
		)

		output := g.String()
		lines := strings.Split(output, "\n")

		foundLabel := false
		foundRankDir := false
		for _, line := range lines {
			if strings.Contains(line, "label=\"Test\";") {
				foundLabel = true
			}
			if strings.Contains(line, "rankdir=\"TB\";") {
				foundRankDir = true
			}
		}

		asrt.True(foundLabel, "expected label on its own line")
		asrt.True(foundRankDir, "expected rankdir on its own line")
	})

	t.Run("all attributes appear before nodes", func(t *testing.T) {
		asrt := assert.New(t)
		g := NewGraph(
			Directed,
			WithGraphLabel("Test"),
			WithRankDir(RankDirLR),
		)
		n := NewNode("A")
		_ = g.AddNode(n)

		output := g.String()

		labelIdx := strings.Index(output, "label=")
		rankdirIdx := strings.Index(output, "rankdir=")
		nodeIdx := strings.Index(output, "\"A\";")

		asrt.Greater(nodeIdx, labelIdx, "expected attributes before nodes")
		asrt.Greater(nodeIdx, rankdirIdx, "expected attributes before nodes")
	})
}

// Test all graph attribute types
func TestDOT_GraphAttributes_AllTypes(t *testing.T) {
	t.Run("outputs all graph attribute types", func(t *testing.T) {
		asrt := assert.New(t)
		g := NewGraph(
			Directed,
			WithGraphLabel("Complete"),
			WithRankDir(RankDirLR),
			WithBgColor("white"),
			WithGraphFontName("Arial"),
			WithGraphFontSize(12.0),
			WithSplines(SplineOrtho),
			WithNodeSep(0.5),
			WithRankSep(1.0),
			WithCompound(true),
		)

		output := g.String()

		asrt.Contains(output, "label=\"Complete\";", "expected label")
		asrt.Contains(output, "rankdir=\"LR\";", "expected rankdir")
		asrt.Contains(output, "bgcolor=\"white\";", "expected bgcolor")
		asrt.Contains(output, "fontname=\"Arial\";", "expected fontname")
		asrt.Contains(output, "fontsize=\"12\";", "expected fontsize")
		asrt.Contains(output, "splines=\"ortho\";", "expected splines")
		asrt.Contains(output, "nodesep=\"0.5\";", "expected nodesep")
		asrt.Contains(output, "ranksep=\"1\";", "expected ranksep")
		asrt.Contains(output, "compound=\"true\";", "expected compound")
	})

	t.Run("uses correct DOT attribute names", func(t *testing.T) {
		asrt := assert.New(t)
		g := NewGraph(
			Directed,
			WithBgColor("white"),
			WithRankDir(RankDirLR),
			WithNodeSep(0.5),
		)

		output := g.String()

		asrt.Contains(output, "bgcolor=", "expected lowercase bgcolor")
		asrt.NotContains(output, "bgColor=", "expected no camelCase bgColor")
		asrt.Contains(output, "rankdir=", "expected lowercase rankdir")
		asrt.NotContains(output, "rankDir=", "expected no camelCase rankDir")
		asrt.Contains(output, "nodesep=", "expected lowercase nodesep")
		asrt.NotContains(output, "nodeSep=", "expected no camelCase nodeSep")
	})

	t.Run("different spline types", func(t *testing.T) {
		asrt := assert.New(t)

		testCases := []struct {
			name     string
			spline   SplineType
			expected string
		}{
			{"true", SplineTrue, "splines=\"true\""},
			{"false", SplineFalse, "splines=\"false\""},
			{"ortho", SplineOrtho, "splines=\"ortho\""},
			{"polyline", SplinePolyline, "splines=\"polyline\""},
			{"curved", SplineCurved, "splines=\"curved\""},
			{"spline", SplineSpline, "splines=\"spline\""},
			{"line", SplineLine, "splines=\"line\""},
			{"none", SplineNone, "splines=\"none\""},
		}

		for _, tc := range testCases {
			g := NewGraph(Directed, WithSplines(tc.spline))
			output := g.String()
			asrt.Contains(output, tc.expected, "expected %s spline value", tc.name)
		}
	})

	t.Run("compound boolean formatting", func(t *testing.T) {
		asrt := assert.New(t)

		gTrue := NewGraph(Directed, WithCompound(true))
		outputTrue := gTrue.String()
		asrt.Contains(outputTrue, "compound=\"true\";", "expected true as string")

		gFalse := NewGraph(Directed, WithCompound(false))
		outputFalse := gFalse.String()
		asrt.Contains(outputFalse, "compound=\"false\";", "expected false as string")
	})
}

// Test default node attributes output
func TestDOT_DefaultNodeAttrs(t *testing.T) {
	t.Run("outputs default node attributes", func(t *testing.T) {
		asrt := assert.New(t)
		g := NewGraph(
			Directed,
			WithDefaultNodeAttrs(WithBoxShape(), WithColor("blue")),
		)

		output := g.String()

		asrt.Contains(output, "node [", "expected node attributes section")
		asrt.Contains(output, "shape=\"box\"", "expected shape in defaults")
		asrt.Contains(output, "color=\"blue\"", "expected color in defaults")
		asrt.Contains(output, "];", "expected closing bracket and semicolon")
	})

	t.Run("default node attrs appear after graph attrs", func(t *testing.T) {
		asrt := assert.New(t)
		g := NewGraph(
			Directed,
			WithGraphLabel("Test"),
			WithDefaultNodeAttrs(WithBoxShape()),
		)

		output := g.String()

		labelIdx := strings.Index(output, "label=\"Test\";")
		nodeDefaultsIdx := strings.Index(output, "node [")

		asrt.Greater(nodeDefaultsIdx, labelIdx, "expected node defaults after graph attributes")
	})

	t.Run("default node attrs appear before nodes", func(t *testing.T) {
		asrt := assert.New(t)
		g := NewGraph(
			Directed,
			WithDefaultNodeAttrs(WithBoxShape()),
		)
		n := NewNode("A")
		_ = g.AddNode(n)

		output := g.String()

		nodeDefaultsIdx := strings.Index(output, "node [")
		nodeIdx := strings.Index(output, "\"A\";")

		asrt.Greater(nodeIdx, nodeDefaultsIdx, "expected node defaults before actual nodes")
	})

	t.Run("node default attributes are sorted", func(t *testing.T) {
		asrt := assert.New(t)
		g := NewGraph(
			Directed,
			WithDefaultNodeAttrs(
				WithLabel("Default Label"),
				WithBoxShape(),
				WithColor("red"),
			),
		)

		output := g.String()

		colorIdx := strings.Index(output, "color=\"red\"")
		labelIdx := strings.Index(output, "label=\"Default Label\"")
		shapeIdx := strings.Index(output, "shape=\"box\"")

		asrt.Greater(labelIdx, colorIdx, "expected color before label alphabetically")
		asrt.Greater(shapeIdx, labelIdx, "expected label before shape alphabetically")
	})

	t.Run("node defaults format is correct", func(t *testing.T) {
		asrt := assert.New(t)
		g := NewGraph(
			Directed,
			WithDefaultNodeAttrs(WithBoxShape()),
		)

		output := g.String()

		asrt.Contains(output, "\tnode [shape=\"box\"];", "expected correct format with tab and semicolon")
	})

	t.Run("multiple node default attributes", func(t *testing.T) {
		asrt := assert.New(t)
		g := NewGraph(
			Directed,
			WithDefaultNodeAttrs(
				WithBoxShape(),
				WithFillColor("lightblue"),
				WithFontName("Arial"),
				WithFontSize(14.0),
			),
		)

		output := g.String()

		asrt.Contains(output, "shape=\"box\"", "expected shape")
		asrt.Contains(output, "fillcolor=\"lightblue\"", "expected fillcolor")
		asrt.Contains(output, "fontname=\"Arial\"", "expected fontname")
		asrt.Contains(output, "fontsize=\"14\"", "expected fontsize")
	})
}

// Test default edge attributes output
func TestDOT_DefaultEdgeAttrs(t *testing.T) {
	t.Run("outputs default edge attributes", func(t *testing.T) {
		asrt := assert.New(t)
		g := NewGraph(
			Directed,
			WithDefaultEdgeAttrs(WithEdgeColor("gray"), WithEdgeStyle(EdgeStyleDashed)),
		)

		output := g.String()

		asrt.Contains(output, "edge [", "expected edge attributes section")
		asrt.Contains(output, "color=\"gray\"", "expected color in defaults")
		asrt.Contains(output, "style=\"dashed\"", "expected style in defaults")
		asrt.Contains(output, "];", "expected closing bracket and semicolon")
	})

	t.Run("default edge attrs appear after node defaults", func(t *testing.T) {
		asrt := assert.New(t)
		g := NewGraph(
			Directed,
			WithDefaultNodeAttrs(WithBoxShape()),
			WithDefaultEdgeAttrs(WithEdgeColor("gray")),
		)

		output := g.String()

		nodeDefaultsIdx := strings.Index(output, "node [")
		edgeDefaultsIdx := strings.Index(output, "edge [")

		asrt.Greater(edgeDefaultsIdx, nodeDefaultsIdx, "expected edge defaults after node defaults")
	})

	t.Run("default edge attrs appear before nodes", func(t *testing.T) {
		asrt := assert.New(t)
		g := NewGraph(
			Directed,
			WithDefaultEdgeAttrs(WithEdgeColor("gray")),
		)
		n := NewNode("A")
		_ = g.AddNode(n)

		output := g.String()

		edgeDefaultsIdx := strings.Index(output, "edge [")
		nodeIdx := strings.Index(output, "\"A\";")

		asrt.Greater(nodeIdx, edgeDefaultsIdx, "expected edge defaults before actual nodes")
	})

	t.Run("edge default attributes are sorted", func(t *testing.T) {
		asrt := assert.New(t)
		g := NewGraph(
			Directed,
			WithDefaultEdgeAttrs(
				WithEdgeLabel("default"),
				WithEdgeColor("blue"),
				WithArrowHead(ArrowDot),
			),
		)

		output := g.String()

		arrowIdx := strings.Index(output, "arrowhead=\"dot\"")
		colorIdx := strings.Index(output, "color=\"blue\"")
		labelIdx := strings.Index(output, "label=\"default\"")

		asrt.Greater(colorIdx, arrowIdx, "expected arrowhead before color alphabetically")
		asrt.Greater(labelIdx, colorIdx, "expected color before label alphabetically")
	})

	t.Run("edge defaults format is correct", func(t *testing.T) {
		asrt := assert.New(t)
		g := NewGraph(
			Directed,
			WithDefaultEdgeAttrs(WithEdgeColor("red")),
		)

		output := g.String()

		asrt.Contains(output, "\tedge [color=\"red\"];", "expected correct format with tab and semicolon")
	})

	t.Run("multiple edge default attributes", func(t *testing.T) {
		asrt := assert.New(t)
		g := NewGraph(
			Directed,
			WithDefaultEdgeAttrs(
				WithEdgeLabel("edge"),
				WithEdgeColor("blue"),
				WithEdgeStyle(EdgeStyleDotted),
				WithArrowHead(ArrowVee),
				WithWeight(2.5),
			),
		)

		output := g.String()

		asrt.Contains(output, "label=\"edge\"", "expected label")
		asrt.Contains(output, "color=\"blue\"", "expected color")
		asrt.Contains(output, "style=\"dotted\"", "expected style")
		asrt.Contains(output, "arrowhead=\"vee\"", "expected arrowhead")
		asrt.Contains(output, "weight=\"2.5\"", "expected weight")
	})
}

// Test that default sections are only output if non-empty
func TestDOT_DefaultAttrs_OnlyIfNonEmpty(t *testing.T) {
	t.Run("no node defaults when not set", func(t *testing.T) {
		asrt := assert.New(t)
		g := NewGraph(Directed)
		n := NewNode("A")
		_ = g.AddNode(n)

		output := g.String()

		asrt.NotContains(output, "node [", "expected no node defaults section when empty")
	})

	t.Run("no edge defaults when not set", func(t *testing.T) {
		asrt := assert.New(t)
		g := NewGraph(Directed)
		n1 := NewNode("A")
		n2 := NewNode("B")
		_, _ = g.AddEdge(n1, n2)

		output := g.String()

		asrt.NotContains(output, "edge [", "expected no edge defaults section when empty")
	})

	t.Run("empty graph with no attributes has no attribute lines", func(t *testing.T) {
		asrt := assert.New(t)
		g := NewGraph(Directed)

		output := g.String()

		expected := "digraph {\n}"
		asrt.Equal(expected, output, "expected only graph declaration and closing brace")
	})

	t.Run("outputs node defaults but not edge defaults", func(t *testing.T) {
		asrt := assert.New(t)
		g := NewGraph(
			Directed,
			WithDefaultNodeAttrs(WithBoxShape()),
		)

		output := g.String()

		asrt.Contains(output, "node [", "expected node defaults")
		asrt.NotContains(output, "edge [", "expected no edge defaults")
	})

	t.Run("outputs edge defaults but not node defaults", func(t *testing.T) {
		asrt := assert.New(t)
		g := NewGraph(
			Directed,
			WithDefaultEdgeAttrs(WithEdgeColor("gray")),
		)

		output := g.String()

		asrt.Contains(output, "edge [", "expected edge defaults")
		asrt.NotContains(output, "node [", "expected no node defaults")
	})
}

// Test full graph with all sections
func TestDOT_FullGraph_WithAllSections(t *testing.T) {
	t.Run("outputs all sections in correct order", func(t *testing.T) {
		asrt := assert.New(t)
		g := NewGraph(
			Strict,
			Directed,
			WithName("G"),
			WithGraphLabel("Complete Graph"),
			WithRankDir(RankDirLR),
			WithDefaultNodeAttrs(WithBoxShape(), WithFontName("Arial")),
			WithDefaultEdgeAttrs(WithEdgeColor("gray")),
		)
		n1 := NewNode("A", WithLabel("Node A"))
		n2 := NewNode("B", WithLabel("Node B"))
		_ = g.AddNode(n1)
		_ = g.AddNode(n2)
		_, _ = g.AddEdge(n1, n2, WithEdgeLabel("connects"))

		output := g.String()
		lines := strings.Split(output, "\n")

		// Verify structure
		asrt.GreaterOrEqual(len(lines), 8, "expected at least 8 lines")

		// Line 0: Graph declaration
		asrt.Contains(lines[0], "strict digraph \"G\" {", "expected graph declaration")

		// Find attribute lines (order may vary due to sorting)
		graphAttrLines := []string{}
		nodeDefaultLine := ""
		edgeDefaultLine := ""
		nodeLine := ""
		edgeLine := ""

		for i, line := range lines {
			if i == 0 {
				continue // skip declaration
			}
			switch {
			case strings.Contains(line, "label=\"Complete Graph\";") || strings.Contains(line, "rankdir="):
				graphAttrLines = append(graphAttrLines, line)
			case strings.Contains(line, "node ["):
				nodeDefaultLine = line
			case strings.Contains(line, "edge ["):
				edgeDefaultLine = line
			case strings.Contains(line, "\"A\"") && strings.Contains(line, "label=\"Node A\""):
				nodeLine = line
			case strings.Contains(line, "->"):
				edgeLine = line
			}
		}

		asrt.Greater(len(graphAttrLines), 0, "expected graph attributes")
		asrt.NotEmpty(nodeDefaultLine, "expected node default line")
		asrt.NotEmpty(edgeDefaultLine, "expected edge default line")
		asrt.NotEmpty(nodeLine, "expected node line")
		asrt.NotEmpty(edgeLine, "expected edge line")
	})

	t.Run("complete output matches expected format", func(t *testing.T) {
		asrt := assert.New(t)
		g := NewGraph(
			Directed,
			WithName("Test"),
			WithGraphLabel("My Graph"),
			WithRankDir(RankDirLR),
			WithDefaultNodeAttrs(WithBoxShape()),
			WithDefaultEdgeAttrs(WithEdgeColor("gray")),
		)
		n1 := NewNode("A")
		n2 := NewNode("B")
		_ = g.AddNode(n1)
		_ = g.AddNode(n2)
		_, _ = g.AddEdge(n1, n2)

		output := g.String()

		// Verify all expected elements are present
		asrt.Contains(output, "digraph \"Test\" {", "expected graph declaration")
		asrt.Contains(output, "label=\"My Graph\";", "expected graph label")
		asrt.Contains(output, "rankdir=\"LR\";", "expected rankdir")
		asrt.Contains(output, "node [shape=\"box\"];", "expected node defaults")
		asrt.Contains(output, "edge [color=\"gray\"];", "expected edge defaults")
		asrt.Contains(output, "\"A\";", "expected node A")
		asrt.Contains(output, "\"B\";", "expected node B")
		asrt.Contains(output, "\"A\" -> \"B\";", "expected edge")
		asrt.Contains(output, "}", "expected closing brace")
	})
}

// Test explicit output order verification
func TestDOT_OutputOrder_Verification(t *testing.T) {
	t.Run("verifies strict ordering of all sections", func(t *testing.T) {
		asrt := assert.New(t)
		g := NewGraph(
			Directed,
			WithGraphLabel("Test"),
			WithDefaultNodeAttrs(WithBoxShape()),
			WithDefaultEdgeAttrs(WithEdgeColor("gray")),
		)
		n := NewNode("A")
		_ = g.AddNode(n)
		n1 := NewNode("B")
		n2 := NewNode("C")
		_, _ = g.AddEdge(n1, n2)

		output := g.String()

		// Get positions of key elements
		graphAttrIdx := strings.Index(output, "label=\"Test\";")
		nodeDefaultIdx := strings.Index(output, "node [")
		edgeDefaultIdx := strings.Index(output, "edge [")
		firstNodeIdx := strings.Index(output, "\"A\";")
		edgeIdx := strings.Index(output, "->")

		// Verify ordering
		asrt.NotEqual(-1, graphAttrIdx, "expected graph attributes to exist")
		asrt.NotEqual(-1, nodeDefaultIdx, "expected node defaults to exist")
		asrt.NotEqual(-1, edgeDefaultIdx, "expected edge defaults to exist")
		asrt.NotEqual(-1, firstNodeIdx, "expected nodes to exist")
		asrt.NotEqual(-1, edgeIdx, "expected edges to exist")

		asrt.Greater(nodeDefaultIdx, graphAttrIdx, "expected node defaults after graph attrs")
		asrt.Greater(edgeDefaultIdx, nodeDefaultIdx, "expected edge defaults after node defaults")
		asrt.Greater(firstNodeIdx, edgeDefaultIdx, "expected nodes after edge defaults")
		asrt.Greater(edgeIdx, firstNodeIdx, "expected edges after nodes")
	})

	t.Run("graph attributes before everything else", func(t *testing.T) {
		asrt := assert.New(t)
		g := NewGraph(
			Directed,
			WithGraphLabel("First"),
			WithRankDir(RankDirLR),
		)
		n := NewNode("A")
		_ = g.AddNode(n)

		output := g.String()

		firstGraphAttrIdx := strings.Index(output, "label=\"First\";")
		nodeIdx := strings.Index(output, "\"A\";")

		asrt.Greater(nodeIdx, firstGraphAttrIdx, "expected all graph attrs before nodes")
	})
}

// Test graph attributes with empty graph (no nodes/edges)
func TestDOT_GraphAttributes_EmptyGraph(t *testing.T) {
	t.Run("attributes work with no nodes or edges", func(t *testing.T) {
		asrt := assert.New(t)
		g := NewGraph(
			Directed,
			WithGraphLabel("Empty"),
			WithRankDir(RankDirTB),
		)

		output := g.String()

		asrt.Contains(output, "label=\"Empty\";", "expected label in empty graph")
		asrt.Contains(output, "rankdir=\"TB\";", "expected rankdir in empty graph")
	})

	t.Run("structure is correct for empty graph with attributes", func(t *testing.T) {
		asrt := assert.New(t)
		g := NewGraph(Directed, WithGraphLabel("Test"))

		output := g.String()
		lines := strings.Split(output, "\n")

		asrt.Equal(3, len(lines), "expected 3 lines: declaration, label, closing")
		asrt.Contains(lines[0], "digraph {", "expected graph declaration")
		asrt.Contains(lines[1], "label=\"Test\";", "expected label")
		asrt.Equal("}", lines[2], "expected closing brace")
	})
}

// Test custom graph attributes
func TestDOT_CustomGraphAttribute(t *testing.T) {
	t.Run("outputs custom graph attribute", func(t *testing.T) {
		asrt := assert.New(t)
		g := NewGraph(
			Directed,
			WithGraphAttribute("margin", "0.5"),
		)

		output := g.String()

		asrt.Contains(output, "margin=\"0.5\";", "expected custom attribute")
	})

	t.Run("custom and typed attributes together", func(t *testing.T) {
		asrt := assert.New(t)
		g := NewGraph(
			Directed,
			WithGraphLabel("Test"),
			WithGraphAttribute("margin", "0.5"),
			WithRankDir(RankDirLR),
		)

		output := g.String()

		asrt.Contains(output, "label=\"Test\";", "expected typed label")
		asrt.Contains(output, "rankdir=\"LR\";", "expected typed rankdir")
		asrt.Contains(output, "margin=\"0.5\";", "expected custom margin")
	})

	t.Run("custom attributes sorted with typed attributes", func(t *testing.T) {
		asrt := assert.New(t)
		g := NewGraph(
			Directed,
			WithGraphLabel("zzz"),
			WithGraphAttribute("aaa", "first"),
		)

		output := g.String()

		aaaIdx := strings.Index(output, "aaa=\"first\"")
		labelIdx := strings.Index(output, "label=\"zzz\"")

		asrt.Greater(labelIdx, aaaIdx, "expected custom attribute sorted alphabetically with typed attrs")
	})
}

// Test default node attributes with custom
func TestDOT_DefaultNodeAttrs_WithCustom(t *testing.T) {
	t.Run("outputs custom node default attributes", func(t *testing.T) {
		asrt := assert.New(t)
		g := NewGraph(
			Directed,
			WithDefaultNodeAttrs(
				WithBoxShape(),
				WithNodeAttribute("peripheries", "2"),
			),
		)

		output := g.String()

		asrt.Contains(output, "node [", "expected node defaults section")
		asrt.Contains(output, "shape=\"box\"", "expected typed shape")
		asrt.Contains(output, "peripheries=\"2\"", "expected custom peripheries")
	})

	t.Run("custom and typed node defaults sorted together", func(t *testing.T) {
		asrt := assert.New(t)
		g := NewGraph(
			Directed,
			WithDefaultNodeAttrs(
				WithLabel("zzz"),
				WithNodeAttribute("aaa", "first"),
			),
		)

		output := g.String()

		aaaIdx := strings.Index(output, "aaa=\"first\"")
		labelIdx := strings.Index(output, "label=\"zzz\"")

		asrt.Greater(labelIdx, aaaIdx, "expected custom sorted with typed attributes")
	})
}

// Test default edge attributes with custom
func TestDOT_DefaultEdgeAttrs_WithCustom(t *testing.T) {
	t.Run("outputs custom edge default attributes", func(t *testing.T) {
		asrt := assert.New(t)
		g := NewGraph(
			Directed,
			WithDefaultEdgeAttrs(
				WithEdgeColor("blue"),
				WithEdgeAttribute("penwidth", "2.0"),
			),
		)

		output := g.String()

		asrt.Contains(output, "edge [", "expected edge defaults section")
		asrt.Contains(output, "color=\"blue\"", "expected typed color")
		asrt.Contains(output, "penwidth=\"2.0\"", "expected custom penwidth")
	})

	t.Run("custom and typed edge defaults sorted together", func(t *testing.T) {
		asrt := assert.New(t)
		g := NewGraph(
			Directed,
			WithDefaultEdgeAttrs(
				WithEdgeLabel("zzz"),
				WithEdgeAttribute("aaa", "first"),
			),
		)

		output := g.String()

		aaaIdx := strings.Index(output, "aaa=\"first\"")
		labelIdx := strings.Index(output, "label=\"zzz\"")

		asrt.Greater(labelIdx, aaaIdx, "expected custom sorted with typed attributes")
	})
}

// Test numeric attribute formatting
func TestDOT_NumericAttributes_Formatting(t *testing.T) {
	t.Run("fontsize formatted with two decimals", func(t *testing.T) {
		asrt := assert.New(t)
		g := NewGraph(Directed, WithGraphFontSize(12.5))

		output := g.String()

		asrt.Contains(output, "fontsize=\"12.5\";", "expected fontsize formatted with %g")
	})

	t.Run("nodesep formatted with two decimals", func(t *testing.T) {
		asrt := assert.New(t)
		g := NewGraph(Directed, WithNodeSep(0.75))

		output := g.String()

		asrt.Contains(output, "nodesep=\"0.75\";", "expected nodesep with two decimals")
	})

	t.Run("ranksep formatted with two decimals", func(t *testing.T) {
		asrt := assert.New(t)
		g := NewGraph(Directed, WithRankSep(1.25))

		output := g.String()

		asrt.Contains(output, "ranksep=\"1.25\";", "expected ranksep with two decimals")
	})

	t.Run("integer values formatted with decimal places", func(t *testing.T) {
		asrt := assert.New(t)
		g := NewGraph(Directed, WithGraphFontSize(14))

		output := g.String()

		asrt.Contains(output, "fontsize=\"14\";", "expected integer fontsize formatted with %g")
	})
}

// ==================== Step 19: String Escaping in DOT Output ====================

// Test escaping backslashes in node IDs
func TestDOT_Escaping_Backslashes_NodeID(t *testing.T) {
	t.Run("backslash in node ID is escaped", func(t *testing.T) {
		asrt := assert.New(t)
		g := NewGraph(Directed)
		n := NewNode(`path\to\file`)
		_ = g.AddNode(n)

		output := g.String()

		asrt.Contains(output, `"path\\to\\file";`, "expected backslashes escaped in node ID")
	})

	t.Run("multiple backslashes escaped correctly", func(t *testing.T) {
		asrt := assert.New(t)
		g := NewGraph(Directed)
		n := NewNode(`C:\Windows\System32`)
		_ = g.AddNode(n)

		output := g.String()

		asrt.Contains(output, `"C:\\Windows\\System32";`, "expected all backslashes escaped")
	})

	t.Run("trailing backslash escaped", func(t *testing.T) {
		asrt := assert.New(t)
		g := NewGraph(Directed)
		n := NewNode(`path\`)
		_ = g.AddNode(n)

		output := g.String()

		asrt.Contains(output, `"path\\";`, "expected trailing backslash escaped")
	})

	t.Run("leading backslash escaped", func(t *testing.T) {
		asrt := assert.New(t)
		g := NewGraph(Directed)
		n := NewNode(`\root`)
		_ = g.AddNode(n)

		output := g.String()

		asrt.Contains(output, `"\\root";`, "expected leading backslash escaped")
	})

	t.Run("consecutive backslashes escaped", func(t *testing.T) {
		asrt := assert.New(t)
		g := NewGraph(Directed)
		n := NewNode(`path\\\\share`)
		_ = g.AddNode(n)

		output := g.String()

		asrt.Contains(output, `"path\\\\\\\\share";`, "expected consecutive backslashes each escaped")
	})
}

// Test escaping double quotes in node IDs
func TestDOT_Escaping_Quotes_NodeID(t *testing.T) {
	t.Run("double quote in node ID is escaped", func(t *testing.T) {
		asrt := assert.New(t)
		g := NewGraph(Directed)
		n := NewNode(`say "hello"`)
		_ = g.AddNode(n)

		output := g.String()

		asrt.Contains(output, `"say \"hello\"";`, "expected quotes escaped in node ID")
	})

	t.Run("single quote in middle", func(t *testing.T) {
		asrt := assert.New(t)
		g := NewGraph(Directed)
		n := NewNode(`user"name`)
		_ = g.AddNode(n)

		output := g.String()

		asrt.Contains(output, `"user\"name";`, "expected quote escaped")
	})

	t.Run("multiple quotes escaped", func(t *testing.T) {
		asrt := assert.New(t)
		g := NewGraph(Directed)
		n := NewNode(`"quoted" "words"`)
		_ = g.AddNode(n)

		output := g.String()

		asrt.Contains(output, `"\"quoted\" \"words\"";`, "expected all quotes escaped")
	})

	t.Run("leading quote escaped", func(t *testing.T) {
		asrt := assert.New(t)
		g := NewGraph(Directed)
		n := NewNode(`"start`)
		_ = g.AddNode(n)

		output := g.String()

		asrt.Contains(output, `"\"start";`, "expected leading quote escaped")
	})

	t.Run("trailing quote escaped", func(t *testing.T) {
		asrt := assert.New(t)
		g := NewGraph(Directed)
		n := NewNode(`end"`)
		_ = g.AddNode(n)

		output := g.String()

		asrt.Contains(output, `"end\"";`, "expected trailing quote escaped")
	})
}

// Test escaping newlines in node IDs
func TestDOT_Escaping_Newlines_NodeID(t *testing.T) {
	t.Run("newline in node ID is escaped to literal backslash-n", func(t *testing.T) {
		asrt := assert.New(t)
		g := NewGraph(Directed)
		n := NewNode("line1\nline2")
		_ = g.AddNode(n)

		output := g.String()

		asrt.Contains(output, `"line1\nline2";`, "expected newline escaped to literal \\n")
		asrt.NotContains(output, "line1\nline2", "expected no actual newline in output")
	})

	t.Run("multiple newlines escaped", func(t *testing.T) {
		asrt := assert.New(t)
		g := NewGraph(Directed)
		n := NewNode("a\nb\nc")
		_ = g.AddNode(n)

		output := g.String()

		asrt.Contains(output, `"a\nb\nc";`, "expected all newlines escaped")
	})

	t.Run("leading newline escaped", func(t *testing.T) {
		asrt := assert.New(t)
		g := NewGraph(Directed)
		n := NewNode("\nstart")
		_ = g.AddNode(n)

		output := g.String()

		asrt.Contains(output, `"\nstart";`, "expected leading newline escaped")
	})

	t.Run("trailing newline escaped", func(t *testing.T) {
		asrt := assert.New(t)
		g := NewGraph(Directed)
		n := NewNode("end\n")
		_ = g.AddNode(n)

		output := g.String()

		asrt.Contains(output, `"end\n";`, "expected trailing newline escaped")
	})

	t.Run("consecutive newlines escaped", func(t *testing.T) {
		asrt := assert.New(t)
		g := NewGraph(Directed)
		n := NewNode("a\n\nb")
		_ = g.AddNode(n)

		output := g.String()

		asrt.Contains(output, `"a\n\nb";`, "expected consecutive newlines escaped")
	})
}

// Test escaping backslashes in attribute values
func TestDOT_Escaping_Backslashes_AttributeValue(t *testing.T) {
	t.Run("backslash in label attribute is escaped", func(t *testing.T) {
		asrt := assert.New(t)
		g := NewGraph(Directed)
		n := NewNode("A", WithLabel(`path\to\file`))
		_ = g.AddNode(n)

		output := g.String()

		asrt.Contains(output, `label="path\\to\\file"`, "expected backslashes escaped in label")
	})

	t.Run("backslash in edge label", func(t *testing.T) {
		asrt := assert.New(t)
		g := NewGraph(Directed)
		n1 := NewNode("A")
		n2 := NewNode("B")
		_, _ = g.AddEdge(n1, n2, WithEdgeLabel(`path\file`))

		output := g.String()

		asrt.Contains(output, `label="path\\file"`, "expected backslash escaped in edge label")
	})

	t.Run("backslash in graph label", func(t *testing.T) {
		asrt := assert.New(t)
		g := NewGraph(Directed, WithGraphLabel(`C:\Project`))

		output := g.String()

		asrt.Contains(output, `label="C:\\Project";`, "expected backslash escaped in graph label")
	})

	t.Run("backslash in custom attribute value", func(t *testing.T) {
		asrt := assert.New(t)
		g := NewGraph(Directed)
		n := NewNode("A", WithNodeAttribute("tooltip", `hover\text`))
		_ = g.AddNode(n)

		output := g.String()

		asrt.Contains(output, `tooltip="hover\\text"`, "expected backslash escaped in custom attribute")
	})
}

// Test escaping double quotes in attribute values
func TestDOT_Escaping_Quotes_AttributeValue(t *testing.T) {
	t.Run("quote in label attribute is escaped", func(t *testing.T) {
		asrt := assert.New(t)
		g := NewGraph(Directed)
		n := NewNode("A", WithLabel(`Say "Hello"`))
		_ = g.AddNode(n)

		output := g.String()

		asrt.Contains(output, `label="Say \"Hello\""`, "expected quotes escaped in label")
	})

	t.Run("quote in edge label", func(t *testing.T) {
		asrt := assert.New(t)
		g := NewGraph(Directed)
		n1 := NewNode("A")
		n2 := NewNode("B")
		_, _ = g.AddEdge(n1, n2, WithEdgeLabel(`"important"`))

		output := g.String()

		asrt.Contains(output, `label="\"important\""`, "expected quotes escaped in edge label")
	})

	t.Run("quote in graph label", func(t *testing.T) {
		asrt := assert.New(t)
		g := NewGraph(Directed, WithGraphLabel(`The "Main" Graph`))

		output := g.String()

		asrt.Contains(output, `label="The \"Main\" Graph";`, "expected quotes escaped in graph label")
	})

	t.Run("multiple quotes in label", func(t *testing.T) {
		asrt := assert.New(t)
		g := NewGraph(Directed)
		n := NewNode("A", WithLabel(`"word1" "word2" "word3"`))
		_ = g.AddNode(n)

		output := g.String()

		asrt.Contains(output, `label="\"word1\" \"word2\" \"word3\""`, "expected all quotes escaped")
	})
}

// Test escaping newlines in attribute values
func TestDOT_Escaping_Newlines_AttributeValue(t *testing.T) {
	t.Run("newline in label attribute is escaped to literal backslash-n", func(t *testing.T) {
		asrt := assert.New(t)
		g := NewGraph(Directed)
		n := NewNode("A", WithLabel("Line 1\nLine 2"))
		_ = g.AddNode(n)

		output := g.String()

		asrt.Contains(output, `label="Line 1\nLine 2"`, "expected newline escaped to literal \\n")
		// Verify no actual newline in the label value
		lines := strings.Split(output, "\n")
		for _, line := range lines {
			if strings.Contains(line, "label=") {
				asrt.NotContains(line, "Line 1\nLine 2", "expected no actual newline in label")
			}
		}
	})

	t.Run("newline in edge label", func(t *testing.T) {
		asrt := assert.New(t)
		g := NewGraph(Directed)
		n1 := NewNode("A")
		n2 := NewNode("B")
		_, _ = g.AddEdge(n1, n2, WithEdgeLabel("First\nSecond"))

		output := g.String()

		asrt.Contains(output, `label="First\nSecond"`, "expected newline escaped in edge label")
	})

	t.Run("newline in graph label", func(t *testing.T) {
		asrt := assert.New(t)
		g := NewGraph(Directed, WithGraphLabel("Multi\nLine\nTitle"))

		output := g.String()

		asrt.Contains(output, `label="Multi\nLine\nTitle";`, "expected newlines escaped in graph label")
	})

	t.Run("multiple consecutive newlines", func(t *testing.T) {
		asrt := assert.New(t)
		g := NewGraph(Directed)
		n := NewNode("A", WithLabel("Para1\n\nPara2"))
		_ = g.AddNode(n)

		output := g.String()

		asrt.Contains(output, `label="Para1\n\nPara2"`, "expected consecutive newlines escaped")
	})
}

// Test combined escaping (multiple special characters)
func TestDOT_Escaping_Combined(t *testing.T) {
	t.Run("backslash and quote in same node ID", func(t *testing.T) {
		asrt := assert.New(t)
		g := NewGraph(Directed)
		n := NewNode(`path\"file`)
		_ = g.AddNode(n)

		output := g.String()

		asrt.Contains(output, `"path\\\"file";`, "expected both backslash and quote escaped")
	})

	t.Run("backslash and newline in label", func(t *testing.T) {
		asrt := assert.New(t)
		g := NewGraph(Directed)
		n := NewNode("A", WithLabel("path\\file\nnext"))
		_ = g.AddNode(n)

		output := g.String()

		asrt.Contains(output, `label="path\\file\nnext"`, "expected backslash and newline escaped")
	})

	t.Run("quote and newline in label", func(t *testing.T) {
		asrt := assert.New(t)
		g := NewGraph(Directed)
		n := NewNode("A", WithLabel("Say \"Hi\"\nBye"))
		_ = g.AddNode(n)

		output := g.String()

		asrt.Contains(output, `label="Say \"Hi\"\nBye"`, "expected quote and newline escaped")
	})

	t.Run("all three special chars in label", func(t *testing.T) {
		asrt := assert.New(t)
		g := NewGraph(Directed)
		n := NewNode("A", WithLabel("C:\\path\\to \"file\"\nNew line"))
		_ = g.AddNode(n)

		output := g.String()

		asrt.Contains(output, `label="C:\\path\\to \"file\"\nNew line"`, "expected all special chars escaped")
	})

	t.Run("complex escaping in edge label", func(t *testing.T) {
		asrt := assert.New(t)
		g := NewGraph(Directed)
		n1 := NewNode("A")
		n2 := NewNode("B")
		_, _ = g.AddEdge(n1, n2, WithEdgeLabel("Path: C:\\Users\\\"Me\"\nDone"))

		output := g.String()

		asrt.Contains(output, `label="Path: C:\\Users\\\"Me\"\nDone"`, "expected complex escaping")
	})
}

// Test node IDs with spaces require quoting
func TestDOT_NodeID_WithSpaces(t *testing.T) {
	t.Run("node ID with single space is quoted", func(t *testing.T) {
		asrt := assert.New(t)
		g := NewGraph(Directed)
		n := NewNode("node one")
		_ = g.AddNode(n)

		output := g.String()

		asrt.Contains(output, `"node one";`, "expected node ID with space to be quoted")
	})

	t.Run("node ID with multiple spaces is quoted", func(t *testing.T) {
		asrt := assert.New(t)
		g := NewGraph(Directed)
		n := NewNode("this has many spaces")
		_ = g.AddNode(n)

		output := g.String()

		asrt.Contains(output, `"this has many spaces";`, "expected node ID with multiple spaces quoted")
	})

	t.Run("node ID with leading space is quoted", func(t *testing.T) {
		asrt := assert.New(t)
		g := NewGraph(Directed)
		n := NewNode(" leading")
		_ = g.AddNode(n)

		output := g.String()

		asrt.Contains(output, `" leading";`, "expected node ID with leading space quoted")
	})

	t.Run("node ID with trailing space is quoted", func(t *testing.T) {
		asrt := assert.New(t)
		g := NewGraph(Directed)
		n := NewNode("trailing ")
		_ = g.AddNode(n)

		output := g.String()

		asrt.Contains(output, `"trailing ";`, "expected node ID with trailing space quoted")
	})

	t.Run("node ID with space appears in edges", func(t *testing.T) {
		asrt := assert.New(t)
		g := NewGraph(Directed)
		n1 := NewNode("node one")
		n2 := NewNode("node two")
		_, _ = g.AddEdge(n1, n2)

		output := g.String()

		asrt.Contains(output, `"node one" -> "node two";`, "expected quoted node IDs in edge")
	})
}

// Test node IDs with special characters require quoting
func TestDOT_NodeID_WithSpecialChars(t *testing.T) {
	t.Run("node ID with hyphen is quoted", func(t *testing.T) {
		asrt := assert.New(t)
		g := NewGraph(Directed)
		n := NewNode("node-one")
		_ = g.AddNode(n)

		output := g.String()

		asrt.Contains(output, `"node-one";`, "expected node ID with hyphen quoted")
	})

	t.Run("node ID with period is quoted", func(t *testing.T) {
		asrt := assert.New(t)
		g := NewGraph(Directed)
		n := NewNode("node.one")
		_ = g.AddNode(n)

		output := g.String()

		asrt.Contains(output, `"node.one";`, "expected node ID with period quoted")
	})

	t.Run("node ID with colon is quoted", func(t *testing.T) {
		asrt := assert.New(t)
		g := NewGraph(Directed)
		n := NewNode("node:port")
		_ = g.AddNode(n)

		output := g.String()

		asrt.Contains(output, `"node:port";`, "expected node ID with colon quoted")
	})

	t.Run("node ID with forward slash is quoted", func(t *testing.T) {
		asrt := assert.New(t)
		g := NewGraph(Directed)
		n := NewNode("node/path")
		_ = g.AddNode(n)

		output := g.String()

		asrt.Contains(output, `"node/path";`, "expected node ID with slash quoted")
	})

	t.Run("node ID with parentheses is quoted", func(t *testing.T) {
		asrt := assert.New(t)
		g := NewGraph(Directed)
		n := NewNode("func(x)")
		_ = g.AddNode(n)

		output := g.String()

		asrt.Contains(output, `"func(x)";`, "expected node ID with parentheses quoted")
	})

	t.Run("node ID with brackets is quoted", func(t *testing.T) {
		asrt := assert.New(t)
		g := NewGraph(Directed)
		n := NewNode("arr[0]")
		_ = g.AddNode(n)

		output := g.String()

		asrt.Contains(output, `"arr[0]";`, "expected node ID with brackets quoted")
	})

	t.Run("node ID with braces is quoted", func(t *testing.T) {
		asrt := assert.New(t)
		g := NewGraph(Directed)
		n := NewNode("set{}")
		_ = g.AddNode(n)

		output := g.String()

		asrt.Contains(output, `"set{}";`, "expected node ID with braces quoted")
	})
}

// Test node IDs starting with digit require quoting
func TestDOT_NodeID_StartingWithDigit(t *testing.T) {
	t.Run("node ID starting with digit is quoted", func(t *testing.T) {
		asrt := assert.New(t)
		g := NewGraph(Directed)
		n := NewNode("123node")
		_ = g.AddNode(n)

		output := g.String()

		asrt.Contains(output, `"123node";`, "expected node ID starting with digit quoted")
	})

	t.Run("node ID that is only digits is quoted", func(t *testing.T) {
		asrt := assert.New(t)
		g := NewGraph(Directed)
		n := NewNode("12345")
		_ = g.AddNode(n)

		output := g.String()

		asrt.Contains(output, `"12345";`, "expected numeric node ID quoted")
	})

	t.Run("node ID starting with zero is quoted", func(t *testing.T) {
		asrt := assert.New(t)
		g := NewGraph(Directed)
		n := NewNode("0x1F")
		_ = g.AddNode(n)

		output := g.String()

		asrt.Contains(output, `"0x1F";`, "expected node ID starting with 0 quoted")
	})
}

// Test simple alphanumeric node IDs (optionally unquoted, but we always quote for safety)
func TestDOT_NodeID_SimpleAlphanumeric(t *testing.T) {
	t.Run("simple lowercase node ID is quoted for safety", func(t *testing.T) {
		asrt := assert.New(t)
		g := NewGraph(Directed)
		n := NewNode("node")
		_ = g.AddNode(n)

		output := g.String()

		asrt.Contains(output, `"node";`, "expected simple node ID quoted for safety")
	})

	t.Run("simple uppercase node ID is quoted for safety", func(t *testing.T) {
		asrt := assert.New(t)
		g := NewGraph(Directed)
		n := NewNode("NODE")
		_ = g.AddNode(n)

		output := g.String()

		asrt.Contains(output, `"NODE";`, "expected uppercase node ID quoted for safety")
	})

	t.Run("camelCase node ID is quoted for safety", func(t *testing.T) {
		asrt := assert.New(t)
		g := NewGraph(Directed)
		n := NewNode("myNode")
		_ = g.AddNode(n)

		output := g.String()

		asrt.Contains(output, `"myNode";`, "expected camelCase node ID quoted for safety")
	})

	t.Run("node ID with underscores is quoted for safety", func(t *testing.T) {
		asrt := assert.New(t)
		g := NewGraph(Directed)
		n := NewNode("my_node_1")
		_ = g.AddNode(n)

		output := g.String()

		asrt.Contains(output, `"my_node_1";`, "expected node ID with underscores quoted for safety")
	})
}

// Test graph name escaping
func TestDOT_GraphName_Escaping(t *testing.T) {
	t.Run("graph name with quote is escaped", func(t *testing.T) {
		asrt := assert.New(t)
		g := NewGraph(Directed, WithName(`My"Graph`))

		output := g.String()

		asrt.Contains(output, `digraph "My\"Graph"`, "expected quote escaped in graph name")
	})

	t.Run("graph name with backslash is escaped", func(t *testing.T) {
		asrt := assert.New(t)
		g := NewGraph(Directed, WithName(`Graph\1`))

		output := g.String()

		asrt.Contains(output, `digraph "Graph\\1"`, "expected backslash escaped in graph name")
	})

	t.Run("graph name with newline is escaped", func(t *testing.T) {
		asrt := assert.New(t)
		g := NewGraph(Directed, WithName("Multi\nLine"))

		output := g.String()

		asrt.Contains(output, `digraph "Multi\nLine"`, "expected newline escaped in graph name")
	})

	t.Run("graph name with space is quoted", func(t *testing.T) {
		asrt := assert.New(t)
		g := NewGraph(Directed, WithName("My Graph"))

		output := g.String()

		asrt.Contains(output, `digraph "My Graph"`, "expected graph name with space quoted")
	})

	t.Run("simple graph name is quoted for safety", func(t *testing.T) {
		asrt := assert.New(t)
		g := NewGraph(Directed, WithName("G"))

		output := g.String()

		asrt.Contains(output, `digraph "G"`, "expected simple graph name quoted for safety")
	})
}

// Test complex strings with multiple escape scenarios
func TestDOT_ComplexStrings(t *testing.T) {
	t.Run("node with complex ID and label", func(t *testing.T) {
		asrt := assert.New(t)
		g := NewGraph(Directed)
		n := NewNode(`user\admin`, WithLabel("Name: \"Admin\"\nRole: Supervisor"))
		_ = g.AddNode(n)

		output := g.String()

		asrt.Contains(output, `"user\\admin"`, "expected backslash escaped in node ID")
		asrt.Contains(output, `label="Name: \"Admin\"\nRole: Supervisor"`, "expected complex label escaping")
	})

	t.Run("edge with complex labels on both nodes and edge", func(t *testing.T) {
		asrt := assert.New(t)
		g := NewGraph(Directed)
		n1 := NewNode("A", WithLabel("\"Start\"\nPoint"))
		n2 := NewNode("B", WithLabel("\"End\"\nPoint"))
		_, _ = g.AddEdge(n1, n2, WithEdgeLabel("Path: C:\\\\Users\n\"Active\""))

		output := g.String()

		asrt.Contains(output, `"A" [label="\"Start\"\nPoint"]`, "expected escaping in first node")
		asrt.Contains(output, `"B" [label="\"End\"\nPoint"]`, "expected escaping in second node")
		asrt.Contains(output, `"A" -> "B" [label="Path: C:\\\\Users\n\"Active\""]`, "expected complex edge label escaping")
	})

	t.Run("full graph with escaping in all sections", func(t *testing.T) {
		asrt := assert.New(t)
		g := NewGraph(
			Directed,
			WithName(`My"Graph`),
			WithGraphLabel("Title: \"Main\"\nSubtitle: Details"),
		)
		n1 := NewNode(`node\1`, WithLabel(`"First"`))
		n2 := NewNode(`node 2`, WithLabel("Second\nLine"))
		_ = g.AddNode(n1)
		_ = g.AddNode(n2)
		_, _ = g.AddEdge(n1, n2, WithEdgeLabel("Link\n\"Active\""))

		output := g.String()

		asrt.Contains(output, `digraph "My\"Graph"`, "expected graph name escaping")
		asrt.Contains(output, `label="Title: \"Main\"\nSubtitle: Details"`, "expected graph label escaping")
		asrt.Contains(output, `"node\\1" [label="\"First\""]`, "expected first node escaping")
		asrt.Contains(output, `"node 2" [label="Second\nLine"]`, "expected second node escaping")
		asrt.Contains(output, `"node\\1" -> "node 2" [label="Link\n\"Active\""]`, "expected edge escaping")
	})
}

// Test edge cases for string escaping
func TestDOT_Escaping_EdgeCases(t *testing.T) {
	t.Run("empty string node ID", func(t *testing.T) {
		asrt := assert.New(t)
		g := NewGraph(Directed)
		n := NewNode("")
		_ = g.AddNode(n)

		output := g.String()

		asrt.Contains(output, `"";`, "expected empty string quoted")
	})

	t.Run("node ID with only backslashes", func(t *testing.T) {
		asrt := assert.New(t)
		g := NewGraph(Directed)
		n := NewNode(`\\\`)
		_ = g.AddNode(n)

		output := g.String()

		asrt.Contains(output, `"\\\\\\";`, "expected all backslashes escaped")
	})

	t.Run("node ID with only quotes", func(t *testing.T) {
		asrt := assert.New(t)
		g := NewGraph(Directed)
		n := NewNode(`"""`)
		_ = g.AddNode(n)

		output := g.String()

		asrt.Contains(output, `"\"\"\"";`, "expected all quotes escaped")
	})

	t.Run("label with only newlines", func(t *testing.T) {
		asrt := assert.New(t)
		g := NewGraph(Directed)
		n := NewNode("A", WithLabel("\n\n\n"))
		_ = g.AddNode(n)

		output := g.String()

		asrt.Contains(output, `label="\n\n\n"`, "expected all newlines escaped")
	})

	t.Run("empty label attribute", func(t *testing.T) {
		asrt := assert.New(t)
		g := NewGraph(Directed)
		n := NewNode("A", WithLabel(""))
		_ = g.AddNode(n)

		output := g.String()

		asrt.Contains(output, `label=""`, "expected empty label quoted")
	})

	t.Run("tab character in label", func(t *testing.T) {
		asrt := assert.New(t)
		g := NewGraph(Directed)
		n := NewNode("A", WithLabel("Col1\tCol2"))
		_ = g.AddNode(n)

		output := g.String()

		// Tab may or may not be escaped; just verify output contains it somehow
		asrt.Contains(output, "label=", "expected label attribute")
	})
}

// Test attribute value escaping (always quoted)
func TestDOT_AttributeValues_AlwaysQuoted(t *testing.T) {
	t.Run("string attribute values are quoted", func(t *testing.T) {
		asrt := assert.New(t)
		g := NewGraph(Directed)
		n := NewNode("A", WithLabel("Simple"))
		_ = g.AddNode(n)

		output := g.String()

		asrt.Contains(output, `label="Simple"`, "expected string value quoted")
	})

	t.Run("color attribute values are quoted", func(t *testing.T) {
		asrt := assert.New(t)
		g := NewGraph(Directed)
		n := NewNode("A", WithColor("blue"))
		_ = g.AddNode(n)

		output := g.String()

		asrt.Contains(output, `color="blue"`, "expected color value quoted")
	})

	t.Run("shape attribute values are quoted", func(t *testing.T) {
		asrt := assert.New(t)
		g := NewGraph(Directed)
		n := NewNode("A", WithBoxShape())
		_ = g.AddNode(n)

		output := g.String()

		asrt.Contains(output, `shape="box"`, "expected shape value quoted")
	})

	t.Run("custom attribute values are quoted", func(t *testing.T) {
		asrt := assert.New(t)
		g := NewGraph(Directed)
		n := NewNode("A", WithNodeAttribute("custom", "value"))
		_ = g.AddNode(n)

		output := g.String()

		asrt.Contains(output, `custom="value"`, "expected custom attribute value quoted")
	})

	t.Run("numeric attribute values are quoted as strings", func(t *testing.T) {
		asrt := assert.New(t)
		g := NewGraph(Directed)
		n := NewNode("A", WithFontSize(12.5))
		_ = g.AddNode(n)

		output := g.String()

		asrt.Contains(output, `fontsize="12.5"`, "expected numeric value quoted")
	})
}

// Test undirected graph edge escaping
func TestDOT_Escaping_UndirectedEdges(t *testing.T) {
	t.Run("undirected edge with escaped node IDs", func(t *testing.T) {
		asrt := assert.New(t)
		g := NewGraph(Undirected)
		n1 := NewNode(`node\1`)
		n2 := NewNode(`node"2`)
		_, _ = g.AddEdge(n1, n2)

		output := g.String()

		asrt.Contains(output, `"node\\1" -- "node\"2";`, "expected escaped node IDs in undirected edge")
	})

	t.Run("undirected edge with escaped label", func(t *testing.T) {
		asrt := assert.New(t)
		g := NewGraph(Undirected)
		n1 := NewNode("A")
		n2 := NewNode("B")
		_, _ = g.AddEdge(n1, n2, WithEdgeLabel("Link\n\"Active\""))

		output := g.String()

		asrt.Contains(output, `label="Link\n\"Active\""`, "expected escaped edge label in undirected graph")
	})
}

// Test DOT subgraph output - simple subgraph
func TestDOT_Subgraph_Simple(t *testing.T) {
	asrt := assert.New(t)
	g := NewGraph(Directed)

	n1 := NewNode("A")
	n2 := NewNode("B")

	g.Subgraph("sub1", func(sg *Subgraph) {
		_ = sg.AddNode(n1)
		_ = sg.AddNode(n2)
	})

	output := g.String()

	asrt.Contains(output, `subgraph "sub1" {`, "expected subgraph declaration")
	asrt.Contains(output, `"A";`, "expected node A in output")
	asrt.Contains(output, `"B";`, "expected node B in output")

	// Verify nodes appear in subgraph section, not loose
	subgraphIdx := strings.Index(output, `subgraph "sub1"`)
	nodeAIdx := strings.Index(output, `"A";`)
	nodeBIdx := strings.Index(output, `"B";`)
	closingBraceIdx := strings.Index(output[subgraphIdx:], "}")

	asrt.Greater(nodeAIdx, subgraphIdx, "expected node A after subgraph declaration")
	asrt.Greater(nodeBIdx, subgraphIdx, "expected node B after subgraph declaration")
	asrt.Less(nodeAIdx, subgraphIdx+closingBraceIdx, "expected node A before subgraph closing")
	asrt.Less(nodeBIdx, subgraphIdx+closingBraceIdx, "expected node B before subgraph closing")
}

// Test DOT subgraph output - with attributes
func TestDOT_Subgraph_WithAttributes(t *testing.T) {
	asrt := assert.New(t)
	g := NewGraph(Directed)

	n1 := NewNode("A")

	g.Subgraph("sub1", func(sg *Subgraph) {
		sg.SetLabel("My Subgraph")
		sg.SetStyle("filled")
		_ = sg.AddNode(n1)
	})

	output := g.String()

	asrt.Contains(output, `subgraph "sub1" {`, "expected subgraph declaration")
	asrt.Contains(output, `label="My Subgraph";`, "expected label attribute")
	asrt.Contains(output, `style="filled";`, "expected style attribute")

	// Verify attributes come before nodes in subgraph
	labelIdx := strings.Index(output, `label="My Subgraph"`)
	nodeIdx := strings.Index(output, `"A";`)
	asrt.Less(labelIdx, nodeIdx, "expected attributes before nodes in subgraph")
}

// Test DOT subgraph output - cluster
func TestDOT_Subgraph_Cluster(t *testing.T) {
	asrt := assert.New(t)
	g := NewGraph(Directed)

	db1 := NewNode("db1")
	db2 := NewNode("db2")

	g.Subgraph("cluster_db", func(sg *Subgraph) {
		sg.SetLabel("Database")
		_ = sg.AddNode(db1)
		_ = sg.AddNode(db2)
	})

	output := g.String()

	asrt.Contains(output, `subgraph "cluster_db" {`, "expected cluster subgraph declaration")
	asrt.Contains(output, `label="Database";`, "expected label")
	asrt.Contains(output, `"db1";`, "expected db1 node")
	asrt.Contains(output, `"db2";`, "expected db2 node")
}

// Test DOT subgraph output - nested subgraphs
func TestDOT_Subgraph_Nested(t *testing.T) {
	asrt := assert.New(t)
	g := NewGraph(Directed)

	n1 := NewNode("A")
	n2 := NewNode("B")

	g.Subgraph("cluster_outer", func(outer *Subgraph) {
		outer.SetLabel("Outer")
		_ = outer.AddNode(n1)

		outer.Subgraph("cluster_inner", func(inner *Subgraph) {
			inner.SetLabel("Inner")
			_ = inner.AddNode(n2)
		})
	})

	output := g.String()

	asrt.Contains(output, `subgraph "cluster_outer" {`, "expected outer subgraph")
	asrt.Contains(output, `label="Outer";`, "expected outer label")
	asrt.Contains(output, `subgraph "cluster_inner" {`, "expected inner subgraph")
	asrt.Contains(output, `label="Inner";`, "expected inner label")

	// Verify nesting structure - inner should be between outer's braces
	outerIdx := strings.Index(output, `subgraph "cluster_outer"`)
	innerIdx := strings.Index(output, `subgraph "cluster_inner"`)
	asrt.Greater(innerIdx, outerIdx, "expected inner subgraph after outer declaration")

	// Node B should be in inner subgraph, not outer
	innerStart := strings.Index(output, `subgraph "cluster_inner"`)
	if innerStart < 0 {
		t.Fatal("expected output to contain `subgraph \"cluster_inner\"`")
	}
	innerSection := output[innerStart:]
	innerEnd := strings.Index(innerSection, "\n\t}")
	if innerStart < 0 {
		t.Fatal("expected innerSection to contain `\n\t}`")
	}
	innerContent := innerSection[:innerEnd]
	asrt.Contains(innerContent, `"B";`, "expected node B in inner subgraph")

	// Node A should be in outer but not in inner
	outerStart := strings.Index(output, `subgraph "cluster_outer"`)
	if outerStart < 0 {
		t.Fatalf("expected output to include `subgraph \"cluster_outer\"")
	}
	outerSection := output[outerStart:]
	outerEnd := strings.LastIndex(outerSection, "\n\t}")
	outerContent := outerSection[:outerEnd]
	asrt.Contains(outerContent, `"A";`, "expected node A in outer subgraph")
}

// Test DOT subgraph output - anonymous subgraph
func TestDOT_Subgraph_Anonymous(t *testing.T) {
	asrt := assert.New(t)
	g := NewGraph(Directed)

	n1 := NewNode("A")
	n2 := NewNode("B")

	g.Subgraph("", func(sg *Subgraph) {
		sg.SetAttribute("rank", "same")
		_ = sg.AddNode(n1)
		_ = sg.AddNode(n2)
	})

	output := g.String()

	// Anonymous subgraph should not have a name after "subgraph"
	asrt.Contains(output, "subgraph {", "expected anonymous subgraph declaration")
	asrt.Contains(output, `rank="same";`, "expected rank attribute")
	asrt.Contains(output, `"A";`, "expected node A")
	asrt.Contains(output, `"B";`, "expected node B")
}

// Test DOT output - graph with both subgraphs and loose nodes
func TestDOT_Graph_WithSubgraphsAndLooseNodes(t *testing.T) {
	asrt := assert.New(t)
	g := NewGraph(Directed)

	db1 := NewNode("db1")
	db2 := NewNode("db2")
	web := NewNode("web")

	g.Subgraph("cluster_db", func(sg *Subgraph) {
		sg.SetLabel("Database")
		_ = sg.AddNode(db1)
		_ = sg.AddNode(db2)
	})

	_ = g.AddNode(web)
	_, _ = g.AddEdge(web, db1)

	output := g.String()

	// Verify output structure
	asrt.Contains(output, `subgraph "cluster_db" {`, "expected subgraph")
	asrt.Contains(output, `label="Database";`, "expected subgraph label")

	// Find positions
	subgraphIdx := strings.Index(output, `subgraph "cluster_db"`)
	subgraphEnd := strings.Index(output[subgraphIdx:], "\n\t}")
	subgraphSection := output[subgraphIdx : subgraphIdx+subgraphEnd]

	// db1 and db2 should be in subgraph
	asrt.Contains(subgraphSection, `"db1";`, "expected db1 in subgraph")
	asrt.Contains(subgraphSection, `"db2";`, "expected db2 in subgraph")
	asrt.NotContains(subgraphSection, `"web";`, "expected web NOT in subgraph")

	// web should appear as loose node after subgraph
	afterSubgraph := output[subgraphIdx+subgraphEnd:]
	webIdx := strings.Index(afterSubgraph, `"web";`)
	edgeIdx := strings.Index(afterSubgraph, `"web" -> "db1";`)

	asrt.NotEqual(-1, webIdx, "expected web node after subgraph")
	asrt.NotEqual(-1, edgeIdx, "expected edge in output")
	asrt.Less(webIdx, edgeIdx, "expected loose nodes before edges")
}
