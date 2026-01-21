// ABOUTME: Provides rendering functionality for graphs using Graphviz.
// ABOUTME: Defines Format and Layout enums for controlling output format and layout algorithm.
package goraffe

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"os"
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
	cmd := exec.CommandContext(context.TODO(), "dot", "-V")
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

// RenderOption configures rendering behavior.
type RenderOption interface {
	applyRender(*renderConfig)
}

// renderConfig holds rendering configuration.
type renderConfig struct {
	layout Layout
}

// layoutOption implements RenderOption to set the layout engine.
type layoutOption struct {
	layout Layout
}

func (o layoutOption) applyRender(cfg *renderConfig) {
	cfg.layout = o.layout
}

// WithLayout sets the Graphviz layout engine to use for rendering.
// Default is LayoutDot if not specified.
func WithLayout(l Layout) RenderOption {
	return layoutOption{layout: l}
}

// Render renders the graph to the given writer in the specified format.
// Uses the Graphviz layout engine specified by options (default: dot).
func (g *Graph) Render(format Format, w io.Writer, opts ...RenderOption) error {
	// Build config with defaults
	config := &renderConfig{
		layout: LayoutDot,
	}
	for _, opt := range opts {
		opt.applyRender(config)
	}

	// Find the Graphviz binary
	binary, err := findGraphviz(config.layout)
	if err != nil {
		return err
	}

	// Generate DOT string
	dotString := g.String()

	// Execute Graphviz command: binary -Tformat
	//nolint:gosec // G204: binary path is validated via exec.LookPath in findGraphviz
	cmd := exec.CommandContext(context.TODO(), binary, "-T"+string(format))
	cmd.Stdin = strings.NewReader(dotString)

	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	// Run the command
	if err := cmd.Run(); err != nil {
		// Wrap error with stderr output
		return &RenderError{
			Err:      ErrRenderFailed,
			Stderr:   stderr.String(),
			ExitCode: cmd.ProcessState.ExitCode(),
		}
	}

	// Write output to writer
	_, err = io.Copy(w, &stdout)
	return err
}

// RenderToFile renders the graph to a file in the specified format.
// Creates the file, renders to it, and closes it. On error, attempts to clean up the partial file.
func (g *Graph) RenderToFile(format Format, path string, opts ...RenderOption) error {
	// Create the file
	file, err := os.Create(path)
	if err != nil {
		return fmt.Errorf("failed to create file: %w", err)
	}

	// Render to file
	renderErr := g.Render(format, file, opts...)

	// Close the file
	closeErr := file.Close()

	// If rendering failed, clean up the partial file
	if renderErr != nil {
		_ = os.Remove(path)
		return renderErr
	}

	// Return close error if any
	return closeErr
}

// RenderBytes renders the graph and returns the output as a byte slice.
func (g *Graph) RenderBytes(format Format, opts ...RenderOption) ([]byte, error) {
	var buf bytes.Buffer
	err := g.Render(format, &buf, opts...)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}
