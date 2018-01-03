package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSplitGrid(t *testing.T) {
	input := map[int]map[int]bool{
		0: map[int]bool{
			0: false,
			1: true,
			2: false,
		},
		1: map[int]bool{
			0: false,
			1: false,
			2: true,
		},
		2: map[int]bool{
			0: true,
			1: true,
			2: true,
		},
	}

	output := splitGrid(input)

	require.Len(t, output, 1)

	expected := []bool{false, true, false, false, false, true, true, true, true}

	require.Equal(t, expected, output[0])
}

func TestSplitGridBiggerInput(t *testing.T) {
	input := map[int]map[int]bool{
		0: map[int]bool{
			0: false,
			1: true,
			2: false,
			3: true,
		},
		1: map[int]bool{
			0: false,
			1: false,
			2: true,
			3: true,
		},
		2: map[int]bool{
			0: false,
			1: false,
			2: false,
			3: true,
		},
		3: map[int]bool{
			0: false,
			1: true,
			2: true,
			3: true,
		},
	}

	output := splitGrid(input)

	require.Len(t, output, 4)

	expected := []bool{false, true, false, false}
	require.Len(t, output[0], 4)
	require.Equal(t, expected, output[0])

	expected = []bool{false, true, true, true}
	require.Len(t, output[1], 4)
	require.Equal(t, expected, output[1])

	expected = []bool{false, false, false, true}
	require.Len(t, output[2], 4)
	require.Equal(t, expected, output[2])

	expected = []bool{false, true, true, true}
	require.Len(t, output[3], 4)
	require.Equal(t, expected, output[3])
}

func TestSplitGridEvenBiggerInput(t *testing.T) {
	input := map[int]map[int]bool{
		0: map[int]bool{
			0: false,
			1: true,
			2: false,
			3: true,
			4: false,
			5: true,
			6: true,
			7: false,
			8: false,
		},
		1: map[int]bool{
			0: false,
			1: false,
			2: true,
			3: true,
			4: true,
			5: true,
			6: true,
			7: false,
			8: true,
		},
		2: map[int]bool{
			0: false,
			1: false,
			2: false,
			3: true,
			4: false,
			5: false,
			6: false,
			7: true,
			8: true,
		},
		3: map[int]bool{
			0: false,
			1: true,
			2: true,
			3: true,
			4: false,
			5: true,
			6: false,
			7: false,
			8: true,
		},
		4: map[int]bool{
			0: false,
			1: true,
			2: true,
			3: true,
			4: false,
			5: true,
			6: false,
			7: false,
			8: false,
		},
		5: map[int]bool{
			0: false,
			1: false,
			2: true,
			3: true,
			4: false,
			5: true,
			6: true,
			7: true,
			8: true,
		},
		6: map[int]bool{
			0: false,
			1: true,
			2: true,
			3: true,
			4: false,
			5: false,
			6: false,
			7: true,
			8: true,
		},
		7: map[int]bool{
			0: false,
			1: true,
			2: true,
			3: false,
			4: false,
			5: true,
			6: true,
			7: false,
			8: true,
		},
		8: map[int]bool{
			0: true,
			1: true,
			2: false,
			3: false,
			4: false,
			5: false,
			6: true,
			7: false,
			8: true,
		},
	}

	output := splitGrid(input)

	require.Len(t, output, 9)

	expected := []bool{false, true, false, false, false, true, false, false, false}
	require.Len(t, output[0], 9)
	require.Equal(t, expected, output[0])

	expected = []bool{true, false, true, true, true, true, true, false, false}
	require.Len(t, output[1], 9)
	require.Equal(t, expected, output[1])

	expected = []bool{true, false, false, true, false, true, false, true, true}
	require.Len(t, output[2], 9)
	require.Equal(t, expected, output[2])

	expected = []bool{false, true, true, false, true, true, false, false, true}
	require.Len(t, output[3], 9)
	require.Equal(t, expected, output[3])
}

func TestSplitGridRealInput(t *testing.T) {
	input := map[int]map[int]bool{
		0: map[int]bool{
			0: true,
			1: true,
			2: true,
			3: true,
			4: true,
			5: true,
		},
		1: map[int]bool{
			0: true,
			1: true,
			2: false,
			3: true,
			4: true,
			5: false,
		},
		2: map[int]bool{
			0: true,
			1: false,
			2: true,
			3: true,
			4: false,
			5: true,
		},
		3: map[int]bool{
			0: false,
			1: false,
			2: false,
			3: true,
			4: true,
			5: true,
		},
		4: map[int]bool{
			0: false,
			1: false,
			2: false,
			3: true,
			4: true,
			5: false,
		},
		5: map[int]bool{
			0: false,
			1: true,
			2: false,
			3: true,
			4: false,
			5: true,
		},
	}

	output := splitGrid(input)

	require.Len(t, output, 9)

	expected := []bool{true, true, true, true}
	require.Len(t, output[0], 4)
	require.Equal(t, expected, output[0])

	expected = []bool{true, true, false, true}
	require.Len(t, output[1], 4)
	require.Equal(t, expected, output[1])

	expected = []bool{true, true, true, false}
	require.Len(t, output[2], 4)
	require.Equal(t, expected, output[2])

	expected = []bool{true, false, false, false}
	require.Len(t, output[3], 4)
	require.Equal(t, expected, output[3])

	expected = []bool{true, true, false, true}
	require.Len(t, output[3], 4)
	require.Equal(t, expected, output[4])

	expected = []bool{false, true, true, true}
	require.Len(t, output[3], 4)
	require.Equal(t, expected, output[5])

	expected = []bool{false, false, false, true}
	require.Len(t, output[3], 4)
	require.Equal(t, expected, output[6])

	expected = []bool{false, true, false, true}
	require.Len(t, output[3], 4)
	require.Equal(t, expected, output[7])

	expected = []bool{true, false, false, true}
	require.Len(t, output[3], 4)
	require.Equal(t, expected, output[8])
}

func TestAssembleGrid(t *testing.T) {
	gridLets := [][]bool{
		[]bool{false, true, false, false},
		[]bool{false, true, true, true},
		[]bool{false, false, false, true},
		[]bool{false, true, true, true},
	}

	grid := assembleGrid(gridLets)

	expected := map[int]map[int]bool{
		0: map[int]bool{
			0: false,
			1: true,
			2: false,
			3: true,
		},
		1: map[int]bool{
			0: false,
			1: false,
			2: true,
			3: true,
		},
		2: map[int]bool{
			0: false,
			1: false,
			2: false,
			3: true,
		},
		3: map[int]bool{
			0: false,
			1: true,
			2: true,
			3: true,
		},
	}

	require.Len(t, grid, 4)

	require.Equal(t, expected, grid)
}

func TestAssembleGridLarger(t *testing.T) {
	input := [][]bool{
		[]bool{
			false,
			true,
			true,
			true,
			true,
			true,
			false,
			true,
			true,
			true,
			false,
			true,
			true,
			true,
			true,
			true,
		},
	}

	grid := assembleGrid(input)

	require.Len(t, grid, 4)
}

func TestAssembleGridEvenLarger(t *testing.T) {
	input := [][]bool{
		[]bool{
			true,
			true,
			true,
			true,
			true,
			false,
			true,
			false,
			true,
		},
		[]bool{
			true,
			true,
			true,
			true,
			true,
			false,
			true,
			false,
			true,
		},
		[]bool{
			false,
			false,
			false,
			false,
			false,
			false,
			false,
			true,
			false,
		},
		[]bool{
			true,
			true,
			true,
			true,
			true,
			false,
			true,
			false,
			true,
		},
	}

	expected := map[int]map[int]bool{
		0: map[int]bool{
			0: true,
			1: true,
			2: true,
			3: true,
			4: true,
			5: true,
		},
		1: map[int]bool{
			0: true,
			1: true,
			2: false,
			3: true,
			4: true,
			5: false,
		},
		2: map[int]bool{
			0: true,
			1: false,
			2: true,
			3: true,
			4: false,
			5: true,
		},
		3: map[int]bool{
			0: false,
			1: false,
			2: false,
			3: true,
			4: true,
			5: true,
		},
		4: map[int]bool{
			0: false,
			1: false,
			2: false,
			3: true,
			4: true,
			5: false,
		},
		5: map[int]bool{
			0: false,
			1: true,
			2: false,
			3: true,
			4: false,
			5: true,
		},
	}

	actual := assembleGrid(input)

	require.Equal(t, expected, actual)
}
