package main

import (
	"log"
	"math"
)

const (
	offset = 382
)

func main() {
	c := &CircularBuffer{
		contents: []int{},
		position: 0,
	}

	for i := 0; i <= 2017; i++ {
		c.Insert(offset, i)
	}

	log.Println("Part one answer is", c.Get(c.position+1))

	// For part two we know where zero starts so any time we move it or insert after just track that
	var position, value, zIndex int

	for i := 1; i < 50000000; i++ {
		position = int(math.Mod(float64(position+offset), float64(i))) + 1
		if position <= zIndex {
			zIndex++
		}

		if position == zIndex+1 {
			value = i
		}
	}

	log.Println("Part two answer is", value)
}

type CircularBuffer struct {
	contents []int
	position int
}

func (c *CircularBuffer) Insert(offset, value int) {
	var realPosition int
	if len(c.contents) != 0 {
		realPosition = int(math.Mod(float64(c.position+offset), float64(len(c.contents)))) + 1
	}

	newContents := make([]int, len(c.contents[:realPosition]))
	copy(newContents, c.contents[:realPosition])

	newContents = append(newContents, value)
	newContents = append(newContents, c.contents[realPosition:]...)

	c.position = realPosition
	c.contents = newContents

}

func (c *CircularBuffer) Get(position int) int {
	realPosition := int(math.Mod(float64(position), float64(len(c.contents))))

	log.Println("GET: Real position is", realPosition, "requested", position)

	return c.contents[realPosition]
}
