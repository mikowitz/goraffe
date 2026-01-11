package goraffe

import (
	"fmt"
	"maps"
)

// EdgeStyle represents the visual style of an edge line.
type EdgeStyle string

// ArrowType represents the type of arrowhead or arrowtail on an edge.
// See https://www.graphviz.org/docs/attr-types/arrowType/ for all available types.
type ArrowType string

// Predefined edge styles supported by Graphviz.
const (
	EdgeStyleSolid     EdgeStyle = "solid"  // Solid line (default)
	EdgeStyleDashed    EdgeStyle = "dashed" // Dashed line
	EdgeStyleDotted    EdgeStyle = "dotted" // Dotted line
	EdgeStyleBold      EdgeStyle = "bold"   // Bold solid line
	EdgeStyleInvisible EdgeStyle = "invis"  // Invisible edge (affects layout but not visible)
)

// Predefined arrow types supported by Graphviz.
// This is not an exhaustive list; see Graphviz documentation for all available arrow types.
const (
	ArrowNormal ArrowType = "normal" // Standard arrow (default)
	ArrowDot    ArrowType = "dot"    // Circular dot
	ArrowNone   ArrowType = "none"   // No arrow
	ArrowVee    ArrowType = "vee"    // V-shaped arrow
)

// EdgeAttributes holds the visual and structural properties of an edge.
// All fields use pointer types to distinguish between "not set" and "explicitly set to zero value".
// Use the getter methods (Label(), Color(), etc.) to access values safely.
type EdgeAttributes struct {
	label     *string
	color     *string
	style     *EdgeStyle
	arrowHead *ArrowType
	arrowTail *ArrowType
	weight    *float64
	custom    map[string]string
}

// Custom returns a copy of all custom attributes set via WithEdgeAttribute.
// Returns an empty map if no custom attributes are set.
// The returned map is a copy and can be safely modified without affecting the edge.
func (a EdgeAttributes) Custom() map[string]string {
	ret := make(map[string]string)
	if a.custom != nil {
		maps.Copy(ret, a.custom)
	}

	return ret
}

func (a *EdgeAttributes) setCustom(key, value string) {
	if a.custom == nil {
		a.custom = make(map[string]string)
	}

	a.custom[key] = value
}

// Label returns the edge label. Returns empty string if unset.
// Note: An empty string return value may indicate either an unset label or a label
// explicitly set to empty string.
func (a *EdgeAttributes) Label() string {
	if a.label == nil {
		return ""
	}

	return *a.label
}

// Color returns the edge color. Returns empty string if unset.
// Note: An empty string return value may indicate either an unset color or a color
// explicitly set to empty string.
func (a *EdgeAttributes) Color() string {
	if a.color == nil {
		return ""
	}

	return *a.color
}

// Style returns the edge style. Returns empty string if unset.
// Note: An empty EdgeStyle may indicate either an unset value or an explicitly set empty value.
func (a *EdgeAttributes) Style() EdgeStyle {
	if a.style == nil {
		return ""
	}

	return *a.style
}

// ArrowHead returns the arrowhead type. Returns empty string if unset.
// Note: An empty ArrowType may indicate either an unset value or an explicitly set empty value.
func (a *EdgeAttributes) ArrowHead() ArrowType {
	if a.arrowHead == nil {
		return ""
	}

	return *a.arrowHead
}

// ArrowTail returns the arrowtail type. Returns empty string if unset.
// Note: An empty ArrowType may indicate either an unset value or an explicitly set empty value.
func (a *EdgeAttributes) ArrowTail() ArrowType {
	if a.arrowTail == nil {
		return ""
	}

	return *a.arrowTail
}

// Weight returns the edge weight. Returns 0.0 if unset.
// Note: A zero return value may indicate either an unset weight or a weight
// explicitly set to 0.0.
func (a *EdgeAttributes) Weight() float64 {
	if a.weight == nil {
		return 0.0
	}

	return *a.weight
}

// applyEdge implements the EdgeOption interface, allowing EdgeAttributes
// to be used as a reusable template. Only non-nil pointer fields are copied.
//
// NOTE: The unexported custom field is intentionally NOT copied. Custom fields
// are treated as per-instance customizations, not template values to be shared.
// Use Custom() to read custom fields and WithEdgeAttribute() to set them.
func (a EdgeAttributes) applyEdge(dst *EdgeAttributes) {
	if a.label != nil {
		dst.label = a.label
	}

	if a.color != nil {
		dst.color = a.color
	}

	if a.style != nil {
		dst.style = a.style
	}

	if a.arrowHead != nil {
		dst.arrowHead = a.arrowHead
	}

	if a.arrowTail != nil {
		dst.arrowTail = a.arrowTail
	}

	if a.weight != nil {
		dst.weight = a.weight
	}
}

func (a EdgeAttributes) List() []string {
	attrs := make([]string, 0)

	if a.label != nil {
		attrs = append(attrs, fmt.Sprintf(`label="%s"`, escapeDOTString(a.Label())))
	}
	if a.color != nil {
		attrs = append(attrs, fmt.Sprintf(`color="%s"`, escapeDOTString(a.Color())))
	}
	if a.style != nil {
		attrs = append(attrs, fmt.Sprintf(`style="%s"`, escapeDOTString(string(a.Style()))))
	}
	if a.arrowHead != nil {
		attrs = append(attrs, fmt.Sprintf(`arrowhead="%s"`, escapeDOTString(string(a.ArrowHead()))))
	}
	if a.arrowTail != nil {
		attrs = append(attrs, fmt.Sprintf(`arrowtail="%s"`, escapeDOTString(string(a.ArrowTail()))))
	}
	if a.weight != nil {
		attrs = append(attrs, fmt.Sprintf(`weight="%g"`, a.Weight()))
	}

	for k, v := range a.custom {
		attrs = append(attrs, fmt.Sprintf(`%s="%s"`, k, escapeDOTString(v)))
	}

	return attrs
}
