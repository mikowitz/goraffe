package goraffe

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewNode_SetsID(t *testing.T) {
	asrt := assert.New(t)

	n := NewNode("node1")

	asrt.Equal("node1", n.id, "expected node ID to be set to 'node1'")
}

func TestNewNode_EmptyID(t *testing.T) {
	asrt := assert.New(t)

	n := NewNode("")

	asrt.NotNil(n, "expected NewNode to return a non-nil node even with empty ID")
	asrt.Empty(n.id, "expected node ID to be empty string")
}

func TestNode_ID_ReturnsCorrectValue(t *testing.T) {
	t.Run("returns simple ID", func(t *testing.T) {
		asrt := assert.New(t)

		n := NewNode("simple")

		asrt.Equal("simple", n.ID(), "expected ID() to return 'simple'")
	})

	t.Run("returns ID with special characters", func(t *testing.T) {
		asrt := assert.New(t)

		n := NewNode("node-with-dashes")

		asrt.Equal("node-with-dashes", n.ID(), "expected ID() to return exact ID with dashes")
	})

	t.Run("returns ID with spaces", func(t *testing.T) {
		asrt := assert.New(t)

		n := NewNode("node with spaces")

		asrt.Equal("node with spaces", n.ID(), "expected ID() to return exact ID with spaces")
	})

	t.Run("returns ID with numbers", func(t *testing.T) {
		asrt := assert.New(t)

		n := NewNode("node123")

		asrt.Equal("node123", n.ID(), "expected ID() to return exact ID with numbers")
	})

	t.Run("returns empty ID unchanged", func(t *testing.T) {
		asrt := assert.New(t)

		n := NewNode("")

		asrt.Empty(n.ID(), "expected ID() to return empty string")
	})
}

func TestNode_Attrs_ReturnsAttributes(t *testing.T) {
	t.Run("returns non-nil attributes", func(t *testing.T) {
		asrt := assert.New(t)

		n := NewNode("test")

		attrs := n.Attrs()
		asrt.NotNil(attrs, "expected Attrs() to return non-nil NodeAttributes")
	})

	t.Run("returns same instance on multiple calls", func(t *testing.T) {
		asrt := assert.New(t)

		n := NewNode("test")

		attrs1 := n.Attrs()
		attrs2 := n.Attrs()

		asrt.Same(attrs1, attrs2, "expected Attrs() to return the same NodeAttributes instance")
	})

	t.Run("returns attributes with zero values for new node", func(t *testing.T) {
		asrt := assert.New(t)

		n := NewNode("test")

		attrs := n.Attrs()
		asrt.Empty(attrs.Label, "expected Label to be empty for new node")
		asrt.Empty(attrs.Shape, "expected Shape to be empty for new node")
		asrt.Empty(attrs.Color, "expected Color to be empty for new node")
		asrt.Empty(attrs.FillColor, "expected FillColor to be empty for new node")
		asrt.Empty(attrs.FontName, "expected FontName to be empty for new node")
		asrt.Equal(0.0, attrs.FontSize, "expected FontSize to be zero for new node")
	})
}

func TestNewNode_WithOptions(t *testing.T) {
	t.Run("applies single option", func(t *testing.T) {
		asrt := assert.New(t)

		n := NewNode("test", WithShape(ShapeBox))

		asrt.Equal(ShapeBox, n.Attrs().Shape, "expected Shape to be set to ShapeBox")
	})

	t.Run("applies WithLabel option", func(t *testing.T) {
		asrt := assert.New(t)

		n := NewNode("test", WithLabel("My Label"))

		asrt.Equal("My Label", n.Attrs().Label, "expected Label to be set")
	})

	t.Run("applies WithColor option", func(t *testing.T) {
		asrt := assert.New(t)

		n := NewNode("test", WithColor("red"))

		asrt.Equal("red", n.Attrs().Color, "expected Color to be set to red")
	})

	t.Run("applies WithFillColor option", func(t *testing.T) {
		asrt := assert.New(t)

		n := NewNode("test", WithFillColor("blue"))

		asrt.Equal("blue", n.Attrs().FillColor, "expected FillColor to be set to blue")
	})

	t.Run("applies WithFontName option", func(t *testing.T) {
		asrt := assert.New(t)

		n := NewNode("test", WithFontName("Arial"))

		asrt.Equal("Arial", n.Attrs().FontName, "expected FontName to be set to Arial")
	})

	t.Run("applies WithFontSize option", func(t *testing.T) {
		asrt := assert.New(t)

		n := NewNode("test", WithFontSize(14.0))

		asrt.Equal(14.0, n.Attrs().FontSize, "expected FontSize to be set to 14.0")
	})

	t.Run("works with no options", func(t *testing.T) {
		asrt := assert.New(t)

		n := NewNode("test")

		asrt.NotNil(n, "expected NewNode to work with no options")
		asrt.Equal("test", n.ID(), "expected ID to be set correctly")
	})
}

func TestNewNode_WithMultipleOptions(t *testing.T) {
	t.Run("applies multiple different options", func(t *testing.T) {
		asrt := assert.New(t)

		n := NewNode("test",
			WithShape(ShapeBox),
			WithLabel("My Node"),
			WithColor("red"),
		)

		asrt.Equal(ShapeBox, n.Attrs().Shape, "expected Shape to be set")
		asrt.Equal("My Node", n.Attrs().Label, "expected Label to be set")
		asrt.Equal("red", n.Attrs().Color, "expected Color to be set")
	})

	t.Run("applies all available option types", func(t *testing.T) {
		asrt := assert.New(t)

		n := NewNode("test",
			WithShape(ShapeCircle),
			WithLabel("All Options"),
			WithColor("blue"),
			WithFillColor("lightblue"),
			WithFontName("Helvetica"),
			WithFontSize(16.0),
		)

		asrt.Equal(ShapeCircle, n.Attrs().Shape, "expected Shape to be set")
		asrt.Equal("All Options", n.Attrs().Label, "expected Label to be set")
		asrt.Equal("blue", n.Attrs().Color, "expected Color to be set")
		asrt.Equal("lightblue", n.Attrs().FillColor, "expected FillColor to be set")
		asrt.Equal("Helvetica", n.Attrs().FontName, "expected FontName to be set")
		asrt.Equal(16.0, n.Attrs().FontSize, "expected FontSize to be set")
	})
}

func TestNewNode_WithNodeAttributesStruct(t *testing.T) {
	t.Run("applies NodeAttributes as an option", func(t *testing.T) {
		asrt := assert.New(t)

		commonAttrs := NodeAttributes{
			Shape:    ShapeBox,
			FontName: "Arial",
			FontSize: 12.0,
		}

		n := NewNode("test", commonAttrs)

		asrt.Equal(ShapeBox, n.Attrs().Shape, "expected Shape to be set from NodeAttributes")
		asrt.Equal("Arial", n.Attrs().FontName, "expected FontName to be set from NodeAttributes")
		asrt.Equal(12.0, n.Attrs().FontSize, "expected FontSize to be set from NodeAttributes")
	})

	t.Run("applies NodeAttributes with additional options", func(t *testing.T) {
		asrt := assert.New(t)

		commonAttrs := NodeAttributes{
			Shape:    ShapeBox,
			FontName: "Arial",
		}

		n := NewNode("test", commonAttrs, WithLabel("Custom Label"))

		asrt.Equal(ShapeBox, n.Attrs().Shape, "expected Shape from NodeAttributes")
		asrt.Equal("Arial", n.Attrs().FontName, "expected FontName from NodeAttributes")
		asrt.Equal("Custom Label", n.Attrs().Label, "expected Label from WithLabel option")
	})

	t.Run("NodeAttributes only copies non-zero fields", func(t *testing.T) {
		asrt := assert.New(t)

		// Create attrs with only some fields set
		commonAttrs := NodeAttributes{
			Shape: ShapeBox,
			// Label, Color, etc. are zero values
		}

		n := NewNode("test", commonAttrs, WithLabel("Test"))

		asrt.Equal(ShapeBox, n.Attrs().Shape, "expected Shape to be set")
		asrt.Equal("Test", n.Attrs().Label, "expected Label to be set")
		asrt.Empty(n.Attrs().Color, "expected Color to remain empty (zero value not copied)")
	})
}

func TestNewNode_OptionsAppliedInOrder(t *testing.T) {
	t.Run("later options override earlier ones", func(t *testing.T) {
		asrt := assert.New(t)

		n := NewNode("test",
			WithShape(ShapeCircle),
			WithShape(ShapeBox),
		)

		asrt.Equal(ShapeBox, n.Attrs().Shape, "expected later WithShape to override earlier one")
	})

	t.Run("later label overrides earlier label", func(t *testing.T) {
		asrt := assert.New(t)

		n := NewNode("test",
			WithLabel("First"),
			WithLabel("Second"),
		)

		asrt.Equal("Second", n.Attrs().Label, "expected later label to override earlier one")
	})

	t.Run("NodeAttributes then individual options", func(t *testing.T) {
		asrt := assert.New(t)

		commonAttrs := NodeAttributes{
			Shape: ShapeCircle,
			Label: "From Attrs",
		}

		n := NewNode("test", commonAttrs, WithLabel("Overridden"))

		asrt.Equal(ShapeCircle, n.Attrs().Shape, "expected Shape from NodeAttributes")
		asrt.Equal("Overridden", n.Attrs().Label, "expected Label to be overridden by later option")
	})

	t.Run("individual options then NodeAttributes", func(t *testing.T) {
		asrt := assert.New(t)

		commonAttrs := NodeAttributes{
			Shape: ShapeBox,
			Label: "From Attrs",
		}

		n := NewNode("test", WithShape(ShapeCircle), commonAttrs)

		asrt.Equal(ShapeBox, n.Attrs().Shape, "expected Shape from later NodeAttributes to override")
		asrt.Equal("From Attrs", n.Attrs().Label, "expected Label from NodeAttributes")
	})

	t.Run("NodeAttributes with zero values should not override non-zero values", func(t *testing.T) {
		asrt := assert.New(t)

		// This test demonstrates the selective copying requirement
		// A reusable NodeAttributes template with only some fields set
		template := NodeAttributes{
			Shape:    ShapeBox,
			FontName: "Arial",
			// Label, Color, FillColor, FontSize are zero values
		}

		// Apply options first, then the template
		// The template should only override Shape and FontName, not Label
		n := NewNode("test",
			WithLabel("Important Label"),
			WithColor("red"),
			template, // This should NOT override Label and Color to empty strings
		)

		asrt.Equal(ShapeBox, n.Attrs().Shape, "expected Shape from template")
		asrt.Equal("Arial", n.Attrs().FontName, "expected FontName from template")
		asrt.Equal("Important Label", n.Attrs().Label, "expected Label to NOT be overridden by template's zero value")
		asrt.Equal("red", n.Attrs().Color, "expected Color to NOT be overridden by template's zero value")
		asrt.Empty(n.Attrs().FillColor, "expected FillColor to remain empty")
		asrt.Equal(0.0, n.Attrs().FontSize, "expected FontSize to remain zero")
	})
}
