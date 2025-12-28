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
