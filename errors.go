// ABOUTME: Defines error types for rendering operations.
// ABOUTME: Provides RenderError with stderr output and sentinel error values.
package goraffe

import (
	"errors"
	"fmt"
)

// RenderError represents an error that occurred during graph rendering.
type RenderError struct {
	// Err is the underlying error.
	Err error
	// Stderr contains the stderr output from Graphviz.
	Stderr string
	// ExitCode is the exit code from the Graphviz process.
	ExitCode int
}

// Error implements the error interface.
func (e *RenderError) Error() string {
	if e.Stderr != "" {
		// Include a snippet of stderr (first 200 chars) for context
		stderr := e.Stderr
		if len(stderr) > 200 {
			stderr = stderr[:200] + "..."
		}
		return fmt.Sprintf("%v (exit code %d): %s", e.Err, e.ExitCode, stderr)
	}
	return fmt.Sprintf("%v (exit code %d)", e.Err, e.ExitCode)
}

// Unwrap returns the underlying error.
func (e *RenderError) Unwrap() error {
	return e.Err
}

// Sentinel errors for common rendering failures.
var (
	// ErrGraphvizNotFound indicates that Graphviz is not installed or not in PATH.
	ErrGraphvizNotFound = errors.New("goraffe: graphviz not found in PATH")
	// ErrInvalidDOT indicates that the DOT syntax is invalid.
	ErrInvalidDOT = errors.New("goraffe: invalid DOT syntax")
	// ErrRenderFailed indicates that rendering failed for an unknown reason.
	ErrRenderFailed = errors.New("goraffe: rendering failed")
)
