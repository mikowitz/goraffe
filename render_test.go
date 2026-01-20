package goraffe

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFormat_StringValues(t *testing.T) {
	tests := []struct {
		name     string
		format   Format
		expected string
	}{
		{"PNG format", PNG, "png"},
		{"SVG format", SVG, "svg"},
		{"PDF format", PDF, "pdf"},
		{"DOT format", DOT, "dot"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expected, string(tt.format))
		})
	}
}

func TestLayout_StringValues(t *testing.T) {
	tests := []struct {
		name     string
		layout   Layout
		expected string
	}{
		{"dot layout", LayoutDot, "dot"},
		{"neato layout", LayoutNeato, "neato"},
		{"fdp layout", LayoutFdp, "fdp"},
		{"sfdp layout", LayoutSfdp, "sfdp"},
		{"twopi layout", LayoutTwopi, "twopi"},
		{"circo layout", LayoutCirco, "circo"},
		{"osage layout", LayoutOsage, "osage"},
		{"patchwork layout", LayoutPatchwork, "patchwork"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expected, string(tt.layout))
		})
	}
}
