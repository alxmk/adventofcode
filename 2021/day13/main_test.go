package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestFold(t *testing.T) {
	tests := []struct {
		name   string
		input  string
		folds  int
		expect int
	}{
		{
			name: "Example 1",
			input: `6,10
0,14
9,10
0,3
10,4
4,11
6,0
6,12
4,1
0,13
10,12
3,4
3,0
8,4
1,10
2,14
8,10
9,0

fold along y=7
fold along x=5`,
			folds:  1,
			expect: 17,
		},
		{
			name: "Example 2",
			input: `6,10
0,14
9,10
0,3
10,4
4,11
6,0
6,12
4,1
0,13
10,12
3,4
3,0
8,4
1,10
2,14
8,10
9,0

fold along y=7
fold along x=5`,
			folds:  2,
			expect: 16,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			paper, folds, err := parse(tt.input)
			require.NoError(t, err)

			for i := 0; i < tt.folds; i++ {
				paper = paper.Fold(folds[i])
			}
			assert.Equal(t, tt.expect, len(paper.data))
		})
	}
}
