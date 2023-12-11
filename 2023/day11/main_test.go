package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolve(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		distance int
		expect   int
	}{
		{
			name: "Example 1",
			input: `...#......
.......#..
#.........
..........
......#...
.#........
.........#
..........
.......#..
#...#.....`,
			distance: 2,
			expect:   374,
		},
		{
			name: "Example 2",
			input: `...#......
.......#..
#.........
..........
......#...
.#........
.........#
..........
.......#..
#...#.....`,
			distance: 10,
			expect:   1030,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expect, solve(parseGalaxies(tt.input), tt.distance))
		})
	}
}
