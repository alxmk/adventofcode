package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestModulesTick(t *testing.T) {
	tests := []struct {
		name   string
		input  string
		expect int
	}{
		{
			name: "Ex1",
			input: `broadcaster -> a, b, c
%a -> b
%b -> c
%c -> inv
&inv -> a`,
			expect: 32,
		},
		{
			name: "Ex2",
			input: `broadcaster -> a
%a -> inv, con
&inv -> b
%b -> con
&con -> output`,
			expect: 16,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			counts := make(map[bool]int)
			counts = parse([]byte(tt.input)).tick(0, counts, make(map[string]int))
			assert.Equal(t, tt.expect, counts[true]*counts[false])
		})
	}
}

func TestModulesPartOne(t *testing.T) {
	tests := []struct {
		name   string
		input  string
		expect int
	}{
		{
			name: "Ex1",
			input: `broadcaster -> a, b, c
%a -> b
%b -> c
%c -> inv
&inv -> a`,
			expect: 32000000,
		},
		{
			name: "Ex2",
			input: `broadcaster -> a
%a -> inv, con
&inv -> b
%b -> con
&con -> output`,
			expect: 11687500,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expect, parse([]byte(tt.input)).partOne())
		})
	}
}
