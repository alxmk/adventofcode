package main

import (
	_ "embed"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

func main() {
	lines := strings.Split(input, "\n")
	left, right, rmap := make([]int, len(lines)), make([]int, len(lines)), make(map[int]int)
	for i, line := range lines {
		parts := strings.Fields(line)
		left[i], _ = strconv.Atoi(parts[0])
		right[i], _ = strconv.Atoi(parts[1])
		rmap[right[i]]++
	}
	sort.Ints(left)
	sort.Ints(right)
	var difference int
	var similarity int
	for i, l := range left {
		similarity += rmap[l] * l
		d := l - right[i]
		if d < 0 {
			difference -= d
			continue
		}
		difference += d
	}
	fmt.Println("Part one:", difference)
	fmt.Println("Part two:", similarity)
}
