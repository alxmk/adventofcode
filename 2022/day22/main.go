package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	data, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatalln("Failed to read input:", err)
	}

	w, r := parse(data)
	log.Println("Part one:", solve(w, r, p1))
	log.Println("Part two:", solve(w, r, p2(seams)))
}

func solve(w *world, r []instruction, f wrappingFunc) int {
	c := character{position: w.Start(), bearing: right}

	for _, i := range r {
		if i.move != 0 {
			c.position, c.bearing = w.Move(c.position, c.bearing, i.move, f)
			continue
		}
		c.bearing = c.bearing.Turn(i.turn)
	}

	return ((c.position.y+1)*1000 + ((c.position.x + 1) * 4) + c.bearing.index())
}

func parse(input []byte) (*world, []instruction) {
	world, route, _ := bytes.Cut(input, []byte{'\n', '\n'})
	return parseWorld(world), parseRoute(route)
}

func parseWorld(input []byte) *world {
	lines := bytes.Split(input, []byte{'\n'})
	w := world{tiles: make(map[xy]tile), max: xy{len(lines[0]), len(lines)}}

	for y, line := range lines {
		for x := range line {
			var t tile
			switch lines[y][x] {
			case ' ':
			case '.':
				t = floor
			case '#':
				t = wall
			}
			w.tiles[xy{x, y}] = t
		}
	}

	return &w
}

func parseRoute(input []byte) []instruction {
	var is []instruction

	var b bytes.Buffer
	for _, r := range input {
		switch r {
		case 'R':
			if b.Len() != 0 {
				v, _ := strconv.Atoi(b.String())
				is = append(is, instruction{move: v})
				b.Reset()
			}
			is = append(is, instruction{turn: right})
		case 'L':
			if b.Len() != 0 {
				v, _ := strconv.Atoi(b.String())
				is = append(is, instruction{move: v})
				b.Reset()
			}
			is = append(is, instruction{turn: left})
		default:
			b.WriteByte(r)
		}
	}
	if b.Len() != 0 {
		v, _ := strconv.Atoi(b.String())
		is = append(is, instruction{move: v})
		b.Reset()
	}

	return is
}

type xy struct {
	x, y int
}

type direction xy

var (
	right = direction{1, 0}
	down  = direction{0, 1}
	left  = direction{-1, 0}
	up    = direction{0, -1}

	dirs = []direction{right, down, left, up}
)

func (d direction) index() int {
	switch d {
	case right:
		return 0
	case down:
		return 1
	case left:
		return 2
	case up:
		return 3
	}
	return -1
}

func (d direction) Turn(t direction) direction {
	var i int
	switch t {
	case right:
		i = d.index() + 1
	case left:
		i = d.index() - 1
	}
	switch i {
	case -1:
		return dirs[3]
	case 4:
		return dirs[0]
	}
	return dirs[i]
}

func (d direction) String() string {
	switch d {
	case right:
		return "right"
	case left:
		return "left"
	case down:
		return "down"
	case up:
		return "up"
	}
	return fmt.Sprintf("{%d,%d}", d.x, d.y)
}

type instruction struct {
	move int
	turn direction
}

type tile int

const (
	empty tile = iota
	floor
	wall
)

type world struct {
	tiles map[xy]tile
	max   xy
}

func (w world) String(c character) string {
	var b strings.Builder
	for y := 0; y <= w.max.y; y++ {
		for x := 0; x <= w.max.x; x++ {
			if c.position.x == x && c.position.y == y {
				switch c.bearing {
				case right:
					b.WriteRune('>')
				case left:
					b.WriteRune('<')
				case up:
					b.WriteRune('^')
				case down:
					b.WriteRune('v')
				}
				continue
			}
			switch w.tiles[xy{x, y}] {
			case floor:
				b.WriteRune('.')
			case wall:
				b.WriteRune('#')
			case empty:
				b.WriteRune(' ')
			}
		}
		b.WriteRune('\n')
	}
	return b.String()
}

func (w *world) Start() xy {
	for x := 0; x <= w.max.x; x++ {
		p := xy{x, 0}
		if t := w.tiles[p]; t == floor {
			return p
		}
	}
	return xy{-1, -1}
}

func (w *world) Move(from xy, h direction, amount int, f wrappingFunc) (xy, direction) {
	var ph direction
	p, prev := from, from
	for i := 0; i < amount; i++ {
		ph = h
		p.x, p.y = p.x+h.x, p.y+h.y
		if w.tiles[p] == empty {
			p, h = f(w, p, prev, h)
			if w.tiles[p] == empty {
				panic("Bad transform")
			}
		}
		if w.tiles[p] == wall {
			p, h = prev, ph
			break
		}
		prev = p
	}
	return p, h
}

type wrappingFunc func(*world, xy, xy, direction) (xy, direction)

var (
	p1 wrappingFunc = func(w *world, p, _ xy, h direction) (xy, direction) {
		switch h {
		case right:
			for x := 0; ; x++ {
				if w.tiles[xy{x, p.y}] != empty {
					return xy{x, p.y}, h
				}
			}
		case left:
			for x := w.max.x; ; x-- {
				if w.tiles[xy{x, p.y}] != empty {
					return xy{x, p.y}, h
				}
			}
		case up:
			for y := w.max.y; ; y-- {
				if w.tiles[xy{p.x, y}] != empty {
					return xy{p.x, y}, h
				}
			}
		case down:
			for y := 0; ; y++ {
				if w.tiles[xy{p.x, y}] != empty {
					return xy{p.x, y}, h
				}
			}
		}
		return xy{-1, -1}, h
	}
)

func p2(s []seam) wrappingFunc {
	return func(w *world, p xy, prev xy, h direction) (xy, direction) {
		for _, seam := range s {
			if seam.On(prev, h) {
				if ok, np, nh := seam.Traverse(prev, h); ok {
					return np, nh
				}
			}
		}
		panic("Seam failed")
	}
}

type character struct {
	position xy
	bearing  direction
}

type seam struct {
	a, b edge
	name string
}

type edge struct {
	vertices  [2]xy
	transform map[direction]direction
}

func (e edge) On(p xy, d direction) bool {
	if _, ok := e.transform[d]; !ok {
		return false
	}
	return ((p.x <= e.vertices[1].x && p.x >= e.vertices[0].x) || (p.x >= e.vertices[1].x && p.x <= e.vertices[0].x)) && ((p.y <= e.vertices[1].y && p.y >= e.vertices[0].y) || (p.y >= e.vertices[1].y && p.y <= e.vertices[0].y))
}

func (e edge) Delta(p xy) xy {
	if e.vertices[0].x != e.vertices[1].x {
		return xy{p.x - e.vertices[0].x, 0}
	}
	return xy{0, p.y - e.vertices[0].y}
}

func (s seam) Traverse(p xy, h direction) (bool, xy, direction) {
	if s.a.On(p, h) {
		delta := s.a.Delta(p)
		if (s.a.vertices[0].x > s.a.vertices[1].x) != (s.b.vertices[0].x > s.b.vertices[1].x) {
			delta.x *= -1
		}
		if (s.a.vertices[0].y > s.a.vertices[1].y) != (s.b.vertices[0].y > s.b.vertices[1].y) {
			delta.y *= -1
		}
		if (s.a.vertices[0].x == s.a.vertices[1].x) != (s.b.vertices[0].x == s.b.vertices[1].x) {
			delta = xy{delta.y, delta.x}
		}
		return true, xy{s.b.vertices[0].x + delta.x, s.b.vertices[0].y + delta.y}, s.a.transform[h]
	}
	delta := s.b.Delta(p)
	if (s.a.vertices[0].x > s.a.vertices[1].x) != (s.b.vertices[0].x > s.b.vertices[1].x) {
		delta.x *= -1
	}
	if (s.a.vertices[0].y > s.a.vertices[1].y) != (s.b.vertices[0].y > s.b.vertices[1].y) {
		delta.y *= -1
	}
	if (s.a.vertices[0].x == s.a.vertices[1].x) != (s.b.vertices[0].x == s.b.vertices[1].x) {
		delta = xy{delta.y, delta.x}
	}
	return true, xy{s.a.vertices[0].x + delta.x, s.a.vertices[0].y + delta.y}, s.b.transform[h]
}

func (s seam) On(p xy, d direction) bool {
	return s.a.On(p, d) || s.b.On(p, d)
}

var seams = []seam{
	{ // A
		a:    edge{vertices: [2]xy{{50, 0}, {50, 49}}, transform: map[direction]direction{left: right}},
		b:    edge{vertices: [2]xy{{0, 149}, {0, 100}}, transform: map[direction]direction{left: right}},
		name: "A",
	},
	{ // B
		a:    edge{vertices: [2]xy{{50, 0}, {99, 0}}, transform: map[direction]direction{up: right}},
		b:    edge{vertices: [2]xy{{0, 150}, {0, 199}}, transform: map[direction]direction{left: down}},
		name: "B",
	},
	{ // C
		a:    edge{vertices: [2]xy{{50, 50}, {50, 99}}, transform: map[direction]direction{left: down}},
		b:    edge{vertices: [2]xy{{0, 100}, {49, 100}}, transform: map[direction]direction{up: right}},
		name: "C",
	},
	{ // D
		a:    edge{vertices: [2]xy{{49, 150}, {49, 199}}, transform: map[direction]direction{right: up}},
		b:    edge{vertices: [2]xy{{50, 149}, {99, 149}}, transform: map[direction]direction{down: left}},
		name: "D",
	},
	{ // E
		a:    edge{vertices: [2]xy{{100, 49}, {149, 49}}, transform: map[direction]direction{down: left}},
		b:    edge{vertices: [2]xy{{99, 50}, {99, 99}}, transform: map[direction]direction{right: up}},
		name: "E",
	},
	{ // F
		a:    edge{vertices: [2]xy{{149, 0}, {149, 49}}, transform: map[direction]direction{right: left}},
		b:    edge{vertices: [2]xy{{99, 149}, {99, 100}}, transform: map[direction]direction{right: left}},
		name: "F",
	},
	{ // G
		a:    edge{vertices: [2]xy{{149, 0}, {100, 0}}, transform: map[direction]direction{up: up}},
		b:    edge{vertices: [2]xy{{49, 199}, {0, 199}}, transform: map[direction]direction{down: down}},
		name: "G",
	},
}
