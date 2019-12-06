package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsPasswordAnyAdjacent(t *testing.T) {
	tests := []struct {
		name             string
		num              int
		expectIsPassword bool
	}{
		{
			name:             "Ex1",
			num:              111111,
			expectIsPassword: true,
		},
		{
			name:             "Ex2",
			num:              223450,
			expectIsPassword: false,
		},
		{
			name:             "Ex3",
			num:              123789,
			expectIsPassword: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expectIsPassword, isPasswordAnyAdjacent(tt.num))
		})
	}
}

func TestIsPasswordExactlyDouble(t *testing.T) {
	tests := []struct {
		name             string
		num              int
		expectIsPassword bool
	}{
		{
			name:             "Ex1",
			num:              112233,
			expectIsPassword: true,
		},
		{
			name:             "Ex2",
			num:              123444,
			expectIsPassword: false,
		},
		{
			name:             "Ex3",
			num:              111122,
			expectIsPassword: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expectIsPassword, isPasswordExactlyDouble(tt.num))
		})
	}
}
