package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	data, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatalln("Failed to read input:", err)
	}

	log.Println("Part one:", calculate(string(data), 2))
	log.Println("Part two:", calculate(string(data), 10))
}

func calculate(input string, length int) int {
	r := rope{tailCache: make(map[xy]struct{}), knots: make([]xy, length)}

	for _, line := range strings.Split(input, "\n") {
		var direction string
		var magnitude int
		fmt.Sscanf(line, "%s %d", &direction, &magnitude)
		r.Move(cardinal[direction], magnitude)
	}

	return len(r.tailCache)
}

type xy struct {
	x, y int
}

type rope struct {
	knots     []xy
	tailCache map[xy]struct{}
}

func (r *rope) Move(dh dir, count int) {
	for c := 0; c < count; c++ {
		for i, move := 0, dh; i < len(r.knots); i++ {
			r.knots[i].x, r.knots[i].y = r.knots[i].x+move.x, r.knots[i].y+move.y
			if i+1 < len(r.knots) {
				move = resolve(r.knots[i], r.knots[i+1])
			}
		}
		r.tailCache[r.knots[len(r.knots)-1]] = struct{}{}
	}
}

var resolveMap = map[xy]dir{
	{0, 2}:   {0, 1},   // up
	{0, -2}:  {0, -1},  // down
	{2, 0}:   {1, 0},   // right
	{-2, 0}:  {-1, 0},  // left
	{2, 1}:   {1, 1},   // up + right
	{2, 2}:   {1, 1},   // up + right
	{1, 2}:   {1, 1},   // up + right
	{2, -1}:  {1, -1},  // down + right
	{2, -2}:  {1, -1},  // down + right
	{1, -2}:  {1, -1},  // down + right
	{-2, 1}:  {-1, 1},  // up + left
	{-2, 2}:  {-1, 1},  // up + left
	{-1, 2}:  {-1, 1},  // up + left
	{-2, -1}: {-1, -1}, // down + left
	{-2, -2}: {-1, -1}, // down + left
	{-1, -2}: {-1, -1}, // down + left
}

func resolve(a, b xy) dir {
	return resolveMap[xy{a.x - b.x, a.y - b.y}]
}

type dir xy

var cardinal = map[string]dir{
	"U": {0, 1},
	"D": {0, -1},
	"L": {-1, 0},
	"R": {1, 0},
}
