package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

func main() {
	data, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatalln("Failed to read input:", err)
	}

	t, err := parse(string(data))
	if err != nil {
		log.Fatalln("Failed to parse input:", err)
	}

	maxy, hits := findHits(t)
	log.Println("Part one:", maxy)
	log.Println("Part two:", hits)
}

func parse(input string) (*target, error) {
	var t target
	if _, err := fmt.Sscanf(input, "target area: x=%d..%d, y=%d..%d", &t.min.x, &t.max.x, &t.min.y, &t.max.y); err != nil {
		return nil, fmt.Errorf("malformed input: %s", err)
	}
	return &t, nil
}

type target struct {
	min, max coord
}

type coord struct {
	x, y int
}

func (c coord) String() string {
	return fmt.Sprintf("%d,%d", c.x, c.y)
}

func (t target) Contains(c coord) bool {
	return c.x >= t.min.x && c.x <= t.max.x &&
		c.y >= t.min.y && c.y <= t.max.y
}

func (t target) PassedBy(c coord) bool {
	return c.x > t.max.x || c.y < t.min.y
}

func findHits(t *target) (int, int) {
	var maxy, hits int
	for x := 1; x < 1000; x++ {
		for y := -1000; y < 1000; y++ {
			if ok, h := shoot(coord{x, y}, *t); ok {
				hits++
				if h > maxy {
					maxy = h
				}
			}
		}
	}
	return maxy, hits
}

func shoot(v coord, t target) (bool, int) {
	// log.Println("Shooting", v)
	var p coord
	maxy := p.y
	for i := 0; i < 10000; i++ {
		p.x += v.x
		p.y += v.y
		if p.y > maxy {
			maxy = p.y
		}
		if t.Contains(p) {
			return true, maxy
		}
		if t.PassedBy(p) {
			return false, 0
		}
		if v.x > 0 {
			v.x--
		}
		if v.x < 0 {
			v.x++
		}
		v.y--
	}
	return false, 0
}
