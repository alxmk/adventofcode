package intcode

import (
	"fmt"
	"strconv"
	"strings"
)

// Parse input string as an intcode Programme
func Parse(input string) (Programme, error) {
	p := make(Programme)
	for i, value := range strings.Split(input, ",") {
		parsed, err := strconv.Atoi(value)
		if err != nil {
			return nil, fmt.Errorf("failed to parse %s as int: %v", value, err)
		}
		p[i] = parsed
	}
	return p, nil
}

// Run a Programme
func (p Programme) Run(input, output chan int) error {
	var pointer int
	for {
		op, modeA, modeB, _ := parseOpcode(p[pointer])
		switch op {
		case 99: // Exit
			return nil
		case 1: // Add
			p[p[pointer+3]] = value(p, pointer+1, modeA) + value(p, pointer+2, modeB)
			pointer += 4
		case 2: // Multiply
			p[p[pointer+3]] = value(p, pointer+1, modeA) * value(p, pointer+2, modeB)
			pointer += 4
		case 3: // Input
			p[p[pointer+1]] = <-input
			pointer += 2
		case 4: // Output
			output <- value(p, pointer+1, modeA)
			pointer += 2
		case 5: // Jump if true
			if value(p, pointer+1, modeA) != 0 {
				pointer = value(p, pointer+2, modeB)
			} else {
				pointer += 3
			}
		case 6: // Jump if false
			if value(p, pointer+1, modeA) == 0 {
				pointer = value(p, pointer+2, modeB)
			} else {
				pointer += 3
			}
		case 7: // Less than
			if value(p, pointer+1, modeA) < value(p, pointer+2, modeB) {
				p[p[pointer+3]] = 1
			} else {
				p[p[pointer+3]] = 0
			}
			pointer += 4
		case 8: // Equals
			if value(p, pointer+1, modeA) == value(p, pointer+2, modeB) {
				p[p[pointer+3]] = 1
			} else {
				p[p[pointer+3]] = 0
			}
			pointer += 4
		default:
			return fmt.Errorf("bad op code: %d", op)
		}
	}
}

func value(p Programme, index, mode int) int {
	switch mode {
	case 0:
		return p[p[index]]
	case 1:
		return p[index]
	default:
		panic("Unsupported op code")
	}
}

// Copy a Programme
func (p Programme) Copy() Programme {
	prog := make(Programme)
	for k, v := range p {
		prog[k] = v
	}
	return prog
}

func (p Programme) String() string {
	var values []string
	for k := 0; k < len(p); k++ {
		values = append(values, fmt.Sprintf("%d", p[k]))
	}
	return strings.Join(values, ",")
}

// Programme is an intcode programme
type Programme map[int]int

func parseOpcode(opcode int) (int, int, int, int) {
	var digits []int
	for i := 10000; i >= 1; i /= 10 {
		digits = append(digits, opcode/i)
		opcode -= i * (opcode / i)
	}
	return digits[3]*10 + digits[4], digits[2], digits[1], digits[0]
}
