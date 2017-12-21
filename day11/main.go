package main

import (
	"io/ioutil"
	"log"
	"math"
	"strings"
)

var (
	directions = []string{
		"n", "ne", "nw", "s", "se", "sw",
	}
)

func main() {
	input, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatalf("Error reading input file %v", err)
	}

	moves := strings.Split(string(input), ",")

	// g := grid{
	// 	tiles: []*tile{
	// 		newTile(),
	// 	},
	// }

	moveCount := make(map[string]int)

	for _, m := range moves {
		moveCount[m]++
	}

	consolidated := reduce(moveCount)

	var xCount int
	var yCount int

	for dir, count := range consolidated {
		if strings.Contains(dir, "n") {
			yCount += count
		}
		if strings.Contains(dir, "s") {
			yCount -= count
		}
		if strings.Contains(dir, "e") {
			xCount += count
		}
		if strings.Contains(dir, "w") {
			xCount -= count
		}
		log.Println(dir, count)
	}

	log.Println(xCount, yCount)

	xAbs := math.Abs(float64(xCount))
	yAbs := math.Abs(float64(yCount))

	var totalMoves float64

	if xAbs > yAbs {
		totalMoves = xAbs
	} else {
		totalMoves = yAbs
	}

	log.Println("The answer is", totalMoves)
}

// reduce cancels out moves in opposite directions to each other
func reduce(moves map[string]int) map[string]int {
	consolidated := make(map[string]int)

	directionsToCheck := []string{"n", "nw", "ne"}

	for _, d := range directionsToCheck {
		dirCount := moves[d]
		oppositeCount := moves[oppositeDirection(d)]

		if dirCount >= oppositeCount {
			consolidated[oppositeDirection(d)] = 0
			consolidated[d] = dirCount - oppositeCount
		} else {
			consolidated[d] = 0
			consolidated[oppositeDirection(d)] = oppositeCount - dirCount
		}
	}

	return consolidated
}

func oppositeDirection(direction string) string {
	switch direction {
	case "n":
		return "s"
	case "ne":
		return "sw"
	case "nw":
		return "se"
	case "s":
		return "n"
	case "se":
		return "nw"
	case "sw":
		return "ne"
	default:
		return "UNKNOWN"
	}
}
