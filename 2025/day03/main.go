package main

import (
	_ "embed"
	"fmt"
	"math"
	"strings"
)

//go:embed input.txt
var input string

func main() {
	var partOne, partTwo int
	for _, line := range strings.Split(input, "\n") {
		partOne += getDigits(line, 2)
		partTwo += getDigits(line, 12)
	}
	fmt.Println("Part one:", partOne)
	fmt.Println("Part two:", partTwo)
}

func getDigits(line string, n int) (digits int) {
	var digit, start int
	for pow := n - 1; pow >= 0; pow-- {
		digit, start = getDigit(line, start, pow)
		digits += digit * int(math.Pow10(pow))
	}
	return digits
}

func getDigit(line string, start, pow int) (digit, idx int) {
	for i := start; i < len(line)-pow; i++ {
		if v := int(line[i] - '0'); v > digit {
			idx, digit = i, v
			if digit == 9 {
				return digit, idx + 1
			}
		}
	}
	return digit, idx + 1
}
