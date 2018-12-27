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

	nanobots := newBots(strings.Split(string(data), "\n"))

	strongest := nanobots.GetStrongest()

	var count int

	for _, n := range nanobots {
		if strongest.InRange(n.c) {
			count++
		}
	}

	log.Println("Part one:", count)

	// bc, most := nanobots.FindClosest(1<<(strconv.IntSize-2), boundingbox{})
	log.Println("Part two:", partTwo(nanobots))
}

func partTwo(nanobots bots) int {
	neighbours := findNeighbours(nanobots)

	cliques := bronKerbosch(neighbours)

	if len(cliques) != 1 {
		panic("unhandled")
	}

	clique := cliques[0]

	var maxMin int
	var origin coord
	for _, i := range clique {
		bot := nanobots[i]
		d := origin.Distance(bot.c) - bot.r
		if d > maxMin {
			maxMin = d
		}
	}

	return maxMin
}

func findNeighbours(nanobots bots) map[int][]int {
	neighbours := make(map[int][]int)
	for i, bot0 := range nanobots {
		for j, bot1 := range nanobots {
			if i == j || bot0.c.Distance(bot1.c) > bot0.r+bot1.r {
				continue
			}
			neighbours[j] = append(neighbours[j], i)
		}
	}
	return neighbours
}

type grid struct {
	minx, miny, minz, maxx, maxy, maxz int
}

type nanobot struct {
	r int
	c coord
}

func (n *nanobot) DistanceTo(o coord) int {
	return n.c.Distance(o)
}

func (n *nanobot) InRange(o coord) bool {
	return n.DistanceTo(o) <= n.r
}

type bots []*nanobot

func newBots(input []string) bots {
	var nanobots bots

	for _, line := range input {
		var x, y, z, r int
		fmt.Sscanf(line, "pos=<%d,%d,%d>, r=%d", &x, &y, &z, &r)

		newbot := &nanobot{c: coord{x: x, y: y, z: z}, r: r}
		nanobots = append(nanobots, newbot)
	}

	return nanobots
}

func (b bots) NumInRange(c coord) int {
	var num int

	for _, n := range b {
		if n.InRange(c) {
			num++
		}
	}

	return num
}

func (b bots) GetStrongest() *nanobot {
	var maxr int
	var strongest *nanobot

	for _, n := range b {
		if n.r > maxr {
			maxr = n.r
			strongest = n
		}
	}

	return strongest
}

type boundingbox struct {
	topl coord
	botr coord
}

var origin = coord{}

func (b bots) FindClosest(zoom int, bounding boundingbox) (coord, int) {
	var current coord

	var zoomed bots
	var bestcount int
	bestCoords := []coord{coord{}}

	for _, n := range b {
		zoomed = append(zoomed, &nanobot{c: coord{x: n.c.x / zoom, y: n.c.y / zoom, z: n.c.z / zoom}, r: n.r / zoom})
	}

	for current.x = bounding.topl.x; current.x <= bounding.botr.x; current.x++ {
		for current.y = bounding.topl.y; current.y <= bounding.botr.y; current.y++ {
			for current.z = bounding.topl.z; current.z <= bounding.botr.z; current.z++ {
				count := zoomed.NumInRange(current)

				if count < bestcount {
					continue
				}

				if count == bestcount {
					bestCoords = append(bestCoords, current)
				}

				// if count == bestcount && origin.Distance(current) > origin.Distance(bestCoords[0]) {
				// 	continue
				// }

				// if count == bestcount && origin.Distance(current) == origin.Distance(bestCoords[0]) {
				// 	bestCoords = append(bestCoords, current)
				// }

				bestCoords = []coord{current}
				bestcount = count
			}
		}
	}

	zoom >>= 1

	if zoom == 0 {
		return bestCoords[0], origin.Distance(bestCoords[0])
	}

	var most int
	var bestCoord coord

	for _, c := range bestCoords {
		bounding.topl.x, bounding.topl.y, bounding.topl.z = (c.x-1)<<1, (c.y-1)<<1, (c.z-1)<<1
		bounding.botr.x, bounding.botr.y, bounding.botr.z = (c.x+1)<<1, (c.y+1)<<1, (c.z+1)<<1

		if bc, adjacent := b.FindClosest(zoom, bounding); adjacent > most {
			most = adjacent
			bestCoord = bc
		}
	}

	return bestCoord, most
}

type coord struct {
	x, y, z int
}

func (c coord) Distance(o coord) int {
	return int(math.Abs(float64(c.x-o.x)) + math.Abs(float64(c.y-o.y)) + math.Abs(float64(c.z-o.z)))
}
