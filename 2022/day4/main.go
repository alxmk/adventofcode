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

	var contains, overlaps int
	for _, line := range strings.Split(string(data), "\n") {
		var p pair
		fmt.Sscanf(line, "%d-%d,%d-%d", &p[0].a, &p[0].b, &p[1].a, &p[1].b)
		if p.Contains() {
			contains++
		}
		if p.Overlaps() {
			overlaps++
		}
	}
	log.Println("Part one:", contains)
	log.Println("Part two:", overlaps)
}

type pair [2]assignment

func (p pair) Contains() bool {
	return p[0].Contains(p[1]) || p[1].Contains(p[0])
}

func (p pair) Overlaps() bool {
	return p[0].Overlaps(p[1])
}

type assignment struct {
	a, b int
}

func (a assignment) Contains(b assignment) bool {
	return a.a <= b.a && a.b >= b.b
}

func (a assignment) Overlaps(b assignment) bool {
	return !(a.a > b.b || b.a > a.b)
}
