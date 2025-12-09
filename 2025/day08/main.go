package main

import (
	_ "embed"
	"fmt"
	"math"
	"sort"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

func main() {
	var points []xyz
	var circuits circuits
	for line := range strings.SplitSeq(input, "\n") {
		var p xyz
		for i, v := range strings.Split(line, ",") {
			p[i], _ = strconv.Atoi(v)
		}
		points = append(points, p)
		circuits = append(circuits, map[xyz]struct{}{p: {}})
	}
	for i, connection := range connections(points) {
		if i == 1000 {
			partOne := 1
			for i := range 3 {
				partOne *= len(circuits[i])
			}
			fmt.Println("Part one:", partOne)
		}
		circuits = circuits.add(connection)
		sort.Slice(circuits, func(i, j int) bool {
			return len(circuits[j]) < len(circuits[i])
		})
		if len(circuits) == 1 {
			fmt.Println("Part two:", connection.a[0]*connection.b[0])
			break
		}
	}
}

type xyz [3]int

func (a xyz) distance(b xyz) float64 {
	return math.Sqrt(
		math.Pow(float64(a[0]-b[0]), 2) +
			math.Pow(float64(a[1]-b[1]), 2) +
			math.Pow(float64(a[2]-b[2]), 2),
	)
}

type connection struct {
	a, b   xyz
	length float64
}

func connections(points []xyz) []connection {
	var connections []connection
	for i := range len(points) {
		for j := i + 1; j < len(points); j++ {
			connections = append(connections, connection{points[i], points[j], points[i].distance(points[j])})
		}
	}
	sort.Slice(connections, func(i, j int) bool {
		return connections[i].length < connections[j].length
	})
	return connections
}

type circuits []map[xyz]struct{}

func (circuits circuits) add(connection connection) circuits {
	ia, ib := -1, -1
	for i, circuit := range circuits {
		if _, ok := circuit[connection.a]; ok {
			ia = i
		}
		if _, ok := circuit[connection.b]; ok {
			ib = i
		}
		if ia != -1 && ib != -1 {
			break
		}
	}
	if ia != ib {
		for connection := range circuits[ib] {
			circuits[ia][connection] = struct{}{}
		}
		circuits = append(circuits[:ib], circuits[ib+1:]...)
	}
	return circuits
}
