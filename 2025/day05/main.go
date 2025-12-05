package main

import (
	_ "embed"
	"fmt"
	"slices"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

func main() {
	parts := strings.Split(input, "\n\n")
	var ranges []ir
	for _, line := range strings.Split(parts[0], "\n") {
		hyphen := strings.Index(line, "-")
		a, _ := strconv.Atoi(line[:hyphen])
		b, _ := strconv.Atoi(line[hyphen+1:])
		ranges = append(ranges, ir{a, b})
	}
	var partOne int
	for _, line := range strings.Split(parts[1], "\n") {
		v, _ := strconv.Atoi(line)
		for _, r := range ranges {
			if r.contains(v) {
				partOne++
				break
			}
		}
	}
	fmt.Println("Part one", partOne)
	var partTwo int
	for _, r := range consolidate(ranges) {
		partTwo += 1 + r[1] - r[0]
	}
	fmt.Println("Part two:", partTwo)
}

func consolidate(ranges []ir) (consolidated []ir) {
	copy := slices.Clone(ranges)
	for i := 0; i < len(copy); i++ {
		var reduced bool
		for j := i + 1; j < len(copy); j++ {
			if copy[i].overlaps(copy[j]) {
				consolidated = append(consolidated, copy[i].reduce(copy[j]))
				copy = append(copy[:j], copy[j+1:]...)
				reduced = true
				break
			}
		}
		if reduced {
			continue
		}
		consolidated = append(consolidated, copy[i])
	}
	if len(consolidated) == len(ranges) {
		return consolidated
	}
	return consolidate(consolidated)
}

type ir [2]int

func (i ir) contains(v int) bool {
	return v >= i[0] && v <= i[1]
}

func (i ir) overlaps(j ir) bool {
	return (i[0] >= j[0] && i[0] <= j[1]) ||
		(j[0] >= i[0] && j[0] <= i[1])
}

func (i ir) reduce(j ir) ir {
	return ir{min(i[0], j[0]), max(i[1], j[1])}
}
