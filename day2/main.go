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
	var divsum int

	for _, r := range rows {
		cells := strings.Fields(r)

		var max int
		min := math.MaxInt16

		for i, c := range cells {
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

			for j, cell := range cells {
				if i == j {
					continue
				}

				value2, err := strconv.Atoi(cell)
				if err != nil {
					log.Fatalf("Couldn't parse cell %s as int", cell)
				}

				if math.Mod(float64(value), float64(value2)) == 0 {
					divsum += value / value2
					break
				}
			}
		}

		checksum += max - min
	}

	log.Println("Part 1 answer is", checksum)
	log.Println("Part 2 answer is", divsum)
}
