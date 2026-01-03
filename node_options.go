package goraffe

type NodeOption interface {
	applyNode(*NodeAttributes)
}

type nodeOptionFunc func(*NodeAttributes)

func (f nodeOptionFunc) applyNode(a *NodeAttributes) {
	f(a)
}

func newNodeOption(fn func(*NodeAttributes)) NodeOption {
	return nodeOptionFunc(fn)
}

func withShape(s Shape) NodeOption {
	return newNodeOption(func(a *NodeAttributes) {
		a.shape = &s
	})
}

func WithBoxShape() NodeOption {
	return withShape(ShapeBox)
}

func WithCircleShape() NodeOption {
	return withShape(ShapeCircle)
}

func WithEllipseShape() NodeOption {
	return withShape(ShapeEllipse)
}

func WithDiamondShape() NodeOption {
	return withShape(ShapeDiamond)
}

func WithRecordShape() NodeOption {
	return withShape(ShapeRecord)
}

func WithPlaintextShape() NodeOption {
	return withShape(ShapePlaintext)
}

func WithLabel(l string) NodeOption {
	return newNodeOption(func(a *NodeAttributes) {
		a.label = &l
	})
}

func WithColor(c string) NodeOption {
	return newNodeOption(func(a *NodeAttributes) {
		a.color = &c
	})
}

func WithFillColor(c string) NodeOption {
	return newNodeOption(func(a *NodeAttributes) {
		a.fillColor = &c
	})
}

func WithFontName(n string) NodeOption {
	return newNodeOption(func(a *NodeAttributes) {
		a.fontName = &n
	})
}

func WithFontSize(s float64) NodeOption {
	return newNodeOption(func(a *NodeAttributes) {
		a.fontSize = &s
	})
}

// WithNodeAttribute sets a custom attribute on a node.
// This is an escape hatch for Graphviz attributes that don't have typed options.
//
// Example:
//
//	n := NewNode("A",
//	    WithShape(ShapeBox),
//	    WithNodeAttribute("peripheries", "2"),
//	    WithNodeAttribute("tooltip", "Hover text"),
//	)
func WithNodeAttribute(k, v string) NodeOption {
	return newNodeOption(func(a *NodeAttributes) {
		a.setCustom(k, v)
	})
}
