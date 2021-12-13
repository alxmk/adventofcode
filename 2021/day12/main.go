package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

func main() {
	data, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatalln("Failed to read input:", err)
	}

	start, err := parse(string(data))
	if err != nil {
		log.Fatalln("Failed to parse input:", err)
	}

	log.Println("Part one:", len(traverse(start, "start", partOne)))
	log.Println("Part two:", len(traverse(start, "start", partTwo)))
}

func parse(input string) (*node, error) {
	nodes := make(map[string]*node)
	for _, line := range strings.Split(input, "\n") {
		parts := strings.Split(line, "-")
		if len(parts) != 2 {
			return nil, fmt.Errorf("malformed line: %s", line)
		}
		from, ok := nodes[parts[0]]
		if !ok {
			from = &node{name: parts[0], paths: make(map[string]*node)}
			nodes[parts[0]] = from
		}
		to, ok := nodes[parts[1]]
		if !ok {
			to = &node{name: parts[1], paths: make(map[string]*node)}
			nodes[parts[1]] = to
		}
		from.paths[to.name], to.paths[from.name] = to, from
	}
	return nodes["start"], nil
}

type node struct {
	name  string
	paths map[string]*node
}

func (n node) IsSmall() bool {
	return strings.ToLower(n.name) == n.name
}

func partOne(a *node, path string) bool {
	return a.IsSmall() && strings.Contains(path, a.name)
}

func partTwo(a *node, path string) bool {
	if !a.IsSmall() || !strings.Contains(path, a.name) {
		return false
	}
	if a.name == "start" {
		return true
	}
	var containsTwo bool
	for _, name := range strings.Split(path, ",") {
		if strings.ToLower(name) != name {
			continue
		}
		if strings.Count(path, name) == 2 {
			containsTwo = true
			break
		}
	}
	return strings.Contains(path, a.name) && containsTwo
}

func traverse(start *node, path string, traversalFunc func(*node, string) bool) []string {
	var paths []string
	for _, a := range start.paths {
		if a.name == "end" {
			paths = append(paths, path+","+a.name)
			continue
		}
		if traversalFunc(a, path) {
			continue
		}
		paths = append(paths, traverse(a, path+","+a.name, traversalFunc)...)
	}
	return paths
}
