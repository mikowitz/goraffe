package goraffe

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Test GraphAttributes zero value
func TestGraphAttributes_ZeroValue(t *testing.T) {
	t.Run("all fields are empty or zero", func(t *testing.T) {
		asrt := assert.New(t)
		var attrs GraphAttributes

		asrt.Empty(attrs.Label(), "expected Label to be empty")
		asrt.Empty(attrs.RankDir(), "expected RankDir to be empty")
		asrt.Empty(attrs.BgColor(), "expected BgColor to be empty")
		asrt.Empty(attrs.FontName(), "expected FontName to be empty")
		asrt.Equal(0.0, attrs.FontSize(), "expected FontSize to be zero")
		asrt.Empty(attrs.Splines(), "expected Splines to be empty")
		asrt.Equal(0.0, attrs.NodeSep(), "expected NodeSep to be zero")
		asrt.Equal(0.0, attrs.RankSep(), "expected RankSep to be zero")
		asrt.False(attrs.Compound(), "expected Compound to be false")
	})
}

// Test GraphAttributes.Custom() method
func TestGraphAttributes_Custom_ReturnsCopy(t *testing.T) {
	t.Run("returns empty map when custom is nil", func(t *testing.T) {
		asrt := assert.New(t)
		var attrs GraphAttributes

		custom := attrs.Custom()
		asrt.NotNil(custom, "expected Custom() to return empty map, not nil")
		asrt.Empty(custom, "expected Custom() to return empty map when custom field is nil")
	})

	t.Run("each call returns a different map instance", func(t *testing.T) {
		asrt := assert.New(t)
		attrs := GraphAttributes{}

		custom1 := attrs.Custom()
		custom2 := attrs.Custom()

		// Modify one map
		custom1["test"] = "value"

		// The other map should not have this modification
		_, exists := custom2["test"]
		asrt.False(exists, "expected each call to Custom() to return a new map instance")
	})

	t.Run("modifying returned map does not affect original", func(t *testing.T) {
		asrt := assert.New(t)
		attrs := GraphAttributes{}

		// Get the custom map and modify it
		custom1 := attrs.Custom()
		custom1["test"] = "value"

		// Get custom again - should not have the modification
		custom2 := attrs.Custom()
		asrt.Empty(custom2, "expected modification to returned copy to not affect original")

		_, exists := custom2["test"]
		asrt.False(exists, "expected modification to first copy to not appear in second copy")
	})
}

// Test Graph.Attrs() method
func TestGraph_Attrs_ReturnsGraphAttributes(t *testing.T) {
	t.Run("returns non-nil GraphAttributes", func(t *testing.T) {
		asrt := assert.New(t)
		g := NewGraph()

		attrs := g.Attrs()
		asrt.NotNil(attrs, "expected Attrs() to return non-nil GraphAttributes")
	})

	t.Run("returns same instance on multiple calls", func(t *testing.T) {
		asrt := assert.New(t)
		g := NewGraph()

		attrs1 := g.Attrs()
		attrs2 := g.Attrs()

		asrt.Same(attrs1, attrs2, "expected Attrs() to return the same instance")
	})
}

// Test WithGraphLabel option
func TestWithGraphLabel_SetsLabel(t *testing.T) {
	t.Run("sets label on new graph", func(t *testing.T) {
		asrt := assert.New(t)
		g := NewGraph(WithGraphLabel("My Graph"))

		asrt.Equal("My Graph", g.Attrs().Label(), "expected WithGraphLabel to set Label")
	})

	t.Run("sets empty label", func(t *testing.T) {
		asrt := assert.New(t)
		g := NewGraph(WithGraphLabel(""))

		asrt.Equal("", g.Attrs().Label(), "expected WithGraphLabel to set empty label")
	})

	t.Run("sets label with special characters", func(t *testing.T) {
		asrt := assert.New(t)
		g := NewGraph(WithGraphLabel("Graph: A -> B"))

		asrt.Equal("Graph: A -> B", g.Attrs().Label(), "expected WithGraphLabel to set label with special characters")
	})

	t.Run("last option wins", func(t *testing.T) {
		asrt := assert.New(t)
		g := NewGraph(WithGraphLabel("First"), WithGraphLabel("Second"))

		asrt.Equal("Second", g.Attrs().Label(), "expected last WithGraphLabel to override earlier ones")
	})
}

// Test WithRankDir option
func TestWithRankDir_SetsRankDir(t *testing.T) {
	t.Run("sets RankDirTB", func(t *testing.T) {
		asrt := assert.New(t)
		g := NewGraph(WithRankDir(RankDirTB))

		asrt.Equal(RankDirTB, g.Attrs().RankDir(), "expected WithRankDir to set RankDirTB")
	})

	t.Run("sets RankDirBT", func(t *testing.T) {
		asrt := assert.New(t)
		g := NewGraph(WithRankDir(RankDirBT))

		asrt.Equal(RankDirBT, g.Attrs().RankDir(), "expected WithRankDir to set RankDirBT")
	})

	t.Run("sets RankDirLR", func(t *testing.T) {
		asrt := assert.New(t)
		g := NewGraph(WithRankDir(RankDirLR))

		asrt.Equal(RankDirLR, g.Attrs().RankDir(), "expected WithRankDir to set RankDirLR")
	})

	t.Run("sets RankDirRL", func(t *testing.T) {
		asrt := assert.New(t)
		g := NewGraph(WithRankDir(RankDirRL))

		asrt.Equal(RankDirRL, g.Attrs().RankDir(), "expected WithRankDir to set RankDirRL")
	})

	t.Run("last option wins", func(t *testing.T) {
		asrt := assert.New(t)
		g := NewGraph(WithRankDir(RankDirTB), WithRankDir(RankDirLR))

		asrt.Equal(RankDirLR, g.Attrs().RankDir(), "expected last WithRankDir to override earlier ones")
	})
}

// Test WithBgColor option
func TestWithBgColor_SetsBgColor(t *testing.T) {
	t.Run("sets background color", func(t *testing.T) {
		asrt := assert.New(t)
		g := NewGraph(WithBgColor("white"))

		asrt.Equal("white", g.Attrs().BgColor(), "expected WithBgColor to set BgColor")
	})

	t.Run("sets hex color", func(t *testing.T) {
		asrt := assert.New(t)
		g := NewGraph(WithBgColor("#FF0000"))

		asrt.Equal("#FF0000", g.Attrs().BgColor(), "expected WithBgColor to set hex color")
	})

	t.Run("sets rgb color", func(t *testing.T) {
		asrt := assert.New(t)
		g := NewGraph(WithBgColor("rgb(255,0,0)"))

		asrt.Equal("rgb(255,0,0)", g.Attrs().BgColor(), "expected WithBgColor to set rgb color")
	})

	t.Run("last option wins", func(t *testing.T) {
		asrt := assert.New(t)
		g := NewGraph(WithBgColor("white"), WithBgColor("black"))

		asrt.Equal("black", g.Attrs().BgColor(), "expected last WithBgColor to override earlier ones")
	})
}

// Test WithGraphFontName option
func TestWithGraphFontName_SetsFontName(t *testing.T) {
	t.Run("sets font name", func(t *testing.T) {
		asrt := assert.New(t)
		g := NewGraph(WithGraphFontName("Arial"))

		asrt.Equal("Arial", g.Attrs().FontName(), "expected WithGraphFontName to set FontName")
	})

	t.Run("sets different font names", func(t *testing.T) {
		asrt := assert.New(t)

		fonts := []string{"Helvetica", "Times New Roman", "Courier", "Verdana"}
		for _, font := range fonts {
			g := NewGraph(WithGraphFontName(font))
			asrt.Equal(font, g.Attrs().FontName(), "expected WithGraphFontName to set FontName to %s", font)
		}
	})

	t.Run("last option wins", func(t *testing.T) {
		asrt := assert.New(t)
		g := NewGraph(WithGraphFontName("Arial"), WithGraphFontName("Helvetica"))

		asrt.Equal("Helvetica", g.Attrs().FontName(), "expected last WithGraphFontName to override earlier ones")
	})
}

// Test WithGraphFontSize option
func TestWithGraphFontSize_SetsFontSize(t *testing.T) {
	t.Run("sets font size", func(t *testing.T) {
		asrt := assert.New(t)
		g := NewGraph(WithGraphFontSize(12.0))

		asrt.Equal(12.0, g.Attrs().FontSize(), "expected WithGraphFontSize to set FontSize")
	})

	t.Run("sets different font sizes", func(t *testing.T) {
		asrt := assert.New(t)

		sizes := []float64{8.0, 10.0, 12.0, 14.0, 16.0, 24.0}
		for _, size := range sizes {
			g := NewGraph(WithGraphFontSize(size))
			asrt.Equal(size, g.Attrs().FontSize(), "expected WithGraphFontSize to set FontSize to %f", size)
		}
	})

	t.Run("sets zero font size", func(t *testing.T) {
		asrt := assert.New(t)
		g := NewGraph(WithGraphFontSize(0.0))

		asrt.Equal(0.0, g.Attrs().FontSize(), "expected WithGraphFontSize to set zero FontSize")
	})

	t.Run("last option wins", func(t *testing.T) {
		asrt := assert.New(t)
		g := NewGraph(WithGraphFontSize(12.0), WithGraphFontSize(16.0))

		asrt.Equal(16.0, g.Attrs().FontSize(), "expected last WithGraphFontSize to override earlier ones")
	})
}

// Test WithSplines option
func TestWithSplines_SetsSplines(t *testing.T) {
	t.Run("sets SplineTrue", func(t *testing.T) {
		asrt := assert.New(t)
		g := NewGraph(WithSplines(SplineTrue))

		asrt.Equal(SplineTrue, g.Attrs().Splines(), "expected WithSplines to set SplineTrue")
	})

	t.Run("sets SplineFalse", func(t *testing.T) {
		asrt := assert.New(t)
		g := NewGraph(WithSplines(SplineFalse))

		asrt.Equal(SplineFalse, g.Attrs().Splines(), "expected WithSplines to set SplineFalse")
	})

	t.Run("sets SplineOrtho", func(t *testing.T) {
		asrt := assert.New(t)
		g := NewGraph(WithSplines(SplineOrtho))

		asrt.Equal(SplineOrtho, g.Attrs().Splines(), "expected WithSplines to set SplineOrtho")
	})

	t.Run("sets SplinePolyline", func(t *testing.T) {
		asrt := assert.New(t)
		g := NewGraph(WithSplines(SplinePolyline))

		asrt.Equal(SplinePolyline, g.Attrs().Splines(), "expected WithSplines to set SplinePolyline")
	})

	t.Run("sets SplineCurved", func(t *testing.T) {
		asrt := assert.New(t)
		g := NewGraph(WithSplines(SplineCurved))

		asrt.Equal(SplineCurved, g.Attrs().Splines(), "expected WithSplines to set SplineCurved")
	})

	t.Run("last option wins", func(t *testing.T) {
		asrt := assert.New(t)
		g := NewGraph(WithSplines(SplineTrue), WithSplines(SplineOrtho))

		asrt.Equal(SplineOrtho, g.Attrs().Splines(), "expected last WithSplines to override earlier ones")
	})
}

// Test WithNodeSep option
func TestWithNodeSep_SetsNodeSep(t *testing.T) {
	t.Run("sets node separation", func(t *testing.T) {
		asrt := assert.New(t)
		g := NewGraph(WithNodeSep(0.5))

		asrt.Equal(0.5, g.Attrs().NodeSep(), "expected WithNodeSep to set NodeSep")
	})

	t.Run("sets different values", func(t *testing.T) {
		asrt := assert.New(t)

		values := []float64{0.1, 0.25, 0.5, 1.0, 2.0}
		for _, val := range values {
			g := NewGraph(WithNodeSep(val))
			asrt.Equal(val, g.Attrs().NodeSep(), "expected WithNodeSep to set NodeSep to %f", val)
		}
	})

	t.Run("sets zero node separation", func(t *testing.T) {
		asrt := assert.New(t)
		g := NewGraph(WithNodeSep(0.0))

		asrt.Equal(0.0, g.Attrs().NodeSep(), "expected WithNodeSep to set zero NodeSep")
	})

	t.Run("last option wins", func(t *testing.T) {
		asrt := assert.New(t)
		g := NewGraph(WithNodeSep(0.5), WithNodeSep(1.0))

		asrt.Equal(1.0, g.Attrs().NodeSep(), "expected last WithNodeSep to override earlier ones")
	})
}

// Test WithRankSep option
func TestWithRankSep_SetsRankSep(t *testing.T) {
	t.Run("sets rank separation", func(t *testing.T) {
		asrt := assert.New(t)
		g := NewGraph(WithRankSep(0.75))

		asrt.Equal(0.75, g.Attrs().RankSep(), "expected WithRankSep to set RankSep")
	})

	t.Run("sets different values", func(t *testing.T) {
		asrt := assert.New(t)

		values := []float64{0.1, 0.5, 1.0, 1.5, 2.0}
		for _, val := range values {
			g := NewGraph(WithRankSep(val))
			asrt.Equal(val, g.Attrs().RankSep(), "expected WithRankSep to set RankSep to %f", val)
		}
	})

	t.Run("sets zero rank separation", func(t *testing.T) {
		asrt := assert.New(t)
		g := NewGraph(WithRankSep(0.0))

		asrt.Equal(0.0, g.Attrs().RankSep(), "expected WithRankSep to set zero RankSep")
	})

	t.Run("last option wins", func(t *testing.T) {
		asrt := assert.New(t)
		g := NewGraph(WithRankSep(0.75), WithRankSep(1.5))

		asrt.Equal(1.5, g.Attrs().RankSep(), "expected last WithRankSep to override earlier ones")
	})
}

// Test WithCompound option
func TestWithCompound_SetsCompound(t *testing.T) {
	t.Run("sets compound to true", func(t *testing.T) {
		asrt := assert.New(t)
		g := NewGraph(WithCompound(true))

		asrt.True(g.Attrs().Compound(), "expected WithCompound(true) to set Compound to true")
	})

	t.Run("sets compound to false", func(t *testing.T) {
		asrt := assert.New(t)
		g := NewGraph(WithCompound(false))

		asrt.False(g.Attrs().Compound(), "expected WithCompound(false) to set Compound to false")
	})

	t.Run("last option wins", func(t *testing.T) {
		asrt := assert.New(t)
		g := NewGraph(WithCompound(true), WithCompound(false))

		asrt.False(g.Attrs().Compound(), "expected last WithCompound to override earlier ones")
	})

	t.Run("true then false", func(t *testing.T) {
		asrt := assert.New(t)
		g := NewGraph(WithCompound(false), WithCompound(true))

		asrt.True(g.Attrs().Compound(), "expected last WithCompound to override earlier ones")
	})
}

// Test combining multiple GraphOptions
func TestNewGraph_WithMultipleGraphOptions(t *testing.T) {
	t.Run("combines all graph attribute options", func(t *testing.T) {
		asrt := assert.New(t)
		g := NewGraph(
			WithGraphLabel("Test Graph"),
			WithRankDir(RankDirLR),
			WithBgColor("lightgray"),
			WithGraphFontName("Helvetica"),
			WithGraphFontSize(14.0),
			WithSplines(SplineOrtho),
			WithNodeSep(0.5),
			WithRankSep(1.0),
			WithCompound(true),
		)

		attrs := g.Attrs()
		asrt.Equal("Test Graph", attrs.Label(), "expected Label to be set")
		asrt.Equal(RankDirLR, attrs.RankDir(), "expected RankDir to be set")
		asrt.Equal("lightgray", attrs.BgColor(), "expected BgColor to be set")
		asrt.Equal("Helvetica", attrs.FontName(), "expected FontName to be set")
		asrt.Equal(14.0, attrs.FontSize(), "expected FontSize to be set")
		asrt.Equal(SplineOrtho, attrs.Splines(), "expected Splines to be set")
		asrt.Equal(0.5, attrs.NodeSep(), "expected NodeSep to be set")
		asrt.Equal(1.0, attrs.RankSep(), "expected RankSep to be set")
		asrt.True(attrs.Compound(), "expected Compound to be set")
	})

	t.Run("combines graph options with existing graph options", func(t *testing.T) {
		asrt := assert.New(t)
		g := NewGraph(
			Directed,
			Strict,
			WithGraphLabel("My Graph"),
			WithRankDir(RankDirTB),
		)

		asrt.True(g.IsDirected(), "expected graph to be directed")
		asrt.True(g.IsStrict(), "expected graph to be strict")
		asrt.Equal("My Graph", g.Attrs().Label(), "expected Label to be set")
		asrt.Equal(RankDirTB, g.Attrs().RankDir(), "expected RankDir to be set")
	})

	t.Run("graph attribute options do not affect graph structure options", func(t *testing.T) {
		asrt := assert.New(t)
		g := NewGraph(
			WithGraphLabel("Test"),
			WithRankDir(RankDirLR),
		)

		// Graph should still have default structure values
		asrt.False(g.IsDirected(), "expected default graph to be undirected")
		asrt.False(g.IsStrict(), "expected default graph to be non-strict")

		// But should have the graph attributes set
		asrt.Equal("Test", g.Attrs().Label(), "expected Label to be set")
		asrt.Equal(RankDirLR, g.Attrs().RankDir(), "expected RankDir to be set")
	})

	t.Run("options can be applied in any order", func(t *testing.T) {
		asrt := assert.New(t)
		g := NewGraph(
			WithRankDir(RankDirLR),
			Directed,
			WithGraphLabel("Graph 1"),
			Strict,
			WithBgColor("white"),
		)

		asrt.True(g.IsDirected(), "expected graph to be directed")
		asrt.True(g.IsStrict(), "expected graph to be strict")
		asrt.Equal("Graph 1", g.Attrs().Label(), "expected Label to be set")
		asrt.Equal(RankDirLR, g.Attrs().RankDir(), "expected RankDir to be set")
		asrt.Equal("white", g.Attrs().BgColor(), "expected BgColor to be set")
	})
}

// Test that graph attributes don't interfere with other graph functionality
func TestNewGraph_GraphAttributesDoNotAffectNodeEdgeOperations(t *testing.T) {
	t.Run("graph with attributes can add nodes", func(t *testing.T) {
		asrt := assert.New(t)
		g := NewGraph(
			WithGraphLabel("Test Graph"),
			WithRankDir(RankDirLR),
		)

		n := NewNode("A")
		g.AddNode(n)

		asrt.Len(g.Nodes(), 1, "expected graph to have 1 node")
		asrt.Same(n, g.Nodes()[0], "expected node to be added")
	})

	t.Run("graph with attributes can add edges", func(t *testing.T) {
		asrt := assert.New(t)
		g := NewGraph(
			WithGraphLabel("Test Graph"),
			WithBgColor("white"),
		)

		n1 := NewNode("A")
		n2 := NewNode("B")
		e := g.AddEdge(n1, n2)

		asrt.Len(g.Edges(), 1, "expected graph to have 1 edge")
		asrt.Same(e, g.Edges()[0], "expected edge to be added")
	})
}
