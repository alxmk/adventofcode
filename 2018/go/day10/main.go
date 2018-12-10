package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"strings"
)

func main() {
	data, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatalln("Failed to read input", err)
	}

	var stars []*star

	for _, line := range strings.Split(string(data), "\n") {
		stars = append(stars, parse(line))
	}

	// print(stars, -6, -4, 15, 11)
	t := 1

	for {
		var maxY, maxX int
		minY, minX := math.MaxInt32, math.MaxInt32
		for _, star := range stars {
			star.Tick()
			if star.y < minY {
				minY = star.y
			}
			if star.y > maxY {
				maxY = star.y
			}
			if star.x < minX {
				minX = star.x
			}
			if star.x > maxX {
				maxX = star.x
			}
		}
		// print(stars, minX, minY, maxX, maxY)
		if maxY-minY <= 10 {
			fmt.Println(t, maxX-minX, maxY-minY)
			print(stars, minX, minY, maxX, maxY)
			break
		}
		t++
	}
}

type star struct {
	x, y, dx, dy int
}

func (s *star) Tick() {
	s.x += s.dx
	s.y += s.dy
}

func parse(line string) *star {
	s := &star{}
	if _, err := fmt.Sscanf(line, "position=<%d,  %d> velocity=<%d,  %d>", &s.x, &s.y, &s.dx, &s.dy); err != nil {
		log.Fatalln("Failed to parse", line, err)
	}
	return s
}

func print(stars []*star, minX, minY, maxX, maxY int) {
	for y := minY; y <= maxY; y++ {
		for x := minX; x <= maxX; x++ {
			var found bool
			for _, s := range stars {
				if s.x == x && s.y == y {
					found = true
					break
				}
			}
			if found {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Print("\n")
	}
}
