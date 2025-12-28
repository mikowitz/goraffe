package goraffe

type Graph struct {
	name             string
	directed, strict bool
	nodeOrder        []*Node
	nodes            map[string]int
	edges            []*Edge
}

func NewGraph(options ...GraphOption) *Graph {
	g := &Graph{
		nodeOrder: make([]*Node, 0),
		nodes:     make(map[string]int),
		edges:     make([]*Edge, 0),
	}

	for _, option := range options {
		option.applyGraph(g)
	}

	return g
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

// AddNode adds a node to the graph. If a node with the same ID already exists,
// it will be replaced in place, preserving its original position in the node order.
// This ensures that the insertion order of nodes is maintained for DOT output.
func (g *Graph) AddNode(n *Node) {
	if idx, exists := g.nodes[n.ID()]; exists {
		// Replace at existing position
		g.nodeOrder[idx] = n
	} else {
		// Add new node
		g.nodes[n.ID()] = len(g.nodeOrder)
		g.nodeOrder = append(g.nodeOrder, n)
	}
}

func (g *Graph) GetNode(id string) *Node {
	if idx, exists := g.nodes[id]; exists {
		return g.nodeOrder[idx]
	}
	return nil
}

func (g *Graph) Nodes() []*Node {
	return g.nodeOrder
}

func (g *Graph) AddEdge(from, to *Node) *Edge {
	edge := &Edge{from: from, to: to}

	if _, exists := g.nodes[from.ID()]; !exists {
		g.AddNode(from)
	}
	if _, exists := g.nodes[to.ID()]; !exists {
		g.AddNode(to)
	}

	g.edges = append(g.edges, edge)

	return edge
}

func (g *Graph) Edges() []*Edge {
	return g.edges
}
