// ABOUTME: Defines attributes for subgraphs and clusters in Graphviz DOT graphs.
// ABOUTME: SubgraphAttributes control visual properties like labels, colors, and styles.
package goraffe

import (
	"fmt"
	"maps"
)

// SubgraphAttributes holds the visual and structural properties of a subgraph.
// All fields use pointer types to distinguish between "not set" and "explicitly set to zero value".
// Use the getter methods (Label(), Style(), etc.) to access values safely.
//
// Note: Some attributes like FillColor and Color are typically only rendered for cluster subgraphs
// (those with names starting with "cluster"). Regular subgraphs may not visually display these attributes.
type SubgraphAttributes struct {
	label     *string
	style     *string
	color     *string
	fillColor *string
	fontName  *string
	fontSize  *float64
	custom    map[string]string
}

// Custom returns a copy of all custom attributes set via SetAttribute.
// Returns an empty map if no custom attributes are set.
// The returned map is a copy and can be safely modified without affecting the subgraph.
func (a SubgraphAttributes) Custom() map[string]string {
	ret := make(map[string]string)
	if a.custom != nil {
		maps.Copy(ret, a.custom)
	}

	return ret
}

func (a *SubgraphAttributes) setCustom(key, value string) {
	if a.custom == nil {
		a.custom = make(map[string]string)
	}

	a.custom[key] = value
}

// Label returns the subgraph label. Returns empty string if unset.
// Note: An empty string return value may indicate either an unset label or a label
// explicitly set to empty string.
func (a *SubgraphAttributes) Label() string {
	if a.label == nil {
		return ""
	}

	return *a.label
}

// Style returns the subgraph style. Returns empty string if unset.
// Note: An empty string return value may indicate either an unset style or a style
// explicitly set to empty string.
func (a *SubgraphAttributes) Style() string {
	if a.style == nil {
		return ""
	}

	return *a.style
}

// Color returns the subgraph border color. Returns empty string if unset.
// Note: An empty string return value may indicate either an unset color or a color
// explicitly set to empty string.
// Typically only visible for cluster subgraphs.
func (a *SubgraphAttributes) Color() string {
	if a.color == nil {
		return ""
	}

	return *a.color
}

// FillColor returns the subgraph fill color. Returns empty string if unset.
// Note: An empty string return value may indicate either an unset fill color or a fill color
// explicitly set to empty string.
// Typically only visible for cluster subgraphs.
func (a *SubgraphAttributes) FillColor() string {
	if a.fillColor == nil {
		return ""
	}

	return *a.fillColor
}

// FontName returns the font name. Returns empty string if unset.
// Note: An empty string return value may indicate either an unset font or a font
// explicitly set to empty string.
func (a *SubgraphAttributes) FontName() string {
	if a.fontName == nil {
		return ""
	}

	return *a.fontName
}

// FontSize returns the font size. Returns 0.0 if unset.
// Note: A zero return value may indicate either an unset font size or a font size
// explicitly set to 0.0.
func (a *SubgraphAttributes) FontSize() float64 {
	if a.fontSize == nil {
		return 0.0
	}

	return *a.fontSize
}

// List returns a slice of DOT attribute strings for rendering.
// Only attributes that have been explicitly set are included.
func (a SubgraphAttributes) List() []string {
	attrs := make([]string, 0)

	if a.label != nil {
		attrs = append(attrs, fmt.Sprintf(`label="%s"`, escapeDOTString(a.Label())))
	}

	if a.style != nil {
		attrs = append(attrs, fmt.Sprintf(`style="%s"`, escapeDOTString(a.Style())))
	}

	if a.color != nil {
		attrs = append(attrs, fmt.Sprintf(`color="%s"`, escapeDOTString(a.Color())))
	}

	if a.fillColor != nil {
		attrs = append(attrs, fmt.Sprintf(`fillcolor="%s"`, escapeDOTString(a.FillColor())))
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
