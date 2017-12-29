package main

import (
	"io/ioutil"
	"log"
	"strings"
)

func main() {
	input, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatalf("Error reading input file %v", err)
	}

	lines := strings.Split(string(input), "\n")

	myMap := &Map{
		tiles: make(map[int]map[int]rune),
	}

	for i, line := range lines {
		myMap.tiles[i] = make(map[int]rune)
		for j, c := range []rune(line) {
			myMap.tiles[i][j] = c
		}
	}

	_, _, numSteps, letters := myMap.Navigate()

	log.Println("Part one answer is", string(letters))

	log.Println("Part two answer is", numSteps)
}

type Map struct {
	tiles map[int]map[int]rune
}

func (m *Map) Navigate() (int, int, int, []rune) {
	// Find the entry point
	var currentX, currentY int
	for i, c := range m.tiles[currentY] {
		if c != ' ' {
			currentX = i
			break
		}
	}

	// currentDirection is the current direction of travel - we always maintain
	// it if possible. 0 North, 1 East, 2 South, 3 West.
	currentDirection := 2

	// Next tile we're looking at
	nextX, nextY := currentX, currentY

	// Track letters acquired
	var letters []rune

	// Track the number of steps taken
	var numSteps int

	for {
		numSteps++

		// Try straight on
		nextX, nextY = getNext(currentX, currentY, currentDirection)
		if letter, ok := m.checkNext(nextX, nextY); ok {
			if letter != ' ' {
				letters = append(letters, letter)
			}
			currentX, currentY = nextX, nextY
			continue
		}

		// Try left
		nextX, nextY = getNext(currentX, currentY, getLeft(currentDirection))
		if letter, ok := m.checkNext(nextX, nextY); ok {
			if letter != ' ' {
				letters = append(letters, letter)
			}
			currentX, currentY, currentDirection = nextX, nextY, getLeft(currentDirection)
			continue
		}

		// Try right
		nextX, nextY = getNext(currentX, currentY, getRight(currentDirection))
		if letter, ok := m.checkNext(nextX, nextY); ok {
			if letter != ' ' {
				letters = append(letters, letter)
			}
			currentX, currentY, currentDirection = nextX, nextY, getRight(currentDirection)
			continue
		}

		// We're at the end (hopefully!)
		break
	}

	return currentX, currentY, numSteps, letters
}

func (m *Map) checkNext(nextX, nextY int) (rune, bool) {
	if c, ok := m.tiles[nextY][nextX]; ok && c != ' ' {
		// If it's none of the standard characters we've hit a letter
		if c != '|' && c != '-' && c != '+' {
			return c, true
		}
		return ' ', true
	}
	return ' ', false
}

func getNext(nextX, nextY, direction int) (int, int) {
	switch direction {
	case 0:
		nextY--
	case 1:
		nextX++
	case 2:
		nextY++
	case 3:
		nextX--
	}

	return nextX, nextY
}

func getLeft(direction int) int {
	if left := direction - 1; left >= 0 {
		return left
	}

	return 3
}

func getRight(direction int) int {
	if right := direction + 1; right <= 3 {
		return right
	}

	return 0
}
