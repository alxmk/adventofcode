package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
)

func main() {
	data, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatalln("Failed to read input:", err)
	}

	bricks := parseBricks(data)

	log.Println("Part one:", partOne(bricks))
	log.Println("Part two:", partTwo(bricks))
}

func partOne(bricks []brick) int {
	var count int
	for i := 0; i < len(bricks); i++ {
		cpy := make([]brick, len(bricks))
		copy(cpy, bricks)
		if _, changed := tick(append(cpy[:i], cpy[i+1:]...)); len(changed) == 0 {
			count++
		}
	}
	return count
}

func partTwo(bricks []brick) int {
	var sum int
	for i := 0; i < len(bricks); i++ {
		cpy := make([]brick, len(bricks))
		copy(cpy, bricks)
		overall := make(map[int]struct{})
		for cpy, changed := tick(append(cpy[:i], cpy[i+1:]...)); len(changed) != 0; cpy, changed = tick(cpy) {
			for k := range changed {
				overall[k] = struct{}{}
			}
		}
		sum += len(overall)
	}
	return sum
}

func tick(bricks []brick) ([]brick, map[int]struct{}) {
	occupied := resolve(bricks)
	var newbricks []brick
	changed := make(map[int]struct{})
outer:
	for i, b := range bricks {
		if min(b[0].z, b[1].z) == 1 {
			newbricks = append(newbricks, b)
			continue
		}
		for x := min(b[0].x, b[1].x); x <= max(b[0].x, b[1].x); x++ {
			for y := min(b[0].y, b[1].y); y <= max(b[0].y, b[1].y); y++ {
				if _, ok := occupied[xyz{x, y, min(b[0].z-1, b[1].z-1)}]; ok {
					newbricks = append(newbricks, b)
					continue outer
				}
			}
		}
		changed[i] = struct{}{}
		newbricks = append(newbricks, brick{{b[0].x, b[0].y, b[0].z - 1}, {b[1].x, b[1].y, b[1].z - 1}})
	}
	return newbricks, changed
}

func resolve(bricks []brick) map[xyz]struct{} {
	occupied := make(map[xyz]struct{})
	for _, b := range bricks {
		for x := min(b[0].x, b[1].x); x <= max(b[0].x, b[1].x); x++ {
			for y := min(b[0].y, b[1].y); y <= max(b[0].y, b[1].y); y++ {
				for z := min(b[0].z, b[1].z); z <= max(b[0].z, b[1].z); z++ {
					occupied[xyz{x, y, z}] = struct{}{}
				}
			}
		}
	}
	return occupied
}

func parseBricks(input []byte) []brick {
	var bricks []brick
	for _, line := range bytes.Split(input, []byte{'\n'}) {
		var b brick
		fmt.Sscanf(string(line), "%d,%d,%d~%d,%d,%d", &b[0].x, &b[0].y, &b[0].z, &b[1].x, &b[1].y, &b[1].z)
		bricks = append(bricks, b)
	}
	var moved map[int]struct{}
	for bricks, moved = tick(bricks); len(moved) != 0; bricks, moved = tick(bricks) {
	}
	return bricks
}

type brick [2]xyz

type xyz struct {
	x, y, z int
}
