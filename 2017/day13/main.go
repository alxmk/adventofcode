package main

import (
	"io/ioutil"
	"log"
	"math"
	"strconv"
	"strings"
)

func main() {
	input, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatalf("Error reading input file %v", err)
	}

	lines := strings.Split(string(input), "\n")
	layers := make(map[int]int)

	maxDepth := 0

	for _, l := range lines {
		parts := strings.Split(l, ": ")

		depth, err := strconv.Atoi(parts[0])
		if err != nil {
			log.Fatalf("Couldn't convert %s to int", parts[0])
		}

		scanRange, err := strconv.Atoi(parts[1])
		if err != nil {
			log.Fatalf("Couldn't convert %s to int", parts[1])
		}

		layers[depth] = scanRange

		if depth > maxDepth {
			maxDepth = depth
		}
	}

	score := 0

	for depth, scanrange := range layers {
		// If there's no firewall at this level just continue
		if math.Mod(float64(depth), float64(2*scanrange-2)) == 0 {
			score += depth * scanrange
		}
	}

	log.Println("Part one answer is", score)

	delay := 0
	deepest := 0

nextDelay:
	for {
		delay++

		for depth, scanrange := range layers {
			if depth > deepest {
				deepest = depth
			}

			// If we ever hit a scanner try next delay value
			if math.Mod(float64(depth+delay), float64(2*scanrange-2)) == 0 {
				log.Println("Delay", delay, "hit scanner at depth", depth, "max depth", deepest)
				continue nextDelay
			}
		}

		// Success!
		break
	}

	log.Println("Part two answer is", delay)
}
