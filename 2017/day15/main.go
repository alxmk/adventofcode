package main

import (
	"bytes"
	"log"
	"math"
)

const (
	modValue = 2147483647
)

var (
	bufferA = new(bytes.Buffer)
	bufferB = new(bytes.Buffer)
)

func main() {
	generatorA := NewGenerator(783, 16807, 1)
	generatorB := NewGenerator(325, 48271, 1)

	var matched int

	for i := 0; i < 40000000; i++ {
		if match(int(generatorA.NextValue()), int(generatorB.NextValue())) {
			matched++
		}
	}

	log.Println("Part one answer is", matched)

	generatorA = NewGenerator(783, 16807, 4)
	generatorB = NewGenerator(325, 48271, 8)

	matched = 0

	for i := 0; i < 5000000; i++ {
		if match(int(generatorA.NextValue()), int(generatorB.NextValue())) {
			matched++
		}
	}

	log.Println("Part two answer is", matched)
}

func match(a, b int) bool {
	return a&0xffff == b&0xffff
}

type Generator struct {
	previous float64
	factor   float64
	multiple float64
}

func NewGenerator(seed, factor, multiple float64) *Generator {
	return &Generator{
		previous: seed,
		factor:   factor,
		multiple: multiple,
	}
}

func (g *Generator) NextValue() float64 {
	for {
		g.previous = math.Mod(g.previous*g.factor, modValue)

		if g.multiple == 1 || math.Mod(g.previous, g.multiple) == 0 {
			return g.previous
		}
	}
}
