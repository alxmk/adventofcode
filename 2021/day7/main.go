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

	positions, err := parse(string(data))
	if err != nil {
		log.Fatalln("Failed to parse input:", err)
	}

	log.Println("Part one:", solve(positions, partOne))
	log.Println("Part two:", solve(positions, partTwo))
}

func parse(input string) ([]int, error) {
	var positions []int
	for _, raw := range strings.Split(input, ",") {
		v, err := strconv.Atoi(raw)
		if err != nil {
			return nil, fmt.Errorf("failed to parse %s: %s", raw, err)
		}
		positions = append(positions, v)
	}
	return positions, nil
}

func solve(positions []int, solver func([]int, int) int) int {
	min, max := math.MaxInt, 0
	for _, p := range positions {
		if p < min {
			min = p
		}
		if p > max {
			max = p
		}
	}
	best := math.MaxInt
	for i := min; i <= max; i++ {
		if n := solver(positions, i); n < best {
			best = n
		}
	}

	return best
}

func partOne(positions []int, position int) int {
	var fuel int
	for _, p := range positions {
		thisFuel := p - position
		if thisFuel < 0 {
			thisFuel *= -1
		}
		fuel += thisFuel
	}
	return fuel
}

func partTwo(positions []int, position int) int {
	var fuel int
	for _, p := range positions {
		n := float64(p - position)
		if n < 0 {
			n *= -1
		}
		fuel += int((n / 2) * (2 + (n - 1)))
	}
	return fuel
}
