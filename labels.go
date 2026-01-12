package goraffe

type Alignment string

const (
	AlignLeft   Alignment = "left"
	AlignRight  Alignment = "right"
	AlignCenter Alignment = "center"
	AlignText   Alignment = "text"
)

type HTMLCell struct {
	content   string
	port      string
	bold      bool
	italic    bool
	underline bool
	colSpan   int
	rowSpan   int
	bgColor   string
	align     Alignment
}

type HTMLRow struct {
	cells []*HTMLCell
}

func Cell(content string) *HTMLCell {
	return &HTMLCell{
		content: content,
	}
}

func (c *HTMLCell) Port(port string) *HTMLCell {
	c.port = port
	return c
}

func (c *HTMLCell) Bold() *HTMLCell {
	c.bold = true
	return c
}

func (c *HTMLCell) Italic() *HTMLCell {
	c.italic = true
	return c
}

func (c *HTMLCell) Underline() *HTMLCell {
	c.underline = true
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
