package goraffe

import (
	"fmt"
	"maps"
)

// Shape represents the visual shape of a node.
// See https://www.graphviz.org/doc/info/shapes.html for all available shapes.
type Shape string

// Predefined node shapes supported by Graphviz.
// This is not an exhaustive list; see Graphviz documentation for all available shapes.
const (
	ShapeBox       Shape = "box"       // Rectangular box shape
	ShapeCircle    Shape = "circle"    // Circular shape
	ShapeEllipse   Shape = "ellipse"   // Elliptical shape (default)
	ShapeDiamond   Shape = "diamond"   // Diamond shape
	ShapeRecord    Shape = "record"    // Record-based shape for structured nodes
	ShapePlaintext Shape = "plaintext" // Text with no surrounding shape
)

// NodeAttributes holds the visual and structural properties of a node.
// All fields use pointer types to distinguish between "not set" and "explicitly set to zero value".
// Use the getter methods (Label(), Shape(), etc.) to access values safely.
type NodeAttributes struct {
	label     *string
	shape     *Shape
	color     *string
	fillColor *string
	fontName  *string
	fontSize  *float64
	custom    map[string]string
}

// Custom returns a copy of all custom attributes set via WithNodeAttribute.
// Returns an empty map if no custom attributes are set.
// The returned map is a copy and can be safely modified without affecting the node.
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

func (a NodeAttributes) List() []string {
	attrs := make([]string, 0)

	if a.label != nil {
		attrs = append(attrs, fmt.Sprintf(`label="%s"`, escapeDOTString(a.Label())))
	}

	if a.shape != nil {
		attrs = append(attrs, fmt.Sprintf(`shape="%s"`, escapeDOTString(string(a.Shape()))))
	}

	if a.color != nil {
		attrs = append(attrs, fmt.Sprintf(`color="%s"`, escapeDOTString(a.Color())))
	}

	if a.fillColor != nil {
		attrs = append(attrs, fmt.Sprintf(`fillcolor="%s"`, escapeDOTString(a.FillColor())))
		// HACK: this is a temporary hack to ensure a set fillcolor appears as expected
		// When we support the `style` attribute for nodes, we'll allow this to be set
		// when the fillcolor is defined, but overridden later. For now, this.
		// -- MRB, 2026-01-03
		attrs = append(attrs, `style="filled"`)
	}

	if a.fontName != nil {
		attrs = append(attrs, fmt.Sprintf(`fontname="%s"`, escapeDOTString(a.FontName())))
	}

	if a.fontSize != nil {
		attrs = append(attrs, fmt.Sprintf(`fontsize="%g"`, a.FontSize()))
	}

	for k, v := range a.custom {
		attrs = append(attrs, fmt.Sprintf(`%s="%s"`, k, escapeDOTString(v)))
	}

	return attrs
}
