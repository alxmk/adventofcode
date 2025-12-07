package main

import (
	_ "embed"
	"fmt"
	"strings"
)

//go:embed input.txt
var input string

func main() {
	var startx int
	splitters := make(map[[2]int]struct{})
	lines := strings.Split(input, "\n")
	for y, line := range lines {
		for x, r := range line {
			switch r {
			case 'S':
				startx = x
			case '^':
				splitters[[2]int{x, y}] = struct{}{}
			}
		}
	}
	beams := map[int]int{startx: 1}
	partOne, partTwo := 0, 1
	for y := range len(lines) - 1 {
		for beam, count := range beams {
			if _, ok := splitters[[2]int{beam, y}]; ok {
				partOne++
				partTwo += count
				delete(beams, beam)
				beams[beam-1] += count
				beams[beam+1] += count
			}
		}
	}
	fmt.Println("Part one:", partOne)
	fmt.Println("Part two:", partTwo)
}
