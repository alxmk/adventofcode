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

	p1, p2 := solve(parseSeqs(string(data)))

	log.Println("Part one:", p1)
	log.Println("Part two:", p2)
}

func solve(seqs []seq) (int, int) {
	var p1, p2 int
	for _, s := range seqs {
		p1 += s.Next()
		p2 += s.Prev()
	}
	return p1, p2
}

type seq []int

func (s seq) Extrapolate() (seq, int) {
	var sum int
	var diffs seq
	for i := 1; i < len(s); i++ {
		sum += s[i]
		diffs = append(diffs, s[i]-s[i-1])
	}
	return diffs, sum
}

func (s seq) Next() int {
	if diffs, sum := s.Extrapolate(); sum != 0 {
		return s[len(s)-1] + diffs.Next()
	}
	return 0
}

func (s seq) Prev() int {
	if diffs, sum := s.Extrapolate(); sum != 0 {
		return s[0] - diffs.Prev()
	}
	return 0
}

func parseSeqs(input string) []seq {
	var seqs []seq
	for _, line := range strings.Split(input, "\n") {
		var s seq
		for _, w := range strings.Fields(line) {
			v, _ := strconv.Atoi(w)
			s = append(s, v)
		}
		seqs = append(seqs, s)
	}
	return seqs
}
