package main

import (
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	data, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatalln("Failed to read input:", err)
	}

	log.Println("Part one:", solve(parseRaces(string(data))))
	log.Println("Part two:", solve([]*race{parseRace(string(data))}))
}

func solve(races []*race) int {
	product := 1
	for _, r := range races {
		var ways int
		for i := 1; ; i++ {
			if r.Win(i) {
				ways++
				continue
			}
			if ways > 0 {
				break
			}
		}
		product *= ways
	}
	return product
}

func parseRaces(input string) []*race {
	lines := strings.Split(input, "\n")
	var races []*race
	for _, t := range strings.Fields(lines[0])[1:] {
		v, _ := strconv.Atoi(t)
		races = append(races, &race{time: v})
	}
	for i, d := range strings.Fields(lines[1])[1:] {
		v, _ := strconv.Atoi(d)
		races[i].distance = v
	}
	return races
}

func parseRace(input string) *race {
	var r race
	lines := strings.Split(input, "\n")
	r.time, _ = strconv.Atoi(strings.Join(strings.Fields(lines[0])[1:], ""))
	r.distance, _ = strconv.Atoi(strings.Join(strings.Fields(lines[1])[1:], ""))
	return &r
}

type race struct {
	time, distance int
}

func (r race) Simulate(hold int) int {
	return (r.time - hold) * hold
}

func (r race) Win(hold int) bool {
	return r.Simulate(hold) > r.distance
}
