package goraffe

type GraphOption interface {
	applyGraph(*Graph)
}

type graphOptionFunc func(*Graph)

func (f graphOptionFunc) applyGraph(g *Graph) {
	f(g)
}

var Directed GraphOption = graphOptionFunc(func(g *Graph) {
	g.directed = true
})

var Undirected GraphOption = graphOptionFunc(func(g *Graph) {
	g.directed = false
})

var Strict GraphOption = graphOptionFunc(func(g *Graph) {
	g.strict = true
})

func newGraphOption(fn func(*Graph)) GraphOption {
	return graphOptionFunc(fn)
}

func WithGraphLabel(l string) GraphOption {
	return newGraphOption(func(g *Graph) {
		g.attrs.label = &l
	})
}

func WithRankDir(d RankDir) GraphOption {
	return newGraphOption(func(g *Graph) {
		g.attrs.rankDir = &d
	})
}

func WithBgColor(c string) GraphOption {
	return newGraphOption(func(g *Graph) {
		g.attrs.bgColor = &c
	})
}

func WithGraphFontName(n string) GraphOption {
	return newGraphOption(func(g *Graph) {
		g.attrs.fontName = &n
	})
}

func WithGraphFontSize(s float64) GraphOption {
	return newGraphOption(func(g *Graph) {
		g.attrs.fontSize = &s
	})
}

func WithSplines(s SplineType) GraphOption {
	return newGraphOption(func(g *Graph) {
		g.attrs.splines = &s
	})
}

func WithNodeSep(s float64) GraphOption {
	return newGraphOption(func(g *Graph) {
		g.attrs.nodeSep = &s
	})
}

func WithRankSep(s float64) GraphOption {
	return newGraphOption(func(g *Graph) {
		g.attrs.rankSep = &s
	})
}

func WithCompound(c bool) GraphOption {
	return newGraphOption(func(g *Graph) {
		g.attrs.compound = &c
	})
}

func WithDefaultNodeAttrs(options ...NodeOption) GraphOption {
	return newGraphOption(func(g *Graph) {
		nodeAttrs := &NodeAttributes{}

		for _, option := range options {
			option.applyNode(nodeAttrs)
		}

		g.defaultNodeAttrs = nodeAttrs
	})
}

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
