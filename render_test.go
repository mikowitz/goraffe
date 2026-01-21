package goraffe

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// requireGraphviz skips the test if Graphviz is not installed.
func requireGraphviz(t *testing.T) {
	t.Helper()
	if err := checkGraphvizInstalled(); err != nil {
		t.Skip("Graphviz not installed, skipping test")
	}
}

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

func TestFindGraphviz_Dot(t *testing.T) {
	requireGraphviz(t)

	path, err := findGraphviz(LayoutDot)
	assert.NoError(t, err)
	assert.NotEmpty(t, path)
	assert.Contains(t, path, "dot")
}

func TestFindGraphviz_AllLayouts(t *testing.T) {
	requireGraphviz(t)

	layouts := []Layout{
		LayoutDot, LayoutNeato, LayoutFdp, LayoutSfdp,
		LayoutTwopi, LayoutCirco, LayoutOsage, LayoutPatchwork,
	}

	for _, layout := range layouts {
		t.Run(string(layout), func(t *testing.T) {
			path, err := findGraphviz(layout)
			// Note: Not all layouts may be installed
			if err != nil {
				assert.ErrorIs(t, err, ErrGraphvizNotFound)
			} else {
				assert.NotEmpty(t, path)
			}
		})
	}
}

func TestFindGraphviz_InvalidLayout(t *testing.T) {
	path, err := findGraphviz(Layout("nonexistent"))
	assert.ErrorIs(t, err, ErrGraphvizNotFound)
	assert.Empty(t, path)
}

func TestGraphvizVersion_ReturnsVersion(t *testing.T) {
	requireGraphviz(t)

	version, err := GraphvizVersion()
	assert.NoError(t, err)
	assert.NotEmpty(t, version)
	// Version string should contain "graphviz" or "dot"
	containsGraphviz := assert.Contains(t, version, "graphviz") ||
		assert.Contains(t, version, "dot")
	assert.True(t, containsGraphviz, "version string should contain 'graphviz' or 'dot'")
}
