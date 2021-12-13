package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestPartOne(t *testing.T) {
	tests := []struct {
		name   string
		input  string
		expect int
	}{
		{
			name: "Example 1",
			input: `start-A
start-b
A-c
A-b
b-d
A-end
b-end`,
			expect: 10,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			start, err := parse(tt.input)
			require.NoError(t, err)
			assert.Equal(t, tt.expect, len(traverse(start, "start", partOne)))
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
			name: "Example 1",
			input: `start-A
start-b
A-c
A-b
b-d
A-end
b-end`,
			expect: 36,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			start, err := parse(tt.input)
			require.NoError(t, err)
			assert.Equal(t, tt.expect, len(traverse(start, "start", partTwo)))
		})
	}
}

func TestPartTwoTraverseFunc(t *testing.T) {
	tests := []struct {
		name   string
		node   *node
		path   string
		expect bool
	}{
		{
			name:   "Big cave",
			node:   &node{name: "AA"},
			path:   "start,AA,bb",
			expect: false,
		},
		{
			name:   "Small cave unvisited",
			node:   &node{name: "cc"},
			path:   "start,AA,bb",
			expect: false,
		},
		{
			name:   "Small cave visited once",
			node:   &node{name: "bb"},
			path:   "start,AA,bb,CC",
			expect: false,
		},
		{
			name:   "Small cave visited twice",
			node:   &node{name: "bb"},
			path:   "start,AA,bb,CC,bb,AA",
			expect: true,
		},
		{
			name:   "Small cave, another visited twice",
			node:   &node{name: "bb"},
			path:   "start,AA,ee,CC,bb,AA,ee,CC",
			expect: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expect, partTwo(tt.node, tt.path))
		})
	}
}
