package goraffe

import (
	"errors"
	"fmt"
	"io"
	"strings"
)

// ErrNilNode is returned when a nil node is passed to a function that requires a non-nil node.
var ErrNilNode = errors.New("node cannot be nil")

// Graph represents a Graphviz graph structure that can contain nodes and edges.
// A graph can be directed or undirected, and optionally strict (preventing duplicate edges).
// Use NewGraph to create a new graph instance.
type Graph struct {
	name             string
	directed, strict bool
	nodeOrder        []*Node
	nodes            map[string]int
	edges            []*Edge
	attrs            *GraphAttributes
	defaultNodeAttrs *NodeAttributes
	defaultEdgeAttrs *EdgeAttributes
}

// NewGraph creates a new graph with the specified options.
// By default, graphs are undirected and non-strict. Use the Directed, Undirected,
// and Strict options to configure the graph type.
//
// Example:
//
//	g := goraffe.NewGraph(goraffe.Directed, goraffe.WithName("MyGraph"))
func NewGraph(options ...GraphOption) *Graph {
	g := &Graph{
		nodeOrder:        make([]*Node, 0),
		nodes:            make(map[string]int),
		edges:            make([]*Edge, 0),
		attrs:            &GraphAttributes{},
		defaultNodeAttrs: &NodeAttributes{},
		defaultEdgeAttrs: &EdgeAttributes{},
	}

	for _, option := range options {
		option.applyGraph(g)
	}

	return g
}

// IsDirected returns true if the graph is directed, false if undirected.
// Directed graphs use arrows (->) while undirected graphs use lines (--) in DOT output.
func (g *Graph) IsDirected() bool {
	return g.directed
}

// IsStrict returns true if the graph is strict.
// Strict graphs do not allow duplicate edges between the same pair of nodes.
func (g *Graph) IsStrict() bool {
	return g.strict
}

// Name returns the name of the graph.
// Returns empty string if no name was set via WithName option.
func (g *Graph) Name() string {
	return g.name
}

// AddNode adds a node to the graph. If a node with the same ID already exists,
// it will be replaced in place, preserving its original position in the node order.
// This ensures that the insertion order of nodes is maintained for DOT output.
// Returns an error if the node is nil.
func (g *Graph) AddNode(n *Node) error {
	if n == nil {
		return fmt.Errorf("could not add node: %w", ErrNilNode)
	}

	if idx, exists := g.nodes[n.ID()]; exists {
		// Replace at existing position
		g.nodeOrder[idx] = n
	} else {
		// Add new node
		g.nodes[n.ID()] = len(g.nodeOrder)
		g.nodeOrder = append(g.nodeOrder, n)
	}
	return nil
}

// GetNode retrieves a node from the graph by its ID.
// Returns nil if no node with the given ID exists in the graph.
func (g *Graph) GetNode(id string) *Node {
	if idx, exists := g.nodes[id]; exists {
		return g.nodeOrder[idx]
	}

	return nil
}

// Nodes returns all nodes in the graph in insertion order.
// The returned slice should not be modified.
func (g *Graph) Nodes() []*Node {
	return g.nodeOrder
}

// AddEdge creates and adds an edge from one node to another with optional attributes.
// If either node is not already in the graph, it will be automatically added.
// Returns the created edge and an error if either node is nil.
//
// Example:
//
//	n1 := goraffe.NewNode("A")
//	n2 := goraffe.NewNode("B")
//	e, err := g.AddEdge(n1, n2, goraffe.WithEdgeLabel("connects"))
func (g *Graph) AddEdge(from, to *Node, options ...EdgeOption) (*Edge, error) {
	errs := []error{}
	if from == nil {
		// return nil, fmt.Errorf("could not add edge with nil source node: %w", ErrNilNode)
		errs = append(errs, fmt.Errorf("edge requires source node: %w", ErrNilNode))
	}
	if to == nil {
		errs = append(errs, fmt.Errorf("edge requires target node: %w", ErrNilNode))
	}
	if len(errs) > 0 {
		return nil, errors.Join(errs...)
	}

	attrs := &EdgeAttributes{}

	for _, option := range options {
		option.applyEdge(attrs)
	}

	edge := &Edge{
		from:  from,
		to:    to,
		attrs: attrs,
	}

	if _, exists := g.nodes[from.ID()]; !exists {
		err := g.AddNode(from)
		if err != nil {
			return nil, err
		}
	}

	if _, exists := g.nodes[to.ID()]; !exists {
		err := g.AddNode(to)
		if err != nil {
			return nil, err
		}
	}

	g.edges = append(g.edges, edge)

	return edge, nil
}

// Edges returns all edges in the graph in insertion order.
// The returned slice should not be modified.
func (g *Graph) Edges() []*Edge {
	return g.edges
}

// Attrs returns the graph's attributes (label, rank direction, colors, etc.).
// The returned attributes can be modified to change graph-level properties.
func (g *Graph) Attrs() *GraphAttributes {
	return g.attrs
}

// DefaultNodeAttrs returns the default attributes applied to all nodes in the graph.
// These defaults can be overridden by individual node attributes.
func (g *Graph) DefaultNodeAttrs() *NodeAttributes {
	return g.defaultNodeAttrs
}

// DefaultEdgeAttrs returns the default attributes applied to all edges in the graph.
// These defaults can be overridden by individual edge attributes.
func (g *Graph) DefaultEdgeAttrs() *EdgeAttributes {
	return g.defaultEdgeAttrs
}

// String returns the DOT representation of the graph.
// The output is valid Graphviz DOT format that can be rendered with dot, neato, etc.
// Note: Currently only outputs nodes; edge output is not yet implemented.
func (g *Graph) String() string {
	builder := strings.Builder{}

	if g.strict {
		builder.WriteString("strict ")
	}

	if g.directed {
		builder.WriteString("digraph")
	} else {
		builder.WriteString("graph")
	}

	if g.name != "" {
		builder.WriteString(fmt.Sprintf(" %s", g.name))
	}

	builder.WriteString(" {\n")

	for _, node := range g.nodeOrder {
		builder.WriteString(fmt.Sprintf("\t%s;\n", node))
	}

	builder.WriteString("}")

	return builder.String()
}

// WriteDOT writes the graph's DOT representation to the given writer.
// This is a convenience method that writes the output of String() to a writer.
// Returns any error encountered during writing.
//
// Example:
//
//	f, _ := os.Create("graph.dot")
//	defer f.Close()
//	g.WriteDOT(f)
func (g *Graph) WriteDOT(w io.Writer) error {
	_, err := w.Write([]byte(g.String()))

	return err
}
