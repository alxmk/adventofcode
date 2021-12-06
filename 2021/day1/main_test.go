package main

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestPartTwo(t *testing.T) {
	tests := []struct {
		name   string
		input  string
		expect int
	}{
		{
			name: "Example",
			input: `199
200
208
210
200
207
240
269
260
263`,
			expect: 5,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual, err := partTwo(strings.Split(tt.input, "\n"))
			require.NoError(t, err)
			assert.Equal(t, tt.expect, actual)
		})
	}
}
