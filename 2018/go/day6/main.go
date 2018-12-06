package main

import (
	"io/ioutil"
	"log"
	"math"
	"strconv"
	"strings"
)

type node struct {
	closest     rune
	point       rune
	distanceSum int
}

type coordinate struct {
	x    int
	y    int
	name rune
}

var maxDist int

func main() {
	data, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatalln("Failed to read input file", err)
	}

	// Coordinates as described in the input file
	coords := []coordinate{}

	// Track size that the grid needs to be
	var xMax, yMax int

	point := 'a'
	for _, line := range strings.Split(string(data), "\n") {
		x, y := parsecoords(line)

		coords = append(coords, coordinate{name: point, x: x, y: y})
		point++

		if x > xMax {
			xMax = x
		}
		if y > yMax {
			yMax = y
		}
	}

	// Maximum distance anything can be from anything is half the size of the perimeter
	maxDist = xMax + yMax

	// Create the grid
	grid := make([][]*node, xMax+1)
	for i := range grid {
		grid[i] = make([]*node, yMax+1)
	}

	// Put the coordinates onto the grid
	for _, c := range coords {
		grid[c.x][c.y] = &node{point: c.name}
	}

	// Track the size for part 2
	var size int

	// For each node on the grid, find the closest coordinate and how far it is from
	// all coordinates
	for x, row := range grid {
		for y := range row {
			if row[y] == nil {
				row[y] = &node{}
			}
			row[y].closest, row[y].distanceSum = findClosest(x, y, coords)

			if row[y].distanceSum < 10000 {
				size++
			}
		}
	}

	// Find the areas closest per coordinate name
	areas := make(map[rune]int)

	// Track the unbounded coordinates (i.e. those with nodes closest on the boundaries)
	infinite := []rune{}

	for x, row := range grid {
		for y, node := range row {
			// Skip contested nodes
			if node.closest == -1 {
				continue
			}
			areas[node.closest]++

			// If we're on an edge then this is an unbounded area
			if x == xMax || x == 0 || y == yMax || y == 0 {
				infinite = append(infinite, node.closest)
			}
		}
	}

	var largestSize int
	var largestCoord rune

	for point, size := range areas {
		if strings.ContainsRune(string(infinite), point) {
			continue
		}
		if size > largestSize {
			largestSize = size
			largestCoord = point
		}
	}

	log.Println(string(largestCoord), largestSize, size)
}

// parse cartesian coords from line
func parsecoords(line string) (int, int) {
	parts := strings.Split(line, ", ")
	x, _ := strconv.Atoi(parts[0])
	y, _ := strconv.Atoi(parts[1])

	return x, y
}

func findClosest(x, y int, coords []coordinate) (rune, int) {
	// Set the initial min distance to the maximum possible
	mindist := maxDist

	// Track the points which are closest
	closestPoints := make(map[rune]struct{})

	// Track the sum of all distances
	var distSum int
	for _, c := range coords {
		dist := int(math.Abs(float64(x-c.x)) + math.Abs(float64(y-c.y)))
		switch {
		case dist < mindist:
			closestPoints = map[rune]struct{}{c.name: struct{}{}}
			mindist = dist
		case dist == mindist:
			closestPoints[c.name] = struct{}{}
		}
		distSum += dist
	}

	// If there's only one closest point, return that, otherwise return -1 to indicate a contested coordinate
	if len(closestPoints) == 1 {
		for k := range closestPoints {
			log.Println(x, y, k)
			return k, distSum
		}
	}

	return -1, distSum
}
