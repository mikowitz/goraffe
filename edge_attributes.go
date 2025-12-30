package goraffe

import "maps"

type (
	EdgeStyle string
	ArrowType string
)

const (
	EdgeStyleSolid     EdgeStyle = "solid"
	EdgeStyleDashed    EdgeStyle = "dashed"
	EdgeStyleDotted    EdgeStyle = "dotted"
	EdgeStyleBold      EdgeStyle = "bold"
	EdgeStyleInvisible EdgeStyle = "invis"

	// NOTE: https://www.graphviz.org/docs/attr-types/arrowType/
	ArrowNormal ArrowType = "normal"
	ArrowDot    ArrowType = "dot"
	ArrowNone   ArrowType = "none"
	ArrowVee    ArrowType = "vee"
)

type EdgeAttributes struct {
	Label     string
	Color     string
	Style     EdgeStyle
	ArrowHead ArrowType
	ArrowTail ArrowType
	Weight    float64
	custom    map[string]string
}

func (a EdgeAttributes) Custom() map[string]string {
	ret := make(map[string]string)
	if a.custom != nil {
		maps.Copy(ret, a.custom)
	}
	return ret
}
