// ABOUTME: This file defines record label building blocks for DOT graph nodes.
// ABOUTME: Record labels allow creating structured, port-based nodes with fields and groups.
package goraffe

import (
	"strings"
)

// RecordElement is a marker interface for elements that can be part of a record label.
// Both RecordField and RecordGroup implement this interface.
type RecordElement interface {
	recordElement()
	renderRecord() string
}

// RecordField represents a single field in a record label.
// Fields can have content and optionally a port identifier.
type RecordField struct {
	content string
	port    string
	portRef *Port
}

// Field creates a new RecordField with the given content.
func Field(content string) *RecordField {
	return &RecordField{
		content: content,
	}
}

// Port sets the port identifier for this field and returns the field for chaining.
// This creates a Port reference that can be used for edge connections.
func (f *RecordField) Port(id string) *RecordField {
	f.port = id
	f.portRef = &Port{id: id}
	return f
}

// GetPort returns the Port reference associated with this field, or nil if no port was set.
func (f *RecordField) GetPort() *Port {
	return f.portRef
}

// recordElement implements the RecordElement interface marker method.
func (f *RecordField) recordElement() {}

// renderRecord renders this field to record label syntax.
func (f *RecordField) renderRecord() string {
	content := escapeRecordString(f.content)
	if f.port != "" {
		return "<" + f.port + "> " + content
	}
	return content
}

// RecordGroup represents a grouped set of record elements.
// Groups are rendered wrapped in braces { }.
type RecordGroup struct {
	elements []RecordElement
}

// FieldGroup creates a new RecordGroup containing the given elements.
func FieldGroup(elements ...RecordElement) *RecordGroup {
	return &RecordGroup{
		elements: elements,
	}
}

// recordElement implements the RecordElement interface marker method.
func (g *RecordGroup) recordElement() {}

// renderRecord renders this group to record label syntax.
func (g *RecordGroup) renderRecord() string {
	parts := make([]string, len(g.elements))
	for i, elem := range g.elements {
		parts[i] = elem.renderRecord()
	}
	return "{ " + strings.Join(parts, " | ") + " }"
}

// RecordLabel represents a complete record label containing multiple elements.
type RecordLabel struct {
	elements []RecordElement
}

// Record creates a new RecordLabel with the given elements.
func Record(elements ...RecordElement) *RecordLabel {
	return &RecordLabel{
		elements: elements,
	}
}

// String renders the record label to DOT record syntax.
// Fields are separated by |, groups are wrapped in { }.
func (l *RecordLabel) String() string {
	parts := make([]string, len(l.elements))
	for i, elem := range l.elements {
		parts[i] = elem.renderRecord()
	}
	return strings.Join(parts, " | ")
}

// setNodeContext is an internal method that associates all ports in this label with a node ID.
// This is called when the label is attached to a node.
func (l *RecordLabel) setNodeContext(nodeID string) {
	for _, elem := range l.elements {
		setPortContextRecursive(elem, nodeID)
	}
}

// setPortContextRecursive recursively sets the node context for all ports in a record element tree.
func setPortContextRecursive(elem RecordElement, nodeID string) {
	switch e := elem.(type) {
	case *RecordField:
		if e.portRef != nil {
			e.portRef.nodeID = nodeID
		}
	case *RecordGroup:
		for _, child := range e.elements {
			setPortContextRecursive(child, nodeID)
		}
	}
}

// escapeRecordString escapes special characters in record label strings.
// The special characters |, {, }, <, > must be escaped with backslash.
func escapeRecordString(s string) string {
	s = strings.ReplaceAll(s, "\\", "\\\\")
	s = strings.ReplaceAll(s, "|", "\\|")
	s = strings.ReplaceAll(s, "{", "\\{")
	s = strings.ReplaceAll(s, "}", "\\}")
	s = strings.ReplaceAll(s, "<", "\\<")
	s = strings.ReplaceAll(s, ">", "\\>")
	return s
}
