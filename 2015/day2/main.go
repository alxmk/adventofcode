package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"sort"
	"strings"
)

func main() {
	data, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatalln("Failed to read input file", err)
	}

	var totalSize int
	var totalBow int
	for _, line := range strings.Split(string(data), "\n") {
		var l, w, h int
		fmt.Sscanf(line, "%dx%dx%d", &l, &w, &h)
		totalSize += sizeFor(l, w, h)
		totalBow += bowFor(l, w, h)
	}

	log.Println("Part one:", totalSize)
	log.Println("Part two:", totalBow)
}

func sizeFor(l, w, h int) int {
	base := 2*l*w + 2*w*h + 2*h*l

	lengths := []int{l, w, h}
	sort.Sort(sort.IntSlice(lengths))

	return base + lengths[0]*lengths[1]
}

func bowFor(l, w, h int) int {
	lengths := []int{l, w, h}
	sort.Sort(sort.IntSlice(lengths))

	return l*w*h + 2*(lengths[0]+lengths[1])
}
