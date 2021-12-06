package main

import (
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func main() {
	data, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatalln("Failed to read input:", err)
	}

	log.Println("Part one:", newGame(string(data)).RunToTurn(2020))
	log.Println("Part two:", newGame(string(data)).RunToTurn(30000000))
}

type game struct {
	latest, secondLatest map[int]int
	turn                 int
	previous             int
}

func newGame(input string) *game {
	g := &game{latest: make(map[int]int), secondLatest: make(map[int]int)}
	for _, numStr := range strings.Split(input, ",") {
		v, err := strconv.Atoi(numStr)
		if err != nil {
			log.Fatalln("Failed to convert to number:", err)
		}
		g.Input(v)
	}
	return g
}

func (g *game) Input(num int) {
	g.turn++
	g.latest[num] = g.turn
	g.previous = num
}

func (g *game) TakeTurn() int {
	g.turn++
	next := g.latest[g.previous] - g.secondLatest[g.previous]
	if g.secondLatest[g.previous] == 0 {
		next = 0
	}
	g.latest[next], g.secondLatest[next] = g.turn, g.latest[next]
	g.previous = next
	g.latest[next] = g.turn
	return next
}

func (g *game) RunToTurn(turn int) int {
	var last int
	for g.turn != turn {
		last = g.TakeTurn()
	}
	return last
}
