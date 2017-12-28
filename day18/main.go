package main

import (
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

var (
	r = Registers{
		registers: make(map[string]int64),
	}
)

func main() {
	input, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatalf("Error reading input file %v", err)
	}

	instructions := strings.Split(string(input), "\n")

	var instructionNum int

instructionLoop:
	for {
		if instructionNum < 0 || instructionNum >= len(instructions) {
			log.Println("Breaking, instruction number is", instructionNum)
			break
		}

		instruction := instructions[instructionNum]

		switch {
		case strings.HasPrefix(instruction, "snd"):
			sound(instruction)
		case strings.HasPrefix(instruction, "set"):
			set(instruction)
		case strings.HasPrefix(instruction, "add"):
			add(instruction)
		case strings.HasPrefix(instruction, "mul"):
			mul(instruction)
		case strings.HasPrefix(instruction, "mod"):
			mod(instruction)
		case strings.HasPrefix(instruction, "rcv"):
			if freq := rcv(instruction); freq > 0 {
				log.Println("Part one answer is", freq)
				break instructionLoop
			}
		case strings.HasPrefix(instruction, "jgz"):
			if jump := jgz(instruction); jump != 0 {
				instructionNum += int(jump)
				continue instructionLoop
			}
		default:
			log.Fatalln("Unexpected instruction", instruction)
		}

		instructionNum++
	}
}

func sound(instruction string) {
	value := strings.TrimPrefix(instruction, "snd ")

	// If it's a number then set the frequency appropriately, otherwise look up from the register named
	if frequency, err := strconv.Atoi(value); err == nil {
		r.SetLastSound(int64(frequency))
	} else {
		r.SetLastSound(r.GetValue(value))
	}
}

func set(instruction string) {
	values := strings.Fields(strings.TrimPrefix(instruction, "set "))

	// If it's a number then set the value appropriately
	if value, err := strconv.Atoi(values[1]); err == nil {
		r.SetValue(values[0], int64(value))
	} else {
		r.SetValue(values[0], r.GetValue(values[1]))
	}
}

func add(instruction string) {
	values := strings.Fields(strings.TrimPrefix(instruction, "add "))

	currentVal := r.GetValue(values[0])

	// If it's a number then set the value appropriately
	if value, err := strconv.Atoi(values[1]); err == nil {
		r.SetValue(values[0], currentVal+int64(value))
	} else {
		r.SetValue(values[0], currentVal+r.GetValue(values[1]))
	}
}

func mul(instruction string) {
	values := strings.Fields(strings.TrimPrefix(instruction, "mul "))

	currentVal := r.GetValue(values[0])

	// If it's a number then set the value appropriately
	if value, err := strconv.Atoi(values[1]); err == nil {
		r.SetValue(values[0], currentVal*int64(value))
	} else {
		r.SetValue(values[0], currentVal*r.GetValue(values[1]))
	}
}

func mod(instruction string) {
	values := strings.Fields(strings.TrimPrefix(instruction, "mod "))

	currentVal := r.GetValue(values[0])

	// If it's a number then set the value appropriately
	if value, err := strconv.Atoi(values[1]); err == nil {
		r.SetValue(values[0], currentVal%int64(value))
	} else {
		r.SetValue(values[0], currentVal%r.GetValue(values[1]))
	}
}

func rcv(instruction string) int64 {
	value := strings.TrimPrefix(instruction, "rcv ")

	var checkVal int64

	// If it's a number then set the value appropriately
	if val, err := strconv.Atoi(value); err == nil {
		checkVal = int64(val)
	} else {
		checkVal = r.GetValue(value)
	}

	if checkVal != 0 {
		return r.GetLastSound()
	}

	return checkVal
}

func jgz(instruction string) int64 {
	values := strings.Fields(strings.TrimPrefix(instruction, "jgz "))

	var checkVal int64

	// If it's a number then set the value appropriately
	if val, err := strconv.Atoi(values[0]); err == nil {
		checkVal = int64(val)
	} else {
		checkVal = r.GetValue(values[0])
	}

	if checkVal <= 0 {
		return 0
	}

	var jumpVal int64

	// If it's a number then set the value appropriately
	if val, err := strconv.Atoi(values[1]); err == nil {
		jumpVal = int64(val)
	} else {
		jumpVal = r.GetValue(values[1])
	}

	return jumpVal
}

type Registers struct {
	registers map[string]int64
	lastSound int64
}

func (r *Registers) GetValue(register string) int64 {
	return r.registers[register]
}

func (r *Registers) SetValue(register string, value int64) {
	r.registers[register] = value
}

func (r *Registers) GetLastSound() int64 {
	return r.lastSound
}

func (r *Registers) SetLastSound(frequency int64) {
	r.lastSound = frequency
}
