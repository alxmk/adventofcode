package main

import (
	"io/ioutil"
	"log"
	"sort"
	"strings"
)

func main() {
	data, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatalln("Failed to read input:", err)
	}

	g := parse(string(data))

	log.Println("Part one:", partOne(g))
	log.Println("Part two:", partTwo(g))
}

func newGrid() *grid {
	return &grid{data: make(map[coord]int)}
}

type grid struct {
	data       map[coord]int
	xmax, ymax int
}

type coord struct {
	x, y int
}

func (c coord) adjacent(xmax, ymax int, exclude ...coord) []coord {
	var coords []coord
	exclusionMap := make(map[coord]struct{})
	for _, e := range exclude {
		exclusionMap[e] = struct{}{}
	}
	down := coord{x: c.x - 1, y: c.y}
	if _, ok := exclusionMap[down]; !ok && c.x-1 >= 0 {
		coords = append(coords, down)
	}
	up := coord{x: c.x + 1, y: c.y}
	if _, ok := exclusionMap[up]; !ok && c.x+1 <= xmax {
		coords = append(coords, up)
	}
	left := coord{x: c.x, y: c.y - 1}
	if _, ok := exclusionMap[left]; !ok && c.y-1 >= 0 {
		coords = append(coords, left)
	}
	right := coord{x: c.x, y: c.y + 1}
	if _, ok := exclusionMap[right]; !ok && c.y+1 <= ymax {
		coords = append(coords, right)
	}
	return coords
}

func parse(input string) *grid {
	g := newGrid()
	lines := strings.Split(input, "\n")
	g.xmax, g.ymax = len(lines[0])-1, len(lines)-1
	for y, line := range lines {
		for x, raw := range line {
			g.data[coord{x: x, y: y}] = int(raw - '0')
		}
	}
	return g
}

func findLowPoints(g *grid) []coord {
	var lowPoints []coord
	for x := 0; x <= g.xmax; x++ {
		for y := 0; y <= g.ymax; y++ {
			c := coord{x: x, y: y}
			v := g.data[c]
			var notLowPoint bool
			for _, a := range c.adjacent(g.xmax, g.ymax) {
				if g.data[a] <= v {
					notLowPoint = true
					break
				}
			}
			if notLowPoint {
				continue
			}
			lowPoints = append(lowPoints, c)
		}
	}
	return lowPoints
}

func findBasins(g *grid, lowPoints []coord) []int {
	var basins []int
	for _, p := range lowPoints {
		inBasin := findBasin(g, p, map[coord]struct{}{p: {}})
		basins = append(basins, len(inBasin))
	}
	return basins
}

func findBasin(g *grid, start coord, inBasin map[coord]struct{}) map[coord]struct{} {
	var basinSlice []coord
	for b := range inBasin {
		basinSlice = append(basinSlice, b)
	}
	for _, a := range start.adjacent(g.xmax, g.ymax, basinSlice...) {
		if v := g.data[a]; v != 9 {
			inBasin[a] = struct{}{}
			findBasin(g, a, inBasin)
			continue
		}
	}
	return inBasin
}

func partOne(g *grid) int {
	var sum int
	for _, c := range findLowPoints(g) {
		sum += g.data[c] + 1
	}
	return sum
}

func partTwo(g *grid) int {
	sizes := findBasins(g, findLowPoints(g))
	sort.Ints(sizes)
	return sizes[len(sizes)-1] * sizes[len(sizes)-2] * sizes[len(sizes)-3]
}
