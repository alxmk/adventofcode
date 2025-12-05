package main

import (
	_ "embed"
	"fmt"
	"strings"
)

//go:embed input.txt
var input string

func main() {
	p1, p2 := solve(input)
	fmt.Println("Part one:", p1)
	fmt.Println("Part two:", p2)
}

func solve(input string) (partOne, partTwo int) {
	topology := make(map[[2]int]rune)
	for y, line := range strings.Split(input, "\n") {
		for x, r := range line {
			topology[[2]int{x, y}] = r
		}
	}
	for xy, r := range topology {
		if r != '0' {
			continue
		}
		reached := make(map[[2]int]struct{})
		partTwo += reachableSummits(xy, topology, reached)
		partOne += len(reached)
	}
	return
}

func reachableSummits(trailhead [2]int, topology map[[2]int]rune, reached map[[2]int]struct{}) (sum int) {
	height := topology[trailhead]
	x, y := trailhead[0], trailhead[1]
	for _, adj := range [][2]int{{x - 1, y}, {x, y - 1}, {x + 1, y}, {x, y + 1}} {
		if topology[adj] != height+1 {
			continue
		}
		if height+1 == '9' {
			reached[adj] = struct{}{}
			sum += 1
			continue
		}
		sum += reachableSummits(adj, topology, reached)
	}
	return sum
}
