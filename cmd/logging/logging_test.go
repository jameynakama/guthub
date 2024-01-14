package logging

import (
	"bytes"
	"testing"

	"github.com/jameynakama/assert"
)

func NewTestLogger(t *testing.T, buf *bytes.Buffer) *GutHubLogger {
	return NewGutHubLogger(buf, buf, buf, "[TEST] ", 0)
}

func TestLogger(t *testing.T) {
	tcs := []struct {
		name     string
		method   func(l *GutHubLogger, v ...any)
		input    any
		expected string
	}{
		{
			name:     "TestInfoLog",
			method:   (*GutHubLogger).Info,
			input:    "test info",
			expected: "[TEST] INFO: test info\n",
		},
		{
			name:     "TestErrorLog",
			method:   (*GutHubLogger).Error,
			input:    "test error",
			expected: "[TEST] ERROR: test error\n",
		},
		{
			name:     "TestDebugLog",
			method:   (*GutHubLogger).Debug,
			input:    "test debug",
			expected: "[TEST] >>> DEBUG: test debug\n",
		},
	}

	for _, tc := range tcs {
		var buf bytes.Buffer
		l := NewTestLogger(t, &buf)

		tc.method(l, tc.input)

		assert.Equal(t, buf.String(), tc.expected)
	}
}
