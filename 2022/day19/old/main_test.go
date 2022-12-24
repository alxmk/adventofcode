package main

import (
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBlueprintQuality(t *testing.T) {
	tests := []struct {
		name   string
		input  string
		expect int
	}{
		{
			name:   "Example 1",
			input:  example1,
			expect: 9,
		},
		{
			name:   "Example 2",
			input:  example2,
			expect: 24,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expect, partOne(parse(tt.input)))
		})
	}
}

var example1 = `Blueprint 1: Each ore robot costs 4 ore. Each clay robot costs 2 ore. Each obsidian robot costs 3 ore and 14 clay. Each geode robot costs 2 ore and 7 obsidian.`
var example2 = `Blueprint 2: Each ore robot costs 2 ore. Each clay robot costs 3 ore. Each obsidian robot costs 3 ore and 8 clay. Each geode robot costs 3 ore and 12 obsidian.`

func TestEarliestGeode(t *testing.T) {
	tests := []struct {
		name   string
		input  string
		expect int
	}{
		{
			name:   "Example 1",
			input:  example1,
			expect: 17,
		},
		{
			name:   "Example 2",
			input:  example2,
			expect: 18,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual, mrs := parse(tt.input)[0].earliestGeode(0, &state{robots: resources{ore: 1}}, clay)
			assert.Equal(t, tt.expect, actual)
			log.Println(len(mrs))
			log.Println(j)
		})
	}
}
