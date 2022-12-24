package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPartOne(t *testing.T) {
	tests := []struct {
		name   string
		input  string
		expect int
	}{
		{
			name: "Example",
			input: `#.######
#>>.<^<#
#.<..<<#
#>v.><>#
#<^v^^>#
######.#`,
			expect: 18,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expect, partOne(parse([]byte(tt.input))))
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
			name: "Example",
			input: `#.######
#>>.<^<#
#.<..<<#
#>v.><>#
#<^v^^>#
######.#`,
			expect: 54,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expect, partTwo(parse([]byte(tt.input))))
		})
	}
}
