package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"math/rand"
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

	d := &droid{brain: programme.Copy(), in: make(chan int64), out: make(chan int64), position: coordinate{0, 0}}
	w := d.Explore()

	fmt.Println(w)

	path, _ := w.Astar(coordinate{0, 0}, w.oxygenSystem())
	log.Println("Part one:", path-1)
	_, maxPath := w.Astar(w.oxygenSystem(), coordinate{0, 0})
	log.Println("Part two:", maxPath)
}

type droid struct {
	brain    intcode.Programme
	in, out  chan int64
	position coordinate
}

func (d *droid) Explore() *world {
	w := &world{tiles: map[coordinate]tile{coordinate{0, 0}: floor}}

	go func() {
		if err := d.brain.Run(d.in, d.out); err != nil {
			log.Fatalln("Droid brain failed:", err)
		}
	}()

	for i := 0; !w.Traversed(); i++ {
		dir := w.nextDir(d.position)
		d.Move(dir)
		next := d.position.Next(dir)

		rc := <-d.out
		switch returnCode(rc) {
		case hitWall:
			// log.Println("Hit wall", dir)
			w.tiles[next] = wall
		case moved:
			// log.Println("Moved", dir, "to", next)
			w.tiles[next] = floor
			d.position = next
		case found:
			// log.Println("Found O2", dir)
			w.tiles[next] = oxygenSystem
			d.position = next
		}
		if next.x > w.xmax {
			w.xmax = next.x
		}
		if next.x < w.xmin {
			w.xmin = next.x
		}
		if next.y > w.ymax {
			w.ymax = next.y
		}
		if next.y < w.ymin {
			w.ymin = next.y
		}
		// log.Println(d.position, next)
		// if i%100000 == 0 {
		// 	fmt.Println(w)
		// }
	}

	return w
}

type world struct {
	tiles                  map[coordinate]tile
	xmax, ymax, xmin, ymin int64
}

func (w *world) Traversed() bool {
	for c, t := range w.tiles {
		if t == wall {
			continue
		}
		for _, dir := range directions {
			if _, ok := w.tiles[c.Next(dir)]; !ok {
				return false
			}
		}
	}
	return true
}

func (w *world) nextDir(current coordinate) direction {
	// Prioritise immediately proximate unexplored tiles
	for _, d := range directions {
		if _, ok := w.tiles[current.Next(d)]; !ok {
			return d
		}
	}

	// Fall back to random
	return directions[rand.Intn(len(directions))]
}

func (w *world) oxygenSystem() coordinate {
	for c, t := range w.tiles {
		if t == oxygenSystem {
			return c
		}
	}
	return coordinate{0, 0}
}

func (w world) String() string {
	var out strings.Builder
	for y := w.ymax; y >= w.ymin; y-- {
		for x := w.xmin; x <= w.xmax; x++ {
			if t, ok := w.tiles[coordinate{x, y}]; ok {
				switch t {
				case wall:
					out.WriteRune('#')
				case floor:
					if y == 0 && x == 0 {
						out.WriteRune('X')
					} else {
						out.WriteRune('.')
					}
				case oxygenSystem:
					out.WriteRune('O')
				}
				continue
			}
			out.WriteRune(' ')
		}
		out.WriteRune('\n')
	}
	return out.String()
}

type coordinate struct {
	x, y int64
}

func (c coordinate) String() string {
	return fmt.Sprintf("(%d, %d)", c.x, c.y)
}

func (c coordinate) Next(dir direction) coordinate {
	return coordinate{c.x + dir.x, c.y + dir.y}
}

func (c coordinate) ManhattanDistance(to coordinate) int64 {
	return absint(to.x-c.x) + absint(to.y-c.y)
}

func absint(i int64) int64 {
	if i < 0 {
		return i * -1
	}
	return i
}

func (d *droid) Move(dir direction) {
	switch dir {
	case north:
		d.in <- 1
	case south:
		d.in <- 2
	case west:
		d.in <- 3
	case east:
		d.in <- 4
	default:
		panic("undefined direction")
	}
}

type direction coordinate

var (
	north = direction{0, 1}
	south = direction{0, -1}
	west  = direction{-1, 0}
	east  = direction{1, 0}

	directions = []direction{north, south, east, west}
)

func (d direction) String() string {
	switch d {
	case north:
		return "north"
	case south:
		return "south"
	case west:
		return "west"
	case east:
		return "east"
	}
	return ""
}

type tile int64

const (
	unknown tile = iota
	floor
	wall
	oxygenSystem
)

type returnCode int64

const (
	hitWall returnCode = iota
	moved
	found
)

func (w *world) Astar(from, to coordinate) (int64, int64) {
	log.Println("Pathing from", from, "to", to)
	openSet := map[coordinate]struct{}{
		from: struct{}{},
	}

	cameFrom := make(map[coordinate]coordinate)

	gScore := defaultMaxMap{
		from: 0,
	}

	fScore := defaultMaxMap{
		from: gScore[from] + from.ManhattanDistance(from),
	}

	var found bool

	for len(openSet) != 0 {
		lowestScore := int64(math.MaxInt64)
		var current coordinate

		for c := range openSet {
			if score := fScore.Get(c); score < lowestScore {
				lowestScore = score
				current = c
			}
		}

		// log.Println("Pathing", current)

		if current == to {
			log.Println("Path found")
			found = true
		}

		delete(openSet, current)
		for _, d := range directions {
			neighbour := current.Next(d)
			tentativeG := gScore.Get(current)
			if tentativeG != math.MaxInt64 {
				tentativeG++
			}
			if w.tiles[neighbour] == wall {
				tentativeG = math.MaxInt64
			}
			// log.Println("TentativeG for", neighbour, tentativeG, "current G", gScore.Get(neighbour))

			if tentativeG < gScore.Get(neighbour) {
				cameFrom[neighbour] = current
				gScore[neighbour] = tentativeG
				fScore[neighbour] = tentativeG + from.ManhattanDistance(neighbour)
				openSet[neighbour] = struct{}{}
			}
		}
	}

	if !found {
		return -1, -1
	}

	var maxG int64
	for _, g := range gScore {
		if g > maxG {
			maxG = g
		}
	}

	return reconstructPath(cameFrom, to), maxG
}

type defaultMaxMap map[coordinate]int64

func (d defaultMaxMap) Get(c coordinate) int64 {
	if v, ok := d[c]; ok {
		return v
	}
	return math.MaxInt64
}

func reconstructPath(cameFrom map[coordinate]coordinate, current coordinate) int64 {
	totalPath := []coordinate{current}
	// log.Println(current)
	for current, ok := cameFrom[current]; ok; current, ok = cameFrom[current] {
		// log.Println(current)
		totalPath = append(totalPath, current)
	}
	return int64(len(totalPath))
}
