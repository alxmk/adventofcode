package main

import (
	_ "embed"
	"fmt"
	"strings"
)

//go:embed input.txt
var input string

func main() {
	one, two := solve(input)
	fmt.Println("Part one:", one)
	fmt.Println("Part two:", two)
}

func solve(input string) (int, int) {
	var count1, count2 int
	grid := make(map[xy]rune)
	var xmax, ymax int
	for y, line := range strings.Split(input, "\n") {
		xmax = len(line) - 1
		for x, r := range line {
			grid[xy{x, y}] = r
		}
		count1 += strings.Count(line, "XMAS") + strings.Count(line, "SAMX")
		ymax = y
	}

	// Transform 90 degrees
	for x := 0; x <= xmax; x++ {
		var b strings.Builder
		for y := 0; y <= ymax; y++ {
			b.WriteRune(grid[xy{x, y}])
		}
		line := b.String()
		count1 += strings.Count(line, "XMAS") + strings.Count(line, "SAMX")
	}

	centres := make(map[xy]struct{})

	// Transform -45 degrees
	for y := 0; y <= ymax+xmax; y++ {
		var b strings.Builder
		i := y
		var endx, endy int
		for x := 0; x <= xmax; x++ {
			if _, ok := grid[xy{x, i}]; !ok {
				i--
				continue
			}
			endx, endy = x, i
			b.WriteRune(grid[xy{x, i}])
			i--
		}
		line := b.String()
		count1 += strings.Count(line, "XMAS") + strings.Count(line, "SAMX")
		startx, starty := endx-len(line)+1, endy+len(line)-1
		for idx := strings.LastIndex(line, "MAS"); idx != -1; idx = strings.LastIndex(line[:idx], "MAS") {
			centres[xy{startx + idx + 1, starty - idx - 1}] = struct{}{}
		}
		for idx := strings.LastIndex(line, "SAM"); idx != -1; idx = strings.LastIndex(line[:idx], "SAM") {
			centres[xy{startx + idx + 1, starty - idx - 1}] = struct{}{}
		}
	}

	// Transform 45 degrees
	for y := -1 * xmax; y <= ymax; y++ {
		var b strings.Builder
		i := y
		var endx, endy int
		for x := 0; x <= xmax; x++ {
			if _, ok := grid[xy{x, i}]; !ok {
				i++
				continue
			}
			endx, endy = x, i
			b.WriteRune(grid[xy{x, i}])
			i++
		}
		line := b.String()
		count1 += strings.Count(line, "XMAS") + strings.Count(line, "SAMX")
		startx, starty := endx-len(line)+1, endy-len(line)+1
		for idx := strings.LastIndex(line, "MAS"); idx != -1; idx = strings.LastIndex(line[:idx], "MAS") {
			if _, ok := centres[xy{startx + idx + 1, starty + idx + 1}]; ok {
				count2++
			}
		}
		for idx := strings.LastIndex(line, "SAM"); idx != -1; idx = strings.LastIndex(line[:idx], "SAM") {
			if _, ok := centres[xy{startx + idx + 1, starty + idx + 1}]; ok {
				count2++
			}
		}
	}
	return count1, count2
}

type xy [2]int
