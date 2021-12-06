package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func main() {
	data, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatalln("Failed to read input:", err)
	}

	fish, err := parse(string(data))
	if err != nil {
		log.Fatalln("Failed to parse input:", err)
	}

	log.Println("Part one:", naive(fish, 80))
	log.Println("Part two:", optimised(fish, 256))
}

func parse(input string) ([]int, error) {
	var fish []int
	for _, raw := range strings.Split(input, ",") {
		v, err := strconv.Atoi(raw)
		if err != nil {
			return nil, fmt.Errorf("failed to parse fish %s: %s", raw, err)
		}
		fish = append(fish, v)
	}
	return fish, nil
}

func naive(fish []int, days int) int {
	fishCopy := make([]int, len(fish))
	copy(fishCopy, fish)
	for day := 0; day < days; day++ {
		var newCount int
		for i := range fishCopy {
			if fishCopy[i] == 0 {
				fishCopy[i] = 6
				newCount++
				continue
			}
			fishCopy[i]--
		}
		for i := 0; i < newCount; i++ {
			fishCopy = append(fishCopy, 8)
		}
	}
	return len(fishCopy)
}

func optimised(fish []int, days int) int {
	// Initialise
	counts := make([]int, 9)
	for i := range fish {
		counts[fish[i]]++
	}

	for day := 0; day < days; day++ {
		// Track the number of both new fish and resetting fish
		newCount := counts[0]
		// Tick everything else down one
		for i := 1; i < 9; i++ {
			counts[i-1] = counts[i]
		}
		// Add the newly spawned fish at count 8 and reset the spawning fish at 6
		counts[8] = newCount
		counts[6] += newCount
	}

	var finalCount int
	for i := range counts {
		finalCount += counts[i]
	}
	return finalCount
}
