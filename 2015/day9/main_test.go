package main

import "testing"

import "github.com/stretchr/testify/assert"

func TestPermutations(t *testing.T) {
	tests := []struct {
		name               string
		cities             []string
		expectPermutations [][]string
	}{
		{
			name: "Ex1",
			cities: []string{
				"London",
				"Dublin",
				"Belfast",
			},
			expectPermutations: [][]string{
				[]string{
					"Dublin", "London", "Belfast",
				},
				[]string{
					"London", "Dublin", "Belfast",
				},
				[]string{
					"London", "Belfast", "Dublin",
				},
				[]string{
					"Dublin", "Belfast", "London",
				},
				[]string{
					"Belfast", "Dublin", "London",
				},
				[]string{
					"Belfast", "London", "Dublin",
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.ElementsMatch(t, tt.expectPermutations, permutations(tt.cities))
		})
	}
}
