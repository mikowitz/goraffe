// ABOUTME: This file defines the Port type for type-safe port references in HTML labels.
// ABOUTME: Ports allow edges to connect to specific cells within HTML table labels.
package goraffe

// Port represents a type-safe reference to a port within an HTML label cell.
// Ports are used to specify connection points for edges within HTML table labels.
type Port struct {
	id     string
	nodeID string
}

// ID returns the port's identifier string.
// This is the name used in the HTML label's port attribute.
func (p *Port) ID() string {
	return p.id
}
