package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsNicePartTwo(t *testing.T) {
	tests := []struct {
		name       string
		line       string
		expectNice bool
	}{
		{
			name:       "Ex1",
			line:       "qjhvhtzxzqqjkmpb",
			expectNice: true,
		},
		{
			name:       "Ex2",
			line:       "xxyxx",
			expectNice: true,
		},
		{
			name:       "Ex3",
			line:       "uurcxstgmygtbstg",
			expectNice: false,
		},
		{
			name:       "Ex4",
			line:       "ieodomkazucvgmuy",
			expectNice: false,
		},
		{
			name:       "Ex5",
			line:       "sknufchjdvccccta",
			expectNice: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expectNice, isNicePartTwo(tt.line))
		})
	}
}
