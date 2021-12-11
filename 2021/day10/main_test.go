package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPartOne(t *testing.T) {
	tests := []struct {
		name   string
		input  string
		expect rune
	}{
		{
			name:   "Example 1",
			input:  "{([(<{}[<>[]}>{[]{[(<()>",
			expect: '}',
		},
		{
			name:   "Example 2",
			input:  "[[<[([]))<([[{}[[()]]]",
			expect: ')',
		},
		{
			name:   "Example 3",
			input:  "[{[{({}]{}}([{[{{{}}([]",
			expect: ']',
		},
		{
			name:   "Example 4",
			input:  "[<(<(<(<{}))><([]([]()",
			expect: ')',
		},
		{
			name:   "Example 5",
			input:  "<{([([[(<>()){}]>(<<{{",
			expect: '>',
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual, _ := chunk(tt.input)
			assert.Equal(t, tt.expect, actual)
		})
	}
}

func TestChunk(t *testing.T) {
	tests := []struct {
		name             string
		input            string
		expectRune       rune
		expectIncomplete string
	}{
		{
			name:             "Example 1",
			input:            "[({(<(())[]>[[{[]{<()<>>",
			expectRune:       0,
			expectIncomplete: "}}]])})]",
		},
		{
			name:             "Example 2",
			input:            "[(()[<>])]({[<{<<[]>>(",
			expectRune:       0,
			expectIncomplete: ")}>]})",
		},
		{
			name:             "Example 3",
			input:            "(((({<>}<{<{<>}{[]{[]{}",
			expectRune:       0,
			expectIncomplete: "}}>}>))))",
		},
		{
			name:             "Example 4",
			input:            "{<[[]]>}<{[{[{[]{()[[[]",
			expectRune:       0,
			expectIncomplete: "]]}}]}]}>",
		},
		{
			name:             "Example 5",
			input:            "<{([{{}}[<[[[<>{}]]]>[]]",
			expectRune:       0,
			expectIncomplete: "])}>",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actualRune, actualIncomplete := chunk(tt.input)
			assert.Equal(t, tt.expectRune, actualRune)
			assert.Equal(t, tt.expectIncomplete, string(actualIncomplete))
		})
	}
}
