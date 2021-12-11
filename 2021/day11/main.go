package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

func main() {
	data, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatalln("Failed to read input:", err)
	}

	o := parse(string(data))

	log.Println("Part one:", o.partOne(100))
	log.Println("Part two:", o.partTwo()+100)
}

type coord struct {
	x, y int
}

func (c coord) String() string {
	return fmt.Sprintf("%d,%d", c.x, c.y)
}

func (c coord) adjacent(o *octopi) []coord {
	var adj []coord
	for x := c.x - 1; x <= c.x+1; x++ {
		for y := c.y - 1; y <= c.y+1; y++ {
			if c.x == x && c.y == y {
				continue
			}
			if _, ok := o.data[coord{x, y}]; !ok {
				continue
			}
			adj = append(adj, coord{x, y})
		}
	}
	return adj
}

type octopi struct {
	data       map[coord]int
	xmax, ymax int
}

func newOctopi() *octopi { return &octopi{data: make(map[coord]int)} }

func parse(input string) *octopi {
	o := newOctopi()
	for y, line := range strings.Split(input, "\n") {
		for x, raw := range line {
			o.data[coord{x, y}] = int(raw - '0')
			if x > o.xmax {
				o.xmax = x
			}
		}
		if y > o.ymax {
			o.ymax = y
		}
	}
	return o
}

func (o *octopi) partOne(steps int) int {
	var flashes int
	for i := 0; i < steps; i++ {
		flashes += o.Step()
	}
	return flashes
}

func (o *octopi) partTwo() int {
	for i := 1; ; i++ {
		if flashes := o.Step(); flashes == len(o.data) {
			return i
		}
	}
}

func (o *octopi) Step() int {
	// First increment everything by one
	for c := range o.data {
		o.data[c]++
	}
	// Then loop until we didn't get another flash
	flashes := make(map[coord]struct{})
	for {
		var flash bool
		for c := range o.data {
			if o.data[c] > 9 {
				if _, ok := flashes[c]; !ok {
					flashes[c] = struct{}{}
					flash = true
					for _, a := range c.adjacent(o) {
						o.data[a]++
					}
				}
			}
		}
		if !flash {
			break
		}
	}
	// Set all flashees to 0
	for c := range flashes {
		o.data[c] = 0
	}
	return len(flashes)
}

func (o octopi) String() string {
	var b strings.Builder
	for y := 0; y <= o.ymax; y++ {
		for x := 0; x <= o.xmax; x++ {
			b.WriteRune(rune(o.data[coord{x, y}]) + '0')
		}
		b.WriteRune('\n')
	}
	return b.String()
}
