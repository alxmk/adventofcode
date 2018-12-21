package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func main() {
	data, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatalln("Failed to read input", err)
	}

	var instructionNum, ipreg int

	programme := make(map[int]cmd)
	for _, line := range strings.Split(string(data), "\n") {
		words := strings.Fields(line)

		switch words[0] {
		case "#ip":
			ipreg, _ = strconv.Atoi(words[1])
		default:
			var a, b, c int
			var code string

			fmt.Sscanf(line, "%s %d %d %d", &code, &a, &b, &c)

			programme[instructionNum] = cmd{op: opcode(code), a: a, b: b, c: c}

			instructionNum++
		}
	}

	reg := make(registers, 6)

	var part1, part2 int

	var previous int
	seen := make(map[int]struct{})

	for {
		ip := reg[ipreg]
		if cmd, ok := programme[ip]; ok {
			// fmt.Println(ip, cmd.String())
			after := cmd.Do(reg)
			if ip == 28 {
				// fmt.Printf("ip=%d %s %s %s\n", ip, reg.String(), cmd.String(), after.String())
				// fmt.Println(reg[1])
				if part1 == 0 {
					part1 = reg[1]
					// break
				}
				if _, ok := seen[reg[1]]; ok {
					part2 = previous
					break
				}
				previous = reg[1]
				seen[reg[1]] = struct{}{}
			}
			reg = after
		} else {
			break
		}

		reg[ipreg]++
	}

	fmt.Println("Part one:", part1)
	fmt.Println("Part two:", part2)
}

type registers []int

func (r registers) String() string {
	var values []string

	for _, reg := range r {
		values = append(values, fmt.Sprintf("%d", reg))
	}

	return fmt.Sprintf("[%s]", strings.Join(values, ", "))
}

type cmd struct {
	op      opcode
	a, b, c int
}

func (c cmd) Do(regs registers) registers {
	return c.op.Do(c.a, c.b, c.c, regs)
}

func (c cmd) String() string {
	return fmt.Sprintf("%s %d %d %d", c.op, c.a, c.b, c.c)
}

type opcode string

const (
	addr opcode = "addr"
	addi opcode = "addi"

	mulr opcode = "mulr"
	muli opcode = "muli"

	banr opcode = "banr"
	bani opcode = "bani"

	borr opcode = "borr"
	bori opcode = "bori"

	setr opcode = "setr"
	seti opcode = "seti"

	gtir opcode = "gtir"
	gtri opcode = "gtri"
	gtrr opcode = "gtrr"

	eqir opcode = "eqir"
	eqri opcode = "eqri"
	eqrr opcode = "eqrr"
)

func (o opcode) Do(a, b, c int, before []int) []int {
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
