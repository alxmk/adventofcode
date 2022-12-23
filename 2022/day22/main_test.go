package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolve(t *testing.T) {
	tests := []struct {
		name   string
		input  string
		wf     wrappingFunc
		expect int
	}{
		{
			name:   "Example p1",
			input:  example,
			wf:     p1,
			expect: 6032,
		},
		{
			name:   "Example p2",
			input:  example,
			wf:     p2(testSeams),
			expect: 5031,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w, r := parse([]byte(tt.input))
			assert.Equal(t, tt.expect, solve(w, r, tt.wf))
		})
	}
}

var testSeams = []seam{
	{
		name: "A",
		a:    edge{vertices: [2]xy{{8, 0}, {8, 3}}, transform: map[direction]direction{left: down}},
		b:    edge{vertices: [2]xy{{4, 4}, {8, 4}}, transform: map[direction]direction{up: right}},
	},
	{
		name: "B",
		a:    edge{vertices: [2]xy{{8, 0}, {12, 0}}, transform: map[direction]direction{up: down}},
		b:    edge{vertices: [2]xy{{4, 4}, {0, 4}}, transform: map[direction]direction{up: down}},
	},
	{
		name: "C",
		a:    edge{vertices: [2]xy{{4, 7}, {7, 7}}, transform: map[direction]direction{down: right}},
		b:    edge{vertices: [2]xy{{8, 11}, {8, 8}}, transform: map[direction]direction{left: up}},
	},
	{
		name: "D",
		a:    edge{vertices: [2]xy{{11, 7}, {11, 4}}, transform: map[direction]direction{right: down}},
		b:    edge{vertices: [2]xy{{12, 8}, {15, 8}}, transform: map[direction]direction{up: left}},
	},
	{
		name: "E",
		a:    edge{vertices: [2]xy{{3, 7}, {0, 7}}, transform: map[direction]direction{down: up}},
		b:    edge{vertices: [2]xy{{8, 11}, {11, 11}}, transform: map[direction]direction{down: up}},
	},
	{
		name: "F",
		a:    edge{vertices: [2]xy{{0, 7}, {0, 4}}, transform: map[direction]direction{left: up}},
		b:    edge{vertices: [2]xy{{12, 11}, {15, 11}}, transform: map[direction]direction{down: right}},
	},
	{
		name: "G",
		a:    edge{vertices: [2]xy{{11, 0}, {11, 3}}, transform: map[direction]direction{right: left}},
		b:    edge{vertices: [2]xy{{15, 11}, {15, 8}}, transform: map[direction]direction{right: left}},
	},
}

func TestParseRoute(t *testing.T) {
	tests := []struct {
		name   string
		input  string
		expect []instruction
	}{
		{
			name:  "Example",
			input: "10R5L5R10L4R5L5",
			expect: []instruction{
				{move: 10},
				{turn: right},
				{move: 5},
				{turn: left},
				{move: 5},
				{turn: right},
				{move: 10},
				{turn: left},
				{move: 4},
				{turn: right},
				{move: 5},
				{turn: left},
				{move: 5},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expect, parseRoute([]byte(tt.input)))
		})
	}
}

var example = `        ...#
        .#..
        #...
        ....
...#.......#
........#...
..#....#....
..........#.
        ...#....
        .....#..
        .#......
        ......#.

10R5L5R10L4R5L5`

func TestTraverse(t *testing.T) {
	tests := []struct {
		name    string
		seam    seam
		p       xy
		h       direction
		expect  xy
		expectH direction
	}{
		{
			name: "Seam B1",
			seam: seam{ // B
				a:    edge{vertices: [2]xy{{50, 0}, {99, 0}}, transform: map[direction]direction{up: right}},
				b:    edge{vertices: [2]xy{{0, 150}, {0, 199}}, transform: map[direction]direction{left: down}},
				name: "B",
			},
			p:       xy{50, 0},
			h:       up,
			expect:  xy{0, 150},
			expectH: right,
		},
		{
			name: "Seam B2",
			seam: seam{ // B
				a:    edge{vertices: [2]xy{{50, 0}, {99, 0}}, transform: map[direction]direction{up: right}},
				b:    edge{vertices: [2]xy{{0, 150}, {0, 199}}, transform: map[direction]direction{left: down}},
				name: "B",
			},
			p:       xy{0, 150},
			h:       left,
			expect:  xy{50, 0},
			expectH: down,
		},
		{
			name: "Seam B3",
			seam: seam{ // B
				a:    edge{vertices: [2]xy{{50, 0}, {99, 0}}, transform: map[direction]direction{up: right}},
				b:    edge{vertices: [2]xy{{0, 150}, {0, 199}}, transform: map[direction]direction{left: down}},
				name: "B",
			},
			p:       xy{0, 155},
			h:       left,
			expect:  xy{55, 0},
			expectH: down,
		},
		{
			name: "Seam A1",
			seam: seam{ // A
				a:    edge{vertices: [2]xy{{50, 0}, {50, 49}}, transform: map[direction]direction{left: right}},
				b:    edge{vertices: [2]xy{{0, 149}, {0, 100}}, transform: map[direction]direction{left: right}},
				name: "A",
			},
			p:       xy{50, 30},
			h:       left,
			expect:  xy{0, 119},
			expectH: right,
		},
		{
			name: "Seam ED1",
			seam: seam{
				name: "D",
				a:    edge{vertices: [2]xy{{11, 7}, {11, 4}}, transform: map[direction]direction{right: down}},
				b:    edge{vertices: [2]xy{{12, 8}, {15, 8}}, transform: map[direction]direction{up: left}},
			},
			p:       xy{11, 5},
			h:       right,
			expect:  xy{14, 8},
			expectH: down,
		},
		{
			name: "Seam G1",
			seam: seam{ // G
				a:    edge{vertices: [2]xy{{149, 0}, {100, 0}}, transform: map[direction]direction{up: up}},
				b:    edge{vertices: [2]xy{{49, 199}, {0, 199}}, transform: map[direction]direction{down: down}},
				name: "G",
			},
			p:       xy{49, 199},
			h:       down,
			expect:  xy{149, 0},
			expectH: down,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, actual, actualH := tt.seam.Traverse(tt.p, tt.h)
			assert.Equal(t, tt.expect, actual)
			assert.Equal(t, tt.expectH, actualH)
		})
	}
}
