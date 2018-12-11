package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_calculatePower(t *testing.T) {
	tests := []struct {
		x, y, serial, expect int
	}{
		{
			3, 5, 8, 4,
		},
		{
			122, 79, 57, -5,
		},
		{
			217, 196, 39, 0,
		},
		{
			101, 153, 71, 4,
		},
	}

	for _, tt := range tests {
		require.Equal(t, tt.expect, calculatePower(tt.x, tt.y, tt.serial))
	}
}
