package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"strings"
)

func main() {
	data, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatalln("Failed to read input:", err)
	}

	polymer, rules, err := parse(string(data))
	if err != nil {
		log.Fatalln("Failed to parse input:", err)
	}

	log.Println("Part one:", evaluate(polymer, rules, 10))
	log.Println("Part two:", evaluate(polymer, rules, 40))
}

func parse(input string) (string, map[string][]string, error) {
	parts := strings.Split(input, "\n\n")
	if l := len(parts); l != 2 {
		return "", nil, fmt.Errorf("malformed input, expected 2 sections got %d", l)
	}
	rules := make(map[string][]string)
	for _, line := range strings.Split(parts[1], "\n") {
		var a, b string
		if _, err := fmt.Sscanf(line, "%s -> %s", &a, &b); err != nil {
			return "", nil, fmt.Errorf("malformed line %s: %s", line, err)
		}
		rules[a] = []string{string(a[0]) + b, b + string(a[1])}
	}
	return parts[0], rules, nil
}

type pairs struct {
	all        map[string]int64
	head, tail string
}

func newPairs() *pairs {
	return &pairs{all: make(map[string]int64)}
}

func stringToPairs(input string) *pairs {
	p := newPairs()
	for i := 0; i < len(input)-1; i++ {
		p.all[input[i:i+2]]++
	}
	p.head, p.tail = input[0:2], input[len(input)-2:]
	return p
}

func step(p *pairs, rules map[string][]string) *pairs {
	newPairs := newPairs()
	for q, c := range p.all {
		for _, r := range rules[q] {
			newPairs.all[r] += c
		}
	}
	newPairs.head, newPairs.tail = rules[p.head][0], rules[p.tail][1]
	return newPairs
}

func evaluate(input string, rules map[string][]string, steps int) int64 {
	p := stringToPairs(input)
	for i := 0; i < steps; i++ {
		p = step(p, rules)
	}
	return diff(p)
}

func diff(p *pairs) int64 {
	counts := make(map[rune]int64)
	for q, c := range p.all {
		if q == p.head {
			counts[rune(q[0])] += 1
		}
		if q == p.tail {
			counts[rune(q[1])] += 1
		}
		for _, r := range q {
			counts[r] += c
		}
	}
	min, max := int64(math.MaxInt64), int64(0)
	for _, c := range counts {
		if c > max {
			max = c
		}
		if c < min {
			min = c
		}
	}
	return max/2 - min/2
}
