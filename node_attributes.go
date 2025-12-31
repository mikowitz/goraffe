package goraffe

import "maps"

type Shape string

// NOTE: incomplete list of possible node shapes
const (
	ShapeBox       Shape = "box"
	ShapeCircle    Shape = "circle"
	ShapeEllipse   Shape = "ellipse"
	ShapeDiamond   Shape = "diamond"
	ShapeRecord    Shape = "record"
	ShapePlaintext Shape = "plaintext"
)

type NodeAttributes struct {
	Label     string
	Shape     Shape
	Color     string
	FillColor string
	FontName  string
	FontSize  float64
	custom    map[string]string
}

func (a NodeAttributes) Custom() map[string]string {
	ret := make(map[string]string)
	if a.custom != nil {
		maps.Copy(ret, a.custom)
	}
	return ret
}

// applyNode implements the NodeOption interface, allowing NodeAttributes
// to be used as a reusable template. Only non-zero exported fields are copied.
//
// NOTE: The unexported custom field is intentionally NOT copied. Custom fields
// are treated as per-instance customizations, not template values to be shared.
// Use Custom() to read and SetCustom() to write custom fields on individual nodes.
func (a NodeAttributes) applyNode(dst *NodeAttributes) {
	if a.Label != "" {
		dst.Label = a.Label
	}
	dst.Shape = a.Shape
	if a.Color != "" {
		dst.Color = a.Color
	}
	if a.FillColor != "" {
		dst.FillColor = a.FillColor
	}
	if a.FontName != "" {
		dst.FontName = a.FontName
	}
	if a.FontSize != 0.0 {
		dst.FontSize = a.FontSize
	}
}
