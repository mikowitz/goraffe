// ABOUTME: This file contains tests for the Port type and port-related functionality.
// ABOUTME: Tests verify port ID retrieval and cell port reference behavior.
package goraffe

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestPort_ID(t *testing.T) {
	asrt := assert.New(t)

	port := &Port{id: "test_port"}

	asrt.Equal("test_port", port.ID(), "expected Port.ID() to return 'test_port'")
}

func TestCell_GetPort_ReturnsPort(t *testing.T) {
	asrt := assert.New(t)
	req := require.New(t)

	cell := Cell(Text("content")).Port("my_port")

	port := cell.GetPort()
	req.NotNil(port, "expected GetPort() to return non-nil Port")

	asrt.Equal("my_port", port.ID(), "expected port.ID() to return 'my_port'")
}

func TestCell_GetPort_NilIfNoPort(t *testing.T) {
	asrt := assert.New(t)

	cell := Cell(Text("content"))

	port := cell.GetPort()
	asrt.Nil(port, "expected GetPort() to return nil when no port is set")
}
