package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSumBasic(t *testing.T) {
	tests := []struct {
		name   string
		input  string
		expect int64
	}{
		{
			name:   "Ex1",
			input:  "1 + 2 * 3 + 4 * 5 + 6",
			expect: 71,
		},
		{
			name:   "Ex2",
			input:  "1 + (2 * 3) + (4 * (5 + 6))",
			expect: 51,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expect, sumBasic(tt.input))
		})
	}
}

func TestSumAdvanced(t *testing.T) {
	tests := []struct {
		name   string
		input  string
		expect int64
	}{
		{
			name:   "Ex1",
			input:  "1 + 2 * 3 + 4 * 5 + 6",
			expect: 231,
		},
		{
			name:   "Ex2",
			input:  "1 + (2 * 3) + (4 * (5 + 6))",
			expect: 51,
		},
		{
			name:   "Ex3",
			input:  "2 * 3 + (4 * 5)",
			expect: 46,
		},
		{
			name:   "Ex4",
			input:  "5 + (8 * 3 + 9 + 3 * 4 * 3)",
			expect: 1445,
		},
		{
			name:   "Ex5",
			input:  "5 * 9 * (7 * 3 * 3 + 9 * 3 + (8 + 6 * 4))",
			expect: 669060,
		},
		{
			name:   "Ex6",
			input:  "((2 + 4 * 9) * (6 + 9 * 8 + 6) + 6) + 2 + 4 * 2",
			expect: 23340,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expect, sumAdvanced(tt.input))
		})
	}
}
