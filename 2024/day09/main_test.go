package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPartOne(t *testing.T) {
	tests := []struct {
		name   string
		input  string
		expect int64
	}{
		{
			name:   "Ex1",
			input:  "2333133121414131402",
			expect: 1928,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expect, partOne(tt.input))
		})
	}
}

func TestPartTwo(t *testing.T) {
	tests := []struct {
		name   string
		input  string
		expect int64
	}{
		{
			name:   "Ex1",
			input:  "2333133121414131402",
			expect: 2858,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expect, partTwo(tt.input))
		})
	}
}
