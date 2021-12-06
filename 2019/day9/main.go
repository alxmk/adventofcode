package main

import (
	"io/ioutil"
	"log"

	"github.com/alxmk/adventofcode/2019/day2/intcode"
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

	in, out := make(chan int64), make(chan int64)
	go func() {
		in <- 1
		close(in)
	}()
	go func() {
		if err := programme.Copy().Run(in, out); err != nil {
			log.Fatalln("Failed to run:", err)
		}
	}()

	for output := range out {
		log.Println("Part one:", output)
	}

	in, out = make(chan int64), make(chan int64)
	go func() {
		in <- 2
		close(in)
	}()
	go func() {
		if err := programme.Copy().Run(in, out); err != nil {
			log.Fatalln("Failed to run:", err)
		}
	}()

	for output := range out {
		log.Println("Part two:", output)
	}
}
