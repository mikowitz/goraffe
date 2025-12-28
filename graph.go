package goraffe

type Graph struct {
	name             string
	directed, strict bool
}

func NewGraph() *Graph {
	return &Graph{}
}

func (g *Graph) IsDirected() bool {
	return g.directed
}

func (g *Graph) IsStrict() bool {
	return g.strict
}

func (g *Graph) Name() string {
	return g.name
}
