package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"strconv"
	"strings"
)

func main() {
	data, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatalln("Failed to read input:", err)
	}

	lines := strings.Split(string(data), "\n")

	p1, err := partOne(lines)
	if err != nil {
		log.Fatalln("Part one failed:", err)
	}
	log.Println("Part one:", p1)

	p2, err := partTwo(lines)
	if err != nil {
		log.Fatalln("Part two failed:", err)
	}
	log.Println("Part two:", p2)
}

func partOne(lines []string) (int, error) {
	var count int
	previous := math.MaxInt
	for _, line := range lines {
		v, err := strconv.Atoi(line)
		if err != nil {
			return -1, fmt.Errorf("malformed line %s: %s", line, err)
		}
		if v > previous {
			count++
		}
		previous = v
	}
	return count, nil
}

func partTwo(lines []string) (int, error) {
	var count int
	previous := math.MaxInt
	for i := 0; i < len(lines)-2; i++ {
		var sum int
		for j := i; j < i+3; j++ {
			v, err := strconv.Atoi(lines[j])
			if err != nil {
				return -1, fmt.Errorf("malformed line %s: %s", lines[j], err)
			}
			sum += v
		}
		if sum > previous {
			count++
		}
		previous = sum
	}
	return count, nil
}
