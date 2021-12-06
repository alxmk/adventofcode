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

	lines := strings.Split(string(data), "\n")

	gamma, epsilon, err := partOne(lines)
	if err != nil {
		log.Fatalln("Failed part one:", err)
	}
	log.Println("Part one:", gamma*epsilon)

	o2, co2, err := partTwo(lines)
	if err != nil {
		log.Fatalln("Failed part two:", err)
	}
	log.Println("Part two:", o2*co2)
}

func partOne(lines []string) (uint, uint, error) {
	counts := make([]uint, len(lines[0]))
	for _, line := range lines {
		v, err := strconv.ParseUint(line, 2, 64)
		if err != nil {
			return 0, 0, fmt.Errorf("malformed line %s: %s", line, err)
		}
		for i := 0; i < len(line); i++ {
			if v&(1<<i) != 0 {
				counts[i]++
			}
		}
	}

	var gamma, epsilon uint
	for i, v := range counts {
		if v > uint(len(lines)/2) {
			gamma += 1 << i
		} else {
			epsilon += 1 << i
		}
	}

	return gamma, epsilon, nil
}

func partTwo(lines []string) (uint64, uint64, error) {
	nummap := make(map[uint64]struct{})

	for _, line := range lines {
		v, err := strconv.ParseUint(line, 2, 64)
		if err != nil {
			return 0, 0, fmt.Errorf("malformed line %s: %s", line, err)
		}
		nummap[v] = struct{}{}
	}

	o2, err := bitCriteria(lines, nummap, func(a map[uint64]struct{}, b map[uint64]struct{}) bool { return len(a) > len(b) })
	if err != nil {
		return 0, 0, fmt.Errorf("failed to calculate o2 rating: %s", err)
	}

	co2, err := bitCriteria(lines, nummap, func(a map[uint64]struct{}, b map[uint64]struct{}) bool { return len(a) <= len(b) })
	if err != nil {
		return 0, 0, fmt.Errorf("failed to calculate co2 rating: %s", err)
	}

	return o2, co2, nil
}

func bitCriteria(lines []string, nummap map[uint64]struct{}, compare func(a, b map[uint64]struct{}) bool) (uint64, error) {
	copyMap := make(map[uint64]struct{})
	for k := range nummap {
		copyMap[k] = struct{}{}
	}

	for i := len(lines[0]) - 1; i >= 0 && len(copyMap) > 1; i-- {
		map0, map1 := make(map[uint64]struct{}), make(map[uint64]struct{})
		for n := range copyMap {
			if n&(1<<i) == 0 {
				map0[n] = struct{}{}
				continue
			}
			map1[n] = struct{}{}
		}
		if compare(map0, map1) {
			copyMap = map0
			continue
		}
		copyMap = map1
	}

	if len(copyMap) != 1 {
		return 0, fmt.Errorf("ended up with more than one value")
	}

	for k := range copyMap {
		return k, nil
	}

	// Can't get here
	return 0, fmt.Errorf("whaaaa")
}
