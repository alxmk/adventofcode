package main

import (
	"log"
	"math/rand"
	"os"
	"strings"
)

func main() {
	data, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatalln("Failed to read input:", err)
	}

	log.Println("Part one:", solve(parseWires(string(data))))
}

func solve(g graph) int {
	for c, v := kargers(g); ; c, v = kargers(g) {
		if c == 3 {
			return v
		}
	}
}

func findIndexes(subsets []map[string]struct{}, e edge) (int, int) {
	var ss, sd int
	for i, sub := range subsets {
		if _, ok := sub[e.s]; ok {
			ss = i
		}
		if _, ok := sub[e.d]; ok {
			sd = i
		}
	}
	return ss, sd
}

func kargers(g graph) (int, int) {
	var subsets []map[string]struct{}
	for _, v := range g.vertices {
		subsets = append(subsets, map[string]struct{}{v: {}})
	}
	v := len(g.vertices)
	for v > 2 {
		ss, sd := findIndexes(subsets, g.edges[rand.Intn(len(g.edges))])
		if ss == sd {
			continue
		}
		v--
		for j := range subsets[sd] {
			subsets[ss][j] = struct{}{}
		}
		subsets = append(subsets[:sd], subsets[sd+1:]...)
	}
	var cuts int
	for _, e := range g.edges {
		ss, sd := findIndexes(subsets, e)
		if ss != sd {
			cuts++
		}
	}
	return cuts, len(subsets[0]) * len(subsets[1])
}

type edge struct {
	s, d string
}

type graph struct {
	edges    []edge
	vertices []string
}

func parseWires(input string) graph {
	var g graph
	vertices := make(map[string]struct{})
	for _, line := range strings.Split(input, "\n") {
		var a string
		for i, f := range strings.Fields(line) {
			if i == 0 {
				a = f[:len(f)-1]
				vertices[a] = struct{}{}
				continue
			}
			g.edges = append(g.edges, edge{a, f})
			vertices[f] = struct{}{}
		}
	}
	for v := range vertices {
		g.vertices = append(g.vertices, v)
	}
	return g
}
