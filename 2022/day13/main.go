package main

import (
	"bytes"
	"log"
	"os"
	"sort"
	"strconv"
)

func main() {
	data, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatalln("Failed to read input:", err)
	}

	log.Println("Part one:", partOne(parsePairs(data)))
	log.Println("Part two:", partTwo(parsePackets(data)))
}

func parsePackets(data []byte) packets {
	var p packets
	for _, line := range bytes.Split(data, []byte{'\n'}) {
		if len(line) == 0 {
			continue
		}
		_, pkt := parse(line)
		p = append(p, pkt)
	}
	return p
}

func parsePairs(data []byte) []pair {
	var pairs []pair
	for _, p := range bytes.Split(data, []byte{'\n', '\n'}) {
		pairs = append(pairs, parsePair(p))
	}
	return pairs
}

func partOne(pairs []pair) int {
	var sum int
	for i, p := range pairs {
		if p.Valid() {
			sum += i + 1
		}
	}
	return sum
}

func partTwo(pkts packets) int {
	divX, divY := packet{values: []interface{}{packet{values: []interface{}{2}}}}, packet{values: []interface{}{packet{values: []interface{}{6}}}}
	pkts = append(pkts, divX, divY)
	sort.Slice(pkts, pkts.Less)
	var x, y int
	for i, p := range pkts {
		if p.Equal(divX) {
			x = i + 1
		}
		if p.Equal(divY) {
			y = i + 1
		}
	}
	return x * y
}

type pair [2]packet

func parsePair(data []byte) pair {
	a, b, _ := bytes.Cut(data, []byte{'\n'})
	_, first := parse(a)
	_, second := parse(b)
	return pair{first, second}
}

type packets []packet

func (p packets) Less(i, j int) bool {
	return pair{p[i], p[j]}.Valid()
}

func (p pair) Valid() bool {
	v, _ := p.innerValid()
	return v
}

func (p pair) innerValid() (bool, bool) {
	if len(p[0].values) == 0 && len(p[1].values) != 0 {
		return true, true
	}
	for i, a := range p[0].values {
		if i == len(p[1].values) {
			return false, true
		}
		b := p[1].values[i]
		switch t := a.(type) {
		case int:
			switch s := b.(type) {
			case int:
				if t > s {
					return false, true
				}
				if t < s {
					return true, true
				}
			case packet:
				valid, done := (pair{packet{values: []interface{}{p[0].values[i]}}, s}).innerValid()
				if !valid || done {
					return valid, done
				}
			}
		case packet:
			switch s := b.(type) {
			case int:
				valid, done := (pair{t, packet{values: []interface{}{p[1].values[i]}}}).innerValid()
				if !valid || done {
					return valid, done
				}
			case packet:
				valid, done := (pair{t, s}).innerValid()
				if !valid || done {
					return valid, done
				}
			}
		}
	}
	if len(p[1].values) > len(p[0].values) {
		return true, true
	}
	return true, false
}

type packet struct {
	values []interface{}
}

func (p packet) Equal(q packet) bool {
	if len(p.values) != len(q.values) {
		return false
	}
	for i, v := range p.values {
		switch x := v.(type) {
		case int:
			switch y := q.values[i].(type) {
			case int:
				if x != y {
					return false
				}
			default:
				return false
			}
		case packet:
			switch y := q.values[i].(type) {
			case packet:
				return x.Equal(y)
			default:
				return false
			}
		}
	}
	return true
}

func parse(data []byte) (int, packet) {
	p := packet{values: make([]interface{}, 0)}
	for i := 1; i < len(data); i++ {
		switch data[i] {
		case '[':
			j, sub := parse(data[i:])
			i += j
			p.values = append(p.values, sub)
		case ']':
			return i, p
		case ',':
			continue
		default:
			var done bool
			var j int
			for j = i; !done; j++ {
				switch data[j] {
				case ',', ']':
					done = true
				}
			}
			v, _ := strconv.Atoi(string(data[i : j-1]))
			p.values = append(p.values, v)
			i = j - 2
		}
	}
	return len(data) - 1, p
}
