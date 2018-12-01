package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func main() {
	input, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatalf("Error opening input file %v", err)
	}

	sizes := strings.Fields(string(input))

	var banks []int

	for _, s := range sizes {
		value, err := strconv.Atoi(s)
		if err != nil {
			log.Fatalf("Couldn't convert size %s to int", s)
		}

		banks = append(banks, value)
	}

	seen := make(map[string]int)

	var cycles int
	var loopSize int

	for {
		stringRepresentation := fmt.Sprint(banks)
		if c, ok := seen[stringRepresentation]; ok {
			loopSize = cycles - c
			break
		}

		seen[stringRepresentation] = cycles

		banks = redistribute(banks)
		cycles++
	}

	log.Println("Part 1 answer is", cycles)
	log.Println("Part 2 answer is", loopSize)
}

func redistribute(banks []int) []int {
	var max, bank int
	for b, size := range banks {
		if size > max {
			max = size
			bank = b
		}
	}

	banks[bank] = 0
	index := bank

	for i := 0; i < max; i++ {
		index++

		if index > len(banks)-1 {
			index = 0
		}

		banks[index]++
	}

	return banks
}
