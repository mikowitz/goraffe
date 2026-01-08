package goraffe

import (
	"fmt"
	"sort"
	"strings"
)

// Edge represents a connection between two nodes in a graph.
// Edges can be directed (arrows) or undirected (lines) depending on the graph type.
// Use Graph.AddEdge to create edges.
type Edge struct {
	from, to *Node
	attrs    *EdgeAttributes
}

// From returns the source node of the edge.
// In directed graphs, this is the tail of the arrow.
func (e *Edge) From() *Node {
	return e.from
}

// To returns the destination node of the edge.
// In directed graphs, this is the head of the arrow.
func (e *Edge) To() *Node {
	return e.to
}

// Attrs returns the edge's visual attributes (label, color, style, arrows, etc.).
// The returned attributes can be modified to change the edge's appearance.
func (e *Edge) Attrs() *EdgeAttributes {
	return e.attrs
}

func (e *Edge) ToString(directed bool) string {
	builder := strings.Builder{}
	builder.WriteString(fmt.Sprintf(`"%s"`, e.from.ID()))

	if directed {
		builder.WriteString(" -> ")
	} else {
		builder.WriteString(" -- ")
	}
	builder.WriteString(fmt.Sprintf(`"%s"`, e.to.ID()))

	attrs := e.attrs.List()
	var attrsStr string

	if len(attrs) > 0 {
		sort.Strings(attrs)
		attrsStr = "[" + strings.Join(attrs, ", ") + "]"

		builder.WriteString(" ")
		builder.WriteString(attrsStr)
	}

	return builder.String()
}
