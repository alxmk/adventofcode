package main

import (
	"bytes"
	"log"
	"os"
	"strings"
)

func main() {
	data, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatalln("Failed to read input:", err)
	}

	log.Println("Part one:", partOne(parse(data)))
	log.Println("Part two:", partTwo(parse(data)))
}

func partOne(w *world) int {
	for i := 0; i < 10; i++ {
		w.Tick(i)
	}
	return ((w.max.x - w.min.x + 1) * (w.max.y - w.min.y + 1)) - len(w.elves)
}

func partTwo(w *world) int {
	for i := 0; ; i++ {
		if ok := w.Tick(i); !ok {
			return i + 1
		}
	}
}

func parse(input []byte) *world {
	lines := bytes.Split(input, []byte{'\n'})
	w := world{elves: make(map[xy]struct{}), max: xy{len(lines[0]) - 1, len(lines) - 1}}
	for y, line := range lines {
		for x, r := range line {
			if r == '#' {
				w.elves[xy{x, y}] = struct{}{}
			}
		}
	}
	return &w
}

type xy struct {
	x, y int
}

var (
	n  = xy{0, -1}
	ne = xy{1, -1}
	nw = xy{-1, -1}
	w  = xy{-1, 0}
	e  = xy{1, 0}
	s  = xy{0, 1}
	se = xy{1, 1}
	sw = xy{-1, 1}

	alldirs = []xy{n, ne, nw, w, e, s, se, sw}

	checkdirs = [][]xy{{n, ne, nw}, {s, se, sw}, {w, nw, sw}, {e, ne, se}}

	movedirs = []xy{n, s, w, e}
)

type world struct {
	elves    map[xy]struct{}
	min, max xy
}

func (w *world) Tick(i int) bool {
	var moved bool
	// Proposal stage
	proposed := make(map[xy]xy)
	duplicates := make(map[xy]int)
	for e := range w.elves {
		// First check if the elf is isolated
		isolated := true
		for _, d := range alldirs {
			if _, ok := w.elves[xy{e.x + d.x, e.y + d.y}]; ok {
				isolated = false
				break
			}
		}
		// If so then no move proposed
		if isolated {
			continue
		}
		// Otherwise try to identify the best move
		for j := i; j < i+4; j++ {
			found := true
			for _, d := range checkdirs[j%4] {
				if _, ok := w.elves[xy{e.x + d.x, e.y + d.y}]; ok {
					found = false
					break
				}
			}
			if !found {
				continue
			}
			newpos := xy{e.x + movedirs[j%4].x, e.y + movedirs[j%4].y}
			proposed[e], duplicates[newpos] = newpos, duplicates[newpos]+1
			break
		}
	}
	// Move stage
	for f, t := range proposed {
		if count := duplicates[t]; count > 1 {
			continue
		}
		delete(w.elves, f)
		w.elves[t] = struct{}{}
		moved = true
		if t.x > w.max.x {
			w.max.x = t.x
		}
		if t.y > w.max.y {
			w.max.y = t.y
		}
		if t.x < w.min.x {
			w.min.x = t.x
		}
		if t.y < w.min.y {
			w.min.y = t.y
		}
	}
	return moved
}

func (w world) String() string {
	var b strings.Builder
	for y := w.min.y; y <= w.max.y; y++ {
		for x := w.min.x; x <= w.max.x; x++ {
			if _, ok := w.elves[xy{x, y}]; ok {
				b.WriteRune('#')
				continue
			}
			b.WriteRune('.')
		}
		b.WriteRune('\n')
	}
	return b.String()
}
