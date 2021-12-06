package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

func main() {
	data, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatalln("Failed to read input:", err)
	}

	p := parse(string(data))
	p1, _ := p.Run()
	log.Println("Part one:", p1)

	for _, i := range p.instructions {
		switch i.op {
		case "acc":
			continue
		case "jmp":
			i.op = "nop"
			if at, ok := p.Run(); ok {
				log.Println("Part two:", at)
				return
			}
			i.op = "jmp"
		case "nop":
			i.op = "jmp"
			if at, ok := p.Run(); ok {
				log.Println("Part two:", at)
				return
			}
			i.op = "nop"
		}
	}
}

func parse(input string) programme {
	var p programme
	for _, line := range strings.Split(input, "\n") {
		var i instruction
		fmt.Sscanf(line, "%s %d", &i.op, &i.arg)
		p.instructions = append(p.instructions, &i)
	}
	return p
}

type programme struct {
	instructions []*instruction
	accumulator  int
}

type instruction struct {
	op  string
	arg int
}

// Run runs the programme until detecting an infinite loop, or exiting
// either cleanly when the pointer is one past the end of the programme
// or with an error when the pointer is beyond the end of the programme
func (p programme) Run() (int, bool) {
	p.accumulator = 0
	ran := make(map[int]struct{})
	var pointer int
	for pointer < len(p.instructions) {
		if _, ok := ran[pointer]; ok {
			return p.accumulator, false
		}

		ran[pointer] = struct{}{}

		i := p.instructions[pointer]
		switch i.op {
		case "acc":
			p.accumulator += i.arg
		case "jmp":
			pointer += i.arg
			continue
		}
		pointer++
	}
	return p.accumulator, pointer == len(p.instructions)
}
