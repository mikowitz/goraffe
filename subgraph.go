// ABOUTME: Implements subgraph support for Graphviz DOT graphs.
// ABOUTME: Subgraphs can group nodes and edges, with cluster support for visual grouping.
package goraffe

import "strings"

// Subgraph represents a subgraph within a Graph.
// Subgraphs can be used to group nodes and edges together.
// If the name starts with "cluster", it will be rendered as a visual cluster in Graphviz.
type Subgraph struct {
	name   string
	nodes  map[string]*Node
	edges  []*Edge
	parent *Graph
}

// Name returns the name of the subgraph.
func (sg *Subgraph) Name() string {
	return sg.name
}

// IsCluster returns true if the subgraph name starts with "cluster".
// Cluster subgraphs are rendered with a bounding box in Graphviz.
func (sg *Subgraph) IsCluster() bool {
	return strings.HasPrefix(sg.name, "cluster")
}

// AddNode adds a node to the subgraph and also adds it to the parent graph.
// This ensures that nodes in subgraphs are also part of the overall graph structure.
func (sg *Subgraph) AddNode(n *Node) error {
	if n == nil {
		return ErrNilNode
	}

	sg.nodes[n.ID()] = n
	return sg.parent.AddNode(n)
}

// Nodes returns all nodes in the subgraph.
// The returned slice contains nodes in no particular order.
func (sg *Subgraph) Nodes() []*Node {
	nodes := make([]*Node, 0, len(sg.nodes))
	for _, node := range sg.nodes {
		nodes = append(nodes, node)
	}
	return nodes
}

// AddEdge creates and adds an edge between two nodes, delegating to the parent graph.
// This ensures edges are managed at the graph level while allowing subgraph-scoped edge creation.
func (sg *Subgraph) AddEdge(from, to *Node, opts ...EdgeOption) (*Edge, error) {
	edge, err := sg.parent.AddEdge(from, to, opts...)
	if err != nil {
		return nil, err
	}
	sg.edges = append(sg.edges, edge)
	return edge, nil
}
