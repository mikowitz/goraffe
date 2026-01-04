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
// # Current Status
//
// Note: This library is in active development. The graph building API is functional,
// but DOT parsing and rendering features are not yet implemented.
//
// # Requirements
//
// Graphviz must be installed on your system for rendering functionality to work.
// You can install Graphviz from https://graphviz.org/download/ or use your system's
// package manager (e.g., brew install graphviz, apt-get install graphviz).
package goraffe
