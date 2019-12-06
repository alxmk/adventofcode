package main

import (
	"io/ioutil"
	"log"
	"strings"
)

func main() {
	data, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatalln("Failed to read input file", err)
	}

	var nicePartOne, nicePartTwo int
	for _, line := range strings.Split(string(data), "\n") {
		if isNicePartOne(line) {
			nicePartOne++
		}
		if isNicePartTwo(line) {
			log.Println(line)
			nicePartTwo++
		}
	}

	log.Println("Part one:", nicePartOne)
	log.Println("Part two:", nicePartTwo)
}

func isNicePartOne(line string) bool {
	var vowels int
	var double bool

	var previous rune
	for _, r := range line {
		switch r {
		case 'a', 'e', 'i', 'o', 'u':
			vowels++
		case 'b':
			if previous == 'a' {
				return false
			}
		case 'd':
			if previous == 'c' {
				return false
			}
		case 'q':
			if previous == 'p' {
				return false
			}
		case 'y':
			if previous == 'x' {
				return false
			}
		}
		if r == previous {
			double = true
		}
		previous = r
	}

	return vowels >= 3 && double
}

func isNicePartTwo(line string) bool {
	var previous, previouser rune
	var oneLetterBack, pairsFound bool

	pairs := make(map[pair]int)
	for i, r := range line {
		if r == previouser {
			oneLetterBack = true
		}
		thisPair := pair{r, previous}
		if j, ok := pairs[thisPair]; ok {
			if j+1 != i {
				pairsFound = true
			}
		} else {
			pairs[thisPair] = i
		}

		previous, previouser = r, previous
	}
	return oneLetterBack && pairsFound
}

type pair struct {
	a, b rune
}
