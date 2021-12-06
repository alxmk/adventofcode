package main

import (
	"fmt"
	"io/ioutil"
	"log"
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
		log.Fatalln("Failed part one:", err)
	}
	log.Println("Part one:", p1)

	p2, err := partTwo(lines)
	if err != nil {
		log.Fatalln("Failed part two:", err)
	}
	log.Println("Part two:", p2)
}

func partOne(lines []string) (int, error) {
	var x, y int
	for _, line := range lines {
		parts := strings.Fields(line)
		if len(parts) != 2 {
			return -1, fmt.Errorf("malformed line: %s", line)
		}
		v, err := strconv.Atoi(parts[1])
		if err != nil {
			return -1, fmt.Errorf("malformed magnitude %s: %s", parts[1], err)
		}
		switch parts[0] {
		case "forward":
			x += v
		case "down":
			y += v
		case "up":
			y -= v
		default:
			return -1, fmt.Errorf("unknown directive %s", parts[0])
		}
	}
	return x * y, nil
}

func partTwo(lines []string) (int, error) {
	var x, y, aim int
	for _, line := range lines {
		parts := strings.Fields(line)
		if len(parts) != 2 {
			return -1, fmt.Errorf("malformed line: %s", line)
		}
		v, err := strconv.Atoi(parts[1])
		if err != nil {
			return -1, fmt.Errorf("malformed magnitude %s: %s", parts[1], err)
		}
		switch parts[0] {
		case "forward":
			x, y = x+v, y+(v*aim)
		case "down":
			aim += v
		case "up":
			aim -= v
		default:
			return -1, fmt.Errorf("unknown directive %s", parts[0])
		}
	}
	return x * y, nil
}
