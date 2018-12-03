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

	sheet := make(map[int]map[int]int)

	for _, line := range strings.Split(string(data), "\n") {
		xstart, ystart, width, depth := parseCoords(line)

		for x := xstart; x < xstart+width; x++ {
			if _, ok := sheet[x]; !ok {
				sheet[x] = make(map[int]int)
			}
			for y := ystart; y < ystart+depth; y++ {
				sheet[x][y]++
			}
		}
	}

	var contested int
	for _, ymap := range sheet {
		for _, count := range ymap {
			if count > 1 {
				contested++
			}
		}
	}

outer:
	for _, line := range strings.Split(string(data), "\n") {
		xstart, ystart, width, depth := parseCoords(line)

		for x := xstart; x < xstart+width; x++ {
			for y := ystart; y < ystart+depth; y++ {
				if sheet[x][y] != 1 {
					continue outer
				}
			}
		}

		log.Println(line)
	}

	log.Println(contested)
}

func parseCoords(line string) (int, int, int, int) {
	parts := strings.Fields(line)
	xy := strings.Split(parts[2], ",")

	xstart, err := strconv.Atoi(xy[0])
	if err != nil {
		log.Fatalln("Failed to parse as int", xy[0], err)
	}
	ystart, err := strconv.Atoi(xy[1][:len(xy[1])-1])
	if err != nil {
		log.Fatalln("Failed to parse as int", xy[1], err)
	}

	wd := strings.Split(parts[3], "x")

	width, err := strconv.Atoi(wd[0])
	if err != nil {
		log.Fatalln("Failed to parse as int", wd[0], err)
	}

	depth, err := strconv.Atoi(wd[1])
	if err != nil {
		log.Fatalln("Failed to parse as int", wd[1], err)
	}

	return xstart, ystart, width, depth
}
