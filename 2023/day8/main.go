package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	data, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatalln("Failed to read input:", err)
	}

	log.Println("Part one:", partOne(parseNodes(string(data))))
	log.Println("Part two:", partTwo(parseNodes(string(data))))
}

func partOne(route string, nodes map[string]*node) int {
	var step int
	loc := "AAA"
	for i := 0; ; i++ {
		step++
		switch route[i] {
		case 'L':
			loc = nodes[loc].l
		case 'R':
			loc = nodes[loc].r
		}
		if loc == "ZZZ" {
			break
		}
		if i == len(route)-1 {
			i = -1
		}
	}
	return step
}

func partTwo(route string, nodes map[string]*node) int {
	var step int
	locs := make(map[string]struct{})
	for n := range nodes {
		if n[2] == 'A' {
			locs[n] = struct{}{}
		}
	}
	var periods []int
	for i := 0; ; i++ {
		step++
		newlocs := make(map[string]struct{})
		for l := range locs {
			var newl string
			switch route[i] {
			case 'L':
				newl = nodes[l].l
			case 'R':
				newl = nodes[l].r
			}
			if newl[2] == 'Z' {
				periods = append(periods, step)
				continue
			}
			newlocs[newl] = struct{}{}
		}
		if len(newlocs) == 0 {
			break
		}
		locs = newlocs
		if i == len(route)-1 {
			i = -1
		}
	}
	return lowestCommonMultiple(periods...)
}

func parseNodes(input string) (string, map[string]*node) {
	parts := strings.Split(input, "\n\n")
	nodes := make(map[string]*node)
	for _, line := range strings.Split(parts[1], "\n") {
		n := parseNode(line)
		nodes[n.name] = n
	}
	return parts[0], nodes
}

func parseNode(line string) *node {
	var n node
	fmt.Sscanf(line, "%s = (%s %s)", &n.name, &n.l, &n.r)
	n.l, n.r = strings.TrimSuffix(n.l, ","), strings.TrimSuffix(n.r, ")")
	return &n
}

type node struct {
	name string
	l, r string
}

func lowestCommonMultiple(is ...int) int {
	factors := make(map[int]int)
	for _, i := range is {
		for k, n := range primeFactors(i) {
			if factors[k] < n {
				factors[k] = n
			}
		}
	}

	lcm := 1
	for k, n := range factors {
		for i := 0; i < n; i++ {
			lcm *= k
		}
	}

	return lcm
}

func primeFactors(n int) map[int]int {
	pf := make(map[int]int)
	for i := n % 2; i == 0; i = n % 2 {
		pf[2]++
		n /= 2
	}
	for i := 3; i*i <= n; i += 2 {
		for j := n % i; j == 0; j = n % i {
			pf[i]++
			n /= i
		}
		if n == 1 {
			return pf
		}
	}
	if n > 2 {
		pf[n]++
	}
	return pf
}
