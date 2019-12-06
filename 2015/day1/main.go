package main

import (
	"io/ioutil"
	"log"
)

func main() {
	data, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatalln("Failed to read input file", err)
	}

	var floor int
	for i, c := range string(data) {
		switch c {
		case '(':
			floor++
		case ')':
			floor--
		default:
			log.Fatalln("Unknown instruction:", c)
		}
		if floor == -1 {
			log.Println("Part two:", i+1)
		}
	}

	log.Println("Part one:", floor)
}
