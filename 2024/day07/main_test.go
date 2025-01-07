package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPartOne(t *testing.T) {
	tests := []struct {
		name      string
		input     string
		expectOne int
		expectTwo int
	}{
		{
			name: "Ex1",
			input: `190: 10 19
3267: 81 40 27
83: 17 5
156: 15 6
7290: 6 8 6 15
161011: 16 10 13
192: 17 8 14
21037: 9 7 18 13
292: 11 6 16 20`,
			expectOne: 3749,
			expectTwo: 11387,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			one, two := exec(tt.input)
			assert.Equal(t, tt.expectOne, one)
			assert.Equal(t, tt.expectTwo, two)
		})
	}
}
