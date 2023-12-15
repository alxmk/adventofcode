package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	data, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatalln("Failed to read input:", err)
	}

	p1, p2 := solve(parseScanners(string(data)))

	log.Println("Part one:", p1)
	log.Println("Part two:", p2)
}

func solve(scanners []scanner) (int, int) {
	known := []scanner{scanners[0]}
	unknown := scanners[1:]

	known[0].oriented = known[0].readings

outer:
	for l := len(unknown); l > 0; l = len(unknown) {
		for i := 0; i < len(unknown); i++ {
			for j := 0; j < len(known); j++ {
				if s, ok := unknown[i].Overlaps(known[j]); ok {
					known = append(known, s)
					unknown = append(unknown[:i], unknown[i+1:]...)
					continue outer
				}
			}
		}
	}

	oriented := make(map[xyz]struct{})
	for _, s := range known {
		for b := range s.oriented {
			oriented[b] = struct{}{}
		}
	}

	var m int
	for i := 0; i < len(known); i++ {
		for j := i; j < len(known); j++ {
			m = max(m, known[i].translation.distance(known[j].translation))
		}
	}

	return len(oriented), m
}

func parseScanners(input string) []scanner {
	var scanners []scanner
	for _, block := range strings.Split(input, "\n\n") {
		scanners = append(scanners, parseScanner(block))
	}
	return scanners
}

func parseScanner(block string) scanner {
	s := scanner{
		readings:     make(map[xyz]struct{}),
		orientations: make([]map[xyz]struct{}, 24),
		oriented:     make(map[xyz]struct{}),
		distances:    make(map[int]int),
	}
	var readings []xyz
	for i, line := range strings.Split(block, "\n") {
		if i == 0 {
			continue
		}
		c := xyz{}
		coords := strings.Split(line, ",")
		c.x, _ = strconv.Atoi(coords[0])
		c.y, _ = strconv.Atoi(coords[1])
		c.z, _ = strconv.Atoi(coords[2])

		for j, o := range c.Transformations() {
			if s.orientations[j] == nil {
				s.orientations[j] = make(map[xyz]struct{})
			}
			s.orientations[j][o] = struct{}{}
		}

		s.readings[c] = struct{}{}
		readings = append(readings, c)
	}
	for i := 0; i < len(readings); i++ {
		for j := i + 1; j < len(readings); j++ {
			s.distances[readings[i].distance(readings[j])]++
		}
	}
	return s
}

type xyz struct {
	x, y, z int
}

func (a xyz) String() string {
	return fmt.Sprintf("%d,%d,%d", a.x, a.y, a.z)
}

func (a xyz) roll() xyz {
	return xyz{a.x, a.z, -a.y}
}

func (a xyz) turn() xyz {
	return xyz{-a.y, a.x, a.z}
}

func (a xyz) distance(b xyz) int {
	return max(a.x-b.x, b.x-a.x) + max(a.y-b.y, b.y-a.y) + max(a.z-b.z, b.z-a.z)
}

func (a xyz) offset(b xyz) xyz {
	return xyz{a.x - b.x, a.y - b.y, a.z - b.z}
}

func (a xyz) translate(b xyz) xyz {
	return xyz{a.x + b.x, a.y + b.y, a.z + b.z}
}

func (a xyz) Transformations() []xyz {
	t := make([]xyz, 0, 24)
	for c := 0; c < 2; c++ {
		for s := 0; s < 3; s++ {
			t = append(t, a)
			a = a.roll()
			for turn := 0; turn < 3; turn++ {
				t = append(t, a)
				a = a.turn()
			}
		}
		a = a.roll().turn().roll()
	}
	return t
}

type scanner struct {
	readings     map[xyz]struct{}
	orientations []map[xyz]struct{}
	oriented     map[xyz]struct{}
	distances    map[int]int
	translation  xyz
}

func (s scanner) Overlaps(t scanner) (scanner, bool) {
	var count int
	distances := make(map[int]struct{})
	for d, sc := range s.distances {
		if tc, ok := t.distances[d]; ok {
			count += min(tc, sc)
			distances[d] = struct{}{}
		}
	}
	// Number of matching lengths between the points Sn = n(n+1)/2
	if count < (11 * 12 / 2) {
		return s, false
	}

	lookup := make(map[xyz]struct{})
	for a := range t.oriented {
		for b := range t.oriented {
			if a == b {
				continue
			}
			if _, ok := distances[a.distance(b)]; ok {
				lookup[a], lookup[b] = struct{}{}, struct{}{}
			}
		}
	}

	// For each orientation of the unfixed scanner s1
	for _, orientation := range s.orientations {
		// For each point we're trying to match up relative to fixed scanner s0
		for l := range lookup {
			// For each point in this orientation of the unfixed scanner s1 (which we hope matches fixed
			// scanner s0)
			for o := range orientation {
				// Find the offset between the point on s0 we're looking up and the point on s1
				offset := l.offset(o)
				var count int
				for m := range lookup {
					if _, ok := orientation[m.offset(offset)]; ok {
						count++
					}
				}
				if count < 12 {
					continue
				}
				// We found it!
				reverse := o.offset(l)
				for p := range orientation {
					s.oriented[p.offset(reverse)] = struct{}{}
				}
				s.translation = offset
				return s, true
			}
		}
	}

	return s, false
}
