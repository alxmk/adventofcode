package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
	"text/scanner"
)

var garbageCount int

func main() {
	input, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatalf("Error reading input file %v", err)
	}

	scan := &scanner.Scanner{}

	scan.Init(strings.NewReader(string(input)))
	// Read off the first {
	if scan.Scan() != '{' {
		log.Fatalf("Error - input didn't start with {")
	}

	group := parse(scan, 1)

	log.Println("Part one answer is", group.Score())
	log.Println("Part two answer is", garbageCount)
}

func parse(scan *scanner.Scanner, score int) *group {
	group := &group{
		score: score,
	}

	garbage := false

	for {
		c := scan.Next()

		group.me += string(c)

		switch {
		case c == '<' && !garbage:
			garbage = true
		case c == '>':
			garbage = false
		case c == '{' && !garbage:
			group.children = append(group.children, parse(scan, score+1))
		case c == '}' && !garbage:
			return group
		case c == '!':
			// Toss the next one and continue
			scan.Next()
		case c == scanner.EOF:
			return group
		default:
			// If it's a garbage character increase the count
			if garbage {
				garbageCount++
			}
		}
	}
}

type group struct {
	children []*group
	score    int
	me       string
}

func (g *group) Score() int {
	total := g.score

	for _, c := range g.children {
		total += c.Score()
	}

	return total
}

func (g *group) Print() string {
	out := fmt.Sprintf("%s I have score %d and %d children.\n", g.me, g.score, len(g.children))

	for _, c := range g.children {
		out += c.Print()
	}

	return out
}
