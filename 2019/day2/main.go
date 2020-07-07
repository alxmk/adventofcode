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

	init, err := intcode.Parse(string(data))
	if err != nil {
		log.Fatalln("Bad input:", err)
	}

	prog := init.Copy()

	prog[1] = 12
	prog[2] = 2

	if err := prog.Run(nil, make(chan int64)); err != nil {
		log.Println("Failed to run part one:", err)
	}

	log.Println("Part one:", prog[0])

	noun, verb, err := findNV(init, 19690720)
	if err != nil {
		log.Fatalln("Failed to find NV:", err)
	}

	log.Println("Part two:", (100*noun)+verb)
}

func findNV(prog intcode.Programme, target int64) (int64, int64, error) {
	var noun, verb int64
	for ; noun < 100; noun++ {
		for verb = 0; verb < 100; verb++ {
			p := prog.Copy()
			p[1] = noun
			p[2] = verb
			if err := p.Run(nil, make(chan int64)); err != nil {
				return -1, -1, fmt.Errorf("failed to run programme with noun %d verb %d: %v", noun, verb, err)
			}

			if p[0] == target {
				return noun, verb, nil
			}
		}
	}
	return -1, -1, fmt.Errorf("couldn't hit the target for any values of NV")
}
