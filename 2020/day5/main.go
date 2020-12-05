package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"sort"
	"strings"
)

func main() {
	data, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatalln("Failed to load input:", err)
	}

	var maxID int
	existingIDs := make(map[int]struct{})
	p := make(plane)
	for _, line := range strings.Split(string(data), "\n") {
		s := parseSeat(line)
		id := s.ID()
		existingIDs[id] = struct{}{}
		if id > maxID {
			maxID = s.ID()
		}
		if _, ok := p[s.column]; !ok {
			p[s.column] = make(map[int]struct{})
		}
		p[s.column][s.row] = struct{}{}
	}

	log.Println("Part one:", maxID)

	for _, open := range p.OpenSeatIDs() {
		_, okprev := existingIDs[open-1]
		_, oknext := existingIDs[open+1]
		if oknext && okprev {
			log.Println("Part two:", open)
		}
	}
}

type plane map[int]map[int]struct{}

func (p plane) String() string {
	var b strings.Builder
	for y := 0; y < 128; y++ {
		b.WriteString(fmt.Sprintf("%03d", y))
		for x := 0; x < 8; x++ {
			if col, ok := p[x]; ok {
				if _, ok := col[y]; ok {
					b.WriteRune('X')
					continue
				}
			}
			b.WriteRune('O')
		}
		b.WriteRune('\n')
	}
	return b.String()
}

func (p plane) OpenSeatIDs() []int {
	var open []int
	for y := 0; y < 128; y++ {
		for x := 0; x < 8; x++ {
			if col, ok := p[x]; ok {
				if _, ok := col[y]; ok {
					continue
				}
			}
			open = append(open, seat{y, x}.ID())
		}
	}
	sort.Sort(sort.IntSlice(open))
	return open
}

type seat struct {
	row, column int
}

func (s seat) ID() int {
	return (s.row * 8) + s.column
}

func parseSeat(input string) seat {
	rowMin, rowMax := 0, 127
	for i, c := range input {
		if i == 7 {
			break
		}
		switch c {
		case 'F':
			rowMax = rowMax - ((rowMax - rowMin + 1) / 2)
		case 'B':
			rowMin = rowMin + ((rowMax - rowMin + 1) / 2)
		}
	}
	columnMin, columnMax := 0, 7
	for i := 7; i < len(input); i++ {
		switch input[i] {
		case 'L':
			columnMax = columnMax - ((columnMax - columnMin + 1) / 2)
		case 'R':
			columnMin = columnMin + ((columnMax - columnMin + 1) / 2)
		}
	}
	return seat{row: rowMin, column: columnMin}
}
