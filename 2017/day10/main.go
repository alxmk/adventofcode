package main

import (
	"io/ioutil"
	"log"
	"strconv"
	"strings"

	"github.com/alxmk/adventofcode/2017/day10/knot"
)

func main() {
	input, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatalf("Error reading input file %v", err)
	}

	lengths := strings.Split(string(input), ",")

	loop := knot.NewLoop(256)
	skip := 0
	pos := 0

	lengthInts := []int{}

	for _, l := range lengths {
		value, err := strconv.Atoi(l)
		if err != nil {
			log.Fatalf("Couldn't convert %s to int", l)
		}

		lengthInts = append(lengthInts, value)
	}

	loop, skip, pos = knot.OneRound(loop, skip, pos, lengthInts)

	log.Println("Part one answer is", loop.Contents[0]*loop.Contents[1])
	log.Println("Part two answer is", knot.KnotHash(string(input)))
}
