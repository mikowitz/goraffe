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

func TestSubgraph_NestedSubgraph(t *testing.T) {
	asrt := assert.New(t)

	g := NewGraph()
	var outer *Subgraph
	var inner *Subgraph

	g.Subgraph("cluster_outer", func(o *Subgraph) {
		outer = o
		o.SetLabel("Outer")
		o.Subgraph("cluster_inner", func(i *Subgraph) {
			inner = i
			i.SetLabel("Inner")
		})
	})

	asrt.NotNil(outer, "expected outer subgraph to exist")
	asrt.NotNil(inner, "expected inner subgraph to exist")
	asrt.Equal("cluster_outer", outer.Name(), "expected outer subgraph name to be 'cluster_outer'")
	asrt.Equal("cluster_inner", inner.Name(), "expected inner subgraph name to be 'cluster_inner'")
	asrt.Equal("Outer", outer.Attrs().Label(), "expected outer label to be 'Outer'")
	asrt.Equal("Inner", inner.Attrs().Label(), "expected inner label to be 'Inner'")

	// Verify inner is in outer's subgraphs
	outerSubs := outer.Subgraphs()
	asrt.Len(outerSubs, 1, "expected outer to have 1 nested subgraph")
	asrt.Equal(inner, outerSubs[0], "expected outer's nested subgraph to be inner")
}

func TestSubgraph_NestedSubgraph_NodesInRoot(t *testing.T) {
	asrt := assert.New(t)

	g := NewGraph()
	n1 := NewNode("n1")
	n2 := NewNode("n2")

	g.Subgraph("cluster_outer", func(outer *Subgraph) {
		outer.SetLabel("Outer")
		_ = outer.AddNode(n1)

		outer.Subgraph("cluster_inner", func(inner *Subgraph) {
			inner.SetLabel("Inner")
			_ = inner.AddNode(n2)
		})
	})

	// Verify both nodes are in the root graph
	rootNodes := g.Nodes()
	asrt.Len(rootNodes, 2, "expected 2 nodes in root graph")

	nodeIDs := make(map[string]bool)
	for _, n := range rootNodes {
		nodeIDs[n.ID()] = true
	}

	asrt.True(nodeIDs["n1"], "expected node 'n1' in root graph")
	asrt.True(nodeIDs["n2"], "expected node 'n2' in root graph")
}

func TestSubgraph_DeeplyNested(t *testing.T) {
	asrt := assert.New(t)

	g := NewGraph()
	var level1, level2, level3 *Subgraph

	g.Subgraph("cluster_1", func(l1 *Subgraph) {
		level1 = l1
		l1.SetLabel("Level 1")

		l1.Subgraph("cluster_2", func(l2 *Subgraph) {
			level2 = l2
			l2.SetLabel("Level 2")

			l2.Subgraph("cluster_3", func(l3 *Subgraph) {
				level3 = l3
				l3.SetLabel("Level 3")
			})
		})
	})

	asrt.NotNil(level1, "expected level1 subgraph to exist")
	asrt.NotNil(level2, "expected level2 subgraph to exist")
	asrt.NotNil(level3, "expected level3 subgraph to exist")

	asrt.Equal("cluster_1", level1.Name(), "expected level1 name")
	asrt.Equal("cluster_2", level2.Name(), "expected level2 name")
	asrt.Equal("cluster_3", level3.Name(), "expected level3 name")

	// Verify nesting structure
	level1Subs := level1.Subgraphs()
	asrt.Len(level1Subs, 1, "expected level1 to have 1 nested subgraph")
	asrt.Equal(level2, level1Subs[0], "expected level1's nested subgraph to be level2")

	level2Subs := level2.Subgraphs()
	asrt.Len(level2Subs, 1, "expected level2 to have 1 nested subgraph")
	asrt.Equal(level3, level2Subs[0], "expected level2's nested subgraph to be level3")

	level3Subs := level3.Subgraphs()
	asrt.Len(level3Subs, 0, "expected level3 to have no nested subgraphs")
}

func TestSubgraph_NestedCluster(t *testing.T) {
	asrt := assert.New(t)

	g := NewGraph()
	n1 := NewNode("A")
	n2 := NewNode("B")

	var outer *Subgraph
	g.Subgraph("cluster_outer", func(o *Subgraph) {
		outer = o
		o.SetLabel("Outer Cluster")
		o.SetStyle("filled")
		o.SetFillColor("lightgray")

		_ = o.AddNode(n1)

		o.Subgraph("cluster_inner", func(i *Subgraph) {
			i.SetLabel("Inner Cluster")
			i.SetStyle("filled")
			i.SetFillColor("white")
			_ = i.AddNode(n2)
		})
	})

	asrt.True(outer.IsCluster(), "expected outer to be a cluster")

	innerSubs := outer.Subgraphs()
	asrt.Len(innerSubs, 1, "expected outer to have 1 nested subgraph")

	inner := innerSubs[0]
	asrt.True(inner.IsCluster(), "expected inner to be a cluster")
	asrt.Equal("Inner Cluster", inner.Attrs().Label(), "expected inner label")

	// Verify both nodes are in root graph
	rootNodes := g.Nodes()
	asrt.Len(rootNodes, 2, "expected 2 nodes in root graph")
}
