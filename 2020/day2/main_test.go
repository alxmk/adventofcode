package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPassword_ValidTwo(t *testing.T) {
	tests := []struct {
		name        string
		pass        password
		expectValid bool
	}{
		{
			name: "Ex1",
			pass: password{
				A:        1,
				B:        3,
				Char:     byte('a'),
				Password: "abcde",
			},
			expectValid: true,
		},
		{
			name: "Ex2",
			pass: password{
				A:        1,
				B:        3,
				Char:     byte('b'),
				Password: "cdefg",
			},
			expectValid: false,
		},
		{
			name: "Ex3",
			pass: password{
				A:        2,
				B:        9,
				Char:     byte('c'),
				Password: "ccccccccc",
			},
			expectValid: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.pass.ValidTwo(), tt.expectValid)
		})
	}
}
