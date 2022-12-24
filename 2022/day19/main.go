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

	bps := parse(string(data))
	log.Println("Part one:", partOne(bps))
}

func partOne(bps []blueprint) int {
	var sum int
	for _, b := range bps {
		sum += b.quality()
	}
	return sum
}

func parse(input string) []blueprint {
	var bps []blueprint
	for _, line := range strings.Split(input, "\n") {
		var b blueprint
		var n, o, c, oo, oc, gor, gob int
		fmt.Sscanf(line,
			"Blueprint %d: Each ore robot costs %d ore. Each clay robot costs %d ore. Each obsidian robot costs %d ore and %d clay. Each geode robot costs %d ore and %d obsidian.",
			&n,
			&o,
			&c,
			&oo,
			&oc,
			&gor,
			&gob,
		)
		b.costs = map[resource]rset{
			ore:      {ore: o},
			clay:     {ore: c},
			obsidian: {ore: oo, clay: oc},
			geode:    {ore: gor, obsidian: gob},
		}
		bps = append(bps, b)
	}
	return bps
}

type blueprint struct {
	costs map[resource]rset
}

func (b blueprint) quality() int {
	root := &node{time: 0, s: &state{robots: rset{ore: 1}}}
	root.PopulateTree(b)
	return -1
}

type node struct {
	time int
	s    *state
	next []*node
}

var tmax = 24

func (n *node) PopulateTree(b blueprint) {
	// log.Println("Enter", n.time, n.s.resources, n.s.robots)
	if n.time >= tmax {
		if n.s.resources.geode > 23 {
			log.Println(n.s.resources.geode)
		}
		return
	}
	// Shortcut if we can build a geode now that's the only choice
	if buildable, time := n.s.MinsToBuild(geode, b); buildable && time == 1 {
		// log.Println("Build", geode, "at", n.time+1)
		t := n.s.Copy().Tick(1)
		t.robots.Add(geode, 1)
		t.resources.Build(b.costs[geode])
		next := &node{time: n.time + 1, s: t}
		n.next = []*node{next}
		next.PopulateTree(b)
		return
	}
	options := make(map[resource]struct{})
	for _, x := range resources {
		if buildable, time := n.s.MinsToBuild(x, b); buildable && n.time+time <= 24 {
			options[x] = struct{}{}
			// log.Println("Build", x, "at", n.time+time)
			t := n.s.Copy().Tick(time)
			t.robots.Add(x, 1)
			t.resources.Build(b.costs[x])
			next := &node{time: n.time + time, s: t}
			n.next = append(n.next, next)
			next.PopulateTree(b)
		}
	}
}

type state struct {
	resources rset
	robots    rset
}

func (s *state) Tick(mins int) *state {
	for min := 0; min < mins; min++ {
		s.resources.ore += s.robots.ore
		s.resources.clay += s.robots.clay
		s.resources.obsidian += s.robots.obsidian
		s.resources.geode += s.robots.geode
	}
	return s
}

func (s *state) Copy() *state {
	return &state{rset{s.resources.ore, s.resources.clay, s.resources.obsidian, s.resources.geode}, rset{s.robots.ore, s.robots.clay, s.robots.obsidian, s.robots.geode}}
}

func (s *state) MinsToBuild(r resource, b blueprint) (bool, int) {
	var time int
	for _, x := range resources {
		if ok, t := ttb(b.costs[r].Get(x), s.resources.Get(x), s.robots.Get(x)); ok {
			if t > time {
				time = t
			}
			continue
		}
		return false, -1
	}
	return true, time
}

func ttb(cost, resource, robots int) (bool, int) {
	if cost == 0 || resource >= cost {
		return true, 1
	}
	if robots == 0 {
		return false, -1
	}
	return true, 1 + (cost-resource)/robots + func() int {
		if ((cost - resource) % robots) != 0 {
			return 1 // round up
		}
		return 0
	}()
}

type rset struct {
	ore, clay, obsidian, geode int
}

func (r rset) Get(s resource) int {
	switch s {
	case ore:
		return r.ore
	case clay:
		return r.clay
	case obsidian:
		return r.obsidian
	case geode:
		return r.geode
	}
	return -1
}

func (r *rset) Add(s resource, i int) {
	switch s {
	case ore:
		r.ore += i
	case clay:
		r.clay += i
	case obsidian:
		r.obsidian += i
	case geode:
		r.geode += i
	}
}

func (r *rset) Build(c rset) {
	r.ore -= c.ore
	r.clay -= c.clay
	r.obsidian -= c.obsidian
	r.geode -= c.geode
}

type resource int

const (
	ore resource = iota
	clay
	obsidian
	geode
)

var resources = []resource{ore, clay, obsidian, geode}

func (r resource) String() string {
	switch r {
	case ore:
		return "ore"
	case clay:
		return "clay"
	case obsidian:
		return "obsidian"
	case geode:
		return "geode"
	}
	return "?"
}
