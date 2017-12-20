package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func main() {
	input, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatalf("Error reading input file %v", err)
	}

	lengths := strings.Split(string(input), ",")

	loop := newLoop(256)
	skip := 0
	pos := 0

	for _, l := range lengths {
		value, err := strconv.Atoi(l)
		if err != nil {
			log.Fatalf("Couldn't convert %s to int", l)
		}

		loop.Reverse(pos, value)

		pos += value + skip

		if pos > len(loop.contents) {
			pos = pos - len(loop.contents)
		}

		skip++
	}

	log.Println("The answer is", loop.contents[0]*loop.contents[1])
}

type loop struct {
	contents map[int]int
}

func newLoop(length int) loop {
	l := loop{
		contents: make(map[int]int),
	}
	for i := 0; i < length; i++ {
		l.contents[i] = i
	}

	return l
}

func (l *loop) Reverse(pos, length int) {
	newContents := make(map[int]int)
	// Copy the old contents
	for k, v := range l.contents {
		newContents[k] = v
	}
	for i := 0; i < length; i++ {
		oldIndex := pos + i
		for oldIndex > len(l.contents)-1 {
			oldIndex = oldIndex - len(l.contents)
		}

		newIndex := pos + length - i - 1
		for newIndex > len(l.contents)-1 {
			newIndex = newIndex - len(l.contents)
		}

		newContents[newIndex] = l.contents[oldIndex]
	}

	l.contents = newContents
}

func (l *loop) Print() string {
	out := ""

	for i := 0; i < len(l.contents); i++ {
		out += fmt.Sprintf("%d,", l.contents[i])
	}

	return out
}
