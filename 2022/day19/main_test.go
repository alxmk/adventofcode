package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPopulateTree(t *testing.T) {
	tests := []struct {
		name  string
		input string
	}{
		{
			name:  "Example 1",
			input: example1,
		},
		{
			name:  "Example 2",
			input: example2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			parse(tt.input)[0].quality()
			t.Fail()
		})
	}
}

var example1 = `Blueprint 1: Each ore robot costs 4 ore. Each clay robot costs 2 ore. Each obsidian robot costs 3 ore and 14 clay. Each geode robot costs 2 ore and 7 obsidian.`
var example2 = `Blueprint 2: Each ore robot costs 2 ore. Each clay robot costs 3 ore. Each obsidian robot costs 3 ore and 8 clay. Each geode robot costs 3 ore and 12 obsidian.`

func TestMinsToBuild(t *testing.T) {
	tests := []struct {
		name            string
		input           string
		resource        resource
		state           *state
		expectBuildable bool
		expectTime      int
	}{
		{
			name:            "Example 1 clay",
			input:           example1,
			resource:        clay,
			state:           &state{robots: rset{ore: 1}},
			expectBuildable: true,
			expectTime:      3,
		},
		{
			name:            "Example 1 ore",
			input:           example1,
			resource:        ore,
			state:           &state{robots: rset{ore: 1}},
			expectBuildable: true,
			expectTime:      5,
		},
		{
			name:            "Example 1 turn 7",
			input:           example1,
			resource:        obsidian,
			state:           &state{robots: rset{ore: 1, clay: 3}, resources: rset{ore: 1, clay: 6}},
			expectBuildable: true,
			expectTime:      4,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actualBuildable, actualTime := tt.state.MinsToBuild(tt.resource, parse(tt.input)[0])
			assert.Equal(t, tt.expectBuildable, actualBuildable)
			assert.Equal(t, tt.expectTime, actualTime)
		})
	}
}
