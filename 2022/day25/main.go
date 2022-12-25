package main

import (
	"bytes"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	data, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatalln("Failed to read input:", err)
	}

	log.Println("Part one:", partOne(data))
}

func partOne(input []byte) string {
	var sum int
	for _, line := range bytes.Split(input, []byte{'\n'}) {
		sum += parse(line)
	}
	return convert(sum)
}

func parse(snafu []byte) int {
	var v int
	for i, j := len(snafu)-1, 0; i >= 0; i, j = i-1, j+1 {
		pow := int(math.Pow(5, float64(j)))
		switch snafu[i] {
		case '2':
			v += 2 * pow
		case '1':
			v += pow
		case '0':
		case '-':
			v -= pow
		case '=':
			v -= 2 * pow
		}
	}
	return v
}

func convert(v int) string {
	b5 := []byte(strconv.FormatInt(int64(v), 5))
	for strings.Contains(string(b5), "3") || strings.Contains(string(b5), "4") {
		for i := 0; i < len(b5); i++ {
			switch b5[i] {
			case '0', '1', '2':
				continue
			case '3':
				b5[i] = '='
				if i-1 < 0 {
					b5, i = append([]byte{'1'}, b5...), i+1
					continue
				}
				switch b5[i-1] {
				case '0', '1', '2':
					b5[i-1] = b5[i-1] + 1
				case '-':
					b5[i-1] = '0'
				case '=':
					b5[i-1] = '-'
				}
			case '4':
				b5[i] = '-'
				if i-1 < 0 {
					b5, i = append([]byte{'1'}, b5...), i+1
					continue
				}
				switch b5[i-1] {
				case '0', '1', '2':
					b5[i-1] = b5[i-1] + 1
				case '-':
					b5[i-1] = '0'
				case '=':
					b5[i-1] = '-'
				}
			}
		}
	}
	return string(b5)
}
