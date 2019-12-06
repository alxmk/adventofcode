package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestParseWire(t *testing.T) {
	tests := []struct {
		name   string
		input  string
		expect wire
	}{
		{
			name:  "Ex1",
			input: "R8,U5,L5,D3",
			expect: wire{
				segment{
					start:      coordinate{0, 0},
					end:        coordinate{8, 0},
					stepOffset: 0,
				},
				segment{
					start:      coordinate{8, 0},
					end:        coordinate{8, 5},
					stepOffset: 8,
				},
				segment{
					start:      coordinate{8, 5},
					end:        coordinate{3, 5},
					stepOffset: 13,
				},
				segment{
					start:      coordinate{3, 5},
					end:        coordinate{3, 2},
					stepOffset: 18,
				},
			},
		},
		{
			name:  "Ex2",
			input: "U7,R6,D4,L4",
			expect: wire{
				segment{
					start:      coordinate{0, 0},
					end:        coordinate{0, 7},
					stepOffset: 0,
				},
				segment{
					start:      coordinate{0, 7},
					end:        coordinate{6, 7},
					stepOffset: 7,
				},
				segment{
					start:      coordinate{6, 7},
					end:        coordinate{6, 3},
					stepOffset: 13,
				},
				segment{
					start:      coordinate{6, 3},
					end:        coordinate{2, 3},
					stepOffset: 17,
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w, err := parseWire(tt.input)
			require.NoError(t, err)
			assert.Equal(t, tt.expect, w)
		})
	}
}

func TestClosestIntersection(t *testing.T) {
	tests := []struct {
		name           string
		inputA         string
		inputB         string
		distFunc       distFunc
		expectDistance int
	}{
		{
			name:           "Ex1 Manhattan",
			inputA:         "R8,U5,L5,D3",
			inputB:         "U7,R6,D4,L4",
			distFunc:       manhattan(),
			expectDistance: 6,
		},
		{
			name:           "Ex1 Steps",
			inputA:         "R8,U5,L5,D3",
			inputB:         "U7,R6,D4,L4",
			distFunc:       steps(),
			expectDistance: 30,
		},
		{
			name:           "Ex2 Manhattan",
			inputA:         "R75,D30,R83,U83,L12,D49,R71,U7,L72",
			inputB:         "U62,R66,U55,R34,D71,R55,D58,R83",
			distFunc:       manhattan(),
			expectDistance: 159,
		},
		{
			name:           "Ex2 Steps",
			inputA:         "R75,D30,R83,U83,L12,D49,R71,U7,L72",
			inputB:         "U62,R66,U55,R34,D71,R55,D58,R83",
			distFunc:       steps(),
			expectDistance: 610,
		},
		{
			name:           "Ex3 Manhattan",
			inputA:         "R98,U47,R26,D63,R33,U87,L62,D20,R33,U53,R51",
			inputB:         "U98,R91,D20,R16,D67,R40,U7,R15,U6,R7",
			distFunc:       manhattan(),
			expectDistance: 135,
		},
		{
			name:           "Ex3 Steps",
			inputA:         "R98,U47,R26,D63,R33,U87,L62,D20,R33,U53,R51",
			inputB:         "U98,R91,D20,R16,D67,R40,U7,R15,U6,R7",
			distFunc:       steps(),
			expectDistance: 410,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			wireA, err := parseWire(tt.inputA)
			require.NoError(t, err)

			wireB, err := parseWire(tt.inputB)
			require.NoError(t, err)

			assert.Equal(t, tt.expectDistance, wireA.ClosestIntersection(wireB, tt.distFunc))
		})
	}
}
