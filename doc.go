// Package goraffe provides a unified API for creating, parsing, and rendering Graphviz graphs.
//
// This library allows you to programmatically build graph structures, parse existing DOT format
// graphs, and render them to various output formats using Graphviz.
//
// # Quick Start
//
// Create a simple directed graph:
//
//	package main
//
//	import (
//	    "fmt"
//	    "github.com/mikowitz/goraffe"
//	)
//
//	func main() {
//	    // Create a directed graph
//	    g := goraffe.NewGraph(goraffe.Directed, goraffe.WithName("MyGraph"))
//
//	    // Add nodes with attributes
//	    start := goraffe.NewNode("start", goraffe.WithLabel("Start"), goraffe.WithCircleShape())
//	    end := goraffe.NewNode("end", goraffe.WithLabel("End"), goraffe.WithBoxShape())
//
//	    // Connect nodes with an edge
//	    g.AddEdge(start, end, goraffe.WithEdgeLabel("flow"))
//
//	    // Output DOT format
//	    fmt.Println(g.String())
//	}
//
// # Graph Types
//
// Goraffe supports both directed and undirected graphs:
//
//	directed := goraffe.NewGraph(goraffe.Directed)    // Edges have arrows
//	undirected := goraffe.NewGraph(goraffe.Undirected) // Edges are lines
//	strict := goraffe.NewGraph(goraffe.Strict)        // Prevents duplicate edges
//
// # Nodes
//
// Nodes can be customized with various visual attributes:
//
//	n := goraffe.NewNode("A",
//	    goraffe.WithLabel("Node A"),
//	    goraffe.WithCircleShape(),
//	    goraffe.WithColor("red"),
//	    goraffe.WithFillColor("lightblue"),
//	    goraffe.WithFontSize(14.0),
//	)
//
// # Edges
//
// Edges connect nodes and can have their own attributes:
//
//	e := g.AddEdge(n1, n2,
//	    goraffe.WithEdgeLabel("connects"),
//	    goraffe.WithEdgeColor("blue"),
//	    goraffe.WithEdgeStyle(goraffe.EdgeStyleDashed),
//	    goraffe.WithArrowHead(goraffe.ArrowDot),
//	)
//
// # Graph Attributes
//
// Customize the overall graph appearance:
//
//	g := goraffe.NewGraph(
//	    goraffe.Directed,
//	    goraffe.WithRankDir(goraffe.RankDirLR),  // Left to right layout
//	    goraffe.WithBgColor("white"),
//	    goraffe.WithGraphLabel("My System"),
//	)
//
// # DOT Output
//
// Convert your graph to DOT format for rendering with Graphviz:
//
//	dotString := g.String()
//	// Or write directly to a file
//	f, _ := os.Create("graph.dot")
//	g.WriteDOT(f)
//
// # Rendering
//
// Render graphs to various output formats using Graphviz:
//
//	// Render to PNG file
//	g.RenderToFile(goraffe.PNG, "graph.png")
//
//	// Render to bytes
//	svgData, _ := g.RenderBytes(goraffe.SVG)
//
//	// Render to io.Writer with custom layout
//	var buf bytes.Buffer
//	g.Render(goraffe.PNG, &buf, goraffe.WithLayout(goraffe.LayoutNeato))
//
// Supported formats: PNG, SVG, PDF, DOT
// Supported layouts: dot, neato, fdp, sfdp, twopi, circo, osage, patchwork
//
// # Parsing
//
// Parse existing DOT format files:
//
//	// Parse from string
//	g, _ := goraffe.ParseString("digraph G { A -> B }")
//
//	// Parse from io.Reader
//	f, _ := os.Open("graph.dot")
//	g, _ := goraffe.Parse(f)
//
//	// Parse from file path
//	g, _ := goraffe.ParseFile("graph.dot")
//
// # Requirements
//
// Graphviz must be installed on your system for rendering functionality to work.
// You can install Graphviz from https://graphviz.org/download/ or use your system's
// package manager (e.g., brew install graphviz, apt-get install graphviz).
package goraffe
