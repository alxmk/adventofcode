package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUniverse3D(t *testing.T) {
	tests := []struct {
		name   string
		input  string
		cycles int
		expect int
	}{
		{
			name: "Ex1.0",
			input: `.#.
..#
###`,
			cycles: 0,
			expect: 5,
		},
		{
			name: "Ex1.1",
			input: `.#.
..#
###`,
			cycles: 1,
			expect: 11,
		},
		{
			name: "Ex1.2",
			input: `.#.
..#
###`,
			cycles: 2,
			expect: 21,
		},
		{
			name: "Ex1.3",
			input: `.#.
..#
###`,
			cycles: 3,
			expect: 38,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := parse3d(tt.input)
			for i := 0; i < tt.cycles; i++ {
				u = u.Next()
			}
			assert.Equal(t, tt.expect, u.Count())
		})
	}
}

func TestUniverse4D(t *testing.T) {
	tests := []struct {
		name   string
		input  string
		cycles int
		expect int
	}{
		{
			name: "Ex1.0",
			input: `.#.
..#
###`,
			cycles: 0,
			expect: 5,
		},
		{
			name: "Ex1.1",
			input: `.#.
..#
###`,
			cycles: 1,
			expect: 29,
		},
		{
			name: "Ex1.2",
			input: `.#.
..#
###`,
			cycles: 2,
			expect: 60,
		},
		{
			name: "Ex1.6",
			input: `.#.
..#
###`,
			cycles: 6,
			expect: 848,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := parse4d(tt.input)
			for i := 0; i < tt.cycles; i++ {
				u = u.Next()
			}
			assert.Equal(t, tt.expect, u.Count())
		})
	}
}
