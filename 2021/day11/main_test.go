package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStep(t *testing.T) {
	tests := []struct {
		name   string
		input  string
		expect int
	}{
		{
			name: "Example 1",
			input: `11111
19991
19191
19991
11111`,
			expect: 9,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := parse(tt.input)
			actual := o.Step()
			assert.Equal(t, tt.expect, actual)
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
			name: "Example 1",
			input: `5483143223
2745854711
5264556173
6141336146
6357385478
4167524645
2176841721
6882881134
4846848554
5283751526`,
			expect: 195,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := parse(tt.input)
			actual := o.partTwo()
			assert.Equal(t, tt.expect, actual)
		})
	}
}
