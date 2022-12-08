package main

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestVisible(t *testing.T) {
	tests := []struct {
		name   string
		input  string
		expect int
	}{
		{
			name: "Example",
			input: `30373
25512
65332
33549
35390`,
			expect: 21,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			trees := bytes.Split([]byte(tt.input), []byte{'\n'})
			assert.Equal(t, tt.expect, visible(trees))
		})
	}
}

func TestScenicScore(t *testing.T) {
	tests := []struct {
		name   string
		input  string
		x, y   int
		expect int
	}{
		{
			name: "Example 1",
			input: `30373
25512
65332
33549
35390`,
			x:      2,
			y:      1,
			expect: 4,
		},
		{
			name: "Example 2",
			input: `30373
25512
65332
33549
35390`,
			x:      2,
			y:      3,
			expect: 8,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			trees := bytes.Split([]byte(tt.input), []byte{'\n'})
			assert.Equal(t, tt.expect, scenicScore(tt.x, tt.y, trees))
		})
	}
}
