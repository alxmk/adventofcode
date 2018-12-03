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
		log.Fatalln("Failed to read input file:", err)
	}

	var twoCount int
	var threeCount int

	lines := strings.Split(string(data), "\n")
	sort.Sort(sort.StringSlice(lines))

	for _, line := range lines {
		counts := letterCount(line, 2, 3)
		if counts[2] {
			twoCount++
		}
		if counts[3] {
			threeCount++
		}
	}

	for i, line := range lines {
		if match, ok := findOneCharDiff(line, lines[i:]); ok {
			log.Println(line, match)
		}
	}

	log.Println(twoCount * threeCount)
}

func letterCount(line string, nums ...int) map[int]bool {
	letterMap := make(map[rune]int)
	for _, r := range line {
		letterMap[r]++
	}

	outMap := make(map[int]bool)
	for _, count := range letterMap {
		for _, n := range nums {
			if count == n {
				outMap[n] = true
			}
		}
	}

	return outMap
}

func findOneCharDiff(line string, lines []string) (string, bool) {
	for _, candidate := range lines {
		var diff int
		for i, r := range []byte(line) {
			if r != candidate[i] {
				diff++
			}
			if diff == 2 {
				break
			}
		}
		if diff == 1 {
			return candidate, true
		}
	}
	return "", false
}
