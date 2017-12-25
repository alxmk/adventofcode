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

	fullLayers := make(map[int]int)

	for i := 0; i <= maxDepth; i++ {
		fullLayers[i] = layers[i]
	}

	score := 0

	for picosecond := 0; picosecond <= maxDepth; picosecond++ {
		// If there's no firewall at this level just continue
		if math.Mod(float64(picosecond), float64(2*fullLayers[picosecond]-2)) == 0 {
			score += picosecond * fullLayers[picosecond]
		}
	}

	log.Println("The answer is", score)
}
