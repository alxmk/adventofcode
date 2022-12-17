package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGet(t *testing.T) {
	tests := []struct {
		name   string
		gas    gas
		i      int
		expect byte
	}{
		{
			name:   "Idx 0",
			gas:    []byte{'>', '<', '>', '<'},
			i:      0,
			expect: '>',
		},
		{
			name:   "Idx 2",
			gas:    []byte{'>', '<', '>', '<'},
			i:      2,
			expect: '>',
		},
		{
			name:   "Idx 3",
			gas:    []byte{'>', '<', '>', '<'},
			i:      3,
			expect: '<',
		},
		{
			name:   "Idx 4",
			gas:    []byte{'>', '<', '>', '<'},
			i:      4,
			expect: '>',
		},
		{
			name:   "Idx 5",
			gas:    []byte{'>', '<', '>', '<'},
			i:      5,
			expect: '<',
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expect, Get(tt.gas, tt.i))
		})
	}
}

func TestPartOne(t *testing.T) {
	tests := []struct {
		name   string
		input  string
		expect int
	}{
		{
			name:   "Example",
			input:  ">>><<><>><<<>><>>><<<>>><<<><<<>><>><<>>",
			expect: 3068,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expect, partOne([]byte(tt.input)))
		})
	}
}

func TestSolve(t *testing.T) {
	tests := []struct {
		name               string
		input              string
		head, delta, turns int
		expect             int
	}{
		{
			name:  "Example 2022",
			input: ">>><<><>><<<>><>>><<<>>><<<><<<>><>><<>>",
			head:  34,
			delta: 35,
			turns: 2022,
			// expect: 1514285714288,
			expect: 3067,
		},
		{
			name:   "Example 1000000000000",
			input:  ">>><<><>><<<>><>>><<<>>><<<><<<>><>><<>>",
			head:   34,
			delta:  35,
			turns:  1000000000000,
			expect: 1514285714287,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expect, solve([]byte(tt.input), tt.head, tt.delta, tt.turns))
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
			name:   "Example 1000000000000",
			input:  ">>><<><>><<<>><>>><<<>>><<<><<<>><>><<>>",
			expect: 1514285714288,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expect, partTwo([]byte(tt.input)))
		})
	}
}

func TestFindLoop(t *testing.T) {
	tests := []struct {
		name        string
		input       string
		expectIdx   int
		expectDelta int
	}{
		{
			name:        "Example",
			input:       ">>><<><>><<<>><>>><<<>>><<<><<<>><>><<>>",
			expectIdx:   15,
			expectDelta: 35,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actualIdx, actualDelta := findLoop([]byte(tt.input))
			assert.Equal(t, tt.expectIdx, actualIdx)
			assert.Equal(t, tt.expectDelta, actualDelta)
		})
	}
}

func TestRepeats(t *testing.T) {
	tests := []struct {
		name    string
		deltas  []int
		pattern []int
		expect  int
	}{
		{
			name:    "Example 1",
			deltas:  []int{0, 1, 2, 3, 4, 5, 6, 2, 3, 4, 5, 6},
			pattern: []int{2, 3, 4, 5, 6},
			expect:  2,
		},
		{
			name:    "Example 2",
			deltas:  []int{0, 1, 2, 3, 4, 5, 6, 2, 3, 4, 5, 6},
			pattern: []int{4, 5, 6},
			expect:  4,
		},
		{
			name:    "Example 3",
			deltas:  []int{1, 3, 2, 1, 2, 1, 3, 2, 2, 0, 1, 3, 2, 0, 2, 1, 3, 3, 4, 0, 1, 2, 3, 0, 1, 1, 3, 2, 2, 0, 0, 2, 3, 4, 0, 1, 2, 1, 2, 0, 1, 2, 1, 2, 0, 1, 3, 2, 0, 0, 1, 3, 3, 4, 0, 1, 2, 3, 0, 1, 1, 3, 2, 2, 0, 0, 2, 3, 4, 0, 1, 2, 1, 2},
			pattern: []int{0, 1, 2, 1, 2, 0, 1, 3, 2, 0, 0, 1, 3, 3, 4, 0, 1, 2, 3, 0, 1, 1, 3, 2, 2, 0, 0, 2, 3, 4},
			expect:  -1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, actual := repeats(tt.deltas, tt.pattern)
			assert.Equal(t, tt.expect, actual)
		})
	}
}
