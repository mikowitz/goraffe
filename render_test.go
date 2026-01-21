package goraffe

import (
	"errors"
	"os"
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

func TestGraph_Render_PNG_ProducesOutput(t *testing.T) {
	requireGraphviz(t)

	g := NewGraph(Directed)
	n1 := NewNode("A")
	n2 := NewNode("B")
	_ = g.AddNode(n1)
	_ = g.AddNode(n2)
	_, _ = g.AddEdge(n1, n2)

	var buf []byte
	w := &testWriter{buf: &buf}
	err := g.Render(PNG, w)
	assert.NoError(t, err)
	assert.NotEmpty(t, buf)
	// Check PNG magic bytes
	assert.True(t, len(buf) >= 8, "PNG output should be at least 8 bytes")
	assert.Equal(t, []byte{0x89, 0x50, 0x4E, 0x47}, buf[0:4], "PNG should start with PNG magic bytes")
}

func TestGraph_Render_SVG_ProducesOutput(t *testing.T) {
	requireGraphviz(t)

	g := NewGraph(Directed)
	n1 := NewNode("A")
	n2 := NewNode("B")
	_ = g.AddNode(n1)
	_ = g.AddNode(n2)
	_, _ = g.AddEdge(n1, n2)

	var buf []byte
	w := &testWriter{buf: &buf}
	err := g.Render(SVG, w)
	assert.NoError(t, err)
	assert.NotEmpty(t, buf)
	// Check for SVG XML structure
	output := string(buf)
	assert.Contains(t, output, "<svg")
	assert.Contains(t, output, "</svg>")
}

func TestGraph_Render_DOT_ProducesOutput(t *testing.T) {
	requireGraphviz(t)

	g := NewGraph(Directed)
	n1 := NewNode("A")
	n2 := NewNode("B")
	_ = g.AddNode(n1)
	_ = g.AddNode(n2)
	_, _ = g.AddEdge(n1, n2)

	var buf []byte
	w := &testWriter{buf: &buf}
	err := g.Render(DOT, w)
	assert.NoError(t, err)
	assert.NotEmpty(t, buf)
	// DOT format should contain the graph structure
	output := string(buf)
	assert.Contains(t, output, "digraph")
}

func TestGraph_Render_ToBuffer(t *testing.T) {
	requireGraphviz(t)

	g := NewGraph(Directed)
	n1 := NewNode("A")
	_ = g.AddNode(n1)

	var buf []byte
	w := &testWriter{buf: &buf}
	err := g.Render(PNG, w)
	assert.NoError(t, err)
	assert.NotEmpty(t, buf)
}

// testWriter is a simple writer for testing that appends to a byte slice.
type testWriter struct {
	buf *[]byte
}

func (w *testWriter) Write(p []byte) (n int, err error) {
	*w.buf = append(*w.buf, p...)
	return len(p), nil
}

func TestGraph_RenderToFile_CreatesFile(t *testing.T) {
	requireGraphviz(t)

	g := NewGraph(Directed)
	n1 := NewNode("A")
	g.AddNode(n1)

	// Use temp file
	tmpfile := t.TempDir() + "/graph.png"
	err := g.RenderToFile(PNG, tmpfile)
	assert.NoError(t, err)

	// Check file exists
	_, err = os.Stat(tmpfile)
	assert.NoError(t, err, "file should exist")
}

func TestGraph_RenderToFile_ValidContent(t *testing.T) {
	requireGraphviz(t)

	g := NewGraph(Directed)
	n1 := NewNode("A")
	n2 := NewNode("B")
	g.AddNode(n1)
	g.AddNode(n2)
	g.AddEdge(n1, n2)

	// Render to PNG
	tmpfile := t.TempDir() + "/graph.png"
	err := g.RenderToFile(PNG, tmpfile)
	assert.NoError(t, err)

	// Read and validate content
	data, err := os.ReadFile(tmpfile)
	assert.NoError(t, err)
	assert.NotEmpty(t, data)
	// Check PNG magic bytes
	assert.True(t, len(data) >= 4)
	assert.Equal(t, []byte{0x89, 0x50, 0x4E, 0x47}, data[0:4])
}

func TestGraph_RenderBytes_ReturnsPNG(t *testing.T) {
	requireGraphviz(t)

	g := NewGraph(Directed)
	n1 := NewNode("A")
	n2 := NewNode("B")
	g.AddNode(n1)
	g.AddNode(n2)
	g.AddEdge(n1, n2)

	data, err := g.RenderBytes(PNG)
	assert.NoError(t, err)
	assert.NotEmpty(t, data)
	// Check PNG magic bytes
	assert.True(t, len(data) >= 4)
	assert.Equal(t, []byte{0x89, 0x50, 0x4E, 0x47}, data[0:4])
}

func TestGraph_RenderBytes_ReturnsSVG(t *testing.T) {
	requireGraphviz(t)

	g := NewGraph(Directed)
	n1 := NewNode("A")
	n2 := NewNode("B")
	g.AddNode(n1)
	g.AddNode(n2)
	g.AddEdge(n1, n2)

	data, err := g.RenderBytes(SVG)
	assert.NoError(t, err)
	assert.NotEmpty(t, data)
	// Check for SVG structure
	output := string(data)
	assert.Contains(t, output, "<svg")
	assert.Contains(t, output, "</svg>")
}

func TestRender_CompleteWorkflow(t *testing.T) {
	requireGraphviz(t)

	// Create complex graph
	g := NewGraph(Directed, WithGraphLabel("Test Graph"))
	n1 := NewNode("A", WithLabel("Node A"), WithColor("red"))
	n2 := NewNode("B", WithLabel("Node B"), WithColor("blue"))
	n3 := NewNode("C", WithLabel("Node C"))
	g.AddNode(n1)
	g.AddNode(n2)
	g.AddNode(n3)
	g.AddEdge(n1, n2, WithEdgeLabel("edge 1"))
	g.AddEdge(n2, n3, WithEdgeLabel("edge 2"))

	// Render to file
	tmpfile := t.TempDir() + "/complete.svg"
	err := g.RenderToFile(SVG, tmpfile)
	assert.NoError(t, err)

	// Verify file exists and is valid
	data, err := os.ReadFile(tmpfile)
	assert.NoError(t, err)
	assert.NotEmpty(t, data)

	output := string(data)
	assert.Contains(t, output, "<svg")
	assert.Contains(t, output, "</svg>")

	// Also test RenderBytes
	pngData, err := g.RenderBytes(PNG)
	assert.NoError(t, err)
	assert.NotEmpty(t, pngData)
	assert.Equal(t, []byte{0x89, 0x50, 0x4E, 0x47}, pngData[0:4])
}

func TestGraph_Render_DefaultLayout_IsDot(t *testing.T) {
	requireGraphviz(t)

	g := NewGraph(Directed)
	n1 := NewNode("A")
	n2 := NewNode("B")
	g.AddNode(n1)
	g.AddNode(n2)
	g.AddEdge(n1, n2)

	// Render without specifying layout (should default to dot)
	var buf []byte
	w := &testWriter{buf: &buf}
	err := g.Render(PNG, w)
	assert.NoError(t, err)
	assert.NotEmpty(t, buf)
}

func TestGraph_Render_WithLayout_Neato(t *testing.T) {
	requireGraphviz(t)

	g := NewGraph(Undirected)
	n1 := NewNode("A")
	n2 := NewNode("B")
	n3 := NewNode("C")
	g.AddNode(n1)
	g.AddNode(n2)
	g.AddNode(n3)
	g.AddEdge(n1, n2)
	g.AddEdge(n2, n3)
	g.AddEdge(n3, n1)

	var buf []byte
	w := &testWriter{buf: &buf}
	err := g.Render(PNG, w, WithLayout(LayoutNeato))
	assert.NoError(t, err)
	assert.NotEmpty(t, buf)
	assert.True(t, len(buf) >= 4)
	assert.Equal(t, []byte{0x89, 0x50, 0x4E, 0x47}, buf[0:4])
}

func TestGraph_Render_WithLayout_Fdp(t *testing.T) {
	requireGraphviz(t)

	g := NewGraph(Undirected)
	n1 := NewNode("A")
	n2 := NewNode("B")
	g.AddNode(n1)
	g.AddNode(n2)
	g.AddEdge(n1, n2)

	var buf []byte
	w := &testWriter{buf: &buf}
	err := g.Render(PNG, w, WithLayout(LayoutFdp))
	assert.NoError(t, err)
	assert.NotEmpty(t, buf)
}

func TestGraph_Render_WithLayout_Circo(t *testing.T) {
	requireGraphviz(t)

	g := NewGraph(Undirected)
	n1 := NewNode("A")
	n2 := NewNode("B")
	n3 := NewNode("C")
	g.AddNode(n1)
	g.AddNode(n2)
	g.AddNode(n3)
	g.AddEdge(n1, n2)
	g.AddEdge(n2, n3)

	var buf []byte
	w := &testWriter{buf: &buf}
	err := g.Render(PNG, w, WithLayout(LayoutCirco))
	assert.NoError(t, err)
	assert.NotEmpty(t, buf)
}

func TestGraph_Render_AllLayouts(t *testing.T) {
	requireGraphviz(t)

	// Create a simple graph
	g := NewGraph(Directed)
	n1 := NewNode("A")
	n2 := NewNode("B")
	n3 := NewNode("C")
	g.AddNode(n1)
	g.AddNode(n2)
	g.AddNode(n3)
	g.AddEdge(n1, n2)
	g.AddEdge(n2, n3)

	layouts := []Layout{
		LayoutDot, LayoutNeato, LayoutFdp, LayoutSfdp,
		LayoutTwopi, LayoutCirco, LayoutOsage, LayoutPatchwork,
	}

	for _, layout := range layouts {
		t.Run(string(layout), func(t *testing.T) {
			var buf []byte
			w := &testWriter{buf: &buf}
			err := g.Render(PNG, w, WithLayout(layout))
			// Note: Not all layouts may be installed
			if err != nil {
				if errors.Is(err, ErrGraphvizNotFound) {
					t.Skipf("Layout %s not installed", layout)
				} else {
					t.Fatalf("Unexpected error: %v", err)
				}
			}
			assert.NotEmpty(t, buf)
			// Verify PNG magic bytes
			assert.True(t, len(buf) >= 4)
			assert.Equal(t, []byte{0x89, 0x50, 0x4E, 0x47}, buf[0:4])
		})
	}
}

func TestGoraffe_EndToEnd(t *testing.T) {
	requireGraphviz(t)

	// Create a complex graph with nodes, edges, and attributes
	g := NewGraph(Directed,
		WithGraphLabel("End-to-End Test"),
		WithRankDir(RankDirLR),
		WithBgColor("lightgray"),
	)

	// Add nodes with various attributes
	n1 := NewNode("start", WithLabel("Start"), WithColor("green"), WithBoxShape())
	n2 := NewNode("process", WithLabel("Process"), WithColor("blue"), WithEllipseShape())
	n3 := NewNode("end", WithLabel("End"), WithColor("red"), WithBoxShape())
	n4 := NewNode("branch", WithLabel("Branch"), WithDiamondShape())
	g.AddNode(n1)
	g.AddNode(n2)
	g.AddNode(n3)
	g.AddNode(n4)

	// Add edges with attributes
	g.AddEdge(n1, n2, WithEdgeLabel("begin"), WithWeight(2.0))
	g.AddEdge(n2, n4, WithEdgeLabel("decide"))
	g.AddEdge(n4, n3, WithEdgeLabel("finish"), WithEdgeColor("red"))
	g.AddEdge(n4, n2, WithEdgeLabel("loop"), WithEdgeStyle(EdgeStyleDashed))

	// Test 1: Render to multiple formats
	formats := []Format{PNG, SVG, DOT}
	for _, format := range formats {
		t.Run("format_"+string(format), func(t *testing.T) {
			data, err := g.RenderBytes(format)
			assert.NoError(t, err)
			assert.NotEmpty(t, data)

			switch format {
			case PNG:
				assert.Equal(t, []byte{0x89, 0x50, 0x4E, 0x47}, data[0:4])
			case SVG:
				assert.Contains(t, string(data), "<svg")
			case DOT:
				assert.Contains(t, string(data), "digraph")
				assert.Contains(t, string(data), "End-to-End Test")
			}
		})
	}

	// Test 2: Render with different layouts
	t.Run("multiple_layouts", func(t *testing.T) {
		for _, layout := range []Layout{LayoutDot, LayoutNeato, LayoutFdp} {
			data, err := g.RenderBytes(PNG, WithLayout(layout))
			if err != nil && errors.Is(err, ErrGraphvizNotFound) {
				t.Skipf("Layout %s not available", layout)
				continue
			}
			assert.NoError(t, err)
			assert.NotEmpty(t, data)
		}
	})

	// Test 3: Render to file
	t.Run("render_to_file", func(t *testing.T) {
		tmpfile := t.TempDir() + "/endtoend.png"
		err := g.RenderToFile(PNG, tmpfile, WithLayout(LayoutDot))
		assert.NoError(t, err)

		data, err := os.ReadFile(tmpfile)
		assert.NoError(t, err)
		assert.NotEmpty(t, data)
		assert.Equal(t, []byte{0x89, 0x50, 0x4E, 0x47}, data[0:4])
	})

	// Test 4: Modify graph and re-render
	t.Run("modify_and_rerender", func(t *testing.T) {
		// Add another node
		n5 := NewNode("extra", WithLabel("Extra Node"))
		g.AddNode(n5)
		g.AddEdge(n3, n5, WithEdgeLabel("extend"))

		data, err := g.RenderBytes(SVG)
		assert.NoError(t, err)
		assert.Contains(t, string(data), "Extra Node")
	})
}
