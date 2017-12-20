package main

import (
	"io/ioutil"
	"log"
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

	log.Println("The answer is", numValid)
}
