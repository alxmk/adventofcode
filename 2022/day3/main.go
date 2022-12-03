package main

import (
	"log"
	"os"
	"strings"
)

func main() {
	data, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatalln("Failed to read input:", err)
	}

	lines := strings.Split(string(data), "\n")
	var sum int
	for _, backpack := range lines {
		sum += score(backpack)
	}
	log.Println("Part one:", sum)
	sum = 0
	for i := 0; i < len(lines); i += 3 {
		sum += itemScore(common(lines[i : i+3]))
	}
	log.Println("Part two:", sum)
}

func score(backpack string) int {
	for _, itemL := range backpack[:len(backpack)/2+1] {
		for _, itemR := range backpack[len(backpack)/2:] {
			if itemL == itemR {
				return itemScore(itemL)
			}
		}
	}
	return -1
}

func itemScore(item rune) int {
	if item >= 'a' {
		return int(item - 'a' + 1)
	}
	return int(item - 'A' + 27)
}

func common(backpacks []string) rune {
	items := make(map[rune]int)
	for i, b := range backpacks {
		for _, item := range b {
			if i == 0 {
				items[item] = i
				continue
			}
			if j, ok := items[item]; ok && j == i-1 {
				if i == 2 {
					return item
				}
				items[item] = i
			}
		}
	}
	return '0'
}
