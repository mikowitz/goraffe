package goraffe_test

// Commented examples below - uncomment and update as needed

import (
	"fmt"

	"github.com/mikowitz/goraffe"
)

// Example_basicDirectedGraph demonstrates creating a simple directed graph with two nodes and an edge.
func Example_basicDirectedGraph() {
	// Create a directed graph
	g := goraffe.NewGraph(goraffe.Directed)

	// Add nodes
	start := goraffe.NewNode("A")
	end := goraffe.NewNode("B")

	// Connect nodes with an edge
	_, _ = g.AddEdge(start, end)

	// Output DOT format
	fmt.Println(g.String())
	// Output:
	// digraph {
	// 	"A";
	// 	"B";
	// 	"A" -> "B";
	// }
}

// Example_styledNodes demonstrates creating nodes with various visual attributes.
func Example_styledNodes() {
	g := goraffe.NewGraph(goraffe.Directed)

	// Create nodes with different shapes and colors
	start := goraffe.NewNode("start",
		goraffe.WithLabel("Start"),
		goraffe.WithCircleShape(),
		goraffe.WithFillColor("lightgreen"),
	)

	process := goraffe.NewNode("process",
		goraffe.WithLabel("Process"),
		goraffe.WithBoxShape(),
		goraffe.WithFillColor("lightblue"),
	)

	end := goraffe.NewNode("end",
		goraffe.WithLabel("End"),
		goraffe.WithCircleShape(),
		goraffe.WithFillColor("pink"),
	)

	// Add nodes to graph
	// Normally you would want to handle any possible errors here
	_ = g.AddNode(start)
	_ = g.AddNode(process)
	_ = g.AddNode(end)

	fmt.Println(g.String())
	// Output:
	// digraph {
	// 	"start" [fillcolor="lightgreen", label="Start", shape="circle", style="filled"];
	// 	"process" [fillcolor="lightblue", label="Process", shape="box", style="filled"];
	// 	"end" [fillcolor="pink", label="End", shape="circle", style="filled"];
	// }
}

// Example_edgeAttributes demonstrates creating edges with labels, colors, and styles.
func Example_edgeAttributes() {
	g := goraffe.NewGraph(goraffe.Directed, goraffe.WithName("Workflow"))

	n1 := goraffe.NewNode("A", goraffe.WithLabel("Start"))
	n2 := goraffe.NewNode("B", goraffe.WithLabel("Process"))
	n3 := goraffe.NewNode("C", goraffe.WithLabel("End"))

	// Add edges with different styles
	_, _ = g.AddEdge(n1, n2,
		goraffe.WithEdgeLabel("begin"),
		goraffe.WithEdgeColor("blue"),
	)

	_, _ = g.AddEdge(n2, n3,
		goraffe.WithEdgeLabel("complete"),
		goraffe.WithEdgeStyle(goraffe.EdgeStyleDashed),
		goraffe.WithArrowHead(goraffe.ArrowDot),
	)

	fmt.Println(g.String())
	// Output:
	// digraph "Workflow" {
	// 	"A" [label="Start"];
	// 	"B" [label="Process"];
	// 	"C" [label="End"];
	// 	"A" -> "B" [color="blue", label="begin"];
	// 	"B" -> "C" [arrowhead="dot", label="complete", style="dashed"];
	// }
}

// Example_undirectedGraph demonstrates creating an undirected graph.
func Example_undirectedGraph() {
	// Create an undirected graph
	g := goraffe.NewGraph(goraffe.Undirected, goraffe.WithName("Network"))

	// Add nodes
	a := goraffe.NewNode("A")
	b := goraffe.NewNode("B")
	c := goraffe.NewNode("C")

	// Connect nodes (edges are bidirectional)
	_, _ = g.AddEdge(a, b)
	_, _ = g.AddEdge(b, c)
	_, _ = g.AddEdge(c, a)

	fmt.Println(g.String())
	// Output:
	// graph "Network" {
	// 	"A";
	// 	"B";
	// 	"C";
	// 	"A" -- "B";
	// 	"B" -- "C";
	// 	"C" -- "A";
	// }
}

// Example_graphLayout demonstrates customizing graph layout direction.
func Example_graphLayout() {
	// Create a graph with left-to-right layout
	g := goraffe.NewGraph(
		goraffe.Directed,
		goraffe.WithName("LRGraph"),
		goraffe.WithRankDir(goraffe.RankDirLR),
		goraffe.WithGraphLabel("System Flow"),
	)

	n1 := goraffe.NewNode("input", goraffe.WithLabel("Input"))
	n2 := goraffe.NewNode("process", goraffe.WithLabel("Process"))
	n3 := goraffe.NewNode("output", goraffe.WithLabel("Output"))

	_, _ = g.AddEdge(n1, n2)
	_, _ = g.AddEdge(n2, n3)

	fmt.Println(g.String())
	// Output:
	// digraph "LRGraph" {
	// 	label="System Flow";
	// 	rankdir="LR";
	// 	"input" [label="Input"];
	// 	"process" [label="Process"];
	// 	"output" [label="Output"];
	// 	"input" -> "process";
	// 	"process" -> "output";
	// }
}

// Example_defaultAttributes demonstrates setting default attributes for all nodes or edges.
func Example_defaultAttributes() {
	// Create a graph with default node attributes
	g := goraffe.NewGraph(
		goraffe.Directed,
		goraffe.WithDefaultNodeAttrs(
			goraffe.WithCircleShape(),
			goraffe.WithFillColor("lightyellow"),
		),
		goraffe.WithDefaultEdgeAttrs(
			goraffe.WithEdgeColor("gray"),
		),
	)

	// These nodes will inherit the default attributes
	n1 := goraffe.NewNode("A")
	n2 := goraffe.NewNode("B")

	// This node overrides the default fill color
	n3 := goraffe.NewNode("C", goraffe.WithFillColor("lightblue"))

	_, _ = g.AddEdge(n1, n2)
	_, _ = g.AddEdge(n2, n3)

	fmt.Println(g.String())
	// Output:
	// digraph {
	// 	node [fillcolor="lightyellow", shape="circle", style="filled"];
	// 	edge [color="gray"];
	// 	"A";
	// 	"B";
	// 	"C" [fillcolor="lightblue", style="filled"];
	// 	"A" -> "B";
	// 	"B" -> "C";
	// }
}

// ExampleNewGraph demonstrates creating different types of graphs.
func ExampleNewGraph() {
	// Create a simple directed graph
	directed := goraffe.NewGraph(goraffe.Directed)
	fmt.Println("Directed:", directed.IsDirected())

	// Create an undirected graph
	undirected := goraffe.NewGraph(goraffe.Undirected)
	fmt.Println("Undirected:", undirected.IsDirected())

	// Create a strict directed graph (no duplicate edges)
	strict := goraffe.NewGraph(goraffe.Directed, goraffe.Strict)
	fmt.Println("Strict:", strict.IsStrict())

	// Output:
	// Directed: true
	// Undirected: false
	// Strict: true
}

// ExampleNewNode demonstrates creating nodes with various attributes.
func ExampleNewNode() {
	// Basic node
	n1 := goraffe.NewNode("A")
	fmt.Println(n1.ID())

	// Node with label and shape
	n2 := goraffe.NewNode("B",
		goraffe.WithLabel("Node B"),
		goraffe.WithBoxShape(),
	)
	fmt.Println(n2.Attrs().Label())
	fmt.Println(n2.Attrs().Shape())

	// Output:
	// A
	// Node B
	// box
}

// ExampleGraph_AddEdge demonstrates adding edges between nodes.
func ExampleGraph_AddEdge() {
	g := goraffe.NewGraph(goraffe.Directed)

	// Create nodes
	start := goraffe.NewNode("start")
	end := goraffe.NewNode("end")

	// Add an edge
	edge, _ := g.AddEdge(start, end, goraffe.WithEdgeLabel("flow"))

	fmt.Println("From:", edge.From().ID())
	fmt.Println("To:", edge.To().ID())
	fmt.Println("Label:", edge.Attrs().Label())

	// Output:
	// From: start
	// To: end
	// Label: flow
}

// Example_recordLabel demonstrates creating a record label with ports.
// Record labels provide a simpler alternative to HTML tables for structured nodes.
// They support fields and nested groups with port connections.
func Example_recordLabel() {
	g := goraffe.NewGraph(goraffe.Directed)

	// Create fields with ports for edge connections
	inputPort := goraffe.Field("input").Port("in")
	outputPort := goraffe.Field("output").Port("out")

	// Create a record label with multiple fields and a nested group
	processLabel := goraffe.Record(
		goraffe.Field("Process Node"),
		goraffe.FieldGroup(
			inputPort,
			goraffe.Field("data"),
			outputPort,
		),
	)

	// Create a node with the record label
	processNode := goraffe.NewNode("process", goraffe.WithRecordLabel(processLabel))

	// Create source and destination nodes
	source := goraffe.NewNode("source", goraffe.WithLabel("Data Source"))
	sink := goraffe.NewNode("sink", goraffe.WithLabel("Data Sink"))

	// Connect edges to specific ports on the record node
	_, _ = g.AddEdge(source, processNode,
		goraffe.ToPort(inputPort.GetPort()),
		goraffe.WithEdgeLabel("send"),
	)
	_, _ = g.AddEdge(processNode, sink,
		goraffe.FromPort(outputPort.GetPort()),
		goraffe.WithEdgeLabel("receive"),
	)

	fmt.Println(g.String())
	// Output:
	// digraph {
	// 	"source" [label="Data Source"];
	// 	"process" [label="Process Node | { <in> input | data | <out> output }", shape="record"];
	// 	"sink" [label="Data Sink"];
	// 	"source" -> "process":"in" [label="send"];
	// 	"process":"out" -> "sink" [label="receive"];
	// }
}

// Example_htmlTableLabel demonstrates creating an HTML table label with ports.
// HTML tables are used for complex node labels in Graphviz with structured data,
// ports for edge connections, and rich formatting.
func Example_htmlTableLabel() {
	g := goraffe.NewGraph(goraffe.Directed)

	// Create cells with ports for edge connections
	idCell := goraffe.Cell(goraffe.Text("12345")).Port("id")
	nameCell := goraffe.Cell(goraffe.Text("Alice")).Port("name")

	// Create an HTML table label with a header and data rows
	table := goraffe.HTMLTable(
		goraffe.Row(
			goraffe.Cell(goraffe.Text("User Record").Bold()).
				Port("title").
				ColSpan(2).
				BgColor("lightblue").
				Align(goraffe.AlignCenter),
		),
		goraffe.Row(
			goraffe.Cell(goraffe.Text("ID:").Bold()),
			idCell,
		),
		goraffe.Row(
			goraffe.Cell(goraffe.Text("Name:").Bold()),
			nameCell,
		),
		goraffe.Row(
			goraffe.Cell(goraffe.Text("Status:").Bold()),
			goraffe.Cell(goraffe.Text("Active").Italic()),
		),
	).Border(1).CellBorder(1).CellSpacing(0).CellPadding(4)

	// Create a node with the HTML table label
	user := goraffe.NewNode("user", goraffe.WithHTMLLabel(table))

	// Create other nodes
	database := goraffe.NewNode("db", goraffe.WithLabel("Database"))
	display := goraffe.NewNode("display", goraffe.WithLabel("Display"))

	// Connect edges to specific ports on the HTML table
	_, _ = g.AddEdge(database, user,
		goraffe.ToPort(idCell.GetPort()),
		goraffe.WithEdgeLabel("fetch"),
	)
	_, _ = g.AddEdge(user, display,
		goraffe.FromPort(nameCell.GetPort()),
		goraffe.WithEdgeLabel("show"),
	)

	fmt.Println(g.String())
	// Output:
	// digraph {
	// 	"db" [label="Database"];
	// 	"user" [label=<<table border="1" cellborder="1" cellspacing="0" cellpadding="4"><tr><td port="title" colspan="2" bgcolor="lightblue" align="center"><b>User Record</b></td></tr><tr><td><b>ID:</b></td><td port="id">12345</td></tr><tr><td><b>Name:</b></td><td port="name">Alice</td></tr><tr><td><b>Status:</b></td><td><i>Active</i></td></tr></table>>];
	// 	"display" [label="Display"];
	// 	"db" -> "user":"id" [label="fetch"];
	// 	"user":"name" -> "display" [label="show"];
	// }
}
