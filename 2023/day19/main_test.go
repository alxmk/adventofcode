package main

import (
	"fmt"
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
			name: "Ex1",
			input: `px{a<2006:qkq,m>2090:A,rfg}
pv{a>1716:R,A}
lnx{m>1548:A,A}
rfg{s<537:gd,x>2440:R,A}
qs{s>3448:A,lnx}
qkq{x<1416:A,crn}
crn{x>2662:A,R}
in{s<1351:px,qqz}
qqz{s>2770:qs,m<1801:hdj,R}
gd{a>3333:R,R}
hdj{m>838:A,pv}

{x=787,m=2655,a=1222,s=2876}
{x=1679,m=44,a=2067,s=496}
{x=2036,m=264,a=79,s=2244}
{x=2461,m=1339,a=466,s=291}
{x=2127,m=1623,a=2188,s=1013}`,
			expect: 19114,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w, p := parse([]byte(tt.input))
			assert.Equal(t, tt.expect, partOne(w, p))
		})
	}
}

func TestPartTwo(t *testing.T) {
	tests := []struct {
		name   string
		input  string
		expect int64
	}{
		{
			name: "Ex1",
			input: `px{a<2006:qkq,m>2090:A,rfg}
pv{a>1716:R,A}
lnx{m>1548:A,A}
rfg{s<537:gd,x>2440:R,A}
qs{s>3448:A,lnx}
qkq{x<1416:A,crn}
crn{x>2662:A,R}
in{s<1351:px,qqz}
qqz{s>2770:qs,m<1801:hdj,R}
gd{a>3333:R,R}
hdj{m>838:A,pv}

{x=787,m=2655,a=1222,s=2876}
{x=1679,m=44,a=2067,s=496}
{x=2036,m=264,a=79,s=2244}
{x=2461,m=1339,a=466,s=291}
{x=2127,m=1623,a=2188,s=1013}`,
			expect: 167409079868000,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w, _ := parse([]byte(tt.input))
			assert.Equal(t, tt.expect, partTwo(w))
		})
	}
}

func TestDFS(t *testing.T) {
	w, _ := parse([]byte(`px{a<2006:qkq,m>2090:A,rfg}
pv{a>1716:R,A}
lnx{m>1548:A,A}
rfg{s<537:gd,x>2440:R,A}
qs{s>3448:A,lnx}
qkq{x<1416:A,crn}
crn{x>2662:A,R}
in{s<1351:px,qqz}
qqz{s>2770:qs,m<1801:hdj,R}
gd{a>3333:R,R}
hdj{m>838:A,pv}

{x=787,m=2655,a=1222,s=2876}
{x=1679,m=44,a=2067,s=496}
{x=2036,m=264,a=79,s=2244}
{x=2461,m=1339,a=466,s=291}
{x=2127,m=1623,a=2188,s=1013}`))

	tests := []struct {
		name   string
		input  string
		expect []workflow
	}{
		{
			name: "Ex1",
			input: `px{a<2006:qkq,m>2090:A,rfg}
pv{a>1716:R,A}
lnx{m>1548:A,A}
rfg{s<537:gd,x>2440:R,A}
qs{s>3448:A,lnx}
qkq{x<1416:A,crn}
crn{x>2662:A,R}
in{s<1351:px,qqz}
qqz{s>2770:qs,m<1801:hdj,R}
gd{a>3333:R,R}
hdj{m>838:A,pv}

{x=787,m=2655,a=1222,s=2876}
{x=1679,m=44,a=2067,s=496}
{x=2036,m=264,a=79,s=2244}
{x=2461,m=1339,a=466,s=291}
{x=2127,m=1623,a=2188,s=1013}`,
			expect: []workflow{
				{
					// "in", "px", "qkq", "A",
					w["in"][0],
					w["px"][0],
					w["qkq"][0],
					w["A"][0],
				},
				{
					// "in", "px", "qkq", "crn", "A",
					w["in"][0],
					w["px"][0],
					w["qkq"][0].invert(),
					w["qkq"][1],
					w["crn"][0],
					w["A"][0],
				},
				{
					// "in", "px", "A",
					w["in"][0],
					w["px"][0].invert(),
					w["px"][1],
					w["A"][0],
				},
				{
					// "in", "px", "rfg", "A",
					w["in"][0],
					w["px"][0].invert(),
					w["px"][1].invert(),
					w["px"][2],
					w["rfg"][0].invert(),
					w["rfg"][1].invert(),
					w["rfg"][2],
					w["A"][0],
				},
				{
					// "in", "qqz", "qs", "A",
					w["in"][0].invert(),
					w["in"][1],
					w["qqz"][0],
					w["qs"][0],
					w["A"][0],
				},
				{
					// "in", "qqz", "qs", "lnx", "A",
					w["in"][0].invert(),
					w["in"][1],
					w["qqz"][0],
					w["qs"][0].invert(),
					w["qs"][1],
					w["lnx"][0],
					w["A"][0],
				},
				{
					// "in", "qqz", "qs", "lnx", "A",
					w["in"][0].invert(),
					w["in"][1],
					w["qqz"][0],
					w["qs"][0].invert(),
					w["qs"][1],
					w["lnx"][0].invert(),
					w["lnx"][1],
					w["A"][0],
				},
				{
					// "in", "qqz", "hdj", "A",
					w["in"][0].invert(),
					w["in"][1],
					w["qqz"][0].invert(),
					w["qqz"][1],
					w["hdj"][0],
					w["A"][0],
				},
				{
					// "in", "qqz", "hdj", "pv", "A",
					w["in"][0].invert(),
					w["in"][1],
					w["qqz"][0].invert(),
					w["qqz"][1],
					w["hdj"][0].invert(),
					w["hdj"][1],
					w["pv"][0].invert(),
					w["pv"][1],
					w["A"][0],
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w, _ := parse([]byte(tt.input))
			paths := dfs(w)
			assert.ElementsMatch(t, tt.expect, paths)
			fmt.Println(tt.expect)
			fmt.Println(paths)
		})
	}
}

func TestWorkflowRange(t *testing.T) {
	tests := []struct {
		name                 string
		w                    workflow
		expectmin, expectmax part
	}{
		{
			name: "Ex1",
			w: workflow{
				{
					category:   "s",
					condition:  '<',
					comparator: 1351,
					result:     "px",
				},
				{
					category:   "a",
					condition:  '>',
					comparator: 2005,
					result:     "qkq",
				},
				{
					category:   "m",
					condition:  '>',
					comparator: 2090,
					result:     "A",
				},
			},
			expectmin: part{x: 1, m: 2091, a: 2006, s: 1},
			expectmax: part{x: 4000, m: 4000, a: 4000, s: 1350},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actualmin, actualmax := tt.w.Range()
			assert.Equal(t, tt.expectmin, actualmin)
			assert.Equal(t, tt.expectmax, actualmax)
		})
	}
}
