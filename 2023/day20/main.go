package main

import (
	"bytes"
	"log"
	"os"
	"sort"
)

func main() {
	data, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatalln("Failed to read input:", err)
	}

	log.Println("Part one:", parse(data).partOne())
	log.Println("Part two:", parse(data).partTwo())
}

func parse(input []byte) modules {
	mods := make(modules)
	for _, line := range bytes.Split(input, []byte{'\n'}) {
		parts := bytes.Split(line, []byte(" -> "))
		mod := module{state: make(map[string]bool), name: string(parts[0])}
		switch mod.name[0] {
		case '%', '&':
			mod.kind, mod.name = mod.name[0], mod.name[1:]
		}
		for _, out := range bytes.Split(parts[1], []byte{',', ' '}) {
			mod.outputs = append(mod.outputs, string(out))
		}
		mods[mod.name] = &mod
	}
	mods["button"] = &module{outputs: []string{"broadcaster"}, name: "button"}
	for name, mod := range mods {
		for _, out := range mod.outputs {
			if _, ok := mods[out]; ok {
				mods[out].state[name] = false
			}
		}
	}
	return mods
}

type modules map[string]*module

func (m modules) partOne() int {
	counts := make(map[bool]int)
	for i := 0; i < 1000; i++ {
		counts = m.tick(i, counts, make(map[string]int))
	}
	return counts[true] * counts[false]
}

func (m modules) partTwo() int {
	repeats := make(map[string]int)
	m.backtraceSoleConjunctions("rx", repeats)
outer:
	for i := 1; ; i++ {
		m.tick(i, make(map[bool]int), repeats)
		var cycles []int
		for _, cycle := range repeats {
			if cycle == 0 {
				continue outer
			}
			cycles = append(cycles, cycle)
		}
		return lowestCommonMultiple(cycles...)
	}
}

func (m modules) backtraceSoleConjunctions(start string, conjunctions map[string]int) {
	var found bool
	for n, mod := range m {
		if sort.Search(len(mod.outputs), func(i int) bool {
			return mod.outputs[i] == start
		}) < len(mod.outputs) && len(mod.outputs) == 1 && mod.kind == '&' {
			m.backtraceSoleConjunctions(n, conjunctions)
			found = true
		}
	}
	if !found {
		conjunctions[start] = 0
	}
}

func (m modules) tick(i int, counts map[bool]int, repeats map[string]int) map[bool]int {
	p, _ := m["button"].Receive(pulse{})
	pulses := []pulse{p}
	for ; len(pulses) > 0; pulses = pulses[1:] {
		p := pulses[0]
		counts[p.value] += len(p.to)
		for _, t := range p.to {
			if _, ok := m[t]; !ok {
				continue
			}
			if emit, ok := m[t].Receive(p); ok {
				pulses = append(pulses, emit)
			}
			for name := range repeats {
				if name == t && !p.value {
					repeats[name] = i
				}
			}
		}
	}
	return counts
}

func lowestCommonMultiple(is ...int) int {
	factors := make(map[int]int)
	for _, i := range is {
		for k, n := range primeFactors(i) {
			if factors[k] < n {
				factors[k] = n
			}
		}
	}

	lcm := 1
	for k, n := range factors {
		for i := 0; i < n; i++ {
			lcm *= k
		}
	}

	return lcm
}

func primeFactors(n int) map[int]int {
	pf := make(map[int]int)
	for i := n % 2; i == 0; i = n % 2 {
		pf[2]++
		n /= 2
	}
	for i := 3; i*i <= n; i += 2 {
		for j := n % i; j == 0; j = n % i {
			pf[i]++
			n /= i
		}
		if n == 1 {
			return pf
		}
	}
	if n > 2 {
		pf[n]++
	}
	return pf
}

type pulse struct {
	from  string
	to    []string
	value bool
}

type module struct {
	name    string
	kind    byte
	outputs []string
	state   map[string]bool
	next    bool
}

func (m *module) Receive(p pulse) (pulse, bool) {
	switch m.kind {
	case '%':
		if p.value {
			return p, false
		}
		m.state[""] = !m.state[""]
		m.next = m.state[""]
	case '&':
		m.state[p.from] = p.value
		m.next = true
		for _, s := range m.state {
			m.next = m.next && s
		}
		m.next = !m.next
	default:
	}
	return pulse{from: m.name, to: m.outputs, value: m.next}, true
}
