package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
)

func main() {
	// Read input
	input, err := ioutil.ReadFile("input.txt")

	if err != nil {
		log.Fatalf("Error reading input file: %v", err)
	}

	// Get in string format
	digits := string(input)

	var prev rune
	var sum int

	// Check through the digits and add them up if they match
	for _, c := range digits {
		if c == prev {
			value, err := strconv.Atoi(string(c))
			if err != nil {
				log.Fatalf("Couldn't convert %v to int", c)
			}

			sum += value
		}

		prev = c
	}

	// Special case to check if the last one matches the first
	if digits[0] == digits[len(digits)-1] {
		value, err := strconv.Atoi(string(digits[0]))
		if err != nil {
			log.Fatalf("Couldn't convert %v to int", digits[0])
		}

		sum += value
	}

	fmt.Println("The answer is", sum)
}
