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

	log.Println("Part one:", partOne(parseWorld(data)))
	log.Println("Part two:", partTwo(parseWorld(data)))
}

func partOne(w world, max xy) int {
	return w.Graph(xy{1, 0}, xy{max.x - 1, max.y}, true).longestPath(xy{1, 0}, xy{max.x - 1, max.y}, make(map[xy]struct{}))
}

func partTwo(w world, max xy) int {
	return w.Graph(xy{1, 0}, xy{max.x - 1, max.y}, false).longestPath(xy{1, 0}, xy{max.x - 1, max.y}, make(map[xy]struct{}))
}

func parseWorld(input []byte) (world, xy) {
	w := make(world)
	var max xy
	for y, line := range bytes.Split(input, []byte{'\n'}) {
		for x, b := range line {
			w[xy{x, y}] = b
			max.x = x
		}
		max.y = y
	}
	return w, max
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

func (w world) Score(a, d xy, slippery bool) bool {
	switch w[a] {
	case '#':
		return false
	case '.', '>', '<', '^', 'v':
		if !slippery {
			return true
		}
		prev := xy{a.x + (d.x * -1), a.y + (d.y * -1)}
		switch w[prev] {
		case '.':
			switch w[a] {
			case '>':
				return d != west
			case '<':
				return d != east
			case 'v':
				return d != north
			case '^':
				return d != south
			}
			return true
		case '>':
			return d == east
		case '^':
			return d == north
		case '<':
			return d == west
		case 'v':
			return d == south
		}
	}
	return false
}

func (d xy) NextDir() []xy {
	switch d {
	case north:
		return []xy{north, east, west}
	case east:
		return []xy{north, south, west}
	case west:
		return []xy{north, east, south}
	case south:
		return []xy{south, east, west}
	}
	panic("undefined dir")
}

func copyMap(m map[xy]struct{}) map[xy]struct{} {
	n := make(map[xy]struct{})
	for k, v := range m {
		n[k] = v
	}
	return n
}

type graph map[xy]map[xy]int

func (g graph) longestPath(a, b xy, visited map[xy]struct{}) int {
	if a == b {
		return 0
	}
	visited[a] = struct{}{}
	var options []xy
	for v := range g[a] {
		if _, ok := visited[v]; ok {
			continue
		}
		options = append(options, v)
	}
	best := -1
	for _, o := range options {
		p := g.longestPath(o, b, copyMap(visited))
		if p == -1 {
			continue
		}
		best = max(best, g[a][o]+p)
	}
	return best
}

func (w world) Graph(start, end xy, slippery bool) graph {
	g := make(graph)

	open := map[xy]struct{}{start: {}}
	done := make(map[xy]struct{})

	for len(open) != 0 {
		for o := range open {
			if _, ok := g[o]; !ok {
				g[o] = make(map[xy]int)
			}
			delete(open, o)
			done[o] = struct{}{}
			for v, d := range w.findNextVertices(o, end, slippery) {
				if _, ok := done[v]; ok {
					continue
				}
				if v != end {
					open[v] = struct{}{}
				}
				g[o][v] = d
				if slippery {
					continue
				}
				if _, ok := g[v]; !ok {
					g[v] = make(map[xy]int)
				}
				g[v][o] = d
			}
		}
	}
	return g
}

func (w world) findNextVertices(start, end xy, slippery bool) map[xy]int {
	vertices := make(map[xy]int)
	for _, d := range cardinal {
		n := start.Next(d)
		if !w.Score(n, d, slippery) {
			continue
		}
		end, dist := w.followVertex(start, n, end, slippery)
		vertices[end] = max(vertices[end], dist)
	}
	return vertices
}

func (w world) followVertex(origin, first, end xy, slippery bool) (xy, int) {
	var dist int
	c, prev := first, origin
	for {
		dist++
		if c == end {
			return c, dist
		}
		var valid []xy
		for _, d := range cardinal {
			n := c.Next(d)
			if n == prev || !w.Score(n, d, slippery) {
				continue
			}
			valid = append(valid, n)
		}
		if len(valid) == 1 {
			c, prev = valid[0], c
			continue
		}
		return c, dist
	}
}
