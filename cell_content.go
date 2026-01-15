// ABOUTME: This file defines content types that can be placed inside HTML table cells.
// ABOUTME: Includes text with formatting, line breaks, and horizontal rules.

package goraffe

// Content represents any content piece that can appear inside a cell.
// The unexported contentMarker() method prevents external implementations.
type Content interface {
	contentMarker()
	toHTML() string
}

// TextContent represents text with optional formatting (bold, italic, underline).
type TextContent struct {
	text      string
	bold      bool
	italic    bool
	underline bool
}

// Text creates a new TextContent with the given text.
func Text(text string) *TextContent {
	return &TextContent{
		text: text,
	}
}

// Bold sets the text to be bold and returns the TextContent for chaining.
func (t *TextContent) Bold() *TextContent {
	t.bold = true
	return t
}

// Italic sets the text to be italic and returns the TextContent for chaining.
func (t *TextContent) Italic() *TextContent {
	t.italic = true
	return t
}

// Underline sets the text to be underlined and returns the TextContent for chaining.
func (t *TextContent) Underline() *TextContent {
	t.underline = true
	return t
}

func (t *TextContent) contentMarker() {}

func (t *TextContent) toHTML() string {
	result := t.text

	// Nest tags in order: U → I → B
	if t.underline {
		result = "<u>" + result + "</u>"
	}
	if t.italic {
		result = "<i>" + result + "</i>"
	}
	if t.bold {
		result = "<b>" + result + "</b>"
	}

	return result
}

// LineBreak represents a line break (<BR/>) in HTML.
type LineBreak struct{}

// BR creates a new LineBreak.
func BR() *LineBreak {
	return &LineBreak{}
}

func (l *LineBreak) contentMarker() {}

func (l *LineBreak) toHTML() string {
	return "<br/>"
}

// HorizontalRule represents a horizontal rule (<HR/>) in HTML.
type HorizontalRule struct{}

// HR creates a new HorizontalRule.
func HR() *HorizontalRule {
	return &HorizontalRule{}
}

func (h *HorizontalRule) contentMarker() {}

func (h *HorizontalRule) toHTML() string {
	return "<hr/>"
}
