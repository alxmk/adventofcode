package main

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestInRoom(t *testing.T) {
	tests := []struct {
		name   string
		c      coord
		expect bool
	}{
		{
			name:   "Room A lower",
			c:      coord{3, 3},
			expect: true,
		},
		{
			name:   "Room A upper",
			c:      coord{3, 2},
			expect: true,
		},
		{
			name:   "Room B lower",
			c:      coord{5, 3},
			expect: true,
		},
		{
			name:   "Room B upper",
			c:      coord{5, 2},
			expect: true,
		},
		{
			name:   "Room C lower",
			c:      coord{7, 3},
			expect: true,
		},
		{
			name:   "Room C upper",
			c:      coord{7, 2},
			expect: true,
		},
		{
			name:   "Room D lower",
			c:      coord{9, 3},
			expect: true,
		},
		{
			name:   "Room D upper",
			c:      coord{9, 2},
			expect: true,
		},
		{
			name:   "Corridor",
			c:      coord{8, 1},
			expect: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expect, tt.c.InRoom())
		})
	}
}

func TestInCorrectRoom(t *testing.T) {
	tests := []struct {
		name   string
		c      coord
		r      rune
		expect bool
	}{
		{
			name:   "Room A lower, is A",
			c:      coord{3, 3},
			r:      'A',
			expect: true,
		},
		{
			name:   "Room A upper, is A",
			c:      coord{3, 2},
			r:      'A',
			expect: true,
		},
		{
			name:   "Room A lower, is B",
			c:      coord{3, 3},
			r:      'B',
			expect: false,
		},
		{
			name:   "Room A upper, is B",
			c:      coord{3, 2},
			r:      'B',
			expect: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expect, tt.c.InCorrectRoom(tt.r))
		})
	}
}

func TestValidMovesInner(t *testing.T) {
	tests := []struct {
		name            string
		input           string
		c               coord
		r               rune
		expect          map[move]int
		expectFoundBest bool
	}{
		{
			name:            "Already in own room",
			input:           sampleInput,
			c:               coord{3, 3},
			r:               'A',
			expect:          nil,
			expectFoundBest: false,
		},
		{
			name:            "Move from corridor to own room",
			input:           sampleInputCorridorToRoom,
			c:               coord{4, 1},
			r:               'B',
			expect:          map[move]int{{coord{4, 1}, coord{5, 3}}: 30},
			expectFoundBest: true,
		},
		{
			name:  "Move from room to corridor",
			input: sampleInput,
			c:     coord{3, 2},
			r:     'B',
			expect: map[move]int{
				{coord{3, 2}, coord{1, 1}}:  30,
				{coord{3, 2}, coord{2, 1}}:  20,
				{coord{3, 2}, coord{4, 1}}:  20,
				{coord{3, 2}, coord{6, 1}}:  40,
				{coord{3, 2}, coord{8, 1}}:  60,
				{coord{3, 2}, coord{10, 1}}: 80,
				{coord{3, 2}, coord{11, 1}}: 90,
			},
			expectFoundBest: false,
		},
		{
			name:            "Move from room to room",
			input:           sampleInputMoveRoomToRoom,
			c:               coord{5, 2},
			r:               'C',
			expect:          map[move]int{{coord{5, 2}, coord{7, 2}}: 400},
			expectFoundBest: true,
		},
		{
			name:            "Winning move",
			input:           sampleInputWinningMove,
			c:               coord{10, 1},
			r:               'A',
			expect:          map[move]int{{coord{10, 1}, coord{3, 2}}: 8},
			expectFoundBest: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			state, err := parse(tt.input)
			require.NoError(t, err)

			actual, foundBest := state.validMovesInner(tt.c, tt.r, 0, math.MaxInt)
			assert.Equal(t, tt.expect, actual)
			assert.Equal(t, tt.expectFoundBest, foundBest)
		})
	}
}

func TestDistance(t *testing.T) {
	tests := []struct {
		name   string
		c, d   coord
		expect int
	}{
		{
			name:   "Room to room",
			c:      coord{3, 3},
			d:      coord{5, 3},
			expect: 6,
		},
		{
			name:   "Room to room closer",
			c:      coord{3, 2},
			d:      coord{5, 3},
			expect: 5,
		},
		{
			name:   "Room to room closerer",
			c:      coord{3, 2},
			d:      coord{5, 2},
			expect: 4,
		},
		{
			name:   "Corridor to room",
			c:      coord{2, 1},
			d:      coord{5, 2},
			expect: 4,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expect, tt.c.Distance(tt.d))
		})
	}
}

var sampleInput = `#############
#...........#
###B#C#B#D###
  #A#D#C#A#
  #########`

var sampleInputCorridorToRoom = `#############
#...B.D.....#
###B#.#C#D###
  #A#.#C#A#
  #########`

var sampleInputMoveRoomToRoom = `#############
#...B.......#
###B#C#.#D###
  #A#D#C#A#
  #########`

var sampleInputWinningMove = `#############
#.........A.#
###.#B#C#D###
  #A#B#C#D#
  #########`
