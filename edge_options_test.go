package goraffe

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// NOTE: These tests test the private applyEdge method directly.
// Once we wire EdgeOptions into AddEdge (Prompt 11), we should revisit
// whether these tests are still needed or if they're redundant with
// the public API tests in graph_test.go/edge_test.go.
// They're useful for TDD (Red phase) but may not be needed long-term.

func TestWithEdgeLabel_SetsLabel(t *testing.T) {
	t.Run("sets label on empty attributes", func(t *testing.T) {
		asrt := assert.New(t)
		attrs := &EdgeAttributes{}

		opt := WithEdgeLabel("My Label")
		opt.applyEdge(attrs)

		asrt.Equal("My Label", attrs.Label, "expected WithEdgeLabel to set Label")
	})

	t.Run("sets empty label", func(t *testing.T) {
		asrt := assert.New(t)
		attrs := &EdgeAttributes{}

		opt := WithEdgeLabel("")
		opt.applyEdge(attrs)

		asrt.Equal("", attrs.Label, "expected WithEdgeLabel to set empty label")
	})

	t.Run("overwrites existing label", func(t *testing.T) {
		asrt := assert.New(t)
		attrs := &EdgeAttributes{Label: "Old Label"}

		opt := WithEdgeLabel("New Label")
		opt.applyEdge(attrs)

		asrt.Equal("New Label", attrs.Label, "expected WithEdgeLabel to overwrite existing label")
	})

	t.Run("does not modify other fields", func(t *testing.T) {
		asrt := assert.New(t)
		attrs := &EdgeAttributes{
			Color:     "red",
			Style:     EdgeStyleSolid,
			ArrowHead: ArrowNormal,
			ArrowTail: ArrowDot,
			Weight:    2.5,
		}

		opt := WithEdgeLabel("test")
		opt.applyEdge(attrs)

		asrt.Equal("red", attrs.Color, "expected Color to remain unchanged")
		asrt.Equal(EdgeStyleSolid, attrs.Style, "expected Style to remain unchanged")
		asrt.Equal(ArrowNormal, attrs.ArrowHead, "expected ArrowHead to remain unchanged")
		asrt.Equal(ArrowDot, attrs.ArrowTail, "expected ArrowTail to remain unchanged")
		asrt.Equal(2.5, attrs.Weight, "expected Weight to remain unchanged")
	})
}

func TestWithEdgeColor_SetsColor(t *testing.T) {
	t.Run("sets color on empty attributes", func(t *testing.T) {
		asrt := assert.New(t)
		attrs := &EdgeAttributes{}

		opt := WithEdgeColor("red")
		opt.applyEdge(attrs)

		asrt.Equal("red", attrs.Color, "expected WithEdgeColor to set Color")
	})

	t.Run("sets different color values", func(t *testing.T) {
		asrt := assert.New(t)

		colors := []string{"red", "blue", "#FF0000", "rgb(255,0,0)", ""}
		for _, color := range colors {
			attrs := &EdgeAttributes{}
			opt := WithEdgeColor(color)
			opt.applyEdge(attrs)
			asrt.Equal(color, attrs.Color, "expected WithEdgeColor to set Color to %s", color)
		}
	})

	t.Run("overwrites existing color", func(t *testing.T) {
		asrt := assert.New(t)
		attrs := &EdgeAttributes{Color: "blue"}

		opt := WithEdgeColor("red")
		opt.applyEdge(attrs)

		asrt.Equal("red", attrs.Color, "expected WithEdgeColor to overwrite existing color")
	})

	t.Run("does not modify other fields", func(t *testing.T) {
		asrt := assert.New(t)
		attrs := &EdgeAttributes{
			Label:     "test",
			Style:     EdgeStyleDashed,
			ArrowHead: ArrowVee,
			ArrowTail: ArrowNone,
			Weight:    1.0,
		}

		opt := WithEdgeColor("red")
		opt.applyEdge(attrs)

		asrt.Equal("test", attrs.Label, "expected Label to remain unchanged")
		asrt.Equal(EdgeStyleDashed, attrs.Style, "expected Style to remain unchanged")
		asrt.Equal(ArrowVee, attrs.ArrowHead, "expected ArrowHead to remain unchanged")
		asrt.Equal(ArrowNone, attrs.ArrowTail, "expected ArrowTail to remain unchanged")
		asrt.Equal(1.0, attrs.Weight, "expected Weight to remain unchanged")
	})
}

func TestWithEdgeStyle_SetsStyle(t *testing.T) {
	t.Run("sets style on empty attributes", func(t *testing.T) {
		asrt := assert.New(t)
		attrs := &EdgeAttributes{}

		opt := WithEdgeStyle(EdgeStyleDashed)
		opt.applyEdge(attrs)

		asrt.Equal(EdgeStyleDashed, attrs.Style, "expected WithEdgeStyle to set Style to EdgeStyleDashed")
	})

	t.Run("sets different styles", func(t *testing.T) {
		asrt := assert.New(t)

		styles := []EdgeStyle{EdgeStyleSolid, EdgeStyleDashed, EdgeStyleDotted, EdgeStyleBold}
		for _, style := range styles {
			attrs := &EdgeAttributes{}
			opt := WithEdgeStyle(style)
			opt.applyEdge(attrs)
			asrt.Equal(style, attrs.Style, "expected WithEdgeStyle to set Style to %s", style)
		}
	})

	t.Run("overwrites existing style", func(t *testing.T) {
		asrt := assert.New(t)
		attrs := &EdgeAttributes{Style: EdgeStyleSolid}

		opt := WithEdgeStyle(EdgeStyleDashed)
		opt.applyEdge(attrs)

		asrt.Equal(EdgeStyleDashed, attrs.Style, "expected WithEdgeStyle to overwrite existing style")
	})

	t.Run("does not modify other fields", func(t *testing.T) {
		asrt := assert.New(t)
		attrs := &EdgeAttributes{
			Label:     "test",
			Color:     "red",
			ArrowHead: ArrowDot,
			ArrowTail: ArrowVee,
			Weight:    3.0,
		}

		opt := WithEdgeStyle(EdgeStyleBold)
		opt.applyEdge(attrs)

		asrt.Equal("test", attrs.Label, "expected Label to remain unchanged")
		asrt.Equal("red", attrs.Color, "expected Color to remain unchanged")
		asrt.Equal(ArrowDot, attrs.ArrowHead, "expected ArrowHead to remain unchanged")
		asrt.Equal(ArrowVee, attrs.ArrowTail, "expected ArrowTail to remain unchanged")
		asrt.Equal(3.0, attrs.Weight, "expected Weight to remain unchanged")
	})
}

func TestWithArrowHead_SetsArrowHead(t *testing.T) {
	t.Run("sets arrow head on empty attributes", func(t *testing.T) {
		asrt := assert.New(t)
		attrs := &EdgeAttributes{}

		opt := WithArrowHead(ArrowNormal)
		opt.applyEdge(attrs)

		asrt.Equal(ArrowNormal, attrs.ArrowHead, "expected WithArrowHead to set ArrowHead")
	})

	t.Run("sets different arrow types", func(t *testing.T) {
		asrt := assert.New(t)

		arrows := []ArrowType{ArrowNormal, ArrowDot, ArrowNone, ArrowVee}
		for _, arrow := range arrows {
			attrs := &EdgeAttributes{}
			opt := WithArrowHead(arrow)
			opt.applyEdge(attrs)
			asrt.Equal(arrow, attrs.ArrowHead, "expected WithArrowHead to set ArrowHead to %s", arrow)
		}
	})

	t.Run("overwrites existing arrow head", func(t *testing.T) {
		asrt := assert.New(t)
		attrs := &EdgeAttributes{ArrowHead: ArrowNormal}

		opt := WithArrowHead(ArrowDot)
		opt.applyEdge(attrs)

		asrt.Equal(ArrowDot, attrs.ArrowHead, "expected WithArrowHead to overwrite existing arrow head")
	})

	t.Run("does not modify other fields", func(t *testing.T) {
		asrt := assert.New(t)
		attrs := &EdgeAttributes{
			Label:     "test",
			Color:     "blue",
			Style:     EdgeStyleDotted,
			ArrowTail: ArrowVee,
			Weight:    1.5,
		}

		opt := WithArrowHead(ArrowNone)
		opt.applyEdge(attrs)

		asrt.Equal("test", attrs.Label, "expected Label to remain unchanged")
		asrt.Equal("blue", attrs.Color, "expected Color to remain unchanged")
		asrt.Equal(EdgeStyleDotted, attrs.Style, "expected Style to remain unchanged")
		asrt.Equal(ArrowVee, attrs.ArrowTail, "expected ArrowTail to remain unchanged")
		asrt.Equal(1.5, attrs.Weight, "expected Weight to remain unchanged")
	})
}

func TestWithArrowTail_SetsArrowTail(t *testing.T) {
	t.Run("sets arrow tail on empty attributes", func(t *testing.T) {
		asrt := assert.New(t)
		attrs := &EdgeAttributes{}

		opt := WithArrowTail(ArrowVee)
		opt.applyEdge(attrs)

		asrt.Equal(ArrowVee, attrs.ArrowTail, "expected WithArrowTail to set ArrowTail")
	})

	t.Run("sets different arrow types", func(t *testing.T) {
		asrt := assert.New(t)

		arrows := []ArrowType{ArrowNormal, ArrowDot, ArrowNone, ArrowVee}
		for _, arrow := range arrows {
			attrs := &EdgeAttributes{}
			opt := WithArrowTail(arrow)
			opt.applyEdge(attrs)
			asrt.Equal(arrow, attrs.ArrowTail, "expected WithArrowTail to set ArrowTail to %s", arrow)
		}
	})

	t.Run("overwrites existing arrow tail", func(t *testing.T) {
		asrt := assert.New(t)
		attrs := &EdgeAttributes{ArrowTail: ArrowNone}

		opt := WithArrowTail(ArrowDot)
		opt.applyEdge(attrs)

		asrt.Equal(ArrowDot, attrs.ArrowTail, "expected WithArrowTail to overwrite existing arrow tail")
	})

	t.Run("does not modify other fields", func(t *testing.T) {
		asrt := assert.New(t)
		attrs := &EdgeAttributes{
			Label:     "test",
			Color:     "green",
			Style:     EdgeStyleBold,
			ArrowHead: ArrowDot,
			Weight:    4.0,
		}

		opt := WithArrowTail(ArrowNormal)
		opt.applyEdge(attrs)

		asrt.Equal("test", attrs.Label, "expected Label to remain unchanged")
		asrt.Equal("green", attrs.Color, "expected Color to remain unchanged")
		asrt.Equal(EdgeStyleBold, attrs.Style, "expected Style to remain unchanged")
		asrt.Equal(ArrowDot, attrs.ArrowHead, "expected ArrowHead to remain unchanged")
		asrt.Equal(4.0, attrs.Weight, "expected Weight to remain unchanged")
	})
}

func TestWithWeight_SetsWeight(t *testing.T) {
	t.Run("sets weight on empty attributes", func(t *testing.T) {
		asrt := assert.New(t)
		attrs := &EdgeAttributes{}

		opt := WithWeight(2.5)
		opt.applyEdge(attrs)

		asrt.Equal(2.5, attrs.Weight, "expected WithWeight to set Weight")
	})

	t.Run("sets different weight values", func(t *testing.T) {
		asrt := assert.New(t)

		weights := []float64{0.0, 1.0, 2.5, 10.0, 100.5}
		for _, weight := range weights {
			attrs := &EdgeAttributes{}
			opt := WithWeight(weight)
			opt.applyEdge(attrs)
			asrt.Equal(weight, attrs.Weight, "expected WithWeight to set Weight to %f", weight)
		}
	})

	t.Run("overwrites existing weight", func(t *testing.T) {
		asrt := assert.New(t)
		attrs := &EdgeAttributes{Weight: 1.0}

		opt := WithWeight(5.0)
		opt.applyEdge(attrs)

		asrt.Equal(5.0, attrs.Weight, "expected WithWeight to overwrite existing weight")
	})

	t.Run("does not modify other fields", func(t *testing.T) {
		asrt := assert.New(t)
		attrs := &EdgeAttributes{
			Label:     "test",
			Color:     "yellow",
			Style:     EdgeStyleSolid,
			ArrowHead: ArrowVee,
			ArrowTail: ArrowNone,
		}

		opt := WithWeight(3.5)
		opt.applyEdge(attrs)

		asrt.Equal("test", attrs.Label, "expected Label to remain unchanged")
		asrt.Equal("yellow", attrs.Color, "expected Color to remain unchanged")
		asrt.Equal(EdgeStyleSolid, attrs.Style, "expected Style to remain unchanged")
		asrt.Equal(ArrowVee, attrs.ArrowHead, "expected ArrowHead to remain unchanged")
		asrt.Equal(ArrowNone, attrs.ArrowTail, "expected ArrowTail to remain unchanged")
	})
}

func TestEdgeOption_MultipleOptionsCanBeApplied(t *testing.T) {
	t.Run("applies multiple options in sequence", func(t *testing.T) {
		asrt := assert.New(t)
		attrs := &EdgeAttributes{}

		WithEdgeLabel("Test").applyEdge(attrs)
		WithEdgeColor("red").applyEdge(attrs)
		WithEdgeStyle(EdgeStyleDashed).applyEdge(attrs)
		WithArrowHead(ArrowNormal).applyEdge(attrs)
		WithArrowTail(ArrowDot).applyEdge(attrs)
		WithWeight(2.5).applyEdge(attrs)

		asrt.Equal("Test", attrs.Label, "expected Label to be set")
		asrt.Equal("red", attrs.Color, "expected Color to be set")
		asrt.Equal(EdgeStyleDashed, attrs.Style, "expected Style to be set")
		asrt.Equal(ArrowNormal, attrs.ArrowHead, "expected ArrowHead to be set")
		asrt.Equal(ArrowDot, attrs.ArrowTail, "expected ArrowTail to be set")
		asrt.Equal(2.5, attrs.Weight, "expected Weight to be set")
	})

	t.Run("later options override earlier ones for same field", func(t *testing.T) {
		asrt := assert.New(t)
		attrs := &EdgeAttributes{}

		WithEdgeStyle(EdgeStyleSolid).applyEdge(attrs)
		WithEdgeStyle(EdgeStyleDashed).applyEdge(attrs)

		asrt.Equal(EdgeStyleDashed, attrs.Style, "expected later option to override earlier one")
	})
}

func TestEdgeAttributes_AsOption(t *testing.T) {
	t.Run("applies EdgeAttributes as an option", func(t *testing.T) {
		asrt := assert.New(t)

		template := EdgeAttributes{
			Style:     EdgeStyleDashed,
			Color:     "gray",
			ArrowHead: ArrowNormal,
		}

		target := &EdgeAttributes{}
		template.applyEdge(target)

		asrt.Equal(EdgeStyleDashed, target.Style, "expected Style to be set from EdgeAttributes")
		asrt.Equal("gray", target.Color, "expected Color to be set from EdgeAttributes")
		asrt.Equal(ArrowNormal, target.ArrowHead, "expected ArrowHead to be set from EdgeAttributes")
	})

	t.Run("only copies non-zero fields", func(t *testing.T) {
		asrt := assert.New(t)

		// Create attrs with only some fields set
		template := EdgeAttributes{
			Style: EdgeStyleBold,
			Color: "blue",
			// Label, ArrowHead, ArrowTail, Weight are zero values
		}

		// Set up target with some existing values
		target := &EdgeAttributes{
			Label:  "Important",
			Weight: 3.0,
		}

		template.applyEdge(target)

		asrt.Equal(EdgeStyleBold, target.Style, "expected Style to be set from template")
		asrt.Equal("blue", target.Color, "expected Color to be set from template")
		asrt.Equal("Important", target.Label, "expected Label to NOT be overridden by template's zero value")
		asrt.Equal(3.0, target.Weight, "expected Weight to NOT be overridden by template's zero value")
		asrt.Empty(target.ArrowHead, "expected ArrowHead to remain empty")
		asrt.Empty(target.ArrowTail, "expected ArrowTail to remain empty")
	})

	t.Run("can be combined with other options", func(t *testing.T) {
		asrt := assert.New(t)

		template := EdgeAttributes{
			Style: EdgeStyleDashed,
			Color: "gray",
		}

		target := &EdgeAttributes{}
		template.applyEdge(target)
		WithEdgeLabel("Custom").applyEdge(target)
		WithWeight(2.0).applyEdge(target)

		asrt.Equal(EdgeStyleDashed, target.Style, "expected Style from template")
		asrt.Equal("gray", target.Color, "expected Color from template")
		asrt.Equal("Custom", target.Label, "expected Label from option")
		asrt.Equal(2.0, target.Weight, "expected Weight from option")
	})

	t.Run("later options override template values", func(t *testing.T) {
		asrt := assert.New(t)

		template := EdgeAttributes{
			Style: EdgeStyleDashed,
			Label: "Template Label",
		}

		target := &EdgeAttributes{}
		template.applyEdge(target)
		WithEdgeLabel("Override").applyEdge(target)

		asrt.Equal(EdgeStyleDashed, target.Style, "expected Style from template")
		asrt.Equal("Override", target.Label, "expected Label to be overridden by later option")
	})
}
