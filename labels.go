package goraffe

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
