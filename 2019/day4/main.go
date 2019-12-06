package main

import (
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func main() {
	data, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatalln("Failed to read input file:", err)
	}

	parts := strings.Split(string(data), "-")
	if l := len(parts); l != 2 {
		log.Fatalln("Expected two parts in the range input, got", l)
	}

	from, err := strconv.Atoi(parts[0])
	if err != nil {
		log.Fatalf("Failed to parse %s as int: %s", parts[0], err)
	}

	to, err := strconv.Atoi(parts[1])
	if err != nil {
		log.Fatalf("Failed to parse %s as int: %s", parts[1], err)
	}

	log.Println("Part one:", len(findPasswords(from, to, isPasswordAnyAdjacent)))
	log.Println("Part two:", len(findPasswords(from, to, isPasswordExactlyDouble)))
}

func findPasswords(from, to int, isPassword func(int) bool) []int {
	var passwords []int
	for i := from; i <= to; i++ {
		if isPassword(i) {
			passwords = append(passwords, i)
		}
	}
	return passwords
}

func isPasswordAnyAdjacent(num int) bool {
	var double bool
	var previous int
	for i := 100000; i >= 1; i /= 10 {
		digit := num / i
		if digit == previous {
			double = true
		}
		if digit < previous {
			return false
		}
		num -= i * digit
		previous = digit
	}
	return double
}

func isPasswordExactlyDouble(num int) bool {
	var double bool
	var previous int
	values := make(map[int]int)
	for i := 100000; i >= 1; i /= 10 {
		digit := num / i
		if digit == previous {
			double = true
		}
		if digit < previous {
			return false
		}
		num -= i * digit
		previous = digit
		values[digit]++
	}
	if !double {
		return false
	}

	for _, n := range values {
		if n == 2 {
			return true
		}
	}

	return false
}
