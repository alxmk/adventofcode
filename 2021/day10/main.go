package main

import (
	"io/ioutil"
	"log"
	"sort"
	"strings"
)

func main() {
	data, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatalln("Failed to read input:", err)
	}

	p1, p2 := solve(string(data))
	log.Println("Part one:", p1)
	log.Println("Part two:", p2)
}

func solve(input string) (int, int) {
	var p1 int
	var scores []int
	for _, line := range strings.Split(input, "\n") {
		syntaxError, incomplete := chunk(line)
		if syntaxError != 0 {
			p1 += scoreMap[syntaxError]
			continue
		}
		var score int
		for _, r := range incomplete {
			score = score*5 + scoreMap2[r]
		}
		scores = append(scores, score)
	}
	sort.Ints(scores)
	return p1, scores[len(scores)/2]
}

type stack []rune

func (s *stack) IsEmpty() bool {
	return len(*s) == 0
}

func (s *stack) Push(r rune) {
	*s = append(*s, r)
}

func (s *stack) Pop() (rune, bool) {
	if s.IsEmpty() {
		return 0, false
	}
	i := len(*s) - 1
	r := (*s)[i]
	*s = (*s)[:i]
	return r, true
}

func chunk(line string) (rune, string) {
	var s stack
	for _, r := range line {
		switch r {
		case '(', '{', '<', '[':
			s.Push(r)
		case ')', '}', '>', ']':
			q, ok := s.Pop()
			if !ok || r != closerMap[q] {
				return r, ""
			}
		}
	}
	if s.IsEmpty() {
		return 0, ""
	}
	var missing []rune
	for r, ok := s.Pop(); ok; r, ok = s.Pop() {
		missing = append(missing, closerMap[r])
	}
	return 0, string(missing)
}

var (
	closerMap = map[rune]rune{'(': ')', '{': '}', '[': ']', '<': '>'}

	scoreMap = map[rune]int{')': 3, '}': 1197, ']': 57, '>': 25137}

	scoreMap2 = map[rune]int{')': 1, '}': 3, ']': 2, '>': 4}
)
