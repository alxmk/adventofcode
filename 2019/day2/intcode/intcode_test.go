package intcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestProgramme_Run(t *testing.T) {
	tests := []struct {
		name   string
		prog   Programme
		expect Programme
	}{
		{
			name: "Ex1",
			prog: map[int64]int64{
				0: 1,
				1: 0,
				2: 0,
				3: 0,
				4: 99,
			},
			expect: map[int64]int64{
				0: 2,
				1: 0,
				2: 0,
				3: 0,
				4: 99,
			},
		},
		{
			name: "Ex2",
			prog: map[int64]int64{
				0: 2,
				1: 3,
				2: 0,
				3: 3,
				4: 99,
			},
			expect: map[int64]int64{
				0: 2,
				1: 3,
				2: 0,
				3: 6,
				4: 99,
			},
		},
		{
			name: "Ex3",
			prog: map[int64]int64{
				0: 2,
				1: 4,
				2: 4,
				3: 5,
				4: 99,
				5: 0,
			},
			expect: map[int64]int64{
				0: 2,
				1: 4,
				2: 4,
				3: 5,
				4: 99,
				5: 9801,
			},
		},
		{
			name: "Ex4",
			prog: map[int64]int64{
				0: 1,
				1: 1,
				2: 1,
				3: 4,
				4: 99,
				5: 5,
				6: 6,
				7: 0,
				8: 99,
			},
			expect: map[int64]int64{
				0: 30,
				1: 1,
				2: 1,
				3: 4,
				4: 2,
				5: 5,
				6: 6,
				7: 0,
				8: 99,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			require.NoError(t, tt.prog.Run(nil, make(chan int64)))
			assert.Equal(t, tt.expect, tt.prog)
		})
	}
}

func TestProgramme_Copy(t *testing.T) {
	tests := []struct {
		name   string
		prog   Programme
		expect Programme
	}{
		{
			name: "Ex1",
			prog: map[int64]int64{
				0: 1,
				1: 0,
				2: 0,
				3: 0,
				4: 99,
			},
			expect: map[int64]int64{
				0: 1,
				1: 0,
				2: 0,
				3: 0,
				4: 99,
			},
		},
		{
			name: "Ex2",
			prog: map[int64]int64{
				0: 2,
				1: 3,
				2: 0,
				3: 3,
				4: 99,
			},
			expect: map[int64]int64{
				0: 2,
				1: 3,
				2: 0,
				3: 3,
				4: 99,
			},
		},
		{
			name: "Ex3",
			prog: map[int64]int64{
				0: 2,
				1: 4,
				2: 4,
				3: 5,
				4: 99,
				5: 0,
			},
			expect: map[int64]int64{
				0: 2,
				1: 4,
				2: 4,
				3: 5,
				4: 99,
				5: 0,
			},
		},
		{
			name: "Ex4",
			prog: map[int64]int64{
				0: 1,
				1: 1,
				2: 1,
				3: 4,
				4: 99,
				5: 5,
				6: 6,
				7: 0,
				8: 99,
			},
			expect: map[int64]int64{
				0: 1,
				1: 1,
				2: 1,
				3: 4,
				4: 99,
				5: 5,
				6: 6,
				7: 0,
				8: 99,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expect, tt.prog.Copy())
		})
	}
}

func TestParseOpcode(t *testing.T) {
	tests := []struct {
		name        string
		opcode      int64
		expectOp    int64
		expectModeA int64
		expectModeB int64
		expectModeC int64
	}{
		{
			name:        "Ex1",
			opcode:      1002,
			expectOp:    2,
			expectModeA: 0,
			expectModeB: 1,
			expectModeC: 0,
		},
		{
			name:        "Ex2",
			opcode:      1101,
			expectOp:    1,
			expectModeA: 1,
			expectModeB: 1,
			expectModeC: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actualOp, actualModeA, actualModeB, actualModeC := parseOpcode(tt.opcode)
			assert.Equal(t, tt.expectOp, actualOp)
			assert.Equal(t, tt.expectModeA, actualModeA)
			assert.Equal(t, tt.expectModeB, actualModeB)
			assert.Equal(t, tt.expectModeC, actualModeC)
		})
	}
}

func TestProgramme_RunIO(t *testing.T) {
	tests := []struct {
		name         string
		programme    string
		input        []int64
		expectOutput []int64
	}{
		{
			name:         "Ex1 input == 8",
			programme:    "3,9,8,9,10,9,4,9,99,-1,8",
			input:        []int64{8},
			expectOutput: []int64{1},
		},
		{
			name:         "Ex1 input != 8",
			programme:    "3,9,8,9,10,9,4,9,99,-1,8",
			input:        []int64{7},
			expectOutput: []int64{0},
		},
		{
			name:         "Ex2 input < 8",
			programme:    "3,9,7,9,10,9,4,9,99,-1,8",
			input:        []int64{7},
			expectOutput: []int64{1},
		},
		{
			name:         "Ex2 input >= 8",
			programme:    "3,9,7,9,10,9,4,9,99,-1,8",
			input:        []int64{9},
			expectOutput: []int64{0},
		},
		{
			name:         "Ex3 input == 8",
			programme:    "3,3,1108,-1,8,3,4,3,99",
			input:        []int64{8},
			expectOutput: []int64{1},
		},
		{
			name:         "Ex3 input != 8",
			programme:    "3,3,1108,-1,8,3,4,3,99",
			input:        []int64{7},
			expectOutput: []int64{0},
		},
		{
			name:         "Ex4 input < 8",
			programme:    "3,3,1107,-1,8,3,4,3,99",
			input:        []int64{7},
			expectOutput: []int64{1},
		},
		{
			name:         "Ex4 input >= 8",
			programme:    "3,3,1107,-1,8,3,4,3,99",
			input:        []int64{9},
			expectOutput: []int64{0},
		},
		{
			name:         "Ex5 self output",
			programme:    "109,1,204,-1,1001,100,1,100,1008,100,16,101,1006,101,0,99",
			input:        []int64{},
			expectOutput: []int64{109, 1, 204, -1, 1001, 100, 1, 100, 1008, 100, 16, 101, 1006, 101, 0, 99},
		},
		{
			name:         "Ex6 16 digit output",
			programme:    "1102,34915192,34915192,7,4,7,99,0",
			input:        []int64{},
			expectOutput: []int64{1219070632396864},
		},
		{
			name:         "Ex7 large output",
			programme:    "104,1125899906842624,99",
			input:        []int64{},
			expectOutput: []int64{1125899906842624},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p, err := Parse(tt.programme)
			require.NoError(t, err)

			in, out := make(chan int64), make(chan int64)
			go func() {
				for _, elem := range tt.input {
					in <- elem
				}
				close(in)
			}()

			go func() {
				require.NoError(t, p.Run(in, out))
			}()

			var actualOutput []int64
			for elem := range out {
				actualOutput = append(actualOutput, elem)
			}

			assert.Equal(t, tt.expectOutput, actualOutput)
		})
	}
}
