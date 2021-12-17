package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestShoot(t *testing.T) {
	tests := []struct {
		name       string
		input      string
		v          coord
		expectHit  bool
		expectMaxY int
	}{
		{
			name:       "Example",
			input:      "target area: x=20..30, y=-10..-5",
			v:          coord{x: 7, y: 2},
			expectHit:  true,
			expectMaxY: 3,
		},
		{
			name:       "Example",
			input:      "target area: x=20..30, y=-10..-5",
			v:          coord{x: 6, y: 3},
			expectHit:  true,
			expectMaxY: 6,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			target, err := parse(tt.input)
			require.NoError(t, err)
			actualHit, actualMaxY := shoot(tt.v, *target)
			assert.Equal(t, tt.expectHit, actualHit)
			assert.Equal(t, tt.expectMaxY, actualMaxY)
		})
	}
}

func TestFindHits(t *testing.T) {
	tests := []struct {
		name       string
		input      string
		expectMaxY int
		expectHits int
	}{
		{
			name:       "Example",
			input:      "target area: x=20..30, y=-10..-5",
			expectMaxY: 45,
			expectHits: 112,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			target, err := parse(tt.input)
			require.NoError(t, err)
			actualMaxY, actualHits := findHits(target)
			assert.Equal(t, tt.expectMaxY, actualMaxY)
			assert.Equal(t, tt.expectHits, actualHits)
		})
	}
}
