package main

import (
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	data, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatalln("Failed to read input:", err)
	}

	rows := parseRows(string(data))

	log.Println("Part one:", partOne(rows))
	log.Println("Part two:", partTwo(rows))
}

func partOne(rows []row) int {
	var sum int
	for _, r := range rows {
		p := r.Permutations(0, 0, 0, make(map[[3]int]int))
		sum += p
	}
	return sum
}

func partTwo(rows []row) int {
	var sum int
	for _, r := range rows {
		p := r.Unfold().Permutations(0, 0, 0, make(map[[3]int]int))
		sum += p
	}
	return sum
}

func parseRows(input string) []row {
	var rows []row
	for _, line := range strings.Split(input, "\n") {
		rows = append(rows, parseRow(line))
	}
	return rows
}

func parseRow(line string) row {
	fields := strings.Fields(line)
	r := row{report: []rune(fields[0])}
	for _, group := range strings.Split(fields[1], ",") {
		v, _ := strconv.Atoi(group)
		r.groups = append(r.groups, v)
		r.totaldamaged += v
	}
	return r
}

type row struct {
	report       []rune
	groups       []int
	totaldamaged int
}

func (r row) Unfold() row {
	r.report = append([]rune{'?'}, r.report...)
	l, m := len(r.report), len(r.groups)
	for i := 0; i < 4; i++ {
		for j := 0; j < l; j++ {
			r.report = append(r.report, r.report[j])
		}
		for j := 0; j < m; j++ {
			r.groups = append(r.groups, r.groups[j])
		}
	}
	r.report = r.report[1:]
	return r
}

func (r row) Permutations(ir, ig, sequence int, cache map[[3]int]int) int {
	if p, ok := cache[[3]int{ir, ig, sequence}]; ok {
		return p
	}
	if ir == len(r.report) {
		if ig == len(r.groups) && sequence == 0 || ig == len(r.groups)-1 && r.groups[ig] == sequence {
			return 1
		}
		return 0
	}
	var perms int
	switch r.report[ir] {
	case '.':
		if sequence == 0 {
			perms = r.Permutations(ir+1, ig, 0, cache)
		}
		if sequence > 0 && ig < len(r.groups) && r.groups[ig] == sequence {
			perms = r.Permutations(ir+1, ig+1, 0, cache)
		}
	case '#':
		perms = r.Permutations(ir+1, ig, sequence+1, cache)
	case '?':
		if sequence == 0 {
			perms = r.Permutations(ir+1, ig, 0, cache) + r.Permutations(ir+1, ig, sequence+1, cache)
		} else {
			if ig < len(r.groups) && r.groups[ig] == sequence {
				perms = r.Permutations(ir+1, ig+1, 0, cache)
			}
			perms += r.Permutations(ir+1, ig, sequence+1, cache)
		}
	}
	cache[[3]int{ir, ig, sequence}] = perms
	return perms
}
