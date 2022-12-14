package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"strings"
)

func main() {
	data, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatalln("Failed to read input:", err)
	}

	log.Println("Part one:", partOne(parse(strings.Split(string(data), "\n"))))
	log.Println("Part two:", partTwo(parse(strings.Split(string(data), "\n"))))
}

func partOne(w *world) int {
	var i int
	for done := false; !done; done = w.InsertSand() {
		i++
	}
	return i - 1
}

func partTwo(w *world) int {
	var i int
	for done := false; !done; done = w.InsertSandFloor() {
		i++
	}
	return i
}

type tile int

const (
	air tile = iota
	rock
	sand
)

type xy struct {
	x, y int
}

type world struct {
	tiles      map[xy]tile
	xmin, xmax int
	ymax       int
}

func (w world) String() string {
	var b strings.Builder
	for y := 0; y <= w.ymax+2; y++ {
		for x := w.xmin; x <= w.xmax; x++ {
			switch w.tiles[xy{x, y}] {
			case air:
				b.WriteRune('.')
			case rock:
				b.WriteRune('#')
			case sand:
				b.WriteRune('o')
			}
		}
		b.WriteRune('\n')
	}
	return b.String()
}

var (
	directions = []xy{{0, 1}, {-1, 1}, {1, 1}}
)

func (w *world) InsertSand() bool {
	at := xy{500, 0}
	for {
		var moves bool
		for _, d := range directions {
			next := xy{at.x + d.x, at.y + d.y}
			switch w.tiles[next] {
			case air:
				at, moves = next, true
			default:
			}
			if moves {
				break
			}
		}
		if !moves {
			w.tiles[at] = sand
			return false
		}
		if at.y > w.ymax {
			return true
		}
	}
}

func (w *world) InsertSandFloor() bool {
	at := xy{500, 0}
	for {
		var moves bool
		for _, d := range directions {
			next := xy{at.x + d.x, at.y + d.y}
			switch w.tiles[next] {
			case air:
				at, moves = next, true
			default:
			}
			if moves {
				break
			}
		}
		if !moves || at.y == w.ymax+1 {
			w.tiles[at] = sand
			return (at.x == 500 && at.y == 0)
		}
	}
}

func parse(input []string) *world {
	w := world{tiles: make(map[xy]tile), xmin: math.MaxInt}
	for _, line := range input {
		var current xy
		for i, cstr := range strings.Split(line, " -> ") {
			var coord xy
			fmt.Sscanf(cstr, "%d,%d", &coord.x, &coord.y)
			if coord.y > w.ymax {
				w.ymax = coord.y
			}
			if coord.x > w.xmax {
				w.xmax = coord.x
			}
			if coord.x < w.xmin {
				w.xmin = coord.x
			}
			if i != 0 {
				iter := &iterator{a: current, b: coord}
				for ok, next := iter.Next(); ok; ok, next = iter.Next() {
					w.tiles[next] = rock
				}
			}
			current = coord
		}
	}
	return &w
}

type iterator struct {
	a, b xy
	i    int
}

func (i *iterator) Next() (bool, xy) {
	defer func() { i.i++ }()
	if i.a.x == i.b.x {
		if i.a.y < i.b.y {
			return i.a.y+i.i <= i.b.y, xy{i.a.x, i.a.y + i.i}
		}
		return i.a.y-i.i >= i.b.y, xy{i.a.x, i.a.y - i.i}
	}
	if i.a.y != i.b.y {
		panic("uh oh")
	}
	if i.a.x < i.b.x {
		return i.a.x+i.i <= i.b.x, xy{i.a.x + i.i, i.a.y}
	}
	return i.a.x-i.i >= i.b.x, xy{i.a.x - i.i, i.a.y}
}
