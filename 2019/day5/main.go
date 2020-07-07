package main

import (
	"flag"
	"io/ioutil"
	"log"

	"github.ibm.com/alexmk/adventofcode/2019/day2/intcode"
)

func main() {
	var input int64
	flag.Int64Var(&input, "input", -1, "Input value to programme")
	flag.Parse()

	data, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatalln("Failed to read input file:", err)
	}

	p, err := intcode.Parse(string(data))
	if err != nil {
		log.Fatalln("Bad input:", err)
	}

	in, out := make(chan int64), make(chan int64)

	go func() {
		in <- input
		close(in)
	}()

	go func() {
		if err := p.Run(in, out); err != nil {
			log.Fatalln("Error running programme:", err)
		}
	}()

	for output := range out {
		log.Println(output)
	}
}
