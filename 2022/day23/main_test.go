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
			name: "Example",
			input: `
....#..
..###.#
#...#.#
.#...##
#.###..
##.#.##
.#..#..`,
			expect: 110,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expect, partOne(parse([]byte(tt.input))))
		})
	}
}

func TestWorldTick(t *testing.T) {
	tests := []struct {
		name   string
		input  string
		i      int
		expect string
	}{
		{
			name: "Example 1 Turn 1",
			input: `.....
..##.
..#..
.....
..##.
.....`,
			i: 0,
			expect: `..##.
.....
..#..
...#.
..#..
.....
`,
		},
		{
			name: "Example 1 Turn 2",
			input: `..##.
.....
..#..
...#.
..#..
.....`,
			i: 1,
			expect: `.....
..##.
.#...
....#
.....
..#..
`,
		},
		{
			name: "Example 1 Turn 3",
			input: `.....
..##.
.#...
....#
.....
..#..`,
			i: 2,
			expect: `..#..
....#
#....
....#
.....
..#..
`,
		},
		{
			name: "Example 1 Turn 4",
			input: `..#..
....#
#....
....#
.....
..#..`,
			i: 3,
			expect: `..#..
....#
#....
....#
.....
..#..
`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := parse([]byte(tt.input))
			w.Tick(tt.i)
			assert.Equal(t, tt.expect, w.String())
		})
	}
}
