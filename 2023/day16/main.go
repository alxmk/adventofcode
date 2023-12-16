package main

import (
	"bytes"
	"log"
	"os"
)

func main() {
	data, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatalln("Failed to read input:", err)
	}

	log.Println("Part one:", partOne(parseContraption(data)))
	log.Println("Part two:", partTwo(parseContraption(data)))
}

func partOne(c contraption) int {
	return solve(c, xy{0, 0}, e)
}

func partTwo(c contraption) int {
	var best int
	for x := 0; x <= c.max.x; x++ {
		best = max(best, solve(c, xy{x, 0}, s))
		best = max(best, solve(c, xy{x, c.max.y}, n))
	}
	for y := 0; y <= c.max.y; y++ {
		best = max(best, solve(c, xy{0, y}, e))
		best = max(best, solve(c, xy{c.max.x, y}, w))
	}
	return best
}

func solve(c contraption, s, d xy) int {
	energised := make(map[xy]struct{})
	beams := []beam{{p: s, d: d, path: make(map[xy]int)}}
	// Cache to make sure we don't trace the same beam multiple times
	beamcache := map[key]struct{}{{p: beams[0].p, d: beams[0].d}: {}}
	for len(beams) != 0 {
		b := beams[0]
		for b.p.x >= 0 && b.p.x <= c.max.x && b.p.y >= 0 && b.p.y <= c.max.y {
			// Shortcut loops - can hit a slanted mirror twice
			if b.path[b.p] >= 2 {
				break
			}
			b.path[b.p]++
			energised[b.p] = struct{}{}
			r := c.grid[b.p]
			switch r {
			case '.':
				b.p.x += b.d.x
				b.p.y += b.d.y
			default:
				next := b.d.Turns(r)
				b.d = next[0]
				b.p.x += next[0].x
				b.p.y += next[0].y
				if len(next) == 2 {
					np := xy{b.p.x + next[1].x, b.p.y + next[1].y}
					if _, ok := beamcache[key{p: np, d: next[1]}]; !ok {
						beamcache[key{p: np, d: next[1]}] = struct{}{}
						beams = append(beams, beam{p: np, d: next[1], path: make(map[xy]int)})
					}
				}
			}
		}
		beams = beams[1:]
	}
	return len(energised)
}

type beam struct {
	p    xy
	d    xy
	path map[xy]int
}

type key struct {
	p, d xy
}

type xy struct {
	x, y int
}

func (a xy) Turns(m byte) []xy {
	switch m {
	case '|':
		switch a {
		case e, w:
			return []xy{n, s}
		case n, s:
			return []xy{a}
		}
	case '-':
		switch a {
		case e, w:
			return []xy{a}
		case n, s:
			return []xy{e, w}
		}
	case '\\':
		switch a {
		case e:
			return []xy{s}
		case w:
			return []xy{n}
		case n:
			return []xy{w}
		case s:
			return []xy{e}
		}
	case '/':
		switch a {
		case e:
			return []xy{n}
		case w:
			return []xy{s}
		case n:
			return []xy{e}
		case s:
			return []xy{w}
		}
	}
	panic(m)
}

var (
	n = xy{0, -1}
	s = xy{0, 1}
	e = xy{1, 0}
	w = xy{-1, 0}
)

type contraption struct {
	grid map[xy]byte
	max  xy
}

func parseContraption(input []byte) contraption {
	c := contraption{grid: make(map[xy]byte)}
	for y, line := range bytes.Split(input, []byte{'\n'}) {
		for x, r := range line {
			c.grid[xy{x, y}] = r
			c.max.x = x
		}
		c.max.y = y
	}
	return c
}
