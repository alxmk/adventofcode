package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"strings"
)

func main() {
	data, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatalln("Failed to read input:", err)
	}
	c := parse(string(data))
	log.Println("Part one:", partOne(c))
	log.Println("Part two:", partTwo(c))
}

func partOne(c cloud) int {
	return surfaceArea(lava, c)
}

func surfaceArea(p particle, c cloud) int {
	var sum int
	for x := c.min.x; x <= c.max.x; x++ {
		for y := c.min.y; y <= c.max.y; y++ {
			for z := c.min.z; z <= c.max.z; z++ {
				d := xyz{x, y, z}
				t := c.droplets[d]
				if t != p {
					continue
				}
				var ds int
				for _, a := range d.Adjacent() {
					if t := c.droplets[a]; t != p {
						ds++
					}
				}
				sum += ds
			}
		}
	}
	return sum
}

func partTwo(c cloud) int {
	for c.steamy() {

	}
	return surfaceArea(lava, c) - surfaceArea(air, c)
}

func (c cloud) steamy() bool {
	var changed bool
	for x := c.min.x; x <= c.max.x; x++ {
		for y := c.min.y; y <= c.max.y; y++ {
			for z := c.min.z; z <= c.max.z; z++ {
				t := xyz{x, y, z}
				p := c.droplets[t]
				if p == lava {
					continue
				}
				if p == air {
					// If we're at a boundary and it's not lava it's steam
					if x == c.min.x || x == c.max.x || y == c.min.y || y == c.max.y || z == c.min.z || z == c.max.z {
						p, c.droplets[t], changed = steam, steam, true
					}
				}
				if p == steam {
					for _, a := range t.Adjacent() {
						if c.droplets[a] == air {
							c.droplets[a], changed = steam, true
						}
					}
				}
			}
		}
	}
	return changed
}

func parse(input string) cloud {
	c := cloud{make(map[xyz]particle), xyz{math.MaxInt, math.MaxInt, math.MaxInt}, xyz{}}

	for _, line := range strings.Split(input, "\n") {
		var d xyz
		fmt.Sscanf(line, "%d,%d,%d", &d.x, &d.y, &d.z)
		c.droplets[d] = lava
		c.min.x, c.min.y, c.min.z = min(c.min.x, d.x), min(c.min.y, d.y), min(c.min.z, d.z)
		c.max.x, c.max.y, c.max.z = max(c.max.x, d.x), max(c.max.y, d.y), max(c.max.z, d.z)
	}

	return c
}

type cloud struct {
	droplets map[xyz]particle
	min, max xyz
}

type particle int

const (
	air particle = iota
	lava
	steam
)

type xyz struct {
	x, y, z int
}

func (x xyz) Adjacent() []xyz {
	return []xyz{
		{x.x - 1, x.y, x.z},
		{x.x + 1, x.y, x.z},
		{x.x, x.y - 1, x.z},
		{x.x, x.y + 1, x.z},
		{x.x, x.y, x.z - 1},
		{x.x, x.y, x.z + 1},
	}
}

func min(i, j int) int {
	if i < j {
		return i
	}
	return j
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}
