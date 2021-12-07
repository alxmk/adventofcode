package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSolve(t *testing.T) {
	tests := []struct {
		name   string
		input  string
		solver func([]int, int) int
		expect int
	}{
		{
			name:   "Example part one",
			input:  sample,
			solver: partOne,
			expect: 37,
		},
		{
			name:   "Example part two",
			input:  sample,
			solver: partTwo,
			expect: 168,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			positions, err := parse(tt.input)
			require.NoError(t, err)
			actual := solve(positions, tt.solver)
			assert.Equal(t, tt.expect, actual)
		})
	}
}

var sample = `16,1,2,0,4,2,7,1,2,14`
