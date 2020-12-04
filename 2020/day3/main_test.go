package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestForest_Move(t *testing.T) {
	tests := []struct {
		name   string
		input  string
		dirs   []vec
		expect int
	}{
		{
			name: "Ex1",
			input: `..##.......
#...#...#..
.#....#..#.
..#.#...#.#
.#...##..#.
..#.##.....
.#.#.#....#
.#........#
#.##...#...
#...##....#
.#..#...#.#`,
			dirs:   []vec{{3, 1}},
			expect: 7,
		},
		{
			name: "Ex1",
			input: `..##.......
#...#...#..
.#....#..#.
..#.#...#.#
.#...##..#.
..#.##.....
.#.#.#....#
.#........#
#.##...#...
#...##....#
.#..#...#.#`,
			dirs: []vec{{1, 1},
				{3, 1},
				{5, 1},
				{7, 1},
				{1, 2}},
			expect: 336,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := parse(tt.input)
			total := 1
			for _, dir := range tt.dirs {
				count := f.move(dir)
				fmt.Println(dir, count)
				total *= count
			}
			assert.Equal(t, tt.expect, total)
		})
	}
}
