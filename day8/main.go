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
		log.Fatalf("Error reading input file %v", err)
	}

	instructions := strings.Split(string(input), "\n")

	registers := make(map[string]int)

	for _, i := range instructions {
		register, increment, check, err := parse(i)
		if err != nil {
			log.Fatalf("Error parsing instruction: %v", err)
		}

		if !check.Pass(registers) {
			continue
		}

		if _, ok := registers[register]; ok {
			registers[register] += increment
		} else {
			registers[register] = increment
		}
	}

	var max int

	for _, val := range registers {
		if val > max {
			max = val
		}
	}

	log.Println("The answer is", max)
}

func parse(instruction string) (string, int, *check, error) {
	parts := strings.Fields(instruction)

	if len(parts) != 7 {
		return "", 0, nil, fmt.Errorf("Malformed instruction %s", instruction)
	}

	increment, err := strconv.Atoi(parts[2])
	if err != nil {
		return "", 0, nil, fmt.Errorf("Couldn't parse %s to int in instruction %s", parts[2], instruction)
	}

	switch parts[1] {
	case "inc":
	case "dec":
		increment = increment * -1
	default:
		return "", 0, nil, fmt.Errorf("Unexpected %s in instruction %s", parts[1], instruction)
	}

	checkValue, err := strconv.Atoi(parts[6])
	if err != nil {
		return "", 0, nil, fmt.Errorf("Couldn't parse %s to int in instruction %s", parts[6], instruction)
	}

	return parts[0], increment, &check{
		register:  parts[4],
		value:     checkValue,
		operation: parts[5],
	}, nil
}

type check struct {
	register  string
	value     int
	operation string
}

func (c *check) Pass(registers map[string]int) bool {
	switch c.operation {
	case ">":
		return registers[c.register] > c.value
	case "<":
		return registers[c.register] < c.value
	case ">=":
		return registers[c.register] >= c.value
	case "<=":
		return registers[c.register] <= c.value
	case "==":
		return registers[c.register] == c.value
	case "!=":
		return registers[c.register] != c.value
	default:
		return false
	}
}
