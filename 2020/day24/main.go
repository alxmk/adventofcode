package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
	"text/scanner"
)

func main() {
	data, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatalln("Failed to read input:", err)
	}
	g := make(grid)
	for _, line := range strings.Split(string(data), "\n") {
		g.Flip(line)
	}
	log.Println("Part one:", g.Count())

	for i := 0; i < 100; i++ {
		g = g.Iterate()
	}
	log.Println("Part two:", g.Count())
}

type grid map[int]map[int]bool

func (g grid) Flip(path string) {
	var s scanner.Scanner
	s.Init(strings.NewReader(path))
	var x, y int
	for r := s.Next(); r != scanner.EOF; r = s.Next() {
		switch r {
		case 'w':
			x = x - 2
		case 'e':
			x = x + 2
		case 'n':
			q := s.Peek()
			switch q {
			case 'e':
				s.Next()
				x, y = x+1, y+1
			case 'w':
				s.Next()
				x, y = x-1, y+1
			default:
				y = y + 2
			}
		case 's':
			q := s.Peek()
			switch q {
			case 'e':
				s.Next()
				x, y = x+1, y-1
			case 'w':
				s.Next()
				x, y = x-1, y-1
			default:
				y = y - 2
			}
		}
	}
	if _, ok := g[x]; !ok {
		g[x] = make(map[int]bool)
	}
	g[x][y] = !g[x][y]
}

func (g grid) Count() int {
	var count int
	for x := range g {
		for _, black := range g[x] {
			if black {
				count++
			}
		}
	}
	return count
}

func (g grid) State(x, y int) bool {
	if _, ok := g[x]; ok {
		return g[x][y]
	}
	return false
}

func (g grid) Set(x, y int) {
	if _, ok := g[x]; !ok {
		g[x] = make(map[int]bool)
	}
	g[x][y] = true
}

func (g grid) Iterate() grid {
	cache := make(map[string]struct{})
	newg := make(grid)
	for x := range g {
		for y, s := range g[x] {
			if !s {
				continue
			}
			// Check adjacent tiles
			var adjacentBlack int
			for _, n := range neighbours {
				i, j := x+n[0], y+n[1]
				if g.State(i, j) {
					adjacentBlack++
					continue
				}
				// If they're white we need to see if they themselves need flipping
				// If they're black we'll catch them in the outer loop
				// Check the cache to see if we already checked them
				str := fmt.Sprintf("%d,%d", i, j)
				if _, ok := cache[str]; ok {
					continue
				}
				// Now we will check them so add to cache
				cache[str] = struct{}{}
				var innerAdjBlack int
				for _, o := range neighbours {
					k, l := i+o[0], j+o[1]
					if g.State(k, l) {
						innerAdjBlack++
					}
				}
				if innerAdjBlack == 2 {
					newg.Set(i, j)
				}
			}
			if adjacentBlack == 1 || adjacentBlack == 2 {
				newg.Set(x, y)
			}
		}
	}
	return newg
}

var neighbours = [][2]int{{1, 1}, {2, 0}, {-1, 1}, {-1, -1}, {-2, 0}, {1, -1}}
