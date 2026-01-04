package goraffe

// GraphOption is a functional option for configuring graph properties.
// Options can be passed to NewGraph to configure the graph type and attributes.
type GraphOption interface {
	applyGraph(*Graph)
}

type graphOptionFunc func(*Graph)

func (f graphOptionFunc) applyGraph(g *Graph) {
	f(g)
}

// Directed creates a directed graph where edges have arrows indicating direction.
// Use with NewGraph to create a directed graph (default is undirected).
//
// Example:
//
//	g := goraffe.NewGraph(goraffe.Directed)
var Directed GraphOption = graphOptionFunc(func(g *Graph) {
	g.directed = true
})

// Undirected creates an undirected graph where edges are bidirectional lines.
// This is the default behavior, so this option is only needed for clarity.
//
// Example:
//
//	g := goraffe.NewGraph(goraffe.Undirected)
var Undirected GraphOption = graphOptionFunc(func(g *Graph) {
	g.directed = false
})

// Strict prevents duplicate edges between the same pair of nodes.
// In strict graphs, only one edge is allowed between any two nodes.
//
// Example:
//
//	g := goraffe.NewGraph(goraffe.Directed, goraffe.Strict)
var Strict GraphOption = graphOptionFunc(func(g *Graph) {
	g.strict = true
})

func newGraphOption(fn func(*Graph)) GraphOption {
	return graphOptionFunc(fn)
}

// WithName sets the name of the graph, which appears in the DOT output header.
//
// Example:
//
//	g := goraffe.NewGraph(goraffe.WithName("MyGraph"))
func WithName(n string) GraphOption {
	return newGraphOption(func(g *Graph) {
		g.name = n
	})
}

// WithGraphLabel sets a text label for the entire graph.
// The label is typically displayed at the top or bottom of the rendered graph.
//
// Example:
//
//	g := goraffe.NewGraph(goraffe.WithGraphLabel("System Architecture"))
func WithGraphLabel(l string) GraphOption {
	return newGraphOption(func(g *Graph) {
		g.attrs.label = &l
	})
}

// WithRankDir sets the direction of graph layout (top-to-bottom, left-to-right, etc.).
//
// Example:
//
//	g := goraffe.NewGraph(goraffe.WithRankDir(goraffe.RankDirLR))
func WithRankDir(d RankDir) GraphOption {
	return newGraphOption(func(g *Graph) {
		g.attrs.rankDir = &d
	})
}

// WithBgColor sets the background color of the graph canvas.
// Accepts color names (e.g., "white", "lightgray") or hex values (e.g., "#F0F0F0").
//
// Example:
//
//	g := goraffe.NewGraph(goraffe.WithBgColor("lightgray"))
func WithBgColor(c string) GraphOption {
	return newGraphOption(func(g *Graph) {
		g.attrs.bgColor = &c
	})
}

// WithGraphFontName sets the default font family for graph labels.
//
// Example:
//
//	g := goraffe.NewGraph(goraffe.WithGraphFontName("Arial"))
func WithGraphFontName(n string) GraphOption {
	return newGraphOption(func(g *Graph) {
		g.attrs.fontName = &n
	})
}

// WithGraphFontSize sets the default font size for graph labels in points.
//
// Example:
//
//	g := goraffe.NewGraph(goraffe.WithGraphFontSize(16.0))
func WithGraphFontSize(s float64) GraphOption {
	return newGraphOption(func(g *Graph) {
		g.attrs.fontSize = &s
	})
}

// WithSplines controls how edges are routed between nodes.
//
// Example:
//
//	g := goraffe.NewGraph(goraffe.WithSplines(goraffe.SplineOrtho))
func WithSplines(s SplineType) GraphOption {
	return newGraphOption(func(g *Graph) {
		g.attrs.splines = &s
	})
}

// WithNodeSep sets the minimum space between nodes at the same rank.
// Default is 0.25 inches in Graphviz.
//
// Example:
//
//	g := goraffe.NewGraph(goraffe.WithNodeSep(0.5))
func WithNodeSep(s float64) GraphOption {
	return newGraphOption(func(g *Graph) {
		g.attrs.nodeSep = &s
	})
}

// WithRankSep sets the minimum vertical space between ranks.
// Default is 0.5 inches in dot layout, 1.0 in twopi.
//
// Example:
//
//	g := goraffe.NewGraph(goraffe.WithRankSep(1.0))
func WithRankSep(s float64) GraphOption {
	return newGraphOption(func(g *Graph) {
		g.attrs.rankSep = &s
	})
}

// WithCompound enables compound mode, allowing edges between clusters.
//
// Example:
//
//	g := goraffe.NewGraph(goraffe.WithCompound(true))
func WithCompound(c bool) GraphOption {
	return newGraphOption(func(g *Graph) {
		g.attrs.compound = &c
	})
}

// WithDefaultNodeAttrs sets default attributes applied to all nodes in the graph.
// Individual node attributes can override these defaults.
//
// Example:
//
//	g := goraffe.NewGraph(
//	    goraffe.WithDefaultNodeAttrs(
//	        goraffe.WithCircleShape(),
//	        goraffe.WithFillColor("lightblue"),
//	    ),
//	)
func WithDefaultNodeAttrs(options ...NodeOption) GraphOption {
	return newGraphOption(func(g *Graph) {
		nodeAttrs := &NodeAttributes{}

		for _, option := range options {
			option.applyNode(nodeAttrs)
		}

		g.defaultNodeAttrs = nodeAttrs
	})
}

// WithDefaultEdgeAttrs sets default attributes applied to all edges in the graph.
// Individual edge attributes can override these defaults.
//
// Example:
//
//	g := goraffe.NewGraph(
//	    goraffe.WithDefaultEdgeAttrs(
//	        goraffe.WithEdgeColor("gray"),
//	        goraffe.WithArrowHead(goraffe.ArrowDot),
//	    ),
//	)
func WithDefaultEdgeAttrs(options ...EdgeOption) GraphOption {
	return newGraphOption(func(g *Graph) {
		nodeAttrs := &EdgeAttributes{}

		for _, option := range options {
			option.applyEdge(nodeAttrs)
		}

		g.defaultEdgeAttrs = nodeAttrs
	})
}

// WithGraphAttribute sets a custom attribute on a graph.
// This is an escape hatch for Graphviz attributes that don't have typed options.
//
// Example:
//
//	g := NewGraph(
//	    WithGraphLabel("My Graph"),
//	    WithGraphAttribute("ratio", "fill"),
//	    WithGraphAttribute("concentrate", "true"),
//	)
func WithGraphAttribute(k, v string) GraphOption {
	return newGraphOption(func(a *Graph) {
		a.attrs.setCustom(k, v)
	})
}
