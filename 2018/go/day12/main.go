package main

import (
	"io/ioutil"
	"log"
	"strings"
)

func main() {
	data, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatalln("Failed to read input", err)
	}

	lines := strings.Split(string(data), "\n")

	state := strings.TrimPrefix(lines[0], "initial state: ")

	rules := make(map[string]string)

	for _, line := range lines[2:] {
		rules[line[:5]] = string(line[9])
	}

	iterations := 50000000000

	var zero, i int
	for i = 0; i < iterations; i++ {
		// log.Println(i, state, zero)
		var newState string
		newState, zero = applyRules(state, zero, rules)
		if state == newState {
			break
		}
		state = newState
	}

	log.Println(i, state, zero)

	zero += (i - iterations + 1)

	log.Println(iterations, state, zero)

	var score int

	for i, r := range state {
		switch r {
		case '.':
		case '#':
			score += (i - zero)
		}
	}

	log.Println(score)
}

func applyRules(state string, zero int, rules map[string]string) (string, int) {
	for !strings.HasPrefix(state, ".....#") {
		if strings.HasPrefix(state, ".....") {
			state = state[1:]
			zero--
		} else {
			state = "." + state
			zero++
		}
	}
	for !strings.HasSuffix(state, "#.....") {
		if strings.HasSuffix(state, ".....") {
			state = state[:len(state)-2]
			zero--
		} else {
			state += "."
		}
	}
	zero -= 2

	var newState string

	for i := 0; i < len(state)-5; i++ {
		if match, ok := rules[state[i:i+5]]; ok {
			newState += match
		} else {
			newState += "."
		}
	}

	return newState, zero
}
