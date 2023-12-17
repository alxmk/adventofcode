package main

import (
	"bytes"
	"log"
	"math"
	"os"
)

func main() {
	data, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatalln("Failed to read input:", err)
	}

	w, from, to := parseWorld(data)

	log.Println("Part one:", w.Dijkstra(from, to, 1, 3))
	log.Println("Part two:", w.Dijkstra(from, to, 4, 10))
}

func parseWorld(input []byte) (world, xy, xy) {
	w := make(world)
	var max xy
	for y, line := range bytes.Split(input, []byte{'\n'}) {
		for x, r := range line {
			w[xy{x, y}] = r
			max.x = x
		}
		max.y = y
	}
	return w, xy{0, 0}, max
}

type world map[xy]byte

type xy struct {
	x, y int
}

type dir xy

func (w world) Distance(a xy) int {
	if _, ok := w[a]; !ok {
		return math.MaxInt
	}
	return int(w[a] - '0')
}

func (a dir) NextDir() []dir {
	switch a {
	case n, s:
		return []dir{e, w}
	case e, w:
		return []dir{n, s}
	}
	panic("fak")
}

var (
	n = dir{0, -1}
	s = dir{0, 1}
	e = dir{1, 0}
	w = dir{-1, 0}
)

type defaultMaxMap map[state]int

func (d defaultMaxMap) Get(k state) int {
	if v, ok := d[k]; ok {
		return v
	}
	return math.MaxInt
}

// State is the location and previous direction
type state struct {
	a xy
	d dir
}

func (wo world) Dijkstra(a, b xy, minsteps, maxsteps int) int {
	dist := make(defaultMaxMap)
	open := make(map[state]int)

	open[state{a, e}], open[state{a, s}] = 0, 0

	best := math.MaxInt

	for len(open) != 0 {
		for o, c := range open {
			delete(open, o)
			// If we're at the end, track the best answer so far
			if o.a == b {
				best = min(c, best)
			}
			// Shortcut if this is a worse route here or already
			// longer than the best route to the destination
			if c > dist.Get(state{o.a, o.d}) || c > best {
				continue
			}
			// Try each valid next direction
			for _, dir := range o.d.NextDir() {
				nexto := state{d: dir}
				nextc := c
				// Try each valid number of steps
				for i := 1; i <= maxsteps; i++ {
					nexto.a = xy{o.a.x + dir.x*i, o.a.y + dir.y*i}
					nc := wo.Distance(nexto.a)
					// Going further in this direction is invalid
					if nc == math.MaxInt {
						break
					}
					nextc += nc
					if i < minsteps {
						continue
					}
					// If this is a better route here then add to the open
					// set to investigate further, otherwise it's a dead
					// end or redundant at best so carry on
					if dist.Get(state{nexto.a, nexto.d}) > nextc {
						open[nexto] = nextc
						dist[state{nexto.a, nexto.d}] = nextc
					}
				}
			}
		}
	}
	return best
}
