package goraffe

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNodeAttributes_ZeroValue(t *testing.T) {
	t.Run("all fields are empty or zero", func(t *testing.T) {
		asrt := assert.New(t)

		var attrs NodeAttributes

		asrt.Empty(attrs.Label(), "expected Label to be empty")
		asrt.Empty(attrs.Shape(), "expected Shape to be empty")
		asrt.Empty(attrs.Color(), "expected Color to be empty")
		asrt.Empty(attrs.FillColor(), "expected FillColor to be empty")
		asrt.Empty(attrs.FontName(), "expected FontName to be empty")
		asrt.Equal(0.0, attrs.FontSize(), "expected FontSize to be zero")
	})
}

func TestNodeAttributes_Custom_ReturnsCopy(t *testing.T) {
	t.Run("returns empty map when custom is nil", func(t *testing.T) {
		asrt := assert.New(t)

		var attrs NodeAttributes

		custom := attrs.Custom()
		asrt.NotNil(custom, "expected Custom() to return empty map, not nil")
		asrt.Empty(custom, "expected Custom() to return empty map when custom field is nil")
	})

	t.Run("each call returns a different map instance", func(t *testing.T) {
		asrt := assert.New(t)
		attrs := NodeAttributes{}

		custom1 := attrs.Custom()
		custom2 := attrs.Custom()

		// Modify one map
		custom1["test"] = "value"

		// The other map should not have this modification
		_, exists := custom2["test"]
		asrt.False(exists, "expected each call to Custom() to return a new map instance")
	})

	t.Run("modifying returned map does not affect original", func(t *testing.T) {
		asrt := assert.New(t)
		attrs := NodeAttributes{}

		// Get the custom map and modify it
		custom1 := attrs.Custom()
		custom1["test"] = "value"

		// Get custom again - should not have the modification
		custom2 := attrs.Custom()
		asrt.Empty(custom2, "expected modification to returned copy to not affect original")

		_, exists := custom2["test"]
		asrt.False(exists, "expected modification to first copy to not appear in second copy")
	})
}
