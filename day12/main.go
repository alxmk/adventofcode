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
	allNodes := []string{}

	for _, l := range lines {
		parts := strings.Fields(l)

		srcDest[parts[0]] = parts[2:]
		allNodes = append(allNodes, parts[0])
	}

	// Start at node 0
	start := "0"
	visited := make(map[string]struct{})

	traverse(srcDest, visited, start)

	log.Println("Part one answer is", len(visited))

	index := 0
	numGroups := 1

	log.Println("Total nodes", len(allNodes))

	// Find the next unvisited node
	for i := index; i < len(allNodes); i++ {
		if _, ok := visited[allNodes[i]]; !ok {
			traverse(srcDest, visited, allNodes[i])
			numGroups++
		}
		index = i
	}

	log.Println("Part two answer is", numGroups)
}

func traverse(srcDest map[string][]string, visited map[string]struct{}, current string) {
	visited[current] = struct{}{}

	for _, dest := range srcDest[current] {
		if _, ok := visited[dest]; !ok {
			traverse(srcDest, visited, dest)
		}
	}
}
