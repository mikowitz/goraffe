// ABOUTME: Tests for subgraph functionality in Graphviz DOT graphs.
// ABOUTME: Verifies subgraph creation, cluster identification, and node/edge management.
package goraffe

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSubgraph_Name(t *testing.T) {
	asrt := assert.New(t)

	g := NewGraph()
	var sg *Subgraph
	g.Subgraph("my_subgraph", func(s *Subgraph) {
		sg = s
	})

	asrt.Equal("my_subgraph", sg.Name(), "expected name to be 'my_subgraph'")
}

func TestSubgraph_IsCluster_True(t *testing.T) {
	asrt := assert.New(t)

	g := NewGraph()
	var sg *Subgraph
	g.Subgraph("cluster_group", func(s *Subgraph) {
		sg = s
	})

	asrt.True(sg.IsCluster(), "expected IsCluster() to return true for name starting with 'cluster'")
}

func TestSubgraph_IsCluster_False(t *testing.T) {
	asrt := assert.New(t)

	g := NewGraph()
	var sg *Subgraph
	g.Subgraph("my_subgraph", func(s *Subgraph) {
		sg = s
	})

	asrt.False(sg.IsCluster(), "expected IsCluster() to return false for name not starting with 'cluster'")
}

func TestSubgraph_AddNode(t *testing.T) {
	asrt := assert.New(t)

	g := NewGraph()
	var sg *Subgraph
	g.Subgraph("sub", func(s *Subgraph) {
		sg = s
	})

	n := NewNode("A")
	_ = sg.AddNode(n)

	nodes := sg.Nodes()
	asrt.Len(nodes, 1, "expected 1 node in subgraph")
	asrt.Equal("A", nodes[0].ID(), "expected node to be 'A'")
}

func TestSubgraph_AddNode_AlsoAddsToParent(t *testing.T) {
	asrt := assert.New(t)

	g := NewGraph()
	var sg *Subgraph
	g.Subgraph("sub", func(s *Subgraph) {
		sg = s
	})

	n := NewNode("A")
	_ = sg.AddNode(n)

	// Check that node was also added to parent graph
	parentNodes := g.Nodes()
	asrt.Len(parentNodes, 1, "expected 1 node in parent graph")
	asrt.Equal("A", parentNodes[0].ID(), "expected node 'A' in parent")
}

func TestGraph_Subgraph_CallsFunction(t *testing.T) {
	asrt := assert.New(t)

	g := NewGraph()
	called := false

	g.Subgraph("sub", func(s *Subgraph) {
		called = true
	})

	asrt.True(called, "expected subgraph function to be called")
}

func TestGraph_Subgraph_ReturnsSubgraph(t *testing.T) {
	asrt := assert.New(t)

	g := NewGraph()

	sg := g.Subgraph("sub", func(s *Subgraph) {})

	asrt.NotNil(sg, "expected Subgraph to return non-nil subgraph")
	asrt.Equal("sub", sg.Name(), "expected returned subgraph name to be 'sub'")
}

func TestGraph_Subgraphs_ReturnsAll(t *testing.T) {
	asrt := assert.New(t)

	g := NewGraph()

	g.Subgraph("sub1", func(s *Subgraph) {})
	g.Subgraph("sub2", func(s *Subgraph) {})
	g.Subgraph("cluster_3", func(s *Subgraph) {})

	subs := g.Subgraphs()
	asrt.Len(subs, 3, "expected 3 subgraphs")

	// Verify names
	names := make(map[string]bool)
	for _, sg := range subs {
		names[sg.Name()] = true
	}

	asrt.True(names["sub1"], "expected subgraph 'sub1' to exist")
	asrt.True(names["sub2"], "expected subgraph 'sub2' to exist")
	asrt.True(names["cluster_3"], "expected subgraph 'cluster_3' to exist")
}

func TestSubgraph_SetLabel(t *testing.T) {
	asrt := assert.New(t)

	g := NewGraph()
	var sg *Subgraph
	g.Subgraph("sub", func(s *Subgraph) {
		sg = s
	})

	sg.SetLabel("My Label")

	asrt.Equal("My Label", sg.Attrs().Label(), "expected label to be 'My Label'")
}

func TestSubgraph_SetStyle(t *testing.T) {
	asrt := assert.New(t)

	g := NewGraph()
	var sg *Subgraph
	g.Subgraph("sub", func(s *Subgraph) {
		sg = s
	})

	sg.SetStyle("filled")

	asrt.Equal("filled", sg.Attrs().Style(), "expected style to be 'filled'")
}

func TestSubgraph_SetAttribute(t *testing.T) {
	asrt := assert.New(t)

	g := NewGraph()
	var sg *Subgraph
	g.Subgraph("sub", func(s *Subgraph) {
		sg = s
	})

	sg.SetAttribute("rank", "same")

	custom := sg.Attrs().Custom()
	asrt.Equal("same", custom["rank"], "expected custom attribute 'rank' to be 'same'")
}

func TestSubgraph_Attrs_ReturnsAttributes(t *testing.T) {
	asrt := assert.New(t)

	g := NewGraph()
	var sg *Subgraph
	g.Subgraph("sub", func(s *Subgraph) {
		sg = s
	})

	attrs := sg.Attrs()

	asrt.NotNil(attrs, "expected Attrs() to return non-nil attributes")
	// Calling Attrs() again should return the same instance
	attrs2 := sg.Attrs()
	asrt.Equal(attrs, attrs2, "expected Attrs() to return the same instance")
}

func TestSubgraph_Cluster_CanHaveStyle(t *testing.T) {
	asrt := assert.New(t)

	g := NewGraph()
	var sg *Subgraph
	g.Subgraph("cluster_test", func(s *Subgraph) {
		sg = s
	})

	sg.SetStyle("filled")
	sg.SetFillColor("lightblue")
	sg.SetColor("blue")
	sg.SetLabel("Test Cluster")

	asrt.True(sg.IsCluster(), "expected subgraph to be a cluster")
	asrt.Equal("filled", sg.Attrs().Style(), "expected style to be 'filled'")
	asrt.Equal("lightblue", sg.Attrs().FillColor(), "expected fill color to be 'lightblue'")
	asrt.Equal("blue", sg.Attrs().Color(), "expected color to be 'blue'")
	asrt.Equal("Test Cluster", sg.Attrs().Label(), "expected label to be 'Test Cluster'")
}
