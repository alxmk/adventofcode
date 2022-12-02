package main

import (
	"log"
	"os"
	"strings"
)

var (
	diff = byte('X' - 'A')
)

func main() {
	data, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatalln("Failed to read input:", err)
	}

	var one, two int
	for _, round := range strings.Split(string(data), "\n") {
		one += score(round[0], round[2])
		switch round[2] {
		case 'X':
			// lose
			two += score(round[0], round[0]+diff-1)
		case 'Y':
			// draw
			two += score(round[0], round[0]+diff)
		case 'Z':
			// win
			two += score(round[0], round[0]+diff+1)
		}
	}
	log.Println("Part one:", one)
	log.Println("Part two:", two)
}

func score(a, b byte) int {
	var score int
	switch b {
	case 'X', 'Z' + 1:
		score += 1
	case 'Y':
		score += 2
	case 'Z', 'X' - 1:
		score += 3
	}
	switch b - a {
	// draw
	case diff:
		score += 3
	// win
	case diff + 1, diff - 2:
		score += 6
		// otherwise lose
	}
	return score
}
