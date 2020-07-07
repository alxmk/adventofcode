package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.ibm.com/alexmk/adventofcode/2019/day2/intcode"
)

func TestRunAmps(t *testing.T) {
	tests := []struct {
		name        string
		programme   string
		phases      seq
		expectPower int64
	}{
		{
			name:        "Ex1",
			programme:   "3,15,3,16,1002,16,10,16,1,16,15,15,4,15,99,0,0",
			phases:      seq{4, 3, 2, 1, 0},
			expectPower: 43210,
		},
		{
			name:        "Ex2",
			programme:   "3,23,3,24,1002,24,10,24,1002,23,-1,23,101,5,23,23,1,24,23,23,4,23,99,0,0",
			phases:      seq{0, 1, 2, 3, 4},
			expectPower: 54321,
		},
		{
			name:        "Ex3",
			programme:   "3,31,3,32,1002,32,10,32,1001,31,-2,31,1007,31,0,33,1002,33,7,33,1,33,31,31,1,32,31,31,4,31,99,0,0,0",
			phases:      seq{1, 0, 4, 3, 2},
			expectPower: 65210,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p, err := intcode.Parse(tt.programme)
			require.NoError(t, err)

			assert.Equal(t, tt.expectPower, runAmps(tt.phases, p))
		})
	}
}

func TestRunAmpsFeedback(t *testing.T) {
	tests := []struct {
		name        string
		programme   string
		phases      seq
		expectPower int64
	}{
		{
			name:        "Ex1",
			programme:   "3,26,1001,26,-4,26,3,27,1002,27,2,27,1,27,26,27,4,27,1001,28,-1,28,1005,28,6,99,0,0,5",
			phases:      seq{9, 8, 7, 6, 5},
			expectPower: 139629729,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p, err := intcode.Parse(tt.programme)
			require.NoError(t, err)

			assert.Equal(t, tt.expectPower, runAmpsFeedback(tt.phases, p))
		})
	}
}

func TestGenerateSequences(t *testing.T) {
	tests := []struct {
		name            string
		values          []int64
		expectSequences []seq
	}{
		{
			name:   "Ex1",
			values: []int64{1, 2, 3},
			expectSequences: []seq{
				seq{1, 2, 3},
				seq{1, 3, 2},
				seq{2, 1, 3},
				seq{2, 3, 1},
				seq{3, 1, 2},
				seq{3, 2, 1},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.ElementsMatch(t, generateSequences(tt.values), tt.expectSequences)
		})
	}
}
