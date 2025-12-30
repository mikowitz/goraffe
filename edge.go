package goraffe

type Edge struct {
	from, to *Node
	attrs    *EdgeAttributes
}

func (e *Edge) From() *Node {
	return e.from
}

func (e *Edge) To() *Node {
	return e.to
}

func (e *Edge) Attrs() *EdgeAttributes {
	return e.attrs
}
