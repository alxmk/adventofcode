package main

import (
	"fmt"
	"math"
)

func main() {
	partone()
	parttwo()
}

func partone() {
	const target = 312051
	var i int
	var sideLength int
	var maximum int

	// Calculate the attributes of the layer the target is in
	for i = 1; true; i++ {
		sideLength = (i * 2) - 1

		maximum = int(math.Pow(float64(sideLength), 2))

		if maximum > target {
			break
		}
	}

	// Where are the corners?
	bottomRight := maximum
	bottomLeft := maximum - (sideLength - 1)
	topLeft := maximum - ((sideLength - 1) * 2)
	topRight := maximum - ((sideLength - 1) * 3)

	var x, y int

	switch {
	case target > bottomLeft:
		// We're on the bottom row
		y = -1 * (sideLength - 1) / 2
		x = (sideLength-1)/2 - (bottomRight - target)
	case target > topLeft:
		// We're on the left column
		x = -1 * (sideLength - 1) / 2
		y = -1*(sideLength-1)/2 + (bottomLeft - target)
	case target > topRight:
		// We're on the top row
		y = (sideLength - 1) / 2
		x = -1*(sideLength-1)/2 + (topLeft - target)
	default:
		// We're on the right column
		x = (sideLength - 1) / 2
		y = (sideLength-1)/2 + (topRight - target)
	}

	distance := math.Abs(float64(x)) + math.Abs(float64(y))

	fmt.Println("Part 1 answer is", distance)
}

func parttwo() {
	const target = 312051

	grid := newWorld()

	c := coord{
		x: 1,
		y: 0,
	}

	sidelength := 3
	count := 1

	for {
		if _, ok := grid[c.x]; !ok {
			grid[c.x] = make(map[int]int)
		}

		var newval int
		for _, co := range c.Adjacent() {
			newval += grid.Get(co)
		}
		grid[c.x][c.y] = newval

		if newval > target {
			fmt.Println("Part two:", newval)
			break
		}

		count++
		if sidelength*sidelength == count {
			sidelength += 2
		}

		max := (sidelength - 1) / 2

		if c.x == max {
			if c.y == max {
				c = coord{x: c.x - 1, y: c.y}
			} else {
				c = coord{x: c.x, y: c.y + 1}
			}
		} else if c.x == max*-1 {
			if c.y == max*-1 {
				c = coord{x: c.x + 1, y: c.y}
			} else {
				c = coord{x: c.x, y: c.y - 1}
			}
		} else {
			if c.y == max {
				c = coord{x: c.x - 1, y: c.y}
			} else {
				c = coord{x: c.x + 1, y: c.y}
			}
		}
	}
}

type world map[int]map[int]int

func newWorld() world {
	return map[int]map[int]int{
		0: map[int]int{
			0: 1,
		},
	}
}

func (w world) Get(c coord) int {
	if _, ok := w[c.x]; ok {
		if v, ok := w[c.x][c.y]; ok {
			return v
		}
	}

	return 0
}

type coord struct {
	x, y int
}

func (c coord) Adjacent() []coord {
	return []coord{
		coord{x: c.x, y: c.y + 1},
		coord{x: c.x + 1, y: c.y + 1},
		coord{x: c.x - 1, y: c.y + 1},
		coord{x: c.x, y: c.y - 1},
		coord{x: c.x + 1, y: c.y - 1},
		coord{x: c.x - 1, y: c.y - 1},
		coord{x: c.x + 1, y: c.y},
		coord{x: c.x - 1, y: c.y},
	}
}

type dir int

const (
	east dir = iota
	north
	west
	south
)

func (d dir) Next() dir {
	if d+1 > south {
		return east
	}
	return d + 1
}
