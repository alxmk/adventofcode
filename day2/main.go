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
		log.Fatalf("Failed to read input file: %v", err)
	}

	rows := strings.Split(string(input), "\n")

	var checksum int

	for _, r := range rows {
		cells := strings.Fields(r)

		var max int
		min := math.MaxInt16

		for _, c := range cells {
			value, err := strconv.Atoi(c)
			if err != nil {
				log.Fatalf("Couldn't parse cell %s as int", c)
			}

			if value < min {
				min = value
			}

			if value > max {
				max = value
			}
		}

		checksum += max - min
	}

	log.Println("Answer is", checksum)
}
