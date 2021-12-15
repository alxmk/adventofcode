package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"strings"
)

func main() {
	data, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatalln("Failed to read input:", err)
	}

	g := parse(string(data))
	log.Println("Part one:", g.Path(coord{0, 0}, coord{g.xmax, g.ymax}))
	g = g.Stonks(5)
	log.Println("Part two:", g.Path(coord{0, 0}, coord{g.xmax, g.ymax}))
}

func parse(input string) *grid {
	g := newGrid()
	lines := strings.Split(input, "\n")
	g.ymax = len(lines) - 1
	for y, line := range lines {
		g.xmax = len(line) - 1
		for x, r := range line {
			g.data[coord{x, y}] = int(r - '0')
		}
	}
	return g
}

func (g *grid) Stonks(times int) *grid {
	newG := newGrid()
	for dx := 0; dx <= times; dx++ {
		for dy := 0; dy <= times; dy++ {
			for x := 0; x <= g.xmax; x++ {
				for y := 0; y <= g.ymax; y++ {
					increment := dx + dy
					newG.data[coord{x + dx*(g.xmax+1), y + dy*(g.ymax+1)}] = (g.data[coord{x, y}]+increment)%10 + (g.data[coord{x, y}]+increment)/10
				}
			}
		}
	}
	newG.xmax, newG.ymax = (g.xmax+1)*times-1, (g.ymax+1)*times-1
	return newG
}

type coord struct {
	x, y int
}

func (c coord) String() string {
	return fmt.Sprintf("%d,%d", c.x, c.y)
}

func (c coord) Adjacent(xmax, ymax int) []coord {
	var adjacent []coord
	if x, y := c.x-1, c.y; x >= 0 {
		adjacent = append(adjacent, coord{x, y})
	}
	if x, y := c.x+1, c.y; x <= xmax {
		adjacent = append(adjacent, coord{x, y})
	}
	if x, y := c.x, c.y-1; y >= 0 {
		adjacent = append(adjacent, coord{x, y})
	}
	if x, y := c.x, c.y+1; y <= ymax {
		adjacent = append(adjacent, coord{x, y})
	}
	return adjacent
}

func (c coord) Distance(d coord) int {
	dx, dy := c.x-d.x, c.y-d.y
	if dx < 0 {
		dx *= -1
	}
	if dy < 0 {
		dy *= -1
	}
	return dx + dy
}

type grid struct {
	data       map[coord]int
	xmax, ymax int
}

func (g grid) String() string {
	var b strings.Builder
	for y := 0; y <= g.ymax; y++ {
		for x := 0; x <= g.xmax; x++ {
			b.WriteRune(rune(g.data[coord{x, y}] + '0'))
		}
		b.WriteRune('\n')
	}
	return b.String()
}

func newGrid() *grid { return &grid{data: make(map[coord]int)} }

func (g *grid) Path(from, to coord) int {
	openSet := map[coord]struct{}{
		from: {},
	}

	routes := make(map[coord]coord)

	gScore := defaultMaxMap{
		from: 0,
	}

	fScore := defaultMaxMap{
		from: gScore[from] + from.Distance(from),
	}

	for len(openSet) != 0 {
		lowestScore := math.MaxInt
		var current coord

		for c := range openSet {
			if score := fScore.Get(c); score < lowestScore {
				lowestScore = score
				current = c
			}
		}

		if current == to {
			return scorePath(to, from, routes, g)
		}

		delete(openSet, current)
		for _, neighbour := range current.Adjacent(g.xmax, g.ymax) {
			newG := gScore.Get(current)
			if newG != math.MaxInt64 {
				newG += g.data[neighbour]
			}

			if newG < gScore.Get(neighbour) {
				routes[neighbour], gScore[neighbour], fScore[neighbour], openSet[neighbour] =
					current, newG, newG+from.Distance(neighbour), struct{}{}
			}
		}
	}

	return -1
}

func scorePath(to, from coord, pmap map[coord]coord, g *grid) int {
	score := g.data[to]
	for to, ok := pmap[to]; ok && to != from; to, ok = pmap[to] {
		score += g.data[to]
	}
	return score
}

type defaultMaxMap map[coord]int

func (d defaultMaxMap) Get(c coord) int {
	if v, ok := d[c]; ok {
		return v
	}
	return math.MaxInt64
}
