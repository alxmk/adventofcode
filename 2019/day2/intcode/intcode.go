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
		p[int64(i)] = int64(parsed)
	}
	return p, nil
}

// Run a Programme
func (p Programme) Run(input, output chan int64) error {
	var pointer, relativeBase int64
	for {
		op, modeA, modeB, modeC := parseOpcode(p[pointer])
		// log.Println("op", op)
		switch op {
		case 99: // Exit
			close(output)
			return nil
		case 1: // Add
			p[targetIdx(p, pointer+3, modeC, relativeBase)] = value(p, pointer+1, modeA, relativeBase) + value(p, pointer+2, modeB, relativeBase)
			pointer += 4
		case 2: // Multiply
			p[targetIdx(p, pointer+3, modeC, relativeBase)] = value(p, pointer+1, modeA, relativeBase) * value(p, pointer+2, modeB, relativeBase)
			pointer += 4
		case 3: // Input
			p[targetIdx(p, pointer+1, modeA, relativeBase)] = <-input
			pointer += 2
		case 4: // Output
			output <- value(p, pointer+1, modeA, relativeBase)
			pointer += 2
		case 5: // Jump if true
			if value(p, pointer+1, modeA, relativeBase) != 0 {
				pointer = value(p, pointer+2, modeB, relativeBase)
			} else {
				pointer += 3
			}
		case 6: // Jump if false
			if value(p, pointer+1, modeA, relativeBase) == 0 {
				pointer = value(p, pointer+2, modeB, relativeBase)
			} else {
				pointer += 3
			}
		case 7: // Less than
			if value(p, pointer+1, modeA, relativeBase) < value(p, pointer+2, modeB, relativeBase) {
				p[targetIdx(p, pointer+3, modeC, relativeBase)] = 1
			} else {
				p[targetIdx(p, pointer+3, modeC, relativeBase)] = 0
			}
			pointer += 4
		case 8: // Equals
			if value(p, pointer+1, modeA, relativeBase) == value(p, pointer+2, modeB, relativeBase) {
				p[targetIdx(p, pointer+3, modeC, relativeBase)] = 1
			} else {
				p[targetIdx(p, pointer+3, modeC, relativeBase)] = 0
			}
			pointer += 4
		case 9: // Relative base modify
			relativeBase += value(p, pointer+1, modeA, relativeBase)
			pointer += 2
		default:
			return fmt.Errorf("bad op code: %d", op)
		}
	}
}

func targetIdx(p Programme, index, mode, relativeBase int64) int64 {
	switch mode {
	case 0:
		return p[index]
	case 1:
		panic("immediate mode unsupported for storing")
	case 2:
		return p[index] + relativeBase
	default:
		panic("unsupported mode")
	}
}

func value(p Programme, index, mode, relativeBase int64) int64 {
	switch mode {
	case 0, 2:
		return p[targetIdx(p, index, mode, relativeBase)]
	case 1:
		return p[index]
	default:
		panic("unsupported mode")
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
	for k := int64(0); k < int64(len(p)); k++ {
		values = append(values, fmt.Sprintf("%d", p[k]))
	}
	return strings.Join(values, ",")
}

// Programme is an intcode programme
type Programme map[int64]int64

func parseOpcode(opcode int64) (int64, int64, int64, int64) {
	var digits []int64
	for i := int64(10000); i >= 1; i /= 10 {
		digits = append(digits, opcode/i)
		opcode -= i * (opcode / i)
	}
	return digits[3]*10 + digits[4], digits[2], digits[1], digits[0]
}
