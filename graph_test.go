package goraffe

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewGraph_DefaultValues(t *testing.T) {
	asrt := assert.New(t)

	g := NewGraph()

	asrt.False(g.directed, "expected directed to be false")
	asrt.False(g.strict, "expected strict to be false")
	asrt.Empty(g.name, "expected name to be empty")
}

func TestIsDirected(t *testing.T) {
	t.Run("returns false for undirected graph", func(t *testing.T) {
		asrt := assert.New(t)

		g := NewGraph()
		g.directed = false

		asrt.False(g.IsDirected(), "expected IsDirected to return false")
	})

	t.Run("returns true for directed graph", func(t *testing.T) {
		asrt := assert.New(t)

		g := NewGraph()
		g.directed = true

		asrt.True(g.IsDirected(), "expected IsDirected to return true")
	})
}

func TestIsStrict(t *testing.T) {
	t.Run("returns false for non-strict graph", func(t *testing.T) {
		asrt := assert.New(t)

		g := NewGraph()
		g.strict = false

		asrt.False(g.IsStrict(), "expected IsStrict to return false")
	})

	t.Run("returns true for strict graph", func(t *testing.T) {
		asrt := assert.New(t)

		g := NewGraph()
		g.strict = true

		asrt.True(g.IsStrict(), "expected IsStrict to return true")
	})
}

func TestName(t *testing.T) {
	t.Run("returns empty string for unnamed graph", func(t *testing.T) {
		asrt := assert.New(t)

		g := NewGraph()
		g.name = ""

		asrt.Empty(g.Name(), "expected Name to return empty string")
	})

	t.Run("returns name for named graph", func(t *testing.T) {
		asrt := assert.New(t)

		g := NewGraph()
		g.name = "TestGraph"

		asrt.Equal("TestGraph", g.Name(), "expected Name to return 'TestGraph'")
	})
}
