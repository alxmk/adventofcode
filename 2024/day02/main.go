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
	var nsafeUndamped, nsafeDamped int
	for _, line := range strings.Split(input, "\n") {
		var r report
		for _, level := range strings.Fields(line) {
			l, _ := strconv.Atoi(level)
			r = append(r, l)
		}
		if r.safe(false) {
			nsafeUndamped++
			nsafeDamped++
			continue
		}
		if r.safe(true) {
			nsafeDamped++
		}
	}
	fmt.Println("Part one:", nsafeUndamped)
	fmt.Println("Part two:", nsafeDamped)
}

type report []int

func (r report) safe(damped bool) bool {
	return r.check(damped, true) || r.check(damped, false)
}

func (r report) check(damped, dir bool) bool {
	c := r[0]
	var removed bool
	for i := 1; i < len(r); i++ {
		diff := c - r[i]
		if diff == 0 || diff > 3 || diff < -3 || c > r[i] != dir {
			if !damped || removed {
				// Special case the very first level being the one to remove
				return damped && r[1:].check(false, dir)
			}
			removed = true
			continue
		}
		c = r[i]
	}
	return true
}
