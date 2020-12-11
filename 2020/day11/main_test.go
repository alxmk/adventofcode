package main

import (
	"io/ioutil"
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOccupancy(t *testing.T) {
	tests := []struct {
		name           string
		input          string
		expectAdjacent int
		expectVector   int
	}{
		{
			name: "Ex1",
			input: `L.LL.LL.LL
LLLLLLL.LL
L.L.L..L..
LLLL.LL.LL
L.LL.LL.LL
L.LLLLL.LL
..L.L.....
LLLLLLLLLL
L.LLLLLL.L
L.LLLLL.LL`,
			expectAdjacent: 37,
			expectVector:   26,
		},
		// 		{
		// 			name: "Ex2",
		// 			input: `L.LL
		// LLLL`,
		// 			expectAdjacent: 4,
		// 			expectVector:   6,
		// 		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expectAdjacent, parse(tt.input).SteadyState(adjacentRules))
			assert.Equal(t, tt.expectVector, parse(tt.input).SteadyState(vectorRules))
		})
	}
}

var i string

func init() {
	data, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatalln("Failed to read input:", err)
	}
	i = string(data)
}

func BenchmarkStuff(b *testing.B) {
	for n := 0; n < b.N; n++ {
		parse(i).SteadyState(vectorRules)
	}
}
