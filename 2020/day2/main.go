package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

func main() {
	data, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatalln("Failed to read input:", err)
	}

	var countOne, countTwo int
	for _, line := range strings.Split(string(data), "\n") {
		var pass password
		fmt.Sscanf(line, "%d-%d %c: %s", &pass.A, &pass.B, &pass.Char, &pass.Password)
		if pass.ValidOne() {
			countOne++
		}
		if pass.ValidTwo() {
			countTwo++
		}
	}
	log.Println("Part one:", countOne)
	log.Println("Part two:", countTwo)
}

type password struct {
	A        int
	B        int
	Char     byte
	Password string
}

func (p password) ValidOne() bool {
	c := strings.Count(p.Password, string(p.Char))
	return c >= p.A && c <= p.B
}

func (p password) ValidTwo() bool {
	return (p.Password[p.A-1] == p.Char) != (p.Password[p.B-1] == p.Char) // xor
}
