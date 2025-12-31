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
