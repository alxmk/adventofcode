package main

import (
	"fmt"
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

	numbers, boards, err := parse(string(data))
	if err != nil {
		log.Fatalln("Failed to parse input:", err)
	}

	log.Println("Part one:", findWinner(numbers, boards))
	log.Println("Part two:", findLoser(numbers, boards))
}

type bingoBoard map[coord]square

type coord struct{ x, y int }

type square struct {
	value   int
	checked bool
}

func parse(input string) ([]int, []bingoBoard, error) {
	var numbers []int
	var boards []bingoBoard
	for i, part := range strings.Split(input, "\n\n") {
		if i == 0 {
			for _, n := range strings.Split(part, ",") {
				v, err := strconv.Atoi(n)
				if err != nil {
					return nil, nil, fmt.Errorf("failed to parse number %s: %s", n, err)
				}
				numbers = append(numbers, v)
			}
			continue
		}
		board, err := parseBoard(part)
		if err != nil {
			return nil, nil, fmt.Errorf("failed to parse %s as board: %s", part, err)
		}
		boards = append(boards, board)
	}
	return numbers, boards, nil
}

func parseBoard(input string) (bingoBoard, error) {
	board := make(bingoBoard)
	for y, line := range strings.Split(input, "\n") {
		for x, num := range strings.Fields(line) {
			v, err := strconv.Atoi(num)
			if err != nil {
				return nil, fmt.Errorf("failed to parse %s: %s", num, err)
			}
			board[coord{x: x, y: y}] = square{value: v}
		}
	}
	return board, nil
}

func findWinner(numbers []int, boards []bingoBoard) int {
	for _, number := range numbers {
		for _, b := range boards {
			b.Play(number)
			if b.Wins() {
				return number * b.Score()
			}
		}
	}
	return -1
}

func (b bingoBoard) Play(number int) {
	for k, v := range b {
		if v.value == number {
			b[k] = square{value: number, checked: true}
			return
		}
	}
}

func (b bingoBoard) Wins() bool {
	// Check columns
	for x := 0; x < 5; x++ {
		var missing bool
		for y := 0; y < 5; y++ {
			if !b[coord{x: x, y: y}].checked {
				missing = true
				break
			}
		}
		if !missing {
			return true
		}
	}
	// Check rows
	for y := 0; y < 5; y++ {
		var missing bool
		for x := 0; x < 5; x++ {
			if !b[coord{x: x, y: y}].checked {
				missing = true
				break
			}
		}
		if !missing {
			return true
		}
	}
	return false
}

func (b bingoBoard) Score() int {
	var score int
	for x := 0; x < 5; x++ {
		for y := 0; y < 5; y++ {
			if v := b[coord{x: x, y: y}]; !v.checked {
				score += v.value
			}
		}
	}
	return score
}

func findLoser(numbers []int, boards []bingoBoard) int {
	bMap := make(map[int]bingoBoard)
	for i, b := range boards {
		bMap[i] = b
	}
	for _, number := range numbers {
		for i, b := range bMap {
			b.Play(number)
			if b.Wins() {
				if len(bMap) == 1 {
					return number * b.Score()
				}
				delete(bMap, i)
			}
		}

	}
	return -1
}
