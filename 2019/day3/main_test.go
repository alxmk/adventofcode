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
				// coordinate{0, 0}: 0,
				coordinate{1, 0}: 1,
				coordinate{2, 0}: 2,
				coordinate{3, 0}: 3,
				coordinate{4, 0}: 4,
				coordinate{5, 0}: 5,
				coordinate{6, 0}: 6,
				coordinate{7, 0}: 7,
				coordinate{8, 0}: 8,
				coordinate{8, 1}: 9,
				coordinate{8, 2}: 10,
				coordinate{8, 3}: 11,
				coordinate{8, 4}: 12,
				coordinate{8, 5}: 13,
				coordinate{7, 5}: 14,
				coordinate{6, 5}: 15,
				coordinate{5, 5}: 16,
				coordinate{4, 5}: 17,
				coordinate{3, 5}: 18,
				coordinate{3, 4}: 19,
				coordinate{3, 3}: 20,
				coordinate{3, 2}: 21,
			},
		},
		{
			name:  "Ex2",
			input: "U7,R6,D4,L4",
			expect: wire{
				// coordinate{0, 0}: 0,
				coordinate{0, 1}: 1,
				coordinate{0, 2}: 2,
				coordinate{0, 3}: 3,
				coordinate{0, 4}: 4,
				coordinate{0, 5}: 5,
				coordinate{0, 6}: 6,
				coordinate{0, 7}: 7,
				coordinate{1, 7}: 8,
				coordinate{2, 7}: 9,
				coordinate{3, 7}: 10,
				coordinate{4, 7}: 11,
				coordinate{5, 7}: 12,
				coordinate{6, 7}: 13,
				coordinate{6, 6}: 14,
				coordinate{6, 5}: 15,
				coordinate{6, 4}: 16,
				coordinate{6, 3}: 17,
				coordinate{5, 3}: 18,
				coordinate{4, 3}: 19,
				coordinate{3, 3}: 20,
				coordinate{2, 3}: 21,
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
		distFunc       func(coordinate, int, int) int
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
		// {
		// 	name:           "Ex2 Steps",
		// 	inputA:         "R75,D30,R83,U83,L12,D49,R71,U7,L72",
		// 	inputB:         "U62,R66,U55,R34,D71,R55,D58,R83",
		// 	distFunc:       steps(),
		// 	expectDistance: 610,
		// },
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

			assert.Equal(t, tt.expectDistance, closestIntersection(wireA, wireB, tt.distFunc))

			if tt.name == "Ex2 Steps" {
				t.Fail()
			}
		})
	}
}
