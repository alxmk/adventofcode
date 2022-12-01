package main

import (
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	data, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatalln("Failed to read input:", err)
	}

	var elves []int
	for _, inventory := range strings.Split(string(data), "\n\n") {
		var this int
		for _, item := range strings.Split(inventory, "\n") {
			cal, err := strconv.Atoi(item)
			if err != nil {
				log.Fatalf("Unparseable %s: %s", item, err)
			}
			this += cal
		}
		elves = append(elves, this)
	}
	sort.Ints(elves)
	log.Println("Part one:", elves[len(elves)-1])
	log.Println("Part two:", elves[len(elves)-1]+elves[len(elves)-2]+elves[len(elves)-3])
}
