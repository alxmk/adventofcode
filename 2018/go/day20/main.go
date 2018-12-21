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

	w, score, thousandPlus := solve(string(data))

	fmt.Println(w)
	fmt.Println(score)
	fmt.Println(thousandPlus)
}

func solve(input string) (string, int, int) {
	s := newScanner(input)

	start := &room{}
	origin := coord{}

	w := make(world)
	w[origin] = start

	w.Path(origin, start, s)

	return w.String(), start.ScoreRecursive(0), w.NumThousandPlus()
}

type scanner struct {
	data    []rune
	pointer int
}

func newScanner(data string) *scanner {
	return &scanner{
		data: []rune(data),
	}
}

func (s *scanner) Next() (rune, error) {
	s.pointer++

	if s.pointer >= len(s.data) {
		return '!', fmt.Errorf("end of data")
	}

	return s.data[s.pointer], nil
}

type coord struct {
	x, y int
}

type byMap []coord

func (s byMap) Less(i, j int) bool {
	if s[i].y < s[j].y {
		return true
	}
	if s[i].y == s[j].y {
		return s[i].x < s[j].x
	}
	return false
}
func (s byMap) Swap(i, j int) { s[i], s[j] = s[j], s[i] }
func (s byMap) Len() int      { return len(s) }

type world map[coord]*room

type room struct {
	north, east, south, west *room
	score                    int
	max                      int
	scored                   bool
}

func (r *room) Attached() []*room {
	return []*room{r.north, r.east, r.south, r.west}
}

func (r *room) ScoreRecursive(current int) int {
	if r.scored {
		return r.max
	}

	// fmt.Println("Scoring", current)

	r.scored = true

	r.score = current
	r.max = r.score

	for _, attached := range r.Attached() {
		if attached != nil {
			if s := attached.ScoreRecursive(current + 1); s > r.max {
				r.max = s
			}
		}
	}

	// fmt.Println("rmax", r.max)

	return r.max
}

func (w world) NumThousandPlus() int {
	var num int
	for _, r := range w {
		if r.score >= 1000 {
			num++
		}
	}
	return num
}

func (w world) Follow(c coord, r *room, dir rune) (coord, *room) {
	switch dir {
	case 'N':
		newc := coord{x: c.x, y: c.y + 1}
		if r.north == nil {
			if n, ok := w[newc]; ok {
				r.north = n
				n.south = r
			} else {
				r.north = &room{south: r}
				w[newc] = r.north
			}
		}
		return newc, r.north
	case 'E':
		newc := coord{x: c.x + 1, y: c.y}
		if r.east == nil {
			if e, ok := w[newc]; ok {
				r.east = e
				e.west = r
			} else {
				r.east = &room{west: r}
				w[newc] = r.east
			}
		}
		return newc, r.east
	case 'S':
		newc := coord{x: c.x, y: c.y - 1}
		if r.south == nil {
			if s, ok := w[newc]; ok {
				r.south = s
				s.north = r
			} else {
				r.south = &room{north: r}
				w[newc] = r.south
			}
		}
		return newc, r.south
	case 'W':
		newc := coord{x: c.x - 1, y: c.y}
		if r.west == nil {
			if v, ok := w[newc]; ok {
				r.west = v
				v.east = r
			} else {
				r.west = &room{east: r}
				w[newc] = r.west
			}
		}
		return newc, r.west
	}

	panic("Shiet")
}

func (w world) Path(c coord, starting *room, s *scanner) bool {
	// fmt.Println("Pathing from", c.x, c.y)
	current := starting

	for r, err := s.Next(); err == nil; r, err = s.Next() {
		switch r {
		case '^':
			continue
		case '$', ')':
			return false
		case '(':
			// Recurse through each fork
			for w.Path(c, current, s) {
			}
			// fmt.Println("World size", len(w))
		case '|':
			return true
		default:
			// fmt.Println("Following", string(r))
			c, current = w.Follow(c, current, r)
		}
	}

	return false
}

func (w world) String() string {
	var keys byMap
	var maxx, maxy int
	minx, miny := math.MaxInt32, math.MaxInt32
	for k := range w {
		keys = append(keys, k)
		if k.x > maxx {
			maxx = k.x
		}
		if k.x < minx {
			minx = k.x
		}
		if k.y > maxy {
			maxy = k.y
		}
		if k.y < miny {
			miny = k.y
		}
	}
	sort.Sort(keys)

	var out strings.Builder
	for x := minx; x <= maxx; x++ {
		out.WriteString("##")
	}
	out.WriteString("#\n")
	for y := maxy; y >= miny; y-- {
		for x := minx; x <= maxx; x++ {
			if x == minx {
				out.WriteString("#")
			}
			if r, ok := w[coord{x: x, y: y}]; ok {
				if x == 0 && y == 0 {
					out.WriteString("X")
				} else {
					out.WriteString(".")
				}
				if r.east != nil {
					out.WriteString("|")
					continue
				}
				out.WriteString("#")
				continue
			}
			out.WriteString("##")
		}
		out.WriteString("\n")
		for x := minx; x <= maxx; x++ {
			if x == minx {
				out.WriteString("#")
			}
			if r, ok := w[coord{x: x, y: y}]; ok {
				if r.south != nil {
					out.WriteString("-#")
					continue
				}
				out.WriteString("##")
				continue
			}
			out.WriteString("##")
		}
		out.WriteString("\n")
	}

	return out.String()
}
