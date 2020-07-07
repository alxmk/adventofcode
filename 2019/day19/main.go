package main

import (
	"fmt"
	"io/ioutil"
	"log"

	"github.ibm.com/alexmk/adventofcode/2019/day2/intcode"
)

func main() {
	data, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatalln("Failed to read input file:", err)
	}

	programme, err := intcode.Parse(string(data))
	if err != nil {
		log.Fatalln("Failed to parse input as intcode programme:", err)
	}

	var numInBeam int
	for x := int64(0); x < 50; x++ {
		for y := int64(0); y < 50; y++ {
			if t := checkTile(x, y, programme); t == 1 {
				numInBeam++
			}
		}
	}

	log.Println("Part one:", numInBeam)

	c := squareStartsAt(100, programme)

	log.Println("Part two:", c.x*10000+c.y)
}

type coordinate struct {
	x, y int64
}

func (c coordinate) String() string {
	return fmt.Sprintf("(%d, %d)", c.x, c.y)
}

func squareStartsAt(size int64, programme intcode.Programme) coordinate {
	c := coordinate{0, 100}
	for {
		// Find bottom left
		for t := checkTile(c.x, c.y, programme); t != 1; t = checkTile(c.x, c.y, programme) {
			c.x++
		}

		// Check top right
		if t := checkTile(c.x+99, c.y-99, programme); t == 1 {
			// Return top left
			return coordinate{c.x, c.y - 99}
		}

		c.y++
	}
}

func checkTile(x, y int64, programme intcode.Programme) int64 {
	in, out := make(chan int64), make(chan int64)

	go func() {
		if err := programme.Copy().Run(in, out); err != nil {
			log.Fatalln("Intcode machine failed:", err)
		}
	}()

	in <- x
	in <- y
	close(in)
	return <-out
}
