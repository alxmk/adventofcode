package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSummarise(t *testing.T) {
	tests := []struct {
		name   string
		input  string
		expect int
	}{
		{
			name: "Ex1",
			input: `#.##..##.
..#.##.#.
##......#
##......#
..#.##.#.
..##..##.
#.#.##.#.`,
			expect: 5,
		},
		{
			name: "Ex2",
			input: `#...##..#
#....#..#
..##..###
#####.##.
#####.##.
..##..###
#....#..#`,
			expect: 400,
		},
		{
			name: "Ex3",
			input: `..#...####...#...
....#.#..#.#.....
..##.######.##...
...#.######.#....
#.###..##...##.##
##..##.##.##..###
......####.......
...###.##.###....
....###..###.....
.##.#......#.##..
.#.#........#.#..
.##...#..#...##..
#..#.#....#.#..##`,
			expect: 16,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expect, summarise(tt.input, findSymmetry(0)))
		})
	}
}

func TestSummariseSmudged(t *testing.T) {
	tests := []struct {
		name   string
		input  string
		expect int
	}{
		{
			name: "Ex1",
			input: `#.##..##.
..#.##.#.
##......#
##......#
..#.##.#.
..##..##.
#.#.##.#.`,
			expect: 300,
		},
		{
			name: "Ex2",
			input: `#...##..#
#....#..#
..##..###
#####.##.
#####.##.
..##..###
#....#..#`,
			expect: 100,
		},
		{
			name: "Ex3",
			input: `####.......#..#..
.......####....##
....####.##.##.##
######..####..###
.##...##..######.
.##..##.##.####.#
#..#.###.#.#.##.#`,
			expect: 13,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expect, summarise(tt.input, findSymmetry(1)))
		})
	}
}

func TestLineDiff(t *testing.T) {
	tests := []struct {
		a, b   string
		expect int
	}{
		{
			a:      "#.##..##.",
			b:      "..##..##.",
			expect: 1,
		},
	}

	for _, tt := range tests {
		assert.Equal(t, tt.expect, lineDiff(tt.a, tt.b))
	}
}
