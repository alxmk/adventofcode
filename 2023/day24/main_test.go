package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPartOne(t *testing.T) {
	tests := []struct {
		name   string
		input  string
		expect int
	}{
		{
			name: "Ex1",
			input: `19, 13, 30 @ -2,  1, -2
18, 19, 22 @ -1, -1, -2
20, 25, 34 @ -2, -2, -4
12, 31, 28 @ -1, -2, -1
20, 19, 15 @  1, -5, -3`,
			expect: 2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expect, partOne(parseHailstones(tt.input), [2]float64{7, 7}, [2]float64{27, 27}))
		})
	}
}

func TestVectorIntersects(t *testing.T) {
	tests := []struct {
		name     string
		v, w     hailstone
		expect   [3]float64
		expectok bool
	}{
		{
			name:     "Ex1",
			v:        hailstone{p: vector{19, 13, 30}, v: vector{-2, 1, -2}},
			w:        hailstone{p: vector{18, 19, 22}, v: vector{-1, -1, -2}},
			expect:   [3]float64{14.333333333333332, 15.333333333333334, 0},
			expectok: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual, ok := tt.v.intersects2d(tt.w)
			assert.Equal(t, tt.expect, actual)
			assert.Equal(t, tt.expectok, ok)
		})
	}
}

// Real example
// x1 = 19
// y1 = 13
// x2 = -2 + 19 = 17
// y2 = 1 + 13 = 14
//
// x3 = 18
// y3 = 19
// x4 = -1 + 18 = 17
// y4 = -1 + 19 = 18
//
// t := ((x1-x3)*(y3-y4) - (y1-y3)*(x3-x4)) /
// ((x1-x2)*(y3-y4) - (y1-y2)*(x3-x4))
// u := ((x1-x3)*(y1-y2) - (y1-y3)*(x1-x2)) /
// ((x1-x2)*(y3-y4) - (y1-y2)*(x3-x4))
//
// t = (19-18) * (19-18) - (13 - 19) * (18-17) / (19-17) * (19-18) - (13-14)*(18-17)
// t = 1 * 1 - (-6) * 1 / 2 * 1 - (-1) * 1
// t = 7 / 3
//
// intersection[0] = x1 + t*(x2-x1)
// = 19 + (7/3) * (17 - 19)
// = 19 - 14/3
// = 14.333
// intersection[1] = y1 + t*(y2-y1)
