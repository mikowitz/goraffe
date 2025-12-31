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

// applyEdge implements the EdgeOption interface, allowing EdgeAttributes
// to be used as a reusable template. Only non-zero exported fields are copied.
//
// NOTE: The unexported custom field is intentionally NOT copied. Custom fields
// are treated as per-instance customizations, not template values to be shared.
// Use Custom() to read and SetCustom() to write custom fields on individual edges.
func (a EdgeAttributes) applyEdge(dst *EdgeAttributes) {
	if a.Label != "" {
		dst.Label = a.Label
	}
	if a.Color != "" {
		dst.Color = a.Color
	}
	dst.Style = a.Style
	dst.ArrowHead = a.ArrowHead
	dst.ArrowTail = a.ArrowTail
	if a.Weight != 0.0 {
		dst.Weight = a.Weight
	}
}
