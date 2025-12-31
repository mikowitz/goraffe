package goraffe

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewGraph_Directed(t *testing.T) {
	t.Run("sets directed to true", func(t *testing.T) {
		asrt := assert.New(t)

		g := NewGraph(Directed)

		asrt.True(g.IsDirected(), "expected graph to be directed")
		asrt.False(g.IsStrict(), "expected graph to not be strict by default")
		asrt.Empty(g.Name(), "expected graph to have empty name by default")
	})
}

func TestNewGraph_Undirected(t *testing.T) {
	t.Run("explicitly sets directed to false", func(t *testing.T) {
		asrt := assert.New(t)

		g := NewGraph(Undirected)

		asrt.False(g.IsDirected(), "expected graph to be undirected")
		asrt.False(g.IsStrict(), "expected graph to not be strict by default")
	})

	t.Run("undirected is the default", func(t *testing.T) {
		asrt := assert.New(t)

		g := NewGraph()

		asrt.False(g.IsDirected(), "expected graph to be undirected by default")
	})
}

func TestNewGraph_Strict(t *testing.T) {
	t.Run("sets strict to true", func(t *testing.T) {
		asrt := assert.New(t)

		g := NewGraph(Strict)

		asrt.True(g.IsStrict(), "expected graph to be strict")
		asrt.False(g.IsDirected(), "expected graph to be undirected by default")
	})
}

func TestNewGraph_DirectedAndStrict(t *testing.T) {
	t.Run("applies both options", func(t *testing.T) {
		asrt := assert.New(t)

		g := NewGraph(Directed, Strict)

		asrt.True(g.IsDirected(), "expected graph to be directed")
		asrt.True(g.IsStrict(), "expected graph to be strict")
	})

	t.Run("order of options does not matter", func(t *testing.T) {
		asrt := assert.New(t)

		g := NewGraph(Strict, Directed)

		asrt.True(g.IsDirected(), "expected graph to be directed")
		asrt.True(g.IsStrict(), "expected graph to be strict")
	})
}

func TestNewGraph_MultipleOptions_LastWins(t *testing.T) {
	t.Run("Directed then Undirected results in undirected", func(t *testing.T) {
		asrt := assert.New(t)

		g := NewGraph(Directed, Undirected)

		asrt.False(g.IsDirected(), "expected last option (Undirected) to win")
	})

	t.Run("Undirected then Directed results in directed", func(t *testing.T) {
		asrt := assert.New(t)

		g := NewGraph(Undirected, Directed)

		asrt.True(g.IsDirected(), "expected last option (Directed) to win")
	})
}

func TestNewGraph_NoOptions(t *testing.T) {
	t.Run("creates graph with default values", func(t *testing.T) {
		asrt := assert.New(t)

		g := NewGraph()

		asrt.False(g.IsDirected(), "expected default graph to be undirected")
		asrt.False(g.IsStrict(), "expected default graph to be non-strict")
		asrt.Empty(g.Name(), "expected default graph to have empty name")
		asrt.NotNil(g.Nodes(), "expected nodes to be initialized")
		asrt.NotNil(g.Edges(), "expected edges to be initialized")
		asrt.Len(g.Nodes(), 0, "expected new graph to have no nodes")
		asrt.Len(g.Edges(), 0, "expected new graph to have no edges")
	})
}
