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
		asrt.Contains(output, "fontsize=\"14.00\"", "expected fontsize")
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

		asrt.Contains(output, "fontsize=\"14.50\"", "expected fontsize as quoted string")
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

		expected := "digraph G {\n\t\"A\" [label=\"Node A\"];\n}"
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
		asrt.Contains(output, "weight=\"2.50\"", "expected weight")
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

		asrt.Contains(output, "weight=\"2.50\"", "expected weight formatted as %.2f")
	})

	t.Run("integer weight formatted with decimal places", func(t *testing.T) {
		asrt := assert.New(t)

		g := NewGraph(Directed)
		n1 := NewNode("A")
		n2 := NewNode("B")
		_, _ = g.AddEdge(n1, n2, WithWeight(3))

		output := g.String()

		asrt.Contains(output, "weight=\"3.00\"", "expected integer weight with .00")
	})

	t.Run("weight with many decimal places is truncated", func(t *testing.T) {
		asrt := assert.New(t)

		g := NewGraph(Directed)
		n1 := NewNode("A")
		n2 := NewNode("B")
		_, _ = g.AddEdge(n1, n2, WithWeight(2.12345))

		output := g.String()

		asrt.Contains(output, "weight=\"2.12\"", "expected weight truncated to 2 decimal places")
		asrt.NotContains(output, "weight=\"2.12345\"", "expected no more than 2 decimal places")
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

		asrt.Contains(output, "digraph G {", "expected graph declaration")
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

		asrt.Contains(output, "strict digraph StrictG {", "expected strict graph declaration")
		asrt.Contains(output, "\"A\" -> \"B\";", "expected edge in strict graph")
	})

	t.Run("undirected graph with nodes and edges", func(t *testing.T) {
		asrt := assert.New(t)

		g := NewGraph(Undirected, WithName("UG"))
		n1 := NewNode("A", WithCircleShape())
		n2 := NewNode("B", WithBoxShape())

		_, _ = g.AddEdge(n1, n2, WithEdgeStyle(EdgeStyleDashed))

		output := g.String()

		asrt.Contains(output, "graph UG {", "expected undirected graph declaration")
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
