package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_generatePairs(t *testing.T) {
	tests := []struct {
		name        string
		moons       []*moon
		expectPairs []pair
	}{
		{
			name: "Ex1",
			moons: []*moon{
				&moon{x: 7, y: 10, z: 17},
				&moon{x: -2, y: 7, z: 0},
				&moon{x: 12, y: 5, z: 12},
				&moon{x: 5, y: -8, z: 6},
			},
			expectPairs: []pair{
				pair{&moon{x: 7, y: 10, z: 17}, &moon{x: -2, y: 7, z: 0}},
				pair{&moon{x: 7, y: 10, z: 17}, &moon{x: 12, y: 5, z: 12}},
				pair{&moon{x: 7, y: 10, z: 17}, &moon{x: 5, y: -8, z: 6}},
				pair{&moon{x: -2, y: 7, z: 0}, &moon{x: 12, y: 5, z: 12}},
				pair{&moon{x: -2, y: 7, z: 0}, &moon{x: 5, y: -8, z: 6}},
				pair{&moon{x: 12, y: 5, z: 12}, &moon{x: 5, y: -8, z: 6}},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.ElementsMatch(t, tt.expectPairs, generatePairs(tt.moons))
		})
	}
}

func TestSimulate(t *testing.T) {
	tests := []struct {
		name         string
		input        string
		steps        int
		expectEnergy int
	}{
		{
			name: "Ex1",
			input: `<x=-1, y=0, z=2>
<x=2, y=-10, z=-7>
<x=4, y=-8, z=8>
<x=3, y=5, z=-1>`,
			steps:        10,
			expectEnergy: 179,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expectEnergy, simulate(parseMoons(tt.input), tt.steps))
		})
	}
}

func TestSearchForRepeat(t *testing.T) {
	tests := []struct {
		name        string
		input       string
		expectSteps int
	}{
		{
			name: "Ex1",
			input: `<x=-1, y=0, z=2>
<x=2, y=-10, z=-7>
<x=4, y=-8, z=8>
<x=3, y=5, z=-1>`,
			expectSteps: 2772,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expectSteps, searchForRepeat(parseMoons(tt.input)))
		})
	}
}

func TestPrimeFactors(t *testing.T) {
	tests := []struct {
		name               string
		n                  int
		expectPrimeFactors map[int]int
	}{
		{
			name:               "1",
			n:                  1,
			expectPrimeFactors: map[int]int{},
		},
		{
			name:               "2",
			n:                  2,
			expectPrimeFactors: map[int]int{2: 1},
		},
		{
			name:               "3",
			n:                  3,
			expectPrimeFactors: map[int]int{3: 1},
		},
		{
			name:               "4",
			n:                  4,
			expectPrimeFactors: map[int]int{2: 2},
		},
		{
			name:               "18",
			n:                  18,
			expectPrimeFactors: map[int]int{2: 1, 3: 2},
		},
		{
			name:               "28",
			n:                  28,
			expectPrimeFactors: map[int]int{2: 2, 7: 1},
		},
		{
			name:               "30",
			n:                  30,
			expectPrimeFactors: map[int]int{2: 1, 3: 1, 5: 1},
		},
		{
			name:               "44",
			n:                  44,
			expectPrimeFactors: map[int]int{2: 2, 11: 1},
		},
		{
			name:               "45",
			n:                  45,
			expectPrimeFactors: map[int]int{3: 2, 5: 1},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expectPrimeFactors, primeFactors(tt.n))
		})
	}
}

func TestLowestCommonMultiple(t *testing.T) {
	tests := []struct {
		name      string
		x, y, z   int
		expectLCM int
	}{
		{
			name:      "Ex1",
			x:         18,
			y:         28,
			z:         44,
			expectLCM: 2772,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expectLCM, lowestCommonMultiple(tt.x, tt.y, tt.z))
		})
	}
}
