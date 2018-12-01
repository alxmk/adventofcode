package main

import (
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

var (
	Alphabet = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
)

func main() {
	input, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatalf("Error reading input file %v", err)
	}

	moves := strings.Split(string(input), ",")

	line := NewLine(16)

	dance(line, moves)

	log.Println("Part one answer is", string(line.orderLul))

	line = NewLine(16)
	var danceNum, repeatInterval int
	seen := make(map[string]int)
	ordered := make(map[int]string)

	log.Println("Finding where the sequence repeats")

	for {
		log.Println(danceNum, string(line.orderLul))
		if first, ok := seen[string(line.orderLul)]; !ok {
			seen[string(line.orderLul)] = danceNum
			ordered[danceNum] = string(line.orderLul)
		} else {
			repeatInterval = danceNum - first
			break
		}

		dance(line, moves)

		danceNum++
	}
	log.Println("Found repeat at", danceNum, "interval is", repeatInterval)

	// Which iteration do we need
	n := 1000000000 - (repeatInterval * (1000000000 / repeatInterval))

	log.Println(n)

	log.Println("Part two answer is", ordered[n])
}

func dance(line *Line, moves []string) {
	for _, m := range moves {
		switch {
		case strings.Contains(m, "x"):
			line.Exchange(m)
		case strings.Contains(m, "s"):
			line.Spin(m)
		case strings.Contains(m, "p"):
			line.Partner(m)
		default:
			log.Fatalln("Unrecognised instruction", m)
		}
	}
}

type Line struct {
	orderLul []rune
}

func NewLine(length int) *Line {
	order := make([]rune, 16)
	copy(order, Alphabet[0:length])
	return &Line{
		orderLul: order,
	}
}

func (l *Line) Exchange(move string) {
	swappees := strings.Split(strings.TrimPrefix(move, "x"), "/")

	// Parse the indexes
	s1, _ := strconv.Atoi(swappees[0])
	s2, _ := strconv.Atoi(swappees[1])

	// Store one so we don't overwrite
	cache := l.orderLul[s1]

	// Swap
	l.orderLul[s1] = l.orderLul[s2]
	l.orderLul[s2] = cache
}

func (l *Line) Partner(move string) {
	swappees := strings.Split(strings.TrimPrefix(move, "p"), "/")

	var s1 int
	var s2 int

	var found1 bool
	var found2 bool

	// Find indexes of the specified partners
	for i, c := range l.orderLul {
		if string(c) == swappees[0] {
			s1 = i
			found1 = true
		}
		if string(c) == swappees[1] {
			s2 = i
			found2 = true
		}
		if found1 && found2 {
			break
		}
	}

	// Store one so we don't overwrite
	cache := l.orderLul[s1]

	// Swap
	l.orderLul[s1] = l.orderLul[s2]
	l.orderLul[s2] = cache
}

func (l *Line) Spin(move string) {
	numberToSpin, _ := strconv.Atoi(strings.TrimPrefix(move, "s"))

	// Start with the required number from the end of the old order
	newOrder := l.orderLul[len(l.orderLul)-numberToSpin : len(l.orderLul)]

	// Add the rest on to the end of the new order
	newOrder = append(newOrder, l.orderLul[:len(l.orderLul)-numberToSpin]...)

	l.orderLul = newOrder
}
