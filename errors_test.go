package goraffe

import (
	"errors"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRenderError_Error_IncludesStderr(t *testing.T) {
	tests := []struct {
		name     string
		err      *RenderError
		contains []string
	}{
		{
			name: "with stderr",
			err: &RenderError{
				Err:      ErrRenderFailed,
				Stderr:   "Error: syntax error in line 5",
				ExitCode: 1,
			},
			contains: []string{"rendering failed", "exit code 1", "syntax error"},
		},
		{
			name: "without stderr",
			err: &RenderError{
				Err:      ErrGraphvizNotFound,
				Stderr:   "",
				ExitCode: 127,
			},
			contains: []string{"graphviz not found", "exit code 127"},
		},
		{
			name: "long stderr is truncated",
			err: &RenderError{
				Err:      ErrInvalidDOT,
				Stderr:   strings.Repeat("error ", 100),
				ExitCode: 1,
			},
			contains: []string{"invalid DOT", "exit code 1", "..."},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			errMsg := tt.err.Error()
			for _, substr := range tt.contains {
				assert.Contains(t, errMsg, substr)
			}
		})
	}
}

func TestRenderError_Unwrap(t *testing.T) {
	underlyingErr := errors.New("underlying error")
	renderErr := &RenderError{
		Err:      underlyingErr,
		Stderr:   "some stderr",
		ExitCode: 1,
	}

	assert.Equal(t, underlyingErr, renderErr.Unwrap())
	assert.ErrorIs(t, renderErr, underlyingErr)
}

func TestRenderError_Is_RenderFailed(t *testing.T) {
	renderErr := &RenderError{
		Err:      ErrRenderFailed,
		Stderr:   "error output",
		ExitCode: 1,
	}

	assert.ErrorIs(t, renderErr, ErrRenderFailed)
	assert.True(t, errors.Is(renderErr, ErrRenderFailed))
}

func TestSentinelErrors_Distinct(t *testing.T) {
	// Ensure all sentinel errors are distinct
	assert.NotEqual(t, ErrGraphvizNotFound, ErrInvalidDOT)
	assert.NotEqual(t, ErrGraphvizNotFound, ErrRenderFailed)
	assert.NotEqual(t, ErrInvalidDOT, ErrRenderFailed)

	// Ensure they can be distinguished with errors.Is
	err1 := ErrGraphvizNotFound
	assert.True(t, errors.Is(err1, ErrGraphvizNotFound))
	assert.False(t, errors.Is(err1, ErrInvalidDOT))
	assert.False(t, errors.Is(err1, ErrRenderFailed))
}
