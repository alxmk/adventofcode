package main

import (
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

	nums := parse(string(data))

	d := diffs(nums)
	log.Println("Part one:", d[1]*d[3])
	log.Println("Part two:", newSolver(nums).permutations(0))
}

func parse(input string) []int {
	var out []int
	for _, line := range strings.Split(input, "\n") {
		v, err := strconv.Atoi(line)
		if err != nil {
			log.Fatalf("Failed to convert %s to int: %s", line, err)
		}
		out = append(out, v)
	}
	sort.Ints(out)
	out = append(out, out[len(out)-1]+3)
	return out
}

func diffs(nums []int) [4]int {
	diffs := [4]int{}
	var prev int
	for _, n := range nums {
		diffs[n-prev]++
		prev = n
	}
	return diffs
}

type solver struct {
	nums   map[int]struct{}
	target int
	cache  map[int]int
}

func newSolver(nums []int) *solver {
	numMap := make(map[int]struct{})
	for _, n := range nums {
		numMap[n] = struct{}{}
	}
	return &solver{
		nums:   numMap,
		cache:  make(map[int]int),
		target: nums[len(nums)-1],
	}
}

func (s *solver) permutations(start int) int {
	if start == s.target {
		return 1
	}
	var count int
	for i := 1; i <= 3; i++ {
		if _, ok := s.nums[start+i]; ok {
			if c, ok := s.cache[start+i]; ok {
				count += c
				continue
			}
			s.cache[start+i] = s.permutations(start + i)
			count += s.cache[start+i]
		}
	}
	return count
}
