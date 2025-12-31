# Goraffe Implementation Checklist

## Overview

This checklist tracks the implementation progress of the Goraffe library. Each item corresponds to a prompt in `prompts.md`. Mark items complete as you finish each step.

**Legend:**

- ⬜ Not started
- 🟡 In progress
- ✅ Complete
- ⏸️ Blocked

---

## Phase 1: Foundation (Steps 1-5)

### Step 1: Project Setup and Basic Graph Struct

- ✅ Initialize Go module `github.com/example/goraffe`
- ✅ Create `doc.go` with package documentation
- ✅ Create `graph.go` with Graph struct
  - ✅ Add `name` field (string)
  - ✅ Add `directed` field (bool)
  - ✅ Add `strict` field (bool)
- ✅ Implement `NewGraph()` constructor
- ✅ Implement `IsDirected()` method
- ✅ Implement `IsStrict()` method
- ✅ Implement `Name()` method
- ✅ Create `graph_test.go`
  - ✅ `TestNewGraph_DefaultValues`
  - ✅ Test getter methods

### Step 2: Node Struct and NewNode

- ✅ Create `node.go`
- ✅ Define Node struct with `id` field
- ✅ Implement `NewNode(id string)` constructor
- ✅ Implement `ID()` method
- ✅ Create `node_test.go`
  - ✅ `TestNewNode_SetsID`
  - ✅ `TestNewNode_EmptyID`
  - ✅ `TestNode_ID_ReturnsCorrectValue`

### Step 3: Graph.AddNode Method

- ✅ Add `nodeOrder` field to Graph ([]*Node)
- ✅ Add `nodes` field to Graph (map[string]int)
- ✅ Initialize both in `NewGraph()`
- ✅ Implement `AddNode(n *Node)` method
  - ✅ Replace-in-place for duplicate IDs
  - ✅ Append to nodeOrder and store index for new nodes
- ✅ Implement `GetNode(id string)` method
  - ✅ Lookup index in nodes map
  - ✅ Return nodeOrder[idx] or nil
- ✅ Implement `Nodes()` method (returns nodeOrder)
- ✅ Update `graph_test.go`
  - ✅ `TestGraph_AddNode_SingleNode`
  - ✅ `TestGraph_AddNode_MultipleNodes`
  - ✅ `TestGraph_AddNode_DuplicateID`
  - ✅ `TestGraph_AddNode_PreservesOrder`
  - ✅ `TestGraph_GetNode_Exists`
  - ✅ `TestGraph_GetNode_NotFound`
  - ✅ `TestGraph_Nodes_ReturnsAllNodes`
  - ✅ `TestGraph_Nodes_ReturnsInInsertionOrder`

### Step 4: Edge Struct and Graph.AddEdge

- ✅ Create `edge.go`
- ✅ Define Edge struct with `from` and `to` fields
- ✅ Implement `From()` method
- ✅ Implement `To()` method
- ✅ Add `edges` field to Graph ([]*Edge)
- ✅ Initialize edges slice in `NewGraph()`
- ✅ Implement `AddEdge(from, to *Node)` method
  - ✅ Create new Edge
  - ✅ Implicit node addition
  - ✅ Append to edges slice (allows parallel edges)
  - ✅ Return created edge
- ✅ Implement `Edges()` method
- ✅ Create `edge_test.go`
  - ✅ `TestEdge_FromTo_ReturnsCorrectNodes`
  - ✅ `TestGraph_AddEdge_BothNodesExist`
  - ✅ `TestGraph_AddEdge_ImplicitNodeAdd`
  - ✅ `TestGraph_AddEdge_PartialImplicitAdd`
  - ✅ `TestGraph_AddEdge_AllowsParallelEdges`
  - ✅ `TestGraph_AddEdge_AllowsSelfLoops`
  - ✅ `TestGraph_Edges_ReturnsAllEdges`
  - ✅ `TestGraph_Edges_ReturnsInInsertionOrder`

### Step 5: Directed/Undirected/Strict Graph Options

- ✅ Create `options.go`
- ✅ Define `GraphOption` interface
- ✅ Define `graphOptionFunc` type
- ✅ Implement `Directed` option
- ✅ Implement `Undirected` option
- ✅ Implement `Strict` option
- ✅ Update `NewGraph` to accept variadic `GraphOption`
- ✅ Create `options_test.go`
  - ✅ `TestNewGraph_Directed`
  - ✅ `TestNewGraph_Undirected`
  - ✅ `TestNewGraph_Strict`
  - ✅ `TestNewGraph_DirectedAndStrict`
  - ✅ `TestNewGraph_MultipleOptions_LastWins`
  - ✅ `TestNewGraph_NoOptions`
  - ✅ `TestGraphOption_Interface`
- ✅ Update existing tests for new signature

---

## Phase 2: Attributes (Steps 6-14)

### Step 6: Shape Enum and NodeAttributes Struct

- ✅ Create `attributes.go`
- ✅ Define `Shape` type
- ✅ Add Shape constants
  - ✅ `ShapeBox`
  - ✅ `ShapeCircle`
  - ✅ `ShapeEllipse`
  - ✅ `ShapeDiamond`
  - ✅ `ShapeRecord`
  - ✅ `ShapePlaintext`
- ✅ Define `NodeAttributes` struct
  - ✅ `Label` field
  - ✅ `Shape` field
  - ✅ `Color` field
  - ✅ `FillColor` field
  - ✅ `FontName` field
  - ✅ `FontSize` field
  - ✅ `custom` field (unexported map)
- ✅ Implement `Custom()` method on NodeAttributes
- ✅ Update `node.go` to add `attrs` field
- ✅ Implement `Attrs()` method on Node
- ✅ Create `attributes_test.go`
  - ✅ `TestNodeAttributes_ZeroValue`
  - ✅ `TestNodeAttributes_Custom_ReturnsCopy`
- ✅ Update `node_test.go`
  - ✅ `TestNode_Attrs_ReturnsAttributes`

### Step 7: NodeOption Interface and Basic Functional Options

- ✅ Define `NodeOption` interface
- ✅ Define `nodeOptionFunc` type
- ✅ Implement `newNodeOption` helper
- ✅ Implement `WithShape(s Shape)` option
- ✅ Implement `WithLabel(l string)` option
- ✅ Implement `WithColor(c string)` option
- ✅ Implement `WithFillColor(c string)` option
- ✅ Implement `WithFontName(f string)` option
- ✅ Implement `WithFontSize(s float64)` option
- ✅ Create `node_options_test.go`
  - ✅ `TestWithShape_SetsShape`
  - ✅ `TestWithLabel_SetsLabel`
  - ✅ `TestWithColor_SetsColor`
  - ✅ `TestWithFillColor_SetsFillColor`
  - ✅ `TestWithFontName_SetsFontName`
  - ✅ `TestWithFontSize_SetsFontSize`

### Step 8: Wire NodeOptions into NewNode

- ✅ Update `NewNode` signature to accept `...NodeOption`
- ✅ Apply options to node attributes in constructor
- ✅ Make `NodeAttributes` implement `NodeOption`
  - ✅ Add `applyNode` method
  - ✅ Implement non-zero field merging
- ✅ Update `node_test.go`
  - ✅ `TestNewNode_WithOptions`
  - ✅ `TestNewNode_WithMultipleOptions`
  - ✅ `TestNewNode_WithNodeAttributesStruct`
  - ✅ `TestNewNode_OptionsAppliedInOrder`
- ✅ Update existing tests for new signature (no changes needed - variadic options)
- ✅ Delete `node_options_test.go` (tests private API, redundant with public API tests)

### Step 9: EdgeStyle Enum and EdgeAttributes Struct

- ✅ Define `EdgeStyle` type
- ✅ Add EdgeStyle constants
  - ✅ `EdgeStyleSolid`
  - ✅ `EdgeStyleDashed`
  - ✅ `EdgeStyleDotted`
  - ✅ `EdgeStyleBold`
- ✅ Define `ArrowType` type
- ✅ Add ArrowType constants
  - ✅ `ArrowNormal`
  - ✅ `ArrowDot`
  - ✅ `ArrowNone`
  - ✅ `ArrowVee`
- ✅ Define `EdgeAttributes` struct
  - ✅ `Label` field
  - ✅ `Color` field
  - ✅ `Style` field
  - ✅ `ArrowHead` field
  - ✅ `ArrowTail` field
  - ✅ `Weight` field
  - ✅ `custom` field (unexported map)
- ✅ Implement `Custom()` method on EdgeAttributes
- ✅ Update `edge.go` to add `attrs` field
- ✅ Implement `Attrs()` method on Edge
- ✅ Update `attributes_test.go`
  - ✅ `TestEdgeAttributes_ZeroValue`
  - ✅ `TestEdgeAttributes_Custom_ReturnsCopy`
- ✅ Update `edge_test.go`
  - ✅ `TestEdge_Attrs_ReturnsAttributes`

### Step 10: EdgeOption Interface and Basic Options

- ✅ Define `EdgeOption` interface
- ✅ Define `edgeOptionFunc` type
- ✅ Implement `newEdgeOption` helper
- ✅ Implement `WithEdgeLabel(l string)` option
- ✅ Implement `WithEdgeColor(c string)` option
- ✅ Implement `WithEdgeStyle(s EdgeStyle)` option
- ✅ Implement `WithArrowHead(a ArrowType)` option
- ✅ Implement `WithArrowTail(a ArrowType)` option
- ✅ Implement `WithWeight(w float64)` option
- ✅ Make `EdgeAttributes` implement `EdgeOption`
  - ✅ Implement `applyEdge` method with non-zero field merging
  - ✅ Document that custom fields are NOT copied (per-instance)
- ✅ Create `edge_options_test.go`
  - ✅ `TestWithEdgeLabel_SetsLabel`
  - ✅ `TestWithEdgeColor_SetsColor`
  - ✅ `TestWithEdgeStyle_SetsStyle`
  - ✅ `TestWithArrowHead_SetsArrowHead`
  - ✅ `TestWithArrowTail_SetsArrowTail`
  - ✅ `TestWithWeight_SetsWeight`
  - ✅ `TestEdgeOption_MultipleOptionsCanBeApplied`
  - ✅ `TestEdgeAttributes_AsOption`
  - Note: Revisit whether these private API tests are needed after Step 11

### Step 11: Wire EdgeOptions into AddEdge

- ✅ Update `AddEdge` signature to accept `...EdgeOption`
- ✅ Apply options to edge attributes
- ✅ Update `edge_test.go` and `graph_test.go`
  - ✅ `TestGraph_AddEdge_WithOptions`
  - ✅ `TestGraph_AddEdge_WithMultipleOptions`
  - ✅ `TestGraph_AddEdge_WithEdgeAttributesStruct`
  - ✅ `TestGraph_AddEdge_OptionsAppliedInOrder`
- ✅ Update existing AddEdge calls in tests (no changes needed - variadic parameter)

### Step 12: GraphAttributes and GraphOption

- ✅ Define `RankDir` type
- ✅ Add RankDir constants
  - ✅ `RankDirTB`
  - ✅ `RankDirBT`
  - ✅ `RankDirLR`
  - ✅ `RankDirRL`
- ✅ Define `SplineType` type
- ✅ Add SplineType constants
  - ✅ `SplineTrue`
  - ✅ `SplineFalse`
  - ✅ `SplineOrtho`
  - ✅ `SplinePolyline`
  - ✅ `SplineCurved`
  - ✅ Additional: `SplineSpline`, `SplineLine`, `SplineNone`
- ✅ Define `GraphAttributes` struct (using pointer fields)
  - ✅ `label` field (*string)
  - ✅ `rankDir` field (*RankDir)
  - ✅ `bgColor` field (*string)
  - ✅ `fontName` field (*string)
  - ✅ `fontSize` field (*float64)
  - ✅ `splines` field (*SplineType)
  - ✅ `nodeSep` field (*float64)
  - ✅ `rankSep` field (*float64)
  - ✅ `compound` field (*bool)
  - ✅ `custom` field (unexported map)
  - ✅ Getter methods for all fields with zero-value documentation
- ✅ Add `attrs` field to Graph
- ✅ Implement `Attrs()` method on Graph
- ✅ Implement GraphOption functions
  - ✅ `WithGraphLabel`
  - ✅ `WithRankDir`
  - ✅ `WithBgColor`
  - ✅ `WithGraphFontName`
  - ✅ `WithGraphFontSize`
  - ✅ `WithSplines`
  - ✅ `WithNodeSep`
  - ✅ `WithRankSep`
  - ✅ `WithCompound`
- ✅ Create tests
  - ✅ `TestGraphAttributes_ZeroValue`
  - ✅ `TestGraphAttributes_Custom_ReturnsCopy`
  - ✅ `TestGraph_Attrs_ReturnsGraphAttributes`
  - ✅ `TestWithGraphLabel_SetsLabel`
  - ✅ `TestWithRankDir_SetsRankDir`
  - ✅ `TestWithBgColor_SetsBgColor`
  - ✅ `TestWithGraphFontName_SetsFontName`
  - ✅ `TestWithGraphFontSize_SetsFontSize`
  - ✅ `TestWithSplines_SetsSplines`
  - ✅ `TestWithNodeSep_SetsNodeSep`
  - ✅ `TestWithRankSep_SetsRankSep`
  - ✅ `TestWithCompound_SetsCompound`
  - ✅ `TestNewGraph_WithMultipleGraphOptions`
  - ✅ `TestNewGraph_GraphAttributesDoNotAffectNodeEdgeOperations`

### Step 13: Default Node/Edge Attributes on Graph

- ⬜ Add `defaultNodeAttrs` field to Graph
- ⬜ Add `defaultEdgeAttrs` field to Graph
- ⬜ Initialize both in `NewGraph()`
- ⬜ Implement `DefaultNodeAttrs()` method
- ⬜ Implement `DefaultEdgeAttrs()` method
- ⬜ Implement `WithDefaultNodeAttrs(opts ...NodeOption)` option
- ⬜ Implement `WithDefaultEdgeAttrs(opts ...EdgeOption)` option
- ⬜ Update tests
  - ⬜ `TestGraph_WithDefaultNodeAttrs`
  - ⬜ `TestGraph_WithDefaultEdgeAttrs`
  - ⬜ `TestGraph_DefaultAttrs_AppliesMultipleOptions`

### Step 14: WithAttribute Escape Hatch

- ⬜ Add `SetCustom(key, value string)` to NodeAttributes
- ⬜ Add `SetCustom(key, value string)` to EdgeAttributes
- ⬜ Add `SetCustom(key, value string)` to GraphAttributes
- ⬜ Ensure `Custom()` returns copy on all types
- ⬜ Implement `WithNodeAttribute(key, value string)` option
- ⬜ Implement `WithEdgeAttribute(key, value string)` option
- ⬜ Implement `WithGraphAttribute(key, value string)` option
- ⬜ Create tests
  - ⬜ `TestWithNodeAttribute_SetsCustom`
  - ⬜ `TestWithEdgeAttribute_SetsCustom`
  - ⬜ `TestWithGraphAttribute_SetsCustom`
  - ⬜ `TestCustomAttributes_DoNotOverrideTyped`
  - ⬜ `TestCustomAttributes_MultipleCalls_Accumulate`

---

## Phase 3: DOT Generation (Steps 15-19)

### Step 15: Graph.String() - Basic DOT Output

- ⬜ Create `dot.go`
- ⬜ Implement `String()` method on Graph
- ⬜ Implement `WriteDOT(w io.Writer)` method
- ⬜ Handle digraph vs graph keywords
- ⬜ Handle strict prefix
- ⬜ Handle graph name
- ⬜ Create `dot_test.go`
  - ⬜ `TestGraph_String_EmptyDirected`
  - ⬜ `TestGraph_String_EmptyUndirected`
  - ⬜ `TestGraph_String_Strict`
  - ⬜ `TestGraph_String_WithName`
  - ⬜ `TestGraph_WriteDOT_WritesToWriter`

### Step 16: Node DOT Rendering with Attributes

- ⬜ Add internal node rendering method
- ⬜ Update `String()`/`WriteDOT()` to include nodes
- ⬜ Handle node attribute rendering
  - ⬜ Label → label="value"
  - ⬜ Shape → shape="value"
  - ⬜ Color → color="value"
  - ⬜ FillColor → fillcolor="value"
  - ⬜ FontName → fontname="value"
  - ⬜ FontSize → fontsize="value"
  - ⬜ Custom attributes
- ⬜ Only output non-zero/non-empty attributes
- ⬜ Handle node ID quoting
- ⬜ Update `dot_test.go`
  - ⬜ `TestDOT_SingleNode_NoAttributes`
  - ⬜ `TestDOT_SingleNode_WithLabel`
  - ⬜ `TestDOT_SingleNode_WithShape`
  - ⬜ `TestDOT_SingleNode_MultipleAttributes`
  - ⬜ `TestDOT_SingleNode_CustomAttribute`
  - ⬜ `TestDOT_MultipleNodes`

### Step 17: Edge DOT Rendering with Attributes

- ⬜ Add internal edge rendering method
- ⬜ Update `String()`/`WriteDOT()` to include edges
- ⬜ Handle directed (→) vs undirected (--)
- ⬜ Handle edge attribute rendering
  - ⬜ Label → label="value"
  - ⬜ Color → color="value"
  - ⬜ Style → style="value"
  - ⬜ ArrowHead → arrowhead="value"
  - ⬜ ArrowTail → arrowtail="value"
  - ⬜ Weight → weight="value"
  - ⬜ Custom attributes
- ⬜ Update `dot_test.go`
  - ⬜ `TestDOT_SingleEdge_NoAttributes`
  - ⬜ `TestDOT_SingleEdge_Directed`
  - ⬜ `TestDOT_SingleEdge_Undirected`
  - ⬜ `TestDOT_SingleEdge_WithLabel`
  - ⬜ `TestDOT_SingleEdge_MultipleAttributes`
  - ⬜ `TestDOT_MultipleEdges`
  - ⬜ `TestDOT_CompleteGraph`

### Step 18: Graph and Default Attributes in DOT Output

- ⬜ Output graph attributes after opening brace
- ⬜ Output default node attributes (node [...];)
- ⬜ Output default edge attributes (edge [...];)
- ⬜ Only output if non-zero defaults exist
- ⬜ Implement correct output order
  1. ⬜ Graph declaration
  2. ⬜ Graph attributes
  3. ⬜ Default node attributes
  4. ⬜ Default edge attributes
  5. ⬜ Nodes
  6. ⬜ Edges
  7. ⬜ Closing brace
- ⬜ Update `dot_test.go`
  - ⬜ `TestDOT_GraphAttributes_RankDir`
  - ⬜ `TestDOT_GraphAttributes_Label`
  - ⬜ `TestDOT_GraphAttributes_Multiple`
  - ⬜ `TestDOT_DefaultNodeAttrs`
  - ⬜ `TestDOT_DefaultEdgeAttrs`
  - ⬜ `TestDOT_FullGraph_WithAllSections`

### Step 19: String Escaping in DOT Output

- ⬜ Create DOT string escaping helper
  - ⬜ Escape backslashes
  - ⬜ Escape double quotes
  - ⬜ Escape newlines
  - ⬜ Handle other special characters
- ⬜ Create quoting decision helper
- ⬜ Apply escaping to all DOT output
  - ⬜ Node IDs
  - ⬜ Attribute values
  - ⬜ Graph names
- ⬜ Update `dot_test.go`
  - ⬜ `TestDOT_NodeID_WithSpaces`
  - ⬜ `TestDOT_NodeID_WithSpecialChars`
  - ⬜ `TestDOT_Label_WithQuotes`
  - ⬜ `TestDOT_Label_WithNewlines`
  - ⬜ `TestDOT_Label_WithBackslashes`
  - ⬜ `TestDOT_ComplexStrings`

---

## Phase 4: Labels (Steps 20-26)

### Step 20: HTMLCell and HTMLRow Types

- ⬜ Create `labels.go`
- ⬜ Define `HTMLCell` struct
  - ⬜ `content` field
  - ⬜ `port` field
  - ⬜ `bold` field
  - ⬜ `italic` field
  - ⬜ `underline` field
  - ⬜ `colSpan` field
  - ⬜ `rowSpan` field
  - ⬜ `bgColor` field
  - ⬜ `align` field
- ⬜ Implement `Cell(content string)` constructor
- ⬜ Implement chainable methods
  - ⬜ `Port(id string)`
  - ⬜ `Bold()`
  - ⬜ `Italic()`
  - ⬜ `Underline()`
  - ⬜ `ColSpan(n int)`
  - ⬜ `RowSpan(n int)`
  - ⬜ `BgColor(color string)`
  - ⬜ `Align(a string)`
- ⬜ Define `HTMLRow` struct
- ⬜ Implement `Row(cells ...*HTMLCell)` constructor
- ⬜ Implement `Cells()` method
- ⬜ Create `labels_test.go`
  - ⬜ `TestCell_Content`
  - ⬜ `TestCell_Chaining`
  - ⬜ `TestCell_AllOptions`
  - ⬜ `TestRow_ContainsCells`
  - ⬜ `TestRow_MultipleCells`

### Step 21: HTMLTable Builder

- ⬜ Define `HTMLLabel` struct
  - ⬜ `rows` field
  - ⬜ `border` field
  - ⬜ `cellBorder` field
  - ⬜ `cellSpacing` field
  - ⬜ `cellPadding` field
  - ⬜ `bgColor` field
- ⬜ Implement `HTMLTable(rows ...*HTMLRow)` constructor
- ⬜ Implement chainable methods
  - ⬜ `Border(n int)`
  - ⬜ `CellBorder(n int)`
  - ⬜ `CellSpacing(n int)`
  - ⬜ `CellPadding(n int)`
  - ⬜ `BgColor(color string)`
- ⬜ Implement `String()` method for HTML rendering
  - ⬜ Output wrapped in < >
  - ⬜ TABLE element with attributes
  - ⬜ TR for each row
  - ⬜ TD for each cell
  - ⬜ Formatting tags (B, I, U)
  - ⬜ PORT attribute
- ⬜ Update `labels_test.go`
  - ⬜ `TestHTMLTable_SimpleTable`
  - ⬜ `TestHTMLTable_WithTableAttributes`
  - ⬜ `TestHTMLTable_CellWithPort`
  - ⬜ `TestHTMLTable_CellWithFormatting`
  - ⬜ `TestHTMLTable_CellWithSpan`
  - ⬜ `TestHTMLTable_ComplexTable`

### Step 22: Port Type and Cell Port Reference

- ⬜ Create `port.go`
- ⬜ Define `Port` struct
  - ⬜ `id` field
  - ⬜ `nodeID` field
- ⬜ Implement `ID()` method
- ⬜ Update `HTMLCell`
  - ⬜ Add `portRef` field
  - ⬜ Update `Port()` method to create Port
  - ⬜ Add `GetPort()` method
- ⬜ Add mechanism to associate ports with nodes
  - ⬜ Internal method on HTMLLabel to set node context
  - ⬜ Update Port.nodeID when label attached
- ⬜ Create `port_test.go`
  - ⬜ `TestPort_ID`
  - ⬜ `TestCell_GetPort_ReturnsPort`
  - ⬜ `TestCell_GetPort_NilIfNoPort`
- ⬜ Update `labels_test.go`
  - ⬜ `TestHTMLLabel_PortsKnowNodeID`

### Step 23: FromPort/ToPort Edge Options

- ⬜ Add `fromPort` field to EdgeAttributes
- ⬜ Add `toPort` field to EdgeAttributes
- ⬜ Implement `FromPort()` method on EdgeAttributes
- ⬜ Implement `ToPort()` method on EdgeAttributes
- ⬜ Implement `FromPort(p *Port)` EdgeOption
- ⬜ Implement `ToPort(p *Port)` EdgeOption
- ⬜ Update DOT generation for port syntax
  - ⬜ Handle fromPort: "nodeID":"portID"
  - ⬜ Handle toPort: "nodeID":"portID"
- ⬜ Create/update tests
  - ⬜ `TestFromPort_SetsPort`
  - ⬜ `TestToPort_SetsPort`
  - ⬜ `TestDOT_Edge_WithFromPort`
  - ⬜ `TestDOT_Edge_WithToPort`
  - ⬜ `TestDOT_Edge_WithBothPorts`

### Step 24: HTML Label DOT Output Integration

- ⬜ Add `htmlLabel` field to NodeAttributes
- ⬜ Add `rawHTMLLabel` field to NodeAttributes
- ⬜ Implement `WithHTMLLabel(label *HTMLLabel)` option
- ⬜ Implement `WithRawHTMLLabel(html string)` option
- ⬜ Update DOT generation
  - ⬜ Output label=<...> for HTML labels
  - ⬜ No quotes, angle brackets
  - ⬜ HTML labels take precedence over Label
- ⬜ Wire port node association
- ⬜ Update tests
  - ⬜ `TestWithHTMLLabel_SetsLabel`
  - ⬜ `TestWithRawHTMLLabel_SetsLabel`
  - ⬜ `TestDOT_Node_WithHTMLLabel`
  - ⬜ `TestDOT_Node_WithHTMLLabel_Ports`
  - ⬜ `TestDOT_Node_WithRawHTMLLabel`
  - ⬜ `TestDOT_HTMLLabel_NotDoubleEscaped`

### Step 25: Record Field and FieldGroup

- ⬜ Define `RecordField` struct
  - ⬜ `content` field
  - ⬜ `port` field
  - ⬜ `portRef` field
- ⬜ Implement `Field(content string)` constructor
- ⬜ Implement `Port(id string)` method (chainable)
- ⬜ Implement `GetPort()` method
- ⬜ Define `RecordGroup` struct
- ⬜ Implement `FieldGroup(elements ...RecordElement)` constructor
- ⬜ Define `RecordElement` interface
- ⬜ Make RecordField implement RecordElement
- ⬜ Make RecordGroup implement RecordElement
- ⬜ Define `RecordLabel` struct
- ⬜ Implement `RecordLabel(elements ...RecordElement)` constructor
- ⬜ Implement `String()` method for record rendering
  - ⬜ Fields separated by |
  - ⬜ Groups wrapped in { }
  - ⬜ Ports: <portID> content
  - ⬜ Escape special chars
- ⬜ Create `record_labels_test.go`
  - ⬜ `TestRecordField_Content`
  - ⬜ `TestRecordField_WithPort`
  - ⬜ `TestRecordGroup_Nesting`
  - ⬜ `TestRecordLabel_SimpleFields`
  - ⬜ `TestRecordLabel_WithGroup`
  - ⬜ `TestRecordLabel_Escaping`

### Step 26: WithRecordLabel and DOT Output

- ⬜ Add `recordLabel` field to NodeAttributes
- ⬜ Implement `WithRecordLabel(elements ...RecordElement)` option
  - ⬜ Create RecordLabel from elements
  - ⬜ Set shape to Record
- ⬜ Update DOT generation
  - ⬜ Output label="..." for record labels
  - ⬜ Record labels ARE quoted
  - ⬜ Ensure shape="record" is output
- ⬜ Wire port association for record labels
- ⬜ Update tests
  - ⬜ `TestWithRecordLabel_SetsLabel`
  - ⬜ `TestWithRecordLabel_SetsShape`
  - ⬜ `TestDOT_Node_WithRecordLabel_Simple`
  - ⬜ `TestDOT_Node_WithRecordLabel_WithPorts`
  - ⬜ `TestDOT_Node_WithRecordLabel_Nested`
  - ⬜ `TestDOT_Edge_ToRecordPort`

---

## Phase 5: Subgraphs (Steps 27-32)

### Step 27: Subgraph Struct and Graph.Subgraph()

- ⬜ Create `subgraph.go`
- ⬜ Define `Subgraph` struct
  - ⬜ `name` field
  - ⬜ `nodes` field
  - ⬜ `edges` field
  - ⬜ `isCluster` field
  - ⬜ `parent` field
- ⬜ Implement `Name()` method
- ⬜ Implement `IsCluster()` method
- ⬜ Implement `AddNode(n *Node)` method
  - ⬜ Add to subgraph's nodes
  - ⬜ Add to parent graph
- ⬜ Implement `Nodes()` method
- ⬜ Implement `AddEdge()` method (delegates to parent)
- ⬜ Add `subgraphs` field to Graph
- ⬜ Implement `Subgraph(name string, fn func(*Subgraph))` method
- ⬜ Implement `Subgraphs()` method
- ⬜ Create `subgraph_test.go`
  - ⬜ `TestSubgraph_Name`
  - ⬜ `TestSubgraph_IsCluster_True`
  - ⬜ `TestSubgraph_IsCluster_False`
  - ⬜ `TestSubgraph_AddNode`
  - ⬜ `TestSubgraph_AddNode_AlsoAddsToParent`
  - ⬜ `TestGraph_Subgraph_CallsFunction`
  - ⬜ `TestGraph_Subgraph_ReturnsSubgraph`
  - ⬜ `TestGraph_Subgraphs_ReturnsAll`

### Step 28: Cluster Detection and Subgraph Attributes

- ⬜ Define `SubgraphAttributes` struct
  - ⬜ `Label` field
  - ⬜ `Style` field
  - ⬜ `Color` field
  - ⬜ `FillColor` field
  - ⬜ `FontName` field
  - ⬜ `FontSize` field
  - ⬜ `custom` field
- ⬜ Add `attrs` field to Subgraph
- ⬜ Implement `Attrs()` method
- ⬜ Implement setter methods
  - ⬜ `SetLabel(l string)`
  - ⬜ `SetStyle(s string)`
  - ⬜ `SetColor(c string)`
  - ⬜ `SetFillColor(c string)`
  - ⬜ `SetAttribute(key, value string)`
- ⬜ Document cluster-specific behavior
- ⬜ Update tests
  - ⬜ `TestSubgraph_SetLabel`
  - ⬜ `TestSubgraph_SetStyle`
  - ⬜ `TestSubgraph_SetAttribute`
  - ⬜ `TestSubgraph_Attrs_ReturnsAttributes`
  - ⬜ `TestSubgraph_Cluster_CanHaveStyle`

### Step 29: Nested Subgraphs

- ⬜ Add `subgraphs` field to Subgraph
- ⬜ Implement `Subgraph(name string, fn func(*Subgraph))` on Subgraph
  - ⬜ Create nested subgraph
  - ⬜ Set parent appropriately
  - ⬜ Call fn
  - ⬜ Return subgraph
- ⬜ Implement `Subgraphs()` method on Subgraph
- ⬜ Ensure node tracking works (nodes in root graph)
- ⬜ Update tests
  - ⬜ `TestSubgraph_NestedSubgraph`
  - ⬜ `TestSubgraph_NestedSubgraph_NodesInRoot`
  - ⬜ `TestSubgraph_DeeplyNested`
  - ⬜ `TestSubgraph_NestedCluster`

### Step 30: Subgraph DOT Generation

- ⬜ Add internal subgraph rendering method
- ⬜ Implement subgraph DOT format
  - ⬜ subgraph name { ... }
  - ⬜ Subgraph attributes
  - ⬜ Nodes in subgraph
  - ⬜ Nested subgraphs (recursive)
- ⬜ Update main DOT output order
  1. ⬜ Graph declaration
  2. ⬜ Graph attributes
  3. ⬜ Default node/edge attributes
  4. ⬜ Subgraphs
  5. ⬜ Loose nodes
  6. ⬜ Edges
  7. ⬜ Closing brace
- ⬜ Handle empty subgraph names (anonymous)
- ⬜ Update `dot_test.go`
  - ⬜ `TestDOT_Subgraph_Simple`
  - ⬜ `TestDOT_Subgraph_WithAttributes`
  - ⬜ `TestDOT_Subgraph_Cluster`
  - ⬜ `TestDOT_Subgraph_Nested`
  - ⬜ `TestDOT_Subgraph_Anonymous`
  - ⬜ `TestDOT_Graph_WithSubgraphsAndLooseNodes`

### Step 31: SameRank, MinRank, MaxRank Convenience Methods

- ⬜ Define `Rank` type
- ⬜ Add Rank constants
  - ⬜ `RankSame`
  - ⬜ `RankMin`
  - ⬜ `RankMax`
  - ⬜ `RankSource`
  - ⬜ `RankSink`
- ⬜ Add `rank` field to Subgraph
- ⬜ Implement `SetRank(r Rank)` method
- ⬜ Implement `Rank()` getter
- ⬜ Add internal helper for anonymous rank subgraph
- ⬜ Implement convenience methods on Graph
  - ⬜ `SameRank(nodes ...*Node)`
  - ⬜ `MinRank(nodes ...*Node)`
  - ⬜ `MaxRank(nodes ...*Node)`
  - ⬜ `SourceRank(nodes ...*Node)`
  - ⬜ `SinkRank(nodes ...*Node)`
- ⬜ Update tests
  - ⬜ `TestGraph_SameRank`
  - ⬜ `TestGraph_MinRank`
  - ⬜ `TestGraph_MaxRank`
  - ⬜ `TestGraph_SourceRank`
  - ⬜ `TestGraph_SinkRank`
  - ⬜ `TestSubgraph_SetRank`

### Step 32: Rank Constraint DOT Output

- ⬜ Update subgraph DOT generation for rank
  - ⬜ Output rank="value";
  - ⬜ Place after other subgraph attributes
- ⬜ Handle rank subgraphs from convenience methods
  - ⬜ Anonymous subgraphs
  - ⬜ Only rank attribute and nodes
- ⬜ Update `dot_test.go`
  - ⬜ `TestDOT_Subgraph_WithRank`
  - ⬜ `TestDOT_SameRank_CreatesSubgraph`
  - ⬜ `TestDOT_MinRank_Output`
  - ⬜ `TestDOT_MaxRank_Output`
  - ⬜ `TestDOT_MultipleRankConstraints`
  - ⬜ `TestDOT_ComplexGraph_WithRanks`

---

## Phase 6: Parsing (Steps 33-37)

### Step 33: DOT Lexer

- ⬜ Create `parse.go` (or `lexer.go`)
- ⬜ Define `TokenType` constants
  - ⬜ `TokenEOF`
  - ⬜ `TokenIdent`
  - ⬜ `TokenString`
  - ⬜ `TokenNumber`
  - ⬜ `TokenLBrace`
  - ⬜ `TokenRBrace`
  - ⬜ `TokenLBracket`
  - ⬜ `TokenRBracket`
  - ⬜ `TokenLParen`
  - ⬜ `TokenRParen`
  - ⬜ `TokenSemi`
  - ⬜ `TokenComma`
  - ⬜ `TokenColon`
  - ⬜ `TokenEqual`
  - ⬜ `TokenArrow`
  - ⬜ `TokenHTML`
- ⬜ Define `Token` struct
  - ⬜ `Type` field
  - ⬜ `Value` field
  - ⬜ `Line` field
  - ⬜ `Col` field
- ⬜ Define `Lexer` struct
- ⬜ Implement `Next()` method
- ⬜ Implement `Peek()` method
- ⬜ Implement lexer behavior
  - ⬜ Skip whitespace
  - ⬜ Skip comments (// and /**/)
  - ⬜ Handle quoted strings with escapes
  - ⬜ Handle HTML strings (< >)
  - ⬜ Handle -> and -- tokens
- ⬜ Create `lexer_test.go`
  - ⬜ `TestLexer_SimpleTokens`
  - ⬜ `TestLexer_Identifiers`
  - ⬜ `TestLexer_QuotedStrings`
  - ⬜ `TestLexer_HTMLStrings`
  - ⬜ `TestLexer_Arrows`
  - ⬜ `TestLexer_Comments`
  - ⬜ `TestLexer_CompleteGraph`

### Step 34: DOT Parser - Graph Structure

- ⬜ Define `Parser` struct
  - ⬜ `lexer` field
  - ⬜ `current` field
- ⬜ Implement `advance()` method
- ⬜ Implement `expect(TokenType)` method
- ⬜ Implement `match(TokenType)` method
- ⬜ Implement `parseGraph()` function
  - ⬜ Handle [strict] prefix
  - ⬜ Handle graph/digraph keyword
  - ⬜ Handle optional name
  - ⬜ Handle { } body
- ⬜ Implement `parseStmtList()` helper
- ⬜ Implement `parseStmt()` helper (skeleton)
- ⬜ Create `parser_test.go`
  - ⬜ `TestParse_EmptyDigraph`
  - ⬜ `TestParse_EmptyGraph`
  - ⬜ `TestParse_StrictGraph`
  - ⬜ `TestParse_NamedGraph`
  - ⬜ `TestParse_InvalidSyntax_Error`

### Step 35: DOT Parser - Nodes and Edges

- ⬜ Implement `parseNodeStmt()`
  - ⬜ Parse nodeID
  - ⬜ Parse [attributes]
  - ⬜ Create Node
  - ⬜ Add to graph
- ⬜ Implement `parseEdgeStmt()`
  - ⬜ Parse edge chains
  - ⬜ Parse [attributes]
  - ⬜ Create edges
- ⬜ Implement `parseAttrList()`
  - ⬜ Parse [attr=value, ...]
  - ⬜ Return map[string]string
- ⬜ Implement `parseID()`
  - ⬜ Handle identifier
  - ⬜ Handle quoted string
  - ⬜ Handle number
  - ⬜ Handle HTML string
- ⬜ Implement attribute mapping
  - ⬜ Map known attributes to typed fields
  - ⬜ Store unknown in custom map
- ⬜ Update `parser_test.go`
  - ⬜ `TestParse_SingleNode`
  - ⬜ `TestParse_NodeWithAttributes`
  - ⬜ `TestParse_SingleEdge`
  - ⬜ `TestParse_EdgeWithAttributes`
  - ⬜ `TestParse_EdgeChain`
  - ⬜ `TestParse_MixedNodesAndEdges`

### Step 36: DOT Parser - Subgraphs

- ⬜ Implement `parseSubgraph()`
  - ⬜ Parse subgraph [name] { ... }
  - ⬜ Handle anonymous { ... }
  - ⬜ Recursive content parsing
  - ⬜ Return *Subgraph
- ⬜ Update `parseStmt()` for subgraph handling
  - ⬜ Handle subgraph keyword
  - ⬜ Handle bare { for anonymous
- ⬜ Handle default attribute statements
  - ⬜ node [attr=value]
  - ⬜ edge [attr=value]
  - ⬜ graph [attr=value]
- ⬜ Handle subgraph as edge endpoint
- ⬜ Update `parser_test.go`
  - ⬜ `TestParse_Subgraph_Named`
  - ⬜ `TestParse_Subgraph_Anonymous`
  - ⬜ `TestParse_Subgraph_Cluster`
  - ⬜ `TestParse_Subgraph_Nested`
  - ⬜ `TestParse_Subgraph_WithAttributes`
  - ⬜ `TestParse_DefaultNodeAttrs`
  - ⬜ `TestParse_DefaultEdgeAttrs`
  - ⬜ `TestParse_SubgraphAsEdgeEndpoint`

### Step 37: Parse Functions Public API

- ⬜ Implement `Parse(r io.Reader)` function
- ⬜ Implement `ParseString(dot string)` function
- ⬜ Implement `ParseFile(path string)` function
- ⬜ Create `ParseError` type
  - ⬜ `Message` field
  - ⬜ `Line` field
  - ⬜ `Col` field
  - ⬜ `Snippet` field
- ⬜ Wrap parser errors with location info
- ⬜ Create integration tests
  - ⬜ `TestParse_FromReader`
  - ⬜ `TestParseString_SimpleGraph`
  - ⬜ `TestParseFile_ValidFile`
  - ⬜ `TestParseFile_NotFound_Error`
  - ⬜ `TestParse_SyntaxError_HasLocation`
- ⬜ Create round-trip tests
  - ⬜ `TestParse_RoundTrip_SimpleGraph`
  - ⬜ `TestParse_RoundTrip_ComplexGraph`
- ⬜ Add test fixtures in `testdata/`
  - ⬜ `simple.dot`
  - ⬜ `complex.dot`
  - ⬜ `cluster.dot`

---

## Phase 7: Rendering (Steps 38-43)

### Step 38: Format and Layout Enums

- ⬜ Create `render.go`
- ⬜ Define `Format` type
- ⬜ Add Format constants
  - ⬜ `PNG`
  - ⬜ `SVG`
  - ⬜ `PDF`
  - ⬜ `DOT`
- ⬜ Define `Layout` type
- ⬜ Add Layout constants
  - ⬜ `LayoutDot`
  - ⬜ `LayoutNeato`
  - ⬜ `LayoutFdp`
  - ⬜ `LayoutSfdp`
  - ⬜ `LayoutTwopi`
  - ⬜ `LayoutCirco`
  - ⬜ `LayoutOsage`
  - ⬜ `LayoutPatchwork`
- ⬜ Create `render_test.go`
  - ⬜ `TestFormat_StringValues`
  - ⬜ `TestLayout_StringValues`

### Step 39: RenderError and Sentinel Errors

- ⬜ Create `errors.go`
- ⬜ Define `RenderError` struct
  - ⬜ `Err` field
  - ⬜ `Stderr` field
  - ⬜ `ExitCode` field
- ⬜ Implement `Error()` method
- ⬜ Implement `Unwrap()` method
- ⬜ Define sentinel errors
  - ⬜ `ErrGraphvizNotFound`
  - ⬜ `ErrInvalidDOT`
  - ⬜ `ErrRenderFailed`
- ⬜ Create `errors_test.go`
  - ⬜ `TestRenderError_Error_IncludesStderr`
  - ⬜ `TestRenderError_Unwrap`
  - ⬜ `TestRenderError_Is_RenderFailed`
  - ⬜ `TestSentinelErrors_Distinct`

### Step 40: Graphviz CLI Detection

- ⬜ Implement `findGraphviz(layout Layout)` function
  - ⬜ Use exec.LookPath
  - ⬜ Return full path or ErrGraphvizNotFound
- ⬜ Implement `GraphvizVersion()` function
  - ⬜ Run "dot -V"
  - ⬜ Parse and return version
- ⬜ Implement `checkGraphvizInstalled()` function
- ⬜ Create `requireGraphviz(t *testing.T)` test helper
- ⬜ Update `render_test.go`
  - ⬜ `TestFindGraphviz_Dot`
  - ⬜ `TestFindGraphviz_AllLayouts`
  - ⬜ `TestFindGraphviz_InvalidLayout`
  - ⬜ `TestGraphvizVersion_ReturnsVersion`

### Step 41: Graph.Render to io.Writer

- ⬜ Define `RenderOption` interface
- ⬜ Define `renderConfig` struct
  - ⬜ `layout` field (default: LayoutDot)
- ⬜ Implement `Render(format Format, w io.Writer, opts ...RenderOption)` method
  - ⬜ Build renderConfig from opts
  - ⬜ Find Graphviz binary
  - ⬜ Generate DOT string
  - ⬜ Execute command with stdin pipe
  - ⬜ Write stdout to w
  - ⬜ Handle errors with RenderError
- ⬜ Implement exec.Command handling
  - ⬜ Capture stdout and stderr
  - ⬜ Handle non-zero exit codes
- ⬜ Update `render_test.go`
  - ⬜ `TestGraph_Render_PNG_ProducesOutput`
  - ⬜ `TestGraph_Render_SVG_ProducesOutput`
  - ⬜ `TestGraph_Render_DOT_ProducesOutput`
  - ⬜ `TestGraph_Render_InvalidGraph_Error`
  - ⬜ `TestGraph_Render_ToBuffer`
- ⬜ Add validation helpers
  - ⬜ `assertValidPNG(t, data []byte)`
  - ⬜ `assertValidSVG(t, data []byte)`

### Step 42: RenderToFile and RenderBytes Conveniences

- ⬜ Implement `RenderToFile(format Format, path string, opts ...RenderOption)` method
  - ⬜ Create file
  - ⬜ Call Render with file as writer
  - ⬜ Close file
  - ⬜ Clean up on error
- ⬜ Implement `RenderBytes(format Format, opts ...RenderOption)` method
  - ⬜ Create bytes.Buffer
  - ⬜ Call Render with buffer
  - ⬜ Return buffer.Bytes()
- ⬜ Update `render_test.go`
  - ⬜ `TestGraph_RenderToFile_CreatesFile`
  - ⬜ `TestGraph_RenderToFile_ValidContent`
  - ⬜ `TestGraph_RenderToFile_ErrorCleansUp`
  - ⬜ `TestGraph_RenderBytes_ReturnsPNG`
  - ⬜ `TestGraph_RenderBytes_ReturnsSVG`
- ⬜ Create integration test
  - ⬜ `TestRender_CompleteWorkflow`

### Step 43: WithLayout Render Option

- ⬜ Implement `WithLayout(l Layout)` RenderOption
- ⬜ Verify all layout engines work
  - ⬜ dot
  - ⬜ neato
  - ⬜ fdp
  - ⬜ sfdp
  - ⬜ twopi
  - ⬜ circo
  - ⬜ osage
  - ⬜ patchwork
- ⬜ Update `render_test.go`
  - ⬜ `TestGraph_Render_WithLayout_Neato`
  - ⬜ `TestGraph_Render_WithLayout_Fdp`
  - ⬜ `TestGraph_Render_WithLayout_Circo`
  - ⬜ `TestGraph_Render_AllLayouts`
  - ⬜ `TestGraph_Render_DefaultLayout_IsDot`
- ⬜ Create final integration test
  - ⬜ `TestGoraffe_EndToEnd`
- ⬜ Complete documentation
  - ⬜ Update `doc.go` with complete overview
  - ⬜ Add `example_test.go`
    - ⬜ `Example_simpleGraph`
    - ⬜ `Example_withSubgraphs`
    - ⬜ `Example_htmlLabels`
    - ⬜ `Example_parseAndModify`

---

## Final Verification

### Code Quality

- ⬜ All tests passing
- ⬜ No race conditions (`go test -race`)
- ⬜ Linting passes (`golangci-lint run`)
- ⬜ No unused code
- ⬜ Consistent code formatting (`gofmt`)

### Documentation

- ⬜ Package documentation complete
- ⬜ All public types documented
- ⬜ All public functions documented
- ⬜ Examples for key functionality
- ⬜ README.md with usage examples

### Testing Coverage

- ⬜ Unit test coverage > 80%
- ⬜ Integration tests for parsing
- ⬜ Integration tests for rendering
- ⬜ Round-trip tests passing
- ⬜ Edge cases covered

### Files Created

- ✅ `go.mod`
- ✅ `doc.go`
- ✅ `graph.go`
- ✅ `node.go`
- ✅ `edge.go`
- ⬜ `subgraph.go`
- ✅ `node_attributes.go` (split from original attributes.go)
- ✅ `edge_attributes.go` (split from original attributes.go)
- ✅ `graph_attributes.go`
- ✅ `graph_options.go`
- ✅ `node_options.go`
- ✅ `edge_options.go`
- ⬜ `labels.go`
- ⬜ `port.go`
- ⬜ `dot.go`
- ⬜ `parse.go`
- ⬜ `render.go`
- ⬜ `errors.go`
- ✅ `graph_test.go`
- ✅ `node_test.go`
- ✅ `edge_test.go`
- ⬜ `subgraph_test.go`
- ✅ `node_attributes_test.go` (split from original attributes_test.go)
- ✅ `graph_attributes_test.go`
- ✅ `graph_options_test.go`
- ✅ `node_options_test.go`
- ✅ `edge_options_test.go`
- ⬜ `labels_test.go`
- ⬜ `record_labels_test.go`
- ⬜ `port_test.go`
- ⬜ `dot_test.go`
- ⬜ `lexer_test.go`
- ⬜ `parser_test.go`
- ⬜ `render_test.go`
- ⬜ `errors_test.go`
- ⬜ `example_test.go`
- ⬜ `testdata/simple.dot`
- ⬜ `testdata/complex.dot`
- ⬜ `testdata/cluster.dot`

---

## Progress Summary

| Phase | Steps | Completed | Percentage |
|-------|-------|-----------|------------|
| Foundation | 1-5 | 5/5 | 100% |
| Attributes | 6-14 | 7/9 | 78% |
| DOT Generation | 15-19 | 0/5 | 0% |
| Labels | 20-26 | 0/7 | 0% |
| Subgraphs | 27-32 | 0/6 | 0% |
| Parsing | 33-37 | 0/5 | 0% |
| Rendering | 38-43 | 0/6 | 0% |
| **Total** | **1-43** | **12/43** | **28%** |

---

## Notes

_Use this section to track blockers, decisions, or deviations from the plan._

### Blockers

- None yet

### Decisions Made

- **Step 12 - Pointer Fields for Attributes**: Decided to use pointer fields (*string,*float64, etc.) in GraphAttributes (and will refactor NodeAttributes/EdgeAttributes to match) to distinguish between "not set" vs "explicitly set to zero value". Public API uses getter methods that return zero values, with documentation noting the ambiguity. Internal DOT generation code can access pointer fields directly to check for nil.

### Deviations from Plan

- None yet

### Lessons Learned

- None yet
