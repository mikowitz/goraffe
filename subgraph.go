// ABOUTME: Implements subgraph support for Graphviz DOT graphs.
// ABOUTME: Subgraphs can group nodes and edges, with cluster support for visual grouping.
package goraffe

import (
	"fmt"
	"strings"
)

// Subgraph represents a subgraph within a Graph.
// Subgraphs can be used to group nodes and edges together.
// If the name starts with "cluster", it will be rendered as a visual cluster in Graphviz.
//
// Cluster subgraphs (names starting with "cluster") support visual attributes like colors
// and fill colors. Regular subgraphs may have these attributes set but they typically won't
// be rendered by Graphviz.
type Subgraph struct {
	name      string
	nodes     map[string]*Node
	edges     []*Edge
	parent    *Graph
	attrs     *SubgraphAttributes
	subgraphs []*Subgraph
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

// Attrs returns the subgraph's attribute configuration.
// If attributes haven't been initialized yet, this creates and returns a new SubgraphAttributes.
func (sg *Subgraph) Attrs() *SubgraphAttributes {
	if sg.attrs == nil {
		sg.attrs = &SubgraphAttributes{}
	}
	return sg.attrs
}

// SetLabel sets the label for the subgraph.
// The label is displayed as text associated with the subgraph (typically visible for clusters).
func (sg *Subgraph) SetLabel(l string) {
	sg.Attrs().label = &l
}

// SetStyle sets the style for the subgraph (e.g., "filled", "dashed", "bold").
// Multiple styles can be comma-separated. Typically visible for cluster subgraphs.
func (sg *Subgraph) SetStyle(s string) {
	sg.Attrs().style = &s
}

// SetColor sets the border color for the subgraph.
// Typically only visible for cluster subgraphs.
func (sg *Subgraph) SetColor(c string) {
	sg.Attrs().color = &c
}

// SetFillColor sets the fill/background color for the subgraph.
// Typically only visible for cluster subgraphs.
// When using fillcolor with a cluster, you may also want to set style to "filled".
func (sg *Subgraph) SetFillColor(c string) {
	sg.Attrs().fillColor = &c
}

// SetAttribute sets a custom DOT attribute on the subgraph.
// This allows setting arbitrary Graphviz attributes not covered by dedicated setter methods.
func (sg *Subgraph) SetAttribute(key, value string) {
	sg.Attrs().setCustom(key, value)
}

// Subgraph creates a nested subgraph within this subgraph.
// The nested subgraph will reference the root graph for node tracking, ensuring all nodes
// are registered at the graph level while maintaining the subgraph hierarchy for DOT output.
//
// Example:
//
//	outer := g.Subgraph("cluster_outer", func(o *Subgraph) {
//		o.SetLabel("Outer")
//		o.Subgraph("cluster_inner", func(i *Subgraph) {
//			i.SetLabel("Inner")
//			i.AddNode(NewNode("A"))
//		})
//	})
func (sg *Subgraph) Subgraph(name string, fn func(*Subgraph)) *Subgraph {
	nested := &Subgraph{
		name:      name,
		nodes:     make(map[string]*Node),
		edges:     make([]*Edge, 0),
		parent:    sg.parent, // Reference root graph for node tracking
		subgraphs: make([]*Subgraph, 0),
	}

	fn(nested)

	sg.subgraphs = append(sg.subgraphs, nested)

	return nested
}

// Subgraphs returns all nested subgraphs within this subgraph.
// The returned slice should not be modified.
func (sg *Subgraph) Subgraphs() []*Subgraph {
	return sg.subgraphs
}

// String returns the DOT representation of the subgraph.
// The output includes the subgraph declaration, attributes, nodes, and edges.
func (sg *Subgraph) String() string {
	builder := strings.Builder{}

	// Start subgraph declaration
	builder.WriteString(fmt.Sprintf("subgraph %s {\n", quoteDOTID(sg.name)))

	// Add subgraph attributes
	if sg.attrs != nil {
		attrs := sg.attrs.List()
		if len(attrs) > 0 {
			for _, attr := range attrs {
				builder.WriteString(fmt.Sprintf("\t\t%s;\n", attr))
			}
		}
	}

	// Add nodes
	for _, node := range sg.nodes {
		builder.WriteString(fmt.Sprintf("\t\t%s;\n", node))
	}

	// Add edges
	for _, edge := range sg.edges {
		builder.WriteString(fmt.Sprintf("\t\t%s;\n", edge.ToString(sg.parent.IsDirected())))
	}

	builder.WriteString("\t}")

	return builder.String()
}
