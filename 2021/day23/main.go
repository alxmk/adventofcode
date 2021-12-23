package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"sort"
	"strings"
)

func main() {
	data, err := ioutil.ReadFile("input1.txt")
	if err != nil {
		log.Fatalln("Failed to read input:", err)
	}

	s, err := parse(string(data))
	if err != nil {
		log.Fatalln("Failed to parse input:", err)
	}

	r := solver{bestSolution: math.MaxInt, cache: make(map[string]int)}
	log.Println("Part one:", r.Solve(s, 0))

	data, err = ioutil.ReadFile("input2.txt")
	if err != nil {
		log.Fatalln("Failed to read input:", err)
	}

	s, err = parse(string(data))
	if err != nil {
		log.Fatalln("Failed to parse input:", err)
	}

	r = solver{bestSolution: math.MaxInt, cache: make(map[string]int)}
	log.Println("Part two:", r.Solve(s, 0))
}

func parse(input string) (*state, error) {
	s := state{amphipods: make(map[coord]rune)}
	for y, line := range strings.Split(input, "\n") {
		for x, r := range line {
			switch r {
			case 'A', 'B', 'C', 'D':
				s.amphipods[coord{x, y}] = r
				if y > s.ymax {
					s.ymax = y
				}
			}
		}
	}
	return &s, nil
}

type coord struct {
	x, y int
}

func (c coord) String() string {
	return fmt.Sprintf("%d,%d", c.x, c.y)
}

func (c coord) InRoom() bool {
	switch c.x {
	case 3, 5, 7, 9:
		return c.y != 1
	}
	return false
}

func (c coord) InCorrectRoom(r rune) bool {
	if !c.InRoom() {
		return false
	}
	switch r {
	case 'A':
		return c.x == 3
	case 'B':
		return c.x == 5
	case 'C':
		return c.x == 7
	case 'D':
		return c.x == 9
	}
	return false
}

type state struct {
	amphipods map[coord]rune
	ymax      int
}

type move struct {
	from, to coord
}

func (s state) CheapStringer() string {
	var coordstrs []string
	for c, r := range s.amphipods {
		coordstrs = append(coordstrs, c.String()+string(r))
	}
	sort.Strings(coordstrs)
	return strings.Join(coordstrs, "")
}

func (s state) String() string {
	var b strings.Builder
	for y := 0; y <= s.ymax+1; y++ {
		for x := 0; x < 13; x++ {
			switch y {
			case 0:
				b.WriteRune('#')
			case 1:
				switch x {
				case 0, 12:
					b.WriteRune('#')
				default:
					if r, ok := s.amphipods[coord{x, y}]; ok {
						b.WriteRune(r)
						continue
					}
					b.WriteRune('.')
				}
			case 2:
				switch x {
				case 0, 1, 2, 4, 6, 8, 10, 11, 12:
					b.WriteRune('#')
				default:
					if r, ok := s.amphipods[coord{x, y}]; ok {
						b.WriteRune(r)
						continue
					}
					b.WriteRune('.')

				}
			case s.ymax + 1:
				switch x {
				case 0, 1, 11, 12:
					b.WriteRune(' ')
				default:
					b.WriteRune('#')
				}
			default:
				switch x {
				case 0, 1, 11, 12:
					b.WriteRune(' ')
				case 2, 4, 6, 8, 10:
					b.WriteRune('#')
				default:
					if r, ok := s.amphipods[coord{x, y}]; ok {
						b.WriteRune(r)
						continue
					}
					b.WriteRune('.')
				}
			}
		}
		b.WriteRune('\n')
	}
	return b.String()
}

func (s *state) Apply(m move) *state {
	t := &state{amphipods: make(map[coord]rune), ymax: s.ymax}
	r := s.amphipods[m.from]
	t.amphipods[m.to] = r
	for k, v := range s.amphipods {
		if k == m.from {
			continue
		}
		t.amphipods[k] = v
	}
	return t
}

func (s *state) ValidMoves(current, best int) map[move]int {
	standardMoves := make(map[move]int)
	bestMoves := make(map[move]int)
	for c, r := range s.amphipods {
		valid, foundBest := s.validMovesInner(c, r, current, best)
		// Best moves should always be done
		if foundBest {
			for k, v := range valid {
				bestMoves[k] = v
			}
			continue
		}
		for k, v := range valid {
			standardMoves[k] = v
		}
	}
	if len(bestMoves) != 0 {
		return bestMoves
	}
	return standardMoves
}

func (s *state) validMovesInner(c coord, r rune, current, best int) (map[move]int, bool) {
	// If the amphipod is in the correct room and either it's at the back or the ones behind it
	// are also correct then it has no valid moves
	if c.InCorrectRoom(r) && (c.y == s.ymax || func() bool {
		for y := s.ymax; y > c.y; y-- {
			if s.amphipods[coord{c.x, y}] != r {
				return false
			}
		}
		return true
	}()) {
		return nil, false
	}
	// Is there a valid space in its own room?
	space, d := s.SpaceInRoom(r)
	if space {
		// Check if it's pathable
		pathable, distance := s.Pathable(c, d)
		if pathable {
			// There's space in the room and it's pathable, so this is always the best move
			return map[move]int{{c, d}: int(math.Pow10(int(r-'A'))) * distance}, true
		}
	}
	// If it's not in a room it can't move except to go into its own room
	if !c.InRoom() {
		return nil, false
	}
	// Lastly, check the options it has in the corridor
	moves := make(map[move]int)
	multiplier := int(math.Pow10(int(r - 'A')))
	for _, d := range corridorCoords {
		// Shortcut some stupid moves
		if c.Distance(d)*multiplier+current > best {
			continue
		}
		if pathable, distance := s.Pathable(c, d); pathable {
			moves[move{c, d}] = multiplier * distance
		}
	}
	return moves, false
}

func (c coord) Distance(d coord) int {
	xdist := c.x - d.x
	if xdist < 0 {
		xdist *= -1
	}
	return xdist + (d.y - 1) + (c.y - 1)
}

var corridorCoords = []coord{
	{1, 1},
	{2, 1},
	{4, 1},
	{6, 1},
	{8, 1},
	{10, 1},
	{11, 1},
}

func (s *state) Pathable(a, b coord) (bool, int) {
	// Shortcut trying to path to an occupied space
	if _, ok := s.amphipods[b]; ok {
		return false, -1
	}
	switch a.y {
	case 2, 3, 4, 5:
		// In a room
		// We could be going to a corridor or to a room
		switch b.y {
		case 2, 3, 4, 5:
			// Going to a room we need to path to the corridor then call this again to hit case 1
			for y := a.y - 1; y >= 1; y-- {
				if _, ok := s.amphipods[coord{a.x, y}]; ok {
					return false, -1
				}
			}
			if pathable, distance := s.Pathable(coord{a.x, 1}, b); pathable {
				return true, distance + a.y - 1
			}
			return false, -1
		case 1:
			// Going to a corridor we need to path up then along
			for y := a.y - 1; y >= 1; y-- {
				if _, ok := s.amphipods[coord{a.x, y}]; ok {
					return false, -1
				}
			}
			start, end := a.x, b.x
			if end < start {
				end, start = start, end
			}
			// Check to see if there's anything blocking the corridor
			for x := start + 1; x <= end; x++ {
				if _, ok := s.amphipods[coord{x, 1}]; ok {
					return false, -1
				}
			}
			// Otherwise it's pathable
			return true, a.y - 1 + end - start
		}
	case 1:
		// In corridor
		// We must be moving to a room which will involve going left or right then down
		start, end := a.x, b.x
		if end < start {
			end, start = start, end
		}
		// Check to see if there's anything blocking the corridor
		for x := start; x <= end; x++ {
			// Don't block with ourselves
			if x == a.x || x == b.x {
				continue
			}
			if _, ok := s.amphipods[coord{x, 1}]; ok {
				return false, -1
			}
		}
		// We already know we can get into the room so it's pathable
		return true, end - start + b.y - 1
	}
	return false, -1
}

func (s *state) SpaceInRoom(r rune) (bool, coord) {
	var x int
	switch r {
	case 'A':
		x = 3
	case 'B':
		x = 5
	case 'C':
		x = 7
	case 'D':
		x = 9
	}
	// Check if any occupants are of the wrong kind
	for y := s.ymax; y > 1; y-- {
		q, present := s.amphipods[coord{x, y}]
		if present && q != r {
			return false, coord{}
		}
		if !present {
			return true, coord{x, y}
		}
	}
	// Otherwise it's empty
	return true, coord{x, s.ymax}
}

type solver struct {
	bestSolution int
	cache        map[string]int
}

func (r *solver) Solve(s *state, score int) int {
	// Check the cache, if we have a hit with a lower score then this is a dead end
	str := s.CheapStringer()
	if l, ok := r.cache[str]; ok && l < score {
		return r.bestSolution
	}
	r.cache[str] = score
	for move, addScore := range s.ValidMoves(score, r.bestSolution) {
		if (score + addScore) > r.bestSolution {
			// Dead end, we're already worse than the best solution
			continue
		}
		r.Solve(s.Apply(move), score+addScore)
	}
	// If no valid moves remain and we hit a win condition then update the best solution if it's better
	if s.Win() {
		if score < r.bestSolution {
			r.bestSolution = score
		}
	}
	// Otherwise this isn't a win condition
	return r.bestSolution
}

func (s *state) Win() bool {
	for y := s.ymax; y > 1; y-- {
		if s.amphipods[coord{3, y}] != 'A' {
			return false
		}
		if s.amphipods[coord{5, y}] != 'B' {
			return false
		}
		if s.amphipods[coord{7, y}] != 'C' {
			return false
		}
		if s.amphipods[coord{9, y}] != 'D' {
			return false
		}
	}
	return true
}
