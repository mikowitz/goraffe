package goraffe

import (
	"errors"
	"fmt"
	"io"
	"sort"
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
	subgraphs        []*Subgraph
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
		subgraphs:        make([]*Subgraph, 0),
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
// Output order follows DOT conventions:
// 1. Graph declaration
// 2. Graph attributes
// 3. Default node/edge attributes
// 4. Subgraphs (each contains their nodes)
// 5. Nodes not in any subgraph
// 6. All edges
// 7. Closing brace
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
		builder.WriteString(fmt.Sprintf(" %s", quoteDOTID(g.name)))
	}

	builder.WriteString(" {\n")

	g.addGraphAttributes(&builder)
	g.addDefaultNodeAttributes(&builder)
	g.addDefaultEdgeAttrbiutes(&builder)

	// Output subgraphs with their nodes
	for _, subgraph := range g.subgraphs {
		g.renderSubgraph(&builder, subgraph, 1)
	}

	// Output nodes not in any subgraph
	nodesInSubgraphs := g.collectNodesInSubgraphs()
	for _, node := range g.nodeOrder {
		if !nodesInSubgraphs[node.ID()] {
			builder.WriteString(fmt.Sprintf("\t%s;\n", node))
		}
	}

	// Output all edges
	for _, edge := range g.edges {
		builder.WriteString(fmt.Sprintf("\t%s;\n", edge.ToString(g.directed)))
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

func (g *Graph) addGraphAttributes(builder *strings.Builder) {
	attrs := g.attrs.List()
	var attrsStr string

	if len(attrs) > 0 {
		sort.Strings(attrs)
		attrsStr = strings.Join(attrs, "\n")

		fmt.Fprintf(builder, "%s\n", attrsStr)
	}
}

func (g *Graph) addDefaultNodeAttributes(builder *strings.Builder) {
	attrs := g.defaultNodeAttrs.List()
	var attrsStr string

	if len(attrs) > 0 {
		sort.Strings(attrs)
		attrsStr = "[" + strings.Join(attrs, ", ") + "]"

		fmt.Fprintf(builder, "\tnode %s;\n", attrsStr)
	}
}

func (g *Graph) addDefaultEdgeAttrbiutes(builder *strings.Builder) {
	attrs := g.defaultEdgeAttrs.List()
	var attrsStr string

	if len(attrs) > 0 {
		sort.Strings(attrs)
		attrsStr = "[" + strings.Join(attrs, ", ") + "]"

		fmt.Fprintf(builder, "\tedge %s;\n", attrsStr)
	}
}

// collectNodesInSubgraphs returns a set of all node IDs that are in any subgraph.
// This is used to determine which nodes should be output as "loose" nodes.
func (g *Graph) collectNodesInSubgraphs() map[string]bool {
	nodesInSubgraphs := make(map[string]bool)

	var collectFromSubgraph func(*Subgraph)
	collectFromSubgraph = func(sg *Subgraph) {
		for id := range sg.nodes {
			nodesInSubgraphs[id] = true
		}
		// Recursively collect from nested subgraphs
		for _, nested := range sg.subgraphs {
			collectFromSubgraph(nested)
		}
	}

	for _, sg := range g.subgraphs {
		collectFromSubgraph(sg)
	}

	return nodesInSubgraphs
}

// renderSubgraph renders a subgraph and its nested subgraphs to DOT format.
// The depth parameter controls indentation level.
func (g *Graph) renderSubgraph(builder *strings.Builder, sg *Subgraph, depth int) {
	indent := strings.Repeat("\t", depth)

	// Start subgraph declaration
	if sg.name == "" {
		fmt.Fprintf(builder, "%ssubgraph {\n", indent)
	} else {
		fmt.Fprintf(builder, "%ssubgraph %s {\n", indent, quoteDOTID(sg.name))
	}

	// Add subgraph attributes
	if sg.attrs != nil {
		attrs := sg.attrs.List()
		if len(attrs) > 0 {
			for _, attr := range attrs {
				fmt.Fprintf(builder, "%s\t%s;\n", indent, attr)
			}
		}
	}

	// Recursively render nested subgraphs first
	for _, nested := range sg.subgraphs {
		g.renderSubgraph(builder, nested, depth+1)
	}

	// Add nodes that belong directly to this subgraph (not in nested subgraphs)
	nodesInNested := make(map[string]bool)
	for _, nested := range sg.subgraphs {
		for id := range nested.nodes {
			nodesInNested[id] = true
		}
	}

	for _, node := range sg.nodes {
		if !nodesInNested[node.ID()] {
			fmt.Fprintf(builder, "%s\t%s;\n", indent, node)
		}
	}

	fmt.Fprintf(builder, "%s}\n", indent)
}

// Subgraph creates a new subgraph with the given name and executes the provided function.
// The function receives the created subgraph as a parameter, allowing for subgraph configuration.
// Returns the created subgraph for further use.
//
// Example:
//
//	sg := g.Subgraph("cluster_0", func(s *Subgraph) {
//		s.AddNode(NewNode("A"))
//		s.AddNode(NewNode("B"))
//	})
func (g *Graph) Subgraph(name string, fn func(*Subgraph)) *Subgraph {
	sg := &Subgraph{
		name:      name,
		nodes:     make(map[string]*Node),
		edges:     make([]*Edge, 0),
		parent:    g,
		subgraphs: make([]*Subgraph, 0),
	}

	fn(sg)

	g.subgraphs = append(g.subgraphs, sg)

	return sg
}

// Subgraphs returns all subgraphs in the graph.
// The returned slice should not be modified.
func (g *Graph) Subgraphs() []*Subgraph {
	return g.subgraphs
}

// SameRank creates an anonymous subgraph with rank=same for the given nodes.
// This is a convenience method equivalent to creating a subgraph with SetRank(RankSame).
// All nodes will be placed at the same horizontal level in the graph layout.
// Returns an error if any node is nil.
//
// Example:
//
//	g.SameRank(n1, n2, n3)
func (g *Graph) SameRank(nodes ...*Node) (*Subgraph, error) {
	return g.createRankSubgraph(RankSame, nodes...)
}

// MinRank creates an anonymous subgraph with rank=min for the given nodes.
// This is a convenience method equivalent to creating a subgraph with SetRank(RankMin).
// All nodes will be placed at the minimum rank.
// Returns an error if any node is nil.
func (g *Graph) MinRank(nodes ...*Node) (*Subgraph, error) {
	return g.createRankSubgraph(RankMin, nodes...)
}

// MaxRank creates an anonymous subgraph with rank=max for the given nodes.
// This is a convenience method equivalent to creating a subgraph with SetRank(RankMax).
// All nodes will be placed at the maximum rank.
// Returns an error if any node is nil.
func (g *Graph) MaxRank(nodes ...*Node) (*Subgraph, error) {
	return g.createRankSubgraph(RankMax, nodes...)
}

// SourceRank creates an anonymous subgraph with rank=source for the given nodes.
// This is a convenience method equivalent to creating a subgraph with SetRank(RankSource).
// All nodes will be placed at the source rank (top of graph).
// Returns an error if any node is nil.
func (g *Graph) SourceRank(nodes ...*Node) (*Subgraph, error) {
	return g.createRankSubgraph(RankSource, nodes...)
}

// SinkRank creates an anonymous subgraph with rank=sink for the given nodes.
// This is a convenience method equivalent to creating a subgraph with SetRank(RankSink).
// All nodes will be placed at the sink rank (bottom of graph).
// Returns an error if any node is nil.
func (g *Graph) SinkRank(nodes ...*Node) (*Subgraph, error) {
	return g.createRankSubgraph(RankSink, nodes...)
}

// createRankSubgraph is an internal helper that creates an anonymous subgraph with a rank constraint.
func (g *Graph) createRankSubgraph(rank Rank, nodes ...*Node) (*Subgraph, error) {
	// Validate all nodes upfront
	for _, n := range nodes {
		if n == nil {
			return nil, fmt.Errorf("rank subgraph requires valid nodes: %w", ErrNilNode)
		}
	}

	sg := g.Subgraph("", func(s *Subgraph) {
		s.SetRank(rank)
		for _, n := range nodes {
			_ = s.AddNode(n) // Safe to ignore - already validated
		}
	})

	return sg, nil
}
