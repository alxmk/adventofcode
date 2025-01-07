package main

import (
	_ "embed"
	"fmt"
	"strings"
)

//go:embed input.txt
var input string

func main() {
	l, p := parse(input)
	_, v, path := l.partOne(p)
	fmt.Println("Part one:", v)
	fmt.Println("Part two:", l.partTwo(p, path))
}

type xy [2]int

type dir xy

var (
	left  = dir{-1, 0}
	right = dir{1, 0}
	up    = dir{0, -1}
	down  = dir{0, 1}
)

func (d dir) right() dir {
	switch d {
	case left:
		return up
	case right:
		return down
	case up:
		return right
	case down:
		return left
	}
	panic("unexpected dir")
}

type lab map[xy]rune

func parse(input string) (lab, xy) {
	l := make(lab)
	var start xy
	for y, line := range strings.Split(input, "\n") {
		for x, r := range line {
			switch r {
			case '^', '>', '<', 'v':
				l[xy{x, y}] = '.'
				start = xy{x, y}
			default:
				l[xy{x, y}] = r
			}
		}
	}
	return l, start
}

func (l lab) clone() lab {
	clone := make(lab)
	for k, v := range l {
		clone[k] = v
	}
	return clone
}

func (l lab) partOne(start xy) (bool, int, map[xy]map[dir]struct{}) {
	positions := map[xy]map[dir]struct{}{start: {up: {}}}
	d := up
	pos := start
	for {
		next := xy{pos[0] + d[0], pos[1] + d[1]}
		switch l[next] {
		case rune(0):
			return false, len(positions), positions
		case '.':
			pos = next
			if _, ok := positions[pos]; !ok {
				positions[pos] = make(map[dir]struct{})
			}
			if _, ok := positions[pos][d]; ok {
				return true, 0, positions
			}
			positions[pos][d] = struct{}{}
		case '#':
			d = d.right()
		}
	}
}

func (l lab) partTwo(start xy, path map[xy]map[dir]struct{}) int {
	positions := make(map[xy]struct{})
	for obs := range path {
		if obs == start {
			continue
		}
		c := l.clone()
		c[obs] = '#'
		if loop, _, _ := c.partOne(start); loop {
			positions[obs] = struct{}{}
		}
	}
	return len(positions)
}
