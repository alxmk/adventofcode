package main

import (
	_ "embed"
	"fmt"
	"strings"
)

//go:embed input.txt
var input string

func main() {
	grid := make(map[xy]rune)
	var xmax, ymax int
	for y, line := range strings.Split(input, "\n") {
		for x, r := range line {
			grid[xy{x, y}] = r
			xmax = max(x, xmax)
		}
		ymax = y
	}
	var count int
	count, grid = reduce(grid, xmax, ymax)
	fmt.Println("Part one:", count)
	for removed, grid := reduce(grid, xmax, ymax); removed != 0; removed, grid = reduce(grid, xmax, ymax) {
		count += removed
	}
	fmt.Println("Part two:", count)
}

type xy [2]int

func reduce(grid map[xy]rune, xmax, ymax int) (count int, next map[xy]rune) {
	next = make(map[xy]rune)
	for x := 0; x <= xmax; x++ {
		for y := 0; y <= ymax; y++ {
			if grid[xy{x, y}] == '.' {
				next[xy{x, y}] = '.'
				continue
			}
			var adjacent int
			func() {
				for i := max(x-1, 0); i <= min(xmax, x+1); i++ {
					for j := max(y-1, 0); j <= min(ymax, y+1); j++ {
						if i == x && j == y {
							continue
						}
						if grid[xy{i, j}] == '@' {
							adjacent++
							if adjacent == 4 {
								return
							}
						}
					}
				}
			}()
			if adjacent < 4 {
				next[xy{x, y}] = '.'
				count++
				continue
			}
			next[xy{x, y}] = '@'
		}
	}
	return count, next
}
