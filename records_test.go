package goraffe

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRecordField_Content(t *testing.T) {
	asrt := assert.New(t)

	field := Field("test content")
	got := field.renderRecord()

	asrt.Equal("test content", got, "Field().renderRecord() should return content")
}

func TestRecordField_WithPort(t *testing.T) {
	field := Field("test").Port("p1")

	t.Run("renders with port syntax", func(t *testing.T) {
		asrt := assert.New(t)

		got := field.renderRecord()

		asrt.Equal("<p1> test", got, "Field().Port().renderRecord() should include port syntax")
	})

	t.Run("GetPort returns port reference", func(t *testing.T) {
		asrt := assert.New(t)

		port := field.GetPort()

		asrt.NotNil(port, "GetPort() should return non-nil Port")
		asrt.Equal("p1", port.ID(), "GetPort().ID() should return port ID")
	})

	t.Run("GetPort returns nil when no port set", func(t *testing.T) {
		asrt := assert.New(t)

		fieldNoPort := Field("test")

		asrt.Nil(fieldNoPort.GetPort(), "GetPort() should return nil for field without port")
	})
}

func TestRecordField_Escaping(t *testing.T) {
	tests := []struct {
		name    string
		content string
		want    string
	}{
		{"pipe", "a|b", "a\\|b"},
		{"left brace", "a{b", "a\\{b"},
		{"right brace", "a}b", "a\\}b"},
		{"left angle", "a<b", "a\\<b"},
		{"right angle", "a>b", "a\\>b"},
		{"backslash", "a\\b", "a\\\\b"},
		{"multiple special chars", "a|b{c}d<e>f", "a\\|b\\{c\\}d\\<e\\>f"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			asrt := assert.New(t)

			field := Field(tt.content)
			got := field.renderRecord()

			asrt.Equal(tt.want, got, "Field(%q).renderRecord() should escape special characters", tt.content)
		})
	}
}

func TestRecordGroup_Nesting(t *testing.T) {
	asrt := assert.New(t)

	group := FieldGroup(Field("x"), Field("y"))
	got := group.renderRecord()

	asrt.Equal("{ x | y }", got, "FieldGroup().renderRecord() should wrap fields in braces")
}

func TestRecordGroup_NestedGroups(t *testing.T) {
	asrt := assert.New(t)

	innerGroup := FieldGroup(Field("a"), Field("b"))
	outerGroup := FieldGroup(Field("x"), innerGroup, Field("y"))

	got := outerGroup.renderRecord()

	asrt.Equal("{ x | { a | b } | y }", got, "nested FieldGroup().renderRecord() should handle nested groups")
}

func TestRecordLabel_SimpleFields(t *testing.T) {
	asrt := assert.New(t)

	label := Record(Field("a"), Field("b"))
	got := label.String()

	asrt.Equal("a | b", got, "Record().String() should join fields with pipes")
}

func TestRecordLabel_WithGroup(t *testing.T) {
	asrt := assert.New(t)

	label := Record(
		Field("header"),
		FieldGroup(Field("left"), Field("right")),
		Field("footer"),
	)

	got := label.String()

	asrt.Equal("header | { left | right } | footer", got, "Record().String() should handle groups within fields")
}

func TestRecordLabel_WithPorts(t *testing.T) {
	asrt := assert.New(t)

	label := Record(
		Field("input").Port("in"),
		Field("output").Port("out"),
	)

	got := label.String()

	asrt.Equal("<in> input | <out> output", got, "Record() with ports String() should include port syntax")
}

func TestRecordLabel_Escaping(t *testing.T) {
	asrt := assert.New(t)

	label := Record(
		Field("field|with|pipes"),
		Field("field{with}braces"),
	)

	got := label.String()

	asrt.Equal(
		"field\\|with\\|pipes | field\\{with\\}braces",
		got, "Record() with special chars String() should escape characters",
	)
}

func TestRecordLabel_SetNodeContext(t *testing.T) {
	field1 := Field("a").Port("p1")
	field2 := Field("b").Port("p2")
	nestedField := Field("c").Port("p3")
	group := FieldGroup(nestedField)

	label := Record(field1, group, field2)
	label.setNodeContext("node1")

	tests := []struct {
		name     string
		port     *Port
		wantID   string
		wantNode string
	}{
		{"field1 port", field1.GetPort(), "p1", "node1"},
		{"field2 port", field2.GetPort(), "p2", "node1"},
		{"nested field port", nestedField.GetPort(), "p3", "node1"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			asrt := assert.New(t)

			asrt.NotNil(tt.port, "GetPort() should return non-nil")
			asrt.Equal(tt.wantID, tt.port.ID(), "port.ID() should match")
			asrt.Equal(tt.wantNode, tt.port.NodeID(), "port.NodeID() should match")
		})
	}
}

func TestRecordLabel_ComplexExample(t *testing.T) {
	asrt := assert.New(t)

	// Test a realistic complex record label
	label := Record(
		Field("Title").Port("title"),
		FieldGroup(
			Field("Field 1"),
			Field("Field 2"),
			FieldGroup(
				Field("Nested A"),
				Field("Nested B"),
			),
		),
		Field("Footer"),
	)

	got := label.String()

	asrt.Equal(
		"<title> Title | { Field 1 | Field 2 | { Nested A | Nested B } } | Footer",
		got, "complex Record().String() should handle nested structures",
	)
}
