package main

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestPartOne(t *testing.T) {
	tests := []struct {
		name          string
		input         string
		expectGamma   uint
		expectEpsilon uint
	}{
		{
			name: "Example",
			input: `00100
11110
10110
10111
10101
01111
00111
11100
10000
11001
00010
01010`,
			expectGamma:   22,
			expectEpsilon: 9,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actualGamma, actualEpsilon, err := partOne(strings.Split(tt.input, "\n"))
			require.NoError(t, err)
			assert.Equal(t, tt.expectGamma, actualGamma)
			assert.Equal(t, tt.expectEpsilon, actualEpsilon)
		})
	}
}

func TestPartTwo(t *testing.T) {
	tests := []struct {
		name      string
		input     string
		expectO2  uint64
		expectCO2 uint64
	}{
		{
			name: "Example",
			input: `00100
11110
10110
10111
10101
01111
00111
11100
10000
11001
00010
01010`,
			expectO2:  23,
			expectCO2: 10,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actualO2, actualCO2, err := partTwo(strings.Split(tt.input, "\n"))
			require.NoError(t, err)
			assert.Equal(t, tt.expectO2, actualO2)
			assert.Equal(t, tt.expectCO2, actualCO2)
		})
	}
}
