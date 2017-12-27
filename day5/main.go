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

	log.Println("Part 1 answer is", process(maze, 1))

	log.Println("Part 2 answer is", process(maze, 2))
}

func process(maze []int, part int) int {
	index := 0
	numInstructions := 0

	mymaze := make([]int, len(maze))
	copy(mymaze, maze)

	for {
		numInstructions++

		current := mymaze[index]
		if current >= 3 && part == 2 {
			mymaze[index]--
		} else {
			mymaze[index]++
		}

		index += current

		if index < 0 || index >= len(mymaze) {
			break
		}
	}

	return numInstructions
}
