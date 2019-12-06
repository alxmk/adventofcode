package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"strconv"
	"strings"
)

func main() {
	data, err := ioutil.ReadFile("../input.txt")
	if err != nil {
		log.Fatalln("Failed to read input file:", err)
	}

	var wires []wire
	for _, line := range strings.Split(string(data), "\n") {
		w, err := parseWire(line)
		if err != nil {
			log.Fatalln("Failed to parse wire:", err)
		}
		wires = append(wires, w)
	}

	if len(wires) != 2 {
		log.Fatalln("Expected 2 wires but got", len(wires))
	}

	log.Println("Part one:", wires[0].ClosestIntersection(wires[1], manhattan()))
	log.Println("Part two:", wires[0].ClosestIntersection(wires[1], steps()))
}

type wire []segment

type segment struct {
	start, end coordinate
	stepOffset int
}

func (a segment) String() string {
	return fmt.Sprintf("%s -> %s", a.start, a.end)
}

type coordinate struct {
	x, y int
}

func (c coordinate) Distance() int {
	return absint(c.x) + absint(c.y)
}

func (c coordinate) String() string {
	return fmt.Sprintf("(%d, %d)", c.x, c.y)
}

func absint(i int) int {
	if i < 0 {
		return -1 * i
	}
	return i
}

func (a segment) Intersects(b segment) (bool, *coordinate) {
	// Potential intersections are limited
	c1 := coordinate{a.start.x, b.start.y} // Assume xa = c, yb = d
	if a.Contains(c1) && b.Contains(c1) {
		return true, &c1
	}
	c2 := coordinate{b.start.x, a.start.y} // Assume xb = c, ya = d
	if a.Contains(c2) && b.Contains(c2) {
		return true, &c2
	}
	return false, nil
}

func (a segment) Contains(c coordinate) bool {

	// Determine min/max values
	xmin := a.end.x
	xmax := a.start.x
	if a.start.x < a.end.x {
		xmin = a.start.x
		xmax = a.end.x
	}
	ymin := a.end.y
	ymax := a.start.y
	if a.start.y < a.end.y {
		ymin = a.start.y
		ymax = a.end.y
	}

	// Does the coordinate checked against fall within the min/max
	return xmin <= c.x && xmax >= c.x && ymin <= c.y && ymax >= c.y
}

func (a segment) Steps(c coordinate) int {
	return a.stepOffset + absint(c.x-a.start.x) + absint(c.y-a.start.y)
}

func (a wire) ClosestIntersection(b wire, distance distFunc) int {
	closest := math.MaxInt16
	for _, sa := range a {
		for _, sb := range b {
			if ok, at := sa.Intersects(sb); ok {
				if d := distance(*at, sa, sb); d != 0 && d < closest {
					closest = d
				}
			}
		}
	}
	return closest
}

type distFunc func(coordinate, segment, segment) int

func manhattan() distFunc {
	return func(c coordinate, sa, sb segment) int { return c.Distance() }
}

func steps() distFunc {
	return func(c coordinate, sa, sb segment) int { return sa.Steps(c) + sb.Steps(c) }
}

func parseWire(line string) (wire, error) {
	var w wire
	var x, y int
	var xmul, ymul int
	var steps int
	for _, instruction := range strings.Split(line, ",") {
		switch instruction[0] {
		case 'R':
			xmul, ymul = 1, 0
		case 'L':
			xmul, ymul = -1, 0
		case 'U':
			xmul, ymul = 0, 1
		case 'D':
			xmul, ymul = 0, -1
		default:
			return nil, fmt.Errorf("bad instruction %s", instruction)
		}
		distance, err := strconv.Atoi(instruction[1:])
		if err != nil {
			return nil, fmt.Errorf("bad instruction %s: %v", instruction, err)
		}
		w = append(w, segment{
			start:      coordinate{x, y},
			end:        coordinate{x + (xmul * distance), y + (ymul * distance)},
			stepOffset: steps,
		})
		steps += distance
		x, y = x+(xmul*distance), y+(ymul*distance)
	}
	return w, nil
}
