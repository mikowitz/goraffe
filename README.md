# Goraffe

A Go library for building and rendering Graphviz graphs with a clean, type-safe API.

[![MIT License](https://img.shields.io/badge/license-MIT-blue.svg)](LICENSE)
[![Go Reference](https://pkg.go.dev/badge/github.com/mikowitz/goraffe.svg)](https://pkg.go.dev/github.com/mikowitz/goraffe)

## Overview

Goraffe provides an ergonomic, type-safe API for creating Graphviz graphs in Go. It combines graph construction, DOT format generation, and rendering capabilities into a unified library, with compile-time safety for common attributes and escape hatches for advanced use cases.

## Status

⚠️ **Early Development** - Currently at ~30% implementation

**What works:**

- ✅ Graph, Node, and Edge construction
- ✅ Type-safe attributes (shapes, colors, styles, etc.)
- ✅ Functional options pattern for configuration
- ✅ Graph-level default attributes
- ✅ Custom attribute escape hatches
- ✅ Basic DOT format output (nodes only)

**What's coming:**

- 🚧 Complete DOT generation (edges, subgraphs, labels)
- 🚧 HTML and record labels
- 🚧 Subgraphs and clusters
- 🚧 Rank constraints
- 🚧 DOT parsing
- 🚧 Rendering via Graphviz CLI

## Installation

```bash
go get github.com/mikowitz/goraffe
```

**Requirements:**

- Go 1.23 or later
- Graphviz (for rendering, not yet implemented)

## Quick Start

```go
package main

import (
    "fmt"
    "github.com/mikowitz/goraffe"
)

func main() {
    // Create a directed graph
    g := goraffe.NewGraph(goraffe.Directed)

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
    edge := g.AddEdge(start, end,
        goraffe.WithEdgeLabel("proceed"),
        goraffe.WithEdgeColor("blue"),
    )

    // Output DOT format
    fmt.Println(g.String())
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
- **Phase 3: DOT Generation** 🚧 In Progress (30% complete)
- **Phase 4: Labels** 📋 Planned (HTML and record labels)
- **Phase 5: Subgraphs** 📋 Planned (Clusters and rank constraints)
- **Phase 6: Parsing** 📋 Planned (DOT format parser)
- **Phase 7: Rendering** 📋 Planned (Graphviz CLI integration)

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
