package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCircuitRun(t *testing.T) {
	tests := []struct {
		name          string
		input         string
		expectCircuit circuit
	}{
		{
			name: "Ex1",
			input: `123 -> x
456 -> y
x AND y -> d
x OR y -> e
x LSHIFT 2 -> f
y RSHIFT 2 -> g
NOT x -> h
NOT y -> i`,
			expectCircuit: circuit{
				"d": 72,
				"e": 507,
				"f": 492,
				"g": 114,
				"h": 65412,
				"i": 65079,
				"x": 123,
				"y": 456,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actualCircuit := make(circuit)
			actualCircuit.Run(parseCommands(tt.input))
			assert.Equal(t, tt.expectCircuit, actualCircuit)
		})
	}
}
