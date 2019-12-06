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
	data, err := ioutil.ReadFile("input.txt")
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

	log.Println("Part one:", closestIntersection(wires[0], wires[1], manhattan()))
	log.Println("Part two:", closestIntersection(wires[0], wires[1], steps()))
}

type wire map[coordinate]int

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

func parseWire(line string) (wire, error) {
	w := make(wire)
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
		for i := 0; i < distance; i++ {
			steps++
			x += xmul
			y += ymul
			if _, ok := w[coordinate{x, y}]; !ok {
				w[coordinate{x, y}] = steps
			}
		}
	}
	delete(w, coordinate{0, 0})
	return w, nil
}

func closestIntersection(a, b wire, distance distFunc) int {
	minDistance := math.MaxInt16
	for k, stepsA := range a {
		if stepsB, ok := b[k]; ok {
			if d := distance(k, stepsA, stepsB); d < minDistance {
				minDistance = d
			}
		}
	}
	return minDistance
}

type distFunc func(coordinate, int, int) int

func manhattan() distFunc {
	return func(c coordinate, stepsA, stepsB int) int { return c.Distance() }
}

func steps() distFunc {
	return func(c coordinate, stepsA, stepsB int) int { return stepsA + stepsB }
}
