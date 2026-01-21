// ABOUTME: Provides rendering functionality for graphs using Graphviz.
// ABOUTME: Defines Format and Layout enums for controlling output format and layout algorithm.
package goraffe

import (
	"fmt"
	"os/exec"
	"strings"
)

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

// findGraphviz finds the Graphviz binary for the given layout.
// Returns the full path to the binary or ErrGraphvizNotFound.
func findGraphviz(layout Layout) (string, error) {
	binaryName := string(layout)
	path, err := exec.LookPath(binaryName)
	if err != nil {
		return "", ErrGraphvizNotFound
	}
	return path, nil
}

// GraphvizVersion returns the version of Graphviz installed.
// Returns the version string or an error if Graphviz is not found.
func GraphvizVersion() (string, error) {
	// Try to find dot first
	if _, err := findGraphviz(LayoutDot); err != nil {
		return "", err
	}

	// Run "dot -V" to get version
	cmd := exec.Command("dot", "-V")
	output, err := cmd.CombinedOutput()
	if err != nil {
		return "", fmt.Errorf("failed to get Graphviz version: %w", err)
	}

	// Parse version from output (format: "dot - graphviz version X.Y.Z ...")
	version := strings.TrimSpace(string(output))
	return version, nil
}

// checkGraphvizInstalled checks if Graphviz is installed and available.
// Returns nil if available, or ErrGraphvizNotFound if not.
func checkGraphvizInstalled() error {
	_, err := findGraphviz(LayoutDot)
	return err
}
