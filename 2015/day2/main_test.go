package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSizeFor(t *testing.T) {
	tests := []struct {
		name       string
		l, w, h    int
		expectSize int
	}{
		{
			name:       "2x3x4",
			l:          2,
			w:          3,
			h:          4,
			expectSize: 58,
		},
		{
			name:       "1x1x10",
			l:          1,
			w:          1,
			h:          10,
			expectSize: 43,
		},
		{
			name:       "10x1x1",
			l:          10,
			w:          1,
			h:          1,
			expectSize: 43,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expectSize, sizeFor(tt.l, tt.w, tt.h))
		})
	}
}
