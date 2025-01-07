package main

import (
	_ "embed"
	"fmt"
	"math"
	"strings"
)

//go:embed input.txt
var input string

func main() {
	c, antennae, max := parse(input)
	fmt.Println("Part one:", partOne(c, antennae, max))
	fmt.Println("Part two:", partTwo(c, antennae, max))
}

type xy [2]float64

type city map[xy]rune

func parse(input string) (city, map[rune][]xy, xy) {
	c := make(city)
	antennae := make(map[rune][]xy)
	var max xy
	for y, line := range strings.Split(input, "\n") {
		max = xy{float64(len(line) - 1), float64(y)}
		for x, r := range line {
			if r != '.' {
				antennae[r] = append(antennae[r], xy{float64(x), float64(y)})
			}
			c[xy{float64(x), float64(y)}] = r
		}
	}
	return c, antennae, max
}

func partOne(c city, antennae map[rune][]xy, _ xy) int {
	antinodes := make(map[xy]struct{})
	for _, nodes := range antennae {
		for i, n := range nodes {
			for j, o := range nodes {
				if i == j {
					continue
				}
				potential := xy{2*o[0] - n[0], 2*o[1] - n[1]}
				if _, ok := c[potential]; ok {
					antinodes[potential] = struct{}{}
				}
			}
		}
	}
	return len(antinodes)
}

func partTwo(cty city, antennae map[rune][]xy, max xy) int {
	antinodes := make(map[xy]struct{})
	for _, nodes := range antennae {
		for i, n := range nodes {
			antinodes[n] = struct{}{}
			for j, o := range nodes {
				if i == j {
					continue
				}
				// y1 = mx1 + c - (1)
				// y2 = mx2 + c - (2)
				// (1) - (2)
				// y1 - y2 = m(x1 - x2) - (3)
				// Solve (3) for m
				// m = y1 - y2 / x1 - x2
				// Solve (1) for c
				// c = y1 - mx1
				m := (n[1] - o[1]) / (n[0] - o[0])
				c := n[1] - (m * n[0])

				for x := float64(0); x <= max[0]; x++ {
					y := m*x + c
					// Round to 2 decimal places to account for floating point error
					y = math.Round(y*100) / 100
					// Is it an integer?
					if y != math.Trunc(y) {
						continue
					}
					// Is it on the map?
					if _, ok := cty[xy{x, y}]; !ok {
						continue
					}
					antinodes[xy{x, y}] = struct{}{}
				}
			}
		}
	}
	return len(antinodes)
}
