package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseAnyone(t *testing.T) {
	tests := []struct {
		name   string
		input  string
		expect int
	}{
		{
			name: "Ex2",
			input: `abc

a
b
c

ab
ac

a
a
a
a

b
`,
			expect: 11,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expect, parseAnyone(tt.input).Length())
		})
	}
}

func TestParseEveryone(t *testing.T) {
	tests := []struct {
		name   string
		input  string
		expect int
	}{
		{
			name: "Ex2",
			input: `abc

a
b
c

ab
ac

a
a
a
a

b`,
			expect: 6,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expect, parseEveryone(tt.input).Length())
		})
	}
}
