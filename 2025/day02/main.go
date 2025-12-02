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
	var partOne, partTwo int
	for _, rng := range strings.Split(input, ",") {
		split := strings.Index(rng, "-")
		start, _ := strconv.Atoi(rng[:split])
		end, _ := strconv.Atoi(rng[split+1:])
		for i := start; i <= end; i++ {
			v := strconv.Itoa(i)
			pfs := primeFactors(len(v))
			if isValid(v, pfs) {
				continue
			}
			if slices.Contains(pfs, 2) {
				partOne += i
			}
			partTwo += i
		}
	}
	fmt.Println("Part one:", partOne)
	fmt.Println("Part two:", partTwo)
}

func isValid(v string, pfs []int) bool {
outer:
	for _, pf := range pfs {
		if pf == len(v) {
			continue
		}
		pattern := v[:pf]
		for idx := pf; idx < len(v); idx += pf {
			if v[idx:idx+pf] != pattern {
				continue outer
			}
		}
		return false
	}
	return true
}

func primeFactors(n int) []int {
	var fs []int
	for i := 1; i <= n; i++ {
		if n%i == 0 {
			fs = append(fs, i)
			n /= i
		}
	}
	return fs
}
