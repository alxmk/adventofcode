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
}

func partOne(bps []blueprint) int {
	var sum int
	for _, b := range bps {
		_, n := b.quality(0, &state{robots: resources{ore: 1}})
		sum += n
	}
	log.Println(i)
	return sum
}

type blueprint struct {
	costs map[resource]*resources
}

func (b blueprint) buildable(r resource, s resources) bool {
	return b.costs[r].Less(s)
}

func (b *blueprint) quality(min int, s *state) (int, int) {
	states := map[state]struct{}{*s: {}}
	var quality int
	for {
		best := math.MaxInt
		var next map[state]struct{}
		for s := range states {
			for _, c := range constituents {
				this, these := b.earliestGeode(min, s.Copy(), c)
				if this < best {
					for reduced := true; reduced; reduced, these = reduce(these) {
					}
					best, next = this, these
				}
			}
		}
		if best == 24 {
			break
		}
		quality += 24 - best
		log.Println(quality)
		states = next
	}
	log.Println(quality)
	return 24, quality
}

func reduce(states map[state]struct{}) (bool, map[state]struct{}) {
	reduced := make(map[state]struct{})
	for s := range states {
		var less bool
		for r := range reduced {
			if s.Less(r) {
				less = true
				break
			}
		}
		if !less {
			reduced[s] = struct{}{}
		}
	}
	return len(reduced) < len(states), reduced
}

func (b *blueprint) earliestGeode(min int, s *state, r resource) (int, map[state]struct{}) {
	j++
	// log.Println("Enter", min, r)
	for ; min < 24; min++ {
		// log.Println("Minute", min+1)
		if b.buildable(geode, s.resources) {
			// log.Println("Geode buildable")
			return min, map[state]struct{}{*s: {}}
		}
		var built *resource
		if b.buildable(r, s.resources) {
			// log.Println("Building", r)
			built = &r
			s.resources.Build(*b.costs[r])
		}
		s.resources.Add(ore, s.robots.ore)
		s.resources.Add(clay, s.robots.clay)
		s.resources.Add(obsidian, s.robots.obsidian)
		s.resources.Add(geode, s.robots.geode)
		if built != nil {
			s.robots.Add(*built, 1)
			minimum := math.MaxInt
			minResources := make(map[state]struct{})
			// If we have no obsidian and we can build one then do
			if s.robots.obsidian == 0 && b.buildable(obsidian, s.resources) {
				return b.earliestGeode(min+1, s.Copy(), obsidian)
			}
			for _, p := range []resource{obsidian, clay, ore} {
				switch p {
				case obsidian:
					// If we have no clay then obsidian is unbuildable
					if s.robots.clay == 0 {
						continue
					}
				}
				q, rs := b.earliestGeode(min+1, s.Copy(), p)
				if q < minimum {
					minimum = q
					minResources = make(map[state]struct{})
				}
				if q == minimum {
					for r := range rs {
						minResources[r] = struct{}{}
					}
				}
			}
			return minimum, minResources
		}
	}
	return 24, nil
}

var i, j int

// func (b *blueprint) quality(min int, robots, resources map[resource]int, r resource) int {
// 	// log.Println("Enter", min, r)
// 	for ; min < 24; min++ {
// 		// If we could have built a geode and we aren't then shortcut
// 		if r != geode && b.buildable(geode, resources) {
// 			return 0
// 		}
// 		// log.Println("Minute", min+1)
// 		var built *resource
// 		if b.buildable(r, resources) {
// 			// log.Println("Building", r)
// 			built = &r
// 			for t, n := range b.costs[r] {
// 				resources[t] -= n
// 			}
// 		}
// 		for r, n := range robots {
// 			resources[r] += n
// 			// log.Println(n, "robots collected", r, "we have", resources[r])
// 		}
// 		if built != nil {
// 			robots[*built]++
// 			var max int
// 			// If we can build a geode then do
// 			if b.buildable(geode, resources) {
// 				return b.quality(min+1, robots, resources, geode)
// 			}
// 			// If we have no obsidian and we can build one then do
// 			if robots[obsidian] == 0 && b.buildable(obsidian, resources) {
// 				return b.quality(min+1, robots, resources, obsidian)
// 			}
// 			for _, p := range priority {
// 				switch p {
// 				case geode:
// 					// If we have no obsidian then geode is unbuildable
// 					if robots[obsidian] == 0 {
// 						continue
// 					}
// 				case obsidian:
// 					// If we have no clay then obsidian is unbuildable
// 					if robots[clay] == 0 {
// 						continue
// 					}
// 				}
// 				if q := b.quality(min+1, copymap(robots), copymap(resources), p); q > max {
// 					max = q
// 				}
// 			}
// 			return max
// 		}
// 	}
// 	i++
// 	return resources[geode]
// }

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
		b.costs = map[resource]*resources{
			ore:      {ore: o},
			clay:     {ore: c},
			obsidian: {ore: oo, clay: oc},
			geode:    {ore: gor, obsidian: gob},
		}
		bps = append(bps, b)
	}
	return bps
}

type resource int

const (
	ore resource = iota
	clay
	obsidian
	geode
)

type state struct {
	robots    resources
	resources resources
}

func (s *state) Copy() *state {
	return &state{robots: *s.robots.Copy(), resources: *s.resources.Copy()}
}

func (s *state) Less(t state) bool {
	return s.resources.Less(t.resources) && s.robots.Less(t.robots)
}

type resources struct {
	ore, clay, obsidian, geode int
}

func (r *resources) Copy() *resources {
	return &resources{ore: r.ore, clay: r.clay, obsidian: r.obsidian, geode: r.geode}
}

func (r *resources) Get(s resource) int {
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

func (r *resources) Add(s resource, i int) {
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

func (r *resources) Less(s resources) bool {
	return s.ore >= r.ore && s.clay >= r.clay && s.obsidian >= r.obsidian && s.geode >= r.geode
}

func (r *resources) Build(s resources) {
	r.ore -= s.ore
	r.clay -= s.clay
	r.obsidian -= s.obsidian
	r.geode -= s.geode
}

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

var constituents = []resource{ore, clay, obsidian}
