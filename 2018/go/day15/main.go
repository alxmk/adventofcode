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
	data, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatalln("Failed to read input", err)
	}

	round, hp, score, _ := solve(string(data), 3)

	fmt.Println(round, hp, score)

	for i := 4; i < 1000; i++ {
		round, hp, score, elfcasualties := solve(string(data), i)
		fmt.Println(i, round, hp, score, elfcasualties)
		if elfcasualties == 0 {
			break
		}
	}
}

func solve(input string, attack int) (int, int, int, int) {
	m := load(strings.Split(input, "\n"), attack)

	// fmt.Println("Start:")
	// m.Print()

	initialElves := len(m.elves)

	round := 1
	for {
		if complete, endedEarly := m.Tick(); complete {
			if endedEarly {
				round--
			}
			// m.Print()
			break
		}
		// fmt.Printf("('Round', %d, 'sum', %d)\n", round, m.SumRemainingHP())
		// m.Print()
		// fmt.Println("")
		round++

		// if round == 21 {
		// 	break
		// }
	}

	hp := m.SumRemainingHP()

	return round, hp, hp * round, initialElves - len(m.elves)
}

type grid struct {
	tiles   [][]bool
	elves   []*unit
	goblins []*unit
}

func (g *grid) SumRemainingHP() int {
	var sum int
	for _, e := range g.elves {
		sum += e.hp
	}
	for _, g := range g.goblins {
		sum += g.hp
	}
	return sum
}

func (g *grid) Print() {
	fmt.Print("  ")
	for x := range g.tiles {
		fmt.Print(x % 10)
	}
	fmt.Print("\n")
	for y := range g.tiles[0] {
		fmt.Printf("%d ", y%10)
	row:
		for x := range g.tiles {
			for _, e := range g.elves {
				if e.x == x && e.y == y {
					fmt.Print("E")
					continue row
				}
			}
			for _, g := range g.goblins {
				if g.x == x && g.y == y {
					fmt.Print("G")
					continue row
				}
			}
			if g.tiles[x][y] {
				fmt.Print(".")
			} else {
				fmt.Print("#")
			}
		}
		fmt.Print("\n")
	}
}

func (g *grid) Tick() (bool, bool) {
	var units []*unit
	units = append(units, g.goblins...)
	units = append(units, g.elves...)
	sort.Sort(readingOrder(units))
	sort.Sort(readingOrder(g.goblins))
	sort.Sort(readingOrder(g.elves))

	var deadThisTick []*unit
unitloop:
	for _, u := range units {
		// Don't process anything for stuff that died previously this tick
		// fmt.Println("Dead this loop:", len(deadThisTick))
		for _, d := range deadThisTick {
			if u == d {
				// fmt.Println("Skipping", u.kind, i)
				continue unitloop
			}
		}
		var opponents []*unit
		switch u.kind {
		case "elf":
			opponents = g.goblins
		case "goblin":
			opponents = g.elves
		}

		if len(opponents) == 0 {
			return true, true
		}

		// fmt.Println("Unit moving:", u.kind)

		self := tile{x: u.x, y: u.y}
		// fmt.Println("Unit is at:", self.x, self.y)

		_, next := g.NextMove(self, opponents)

		// Moving
		if next != nil {
			// fmt.Println(u.kind, u.x, u.y, "moving to", next.x, next.y, "en route to", dest.x, dest.y)
			// Make the move
			// fmt.Println(u.kind, i, "moving to", next.x, next.y)
			u.x = next.x
			u.y = next.y
		} else {
			// fmt.Println(u.kind, i, "chilling at", u.x, u.y)
		}

		// Check for attacks
		self = tile{x: u.x, y: u.y}
		attacked, unitAttacked, killed := g.Attack(u, opponents)
		if attacked {
			// fmt.Println(u.kind, u.x, u.y, "attacked", unitAttacked.kind, unitAttacked.x, unitAttacked.y)
		}

		if killed {
			// fmt.Println("Killed", unitAttacked.kind)
			deadThisTick = append(deadThisTick, unitAttacked)

			for i, o := range opponents {
				if o == unitAttacked {
					// fmt.Println("Removing", toAttack.kind, i)
					opponents = append(opponents[:i], opponents[i+1:]...)
					break
				}
			}
			// fmt.Println(len(opponents), unitAttacked.kind, "left")
		}

		switch u.kind {
		case "elf":
			g.goblins = opponents
		case "goblin":
			g.elves = opponents
		}
	}

	// fmt.Println(len(g.elves), "elves left;", len(g.goblins), "goblins left")
	// sort.Sort(readingOrder(units))
	// for _, u := range units {
	// 	fmt.Printf("('%s', %d, %d, %d)\n", strings.ToUpper(string(u.kind[0])), u.hp, u.x, u.y)
	// }
	// printUnits(g.elves)
	// printUnits(g.goblins)

	// If everything on one side is dead then it's over
	return len(g.elves) == 0 || len(g.goblins) == 0, false
}

func (g *grid) Attack(u *unit, opponents []*unit) (bool, *unit, bool) {
	self := tile{x: u.x, y: u.y}
	adjacent := self.Next()
	var inRange []*unit
	for _, o := range opponents {
		oTile := tile{x: o.x, y: o.y}
		for _, a := range adjacent {
			// An opponent is adjacent, so we attack
			if a == oTile {
				inRange = append(inRange, o)
			}
		}
	}
	var toAttack *unit
	minHP := math.MaxInt32
	for _, candidate := range inRange {
		if candidate.hp < minHP {
			minHP = candidate.hp
			toAttack = candidate
		}
	}
	var attacked bool
	var killed bool
	if toAttack != nil {
		// fmt.Println(u.kind, i, "attacked", toAttack.kind, "at", toAttack.x, toAttack.y)
		attacked = true
		if killed = u.Attack(toAttack); killed {
			// fmt.Println(u.kind, "at", u.x, u.y, "killed", toAttack.kind, "at", toAttack.x, toAttack.y)
		}
	}

	return attacked, toAttack, killed
}

func (g *grid) NextMove(self tile, opponents []*unit) (*tile, *tile) {
	var dests []tile
	adjacent := self.Next()
	nexts := make(map[tile]tile)
	routes := make(map[tile][]tile)

	// var target tile
	shortest := math.MaxInt32
	for _, o := range opponents {
		oTile := tile{x: o.x, y: o.y}
		for _, a := range adjacent {
			// An opponent is adjacent, so we will attack rather than move (processed later)
			if a == oTile {
				// fmt.Println(u.kind, i, "found adjacent enemy, not moving")
				return nil, nil
			}
		}
		for _, t := range oTile.Next() {
			// fmt.Println("Pathing:", t.x, t.y)
			if p, _ := g.Passable(t.x, t.y); p {
				if closed, n, ok := g.FindPath(self, t); ok {
					if len(closed) == shortest {
						dests = append(dests, t)
						nexts[t] = n
						routes[t] = closed
					}
					if len(closed) < shortest {
						dests = []tile{t}
						shortest = len(closed)
						nexts[t] = n
						routes[t] = closed
					}
				}
			}
		}
	}

	if len(dests) == 0 {
		return nil, nil
	}

	sort.Sort(tileOrder(dests))
	dest := dests[0]

	var tshortest tile
	short := shortest

	var found bool

	// Found destination, now find which move to make
	for _, t := range adjacent {
		if p, _ := g.Passable(t.x, t.y); !p {
			continue
		}
		if t == dest {
			// fmt.Println("Weird condition")
			return &dest, &dest
		}
		if route, _, ok := g.FindPath(t, dest); ok {
			// fmt.Printf("From %d, %d: %#v %d\n", t.x, t.y, route, len(route))
			if len(route) < short {
				found = true
				tshortest = t
				short = len(route)
			}
		}
	}

	if found {
		return &dest, &tshortest
	}

	fmt.Println("Failed to route", self.x, self.y, "=>", dest.x, dest.y)
	fmt.Printf("%#v\n", routes[dest])
	for _, t := range adjacent {
		if p, _ := g.Passable(t.x, t.y); !p {
			log.Println(t.x, t.y, "not passable")
			continue
		}
		if t == dest {
			// fmt.Println("Weird condition")
			return &dest, &dest
		}
		if route, _, ok := g.FindPath(t, dest); ok {
			fmt.Println(len(route), "via", t.x, t.y, "which is longer than", shortest, "via", nexts[dest].x, nexts[dest].y)
			fmt.Printf("%#v\n", route)
			if len(route) == shortest-1 {
				return &dest, &t
			}
		}
	}
	panic("Fiddlesticks")
}

type tile struct {
	x int
	y int
}

type tileOrder []tile

func (a tileOrder) Len() int      { return len(a) }
func (a tileOrder) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a tileOrder) Less(i, j int) bool {
	if a[i].y == a[j].y {
		return a[i].x < a[j].x
	}
	return a[i].y < a[j].y
}

func (t tile) Next() []tile {
	return []tile{tile{x: t.x, y: t.y - 1}, tile{x: t.x - 1, y: t.y}, tile{x: t.x + 1, y: t.y}, tile{x: t.x, y: t.y + 1}}
}

func (t tile) ReverseNext() []tile {
	return []tile{tile{x: t.x, y: t.y + 1}, tile{x: t.x + 1, y: t.y}, tile{x: t.x - 1, y: t.y}, tile{x: t.x, y: t.y - 1}}
}

func (t tile) Heuristic(to tile) int {
	x := t.x - to.x
	if x < 0 {
		x = x * -1
	}

	y := t.y - to.y
	if y < 0 {
		y = y * -1
	}

	return x + y
}

type path map[tile]entry

type entry struct {
	score int
	g     int
}

func (p path) Next(current tile) (tile, bool) {
	minScore := math.MaxInt32
	var gmin int
	var minTile tile
	var found bool
	for _, t := range current.ReverseNext() {
		if entry, ok := p[t]; ok {
			// fmt.Println("Could pick", t.x, t.y)
			if entry.score < minScore || (entry.score == minScore && entry.g < gmin) {
				gmin = entry.g
				minScore = entry.score
				minTile = t
				found = true
			}
		}
	}

	return minTile, found
}

func (p path) LowestScore() (tile, entry) {
	minScore := math.MaxInt32
	var minTile tile
	var minEntry entry
	for tile, entry := range p {
		if entry.score < minScore {
			minScore = entry.score
			minEntry = entry
			minTile = tile
		}
		if entry.score == minScore {
			if minTile.y == tile.y {
				if minTile.x > tile.x {
					minScore = entry.score
					minEntry = entry
					minTile = tile
					continue
				}
			}
			if minTile.y > tile.y {
				minScore = entry.score
				minEntry = entry
				minTile = tile
				continue
			}
		}
	}
	return minTile, minEntry
}

func (p path) Print() {
	for tile, score := range p {
		fmt.Printf("%d,%d => %d; ", tile.x, tile.y, score)
	}
	fmt.Print("\n")
}

func (p path) Path(from, to tile) []tile {
	// Work backwards from the destination
	current := to
	resultant := []tile{current}

	// fmt.Println(from.x, from.y, "=>", to.x, to.y)
	// fmt.Printf("%#v\n", p)

	var i int
	var ok bool
	for current != from && i < 10000 {
		delete(p, current)
		// fmt.Printf("%#v\n", current)
		if current, ok = p.Next(current); ok {
			if current == from {
				return resultant
			}
			resultant = append(resultant, current)
		} else {
			// fmt.Printf("FUCK %#v\n", current)
			resultant = resultant[:len(resultant)-1]
			current = resultant[len(resultant)-1]
		}
		i++
	}
	if i == 10000 {
		fmt.Println("Pathing from", from.x, from.y, "=>", to.x, to.y)
		fmt.Printf("Resultant: %#v\n", resultant)
		fmt.Printf("P: %#v\n", p)
		panic("ffs")
	}

	return resultant
}

func (gr *grid) FindPath(from, to tile) ([]tile, tile, bool) {
	closed := path{from: entry{score: 0, g: 0}}
	open := path{}

	current := from
	g := 0

	var pathFound bool

	// fmt.Println("Pathing:", from.x, from.y, "->", to.x, to.y)

	for {
		// fmt.Println("Current:", current.x, current.y)
		// closed.Print()
		// open.Print()
		g = g + 1
		for _, t := range current.Next() {
			// fmt.Println(current.x, current.y, "->", t.x, t.y)
			if _, ok := closed[t]; ok {
				// fmt.Println(t.x, t.y, closed)
				continue
			}
			passable, err := gr.Passable(t.x, t.y)
			if err != nil {
				fmt.Printf("Pathing: %d,%d => %d,%d\n", from.x, from.y, to.x, to.y)
				fmt.Println("Current:", current.x, current.y)
				fmt.Print("Closed: ")
				closed.Print()
				fmt.Print("Open: ")
				open.Print()
				log.Fatalln("Failed to check passability:", err)
			}
			if !passable {
				// fmt.Println(t.x, t.y, "impassable")
				continue
			}
			h := t.Heuristic(to)
			if _, ok := open[t]; ok {
				if open[t].score > h+g {
					open[t] = entry{score: h + g, g: g}
				}
			} else {
				open[t] = entry{score: h + g, g: g}
			}

			if h == 0 {
				pathFound = true
			}
		}

		next, entry := open.LowestScore()
		// no score
		if entry.score == 0 && entry.g == 0 {
			break
		}
		delete(open, next)
		closed[next] = entry

		current = next
		g = entry.g

		if pathFound {
			break
		}
	}

	if !pathFound {
		var empty tile
		return nil, empty, false
	}

	// fmt.Printf("Closed: %#v\n", closed)

	route := closed.Path(from, to)
	// fmt.Printf("Route length: %d, route: %#v\n", len(route), route)
	return route, route[len(route)-1], true
}

func load(lines []string, attack int) *grid {
	var goblins []*unit
	var elves []*unit
	m := make([][]bool, len(lines[0]))

	for y, line := range lines {
		for x, r := range line {
			if y == 0 {
				m[x] = make([]bool, len(lines))
			}

			switch r {
			case '#':
				m[x][y] = false
			case 'E':
				elves = append(elves, newUnit(x, y, "elf", attack))
				m[x][y] = true
			case 'G':
				goblins = append(goblins, newUnit(x, y, "goblin", 3))
				m[x][y] = true
			default:
				m[x][y] = true
			}
		}
	}

	return &grid{goblins: goblins, elves: elves, tiles: m}
}

func (g *grid) Passable(x, y int) (bool, error) {
	if x < 0 || x > len(g.tiles) || y < 0 || y > len(g.tiles[x]) {
		return false, fmt.Errorf("Asked whether %d, %d is passable", x, y)
	}

	if !g.tiles[x][y] {
		return false, nil
	}

	for _, gob := range g.goblins {
		if gob.x == x && gob.y == y {
			return false, nil
		}
	}

	for _, elf := range g.elves {
		if elf.x == x && elf.y == y {
			return false, nil
		}
	}

	return true, nil
}

type unit struct {
	attack int
	hp     int
	x      int
	y      int
	kind   string
}

func newUnit(x, y int, kind string, attack int) *unit {
	return &unit{
		attack: attack,
		hp:     200,
		x:      x,
		y:      y,
		kind:   kind,
	}
}

func printUnits(units []*unit) {
	sort.Sort(readingOrder(units))
	for i, u := range units {
		fmt.Printf("%s %d: %d, %d (%d)\n", string(u.kind[0]), i, u.x, u.y, u.hp)
	}
}

func (u *unit) Print() string {
	return fmt.Sprintf("%s: %d, %d (%d)\n", string(u.kind[0]), u.x, u.y, u.hp)
}

func (u *unit) Attack(enemy *unit) bool {
	enemy.hp -= u.attack
	return enemy.hp <= 0
}

type readingOrder []*unit

func (a readingOrder) Len() int      { return len(a) }
func (a readingOrder) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a readingOrder) Less(i, j int) bool {
	if a[i].y == a[j].y {
		return a[i].x < a[j].x
	}
	return a[i].y < a[j].y
}
