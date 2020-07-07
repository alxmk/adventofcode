package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestBelt_getMaxInVision(t *testing.T) {
	tests := []struct {
		name              string
		input             string
		expectBest        coordinate
		expectMaxInVision int
	}{
		{
			name: "Ex1",
			input: `.#..#
.....
#####
....#
...##`,
			expectBest:        coordinate{3, 4},
			expectMaxInVision: 8,
		},
		{
			name: "Ex2",
			input: `......#.#.
#..#.#....
..#######.
.#.#.###..
.#..#.....
..#....#.#
#..#....#.
.##.#..###
##...#..#.
.#....####`,
			expectBest:        coordinate{5, 8},
			expectMaxInVision: 33,
		},
		{
			name: "Ex3",
			input: `#.#...#.#.
.###....#.
.#....#...
##.#.#.#.#
....#.#.#.
.##..###.#
..#...##..
..##....##
......#...
.####.###.`,
			expectBest:        coordinate{1, 2},
			expectMaxInVision: 35,
		},
		{
			name: "Ex4",
			input: `.#..##.###...#######
##.############..##.
.#.######.########.#
.###.#######.####.#.
#####.##.#.##.###.##
..#####..#.#########
####################
#.####....###.#.#.##
##.#################
#####.##.###..####..
..######..##.#######
####.##.####...##..#
.#####..#.######.###
##...#.##########...
#.##########.#######
.####.#.###.###.#.##
....##.##.###..#####
.#.#.###########.###
#.#.#.#####.####.###
###.##.####.##.#..##`,
			expectBest:        coordinate{11, 13},
			expectMaxInVision: 210,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b, err := parseAsteroids(tt.input)
			require.NoError(t, err)

			actualBest, actualMaxInVision := b.getMaxInVision()
			assert.Equal(t, tt.expectBest, actualBest)
			assert.Equal(t, tt.expectMaxInVision, actualMaxInVision)
		})
	}
}

func TestBelt_getNumInVision(t *testing.T) {
	tests := []struct {
		name              string
		input             string
		from              coordinate
		expectNumInVision int
	}{
		{
			name: "Ex1",
			input: `......#.#.
#..#.#....
..#######.
.#.#.###..
.#..#.....
..#....#.#
#..#....#.
.##.#..###
##...#..#.
.#....####`,
			from:              coordinate{5, 8},
			expectNumInVision: 33,
		},
		{
			name: "Ex2",
			input: `#.#...#.#.
.###....#.
.#....#...
##.#.#.#.#
....#.#.#.
.##..###.#
..#...##..
..##....##
......#...
.####.###.`,
			from:              coordinate{1, 2},
			expectNumInVision: 35,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b, err := parseAsteroids(tt.input)
			require.NoError(t, err)

			actualNumInVision := b.getNumInVision(tt.from)
			assert.Equal(t, tt.expectNumInVision, actualNumInVision)
		})
	}
}

func TestVector_parallelTo(t *testing.T) {
	tests := []struct {
		name           string
		a, b           vector
		expectParallel bool
	}{
		{
			name:           "Ex1",
			a:              vector{1, 0},
			b:              vector{2, 0},
			expectParallel: true,
		},
		{
			name:           "Ex2",
			a:              vector{1, 1},
			b:              vector{5, 5},
			expectParallel: true,
		},
		{
			name:           "Ex2",
			a:              vector{1, 1},
			b:              vector{5, 4},
			expectParallel: false,
		},
		{
			name:           "Ex4",
			a:              vector{0, 1},
			b:              vector{0, 3},
			expectParallel: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expectParallel, tt.a.parallelTo(tt.b))
		})
	}
}
