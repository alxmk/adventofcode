package main

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestProgramme(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		version int
		expect  uint64
	}{
		{
			name: "Ex1",
			input: `mask = XXXXXXXXXXXXXXXXXXXXXXXXXXXXX1XXXX0X
mem[8] = 11
mem[7] = 101
mem[8] = 0`,
			version: 1,
			expect:  165,
		},
		{
			name: "Ex2",
			input: `mask = 000000000000000000000000000000X1001X
mem[42] = 100
mask = 00000000000000000000000000000000X0XX
mem[26] = 1`,
			version: 2,
			expect:  208,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &programme{
				mem:          make(memory),
				instructions: strings.Split(tt.input, "\n"),
			}

			require.NoError(t, p.Execute(tt.version))

			assert.Equal(t, tt.expect, p.mem.Sum())
		})
	}
}
