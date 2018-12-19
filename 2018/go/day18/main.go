package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

func main() {
	data, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatalln("Failed to read input", err)
	}

	lines := strings.Split(string(data), "\n")

	w := make(world, len(lines[0]))
	for y, line := range lines {
		for x, r := range line {
			if y == 0 {
				w[x] = make([]tile, len(lines))
			}

			switch r {
			case '.':
				w[x][y] = open
			case '|':
				w[x][y] = trees
			case '#':
				w[x][y] = lumberyard
			}
		}
	}

	// fmt.Println(w.String())

	patterns := make(map[string]int)
	scores := make(map[int]int)

	var firstInstance, secondInstance int

	for minute := 1; minute <= 1000000000; minute++ {
		w = w.Next()
		pattern := w.String()
		if m, ok := patterns[pattern]; ok {
			firstInstance = m
			secondInstance = minute
			break
		}
		patterns[pattern] = minute
		scores[minute] = w.Score()

		if minute == 10 {
			fmt.Println("Part one:", scores[minute])
		}
	}

	patternInterval := secondInstance - firstInstance

	increment := (1000000000 - firstInstance) % patternInterval

	fmt.Println("Part two:", scores[increment+firstInstance])
}

type world [][]tile

func (w world) String() string {
	var buf bytes.Buffer

	for y := range w[0] {
		for x := range w {
			switch w[x][y] {
			case open:
				buf.WriteString(".")
			case trees:
				buf.WriteString("|")
			case lumberyard:
				buf.WriteString("#")
			}
		}
		buf.WriteString("\n")
	}

	return buf.String()
}

func (w world) Next() world {
	newWorld := make(world, len(w))

	for x := range w {
		for y := range w[0] {
			if y == 0 {
				newWorld[x] = make([]tile, len(w[0]))
			}

			newWorld[x][y] = w.NextTile(x, y)
		}
	}

	return newWorld
}

func (w world) NextTile(x, y int) tile {
	adjacent := make(map[tile]int)

	for i := x - 1; i <= x+1; i++ {
		for j := y - 1; j <= y+1; j++ {
			if i == x && j == y {
				continue
			}
			if t, ok := w.Get(i, j); ok {
				adjacent[t]++
			}
		}
	}

	switch w[x][y] {
	case open:
		if adjacent[trees] >= 3 {
			return trees
		}
		return open
	case trees:
		if adjacent[lumberyard] >= 3 {
			return lumberyard
		}
		return trees
	case lumberyard:
		if adjacent[trees] >= 1 && adjacent[lumberyard] >= 1 {
			return lumberyard
		}
		return open
	default:
		return w[x][y]
	}
}

func (w world) Get(x, y int) (tile, bool) {
	var t tile

	if x < 0 || x >= len(w) || y < 0 || y >= len(w[0]) {
		return t, false
	}

	return w[x][y], true
}

func (w world) Score() int {
	var numtrees, numlumberyard int
	for y := range w[0] {
		for x := range w {
			switch w[x][y] {
			case trees:
				numtrees++
			case lumberyard:
				numlumberyard++
			}
		}
	}
	return numtrees * numlumberyard
}

type tile int

const (
	open       tile = 1
	trees      tile = 10
	lumberyard tile = 100
)
