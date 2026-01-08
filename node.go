package goraffe

import (
	"fmt"
	"sort"
	"strings"
)

// Node represents a node (vertex) in a graph.
// Each node has a unique ID and optional visual attributes like shape, color, and label.
type Node struct {
	id    string
	attrs *NodeAttributes
}

// NewNode creates a new node with the given ID and optional attributes.
// The ID must be unique within a graph and will be used to identify the node in DOT output.
//
// Example:
//
//	n := goraffe.NewNode("A", goraffe.WithLabel("Start Node"), goraffe.WithCircleShape())
func NewNode(id string, options ...NodeOption) *Node {
	attrs := &NodeAttributes{}

	for _, option := range options {
		option.applyNode(attrs)
	}

	return &Node{
		id:    id,
		attrs: attrs,
	}
}

// ID returns the unique identifier for this node.
func (n *Node) ID() string {
	return n.id
}

// Attrs returns the node's visual attributes (label, shape, color, etc.).
// The returned attributes can be modified to change the node's appearance.
func (n *Node) Attrs() *NodeAttributes {
	return n.attrs
}

// String returns the DOT representation of the node with its attributes.
// The output includes the node ID and any set attributes in DOT format.
func (n *Node) String() string {
	builder := strings.Builder{}
	builder.WriteString(fmt.Sprintf(`"%s"`, n.ID()))

	attrs := n.attrs.List()
	var attrsStr string

	if len(attrs) > 0 {
		sort.Strings(attrs)
		attrsStr = "[" + strings.Join(attrs, ", ") + "]"

		builder.WriteString(" ")
		builder.WriteString(attrsStr)
	}

	return builder.String()
}
