package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDistance(t *testing.T) {
	tests := []struct {
		name   string
		input  string
		a, b   string
		expect int
	}{
		{
			name:   "Example 1",
			input:  example,
			a:      "AA",
			b:      "DD",
			expect: 1,
		},
		{
			name:   "Example 2",
			input:  example,
			a:      "DD",
			b:      "BB",
			expect: 2,
		},
		{
			name:   "Example 3",
			input:  example,
			a:      "BB",
			b:      "JJ",
			expect: 3,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			n := parse(tt.input)
			assert.Equal(t, tt.expect, n.distance(tt.a, tt.b))
		})
	}
}

var example = `Valve AA has flow rate=0; tunnels lead to valves DD, II, BB
Valve BB has flow rate=13; tunnels lead to valves CC, AA
Valve CC has flow rate=2; tunnels lead to valves DD, BB
Valve DD has flow rate=20; tunnels lead to valves CC, AA, EE
Valve EE has flow rate=3; tunnels lead to valves FF, DD
Valve FF has flow rate=0; tunnels lead to valves EE, GG
Valve GG has flow rate=0; tunnels lead to valves FF, HH
Valve HH has flow rate=22; tunnel leads to valve GG
Valve II has flow rate=0; tunnels lead to valves AA, JJ
Valve JJ has flow rate=21; tunnel leads to valve II`

func TestPartOne(t *testing.T) {
	tests := []struct {
		name   string
		input  string
		expect int
	}{
		{
			name:   "Example",
			input:  example,
			expect: 1651,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expect, partOne(parse(tt.input)))
		})
	}
}

func TestPartTwo(t *testing.T) {
	tests := []struct {
		name   string
		input  string
		expect int
	}{
		{
			name:   "Example",
			input:  example,
			expect: 1707,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expect, partTwo(parse(tt.input)))
		})
	}
}

func TestOpposite(t *testing.T) {
	tests := []struct {
		name        string
		input       string
		combination []int
		expect      []int
	}{
		{
			name:        "Example 1",
			input:       example,
			combination: []int{0, 1, 2},
			expect:      []int{3, 4, 5},
		},
		{
			name:        "Example 2",
			input:       example,
			combination: []int{2, 3, 4},
			expect:      []int{0, 1, 5},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.ElementsMatch(t, tt.expect, parse(tt.input).opposite(tt.combination))
		})
	}
}
