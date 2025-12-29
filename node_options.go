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

func WithShape(s Shape) NodeOption {
	return newNodeOption(func(a *NodeAttributes) {
		a.Shape = s
	})
}

func WithLabel(l string) NodeOption {
	return newNodeOption(func(a *NodeAttributes) {
		a.Label = l
	})
}

func WithColor(c string) NodeOption {
	return newNodeOption(func(a *NodeAttributes) {
		a.Color = c
	})
}

func WithFillColor(c string) NodeOption {
	return newNodeOption(func(a *NodeAttributes) {
		a.FillColor = c
	})
}

func WithFontName(n string) NodeOption {
	return newNodeOption(func(a *NodeAttributes) {
		a.FontName = n
	})
}

func WithFontSize(s float64) NodeOption {
	return newNodeOption(func(a *NodeAttributes) {
		a.FontSize = s
	})
}
