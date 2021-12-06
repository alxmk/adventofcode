package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGenerateGrid(t *testing.T) {
	tests := []struct {
		name             string
		input            string
		includeDiagonals bool
		expectCount      int
	}{
		{
			name:             "Example part two",
			input:            sampleInput,
			includeDiagonals: true,
			expectCount:      12,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			lines, err := parse(tt.input)
			require.NoError(t, err)

			actualCount := generateGrid(lines, tt.includeDiagonals).Count()
			assert.Equal(t, tt.expectCount, actualCount)
		})
	}
}

var sampleInput = `0,9 -> 5,9
8,0 -> 0,8
9,4 -> 3,4
2,2 -> 2,1
7,0 -> 7,4
6,4 -> 2,0
0,9 -> 2,9
3,4 -> 1,4
0,0 -> 8,8
5,5 -> 8,2`
