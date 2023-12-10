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

	sys, starting := parseSystem(string(data))
	tunnel := findTunnel(sys, starting)

	log.Println("Part one:", len(tunnel)/2)
	log.Println("Part two:", partTwo(sys, tunnel))
}

func partTwo(sys system, tunnel map[xy]struct{}) int {
	open := make(map[xy]struct{})
	for l := range sys {
		if _, ok := tunnel[l]; ok {
			continue
		}
		open[l] = struct{}{}
	}
	for o := range open {
		var intersects bool
		var previous pipe
		for x := 0; x <= o.x; x++ {
			if _, ok := tunnel[xy{x, o.y}]; !ok {
				continue
			}
			p := sys[xy{x, o.y}]
			switch {
			case p == ns, p == nw && previous == se, p == sw && previous == ne:
				intersects = !intersects
				previous = '.'
			case p == ne || p == se:
				previous = p
			}
		}
		if !intersects {
			delete(open, o)
		}
	}
	return len(open)
}

func findTunnel(sys system, starting xy) map[xy]struct{} {
	tunnel := make(map[xy]struct{})
	var dir xy
	pos := starting
	pipe := sys[starting]
	for _, d := range []xy{n, s, e, w} {
		if pipe.ConnectsTo(sys[starting.Move(d)], d) {
			dir = d
			break
		}
	}
	for i := 0; ; i++ {
		tunnel[pos] = struct{}{}
		pos = pos.Move(dir)
		pipe = sys[pos]
		if i != 0 && pos == starting {
			return tunnel
		}
		dir = pipe.Next(dir)
	}
}

func parseSystem(input string) (system, xy) {
	sys := make(system)
	var starting xy
	for y, line := range strings.Split(input, "\n") {
		for x, r := range line {
			sys[xy{x, y}] = pipe(r)
			if pipe(r) == start {
				starting = xy{x, y}
			}
		}
	}
	connections := make(map[xy]struct{})
	for _, d := range []xy{n, s, e, w} {
		if start.ConnectsTo(sys[starting.Move(d)], d) {
			connections[d] = struct{}{}
		}
	}
	cn := start.ConnectsTo(sys[starting.Move(n)], n)
	cs := start.ConnectsTo(sys[starting.Move(s)], s)
	ce := start.ConnectsTo(sys[starting.Move(e)], e)
	cw := start.ConnectsTo(sys[starting.Move(w)], w)
	switch {
	case cn && cs:
		sys[starting] = ns
	case ce && cw:
		sys[starting] = ew
	case cn && cw:
		sys[starting] = nw
	case cn && ce:
		sys[starting] = ne
	case cs && cw:
		sys[starting] = sw
	case cs && ce:
		sys[starting] = se
	}
	return sys, starting
}

type system map[xy]pipe

type xy struct {
	x, y int
}

func (x xy) Move(d xy) xy {
	return xy{x.x + d.x, x.y + d.y}
}

type pipe rune

func (p pipe) Next(d xy) xy {
	switch {
	case p == ns && d == n, p == nw && d == e, p == ne && d != s:
		return n
	case p == ns && d != n, p == sw && d == e, p == se && d != n:
		return s
	case p == ne && d == s, p == ew && d == e, p == se && d == n:
		return e
	case p == nw && d != e, p == ew && d != e, p == sw && d != e:
		return w
	default:
		panic(fmt.Sprintf("unprocessable %v %s", d, string(p)))
	}
}

func (p pipe) ConnectsTo(q pipe, d xy) bool {
	switch d {
	case n:
		return q == sw || q == se || q == ns
	case s:
		return q == nw || q == ne || q == ns
	case e:
		return q == ew || q == nw || q == sw
	case w:
		return q == ew || q == ne || q == se
	}
	return false
}

var (
	ns    pipe = '|'
	ew    pipe = '-'
	ne    pipe = 'L'
	nw    pipe = 'J'
	sw    pipe = '7'
	se    pipe = 'F'
	start pipe = 'S'

	n = xy{0, -1}
	s = xy{0, 1}
	e = xy{1, 0}
	w = xy{-1, 0}
)
