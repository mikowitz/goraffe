package goraffe

import "maps"

type RankDir string

const (
	RankDirTB RankDir = "TB"
	RankDirBT RankDir = "BT"
	RankDirLR RankDir = "LR"
	RankDirRL RankDir = "RL"
)

type SplineType string

const (
	SplineTrue     SplineType = "true"
	SplineSpline   SplineType = "spline" // synonym for true
	SplineFalse    SplineType = "false"
	SplineLine     SplineType = "line" // synonym for false
	SplineOrtho    SplineType = "ortho"
	SplinePolyline SplineType = "polyline"
	SplineCurved   SplineType = "curved"
	SplineNone     SplineType = "none"
)

type GraphAttributes struct {
	label    *string
	rankDir  *RankDir
	bgColor  *string
	fontName *string
	fontSize *float64
	splines  *SplineType
	nodeSep  *float64 // default 0.25
	rankSep  *float64 // default 0.5 in dot, 1.0 in twopi
	compound *bool
	custom   map[string]string
}

func (a GraphAttributes) Custom() map[string]string {
	ret := make(map[string]string)
	if a.custom != nil {
		maps.Copy(ret, a.custom)
	}
	return ret
}

func (a *GraphAttributes) setCustom(key, value string) {
	if a.custom == nil {
		a.custom = make(map[string]string)
	}
	a.custom[key] = value
}

// Label returns the graph label. Returns empty string if unset.
// Note: An empty string return value may indicate either an unset label or a label
// explicitly set to empty string.
func (a *GraphAttributes) Label() string {
	if a.label == nil {
		return ""
	}
	return *a.label
}

// RankDir returns the rank direction for the graph. Returns empty string if unset.
// Note: An empty RankDir may indicate either an unset value or an explicitly set empty value.
func (a *GraphAttributes) RankDir() RankDir {
	if a.rankDir == nil {
		return ""
	}
	return *a.rankDir
}

// BgColor returns the background color. Returns empty string if unset.
// Note: An empty string return value may indicate either an unset color or a color
// explicitly set to empty string.
func (a *GraphAttributes) BgColor() string {
	if a.bgColor == nil {
		return ""
	}
	return *a.bgColor
}

// FontName returns the font name. Returns empty string if unset.
// Note: An empty string return value may indicate either an unset font or a font
// explicitly set to empty string.
func (a *GraphAttributes) FontName() string {
	if a.fontName == nil {
		return ""
	}
	return *a.fontName
}

// FontSize returns the font size. Returns 0.0 if unset.
// Note: A zero return value may indicate either an unset font size or a font size
// explicitly set to 0.0.
func (a *GraphAttributes) FontSize() float64 {
	if a.fontSize == nil {
		return 0.0
	}
	return *a.fontSize
}

// Splines returns the spline type for edge routing. Returns empty string if unset.
// Note: An empty SplineType may indicate either an unset value or an explicitly set empty value.
func (a *GraphAttributes) Splines() SplineType {
	if a.splines == nil {
		return ""
	}
	return *a.splines
}

// NodeSep returns the node separation distance. Returns 0.0 if unset.
// Note: A zero return value may indicate either an unset node separation or a value
// explicitly set to 0.0. The Graphviz default is 0.25.
func (a *GraphAttributes) NodeSep() float64 {
	if a.nodeSep == nil {
		return 0.0
	}
	return *a.nodeSep
}

// RankSep returns the rank separation distance. Returns 0.0 if unset.
// Note: A zero return value may indicate either an unset rank separation or a value
// explicitly set to 0.0. The Graphviz default is 0.5 in dot, 1.0 in twopi.
func (a *GraphAttributes) RankSep() float64 {
	if a.rankSep == nil {
		return 0.0
	}
	return *a.rankSep
}

// Compound returns whether compound mode is enabled. Returns false if unset.
// Note: A false return value may indicate either an unset compound flag or a flag
// explicitly set to false.
func (a *GraphAttributes) Compound() bool {
	if a.compound == nil {
		return false
	}
	return *a.compound
}
