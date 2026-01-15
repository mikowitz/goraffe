package goraffe

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestCell_Content verifies that a new cell stores and returns its content
func TestCell_Content(t *testing.T) {
	t.Run("stores simple text content", func(t *testing.T) {
		asrt := assert.New(t)

		cell := Cell(Text("Hello"))

		asrt.NotNil(cell, "expected Cell to return non-nil HTMLCell")
		// Content should be accessible via the cell (internal field, tested via behavior)
	})

	t.Run("stores empty content", func(t *testing.T) {
		asrt := assert.New(t)

		cell := Cell(Text(""))

		asrt.NotNil(cell, "expected Cell to return non-nil HTMLCell even with empty content")
	})

	t.Run("stores content with special characters", func(t *testing.T) {
		asrt := assert.New(t)

		cell := Cell(Text("<b>bold</b>"))

		asrt.NotNil(cell, "expected Cell to handle HTML-like content")
	})

	t.Run("stores content with spaces", func(t *testing.T) {
		asrt := assert.New(t)

		cell := Cell(Text("Multi word content"))

		asrt.NotNil(cell, "expected Cell to handle multi-word content")
	})
}

// TestCell_Chaining verifies that all chainable methods work and return the same instance
func TestCell_Chaining(t *testing.T) {
	t.Run("Port returns same cell instance", func(t *testing.T) {
		asrt := assert.New(t)

		cell := Cell(Text("test"))
		result := cell.Port("p1")

		asrt.Same(cell, result, "expected Port to return the same HTMLCell instance for chaining")
	})

	t.Run("ColSpan returns same cell instance", func(t *testing.T) {
		asrt := assert.New(t)

		cell := Cell(Text("test"))
		result := cell.ColSpan(2)

		asrt.Same(cell, result, "expected ColSpan to return the same HTMLCell instance for chaining")
	})

	t.Run("RowSpan returns same cell instance", func(t *testing.T) {
		asrt := assert.New(t)

		cell := Cell(Text("test"))
		result := cell.RowSpan(3)

		asrt.Same(cell, result, "expected RowSpan to return the same HTMLCell instance for chaining")
	})

	t.Run("BgColor returns same cell instance", func(t *testing.T) {
		asrt := assert.New(t)

		cell := Cell(Text("test"))
		result := cell.BgColor("lightblue")

		asrt.Same(cell, result, "expected BgColor to return the same HTMLCell instance for chaining")
	})

	t.Run("Align returns same cell instance", func(t *testing.T) {
		asrt := assert.New(t)

		cell := Cell(Text("test"))
		result := cell.Align(AlignCenter)

		asrt.Same(cell, result, "expected Align to return the same HTMLCell instance for chaining")
	})

	t.Run("chains multiple methods together", func(t *testing.T) {
		asrt := assert.New(t)

		cell := Cell(Text("test"))
		result := cell.Port("p1").ColSpan(2).BgColor("yellow")

		asrt.Same(cell, result, "expected all chained methods to return the same instance")
	})

	t.Run("chains all methods together", func(t *testing.T) {
		asrt := assert.New(t)

		cell := Cell(Text("test"))
		result := cell.
			Port("p1").
			ColSpan(2).
			RowSpan(3).
			BgColor("yellow").
			Align(AlignRight)

		asrt.Same(cell, result, "expected all chained methods to return the same instance")
	})
}

// TestCell_AllOptions verifies that a cell can have all options set
func TestCell_AllOptions(t *testing.T) {
	t.Run("sets port", func(t *testing.T) {
		asrt := assert.New(t)

		cell := Cell(Text("content")).Port("port1")

		asrt.NotNil(cell, "expected cell with port to be created")
	})

	t.Run("sets port with empty string", func(t *testing.T) {
		asrt := assert.New(t)

		cell := Cell(Text("content")).Port("")

		asrt.NotNil(cell, "expected cell with empty port to be created")
	})

	t.Run("sets colspan to positive value", func(t *testing.T) {
		asrt := assert.New(t)

		cell := Cell(Text("content")).ColSpan(3)

		asrt.NotNil(cell, "expected cell with colspan to be created")
	})

	t.Run("sets colspan to 1", func(t *testing.T) {
		asrt := assert.New(t)

		cell := Cell(Text("content")).ColSpan(1)

		asrt.NotNil(cell, "expected cell with colspan=1 to be created")
	})

	t.Run("sets rowspan to positive value", func(t *testing.T) {
		asrt := assert.New(t)

		cell := Cell(Text("content")).RowSpan(2)

		asrt.NotNil(cell, "expected cell with rowspan to be created")
	})

	t.Run("sets rowspan to 1", func(t *testing.T) {
		asrt := assert.New(t)

		cell := Cell(Text("content")).RowSpan(1)

		asrt.NotNil(cell, "expected cell with rowspan=1 to be created")
	})

	t.Run("sets bgcolor with color name", func(t *testing.T) {
		asrt := assert.New(t)

		cell := Cell(Text("content")).BgColor("red")

		asrt.NotNil(cell, "expected cell with bgcolor to be created")
	})

	t.Run("sets bgcolor with hex color", func(t *testing.T) {
		asrt := assert.New(t)

		cell := Cell(Text("content")).BgColor("#FF0000")

		asrt.NotNil(cell, "expected cell with hex bgcolor to be created")
	})

	t.Run("sets align to left", func(t *testing.T) {
		asrt := assert.New(t)

		cell := Cell(Text("content")).Align(AlignLeft)

		asrt.NotNil(cell, "expected cell with left align to be created")
	})

	t.Run("sets align to center", func(t *testing.T) {
		asrt := assert.New(t)

		cell := Cell(Text("content")).Align(AlignCenter)

		asrt.NotNil(cell, "expected cell with center align to be created")
	})

	t.Run("sets align to right", func(t *testing.T) {
		asrt := assert.New(t)

		cell := Cell(Text("content")).Align(AlignRight)

		asrt.NotNil(cell, "expected cell with right align to be created")
	})

	t.Run("sets all options together", func(t *testing.T) {
		asrt := assert.New(t)

		cell := Cell(Text("content").Bold().Italic().Underline()).
			Port("p1").
			ColSpan(2).
			RowSpan(3).
			BgColor("lightblue").
			Align(AlignCenter)

		asrt.NotNil(cell, "expected cell with all options to be created")
	})

	t.Run("overwrites port when set multiple times", func(t *testing.T) {
		asrt := assert.New(t)

		cell := Cell(Text("content")).Port("p1").Port("p2")

		asrt.NotNil(cell, "expected cell with overwritten port to be created")
		asrt.Equal(cell.port, "p2")
	})

	t.Run("overwrites colspan when set multiple times", func(t *testing.T) {
		asrt := assert.New(t)

		cell := Cell(Text("content")).ColSpan(2).ColSpan(5)

		asrt.NotNil(cell, "expected cell with overwritten colspan to be created")
		asrt.Equal(cell.colSpan, 5)
	})

	t.Run("overwrites rowspan when set multiple times", func(t *testing.T) {
		asrt := assert.New(t)

		cell := Cell(Text("content")).RowSpan(2).RowSpan(4)

		asrt.NotNil(cell, "expected cell with overwritten rowspan to be created")
		asrt.Equal(cell.rowSpan, 4)
	})

	t.Run("overwrites bgcolor when set multiple times", func(t *testing.T) {
		asrt := assert.New(t)

		cell := Cell(Text("content")).BgColor("red").BgColor("blue")

		asrt.NotNil(cell, "expected cell with overwritten bgcolor to be created")
		asrt.Equal(cell.bgColor, "blue")
	})

	t.Run("overwrites align when set multiple times", func(t *testing.T) {
		asrt := assert.New(t)

		cell := Cell(Text("content")).Align(AlignLeft).Align(AlignRight)

		asrt.NotNil(cell, "expected cell with overwritten align to be created")
		asrt.Equal(cell.align, AlignRight)
	})
}

// TestRow_ContainsCells verifies that Row stores cells
func TestRow_ContainsCells(t *testing.T) {
	t.Run("creates row with single cell", func(t *testing.T) {
		asrt := assert.New(t)

		cell := Cell(Text("A"))
		row := Row(cell)

		asrt.NotNil(row, "expected Row to return non-nil HTMLRow")
		cells := row.Cells()
		asrt.NotNil(cells, "expected Cells() to return non-nil slice")
		asrt.Len(cells, 1, "expected row to contain 1 cell")
		asrt.Same(cell, cells[0], "expected first cell to be the cell we passed")
	})

	t.Run("creates empty row with no cells", func(t *testing.T) {
		asrt := assert.New(t)

		row := Row()

		asrt.NotNil(row, "expected Row to return non-nil HTMLRow even with no cells")
		cells := row.Cells()
		asrt.NotNil(cells, "expected Cells() to return non-nil slice")
		asrt.Empty(cells, "expected row to contain no cells")
	})

	t.Run("Cells returns the actual cells", func(t *testing.T) {
		asrt := assert.New(t)

		cell1 := Cell(Text("A"))
		cell2 := Cell(Text("B"))
		row := Row(cell1, cell2)

		cells := row.Cells()
		asrt.Len(cells, 2, "expected row to contain 2 cells")
		asrt.Same(cell1, cells[0], "expected first cell to be cell1")
		asrt.Same(cell2, cells[1], "expected second cell to be cell2")
	})

	t.Run("Cells returns cells in order", func(t *testing.T) {
		asrt := assert.New(t)

		cell1 := Cell(Text("First"))
		cell2 := Cell(Text("Second"))
		cell3 := Cell(Text("Third"))
		row := Row(cell1, cell2, cell3)

		cells := row.Cells()
		asrt.Len(cells, 3, "expected row to contain 3 cells")
		asrt.Same(cell1, cells[0], "expected cells[0] to be first cell")
		asrt.Same(cell2, cells[1], "expected cells[1] to be second cell")
		asrt.Same(cell3, cells[2], "expected cells[2] to be third cell")
	})

	t.Run("row with cells that have different options", func(t *testing.T) {
		asrt := assert.New(t)

		cell1 := Cell(Text("A").Bold())
		cell2 := Cell(Text("B").Italic())
		cell3 := Cell(Text("C")).Port("p1")
		row := Row(cell1, cell2, cell3)

		cells := row.Cells()
		asrt.Len(cells, 3, "expected row to contain 3 cells")
		asrt.Same(cell1, cells[0], "expected first cell to be bold cell")
		asrt.Same(cell2, cells[1], "expected second cell to be italic cell")
		asrt.Same(cell3, cells[2], "expected third cell to be port cell")
	})
}

// TestRow_MultipleCells verifies various scenarios with multiple cells
func TestRow_MultipleCells(t *testing.T) {
	t.Run("creates row with two cells", func(t *testing.T) {
		asrt := assert.New(t)

		cell1 := Cell(Text("A"))
		cell2 := Cell(Text("B"))
		row := Row(cell1, cell2)

		asrt.NotNil(row, "expected Row to return non-nil HTMLRow")
		cells := row.Cells()
		asrt.Len(cells, 2, "expected row to contain 2 cells")
	})

	t.Run("creates row with three cells", func(t *testing.T) {
		asrt := assert.New(t)

		cell1 := Cell(Text("A"))
		cell2 := Cell(Text("B"))
		cell3 := Cell(Text("C"))
		row := Row(cell1, cell2, cell3)

		cells := row.Cells()
		asrt.Len(cells, 3, "expected row to contain 3 cells")
		asrt.Same(cell1, cells[0], "expected first cell")
		asrt.Same(cell2, cells[1], "expected second cell")
		asrt.Same(cell3, cells[2], "expected third cell")
	})

	t.Run("creates row with many cells", func(t *testing.T) {
		asrt := assert.New(t)

		cells := []*HTMLCell{
			Cell(Text("1")),
			Cell(Text("2")),
			Cell(Text("3")),
			Cell(Text("4")),
			Cell(Text("5")),
		}
		row := Row(cells...)

		returnedCells := row.Cells()
		asrt.Len(returnedCells, 5, "expected row to contain 5 cells")
		for i, cell := range cells {
			asrt.Same(cell, returnedCells[i], "expected cell %d to match", i)
		}
	})

	t.Run("row with cells with empty content", func(t *testing.T) {
		asrt := assert.New(t)

		cell1 := Cell(Text(""))
		cell2 := Cell(Text("B"))
		cell3 := Cell(Text(""))
		row := Row(cell1, cell2, cell3)

		cells := row.Cells()
		asrt.Len(cells, 3, "expected row to contain 3 cells including empty ones")
	})

	t.Run("row with cells containing special characters", func(t *testing.T) {
		asrt := assert.New(t)

		cell1 := Cell(Text("<html>"))
		cell2 := Cell(Text("A & B"))
		cell3 := Cell(Text("\"quoted\""))
		row := Row(cell1, cell2, cell3)

		cells := row.Cells()
		asrt.Len(cells, 3, "expected row to contain 3 cells with special characters")
	})

	t.Run("row with cells of different spans", func(t *testing.T) {
		asrt := assert.New(t)

		cell1 := Cell(Text("A")).ColSpan(2)
		cell2 := Cell(Text("B"))
		cell3 := Cell(Text("C")).RowSpan(3)
		row := Row(cell1, cell2, cell3)

		cells := row.Cells()
		asrt.Len(cells, 3, "expected row to contain 3 cells with different spans")
	})

	t.Run("row with cells with ports", func(t *testing.T) {
		asrt := assert.New(t)

		cell1 := Cell(Text("Input")).Port("in")
		cell2 := Cell(Text("Middle"))
		cell3 := Cell(Text("Output")).Port("out")
		row := Row(cell1, cell2, cell3)

		cells := row.Cells()
		asrt.Len(cells, 3, "expected row to contain 3 cells including those with ports")
	})

	t.Run("Cells called multiple times returns same cells", func(t *testing.T) {
		asrt := assert.New(t)

		cell1 := Cell(Text("A"))
		cell2 := Cell(Text("B"))
		row := Row(cell1, cell2)

		cells1 := row.Cells()
		cells2 := row.Cells()

		asrt.Len(cells1, 2, "expected first call to return 2 cells")
		asrt.Len(cells2, 2, "expected second call to return 2 cells")
		asrt.Same(cells1[0], cells2[0], "expected same cell instance on multiple calls")
		asrt.Same(cells1[1], cells2[1], "expected same cell instance on multiple calls")
	})

	t.Run("row preserves cell order", func(t *testing.T) {
		asrt := assert.New(t)

		// Create cells with content that would sort differently if we were doing that
		cellZ := Cell(Text("Z"))
		cellA := Cell(Text("A"))
		cellM := Cell(Text("M"))
		row := Row(cellZ, cellA, cellM)

		cells := row.Cells()
		asrt.Same(cellZ, cells[0], "expected first cell to be Z (insertion order)")
		asrt.Same(cellA, cells[1], "expected second cell to be A (insertion order)")
		asrt.Same(cellM, cells[2], "expected third cell to be M (insertion order)")
	})
}

// TestText_Creation verifies that Text creates TextContent correctly
func TestText_Creation(t *testing.T) {
	t.Run("creates text with content", func(t *testing.T) {
		asrt := assert.New(t)

		text := Text("Hello")

		asrt.NotNil(text, "expected Text to return non-nil TextContent")
	})

	t.Run("creates text with empty string", func(t *testing.T) {
		asrt := assert.New(t)

		text := Text("")

		asrt.NotNil(text, "expected Text to return non-nil TextContent even with empty string")
	})
}

// TestText_Bold verifies that Bold formatting works
func TestText_Bold(t *testing.T) {
	t.Run("sets bold flag", func(t *testing.T) {
		asrt := assert.New(t)

		text := Text("Hello").Bold()

		asrt.NotNil(text, "expected Bold to return non-nil TextContent")
	})

	t.Run("returns same instance for chaining", func(t *testing.T) {
		asrt := assert.New(t)

		text := Text("Hello")
		result := text.Bold()

		asrt.Same(text, result, "expected Bold to return the same TextContent instance")
	})
}

// TestText_Italic verifies that Italic formatting works
func TestText_Italic(t *testing.T) {
	t.Run("sets italic flag", func(t *testing.T) {
		asrt := assert.New(t)

		text := Text("Hello").Italic()

		asrt.NotNil(text, "expected Italic to return non-nil TextContent")
	})

	t.Run("returns same instance for chaining", func(t *testing.T) {
		asrt := assert.New(t)

		text := Text("Hello")
		result := text.Italic()

		asrt.Same(text, result, "expected Italic to return the same TextContent instance")
	})
}

// TestText_Underline verifies that Underline formatting works
func TestText_Underline(t *testing.T) {
	t.Run("sets underline flag", func(t *testing.T) {
		asrt := assert.New(t)

		text := Text("Hello").Underline()

		asrt.NotNil(text, "expected Underline to return non-nil TextContent")
	})

	t.Run("returns same instance for chaining", func(t *testing.T) {
		asrt := assert.New(t)

		text := Text("Hello")
		result := text.Underline()

		asrt.Same(text, result, "expected Underline to return the same TextContent instance")
	})
}

// TestText_CombinedFormatting verifies multiple formats on same text
func TestText_CombinedFormatting(t *testing.T) {
	t.Run("combines bold and italic", func(t *testing.T) {
		asrt := assert.New(t)

		text := Text("Hello").Bold().Italic()

		asrt.NotNil(text, "expected combined formatting to work")
	})

	t.Run("combines all three formats", func(t *testing.T) {
		asrt := assert.New(t)

		text := Text("Hello").Bold().Italic().Underline()

		asrt.NotNil(text, "expected all three formats to combine")
	})
}

// TestText_Chaining verifies that methods return self for chaining
func TestText_Chaining(t *testing.T) {
	t.Run("chains bold, italic, underline", func(t *testing.T) {
		asrt := assert.New(t)

		text := Text("Hello")
		result := text.Bold().Italic().Underline()

		asrt.Same(text, result, "expected chaining to return same instance")
	})

	t.Run("chains in different order", func(t *testing.T) {
		asrt := assert.New(t)

		text := Text("Hello")
		result := text.Underline().Bold().Italic()

		asrt.Same(text, result, "expected chaining in different order to work")
	})
}

// TestText_ToHTML verifies HTML output
func TestText_ToHTML(t *testing.T) {
	t.Run("plain text renders without tags", func(t *testing.T) {
		asrt := assert.New(t)

		text := Text("Hello")
		html := text.toHTML()

		asrt.Equal("Hello", html, "expected plain text without tags")
	})

	t.Run("bold text renders with B tag", func(t *testing.T) {
		asrt := assert.New(t)

		text := Text("Hello").Bold()
		html := text.toHTML()

		asrt.Equal("<b>Hello</b>", html, "expected bold text with B tags")
	})

	t.Run("italic text renders with I tag", func(t *testing.T) {
		asrt := assert.New(t)

		text := Text("Hello").Italic()
		html := text.toHTML()

		asrt.Equal("<i>Hello</i>", html, "expected italic text with I tags")
	})

	t.Run("underline text renders with U tag", func(t *testing.T) {
		asrt := assert.New(t)

		text := Text("Hello").Underline()
		html := text.toHTML()

		asrt.Equal("<u>Hello</u>", html, "expected underlined text with U tags")
	})

	t.Run("bold and italic renders nested B>I", func(t *testing.T) {
		asrt := assert.New(t)

		text := Text("Hello").Bold().Italic()
		html := text.toHTML()

		asrt.Equal("<b><i>Hello</i></b>", html, "expected bold and italic with nested tags")
	})

	t.Run("all three formats render nested B>I>U", func(t *testing.T) {
		asrt := assert.New(t)

		text := Text("Hello").Bold().Italic().Underline()
		html := text.toHTML()

		asrt.Equal("<b><i><u>Hello</u></i></b>", html, "expected all formats with nested tags")
	})
}

// TestBR_Creation verifies that BR creates LineBreak correctly
func TestBR_Creation(t *testing.T) {
	t.Run("creates line break", func(t *testing.T) {
		asrt := assert.New(t)

		br := BR()

		asrt.NotNil(br, "expected BR to return non-nil LineBreak")
	})
}

// TestBR_ToHTML verifies that BR renders to <BR/>
func TestBR_ToHTML(t *testing.T) {
	t.Run("renders to BR tag", func(t *testing.T) {
		asrt := assert.New(t)

		br := BR()
		html := br.toHTML()

		asrt.Equal("<br/>", html, "expected BR to render as <br/>")
	})
}

// TestHR_Creation verifies that HR creates HorizontalRule correctly
func TestHR_Creation(t *testing.T) {
	t.Run("creates horizontal rule", func(t *testing.T) {
		asrt := assert.New(t)

		hr := HR()

		asrt.NotNil(hr, "expected HR to return non-nil HorizontalRule")
	})
}

// TestHR_ToHTML verifies that HR renders to <HR/>
func TestHR_ToHTML(t *testing.T) {
	t.Run("renders to HR tag", func(t *testing.T) {
		asrt := assert.New(t)

		hr := HR()
		html := hr.toHTML()

		asrt.Equal("<hr/>", html, "expected HR to render as <hr/>")
	})
}

// TestCell_EmptyCell verifies that Cell() with no contents works
func TestCell_EmptyCell(t *testing.T) {
	t.Run("creates empty cell", func(t *testing.T) {
		asrt := assert.New(t)

		cell := Cell()

		asrt.NotNil(cell, "expected Cell with no contents to return non-nil HTMLCell")
	})
}

// TestCell_SingleText verifies Cell with single text content
func TestCell_SingleText(t *testing.T) {
	t.Run("creates cell with single text", func(t *testing.T) {
		asrt := assert.New(t)

		cell := Cell(Text("Hello"))

		asrt.NotNil(cell, "expected Cell with single text to work")
	})
}

// TestCell_MultipleContents verifies Cell with multiple content pieces
func TestCell_MultipleContents(t *testing.T) {
	t.Run("creates cell with two text contents", func(t *testing.T) {
		asrt := assert.New(t)

		cell := Cell(Text("A"), Text("B"))

		asrt.NotNil(cell, "expected Cell with multiple texts to work")
	})

	t.Run("creates cell with three text contents", func(t *testing.T) {
		asrt := assert.New(t)

		cell := Cell(Text("A"), Text("B"), Text("C"))

		asrt.NotNil(cell, "expected Cell with three texts to work")
	})
}

// TestCell_WithBR verifies Cell with line breaks
func TestCell_WithBR(t *testing.T) {
	t.Run("creates cell with text, BR, text", func(t *testing.T) {
		asrt := assert.New(t)

		cell := Cell(Text("A"), BR(), Text("B"))

		asrt.NotNil(cell, "expected Cell with BR to work")
	})

	t.Run("creates cell with multiple BRs", func(t *testing.T) {
		asrt := assert.New(t)

		cell := Cell(Text("Line1"), BR(), Text("Line2"), BR(), Text("Line3"))

		asrt.NotNil(cell, "expected Cell with multiple BRs to work")
	})
}

// TestCell_WithHR verifies Cell with horizontal rules
func TestCell_WithHR(t *testing.T) {
	t.Run("creates cell with text, HR, text", func(t *testing.T) {
		asrt := assert.New(t)

		cell := Cell(Text("Title"), HR(), Text("Body"))

		asrt.NotNil(cell, "expected Cell with HR to work")
	})

	t.Run("creates cell with multiple HRs", func(t *testing.T) {
		asrt := assert.New(t)

		cell := Cell(Text("Section1"), HR(), Text("Section2"), HR(), Text("Section3"))

		asrt.NotNil(cell, "expected Cell with multiple HRs to work")
	})
}

// TestCell_MixedFormatting verifies Cell with mixed formatted contents
func TestCell_MixedFormatting(t *testing.T) {
	t.Run("creates cell with bold and italic text", func(t *testing.T) {
		asrt := assert.New(t)

		cell := Cell(Text("Bold").Bold(), Text("Italic").Italic())

		asrt.NotNil(cell, "expected Cell with different formatted texts to work")
	})

	t.Run("creates cell with formatted text and BR", func(t *testing.T) {
		asrt := assert.New(t)

		cell := Cell(Text("Bold").Bold(), BR(), Text("Normal"))

		asrt.NotNil(cell, "expected Cell with formatted text and BR to work")
	})

	t.Run("creates cell with all content types", func(t *testing.T) {
		asrt := assert.New(t)

		cell := Cell(
			Text("Title").Bold().Underline(),
			HR(),
			Text("Line1").Italic(),
			BR(),
			Text("Line2"),
		)

		asrt.NotNil(cell, "expected Cell with all content types to work")
	})
}
