package main

import (
	"io/ioutil"
	"log"
	"sort"
	"strings"
)

func main() {
	input, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatalf("Error reading input file %v", err)
	}

	passphrases := strings.Split(string(input), "\n")

	var numValid int

outer:
	for _, passphrase := range passphrases {
		seen := make(map[string]struct{})
		words := strings.Fields(passphrase)
		for _, word := range words {
			if _, ok := seen[word]; ok {
				continue outer
			}
			seen[word] = struct{}{}
		}

		numValid++
	}

	log.Println("Part 1 answer is", numValid)

	numValid = 0

invalid:
	for _, passphrase := range passphrases {
		sizes := make(map[int][]string)
		for _, word := range strings.Fields(passphrase) {
			sizes[len(word)] = append(sizes[len(word)], word)
		}

		for size, words := range sizes {
			if size < 2 {
				continue
			}

			found := map[string]struct{}{}

			for _, word := range words {
				ordered := SortString(word)
				if _, ok := found[ordered]; ok {
					continue invalid
				}
				found[ordered] = struct{}{}
			}
		}

		numValid++
	}

	log.Println("Part 2 answer is", numValid)
}

type sortRunes []rune

func (s sortRunes) Less(i, j int) bool {
	return s[i] < s[j]
}

func (s sortRunes) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s sortRunes) Len() int {
	return len(s)
}

func SortString(s string) string {
	r := []rune(s)
	sort.Sort(sortRunes(r))
	return string(r)
}
