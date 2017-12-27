package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"strconv"
	"strings"
)

var (
	standardSuffix = []int{17, 31, 73, 47, 23}
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

	lengthInts := []int{}

	for _, l := range lengths {
		value, err := strconv.Atoi(l)
		if err != nil {
			log.Fatalf("Couldn't convert %s to int", l)
		}

		lengthInts = append(lengthInts, value)
	}

	loop, skip, pos = oneRound(loop, skip, pos, lengthInts)

	log.Println("Part one answer is", loop.contents[0]*loop.contents[1])
	log.Println("Part two answer is", KnotHash(string(input)))
}

func KnotHash(input string) string {
	lengthsASCII := []int{}

	for _, c := range []rune(input) {
		lengthsASCII = append(lengthsASCII, int(c))
	}

	lengthsASCII = append(lengthsASCII, standardSuffix...)

	loop := newLoop(256)
	skip := 0
	pos := 0

	for i := 0; i < 64; i++ {
		loop, skip, pos = oneRound(loop, skip, pos, lengthsASCII)
	}

	// Will never error in this context
	hash, _ := loop.DenseHash()

	return hash
}

func oneRound(l loop, skip, pos int, lengths []int) (loop, int, int) {
	for _, length := range lengths {
		l.Reverse(pos, length)

		pos += length + skip

		if pos > len(l.contents) {
			pos = pos - len(l.contents)
		}

		skip++
	}

	return l, skip, pos
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

func (l *loop) DenseHash() (string, error) {
	if math.Mod(float64(len(l.contents)), 16) != 0 {
		return "", fmt.Errorf("Cannot compute the dense hash of a loop which is not a factor of 16 in length")
	}

	var denseHash string

	for i := 0; i*16 < len(l.contents); i++ {
		index := i * 16
		xored := l.contents[index] ^ l.contents[index+1] ^ l.contents[index+2] ^ l.contents[index+3] ^ l.contents[index+4] ^ l.contents[index+5] ^ l.contents[index+6] ^
			l.contents[index+7] ^ l.contents[index+8] ^ l.contents[index+9] ^ l.contents[index+10] ^ l.contents[index+11] ^ l.contents[index+12] ^ l.contents[index+13] ^
			l.contents[index+14] ^ l.contents[index+15]

		hex := fmt.Sprintf("%x", xored)
		if len(hex) < 2 {
			hex = "0" + hex
		}

		denseHash += hex
	}

	return denseHash, nil
}
