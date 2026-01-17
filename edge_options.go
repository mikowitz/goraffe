package goraffe

// EdgeOption is a functional option for configuring edge attributes.
// Options can be passed to Graph.AddEdge or used to build reusable attribute templates.
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

// WithEdgeLabel sets the text label displayed on or near the edge.
//
// Example:
//
//	e := g.AddEdge(n1, n2, goraffe.WithEdgeLabel("connects"))
func WithEdgeLabel(l string) EdgeOption {
	return newEdgeOption(func(a *EdgeAttributes) {
		a.label = &l
	})
}

// WithEdgeColor sets the color of the edge line and arrowhead.
// Accepts color names (e.g., "red", "blue") or hex values (e.g., "#FF0000").
//
// Example:
//
//	e := g.AddEdge(n1, n2, goraffe.WithEdgeColor("blue"))
func WithEdgeColor(c string) EdgeOption {
	return newEdgeOption(func(a *EdgeAttributes) {
		a.color = &c
	})
}

// WithEdgeStyle sets the visual style of the edge line (solid, dashed, dotted, etc.).
//
// Example:
//
//	e := g.AddEdge(n1, n2, goraffe.WithEdgeStyle(goraffe.EdgeStyleDashed))
func WithEdgeStyle(s EdgeStyle) EdgeOption {
	return newEdgeOption(func(a *EdgeAttributes) {
		a.style = &s
	})
}

// WithArrowHead sets the style of arrowhead at the edge destination.
// Only applies to directed graphs.
//
// Example:
//
//	e := g.AddEdge(n1, n2, goraffe.WithArrowHead(goraffe.ArrowDot))
func WithArrowHead(t ArrowType) EdgeOption {
	return newEdgeOption(func(a *EdgeAttributes) {
		a.arrowHead = &t
	})
}

// WithArrowTail sets the style of arrowhead at the edge source.
// Only applies to directed graphs.
//
// Example:
//
//	e := g.AddEdge(n1, n2, goraffe.WithArrowTail(goraffe.ArrowVee))
func WithArrowTail(t ArrowType) EdgeOption {
	return newEdgeOption(func(a *EdgeAttributes) {
		a.arrowTail = &t
	})
}

// WithWeight sets the edge weight, affecting edge length and crossing minimization.
// Higher weights make edges shorter and more important in layout optimization.
//
// Example:
//
//	e := g.AddEdge(n1, n2, goraffe.WithWeight(2.0))
func WithWeight(w float64) EdgeOption {
	return newEdgeOption(func(a *EdgeAttributes) {
		a.weight = &w
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

// FromPort specifies which port on the source node this edge connects from.
// The port must be defined in the source node's HTML label.
//
// Example:
//
//	outPort := Cell(Text("output")).Port("out").GetPort()
//	e := g.AddEdge(n1, n2, FromPort(outPort))
func FromPort(p *Port) EdgeOption {
	return newEdgeOption(func(a *EdgeAttributes) {
		a.fromPort = p
	})
}

// ToPort specifies which port on the destination node this edge connects to.
// The port must be defined in the destination node's HTML label.
//
// Example:
//
//	inPort := Cell(Text("input")).Port("in").GetPort()
//	e := g.AddEdge(n1, n2, ToPort(inPort))
func ToPort(p *Port) EdgeOption {
	return newEdgeOption(func(a *EdgeAttributes) {
		a.toPort = p
	})
}
