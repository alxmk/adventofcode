package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStateBuildTime(t *testing.T) {
	tests := []struct {
		name   string
		r      resources
		s      state
		expect int
	}{
		{
			name:   "Example 1",
			r:      resources{ore: 4},
			s:      state{robots: resources{ore: 1}},
			expect: 5,
		},
		{
			name:   "Example 2",
			r:      resources{ore: 2},
			s:      state{robots: resources{ore: 1}},
			expect: 3,
		},
		{
			name:   "Example 3",
			r:      resources{ore: 3, clay: 14},
			s:      state{robots: resources{ore: 1, clay: 3}, inventory: resources{ore: 1, clay: 6}},
			expect: 4,
		},
		{
			name:   "Example 4",
			r:      resources{ore: 3, clay: 14},
			s:      state{robots: resources{ore: 1, clay: 3}, inventory: resources{ore: 1, clay: 2}},
			expect: 5,
		},
		//2022/12/24 19:46:45 {20 {7 26 6 0} {4 10 2 0}}
		{
			name:   "Example 5",
			r:      resources{ore: 3, clay: 17},
			s:      state{robots: resources{ore: 4, clay: 10, obsidian: 2}, inventory: resources{ore: 7, clay: 26, obsidian: 6}},
			expect: 1,
		},
		{
			name:   "Example 6",
			r:      resources{ore: 3, obsidian: 8},
			s:      state{robots: resources{ore: 4, clay: 10, obsidian: 2}, inventory: resources{ore: 4, clay: 19, obsidian: 8}},
			expect: 1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expect, tt.s.BuildTime(tt.r))
		})
	}
}

func TestSolverSolve(t *testing.T) {
	tests := []struct {
		name   string
		s      *solver
		time   int
		expect int
	}{
		{
			name: "Example 1 24",
			s: &solver{
				scache: make(map[state]struct{}),
				bp: blueprint{
					number:   1,
					ore:      resources{ore: 4},
					clay:     resources{ore: 2},
					obsidian: resources{ore: 3, clay: 14},
					geode:    resources{ore: 2, obsidian: 7},
				},
			},
			time:   24,
			expect: 9,
		},
		{
			name: "Example 2 24",
			s: &solver{
				scache: make(map[state]struct{}),
				bp: blueprint{
					number:   2,
					ore:      resources{ore: 2},
					clay:     resources{ore: 3},
					obsidian: resources{ore: 3, clay: 8},
					geode:    resources{ore: 3, obsidian: 12},
				},
			},
			time:   24,
			expect: 12,
		},
		{
			name: "Example 1 32",
			s: &solver{
				scache: make(map[state]struct{}),
				bp: blueprint{
					number:   1,
					ore:      resources{ore: 4},
					clay:     resources{ore: 2},
					obsidian: resources{ore: 3, clay: 14},
					geode:    resources{ore: 2, obsidian: 7},
				},
			},
			time:   32,
			expect: 56,
		},
		{
			name: "Example 2 32",
			s: &solver{
				scache: make(map[state]struct{}),
				bp: blueprint{
					number:   2,
					ore:      resources{ore: 2},
					clay:     resources{ore: 3},
					obsidian: resources{ore: 3, clay: 8},
					geode:    resources{ore: 3, obsidian: 12},
				},
			},
			time:   32,
			expect: 62,
		},
		{
			name: "Real 2 32",
			s: &solver{
				scache: make(map[state]struct{}),
				bp: blueprint{
					number:   2,
					ore:      resources{ore: 3},
					clay:     resources{ore: 4},
					obsidian: resources{ore: 3, clay: 17},
					geode:    resources{ore: 3, obsidian: 8},
				},
			},
			time:   32,
			expect: 31,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expect, tt.s.Solve(tt.time))
		})
	}
}

func TestStateLess(t *testing.T) {
	tests := []struct {
		name   string
		r, o   state
		expect bool
	}{
		{
			name:   "Example 1",
			r:      state{},
			o:      state{},
			expect: true,
		},
		{
			name:   "Example 2",
			r:      state{t: 1},
			o:      state{},
			expect: true,
		},
		{
			name:   "Example 3",
			r:      state{},
			o:      state{t: 1},
			expect: false,
		},
		{
			name:   "Example 4",
			r:      state{robots: resources{ore: 1}},
			o:      state{},
			expect: false,
		},
		{
			name:   "Example 4",
			r:      state{robots: resources{ore: 1}},
			o:      state{robots: resources{clay: 1}},
			expect: false,
		},
		{
			name:   "Example 5",
			r:      state{robots: resources{ore: 1}},
			o:      state{inventory: resources{clay: 1}},
			expect: false,
		},
	}

	for _, tt := range tests {
		assert.Equal(t, tt.expect, tt.r.Less(tt.o))
	}
}
