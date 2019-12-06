package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestUniverse_TraverseOrbits(t *testing.T) {
	input := `COM)B
B)C
C)D
D)E
E)F
B)G
G)H
D)I
E)J
J)K
K)L`

	tests := []struct {
		name         string
		input        string
		origin       string
		expectOrbits int
	}{
		{
			name:         "Ex1",
			input:        input,
			origin:       "D",
			expectOrbits: 3,
		},
		{
			name:         "Ex2",
			input:        input,
			origin:       "L",
			expectOrbits: 7,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u, err := parseObjects(tt.input)
			require.NoError(t, err)

			assert.Equal(t, tt.expectOrbits, u.TraverseOrbits(tt.origin))
		})
	}
}

func TestUniverse_TraverseAllOrbits(t *testing.T) {
	input := `COM)B
B)C
C)D
D)E
E)F
B)G
G)H
D)I
E)J
J)K
K)L`

	tests := []struct {
		name         string
		input        string
		expectOrbits int
	}{
		{
			name:         "Ex1",
			input:        input,
			expectOrbits: 42,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u, err := parseObjects(tt.input)
			require.NoError(t, err)

			assert.Equal(t, tt.expectOrbits, u.TraverseAllOrbits())
		})
	}
}

func TestUniverse_HopsBetween(t *testing.T) {
	input := `COM)B
B)C
C)D
D)E
E)F
B)G
G)H
D)I
E)J
J)K
K)L
K)YOU
I)SAN`

	tests := []struct {
		name       string
		input      string
		from       string
		to         string
		expectHops int
	}{
		{
			name:       "Ex1",
			input:      input,
			from:       "K",
			to:         "I",
			expectHops: 4,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u, err := parseObjects(tt.input)
			require.NoError(t, err)

			assert.Equal(t, tt.expectHops, u.HopsBetween(tt.from, tt.to))
		})
	}
}
