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

	log.Println("Part one:", solve(string(data), findSymmetry(0)))
	log.Println("Part two:", solve(string(data), findSymmetry(1)))
}

func solve(input string, method func([]string) int) int {
	var sum int
	for _, pattern := range strings.Split(input, "\n\n") {
		sum += summarise(pattern, method)
	}
	return sum
}

func summarise(pattern string, f func([]string) int) int {
	lines := strings.Split(pattern, "\n")
	if s := f(lines); s != 0 {
		return 100 * s
	}
	if s := f(transform(lines)); s != 0 {
		return s
	}
	panic("No symmetry:\n" + pattern)
}

func transform(lines []string) []string {
	var columns []string
	for i := 0; i < len(lines[0]); i++ {
		var column strings.Builder
		for j := len(lines) - 1; j >= 0; j-- {
			column.WriteByte(lines[j][i])
		}
		columns = append(columns, column.String())
	}
	return columns
}

func findSymmetry(boundary int) func([]string) int {
	return func(lines []string) int {
		for i := 0; i < len(lines)-1; i++ {
			if func() bool {
				var count int
				for j := 0; i-j >= 0 && i+j+1 < len(lines); j++ {
					if lines[i-j] == lines[i+j+1] {
						continue
					}
					count += lineDiff(lines[i-j], lines[i+j+1])
				}
				return count == boundary
			}() {
				return i + 1
			}
		}
		return 0
	}
}

func lineDiff(a, b string) int {
	var count int
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			count++
		}
	}
	return count
}
