package main

import (
	"io/ioutil"
	"log"
)

func main() {
	data, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatalln("Failed to read input file", err)
	}

	visited1 := map[coordinate]struct{}{
		coordinate{0, 0}: struct{}{},
	}
	visited2 := map[coordinate]struct{}{
		coordinate{0, 0}: struct{}{},
	}

	var loneSanta coordinate
	santa, robot := &coordinate{0, 0}, &coordinate{0, 0}

	for i, dir := range string(data) {
		current := robot
		if i%2 == 0 {
			current = santa
		}
		switch dir {
		case '^':
			loneSanta.y++
			current.y++
		case 'v':
			loneSanta.y--
			current.y--
		case '>':
			loneSanta.x++
			current.x++
		case '<':
			loneSanta.x--
			current.x--
		}

		visited1[loneSanta] = struct{}{}
		visited2[*current] = struct{}{}
	}

	log.Println("Part one:", len(visited1))
	log.Println("Part two:", len(visited2))
}

type coordinate struct {
	x, y int
}
