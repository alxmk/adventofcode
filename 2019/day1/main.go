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
		log.Fatalln("Failed to read input file:", err)
	}

	var basicFuel int
	var totalFuel int

	for _, line := range strings.Split(string(data), "\n") {
		moduleMass, err := strconv.Atoi(line)
		if err != nil {
			log.Fatalf("Failed to parse %s as int: %v", line, err)
		}

		basicFuel += fuelFor(moduleMass)
		totalFuel += fuelRecurse(moduleMass)
	}

	log.Println("Part one:", basicFuel)
	log.Println("Part two:", totalFuel)
}

func fuelFor(mass int) int {
	return (mass / 3) - 2
}

func fuelRecurse(mass int) int {
	if fuelMass := fuelFor(mass); fuelMass > 0 {
		return fuelMass + fuelRecurse(fuelMass)
	}
	return 0
}
