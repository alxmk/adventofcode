package main

import (
	_ "embed"
	"fmt"
	"math"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

func main() {
	fmt.Println("Part one:", partOne(input))
	fmt.Println("Part two:", partTwo(input))
}

func partOne(input string) int {
	var problems []problem
	for line := range strings.SplitSeq(input, "\n") {
		for i, part := range strings.Fields(line) {
			switch part {
			case "*", "+":
				problems[i].operator = part
			default:
				v, _ := strconv.Atoi(part)
				if i >= len(problems) {
					problems = append(problems, problem{values: []int{v}})
					continue
				}
				problems[i].values = append(problems[i].values, v)
			}
		}
	}
	var partOne int
	for _, p := range problems {
		partOne += p.solve()
	}
	return partOne
}

func partTwo(input string) int {
	var problems []problem
	lines := strings.Split(input, "\n")
	var fields [][]string
	for _, line := range lines {
		fields = append(fields, strings.Fields(line))
	}
	var start int
	for i := range len(fields[0]) {
		var sectionLength int
		for _, f := range fields {
			sectionLength = max(sectionLength, len(f[i]))
		}
		problems = append(problems, parseVertical(lines, start, start+sectionLength-1))
		start = 1 + start + sectionLength
	}
	var partTwo int
	for _, p := range problems {
		partTwo += p.solve()
	}
	return partTwo
}

func parseVertical(lines []string, start, end int) (p problem) {
	for x := end; x >= start; x-- {
		var pow, v int
		for y := len(lines) - 1; y >= 0; y-- {
			switch lines[y][x] {
			case ' ':
				continue
			case '*', '+':
				p.operator = string([]byte{lines[y][x]})
			default:
				v += int(lines[y][x]-'0') * int(math.Pow10(pow))
				pow++
			}
		}
		p.values = append(p.values, v)
	}
	return p
}

type problem struct {
	values   []int
	operator string
}

func (p problem) solve() (answer int) {
	if p.operator == "*" {
		answer = 1
	}
	for _, v := range p.values {
		switch p.operator {
		case "+":
			answer += v
		case "*":
			answer *= v
		}
	}
	return answer
}
