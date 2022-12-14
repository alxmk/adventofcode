package main

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParse(t *testing.T) {
	tests := []struct {
		name   string
		input  string
		expect *world
	}{
		{
			name: "Example",
			input: `498,4 -> 498,6 -> 496,6
503,4 -> 502,4 -> 502,9 -> 494,9`,
			expect: &world{
				ymax: 9,
				xmax: 503,
				xmin: 494,
				tiles: map[xy]tile{
					{498, 4}: rock,
					{498, 5}: rock,
					{498, 6}: rock,
					{497, 6}: rock,
					{496, 6}: rock,
					{503, 4}: rock,
					{502, 4}: rock,
					{502, 5}: rock,
					{502, 6}: rock,
					{502, 7}: rock,
					{502, 8}: rock,
					{502, 9}: rock,
					{501, 9}: rock,
					{500, 9}: rock,
					{499, 9}: rock,
					{498, 9}: rock,
					{497, 9}: rock,
					{496, 9}: rock,
					{495, 9}: rock,
					{494, 9}: rock,
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expect, parse(strings.Split(tt.input, "\n")))
		})
	}
}

func TestPartOne(t *testing.T) {
	tests := []struct {
		name   string
		input  string
		expect int
	}{
		{
			name: "Example",
			input: `498,4 -> 498,6 -> 496,6
503,4 -> 502,4 -> 502,9 -> 494,9`,
			expect: 24,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expect, partOne(parse(strings.Split(tt.input, "\n"))))
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
			input: `498,4 -> 498,6 -> 496,6
503,4 -> 502,4 -> 502,9 -> 494,9`,
			expect: 93,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expect, partTwo(parse(strings.Split(tt.input, "\n"))))
		})
	}
}
