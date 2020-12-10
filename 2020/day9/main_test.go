package main

import (
	"testing"

	"gopkg.in/go-playground/assert.v1"
)

func TestScan(t *testing.T) {
	tests := []struct {
		name   string
		input  string
		expect uint64
	}{
		{
			name: "Ex1",
			input: `35
20
15
25
47
40
62
55
65
95
102
117
150
182
127
219
299
277
309
576`,
			expect: 127,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expect, scan(parse(tt.input), 5))
		})
	}
}
