package main

import (
	"bytes"
	"log"
	"os"
	"strconv"
)

func main() {
	data, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatalln("Failed to read input:", err)
	}

	p1, p2 := solve(data)

	log.Println("Part one:", p1)
	log.Println("Part two:", p2)
}

func solve(input []byte) (int64, int64) {
	s1, s2 := state{v: make(map[xy]xy)}, state{v: make(map[xy]xy)}
	for _, line := range bytes.Split(input, []byte{'\n'}) {
		m1, _ := strconv.ParseInt(string(bytes.Fields(line)[1]), 10, 64)
		s1, s2 = s1.Increment(m1, parseDir(line[0])), s2.Increment(parseHex(bytes.Fields(line)[2]))
	}
	return s1.Area(), s2.Area()
}

type state struct {
	a xy
	p int64
	v map[xy]xy
}

func (s state) Increment(m int64, d xy) state {
	last := s.a
	s.a.x, s.a.y = s.a.x+(m*d.x), s.a.y+(m*d.y)
	s.p += max(s.a.x-last.x, last.x-s.a.x) + max(s.a.y-last.y, last.y-s.a.y)
	s.v[last] = s.a
	return s
}

func (s state) Area() int64 {
	var sum int64
	for a, b := range s.v {
		sum += a.x*b.y - a.y*b.x
	}
	return (sum / 2) + (s.p / 2) + 1
}

type xy struct {
	x, y int64
}

func parseDir(i byte) xy {
	switch i {
	case 'R', '0':
		return xy{1, 0}
	case 'L', '2':
		return xy{-1, 0}
	case 'U', '3':
		return xy{0, -1}
	case 'D', '1':
		return xy{0, 1}
	}
	panic("Invalid dir")
}

func parseHex(input []byte) (int64, xy) {
	v, _ := strconv.ParseInt(string(input[2:7]), 16, 64)
	return v, parseDir(input[7])
}
