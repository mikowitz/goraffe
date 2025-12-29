package goraffe

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
