package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIR_Overlaps(t *testing.T) {
	for _, tt := range []struct {
		i, j   ir
		expect bool
	}{
		{ir{10, 14}, ir{12, 18}, true},
		{ir{16, 20}, ir{12, 18}, true},
		{ir{10, 14}, ir{16, 20}, false},
	} {
		assert.Equal(t, tt.expect, tt.i.overlaps(tt.j))
		assert.Equal(t, tt.expect, tt.j.overlaps(tt.i))
	}
}

func TestIR_Reduce(t *testing.T) {
	for _, tt := range []struct {
		i, j   ir
		expect ir
	}{
		{ir{10, 14}, ir{12, 18}, ir{10, 18}},
		{ir{16, 20}, ir{12, 18}, ir{12, 20}},
	} {
		assert.Equal(t, tt.expect, tt.i.reduce(tt.j))
		assert.Equal(t, tt.expect, tt.j.reduce(tt.i))
	}
}

func TestConsolidate(t *testing.T) {
	for _, tt := range []struct {
		input  []ir
		expect []ir
	}{
		{
			[]ir{{3, 5}, {10, 14}, {16, 20}, {12, 18}},
			[]ir{{3, 5}, {10, 20}},
		},
	} {
		assert.Equal(t, tt.expect, consolidate(tt.input))
	}
}
