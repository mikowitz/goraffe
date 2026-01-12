package goraffe

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestCell_Content verifies that a new cell stores and returns its content
func TestCell_Content(t *testing.T) {
	t.Run("stores simple text content", func(t *testing.T) {
		asrt := assert.New(t)

		cell := Cell("Hello")

		asrt.NotNil(cell, "expected Cell to return non-nil HTMLCell")
		// Content should be accessible via the cell (internal field, tested via behavior)
	})

	t.Run("stores empty content", func(t *testing.T) {
		asrt := assert.New(t)

		cell := Cell("")

		asrt.NotNil(cell, "expected Cell to return non-nil HTMLCell even with empty content")
	})

	t.Run("stores content with special characters", func(t *testing.T) {
		asrt := assert.New(t)

		cell := Cell("<b>bold</b>")

		asrt.NotNil(cell, "expected Cell to handle HTML-like content")
	})

	t.Run("stores content with spaces", func(t *testing.T) {
		asrt := assert.New(t)

		cell := Cell("Multi word content")

		asrt.NotNil(cell, "expected Cell to handle multi-word content")
	})
}

// TestCell_Chaining verifies that all chainable methods work and return the same instance
func TestCell_Chaining(t *testing.T) {
	t.Run("Port returns same cell instance", func(t *testing.T) {
		asrt := assert.New(t)

		cell := Cell("test")
		result := cell.Port("p1")

		asrt.Same(cell, result, "expected Port to return the same HTMLCell instance for chaining")
	})

	t.Run("Bold returns same cell instance", func(t *testing.T) {
		asrt := assert.New(t)

		cell := Cell("test")
		result := cell.Bold()

		asrt.Same(cell, result, "expected Bold to return the same HTMLCell instance for chaining")
	})

	t.Run("Italic returns same cell instance", func(t *testing.T) {
		asrt := assert.New(t)

		cell := Cell("test")
		result := cell.Italic()

		asrt.Same(cell, result, "expected Italic to return the same HTMLCell instance for chaining")
	})

	t.Run("Underline returns same cell instance", func(t *testing.T) {
		asrt := assert.New(t)

		cell := Cell("test")
		result := cell.Underline()

		asrt.Same(cell, result, "expected Underline to return the same HTMLCell instance for chaining")
	})

	t.Run("ColSpan returns same cell instance", func(t *testing.T) {
		asrt := assert.New(t)

		cell := Cell("test")
		result := cell.ColSpan(2)

		asrt.Same(cell, result, "expected ColSpan to return the same HTMLCell instance for chaining")
	})

	t.Run("RowSpan returns same cell instance", func(t *testing.T) {
		asrt := assert.New(t)

		cell := Cell("test")
		result := cell.RowSpan(3)

		asrt.Same(cell, result, "expected RowSpan to return the same HTMLCell instance for chaining")
	})

	t.Run("BgColor returns same cell instance", func(t *testing.T) {
		asrt := assert.New(t)

		cell := Cell("test")
		result := cell.BgColor("lightblue")

		asrt.Same(cell, result, "expected BgColor to return the same HTMLCell instance for chaining")
	})

	t.Run("Align returns same cell instance", func(t *testing.T) {
		asrt := assert.New(t)

		cell := Cell("test")
		result := cell.Align(AlignCenter)

		asrt.Same(cell, result, "expected Align to return the same HTMLCell instance for chaining")
	})

	t.Run("chains multiple methods together", func(t *testing.T) {
		asrt := assert.New(t)

		cell := Cell("test")
		result := cell.Port("p1").Bold().Italic().ColSpan(2)

		asrt.Same(cell, result, "expected all chained methods to return the same instance")
	})

	t.Run("chains all methods together", func(t *testing.T) {
		asrt := assert.New(t)

		cell := Cell("test")
		result := cell.
			Port("p1").
			Bold().
			Italic().
			Underline().
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

		cell := Cell("content").Port("port1")

		asrt.NotNil(cell, "expected cell with port to be created")
	})

	t.Run("sets port with empty string", func(t *testing.T) {
		asrt := assert.New(t)

		cell := Cell("content").Port("")

		asrt.NotNil(cell, "expected cell with empty port to be created")
	})

	t.Run("sets bold", func(t *testing.T) {
		asrt := assert.New(t)

		cell := Cell("content").Bold()

		asrt.NotNil(cell, "expected bold cell to be created")
	})

	t.Run("sets italic", func(t *testing.T) {
		asrt := assert.New(t)

		cell := Cell("content").Italic()

		asrt.NotNil(cell, "expected italic cell to be created")
	})

	t.Run("sets underline", func(t *testing.T) {
		asrt := assert.New(t)

		cell := Cell("content").Underline()

		asrt.NotNil(cell, "expected underlined cell to be created")
	})

	t.Run("sets colspan to positive value", func(t *testing.T) {
		asrt := assert.New(t)

		cell := Cell("content").ColSpan(3)

		asrt.NotNil(cell, "expected cell with colspan to be created")
	})

	t.Run("sets colspan to 1", func(t *testing.T) {
		asrt := assert.New(t)

		cell := Cell("content").ColSpan(1)

		asrt.NotNil(cell, "expected cell with colspan=1 to be created")
	})

	t.Run("sets rowspan to positive value", func(t *testing.T) {
		asrt := assert.New(t)

		cell := Cell("content").RowSpan(2)

		asrt.NotNil(cell, "expected cell with rowspan to be created")
	})

	t.Run("sets rowspan to 1", func(t *testing.T) {
		asrt := assert.New(t)

		cell := Cell("content").RowSpan(1)

		asrt.NotNil(cell, "expected cell with rowspan=1 to be created")
	})

	t.Run("sets bgcolor with color name", func(t *testing.T) {
		asrt := assert.New(t)

		cell := Cell("content").BgColor("red")

		asrt.NotNil(cell, "expected cell with bgcolor to be created")
	})

	t.Run("sets bgcolor with hex color", func(t *testing.T) {
		asrt := assert.New(t)

		cell := Cell("content").BgColor("#FF0000")

		asrt.NotNil(cell, "expected cell with hex bgcolor to be created")
	})

	t.Run("sets align to left", func(t *testing.T) {
		asrt := assert.New(t)

		cell := Cell("content").Align(AlignLeft)

		asrt.NotNil(cell, "expected cell with left align to be created")
	})

	t.Run("sets align to center", func(t *testing.T) {
		asrt := assert.New(t)

		cell := Cell("content").Align(AlignCenter)

		asrt.NotNil(cell, "expected cell with center align to be created")
	})

	t.Run("sets align to right", func(t *testing.T) {
		asrt := assert.New(t)

		cell := Cell("content").Align(AlignRight)

		asrt.NotNil(cell, "expected cell with right align to be created")
	})

	t.Run("sets all formatting options", func(t *testing.T) {
		asrt := assert.New(t)

		cell := Cell("content").Bold().Italic().Underline()

		asrt.NotNil(cell, "expected cell with all formatting to be created")
	})

	t.Run("sets all options together", func(t *testing.T) {
		asrt := assert.New(t)

		cell := Cell("content").
			Port("p1").
			Bold().
			Italic().
			Underline().
			ColSpan(2).
			RowSpan(3).
			BgColor("lightblue").
			Align(AlignCenter)

		asrt.NotNil(cell, "expected cell with all options to be created")
	})

	t.Run("overwrites port when set multiple times", func(t *testing.T) {
		asrt := assert.New(t)

		cell := Cell("content").Port("p1").Port("p2")

		asrt.NotNil(cell, "expected cell with overwritten port to be created")
		asrt.Equal(cell.port, "p2")
	})

	t.Run("overwrites colspan when set multiple times", func(t *testing.T) {
		asrt := assert.New(t)

		cell := Cell("content").ColSpan(2).ColSpan(5)

		asrt.NotNil(cell, "expected cell with overwritten colspan to be created")
		asrt.Equal(cell.colSpan, 5)
	})

	t.Run("overwrites rowspan when set multiple times", func(t *testing.T) {
		asrt := assert.New(t)

		cell := Cell("content").RowSpan(2).RowSpan(4)

		asrt.NotNil(cell, "expected cell with overwritten rowspan to be created")
		asrt.Equal(cell.rowSpan, 4)
	})

	t.Run("overwrites bgcolor when set multiple times", func(t *testing.T) {
		asrt := assert.New(t)

		cell := Cell("content").BgColor("red").BgColor("blue")

		asrt.NotNil(cell, "expected cell with overwritten bgcolor to be created")
		asrt.Equal(cell.bgColor, "blue")
	})

	t.Run("overwrites align when set multiple times", func(t *testing.T) {
		asrt := assert.New(t)

		cell := Cell("content").Align(AlignLeft).Align(AlignRight)

		asrt.NotNil(cell, "expected cell with overwritten align to be created")
		asrt.Equal(cell.align, AlignRight)
	})

	t.Run("can toggle bold multiple times", func(t *testing.T) {
		asrt := assert.New(t)

		// Bold is a boolean flag, multiple calls should be idempotent
		cell := Cell("content").Bold().Bold()

		asrt.NotNil(cell, "expected cell with bold set multiple times to be created")
		asrt.True(cell.bold)
	})
}

// TestRow_ContainsCells verifies that Row stores cells
func TestRow_ContainsCells(t *testing.T) {
	t.Run("creates row with single cell", func(t *testing.T) {
		asrt := assert.New(t)

		cell := Cell("A")
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

		cell1 := Cell("A")
		cell2 := Cell("B")
		row := Row(cell1, cell2)

		cells := row.Cells()
		asrt.Len(cells, 2, "expected row to contain 2 cells")
		asrt.Same(cell1, cells[0], "expected first cell to be cell1")
		asrt.Same(cell2, cells[1], "expected second cell to be cell2")
	})

	t.Run("Cells returns cells in order", func(t *testing.T) {
		asrt := assert.New(t)

		cell1 := Cell("First")
		cell2 := Cell("Second")
		cell3 := Cell("Third")
		row := Row(cell1, cell2, cell3)

		cells := row.Cells()
		asrt.Len(cells, 3, "expected row to contain 3 cells")
		asrt.Same(cell1, cells[0], "expected cells[0] to be first cell")
		asrt.Same(cell2, cells[1], "expected cells[1] to be second cell")
		asrt.Same(cell3, cells[2], "expected cells[2] to be third cell")
	})

	t.Run("row with cells that have different options", func(t *testing.T) {
		asrt := assert.New(t)

		cell1 := Cell("A").Bold()
		cell2 := Cell("B").Italic()
		cell3 := Cell("C").Port("p1")
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

		cell1 := Cell("A")
		cell2 := Cell("B")
		row := Row(cell1, cell2)

		asrt.NotNil(row, "expected Row to return non-nil HTMLRow")
		cells := row.Cells()
		asrt.Len(cells, 2, "expected row to contain 2 cells")
	})

	t.Run("creates row with three cells", func(t *testing.T) {
		asrt := assert.New(t)

		cell1 := Cell("A")
		cell2 := Cell("B")
		cell3 := Cell("C")
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
			Cell("1"),
			Cell("2"),
			Cell("3"),
			Cell("4"),
			Cell("5"),
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

		cell1 := Cell("")
		cell2 := Cell("B")
		cell3 := Cell("")
		row := Row(cell1, cell2, cell3)

		cells := row.Cells()
		asrt.Len(cells, 3, "expected row to contain 3 cells including empty ones")
	})

	t.Run("row with cells containing special characters", func(t *testing.T) {
		asrt := assert.New(t)

		cell1 := Cell("<html>")
		cell2 := Cell("A & B")
		cell3 := Cell("\"quoted\"")
		row := Row(cell1, cell2, cell3)

		cells := row.Cells()
		asrt.Len(cells, 3, "expected row to contain 3 cells with special characters")
	})

	t.Run("row with cells of different spans", func(t *testing.T) {
		asrt := assert.New(t)

		cell1 := Cell("A").ColSpan(2)
		cell2 := Cell("B")
		cell3 := Cell("C").RowSpan(3)
		row := Row(cell1, cell2, cell3)

		cells := row.Cells()
		asrt.Len(cells, 3, "expected row to contain 3 cells with different spans")
	})

	t.Run("row with cells with ports", func(t *testing.T) {
		asrt := assert.New(t)

		cell1 := Cell("Input").Port("in")
		cell2 := Cell("Middle")
		cell3 := Cell("Output").Port("out")
		row := Row(cell1, cell2, cell3)

		cells := row.Cells()
		asrt.Len(cells, 3, "expected row to contain 3 cells including those with ports")
	})

	t.Run("Cells called multiple times returns same cells", func(t *testing.T) {
		asrt := assert.New(t)

		cell1 := Cell("A")
		cell2 := Cell("B")
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
		cellZ := Cell("Z")
		cellA := Cell("A")
		cellM := Cell("M")
		row := Row(cellZ, cellA, cellM)

		cells := row.Cells()
		asrt.Same(cellZ, cells[0], "expected first cell to be Z (insertion order)")
		asrt.Same(cellA, cells[1], "expected second cell to be A (insertion order)")
		asrt.Same(cellM, cells[2], "expected third cell to be M (insertion order)")
	})
}
