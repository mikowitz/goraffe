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
