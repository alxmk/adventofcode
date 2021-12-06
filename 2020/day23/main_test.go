package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMove(t *testing.T) {
	tests := []struct {
		name   string
		input  string
		moves  int
		expect string
	}{
		{
			name:   "Test parsing",
			input:  "389125467",
			moves:  0,
			expect: "389125467",
		},
		{
			name:   "1 moves",
			input:  "389125467",
			moves:  1,
			expect: "289154673",
		},
		{
			name:   "2 moves",
			input:  "389125467",
			moves:  2,
			expect: "546789132",
		},
		{
			name:   "10 moves",
			input:  "389125467",
			moves:  10,
			expect: "837419265",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cups := parse(tt.input)
			for i := 0; i < tt.moves; i++ {
				cups.Move()
			}
			assert.Equal(t, tt.expect, cups.String())
		})
	}
}

func TestMoveFromOne(t *testing.T) {
	tests := []struct {
		name   string
		input  string
		moves  int
		expect string
	}{
		{
			name:   "10 moves",
			input:  "389125467",
			moves:  10,
			expect: "92658374",
		},
		{
			name:   "100 moves",
			input:  "389125467",
			moves:  100,
			expect: "67384529",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cups := parse(tt.input)
			for i := 0; i < tt.moves; i++ {
				cups.Move()
			}
			assert.Equal(t, tt.expect, cups.FromOne())
		})
	}
}

func TestParseTwo(t *testing.T) {
	tests := []struct {
		name   string
		input  string
		limit  int
		expect string
	}{
		{
			name:   "Test parsing",
			input:  "389125467",
			limit:  10,
			expect: "3891254671011",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cups := parsePartTwo(tt.input)
			assert.Equal(t, tt.expect, cups.StringLimit(tt.limit))
		})
	}
}

func TestMoveFromOneLimit(t *testing.T) {
	tests := []struct {
		name   string
		input  string
		moves  int
		limit  int
		expect string
	}{
		{
			name:  "100 moves limit",
			input: "389125467",
			moves: 100,
			limit: 3,
			expect: `6
7
3
`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cups := parse(tt.input)
			for i := 0; i < tt.moves; i++ {
				cups.Move()
			}
			assert.Equal(t, tt.expect, cups.FromOneLimit(tt.limit))
		})
	}
}
