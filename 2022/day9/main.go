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

func (x xy) String() string {
	return fmt.Sprintf("%d,%d", x.x, x.y)
}

type rope struct {
	knots     []xy
	tailCache map[xy]struct{}
}

func (r *rope) Move(dh dir, count int) {
	for c := 0; c < count; c++ {
		move := dh
		for i := 0; i < len(r.knots); i++ {
			r.knots[i].x, r.knots[i].y = r.knots[i].x+move.x, r.knots[i].y+move.y
			if i+1 < len(r.knots) {
				move = resolve(r.knots[i], r.knots[i+1])
			}
		}
		r.tailCache[r.knots[len(r.knots)-1]] = struct{}{}
	}
}

func resolve(a, b xy) dir {
	// right and left
	if a.y == b.y {
		switch a.x - b.x {
		case 2:
			return cardinal["R"]
		case -2:
			return cardinal["L"]
		}
	}
	// up and down
	if a.x == b.x {
		switch a.y - b.y {
		case 2:
			return cardinal["U"]
		case -2:
			return cardinal["D"]
		}
	}
	// diagonals
	switch a.x - b.x {
	case 2:
		switch a.y - b.y {
		case 1, 2:
			return sumDirs(cardinal["U"], cardinal["R"])
		case -1, -2:
			return sumDirs(cardinal["D"], cardinal["R"])
		}
	case -2:
		switch a.y - b.y {
		case 1, 2:
			return sumDirs(cardinal["U"], cardinal["L"])
		case -1, -2:
			return sumDirs(cardinal["D"], cardinal["L"])
		}
	case 1:
		switch a.y - b.y {
		case 2:
			return sumDirs(cardinal["U"], cardinal["R"])
		case -2:
			return sumDirs(cardinal["D"], cardinal["R"])
		}
	case -1:
		switch a.y - b.y {
		case 2:
			return sumDirs(cardinal["U"], cardinal["L"])
		case -2:
			return sumDirs(cardinal["D"], cardinal["L"])
		}
	}
	// Else no movement needed
	return dir{}
}

type dir xy

var cardinal = map[string]dir{
	"U": {0, 1},
	"D": {0, -1},
	"L": {-1, 0},
	"R": {1, 0},
}

func sumDirs(dirs ...dir) dir {
	var resultant dir
	for _, d := range dirs {
		resultant.x, resultant.y = resultant.x+d.x, resultant.y+d.y
	}
	return resultant
}
