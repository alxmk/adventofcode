package main

import (
	"bytes"
	"log"
	"os"
)

func main() {
	data, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatalln("Failed to read input:", err)
	}

	trees := bytes.Split(data, []byte{'\n'})

	log.Println("Part one:", visible(trees))
	log.Println("Part two:", scenic(trees))
}

func visible(trees [][]byte) int {
	visible := make([][]bool, len(trees))

	for y := 0; y < len(trees); y++ {
		visible[y] = make([]bool, len(trees[0]))
		// visible from west
		hwm := byte('0')
		for x := 0; x < len(trees[0]); x++ {
			if tree := trees[y][x]; tree > hwm || x == 0 {
				visible[y][x], hwm = true, tree
			}
			if hwm == '9' {
				break
			}
		}
		// visible from east
		hwm = byte('0')
		for x := len(trees[0]) - 1; x >= 0; x-- {
			if tree := trees[y][x]; tree > hwm || x == len(trees[0])-1 {
				visible[y][x], hwm = true, tree
			}
			if hwm == '9' {
				break
			}
		}
	}

	for x := 0; x < len(trees[0]); x++ {
		hwm := byte('0')
		// visible from north
		for y := 0; y < len(trees); y++ {
			if tree := trees[y][x]; tree > hwm || y == 0 {
				visible[y][x], hwm = true, tree
			}
			if hwm == '9' {
				break
			}
		}
		hwm = byte('0')
		// visible from south
		for y := len(trees) - 1; y >= 0; y-- {
			if tree := trees[y][x]; tree > hwm || y == len(trees)-1 {
				visible[y][x], hwm = true, tree
			}
			if hwm == '9' {
				break
			}
		}
	}

	var visibleCount int
	for y := 0; y < len(visible); y++ {
		for x := 0; x < len(visible[0]); x++ {
			if visible[y][x] {
				visibleCount++
			}
		}
	}
	return visibleCount
}

func scenic(trees [][]byte) int {
	var max int
	for y := 0; y < len(trees); y++ {
		for x := 0; x < len(trees[0]); x++ {
			if s := scenicScore(x, y, trees); s > max {
				max = s
			}
		}
	}
	return max
}

func scenicScore(x, y int, trees [][]byte) int {
	score := 1
	height := trees[y][x]
	// Look west
	for i := 1; ; i++ {
		if x-i < 0 {
			score *= i - 1
			break
		}
		if trees[y][x-i] < height {
			continue
		}
		score *= i
		break
	}
	// Look east
	for i := 1; ; i++ {
		if x+i > len(trees[0])-1 {
			score *= i - 1
			break
		}
		if trees[y][x+i] < height {
			continue
		}
		score *= i
		break
	}
	// Look north
	for i := 1; ; i++ {
		if y-i < 0 {
			score *= i - 1
			break
		}
		if trees[y-i][x] < height {
			continue
		}
		score *= i
		break
	}
	// Look south
	for i := 1; ; i++ {
		if y+i > len(trees)-1 {
			score *= i - 1
			break
		}
		if trees[y+i][x] < height {
			continue
		}
		score *= i
		break
	}
	return score
}
