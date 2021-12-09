package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestPartOne(t *testing.T) {
	tests := []struct {
		name   string
		input  string
		expect int
	}{
		{
			name:   "Example",
			input:  sampleInput,
			expect: 26,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, outputLines, err := parse(tt.input)
			require.NoError(t, err)
			assert.Equal(t, tt.expect, partOne(outputLines))
		})
	}
}

func TestPartTwo(t *testing.T) {
	tests := []struct {
		name   string
		input  string
		expect int
	}{
		{
			name:   "Example 1",
			input:  shortInput,
			expect: 5353,
		},
		{
			name:   "Example 2",
			input:  sampleInput,
			expect: 61229,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			uspLines, outputLines, err := parse(tt.input)
			require.NoError(t, err)
			assert.Equal(t, tt.expect, partTwo(uspLines, outputLines))
		})
	}
}

var sampleInput = `be cfbegad cbdgef fgaecd cgeb fdcge agebfd fecdb fabcd edb | fdgacbe cefdb cefbgd gcbe
edbfga begcd cbg gc gcadebf fbgde acbgfd abcde gfcbed gfec | fcgedb cgb dgebacf gc
fgaebd cg bdaec gdafb agbcfd gdcbef bgcad gfac gcb cdgabef | cg cg fdcagb cbg
fbegcd cbd adcefb dageb afcb bc aefdc ecdab fgdeca fcdbega | efabcd cedba gadfec cb
aecbfdg fbg gf bafeg dbefa fcge gcbea fcaegb dgceab fcbdga | gecf egdcabf bgf bfgea
fgeab ca afcebg bdacfeg cfaedg gcfdb baec bfadeg bafgc acf | gebdcfa ecba ca fadegcb
dbcfg fgd bdegcaf fgec aegbdf ecdfab fbedc dacgb gdcebf gf | cefg dcbef fcge gbcadfe
bdfegc cbegaf gecbf dfcage bdacg ed bedf ced adcbefg gebcd | ed bcgafe cdgba cbgef
egadfb cdbfeg cegd fecab cgb gbdefca cg fgcdab egfdb bfceg | gbdfcae bgc cg cgb
gcafb gcf dcaebfg ecagb gf abcdeg gaef cafbge fdbac fegbdc | fgae cfgab fg bagce`

var shortInput = `acedgfb cdfbe gcdfa fbcad dab cefabd cdfgeb eafb cagedb ab | cdfeb fcadb cdfeb cdbaf`

func TestCommonCharacters(t *testing.T) {
	tests := []struct {
		name   string
		a, b   string
		expect int
	}{
		{
			name:   "Example 1",
			a:      "ab",
			b:      "abcdeg",
			expect: 2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expect, commonCharacters(tt.a, tt.b))
		})
	}
}
