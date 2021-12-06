package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

func main() {
	data, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatalln("Failed to read input:", err)
	}

	lines, err := parse(string(data))
	if err != nil {
		log.Fatalln("Failed to parse input:", err)
	}

	log.Println("Part one:", generateGrid(lines, false).Count())
	log.Println("Part two:", generateGrid(lines, true).Count())
}

type line struct {
	a, b coord
}

type coord struct {
	x, y int
}

type grid map[coord]int

func parse(input string) ([]line, error) {
	var lines []line
	for _, raw := range strings.Split(input, "\n") {
		var l line
		if _, err := fmt.Sscanf(raw, "%d,%d -> %d,%d", &l.a.x, &l.a.y, &l.b.x, &l.b.y); err != nil {
			return nil, fmt.Errorf("malformed line %s: %s", raw, err)
		}
		lines = append(lines, l)
	}
	return lines, nil
}

func generateGrid(lines []line, includeDiagonals bool) grid {
	g := make(grid)
	for _, l := range lines {
		// Vertical line
		if l.a.x == l.b.x {
			start, end := l.a.y, l.b.y
			if l.a.y > l.b.y {
				start, end = end, start
			}
			for y := start; y <= end; y++ {
				g[coord{x: l.a.x, y: y}]++
			}
			continue
		}
		// Horizontal line
		if l.a.y == l.b.y {
			start, end := l.a.x, l.b.x
			if l.a.x > l.b.x {
				start, end = end, start
			}
			for x := start; x <= end; x++ {
				g[coord{x: x, y: l.a.y}]++
			}
			continue
		}
		if !includeDiagonals {
			continue
		}
		// Diagonal line
		start, end := l.a, l.b
		if l.a.x > l.b.x {
			start, end = end, start
		}
		y := start.y
		for x := start.x; x <= end.x; x++ {
			g[coord{x: x, y: y}]++
			if start.y < end.y {
				y++
				continue
			}
			y--
		}
	}
	return g
}

func (g grid) Count() int {
	var count int
	for _, v := range g {
		if v > 1 {
			count++
		}
	}
	return count
}
