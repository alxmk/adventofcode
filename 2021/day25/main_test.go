package main

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIterate(t *testing.T) {
	tests := []struct {
		input   string
		expect  string
		changed bool
	}{
		{
			input:   "...>>>>>...",
			expect:  "...>>>>.>..",
			changed: true,
		},
		{
			input:   "...>>>>.>..",
			expect:  "...>>>.>.>.",
			changed: true,
		},
		{
			input: `..........
.>v....v..
.......>..
..........`,
			expect: `..........
.>........
..v....v>.
..........`,
			changed: true,
		},
		{
			input: `v...>>.vv>
.vv>>.vv..
>>.>v>...v
>>v>>.>.v.
v>v.vv.v..
>.>>..v...
.vv..>.>v.
v.v..>>v.v
....v..v.>`,
			expect: `....>.>v.>
v.v>.>v.v.
>v>>..>v..
>>v>v>.>.v
.>v.v...v.
v>>.>vvv..
..v...>>..
vv...>>vv.
>.v.v..v.v`,
			changed: true,
		},
		{
			input: `....>.>v.>
v.v>.>v.v.
>v>>..>v..
>>v>v>.>.v
.>v.v...v.
v>>.>vvv..
..v...>>..
vv...>>vv.
>.v.v..v.v`,
			expect: `>.v.v>>..v
v.v.>>vv..
>v>.>.>.v.
>>v>v.>v>.
.>..v....v
.>v>>.v.v.
v....v>v>.
.vv..>>v..
v>.....vv.`,
			changed: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			actual, c := iterate(bytes.Split([]byte(tt.input), []byte{'\n'}))
			assert.Equal(t, tt.expect, string(bytes.Join(actual, []byte{'\n'})))
			assert.Equal(t, tt.changed, c)
		})
	}
}

func TestStable(t *testing.T) {
	tests := []struct {
		input  string
		expect int
	}{
		{
			input: `v...>>.vv>
.vv>>.vv..
>>.>v>...v
>>v>>.>.v.
v>v.vv.v..
>.>>..v...
.vv..>.>v.
v.v..>>v.v
....v..v.>`,
			expect: 58,
		},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			assert.Equal(t, tt.expect, stable(bytes.Split([]byte(tt.input), []byte{'\n'})))
		})
	}
}
