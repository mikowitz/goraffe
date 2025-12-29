# Goraffe Implementation Blueprint

## High-Level Build Phases

1. **Foundation** - Core types, basic graph construction
2. **Attributes** - Type-safe attribute system with functional options
3. **DOT Generation** - Output valid DOT strings
4. **Labels** - HTML and record label builders
5. **Subgraphs** - Clusters, nesting, rank constraints
6. **Parsing** - DOT input parsing
7. **Rendering** - Graphviz CLI integration

---

## Detailed Chunk Breakdown

### Phase 1: Foundation

1.1. Project scaffolding and basic types
1.2. Node type with ID and basic creation
1.3. Edge type connecting nodes
1.4. Graph type holding nodes and edges
1.5. AddNode and AddEdge methods with implicit node addition

### Phase 2: Attributes

2.1. Attribute enums (Shape, Style basics)
2.2. NodeAttributes struct
2.3. NodeOption interface and functional options
2.4. EdgeAttributes struct and EdgeOption interface
2.5. GraphAttributes and GraphOption interface
2.6. Default node/edge attributes on graph
2.7. Attribute escape hatch (WithAttribute)

### Phase 3: DOT Generation

3.1. Basic DOT output for empty/simple graphs
3.2. Node DOT output with attributes
3.3. Edge DOT output with attributes
3.4. Graph attributes in DOT output
3.5. Default attributes in DOT output
3.6. String escaping and special characters

### Phase 4: Labels

4.1. Simple string labels
4.2. HTML Cell and Row types
4.3. HTMLTable builder
4.4. Port type and cell ports
4.5. Record Field and FieldGroup types
4.6. Record label builder
4.7. Raw HTML label escape hatch

### Phase 5: Subgraphs

5.1. Subgraph type and basic creation
5.2. Cluster subgraphs
5.3. Nested subgraphs
5.4. Subgraph DOT generation
5.5. Rank constraint convenience methods
5.6. Rank constraints via subgraph

### Phase 6: Parsing

6.1. DOT lexer/tokenizer
6.2. Basic graph parsing (digraph/graph)
6.3. Node parsing with attributes
6.4. Edge parsing with attributes
6.5. Subgraph parsing
6.6. Full parser integration

### Phase 7: Rendering

7.1. Format and Layout enums
7.2. Error types for rendering
7.3. Graphviz CLI detection
7.4. Core Render method with io.Writer
7.5. RenderToFile and RenderBytes conveniences
7.6. Layout engine selection

---

## Step Refinement (Final Granularity)

After reviewing the chunks, here's the right-sized step breakdown:

| Step | Description | Builds On |
|------|-------------|-----------|
| 1 | Project setup, go.mod, basic Graph struct | - |
| 2 | Node struct with ID, NewNode function | 1 |
| 3 | Graph.AddNode method | 1, 2 |
| 4 | Edge struct, Graph.AddEdge with implicit node add | 3 |
| 5 | Directed/Undirected/Strict graph options | 4 |
| 6 | Shape enum and NodeAttributes struct | 5 |
| 7 | NodeOption interface, WithShape, WithLabel | 6 |
| 8 | Wire NodeOptions into NewNode | 7 |
| 9 | EdgeStyle enum and EdgeAttributes struct | 8 |
| 10 | EdgeOption interface and basic options | 9 |
| 11 | Wire EdgeOptions into AddEdge | 10 |
| 12 | GraphAttributes and GraphOption | 11 |
| 13 | Default node/edge attrs on Graph | 12 |
| 14 | WithAttribute escape hatch (all types) | 13 |
| 15 | Graph.String() - basic DOT output | 14 |
| 16 | Node DOT rendering with attributes | 15 |
| 17 | Edge DOT rendering with attributes | 16 |
| 18 | Graph/default attrs in DOT output | 17 |
| 19 | String escaping in DOT output | 18 |
| 20 | HTMLCell and HTMLRow types | 19 |
| 21 | HTMLTable builder | 20 |
| 22 | Port type, cell.Port() method | 21 |
| 23 | FromPort/ToPort edge options | 22 |
| 24 | HTML label DOT output | 23 |
| 25 | Record Field and FieldGroup | 24 |
| 26 | WithRecordLabel and DOT output | 25 |
| 27 | Subgraph struct and Graph.Subgraph() | 26 |
| 28 | Cluster detection and subgraph attrs | 27 |
| 29 | Nested subgraphs | 28 |
| 30 | Subgraph DOT generation | 29 |
| 31 | SameRank, MinRank, MaxRank methods | 30 |
| 32 | Rank constraint DOT output | 31 |
| 33 | DOT lexer | 32 |
| 34 | DOT parser - graph structure | 33 |
| 35 | DOT parser - nodes and edges | 34 |
| 36 | DOT parser - subgraphs | 35 |
| 37 | Parse, ParseString, ParseFile functions | 36 |
| 38 | Format and Layout enums | 37 |
| 39 | RenderError and sentinel errors | 38 |
| 40 | Graphviz CLI detection | 39 |
| 41 | Graph.Render to io.Writer | 40 |
| 42 | RenderToFile and RenderBytes | 41 |
| 43 | WithLayout render option | 42 |

---

## Implementation Prompts

---

### Prompt 1: Project Setup and Basic Graph Struct

```text
You are implementing a Go library called "goraffe" for building and rendering Graphviz graphs.

Create the initial project structure with:

1. Initialize a Go module named "github.com/example/goraffe"

2. Create doc.go with package documentation:
   - Package goraffe provides a unified API for creating, parsing, and rendering Graphviz graphs
   - Mention it requires Graphviz installed on the system for rendering

3. Create graph.go with:
   - A Graph struct with unexported fields:
     - name string
     - directed bool
     - strict bool
   - A constructor NewGraph() that returns *Graph
   - Methods:
     - IsDirected() bool
     - IsStrict() bool
     - Name() string

4. Create graph_test.go with tests:
   - TestNewGraph_DefaultValues (directed=false, strict=false, empty name)
   - Test the getter methods return expected values

Focus on clean, idiomatic Go. Use table-driven tests where appropriate. Do not implement anything beyond what's specified - we will add features incrementally.
```

---

### Prompt 2: Node Struct and NewNode

```text
Building on the existing goraffe project, add the Node type.

1. Create node.go with:
   - A Node struct with unexported fields:
     - id string (required, immutable)
   - A constructor NewNode(id string) *Node
   - A method ID() string that returns the node's ID

2. Create node_test.go with tests:
   - TestNewNode_SetsID
   - TestNewNode_EmptyID (should still work, Graphviz allows it)
   - TestNode_ID_ReturnsCorrectValue

3. Ensure node IDs are stored exactly as provided (no transformation yet)

Keep it minimal - attributes and options will be added in later steps.
```

---

### Prompt 3: Graph.AddNode Method

```text
Building on the existing goraffe project, implement adding nodes to a graph.

1. Update graph.go:
   - Add fields to Graph:
     - nodeOrder []*Node (preserves insertion order for DOT output)
     - nodes map[string]int (maps node ID to index in nodeOrder for O(1) lookup)
   - Initialize both in NewGraph()
   - Add method AddNode(n *Node) that:
     - If node ID exists: replaces node at that index in nodeOrder
     - If new: appends to nodeOrder and stores index in nodes map
     - This preserves insertion order while allowing fast lookups
   - Add method GetNode(id string) *Node that:
     - Looks up index in nodes map
     - Returns nodeOrder[idx] if found, nil otherwise
   - Add method Nodes() []*Node that returns nodeOrder (guaranteed insertion order)

2. Update graph_test.go with tests:
   - TestGraph_AddNode_SingleNode
   - TestGraph_AddNode_MultipleNodes
   - TestGraph_AddNode_DuplicateID (verify replace-in-place behavior)
   - TestGraph_AddNode_PreservesOrder (verify nodes returned in insertion order)
   - TestGraph_GetNode_Exists
   - TestGraph_GetNode_NotFound
   - TestGraph_Nodes_ReturnsAllNodes
   - TestGraph_Nodes_ReturnsInInsertionOrder

Make sure the Graph properly owns the nodes after AddNode is called.
Note: This design prioritizes insertion order preservation (critical for DOT output)
and fast lookups, at the expense of expensive node deletion (not needed in this API).
```

---

### Prompt 4: Edge Struct and Graph.AddEdge

```text
Building on the existing goraffe project, implement edges.

1. Create edge.go with:
   - An Edge struct with unexported fields:
     - from *Node
     - to *Node
   - Methods:
     - From() *Node
     - To() *Node

2. Update graph.go:
   - Add an edges field to Graph: edges []*Edge
   - Initialize as empty slice in NewGraph()
   - Add method AddEdge(from, to *Node) *Edge that:
     - Creates a new Edge
     - Adds both nodes to the graph if not already present (implicit add)
     - Appends the edge to the edges slice
     - Returns the created edge
   - Add method Edges() []*Edge

3. Create edge_test.go with tests:
   - TestEdge_FromTo_ReturnsCorrectNodes
   - TestGraph_AddEdge_BothNodesExist
   - TestGraph_AddEdge_ImplicitNodeAdd (nodes not previously added)
   - TestGraph_AddEdge_PartialImplicitAdd (one node exists, one doesn't)
   - TestGraph_Edges_ReturnsAllEdges

4. Update graph_test.go if needed for integration

Note: Edge attributes and options come later. Keep edges simple for now.
```

---

### Prompt 5: Directed/Undirected/Strict Graph Options

```text
Building on the existing goraffe project, add graph type configuration using functional options.

1. Create options.go with:
   - A GraphOption interface with unexported method: applyGraph(*Graph)
   - A graphOptionFunc type that implements GraphOption
   - Public option variables/functions:
     - Directed GraphOption - sets graph.directed = true
     - Undirected GraphOption - sets graph.directed = false (this is default, but explicit)
     - Strict GraphOption - sets graph.strict = true

2. Update graph.go:
   - Modify NewGraph to accept variadic GraphOption: NewGraph(opts ...GraphOption) *Graph
   - Apply all options after setting defaults

3. Update graph_test.go:
   - TestNewGraph_Directed
   - TestNewGraph_Undirected
   - TestNewGraph_Strict
   - TestNewGraph_DirectedAndStrict
   - TestNewGraph_MultipleOptions_LastWins (if Directed then Undirected passed)

Make sure existing tests still pass (they may need updating to use NewGraph() with no args).
```

---

### Prompt 6: Shape Enum and NodeAttributes Struct

```text
Building on the existing goraffe project, add the attribute system foundation for nodes.

1. Create attributes.go with:
   - A Shape type (string-based enum)
   - Shape constants:
     - ShapeBox Shape = "box"
     - ShapeCircle Shape = "circle"
     - ShapeEllipse Shape = "ellipse"
     - ShapeDiamond Shape = "diamond"
     - ShapeRecord Shape = "record"
     - ShapePlaintext Shape = "plaintext"
   - A NodeAttributes struct with exported fields:
     - Label string
     - Shape Shape
     - Color string
     - FillColor string
     - FontName string
     - FontSize float64
     - custom map[string]string (unexported, for escape hatch later)
   - A method on NodeAttributes: Custom() map[string]string (returns copy)

2. Update node.go:
   - Add an attrs field to Node: attrs *NodeAttributes
   - Initialize with empty NodeAttributes in NewNode
   - Add method Attrs() *NodeAttributes

3. Create attributes_test.go:
   - TestShapeConstants_Values (verify string values)
   - TestNodeAttributes_ZeroValue (all fields empty/zero)
   - TestNodeAttributes_Custom_ReturnsCopy

4. Update node_test.go:
   - TestNode_Attrs_ReturnsAttributes
```

---

### Prompt 7: NodeOption Interface and Basic Functional Options

```text
Building on the existing goraffe project, implement the NodeOption pattern.

1. Update options.go (or create node_options.go for organization):
   - A NodeOption interface with unexported method: applyNode(*NodeAttributes)
   - A nodeOptionFunc type that implements NodeOption
   - Helper function to create options: newNodeOption(fn func(*NodeAttributes)) NodeOption
   - Implement these options:
     - WithShape(s Shape) NodeOption
     - WithLabel(l string) NodeOption
     - WithColor(c string) NodeOption
     - WithFillColor(c string) NodeOption
     - WithFontName(f string) NodeOption
     - WithFontSize(s float64) NodeOption

2. Create node_options_test.go:
   - TestWithShape_SetsShape
   - TestWithLabel_SetsLabel
   - TestWithColor_SetsColor
   - TestWithFillColor_SetsFillColor
   - TestWithFontName_SetsFontName
   - TestWithFontSize_SetsFontSize
   - Test each option independently by applying to a NodeAttributes

Note: Don't wire into NewNode yet - that's the next step.
```

---

### Prompt 8: Wire NodeOptions into NewNode

```text
Building on the existing goraffe project, wire NodeOptions into node creation.

1. Update node.go:
   - Modify NewNode signature: NewNode(id string, opts ...NodeOption) *Node
   - Apply all options to the node's attributes after creation

2. Make NodeAttributes implement NodeOption:
   - Add method applyNode(*NodeAttributes) to NodeAttributes
   - It should merge non-zero fields from self into target
   - This allows passing a reusable NodeAttributes struct as an option

3. Update node_test.go:
   - TestNewNode_WithOptions
   - TestNewNode_WithMultipleOptions
   - TestNewNode_WithNodeAttributesStruct (reusable attrs)
   - TestNewNode_OptionsAppliedInOrder (later options override earlier)

4. Update any existing tests that call NewNode to ensure they still pass

5. Revisit node_options_test.go from Prompt 7:
   - Consider whether these tests are still needed now that we have public API tests
   - The tests in node_options_test.go test the private applyNode method directly
   - Now that NewNode accepts options, the tests in node_test.go will exercise the same code paths
   - Decision: Keep them for now if they provide value in isolating failures, or delete if they're redundant
   - These tests made sense for TDD (Red phase), but may not be needed long-term

Example usage after this step:
  n := NewNode("A", WithShape(ShapeBox), WithLabel("Node A"))
  
  commonAttrs := NodeAttributes{Shape: ShapeBox, FontName: "Arial"}
  n2 := NewNode("B", commonAttrs, WithLabel("Node B"))
```

---

### Prompt 9: EdgeStyle Enum and EdgeAttributes Struct

```text
Building on the existing goraffe project, add the attribute system for edges.

1. Update attributes.go:
   - An EdgeStyle type (string-based enum)
   - EdgeStyle constants:
     - EdgeStyleSolid EdgeStyle = "solid"
     - EdgeStyleDashed EdgeStyle = "dashed"
     - EdgeStyleDotted EdgeStyle = "dotted"
     - EdgeStyleBold EdgeStyle = "bold"
   - An ArrowType type (string-based enum)
   - ArrowType constants:
     - ArrowNormal ArrowType = "normal"
     - ArrowDot ArrowType = "dot"
     - ArrowNone ArrowType = "none"
     - ArrowVee ArrowType = "vee"
   - An EdgeAttributes struct with exported fields:
     - Label string
     - Color string
     - Style EdgeStyle
     - ArrowHead ArrowType
     - ArrowTail ArrowType
     - Weight float64
     - custom map[string]string (unexported)
   - Method Custom() map[string]string on EdgeAttributes

2. Update edge.go:
   - Add attrs field to Edge: attrs *EdgeAttributes
   - Initialize with empty EdgeAttributes
   - Add method Attrs() *EdgeAttributes

3. Update attributes_test.go:
   - TestEdgeStyleConstants_Values
   - TestArrowTypeConstants_Values
   - TestEdgeAttributes_ZeroValue

4. Update edge_test.go:
   - TestEdge_Attrs_ReturnsAttributes
```

---

### Prompt 10: EdgeOption Interface and Basic Options

```text
Building on the existing goraffe project, implement EdgeOption pattern.

1. Update options.go (or create edge_options.go):
   - An EdgeOption interface with unexported method: applyEdge(*EdgeAttributes)
   - An edgeOptionFunc type that implements EdgeOption
   - Helper: newEdgeOption(fn func(*EdgeAttributes)) EdgeOption
   - Implement these options:
     - WithEdgeLabel(l string) EdgeOption
     - WithEdgeColor(c string) EdgeOption
     - WithEdgeStyle(s EdgeStyle) EdgeOption
     - WithArrowHead(a ArrowType) EdgeOption
     - WithArrowTail(a ArrowType) EdgeOption
     - WithWeight(w float64) EdgeOption

2. Make EdgeAttributes implement EdgeOption:
   - Add applyEdge method that merges non-zero fields

3. Create edge_options_test.go:
   - Test each option independently
   - TestEdgeAttributes_AsOption (using struct as option)

Note: Don't wire into AddEdge yet - next step.
```

---

### Prompt 11: Wire EdgeOptions into AddEdge

```text
Building on the existing goraffe project, wire EdgeOptions into edge creation.

1. Update graph.go:
   - Modify AddEdge signature: AddEdge(from, to *Node, opts ...EdgeOption) *Edge
   - Apply all options to the edge's attributes

2. Update edge_test.go and graph_test.go:
   - TestGraph_AddEdge_WithOptions
   - TestGraph_AddEdge_WithMultipleOptions
   - TestGraph_AddEdge_WithEdgeAttributesStruct
   - TestGraph_AddEdge_OptionsAppliedInOrder

3. Update any existing AddEdge calls in tests to pass with new signature

Example usage after this step:
  g.AddEdge(n1, n2, WithEdgeLabel("connects"), WithEdgeStyle(EdgeStyleDashed))
  
  commonEdge := EdgeAttributes{Style: EdgeStyleDashed, Color: "gray"}
  g.AddEdge(n3, n4, commonEdge)
```

---

### Prompt 12: GraphAttributes and GraphOption

```text
Building on the existing goraffe project, add graph-level attributes.

1. Update attributes.go:
   - A RankDir type (string-based enum)
   - RankDir constants:
     - RankDirTB RankDir = "TB"
     - RankDirBT RankDir = "BT"
     - RankDirLR RankDir = "LR"
     - RankDirRL RankDir = "RL"
   - A SplineType type (string-based enum)
   - SplineType constants:
     - SplineTrue SplineType = "true"
     - SplineFalse SplineType = "false"
     - SplineOrtho SplineType = "ortho"
     - SplinePolyline SplineType = "polyline"
     - SplineCurved SplineType = "curved"
   - A GraphAttributes struct:
     - Label string
     - RankDir RankDir
     - BgColor string
     - FontName string
     - FontSize float64
     - Splines SplineType
     - NodeSep float64
     - RankSep float64
     - Compound bool
     - custom map[string]string

2. Update graph.go:
   - Add attrs field: attrs *GraphAttributes
   - Initialize in NewGraph
   - Add Attrs() *GraphAttributes method

3. Update options.go with GraphOption implementations:
   - WithGraphLabel(l string) GraphOption
   - WithRankDir(d RankDir) GraphOption
   - WithBgColor(c string) GraphOption
   - WithGraphFontName(f string) GraphOption
   - WithGraphFontSize(s float64) GraphOption
   - WithSplines(s SplineType) GraphOption
   - WithNodeSep(n float64) GraphOption
   - WithRankSep(r float64) GraphOption
   - WithCompound(b bool) GraphOption

4. Update tests:
   - Test all new enums and their values
   - Test each GraphOption
   - TestNewGraph_WithGraphOptions
```

---

### Prompt 13: Default Node/Edge Attributes on Graph

```text
Building on the existing goraffe project, add default attributes for nodes and edges.

1. Update graph.go:
   - Add fields to Graph:
     - defaultNodeAttrs *NodeAttributes
     - defaultEdgeAttrs *EdgeAttributes
   - Initialize both in NewGraph
   - Add methods:
     - DefaultNodeAttrs() *NodeAttributes
     - DefaultEdgeAttrs() *EdgeAttributes

2. Add GraphOption implementations in options.go:
   - WithDefaultNodeAttrs(opts ...NodeOption) GraphOption
     - Creates a NodeAttributes, applies all opts, stores on graph
   - WithDefaultEdgeAttrs(opts ...EdgeOption) GraphOption
     - Creates an EdgeAttributes, applies all opts, stores on graph

3. Update tests:
   - TestGraph_WithDefaultNodeAttrs
   - TestGraph_WithDefaultEdgeAttrs
   - TestGraph_DefaultAttrs_AppliesMultipleOptions

Example usage:
  g := NewGraph(
      Directed,
      WithDefaultNodeAttrs(WithShape(ShapeBox), WithFontName("Arial")),
      WithDefaultEdgeAttrs(WithEdgeColor("gray")),
  )

Note: These defaults will be used in DOT generation later, not applied to individual nodes/edges at creation time.
```

---

### Prompt 14: WithAttribute Escape Hatch

```text
Building on the existing goraffe project, add the escape hatch for arbitrary attributes.

1. Update NodeAttributes, EdgeAttributes, and GraphAttributes:
   - Ensure each has a SetCustom(key, value string) method
   - Ensure Custom() returns a copy of the custom map

2. Add escape hatch options:
   - WithNodeAttribute(key, value string) NodeOption
     - Stores in NodeAttributes.custom
   - WithEdgeAttribute(key, value string) EdgeOption
     - Stores in EdgeAttributes.custom
   - WithGraphAttribute(key, value string) GraphOption
     - Stores in GraphAttributes.custom

3. Create tests:
   - TestWithNodeAttribute_SetsCustom
   - TestWithEdgeAttribute_SetsCustom
   - TestWithGraphAttribute_SetsCustom
   - TestCustomAttributes_DoNotOverrideTyped (custom "shape" vs WithShape)
   - TestCustomAttributes_MultipleCalls_Accumulate

Example usage:
  n := NewNode("A", 
      WithShape(ShapeBox),
      WithNodeAttribute("peripheries", "2"),
  )
```

---

### Prompt 15: Graph.String() - Basic DOT Output

```text
Building on the existing goraffe project, implement basic DOT generation.

1. Create dot.go with:
   - Method (g *Graph) String() string
   - Method (g *Graph) WriteDOT(w io.Writer) error
   - Internal helper to generate DOT

2. Initial implementation should handle:
   - "digraph" vs "graph" based on g.directed
   - "strict" prefix if g.strict
   - Graph name if set
   - Empty body for now (nodes/edges come next)

3. Expected output formats:
   - NewGraph(Directed) → "digraph {\n}\n"
   - NewGraph(Undirected) → "graph {\n}\n"
   - NewGraph(Directed, Strict) → "strict digraph {\n}\n"
   - Graph with name "G" → "digraph G {\n}\n"

4. Create dot_test.go:
   - TestGraph_String_EmptyDirected
   - TestGraph_String_EmptyUndirected
   - TestGraph_String_Strict
   - TestGraph_String_WithName
   - TestGraph_WriteDOT_WritesToWriter

Use a strings.Builder internally for efficiency.
```

---

### Prompt 16: Node DOT Rendering with Attributes

```text
Building on the existing goraffe project, add node output to DOT generation.

1. Update dot.go:
   - Add internal method to render a single node to DOT
   - Update String()/WriteDOT() to include all nodes
   - Handle node attributes:
     - Only output non-zero/non-empty attributes
     - Format: nodeID [attr1="val1", attr2="val2"];
     - Node with no attributes: just nodeID;

2. Attribute rendering:
   - Label → label="value"
   - Shape → shape="value"
   - Color → color="value"
   - FillColor → fillcolor="value"
   - FontName → fontname="value"
   - FontSize → fontsize="value" (only if > 0)
   - Custom attributes: key="value"

3. Create/update dot_test.go:
   - TestDOT_SingleNode_NoAttributes
   - TestDOT_SingleNode_WithLabel
   - TestDOT_SingleNode_WithShape
   - TestDOT_SingleNode_MultipleAttributes
   - TestDOT_SingleNode_CustomAttribute
   - TestDOT_MultipleNodes

4. Node IDs that need quoting:
   - For now, always quote node IDs to be safe
   - Format: "nodeID" [attrs];

Example output:
  digraph {
      "A" [shape="box", label="Node A"];
      "B";
  }
```

---

### Prompt 17: Edge DOT Rendering with Attributes

```text
Building on the existing goraffe project, add edge output to DOT generation.

1. Update dot.go:
   - Add internal method to render a single edge to DOT
   - Update String()/WriteDOT() to include all edges after nodes
   - Handle directed vs undirected:
     - Directed: "A" -> "B"
     - Undirected: "A" -- "B"
   - Handle edge attributes (same format as nodes)

2. Attribute rendering for edges:
   - Label → label="value"
   - Color → color="value"
   - Style → style="value"
   - ArrowHead → arrowhead="value"
   - ArrowTail → arrowtail="value"
   - Weight → weight="value" (only if > 0)
   - Custom attributes

3. Update dot_test.go:
   - TestDOT_SingleEdge_NoAttributes
   - TestDOT_SingleEdge_Directed
   - TestDOT_SingleEdge_Undirected
   - TestDOT_SingleEdge_WithLabel
   - TestDOT_SingleEdge_MultipleAttributes
   - TestDOT_MultipleEdges
   - TestDOT_CompleteGraph (nodes + edges together)

Example output:
  digraph {
      "A" [shape="box"];
      "B";
      "A" -> "B" [label="connects", style="dashed"];
  }
```

---

### Prompt 18: Graph and Default Attributes in DOT Output

```text
Building on the existing goraffe project, add graph attributes and defaults to DOT output.

1. Update dot.go to output graph attributes:
   - After opening brace, before nodes
   - Format: attrname="value";
   - Handle all GraphAttributes fields

2. Update dot.go to output default node/edge attributes:
   - node [attr1="val1", attr2="val2"];
   - edge [attr1="val1", attr2="val2"];
   - Only output if there are non-zero defaults

3. Output order in DOT:
   1. Graph declaration (strict? digraph/graph name {)
   2. Graph attributes
   3. Default node attributes (node [...];)
   4. Default edge attributes (edge [...];)
   5. Nodes
   6. Edges
   7. Closing brace

4. Update dot_test.go:
   - TestDOT_GraphAttributes_RankDir
   - TestDOT_GraphAttributes_Label
   - TestDOT_GraphAttributes_Multiple
   - TestDOT_DefaultNodeAttrs
   - TestDOT_DefaultEdgeAttrs
   - TestDOT_FullGraph_WithAllSections

Example output:
  digraph G {
      rankdir="LR";
      label="My Graph";
      node [shape="box", fontname="Arial"];
      edge [color="gray"];
      "A";
      "B";
      "A" -> "B";
  }
```

---

### Prompt 19: String Escaping in DOT Output

```text
Building on the existing goraffe project, properly handle string escaping in DOT output.

1. Create or update a helper function for DOT string escaping:
   - Escape backslashes: \ → \\
   - Escape double quotes: " → \"
   - Escape newlines: \n → \n (literal backslash-n in output)
   - Handle other special characters as needed

2. Create a helper to determine if a string needs quoting:
   - Node IDs: quote if contains spaces, special chars, or starts with digit
   - Attribute values: always quote for safety

3. Update all DOT output to use proper escaping:
   - Node IDs
   - Attribute values (labels especially)
   - Graph names

4. Update dot_test.go:
   - TestDOT_NodeID_WithSpaces
   - TestDOT_NodeID_WithSpecialChars
   - TestDOT_Label_WithQuotes
   - TestDOT_Label_WithNewlines
   - TestDOT_Label_WithBackslashes
   - TestDOT_ComplexStrings

Example:
  Node with label `Say "Hello"` should output:
  "mynode" [label="Say \"Hello\""];
```

---

### Prompt 20: HTMLCell and HTMLRow Types

```text
Building on the existing goraffe project, add HTML label building blocks.

1. Create labels.go with:
   - HTMLCell struct (unexported fields):
     - content string
     - port string
     - bold bool
     - italic bool
     - underline bool
     - colSpan int
     - rowSpan int
     - bgColor string
     - align string
   - Constructor: Cell(content string) *HTMLCell
   - Chainable methods on *HTMLCell:
     - Port(id string) *HTMLCell
     - Bold() *HTMLCell
     - Italic() *HTMLCell
     - Underline() *HTMLCell
     - ColSpan(n int) *HTMLCell
     - RowSpan(n int) *HTMLCell
     - BgColor(color string) *HTMLCell
     - Align(a string) *HTMLCell

2. HTMLRow struct:
   - cells []*HTMLCell
   - Constructor: Row(cells ...*HTMLCell) *HTMLRow
   - Method: Cells() []*HTMLCell

3. Create labels_test.go:
   - TestCell_Content
   - TestCell_Chaining (Port, Bold, etc.)
   - TestCell_AllOptions
   - TestRow_ContainsCells
   - TestRow_MultipleCells

Note: Don't render to string yet - that comes with HTMLTable.
```

---

### Prompt 21: HTMLTable Builder

```text
Building on the existing goraffe project, add HTMLTable and HTML label rendering.

1. Update labels.go:
   - HTMLLabel struct:
     - rows []*HTMLRow
     - border int
     - cellBorder int
     - cellSpacing int
     - cellPadding int
     - bgColor string
   - Constructor: HTMLTable(rows ...*HTMLRow) *HTMLLabel
   - Chainable methods:
     - Border(n int) *HTMLLabel
     - CellBorder(n int) *HTMLLabel
     - CellSpacing(n int) *HTMLLabel
     - CellPadding(n int) *HTMLLabel
     - BgColor(color string) *HTMLLabel
   - Method: String() string - renders to HTML string

2. HTML rendering:
   - Output wrapped in < > (Graphviz HTML label delimiters)
   - TABLE element with attributes
   - TR for each row
   - TD for each cell with attributes
   - Cell content wrapped in formatting tags (B, I, U)
   - PORT attribute on TD if port is set

3. Update labels_test.go:
   - TestHTMLTable_SimpleTable
   - TestHTMLTable_WithTableAttributes
   - TestHTMLTable_CellWithPort
   - TestHTMLTable_CellWithFormatting
   - TestHTMLTable_CellWithSpan
   - TestHTMLTable_ComplexTable

Example output:
  <<TABLE BORDER="0"><TR><TD PORT="p1"><B>Header</B></TD></TR></TABLE>>
```

---

### Prompt 22: Port Type and Cell Port Reference

```text
Building on the existing goraffe project, add the Port type for type-safe port references.

1. Create port.go:
   - Port struct (unexported fields):
     - id string
     - nodeID string (will be set when label is assigned to node)
   - Method: ID() string

2. Update labels.go / HTMLCell:
   - Add internal portRef *Port field
   - Update Port(id string) method to create and store a *Port
   - Add GetPort() *Port method to retrieve the port reference

3. Add a mechanism to associate ports with nodes:
   - When a node is created with an HTML label, the label's ports should know their node
   - Add internal method on HTMLLabel to set node context
   - Update Port.nodeID when label is attached

4. Create port_test.go:
   - TestPort_ID
   - TestCell_GetPort_ReturnsPort
   - TestCell_GetPort_NilIfNoPort

5. Update labels_test.go:
   - TestHTMLLabel_PortsKnowNodeID (after label attached to node)

Note: We'll wire this into edge connections in the next step.
```

---

### Prompt 23: FromPort/ToPort Edge Options

```text
Building on the existing goraffe project, add port-based edge connections.

1. Update edge.go / EdgeAttributes:
   - Add fields:
     - fromPort *Port
     - toPort *Port
   - Add methods:
     - FromPort() *Port
     - ToPort() *Port

2. Add EdgeOption implementations:
   - FromPort(p *Port) EdgeOption
   - ToPort(p *Port) EdgeOption

3. Update DOT generation in dot.go:
   - If edge has fromPort: "nodeID":"portID" -> ...
   - If edge has toPort: ... -> "nodeID":"portID"
   - Format: "A":"p1" -> "B":"p2"

4. Create/update tests:
   - TestFromPort_SetsPort
   - TestToPort_SetsPort
   - TestDOT_Edge_WithFromPort
   - TestDOT_Edge_WithToPort
   - TestDOT_Edge_WithBothPorts

Example output:
  "A":"out" -> "B":"in" [label="data"];
```

---

### Prompt 24: HTML Label DOT Output Integration

```text
Building on the existing goraffe project, integrate HTML labels into node DOT output.

1. Update NodeAttributes:
   - Add htmlLabel *HTMLLabel field
   - Add rawHTMLLabel string field (for escape hatch)

2. Add NodeOption implementations:
   - WithHTMLLabel(label *HTMLLabel) NodeOption
   - WithRawHTMLLabel(html string) NodeOption

3. Update DOT generation:
   - If node has htmlLabel, output label=<...HTML...> (no quotes, angle brackets)
   - If node has rawHTMLLabel, output label=<...raw...>
   - HTML labels take precedence over regular Label

4. Wire port node association:
   - When WithHTMLLabel is applied, set the node context on the label's ports

5. Update tests:
   - TestWithHTMLLabel_SetsLabel
   - TestWithRawHTMLLabel_SetsLabel
   - TestDOT_Node_WithHTMLLabel
   - TestDOT_Node_WithHTMLLabel_Ports
   - TestDOT_Node_WithRawHTMLLabel
   - TestDOT_HTMLLabel_NotDoubleEscaped

Example output:
  "A" [label=<<TABLE><TR><TD>Cell</TD></TR></TABLE>>];
```

---

### Prompt 25: Record Field and FieldGroup

```text
Building on the existing goraffe project, add record label building blocks.

1. Update labels.go:
   - RecordField struct:
     - content string
     - port string
     - portRef *Port
   - Constructor: Field(content string) *RecordField
   - Methods:
     - Port(id string) *RecordField (chainable, creates Port)
     - GetPort() *Port

   - RecordGroup struct (for nested grouping):
     - elements []RecordElement
   - Constructor: FieldGroup(elements ...RecordElement) *RecordGroup

   - RecordElement interface:
     - recordElement() marker method
   - Both RecordField and RecordGroup implement RecordElement

   - RecordLabel struct:
     - elements []RecordElement
   - Constructor: RecordLabel(elements ...RecordElement) *RecordLabel
   - Method: String() string (renders to record label syntax)

2. Record label rendering:
   - Fields separated by |
   - Groups wrapped in { }
   - Ports: <portID> content
   - Escape special chars: |, {, }, <, >

3. Create record_labels_test.go:
   - TestRecordField_Content
   - TestRecordField_WithPort
   - TestRecordGroup_Nesting
   - TestRecordLabel_SimpleFields
   - TestRecordLabel_WithGroup
   - TestRecordLabel_Escaping

Example output:
  Field("a"), Field("b") → "a | b"
  Field("a").Port("p1") → "<p1> a"
  FieldGroup(Field("x"), Field("y")) → "{ x | y }"
```

---

### Prompt 26: WithRecordLabel and DOT Output

```text
Building on the existing goraffe project, integrate record labels into nodes.

1. Update NodeAttributes:
   - Add recordLabel *RecordLabel field

2. Add NodeOption:
   - WithRecordLabel(elements ...RecordElement) NodeOption
     - Creates RecordLabel from elements
     - Note: Should also ensure shape is set to Record or MRecord

3. Update DOT generation:
   - If node has recordLabel, output label="escaped record syntax"
   - Record labels ARE quoted (unlike HTML labels)
   - Ensure shape is output as "record" if using record label

4. Wire port association for record labels (similar to HTML labels)

5. Update tests:
   - TestWithRecordLabel_SetsLabel
   - TestWithRecordLabel_SetsShape
   - TestDOT_Node_WithRecordLabel_Simple
   - TestDOT_Node_WithRecordLabel_WithPorts
   - TestDOT_Node_WithRecordLabel_Nested
   - TestDOT_Edge_ToRecordPort

Example output:
  "A" [shape="record", label="<in> Input | <out> Output"];
  "B" [shape="record", label="{ Top | Bottom }"];
```

---

### Prompt 27: Subgraph Struct and Graph.Subgraph()

```text
Building on the existing goraffe project, add basic subgraph support.

1. Create subgraph.go:
   - Subgraph struct (similar to Graph but simpler):
     - name string
     - nodes map[string]*Node
     - edges []*Edge
     - isCluster bool (determined by name prefix "cluster")
     - parent *Graph (back-reference)
   - Methods:
     - Name() string
     - IsCluster() bool
     - AddNode(n *Node) - also adds to parent graph
     - Nodes() []*Node
     - AddEdge(from, to *Node, opts ...EdgeOption) *Edge - delegates to parent

2. Update graph.go:
   - Add subgraphs field: subgraphs []*Subgraph
   - Add method:
     - Subgraph(name string, fn func(*Subgraph)) *Subgraph
       - Creates subgraph
       - Calls fn with the subgraph
       - Adds to graph's subgraph list
       - Returns the subgraph
   - Add method: Subgraphs() []*Subgraph

3. Create subgraph_test.go:
   - TestSubgraph_Name
   - TestSubgraph_IsCluster_True (name starts with "cluster")
   - TestSubgraph_IsCluster_False
   - TestSubgraph_AddNode
   - TestSubgraph_AddNode_AlsoAddsToParent
   - TestGraph_Subgraph_CallsFunction
   - TestGraph_Subgraph_ReturnsSubgraph
   - TestGraph_Subgraphs_ReturnsAll
```

---

### Prompt 28: Cluster Detection and Subgraph Attributes

```text
Building on the existing goraffe project, add subgraph attributes.

1. Update attributes.go:
   - SubgraphAttributes struct:
     - Label string
     - Style string
     - Color string
     - FillColor string
     - FontName string
     - FontSize float64
     - custom map[string]string

2. Update subgraph.go:
   - Add attrs field: attrs *SubgraphAttributes
   - Add Attrs() *SubgraphAttributes method
   - Add setter methods (direct setters, not functional options for simplicity):
     - SetLabel(l string)
     - SetStyle(s string)
     - SetColor(c string)
     - SetFillColor(c string)
     - SetAttribute(key, value string) - for custom

3. Cluster-specific behavior:
   - Clusters can have bgcolor/fillcolor (regular subgraphs typically don't render these)
   - Document this difference

4. Update tests:
   - TestSubgraph_SetLabel
   - TestSubgraph_SetStyle
   - TestSubgraph_SetAttribute
   - TestSubgraph_Attrs_ReturnsAttributes
   - TestSubgraph_Cluster_CanHaveStyle
```

---

### Prompt 29: Nested Subgraphs

```text
Building on the existing goraffe project, add support for nested subgraphs.

1. Update subgraph.go:
   - Add subgraphs field to Subgraph: subgraphs []*Subgraph
   - Add method:
     - Subgraph(name string, fn func(*Subgraph)) *Subgraph
       - Creates nested subgraph
       - Sets parent appropriately (should reference root graph for node tracking)
       - Calls fn
       - Returns subgraph
   - Add method: Subgraphs() []*Subgraph

2. Ensure node tracking works correctly:
   - Nodes added to nested subgraphs should still appear in root graph's node map
   - Subgraph maintains its own list for DOT output purposes

3. Update tests:
   - TestSubgraph_NestedSubgraph
   - TestSubgraph_NestedSubgraph_NodesInRoot
   - TestSubgraph_DeeplyNested (3 levels)
   - TestSubgraph_NestedCluster

Example usage:
  g.Subgraph("cluster_outer", func(outer *Subgraph) {
      outer.SetLabel("Outer")
      outer.Subgraph("cluster_inner", func(inner *Subgraph) {
          inner.SetLabel("Inner")
          inner.AddNode(n1)
      })
  })
```

---

### Prompt 30: Subgraph DOT Generation

```text
Building on the existing goraffe project, add subgraph output to DOT generation.

1. Update dot.go:
   - Add internal method to render a subgraph to DOT
   - Subgraph format:
     - subgraph name { ... } or subgraph cluster_name { ... }
     - Include subgraph attributes
     - Include nodes that belong to this subgraph
     - Include nested subgraphs (recursive)
   - Update main DOT output to include subgraphs after graph attrs/defaults, before loose nodes

2. Output order:
   1. Graph declaration
   2. Graph attributes
   3. Default node/edge attributes
   4. Subgraphs (each contains their nodes)
   5. Nodes not in any subgraph
   6. All edges (edges are always at graph level in our output)
   7. Closing brace

3. Handle empty subgraph names:
   - Anonymous subgraph: subgraph { ... }

4. Update dot_test.go:
   - TestDOT_Subgraph_Simple
   - TestDOT_Subgraph_WithAttributes
   - TestDOT_Subgraph_Cluster
   - TestDOT_Subgraph_Nested
   - TestDOT_Subgraph_Anonymous
   - TestDOT_Graph_WithSubgraphsAndLooseNodes

Example output:
  digraph {
      subgraph cluster_db {
          label="Database";
          "db1";
          "db2";
      }
      "web";
      "web" -> "db1";
  }
```

---

### Prompt 31: SameRank, MinRank, MaxRank Convenience Methods

```text
Building on the existing goraffe project, add rank constraint convenience methods.

1. Update attributes.go:
   - Add Rank type (string enum):
     - RankSame Rank = "same"
     - RankMin Rank = "min"
     - RankMax Rank = "max"
     - RankSource Rank = "source"
     - RankSink Rank = "sink"

2. Update subgraph.go:
   - Add rank field to Subgraph: rank Rank
   - Add SetRank(r Rank) method
   - Add Rank() Rank getter

3. Update graph.go:
   - Add internal helper to create anonymous rank subgraph
   - Add convenience methods:
     - SameRank(nodes ...*Node) - creates anonymous subgraph with rank=same
     - MinRank(nodes ...*Node)
     - MaxRank(nodes ...*Node)
     - SourceRank(nodes ...*Node)
     - SinkRank(nodes ...*Node)

4. Update tests:
   - TestGraph_SameRank
   - TestGraph_MinRank
   - TestGraph_MaxRank
   - TestGraph_SourceRank
   - TestGraph_SinkRank
   - TestSubgraph_SetRank

Example usage:
  g.SameRank(n1, n2, n3)
  
  // Equivalent to:
  g.Subgraph("", func(s *Subgraph) {
      s.SetRank(RankSame)
      s.AddNode(n1, n2, n3)
  })
```

---

### Prompt 32: Rank Constraint DOT Output

```text
Building on the existing goraffe project, add rank constraint output to DOT.

1. Update subgraph DOT generation:
   - If subgraph has rank set, output: rank="value";
   - This goes after other subgraph attributes

2. Rank subgraphs from convenience methods:
   - These are anonymous subgraphs (no name)
   - Should only contain rank attribute and node references

3. Update dot_test.go:
   - TestDOT_Subgraph_WithRank
   - TestDOT_SameRank_CreatesSubgraph
   - TestDOT_MinRank_Output
   - TestDOT_MaxRank_Output
   - TestDOT_MultipleRankConstraints
   - TestDOT_ComplexGraph_WithRanks

Example output:
  digraph {
      { rank="same"; "A"; "B"; "C"; }
      { rank="min"; "start"; }
      "start" -> "A";
      "A" -> "B";
  }

Note: Rank subgraphs often use { } shorthand instead of subgraph { } - both are valid DOT.
```

---

### Prompt 33: DOT Lexer

```text
Building on the existing goraffe project, start implementing DOT parsing with a lexer.

1. Create parse.go (or lexer.go):
   - Token type with constants:
     - TokenEOF
     - TokenIdent (identifiers, including keywords detected later)
     - TokenString (quoted strings)
     - TokenNumber
     - TokenLBrace, TokenRBrace ({, })
     - TokenLBracket, TokenRBracket ([, ])
     - TokenLParen, TokenRParen (for subgraph grouping)
     - TokenSemi (;)
     - TokenComma (,)
     - TokenColon (:)
     - TokenEqual (=)
     - TokenArrow (-> or --)
     - TokenHTML (HTML string <>)
   
   - Token struct:
     - Type TokenType
     - Value string
     - Line, Col int (for error messages)

   - Lexer struct:
     - input string
     - pos, line, col int
     - Methods: Next() Token, Peek() Token

2. Lexer behavior:
   - Skip whitespace and comments (// and /* */)
   - Recognize keywords as TokenIdent: graph, digraph, subgraph, node, edge, strict
   - Handle quoted strings with escape sequences
   - Handle HTML strings: < ... > (balanced angle brackets)
   - Handle -> and -- as single tokens

3. Create lexer_test.go:
   - TestLexer_SimpleTokens
   - TestLexer_Identifiers
   - TestLexer_QuotedStrings
   - TestLexer_HTMLStrings
   - TestLexer_Arrows
   - TestLexer_Comments
   - TestLexer_CompleteGraph
```

---

### Prompt 34: DOT Parser - Graph Structure

```text
Building on the existing goraffe project, implement basic DOT graph parsing.

1. Update parse.go:
   - Parser struct:
     - lexer *Lexer
     - current Token
     - Methods: advance(), expect(TokenType), match(TokenType) bool

   - Top-level parse function:
     - parseGraph() (*Graph, error)
     - Handles: [strict] (graph|digraph) [name] { ... }

   - Internal helpers:
     - parseStmtList() - parse statements until }
     - parseStmt() - dispatch to appropriate parser

2. Initial statement parsing (skeleton):
   - Recognize but don't fully implement:
     - Node statements
     - Edge statements
     - Attribute statements (graph/node/edge [...])
     - Subgraph statements
   - For now, just identify statement types and skip

3. Create parser_test.go:
   - TestParse_EmptyDigraph
   - TestParse_EmptyGraph
   - TestParse_StrictGraph
   - TestParse_NamedGraph
   - TestParse_InvalidSyntax_Error

Example parsing:
  "digraph G {}" → Graph{directed: true, name: "G"}
```

---

### Prompt 35: DOT Parser - Nodes and Edges

```text
Building on the existing goraffe project, implement node and edge parsing.

1. Update parse.go:
   - parseNodeStmt():
     - Parse: nodeID [attributes]
     - Create Node with parsed attributes
     - Add to graph
   
   - parseEdgeStmt():
     - Parse: nodeID (->|--) nodeID (->|--) nodeID ... [attributes]
     - Handle edge chains: A -> B -> C creates edges A->B and B->C
     - Create edges with parsed attributes
   
   - parseAttrList():
     - Parse: [attr=value, attr=value, ...]
     - Return map[string]string
   
   - parseID():
     - Handle: identifier, quoted string, number, HTML string

2. Attribute mapping:
   - Map parsed attributes to typed fields where known
   - Store unknown attributes in custom map

3. Update parser_test.go:
   - TestParse_SingleNode
   - TestParse_NodeWithAttributes
   - TestParse_SingleEdge
   - TestParse_EdgeWithAttributes
   - TestParse_EdgeChain
   - TestParse_MixedNodesAndEdges

Example:
  "digraph { A [shape=box]; A -> B [label="edge"]; }"
  → Graph with nodes A, B and edge A->B
```

---

### Prompt 36: DOT Parser - Subgraphs

```text
Building on the existing goraffe project, implement subgraph parsing.

1. Update parse.go:
   - parseSubgraph():
     - Parse: subgraph [name] { ... }
     - Handle anonymous subgraphs: { ... }
     - Recursively parse contents
     - Return *Subgraph
   
   - Update parseStmt() to handle:
     - subgraph keyword
     - Bare { indicating anonymous subgraph
   
   - Handle default attribute statements:
     - node [attr=value] → set default node attrs
     - edge [attr=value] → set default edge attrs
     - graph [attr=value] → set graph attrs (in subgraph context)

2. Subgraph handling:
   - Subgraph can appear as edge endpoint: subgraph {} -> B
   - Means all nodes in subgraph connect to B

3. Update parser_test.go:
   - TestParse_Subgraph_Named
   - TestParse_Subgraph_Anonymous
   - TestParse_Subgraph_Cluster
   - TestParse_Subgraph_Nested
   - TestParse_Subgraph_WithAttributes
   - TestParse_DefaultNodeAttrs
   - TestParse_DefaultEdgeAttrs
   - TestParse_SubgraphAsEdgeEndpoint
```

---

### Prompt 37: Parse Functions Public API

```text
Building on the existing goraffe project, expose the public parsing API.

1. Update parse.go with public functions:
   - Parse(r io.Reader) (*Graph, error)
     - Read all from reader, parse as DOT
   - ParseString(dot string) (*Graph, error)
     - Parse DOT string directly
   - ParseFile(path string) (*Graph, error)
     - Open file, read contents, parse

2. Error handling:
   - Create ParseError type:
     - Message string
     - Line, Col int
     - Snippet string (surrounding context)
   - Wrap parser errors with location info

3. Integration tests:
   - TestParse_FromReader
   - TestParseString_SimpleGraph
   - TestParseFile_ValidFile
   - TestParseFile_NotFound_Error
   - TestParse_SyntaxError_HasLocation

4. Round-trip tests:
   - TestParse_RoundTrip_SimpleGraph
     - Create graph, output DOT, parse DOT, compare
   - TestParse_RoundTrip_ComplexGraph
     - Include subgraphs, attributes, etc.
   - Note: Not exact string match, but semantic equivalence

5. Add test fixtures in testdata/:
   - simple.dot
   - complex.dot
   - cluster.dot
```

---

### Prompt 38: Format and Layout Enums

```text
Building on the existing goraffe project, add rendering-related types.

1. Create render.go:
   - Format type (string enum):
     - PNG Format = "png"
     - SVG Format = "svg"
     - PDF Format = "pdf"
     - DOT Format = "dot"
   
   - Layout type (string enum):
     - LayoutDot Layout = "dot"
     - LayoutNeato Layout = "neato"
     - LayoutFdp Layout = "fdp"
     - LayoutSfdp Layout = "sfdp"
     - LayoutTwopi Layout = "twopi"
     - LayoutCirco Layout = "circo"
     - LayoutOsage Layout = "osage"
     - LayoutPatchwork Layout = "patchwork"

2. Create render_test.go:
   - TestFormat_StringValues
   - TestLayout_StringValues
   - Verify all constants have expected string values

Keep this simple - just the types. Actual rendering comes next.
```

---

### Prompt 39: RenderError and Sentinel Errors

```text
Building on the existing goraffe project, add rendering error types.

1. Create errors.go:
   - RenderError struct:
     - Err error (underlying error)
     - Stderr string (Graphviz stderr output)
     - ExitCode int
   - Methods:
     - Error() string - format nice message including stderr snippet
     - Unwrap() error - return underlying Err

   - Sentinel errors:
     - ErrGraphvizNotFound = errors.New("goraffe: graphviz not found in PATH")
     - ErrInvalidDOT = errors.New("goraffe: invalid DOT syntax")
     - ErrRenderFailed = errors.New("goraffe: rendering failed")

2. Create errors_test.go:
   - TestRenderError_Error_IncludesStderr
   - TestRenderError_Unwrap
   - TestRenderError_Is_RenderFailed
   - TestSentinelErrors_Distinct

Example usage:
  err := g.Render(PNG, w)
  if errors.Is(err, ErrGraphvizNotFound) {
      log.Fatal("Please install Graphviz")
  }
  if renderErr, ok := err.(*RenderError); ok {
      fmt.Println("Graphviz said:", renderErr.Stderr)
  }
```

---

### Prompt 40: Graphviz CLI Detection

```text
Building on the existing goraffe project, add Graphviz CLI detection.

1. Update render.go:
   - Internal function: findGraphviz(layout Layout) (string, error)
     - Use exec.LookPath to find the binary
     - Binary name matches layout: "dot", "neato", etc.
     - Return full path or ErrGraphvizNotFound

   - Public function: GraphvizVersion() (string, error)
     - Run "dot -V" and parse output
     - Return version string or error
     - Useful for debugging/diagnostics

   - Internal function: checkGraphvizInstalled() error
     - Quick check if any Graphviz binary is available
     - Used to provide early errors

2. Update render_test.go:
   - TestFindGraphviz_Dot (may need to skip if not installed)
   - TestFindGraphviz_AllLayouts
   - TestFindGraphviz_InvalidLayout
   - TestGraphvizVersion_ReturnsVersion

3. Add test helper:
   - requireGraphviz(t *testing.T) - skip test if Graphviz not installed

Note: Tests should be skippable on systems without Graphviz.
```

---

### Prompt 41: Graph.Render to io.Writer

```text
Building on the existing goraffe project, implement the core rendering method.

1. Update render.go:
   - RenderOption interface:
     - applyRender(*renderConfig)
   - renderConfig struct:
     - layout Layout (default: LayoutDot)
   
   - Method on Graph:
     - Render(format Format, w io.Writer, opts ...RenderOption) error
       1. Build renderConfig from opts (default layout = dot)
       2. Find Graphviz binary for the layout
       3. Generate DOT string from graph
       4. Execute: echo $DOT | $BINARY -T$FORMAT
       5. Write stdout to w
       6. If error, wrap in RenderError with stderr

2. Implementation details:
   - Use exec.Command with Stdin pipe
   - Capture both stdout and stderr
   - Handle non-zero exit codes

3. Update render_test.go (require Graphviz):
   - TestGraph_Render_PNG_ProducesOutput
   - TestGraph_Render_SVG_ProducesOutput
   - TestGraph_Render_DOT_ProducesOutput
   - TestGraph_Render_InvalidGraph_Error (if possible to trigger)
   - TestGraph_Render_ToBuffer

4. Add validation helpers:
   - assertValidPNG(t, data []byte) - check PNG magic bytes
   - assertValidSVG(t, data []byte) - check XML/SVG structure
```

---

### Prompt 42: RenderToFile and RenderBytes Conveniences

```text
Building on the existing goraffe project, add convenience rendering methods.

1. Update render.go:
   - Method: RenderToFile(format Format, path string, opts ...RenderOption) error
     - Create file
     - Call Render with file as writer
     - Close file
     - Handle errors (clean up partial file on error)
   
   - Method: RenderBytes(format Format, opts ...RenderOption) ([]byte, error)
     - Create bytes.Buffer
     - Call Render with buffer as writer
     - Return buffer.Bytes()

2. Update render_test.go:
   - TestGraph_RenderToFile_CreatesFile
   - TestGraph_RenderToFile_ValidContent
   - TestGraph_RenderToFile_ErrorCleansUp (if rendering fails)
   - TestGraph_RenderBytes_ReturnsPNG
   - TestGraph_RenderBytes_ReturnsSVG

3. Integration test:
   - TestRender_CompleteWorkflow
     - Create complex graph
     - Render to file
     - Verify file exists and is valid
     - Clean up
```

---

### Prompt 43: WithLayout Render Option

```text
Building on the existing goraffe project, add layout engine selection.

1. Update render.go:
   - Implement WithLayout(l Layout) RenderOption
     - Sets layout in renderConfig

2. Ensure all layout engines work:
   - dot (default, hierarchical)
   - neato (spring model)
   - fdp (force-directed)
   - sfdp (scalable force-directed)
   - twopi (radial)
   - circo (circular)
   - osage (clustered)
   - patchwork (treemap)

3. Update render_test.go:
   - TestGraph_Render_WithLayout_Neato
   - TestGraph_Render_WithLayout_Fdp
   - TestGraph_Render_WithLayout_Circo
   - TestGraph_Render_AllLayouts (table-driven test for all)
   - TestGraph_Render_DefaultLayout_IsDot

4. Final integration test:
   - TestGoraffe_EndToEnd
     - Create graph with nodes, edges, subgraphs, attributes
     - Render to multiple formats
     - Parse a DOT file
     - Modify and re-render
     - Verify everything works together

5. Documentation:
   - Ensure doc.go has complete package overview
   - Add examples in example_test.go:
     - Example_simpleGraph
     - Example_withSubgraphs
     - Example_htmlLabels
     - Example_parseAndModify
```

---

## Summary

This prompt series takes the developer through 43 incremental steps, each building on the previous:

1. **Steps 1-5**: Foundation (Graph, Node, Edge, basic options)
2. **Steps 6-14**: Attribute system (typed attrs, functional options, escape hatches)
3. **Steps 15-19**: DOT generation (output formatting, escaping)
4. **Steps 20-26**: Label builders (HTML tables, record labels, ports)
5. **Steps 27-32**: Subgraphs (clusters, nesting, rank constraints)
6. **Steps 33-37**: DOT parsing (lexer, parser, public API)
7. **Steps 38-43**: Rendering (CLI integration, all formats/layouts)

Each prompt:
- Has clear scope
- Builds on previous work
- Includes specific test requirements
- Avoids orphaned code
- Maintains working software at each step
