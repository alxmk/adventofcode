package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

func main() {
	data, err := ioutil.ReadFile("part1.txt")
	if err != nil {
		log.Fatalln("Failed to read input", err)
	}

	// Splitting on double new line should give us the blocks
	blocks := strings.Split(string(data), "\n\n")

	possibles := make(map[int]map[opcode]struct{})

	// Before: [2, 0, 0, 3]
	// 3 0 3 1
	// After:  [2, 0, 0, 3]
	var num int
	for _, b := range blocks {
		var b0, b1, b2, b3, i0, i1, i2, i3, a0, a1, a2, a3 int
		fmt.Sscanf(b, `Before: [%d, %d, %d, %d]
%d %d %d %d
After:  [%d, %d, %d, %d`, &b0, &b1, &b2, &b3, &i0, &i1, &i2, &i3, &a0, &a1, &a2, &a3)

		matches := analyseInstructions([]int{b0, b1, b2, b3}, []int{i0, i1, i2, i3}, []int{a0, a1, a2, a3})

		if len(matches) >= 3 {
			num++
		}

		// Track the possible matches for each opcode
		if _, ok := possibles[i0]; !ok {
			possibles[i0] = make(map[opcode]struct{})
		}

		for _, m := range matches {
			possibles[i0][m] = struct{}{}
		}
	}

	dict := reduce(possibles)

	log.Println("Part one:", num)

	log.Printf("Dict: %#v", dict)

	data2, err := ioutil.ReadFile("part2.txt")
	if err != nil {
		log.Fatalln("Failed to read input", err)
	}

	register := make([]int, 4)
	for _, line := range strings.Split(string(data2), "\n") {
		var i0, i1, i2, i3 int
		fmt.Sscanf(line, "%d %d %d %d", &i0, &i1, &i2, &i3)
		op, ok := dict[i0]
		if !ok {
			log.Fatalln("Failed to find opcode for", i0, i1, i2, i3, line)
		}

		register = op.Do([]int{i0, i1, i2, i3}, register)
	}

	log.Println("Part two:", register[0])
}

func reduce(possibles map[int]map[opcode]struct{}) map[int]opcode {
	reduced := make(map[int]opcode)

	size := len(possibles)

	i := 0
	for len(reduced) != size && i < 1000 {
		for realcode, opmap := range possibles {
			if len(opmap) == 1 {
				for op := range opmap {
					reduced[realcode] = op
				}
			}
		}

		var todelete []int
		for rc, m := range reduced {
			for realcode, opmap := range possibles {
				if realcode == rc {
					todelete = append(todelete, realcode)
				}
				delete(opmap, m)
			}
		}

		for _, d := range todelete {
			delete(possibles, d)
		}
		i++
	}

	return reduced
}

func analyseInstructions(before, instruction, after []int) []opcode {
	var matching []opcode
	for _, o := range opcodes {
		if result := o.Do(instruction, before); matchRegisters(after, result) {
			matching = append(matching, o)
		}
	}
	return matching
}

func matchRegisters(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}

	for i, e := range a {
		if b[i] != e {
			return false
		}
	}

	return true
}

type opcode int

const (
	addr opcode = iota
	addi

	mulr
	muli

	banr
	bani

	borr
	bori

	setr
	seti

	gtir
	gtri
	gtrr

	eqir
	eqri
	eqrr
)

var opcodes = []opcode{
	addr, addi, mulr, muli, banr, bani, borr, bori, setr, seti, gtir, gtri, gtrr, eqir, eqri, eqrr,
}

func (o opcode) Do(instruction, before []int) []int {
	a := instruction[1]
	b := instruction[2]
	c := instruction[3]

	register := make([]int, len(before))
	copy(register, before)

	switch o {
	case addr:
		return faddr(a, b, c, register)
	case addi:
		return faddi(a, b, c, register)
	case mulr:
		return fmulr(a, b, c, register)
	case muli:
		return fmuli(a, b, c, register)
	case banr:
		return fbanr(a, b, c, register)
	case bani:
		return fbani(a, b, c, register)
	case borr:
		return fborr(a, b, c, register)
	case bori:
		return fbori(a, b, c, register)
	case setr:
		return fsetr(a, b, c, register)
	case seti:
		return fseti(a, b, c, register)
	case gtir:
		return fgtir(a, b, c, register)
	case gtri:
		return fgtri(a, b, c, register)
	case gtrr:
		return fgtrr(a, b, c, register)
	case eqir:
		return feqir(a, b, c, register)
	case eqri:
		return feqri(a, b, c, register)
	case eqrr:
		return feqrr(a, b, c, register)
	}

	panic("Undefined opcode")
}

func insertAt(value, pos int, slice []int) []int {
	output := append(slice[:pos], value)

	return append(output, slice[pos+1:]...)
}

func faddr(a, b, c int, before []int) []int {
	return insertAt(before[a]+before[b], c, before)
}
func faddi(a, b, c int, before []int) []int {
	return insertAt(before[a]+b, c, before)
}
func fmulr(a, b, c int, before []int) []int {
	return insertAt(before[a]*before[b], c, before)
}
func fmuli(a, b, c int, before []int) []int {
	return insertAt(before[a]*b, c, before)
}
func fbanr(a, b, c int, before []int) []int {
	return insertAt(before[a]&before[b], c, before)
}
func fbani(a, b, c int, before []int) []int {
	return insertAt(before[a]&b, c, before)
}
func fborr(a, b, c int, before []int) []int {
	return insertAt(before[a]|before[b], c, before)
}
func fbori(a, b, c int, before []int) []int {
	return insertAt(before[a]|b, c, before)
}
func fsetr(a, b, c int, before []int) []int {
	return insertAt(before[a], c, before)
}
func fseti(a, b, c int, before []int) []int {
	return insertAt(a, c, before)
}
func fgtir(a, b, c int, before []int) []int {
	if a > before[b] {
		return insertAt(1, c, before)
	}
	return insertAt(0, c, before)
}
func fgtri(a, b, c int, before []int) []int {
	if before[a] > b {
		return insertAt(1, c, before)
	}
	return insertAt(0, c, before)
}
func fgtrr(a, b, c int, before []int) []int {
	if before[a] > before[b] {
		return insertAt(1, c, before)
	}
	return insertAt(0, c, before)
}
func feqir(a, b, c int, before []int) []int {
	if a == before[b] {
		return insertAt(1, c, before)
	}
	return insertAt(0, c, before)
}
func feqri(a, b, c int, before []int) []int {
	if before[a] == b {
		return insertAt(1, c, before)
	}
	return insertAt(0, c, before)
}
func feqrr(a, b, c int, before []int) []int {
	if before[a] == before[b] {
		return insertAt(1, c, before)
	}
	return insertAt(0, c, before)
}
