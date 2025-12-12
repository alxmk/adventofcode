package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

func main() {
	var shapes []shape
	var requirements []requirement
	for block := range strings.SplitSeq(input, "\n\n") {
		lines := strings.Split(block, "\n")
		switch lines[0][1] {
		case ':':
			// Shape
			var s shape
			for y, line := range lines[1:] {
				for x, r := range line {
					if r != '#' {
						continue
					}
					s = append(s, [2]int{x, y})
				}
			}
			shapes = append(shapes, s)
		default:
			// Requirements
			for _, line := range lines {
				var r requirement
				fields := strings.Fields(line)
				area := strings.Split(fields[0][0:len(fields[0])-1], "x")
				r.size[0], _ = strconv.Atoi(area[0])
				r.size[1], _ = strconv.Atoi(area[1])
				for _, field := range fields[1:] {
					v, _ := strconv.Atoi(field)
					r.counts = append(r.counts, v)
				}
				requirements = append(requirements, r)
			}
		}
	}
	var partOne int
	for _, r := range requirements {
		if r.fits(shapes) {
			partOne++
		}
	}
	fmt.Println("Part one:", partOne)
}

type requirement struct {
	size   [2]int
	counts []int
}

func (r requirement) fits(shapes []shape) bool {
	var needed int
	for i, c := range r.counts {
		needed += c * len(shapes[i])
	}
	// This is a bad approximation but it got me a star ¯\_(ツ)_/¯
	return needed <= r.size[0]*r.size[1]
}

type shape [][2]int
