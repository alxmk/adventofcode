package main

import (
	"bytes"
	"log"
	"math"
	"os"
	"strings"
)

func main() {
	data, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatalln("Failed to read input:", err)
	}

	p, im := parseInput(data)

	log.Println("Part one:", solve(p, im, 2))
	log.Println("Part two:", solve(p, im, 50))
}

func solve(programme []byte, im image, iterations int) int {
	for i := 0; i < iterations; i++ {
		im = im.Enhance(programme)
	}
	return len(im.pixels)
}

func parseInput(input []byte) ([]byte, image) {
	i := image{pixels: make(map[xy]bool)}
	blocks := bytes.Split(input, []byte{'\n', '\n'})

	for y, line := range bytes.Split(blocks[1], []byte{'\n'}) {
		for x, r := range line {
			if r == '#' {
				i.pixels[xy{x, y}] = true
			}
			i.max.x = x
		}
		i.max.y = y
	}
	return blocks[0], i
}

type image struct {
	pixels   map[xy]bool
	min, max xy
	border   bool
}

func (i image) String() string {
	var b strings.Builder
	for y := i.min.y - 1; y <= i.max.y+1; y++ {
		for x := i.min.x - 1; x <= i.max.x+1; x++ {
			if i.pixels[xy{x, y}] {
				b.WriteRune('#')
				continue
			}
			b.WriteRune('.')
		}
		b.WriteRune('\n')
	}
	return b.String()
}

func (i image) Enhance(programme []byte) image {
	j := image{
		pixels: make(map[xy]bool),
		border: nextBorder(programme, i.border),
		min:    xy{math.MaxInt, math.MaxInt},
	}

	for y := i.min.y - 1; y <= i.max.y+1; y++ {
		for x := i.min.x - 1; x <= i.max.x+1; x++ {
			if !func() bool {
				var idx int
				exp := 8
				for b := y - 1; b <= y+1; b++ {
					for a := x - 1; a <= x+1; a++ {
						if i.pixels[xy{a, b}] || (i.border && (a < i.min.x || b < i.min.y || a > i.max.x || b > i.max.y)) {
							idx += 1 << exp
						}
						exp--
					}
				}
				return programme[idx] == '#'
			}() {
				continue
			}
			j.pixels[xy{x, y}] = true
			j.min.x, j.min.y = min(j.min.x, x), min(j.min.y, y)
			j.max.x, j.max.y = max(j.max.x, x), max(j.max.y, y)
		}
	}

	return j
}

func nextBorder(programme []byte, current bool) bool {
	if current {
		return programme[511] == '#'
	}
	return programme[0] == '#'
}

type xy struct {
	x, y int
}
