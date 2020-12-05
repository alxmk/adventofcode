package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseSeat(t *testing.T) {
	tests := []struct {
		name       string
		input      string
		expectSeat seat
		expectID   int
	}{
		{
			name:       "Ex1",
			input:      "BFFFBBFRRR",
			expectSeat: seat{row: 70, column: 7},
			expectID:   567,
		},
		{
			name:       "Ex2",
			input:      "FFFBBBFRRR",
			expectSeat: seat{row: 14, column: 7},
			expectID:   119,
		},
		{
			name:       "Ex3",
			input:      "BBFFBBFRLL",
			expectSeat: seat{row: 102, column: 4},
			expectID:   820,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actualSeat := parseSeat(tt.input)
			assert.Equal(t, tt.expectSeat, actualSeat)
			assert.Equal(t, tt.expectID, actualSeat.ID())
		})
	}
}
