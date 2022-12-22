package main

import (
	"log"
	"os"
)

func main() {
	data, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatalln("Failed to read input:", err)
	}
}

type xy struct {
	x, y int
}

type direction int

const (
	right direction = iota
	down
	left
	up
)

type tile int

const (
	empty tile = iota
	floor
	wall
)
