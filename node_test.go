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
