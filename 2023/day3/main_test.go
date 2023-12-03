package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSchematic_PartOne(t *testing.T) {
	tests := []struct {
		name   string
		input  string
		expect int
	}{
		{
			name: "Example 1",
			input: `467..114..
...*......
..35..633.
......#...
617*......
.....+.58.
..592.....
......755.
...$.*....
.664.598..`,
			expect: 4361,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expect, parseSchematic(tt.input).PartOne())
		})
	}
}

func TestSchematic_PartTwo(t *testing.T) {
	tests := []struct {
		name   string
		input  string
		expect int64
	}{
		{
			name: "Example 1",
			input: `467..114..
...*......
..35..633.
......#...
617*......
.....+.58.
..592.....
......755.
...$.*....
.664.598..`,
			expect: 467835,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expect, parseSchematic(tt.input).PartTwo())
		})
	}
}
