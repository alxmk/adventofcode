package main

import (
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	data, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatalln("Failed to read input:", err)
	}

	signal, out := calculate(string(data))
	log.Println("Part one:", signal)
	log.Println("Part two:", "\n"+out)
}

func calculate(input string) (int, string) {
	c := &Clock{}
	x := 1
	for _, line := range strings.Split(input, "\n") {
		fields := strings.Fields(line)
		switch fields[0] {
		case "noop":
			c.Increment(x)
		case "addx":
			c.Increment(x).Increment(x)
			v, _ := strconv.Atoi(fields[1])
			x += v
		}
	}
	return c.signal, c.b.String()
}

type Clock struct {
	cycle, signal int
	b             strings.Builder
}

func (c *Clock) Increment(x int) *Clock {
	// Render
	pos := c.cycle % 40
	if pos == 0 {
		c.b.WriteRune('\n')
	}
	switch pos - x {
	case -1, 0, 1:
		c.b.WriteRune('#')
	default:
		c.b.WriteRune('.')
	}
	// Increment
	c.cycle++
	if c.cycle%40 == 20 {
		c.signal += x * c.cycle
	}
	return c
}
