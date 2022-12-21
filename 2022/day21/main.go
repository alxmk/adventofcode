package main

import (
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	data, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatalln("Failed to read input:", err)
	}

	monkeys := parse(string(data))
	log.Println("Part one:", monkeys["root"].eval(monkeys, &suspiciousCache{data: make(map[string]*cacheEntry)}))

	log.Println("Part two:", partTwo(parse(string(data))))
}

func partTwo(monkeys map[string]*monkey) int {
	monkeys["root"].op = func(i, j int) int {
		if i == j {
			return 1
		}
		return 0
	}

	// Populate the smart cache
	cache := &suspiciousCache{data: make(map[string]*cacheEntry)}
	for i := 0; i < 5; i++ {
		monkeys["humn"].value = &i
		if monkeys["root"].eval(monkeys, cache) == 1 {
			// Hey we might get lucky right
			return i
		}
	}
	// Find out which side of the equation changes
	if cache.data[monkeys["root"].a].static {
		_, target := cache.Get(monkeys["root"].a)
		return binarySearch(monkeys["root"].b, monkeys, target, cache)
	}
	_, target := cache.Get(monkeys["root"].b)
	return binarySearch(monkeys["root"].a, monkeys, target, cache)
}

func binarySearch(name string, monkeys map[string]*monkey, target int, cache *suspiciousCache) int {
	// Figure out if we're searching up or down
	this := monkeys[name].eval(monkeys, cache)
	op := gt
	if this < target {
		op = lt
	}
	return innersearch(name, monkeys, target, cache, 4, 100000000, op)
}

func innersearch(name string, monkeys map[string]*monkey, target int, cache *suspiciousCache, lower, increment int, op func(int, int) bool) int {
	for i := lower; ; i += increment {
		monkeys["humn"].value = &i
		this := monkeys[name].eval(monkeys, cache)
		if op(this, target) {
			lower = i
			continue
		}
		if this == target {
			if increment == 1 {
				return i
			}
		}
		break
	}
	// Can't reduce increment lower than 1 so something went wrong
	if increment == 1 {
		panic("failed")
	}
	return innersearch(name, monkeys, target, cache, lower, increment/10, op)
}

func lt(i, j int) bool {
	return i < j
}

func gt(i, j int) bool {
	return i > j
}

func parse(input string) map[string]*monkey {
	monkeys := make(map[string]*monkey)
	for _, line := range strings.Split(input, "\n") {
		fields := strings.Fields(line)
		name := fields[0][:len(fields[0])-1]
		switch len(fields) {
		case 2:
			// Yell monkey
			v, _ := strconv.Atoi(fields[1])
			monkeys[name] = &monkey{name: name, value: &v}
		case 4:
			// Operator monkey
			monkeys[name] = &monkey{name: name, a: fields[1], b: fields[3], op: getOp(fields[2])}
		}
	}
	return monkeys
}

type monkey struct {
	name  string
	value *int
	op    operator
	a, b  string
}

type operator func(int, int) int

var (
	add   = func(i, j int) int { return i + j }
	minus = func(i, j int) int { return i - j }
	mul   = func(i, j int) int { return i * j }
	div   = func(i, j int) int { return i / j }
)

func getOp(i string) operator {
	switch i {
	case "+":
		return add
	case "-":
		return minus
	case "*":
		return mul
	case "/":
		return div
	}
	panic(i)
}

func (m *monkey) eval(monkeys map[string]*monkey, cache *suspiciousCache) int {
	if m.value != nil {
		return *m.value
	}
	if ok, v := cache.Get(m.name); ok {
		return v
	}
	v := m.op(monkeys[m.a].eval(monkeys, cache), monkeys[m.b].eval(monkeys, cache))
	cache.Set(m.name, v)
	return v
}

// suspiciousCache thinks everything is sus so waits for 5 identical results before adding to cache
type suspiciousCache struct {
	data map[string]*cacheEntry
}

func (s *suspiciousCache) Get(k string) (bool, int) {
	if v, ok := s.data[k]; ok && v.static {
		return true, v.results[0]
	}
	return false, -1
}

func (s *suspiciousCache) Set(k string, v int) {
	if _, ok := s.data[k]; !ok {
		s.data[k] = &cacheEntry{}
	}
	if len(s.data[k].results) < 5 {
		s.data[k].results = append(s.data[k].results, v)
	}
	if len(s.data[k].results) == 5 {
		f := s.data[k].results[0]
		for _, r := range s.data[k].results {
			if r != f {
				return
			}
		}
		s.data[k].static = true
	}
}

func (s *suspiciousCache) Len(static bool) int {
	var len int
	for _, v := range s.data {
		if v.static == static {
			len++
		}
	}
	return len
}

type cacheEntry struct {
	results []int
	static  bool
}
