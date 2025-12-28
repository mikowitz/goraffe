package goraffe

type Node struct {
	id string
}

func NewNode(id string) *Node {
	return &Node{
		id: id,
	}
}

func (n *Node) ID() string {
	return n.id
}
