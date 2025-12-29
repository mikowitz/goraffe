package goraffe

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWithShape_SetsShape(t *testing.T) {
	t.Run("sets shape on empty attributes", func(t *testing.T) {
		asrt := assert.New(t)
		attrs := &NodeAttributes{}

		opt := WithShape(ShapeBox)
		opt.applyNode(attrs)

		asrt.Equal(ShapeBox, attrs.Shape, "expected WithShape to set Shape to ShapeBox")
	})

	t.Run("sets different shapes", func(t *testing.T) {
		asrt := assert.New(t)

		shapes := []Shape{ShapeBox, ShapeCircle, ShapeEllipse, ShapeDiamond, ShapeRecord, ShapePlaintext}
		for _, shape := range shapes {
			attrs := &NodeAttributes{}
			opt := WithShape(shape)
			opt.applyNode(attrs)
			asrt.Equal(shape, attrs.Shape, "expected WithShape to set Shape to %s", shape)
		}
	})

	t.Run("overwrites existing shape", func(t *testing.T) {
		asrt := assert.New(t)
		attrs := &NodeAttributes{Shape: ShapeCircle}

		opt := WithShape(ShapeBox)
		opt.applyNode(attrs)

		asrt.Equal(ShapeBox, attrs.Shape, "expected WithShape to overwrite existing shape")
	})

	t.Run("does not modify other fields", func(t *testing.T) {
		asrt := assert.New(t)
		attrs := &NodeAttributes{
			Label:     "test",
			Color:     "red",
			FillColor: "blue",
			FontName:  "Arial",
			FontSize:  12.0,
		}

		opt := WithShape(ShapeBox)
		opt.applyNode(attrs)

		asrt.Equal("test", attrs.Label, "expected Label to remain unchanged")
		asrt.Equal("red", attrs.Color, "expected Color to remain unchanged")
		asrt.Equal("blue", attrs.FillColor, "expected FillColor to remain unchanged")
		asrt.Equal("Arial", attrs.FontName, "expected FontName to remain unchanged")
		asrt.Equal(12.0, attrs.FontSize, "expected FontSize to remain unchanged")
	})
}

func TestWithLabel_SetsLabel(t *testing.T) {
	t.Run("sets label on empty attributes", func(t *testing.T) {
		asrt := assert.New(t)
		attrs := &NodeAttributes{}

		opt := WithLabel("My Label")
		opt.applyNode(attrs)

		asrt.Equal("My Label", attrs.Label, "expected WithLabel to set Label")
	})

	t.Run("sets empty label", func(t *testing.T) {
		asrt := assert.New(t)
		attrs := &NodeAttributes{}

		opt := WithLabel("")
		opt.applyNode(attrs)

		asrt.Equal("", attrs.Label, "expected WithLabel to set empty label")
	})

	t.Run("sets label with special characters", func(t *testing.T) {
		asrt := assert.New(t)
		attrs := &NodeAttributes{}

		opt := WithLabel("Label with \"quotes\" and\nnewlines")
		opt.applyNode(attrs)

		asrt.Equal("Label with \"quotes\" and\nnewlines", attrs.Label, "expected WithLabel to set label with special characters")
	})

	t.Run("overwrites existing label", func(t *testing.T) {
		asrt := assert.New(t)
		attrs := &NodeAttributes{Label: "Old Label"}

		opt := WithLabel("New Label")
		opt.applyNode(attrs)

		asrt.Equal("New Label", attrs.Label, "expected WithLabel to overwrite existing label")
	})

	t.Run("does not modify other fields", func(t *testing.T) {
		asrt := assert.New(t)
		attrs := &NodeAttributes{
			Shape:     ShapeBox,
			Color:     "red",
			FillColor: "blue",
			FontName:  "Arial",
			FontSize:  12.0,
		}

		opt := WithLabel("test")
		opt.applyNode(attrs)

		asrt.Equal(ShapeBox, attrs.Shape, "expected Shape to remain unchanged")
		asrt.Equal("red", attrs.Color, "expected Color to remain unchanged")
		asrt.Equal("blue", attrs.FillColor, "expected FillColor to remain unchanged")
		asrt.Equal("Arial", attrs.FontName, "expected FontName to remain unchanged")
		asrt.Equal(12.0, attrs.FontSize, "expected FontSize to remain unchanged")
	})
}

func TestWithColor_SetsColor(t *testing.T) {
	t.Run("sets color on empty attributes", func(t *testing.T) {
		asrt := assert.New(t)
		attrs := &NodeAttributes{}

		opt := WithColor("red")
		opt.applyNode(attrs)

		asrt.Equal("red", attrs.Color, "expected WithColor to set Color")
	})

	t.Run("sets different color values", func(t *testing.T) {
		asrt := assert.New(t)

		colors := []string{"red", "blue", "#FF0000", "rgb(255,0,0)", ""}
		for _, color := range colors {
			attrs := &NodeAttributes{}
			opt := WithColor(color)
			opt.applyNode(attrs)
			asrt.Equal(color, attrs.Color, "expected WithColor to set Color to %s", color)
		}
	})

	t.Run("overwrites existing color", func(t *testing.T) {
		asrt := assert.New(t)
		attrs := &NodeAttributes{Color: "blue"}

		opt := WithColor("red")
		opt.applyNode(attrs)

		asrt.Equal("red", attrs.Color, "expected WithColor to overwrite existing color")
	})

	t.Run("does not modify other fields", func(t *testing.T) {
		asrt := assert.New(t)
		attrs := &NodeAttributes{
			Label:     "test",
			Shape:     ShapeBox,
			FillColor: "blue",
			FontName:  "Arial",
			FontSize:  12.0,
		}

		opt := WithColor("red")
		opt.applyNode(attrs)

		asrt.Equal("test", attrs.Label, "expected Label to remain unchanged")
		asrt.Equal(ShapeBox, attrs.Shape, "expected Shape to remain unchanged")
		asrt.Equal("blue", attrs.FillColor, "expected FillColor to remain unchanged")
		asrt.Equal("Arial", attrs.FontName, "expected FontName to remain unchanged")
		asrt.Equal(12.0, attrs.FontSize, "expected FontSize to remain unchanged")
	})
}

func TestWithFillColor_SetsFillColor(t *testing.T) {
	t.Run("sets fill color on empty attributes", func(t *testing.T) {
		asrt := assert.New(t)
		attrs := &NodeAttributes{}

		opt := WithFillColor("blue")
		opt.applyNode(attrs)

		asrt.Equal("blue", attrs.FillColor, "expected WithFillColor to set FillColor")
	})

	t.Run("sets different fill color values", func(t *testing.T) {
		asrt := assert.New(t)

		colors := []string{"red", "blue", "#00FF00", "lightgray", ""}
		for _, color := range colors {
			attrs := &NodeAttributes{}
			opt := WithFillColor(color)
			opt.applyNode(attrs)
			asrt.Equal(color, attrs.FillColor, "expected WithFillColor to set FillColor to %s", color)
		}
	})

	t.Run("overwrites existing fill color", func(t *testing.T) {
		asrt := assert.New(t)
		attrs := &NodeAttributes{FillColor: "red"}

		opt := WithFillColor("blue")
		opt.applyNode(attrs)

		asrt.Equal("blue", attrs.FillColor, "expected WithFillColor to overwrite existing fill color")
	})

	t.Run("does not modify other fields", func(t *testing.T) {
		asrt := assert.New(t)
		attrs := &NodeAttributes{
			Label:    "test",
			Shape:    ShapeBox,
			Color:    "red",
			FontName: "Arial",
			FontSize: 12.0,
		}

		opt := WithFillColor("blue")
		opt.applyNode(attrs)

		asrt.Equal("test", attrs.Label, "expected Label to remain unchanged")
		asrt.Equal(ShapeBox, attrs.Shape, "expected Shape to remain unchanged")
		asrt.Equal("red", attrs.Color, "expected Color to remain unchanged")
		asrt.Equal("Arial", attrs.FontName, "expected FontName to remain unchanged")
		asrt.Equal(12.0, attrs.FontSize, "expected FontSize to remain unchanged")
	})
}

func TestWithFontName_SetsFontName(t *testing.T) {
	t.Run("sets font name on empty attributes", func(t *testing.T) {
		asrt := assert.New(t)
		attrs := &NodeAttributes{}

		opt := WithFontName("Arial")
		opt.applyNode(attrs)

		asrt.Equal("Arial", attrs.FontName, "expected WithFontName to set FontName")
	})

	t.Run("sets different font names", func(t *testing.T) {
		asrt := assert.New(t)

		fonts := []string{"Arial", "Helvetica", "Times New Roman", "Courier", ""}
		for _, font := range fonts {
			attrs := &NodeAttributes{}
			opt := WithFontName(font)
			opt.applyNode(attrs)
			asrt.Equal(font, attrs.FontName, "expected WithFontName to set FontName to %s", font)
		}
	})

	t.Run("overwrites existing font name", func(t *testing.T) {
		asrt := assert.New(t)
		attrs := &NodeAttributes{FontName: "Arial"}

		opt := WithFontName("Helvetica")
		opt.applyNode(attrs)

		asrt.Equal("Helvetica", attrs.FontName, "expected WithFontName to overwrite existing font name")
	})

	t.Run("does not modify other fields", func(t *testing.T) {
		asrt := assert.New(t)
		attrs := &NodeAttributes{
			Label:     "test",
			Shape:     ShapeBox,
			Color:     "red",
			FillColor: "blue",
			FontSize:  12.0,
		}

		opt := WithFontName("Arial")
		opt.applyNode(attrs)

		asrt.Equal("test", attrs.Label, "expected Label to remain unchanged")
		asrt.Equal(ShapeBox, attrs.Shape, "expected Shape to remain unchanged")
		asrt.Equal("red", attrs.Color, "expected Color to remain unchanged")
		asrt.Equal("blue", attrs.FillColor, "expected FillColor to remain unchanged")
		asrt.Equal(12.0, attrs.FontSize, "expected FontSize to remain unchanged")
	})
}

func TestWithFontSize_SetsFontSize(t *testing.T) {
	t.Run("sets font size on empty attributes", func(t *testing.T) {
		asrt := assert.New(t)
		attrs := &NodeAttributes{}

		opt := WithFontSize(14.0)
		opt.applyNode(attrs)

		asrt.Equal(14.0, attrs.FontSize, "expected WithFontSize to set FontSize")
	})

	t.Run("sets different font sizes", func(t *testing.T) {
		asrt := assert.New(t)

		sizes := []float64{8.0, 10.0, 12.0, 14.5, 16.0, 24.0, 0.0}
		for _, size := range sizes {
			attrs := &NodeAttributes{}
			opt := WithFontSize(size)
			opt.applyNode(attrs)
			asrt.Equal(size, attrs.FontSize, "expected WithFontSize to set FontSize to %f", size)
		}
	})

	t.Run("overwrites existing font size", func(t *testing.T) {
		asrt := assert.New(t)
		attrs := &NodeAttributes{FontSize: 12.0}

		opt := WithFontSize(16.0)
		opt.applyNode(attrs)

		asrt.Equal(16.0, attrs.FontSize, "expected WithFontSize to overwrite existing font size")
	})

	t.Run("does not modify other fields", func(t *testing.T) {
		asrt := assert.New(t)
		attrs := &NodeAttributes{
			Label:     "test",
			Shape:     ShapeBox,
			Color:     "red",
			FillColor: "blue",
			FontName:  "Arial",
		}

		opt := WithFontSize(14.0)
		opt.applyNode(attrs)

		asrt.Equal("test", attrs.Label, "expected Label to remain unchanged")
		asrt.Equal(ShapeBox, attrs.Shape, "expected Shape to remain unchanged")
		asrt.Equal("red", attrs.Color, "expected Color to remain unchanged")
		asrt.Equal("blue", attrs.FillColor, "expected FillColor to remain unchanged")
		asrt.Equal("Arial", attrs.FontName, "expected FontName to remain unchanged")
	})
}

func TestNodeOption_MultipleOptionsCanBeApplied(t *testing.T) {
	t.Run("applies multiple options in sequence", func(t *testing.T) {
		asrt := assert.New(t)
		attrs := &NodeAttributes{}

		WithShape(ShapeBox).applyNode(attrs)
		WithLabel("Test").applyNode(attrs)
		WithColor("red").applyNode(attrs)
		WithFillColor("blue").applyNode(attrs)
		WithFontName("Arial").applyNode(attrs)
		WithFontSize(14.0).applyNode(attrs)

		asrt.Equal(ShapeBox, attrs.Shape, "expected Shape to be set")
		asrt.Equal("Test", attrs.Label, "expected Label to be set")
		asrt.Equal("red", attrs.Color, "expected Color to be set")
		asrt.Equal("blue", attrs.FillColor, "expected FillColor to be set")
		asrt.Equal("Arial", attrs.FontName, "expected FontName to be set")
		asrt.Equal(14.0, attrs.FontSize, "expected FontSize to be set")
	})

	t.Run("later options override earlier ones for same field", func(t *testing.T) {
		asrt := assert.New(t)
		attrs := &NodeAttributes{}

		WithShape(ShapeCircle).applyNode(attrs)
		WithShape(ShapeBox).applyNode(attrs)

		asrt.Equal(ShapeBox, attrs.Shape, "expected later option to override earlier one")
	})
}
