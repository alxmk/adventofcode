package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

func main() {
	one, two := exec(input)
	fmt.Println("Part one:", one)
	fmt.Println("Part two:", two)
}

func exec(input string) (int, int) {
	var sumOne, sumTwo int
	for _, line := range strings.Split(input, "\n") {
		parts := strings.Split(line, ": ")
		y, _ := strconv.Atoi(parts[0])
		var vs []int
		for _, part := range strings.Fields(parts[1]) {
			v, _ := strconv.Atoi(part)
			vs = append(vs, v)
		}
		if solve(y, vs[0], vs[1:]) {
			sumOne += y
			sumTwo += y
			continue
		}
		if solveTwo(y, vs[0], vs[1:]) {
			sumTwo += y
		}
	}
	return sumOne, sumTwo
}

func solve(y, x int, vs []int) bool {
	if len(vs) == 0 {
		return y == x
	}
	return solve(y, x+vs[0], vs[1:]) || solve(y, x*vs[0], vs[1:])
}

func solveTwo(y, x int, vs []int) bool {
	if len(vs) == 0 {
		return y == x
	}
	return solveTwo(y, x+vs[0], vs[1:]) || solveTwo(y, x*vs[0], vs[1:]) || solveTwo(y, func() int {
		c, _ := strconv.Atoi(fmt.Sprintf("%d%d", x, vs[0]))
		return c
	}(), vs[1:])
}
