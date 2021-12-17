package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestParsePacket(t *testing.T) {
	tests := []struct {
		name         string
		input        string
		expectPacket *packet
	}{
		{
			name:         "Example literal",
			input:        "D2FE28",
			expectPacket: &packet{version: 6, ptype: 4, value: 2021},
		},
		{
			name:  "Example operator two subpackets",
			input: "38006F45291200",
			expectPacket: &packet{
				version: 1,
				ptype:   6,
				subpackets: []packet{
					{
						version: 6,
						ptype:   4,
						value:   10,
					},
					{
						version: 2,
						ptype:   4,
						value:   20,
					},
				},
			},
		},
		{
			name:  "Example operator three subpackets",
			input: "EE00D40C823060",
			expectPacket: &packet{
				version: 7,
				ptype:   3,
				subpackets: []packet{
					{
						version: 2,
						ptype:   4,
						value:   1,
					},
					{
						version: 4,
						ptype:   4,
						value:   2,
					},
					{
						version: 1,
						ptype:   4,
						value:   3,
					},
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actualPacket, _, err := parsePacket(parseAsBinary(tt.input), 0, false)
			require.NoError(t, err)
			assert.Equal(t, tt.expectPacket, actualPacket)
		})
	}
}

func TestParseAsBinary(t *testing.T) {
	tests := []struct {
		name   string
		input  string
		expect string
	}{
		{
			name:   "Example literal",
			input:  "D2FE28",
			expect: "110100101111111000101000",
		},
		{
			name:   "Example operator",
			input:  "38006F45291200",
			expect: "00111000000000000110111101000101001010010001001000000000",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expect, parseAsBinary(tt.input))
		})
	}
}

func TestEvaluate(t *testing.T) {
	tests := []struct {
		name   string
		input  string
		expect uint64
	}{
		{
			name:   "Example 1",
			input:  "C200B40A82",
			expect: 3,
		},
		{
			name:   "Example 2",
			input:  "04005AC33890",
			expect: 54,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p, err := parse(tt.input)
			require.NoError(t, err)
			assert.Equal(t, tt.expect, p.Evaluate())
		})
	}
}
