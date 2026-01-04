package goraffe

import (
	"fmt"
	"sort"
	"strings"
)

type Node struct {
	id    string
	attrs *NodeAttributes
}

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

func (n *Node) ID() string {
	return n.id
}

func (n *Node) Attrs() *NodeAttributes {
	return n.attrs
}

func (n *Node) String() string {
	builder := strings.Builder{}
	builder.WriteString(fmt.Sprintf(`"%s"`, n.ID()))
	attrs := make([]string, 0)

	if n.attrs.label != nil {
		attrs = append(attrs, fmt.Sprintf(`label="%s"`, n.attrs.Label()))
	}
	if n.attrs.shape != nil {
		attrs = append(attrs, fmt.Sprintf(`shape="%s"`, n.attrs.Shape()))
	}
	if n.attrs.color != nil {
		attrs = append(attrs, fmt.Sprintf(`color="%s"`, n.attrs.Color()))
	}
	if n.attrs.fillColor != nil {
		attrs = append(attrs, fmt.Sprintf(`fillcolor="%s"`, n.attrs.FillColor()))
		// HACK: this is a temporary hack to ensure a set fillcolor appears as expected
		// When we support the `style` attribute for nodes, we'll allow this to be set
		// when the fillcolor is defined, but overridden later. For now, this.
		// -- MRB, 2026-01-03
		attrs = append(attrs, `style="filled"`)
	}
	if n.attrs.fontName != nil {
		attrs = append(attrs, fmt.Sprintf(`fontname="%s"`, n.attrs.FontName()))
	}
	if n.attrs.fontSize != nil {
		attrs = append(attrs, fmt.Sprintf(`fontsize="%0.2f"`, n.attrs.FontSize()))
	}

	for k, v := range n.attrs.custom {
		attrs = append(attrs, fmt.Sprintf(`%s="%s"`, k, v))
	}

	var attrsStr string

	if len(attrs) > 0 {
		sort.Strings(attrs)
		attrsStr = "[" + strings.Join(attrs, ", ") + "]"
		builder.WriteString(" ")
		builder.WriteString(attrsStr)
	}
	return builder.String()
}
