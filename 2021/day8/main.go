package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"sort"
	"strconv"
	"strings"
)

func main() {
	data, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatalln("Failed to read input:", err)
	}

	uspLines, outputLines, err := parse(string(data))
	if err != nil {
		log.Fatalln("Failed to parse input:", err)
	}

	log.Println("Part one:", partOne(outputLines))
	log.Println("Part two:", partTwo(uspLines, outputLines))
}

type digit map[rune]struct{}

func (d digit) String() string {
	var slice []string
	for r := range d {
		slice = append(slice, string(r))
	}
	sort.Strings(slice)
	return strings.Join(slice, "")
}

func parse(input string) ([][]digit, [][]digit, error) {
	var uspLines, outputLines [][]digit
	for _, line := range strings.Split(input, "\n") {
		parts := strings.Split(line, " | ")
		if len(parts) != 2 {
			return nil, nil, fmt.Errorf("malformed line %s", line)
		}
		uspLines = append(uspLines, parseDigits(parts[0]))
		outputLines = append(outputLines, parseDigits(parts[1]))
	}
	return uspLines, outputLines, nil
}

func parseDigits(raw string) []digit {
	var digits []digit
	for _, rawDigit := range strings.Fields(raw) {
		digits = append(digits, parseDigit(rawDigit))
	}
	return digits
}

func parseDigit(raw string) digit {
	d := make(map[rune]struct{})
	for _, r := range raw {
		d[r] = struct{}{}
	}
	return d
}

func partOne(outputLines [][]digit) int {
	var count int
	for _, output := range outputLines {
		for _, d := range output {
			switch len(d) {
			case 2, 3, 4, 7:
				// log.Println(d)
				count++
			}
		}
	}
	return count
}

func partTwo(uspLines, outputLines [][]digit) int {
	var sum int
	for i := range uspLines {
		sum += outputValue(uspLines[i], outputLines[i])
	}
	return sum
}

func outputValue(usp []digit, output []digit) int {
	numbers := make(map[int]digit)
	six, five := make(map[string]struct{}), make(map[string]struct{})
	for _, d := range append(usp, output...) {
		switch len(d) {
		// 0, 6, 9
		case 6:
			six[d.String()] = struct{}{}
		// 1
		case 2:
			numbers[1] = d
		// 2, 3, 5
		case 5:
			five[d.String()] = struct{}{}
		// 7
		case 3:
			numbers[7] = d
		// 4
		case 4:
			numbers[4] = d
		// 8
		case 7:
			numbers[8] = d
		}
	}

	for k := range six {
		// 9 is 4 with two extra bars
		if commonCharacters(k, numbers[4].String()) == 4 {
			numbers[9] = parseDigit(k)
			continue
		}
		// 8 is 0 minus one bar, and has only one common with 1
		if commonCharacters(numbers[8].String(), k) == 6 && commonCharacters(numbers[1].String(), k) == 2 {
			numbers[0] = parseDigit(k)
			continue
		}
		// Only one left is 6
		numbers[6] = parseDigit(k)
	}

	for k := range five {
		// 3 is 1 with 3 extra bars
		if commonCharacters(k, numbers[1].String()) == 2 {
			numbers[3] = parseDigit(k)
			continue
		}
		// 5 is 6 minus one bar
		if commonCharacters(k, numbers[6].String()) == 5 {
			numbers[5] = parseDigit(k)
			continue
		}
		// Only one left is 2
		numbers[2] = parseDigit(k)
	}

	var result string
	for _, o := range output {
		for k, n := range numbers {
			if n.String() == o.String() {
				result += fmt.Sprintf("%d", k)
			}
		}
	}

	n, _ := strconv.Atoi(result)

	return n
}

func commonCharacters(a, b string) int {
	if a == b {
		return 0
	}
	var common int
	for _, ra := range a {
		var contained bool
		for _, rb := range b {
			if ra == rb {
				contained = true
				break
			}
		}
		if contained {
			common++
		}
	}
	return common
}
