# Goraffe: Go Library for Graphviz Graph Building and Rendering

## Overview

Goraffe is a Go library that provides a unified, ergonomic API for creating, parsing, and rendering Graphviz graphs. It combines the best aspects of existing Go Graphviz libraries while addressing their limitations.

### Core Value Proposition

- **Unified library**: Clean builder API, DOT parsing, and rendering in one package
- **Type-safe attributes**: Compile-time safety for common attributes with escape hatches for advanced use
- **Ergonomic complex graphs**: First-class support for subgraphs, clusters, rank constraints, HTML labels, and record shapes
- **Simple rendering**: Shell out to Graphviz CLI (system dependency acceptable)

### Module Path

```
github.com/<owner>/goraffe
```

---

## Architecture

### Package Structure

Single package design. All types, builders, and rendering functionality live in the root package:

```go
import "github.com/<owner>/goraffe"
```

**Rationale**: The workflow is linear (build graph → render), the API surface is cohesive, and splitting packages would add import friction without meaningful benefit.

### System Dependencies

- **Go**: 1.21+ (for modern generics and standard library features)
- **Graphviz**: Must be installed on the system (`dot`, `neato`, etc. available in PATH)

---

## Core Types

### Graph

```go
type Graph struct {
    name          string
    directed      bool
    strict        bool
    nodes         map[string]*Node
    edges         []*Edge
    subgraphs     []*Subgraph
    attrs         *GraphAttributes
    defaultNode   *NodeAttributes
    defaultEdge   *EdgeAttributes
    rankGroups    []rankGroup // for SameRank, MinRank, etc.
}

type GraphOption interface {
    applyGraph(*Graph)
}
```

**Creation**:

```go
// Directed graph (default layout: dot)
g := goraffe.NewGraph(goraffe.Directed)

// Undirected graph
g := goraffe.NewGraph(goraffe.Undirected)

// Strict directed graph (duplicate edges ignored)
g := goraffe.NewGraph(goraffe.Directed, goraffe.Strict)

// With default attributes
g := goraffe.NewGraph(
    goraffe.Directed,
    goraffe.WithDefaultNodeAttrs(goraffe.WithShape(goraffe.Box)),
    goraffe.WithDefaultEdgeAttrs(goraffe.WithColor(goraffe.Gray)),
    goraffe.WithRankDir(goraffe.LeftToRight),
    goraffe.WithLabel("My Graph"),
)
```

### Node

```go
type Node struct {
    id    string
    attrs *NodeAttributes
}

type NodeOption interface {
    applyNode(*NodeAttributes)
}

type NodeAttributes struct {
    Label     string
    Shape     Shape
    Color     string
    FillColor string
    FontName  string
    FontSize  float64
    Style     NodeStyle
    Width     float64
    Height    float64
    // ... additional typed fields
    custom    map[string]string // escape hatch
}

// NodeAttributes implements NodeOption
func (a NodeAttributes) applyNode(target *NodeAttributes)
```

**Creation and Usage**:

```go
// Create node with options
n := goraffe.NewNode("node_id", 
    goraffe.WithShape(goraffe.Box),
    goraffe.WithColor("red"),
    goraffe.WithLabel("My Node"),
)

// Add to graph
g.AddNode(n)

// Reusable attributes
commonAttrs := goraffe.NodeAttributes{
    Shape:    goraffe.Box,
    FontName: "Arial",
}
n1 := goraffe.NewNode("a", commonAttrs)
n2 := goraffe.NewNode("b", commonAttrs)
```

### Edge

```go
type Edge struct {
    from  *Node
    to    *Node
    attrs *EdgeAttributes
}

type EdgeOption interface {
    applyEdge(*EdgeAttributes)
}

type EdgeAttributes struct {
    Label     string
    Color     string
    Style     EdgeStyle
    ArrowHead ArrowType
    ArrowTail ArrowType
    Dir       EdgeDirection
    Weight    float64
    FromPort  *Port
    ToPort    *Port
    // ... additional typed fields
    custom    map[string]string // escape hatch
}

// EdgeAttributes implements EdgeOption
func (a EdgeAttributes) applyEdge(target *EdgeAttributes)
```

**Creation and Usage**:

```go
// Add edge (implicitly adds nodes if not present)
g.AddEdge(n1, n2)

// With options
g.AddEdge(n1, n2, 
    goraffe.WithLabel("connects"),
    goraffe.WithStyle(goraffe.Dashed),
    goraffe.WithColor("blue"),
)

// Reusable edge attributes
dashedGray := goraffe.EdgeAttributes{
    Style: goraffe.Dashed,
    Color: "gray",
}
g.AddEdge(n1, n2, dashedGray)
g.AddEdge(n2, n3, dashedGray)

// Override direction on specific edge
g.AddEdge(n1, n2, goraffe.Directed())   // force directed in undirected graph
g.AddEdge(n3, n4, goraffe.Undirected()) // force undirected in directed graph

// Port connections
g.AddEdge(n1, n2, 
    goraffe.FromPort(cell1.Port()),
    goraffe.ToPort(field2.Port()),
)
```

### Subgraph

```go
type Subgraph struct {
    name      string
    isCluster bool
    nodes     map[string]*Node
    edges     []*Edge
    subgraphs []*Subgraph
    attrs     *SubgraphAttributes
    rank      Rank // for rank constraints
}
```

**Creation and Usage**:

```go
// Cluster subgraph
g.Subgraph("cluster_db", func(s *goraffe.Subgraph) {
    s.SetLabel("Database Layer")
    s.SetStyle(goraffe.Filled)
    s.SetFillColor("lightgray")
    
    s.AddNode(db1)
    s.AddNode(db2)
    g.AddEdge(db1, db2) // edges can reference subgraph nodes
})

// Non-cluster subgraph
g.Subgraph("group1", func(s *goraffe.Subgraph) {
    s.AddNode(n1, n2, n3)
})

// Nested subgraphs
g.Subgraph("cluster_outer", func(outer *goraffe.Subgraph) {
    outer.Subgraph("cluster_inner", func(inner *goraffe.Subgraph) {
        inner.AddNode(n1)
    })
})
```

### Port

```go
type Port struct {
    id   string
    node *Node // back-reference for validation
}
```

Ports are not created directly; they are returned from HTML cells or record fields:

```go
cell := goraffe.Cell("input").Port("in")
port := cell.Port() // returns *Port

// Used in edge connections
g.AddEdge(n1, n2, goraffe.FromPort(port), goraffe.ToPort(otherPort))
```

---

## Type-Safe Attributes

### Enums and Constants

```go
// Shapes
type Shape string

const (
    Box       Shape = "box"
    Circle    Shape = "circle"
    Ellipse   Shape = "ellipse"
    Diamond   Shape = "diamond"
    Record    Shape = "record"
    MRecord   Shape = "Mrecord"
    Plaintext Shape = "plaintext"
    // ... expand to full Graphviz coverage over time
)

// Edge Styles
type EdgeStyle string

const (
    Solid  EdgeStyle = "solid"
    Dashed EdgeStyle = "dashed"
    Dotted EdgeStyle = "dotted"
    Bold   EdgeStyle = "bold"
)

// Node Styles
type NodeStyle string

const (
    StyleFilled  NodeStyle = "filled"
    StyleRounded NodeStyle = "rounded"
    StyleDashed  NodeStyle = "dashed"
    // ...
)

// Arrow Types
type ArrowType string

const (
    ArrowNormal ArrowType = "normal"
    ArrowDot    ArrowType = "dot"
    ArrowODot   ArrowType = "odot"
    ArrowNone   ArrowType = "none"
    ArrowVee    ArrowType = "vee"
    // ...
)

// Rank Direction
type RankDir string

const (
    TopToBottom RankDir = "TB"
    BottomToTop RankDir = "BT"
    LeftToRight RankDir = "LR"
    RightToLeft RankDir = "RL"
)

// Rank Constraints
type Rank string

const (
    SameRank   Rank = "same"
    MinRank    Rank = "min"
    MaxRank    Rank = "max"
    SourceRank Rank = "source"
    SinkRank   Rank = "sink"
)

// Spline Types
type SplineType string

const (
    SplineTrue     SplineType = "true"
    SplineFalse    SplineType = "false"
    SplineOrtho    SplineType = "ortho"
    SplinePolyline SplineType = "polyline"
    SplineCurved   SplineType = "curved"
)

// Edge Direction
type EdgeDirection string

const (
    DirForward EdgeDirection = "forward"
    DirBack    EdgeDirection = "back"
    DirBoth    EdgeDirection = "both"
    DirNone    EdgeDirection = "none"
)
```

### Functional Options

```go
// Node options
func WithShape(s Shape) NodeOption
func WithLabel(l string) NodeOption  // works for Node, Edge, Graph
func WithColor(c string) NodeOption
func WithFillColor(c string) NodeOption
func WithFontName(f string) NodeOption
func WithFontSize(s float64) NodeOption
func WithStyle(s NodeStyle) NodeOption
func WithWidth(w float64) NodeOption
func WithHeight(h float64) NodeOption

// Edge options
func WithEdgeLabel(l string) EdgeOption
func WithEdgeColor(c string) EdgeOption
func WithEdgeStyle(s EdgeStyle) EdgeOption
func WithArrowHead(a ArrowType) EdgeOption
func WithArrowTail(a ArrowType) EdgeOption
func WithWeight(w float64) EdgeOption
func FromPort(p *Port) EdgeOption
func ToPort(p *Port) EdgeOption
func Directed() EdgeOption
func Undirected() EdgeOption

// Graph options
func WithRankDir(d RankDir) GraphOption
func WithGraphLabel(l string) GraphOption
func WithBgColor(c string) GraphOption
func WithFontName(f string) GraphOption
func WithFontSize(s float64) GraphOption
func WithSplines(s SplineType) GraphOption
func WithNodeSep(n float64) GraphOption
func WithRankSep(r float64) GraphOption
func WithCompound(b bool) GraphOption
func WithDefaultNodeAttrs(opts ...NodeOption) GraphOption
func WithDefaultEdgeAttrs(opts ...EdgeOption) GraphOption

// Escape hatch for arbitrary attributes (all types)
func WithAttribute(key, value string) NodeOption  // also EdgeOption, GraphOption
```

---

## Graph Attributes

### Initial Supported Attributes

| Attribute | Type | Description |
|-----------|------|-------------|
| `rankdir` | RankDir | Direction of graph layout (TB, LR, BT, RL) |
| `label` | string | Graph title |
| `bgcolor` | string | Background color |
| `fontname` | string | Default font family |
| `fontsize` | float64 | Default font size |
| `splines` | SplineType | Edge routing style |
| `nodesep` | float64 | Horizontal spacing between nodes |
| `ranksep` | float64 | Vertical spacing between ranks |
| `compound` | bool | Allow edges between clusters |

### Escape Hatch

```go
g := goraffe.NewGraph(
    goraffe.Directed,
    goraffe.WithAttribute("ratio", "fill"),
    goraffe.WithAttribute("concentrate", "true"),
)
```

---

## Rank Constraints

### Convenience Methods

```go
// Place nodes on same rank
g.SameRank(n1, n2, n3)

// Place node(s) at minimum rank (top in TB)
g.MinRank(n1)

// Place node(s) at maximum rank (bottom in TB)
g.MaxRank(n2)

// Source/sink ranks
g.SourceRank(startNode)
g.SinkRank(endNode)
```

### Via Subgraph

```go
g.Subgraph("", func(s *goraffe.Subgraph) {
    s.SetRank(goraffe.SameRank)
    s.AddNode(n1, n2, n3)
})
```

---

## Labels

### Simple Labels

```go
n := goraffe.NewNode("id", goraffe.WithLabel("Display Name"))
g.AddEdge(n1, n2, goraffe.WithLabel("relationship"))
```

### HTML Labels

**Builder API**:

```go
// Simple table
label := goraffe.HTMLTable(
    goraffe.Row(
        goraffe.Cell("Header 1").Bold(),
        goraffe.Cell("Header 2").Bold(),
    ),
    goraffe.Row(
        goraffe.Cell("Value 1"),
        goraffe.Cell("Value 2"),
    ),
)
n := goraffe.NewNode("id", goraffe.WithHTMLLabel(label))

// Building rows separately
headerRow := goraffe.Row(
    goraffe.Cell("Name").Bold().BgColor("lightgray"),
    goraffe.Cell("Type").Bold().BgColor("lightgray"),
)

var dataRows []goraffe.HTMLRow
for _, item := range items {
    dataRows = append(dataRows, goraffe.Row(
        goraffe.Cell(item.Name),
        goraffe.Cell(item.Type),
    ))
}

label := goraffe.HTMLTable(append([]goraffe.HTMLRow{headerRow}, dataRows...)...)
```

**Cell Options**:

```go
type HTMLCell struct { ... }

func Cell(content string) *HTMLCell
func (c *HTMLCell) Port(id string) *HTMLCell  // returns self for chaining
func (c *HTMLCell) Bold() *HTMLCell
func (c *HTMLCell) Italic() *HTMLCell
func (c *HTMLCell) Underline() *HTMLCell
func (c *HTMLCell) ColSpan(n int) *HTMLCell
func (c *HTMLCell) RowSpan(n int) *HTMLCell
func (c *HTMLCell) BgColor(color string) *HTMLCell
func (c *HTMLCell) Align(a Alignment) *HTMLCell
func (c *HTMLCell) Port() *Port  // get the port reference
```

**Table Options**:

```go
func HTMLTable(rows ...HTMLRow) *HTMLLabel
func (t *HTMLLabel) Border(n int) *HTMLLabel
func (t *HTMLLabel) CellBorder(n int) *HTMLLabel
func (t *HTMLLabel) CellSpacing(n int) *HTMLLabel
func (t *HTMLLabel) CellPadding(n int) *HTMLLabel
func (t *HTMLLabel) BgColor(color string) *HTMLLabel
```

**Raw HTML Escape Hatch**:

```go
n := goraffe.NewNode("id", goraffe.WithRawHTMLLabel(`
    <TABLE BORDER="0">
        <TR><TD PORT="in">input</TD></TR>
        <TR><TD PORT="out">output</TD></TR>
    </TABLE>
`))
```

### Record Labels

```go
// Simple record
n := goraffe.NewNode("id",
    goraffe.WithShape(goraffe.Record),
    goraffe.WithRecordLabel(
        goraffe.Field("field1"),
        goraffe.Field("field2"),
        goraffe.Field("field3"),
    ),
)

// With ports
f1 := goraffe.Field("input").Port("in")
f2 := goraffe.Field("output").Port("out")
n := goraffe.NewNode("id",
    goraffe.WithShape(goraffe.Record),
    goraffe.WithRecordLabel(f1, f2),
)

// Later, use ports
g.AddEdge(other, n, goraffe.ToPort(f1.Port()))

// Nested groups (vertical | horizontal alternation)
n := goraffe.NewNode("id",
    goraffe.WithShape(goraffe.Record),
    goraffe.WithRecordLabel(
        goraffe.Field("left"),
        goraffe.FieldGroup(
            goraffe.Field("top"),
            goraffe.Field("bottom"),
        ),
        goraffe.Field("right"),
    ),
)
// Produces: left | { top | bottom } | right
```

---

## DOT Parsing

### API

```go
// Parse DOT from various sources
func Parse(r io.Reader) (*Graph, error)
func ParseString(dot string) (*Graph, error)
func ParseFile(path string) (*Graph, error)
```

### Behavior

- **Semantic parsing**: Parses DOT into Graph/Node/Edge structures
- **Not round-trip preserving**: Re-serialization produces normalized output; comments and original formatting are lost
- **Attribute handling**: Known attributes mapped to typed fields; unknown attributes preserved in custom map

### Example

```go
g, err := goraffe.ParseString(`
    digraph G {
        rankdir=LR;
        node [shape=box];
        A -> B -> C;
        B -> D;
    }
`)
if err != nil {
    log.Fatal(err)
}

// Modify the parsed graph
g.AddNode(goraffe.NewNode("E", goraffe.WithColor("red")))
g.AddEdge(g.GetNode("C"), g.GetNode("E"))

// Render
g.RenderToFile(goraffe.PNG, "output.png")
```

### Graph Inspection

```go
func (g *Graph) GetNode(id string) *Node           // nil if not found
func (g *Graph) Nodes() []*Node                    // all nodes
func (g *Graph) Edges() []*Edge                    // all edges
func (g *Graph) Subgraphs() []*Subgraph            // top-level subgraphs
func (g *Graph) IsDirected() bool
func (g *Graph) IsStrict() bool
```

---

## Rendering

### Output Formats

```go
type Format string

const (
    PNG Format = "png"
    SVG Format = "svg"
    PDF Format = "pdf"
    DOT Format = "dot"
)
```

### Layout Engines

```go
type Layout string

const (
    LayoutDot      Layout = "dot"       // hierarchical (default)
    LayoutNeato    Layout = "neato"     // spring model
    LayoutFdp      Layout = "fdp"       // force-directed
    LayoutSfdp     Layout = "sfdp"      // scalable force-directed
    LayoutTwopi    Layout = "twopi"     // radial
    LayoutCirco    Layout = "circo"     // circular
    LayoutOsage    Layout = "osage"     // clustered
    LayoutPatchwork Layout = "patchwork" // squarified treemap
)
```

### Render Options

```go
type RenderOption interface {
    applyRender(*renderConfig)
}

func WithLayout(l Layout) RenderOption
```

### Render Methods

```go
// Core method: render to io.Writer
func (g *Graph) Render(format Format, w io.Writer, opts ...RenderOption) error

// Convenience: render to file
func (g *Graph) RenderToFile(format Format, path string, opts ...RenderOption) error

// Convenience: render to bytes
func (g *Graph) RenderBytes(format Format, opts ...RenderOption) ([]byte, error)
```

### Examples

```go
// Render to file with default layout (dot)
err := g.RenderToFile(goraffe.PNG, "graph.png")

// Render to file with specific layout
err := g.RenderToFile(goraffe.SVG, "graph.svg", goraffe.WithLayout(goraffe.Neato))

// Render to writer
var buf bytes.Buffer
err := g.Render(goraffe.PDF, &buf)

// Render to bytes
data, err := g.RenderBytes(goraffe.PNG, goraffe.WithLayout(goraffe.Circo))
```

### DOT Output

```go
// Get DOT string representation
func (g *Graph) String() string

// Write DOT to writer
func (g *Graph) WriteDOT(w io.Writer) error
```

---

## Error Handling

### Error Types

```go
// Base error type wrapping Graphviz CLI errors
type RenderError struct {
    Err      error  // underlying error
    Stderr   string // Graphviz stderr output
    ExitCode int    // process exit code
}

func (e *RenderError) Error() string
func (e *RenderError) Unwrap() error

// Sentinel errors
var (
    ErrGraphvizNotFound = errors.New("goraffe: graphviz not found in PATH")
    ErrInvalidDOT       = errors.New("goraffe: invalid DOT syntax")
    ErrRenderFailed     = errors.New("goraffe: rendering failed")
)
```

### Usage

```go
err := g.RenderToFile(goraffe.PNG, "out.png")
if err != nil {
    if errors.Is(err, goraffe.ErrGraphvizNotFound) {
        log.Fatal("Please install Graphviz: brew install graphviz")
    }
    if renderErr, ok := err.(*goraffe.RenderError); ok {
        log.Printf("Graphviz error: %s", renderErr.Stderr)
    }
    log.Fatal(err)
}
```

---

## Internal Implementation

### DOT Generation

Internal method to generate DOT string from graph structure:

```go
func (g *Graph) generateDOT() string
```

Must handle:
- Graph type (`digraph` vs `graph`, `strict` prefix)
- Graph attributes
- Default node/edge attributes
- All nodes with their attributes
- All edges with their attributes
- Subgraphs (recursive)
- Rank constraint subgraphs
- Proper escaping of strings and special characters
- HTML labels (no escaping, wrapped in `< >`)
- Record labels (proper escaping of `|`, `{`, `}`, `<`, `>`)

### CLI Invocation

```go
func (g *Graph) render(format Format, layout Layout, w io.Writer) error {
    // 1. Generate DOT
    dot := g.generateDOT()
    
    // 2. Find Graphviz binary
    binary := string(layout)  // "dot", "neato", etc.
    path, err := exec.LookPath(binary)
    if err != nil {
        return ErrGraphvizNotFound
    }
    
    // 3. Execute: echo $DOT | dot -T$FORMAT
    cmd := exec.Command(path, "-T"+string(format))
    cmd.Stdin = strings.NewReader(dot)
    cmd.Stdout = w
    
    var stderr bytes.Buffer
    cmd.Stderr = &stderr
    
    // 4. Run and handle errors
    if err := cmd.Run(); err != nil {
        return &RenderError{
            Err:      err,
            Stderr:   stderr.String(),
            ExitCode: cmd.ProcessState.ExitCode(),
        }
    }
    return nil
}
```

### DOT Parsing

Use a proper parser (consider adapting from `awalterschulze/gographviz` or writing a new one with a parser generator). The parser should:

1. Tokenize DOT input
2. Build AST
3. Analyze AST into Graph structures
4. Map known attributes to typed fields
5. Preserve unknown attributes in custom maps

---

## Testing Plan

### Unit Tests

**Graph Construction**:
- Create directed/undirected graphs
- Create strict graphs
- Set graph attributes
- Set default node/edge attributes

**Node Operations**:
- Create nodes with various options
- Add nodes to graph
- Verify duplicate node handling
- Node attribute merging (options + reusable structs)

**Edge Operations**:
- Create edges between nodes
- Implicit node addition
- Edge with ports
- Direction override
- Attribute merging

**Subgraph Operations**:
- Create clusters
- Create non-cluster subgraphs
- Nested subgraphs
- Subgraph attributes
- Rank constraints via subgraph

**Rank Constraints**:
- SameRank, MinRank, MaxRank
- SourceRank, SinkRank
- Multiple rank groups

**Labels**:
- Simple string labels
- HTML table construction
- Cell options (port, colspan, styling)
- Record label construction
- Nested record groups
- Raw HTML escape hatch

**DOT Generation**:
- Simple graph output
- All attribute types
- Proper escaping
- HTML label output (no escaping)
- Record label output
- Subgraph output
- Rank constraint output

### Integration Tests

**DOT Parsing**:
- Parse simple graphs
- Parse complex graphs with subgraphs
- Parse graphs with HTML labels
- Parse graphs with record labels
- Round-trip: generate → parse → generate (semantic equivalence)

**Rendering** (requires Graphviz installed):
- Render to each format (PNG, SVG, PDF, DOT)
- Render with each layout engine
- Verify output is valid (non-empty, correct format magic bytes)
- Error handling for invalid graphs

### Test Fixtures

Create a `testdata/` directory with:
- Sample DOT files for parsing tests
- Expected DOT output for generation tests
- Golden files for rendering tests (optional, can be fragile)

### Test Helpers

```go
// Helper to compare DOT output (normalize whitespace)
func assertDOTEqual(t *testing.T, expected, actual string)

// Helper to verify rendered output format
func assertValidPNG(t *testing.T, data []byte)
func assertValidSVG(t *testing.T, data []byte)

// Helper to skip tests if Graphviz not installed
func requireGraphviz(t *testing.T)
```

### Example Test

```go
func TestSimpleGraph(t *testing.T) {
    g := goraffe.NewGraph(goraffe.Directed)
    
    n1 := goraffe.NewNode("A", goraffe.WithShape(goraffe.Box))
    n2 := goraffe.NewNode("B", goraffe.WithShape(goraffe.Circle))
    
    g.AddNode(n1)
    g.AddEdge(n1, n2, goraffe.WithLabel("connects"))
    
    expected := `digraph {
    A [shape="box"];
    B [shape="circle"];
    A -> B [label="connects"];
}`
    assertDOTEqual(t, expected, g.String())
}

func TestRenderPNG(t *testing.T) {
    requireGraphviz(t)
    
    g := goraffe.NewGraph(goraffe.Directed)
    g.AddEdge(
        goraffe.NewNode("A"),
        goraffe.NewNode("B"),
    )
    
    data, err := g.RenderBytes(goraffe.PNG)
    require.NoError(t, err)
    assertValidPNG(t, data)
}
```

---

## Project Structure

```
goraffe/
├── go.mod
├── go.sum
├── README.md
├── LICENSE
├── doc.go              # Package documentation
├── graph.go            # Graph type and methods
├── node.go             # Node type and options
├── edge.go             # Edge type and options
├── subgraph.go         # Subgraph type and methods
├── attributes.go       # Attribute types, enums, constants
├── options.go          # Functional option implementations
├── labels.go           # HTML and record label builders
├── port.go             # Port type
├── dot.go              # DOT generation
├── parse.go            # DOT parsing
├── render.go           # Rendering via Graphviz CLI
├── errors.go           # Error types
└── testdata/
    ├── simple.dot
    ├── complex.dot
    └── ...
```

---

## Future Considerations (Out of Scope for v1)

- **Graph manipulation**: Traverse, query, modify graphs (find edges from node, remove node and edges, etc.)
- **Pure Go rendering**: Implement layout algorithms natively (massive undertaking)
- **Additional output formats**: JPG, GIF, PS, EPS, JSON, etc.
- **Validation**: Pre-render validation of graph structure
- **Full attribute coverage**: Expand typed attributes to full Graphviz parity
- **Graph algorithms**: Integration with gonum/graph for analysis

---

## Appendix: Quick Reference

### Minimal Example

```go
package main

import (
    "log"
    "github.com/<owner>/goraffe"
)

func main() {
    g := goraffe.NewGraph(goraffe.Directed)
    
    a := goraffe.NewNode("A", goraffe.WithLabel("Start"))
    b := goraffe.NewNode("B", goraffe.WithLabel("End"))
    
    g.AddEdge(a, b)
    
    if err := g.RenderToFile(goraffe.PNG, "graph.png"); err != nil {
        log.Fatal(err)
    }
}
```

### Complex Example

```go
package main

import (
    "log"
    "github.com/<owner>/goraffe"
)

func main() {
    g := goraffe.NewGraph(
        goraffe.Directed,
        goraffe.WithRankDir(goraffe.LeftToRight),
        goraffe.WithDefaultNodeAttrs(
            goraffe.WithShape(goraffe.Box),
            goraffe.WithFontName("Arial"),
        ),
    )
    
    // Create nodes with HTML labels
    clientLabel := goraffe.HTMLTable(
        goraffe.Row(goraffe.Cell("Client").Bold()),
        goraffe.Row(goraffe.Cell("Browser").Port("out")),
    )
    client := goraffe.NewNode("client", goraffe.WithHTMLLabel(clientLabel))
    
    serverLabel := goraffe.HTMLTable(
        goraffe.Row(goraffe.Cell("Server").Bold()),
        goraffe.Row(
            goraffe.Cell("HTTP").Port("http"),
            goraffe.Cell("WS").Port("ws"),
        ),
    )
    server := goraffe.NewNode("server", goraffe.WithHTMLLabel(serverLabel))
    
    db := goraffe.NewNode("db",
        goraffe.WithShape(goraffe.Cylinder),
        goraffe.WithLabel("PostgreSQL"),
    )
    
    // Create cluster
    g.Subgraph("cluster_backend", func(s *goraffe.Subgraph) {
        s.SetLabel("Backend Services")
        s.SetStyle(goraffe.StyleFilled)
        s.SetFillColor("lightgray")
        s.AddNode(server, db)
    })
    
    g.AddNode(client)
    
    // Connect with ports
    g.AddEdge(client, server,
        goraffe.FromPort(clientLabel.Rows[1].Cells[0].Port()),
        goraffe.ToPort(serverLabel.Rows[1].Cells[0].Port()),
        goraffe.WithLabel("HTTPS"),
    )
    g.AddEdge(server, db, goraffe.WithStyle(goraffe.Dashed))
    
    // Render
    if err := g.RenderToFile(goraffe.SVG, "architecture.svg"); err != nil {
        log.Fatal(err)
    }
}
```
