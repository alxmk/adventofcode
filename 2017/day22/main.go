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
		log.Fatalln("Failed to read input", err)
	}

	lines := strings.Split(string(data), "\n")

	world, ymax, xmax := parseWorld(lines)

	fmt.Println(world.String(xmax, ymax))

	v := &virus{
		facing: north,
	}

	var infectioncount int
	for i := 0; i < 10000; i++ {
		if v.Tick(world, true) {
			infectioncount++
		}
		if v.x > xmax {
			xmax = v.x
		}
		if v.x*-1 > xmax {
			xmax = v.x * -1
		}
		if v.y > ymax {
			ymax = v.y
		}
		if v.y*-1 > ymax {
			ymax = v.y * -1
		}
		// fmt.Println(world.String(xmax, ymax))
	}

	log.Println("Part one:", infectioncount)

	// Part two
	infectioncount = 0
	world, ymax, xmax = parseWorld(lines)
	v = &virus{
		facing: north,
	}
	for i := 0; i < 10000000; i++ {
		if v.Tick(world, false) {
			infectioncount++
		}
		if v.x > xmax {
			xmax = v.x
		}
		if v.x*-1 > xmax {
			xmax = v.x * -1
		}
		if v.y > ymax {
			ymax = v.y
		}
		if v.y*-1 > ymax {
			ymax = v.y * -1
		}
		// fmt.Println(world.String(xmax, ymax))
	}

	log.Println("Part two:", infectioncount)
}

func parseWorld(lines []string) (grid, int, int) {
	world := make(grid)

	ymax := (len(lines) - 1) / 2
	xmax := (len(lines[0]) - 1) / 2

	for ly, line := range lines {
		y := ly - ymax
		for lx, r := range line {
			x := lx - xmax
			var s state
			switch r {
			case '#':
				s = infected
			default:
				s = clean
			}
			world.Set(x, y, s)
		}
	}

	return world, xmax, ymax
}

type virus struct {
	x, y   int
	facing dir
}

// Tick performs one burst of activity, returning whether or not we infected a node
func (v *virus) Tick(g grid, partone bool) bool {
	var current state
	defer func() {
		g.Set(v.x, v.y, current.Next(partone))
		v.Move()
	}()
	// If infected, turn right and become uninfected
	current = g.Get(v.x, v.y)
	switch current {
	case infected:
		v.facing = v.facing.Right()
		return false
	case clean:
		v.facing = v.facing.Left()
		return partone
	case flagged:
		v.facing = v.facing.Left().Left()
		return false
	case weakened:
		return true
	}
	panic("unknown state")
}

func (v *virus) Move() {
	switch v.facing {
	case north:
		v.y--
	case east:
		v.x++
	case south:
		v.y++
	case west:
		v.x--
	}
}

type grid map[int]map[int]state

func (g grid) Get(x, y int) state {
	if _, ok := g[x]; !ok {
		g[x] = make(map[int]state)
	}
	if v, ok := g[x][y]; ok {
		return v
	}
	return clean
}

func (g grid) Set(x, y int, val state) {
	if _, ok := g[x]; !ok {
		g[x] = make(map[int]state)
	}
	g[x][y] = val
}

func (g grid) String(xmax, ymax int) string {
	var out strings.Builder
	for y := ymax * -1; y <= ymax; y++ {
		for x := xmax * -1; x <= xmax; x++ {
			s := g.Get(x, y)
			switch s {
			case infected:
				out.WriteRune('#')
			case clean:
				out.WriteRune('.')
			case weakened:
				out.WriteRune('W')
			case flagged:
				out.WriteRune('F')
			}
		}
		out.WriteString("\n")
	}
	return out.String()
}

type dir int

const (
	north dir = iota
	east
	south
	west
)

func (d dir) Right() dir {
	d++
	if d > west {
		d = north
	}
	return d
}

func (d dir) Left() dir {
	d--
	if d < north {
		d = west
	}
	return d
}

type state int

const (
	clean state = iota
	weakened
	infected
	flagged
)

func (s state) Next(partone bool) state {
	if partone {
		if s == clean {
			return infected
		}
		return clean
	}

	s++
	if s > flagged {
		s = clean
	}
	return s
}
