package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

func main() {
	data, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatalln("Failed to read input:", err)
	}

	f := parse(string(data))

	dir := vec{3, 1}
	log.Println("Part one:", f.move(dir))

	dirs := []vec{
		{1, 1},
		{3, 1},
		{5, 1},
		{7, 1},
		{1, 2},
	}

	total := 1
	for _, dir := range dirs {
		count := f.move(dir)
		log.Println(dir, count)
		total *= count
	}
	log.Println("Part two:", total)
}

type forest struct {
	trees      trees
	xmax, ymax int
}

func (f forest) String() string {
	var b strings.Builder
	for y := 0; y <= f.ymax; y++ {
		for x := 0; x <= f.xmax; x++ {
			if f.trees[x][y] {
				b.WriteRune('#')
			} else {
				b.WriteRune('.')
			}
		}
		b.WriteRune('\n')
	}
	return b.String()
}

func parse(input string) forest {
	f := forest{trees: make(trees)}
	for y, line := range strings.Split(input, "\n") {
		for x, char := range line {
			if _, ok := f.trees[x]; !ok {
				f.trees[x] = make(map[int]bool)
			}
			f.trees[x][y] = char == '#'
			if x > f.xmax {
				f.xmax = x
			}
		}
		if y > f.ymax {
			f.ymax = y
		}
	}
	return f
}

type trees map[int]map[int]bool

type vec struct{ x, y int }

func (v vec) String() string {
	return fmt.Sprintf("%d, %d", v.x, v.y)
}

func (f forest) move(dir vec) int {
	var pos vec
	var count int
	for pos.y <= f.ymax {
		if f.trees[pos.x%(f.xmax+1)][pos.y] {
			count++
		}
		pos.x, pos.y = pos.x+dir.x, pos.y+dir.y
	}
	return count
}
