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

	b, err := parseAsteroids(string(data))
	if err != nil {
		log.Fatalln("Failed to parse input as asteroid belt:", err)
	}

	best, maxInVision := b.getMaxInVision()
	log.Println("Part one:", maxInVision, "@", best)

	destroyed := b.FIRETHELASER(best)
	log.Println("Part two:", destroyed[199].x*100+destroyed[199].y)
}

type belt struct {
	asteroids  map[coordinate]struct{}
	xmax, ymax int
}

func (b belt) String() string {
	var out strings.Builder
	for y := 0; y <= b.ymax; y++ {
		for x := 0; x <= b.xmax; x++ {
			if _, ok := b.asteroids[coordinate{x, y}]; ok {
				out.WriteRune('#')
			} else {
				out.WriteRune('.')
			}
		}
		out.WriteRune('\n')
	}
	return out.String()
}

func (b belt) getMaxInVision() (coordinate, int) {
	var maxInVision int
	var best coordinate
	for a := range b.asteroids {
		if l := b.getNumInVision(a); l > maxInVision {
			maxInVision = l
			best = a
		}
	}
	return best, maxInVision
}

func (b belt) getNumInVision(from coordinate) int {
	return len(b.getInVision(from))
}

func (b belt) getInVision(from coordinate) visionMap {
	inVision := make(visionMap)
	for other := range b.asteroids {
		if from == other {
			continue
		}

		vec := from.vectorTo(other)
		newDirVec := true
		for c, v := range inVision {
			if vec.parallelTo(v) && !vec.oppositeTo(v) {
				newDirVec = false
				if c.Distance(from) > other.Distance(from) {
					delete(inVision, c)
					inVision[other] = vec
				}
				break
			}
		}
		if newDirVec {
			inVision[other] = vec
		}
	}
	return inVision
}

// returns the destruction order
func (b belt) FIRETHELASER(from coordinate) []coordinate {
	inv := b.getInVision(from)
	angleMap := make(map[float64]coordinate)
	var angleSlice []float64

	origin := vector{0, -1}
	for c, v := range inv {
		a := (v.angle() - origin.angle())
		if a < 0 {
			a += 2 * math.Pi // Adjust negative angles into positive space to make sort easier
		}
		angleSlice = append(angleSlice, a)
		angleMap[a] = c
	}

	sort.Sort(sort.Float64Slice(angleSlice))

	var orderedDestroyed []coordinate
	for _, a := range angleSlice {
		orderedDestroyed = append(orderedDestroyed, angleMap[a])
	}

	return orderedDestroyed
}

type visionMap map[coordinate]vector

func (v visionMap) String() string {
	var output []string
	for c := range v {
		output = append(output, c.String())
	}
	return strings.Join(output, "\n")
}

func newBelt() *belt {
	return &belt{asteroids: make(map[coordinate]struct{})}
}

func parseAsteroids(input string) (*belt, error) {
	b := newBelt()
	var x, y int
	for _, r := range input {
		switch r {
		case '.':
		case '#':
			b.asteroids[coordinate{x, y}] = struct{}{}
		case '\n':
			b.xmax = x - 1
			y++
			x = 0
			continue
		default:
			return nil, fmt.Errorf("unparseable character in input: %v", r)
		}
		x++
	}
	b.ymax = y
	return b, nil
}

type coordinate struct {
	x, y int
}

func (c coordinate) Distance(from coordinate) int {
	return absint(c.x-from.x) + absint(c.y-from.y)
}

func (c coordinate) String() string {
	return fmt.Sprintf("(%d, %d)", c.x, c.y)
}

func absint(i int) int {
	if i < 0 {
		return -1 * i
	}
	return i
}

func (c coordinate) vectorTo(other coordinate) vector {
	return vector{float64(other.x - c.x), float64(other.y - c.y)}
}

type vector struct {
	x, y float64
}

func (v vector) String() string {
	return fmt.Sprintf("V(%f, %f)", v.x, v.y)
}

func (v vector) parallelTo(other vector) bool {
	return v.y/other.y == v.x/other.x || (v.y == 0 && other.y == 0) || (v.x == 0 && other.x == 0)
}

func (v vector) oppositeTo(otherParallel vector) bool {
	return (v.x/otherParallel.x) < 0 || (v.x == 0 && otherParallel.x == 0 && (v.y/otherParallel.y < 0))
}

func (v vector) angle() float64 {
	return math.Atan2(v.y, v.x)
}
