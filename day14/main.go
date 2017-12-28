package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.ibm.com/alexmk/adventofcode/day10/knot"
)

func main() {
	input := "ugkiagan"

	usedCount := 0
	regions := make(map[int]map[int]bool)

	for i := 0; i < 128; i++ {
		key := fmt.Sprintf("%s-%d", input, i)

		hash := knot.KnotHash(key)

		binary := parseToBinary(hash)

		usedCount += strings.Count(binary, "1")

		regions[i] = make(map[int]bool)

		for index, c := range binary {
			regions[i][index] = (c == '1')
		}
	}

	fmt.Println("Part 1 answer is", usedCount)

	visited := make(map[int]map[int]struct{})

	numRegions := 0

	for row, data := range regions {
		visited[row] = make(map[int]struct{})
		for index, value := range data {
			// If it's a zero just add it to the visited and carry on
			if !value {
				visited[row][index] = struct{}{}
				continue
			}
			if _, ok := visited[row][index]; !ok {
				// If we've not been here before it's a new region
				numRegions++
			}

			// Check to the right and below, unless we're at the far right/bottom
			if index != len(data) {
				// If it's a 1 then visit it
				if _, value := data[index+1]; value {
					visited[row][index+1] = struct{}{}
				}
			}
			if row != len(regions) {
				// If it's a 1 then visit it
				if _, value := regions[row+1][index]; value {
					visited[row+1][index] = struct{}{}
				}
			}
		}
	}
}

func parseToBinary(hash string) string {
	output := ""
	for i := 0; i < len(hash); i += 2 {
		chunk := hash[i : i+2]
		value, _ := strconv.ParseInt(chunk, 16, 32)
		output += strconv.FormatInt(value, 2)
	}

	return output
}
