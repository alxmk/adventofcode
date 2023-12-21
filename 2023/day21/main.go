package main

import (
	"bytes"
	"log"
	"math"
	"math/big"
	"os"
)

func main() {
	data, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatalln("Failed to read input:", err)
	}

	log.Println("Part one:", partOne(parseGarden(data)))
	log.Println("Part two:", partTwo(parseGarden(data)))
}

func partOne(g infiniteGarden, start xy) int64 {
	return g.Steps(start, 64)[64]
}

func partTwo(g infiniteGarden, start xy) string {
	offset := g.size.x - start.x - 1
	results := g.Steps(start, offset+(2*g.size.x))

	y0 := results[offset]
	y1 := results[offset+g.size.x]
	y2 := results[offset+g.size.x*2]

	n := int64(26501365 / g.size.x)

	y := big.NewInt(y0 + (y1-y0)*n)
	y.Add(y, big.NewInt(0).Mul(big.NewInt(n*(n-1)/2), big.NewInt(y2-y1-y1+y0)))

	return y.String()
}

func parseGarden(input []byte) (infiniteGarden, xy) {
	i := infiniteGarden{g: make(garden)}
	var start xy
	for y, line := range bytes.Split(input, []byte{'\n'}) {
		for x, r := range line {
			i.g[xy{x, y}] = r
			if r == 'S' {
				start = xy{x, y}
			}
			i.size.x = x
		}
		i.size.y = y
	}
	i.size = xy{i.size.x + 1, i.size.y + 1}
	return i, start
}

var (
	north = xy{0, -1}
	south = xy{0, 1}
	west  = xy{-1, 0}
	east  = xy{1, 0}

	cardinal = []xy{north, south, west, east}
)

type xy struct {
	x, y int
}

func (x xy) Next(dir xy) xy {
	return xy{x.x + dir.x, x.y + dir.y}
}

type infiniteGarden struct {
	g    garden
	size xy
}

func (i infiniteGarden) Get(a xy) byte {
	resolved := xy{a.x % (i.size.x), a.y % (i.size.y)}
	if resolved.x < 0 {
		resolved.x += i.size.x
	}
	if resolved.y < 0 {
		resolved.y += i.size.y
	}
	return i.g[resolved]
}

type garden map[xy]byte

func (i infiniteGarden) Score(x xy) int {
	switch i.Get(x) {
	case '#':
		return math.MaxInt
	default:
		return 1
	}
}

func (i infiniteGarden) Tick(n int, locations map[xy]struct{}) map[xy]struct{} {
	next := make(map[xy]struct{})
	for l := range locations {
		for _, d := range cardinal {
			n := l.Next(d)
			if i.Score(n) == math.MaxInt {
				continue
			}
			next[n] = struct{}{}
		}
	}
	return next
}

func (i infiniteGarden) Steps(start xy, count int) []int64 {
	locations := map[xy]struct{}{start: {}}
	results := []int64{1}
	for j := 1; j <= count; j++ {
		locations = i.Tick(j, locations)
		results = append(results, int64(len(locations)))
	}
	return results
}
