package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCalculate(t *testing.T) {
	tests := []struct {
		name   string
		input  string
		length int
		expect int
	}{
		{
			name: "Example",
			input: `R 4
U 4
L 3
D 1
R 4
D 1
L 5
R 2`,
			length: 2,
			expect: 13,
		},
		{
			name: "Example 2",
			input: `R 4
U 4
L 3
D 1
R 4
D 1
L 5
R 2`,
			length: 10,
			expect: 1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expect, calculate(tt.input, tt.length))
		})
	}
}
