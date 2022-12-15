package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	data, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatalln("Failed to read input:", err)
	}

	log.Println("Part one:", partOne(string(data), 2000000))
	log.Println("Part two:", partTwo(string(data), xy{4000000, 4000000}))
}

func partOne(input string, target int) int {
	emptySet := make(map[xy]struct{})
	var items []xy
	for _, line := range strings.Split(input, "\n") {
		var s, b xy
		fmt.Sscanf(line, "Sensor at x=%d, y=%d: closest beacon is at x=%d, y=%d", &s.x, &s.y, &b.x, &b.y)
		empty(s, b, target, emptySet)
		items = append(items, s, b)
	}
	for _, c := range items {
		delete(emptySet, c)
	}
	return len(emptySet)
}

func partTwo(input string, space xy) int {
	var areas []area
	for _, line := range strings.Split(input, "\n") {
		var s, b xy
		fmt.Sscanf(line, "Sensor at x=%d, y=%d: closest beacon is at x=%d, y=%d", &s.x, &s.y, &b.x, &b.y)
		areas = append(areas, area{centre: s, size: s.Distance(b)})
	}
	for y := 0; y <= space.y; y++ {
		for x := 0; x <= space.x; x++ {
			var dx int
			var found bool
			this := xy{x, y}
			for _, a := range areas {
				if found, dx = a.Next(this); found {
					x += dx - 1
					break
				}
			}
			if !found {
				return x*4000000 + y
			}
		}
	}
	return -1
}

type xy struct {
	x, y int
}

type area struct {
	centre xy
	size   int
}

func (a area) Next(c xy) (bool, int) {
	if a.Contains(c) {
		return true, a.centre.x - c.x + a.size - absint(c.y-a.centre.y) + 1
	}
	return false, 0
}

func (a area) Contains(c xy) bool {
	return c.Distance(a.centre) <= a.size
}

func (c xy) Distance(to xy) int {
	return absint(to.x-c.x) + absint(to.y-c.y)
}

func empty(s, b xy, target int, emptySet map[xy]struct{}) {
	dist := s.Distance(b)
	for x := s.x - dist; x <= s.x+dist; x++ {
		this := xy{x, target}
		if this.Distance(s) <= dist {
			emptySet[this] = struct{}{}
		}
	}
}

func absint(i int) int {
	if i < 0 {
		return -1 * i
	}
	return i
}
