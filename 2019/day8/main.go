package main

import (
	"io/ioutil"
	"log"
	"math"
	"strings"
)

func main() {
	data, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatalln("Failed to read input file:", err)
	}

	zeroCount := make(map[int]int)
	oneCount := make(map[int]int)
	twoCount := make(map[int]int)

	img := make(image)

	for i, r := range string(data) {
		layerNum := i / (25 * 6)
		x := i % 25
		y := (i / 25) % 6
		switch r {
		case '0':
			zeroCount[layerNum]++
		case '1':
			oneCount[layerNum]++
		case '2':
			twoCount[layerNum]++
			continue
		}

		if _, ok := img[y]; !ok {
			img[y] = make(map[int]rune)
		}
		if _, ok := img[y][x]; !ok {
			img[y][x] = r
		}
	}

	fewest := math.MaxInt32
	var layerWithFewest int
	for layer, count := range zeroCount {
		if count < fewest {
			fewest = count
			layerWithFewest = layer
		}
	}

	log.Println("Part one:", oneCount[layerWithFewest]*twoCount[layerWithFewest])
	log.Printf("Part two:\n%s", img)
}

type image map[int]map[int]rune

func (i image) String() string {
	var b strings.Builder

	for y := 0; y < len(i); y++ {
		for x := 0; x < len(i[y]); x++ {
			switch i[y][x] {
			case '0':
				b.WriteRune(' ')
			case '1':
				b.WriteRune('#')
			}

		}
		b.WriteRune('\n')
	}

	return b.String()
}
