package main

import (
	"fmt"
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

	var partOne, partTwo int
	for _, game := range parseGames(string(data)) {
		if game.Possible() {
			partOne += game.id
		}
		partTwo += game.Power()
	}
	log.Println("Part one:", partOne)
	log.Println("Part two:", partTwo)
}

func parseGames(input string) []game {
	var games []game
	for _, line := range strings.Split(input, "\n") {
		games = append(games, parseGame(line))
	}
	return games
}

func parseGame(line string) game {
	var g game
	parts := strings.Split(line, ": ")
	fmt.Sscanf(parts[0], "Game %d", &g.id)
	for _, r := range strings.Split(parts[1], "; ") {
		g.rounds = append(g.rounds, parseRound(r))
	}
	return g
}

func parseRound(r string) round {
	var rnd round
	for _, tuple := range strings.Split(r, ", ") {
		parts := strings.Fields(tuple)
		v, _ := strconv.Atoi(parts[0])
		switch parts[1] {
		case "red":
			rnd.r = v
		case "green":
			rnd.g = v
		case "blue":
			rnd.b = v
		}
	}
	return rnd
}

type game struct {
	id     int
	rounds []round
}

func (g game) Possible() bool {
	for _, r := range g.rounds {
		if !r.Possible() {
			return false
		}
	}
	return true
}

func (g game) Power() int {
	var minr, ming, minb int
	for _, r := range g.rounds {
		if r.r > minr {
			minr = r.r
		}
		if r.g > ming {
			ming = r.g
		}
		if r.b > minb {
			minb = r.b
		}
	}
	return minr * ming * minb
}

type round struct {
	r, g, b int
}

func (r round) Possible() bool {
	return r.r <= 12 && r.g <= 13 && r.b <= 14
}
