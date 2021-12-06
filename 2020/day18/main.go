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

	log.Println("Part one:", total(string(data), sumBasic))
	log.Println("Part two:", total(string(data), sumAdvanced))
}

func total(input string, f func(string) int64) int64 {
	var total int64
	for _, line := range strings.Split(input, "\n") {
		total += f(line)
	}
	return total
}

func sumBasic(line string) int64 {
	var num int64
	var op rune
	for i := 0; i < len(line); i++ {
		r := rune(line[i])
		switch r {
		case ' ':
			continue
		case '+', '*':
			op = r
		case '(':
			var depth int
			for j, c := range line[i+1:] {
				switch c {
				case '(':
					depth++
				case ')':
					if depth == 0 {
						v := sumBasic(line[i+1 : j+i+1])
						switch op {
						case '*':
							num *= v
						default:
							num += v
						}
						i = j + i + 1
					} else {
						depth--
					}
				}
			}
		default:
			switch op {
			case '*':
				num *= int64(r - '0')
			default:
				num += int64(r - '0')
			}
		}
	}
	return num
}

func sumAdvanced(line string) int64 {
	// brackets first
	bracketsFound := true
brackets:
	for bracketsFound {
		var closeIdx int
		bracketsFound = false
		for i := len(line) - 1; i >= 0; i-- {
			r := rune(line[i])
			switch r {
			case ')':
				bracketsFound = true
				closeIdx = i
			case '(':
				subExp := line[i+1 : closeIdx]
				val := sumAdvanced(subExp)
				line = strings.ReplaceAll(line, "("+subExp+")", fmt.Sprintf("%d", val))
				continue brackets
			}
		}
	}

	// now adds
outer:
	for strings.Contains(line, "+") {
		var prev int64
		var op rune
		for _, f := range strings.Fields(line) {
			switch f {
			case "+":
				op = '+'
			case "*":
				op = '*'
			default:
				v, err := strconv.ParseInt(f, 10, 64)
				if err != nil {
					log.Fatalf("Failed to parse %s as int64: %s", f, err)
				}
				if op == '+' {
					line = strings.ReplaceAll(line, fmt.Sprintf("%d + %d", prev, v), fmt.Sprintf("%d", prev+v))
					continue outer
				}
				prev = v
			}
		}
	}

	// Finally multiply through
	var sum int64 = 1
	for _, f := range strings.Fields(line) {
		switch f {
		case "*":
			continue
		default:
			v, err := strconv.ParseInt(f, 10, 64)
			if err != nil {
				log.Fatalf("Failed to parse %s as int64: %s", f, err)
			}
			sum *= v
		}
	}
	return sum
}
