package main

import (
	_ "embed"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

func main() {
	points := parse(input)
	// [i, j, area]
	var sizes [][3]int
	for i := range len(points) {
		for j := i + 1; j < len(points); j++ {
			sizes = append(sizes, [3]int{i, j, area(points[i], points[j])})
		}
	}
	sort.Slice(sizes, func(i, j int) bool {
		return sizes[i][2] > sizes[j][2]
	})
	fmt.Println("Part one:", sizes[0][2])
	lines := calculateLines(points)
	for _, values := range sizes {
		i, j := values[0], values[1]
		if intersects(points[i], points[j], lines) {
			continue
		}
		fmt.Println("Part two:", values[2])
		break
	}
}

func parse(input string) (points []xy) {
	for line := range strings.SplitSeq(input, "\n") {
		i := strings.Index(line, ",")
		var p xy
		p[0], _ = strconv.Atoi(line[:i])
		p[1], _ = strconv.Atoi(line[i+1:])
		points = append(points, p)
	}
	return points
}

func calculateLines(points []xy) (lines []line) {
	point := points[0]
	for i := 1; i < len(points); i++ {
		lines = append(lines, line{point, points[i]})
		point = points[i]
	}
	lines = append(lines, line{point, points[0]})
	return lines
}

func intersects(i, j xy, lines []line) bool {
	rmin, rmax := xy{min(i[0], j[0]), min(i[1], j[1])}, xy{max(i[0], j[0]), max(i[1], j[1])}
	for _, l := range lines {
		lmin, lmax := l.min(), l.max()
		if lmax[0] > rmin[0] && lmin[0] < rmax[0] &&
			lmax[1] > rmin[1] && lmin[1] < rmax[1] {
			return true
		}
	}
	return false
}

type xy [2]int

func area(a, b xy) int {
	return (mod(a[0]-b[0]) + 1) * (mod(a[1]-b[1]) + 1)
}

func mod(a int) int {
	if a < 0 {
		return a * -1
	}
	return a
}

type line [2]xy

func (l line) min() xy {
	return xy{min(l[0][0], l[1][0]), min(l[0][1], l[1][1])}
}

func (l line) max() xy {
	return xy{max(l[0][0], l[1][0]), max(l[0][1], l[1][1])}
}
