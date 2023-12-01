package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCalculate(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		partTwo bool
		expect  int
	}{
		{
			name: "Example 1",
			input: `1abc2
pqr3stu8vwx
a1b2c3d4e5f
treb7uchet`,
			expect: 142,
		},
		{
			name: "Example 2",
			input: `two1nine
eightwothree
abcone2threexyz
xtwone3four
4nineeightseven2
zoneight234
7pqrstsixteen`,
			partTwo: true,
			expect:  281,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expect, calculate(tt.input, tt.partTwo))
		})
	}
}
