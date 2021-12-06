package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGame(t *testing.T) {
	tests := []struct {
		name   string
		input  string
		turns  int
		expect int
	}{
		{
			name:   "Ex1",
			input:  "0,3,6",
			turns:  2020,
			expect: 436,
		},
		{
			name:   "Ex2",
			input:  "0,3,6",
			turns:  30000000,
			expect: 175594,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expect, newGame(tt.input).RunToTurn(tt.turns))
		})
	}
}
