package main

import (
	"io/ioutil"
	"log"
	"strconv"
	"strings"
	"time"
)

var (
	registers = map[int64]*Registers{
		0: &Registers{
			registers: map[string]int64{
				"p": 0,
			},
			receive: make(chan Message, 128),
			sendto:  1,
			me:      0,
		},
		1: &Registers{
			registers: map[string]int64{
				"p": 1,
			},
			receive: make(chan Message, 128),
			sendto:  0,
			me:      1,
		},
	}
	done      = make(chan struct{})
	doneCount int
)

func main() {
	input, err := ioutil.ReadFile("../input.txt")
	if err != nil {
		log.Fatalf("Error reading input file %v", err)
	}

	instructions := strings.Split(string(input), "\n")

	for pid, reg := range registers {
		log.Println("Starting register", pid)
		go func(reg *Registers, pid int64) {
			reg.Run(instructions, pid)
		}(reg, pid)
	}

outer:
	for {
		select {
		case <-done:
			log.Println("Done")
			doneCount++
			if doneCount == 2 {
				break outer
			}
		}
	}

	log.Println("Part two answer is", registers[1].sendCount)
}

func (r *Registers) Run(instructions []string, pid int64) {
	var instructionNum int

instructionLoop:
	for {
		if instructionNum < 0 || instructionNum >= len(instructions) {
			log.Println("Breaking, instruction number is", instructionNum)
			r.snd("close")
			break
		}

		instruction := instructions[instructionNum]

		switch {
		case strings.HasPrefix(instruction, "snd"):
			r.snd(instruction)
		case strings.HasPrefix(instruction, "set"):
			log.Println(r.me, "Calling set with instruction", instruction, "instruction num", instructionNum)
			r.set(instruction)
		case strings.HasPrefix(instruction, "add"):
			r.add(instruction)
		case strings.HasPrefix(instruction, "mul"):
			r.mul(instruction)
		case strings.HasPrefix(instruction, "mod"):
			r.mod(instruction)
		case strings.HasPrefix(instruction, "rcv"):
			if r.rcv(instruction) {
				// We're done
				break instructionLoop
			}
		case strings.HasPrefix(instruction, "jgz"):
			if jump := r.jgz(instruction); jump != 0 {
				instructionNum += int(jump)
				continue instructionLoop
			}
		default:
			log.Fatalln("Unexpected instruction", instruction)
		}

		instructionNum++
	}

	done <- struct{}{}

	log.Println("Register done", pid)
}

func (r *Registers) snd(instruction string) {
	value := strings.TrimPrefix(instruction, "snd ")

	var toSend Message

	if value == "close" {
		registers[r.sendto].receive <- Message{
			close: true,
		}
		return
	}

	// If it's a number then set the frequency appropriately, otherwise look up from the register named
	if val, err := strconv.Atoi(value); err == nil {
		toSend = Message{
			value: int64(val),
		}
	} else {
		toSend = Message{
			value: r.GetValue(value),
		}
	}

	r.sendCount = r.sendCount + 1
	log.Println(r.me, "Sending++", r.sendCount)

	registers[r.sendto].receive <- toSend
}

func (r *Registers) rcv(instruction string) bool {
	value := strings.TrimPrefix(instruction, "rcv ")

	timeout := make(chan struct{})
	go func() {
		time.Sleep(1 * time.Second)
		timeout <- struct{}{}
	}()

	select {
	case receivedValue := <-r.receive:

		if receivedValue.close {
			return true
		}

		r.SetValue(value, receivedValue.value)

		return false
	case <-timeout:
		log.Println(r.me, "timed out receiving")
		return true
	}

}

func (r *Registers) set(instruction string) {
	values := strings.Fields(strings.TrimPrefix(instruction, "set "))

	// If it's a number then set the value appropriately
	if value, err := strconv.Atoi(values[1]); err == nil {
		r.SetValue(values[0], int64(value))
	} else {
		r.SetValue(values[0], r.GetValue(values[1]))
	}
}

func (r *Registers) add(instruction string) {
	values := strings.Fields(strings.TrimPrefix(instruction, "add "))

	currentVal := r.GetValue(values[0])

	// If it's a number then set the value appropriately
	if value, err := strconv.Atoi(values[1]); err == nil {
		r.SetValue(values[0], currentVal+int64(value))
	} else {
		r.SetValue(values[0], currentVal+r.GetValue(values[1]))
	}
}

func (r *Registers) mul(instruction string) {
	values := strings.Fields(strings.TrimPrefix(instruction, "mul "))

	currentVal := r.GetValue(values[0])

	// If it's a number then set the value appropriately
	if value, err := strconv.Atoi(values[1]); err == nil {
		r.SetValue(values[0], currentVal*int64(value))
	} else {
		r.SetValue(values[0], currentVal*r.GetValue(values[1]))
	}
}

func (r *Registers) mod(instruction string) {
	values := strings.Fields(strings.TrimPrefix(instruction, "mod "))

	currentVal := r.GetValue(values[0])

	// If it's a number then set the value appropriately
	if value, err := strconv.Atoi(values[1]); err == nil {
		r.SetValue(values[0], currentVal%int64(value))
	} else {
		r.SetValue(values[0], currentVal%r.GetValue(values[1]))
	}
}

func (r *Registers) jgz(instruction string) int64 {
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
	receive   chan Message
	sendto    int64
	waiting   bool
	sendCount int64
	me        int64
}

type Message struct {
	value int64
	close bool
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
