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

// WithEdgeAttribute sets a custom attribute on an edge.
// This is an escape hatch for Graphviz attributes that don't have typed options.
//
// Example:
//
//	e := g.AddEdge(n1, n2,
//	    WithEdgeLabel("connects"),
//	    WithEdgeAttribute("penwidth", "2.0"),
//	    WithEdgeAttribute("constraint", "false"),
//	)
func WithEdgeAttribute(k, v string) EdgeOption {
	return newEdgeOption(func(a *EdgeAttributes) {
		a.setCustom(k, v)
	})
}
