package main

import (
	"log"
	"os"
	"strings"
)

func main() {
	data, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatalln("Failed to read input:", err)
	}

	log.Println("Part one:", calculate(string(data), false))
	log.Println("Part two:", calculate(string(data), true))
}

func calculate(input string, partTwo bool) int {
	var sum int
	for _, line := range strings.Split(input, "\n") {
		var calibration int
		for i := range line {
			if val, ok := isDigit(line, i, partTwo); ok {
				calibration += val * 10
				break
			}
		}
		for i := len(line) - 1; i >= 0; i-- {
			if val, ok := isDigit(line, i, partTwo); ok {
				calibration += val
				break
			}
		}
		sum += calibration
	}
	return sum
}

var digits = map[string]int{
	"one":   1,
	"two":   2,
	"three": 3,
	"four":  4,
	"five":  5,
	"six":   6,
	"seven": 7,
	"eight": 8,
	"nine":  9,
}

func isDigit(line string, idx int, partTwo bool) (int, bool) {
	if line[idx] >= '0' && line[idx] <= '9' {
		return int(line[idx] - '0'), true
	}
	if !partTwo {
		return 0, false
	}
	for digit, val := range digits {
		if strings.HasPrefix(line[idx:], digit) {
			return val, true
		}
	}
	return 0, false
}
