package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"strings"
	"sync"

	"gonum.org/v1/gonum/stat/combin"
)

func main() {
	data, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatalln("Failed to read input:", err)
	}

	valves := parse(string(data))

	log.Println("Part one:", partOne(valves))
	log.Println("Part two:", partTwo(valves))
}

func partOne(n *network) int {
	return n.open("AA", n.permutations(len(n.unopen))[0], 0, 30, 0, 0)
}

func partTwo(n *network) int {
	var best int
	var m sync.Mutex
	var wg sync.WaitGroup
	for i := 1; i < len(n.unopen)/2+1; i++ {
		for _, perms := range n.permutations(i) {
			wg.Add(1)
			go func(perms []int) {
				sA, sB := n.open("AA", perms, 0, 26, 0, 0), n.open("AA", n.opposite(perms), 0, 26, 0, 0)
				m.Lock()
				if this := sA + sB; this > best {
					best = this
				}
				m.Unlock()
				wg.Done()
			}(perms)
		}
	}
	wg.Wait()
	return best
}

func (n *network) open(name string, unopened []int, i, max, minute, score int) int {
	if len(unopened) == 0 || minute >= max {
		return score
	}
	var best int
	for i, idx := range unopened {
		toopen := n.unopen[idx]
		relief := n.relief(name, toopen, minute, max)
		newunopened := make([]int, len(unopened))
		copy(newunopened, unopened)
		if this := n.open(toopen, append(newunopened[:i], newunopened[i+1:]...), i+1, max, minute+n.distance(name, toopen)+1, relief+score); this > best {
			best = this
		}
	}

	return best
}

func (n *network) permutations(i int) [][]int {
	return combin.Combinations(len(n.unopen), i)
}

func (n *network) opposite(perms []int) []int {
	var opp []int
outer:
	for i := range n.unopen {
		for _, j := range perms {
			if i == j {
				continue outer
			}
		}
		opp = append(opp, i)
	}
	return opp
}

func (n *network) relief(at, open string, minute, max int) int {
	return (max - minute - n.distance(at, open) - 1) * n.valves[open].rate
}

func (n *network) unopened() []string {
	var names []string
	for _, v := range n.valves {
		if !v.open && v.rate > 0 {
			names = append(names, v.name)
		}
	}
	return names
}

type network struct {
	sync.RWMutex
	valves map[string]*valve
	routes map[string]int
	unopen []string
}

type valve struct {
	name   string
	rate   int
	open   bool
	routes []string
}

func parse(input string) *network {
	n := network{valves: make(map[string]*valve)}
	for _, line := range strings.Split(input, "\n") {
		var v valve
		fmt.Sscanf(line, "Valve %s has flow rate=%d", &v.name, &v.rate)
		v.routes = strings.Split(strings.TrimSpace(strings.TrimPrefix(strings.Split(line, "valve")[1], "s")), ", ") // >:(
		n.valves[v.name] = &v
	}
	n.unopen = n.unopened()
	n.enumerateRoutes()
	return &n
}

func (n *network) enumerateRoutes() {
	n.routes = make(map[string]int)
	nodes := append(n.unopen, "AA")
	for i, a := range nodes {
		for j, b := range nodes {
			if i == j {
				continue
			}
			n.routes[a+b] = n.distance(a, b)
		}
	}
}

func (n *network) distance(a, b string) int {
	if r, ok := n.routes[a+b]; ok {
		return r
	}

	open := map[string]struct{}{a: {}}
	path := make(map[string]string)
	g := defaultMaxMap{a: 0}
	f := defaultMaxMap{a: g[a]}

	for len(open) != 0 {
		min := math.MaxInt
		var c string

		for o := range open {
			if score := f.Get(o); score < min {
				min, c = score, o
			}
		}

		if c == b {
			d := reconstructPath(path, c) - 1
			return d
		}

		delete(open, c)
		for _, d := range n.valves[c].routes {
			if d == c {
				continue
			}
			tg := g.Get(c)
			if tg != math.MaxInt64 {
				tg++
			}
			if tg < g.Get(d) {
				path[d], g[d], f[d], open[d] = c, tg, tg+1, struct{}{}
			}
		}
	}

	return -1
}

func (d defaultMaxMap) Get(c string) int {
	if v, ok := d[c]; ok {
		return v
	}
	return math.MaxInt
}

type defaultMaxMap map[string]int

func reconstructPath(path map[string]string, c string) int {
	totalPath := []string{c}
	for c, ok := path[c]; ok; c, ok = path[c] {
		totalPath = append(totalPath, c)
	}
	return len(totalPath)
}
