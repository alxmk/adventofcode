package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"strconv"
	"strings"
)

var (
	xmax, xmin, ymax, ymin int
)

func main() {
	data, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatalln("Failed to read input", err)
	}

	var clay []coord
	xmin, ymin = math.MaxInt32, math.MaxInt32

	log.Println("Building grid")

	for _, line := range strings.Split(string(data), "\n") {
		a, b, c, d, e := parseLine(line)
		clay = append(clay, clayCoords(a, b, c, d, e)...)

		switch d {
		case "x":
			if a > xmax {
				xmax = a
			}
			if a < xmin {
				xmin = a
			}
			for i := b; i <= c; i++ {
				if i > ymax {
					ymax = i
				}
				if i < ymin {
					ymin = i
				}
			}
		case "y":
			if a > ymax {
				ymax = a
			}
			if a < ymin {
				ymin = a
			}
			for i := b; i <= c; i++ {
				if i > xmax {
					xmax = i
				}
				if i < xmin {
					xmin = i
				}
			}
		}
		log.Println("Parsed", line)
	}

	log.Println(xmax, xmin, ymax)

	grid := make(world, xmax-xmin+3)

	for x := range grid {
		grid[x] = make([]*tile, ymax+1)
	}

	for _, c := range clay {
		grid[c.x-xmin+1][c.y] = &tile{clay: true}
	}

	log.Println("Grid assembled")

	source := coord{x: 500 - xmin, y: 0}

	state := grid.String(source)
	fmt.Println(state)

	// ioutil.WriteFile("out.txt", []byte(state), os.ModePerm)

	// return

	for i := 0; i < 20; i++ {
		log.Println("Iteration", i)
		done := grid.Drop(source)
		if done {
			fmt.Println("Drop done")
			break
		}
		fmt.Println(grid.String(source))
	}

	fmt.Println(grid.String(source))

	fmt.Println(grid.CountWet() - ymin + 1)

	fmt.Println(grid.CountWater())
}

func (w world) CountWater() int {
	var sum int
	for x := range w {
		for y := range w[x] {
			if t, ok := w.Get(x, y); ok && t.water {
				sum++
			}
		}
	}

	return sum
}

func (w world) CountWet() int {
	var sum int
	for x := range w {
		for y := range w[x] {
			if t, ok := w.Get(x, y); ok && (t.wet || t.water) {
				sum++
			}
		}
	}

	return sum
}

func (w world) Drop(source coord) bool {
	done := true

	c := coord{x: source.x, y: source.y + 1}

	// log.Println("Drop from", source.x, source.y)

	for {
		t, ok := w.Get(c.x, c.y)
		if !ok {
			return done
		}
		if !t.Permeable() {
			c.y--
			break
		}

		if !t.wet {
			done = false
		}
		t.wet = true
		c.y++
	}

	// Look left
	var leftextent int
	var leftbounded bool
	for i := c.x - 1; i >= 0; i-- {
		next, _ := w.Get(i, c.y)
		if !next.Permeable() {
			leftextent = i + 1
			leftbounded = true
			break
		}
		if !next.wet {
			done = false
		}
		next.wet = true
		below, ok := w.Get(i, c.y+1)
		if !ok {
			leftextent = i
			break
		}
		if below.Permeable() {
			leftextent = i
			for j := 0; j < 10; j++ {
				if w.Drop(coord{x: i, y: c.y}) {
					break
				}
			}
			break
		}
	}

	// Look right
	var rightextent int
	var rightbounded bool
	for i := c.x + 1; i < len(w); i++ {
		next, _ := w.Get(i, c.y)
		if !next.Permeable() {
			rightextent = i - 1
			rightbounded = true
			break
		}
		if !next.wet {
			done = false
		}
		next.wet = true
		below, ok := w.Get(i, c.y+1)
		if !ok {
			rightextent = i
			break
		}
		if below.Permeable() {
			rightextent = i
			for j := 0; j < 10; j++ {
				if w.Drop(coord{x: i, y: c.y}) {
					break
				}
			}
			break
		}
	}

	if leftbounded && rightbounded {
		// log.Println("Drop", source.x, source.y, "not done - filling at", c.y)
		done = false
		for i := leftextent; i <= rightextent; i++ {
			t, _ := w.Get(i, c.y)
			t.water = true
		}
	}

	return done
}

type world [][]*tile

func (w world) Get(x, y int) (*tile, bool) {
	if y >= ymax+1 || x >= xmax || x < 0 || y < 0 {
		return nil, false
	}

	if t := w[x][y]; t == nil {
		w[x][y] = &tile{}
	}
	return w[x][y], true
}

func (w world) String(source coord) string {
	var buf bytes.Buffer

	for y := range w[0] {
		for x := range w {
			if source.x == x && source.y == y {
				buf.WriteString("+")
				continue
			}
			t, ok := w.Get(x, y)
			if !ok {
				buf.WriteString(".")
				continue
			}
			if t.water {
				buf.WriteString("~")
				continue
			}
			if t.wet {
				buf.WriteString("|")
				continue
			}
			if t.clay {
				buf.WriteString("#")
				continue
			}
			buf.WriteString(".")
		}
		buf.WriteString("\n")
	}

	return buf.String()
}

type coord struct {
	x, y int
}

type tile struct {
	clay  bool // true for clay, false for sand
	water bool // whether there's water in the tile
	wet   bool // whether water has ever been in the tile
}

func (t *tile) Permeable() bool {
	// If it's sand and doesn't have water in it it's permeable
	return !t.clay && !t.water
}

func clayCoords(a, b, c int, d, e string) []coord {
	var coords []coord

	switch d {
	case "x":
		for i := b; i <= c; i++ {
			coords = append(coords, coord{x: a, y: i})
		}
	case "y":
		for i := b; i <= c; i++ {
			coords = append(coords, coord{x: i, y: a})
		}
	}

	return coords
}

func parseLine(line string) (int, int, int, string, string) {
	parts := strings.Split(line, ", ")

	d := string(parts[0][0])
	a, _ := strconv.Atoi(parts[0][2:])
	e := string(parts[1][0])

	subparts := strings.Split(parts[1][2:], "..")
	b, _ := strconv.Atoi(subparts[0])
	c, _ := strconv.Atoi(subparts[1])

	return a, b, c, d, e
}
