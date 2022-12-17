package main

import (
	"log"
	"os"
	"strings"
)

func main() {
	data, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatalln("Failed to read input:", err)
	}
	log.Println("Part one:", partOne(data))
	log.Println("Part two:", partTwo(data))
}

func partOne(g gas) int {
	t := tunnel{tiles: make(map[xy]tile), ymax: -1}
	ymax := t.Simulate(2022, g)
	return ymax + 1 // height is +1 because the bottom row of the tunnel is y=0
}

func partTwo(g gas) int {
	head, delta := findLoop(g)
	return solve(g, head, delta, 1000000000000) + 1
}

func findLoop(g gas) (int, int) {
	// chunk is our initial pattern size we look for, I arbitrarily chose 5 it works fine
	chunk := 5
	t := &tunnel{tiles: make(map[xy]tile), ymax: -1}
	var deltas []int
	for turn := 0; ; turn++ {
		// Tick the simulation (i.e. drop one shape)
		deltas = append(deltas, t.Tick(g, turn))
		// Skip the search unless there's a chance the chunk can repeat
		if len(deltas) < 2*chunk {
			continue
		}
		// Search for the last chunk
		s := len(deltas) - chunk
		if ok, r := repeats(deltas, deltas[s:]); ok {
			// If the last chunk repeats, see if the whole sequence between chunks repeats
			if ok, l := repeats(deltas, deltas[r:s]); ok && r-l == s-r {
				// If it did we found the loop
				return l, s - r
			}
		}
	}
}

func repeats(search []int, pattern []int) (bool, int) {
	for i := len(search) - (2 * len(pattern)); i > 0; i-- {
		found := true
		for j, v := range pattern {
			if search[i+j] != v {
				found = false
				break
			}
		}
		if found {
			return true, i
		}
	}
	return false, -1
}

func solve(g gas, head, delta, turns int) int {
	yhead := (&tunnel{tiles: make(map[xy]tile), ymax: -1}).Simulate(head, g)
	ydelta := (&tunnel{tiles: make(map[xy]tile), ymax: -1}).Simulate(head+delta, g) - yhead
	ytail := (&tunnel{tiles: make(map[xy]tile), ymax: -1}).Simulate(head+((turns-head)%delta), g)

	return (ydelta * ((turns - head) / delta)) + ytail
}

func (t *tunnel) Tick(g gas, turn int) int {
	ymax := t.ymax
	pos := xy{2, t.ymax + 4}
	r := Get(rocks, turn).(rock)
	for {
		w := Get(g, t.i)
		switch w.(byte) {
		case '<':
			if !t.CheckCollisions(xy{pos.x - 1, pos.y}, r) {
				pos.x--
			}
		case '>':
			if !t.CheckCollisions(xy{pos.x + 1, pos.y}, r) {
				pos.x++
			}
		}
		t.i++
		// Move vertically
		if t.CheckCollisions(xy{pos.x, pos.y - 1}, r) {
			// If we've collided vertically we're done so turn the rock into solid material in the tunnel
			for _, part := range r {
				ppos := pos.Add(part)
				t.tiles[ppos] = solid
				// Check if ymax has increased
				if ppos.y > ymax {
					ymax = ppos.y
				}
			}
			break
		}
		pos.y--
	}
	delta := ymax - t.ymax
	t.ymax = ymax
	return delta
}

func (t *tunnel) Simulate(turns int, g gas) int {
	for turn := 0; turn < turns; turn++ {
		t.Tick(g, turn)
	}
	return t.ymax
}

type tunnel struct {
	tiles map[xy]tile
	ymax  int
	i     int
}

func (t tunnel) CheckCollisions(p xy, r rock) bool {
	for _, part := range r {
		if t.Occupied(p.Add(part)) {
			return true
		}
	}
	return false
}

func (t tunnel) Occupied(p xy) bool {
	// Check fallen rocks
	if this, ok := t.tiles[p]; ok && this == solid {
		return true
	}
	// Check the tunnel edges
	if p.x < 0 || p.x > 6 || p.y < 0 {
		return true
	}
	return false
}

func (t tunnel) String() string {
	var b strings.Builder
	for y := t.ymax; y >= 0; y-- {
		for x := 0; x < 7; x++ {
			if this, ok := t.tiles[xy{x, y}]; ok && this == solid {
				b.WriteRune('#')
				continue
			}
			b.WriteRune('.')
		}
		b.WriteRune('\n')
	}
	return b.String()
}

type gas []byte

func Get[T any](s []T, i int) any {
	return s[i%(len(s))]
}

type xy struct {
	x, y int
}

func (x xy) Add(y xy) xy {
	return xy{x.x + y.x, x.y + y.y}
}

type rock []xy

var (
	hline  = rock{{0, 0}, {1, 0}, {2, 0}, {3, 0}}
	plus   = rock{{1, 0}, {0, 1}, {1, 1}, {2, 1}, {1, 2}}
	bl     = rock{{0, 0}, {1, 0}, {2, 0}, {2, 1}, {2, 2}}
	vline  = rock{{0, 0}, {0, 1}, {0, 2}, {0, 3}}
	square = rock{{0, 0}, {1, 0}, {0, 1}, {1, 1}}

	rocks = []rock{hline, plus, bl, vline, square}
)

type tile int

const (
	air tile = iota
	solid
)
