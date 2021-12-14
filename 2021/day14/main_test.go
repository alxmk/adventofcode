package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestStep(t *testing.T) {
	tests := []struct {
		name   string
		input  string
		steps  int
		expect *pairs
	}{
		{
			name:   "Example 1",
			input:  sampleInput,
			steps:  1,
			expect: &pairs{all: map[string]int64{"NC": 1, "CN": 1, "NB": 1, "BC": 1, "CH": 1, "HB": 1}, head: "NC", tail: "HB"}, //"NCNBCHB"
		},
		{
			name:   "Example 2",
			input:  sampleInput,
			steps:  2,
			expect: stringToPairs("NBCCNBBBCBHCB"),
		},
		{
			name:   "Example 3",
			input:  sampleInput,
			steps:  3,
			expect: stringToPairs("NBBBCNCCNBBNBNBBCHBHHBCHB"),
		},
		{
			name:   "Example 4",
			input:  sampleInput,
			steps:  4,
			expect: stringToPairs("NBBNBNBBCCNBCNCCNBBNBBNBBBNBBNBBCBHCBHHNHCBBCBHCB"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			polymer, rules, err := parse(tt.input)
			require.NoError(t, err)
			p := stringToPairs(polymer)
			for i := 0; i < tt.steps; i++ {
				p = step(p, rules)
			}
			assert.Equal(t, tt.expect, p)
		})
	}
}

func TestEvaluate(t *testing.T) {
	tests := []struct {
		name   string
		input  string
		steps  int
		expect int64
	}{
		{
			name:   "Example p1",
			input:  sampleInput,
			steps:  10,
			expect: 1588,
		},
		{
			name:   "Example p2",
			input:  sampleInput,
			steps:  40,
			expect: 2188189693529,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			polymer, rules, err := parse(tt.input)
			require.NoError(t, err)
			assert.Equal(t, tt.expect, evaluate(polymer, rules, tt.steps))
		})
	}
}

var sampleInput = `NNCB

CH -> B
HH -> N
CB -> H
NH -> C
HB -> C
HC -> B
HN -> C
NN -> C
BH -> H
NC -> B
NB -> B
BN -> B
BB -> N
BC -> B
CC -> N
CN -> C`
