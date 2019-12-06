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
		log.Fatalln("Failed to read input file", err)
	}

	commands := parseCommands(string(data))

	wires := make(circuit)
	wires.Run(commands)

	log.Println("Part one:", wires["a"])

	wires2 := make(circuit)
	wires2["b"] = wires["a"]

	wires2.Run(commands, override("b", wires["a"]))
	log.Println("Part two:", wires2["a"])
}

func override(match string, value uint16) func(string) (bool, uint16) {
	return func(wire string) (bool, uint16) {
		if wire == match {
			return true, value
		}
		return false, 0
	}
}

func parseCommands(input string) cmds {
	commands := make(cmds)
	for _, line := range strings.Split(input, "\n") {
		commands[parseCommand(line)] = struct{}{}
	}
	return commands
}

type circuit map[string]uint16

func (c circuit) String() string {
	var elements []string
	for wireName, value := range c {
		elements = append(elements, fmt.Sprintf("%s: %d", wireName, value))
	}
	return strings.Join(elements, "\n")
}

type cmds map[command]struct{}

func (c circuit) Run(commands cmds, overrides ...func(string) (bool, uint16)) {
	commandsCopy := make(cmds)
	for k, v := range commands {
		commandsCopy[k] = v
	}
	for {
		var ran []command
		for cmd := range commandsCopy {
			for _, o := range overrides {
				if ok, value := o(cmd.output); ok {
					ran = append(ran, cmd)
					c[cmd.output] = value
					continue
				}
			}
			if cmd.Run(c) {
				ran = append(ran, cmd)
			}
		}

		for _, r := range ran {
			delete(commandsCopy, r)
		}

		if len(commandsCopy) == 0 || len(ran) == 0 {
			break
		}
	}
}

type command struct {
	operandA, operandB string
	operator           string
	output             string
}

func (c command) String() string {
	return fmt.Sprintf("%s %s %s -> %s", c.operandA, c.operator, c.operandB, c.output)
}

func (c command) Run(wires circuit) bool {
	v, err := strconv.Atoi(c.operandA)
	valueA := uint16(v)
	if err != nil {
		// Not a number so see if the wire has a signal
		var ok bool
		if valueA, ok = wires[c.operandA]; !ok {
			return false
		}
	}
	switch c.operator {
	case "": // Assignment operator
		wires[c.output] = valueA
		return true
	case "NOT":
		wires[c.output] = ^valueA
		return true
	default:
		// We have two operands now
		v, err := strconv.Atoi(c.operandB)
		valueB := uint16(v)
		if err != nil {
			// Not a number so see if the wire has a signal
			var ok bool
			if valueB, ok = wires[c.operandB]; !ok {
				return false
			}
		}
		switch c.operator {
		case "AND":
			wires[c.output] = valueA & valueB
		case "OR":
			wires[c.output] = valueA | valueB
		case "LSHIFT":
			wires[c.output] = valueA << valueB
		case "RSHIFT":
			wires[c.output] = valueA >> valueB
		}
		return true
	}
}

func parseCommand(line string) command {
	var cmd command
	switch {
	case strings.Contains(line, "AND"), strings.Contains(line, "OR"), strings.Contains(line, "LSHIFT"), strings.Contains(line, "RSHIFT"):
		fmt.Sscanf(line, "%s %s %s -> %s", &cmd.operandA, &cmd.operator, &cmd.operandB, &cmd.output)
		return cmd
	case strings.Contains(line, "NOT"):
		fmt.Sscanf(line, "%s %s -> %s", &cmd.operator, &cmd.operandA, &cmd.output)
		return cmd
	default:
		fmt.Sscanf(line, "%s -> %s", &cmd.operandA, &cmd.output)
		return cmd
	}
}
