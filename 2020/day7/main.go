package main

import (
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func main() {
	data, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatalln("Failed to read input:", err)
	}

	bags := parse(string(data))

	var count int
	for _, b := range bags {
		if b.find("shiny gold") {
			count++
		}
	}
	log.Println("Part one:", count)
	log.Println("Part two:", bags["shiny gold"].inside())
}

func parse(input string) map[string]*bag {
	bags := make(map[string]*bag)
	for _, line := range strings.Split(input, "\n") {
		halves := strings.Split(line, " bags contain ")
		name := halves[0]
		thisBag, ok := bags[name]
		if !ok {
			thisBag = &bag{name: name, contains: make(map[*bag]int)}
			bags[name] = thisBag
		}
		for _, contains := range strings.Split(strings.TrimSuffix(halves[1], "."), ", ") {
			if contains == "no other bags" {
				continue
			}
			parts := strings.Fields(contains)
			v, err := strconv.Atoi(parts[0])
			if err != nil {
				log.Fatalln("Failed to convert to int:", err)
			}
			subBagName := strings.Join(parts[1:3], " ")
			subBag, ok := bags[subBagName]
			if !ok {
				subBag = &bag{name: subBagName, contains: make(map[*bag]int)}
				bags[subBagName] = subBag
			}
			thisBag.contains[subBag] = v
		}
	}
	return bags
}

type bag struct {
	name     string
	contains map[*bag]int
}

func (b bag) find(target string) bool {
	for c := range b.contains {
		if c.name == target || c.find(target) {
			return true
		}
	}
	return false
}

func (b bag) inside() int {
	var inner int
	for c, count := range b.contains {
		inner += count + (count * c.inside())
	}
	return inner
}
