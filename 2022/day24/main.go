package main

import (
	"bytes"
	"fmt"
	"log"
	"math"
	"os"
)

func main() {
	data, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatalln("Failed to read input:", err)
	}
	log.Println("Part one:", partOne(parse(data)))
	log.Println("Part two:", partTwo(parse(data)))
}

func partOne(s *solver) int {
	return s.Solve(xy{1, 0}, xy{s.w.max.x - 1, s.w.max.y}, 0)
}

func partTwo(s *solver) int {
	c0, c1 := xy{1, 0}, xy{s.w.max.x - 1, s.w.max.y}
	t1 := s.Solve(c0, c1, 0)
	t2 := s.Solve(c1, c0, t1)
	return s.Solve(c0, c1, t2)
}

func parse(input []byte) *solver {
	bz := make(blizzards)
	lines := bytes.Split(input, []byte{'\n'})
	w := world{tiles: make(map[xy]tile), max: xy{len(lines[0]) - 1, len(lines) - 1}}
	for y, line := range lines {
		for x, r := range line {
			switch r {
			case '#':
				w.tiles[xy{x, y}] = wall
				continue
			case '^':
				bz[blizzard{c: xy{x, y}, d: north}] = struct{}{}
			case '>':
				bz[blizzard{c: xy{x, y}, d: east}] = struct{}{}
			case '<':
				bz[blizzard{c: xy{x, y}, d: west}] = struct{}{}
			case 'v':
				bz[blizzard{c: xy{x, y}, d: south}] = struct{}{}
			}
			w.tiles[xy{x, y}] = floor
		}
	}
	s := &solver{w: &w, bscache: make(map[int]blizzards), blcache: make(map[int]map[xy]struct{}), scache: make(map[state]struct{})}
	s.blcache[1], s.bscache[1] = bz.Tick(&w)
	return s
}

type world struct {
	tiles map[xy]tile
	max   xy
}

type blizzards map[blizzard]struct{}

func (z blizzards) Tick(w *world) (map[xy]struct{}, blizzards) {
	bz := make(blizzards)
	bl := make(map[xy]struct{})
	for b := range z {
		next := xy{b.c.x + b.d.x, b.c.y + b.d.y}
		if w.tiles[next] == wall {
			switch b.d {
			case south:
				next.y = 1
			case north:
				next.y = w.max.y - 1
			case east:
				next.x = 1
			case west:
				next.x = w.max.x - 1
			}
		}
		b.c = next
		bl[next] = struct{}{}
		bz[b] = struct{}{}
	}
	return bl, bz
}

type solver struct {
	w       *world
	bscache map[int]blizzards
	blcache map[int]map[xy]struct{}
	scache  map[state]struct{}
}

func (s *solver) Solve(c0, cN xy, t0 int) int {
	start := state{c: c0, t: t0}
	s.scache[start] = struct{}{}

	queue := []state{start}

	var this state
	for len(queue) != 0 {
		this, queue = queue[0], queue[1:]
		for _, d := range dirs {
			next := state{c: xy{this.c.x + d.x, this.c.y + d.y}, t: this.t + 1}
			if _, ok := s.scache[next]; ok {
				// log.Println(next, "state cache hit")
				continue
			}
			if s.w.tiles[next.c] == wall {
				// log.Println(next, "wall hit")
				continue
			}
			if _, ok := s.blcache[next.t]; !ok {
				s.blcache[next.t], s.bscache[next.t] = s.bscache[next.t-1].Tick(s.w)
			}
			if _, ok := s.blcache[next.t][next.c]; ok {
				// log.Println(next, "blizzard hit")
				continue
			}
			if next.c == cN {
				return next.t
			}
			// log.Println(next, "viable")
			s.scache[next] = struct{}{}
			queue = append(queue, next)
		}
	}
	return math.MaxInt
}

type tile int

const (
	wall tile = iota
	floor
)

type xy struct {
	x, y int
}

type direction xy

var (
	east  = direction{1, 0}
	south = direction{0, 1}
	west  = direction{-1, 0}
	north = direction{0, -1}
	stay  = direction{0, 0}

	dirs = []direction{east, south, west, north, stay}
)

func (d direction) String() string {
	switch d {
	case east:
		return "east"
	case south:
		return "south"
	case west:
		return "west"
	case north:
		return "north"
	}
	return fmt.Sprintf("%d,%d", d.x, d.y)
}

type blizzard struct {
	c xy
	d direction
}

type state struct {
	c xy
	t int
}
