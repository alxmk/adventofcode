package main

import (
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func main() {
	input, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatalf("Error reading input file %v", err)
	}

	instructions := strings.Split(string(input), "\n")

	var maze []int

	for _, instruction := range instructions {
		value, err := strconv.Atoi(instruction)
		if err != nil {
			log.Fatalf("Couldn't convert instruction %s to int", instruction)
		}

		maze = append(maze, value)
	}

	index := 0
	numInstructions := 0

	for {
		numInstructions++

		current := maze[index]
		maze[index]++

		index += current

		if index < 0 || index >= len(maze) {
			break
		}
	}

	log.Println("The answer is", numInstructions)
}
