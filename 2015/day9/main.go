package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"strings"
)

func main() {
	data, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatalln("Failed to read input file", err)
	}

	r := make(routes)
	citiesMap := make(map[string]struct{})
	for _, line := range strings.Split(string(data), "\n") {
		var from, to string
		var distance int
		fmt.Sscanf(line, "%s to %s = %d", &from, &to, &distance)
		if _, ok := r[from]; !ok {
			r[from] = make(map[string]int)
		}
		r[from][to] = distance
		if _, ok := r[to]; !ok {
			r[to] = make(map[string]int)
		}
		r[to][from] = distance
		citiesMap[from], citiesMap[to] = struct{}{}, struct{}{}
	}

	var cities []string
	for k := range citiesMap {
		cities = append(cities, k)
	}

	minDistance, maxDistance := math.MaxInt16, 0
	for _, permutation := range permutations(cities) {
		d := r.Try(permutation)
		if d < minDistance {
			minDistance = d
		}
		if d > maxDistance {
			maxDistance = d
		}
	}
	log.Println("Part one:", minDistance)
	log.Println("Part two:", maxDistance)
}

type routes map[string]map[string]int

func (r routes) Try(order []string) int {
	var distance int
	var previous string
	for _, city := range order {
		if previous != "" {
			distance += r[previous][city]
		}
		previous = city
	}
	return distance
}

func permutations(cities []string) [][]string {
	var helper func([]string, int)
	var result [][]string

	helper = func(arr []string, n int) {
		if n == 1 {
			tmp := make([]string, len(arr))
			copy(tmp, arr)
			result = append(result, tmp)
		} else {
			for i := 0; i < n; i++ {
				helper(arr, n-1)
				if n%2 == 1 {
					tmp := arr[i]
					arr[i] = arr[n-1]
					arr[n-1] = tmp
				} else {
					tmp := arr[0]
					arr[0] = arr[n-1]
					arr[n-1] = tmp
				}
			}
		}
	}

	helper(cities, len(cities))
	return result
}
