# Goraffe

A Go library for building and rendering Graphviz graphs with a clean, type-safe API.

[![MIT License](https://img.shields.io/badge/license-MIT-blue.svg)](LICENSE)
[![Go Reference](https://pkg.go.dev/badge/github.com/mikowitz/goraffe.svg)](https://pkg.go.dev/github.com/mikowitz/goraffe)

## Overview

Goraffe provides an ergonomic, type-safe API for creating Graphviz graphs in Go. It combines graph construction, DOT format generation, and rendering capabilities into a unified library, with compile-time safety for common attributes and escape hatches for advanced use cases.

## Status

✨ **Active Development** - Core features complete, parsing in progress

**What works:**

- ✅ Graph, Node, and Edge construction
- ✅ Type-safe attributes (shapes, colors, styles, etc.)
- ✅ Functional options pattern for configuration
- ✅ Graph-level default attributes
- ✅ Custom attribute escape hatches
- ✅ Complete DOT format output
- ✅ HTML and record labels
- ✅ Subgraphs and clusters
- ✅ Rendering via Graphviz CLI (PNG, SVG, PDF, DOT)
- ✅ Multiple layout engines (dot, neato, fdp, sfdp, twopi, circo, osage, patchwork)

## Installation

```bash
go get github.com/mikowitz/goraffe
```

**Requirements:**

- Go 1.23 or later
- Graphviz (for rendering - install via `brew install graphviz`, `apt-get install graphviz`, or [graphviz.org](https://graphviz.org/download/))

## Quick Start

### Creating and Rendering a Graph

```go
package main

import (
    "fmt"
    "log"
    "github.com/mikowitz/goraffe"
)

func main() {
    // Create a directed graph
    g := goraffe.NewGraph(goraffe.Directed,
        goraffe.WithGraphLabel("My Workflow"),
        goraffe.WithRankDir(goraffe.RankDirLR),
    )

    // Create nodes with attributes
    start := goraffe.NewNode("start",
        goraffe.WithLabel("Start"),
        goraffe.WithBoxShape(),
        goraffe.WithFillColor("lightblue"),
    )

    end := goraffe.NewNode("end",
        goraffe.WithLabel("End"),
        goraffe.WithCircleShape(),
        goraffe.WithFillColor("lightgreen"),
    )

    // Add nodes and edges
    g.AddNode(start)
    g.AddNode(end)
    g.AddEdge(start, end,
        goraffe.WithEdgeLabel("proceed"),
        goraffe.WithEdgeColor("blue"),
    )

    // Output DOT format
    fmt.Println(g.String())

    // Render to PNG file
    if err := g.RenderToFile(goraffe.PNG, "workflow.png"); err != nil {
        log.Fatal(err)
    }

    // Render to SVG with custom layout
    if err := g.RenderToFile(goraffe.SVG, "workflow.svg",
        goraffe.WithLayout(goraffe.LayoutNeato)); err != nil {
        log.Fatal(err)
    }

    // Render to bytes (useful for web servers)
    pngBytes, err := g.RenderBytes(goraffe.PNG)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Printf("Generated PNG: %d bytes\n", len(pngBytes))
}
```

## Features

### Graph Construction

```go
// Directed graph (default)
g := goraffe.NewGraph(goraffe.Directed)

// Undirected graph
g := goraffe.NewGraph(goraffe.Undirected)

// Strict graph (no duplicate edges)
g := goraffe.NewGraph(goraffe.Directed, goraffe.Strict)

// With graph attributes
g := goraffe.NewGraph(
    goraffe.Directed,
    goraffe.WithRankDir(goraffe.RankDirLR),
    goraffe.WithGraphLabel("My Graph"),
)
```

### Node Attributes

```go
n := goraffe.NewNode("id",
    goraffe.WithLabel("Display Name"),
    goraffe.WithBoxShape(),           // or WithCircleShape(), WithDiamondShape(), etc.
    goraffe.WithColor("red"),
    goraffe.WithFillColor("pink"),
    goraffe.WithFontName("Arial"),
    goraffe.WithFontSize(12.0),
)
```

### Edge Attributes

```go
g.AddEdge(n1, n2,
    goraffe.WithEdgeLabel("connects"),
    goraffe.WithEdgeColor("blue"),
    goraffe.WithEdgeStyle(goraffe.EdgeStyleDashed),
    goraffe.WithArrowHead(goraffe.ArrowDot),
    goraffe.WithWeight(2.0),
)
```

### Default Attributes

```go
g := goraffe.NewGraph(
    goraffe.Directed,
    goraffe.WithDefaultNodeAttrs(
        goraffe.WithBoxShape(),
        goraffe.WithFontName("Arial"),
    ),
    goraffe.WithDefaultEdgeAttrs(
        goraffe.WithEdgeColor("gray"),
    ),
)
```

### Reusable Attribute Templates

```go
// Create reusable attribute sets
primaryStyle := goraffe.NodeAttributes{
    Shape:     goraffe.ShapeBox,
    FillColor: "lightblue",
    FontName:  "Arial",
}

n1 := goraffe.NewNode("a", primaryStyle)
n2 := goraffe.NewNode("b", primaryStyle, goraffe.WithLabel("Custom"))
```

### Custom Attributes (Escape Hatch)

```go
// For Graphviz attributes not yet supported
n := goraffe.NewNode("id",
    goraffe.WithBoxShape(),
    goraffe.WithNodeAttribute("peripheries", "2"),
    goraffe.WithNodeAttribute("tooltip", "Hover text"),
)
```

### Rendering Graphs

```go
// Render to file in various formats
g.RenderToFile(goraffe.PNG, "graph.png")
g.RenderToFile(goraffe.SVG, "graph.svg")
g.RenderToFile(goraffe.PDF, "graph.pdf")

// Render to bytes (useful for web servers, APIs, etc.)
pngData, err := g.RenderBytes(goraffe.PNG)
svgData, err := g.RenderBytes(goraffe.SVG)

// Render with custom layout engine
g.RenderToFile(goraffe.PNG, "graph.png",
    goraffe.WithLayout(goraffe.LayoutNeato))  // Spring model layout
g.RenderToFile(goraffe.PNG, "graph.png",
    goraffe.WithLayout(goraffe.LayoutCirco))  // Circular layout

// Available formats: PNG, SVG, PDF, DOT
// Available layouts: LayoutDot, LayoutNeato, LayoutFdp, LayoutSfdp,
//                    LayoutTwopi, LayoutCirco, LayoutOsage, LayoutPatchwork

// Render to io.Writer
var buf bytes.Buffer
g.Render(goraffe.PNG, &buf, goraffe.WithLayout(goraffe.LayoutDot))
```

### Error Handling

```go
err := g.RenderToFile(goraffe.PNG, "output.png")
if err != nil {
    if errors.Is(err, goraffe.ErrGraphvizNotFound) {
        log.Fatal("Graphviz is not installed")
    }

    // Get detailed error information
    if renderErr, ok := err.(*goraffe.RenderError); ok {
        log.Printf("Rendering failed (exit code %d): %s",
            renderErr.ExitCode, renderErr.Stderr)
    }
}

// Check Graphviz version
version, err := goraffe.GraphvizVersion()
if err == nil {
    fmt.Printf("Using Graphviz: %s\n", version)
}
```

## Documentation

- [Package Documentation](https://pkg.go.dev/github.com/mikowitz/goraffe)

## Architecture

Goraffe uses a single-package design with a clean builder API:

- **Type-safe attributes**: Enums and constants for shapes, colors, styles, arrow types, etc.
- **Functional options**: Composable configuration using the options pattern
- **Pointer-based attributes**: Internal use of pointers to distinguish "not set" from "zero value"
- **Escape hatches**: Custom attribute support for advanced Graphviz features

## Development Status

The project is being built in phases:

- **Phase 1: Foundation** ✅ Complete (Graph, Node, Edge structs)
- **Phase 2: Attributes** ✅ Complete (Type-safe attributes and options)
- **Phase 3: DOT Generation** ✅ Complete (Full DOT output with all features)
- **Phase 4: Labels** ✅ Complete (HTML and record labels)
- **Phase 5: Subgraphs** ✅ Complete (Clusters and subgraphs)
- **Phase 6: Parsing** 🚧 In Progress (DOT format parser on parser branch)
- **Phase 7: Rendering** ✅ Complete (Graphviz CLI integration with multiple formats and layouts)

## Contributing

This project is in early development. Contributions are welcome, but the API is still evolving.

Before contributing:

1. Check [dev/todo.md](dev/todo.md) for current progress
2. Review [dev/spec.md](dev/spec.md) for design decisions
3. Ensure tests pass with 100% coverage: `go test -cover ./...`
4. Run linting: `golangci-lint run` (once configured)

## License

MIT License - see [LICENSE](LICENSE) file for details.

Copyright (c) 2026 Michael Berkowitz

## Acknowledgments

Goraffe aims to combine the best aspects of existing Go Graphviz libraries while addressing their limitations through:

- Unified API (building, parsing, and rendering in one package)
- Type-safe attributes with escape hatches
- First-class support for complex graph features
- Simple rendering via Graphviz CLI

Inspired by the Go community's excellent graph libraries and the power of Graphviz.

Built with [💜](https://songsaboutsnow.com) and [Go](https://go.dev).
