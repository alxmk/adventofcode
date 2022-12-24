package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"strings"
)

func main() {
	data, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatalln("Failed to read input:", err)
	}

	log.Println("Part one:", partOne(parse(string(data))))
	log.Println("Part two:", partTwo(parse(string(data))))
}

func partOne(bps []blueprint) int {
	var sum int
	for _, bp := range bps {
		sum += (&solver{scache: make(map[state]struct{}), bp: bp}).Solve(24) * bp.number
	}
	return sum
}

func partTwo(bps []blueprint) int {
	product := 1
	for _, bp := range bps[:3] {
		product *= (&solver{scache: make(map[state]struct{}), bp: bp}).Solve(32)
	}
	return product
}

const (
	bpfmt = "Blueprint %d: Each ore robot costs %d ore. Each clay robot costs %d ore. Each obsidian robot costs %d ore and %d clay. Each geode robot costs %d ore and %d obsidian."
)

func parse(input string) []blueprint {
	var bps []blueprint
	for _, line := range strings.Split(input, "\n") {
		var bp blueprint
		fmt.Sscanf(line, bpfmt, &bp.number, &bp.ore.ore, &bp.clay.ore, &bp.obsidian.ore, &bp.obsidian.clay, &bp.geode.ore, &bp.geode.obsidian)
		bps = append(bps, bp)
	}
	return bps
}

type blueprint struct {
	number                     int
	ore, clay, obsidian, geode resources
}

type resources struct {
	ore, clay, obsidian, geode int
}

func (r resources) Less(o resources) bool {
	return !(r.ore > o.ore) && !(r.clay > o.clay) && !(r.obsidian > o.obsidian) && !(r.geode > o.geode)
}

type state struct {
	t                 int
	inventory, robots resources
}

func (s state) Less(o state) bool {
	return s.t >= o.t && s.inventory.Less(o.inventory) && s.robots.Less(o.robots)
}

func (s state) Tick(t int) state {
	return state{
		robots: s.robots,
		inventory: resources{
			ore:      s.inventory.ore + (t * s.robots.ore),
			clay:     s.inventory.clay + (t * s.robots.clay),
			obsidian: s.inventory.obsidian + (t * s.robots.obsidian),
			geode:    s.inventory.geode + (t * s.robots.geode),
		},
		t: s.t + t,
	}
}

func (s state) Build(i int, r resources) state {
	n := state{
		inventory: resources{
			ore:      s.inventory.ore - r.ore,
			clay:     s.inventory.clay - r.clay,
			obsidian: s.inventory.obsidian - r.obsidian,
			geode:    s.inventory.geode - r.geode,
		},
		robots: s.robots,
		t:      s.t,
	}
	switch i {
	case 0:
		n.robots.ore++
	case 1:
		n.robots.clay++
	case 2:
		n.robots.obsidian++
	case 3:
		n.robots.geode++
	}
	return n
}

func (s state) BuildTime(r resources) int {
	var time int
	if t := calculateTime(s.inventory.ore, s.robots.ore, r.ore); t > time {
		time = t
	}
	if t := calculateTime(s.inventory.clay, s.robots.clay, r.clay); t > time {
		time = t
	}
	if t := calculateTime(s.inventory.obsidian, s.robots.obsidian, r.obsidian); t > time {
		time = t
	}
	return time
}

func calculateTime(inv, rate, cost int) int {
	if cost == 0 || inv >= cost {
		return 1
	}
	if rate == 0 {
		return math.MaxInt
	}
	return ((cost - inv) / rate) + 1 + func() int {
		if ((cost - inv) % rate) == 0 {
			return 0
		}
		return 1
	}()
}

type solver struct {
	scache map[state]struct{}
	bp     blueprint
}

func max(i ...int) int {
	var max int
	for _, n := range i {
		if n > max {
			max = n
		}
	}
	return max
}

type resource int

const (
	ore resource = iota
	clay
	obsidian
	geode
)

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

func (s *solver) Solve(time int) int {
	start := state{t: 0, robots: resources{ore: 1}}
	s.scache[start] = struct{}{}
	queue := []state{start}

	var best int

	var this state
	for len(queue) != 0 {
		this, queue = queue[0], queue[1:]
		// See if we can possibly get more than best
		remaining := time - this.t
		gs := this.inventory.geode + (this.robots.geode * remaining)
		for i := remaining; i > 0; i-- {
			gs += (remaining - 1)
		}
		if gs <= best {
			continue
		}
		if this.BuildTime(s.bp.geode) == 1 {
			// If we can always build a geode then we can just solve
			if this.robots.ore >= s.bp.geode.ore && this.robots.obsidian >= s.bp.geode.obsidian {
				remaining := time - this.t
				gs := this.inventory.geode + (this.robots.geode * remaining)
				for i := remaining; i > 0; i-- {
					gs += (remaining - 1)
				}
				if gs > best {
					best = gs
				}
				continue
			}
		}
		for i, r := range []resources{s.bp.ore, s.bp.clay, s.bp.obsidian, s.bp.geode} {
			// Try to optimise a bit idk
			switch i {
			case 0:
				// If we have as much ore as we could ever use then don't build any more
				if this.robots.ore >= max(s.bp.clay.ore, s.bp.obsidian.ore, s.bp.geode.ore) {
					continue
				}
			case 1:
				if this.robots.clay >= s.bp.obsidian.clay {
					continue
				}
			case 2:
				if this.robots.obsidian >= s.bp.geode.obsidian {
					continue
				}
			}
			t := this.BuildTime(r)
			if remaining := time - this.t; t > remaining {
				if score := this.Tick(remaining).inventory.geode; score > best {
					best = score
				}
				continue
			}
			next := this.Tick(t).Build(i, r)
			if _, ok := s.scache[next]; ok {
				continue
			}
			s.scache[next] = struct{}{}
			queue = append(queue, next)
		}
	}

	return best
}
