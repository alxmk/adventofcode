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
		log.Fatalln("Failed to read input data:", err)
	}

	var part1, part2, i, freq int
	tracker := make(map[int]struct{})
	lines := strings.Split(string(data), "\n")
	for {
		line := lines[i%len(lines)]
		num, err := strconv.Atoi(line)
		if err != nil {
			log.Fatalln("Failed to parse", line, err)
		}

		freq += num

		if i/len(lines) == 0 {
			part1 = freq
		}

		if _, ok := tracker[freq]; ok {
			part2 = freq
			break
		}

		tracker[part1] = struct{}{}

		i++
	}

	log.Println(part1, part2)
}
