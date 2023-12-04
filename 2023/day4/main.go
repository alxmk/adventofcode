package main

import (
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	data, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatalln("Failed to read input:", err)
	}

	log.Println("Part one:", partOne(parseGames(string(data))))
	log.Println("Part two:", partTwo(parseGames(string(data))))
}

func partOne(games map[int]*game) float64 {
	var sum float64
	for _, g := range games {
		sum += g.Points()
	}
	return sum
}

func partTwo(games map[int]*game) int {
	for id := 1; ; id++ {
		g, ok := games[id]
		if !ok {
			break
		}
		matches := g.Matches()
		for jd := id + 1; jd <= id+matches; jd++ {
			if j, ok := games[jd]; ok {
				j.count += g.count
			}
		}
	}
	var sum int
	for _, g := range games {
		sum += g.count
	}
	return sum
}

func parseGames(input string) map[int]*game {
	games := make(map[int]*game)
	for _, line := range strings.Split(input, "\n") {
		g := parseGame(line)
		games[g.id] = &g
	}
	return games
}

func parseGame(line string) game {
	parts := strings.Split(line, ": ")
	id, _ := strconv.Atoi(strings.Fields(parts[0])[1])
	numbers := strings.Split(parts[1], " | ")
	g := game{numbers: make(map[string]struct{}), winners: strings.Fields(numbers[0]), count: 1, id: id}
	for _, n := range strings.Fields(numbers[1]) {
		g.numbers[n] = struct{}{}
	}
	return g
}

type game struct {
	id      int
	count   int
	winners []string
	numbers map[string]struct{}
}

func (g game) Points() float64 {
	if matches := float64(g.Matches()); matches > 0 {
		return math.Pow(2, matches-1)
	}
	return 0.0
}

func (g game) Matches() int {
	var matches int
	for _, w := range g.winners {
		if _, ok := g.numbers[w]; ok {
			matches++
		}
	}
	return matches
}
