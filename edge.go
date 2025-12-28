package goraffe

type Edge struct {
	from, to *Node
}

func (e *Edge) From() *Node {
	return e.from
}

func (e *Edge) To() *Node {
	return e.to
}
