package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidation(t *testing.T) {
	tests := []struct {
		name   string
		input  string
		expect int
	}{
		{
			name: "Ex1",
			input: `class: 1-3 or 5-7
row: 6-11 or 33-44
seat: 13-40 or 45-50

your ticket:
7,1,14

nearby tickets:
7,3,47
40,4,50
55,2,20
38,6,12`,
			expect: 71,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rules, _, nearby := parse(tt.input)
			var er int
			for _, ticket := range nearby {
				er += ticket.IsPossiblyValid(rules)
			}
			assert.Equal(t, tt.expect, er)
		})
	}
}
