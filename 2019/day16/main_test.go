package main

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPatternIdx(t *testing.T) {
	tests := []struct {
		name          string
		pattern       []int
		inIdx, outIdx int
		expectIdx     int
	}{
		{
			name:      "Ex1",
			pattern:   []int{0, 1, 0, -1},
			inIdx:     0,
			outIdx:    0,
			expectIdx: 1,
		},
		{
			name:      "Ex2",
			pattern:   []int{0, 1, 0, -1},
			inIdx:     3,
			outIdx:    0,
			expectIdx: 0,
		},
		{
			name:      "Ex3",
			pattern:   []int{0, 1, 0, -1},
			inIdx:     4,
			outIdx:    1,
			expectIdx: 2,
		},
		{
			name:      "Ex4",
			pattern:   []int{0, 1, 0, -1},
			inIdx:     7,
			outIdx:    1,
			expectIdx: 0,
		},
		{
			name:      "Ex5",
			pattern:   []int{0, 1, 0, -1},
			inIdx:     8,
			outIdx:    1,
			expectIdx: 0,
		},
		{
			name:      "Ex6",
			pattern:   []int{0, 1, 0, -1},
			inIdx:     6,
			outIdx:    1,
			expectIdx: 3,
		},
		{
			name:      "Ex7",
			pattern:   []int{0, 1, 0, -1},
			inIdx:     10,
			outIdx:    1,
			expectIdx: 1,
		},
		{
			name:      "Ex8",
			pattern:   []int{0, 1, 0, -1},
			inIdx:     4,
			outIdx:    2,
			expectIdx: 1,
		},
		{
			name:      "Ex9",
			pattern:   []int{0, 1, 0, -1},
			inIdx:     5,
			outIdx:    2,
			expectIdx: 2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expectIdx, patternIdx(tt.inIdx, tt.outIdx, len(tt.pattern)))
		})
	}
}

func TestRunPhase(t *testing.T) {
	tests := []struct {
		name      string
		input     string
		numPhases int
		expect    string
	}{
		{
			name:      "Ex1",
			input:     "12345678",
			numPhases: 1,
			expect:    "48226158",
		},
		{
			name:      "Ex2",
			input:     "12345678",
			numPhases: 2,
			expect:    "34040438",
		},
		{
			name:      "Ex3",
			input:     "12345678",
			numPhases: 3,
			expect:    "03415518",
		},
		{
			name:      "Ex4",
			input:     "12345678",
			numPhases: 4,
			expect:    "01029498",
		},
		{
			name:      "Ex5",
			input:     "80871224585914546619083218645595",
			numPhases: 100,
			expect:    "24176176",
		},
		{
			name:      "Ex6",
			input:     "19617804207202209144916044189917",
			numPhases: 100,
			expect:    "73745418",
		},
		{
			name:      "Ex7",
			input:     "69317163492948606335995924319873",
			numPhases: 100,
			expect:    "52432133",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := parseOutputList(tt.input)
			for i := 0; i < tt.numPhases; i++ {
				actual = runPhase(actual, []int{0, 1, 0, -1})
			}
			assert.True(t, strings.HasPrefix(actual.String(), tt.expect))
		})
	}
}
