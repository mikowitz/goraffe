package goraffe

type EdgeOption interface {
	applyEdge(*EdgeAttributes)
}

type edgeOptionFunc func(*EdgeAttributes)

func (f edgeOptionFunc) applyEdge(a *EdgeAttributes) {
	f(a)
}

func newEdgeOption(fn func(*EdgeAttributes)) EdgeOption {
	return edgeOptionFunc(fn)
}

func WithEdgeLabel(l string) EdgeOption {
	return newEdgeOption(func(a *EdgeAttributes) {
		a.Label = l
	})
}

func WithEdgeColor(c string) EdgeOption {
	return newEdgeOption(func(a *EdgeAttributes) {
		a.Color = c
	})
}

func WithEdgeStyle(s EdgeStyle) EdgeOption {
	return newEdgeOption(func(a *EdgeAttributes) {
		a.Style = s
	})
}

func WithArrowHead(t ArrowType) EdgeOption {
	return newEdgeOption(func(a *EdgeAttributes) {
		a.ArrowHead = t
	})
}

func WithArrowTail(t ArrowType) EdgeOption {
	return newEdgeOption(func(a *EdgeAttributes) {
		a.ArrowTail = t
	})
}

func WithWeight(w float64) EdgeOption {
	return newEdgeOption(func(a *EdgeAttributes) {
		a.Weight = w
	})
}
