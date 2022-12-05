package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	data, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatalln("Failed to read input:", err)
	}

	parts := strings.Split(string(data), "\n\n")

	one := parseInstructions(parseStacks(parts[0]), parts[1], nil)
	var output string
	for _, s := range one {
		r, _ := s.Pop()
		output += string(r)
	}
	log.Println("Part one:", output)

	two := parseInstructions(parseStacks(parts[0]), parts[1], stack{})
	output = ""
	for _, s := range two {
		r, _ := s.Pop()
		output += string(r)
	}
	log.Println("Part two:", output)
}

func parseInstructions(stacks [9]stack, input string, crane stack) [9]stack {
	for _, line := range strings.Split(input, "\n") {
		var q, f, t int
		fmt.Sscanf(line, "move %d from %d to %d", &q, &f, &t)
		for i := 0; i < q; i++ {
			r, _ := stacks[f-1].Pop()
			if crane != nil {
				crane.Push(r)
				continue
			}
			stacks[t-1].Push(r)
		}
		if crane != nil {
			for i := 0; i < q; i++ {
				r, _ := crane.Pop()
				stacks[t-1].Push(r)
			}
		}
	}
	return stacks
}

func parseStacks(input string) [9]stack {
	var stacks [9]stack
	lines := strings.Split(input, "\n")
	for i := len(lines) - 2; i >= 0; i-- {
		for j := 0; j < 9; j++ {
			if r := lines[i][1+(4*j)]; r >= 'A' && r <= 'Z' {
				stacks[j].Push(r)
			}
		}
	}
	return stacks
}

type stack []byte

func (s *stack) IsEmpty() bool {
	return len(*s) == 0
}

func (s *stack) Push(r byte) {
	*s = append(*s, r)
}

func (s *stack) Pop() (byte, bool) {
	if s.IsEmpty() {
		return 0, false
	}
	i := len(*s) - 1
	r := (*s)[i]
	*s = (*s)[:i]
	return r, true
}
