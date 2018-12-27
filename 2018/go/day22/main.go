package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"os"
	"strings"
)

func main() {
	// depth: 10914
	// target: 9,739
	generated, score := solve(10914, 9, 739)
	fmt.Println(generated)
	fmt.Println("Part one:", score)

	length := findpath(10914, 9, 739)

	fmt.Println("Part two:", length)
}

func findpath(depth, tx, ty int) int {
	w := generateWorld(depth, tx, ty, 40, 150)

	// fmt.Println(w.String())

	return w.Path(0, 0, tx, ty)
}

func solve(depth, tx, ty int) (string, int) {
	w := generateWorld(depth, tx, ty, 0, 0)

	return w.String(), w.Risk()
}

func generateWorld(depth, tx, ty, xoffset, yoffset int) world {
	w := make(world, tx+1+xoffset)

	for x := 0; x < len(w); x++ {
		w[x] = make([]*region, ty+1+yoffset)
	}

	w.Populate(depth, tx, ty)

	return w
}

type world [][]*region

func (w world) Populate(depth, tx, ty int) {
	for x := range w {
		for y := range w[x] {
			w[x][y] = w.regionFor(x, y, depth, tx, ty)
		}
	}
}

func (w world) regionFor(x, y, depth, tx, ty int) *region {
	var g, e int
	switch {
	case (x == 0 && y == 0) || (x == tx && y == ty):
		g = 0
	case y == 0:
		g = x * 16807
	case x == 0:
		g = y * 48271
	default:
		g = w[x-1][y].erosion * w[x][y-1].erosion
	}
	e = (g + depth) % 20183
	var t kind
	switch e % 3 {
	case 0:
		t = rocky
	case 1:
		t = wet
	case 2:
		t = narrow
	default:
		panic("wah")
	}
	return &region{terrain: t, geologic: g, erosion: e}
}

func (w world) String() string {
	var out strings.Builder

	for y := range w[0] {
		for x := range w {
			switch w[x][y].terrain {
			case rocky:
				out.WriteString(".")
			case wet:
				out.WriteString("=")
			case narrow:
				out.WriteString("|")
			}
		}
		out.WriteString("\n")
	}

	return out.String()
}

func (w world) Risk() int {
	var risk int

	for y := range w[0] {
		for x := range w {
			risk += int(w[x][y].terrain)
		}
	}

	return risk
}

func (w world) Path(sx, sy, dx, dy int) int {
	visited := make(nodes, len(w))
	unvisited := make(nodes, len(w))
	for x := range w {
		visited[x] = make([]*node, len(w[x]))
		unvisited[x] = make([]*node, len(w[x]))
		for y := range w[x] {
			unvisited[x][y] = newNode()
		}
	}

	initial := coord{x: 0, y: 0}
	unvisited[initial.x][initial.y] = &node{tentative: 0, equip: map[loadout]struct{}{torch: struct{}{}}}

	current := initial
	var ok bool
	// for i := 0; i < 50; i++ {
	for {
		w.Distances(visited, unvisited, current)

		if current, ok = unvisited.Next(); !ok {
			return 0
		}

		// fmt.Println(visited.String())

		if target := visited[dx][dy]; target != nil {
			ioutil.WriteFile("debug.out", []byte(visited.String()), os.ModePerm)
			ioutil.WriteFile("route.out", []byte(visited.Route(coord{x: sx, y: sy}, coord{x: dx, y: dy}, w)), os.ModePerm)
			if _, ok := target.equip[torch]; !ok {
				return target.tentative + 7
			}
			return target.tentative
		}
	}

	return -1
}

type nodes [][]*node

func (n nodes) Next() (coord, bool) {
	lowestTentative := math.MaxInt32
	var nextCoord coord
	var found bool

	for x := range n {
		for y := range n[x] {
			if u := n[x][y]; u != nil {
				if u.tentative < lowestTentative {
					found = true
					lowestTentative = u.tentative
					nextCoord.x = x
					nextCoord.y = y
				}
			}
		}
	}

	return nextCoord, found
}

func (n nodes) String() string {
	var out strings.Builder

	out.WriteString("  ")

	for x := range n {
		out.WriteString(fmt.Sprintf("%d     ", x%10))
	}

	out.WriteString("\n")

	for y := range n[0] {
		out.WriteString(fmt.Sprintf("%d ", y%10))
		for x := range n {
			node := n[x][y]
			if node == nil {
				out.WriteString("      ")
				continue
			}
			out.WriteString(fmt.Sprintf("%d", node.tentative))
			var count int
			for e := range node.equip {
				switch e {
				case torch:
					out.WriteString("t")
				case climbing:
					out.WriteString("c")
				case neither:
					out.WriteString("n")
				}
				count++
			}
			for i := 0; i < 3-count; i++ {
				out.WriteString(" ")
			}
			if node.tentative < 100 {
				out.WriteString(" ")
				if node.tentative < 10 {
					out.WriteString(" ")
				}
			}
		}
		out.WriteString("\n")
	}
	out.WriteString("\n")

	return out.String()
}

type routePoint struct {
	node  *node
	coord coord
}

func (n nodes) Route(from, to coord, w world) string {
	route := []routePoint{routePoint{node: n[from.x][from.y], coord: from}}

	current := from

	for current != to {
		lowest := math.MaxInt32
		var nextCoord coord
		for _, next := range current.Adjacent() {
			if !next.Valid(w) {
				continue
			}

			node := n[next.x][next.y]
			if node == nil {
				continue
			}

			if node.tentative < lowest {
				lowest = node.tentative
				nextCoord = next
			}
		}

		route = append(route, routePoint{node: n[nextCoord.x][nextCoord.y], coord: nextCoord})
		current = nextCoord
	}

	var out strings.Builder
	for i := len(route) - 1; i >= 0; i-- {
		r := route[i]
		out.WriteString(fmt.Sprintf("Path (%d,%d) cost %d equipped %d\n", r.coord.x, r.coord.y, r.node.tentative, r.node.equip))
	}

	return out.String()
}

func (w world) Distances(visited, unvisited [][]*node, current coord) {
	from := unvisited[current.x][current.y]

	// Visit all adjacent nodes to the current
	for _, c := range current.Adjacent() {
		if !c.Valid(w) {
			continue
		}

		// If they've been visited then ignore
		n := visited[c.x][c.y]
		if n != nil {
			continue
		}
		// Calculate the movement costs to the destination node
		dist, equip := w.calculateCost(current, c, from)

		// Update the unvisited list
		u := unvisited[c.x][c.y]
		if dist < u.tentative {
			u.tentative = dist
			u.equip = make(map[loadout]struct{})
			for _, e := range equip {
				u.equip[e] = struct{}{}
			}
		}
		if dist == u.tentative {
			for _, e := range equip {
				u.equip[e] = struct{}{}
			}
		}

		// if c.x == 3 && c.y == 1 {
		// 	fmt.Println(current.x, current.y, w[current.x][current.y], "=>", c.x, c.y, w[c.x][c.y], from, u, dist, equip)
		// }
	}

	// We've now visited the current node
	visited[current.x][current.y] = unvisited[current.x][current.y]

	unvisited[current.x][current.y] = nil
}

func (w world) calculateCost(current, next coord, previous *node) (int, []loadout) {
	terrainnow := w[current.x][current.y].terrain
	terrainnext := w[next.x][next.y].terrain

	// If we can keep using the same loadout then we just increment the distance and return
	options := make(map[loadout]int)
	for e := range previous.equip {
		if terrainnext.Valid(e) {
			options[e] = previous.tentative + 1
			continue
		}

		// Otherwise work out if we can switch to a valid one
		newLoadout := terrainnow.OtherValid(e)
		if terrainnext.Valid(newLoadout) {
			if _, ok := options[newLoadout]; ok {
				if options[newLoadout] < previous.tentative+8 {
					continue
				}
			}
			options[newLoadout] = previous.tentative + 8
			continue
		}

		// Impassable
	}

	minDist := math.MaxInt32
	var equipment []loadout
	for k, v := range options {
		if v < minDist {
			minDist = v
			equipment = []loadout{k}
		}
		if v == minDist {
			equipment = append(equipment, k)
		}
	}

	return minDist, equipment
}

type region struct {
	terrain  kind
	geologic int
	erosion  int
}

type kind int

const (
	rocky kind = iota
	wet
	narrow
)

// Returns whether the specified loadout is valid in the terrain
func (k kind) Valid(equip loadout) bool {
	switch k {
	case rocky:
		return equip == torch || equip == climbing
	case wet:
		return equip == climbing || equip == neither
	case narrow:
		return equip == neither || equip == torch
	}
	return false
}

// What is the other loadout we could use in this terrain
func (k kind) OtherValid(equip loadout) loadout {
	switch k {
	case rocky:
		if equip == torch {
			return climbing
		}
		return torch
	case wet:
		if equip == climbing {
			return neither
		}
		return climbing
	case narrow:
		if equip == neither {
			return torch
		}
		return neither
	}
	panic("bums")
}

type node struct {
	tentative int
	equip     map[loadout]struct{}
}

func newNode() *node {
	return &node{tentative: math.MaxInt32}
}

type coord struct {
	x, y int
}

func (c coord) Valid(w world) bool {
	if c.x < 0 || c.x >= len(w) || c.y < 0 || c.y >= len(w[0]) {
		return false
	}
	return true
}

func (c coord) Adjacent() []coord {
	return []coord{coord{x: c.x, y: c.y - 1}, coord{x: c.x + 1, y: c.y}, coord{x: c.x, y: c.y + 1}, coord{x: c.x - 1, y: c.y}}
}

type loadout int

const (
	unknown loadout = iota
	torch
	climbing
	neither
)
