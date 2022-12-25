package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParse(t *testing.T) {
	tests := []struct {
		snafu  string
		expect int
	}{
		{
			expect: 1,
			snafu:  "1",
		},
		{
			expect: 2,
			snafu:  "2",
		},
		{
			expect: 3,
			snafu:  "1=",
		},
		{
			expect: 4,
			snafu:  "1-",
		},
		{
			expect: 5,
			snafu:  "10",
		},
		{
			expect: 6,
			snafu:  "11",
		},
		{
			expect: 7,
			snafu:  "12",
		},
		{
			expect: 8,
			snafu:  "2=",
		},
		{
			expect: 9,
			snafu:  "2-",
		},
		{
			expect: 10,
			snafu:  "20",
		},
		{
			expect: 15,
			snafu:  "1=0",
		},
		{
			expect: 20,
			snafu:  "1-0",
		},
		{
			expect: 2022,
			snafu:  "1=11-2",
		},
		{
			expect: 12345,
			snafu:  "1-0---0",
		},
		{
			expect: 314159265,
			snafu:  "1121-1110-1=0",
		},
		{
			snafu:  "1=-0-2",
			expect: 1747,
		},
		{
			snafu:  "12111",
			expect: 906,
		},
		{
			snafu:  "2=0=",
			expect: 198,
		},
		{
			snafu:  "21",
			expect: 11,
		},
		{
			snafu:  "2=01",
			expect: 201,
		},
		{
			snafu:  "111",
			expect: 31,
		},
		{
			snafu:  "20012",
			expect: 1257,
		},
		{
			snafu:  "112",
			expect: 32,
		},
		{
			snafu:  "1=-1=",
			expect: 353,
		},
		{
			snafu:  "1-12",
			expect: 107,
		},
		{
			snafu:  "12",
			expect: 7,
		},
		{
			snafu:  "1=",
			expect: 3,
		},
		{
			snafu:  "122",
			expect: 37,
		},
	}

	for _, tt := range tests {
		t.Run(tt.snafu, func(t *testing.T) {
			assert.Equal(t, tt.expect, parse([]byte(tt.snafu)))
		})
	}
}

func TestConvert(t *testing.T) {
	tests := []struct {
		snafu  string
		expect int
	}{
		{
			expect: 1,
			snafu:  "1",
		},
		{
			expect: 2,
			snafu:  "2",
		},
		{
			expect: 3,
			snafu:  "1=",
		},
		{
			expect: 4,
			snafu:  "1-",
		},
		{
			expect: 5,
			snafu:  "10",
		},
		{
			expect: 6,
			snafu:  "11",
		},
		{
			expect: 7,
			snafu:  "12",
		},
		{
			expect: 8,
			snafu:  "2=",
		},
		{
			expect: 9,
			snafu:  "2-",
		},
		{
			expect: 10,
			snafu:  "20",
		},
		{
			expect: 15,
			snafu:  "1=0",
		},
		{
			expect: 20,
			snafu:  "1-0",
		},
		{
			expect: 2022,
			snafu:  "1=11-2",
		},
		{
			expect: 12345,
			snafu:  "1-0---0",
		},
		{
			expect: 314159265,
			snafu:  "1121-1110-1=0",
		},
		{
			snafu:  "1=-0-2",
			expect: 1747,
		},
		{
			snafu:  "12111",
			expect: 906,
		},
		{
			snafu:  "2=0=",
			expect: 198,
		},
		{
			snafu:  "21",
			expect: 11,
		},
		{
			snafu:  "2=01",
			expect: 201,
		},
		{
			snafu:  "111",
			expect: 31,
		},
		{
			snafu:  "20012",
			expect: 1257,
		},
		{
			snafu:  "112",
			expect: 32,
		},
		{
			snafu:  "1=-1=",
			expect: 353,
		},
		{
			snafu:  "1-12",
			expect: 107,
		},
		{
			snafu:  "12",
			expect: 7,
		},
		{
			snafu:  "1=",
			expect: 3,
		},
		{
			snafu:  "122",
			expect: 37,
		},
	}

	for _, tt := range tests {
		t.Run(tt.snafu, func(t *testing.T) {
			assert.Equal(t, tt.snafu, convert(tt.expect))
		})
	}
}
