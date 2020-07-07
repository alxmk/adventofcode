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
		log.Fatalln("Failed to read input file:", err)
	}

	w := parseWorld(string(data))

	w.TrimDeadEnds()

	fmt.Println(w)

	log.Println("Part one:", newMonkey(w).Go())
}

type monkey struct {
	environment *world
	location    coordinate
	target      *coordinate
	keys        map[rune]struct{}
	travelled   int64
}

func newMonkey(w *world) *monkey {
	return &monkey{
		environment: w.Copy(),
		location:    w.entrance,
		keys:        make(map[rune]struct{}),
	}
}

func (m *monkey) Fork(target coordinate) *monkey {
	keysCopy := make(map[rune]struct{})

	for k := range m.keys {
		keysCopy[k] = struct{}{}
	}

	return &monkey{
		environment: m.environment.Copy(),
		location:    m.location,
		target:      &target,
		keys:        keysCopy,
		travelled:   m.travelled,
	}
}

func (m *monkey) Keys() string {
	var keys []string
	for k := range m.keys {
		keys = append(keys, string(k))
	}
	sort.Strings(keys)
	return strings.Join(keys, "")
}

var (
	monkeys int
	cache   = make(map[rune]map[string]int64) // cache maps current key to the shortest path to remaining keys
)

func (m *monkey) Go() int64 {
	monkeys++
	defer func() {
		monkeys--
	}()
	var collectedKey rune
	if m.target != nil {
		if len, ok := m.environment.Astar(m.location, *m.target, m.keys); ok {
			m.travelled += len
			m.location = *m.target
			if r := m.environment.tiles[m.location]; isKey(r) {
				collectedKey = r
				m.keys[r] = struct{}{}
				delete(m.environment.keys, r)
			}
		}
	}

	targets := m.environment.PathableKeys(m.location, m.keys)
	switch len(targets) {
	case 0:
		return m.travelled
	case 1:
		m.target = &targets[0]
		return m.Go()
	default:
		// if c, ok := cache[collectedKey]; ok {
		// 	if i, ok := c[m.environment.RemainingKeys()]; ok {
		// 		log.Println("Cache hit", m.Keys(), string(collectedKey), m.environment.RemainingKeys(), i)
		// 		return i
		// 	}
		// }
		results := []int{math.MaxInt64}
		for _, t := range targets {
			if r := int(m.Fork(t).Go()); r < math.MaxInt64 {
				results = append(results, r)
			}
		}
		sort.Ints(results)
		result := int64(results[0])

		if _, ok := cache[collectedKey]; !ok {
			cache[collectedKey] = make(map[string]int64)
		}
		cache[collectedKey][m.environment.RemainingKeys()] = result

		return result
	}
}

func isKey(r rune) bool {
	return r >= 'a' && r <= 'z'
}

func isDoor(r rune) bool {
	return r >= 'A' && r <= 'Z'
}

type coordinate struct {
	x, y int64
}

func (c coordinate) String() string {
	return fmt.Sprintf("(%d, %d)", c.x, c.y)
}

func (c coordinate) ManhattanDistance(to coordinate) int64 {
	return absint(to.x-c.x) + absint(to.y-c.y)
}

func (c coordinate) Next(dir coordinate) coordinate {
	return coordinate{c.x + dir.x, c.y + dir.y}
}

var (
	north = coordinate{0, -1}
	south = coordinate{0, 1}
	west  = coordinate{-1, 0}
	east  = coordinate{1, 0}

	directions = []coordinate{north, south, west, east}

	wall     = '#'
	entrance = '@'
	floor    = '.'
)

func absint(i int64) int64 {
	if i < 0 {
		return i * -1
	}
	return i
}

type world struct {
	tiles       map[coordinate]rune
	xmax, ymax  int64
	doors, keys map[rune]coordinate
	entrance    coordinate
}

func newWorld() *world {
	return &world{tiles: make(map[coordinate]rune), doors: make(map[rune]coordinate), keys: make(map[rune]coordinate)}
}

func (w *world) Copy() *world {
	copy := newWorld()

	copy.xmax, copy.ymax, copy.entrance = w.xmax, w.ymax, w.entrance

	for k, v := range w.tiles {
		copy.tiles[k] = v
	}

	for k, v := range w.doors {
		copy.doors[k] = v
	}

	for k, v := range w.keys {
		copy.keys[k] = v
	}

	return copy
}

func parseWorld(input string) *world {
	w := newWorld()
	var x, y int64
	for _, r := range input {
		if x > w.xmax {
			w.xmax = x
		}
		if y > w.ymax {
			w.ymax = y
		}
		switch r {
		case '\n':
			y++
			x = 0
		case '#', '.':
			w.tiles[coordinate{x, y}] = r
			x++
		case '@':
			w.tiles[coordinate{x, y}] = r
			w.entrance = coordinate{x, y}
			x++
		default:
			w.tiles[coordinate{x, y}] = r
			if r >= 'a' {
				w.keys[r] = coordinate{x, y}
			} else {
				w.doors[r] = coordinate{x, y}
			}
			x++
		}
	}
	return w
}

func (w world) String() string {
	var out strings.Builder
	for y := int64(0); y <= w.ymax; y++ {
		for x := int64(0); x <= w.xmax; x++ {
			out.WriteRune(w.tiles[coordinate{x, y}])
		}
		out.WriteRune('\n')
	}
	return out.String()
}

func (w *world) Astar(from, to coordinate, keys map[rune]struct{}) (int64, bool) {
	openSet := map[coordinate]struct{}{
		from: struct{}{},
	}

	cameFrom := make(map[coordinate]coordinate)

	gScore := defaultMaxMap{
		from: 0,
	}

	fScore := defaultMaxMap{
		from: gScore[from] + from.ManhattanDistance(from),
	}

	for len(openSet) != 0 {
		lowestScore := int64(math.MaxInt64)
		var current coordinate

		for c := range openSet {
			if score := fScore.Get(c); score < lowestScore {
				lowestScore = score
				current = c
			}
		}

		if current == to {
			return reconstructPath(cameFrom, current) - 1, true
		}

		delete(openSet, current)
		for _, d := range directions {
			neighbour := current.Next(d)
			tentativeG := gScore.Get(current)
			if tentativeG != math.MaxInt64 {
				tentativeG++
			}
			t := w.tiles[neighbour]
			if t == wall {
				tentativeG = math.MaxInt64
			}
			if _, haveKey := keys[t+('a'-'A')]; isDoor(t) && !haveKey {
				tentativeG = math.MaxInt64
			}

			if tentativeG < gScore.Get(neighbour) {
				cameFrom[neighbour] = current
				gScore[neighbour] = tentativeG
				fScore[neighbour] = tentativeG + from.ManhattanDistance(neighbour)
				openSet[neighbour] = struct{}{}
			}
		}
	}

	return -1, false
}

type defaultMaxMap map[coordinate]int64

func (d defaultMaxMap) Get(c coordinate) int64 {
	if v, ok := d[c]; ok {
		return v
	}
	return math.MaxInt64
}

func reconstructPath(cameFrom map[coordinate]coordinate, current coordinate) int64 {
	totalPath := []coordinate{current}
	for current, ok := cameFrom[current]; ok; current, ok = cameFrom[current] {
		totalPath = append(totalPath, current)
	}
	return int64(len(totalPath))
}

func (w *world) PathableKeys(from coordinate, haveKeys map[rune]struct{}) []coordinate {
	var pathable []coordinate
	for _, c := range w.keys {
		if _, ok := w.Astar(from, c, haveKeys); ok {
			pathable = append(pathable, c)
			// log.Println(string(k), "pathable", c, "from", from)
		} else {
			// log.Println(string(k), "unpathable", c, "from", from)
		}
	}
	return pathable
}

func (w *world) TrimDeadEnds() {
	for {
		var diff bool
		for c, r := range w.tiles {
			if r != '.' {
				continue
			}
			var walledNeighbourCount int
			for _, dir := range directions {
				if n, ok := w.tiles[c.Next(dir)]; ok && n == '#' {
					walledNeighbourCount++
				}
			}
			if walledNeighbourCount == 3 {
				w.tiles[c] = '#'
				diff = true
			}
		}
		if !diff {
			break
		}
	}
}

func (w *world) RemainingKeys() string {
	var keys []string
	for r := range w.keys {
		keys = append(keys, string(r))
	}
	sort.Strings(keys)
	return strings.Join(keys, "")
}
