package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPath(t *testing.T) {
	tests := []struct {
		name   string
		input  string
		expect int
	}{
		{
			name: "Example",
			input: `1163751742
1381373672
2136511328
3694931569
7463417111
1319128137
1359912421
3125421639
1293138521
2311944581`,
			expect: 40,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := parse(tt.input)
			assert.Equal(t, tt.expect, g.Path(coord{0, 0}, coord{g.xmax, g.ymax}))
		})
	}
}

func TestPathStonksed(t *testing.T) {
	tests := []struct {
		name   string
		input  string
		stonks int
		expect int
	}{
		{
			name: "Example",
			input: `1163751742
1381373672
2136511328
3694931569
7463417111
1319128137
1359912421
3125421639
1293138521
2311944581`,
			stonks: 5,
			expect: 315,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := parse(tt.input)
			g = g.Stonks(tt.stonks)
			assert.Equal(t, tt.expect, g.Path(coord{0, 0}, coord{g.xmax, g.ymax}))
		})
	}
}
