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
		log.Fatalln("Failed to read input file", err)
	}

	lights := make([][]bool, 1000)
	brightnesses := make([][]int, 1000)
	for column := 0; column < 1000; column++ {
		lights[column] = make([]bool, 1000)
		brightnesses[column] = make([]int, 1000)
	}

	for _, line := range strings.Split(string(data), "\n") {
		var modFunc modifier
		var brightFunc brightnessModifier
		switch {
		case strings.HasPrefix(line, "turn off"):
			modFunc = turnOff
			brightFunc = turnDown
			line = strings.TrimPrefix(line, "turn off ")
		case strings.HasPrefix(line, "turn on"):
			modFunc = turnOn
			brightFunc = turnUp
			line = strings.TrimPrefix(line, "turn on ")
		case strings.HasPrefix(line, "toggle"):
			modFunc = toggle
			brightFunc = turnUpTwice
			line = strings.TrimPrefix(line, "toggle ")
		default:
			log.Fatalln("Undefined instruction:", line)
		}
		var xmin, ymin, xmax, ymax int
		fmt.Sscanf(line, "%d,%d through %d,%d", &xmin, &ymin, &xmax, &ymax)

		for x := xmin; x <= xmax; x++ {
			for y := ymin; y <= ymax; y++ {
				lights[x][y] = modFunc(lights[x][y])
				brightnesses[x][y] = brightFunc(brightnesses[x][y])
			}
		}
	}

	var numOn int
	var totalBrightness int
	for x := 0; x < len(lights); x++ {
		for y := 0; y < len(lights[x]); y++ {
			if lights[x][y] {
				numOn++
			}
			totalBrightness += brightnesses[x][y]
		}
	}

	log.Println("Part one:", numOn)
	log.Println("Part two:", totalBrightness)
}

type modifier func(bool) bool

func turnOff(current bool) bool {
	return false
}

func turnOn(current bool) bool {
	return true
}

func toggle(current bool) bool {
	return !current
}

type brightnessModifier func(int) int

func turnDown(current int) int {
	if current > 0 {
		return current - 1
	}
	return 0
}

func turnUp(current int) int {
	return current + 1
}

func turnUpTwice(current int) int {
	return current + 2
}
