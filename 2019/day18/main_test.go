package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMonkey_Go(t *testing.T) {
	tests := []struct {
		name   string
		input  string
		expect int64
	}{
		{
			name: "Ex1",
			input: `#########
#b.A.@.a#
#########`,
			expect: 8,
		},
		{
			name: "Ex2",
			input: `########################
#f.D.E.e.C.b.A.@.a.B.c.#
######################.#
#d.....................#
########################`,
			expect: 86,
		},
		{
			name: "Ex3",
			input: `########################
#...............b.C.D.f#
#.######################
#.....@.a.B.c.d.A.e.F.g#
########################`,
			expect: 132,
		},
		{
			name: "Ex5",
			input: `########################
#@..............ac.GI.b#
###d#e#f################
###A#B#C################
###g#h#i################
########################`,
			expect: 81,
		},
		{
			name: "Ex4",
			input: `#################
#i.G..c...e..H.p#
########.########
#j.A..b...f..D.o#
########@########
#k.E..a...g..B.n#
########.########
#l.F..d...h..C.m#
#################`,
			expect: 136,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expect, newMonkey(parseWorld(tt.input)).Go())
		})
	}
}
