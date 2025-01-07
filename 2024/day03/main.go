package main

import (
	_ "embed"
	"fmt"
	"regexp"
	"strconv"
)

//go:embed input.txt
var input string

func main() {
	r := regexp.MustCompile(`mul\((?P<a>\d+),(?P<b>\d+)\)|do\(\)|don't\(\)`)
	var sum, csum int
	conditional := true
	for _, match := range r.FindAllStringSubmatch(input, -1) {
		switch match[0] {
		case "do()":
			conditional = true
		case "don't()":
			conditional = false
		default:
			a, _ := strconv.Atoi(match[1])
			b, _ := strconv.Atoi(match[2])
			sum += a * b
			if conditional {
				csum += a * b
			}
		}
	}
	fmt.Println("Part one:", sum)
	fmt.Println("Part two:", csum)
}
