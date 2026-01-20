// ABOUTME: Provides rendering functionality for graphs using Graphviz.
// ABOUTME: Defines Format and Layout enums for controlling output format and layout algorithm.
package goraffe

// Format represents the output format for rendered graphs.
type Format string

const (
	// PNG produces PNG (Portable Network Graphics) raster images.
	PNG Format = "png"
	// SVG produces SVG (Scalable Vector Graphics) vector images.
	SVG Format = "svg"
	// PDF produces PDF (Portable Document Format) documents.
	PDF Format = "pdf"
	// DOT produces DOT language source code.
	DOT Format = "dot"
)

// Layout represents the graph layout algorithm to use.
type Layout string

const (
	// LayoutDot uses the hierarchical "dot" layout for directed graphs.
	LayoutDot Layout = "dot"
	// LayoutNeato uses the "neato" spring model layout.
	LayoutNeato Layout = "neato"
	// LayoutFdp uses the "fdp" force-directed layout.
	LayoutFdp Layout = "fdp"
	// LayoutSfdp uses the "sfdp" scalable force-directed layout for large graphs.
	LayoutSfdp Layout = "sfdp"
	// LayoutTwopi uses the "twopi" radial layout.
	LayoutTwopi Layout = "twopi"
	// LayoutCirco uses the "circo" circular layout.
	LayoutCirco Layout = "circo"
	// LayoutOsage uses the "osage" clustered layout.
	LayoutOsage Layout = "osage"
	// LayoutPatchwork uses the "patchwork" treemap layout.
	LayoutPatchwork Layout = "patchwork"
)
