package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSafe(t *testing.T) {
	tests := []struct {
		name    string
		reports []report
		expect  int
	}{
		{
			name: "Ex1",
			reports: []report{
				{7, 6, 4, 2, 1},
				{1, 2, 7, 8, 9},
				{9, 7, 6, 2, 1},
				{1, 3, 2, 4, 5},
				{8, 6, 4, 4, 1},
				{1, 3, 6, 7, 9},
			},
			expect: 4,
		},
		{
			name: "Ex2",
			reports: []report{
				{43, 41, 43, 44, 45, 47, 49},
				{32, 34, 29, 26, 24},
			},
			expect: 2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var actual int
			for _, r := range tt.reports {
				if r.safe(true) {
					actual++
				}
			}
			assert.Equal(t, tt.expect, actual)
		})
	}
}
