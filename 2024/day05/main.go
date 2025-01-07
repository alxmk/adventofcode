package main

import (
	_ "embed"
	"fmt"
	"slices"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

func main() {
	rules, updates := parse(input)
	fmt.Println("Part one:", partOne(rules, updates))
	fmt.Println("Part two:", partTwo(rules, updates))
}

func partOne(rules ruleList, updates []string) int {
	var count int
	for _, line := range updates {
		if v, ok, _, _ := assess(rules, line); ok {
			count += v
		}
	}
	return count
}

func partTwo(rules ruleList, updates []string) int {
	var count int
	for _, line := range updates {
		v, ok, i, j := assess(rules, line)
		if ok {
			continue
		}
		for ; !ok; v, ok, i, j = assess(rules, line) {
			pages := strings.Split(line, ",")
			cut := pages[i]
			line = strings.Join(slices.Insert(slices.Delete(pages, i, i+1), j, cut), ",")
		}
		count += v
	}
	return count
}

type ruleList map[string][]string

func parse(input string) (ruleList, []string) {
	parts := strings.Split(input, "\n\n")
	rules := make(ruleList)
	for _, line := range strings.Split(parts[0], "\n") {
		split := strings.Split(line, "|")
		rules[split[1]] = append(rules[split[1]], split[0])
	}
	return rules, strings.Split(parts[1], "\n")
}

func assess(rules ruleList, line string) (int, bool, int, int) {
	pages := strings.Split(line, ",")
	for i, page := range pages {
		if expects, ok := rules[page]; ok {
			for _, expect := range expects {
				if !strings.Contains(line, expect) {
					continue
				}

				j := slices.Index(pages, expect)
				if j < i {
					continue
				}
				return 0, false, i, j
			}
		}
	}
	v, _ := strconv.Atoi(pages[len(pages)/2])
	return v, true, -1, -1
}
