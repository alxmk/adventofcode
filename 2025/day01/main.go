package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

func main() {
	position := 50
	var partOne, partTwo int
	for _, line := range strings.Split(input, "\n") {
		clicks, _ := strconv.Atoi(line[1:])
		var next int
		switch line[0] {
		case 'R':
			next = position + clicks
			partTwo += next / 100
		case 'L':
			next = position - clicks
			switch {
			case position == 0:
				partTwo += clicks / 100
			case clicks >= position:
				partTwo += 1 + (clicks-position)/100
			}
		}
		next %= 100
		if next < 0 {
			next += 100
		}
		if next == 0 {
			partOne++
		}
		position = next
	}
	fmt.Println("Part one:", partOne)
	fmt.Println("Part two:", partTwo)
}
