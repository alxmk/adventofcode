package main

import (
	"fmt"
	"strings"
)

func main() {
	reg := make(registers, 6)

	matches := make(map[int]struct{})

START:
	reg[4] = reg[1] | 65536
	reg[1] = 16298264

LABELD:
	reg[5] = (reg[4] & 255)
	reg[1] = reg[5] + reg[1]
	reg[1] = reg[1] & 16777215
	reg[1] = reg[1] * 65899
	reg[1] = reg[1] & 16777215
	if 256 > reg[4] {
		reg[5] = 1
		goto LABLEC
	} else {
		reg[5] = 0
	}
	reg[5] = 0
LABELA:
	reg[3] = reg[5] + 1
	reg[3] = reg[3] * 256
	if reg[3] > reg[4] {
		reg[3] = 1
		goto LABELB
	} else {
		reg[3] = 0
	}

	reg[5] = reg[5] + 1
	goto LABELA
LABELB:
	reg[4] = reg[5]
	goto LABELD

LABLEC:
	if _, ok := matches[reg[1]]; ok {
		fmt.Println(reg[1])
		return
	}
	fmt.Println(reg[1])
	// return
	matches[reg[1]] = struct{}{}
	if reg[1] == reg[0] {
		reg[5] = 1
		return
	} else {
		reg[5] = 0
		goto START
	}
}

type registers []int

func (r registers) String() string {
	var values []string

	for _, reg := range r {
		values = append(values, fmt.Sprintf("%d", reg))
	}

	return fmt.Sprintf("[%s]", strings.Join(values, ", "))
}
