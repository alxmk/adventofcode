package main

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestDistance(t *testing.T) {
	tests := []struct {
		from, to *point
		expect   int
	}{
		{
			from:   &point{x: 1, y: -1, z: 0, t: -1},
			to:     &point{x: 0, y: 0, z: -1, t: -1},
			expect: 3,
		},
	}

	for _, tt := range tests {
		require.Equal(t, tt.expect, tt.from.Distance(tt.to))
	}
}

func TestE2E(t *testing.T) {
	tests := []struct {
		input  string
		expect int
	}{
		{
			input: `0,0,0,0
3,0,0,0
0,3,0,0
0,0,3,0
0,0,0,3
0,0,0,6
9,0,0,0
12,0,0,0`,
			expect: 2,
		},
		{
			input: `-1,2,2,0
0,0,2,-2
0,0,0,-2
-1,2,0,0
-2,-2,-2,2
3,0,2,-1
-1,3,2,2
-1,0,-1,0
0,2,1,-2
3,0,0,0`,
			expect: 4,
		},
		{
			input: `1,-1,0,1
2,0,-1,0
3,2,-1,0
0,0,3,1
0,0,-1,-1
2,3,-2,0
-2,2,0,0
2,-2,0,-1
1,-1,0,-1
3,2,0,2`,
			expect: 3,
		},
		{
			input: `1,-1,-1,-2
-2,-2,0,1
0,2,1,3
-2,3,-2,1
0,2,3,-2
-1,-1,1,-2
0,-2,-1,0
-2,2,3,-1
1,2,2,0
-1,-2,0,-2`,
			expect: 8,
		},
	}

	for _, tt := range tests {
		ps := parsePoints(strings.Split(tt.input, "\n"))
		cons := ps.Constellations(3)

		require.Len(t, cons, tt.expect, cons.String())
	}
}
