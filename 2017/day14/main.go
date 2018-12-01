package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.ibm.com/alexmk/adventofcode/day10/knot"
)

var visitCount int

func main() {
	input := "ugkiagan"

	usedCount := 0
	regions := make(map[int]map[int]bool)

	for i := 0; i < 128; i++ {
		key := fmt.Sprintf("%s-%d", input, i)

		hash := knot.KnotHash(key)

		binary := parseToBinary(hash)

		if len(binary) != 128 {
			log.Fatalln("Binary was not 128 bits", binary, len(binary))
		}

		usedCount += strings.Count(binary, "1")

		regions[i] = make(map[int]bool)

		for index, c := range binary {
			regions[i][index] = (c == '1')
		}
	}

	fmt.Println("Part 1 answer is", usedCount)

	visited := make(map[int]map[int]struct{})

	numRegions := 0

	// Initialise everything
	for row := range regions {
		visited[row] = make(map[int]struct{})
	}

	for row := 0; row < len(regions); row++ {
		log.Println("Starting row", row, "regions", numRegions, "visit count", visitCount)
		for index := 0; index < len(regions[row]); index++ {
			// Skip zeroes and cells we've already visited
			if _, ok := visited[row][index]; !regions[row][index] || ok {
				continue
			}

			// If we get here we're at a new region
			numRegions++

			// Mark as visited
			visited[row][index] = struct{}{}
			visitCount++

			// Explore up, down, left, right and mark as visited if appropriate
			regions, visited = explore(row-1, index, regions, visited)
			regions, visited = explore(row+1, index, regions, visited)
			regions, visited = explore(row, index-1, regions, visited)
			regions, visited = explore(row, index+1, regions, visited)
		}
	}

	log.Println("Part two answer is", numRegions)
}

func explore(row, index int, regions map[int]map[int]bool, visited map[int]map[int]struct{}) (map[int]map[int]bool, map[int]map[int]struct{}) {
	// If we have invalid coordinates drop out
	if row < 0 || row >= len(regions) || index < 0 || index >= len(regions[row]) {
		return regions, visited
	}

	// If we already visited or it's a 0 drop out
	if _, ok := visited[row][index]; !regions[row][index] || ok {
		return regions, visited
	}

	// If we got here we haven't visited and it's a 1, so mark as visited
	visited[row][index] = struct{}{}
	visitCount++

	// Explore up, down, left, right and mark as visited if appropriate
	regions, visited = explore(row-1, index, regions, visited)
	regions, visited = explore(row+1, index, regions, visited)
	regions, visited = explore(row, index-1, regions, visited)
	return explore(row, index+1, regions, visited)
}

func parseToBinary(hash string) string {
	output := ""
	for i := 0; i < len(hash); i += 2 {
		chunk := hash[i : i+2]
		value, _ := strconv.ParseInt(chunk, 16, 32)
		bin := strconv.FormatInt(value, 2)
		for {
			if len(bin) == 8 {
				break
			}
			bin = "0" + bin
		}
		output += bin
	}

	return output
}
