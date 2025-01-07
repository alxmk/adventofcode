package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSearch(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		expectA int
		expectB int
	}{
		{
			name: "Ex1",
			input: `MMMSXXMASM
MSAMXMSMSA
AMXSXMAAMM
MSAMASMSMX
XMASAMXAMM
XXAMMXXAMA
SMSMSASXSS
SAXAMASAAA
MAMMMXMMMM
MXMXAXMASX`,
			expectA: 18,
			expectB: 9,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actualA, actualB := solve(tt.input)
			assert.Equal(t, tt.expectA, actualA)
			assert.Equal(t, tt.expectB, actualB)
		})
	}
}
