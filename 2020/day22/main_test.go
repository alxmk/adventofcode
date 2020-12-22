package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPlayCombat(t *testing.T) {
	tests := []struct {
		name   string
		input  string
		expect int
	}{
		{
			name: "Ex1",
			input: `Player 1:
9
2
6
3
1

Player 2:
5
8
4
7
10`,
			expect: 306,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expect, playCombat(parse(tt.input)).Score())
		})
	}
}

func TestPlayRecursiveCombat(t *testing.T) {
	tests := []struct {
		name   string
		input  string
		expect int
	}{
		{
			name: "Ex1",
			input: `Player 1:
9
2
6
3
1

Player 2:
5
8
4
7
10`,
			expect: 291,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, winner := playRecursiveCombat(parse(tt.input))
			assert.Equal(t, tt.expect, winner.Score())
			t.FailNow()
		})
	}
}
