package main

import (
	"io/ioutil"
	"log"
	"strings"
)

func main() {
	data, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatalln("Failed to read input:", err)
	}

	log.Println("Part one:", parse(string(data)).SteadyState(adjacentRules))
	log.Println("Part two:", parse(string(data)).SteadyState(vectorRules))
}

type seats struct {
	data       map[int]map[int]rune
	xmax, ymax int
	state      string
}

func (s seats) String() string {
	var b strings.Builder
	for y := 0; y <= s.ymax; y++ {
		for x := 0; x <= s.xmax; x++ {
			b.WriteRune(s.data[x][y])
		}
		b.WriteRune('\n')
	}
	return b.String()
}

// Iterate runs through one iteration of the occupancy rules and
// then returns a boolean corresponding to whether the state has changed following
// the rules or not
func (s *seats) Iterate(occupancyRules rules) bool {
	newData := make(map[int]map[int]rune)
	for y := 0; y <= s.ymax; y++ {
		for x := 0; x <= s.xmax; x++ {
			if _, ok := newData[x]; !ok {
				newData[x] = make(map[int]rune)
			}
			tile := s.data[x][y]
			occupied, threshold := occupancyRules(x, y, s)
			if occupied == 0 && tile == 'L' {
				newData[x][y] = '#'
				continue
			}
			if occupied >= threshold && tile == '#' {
				newData[x][y] = 'L'
				continue
			}
			newData[x][y] = tile
		}
	}

	s.data = newData
	if newState := s.String(); newState != s.state {
		s.state = newState
		return true
	}
	return false
}

type rules func(x, y int, s *seats) (int, int)

func adjacentRules(x, y int, s *seats) (int, int) {
	var occupied int
	for i := -1; i <= 1; i++ {
		for j := -1; j <= 1; j++ {
			if i == 0 && j == 0 {
				continue
			}
			if column, ok := s.data[x+i]; ok {
				if seat, ok := column[y+j]; ok && seat == '#' {
					occupied++
				}
			}
		}
	}

	return occupied, 4
}

func vectorRules(x, y int, s *seats) (int, int) {
	var occupied int
	for i := -1; i <= 1; i++ {
		for j := -1; j <= 1; j++ {
			if i == 0 && j == 0 {
				continue
			}
			for scale := 1; ; scale++ {
				thisX, thisY := x+(j*scale), y+(i*scale)
				// Check occupied
				if column, ok := s.data[thisX]; ok {
					if seat, ok := column[thisY]; ok {
						switch seat {
						case '#':
							occupied++
						case 'L':
						default:
							continue
						}
						break
					}
				}
				// Check bounds
				if thisX > s.xmax || thisX < 0 || thisY > s.ymax || thisY < 0 {
					break
				}
			}
		}
	}
	return occupied, 5
}

func (s *seats) SteadyState(r rules) int {
	for s.Iterate(r) {
	}
	return strings.Count(s.state, "#")
}

func parse(input string) *seats {
	s := seats{data: make(map[int]map[int]rune)}
	for y, line := range strings.Split(input, "\n") {
		for x, r := range line {
			if _, ok := s.data[x]; !ok {
				s.data[x] = make(map[int]rune)
			}
			s.data[x][y] = r
			if x > s.xmax {
				s.xmax = x
			}
		}
		if y > s.ymax {
			s.ymax = y
		}
	}
	return &s
}
