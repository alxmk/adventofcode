package main

import (
	"fmt"
	"math"
)

func main() {
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

	fmt.Println("The answer is", distance)
}
