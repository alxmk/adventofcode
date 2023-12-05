package main

import (
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	data, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatalln("Failed to read input:", err)
	}

	a := parseAlmanac(string(data))

	log.Println("Part one:", partOne(a))
	log.Println("Part two:", partTwo(a))
}

func parseAlmanac(input string) almanac {
	a := almanac{mappings: make(map[string][]mapping), reverse: make(map[string][]mapping)}
	for _, chunk := range strings.Split(input, "\n\n") {
		if strings.HasPrefix(chunk, "seeds:") {
			seedsraw := strings.Fields(strings.TrimPrefix(chunk, "seeds: "))
			for _, s := range seedsraw {
				v, _ := strconv.Atoi(s)
				a.seeds = append(a.seeds, v)
			}
			continue
		}
		var m mapping
		for _, line := range strings.Split(chunk, "\n") {
			parts := strings.Fields(line)
			switch len(parts) {
			case 2:
				subparts := strings.Split(parts[0], "-")
				m.from, m.to = subparts[0], subparts[2]
			case 3:
				for i, p := range parts {
					v, _ := strconv.Atoi(p)
					switch i {
					case 0:
						m.dst = v
					case 1:
						m.src = v
					case 2:
						m.rng = v
					}
				}
				a.mappings[m.from] = append(a.mappings[m.from], m)
				a.reverse[m.to] = append(a.reverse[m.to], m)
			}
		}
	}
	return a
}

func partOne(a almanac) int {
	closest := math.MaxInt
	for _, s := range a.seeds {
		thing := "seed"
		for s, thing = a.Corresponding(s, thing); thing != "location"; s, thing = a.Corresponding(s, thing) {
		}
		closest = min(s, closest)
	}
	return closest
}

func partTwo(a almanac) int {
	var rngs []rng
	for i := 0; i < len(a.seeds); i += 2 {
		rngs = append(rngs, rng{a.seeds[i], a.seeds[i+1]})
	}
	for i := 0; ; i++ {
		s := i
		thing := "location"
		for s, thing = a.Reverse(s, thing); thing != "seed"; s, thing = a.Reverse(s, thing) {
		}
		for _, r := range rngs {
			if r.Within(s) {
				return i
			}
		}
	}
}

type rng struct {
	start, size int
}

func (r rng) Within(v int) bool {
	return v >= r.start && v <= r.start+r.size
}

type mapping struct {
	src, dst, rng int
	from, to      string
}

func (m mapping) Corresponding(input int) (int, bool) {
	if input >= m.src && input <= m.src+m.rng {
		return m.dst + (input - m.src), true
	}
	return 0, false
}

func (m mapping) Reverse(input int) (int, bool) {
	if input >= m.dst && input <= m.dst+m.rng {
		return m.src + (input - m.dst), true
	}
	return 0, false
}

type almanac struct {
	seeds    []int
	mappings map[string][]mapping
	reverse  map[string][]mapping
}

func (a almanac) Corresponding(input int, thing string) (int, string) {
	for _, m := range a.mappings[thing] {
		v, ok := m.Corresponding(input)
		if ok {
			return v, m.to
		}
	}
	return input, a.mappings[thing][0].to
}

func (a almanac) Reverse(input int, thing string) (int, string) {
	for _, m := range a.reverse[thing] {
		v, ok := m.Reverse(input)
		if ok {
			return v, m.from
		}
	}
	return input, a.reverse[thing][0].from
}
