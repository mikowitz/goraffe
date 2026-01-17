package goraffe

import "strconv"

type Alignment string

const (
	AlignLeft   Alignment = "left"
	AlignRight  Alignment = "right"
	AlignCenter Alignment = "center"
	AlignText   Alignment = "text"
)

type HTMLCell struct {
	contents []Content
	port     string
	portRef  *Port
	colSpan  int
	rowSpan  int
	bgColor  string
	align    Alignment
}

type HTMLRow struct {
	cells []*HTMLCell
}

func Cell(contents ...Content) *HTMLCell {
	return &HTMLCell{
		contents: contents,
	}
}

func (c *HTMLCell) Port(port string) *HTMLCell {
	c.port = port
	c.portRef = &Port{id: port}
	return c
}

func (c *HTMLCell) ColSpan(span int) *HTMLCell {
	c.colSpan = span
	return c
}

func (c *HTMLCell) RowSpan(span int) *HTMLCell {
	c.rowSpan = span
	return c
}

func (c *HTMLCell) BgColor(color string) *HTMLCell {
	c.bgColor = color
	return c
}

func (c *HTMLCell) Align(align Alignment) *HTMLCell {
	c.align = align
	return c
}

// GetPort returns the Port reference associated with this cell, or nil if no port was set.
func (c *HTMLCell) GetPort() *Port {
	return c.portRef
}

func Row(cells ...*HTMLCell) *HTMLRow {
	return &HTMLRow{
		cells: cells,
	}
}

func (r *HTMLRow) Cells() []*HTMLCell {
	ret := make([]*HTMLCell, len(r.cells))
	copy(ret, r.cells)
	return ret
}

type HTMLLabel struct {
	rows        []*HTMLRow
	border      *int
	cellBorder  *int
	cellSpacing *int
	cellPadding *int
	bgColor     string
}

func HTMLTable(rows ...*HTMLRow) *HTMLLabel {
	return &HTMLLabel{
		rows: rows,
	}
}

func (l *HTMLLabel) Border(n int) *HTMLLabel {
	l.border = &n
	return l
}

func (l *HTMLLabel) CellBorder(n int) *HTMLLabel {
	l.cellBorder = &n
	return l
}

func (l *HTMLLabel) CellSpacing(n int) *HTMLLabel {
	l.cellSpacing = &n
	return l
}

func (l *HTMLLabel) CellPadding(n int) *HTMLLabel {
	l.cellPadding = &n
	return l
}

func (l *HTMLLabel) BgColor(color string) *HTMLLabel {
	l.bgColor = color
	return l
}

// setNodeContext is an internal method that associates all ports in this label with a node ID.
// This is called when the label is attached to a node.
func (l *HTMLLabel) setNodeContext(nodeID string) {
	for _, row := range l.rows {
		for _, cell := range row.cells {
			if cell.portRef != nil {
				cell.portRef.nodeID = nodeID
			}
		}
	}
}

func (l *HTMLLabel) String() string {
	result := "<"

	// Open table tag with attributes
	result += "<table"
	if l.border != nil {
		result += " border=\"" + strconv.Itoa(*l.border) + "\""
	}
	if l.cellBorder != nil {
		result += " cellborder=\"" + strconv.Itoa(*l.cellBorder) + "\""
	}
	if l.cellSpacing != nil {
		result += " cellspacing=\"" + strconv.Itoa(*l.cellSpacing) + "\""
	}
	if l.cellPadding != nil {
		result += " cellpadding=\"" + strconv.Itoa(*l.cellPadding) + "\""
	}
	if l.bgColor != "" {
		result += " bgcolor=\"" + l.bgColor + "\""
	}
	result += ">"

	// Add rows
	for _, row := range l.rows {
		result += "<tr>"
		for _, cell := range row.cells {
			result += "<td"
			if cell.port != "" {
				result += " port=\"" + cell.port + "\""
			}
			if cell.colSpan != 0 {
				result += " colspan=\"" + strconv.Itoa(cell.colSpan) + "\""
			}
			if cell.rowSpan != 0 {
				result += " rowspan=\"" + strconv.Itoa(cell.rowSpan) + "\""
			}
			if cell.bgColor != "" {
				result += " bgcolor=\"" + cell.bgColor + "\""
			}
			if cell.align != "" {
				result += " align=\"" + string(cell.align) + "\""
			}
			result += ">"

			// Add cell contents
			for _, content := range cell.contents {
				result += content.toHTML()
			}

			result += "</td>"
		}
		result += "</tr>"
	}

	// Close table tag
	result += "</table>"

	result += ">"
	return result
}

