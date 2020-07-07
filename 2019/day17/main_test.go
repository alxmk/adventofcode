package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWorld_SumOfAlignmentParameters(t *testing.T) {
	tests := []struct {
		name   string
		input  string
		expect int
	}{
		{
			name: "Ex1",
			input: `..#..........
..#..........
#######...###
#.#...#...#.#
#############
..#...#...#..
..#####...^..`,
			expect: 76,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expect, parseWorld(tt.input).SumOfAlignmentParameters())
		})
	}
}

func parseWorld(text string) *world {
	w := &world{tiles: make(map[coordinate]rune)}
	var x, y int
	for _, r := range text {
		if x > w.xmax {
			w.xmax = x
		}
		if y > w.ymax {
			w.ymax = y
		}
		switch r {
		case '\n':
			y++
			x = 0
		default:
			w.tiles[coordinate{x, y}] = r
			x++
		}
	}

	return w
}
