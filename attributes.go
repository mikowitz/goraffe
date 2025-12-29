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
