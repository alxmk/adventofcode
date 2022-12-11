package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSimulate(t *testing.T) {
	tests := []struct {
		name   string
		input  string
		relief func(uint64) func(uint64) uint64
		rounds uint64
		expect uint64
	}{
		{
			name:   "Part one",
			input:  example,
			relief: func(p uint64) func(uint64) uint64 { return func(i uint64) uint64 { return i / 3 } },
			rounds: 20,
			expect: 10605,
		},
		{
			name:   "Part two round one",
			input:  example,
			relief: func(p uint64) func(uint64) uint64 { return func(i uint64) uint64 { return i % p } },
			rounds: 1,
			expect: 6 * 4,
		},
		{
			name:   "Part two round twenty",
			input:  example,
			relief: func(p uint64) func(uint64) uint64 { return func(i uint64) uint64 { return i % p } },
			rounds: 20,
			expect: 99 * 103,
		},
		{
			name:   "Part two round 1000",
			input:  example,
			relief: func(p uint64) func(uint64) uint64 { return func(i uint64) uint64 { return i % p } },
			rounds: 1000,
			expect: 5204 * 5192,
		},
		{
			name:   "Part two round 2000",
			input:  example,
			relief: func(p uint64) func(uint64) uint64 { return func(i uint64) uint64 { return i % p } },
			rounds: 2000,
			expect: 10391 * 10419,
		},
		{
			name:   "Part two round 3000",
			input:  example,
			relief: func(p uint64) func(uint64) uint64 { return func(i uint64) uint64 { return i % p } },
			rounds: 3000,
			expect: 15638 * 15593,
		},
		{
			name:   "Part two round 4000",
			input:  example,
			relief: func(p uint64) func(uint64) uint64 { return func(i uint64) uint64 { return i % p } },
			rounds: 4000,
			expect: 20858 * 20797,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m, p := parse(tt.input)
			assert.Equal(t, tt.expect, simulate(m, tt.relief(p), tt.rounds))
		})
	}
}

var example = `Monkey 0:
Starting items: 79, 98
Operation: new = old * 19
Test: divisible by 23
  If true: throw to monkey 2
  If false: throw to monkey 3

Monkey 1:
Starting items: 54, 65, 75, 74
Operation: new = old + 6
Test: divisible by 19
  If true: throw to monkey 2
  If false: throw to monkey 0

Monkey 2:
Starting items: 79, 60, 97
Operation: new = old * old
Test: divisible by 13
  If true: throw to monkey 1
  If false: throw to monkey 3

Monkey 3:
Starting items: 74
Operation: new = old + 3
Test: divisible by 17
  If true: throw to monkey 0
  If false: throw to monkey 1`
