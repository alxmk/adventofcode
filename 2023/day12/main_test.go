package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPermutations(t *testing.T) {
	tests := []struct {
		name   string
		r      string
		expect int
	}{
		{
			name:   "Example 1",
			r:      "???.### 1,1,3",
			expect: 1,
		},
		{
			name:   "Example 2",
			r:      ".??..??...?##. 1,1,3",
			expect: 4,
		},
		{
			name:   "Example 3",
			r:      "?###???????? 3,2,1",
			expect: 10,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expect, parseRow(tt.r).Permutations(0, 0, 0, make(map[[3]int]int)))
		})
	}
}

func TestUnfoldedPermutations(t *testing.T) {
	tests := []struct {
		name   string
		r      string
		expect int
	}{
		{
			name:   "Example 1",
			r:      "???.### 1,1,3",
			expect: 1,
		},
		{
			name:   "Example 2",
			r:      ".??..??...?##. 1,1,3",
			expect: 16384,
		},
		{
			name:   "Example 3",
			r:      "?###???????? 3,2,1",
			expect: 506250,
		},
		{
			name:   "Example 4",
			r:      "????.#...#... 4,1,1",
			expect: 16,
		},
		{
			name:   "Example 5",
			r:      "?#?#?#?#?#?#?#? 1,3,1,6",
			expect: 1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expect, parseRow(tt.r).Unfold().Permutations(0, 0, 0, make(map[[3]int]int)))
		})
	}
}
