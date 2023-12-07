package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolve(t *testing.T) {
	tests := []struct {
		name   string
		jokers bool
		input  string
		expect int
	}{
		{
			name:   "Example 1",
			jokers: false,
			input: `32T3K 765
T55J5 684
KK677 28
KTJJT 220
QQQJA 483`,
			expect: 6440,
		},
		{
			name:   "Example 2",
			jokers: true,
			input: `32T3K 765
T55J5 684
KK677 28
KTJJT 220
QQQJA 483`,
			expect: 5905,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expect, solve(parseHands(tt.input), tt.jokers))
		})
	}

}
