package main

import (
	"bytes"
	"encoding/json"
	"log"
	"os"
	"sort"
)

func main() {
	data, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatalln("Failed to read input:", err)
	}

	packets := parsePackets(data)

	log.Println("Part one:", partOne(packets))
	log.Println("Part two:", partTwo(packets))
}

func parsePackets(data []byte) packets {
	var p packets
	for _, line := range bytes.Split(data, []byte{'\n'}) {
		if len(line) == 0 {
			continue
		}
		var this []interface{}
		json.Unmarshal(line, &this)
		p = append(p, this)
	}
	return p
}

func partOne(pkts packets) int {
	var sum int
	for i := 0; i < len(pkts); i += 2 {
		if pkts.Less(i, i+1) {
			sum += (i / 2) + 1
		}
	}
	return sum
}

func partTwo(pkts packets) int {
	da, db := []interface{}{[]interface{}{2.0}}, []interface{}{[]interface{}{6.0}}
	pkts = append(pkts, da, db)
	sort.Slice(pkts, pkts.Less)
	var x, y int
	for i, p := range pkts {
		if equal(p, da) {
			x = i + 1
		}
		if equal(p, db) {
			y = i + 1
		}
	}
	return x * y
}

type packets [][]interface{}

func (p packets) Less(i, j int) bool {
	less, _ := p.innerLess(i, j)
	return less
}

func (p packets) innerLess(i, j int) (bool, bool) {
	for k, a := range p[i] {
		if k == len(p[j]) {
			return false, true
		}
		b := p[j][k]
		var n, m []interface{}
		switch t := a.(type) {
		case float64:
			switch s := b.(type) {
			case float64:
				if t == s {
					continue
				}
				return t < s, true
			case []interface{}:
				n, m = []interface{}{p[i][k]}, s
			}
		case []interface{}:
			switch s := b.(type) {
			case float64:
				n, m = t, []interface{}{p[j][k]}
			case []interface{}:
				n, m = t, s
			}
		}
		if less, done := (packets{n, m}).innerLess(0, 1); done {
			return less, done
		}
	}
	if len(p[j]) > len(p[i]) {
		return true, true
	}
	return true, false
}

func equal(p, q []interface{}) bool {
	pp, _ := json.Marshal(p)
	qq, _ := json.Marshal(q)
	return string(qq) == string(pp)
}
