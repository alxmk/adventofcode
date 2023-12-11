package main

import (
	"log"
	"os"
	"strings"
)

func main() {
	data, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatalln("Failed to read input:", err)
	}

	log.Println("Part one:", solve(parseGalaxies(string(data)), 2))
	log.Println("Part two:", solve(parseGalaxies(string(data)), 1000000))
}

func solve(u universe, expansion int) int {
	var sum int
	for i, g := range u.galaxies {
		for j := i + 1; j < len(u.galaxies); j++ {
			h := u.galaxies[j]
			d := g.Distance(h)
			for x := min(g.x, h.x); x <= max(g.x, h.x); x++ {
				if _, ok := u.emptyX[x]; ok {
					d += expansion - 1
				}
			}
			for y := min(g.y, h.y); y <= max(g.y, h.y); y++ {
				if _, ok := u.emptyY[y]; ok {
					d += expansion - 1
				}
			}
			sum += d
		}
	}
	return sum
}

func parseGalaxies(input string) (u universe) {
	u.emptyX, u.emptyY = make(map[int]struct{}), make(map[int]struct{})
	for y, line := range strings.Split(input, "\n") {
		var found bool
		for x, r := range line {
			if y == 0 {
				u.emptyX[x] = struct{}{}
			}
			u.xmax = x
			switch r {
			case '.':
			case '#':
				u.galaxies = append(u.galaxies, xy{x, y})
				found = true
				delete(u.emptyX, x)
			}
		}
		u.ymax = y
		if !found {
			u.emptyY[y] = struct{}{}
		}
	}
	return u
}

type universe struct {
	galaxies       []xy
	emptyX, emptyY map[int]struct{}
	xmax, ymax     int
}

type xy struct {
	x, y int
}

func (a xy) Distance(b xy) int {
	return max(a.x-b.x, b.x-a.x) + max(a.y-b.y, b.y-a.y)
}
