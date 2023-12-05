package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPartOne(t *testing.T) {
	tests := []struct {
		name   string
		input  string
		expect int
	}{
		{
			name: "Example",
			input: `seeds: 79 14 55 13

seed-to-soil map:
50 98 2
52 50 48

soil-to-fertilizer map:
0 15 37
37 52 2
39 0 15

fertilizer-to-water map:
49 53 8
0 11 42
42 0 7
57 7 4

water-to-light map:
88 18 7
18 25 70

light-to-temperature map:
45 77 23
81 45 19
68 64 13

temperature-to-humidity map:
0 69 1
1 0 69

humidity-to-location map:
60 56 37
56 93 4`,
			expect: 35,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expect, partOne(parseAlmanac(tt.input)))
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
			name: "Example",
			input: `seeds: 79 14 55 13

seed-to-soil map:
50 98 2
52 50 48

soil-to-fertilizer map:
0 15 37
37 52 2
39 0 15

fertilizer-to-water map:
49 53 8
0 11 42
42 0 7
57 7 4

water-to-light map:
88 18 7
18 25 70

light-to-temperature map:
45 77 23
81 45 19
68 64 13

temperature-to-humidity map:
0 69 1
1 0 69

humidity-to-location map:
60 56 37
56 93 4`,
			expect: 46,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expect, partTwo(parseAlmanac(tt.input)))
		})
	}
}

func TestMappingCorresponding(t *testing.T) {
	tests := []struct {
		name   string
		m      mapping
		input  int
		expect int
	}{
		{
			name: "Seed to soil",
			m: mapping{
				src: 50,
				dst: 52,
				rng: 48,
			},
			input:  79,
			expect: 81,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual, _ := tt.m.Corresponding(tt.input)
			assert.Equal(t, tt.expect, actual)
		})
	}
}

func TestMappingReverse(t *testing.T) {
	tests := []struct {
		name   string
		m      mapping
		input  int
		expect int
	}{
		{
			name: "Seed to soil",
			m: mapping{
				src: 50,
				dst: 52,
				rng: 48,
			},
			input:  81,
			expect: 79,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual, _ := tt.m.Reverse(tt.input)
			assert.Equal(t, tt.expect, actual)
		})
	}
}

func TestAlmanacReverse(t *testing.T) {
	tests := []struct {
		name        string
		input       string
		s           int
		thing       string
		expect      int
		expectThing string
	}{
		{
			name: "Seed to soil",
			input: `seeds: 79 14 55 13

seed-to-soil map:
50 98 2
52 50 48

soil-to-fertilizer map:
0 15 37
37 52 2
39 0 15

fertilizer-to-water map:
49 53 8
0 11 42
42 0 7
57 7 4

water-to-light map:
88 18 7
18 25 70

light-to-temperature map:
45 77 23
81 45 19
68 64 13

temperature-to-humidity map:
0 69 1
1 0 69

humidity-to-location map:
60 56 37
56 93 4`,
			s:           81,
			thing:       "soil",
			expect:      79,
			expectThing: "seed",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := parseAlmanac(tt.input)
			actual, actualThing := a.Reverse(tt.s, tt.thing)
			assert.Equal(t, tt.expect, actual)
			assert.Equal(t, tt.expectThing, actualThing)
		})
	}
}
