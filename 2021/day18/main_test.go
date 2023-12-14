package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPartTwo(t *testing.T) {
	tests := []struct {
		name   string
		input  string
		expect int
	}{
		{
			name: "Ex1",
			input: `[[[0,[5,8]],[[1,7],[9,6]]],[[4,[1,2]],[[1,4],2]]]
[[[5,[2,8]],4],[5,[[9,9],0]]]
[6,[[[6,2],[5,6]],[[7,6],[4,7]]]]
[[[6,[0,7]],[0,9]],[4,[9,[9,0]]]]
[[[7,[6,4]],[3,[1,3]]],[[[5,5],1],9]]
[[6,[[7,3],[3,2]]],[[[3,8],[5,7]],4]]
[[[[5,4],[7,7]],8],[[8,3],8]]
[[9,3],[[9,9],[6,[4,9]]]]
[[2,[[7,7],7]],[[5,8],[[9,3],[0,2]]]]
[[[[5,2],5],[8,[3,7]]],[[5,[7,5]],[4,4]]]`,
			expect: 3993,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expect, partTwo(parseNumbers([]byte(tt.input))))
		})
	}
}

func TestNumberMagnitude(t *testing.T) {
	tests := []struct {
		name   string
		number string
		expect int
	}{
		{
			name:   "Ex1",
			number: "[[1,2],[[3,4],5]]",
			expect: 143,
		},
		{
			name:   "Ex2",
			number: "[[[[0,7],4],[[7,8],[6,0]]],[8,1]]",
			expect: 1384,
		},
		{
			name:   "Ex3",
			number: "[[[[1,1],[2,2]],[3,3]],[4,4]]",
			expect: 445,
		},
		{
			name:   "Ex4",
			number: "[[[[3,0],[5,3]],[4,4]],[5,5]]",
			expect: 791,
		},
		{
			name:   "Ex5",
			number: "[[[[5,0],[7,4]],[5,5]],[6,6]]",
			expect: 1137,
		},
		{
			name:   "Ex6",
			number: "[[[[8,7],[7,7]],[[8,6],[7,7]]],[[[0,7],[6,6]],[8,7]]]",
			expect: 3488,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expect, number(tt.number).Magnitude())
		})
	}
}

func TestNumberExplode(t *testing.T) {
	tests := []struct {
		name   string
		number string
		index  []int
		expect string
	}{
		{
			name:   "Ex1",
			number: "[[[[[9,8],1],2],3],4]",
			index:  []int{5, 8},
			expect: "[[[[0,9],2],3],4]",
		},
		{
			name:   "Ex2",
			number: "[7,[6,[5,[4,[3,2]]]]]",
			index:  []int{13, 16},
			expect: "[7,[6,[5,[7,0]]]]",
		},
		{
			name:   "Ex3",
			number: "[[6,[5,[4,[3,2]]]],1]",
			index:  []int{11, 14},
			expect: "[[6,[5,[7,0]]],3]",
		},
		{
			name:   "Ex4",
			number: "[[3,[2,[1,[7,3]]]],[6,[5,[4,[3,2]]]]]",
			index:  []int{11, 14},
			expect: "[[3,[2,[8,0]]],[9,[5,[4,[3,2]]]]]",
		},
		{
			name:   "Ex5",
			number: "[[3,[2,[8,0]]],[9,[5,[4,[3,2]]]]]",
			index:  []int{25, 28},
			expect: "[[3,[2,[8,0]]],[9,[5,[7,0]]]]",
		},
		{
			name:   "Ex6",
			number: "[[[[0,7],[[6,6],14]],[[17,6],[9,10]]],[[[9,13],[0,7]],8]]",
			index:  []int{11, 14},
			expect: "[[[[0,13],[0,20]],[[17,6],[9,10]]],[[[9,13],[0,7]],8]]",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expect, string(number(tt.number).Explode(tt.index)))
		})
	}
}

func TestNumberSplit(t *testing.T) {
	tests := []struct {
		name   string
		number string
		index  []int
		expect string
	}{
		{
			name:   "Ex1",
			number: "[[[[0,7],4],[15,[0,13]]],[1,1]]",
			index:  []int{13, 15},
			expect: "[[[[0,7],4],[[7,8],[0,13]]],[1,1]]",
		},
		{
			name:   "Ex2",
			number: "[[[[0,7],4],[[7,8],[0,13]]],[1,1]]",
			index:  []int{22, 24},
			expect: "[[[[0,7],4],[[7,8],[0,[6,7]]]],[1,1]]",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expect, string(number(tt.number).Split(tt.index)))
		})
	}
}

func TestNumberReduce(t *testing.T) {
	tests := []struct {
		name   string
		number string
		expect string
	}{
		{
			name:   "Ex1",
			number: "[[[[[9,8],1],2],3],4]",
			expect: "[[[[0,9],2],3],4]",
		},
		{
			name:   "Ex2",
			number: "[7,[6,[5,[4,[3,2]]]]]",
			expect: "[7,[6,[5,[7,0]]]]",
		},
		{
			name:   "Ex3",
			number: "[[6,[5,[4,[3,2]]]],1]",
			expect: "[[6,[5,[7,0]]],3]",
		},
		{
			name:   "Ex4",
			number: "[[3,[2,[1,[7,3]]]],[6,[5,[4,[3,2]]]]]",
			expect: "[[3,[2,[8,0]]],[9,[5,[7,0]]]]",
		},
		{
			name:   "Ex5",
			number: "[[3,[2,[8,0]]],[9,[5,[4,[3,2]]]]]",
			expect: "[[3,[2,[8,0]]],[9,[5,[7,0]]]]",
		},
		{
			name:   "Ex6",
			number: "[[[[[4,3],4],4],[7,[[8,4],9]]],[1,1]]",
			expect: "[[[[0,7],4],[[7,8],[6,0]]],[8,1]]",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expect, string(number(tt.number).Reduce()))
		})
	}
}
