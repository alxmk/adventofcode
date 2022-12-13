package main

import (
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParse(t *testing.T) {
	tests := []struct {
		name   string
		input  string
		expect packet
	}{
		{
			name:   "Example 1",
			input:  "[1,1,3,1,1]",
			expect: packet{values: []interface{}{1, 1, 3, 1, 1}},
		},
		{
			name:  "Example 2",
			input: "[[1],[2,3,4]]",
			expect: packet{
				values: []interface{}{
					packet{
						values: []interface{}{1},
					},
					packet{
						values: []interface{}{2, 3, 4},
					},
				},
			},
		},
		{
			name:  "Example 3",
			input: "[[10],[2,3,4]]",
			expect: packet{
				values: []interface{}{
					packet{
						values: []interface{}{10},
					},
					packet{
						values: []interface{}{2, 3, 4},
					},
				},
			},
		},
		{
			name:  "Example 4",
			input: "[[[]],[10,5,4]]",
			expect: packet{
				values: []interface{}{
					packet{
						values: []interface{}{
							packet{
								values: []interface{}{},
							},
						},
					},
					packet{
						values: []interface{}{
							10, 5, 4,
						},
					},
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, actual := parse([]byte(tt.input))
			assert.Equal(t, tt.expect, actual)
		})
	}
}

func TestValid(t *testing.T) {
	tests := []struct {
		name   string
		input  string
		expect bool
	}{
		{
			name: "Example 1",
			input: `[1,1,3,1,1]
[1,1,5,1,1]`,
			expect: true,
		},
		{
			name: "Example 2",
			input: `[[1],[2,3,4]]
[[1],4]`,
			expect: true,
		},
		{
			name: "Example 3",
			input: `[9]
[[8,7,6]]`,
			expect: false,
		},
		{
			name: "Example 4",
			input: `[[4,4],4,4]
[[4,4],4,4,4]`,
			expect: true,
		},
		{
			name: "Example 5",
			input: `[7,7,7,7]
[7,7,7]`,
			expect: false,
		},
		{
			name: "Example 6",
			input: `[]
[3]`,
			expect: true,
		},
		{
			name: "Example 7",
			input: `[[[]]]
[[]]`,
			expect: false,
		},
		{
			name: "Example 8",
			input: `[1,[2,[3,[4,[5,6,7]]]],8,9]
[1,[2,[3,[4,[5,6,0]]]],8,9]`,
			expect: false,
		},
		{
			name: "Example 9",
			input: `[1,[2,[3,[4,[5,6,7]]]],8,9]
[1,[2,[3,[4,[5,6,7]]]],7,9]`,
			expect: false,
		},
		{
			name: "Example 10",
			input: `[[[4,6,9,3,1],4,[3],[2,4,[6,10]],6],[]]
[[4,8],[[],6],[3,4]]`,
			expect: false,
		},
		{
			name: "Example 11",
			input: `[[],[],[4,[9],[[3,7]],4],[5,10,[4],[[9,4,7],[2,0,6,8,5],3,5,7],10],[]]
[[4,[2,10,[],[9,5,9]],[[2,8,8],[4,6,10],1,[9],[1,9]],[8,[5,3,10,10],[7,0],0],[]],[],[],[[10,[1,8,3,3],7,7,[4]],4],[[],1,3]]`,
			expect: true,
		},
		{
			name: "Example 12",
			input: `[[[],1,[[8,6,8],[5,9],10],[[10,9,1,10,10],9,[8,3,3,0,5],[2,6,1,3,5],[0,7,8,7]],[[8,6,9,1],4]],[10,[1]],[],[4,[[5,4,0],[1],5]]]
[[[]],[10,5,4]]`,
			expect: false,
		},
		{
			name: "Example 13",
			input: `[[6,2,[[],0,[1]]],[[[2],[6,2,8,5,0],[6,6,5,3],[8,10,8,5,1],[]],[[],[10,2,7],[7,3,4],3],[[3],1],8],[8,[7],[],7],[[[0,4,5,3,0],[0,10]],7]]
[[[[6,9]],0,[1,[6,6,4,6,5],8],[2,6,0,[3]],[]],[[2,1],[8,2,8,3,6],5,10],[1,8],[1,8,1],[[10,[5,5,4,8,2],1,[]]]]`,
			expect: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			log.Println(tt.name)
			assert.Equal(t, tt.expect, parsePair([]byte(tt.input)).Valid())
		})
	}
}

func TestScore(t *testing.T) {
	tests := []struct {
		name   string
		input  string
		expect int
	}{
		{
			name: "Example",
			input: `[1,1,3,1,1]
[1,1,5,1,1]

[[1],[2,3,4]]
[[1],4]

[9]
[[8,7,6]]

[[4,4],4,4]
[[4,4],4,4,4]

[7,7,7,7]
[7,7,7]

[]
[3]

[[[]]]
[[]]

[1,[2,[3,[4,[5,6,7]]]],8,9]
[1,[2,[3,[4,[5,6,0]]]],8,9]`,
			expect: 13,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expect, partOne(parsePairs([]byte(tt.input))))
		})
	}
}
