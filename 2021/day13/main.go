package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func main() {
	data, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatalln("Failed to read input:", err)
	}

	paper, folds, err := parse(string(data))
	if err != nil {
		log.Fatalln("Failed to parse input:", err)
	}

	paper = paper.Fold(folds[0])

	log.Println("Part one:", len(paper.data))
	for i := 1; i < len(folds); i++ {
		paper = paper.Fold(folds[i])
	}
	log.Printf("Part two:\n%s", paper)
}

func parse(input string) (*grid, []coord, error) {
	parts := strings.Split(input, "\n\n")
	if l := len(parts); l != 2 {
		return nil, nil, fmt.Errorf("malformed input, expected 2 sections but got %d", l)
	}
	g := newGrid()
	for _, line := range strings.Split(parts[0], "\n") {
		var c coord
		if _, err := fmt.Sscanf(line, "%d,%d", &c.x, &c.y); err != nil {
			return nil, nil, fmt.Errorf("malformed coordinate %s: %s", err, line)
		}
		g.data[c] = struct{}{}
		if c.x > g.xmax {
			g.xmax = c.x
		}
		if c.y > g.ymax {
			g.ymax = c.y
		}
	}
	var folds []coord
	for _, line := range strings.Split(parts[1], "\n") {
		subparts := strings.Fields(line)
		if len(subparts) != 3 {
			return nil, nil, fmt.Errorf("malformed fold %s", line)
		}
		subsubparts := strings.Split(subparts[2], "=")
		if len(subsubparts) != 2 {
			return nil, nil, fmt.Errorf("malformed fold %s", line)
		}
		v, err := strconv.Atoi(subsubparts[1])
		if err != nil {
			return nil, nil, fmt.Errorf("malformed fold %s: %s", line, err)
		}
		switch subsubparts[0] {
		case "x":
			folds = append(folds, coord{x: v})
		case "y":
			folds = append(folds, coord{y: v})
		default:
			return nil, nil, fmt.Errorf("malformed fold %s", line)
		}
	}
	return g, folds, nil
}

type coord struct{ x, y int }

func (c coord) String() string {
	return fmt.Sprintf("%d,%d", c.x, c.y)
}

type grid struct {
	data       map[coord]struct{}
	xmax, ymax int
}

func newGrid() *grid { return &grid{data: make(map[coord]struct{})} }

func (g grid) String() string {
	var b strings.Builder
	for y := 0; y <= g.ymax; y++ {
		for x := 0; x <= g.xmax; x++ {
			if _, ok := g.data[coord{x, y}]; ok {
				b.WriteRune('#')
				continue
			}
			b.WriteRune('.')
		}
		b.WriteRune('\n')
	}
	return b.String()
}

func (g *grid) Fold(f coord) *grid {
	updated := newGrid()
	switch {
	case f.x == 0:
		// log.Println("Folding along y = ", f.y)
		// Fold up along f.y
		var offset int
		for y := 0; y <= g.ymax; y++ {
			if y > f.y {
				offset += 2
			}
			for x := 0; x <= g.xmax; x++ {
				// log.Println(coord{x, y}, "->", coord{x, y - offset})
				if _, ok := g.data[coord{x, y}]; ok {
					updated.data[coord{x, y - offset}] = struct{}{}
				}
			}
		}
		updated.xmax, updated.ymax = g.xmax, f.y-1
	case f.y == 0:
		// log.Println("Folding along x = ", f.x)
		// Fold left along f.x
		var offset int
		for x := 0; x <= g.xmax; x++ {
			if x > f.x {
				offset += 2
			}
			for y := 0; y <= g.ymax; y++ {
				if _, ok := g.data[coord{x, y}]; ok {
					updated.data[coord{x - offset, y}] = struct{}{}
				}
			}
		}
		updated.xmax, updated.ymax = f.x-1, g.ymax
	}

	return updated
}
