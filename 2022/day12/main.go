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

	w, from, to := parse(data)
	l, _ := w.Astar(from, to)
	log.Println("Part one:", l)

	min := math.MaxInt
	for c, r := range w {
		if r == 'a' {
			if p, ok := w.Astar(c, to); ok {
				if p < min {
					min = p
				}
			}
		}
	}
	log.Println("Part two:", min)
}

func parse(input []byte) (world, xy, xy) {
	w := make(world)
	var from, to xy
	lines := bytes.Split(input, []byte{'\n'})
	for y, line := range lines {
		for x, r := range line {
			switch r {
			case 'S':
				from = xy{x, y}
				w[from] = 'a'
			case 'E':
				to = xy{x, y}
				w[to] = 'z'
			default:
				w[xy{x, y}] = r
			}
		}
	}
	return w, from, to
}

var (
	north = xy{0, -1}
	south = xy{0, 1}
	west  = xy{-1, 0}
	east  = xy{1, 0}

	cardinal = []xy{north, south, west, east}
)

type xy struct {
	x, y int
}

func (x xy) Distance(y xy) int {
	return absint(y.x-x.x) + absint(y.y-x.y)
}

func (x xy) Next(dir xy) xy {
	return xy{x.x + dir.x, x.y + dir.y}
}

func absint(i int) int {
	if i < 0 {
		return i * -1
	}
	return i
}

type world map[xy]byte

func (w world) Score(x xy) int {
	return int(w[x] - 'a')
}

type defaultMaxMap map[xy]int

func (d defaultMaxMap) Get(c xy) int {
	if v, ok := d[c]; ok {
		return v
	}
	return math.MaxInt
}

func (w world) Astar(a, b xy) (int, bool) {
	open := map[xy]struct{}{a: {}}
	path := make(map[xy]xy)
	g := defaultMaxMap{a: 0}
	f := defaultMaxMap{a: g[a] + a.Distance(a)}

	for len(open) != 0 {
		min := math.MaxInt
		var c xy

		for o := range open {
			if score := f.Get(o); score < min {
				min, c = score, o
			}
		}

		if c == b {
			return reconstructPath(path, c) - 1, true
		}

		delete(open, c)
		for _, d := range cardinal {
			n := c.Next(d)
			tg := g.Get(c)
			if tg != math.MaxInt64 {
				tg++
			}
			if w.Score(n) > w.Score(c)+1 {
				tg = math.MaxInt64
			}
			if tg < g.Get(n) {
				path[n], g[n], f[n], open[n] = c, tg, tg+a.Distance(n), struct{}{}
			}
		}
	}

	return -1, false
}

func reconstructPath(path map[xy]xy, c xy) int {
	totalPath := []xy{c}
	for c, ok := path[c]; ok; c, ok = path[c] {
		totalPath = append(totalPath, c)
	}
	return len(totalPath)
}
