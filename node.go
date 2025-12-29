package goraffe

type Node struct {
	id    string
	attrs *NodeAttributes
}

func NewNode(id string) *Node {
	return &Node{
		id:    id,
		attrs: &NodeAttributes{},
	}
}

func (n *Node) ID() string {
	return n.id
}

func (n *Node) Attrs() *NodeAttributes {
	return n.attrs
}
