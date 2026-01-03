package goraffe

import (
	"maps"
)

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
	label     *string
	shape     *Shape
	color     *string
	fillColor *string
	fontName  *string
	fontSize  *float64
	custom    map[string]string
}

func (a NodeAttributes) Custom() map[string]string {
	ret := make(map[string]string)
	if a.custom != nil {
		maps.Copy(ret, a.custom)
	}
	return ret
}

func (a *NodeAttributes) setCustom(key, value string) {
	if a.custom == nil {
		a.custom = make(map[string]string)
	}
	a.custom[key] = value
}

// Label returns the node label. Returns empty string if unset.
// Note: An empty string return value may indicate either an unset label or a label
// explicitly set to empty string.
func (a *NodeAttributes) Label() string {
	if a.label == nil {
		return ""
	}
	return *a.label
}

// Shape returns the node shape. Returns empty string if unset.
// Note: An empty Shape may indicate either an unset value or an explicitly set empty value.
func (a *NodeAttributes) Shape() Shape {
	if a.shape == nil {
		return ""
	}
	return *a.shape
}

// Color returns the node color. Returns empty string if unset.
// Note: An empty string return value may indicate either an unset color or a color
// explicitly set to empty string.
func (a *NodeAttributes) Color() string {
	if a.color == nil {
		return ""
	}
	return *a.color
}

// FillColor returns the node fill color. Returns empty string if unset.
// Note: An empty string return value may indicate either an unset fill color or a fill color
// explicitly set to empty string.
func (a *NodeAttributes) FillColor() string {
	if a.fillColor == nil {
		return ""
	}
	return *a.fillColor
}

// FontName returns the font name. Returns empty string if unset.
// Note: An empty string return value may indicate either an unset font or a font
// explicitly set to empty string.
func (a *NodeAttributes) FontName() string {
	if a.fontName == nil {
		return ""
	}
	return *a.fontName
}

// FontSize returns the font size. Returns 0.0 if unset.
// Note: A zero return value may indicate either an unset font size or a font size
// explicitly set to 0.0.
func (a *NodeAttributes) FontSize() float64 {
	if a.fontSize == nil {
		return 0.0
	}
	return *a.fontSize
}

// applyNode implements the NodeOption interface, allowing NodeAttributes
// to be used as a reusable template. Only non-nil pointer fields are copied.
//
// NOTE: The unexported custom field is intentionally NOT copied. Custom fields
// are treated as per-instance customizations, not template values to be shared.
// Use Custom() to read custom fields and WithNodeAttribute() to set them.
func (a NodeAttributes) applyNode(dst *NodeAttributes) {
	if a.label != nil {
		dst.label = a.label
	}
	if a.shape != nil {
		dst.shape = a.shape
	}
	if a.color != nil {
		dst.color = a.color
	}
	if a.fillColor != nil {
		dst.fillColor = a.fillColor
	}
	if a.fontName != nil {
		dst.fontName = a.fontName
	}
	if a.fontSize != nil {
		dst.fontSize = a.fontSize
	}
}
