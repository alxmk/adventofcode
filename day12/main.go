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

	lines := strings.Split(strings.Replace(string(input), ",", "", -1), "\n")

	srcDest := make(map[string][]string)

	for _, l := range lines {
		parts := strings.Fields(l)

		srcDest[parts[0]] = parts[2:]
	}

	// Start at node 0
	start := "0"
	visited := make(map[string]struct{})

	traverse(srcDest, visited, start)

	log.Println("The answer is", len(visited))
}

func traverse(srcDest map[string][]string, visited map[string]struct{}, current string) {
	visited[current] = struct{}{}

	for _, dest := range srcDest[current] {
		if _, ok := visited[dest]; !ok {
			traverse(srcDest, visited, dest)
		}
	}
}
