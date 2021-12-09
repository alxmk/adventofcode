package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPartOne(t *testing.T) {
	tests := []struct {
		name   string
		input  string
		expect int
	}{
		{
			name:   "Example",
			input:  sampleInput,
			expect: 15,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := parse(tt.input)
			assert.Equal(t, tt.expect, partOne(g))
		})
	}
}

func TestPartTwo(t *testing.T) {
	tests := []struct {
		name   string
		input  string
		expect int
	}{
		{
			name:   "Example",
			input:  sampleInput,
			expect: 1134,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := parse(tt.input)
			assert.Equal(t, tt.expect, partTwo(g))
		})
	}
}

var sampleInput = `2199943210
3987894921
9856789892
8767896789
9899965678`
