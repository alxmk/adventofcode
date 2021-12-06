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

	target := scan(nums, 25)
	log.Println("Part one:", target)
	log.Println("Part two:", weakness(nums, target))
}

func parse(input string) []uint64 {
	var nums []uint64
	for _, line := range strings.Split(input, "\n") {
		v, err := strconv.ParseUint(line, 10, 64)
		if err != nil {
			log.Fatalf("Failed to parse %s as uint64: %s", line, err)
		}
		nums = append(nums, v)
	}
	return nums
}

// scan returns the first number in the list after the preamble
// that is not a sum of two of the previous preambleSize numbers
func scan(nums []uint64, preambleSize int) uint64 {
	for i := preambleSize; i < len(nums); i++ {
		var ok bool
		for j := i - preambleSize; j < i; j++ {
			for k := j + 1; k < i; k++ {
				if nums[k]+nums[j] == nums[i] {
					ok = true
					break
				}
			}
			if ok {
				break
			}
		}
		if !ok {
			return nums[i]
		}
	}
	return 0
}

// weakness finds the first contiguous set of numbers which sum
// to the target number and returns the sum of the smallest and
// largest of those
func weakness(nums []uint64, target uint64) uint64 {
	for i := 0; i < len(nums); i++ {
		var sum uint64
		for j := i; j < len(nums); j++ {
			sum += nums[j]
			if sum == target {
				found := nums[i : j+1]
				sort.Slice(found, func(i, j int) bool { return found[i] < found[j] })
				return found[0] + found[len(found)-1]
			}
			if sum > target {
				break
			}
		}
	}
	return 0
}
