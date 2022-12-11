package main

import (
	"container/list"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	data, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatalln("Failed to read input:", err)
	}

	m, product := parse(string(data))
	log.Println("Part one:", simulate(m, func(i uint64) uint64 { return i / 3 }, 20))
	m, product = parse(string(data))
	log.Println("Part two:", simulate(m, func(i uint64) uint64 { return i % product }, 10000))
}

type Monkey struct {
	items *list.List
	op    func(uint64) uint64
	test  func(uint64) uint64
}

func parse(input string) ([]*Monkey, uint64) {
	var monkeys []*Monkey
	product := uint64(1)
	for _, block := range strings.Split(input, "\n\n") {
		m, prime := parseBlock(block)
		product *= prime
		monkeys = append(monkeys, m)
	}
	return monkeys, product
}

func parseBlock(input string) (*Monkey, uint64) {
	var p uint64
	var m Monkey
	lines := strings.Split(input, "\n")
	for n, line := range lines {
		fields := strings.Fields(line)
		switch fields[0] {
		case "Starting":
			m.items = list.New()
			for _, f := range fields[2:] {
				v, _ := strconv.ParseUint(strings.TrimSuffix(f, ","), 10, 64)
				m.items.PushBack(v)
			}
		case "Operation:":
			switch fields[4] {
			case "*":
				switch fields[5] {
				case "old":
					m.op = func(i uint64) uint64 { return i * i }
				default:
					v, _ := strconv.ParseUint(fields[5], 10, 64)
					m.op = func(i uint64) uint64 { return i * v }
				}
			case "+":
				v, _ := strconv.ParseUint(fields[5], 10, 64)
				m.op = func(i uint64) uint64 { return i + v }
			}
		case "Test:":
			p, _ = strconv.ParseUint(fields[3], 10, 64)
			a, _ := strconv.ParseUint(strings.Fields(lines[n+1])[5], 10, 64)
			b, _ := strconv.ParseUint(strings.Fields(lines[n+2])[5], 10, 64)
			m.test = func(i uint64) uint64 {
				if i%p == 0 {
					return a
				}
				return b
			}
		}
	}
	return &m, p
}

func simulate(monkeys []*Monkey, relief func(uint64) uint64, rounds uint64) uint64 {
	inspectionCount := make([]uint64, len(monkeys))
	for round := uint64(0); round < rounds; round++ {
		for j, m := range monkeys {
			for i := m.items.Front(); i != nil; i = m.items.Front() {
				m.items.Remove(i)
				v := i.Value.(uint64)
				inspectionCount[j]++
				v = relief(m.op(v))
				monkeys[m.test(v)].items.PushBack(v)
			}
		}
	}
	sort.Slice(inspectionCount, func(i, j int) bool { return inspectionCount[i] < inspectionCount[j] })
	return inspectionCount[len(inspectionCount)-2] * inspectionCount[len(inspectionCount)-1]
}
