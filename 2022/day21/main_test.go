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
			name:   "Example",
			input:  example,
			expect: 152,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			monkeys := parse(tt.input)
			assert.Equal(t, tt.expect, monkeys["root"].eval(monkeys, &suspiciousCache{data: make(map[string]*cacheEntry)}))
		})
	}
}

var example = `root: pppw + sjmn
dbpl: 5
cczh: sllz + lgvd
zczc: 2
ptdq: humn - dvpt
dvpt: 3
lfqf: 4
humn: 5
ljgn: 2
sjmn: drzm * dbpl
sllz: 4
pppw: cczh / lfqf
lgvd: ljgn * ptdq
drzm: hmdt - zczc
hmdt: 32`

func TestPartTwo(t *testing.T) {
	tests := []struct {
		name   string
		input  string
		expect int
	}{
		{
			name:   "Example",
			input:  example,
			expect: 301,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			monkeys := parse(tt.input)
			assert.Equal(t, tt.expect, partTwo(monkeys))
		})
	}
}

func TestDebugPartTwo(t *testing.T) {
	tests := []struct {
		name   string
		input  string
		expect int
	}{
		{
			name:   "Example",
			input:  debugExample,
			expect: 300,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			monkeys := parse(tt.input)
			// parse(tt.input)
			assert.Equal(t, tt.expect, monkeys["root"].eval(monkeys, &suspiciousCache{data: make(map[string]*cacheEntry)}))
		})
	}
}

var debugExample = `root: pppw + sjmn
dbpl: 5
cczh: sllz + lgvd
zczc: 2
ptdq: humn - dvpt
dvpt: 3
lfqf: 4
humn: 301
ljgn: 2
sjmn: drzm * dbpl
sllz: 4
pppw: cczh / lfqf
lgvd: ljgn * ptdq
drzm: hmdt - zczc
hmdt: 32`
