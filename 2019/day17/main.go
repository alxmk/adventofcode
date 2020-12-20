package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"

	"github.com/alxmk/adventofcode/2019/day2/intcode"
)

func main() {
	data, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatalln("Failed to read input file:", err)
	}

	programme, err := intcode.Parse(string(data))
	if err != nil {
		log.Fatalln("Failed to parse input as intcode programme:", err)
	}

	out := make(chan int64)

	go func() {
		if err := programme.Copy().Run(nil, out); err != nil {
			log.Fatalln("Intcode computer failed:", err)
		}
	}()

	scaffold := &world{tiles: make(map[coordinate]rune)}

	var x, y int
	for o := range out {
		if x > scaffold.xmax {
			scaffold.xmax = x
		}
		if y > scaffold.ymax {
			scaffold.ymax = y
		}

		switch o {
		case 10:
			y++
			x = 0
		default:
			scaffold.tiles[coordinate{x, y}] = rune(o)
			x++
		}
	}

	log.Println("Part one:", scaffold.SumOfAlignmentParameters())

	// Activate robo-boogie
	programme[0] = 2

	in := make(chan int64)
	out = make(chan int64)
	go func() {
		if err := programme.Copy().Run(in, out); err != nil {
			log.Fatalln("Intcode computer failed:", err)
		}
	}()

	go func() {
		for _, input := range []string{m, a, b, c, "n"} {
			for _, r := range input {
				in <- int64(r)
			}
			in <- int64('\n')
		}
		close(in)
	}()

	var result int64
	for o := range out {
		result = o
	}

	log.Println("Part two:", result)
}

var (
	m = `A,B,B,A,C,B,C,C,B,A`
	a = `R,10,R,8,L,10,L,10`
	b = `R,8,L,6,L,6`
	c = `L,10,R,10,L,6`
)

type coordinate struct {
	x, y int
}

func (c coordinate) Next(dir coordinate) coordinate {
	return coordinate{c.x + dir.x, c.y + dir.y}
}

func (c coordinate) String() string {
	return fmt.Sprintf("(%d, %d)", c.x, c.y)
}

func (c coordinate) AlignmentParameter() int {
	return c.x * c.y
}

var (
	up    = coordinate{0, -1}
	down  = coordinate{0, 1}
	left  = coordinate{-1, 0}
	right = coordinate{1, 0}

	directions = []coordinate{up, down, left, right}
)

type world struct {
	tiles      map[coordinate]rune
	xmax, ymax int
}

func (w world) String() string {
	var out strings.Builder
	for y := 0; y <= w.ymax; y++ {
		for x := 0; x <= w.xmax; x++ {
			out.WriteRune(w.tiles[coordinate{x, y}])
		}
		out.WriteRune('\n')
	}
	return out.String()
}

func (w world) FindIntersections() []coordinate {
	var intersections []coordinate
	for y := 0; y <= w.ymax; y++ {
		for x := 0; x <= w.xmax; x++ {
			loc := coordinate{x, y}
			if w.tiles[loc] != '#' {
				continue
			}
			isIntersection := true
			for _, dir := range directions {
				if r, ok := w.tiles[loc.Next(dir)]; !ok || r != '#' {
					isIntersection = false
					break
				}
			}
			if isIntersection {
				intersections = append(intersections, loc)
			}
		}
	}
	return intersections
}

func (w world) SumOfAlignmentParameters() int {
	var sumOfAPs int
	for _, intersection := range w.FindIntersections() {
		sumOfAPs += intersection.AlignmentParameter()
	}
	return sumOfAPs
}
