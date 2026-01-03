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
		g.AddNode(n)

		output := g.String()

		expected := "digraph {\n\t\"A\";\n}"
		asrt.Equal(expected, output, "expected node to be rendered with quoted ID and semicolon")
	})

	t.Run("node ID is always quoted", func(t *testing.T) {
		asrt := assert.New(t)
		g := NewGraph(Directed)
		n := NewNode("SimpleID")
		g.AddNode(n)

		output := g.String()

		asrt.Contains(output, "\"SimpleID\";", "expected node ID to be quoted even if simple")
		asrt.NotContains(output, "SimpleID;", "expected no unquoted node ID")
	})

	t.Run("node appears inside graph body", func(t *testing.T) {
		asrt := assert.New(t)
		g := NewGraph(Directed)
		n := NewNode("A")
		g.AddNode(n)

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
		g.AddNode(n)

		output := g.String()

		asrt.Contains(output, "\t\"A\";", "expected node line to be indented with tab")
	})

	t.Run("undirected graph renders node same way", func(t *testing.T) {
		asrt := assert.New(t)
		g := NewGraph(Undirected)
		n := NewNode("A")
		g.AddNode(n)

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
		g.AddNode(n)

		output := g.String()

		asrt.Contains(output, "\"A\" [label=\"Node A\"];", "expected node with label attribute")
	})

	t.Run("label value is quoted", func(t *testing.T) {
		asrt := assert.New(t)
		g := NewGraph(Directed)
		n := NewNode("A", WithLabel("My Label"))
		g.AddNode(n)

		output := g.String()

		asrt.Contains(output, "label=\"My Label\"", "expected quoted label value")
		asrt.NotContains(output, "label=My Label", "expected no unquoted label value")
	})

	t.Run("label with special characters", func(t *testing.T) {
		asrt := assert.New(t)
		g := NewGraph(Directed)
		n := NewNode("A", WithLabel("Label with spaces"))
		g.AddNode(n)

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
		g.AddNode(n)

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
			g.AddNode(n)

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
		g.AddNode(n)

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
		g.AddNode(n)

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
		g.AddNode(n)

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
		g1.AddNode(n1)

		g2 := NewGraph(Directed)
		n2 := NewNode("A", WithBoxShape(), WithLabel("A"))
		g2.AddNode(n2)

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
		g.AddNode(n)

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
		g.AddNode(n)

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
		g.AddNode(n)

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
		g.AddNode(n1)
		g.AddNode(n2)
		g.AddNode(n3)

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
		g.AddNode(n1)
		g.AddNode(n2)
		g.AddNode(n3)

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
		g.AddNode(n1)
		g.AddNode(n2)
		g.AddNode(n3) // Replaces n1 in place

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
		g.AddNode(n)

		output := g.String()

		asrt.Contains(output, "\"node1\";", "expected quoted simple ID")
	})

	t.Run("ID with spaces is quoted", func(t *testing.T) {
		asrt := assert.New(t)
		g := NewGraph(Directed)
		n := NewNode("node with spaces")
		g.AddNode(n)

		output := g.String()

		asrt.Contains(output, "\"node with spaces\";", "expected quoted ID with spaces")
	})

	t.Run("ID with special characters is quoted", func(t *testing.T) {
		asrt := assert.New(t)
		g := NewGraph(Directed)
		n := NewNode("node-with-dashes")
		g.AddNode(n)

		output := g.String()

		asrt.Contains(output, "\"node-with-dashes\";", "expected quoted ID with dashes")
	})

	t.Run("ID starting with number is quoted", func(t *testing.T) {
		asrt := assert.New(t)
		g := NewGraph(Directed)
		n := NewNode("123node")
		g.AddNode(n)

		output := g.String()

		asrt.Contains(output, "\"123node\";", "expected quoted ID starting with number")
	})

	t.Run("empty ID is quoted", func(t *testing.T) {
		asrt := assert.New(t)
		g := NewGraph(Directed)
		n := NewNode("")
		g.AddNode(n)

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
		g.AddNode(n)

		output := g.String()

		asrt.Contains(output, "label=\"Label\"", "expected quoted label value")
		asrt.Contains(output, "color=\"red\"", "expected quoted color value")
	})

	t.Run("numeric fontsize is quoted as string", func(t *testing.T) {
		asrt := assert.New(t)
		g := NewGraph(Directed)
		n := NewNode("A", WithFontSize(14.5))
		g.AddNode(n)

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
		g.AddNode(n)

		output := g.String()

		expected := "strict digraph {\n\t\"A\";\n}"
		asrt.Equal(expected, output, "expected node in strict graph")
	})

	t.Run("nodes in named graph", func(t *testing.T) {
		asrt := assert.New(t)
		g := NewGraph(Directed, WithName("G"))
		n := NewNode("A", WithLabel("Node A"))
		g.AddNode(n)

		output := g.String()

		expected := "digraph G {\n\t\"A\" [label=\"Node A\"];\n}"
		asrt.Equal(expected, output, "expected node in named graph")
	})

	t.Run("nodes in undirected graph", func(t *testing.T) {
		asrt := assert.New(t)
		g := NewGraph(Undirected)
		n1 := NewNode("A", WithCircleShape())
		n2 := NewNode("B", WithBoxShape())
		g.AddNode(n1)
		g.AddNode(n2)

		output := g.String()

		asrt.Contains(output, "graph {", "expected undirected graph")
		asrt.Contains(output, "\"A\" [shape=\"circle\"];", "expected first node")
		asrt.Contains(output, "\"B\" [shape=\"box\"];", "expected second node")
	})
}
